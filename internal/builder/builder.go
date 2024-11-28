package builder

import (
	"github.com/apriSinggih/movie-app/internal/http/handler"
	"github.com/apriSinggih/movie-app/internal/http/router"
	"github.com/apriSinggih/movie-app/internal/repository"
	"github.com/apriSinggih/movie-app/internal/service"
	"github.com/apriSinggih/movie-app/pkg/route"
	"gorm.io/gorm"
)

func BuildPublicRoutes(db *gorm.DB) []route.Route {
	//repository
	userRepository := repository.NewUserRepository(db)
	movieRepository := repository.NewMovieRepository(db)
	//----------
	//service
	_ = service.NewUserService(userRepository)
	movieService := service.NewMovieService(movieRepository)
	//----------
	//handler
	movieHandler := handler.NewMovieHandler(movieService)
	//----------
	return router.PublicRoutes(movieHandler)
}

func BuildPrivateRoutes(db *gorm.DB) []route.Route {
	_ = repository.NewUserRepository(db)
	_ = repository.NewMovieRepository(db)
	return router.PrivateRoutes()
}
