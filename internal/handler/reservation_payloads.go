package handler

import (
	"time"

	"github.com/marcos-nsantos/reservation-api/internal/entity"
)

type CreateReservationRequest struct {
	ResourceID uint64    `json:"resourceID" binding:"required"`
	StartTime  time.Time `json:"startTime" binding:"required"`
	EndTime    time.Time `json:"endTime" binding:"required"`
}

type ReservationResponseWithPagination struct {
	Reservations []entity.Reservation `json:"reservations"`
	Total        int64                `json:"total"`
	Page         int                  `json:"page"`
	PerPage      int                  `json:"PerPage"`
}
