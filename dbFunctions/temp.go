package dbFunctions

import (
	_ "github.com/go-sql-driver/mysql"
	"log"
	"strconv"
)

func AddPersonTemp(name string, balance float32, email string) string{
	db := getAccessDb()
	preparedQuery,err := db.Prepare("INSERT INTO people(name, balance, email) values (?,?,?)")
	if err != nil {
		log.Print("Prep Fails")
		log.Fatal(err)
	}

	response, err := preparedQuery.Exec(name,balance,email)
	if err != nil {
		log.Print("Exec Fail")
		log.Fatal(err)
	}

	lastId, err := response.LastInsertId()
	if err != nil {
		log.Print("Last Id fail")
		log.Fatal(err)
	}

	rowsAffected, err := response.RowsAffected()
	if err != nil {
		log.Print("Rows Affected")
		log.Fatal(err)
	}
	db.Close()
	log.Print("ID = " , lastId, ", Rows Affected: " , rowsAffected)
	return "success! ID = " + strconv.Itoa(int(lastId)) + ", Rows Affected: " + strconv.Itoa(int(rowsAffected))
}