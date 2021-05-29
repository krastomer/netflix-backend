package database

import (
	"github.com/krastomer/netflix-backend/models"
)

func AddUser(u models.User) error {
	d := GetDB()
	result := d.Create(&u)
	return result.Error
}

func GetUser(u string) models.User {
	d := GetDB()
	user := models.User{}
	d.First(&user, "email=?", u)
	return user
}

func GetUserPayment(u string) models.UserPayment {
	d := GetDB()
	payment := models.UserPayment{}
	d.First(&payment, "email=?", u)
	return payment
}

func SetUserPayment(payment models.UserPayment) error {
	d := GetDB()
	result := d.Where("email=?", payment.Email).Save(&payment)
	return result.Error
}
