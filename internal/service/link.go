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

func (s *Service) Update(link shortener.Link) (shortener.Link, error) {
	return shortener.Link{}, nil
}

func (s *Service) DeleteLink(userID, linkID int) error {
	if _, err := s.LinkRepository.GetByID(uint(linkID)); err != nil {
		return err
	}

	return s.LinkRepository.Delete(uint(userID), uint(linkID))
}
