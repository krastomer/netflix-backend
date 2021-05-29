package handlers

import (
	"net/http"

	"github.com/krastomer/netflix-backend/database"
	"github.com/krastomer/netflix-backend/models"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func PaymentHandlers(e *echo.Group) {
	e.Use(middleware.JWT([]byte(JWT_KEY)))

	e.GET("", getPaymentHandler)
	e.POST("", setPaymentHandler)
}

func getPaymentHandler(c echo.Context) error {
	u := getTokenUserEmail(c)
	payment := database.GetUserPayment(u)
	return c.JSON(http.StatusOK, payment)
}

func setPaymentHandler(c echo.Context) error {
	u := getTokenUserEmail(c)
	payment := models.UserPayment{Email: u}
	if err := c.Bind(&payment); err != nil {
		return err
	}
	if payment.DataInvalid() {
		return paymentInvalidError
	}
	err := database.SetUserPayment(payment)
	if err != nil {
		return internalServerError
	}
	return c.JSON(http.StatusOK, map[string]string{
		"message": "Update payment successed",
	})
}

// TODO: reBillingHandler
// TODO: cancelMemberShip <- using paymentMethod = NULL
