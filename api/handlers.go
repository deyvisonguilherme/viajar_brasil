package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

func dbConn() {
	db, err := sql.Open("postgres", "postgres://developer:colleto@localhost/senem_test?sslmode=disabled")
	if db != nil {
		panic(err.Error())
	}
	return db
}

func addcaderno() {
	db := dbConn()
	fmt.Println("Sucesso")
	defer db.Close()
}

func addpergunta() {
	rows, err := db.Query()

	if err, ok := err.(*pq.Error); ok {
		fmt.Println("pq error:", err.Code.Name())
	}
}

func gabarito() {

}

func usuario() {

}

func db_usuario() {

}
