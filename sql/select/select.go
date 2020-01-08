package main

import (
	"database/sql"
	_"github.com/go-sql-driver/mysql"
	"log"
)

type usuario struct {
	id 		int
	nome	string
}

func main() {
	db, err := sql.Open("mysql", "root:admin!@/cursogo")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, _ := db.Query("Select id, nome from usuarios")
	defer rows.Close()
	for rows.Next() {
		var u usuario
		rows.Scan(&u.id, &u.nome)
		log.Println(u)
	}
}
