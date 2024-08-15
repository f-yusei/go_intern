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
	timeline struct {
		db *sqlx.DB
	}
)

var _ repository.Timeline = (*timeline)(nil)

// Create status repository
func NewTimeline(db *sqlx.DB) *timeline {
	return &timeline{db: db}
}

// Get : accountとstatusを取得してtimelineとしてまとめる
func (t *timeline) Get(ctx context.Context, limit int) (*object.Timeline, error) {
	entity := new(object.Timeline)
	query := `
		SELECT 
			s.id AS status_id, 
			s.account_id, 
			a.username, 
			a.display_name, 
			a.avatar, 
			s.content, 
			s.url, 
			s.created_at
		FROM 
			status s
		INNER JOIN 
			account a ON s.account_id = a.id
		ORDER BY 
			s.created_at DESC
		LIMIT ?`

	err := t.db.QueryRowxContext(ctx, query, limit).StructScan(entity)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}

		return nil, fmt.Errorf("failed to find status from db: %w", err)
	}

	return entity, nil
}
