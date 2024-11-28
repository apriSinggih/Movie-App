package service

import (
	"context"

	"github.com/apriSinggih/movie-app/internal/entity"
	"github.com/apriSinggih/movie-app/internal/repository"
)

type MovieService interface {
	Create(ctx context.Context, movie *entity.Movie) error
	GetAll(ctx context.Context) ([]entity.Movie, error)
	GetByID(ctx context.Context, id int64) (*entity.Movie, error)
	GetByTitle(ctx context.Context, title string) (*entity.Movie, error)
	Update(ctx context.Context, movie *entity.Movie) error
	Delete(ctx context.Context, id int64) error
}

type movieService struct {
	movieRepository repository.MovieRepository
}

func NewMovieService(movieRepository repository.MovieRepository) MovieService {
	return &movieService{movieRepository}
}


func(s *movieService) Create(ctx context.Context, movie *entity.Movie) error{
	return s.movieRepository.Create(ctx, movie)
}
func (s *movieService) GetAll(ctx context.Context) ([]entity.Movie, error){
	return s.movieRepository.GetAll(ctx)
}
func (s *movieService)	GetByID(ctx context.Context, id int64) (*entity.Movie, error){
	return s.movieRepository.GetByID(ctx, id)
}
func (s *movieService)	GetByTitle(ctx context.Context, title string) (*entity.Movie, error){
	return s.movieRepository.GetByTitle(ctx, title)
}
func (s *movieService)	Update(ctx context.Context, movie *entity.Movie) error{
	return s.movieRepository.Update(ctx, movie)
}
func (s *movieService)	Delete(ctx context.Context, id int64) error{
	return s.movieRepository.Delete(ctx, id)
}