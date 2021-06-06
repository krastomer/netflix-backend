package handlers

import (
	"net/http"
	"strconv"

	"github.com/krastomer/netflix-backend/database"
	"github.com/labstack/echo/v4"
)

func BrowseHandlers(e *echo.Group) {
	// e.GET("", getBrowseHandler)
	e.GET("/mylist", getMyListHandler)
	e.GET("/movie", getMovieBrowseHandler)
	e.GET("/tvshows", getTVShowsBrowseHandler)
	e.GET("/history", getHistoryHandler)
	e.GET("/top10", getTop10Handler)
	e.GET("/genres", getGenresBrowseHandler)
	e.GET("/banner", getBannerMovieHandler)
}

func getMovieBrowseHandler(c echo.Context) error {
	return c.String(http.StatusOK, "")
}

func getTVShowsBrowseHandler(c echo.Context) error {
	return c.String(http.StatusOK, "")
}

func getMyListHandler(c echo.Context) error {
	_, v := getUserFromToken(c)
	myList := database.GetMyList(v)
	return c.JSON(http.StatusOK, myList)
}

func getHistoryHandler(c echo.Context) error {
	_, v := getUserFromToken(c)
	hm := database.GetHistoryMovie(v)
	return c.JSON(http.StatusOK, hm)
}

func getTop10Handler(c echo.Context) error {
	tm := database.GetTop10Movie()
	return c.JSON(http.StatusOK, tm)
}

func getGenresBrowseHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.QueryParam("id"))
	listMovie := database.GetGenresMovie(id)
	if len(listMovie) == 0 {
		return notFoundMovieError
	}
	return c.JSON(http.StatusOK, listMovie)
}

func getBannerMovieHandler(c echo.Context) error {
	_, v := getUserFromToken(c)
	movie := database.GetBannerMovie(v)
	return c.JSON(http.StatusOK, movie)
}

// TODO: getMovieBrowseHandler
// TODO: getTVShowsBrowseHandler
