package handlers

import (
	"net/http"
	"strconv"

	"github.com/krastomer/netflix-backend/database"
	"github.com/labstack/echo/v4"
)

func MovieHandlers(e *echo.Group) {
	e.GET("", getMovieDetailHandler)
}

func getMovieDetailHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.QueryParam("id"))
	md := database.GetMovieDetail(id)
	return c.JSON(http.StatusOK, md)
}

// TODO: getMovieDetail
// TODO: getListMovieFromActor
// TODO: getMovieEpisode
// TODO: getPosterMovie
