package routes

//entonces esto va siempre arriba

//nuestras importaciones
import (
	"Goland/internal/controllers"
	"github.com/gorilla/mux"
	"github.com/swaggo/http-swagger"
)

func SetupRouter() *mux.Router {
	router := mux.NewRouter()

	//¿Que es mux?
	//Es un enrutador HTTP y
	//un envoltorio para la solicitud y respuesta de http.

	//¿Que es NewRouter?
	//Es una función que devuelve un nuevo enrutador.

	//¿Que es *mux.Router?
	//Es un enrutador HTTP y un envoltorio para la solicitud y respuesta de http.

	//¿Que es ese *?
	//Es un puntero, que es una dirección de memoria de una variable.

	//¿Que es := ?
	//Es una declaración corta de variable.
	//Se puede usar para declarar e
	//inicializar variables locales.

	//aqui se asignan las rutas
	//se asigna la ruta /tasks con el metodo GET a la funcion GetTasks
	router.HandleFunc("/tasks", controllers.GetTasks).Methods("GET")
	router.HandleFunc("/tasks", controllers.CreateTask).Methods("POST")
	router.HandleFunc("/users", controllers.GetUsers).Methods("GET")
	router.HandleFunc("/users", controllers.CreateUser).Methods("POST")
	router.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	// Puedes agregar más rutas para actualizar y eliminar tareas
	return router
}
