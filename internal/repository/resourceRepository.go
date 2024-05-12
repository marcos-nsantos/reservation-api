package repository

import (
	"gorm.io/gorm"

	"github.com/marcos-nsantos/reservation-api/internal/entity"
)

type ResourceRepository struct {
	db *gorm.DB
}

func NewResourceRepository(db *gorm.DB) *ResourceRepository {
	return &ResourceRepository{db: db}
}

func (r *ResourceRepository) Create(resource *entity.Resource) error {
	return r.db.Create(resource).Error
}
