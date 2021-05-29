package handlers

import (
	"net/http"

	"github.com/dgrijalva/jwt-go"
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
	paymentInvalidError = echo.NewHTTPError(http.StatusBadRequest, "Payment invalid")
)

func SetHandlers(e *echo.Echo) {
	loginGroup := e.Group("/login")
	LoginHandlers(loginGroup)

	paymentGroup := e.Group("/user/payment")
	PaymentHandlers(paymentGroup)
}

func getTokenUserEmail(c echo.Context) string {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	name := claims["username"].(string)
	return name
}

// TODO: checkUserActive Middleware
