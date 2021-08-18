package models

type User struct {
	ID       uint   `json:"id" gorm:"primary_key"`
	AuthID   string `json:"auth_id" gorm:"unique"`
	Email    string `json:"email" gorm:"unique"`
	Nickname string `json:"nickname"`
	Name     string `json:"name"`
}
