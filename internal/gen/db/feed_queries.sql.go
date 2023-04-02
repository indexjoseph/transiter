// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.2
// source: feed_queries.sql

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const countUpdatesInFeed = `-- name: CountUpdatesInFeed :one
SELECT COUNT(*) FROM feed_update WHERE feed_pk = $1
`

func (q *Queries) CountUpdatesInFeed(ctx context.Context, feedPk int64) (int64, error) {
	row := q.db.QueryRow(ctx, countUpdatesInFeed, feedPk)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const deleteFeed = `-- name: DeleteFeed :exec
DELETE FROM feed WHERE pk = $1
`

func (q *Queries) DeleteFeed(ctx context.Context, pk int64) error {
	_, err := q.db.Exec(ctx, deleteFeed, pk)
	return err
}

const finishFeedUpdate = `-- name: FinishFeedUpdate :exec
UPDATE feed_update
SET finished = true, 
    result = $1,
    finished_at = $2,
    content_length = $3,
    content_hash = $4,
    error_message = $5
WHERE pk = $6
`

type FinishFeedUpdateParams struct {
	Result        pgtype.Text
	FinishedAt    pgtype.Timestamptz
	ContentLength pgtype.Int4
	ContentHash   pgtype.Text
	ErrorMessage  pgtype.Text
	UpdatePk      int64
}

func (q *Queries) FinishFeedUpdate(ctx context.Context, arg FinishFeedUpdateParams) error {
	_, err := q.db.Exec(ctx, finishFeedUpdate,
		arg.Result,
		arg.FinishedAt,
		arg.ContentLength,
		arg.ContentHash,
		arg.ErrorMessage,
		arg.UpdatePk,
	)
	return err
}

const garbageCollectFeedUpdates = `-- name: GarbageCollectFeedUpdates :one
WITH deleted_pks AS (
    DELETE FROM feed_update
    WHERE started_at <= NOW() - INTERVAL '7 days'
    AND NOT feed_update.pk = ANY($1::bigint[])
    RETURNING pk
)
SELECT COUNT(*) FROM deleted_pks
`

func (q *Queries) GarbageCollectFeedUpdates(ctx context.Context, activeFeedUpdatePks []int64) (int64, error) {
	row := q.db.QueryRow(ctx, garbageCollectFeedUpdates, activeFeedUpdatePks)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const getFeed = `-- name: GetFeed :one
SELECT feed.pk, feed.id, feed.system_pk, feed.update_strategy, feed.update_period, feed.config FROM feed
    INNER JOIN system on system.pk = feed.system_pk
    WHERE system.id = $1
    AND feed.id = $2
`

type GetFeedParams struct {
	SystemID string
	FeedID   string
}

func (q *Queries) GetFeed(ctx context.Context, arg GetFeedParams) (Feed, error) {
	row := q.db.QueryRow(ctx, getFeed, arg.SystemID, arg.FeedID)
	var i Feed
	err := row.Scan(
		&i.Pk,
		&i.ID,
		&i.SystemPk,
		&i.UpdateStrategy,
		&i.UpdatePeriod,
		&i.Config,
	)
	return i, err
}

const getFeedForUpdate = `-- name: GetFeedForUpdate :one
SELECT feed.pk, feed.id, feed.system_pk, feed.update_strategy, feed.update_period, feed.config FROM feed
    INNER JOIN feed_update ON feed_update.feed_pk = feed.pk
    WHERE feed_update.pk = $1
`

func (q *Queries) GetFeedForUpdate(ctx context.Context, updatePk int64) (Feed, error) {
	row := q.db.QueryRow(ctx, getFeedForUpdate, updatePk)
	var i Feed
	err := row.Scan(
		&i.Pk,
		&i.ID,
		&i.SystemPk,
		&i.UpdateStrategy,
		&i.UpdatePeriod,
		&i.Config,
	)
	return i, err
}

const getFeedUpdate = `-- name: GetFeedUpdate :one
SELECT pk, feed_pk, started_at, finished, finished_at, result, content_length, content_hash, error_message FROM feed_update WHERE pk = $1
`

func (q *Queries) GetFeedUpdate(ctx context.Context, pk int64) (FeedUpdate, error) {
	row := q.db.QueryRow(ctx, getFeedUpdate, pk)
	var i FeedUpdate
	err := row.Scan(
		&i.Pk,
		&i.FeedPk,
		&i.StartedAt,
		&i.Finished,
		&i.FinishedAt,
		&i.Result,
		&i.ContentLength,
		&i.ContentHash,
		&i.ErrorMessage,
	)
	return i, err
}

const getLastFeedUpdateContentHash = `-- name: GetLastFeedUpdateContentHash :one
SELECT content_hash
FROM feed_update
WHERE feed_pk = $1 AND result = 'UPDATED'
ORDER BY finished_at DESC
LIMIT 1
`

func (q *Queries) GetLastFeedUpdateContentHash(ctx context.Context, feedPk int64) (pgtype.Text, error) {
	row := q.db.QueryRow(ctx, getLastFeedUpdateContentHash, feedPk)
	var content_hash pgtype.Text
	err := row.Scan(&content_hash)
	return content_hash, err
}

const insertFeed = `-- name: InsertFeed :exec
INSERT INTO feed
    (id, system_pk, update_strategy, update_period, config)
VALUES
    ($1, $2, $3, 
     $4, $5)
`

type InsertFeedParams struct {
	ID             string
	SystemPk       int64
	UpdateStrategy string
	UpdatePeriod   pgtype.Float8
	Config         string
}

func (q *Queries) InsertFeed(ctx context.Context, arg InsertFeedParams) error {
	_, err := q.db.Exec(ctx, insertFeed,
		arg.ID,
		arg.SystemPk,
		arg.UpdateStrategy,
		arg.UpdatePeriod,
		arg.Config,
	)
	return err
}

const insertFeedUpdate = `-- name: InsertFeedUpdate :one
INSERT INTO feed_update
    (feed_pk, started_at, finished)
VALUES
    ($1, $2, false)
RETURNING pk
`

type InsertFeedUpdateParams struct {
	FeedPk    int64
	StartedAt pgtype.Timestamptz
}

func (q *Queries) InsertFeedUpdate(ctx context.Context, arg InsertFeedUpdateParams) (int64, error) {
	row := q.db.QueryRow(ctx, insertFeedUpdate, arg.FeedPk, arg.StartedAt)
	var pk int64
	err := row.Scan(&pk)
	return pk, err
}

const listActiveFeedUpdatePks = `-- name: ListActiveFeedUpdatePks :many
      SELECT DISTINCT source_pk FROM agency
UNION SELECT DISTINCT source_pk FROM alert
UNION SELECT DISTINCT source_pk FROM scheduled_service
UNION SELECT DISTINCT source_pk FROM route
UNION SELECT DISTINCT source_pk FROM stop
UNION SELECT DISTINCT source_pk FROM stop_headsign_rule
UNION SELECT DISTINCT source_pk FROM transfer
UNION SELECT DISTINCT source_pk FROM trip
UNION SELECT DISTINCT source_pk FROM vehicle
`

func (q *Queries) ListActiveFeedUpdatePks(ctx context.Context) ([]int64, error) {
	rows, err := q.db.Query(ctx, listActiveFeedUpdatePks)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []int64
	for rows.Next() {
		var source_pk int64
		if err := rows.Scan(&source_pk); err != nil {
			return nil, err
		}
		items = append(items, source_pk)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listFeeds = `-- name: ListFeeds :many
SELECT pk, id, system_pk, update_strategy, update_period, config FROM feed WHERE system_pk = $1 ORDER BY id
`

func (q *Queries) ListFeeds(ctx context.Context, systemPk int64) ([]Feed, error) {
	rows, err := q.db.Query(ctx, listFeeds, systemPk)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Feed
	for rows.Next() {
		var i Feed
		if err := rows.Scan(
			&i.Pk,
			&i.ID,
			&i.SystemPk,
			&i.UpdateStrategy,
			&i.UpdatePeriod,
			&i.Config,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listUpdatesInFeed = `-- name: ListUpdatesInFeed :many
SELECT pk, feed_pk, started_at, finished, finished_at, result, content_length, content_hash, error_message FROM feed_update
WHERE feed_pk = $1
ORDER BY pk DESC
LIMIT 100
`

func (q *Queries) ListUpdatesInFeed(ctx context.Context, feedPk int64) ([]FeedUpdate, error) {
	rows, err := q.db.Query(ctx, listUpdatesInFeed, feedPk)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []FeedUpdate
	for rows.Next() {
		var i FeedUpdate
		if err := rows.Scan(
			&i.Pk,
			&i.FeedPk,
			&i.StartedAt,
			&i.Finished,
			&i.FinishedAt,
			&i.Result,
			&i.ContentLength,
			&i.ContentHash,
			&i.ErrorMessage,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateFeed = `-- name: UpdateFeed :exec
UPDATE feed
SET update_strategy = $1,
    update_period = $2, 
    config = $3
WHERE pk = $4
`

type UpdateFeedParams struct {
	UpdateStrategy string
	UpdatePeriod   pgtype.Float8
	Config         string
	FeedPk         int64
}

func (q *Queries) UpdateFeed(ctx context.Context, arg UpdateFeedParams) error {
	_, err := q.db.Exec(ctx, updateFeed,
		arg.UpdateStrategy,
		arg.UpdatePeriod,
		arg.Config,
		arg.FeedPk,
	)
	return err
}
