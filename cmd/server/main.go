package main

import (
	"log"
	"net/http"

	_ "Goland/docs" // Importa la documentación generada por Swag
	"Goland/internal/routes"

	httpSwagger "github.com/swaggo/http-swagger"
)

func main() {
	router := routes.SetupRouter()

	// Ruta para acceder a Swagger UI
	router.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	log.Println("Servidor corriendo en http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
