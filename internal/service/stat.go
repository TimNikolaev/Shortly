package service

import (
	"shortener"
	"shortener/pkg/event"
	"time"

	"github.com/sirupsen/logrus"
)

func (s *Service) AddClick() {
	for msg := range s.EventBus.Subscribe() {
		if msg.Type == event.EventLinkVisited {
			linkID, ok := msg.Data.(uint)
			if !ok {
				logrus.Fatalln("Bad EventLinkVisited Data", msg.Data)
				continue
			}
			s.StatRepository.AddClick(linkID)
		}
	}
}

func (s *Service) GetStats(by string, from, to time.Time) ([]shortener.StatGetResponse, error) {
	return s.StatRepository.GetStats(by, from, to)
}
