package service

import (
	"context"
	"errors"

	"github.com/apriSinggih/movie-app/internal/entity"
	"github.com/apriSinggih/movie-app/internal/repository"
)

type UserService interface {
	Login(ctx context.Context, username string, password string) (*entity.User, error)
}

type userService struct {
	userRepository repository.UserRepository
}

func NewUserService(userRepository repository.UserRepository) UserService {
	return &userService{userRepository}
}

func (s *userService) Login(ctx context.Context, username string, password string) (*entity.User, error) {
	user, err := s.userRepository.GetByUserName(ctx, username)
	if err != nil {
		return nil, errors.New("username or password is incorrect")
	}
	if user.Password != password {
		return nil, errors.New("username or password is incorrect")
	}
	return user, nil
}