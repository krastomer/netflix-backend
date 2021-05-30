package handlers

import (
	"net/http"
	"strconv"

	"github.com/krastomer/netflix-backend/database"
	"github.com/labstack/echo/v4"
)

func MovieHandlers(e *echo.Group) {
	e.GET("", getMovieDetailHandler)
	e.GET("/actor", getListMovieFromActorHandler)
}

func getMovieDetailHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.QueryParam("id"))
	md := database.GetMovieDetail(id)
	if md.Name == "" {
		return notfoundMovieError
	}
	return c.JSON(http.StatusOK, md)
}

func getListMovieFromActorHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.QueryParam("id"))
	lm := database.GetListMovieFromActor(id)
	if lm.Name == "" {
		return notfoundMovieError
	}
	return c.JSON(http.StatusOK, lm)
}

// TODO: getMovieEpisode
// TODO: getPosterMovie
