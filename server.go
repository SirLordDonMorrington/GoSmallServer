package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/route1", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Accesaron aqui")
	})

	log.Fatal(http.ListenAndServe(":8080", nil))

}
