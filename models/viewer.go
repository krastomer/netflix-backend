package models

type Viewer struct {
	IDViewer  int    `json:"id_viewer"`
	IDAccount int    `json:"id_account"`
	Name      string `json:"name"`
	PinNumber string `json:"pin_number"`
	IsKid     bool   `json:"is_kid"`
}

type BodyViewer struct {
	IDViewer int `json:"id_viewer"`
	Email    string
}

func (BodyViewer) TableName() string {
	return "viewer"
}
