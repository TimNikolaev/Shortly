package service

import (
	"crypto/sha1"
	"fmt"
	"shortly"
	"shortly/pkg/jwt"
)

const salt = "qwerty123456789"

func (s *Service) CreateUser(user *shortly.User) (uint, error) {
	user.Password = generatePasswordHash(user.Password)
	return s.UserRepository.CreateUser(user)
}

func (s *Service) GenerateToken(email, password string) (string, error) {
	user, err := s.UserRepository.GetUser(email, generatePasswordHash(password))
	if err != nil {
		return "", err
	}

	return jwt.NewJWT(s.config.AuthConfig.Secret).Generate(user.ID)
}

func (s *Service) ParseToken(accessToken string) (int, error) {
	return jwt.NewJWT(s.config.AuthConfig.Secret).Parse(accessToken)
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))
	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
