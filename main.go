package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
	"github.com/tanimutomo/sqlfile"
)

func crearBaseDeDatos() {
	db, err := sql.Open("postgres", "user=postgres host=localhost dbname=postgres sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	_, err = db.Exec(`DROP DATABASE IF EXISTS creditos`)
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(`CREATE DATABASE creditos`)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Bases de datos creada.")
}

func crearTablas() {
	db, err := sql.Open("postgres", "user=postgres host=localhost dbname=creditos sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	s := sqlfile.New()

	err = s.File("crear-tablas.sql")
	if err != nil {
		log.Fatal(err)
	}

	_, err = s.Exec(db)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Tablas creadas.")
}

func agregarClaves() {
	db, err := sql.Open("postgres", "user=postgres host=localhost dbname=creditos sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	s := sqlfile.New()

	err = s.File("agregar-claves.sql")
	if err != nil {
		log.Fatal(err)
	}

	_, err = s.Exec(db)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Claves agregadas.")
}

func eliminarClaves() {
	db, err := sql.Open("postgres", "user=postgres host=localhost dbname=creditos sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	s := sqlfile.New()

	err = s.File("eliminar-claves.sql")
	if err != nil {
		log.Fatal(err)
	}

	_, err = s.Exec(db)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Claves eliminadas.")
}

func main() {
	crearBaseDeDatos()
	crearTablas()
	agregarClaves()
	eliminarClaves()
}
