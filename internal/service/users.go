package service

import (
	"context"
	"errors"
	"time"

	"github.com/apriSinggih/movie-app/internal/entity"
	"github.com/apriSinggih/movie-app/internal/http/dto"
	"github.com/apriSinggih/movie-app/internal/repository"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	Login(ctx context.Context, username string, password string) (*entity.JWTCustomeClaims, error)
	Register(ctx context.Context, req dto.UserRegisterRequest) error
	GetAll(ctx context.Context) ([]entity.User, error)
	GetByID(ctx context.Context, id int64) (*entity.User, error)
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

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, errors.New("username or password is incorrect")
	}


	expiredAt := time.Now().Local().Add(time.Minute * 10)

	claims := &entity.JWTCustomeClaims{
		Username: user.Username,
		FullName: user.FullName,
		Role: user.Role,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "movie-app",
			ExpiresAt: jwt.NewNumericDate(expiredAt),
		},
	}
	return claims, nil
}

func (s *userService) Register(ctx context.Context, req dto.UserRegisterRequest) error{
	user := new(entity.User)

	user.Username = req.Username

	exist, err := s.userRepository.GetByUserName(ctx, user.Username)
	if err == nil && exist != nil {
		return errors.New("username already exists")
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)
	user.FullName = req.FullName
	user.Role = "user"
	return s.userRepository.Create(ctx, user)
}

func (s *userService) GetAll(ctx context.Context) ([]entity.User, error) {
	return s.userRepository.GetAll(ctx)
}

func (s *userService) GetByID(ctx context.Context, id int64) (*entity.User, error) {
	return s.userRepository.GetByID(ctx, id)
}