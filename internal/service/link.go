package service

import "shortener"

func (s *Service) CreateLink(userID int, url string) (*shortener.Link, error) {
	link := shortener.NewLink(url)

	createdLink, err := s.LinkRepository.Create(link)
	if err != nil {
		return nil, err
	}

	return createdLink, nil
}

func (s *Service) GetByHash(hash string) (shortener.Link, error) {
	return shortener.Link{}, nil
}

func (s *Service) GetByID(id uint) (shortener.Link, error) {
	return shortener.Link{}, nil
}

func (s *Service) Update(link shortener.Link) (shortener.Link, error) {
	return shortener.Link{}, nil
}

func (s *Service) Delete(id uint) error {
	return nil
}
