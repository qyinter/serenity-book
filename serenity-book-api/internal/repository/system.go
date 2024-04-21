package repository

import (
	"context"
	"serenity-book-api/internal/model"
)

type SystemRepository interface {
	GetSystem(ctx context.Context, id int64) (*model.System, error)
}

func NewSystemRepository(
	repository *Repository,
) SystemRepository {
	return &systemRepository{
		Repository: repository,
	}
}

type systemRepository struct {
	*Repository
}

func (r *systemRepository) GetSystem(ctx context.Context, id int64) (*model.System, error) {
	var system model.System

	return &system, nil
}
