package service

import "shortener"

func (s *Service) Create(link shortener.Link) (shortener.Link, error) {
	return shortener.Link{}, nil
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
