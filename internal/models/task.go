package models

type Task struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

//struct es una coleccion de campos

//¿para que es el type?
//Para definir un nuevo tipo de dato
//Task es un nuevo tipo de dato que es un struct
