package models

type Profile struct {
	ID       uint   `json:"id" gorm:"primary_key"`
	Email    string `json:"email" gorm:"unique"`
	Nickname string `json:"nickname"`
	Name     string `json:"name"`
	Picture  string `json:"picture"`
}
