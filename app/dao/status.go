package dao

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"yatter-backend-go/app/domain/object"
	"yatter-backend-go/app/domain/repository"

	"github.com/jmoiron/sqlx"
)

type (
	// Implementation for repository.Status
	status struct {
		db *sqlx.DB
	}
)

var _ repository.Status = (*status)(nil)

// Create status repository
func NewStatus(db *sqlx.DB) *status {
	return &status{db: db}
}

// FindById : statusIdからstatusを取得
func (s *status) FindById(ctx context.Context, id int) (*object.Status, error) {
	entity := new(object.Status)
	err := s.db.QueryRowxContext(ctx, "select * from status where id = ?", id).StructScan(entity)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}

		return nil, fmt.Errorf("failed to find status from db: %w", err)
	}

	return entity, nil
}

func (s *status) Create(ctx context.Context, tx *sqlx.Tx, sta *object.Status) error {
	_, err := s.db.Exec("insert into status (account_id, url, content, created_at) values (?, ?, ?, ?)",
		sta.AccountID, sta.URL, sta.Content, sta.CreatedAt)
	if err != nil {
		return fmt.Errorf("failed to insert status: %w", err)
	}

	return nil
}
