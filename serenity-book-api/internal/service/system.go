package service

import (
	"context"
	"serenity-book-api/internal/model"
	"serenity-book-api/internal/repository"
)

type SystemService interface {
	GetSystem(ctx context.Context, id int64) (*model.System, error)
}

func NewSystemService(service *Service, systemRepository repository.SystemRepository) SystemService {
	return &systemService{
		Service:          service,
		systemRepository: systemRepository,
	}
}

type systemService struct {
	*Service
	systemRepository repository.SystemRepository
}

func (s *systemService) GetSystem(ctx context.Context, id int64) (*model.System, error) {
	return s.systemRepository.GetSystem(ctx, id)
}
