package models

// User struct
type User struct {
	//gorm.Model
	ID uint `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Name    string `json:"name"`
	Password string `json:"-"`
}
