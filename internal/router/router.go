package router

import (
	"github.com/gin-gonic/gin"

	"github.com/marcos-nsantos/reservation-api/internal/handler"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/health", handler.Health)

	return router
}
