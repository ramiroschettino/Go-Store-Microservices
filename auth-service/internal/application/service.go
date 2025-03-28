package application

import (
	"errors"

	"github.com/ramiroschettino/Go-Store-Microservices/auth-service/internal/domain"
)

type AuthService struct {
	userRepo  domain.UserRepository
	tokenRepo domain.TokenRepository
}

func NewAuthService(userRepo domain.UserRepository, tokenRepo domain.TokenRepository) *AuthService {
	return &AuthService{
		userRepo:  userRepo,
		tokenRepo: tokenRepo,
	}
}

func (s *AuthService) Register(email, password string) (*domain.User, error) {
	existingUser, _ := s.userRepo.FindByEmail(email)
	if existingUser != nil {
		return nil, errors.New("user already exists")
	}

	user := &domain.User{
		Email:    email,
		Password: password,
	}

	if err := user.HashPassword(); err != nil {
		return nil, err
	}

	if err := s.userRepo.Create(user); err != nil {
		return nil, err
	}

	return user, nil
}

func (s *AuthService) Login(email, password string) (*domain.TokenPair, error) {
	user, err := s.userRepo.FindByEmail(email)
	if err != nil {
		return nil, errors.New("invalid credentials")
	}

	if err := user.CheckPassword(password); err != nil {
		return nil, errors.New("invalid credentials")
	}

	accessToken, err := s.tokenRepo.GenerateAccessToken(user.ID)
	if err != nil {
		return nil, err
	}

	return &domain.TokenPair{
		AccessToken: accessToken,
	}, nil
}
