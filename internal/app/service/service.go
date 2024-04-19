package service

import "github.com/XRS0/Sigma-Network/internal/app/repository"

type Authorization interface {
	CreateUser(user repository.User) (int, error)
	GenerateToken(email, password string) (string, error)
}

type Service struct{
	Authorization
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
	}
}