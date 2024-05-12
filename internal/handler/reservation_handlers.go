package handler

import (
	"errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/marcos-nsantos/reservation-api/internal/entity"
	"github.com/marcos-nsantos/reservation-api/internal/service"
)

type ReservationHandler struct {
	ReservationService *service.ReservationService
}

func NewReservationHandler(reservationService *service.ReservationService) *ReservationHandler {
	return &ReservationHandler{
		ReservationService: reservationService,
	}
}

type CreateReservationRequest struct {
	ResourceID uint64    `json:"resourceID" binding:"required"`
	StartTime  time.Time `json:"startTime" binding:"required"`
	EndTime    time.Time `json:"endTime" binding:"required"`
}

func (h *ReservationHandler) CreateReservation(c *gin.Context) {
	var createReservationRequest CreateReservationRequest

	if err := c.ShouldBindJSON(&createReservationRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID := c.GetUint64("userID")
	reservation := &entity.Reservation{
		ResourceID: createReservationRequest.ResourceID,
		UserID:     userID,
		StartTime:  createReservationRequest.StartTime,
		EndTime:    createReservationRequest.EndTime,
	}

	if err := h.ReservationService.CreateReservation(reservation); err != nil {
		if errors.Is(err, entity.ErrInvalidReservationTime) {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if errors.Is(err, entity.ErrResourceNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}

		if errors.Is(err, entity.ErrReservationNotAvailable) {
			c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, reservation)
}
