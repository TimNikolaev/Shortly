package service

import (
	"shortener"
	"shortener/configs"
	"shortener/internal/repository"
)

type Service struct {
	shortener.UserRepository
	shortener.LinkRepository
	Config *configs.Config
}

func NewService(repo *repository.Repository, config *configs.Config) *Service {
	return &Service{
		LinkRepository: repo.LinkRepository,
		UserRepository: repo.UserRepository,
		Config:         config,
	}
}
