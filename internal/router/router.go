package router

import (
	"github.com/gin-gonic/gin"

	"github.com/marcos-nsantos/reservation-api/internal/handler"
)

func SetupRouter(userHandler *handler.UserHandler) *gin.Engine {
	router := gin.Default()

	router.GET("/health", handler.Health)

	userRoutes := router.Group("/users")
	{
		userRoutes.POST("", userHandler.CreateUser)
	}

	return router
}
