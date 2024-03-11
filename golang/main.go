package main

import (
	"encoding/json"
	// "flag"
	"fmt"
	// "log"
	"net/http"
	// "strconv"
)

// var user User

func main() {
    message := "Hello, world!"
    fmt.Println(message)

    mux := http.NewServeMux()

    fmt.Println("Fin, se acabó")

    // Instancia de http.DefaultServerMux
    // mux := http.NewServeMux()
    // Ruta a manejar
    
    mux.HandleFunc("/", IndexHandler)
    mux.HandleFunc("/users", UsersHandler)

    // Servidor escuchando en el puerto 8080
    http.ListenAndServe(":8080", mux)

}

// IndexHandler nos permite manejar la petición a la ruta '/'
// y retornar "hola mundo" como respuesta al cliente.
func IndexHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, "hola mundo keslakeay")
}

func UsersHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		GetUsersHandler(w, r)
	// case http.MethodPost:
	// 	CreateNotesHandler(w, r)
	// case http.MethodPut:
	// 	UpdateNotesHandler(w, r)
	// case http.MethodDelete:
	// 	DeleteNotesHandler(w, r)
	default:
		// Caso por defecto en caso de que se realice una petición con un
		// método deferente a los esperados.
		http.Error(w, "Metodo no permitido", http.StatusBadRequest)
		return
	}
}

// GetNotesHandler nos permite manejar las peticiones a la ruta
// ‘/users’ con el método GET.
func GetUsersHandler(w http.ResponseWriter, r *http.Request) {
    // Puntero a una estructura de tipo Note.
    n := new(User)
    
    // Solicitando todas las notas en la base de datos.
    users, err := n.GetAll()
    if err != nil {
        http.Error(w, err.Error(), http.StatusNotFound)
        return
    }
    // Convirtiendo el slice de notas a formato JSON,
    // retorna un []byte y un error.
    j, err := json.Marshal(users)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
    // Escribiendo el código de respuesta.
    w.WriteHeader(http.StatusOK)
    // Estableciendo el tipo de contenido del cuerpo de la
    // respuesta.
    w.Header().Set("Content-Type", "application/json")
    // Escribiendo la respuesta, es decir nuestro slice de notas
    // en formato JSON.
    w.Write(j)
}

