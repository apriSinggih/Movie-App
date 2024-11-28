package repository

import (
	"context"

	"github.com/apriSinggih/movie-app/internal/entity"
	"gorm.io/gorm"
)

type MovieRepository interface {
	Create(ctx context.Context, movie *entity.Movie) error
	GetAll(ctx context.Context) ([]entity.Movie, error)
	GetByID(ctx context.Context, id int64) (*entity.Movie, error)
	GetByTitle(ctx context.Context, title string) (*entity.Movie, error)
	Update(ctx context.Context, movie *entity.Movie) error
	Delete(ctx context.Context, id int64) error
}


type movieRepository struct {
	db *gorm.DB
}

func NewMovieRepository(db *gorm.DB) MovieRepository {
	return &movieRepository{db}
}

func (r *movieRepository) Create(ctx context.Context, movie *entity.Movie) error {
	return r.db.WithContext(ctx).Create(movie).Error
}

func (r *movieRepository) GetAll(ctx context.Context) ([]entity.Movie, error) {
	movies := make([]entity.Movie, 0)
	if err := r.db.WithContext(ctx).
	Find(&movies).Error; err != nil {
		return nil, err
	}
	return movies, nil
}


func (r *movieRepository) GetByID(ctx context.Context,id int64) (*entity.Movie, error) {
	movie := new(entity.Movie)
	if err := r.db.WithContext(ctx).Where("id = ?", id).First(&movie).Error; err != nil{
		return nil, err
	}
	return movie, nil
}

func (r *movieRepository) GetByTitle(ctx context.Context,title string) (*entity.Movie, error) {
	movie := new(entity.Movie)
	if err := r.db.WithContext(ctx).
	Where("title = ?", title).
	First(&movie).Error; err != nil {
		return nil, err
	}
	return movie, nil
}

func (r *movieRepository) Update(ctx context.Context, movie *entity.Movie) error {
	return r.db.WithContext(ctx).Save(movie).Error
}

func (r *movieRepository) Delete(ctx context.Context,id int64) error {
	return r.db.WithContext(ctx).Delete(id).Error
}