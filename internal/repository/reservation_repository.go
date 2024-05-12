package repository

import (
	"errors"

	"gorm.io/gorm"

	"github.com/marcos-nsantos/reservation-api/internal/entity"
)

type ReservationRepository struct {
	db *gorm.DB
}

func NewReservationRepository(db *gorm.DB) *ReservationRepository {
	return &ReservationRepository{db: db}
}

func (r *ReservationRepository) Create(reservation *entity.Reservation) error {
	return r.db.Create(reservation).Error
}

func (r *ReservationRepository) CheckAvailability(reservation *entity.Reservation) bool {
	var count int64
	r.db.Model(&entity.Reservation{}).
		Where("resource_id = ?", reservation.ResourceID).
		Where("? < end_time AND ? > start_time", reservation.EndTime, reservation.StartTime).
		Where("status <> ?", entity.Cancelled).
		Where("status <> ?", entity.Rejected).
		Count(&count)

	return count == 0
}

func (r *ReservationRepository) GetByID(id uint64) (*entity.Reservation, error) {
	var reservation entity.Reservation
	err := r.db.First(&reservation, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, entity.ErrReservationNotFound
	}

	return &reservation, err
}

func (r *ReservationRepository) GetUserReservations(userID uint64, page, perPage int) ([]entity.Reservation, int64, error) {
	var reservations []entity.Reservation
	var total int64

	r.db.Model(&entity.Reservation{}).
		Where("user_id = ?", userID).
		Count(&total).
		Offset((page - 1) * perPage).
		Limit(perPage).
		Find(&reservations)

	return reservations, total, nil
}
