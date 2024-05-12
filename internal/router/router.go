package router

import (
	"github.com/gin-gonic/gin"

	"github.com/marcos-nsantos/reservation-api/internal/handler"
	"github.com/marcos-nsantos/reservation-api/internal/router/middleware"
)

type Router struct {
	Key                string
	UserHandler        *handler.UserHandler
	ResourceHandler    *handler.ResourceHandler
	ReservationHandler *handler.ReservationHandler
}

func (r *Router) SetupRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/health", handler.Health)

	userRoutes := router.Group("/users")
	{
		userRoutes.POST("", r.UserHandler.CreateUser)
		userRoutes.POST("login", r.UserHandler.Authenticate)
	}

	resourceRoutes := router.Group("/resources")
	{
		resourceRoutes.Use(middleware.AuthMiddleware(r.Key))

		resourceRoutes.POST("", r.ResourceHandler.CreateResource)
	}

	reservationRoutes := router.Group("/reservations")
	{
		reservationRoutes.Use(middleware.AuthMiddleware(r.Key))

		reservationRoutes.POST("", r.ReservationHandler.CreateReservation)
		reservationRoutes.GET("/:id", r.ReservationHandler.GetReservation)
		reservationRoutes.GET("/auth-user", r.ReservationHandler.GetUserReservations)
	}

	return router
}
