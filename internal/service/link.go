package service

import "shortener"

func (s *Service) CreateLink(userID int, url string) (*shortener.Link, error) {
	link := shortener.NewLink(uint(userID), url)

	createdLink, err := s.LinkRepository.Create(link)
	if err != nil {
		return nil, err
	}

	return createdLink, nil
}

func (s *Service) GoToLink(hash string) (shortener.Link, error) {
	return shortener.Link{}, nil
}

func (s *Service) GetAllLinks(id uint) (shortener.Link, error) {
	return shortener.Link{}, nil
}

func (s *Service) Update(link shortener.Link) (shortener.Link, error) {
	return shortener.Link{}, nil
}

func (s *Service) DeleteLink(userID, linkID uint) error {
	return s.LinkRepository.Delete(userID, linkID)
}
