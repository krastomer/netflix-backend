package models

import "time"

type User struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserPayment struct {
	Email        string  `json:"email"`
	Firstname    string  `json:"firstname"`
	Lastname     string  `json:"lastname"`
	CardNumber   string  `json:"card_number"`
	ExpDate      string  `json:"exp_date"`
	SecurityCode string  `json:"cvc_code"`
	NextBilling  []uint8 `json:"next_blling"`
	PlanId       int     `json:"plan_id"`
	PhoneNumber  string  `json:"phone_number"`
}

type UserProfile struct {
	IDAccount   int
	Email       string
	NextBilling []uint8
}

type UserBilling struct {
	BillingDate time.Time `gorm:"column:billing_date"`
	IDAccount   int       `gorm:"column:id_account"`
}

func (u *UserPayment) DataInvalid() bool {
	// TODO: check ExpDate correct
	return len(u.CardNumber) != 16 ||
		len(u.ExpDate) != 5 ||
		u.PlanId < 1 ||
		u.PlanId > 4 ||
		u.Firstname == "" ||
		u.Lastname == "" ||
		len(u.SecurityCode) < 3 ||
		len(u.PhoneNumber) != 10
}

func (User) TableName() string {
	return "user"
}

func (UserPayment) TableName() string {
	return "user"
}

func (UserProfile) TableName() string {
	return "user"
}

func (UserBilling) TableName() string {
	return "billing"
}
