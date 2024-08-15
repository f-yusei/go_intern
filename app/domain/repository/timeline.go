package repository

import (
	"context"

	"yatter-backend-go/app/domain/object"
)

type Timeline interface {
	Get(ctx context.Context, limit int) (*object.Timeline, error)
}
