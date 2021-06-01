package handlers

import (
	"net/http"

	"github.com/krastomer/netflix-backend/database"
	"github.com/labstack/echo/v4"
)

func BrowseHandlers(e *echo.Group) {
	e.GET("", getBrowseHandler)
	e.GET("/mylist", getMyListHandler)
}

func getBrowseHandler(c echo.Context) error {
	return c.String(http.StatusOK, "test")
}

func getMyListHandler(c echo.Context) error {
	_, v := getUserFromToken(c)
	myList := database.GetMyList(v)
	return c.JSON(http.StatusOK, myList)
}
