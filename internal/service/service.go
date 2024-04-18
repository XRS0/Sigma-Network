package service

import "github.com/XRS0/Sigma-Network/internal/repository"

type Service struct{}

func NewService(repos *repository.Repository) *Service {
	return &Service{}
}