package handlers

import (
	"errors"
	"net/http"

	"github.com/krastomer/netflix-backend/database"
	"github.com/krastomer/netflix-backend/models"
	"github.com/labstack/echo/v4"
)

func ViewerHandlers(e *echo.Group) {
	e.GET("", getListViewerHandler)
	e.POST("", createViewerHandler)
	e.DELETE("", deleteViewerHandler)
}

func getListViewerHandler(c echo.Context) error {
	u := getTokenUserEmail(c)
	listViewer := database.GetListViewer(u)
	return c.JSON(http.StatusOK, listViewer)
}

func createViewerHandler(c echo.Context) error {
	u := getTokenUserEmail(c)
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
		"message": "Create Viewer successed",
	})
}

func deleteViewerHandler(c echo.Context) error {
	u := getTokenUserEmail(c)

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
		"message": "Delete Viewer successed",
	})

}

// func updateViewerHandler(c echo.Context) error {
// 	u := getTokenUserEmail()
// }
