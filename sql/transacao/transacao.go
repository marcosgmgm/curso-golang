package main

import (
	"database/sql"
	_"github.com/go-sql-driver/mysql"
	"log"
)

func main() {
	db, err := sql.Open("mysql", "root:admin!@/cursogo")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	tx, _ := db.Begin()
	stmt, _ := tx.Prepare("Insert into usuarios(id, nome) values (?, ?)")

	stmt.Exec(2000, "Bia")
	stmt.Exec(2001, "Carlos")

	_, err = stmt.Exec(1, "Tiago") //Chave duplicada
	if err != nil {
		tx.Rollback()
		log.Fatal(err)
	}

	tx.Commit()

}
