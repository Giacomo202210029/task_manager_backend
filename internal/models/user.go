package models

type User struct {
	ID       int    `json:"id"`
	Email    string `json:"email"`
	Password bool   `json:"password"`
}
