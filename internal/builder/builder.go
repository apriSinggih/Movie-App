package builder

import (
	"github.com/apriSinggih/movie-app/config"
	"github.com/apriSinggih/movie-app/internal/http/handler"
	"github.com/apriSinggih/movie-app/internal/http/router"
	"github.com/apriSinggih/movie-app/internal/repository"
	"github.com/apriSinggih/movie-app/internal/service"
	"github.com/apriSinggih/movie-app/pkg/route"
	"gorm.io/gorm"
)

func BuildPublicRoutes(cfg *config.Config, db *gorm.DB) []route.Route {
	//repository
	userRepository := repository.NewUserRepository(db)
	movieRepository := repository.NewMovieRepository(db)
	//----------
	//service
	userService := service.NewUserService(userRepository)
	tokenService := service.NewTokenService(cfg.JWTConfig.SecretKey)
	movieService := service.NewMovieService(movieRepository)
	//----------
	//handler
	movieHandler := handler.NewMovieHandler(movieService)
	userHandler := handler.NewUserHandler(tokenService, userService)
	//----------
	return router.PublicRoutes(movieHandler, userHandler)
}

func BuildPrivateRoutes(db *gorm.DB) []route.Route {
	_ = repository.NewUserRepository(db)
	_ = repository.NewMovieRepository(db)
	return router.PrivateRoutes()
}
