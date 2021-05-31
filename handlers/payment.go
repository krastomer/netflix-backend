package handlers

import (
	"net/http"
	"strconv"

	"github.com/krastomer/netflix-backend/database"
	"github.com/krastomer/netflix-backend/models"
	"github.com/labstack/echo/v4"
)

func PaymentHandlers(e *echo.Group) {
	e.GET("", getPaymentHandler)
	e.POST("", setPaymentHandler)
	e.GET("/rebill", reBillingHandler)
	e.DELETE("", cancelMemberShipHandler)
}

func getPaymentHandler(c echo.Context) error {
	u := getTokenUserEmail(c)
	payment := database.GetUserPayment(u)
	return c.JSON(http.StatusOK, map[string]string{
		"email":        payment.Email,
		"firstname":    payment.Firstname,
		"lastname":     payment.Lastname,
		"card_number":  payment.CardNumber,
		"exp_date":     payment.ExpDate,
		"cvc_code":     payment.SecurityCode,
		"next_billing": string(payment.NextBilling),
		"plan_id":      strconv.Itoa(payment.PlanId),
		"phone_number": payment.PhoneNumber,
	})
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

func reBillingHandler(c echo.Context) error {
	u := getTokenUserEmail(c)
	payment := database.GetUserPayment(u)
	if payment.CardNumber == "" {
		return paymentInvalidError
	}

	err := database.ReBillingPayment(&payment, u)
	if err != nil {
		return internalServerError
	}

	user := database.GetUserProfile(u)
	defer database.SetReceiptPayment(user)

	return c.JSON(http.StatusOK, map[string]string{
		"message":     "Re-Billing successed",
		"next-blling": string(payment.NextBilling),
	})
}

func cancelMemberShipHandler(c echo.Context) error {
	u := getTokenUserEmail(c)
	payment := database.GetUserPayment(u)
	err := database.CancelMemberShip(payment)

	if err != nil {
		return internalServerError
	}

	return c.JSON(http.StatusOK, map[string]string{
		"message": "Cancel Membership successed",
	})
}
