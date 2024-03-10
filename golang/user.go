package main
// package models

import (
	// "errors"
	// "time"
	_ "github.com/mattn/go-sqlite3"
)

type User struct {
    ID  int `json:"id,omitempty"`
    Name    string  `json:"name"`
    Email   string  `json:"email"`
    Password   string  `json:"password"`
}


func (n *User) GetAll() ([]User, error) {
    db := GetConnection()
    q := `SELECT
            id, name, email
            FROM user`

    // Ejecutamos la query
    rows, err := db.Query(q)
    if err != nil {
        return []User{}, err
    }
    // Cerramos el recurso
    defer rows.Close()
    // Declaramos un slice de notas para que almacene las
    // notas que retorna la petición.
    users := []User{}
    // El método Next retorna un bool, mientras sea true indicará
    // que existe un valor siguiente para leer.
    for rows.Next() {
        // Escaneamos el valor actual de la fila e insertamos el
        // retorno en los correspondientes campos de la nota.
        rows.Scan(
            &n.ID,
            &n.Name,
            &n.Email,
        )
        // Añadimos cada nueva nota al slice de notas que
        // declaramos antes.
        users = append(users, *n)
    }
    return users, nil
}


// func obtenerUsuario(w http.ResponseWriter, r *http.Request) {
//     id := r.URL.Query().Get("id")
//     if user, ok := users[id]; ok {
//         json.NewEncoder(w).Encode(user)
//     } else {
//         w.WriteHeader(http.StatusNotFound)
//         fmt.Fprintf(w, "Usuario no encontrado")
//     }
// }