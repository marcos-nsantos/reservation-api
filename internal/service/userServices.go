package service

import (
	"github.com/marcos-nsantos/reservation-api/internal/entity"
	"github.com/marcos-nsantos/reservation-api/internal/repository"
	"github.com/marcos-nsantos/reservation-api/pkg/password"
)

type UserService struct {
	Repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{Repo: repo}
}

func (s *UserService) CreateUser(user *entity.User) error {
	hashedPassword, err := password.Hash(user.Password)
	if err != nil {
		return err
	}

	user.Password = hashedPassword
	return s.Repo.Create(user)
}
