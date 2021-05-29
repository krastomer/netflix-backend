package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

const (
	JWT_KEY = "september"
	EXP_KEY = 72
)

var (
	incorrectEmailError = echo.NewHTTPError(http.StatusBadRequest, "Incorrect email or password")
	internalServerError = echo.ErrInternalServerError
	duplicateEmailError = echo.NewHTTPError(http.StatusBadRequest, "Email has registered")
)

func SetHandlers(e *echo.Echo) {
	loginGroup := e.Group("/login")
	LoginHandlers(loginGroup)
}
