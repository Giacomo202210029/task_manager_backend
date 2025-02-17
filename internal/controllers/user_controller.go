package controllers

import (
	"Goland/internal/models"
	"Goland/internal/services"
	"encoding/json"
	"net/http"
)

// GetUsers obtiene todos los usuarios
// @Summary Obtener todos los usuarios
// @Description Retorna la lista de usuarios almacenados
// @Tags Users
// @Produce json
// @Success 200 {array} models.User
// @Router /users [get]
func GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	users := services.GetAllUsers()
	err := json.NewEncoder(w).Encode(users)
	if err != nil {
		return
	}
}

// CreateUser crea un nuevo usuario
// @Summary Crear un nuevo usuario
// @Description Crea un nuevo usuario y lo agrega a la lista de usuarios
// @Tags Users
// @Accept json
// @Produce json
// @Param user body models.User true "Datos del nuevo usuario"
// @Success 201 {object} models.User
// @Failure 400 {string} string "Bad Request"
// @Router /users [post]
func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	newuser := services.AddUser(user)
	w.Header().Set("Content-Type", "application/json")
	err2 := json.NewEncoder(w).Encode(newuser)
	if err2 != nil {
		return
	}
}
