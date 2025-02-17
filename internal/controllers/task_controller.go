package controllers

import (
	"Goland/internal/models"
	"Goland/internal/services"
	"encoding/json"
	"net/http"
)

// GetTasks obtiene todas las tareas
// @Summary Obtener todas las tareas
// @Description Retorna la lista de tareas almacenadas
// @Tags Tasks
// @Produce json
// @Success 200 {array} models.Task
// @Router /tasks [get]
func GetTasks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	tasks := services.GetAllTasks()
	err := json.NewEncoder(w).Encode(tasks)
	if err != nil {
		return
	}
}

// CreateTask crea una nueva tarea
// @Summary Crear una nueva tarea
// @Description Crea una nueva tarea en la lista
// @Tags Tasks
// @Accept json
// @Produce json
// @Param task body models.Task true "Datos de la nueva tarea"
// @Success 201 {object} models.Task
// @Router /tasks [post]
func CreateTask(w http.ResponseWriter, r *http.Request) {
	var task models.Task
	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	newTask := services.AddTask(task)
	w.Header().Set("Content-Type", "application/json")
	err2 := json.NewEncoder(w).Encode(newTask)
	if err2 != nil {
		return
	}
}
