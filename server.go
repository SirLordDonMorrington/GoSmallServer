package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type product struct {
	ID    string  `json:"product_id"`
	Name  string  `json:"name"`
	Price float32 `json:"price"`
}

var database []product

func main() {
	http.HandleFunc("/products", ruta1)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func ruta1(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case "GET":
		fmt.Fprintln(w, "List of products => ", database)
	case "POST":
		decoder := json.NewDecoder(r.Body)
		var p product
		err := decoder.Decode(&p)
		if err != nil {
			panic(err)
		}

		database = append(database, p) //Not checking if ID already exists

		fmt.Fprintln(w, "Added new product =>", database)
	case "PUT":
		decoder := json.NewDecoder(r.Body)
		var p product
		err := decoder.Decode(&p)
		if err != nil {
			panic(err)
		}

		for index, element := range database {
			if p.ID == element.ID {
				database[index] = p
			}
		}

		fmt.Fprintln(w, "Updated product =>", database)
	case "DELETE":
		decoder := json.NewDecoder(r.Body)
		var p product
		err := decoder.Decode(&p)
		if err != nil {
			panic(err)
		}

		for index, element := range database {
			if p.ID == element.ID {
				database = append(database[:index], database[index+1:]...)
			}
		}

		fmt.Fprintln(w, "Removed product => ", database)
	default:
		fmt.Fprintln(w, "Method ", r.Method, " not supported!")
	}
}
