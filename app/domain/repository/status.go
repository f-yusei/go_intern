package repository

import (
	"context"

	"yatter-backend-go/app/domain/object"

	"github.com/jmoiron/sqlx"
)

type Status interface {
	// Fetch status which has specified id
	FindById(ctx context.Context, id int) (*object.Status, error)
	Create(ctx context.Context, tx *sqlx.Tx, sta *object.Status) error
}
