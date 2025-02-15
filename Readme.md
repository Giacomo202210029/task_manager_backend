# BackendTaskManager

## Español

### Descripción
BackendTaskManager es una API REST desarrollada en Go para gestionar tareas (to-do list).  
Este proyecto utiliza una arquitectura modular basada en MVC/Clean Architecture, facilitando la separación de responsabilidades entre controladores, servicios, repositorios y modelos.

### Características
- Listar tareas
- Crear nuevas tareas
- Arquitectura modular y escalable
- Desarrollo y pruebas con GoLand de JetBrains

### Estructura del Proyecto
mi-backend/ ├── cmd/ │ └── server/ │ └── main.go // Punto de entrada del servidor └── internal/ ├── controllers/ // Controladores HTTP ├── models/ // Modelos de datos ├── repositories/ // Acceso a datos (ejemplo: repositorio en memoria) ├── routes/ // Configuración de rutas └── services/ // Lógica de negocio



