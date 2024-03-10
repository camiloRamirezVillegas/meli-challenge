package main

import "database/sql"

var db *sql.DB

func GetConnection() *sql.DB {
    // Para evitar realizar una nueva conexión en cada llamada a
    // la función GetConnection.
    if db != nil {
        return db
    }
    // Declaramos la variable err para poder usar el operador
    // de asignación "=" en lugar que el de asignación corta,
    // para evitar que cree una nueva variable db en este scope y
    // en su lugar que inicialice la variable db que declaramos a
    // nivel de paquete.
    var err error
    // Conexión a la base de datos

	dbPath := "../instance/site.db"

    db, err = sql.Open("sqlite3", dbPath)
    // db, err = sql.Open("sqlite3", "site.db")
    if err != nil {
        panic(err)
    }
    return db
}