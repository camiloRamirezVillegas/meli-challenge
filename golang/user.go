package main

import (
	_ "github.com/mattn/go-sqlite3"
)

type User struct {
    ID  int `json:"id,omitempty"`
    Name    string  `json:"name"`
    Email   string  `json:"email"`
}

func (n *User) GetAll() ([]User, error) {
    db := GetConnection()
    q := `SELECT
            id, name, email
            FROM user`

    // Ejecución del query
    rows, err := db.Query(q)
    if err != nil {
        return []User{}, err
    }

    // Cerramos el recurso
    defer rows.Close()

    users := []User{}
    // El método Next retorna un bool, mientras sea true indicará que existe un valor siguiente para leer.
    for rows.Next() {
        // Escaneamos el valor actual de la fila e insertamos el retorno en los correspondientes campos del usuario.
        rows.Scan(
            &n.ID,
            &n.Name,
            &n.Email,
        )
        // Añadimos cada nuevo usuario al slice de usuarios que declaramos antes.
        users = append(users, *n)
    }
    return users, nil
}


func (n *User) GetByID(id int) (User, error) {
	db := GetConnection()
	q := `SELECT
		id, name, email
		FROM user WHERE id=?`

    // Ejecución del query
	err := db.QueryRow(q, id).Scan(
        // Asignación de los valores escaneados al puntero del usuario
		&n.ID, &n.Name, &n.Email,
	)
	if err != nil {
		return User{}, err
	}

	return *n, nil
}