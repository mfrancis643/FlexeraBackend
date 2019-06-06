package dbFunctions

import (
	"log"
)

func GetPeople() string{
	db := getAccessDb()
	rows,err := db.Query("SELECT * FROM people")
	if err != nil {
		log.Fatal(err)
	}
	return parseSqlResToJson(rows)
}