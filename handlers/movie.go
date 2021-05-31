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
	e.GET("/poster", getPosterMovieHandler)
}

func getMovieDetailHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.QueryParam("id"))
	md := database.GetMovieDetail(id)
	if md.Name == "" {
		return notFoundMovieError
	}
	return c.JSON(http.StatusOK, md)
}

func getListMovieFromActorHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.QueryParam("id"))
	lm := database.GetListMovieFromActor(id)
	if lm.Name == "" {
		return notFoundMovieError
	}
	return c.JSON(http.StatusOK, lm)
}

func getPosterMovieHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.QueryParam("id"))
	if id == 0 {
		return notFoundMovieError
	}
	link, err := database.GetPosterMovie(id)
	if err != nil {
		return err
	}
	return c.Redirect(http.StatusMovedPermanently, link)
}

// TODO: getMovieEpisode
