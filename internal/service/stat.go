package service

import (
	"shortly"
	"shortly/pkg/event"
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

func (s *Service) GetStats(linkID uint, by string, from, to time.Time) ([]shortly.GetStatResponse, error) {
	return s.StatRepository.GetStats(linkID, by, from, to)
}
