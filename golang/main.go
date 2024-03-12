package main

import (
	"encoding/json"
	"fmt"
	"net/http"
    "github.com/gorilla/mux"
    "strconv"
)

// Creación del servidor
func main() {
    router := mux.NewRouter()

    // Rutas a manejar
    router.HandleFunc("/", IndexHandler)
    router.HandleFunc("/users", UsersHandler)
    router.HandleFunc("/users/{id}", GetUserByIDHandler)

    // Servidor escuchando en el puerto 8080
    http.ListenAndServe(":8080", router)
}

// IndexHandler nos permite manejar la petición a la ruta '/' y retornar información relevante de la API.
func IndexHandler(w http.ResponseWriter, r *http.Request) {
    message := fmt.Sprintf(`Bienvenidos a la Golang API del reto Meli. Las rutas disponibles son:
- GET /users
- GET /users/<id>

User: {
    "id": "int,
    "name": "string",
    "email": "string"
}`)
    
    fmt.Fprint(w, message)
}

// UsersHandler nos permite manejar las peticiones a la ruta '/users'.
func UsersHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		GetUsersHandler(w, r)
	default:
		// Caso por defecto en caso de que se realice una petición con un método deferente a los esperados.
		http.Error(w, "Metodo no permitido", http.StatusBadRequest)
		return
	}
}

// GetUsersHandler nos permite manejar las peticiones a la ruta ‘/users’ con el método GET.
func GetUsersHandler(w http.ResponseWriter, r *http.Request) {
    // Puntero a una estructura de tipo User.
    n := new(User)
    
    // Solicitando todos los usuarios en la base de datos.
    users, err := n.GetAll()
    if err != nil {
        http.Error(w, err.Error(), http.StatusNotFound)
        return
    }
    // Convirtiendo el slice de usuarios a formato JSON, retorna un []byte y un error.
    j, err := json.Marshal(users)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
    // Escribiendo el código de respuesta.
    w.WriteHeader(http.StatusOK)
    // Estableciendo el tipo de contenido del cuerpo de la respuesta.
    w.Header().Set("Content-Type", "application/json")
    // Escribiendo la respuesta, es decir nuestro slice de usuarios en formato JSON.
    w.Write(j)
}

// GetUserByIDHandler nos permite manejar las peticiones a la ruta ‘/users/id’ con el método GET.
func GetUserByIDHandler(w http.ResponseWriter, r *http.Request) {
    // Puntero a una estructura de tipo User.
    n := new(User)

    // Se toman los parametros del request
    params := mux.Vars(r)

    // Se convierte el id de string a int
    userID, err := strconv.Atoi(params["id"])
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    // Solicitando el usuario en la base de datos.
    n.GetByID(userID)
    // Convirtiendo el usuario a formato JSON,
    j, err := json.Marshal(n)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
    // Escribiendo el código de respuesta.
    w.WriteHeader(http.StatusOK)
    // Estableciendo el tipo de contenido del cuerpo de la respuesta.
    w.Header().Set("Content-Type", "application/json")
    // Escribiendo la respuesta, es decir nuestro slice de usuarios en formato JSON.
    w.Write(j)
}
