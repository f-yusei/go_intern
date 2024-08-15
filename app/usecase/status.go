package usecase

import (
	"context"
	"yatter-backend-go/app/domain/object"
	"yatter-backend-go/app/domain/repository"

	"github.com/jmoiron/sqlx"
)

type Status interface {
	Create(ctx context.Context, status string, accountId int) (*CreateStatusDTO, error)
	FindById(ctx context.Context, id int) (*GetStatusDTO, error)
}

type status struct {
	db         *sqlx.DB
	statusRepo repository.Status
}

type CreateStatusDTO struct {
	Status *object.Status
}

type GetStatusDTO struct {
	Status *object.Status
}

var _ Status = (*status)(nil)

func NewStatus(db *sqlx.DB, statusRepo repository.Status) *status {
	return &status{
		db:         db,
		statusRepo: statusRepo,
	}
}

func (s *status) Create(ctx context.Context, status string, accountId int) (*CreateStatusDTO, error) {
	sta := object.NewStatus(status, accountId)

	tx, err := s.db.Beginx()
	if err != nil {
		return nil, err
	}

	defer func() {
		if err := recover(); err != nil {
			tx.Rollback()
		}

		tx.Commit()
	}()

	if err := s.statusRepo.Create(ctx, tx, sta); err != nil {
		return nil, err
	}

	return &CreateStatusDTO{
		Status: sta,
	}, nil
}

func (s *status) FindById(ctx context.Context, id int) (*GetStatusDTO, error) {
	sta, err := s.statusRepo.FindById(ctx, id)
	if err != nil {
		return nil, err
	}

	return &GetStatusDTO{
		Status: sta,
	}, nil
}
