package handler

import (
	"net/http"

	"github.com/apriSinggih/movie-app/internal/http/dto"
	"github.com/apriSinggih/movie-app/internal/service"
	"github.com/apriSinggih/movie-app/pkg/response"
	"github.com/labstack/echo/v4"
)

type MovieHandler struct {
	movieService service.MovieService
}

func NewMovieHandler(movieService service.MovieService) MovieHandler {
	return MovieHandler{movieService}
}

func (h *MovieHandler) GetMovies(ctx echo.Context) error {
	movies, err := h.movieService.GetAll(ctx.Request().Context())
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, response.ErrorResponse(http.StatusInternalServerError, err.Error()))
	}
	return ctx.JSON(http.StatusOK, response.SuccessResponse("success get all movies", movies))
}

func (h *MovieHandler) GetMovieByID(ctx echo.Context) error {
	var req dto.GetMovieByIDRequest
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, response.ErrorResponse(http.StatusBadRequest, err.Error()))
	}
	movie, err := h.movieService.GetByID(ctx.Request().Context(), req.ID)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, response.ErrorResponse(http.StatusInternalServerError, err.Error()))
	}
	return ctx.JSON(http.StatusOK, response.SuccessResponse("success get movie by id", movie))
}

func (h *MovieHandler) CreateMovie(ctx echo.Context) error {
	var req dto.CreateMovieRequest
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, response.ErrorResponse(http.StatusBadRequest, err.Error()))
	}
	err := h.movieService.Create(ctx.Request().Context(), &req)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, response.ErrorResponse(http.StatusInternalServerError, err.Error()))
	}
	return ctx.JSON(http.StatusOK, response.SuccessResponse("success create movie", nil))
}

func (h *MovieHandler) UpdateMovie(ctx echo.Context) error {
	var req dto.UpdateMovieRequest
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, response.ErrorResponse(http.StatusBadRequest, err.Error()))
	}
	err := h.movieService.Update(ctx.Request().Context(), &req)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, response.ErrorResponse(http.StatusInternalServerError, err.Error()))
	}
	return ctx.JSON(http.StatusOK, response.SuccessResponse("success update movie", nil))
}

func (h *MovieHandler) DeleteMovie(ctx echo.Context) error {
	var req dto.GetMovieByIDRequest
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, response.ErrorResponse(http.StatusBadRequest, err.Error()))
	}
	movie, err := h.movieService.GetByID(ctx.Request().Context(), req.ID)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, response.ErrorResponse(http.StatusInternalServerError, err.Error()))
	}
	err = h.movieService.Delete(ctx.Request().Context(), movie.ID)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, response.ErrorResponse(http.StatusInternalServerError, err.Error()))
	}
	return ctx.JSON(http.StatusOK, response.SuccessResponse("success delete movie", nil))
}
