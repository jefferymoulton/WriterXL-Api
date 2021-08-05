package models

type User struct {
	ID           uint   `json:"id" gorm:"primary_key"`
	EmailAddress string `json:"emailAddress" gorm:"unique"`
	FirstName    string `json:"firstName"`
	LastName     string `json:"lastName"`
}
