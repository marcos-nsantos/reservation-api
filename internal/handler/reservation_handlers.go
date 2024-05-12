package handler

import (
	"errors"
	"net/http"
	"strconv"

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

func (h *ReservationHandler) GetReservation(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)

	reservation, err := h.ReservationService.GetReservationByID(id)
	if err != nil {
		if errors.Is(err, entity.ErrReservationNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, reservation)
}

func (h *ReservationHandler) GetUserReservations(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	perPage, _ := strconv.Atoi(c.DefaultQuery("perPage", "10"))

	userID := c.GetUint64("userID")

	reservations, total, err := h.ReservationService.GetUserReservations(userID, page, perPage)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, ReservationResponseWithPagination{
		Reservations: reservations,
		Total:        total,
		Page:         page,
		PerPage:      perPage,
	})
}
