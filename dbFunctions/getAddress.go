package dbFunctions

import "log"

func GetAddress(id string) string {
	db := getAccessDb()
	rows,err := db.Query("SELECT * FROM address WHERE id = " + string(id))
	if err != nil {
		log.Fatal(err)
	}
	return parseSqlResToJson(rows)
}
