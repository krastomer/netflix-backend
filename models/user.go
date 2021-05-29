package models

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
}

func (User) TableName() string {
	return "user"
}

func (UserPayment) TableName() string {
	return "user"
}

func (u *UserPayment) DataInvalid() bool {
	// TODO: check ExpDate correct
	return !(len(u.CardNumber) != 16 || len(u.ExpDate) != 5 || u.PlanId < 1 || u.PlanId > 4)
}
