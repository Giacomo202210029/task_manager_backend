package services

import (
	"Goland/internal/models"
	"Goland/internal/repositories"
)

func GetAllUsers() []models.User {
	return repositories.GetAllUsers()
}

func AddUser(user models.User) models.User {
	return repositories.CreateUser(user)
}
