package service

import (
	"time"

	"github.com/marcos-nsantos/reservation-api/internal/entity"
	"github.com/marcos-nsantos/reservation-api/internal/repository"
)

type ReservationService struct {
	ReservationRepository *repository.ReservationRepository
	ResourceRepository    *repository.ResourceRepository
}

func NewReservationService(reservationRepository *repository.ReservationRepository, resourceRepository *repository.ResourceRepository) *ReservationService {
	return &ReservationService{
		ReservationRepository: reservationRepository,
		ResourceRepository:    resourceRepository,
	}
}

func (s *ReservationService) CreateReservation(reservation *entity.Reservation) error {
	if reservation.StartTime.After(reservation.EndTime) || reservation.StartTime.Before(time.Now()) {
		return entity.ErrInvalidReservationTime
	}

	_, err := s.ResourceRepository.GetByID(reservation.ResourceID)
	if err != nil {
		return entity.ErrResourceNotFound
	}

	if !s.ReservationRepository.CheckAvailability(reservation) {
		return entity.ErrReservationNotAvailable
	}

	return s.ReservationRepository.Create(reservation)
}
