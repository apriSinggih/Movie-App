package handler

import (
	"net/http"

	"github.com/apriSinggih/movie-app/internal/http/dto"
	"github.com/apriSinggih/movie-app/internal/service"
	"github.com/apriSinggih/movie-app/pkg/response"
	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	tokenService service.TokenService
	userService  service.UserService
}

func NewUserHandler(tokenService service.TokenService, userService service.UserService) UserHandler {
	return UserHandler{tokenService, userService}
}

func (h *UserHandler) Login(ctx echo.Context) error {
	var req dto.UserLoginRequest

	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, response.ErrorResponse(http.StatusBadRequest, err.Error()))
	}

	claims, err := h.userService.Login(ctx.Request().Context(), req.Username, req.Password)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, response.ErrorResponse(http.StatusInternalServerError, err.Error()))
	}

	token, err := h.tokenService.GenerateToken(ctx.Request().Context(), *claims)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, response.ErrorResponse(http.StatusInternalServerError, err.Error()))
	}

	return ctx.JSON(http.StatusOK, response.SuccessResponse("success login", map[string]string{"token": token}))
}
