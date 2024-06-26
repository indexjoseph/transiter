package admin

import (
	"context"
	"time"

	"github.com/jamespfennell/transiter/internal/gen/api"
)

func (s *Service) GetSchedulerStatus(ctx context.Context, req *api.GetSchedulerStatusRequest) (*api.GetSchedulerStatusReply, error) {
	reply := &api.GetSchedulerStatusReply{}
	feeds, err := s.scheduler.Status(ctx)
	if err != nil {
		return nil, err
	}
	for _, feed := range feeds {
		reply.Feeds = append(reply.Feeds, &api.GetSchedulerStatusReply_FeedStatus{
			SystemId:             feed.SystemID,
			FeedConfig:           feed.FeedConfig,
			LastSuccessfulUpdate: convertTime(feed.LastSuccessfulUpdate),
			LastFinishedUpdate:   convertTime(feed.LastFinishedUpdate),
			CurrentlyRunning:     feed.CurrentlyRunning,
		})
	}
	return reply, nil
}

func (s *Service) ResetScheduler(ctx context.Context, req *api.ResetSchedulerRequest) (*api.ResetSchedulerReply, error) {
	if err := s.scheduler.Reset(ctx); err != nil {
		return nil, err
	}
	return &api.ResetSchedulerReply{}, nil
}

func convertTime(t time.Time) int64 {
	if t.IsZero() {
		return 0
	}
	return t.Unix()
}
