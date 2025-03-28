package application

import (
	"context"
	"user-service/api"
	"user-service/internal/domain"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	api.UnimplementedUserServiceServer
	repo domain.UserRepository
}

func NewUserService(repo domain.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) CreateUser(ctx context.Context, req *api.CreateUserRequest) (*api.UserResponse, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := &domain.User{
		Email:    req.Email,
		Password: string(hashedPassword),
		Name:     req.Name,
	}

	if err := s.repo.Create(ctx, user); err != nil {
		return nil, err
	}

	return &api.UserResponse{
		Id:    user.ID, 
		Email: user.Email,
		Name:  user.Name,
	}, nil
}

func (s *UserService) GetUser(ctx context.Context, req *api.GetUserRequest) (*api.UserResponse, error) {
	user, err := s.repo.FindByID(ctx, req.UserId)
	if err != nil {
		return nil, err
	}

	return &api.UserResponse{
		Id:    user.ID,
		Email: user.Email,
		Name:  user.Name,
	}, nil
}