package handler

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/marcos-nsantos/reservation-api/internal/entity"
	"github.com/marcos-nsantos/reservation-api/internal/service"
	"github.com/marcos-nsantos/reservation-api/internal/token"
)

type UserHandler struct {
	UserService *service.UserService
	Key         string
}

func NewUserHandler(userService *service.UserService, key string) *UserHandler {
	return &UserHandler{
		UserService: userService,
		Key:         key,
	}
}

func (h *UserHandler) CreateUser(c *gin.Context) {
	var createUserRequest CreateUserRequest

	if err := c.ShouldBindJSON(&createUserRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := entity.User{
		Name:     createUserRequest.Name,
		Email:    createUserRequest.Email,
		Password: createUserRequest.Password,
	}

	if err := h.UserService.CreateUser(&user); err != nil {
		if errors.Is(err, entity.ErrEmailAlreadyExists) {
			c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, user)
}

func (h *UserHandler) Authenticate(c *gin.Context) {
	var userLoginRequest UserLoginRequest

	if err := c.ShouldBindJSON(&userLoginRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := h.UserService.Authenticate(userLoginRequest.Email, userLoginRequest.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "email or password is incorrect"})
		return
	}

	jwtToken, err := token.GenerateJWT(h.Key, user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": jwtToken})
}
