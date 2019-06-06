package dbFunctions

import (
	"log"
)
func DeletePerson(id string) string {

	db := getAccessDb()
	tx, _ := db.Begin()

	peopleQuery, err := tx.Prepare("DELETE FROM people where id = " + id)
	if err != nil {
		log.Print("Prep for peopleQuery Failed")
		log.Fatal(err)
		return "Prep for peopleQuery Failed"
	}
	addressQuery, err := tx.Prepare("DELETE FROM address where id = " + id)
	if err != nil {
		log.Print("Prep for addressQuery Failed")
		log.Fatal(err)
		return "Prep for addressQuery Failed"
	}
	_, err = peopleQuery.Exec()
	if err != nil{
		log.Print("peopleQuery exec failed")
		log.Fatal(err)
		return "peopleQuery exec failedS"
	}
	_, err = addressQuery.Exec()
	if err != nil{
		log.Print("addressQuery exec failed")
		log.Fatal(err)
		return "addressQuery exec failed"
	}

	err = tx.Commit()
	if err != nil{
		log.Fatal(err)
	}
	db.Close()
	return "successfully deleted person"

}