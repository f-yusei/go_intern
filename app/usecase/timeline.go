package usecase

import (
	"context"
	"yatter-backend-go/app/domain/object"
	"yatter-backend-go/app/domain/repository"

	"github.com/jmoiron/sqlx"
)

type Timeline interface {
	Get(ctx context.Context, limit int) (*GetTimelineDTO, error)
}

type timeline struct {
	db           *sqlx.DB
	timelineRepo repository.Timeline
}

type GetTimelineDTO struct {
	Timeline *object.Timeline
}

var _ Timeline = (*timeline)(nil)

func NewTimeline(db *sqlx.DB, timelineRepo repository.Timeline) *timeline {
	return &timeline{
		db:           db,
		timelineRepo: timelineRepo,
	}
}

func (t *timeline) Get(ctx context.Context, limit int) (*GetTimelineDTO, error) {
	tim, err := t.timelineRepo.Get(ctx, limit)
	if err != nil {
		return nil, err
	}

	return &GetTimelineDTO{
		Timeline: tim,
	}, nil
}
