package service

import (
	"shortener"
	"shortener/configs"
	"shortener/internal/repository"
	"shortener/pkg/event"
)

type Service struct {
	shortener.UserRepository
	shortener.LinkRepository
	shortener.StatRepository
	Config   *configs.Config
	EventBus *event.EventBus
}

func NewService(r *repository.Repository, cfg *configs.Config, eBus *event.EventBus) *Service {
	s := &Service{
		LinkRepository: r.LinkRepository,
		UserRepository: r.UserRepository,
		StatRepository: r.StatRepository,
		Config:         cfg,
		EventBus:       eBus,
	}
	go s.AddClick()
	return s
}
