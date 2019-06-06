package main

import (
	"./dbFunctions"
	"./entities"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
	"strconv"
)

var serverAddress = "/people"

func enableCors(w *http.ResponseWriter){
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}

func Index (w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	switch r.Method {
	case "GET":

		fmt.Fprint(w, dbFunctions.GetPeople())
	case "POST":
		log.Print("hit Post")

		balance := r.FormValue("balance")
		formattedBalance, err := strconv.ParseFloat(balance, 32)
		if err != nil {
			log.Print(err)
			fmt.Fprint(w, err)
		}

		var address entities.AddressStr
		address.StreetNameAndNumber = r.FormValue("addressFirstLine")
		address.Town = r.FormValue("town")
		address.Postcode = r.FormValue("postcode")

		var person entities.PersonStr
		person.Name = r.FormValue("name")
		person.Balance = formattedBalance
		person.Email = r.FormValue("email")
		person.Address = address

		reply := dbFunctions.AddPerson(person)
		enableCors(&w)
		fmt.Fprint(w, reply)
	case "DELETE":
		log.Print("delete hit")
		id, ok := r.URL.Query()["id"]
		if !ok{
			log.Fatal("no id found")
		}

		reply := dbFunctions.DeletePerson(id[0])

		fmt.Fprint(w, reply)
	}
}

func getAddress (w http.ResponseWriter, r *http.Request){
	enableCors(&w)
	switch r.Method {
	case "GET":
		id, ok := r.URL.Query()["id"]
		if !ok{
			log.Fatal("no id found")
		}
		reply := dbFunctions.GetAddress(id[0])

		fmt.Fprint(w, reply)
	}
}

func main() {
	http.HandleFunc( serverAddress +"/", Index)
	http.HandleFunc( serverAddress + "/getAddress", getAddress)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

