package handlers

import (
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/krastomer/netflix-backend/database"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
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
	userInacitveError   = echo.NewHTTPError(http.StatusUnauthorized, "User Inactive")
	notFoundMovieError  = echo.NewHTTPError(http.StatusNotFound, "Not found movie or actor related id")
	maxViewerError      = echo.NewHTTPError(http.StatusBadRequest, "Maximum User Viewer")
	notFoundViewerError = echo.NewHTTPError(http.StatusNotFound, "Not found viewer")
)

func SetHandlers(e *echo.Echo) {
	loginGroup := e.Group("/login")
	LoginHandlers(loginGroup)

	paymentGroup := e.Group("/user/payment")
	paymentGroup.Use(middleware.JWT([]byte(JWT_KEY)))
	PaymentHandlers(paymentGroup)

	movieGroup := e.Group("/movie")
	movieGroup.Use(middleware.JWT([]byte(JWT_KEY)))
	movieGroup.Use(userActiveMiddleware)
	MovieHandlers(movieGroup)

	viewerGroup := e.Group("/viewer")
	viewerGroup.Use(middleware.JWT([]byte(JWT_KEY)))
	viewerGroup.Use(userActiveMiddleware)
	ViewerHandlers(viewerGroup)
}

func getTokenUserEmail(c echo.Context) string {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	name := claims["username"].(string)
	return name
}

func userActiveMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		u_active := getUserActive(c)
		if !u_active {
			return userInacitveError
		}
		return next(c)
	}
}

func getUserActive(c echo.Context) bool {
	u := database.GetUserProfile(getTokenUserEmail(c))
	t, _ := time.Parse("2006-01-02", string(u.NextBilling))
	return t.Unix() >= time.Now().Unix()
}
