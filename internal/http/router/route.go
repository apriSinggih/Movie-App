package router

import (
	"net/http"

	"github.com/apriSinggih/movie-app/internal/http/handler"
	"github.com/apriSinggih/movie-app/pkg/route"
)

func PublicRoutes(movieHandler handler.MovieHandler) []route.Route {
	return []route.Route{
		{
			Method: http.MethodGet,
			Path:    "/movies",
			Handler: movieHandler.GetMovies,
		},
	}
}

func PrivateRoutes() []route.Route {
	return []route.Route{}
}
