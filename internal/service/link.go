package service

import (
	"shortener"
	"shortener/pkg/event"

	"gorm.io/gorm"
)

func (s *Service) CreateLink(userID int, url string) (*shortener.Link, error) {
	link := shortener.NewLink(uint(userID), url)

	createdLink, err := s.LinkRepository.Create(link)
	if err != nil {
		return nil, err
	}

	return createdLink, nil
}

func (s *Service) GoToLink(hash string) (*shortener.Link, error) {
	link, err := s.LinkRepository.GetByHash(hash)
	if err != nil {
		return nil, err
	}
	go s.EventBus.Publish(event.Event{
		Type: event.EventLinkVisited,
		Data: link.ID,
	})

	return link, nil
}

func (s *Service) GetAllLinks(userID, limit, offset int) ([]shortener.Link, int64, error) {
	links, err := s.LinkRepository.GetAll(uint(userID), limit, offset)
	if err != nil {
		return nil, 0, err
	}

	count, err := s.LinkRepository.Count(uint(userID))
	if err != nil {
		return nil, 0, err
	}

	return links, count, nil
}

func (s *Service) UpdateLink(userID, linkID uint, url, hash string) (*shortener.Link, error) {
	return s.LinkRepository.Update(&shortener.Link{
		Model: gorm.Model{ID: linkID},
		URL:   url,
		Hash:  hash,
	}, userID)
}

func (s *Service) DeleteLink(userID, linkID uint) error {
	if _, err := s.LinkRepository.GetByID(linkID); err != nil {
		return err
	}

	return s.LinkRepository.Delete(userID, linkID)
}
