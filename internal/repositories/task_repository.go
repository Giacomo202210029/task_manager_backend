package repositories

import (
	"Goland/internal/models"
	"sync"
)

var (
	tasks = []models.Task{}
	//aqui se define la variable tasks como un slice de Task
	//un slice es una estructura de datos que representa
	//una secuencia de elementos
	//como un arreglo? si, pero con una longitud dinamica
	//dinamica? si, se puede cambiar la longitud de un slice
	taskId = 1
	//se define la variable taskId como un entero
	mu sync.Mutex
	//se define a mu que sirve para abrir y cerrar el acceso a la seccion de codigo
	//esto lo usamos luego
)

func GetAllTasks() []models.Task {
	//funcion GetAllTasks que retorna un slice de Task
	mu.Lock()
	//se bloquea el acceso a la seccion de codigo
	defer mu.Unlock()
	//se desbloquea el acceso a la seccion de codigo
	return tasks
	//retorna la variable tasks
}

func CreateTask(task models.Task) models.Task {
	mu.Lock()
	defer mu.Unlock() //lleva un defer, para que suceda al final de la funcion
	//y si sucede un error en mitad de la funcion, se ejecuta igual
	task.ID = taskId
	taskId++
	tasks = append(tasks, task)
	return task
}
