package services

import (
	"Goland/internal/models"
	"Goland/internal/repositories"
)

func GetAllTasks() []models.Task {
	return repositories.GetAllTasks()
}

func AddTask(task models.Task) models.Task {
	return repositories.CreateTask(task)
}
