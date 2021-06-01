package handlers

import (
	"net/http"
	"net/mail"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/krastomer/netflix-backend/database"
	"github.com/krastomer/netflix-backend/models"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

func LoginHandlers(e *echo.Group) {
	e.POST("/register", registerHandler)
	e.POST("", loginHandler)
	// TODO: recoveryPassword
}

func registerHandler(c echo.Context) error {
	u := models.User{}
	if err := c.Bind(&u); err != nil {
		return err
	}

	if !checkEmailPassword(u) {
		return incorrectEmailError
	}

	u.Password = hashPassword(u.Password)
	err := database.AddUser(u)

	if err != nil {
		if strings.Contains(err.Error(), "Duplicate entry") {
			return duplicateEmailError
		}
		return internalServerError
	}

	return tokenJSON(c, u.Email, -1)
}

func loginHandler(c echo.Context) error {
	u := models.User{Email: c.FormValue("username"), Password: c.FormValue("password")}

	if !checkEmailPassword(u) {
		return incorrectEmailError
	}

	r_u := database.GetUser(u.Email)
	result := checkPassword(r_u.Password, u.Password)
	if result {
		return incorrectEmailError
	}

	return tokenJSON(c, u.Email, -1)
}

func checkEmailPassword(u models.User) bool {
	_, err := mail.ParseAddress(u.Email)
	return !(err != nil || len(u.Password) < 8)
}

func hashPassword(p string) string {
	h, _ := bcrypt.GenerateFromPassword([]byte(p), 10)
	return string(h)
}

func generateToken(u string, v int) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = u
	claims["exp"] = time.Now().Add(time.Hour * EXP_KEY).Unix()
	claims["viewer"] = v

	t, err := token.SignedString([]byte(JWT_KEY))
	return t, err
}

func tokenJSON(c echo.Context, s string, v int) error {
	t, err := generateToken(s, v)
	if err != nil {
		return internalServerError
	}
	return c.JSON(http.StatusOK, map[string]string{
		"token": t,
	})
}

func checkPassword(h string, p string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(h), []byte(p))
	return err != nil
}
