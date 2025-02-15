package main //definimos que esto es el main.

//importamos log para imprimir mensajes
//las funciones http y nuestro router
import (
	"Goland/internal/routes"
	"log"
	"net/http"
)

// func main se ejecuta al iniciar la aplicacion
func main() {
	router := routes.SetupRouter()
	log.Println("Servidor corriendo en http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}

//se llama a routes para asignar y configurar todas las rutas
