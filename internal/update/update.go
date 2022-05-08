package update

import (
	"context"
	"database/sql"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/jamespfennell/gtfs"
	"github.com/jamespfennell/transiter/config"
	"github.com/jamespfennell/transiter/internal/gen/db"
	"github.com/jamespfennell/transiter/internal/update/static"
)

func CreateAndRun(ctx context.Context, database *sql.DB, systemId, feedId string) error {
	runInTx := func(f func(querier db.Querier) error) error {
		tx, err := database.BeginTx(ctx, nil)
		if err != nil {
			return err
		}
		defer tx.Rollback()
		err = f(db.New(tx))
		if err != nil {
			return err
		}
		tx.Commit()
		return nil
	}
	var feedPk *int64
	if err := runInTx(func(querier db.Querier) error {
		var err error
		*feedPk, err = CreateInsideTx(ctx, querier, systemId, feedId)
		return err
	}); err != nil {
		return err
	}
	// TODO: mark update as IN_PROGRESS
	return runInTx(func(querier db.Querier) error {
		return RunInsideTx(ctx, querier, *feedPk)
	})
}

func CreateAndRunInsideTx(ctx context.Context, querier db.Querier, systemId, feedId string) error {
	updatePk, err := CreateInsideTx(ctx, querier, systemId, feedId)
	if err != nil {
		return err
	}
	return RunInsideTx(ctx, querier, updatePk)
}

func CreateInsideTx(ctx context.Context, querier db.Querier, systemId, feedId string) (int64, error) {
	log.Printf("Creating update for %s/%s\n", systemId, feedId)
	feed, err := querier.GetFeedInSystem(ctx, db.GetFeedInSystemParams{
		SystemID: systemId,
		FeedID:   feedId,
	})
	if err != nil {
		return 0, err
	}
	return querier.InsertFeedUpdate(ctx, db.InsertFeedUpdateParams{
		FeedPk: feed.Pk,
		Status: "CREATED",
	})
}

func RunInsideTx(ctx context.Context, querier db.Querier, updatePk int64) error {
	feed, err := querier.GetFeedForUpdate(ctx, updatePk)
	if err != nil {
		log.Printf("Error update for pk=%d\n", updatePk)
		return err
	}
	feedConfig, err := config.UnmarshalFromJson([]byte(feed.Config))
	if err != nil {
		return fmt.Errorf("failed to parse feed config in the DB: %w", err)
	}
	content, err := getFeedContent(ctx, feedConfig)
	if err != nil {
		return err
	}
	switch feedConfig.Parser {
	case config.GtfsStatic:
		// TODO: support custom GTFS static options
		parsedEntities, err := gtfs.ParseStatic(content, gtfs.ParseStaticOptions{})
		if err != nil {
			return err
		}
		return static.Update(ctx, static.UpdateContext{
			Querier:  querier,
			SystemPk: feed.SystemPk,
			FeedPk:   feed.Pk,
			UpdatePk: updatePk,
		}, parsedEntities)
	default:
		return fmt.Errorf("unknown parser %q", feedConfig.Parser)
	}
}

func getFeedContent(ctx context.Context, feedConfig *config.FeedConfig) ([]byte, error) {
	client := http.Client{
		Timeout: 5 * time.Second,
	}
	if feedConfig.HttpTimeout != nil {
		client.Timeout = *feedConfig.HttpTimeout
	}
	req, err := http.NewRequestWithContext(ctx, "GET", feedConfig.Url, nil)
	if err != nil {
		return nil, err
	}
	for key, value := range feedConfig.HttpHeaders {
		req.Header.Add(key, value)
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return io.ReadAll(resp.Body)
}
