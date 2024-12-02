package service

import (
	"context"
	"errors"
	"time"

	"github.com/apriSinggih/movie-app/internal/entity"
	"github.com/apriSinggih/movie-app/internal/repository"
	"github.com/golang-jwt/jwt/v5"
)

type UserService interface {
	Login(ctx context.Context, username string, password string) (*entity.JWTCustomeClaims, error)
}

type userService struct {
	userRepository repository.UserRepository
}

func NewUserService(userRepository repository.UserRepository) UserService {
	return &userService{userRepository}
}

func (s *userService) Login(ctx context.Context, username string, password string) (*entity.JWTCustomeClaims, error) {
	user, err := s.userRepository.GetByUserName(ctx, username)
	if err != nil {
		return nil, errors.New("username or password is incorrect")
	}
	if user.Password != password {
		return nil, errors.New("username or password is incorrect")
	}

	expiredAt := time.Now().Local().Add(time.Minute * 10)

	claims := &entity.JWTCustomeClaims{
		Username: user.Username,
		FullName: user.FullName,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "movie-app",
			ExpiresAt: jwt.NewNumericDate(expiredAt),
		},
	}
	return claims, nil
}
