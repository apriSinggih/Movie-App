package router

import (
	"net/http"

	"github.com/apriSinggih/movie-app/internal/http/handler"
	"github.com/apriSinggih/movie-app/pkg/route"
)

func PublicRoutes(movieHandler handler.MovieHandler, userHandler handler.UserHandler) []route.Route {
	return []route.Route{
		{
			Method: http.MethodPost,
			Path:    "/login",
			Handler: userHandler.Login,
		},
		{
			Method: http.MethodGet,
			Path:    "/movies",
			Handler: movieHandler.GetMovies,
		},
		{
			Method: http.MethodGet,
			Path:    "/movies/:id",
			Handler: movieHandler.GetMovieByID,
		},
		{
			Method: http.MethodPost,
			Path:    "/movies",
			Handler: movieHandler.CreateMovie,
		},
		{
			Method: http.MethodPut,
			Path:    "/movies/:id",
			Handler: movieHandler.UpdateMovie,
		},
		{
			Method: http.MethodDelete,
			Path:    "/movies/:id",
			Handler: movieHandler.DeleteMovie,
		},
	}
}

func PrivateRoutes() []route.Route {
	return []route.Route{}
}
