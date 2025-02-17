package repositories

import (
	"Goland/internal/models"
)

var (
	users  = []models.User{}
	userId = 1
)

func GetAllUsers() []models.User {
	mu.Lock()
	defer mu.Unlock()
	return users
}

func CreateUser(user models.User) models.User {
	mu.Lock()
	defer mu.Unlock()
	user.ID = userId
	userId++
	users = append(users, user)
	return user
}
