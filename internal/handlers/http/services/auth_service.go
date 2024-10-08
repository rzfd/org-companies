package services

import (
	"errors"
	"os"

	"github.com/google/uuid"
	"github.com/rzfd/gorm-ners/internal/handlers/http/entities"
	"github.com/rzfd/gorm-ners/internal/handlers/http/repositories"
	"github.com/rzfd/gorm-ners/internal/handlers/http/security"
)

type AuthService struct {
	userRepo repositories.UserRepository
}

func NewAuthService(userRepo repositories.UserRepository) *AuthService {
	return &AuthService{userRepo: userRepo}
}

func (s *AuthService) RegisterUser(u *entities.Regis) error {
	existingUser, err := s.userRepo.FindByUsername(u.Username)
	if err != nil {
		return err
	}
	if existingUser != nil {
		return errors.New("user already exists")
	}

	hash, err := security.HashPassword(u.Password)
	if err != nil {
		return err
	}

	u.Password = hash
	return s.userRepo.CreateUser(u)
}

func (s *AuthService) AuthenticateUser(username, password string) (*entities.Regis, error) {
	user, err := s.userRepo.FindByUsername(username)
	if err != nil {
		return nil, err
	}
	if user == nil || !security.CheckPasswordHash(password, user.Password) {
		return nil, errors.New("invalid username or password")
	}
	return user, nil
}

func (s *AuthService) GenerateToken(userID uuid.UUID) (string, error) {
	jwtSecret := os.Getenv("JWT_TOKEN")
	if jwtSecret == "" {
		return "", errors.New("JWT secret is not set in environment variables at Generates Token")
	}
	return security.GenerateToken(userID.String(), jwtSecret)
}
