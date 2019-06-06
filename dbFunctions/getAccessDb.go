package dbFunctions

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func getAccessDb () *sql.DB{
	db, err := sql.Open(
		"mysql",
		"root@(localhost:3306)/flexera")

	if err != nil {
		log.Fatal(err)
	} else{
		log.Print("Db connection success")
	}
	return db
}