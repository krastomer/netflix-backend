package handlers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/krastomer/netflix-backend/database"
	"github.com/krastomer/netflix-backend/models"
	"github.com/labstack/echo/v4"
)

func ViewerHandlers(e *echo.Group) {
	e.GET("", getListViewerHandler)
	e.POST("", createViewerHandler)
	e.DELETE("", deleteViewerHandler)
	e.GET("/token", getTokenViewerHandler)
}

func getListViewerHandler(c echo.Context) error {
	u, _ := getUserFromToken(c)
	listViewer := database.GetListViewer(u)
	return c.JSON(http.StatusOK, listViewer)
}

func createViewerHandler(c echo.Context) error {
	u, _ := getUserFromToken(c)
	viewer := models.Viewer{}
	if err := c.Bind(&viewer); err != nil {
		return err
	}
	err := database.CreateViewer(u, viewer)
	if err != nil {
		if errors.Is(err, database.GetMaxViewerError()) {
			return maxViewerError
		}
		return err
	}
	return c.JSON(http.StatusOK, map[string]string{
		"message": "Create Viewer succeed",
	})
}

func deleteViewerHandler(c echo.Context) error {
	u, _ := getUserFromToken(c)

	viewer := models.BodyViewer{Email: u}
	if err := c.Bind(&viewer); err != nil {
		return err
	}
	err := database.DeleteViewer(viewer)
	if err != nil {
		if errors.Is(err, database.GetNotFoundViewerError()) {
			return notFoundViewerError
		}
		return err
	}

	return c.JSON(http.StatusOK, map[string]string{
		"message": "Delete Viewer succeed",
	})

}

func getTokenViewerHandler(c echo.Context) error {
	u, _ := getUserFromToken(c)
	id, _ := strconv.Atoi(c.QueryParam("id"))

	return tokenJSON(c, u, id)
}
