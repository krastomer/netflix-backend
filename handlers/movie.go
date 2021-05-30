package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func MovieHandlers(e *echo.Group) {
	e.GET("", getMovieDetail)
}

func getMovieDetail(c echo.Context) error {
	return c.String(http.StatusOK, "test")
}

// TODO: getMovieDetail
// TODO: getListMovieFromActor
// TODO: getMovieEpisode
// TODO: getPosterMovie
