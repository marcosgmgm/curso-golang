package main

import (
	"database/sql"
	_"github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:admin!@/cursogo")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	//Update
	stmt, _ := db.Prepare("update usuarios set nome = ? where id = ?")

	stmt.Exec("Uóxinton Istive", 1)
	stmt.Exec("Sheristone Uasliska", 2)

	//Delete
	stmt2, _ := db.Prepare("delete from usuarios where id = ?")
	stmt2.Exec(3)
}
