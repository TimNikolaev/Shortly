package service

import (
	"shortly"
	"shortly/internal/configs"
	"shortly/internal/repository"
	"shortly/pkg/event"
)

type Service struct {
	shortly.UserRepository
	shortly.LinkRepository
	shortly.StatRepository
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
