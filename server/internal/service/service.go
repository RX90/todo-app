package service

import (
	"github.com/RX90/Todo-App/server/internal/repository"
	"github.com/RX90/Todo-App/server/internal/user"
)

type Authorization interface {
	CreateUser(user user.User) error
	GetUserId(username, password string) (string, error)
	NewAccessToken(userId string) (string, error)
	NewRefreshToken(userId string) (string, error)
}

type Service struct {
	Authorization
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: newAuthService(repos.Authorization),
	}
}
