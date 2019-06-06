package dbFunctions

import (
	"../entities"
	"log"
)

func AddPerson(person entities.PersonStr) string{
	db := getAccessDb()
	tx, _ := db.Begin()


	peopleQuery, err := tx.Prepare("INSERT INTO people(name, balance, email) values (?,?,?)")
	if err != nil {
		log.Print("Prep for peopleQuery Failed")
		log.Fatal(err)
		return "Prep for peopleQuery Failed"
	}
	addressQuery, err := tx.Prepare("INSERT INTO address(streetNameAndNumber, town, postcode) values (?,?,?)")
	if err != nil {
		log.Print("Prep for addressQuery Failed")
		log.Fatal(err)
		return "Prep for addressQuery Failed"
	}
	_, err = peopleQuery.Exec(person.Name,person.Balance,person.Email)
	if err != nil{
		log.Print("peopleQuery exec failed")
		log.Fatal(err)
		return "peopleQuery exec failedS"
	}
	_, err = addressQuery.Exec(person.Address.StreetNameAndNumber, person.Address.Town, person.Address.Postcode)
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
	return "success"


}