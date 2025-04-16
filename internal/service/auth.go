package service

import "shortener"

func (s *Service) CreateUser(user shortener.User) (int, error) {
	return 0, nil
}

func (s *Service) GetUser(email, password string) (shortener.User, error) {
	return shortener.User{}, nil
}
