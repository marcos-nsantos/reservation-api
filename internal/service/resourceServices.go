package service

import (
	"github.com/marcos-nsantos/reservation-api/internal/entity"
	"github.com/marcos-nsantos/reservation-api/internal/repository"
)

type ResourceService struct {
	Repo *repository.ResourceRepository
}

func NewResourceService(repo *repository.ResourceRepository) *ResourceService {
	return &ResourceService{Repo: repo}
}

func (s *ResourceService) CreateResource(resource *entity.Resource) error {
	return s.Repo.Create(resource)
}
