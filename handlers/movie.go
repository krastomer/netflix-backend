package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/krastomer/netflix-backend/database"
	"github.com/labstack/echo/v4"
)

func MovieHandlers(e *echo.Group) {
	e.GET("", getMovieDetailHandler)
	e.GET("/actor", getListMovieFromActorHandler)
	e.GET("/poster", getPosterMovieHandler)
	e.POST("", addMyListMovieHandler)
	e.DELETE("", removeMyListMovieHandler)
	e.GET("/episode", getMovieEpisodeHandler)
	e.POST("/episode", setEpisodeHistoryHandler)
}

func getMovieDetailHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.QueryParam("id"))
	_, v := getUserFromToken(c)
	md := database.GetMovieDetail(id, v)
	fmt.Printf("%v\t%v\t%v\n", md.Name, md.MyList, v)
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

func addMyListMovieHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.QueryParam("id"))
	_, v := getUserFromToken(c)
	if id == 0 {
		return notFoundMovieError
	}
	err := database.AddMyListMovie(v, id)
	if err == 0 {
		return notMyListError
	}
	return c.JSON(http.StatusOK, map[string]string{
		"message": "Add MyList id_movie: " + strconv.Itoa(id) + " succeed",
	})
}

func removeMyListMovieHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.QueryParam("id"))
	_, v := getUserFromToken(c)
	if id == 0 {
		return notFoundMovieError
	}
	err := database.RemoveMyListMovie(v, id)
	if err == 0 {
		return notMyListError
	}
	return c.JSON(http.StatusOK, map[string]string{
		"message": "Remove MyList id_movie: " + strconv.Itoa(id) + " succeed",
	})
}

func getMovieEpisodeHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.QueryParam("id"))
	if id == 0 {
		return notFoundMovieError
	}
	listEpisode := database.GetMovieEpisode(id)
	return c.JSON(http.StatusOK, listEpisode)
}

func setEpisodeHistoryHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.QueryParam("id"))
	stop_time := c.QueryParam("stop")
	_, v := getUserFromToken(c)
	err := database.SetEpisodeHistory(id, v, stop_time)
	if err == 0 {
		return badEpisodeHistoryError
	}
	return c.JSON(http.StatusOK, map[string]string{
		"message": "Add History succeed",
	})
}
