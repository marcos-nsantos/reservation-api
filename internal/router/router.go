package router

import (
	"github.com/gin-gonic/gin"

	"github.com/marcos-nsantos/reservation-api/internal/handler"
)

type Router struct {
	UserHandler     *handler.UserHandler
	ResourceHandler *handler.ResourceHandler
}

func (r *Router) SetupRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/health", handler.Health)

	userRoutes := router.Group("/users")
	{
		userRoutes.POST("", r.UserHandler.CreateUser)
		userRoutes.POST("/login", r.UserHandler.Authenticate)
	}

	resourceRoutes := router.Group("/resources")
	{
		resourceRoutes.POST("", r.ResourceHandler.CreateResource)
	}

	return router
}
