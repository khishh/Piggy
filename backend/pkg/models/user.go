package models

type User struct {
	ID        int    `gorm:"primary key;autoIncrement" json:"id"`
	Email     string `json:"email"`
	LastName  string `json:"last_name"`
	FirstName string `json:"first_name"`
	Picture   string `json:"picture"`
	Sub       string `json:"sub"`
}
