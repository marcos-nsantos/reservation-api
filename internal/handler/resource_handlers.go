package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/marcos-nsantos/reservation-api/internal/entity"
	"github.com/marcos-nsantos/reservation-api/internal/service"
)

type ResourceHandler struct {
	ResourceService *service.ResourceService
}

func NewResourceHandler(resourceService *service.ResourceService) *ResourceHandler {
	return &ResourceHandler{
		ResourceService: resourceService,
	}
}

func (h *ResourceHandler) CreateResource(c *gin.Context) {
	var createResourceRequest CreateResourceRequest
	if err := c.ShouldBindJSON(&createResourceRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resource := &entity.Resource{
		Name:        createResourceRequest.Name,
		Capacity:    createResourceRequest.Capacity,
		Description: createResourceRequest.Description,
	}

	if err := h.ResourceService.CreateResource(resource); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, resource)
}
