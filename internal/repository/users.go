package repository

import (
	"context"

	"github.com/apriSinggih/movie-app/internal/entity"
	"gorm.io/gorm"
)

type UserRepository interface {
	GetByUserName(ctx context.Context, username string) (*entity.User, error)
	Create(ctx context.Context, user *entity.User) error
	GetAll(ctx context.Context) ([]entity.User, error)
	GetByID(ctx context.Context, id int64) (*entity.User, error)
	UpdateUser(ctx context.Context, user *entity.User) error
	DeleteUser(ctx context.Context, id int64) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db}
}

func (u *userRepository) GetByUserName(ctx context.Context, username string) (*entity.User, error) {
	result := new(entity.User)
	if err := u.db.WithContext(ctx).Where("username = ?", username).First(&result).Error; err != nil {
		return nil, err
	}
	return result, nil
}


func (u *userRepository) Create(ctx context.Context, user *entity.User) error {
	return u.db.WithContext(ctx).Create(user).Error
}

func (u *userRepository) GetAll(ctx context.Context) ([]entity.User, error) {
	users := make([]entity.User, 0)
	if err := u.db.WithContext(ctx).Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (u *userRepository) GetByID(ctx context.Context, id int64) (*entity.User, error) {
	result := new(entity.User)
	if err := u.db.WithContext(ctx).Where("id = ?", id).First(&result).Error; err != nil {
		return nil, err
	}
	return result, nil
}

func (u *userRepository) UpdateUser(ctx context.Context, user *entity.User) error {
	return u.db.WithContext(ctx).Save(user).Error
}
func (u *userRepository) DeleteUser(ctx context.Context, id int64) error {
	return u.db.WithContext(ctx).Delete(&entity.User{}, id).Error
}