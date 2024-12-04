package server

import (
	"net/http"

	"github.com/apriSinggih/movie-app/config"
	"github.com/apriSinggih/movie-app/internal/entity"
	"github.com/apriSinggih/movie-app/pkg/response"
	"github.com/apriSinggih/movie-app/pkg/route"
	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

type Server struct {
	*echo.Echo
}

func NewServer(cfg *config.Config, publicRoutes, privateRoutes []route.Route) *Server {
	e := echo.New()

	v1 := e.Group("/api/v1")
	if len(publicRoutes) > 0 {
		for _, route := range publicRoutes {
			v1.Add(route.Method, route.Path, route.Handler)
		}
	}
	if len(privateRoutes) > 0 {
		for _, route := range privateRoutes {
			v1.Add(route.Method, route.Path, route.Handler, JWTMiddleware(cfg.JWTConfig.SecretKey), RBACMiddlerware(route.Roles))
		}
	}
	return &Server{e}
}

func JWTMiddleware(secretKey string) echo.MiddlewareFunc {
	return echojwt.WithConfig(echojwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(entity.JWTCustomeClaims)
		},
		SigningKey: []byte(secretKey),
		ErrorHandler: func(ctx echo.Context, err error) error {
			return ctx.JSON(http.StatusUnauthorized, response.ErrorResponse(http.StatusUnauthorized, "Unauthorized"))
		},
	})
}

func RBACMiddlerware(roles []string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func (ctx echo.Context) error {
			user := ctx.Get("user").(*jwt.Token)
			claims := user.Claims.(*entity.JWTCustomeClaims)

			allowed := false

			for _, role := range roles {
				if role == claims.Role {
					allowed = true
					break
				}
			}

			if !allowed {
				return ctx.JSON(http.StatusForbidden, response.ErrorResponse(http.StatusForbidden, "Forbidden Access"))
			}
			return next(ctx)
		}
	}
}
