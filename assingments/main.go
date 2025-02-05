package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type person struct {
	FirstName string
}

func main() {
	http.HandleFunc("/encode", encoder)
	http.HandleFunc("/decode", decoder)

	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		fmt.Println("Error: ", err)
	}

}
func encoder(w http.ResponseWriter, r *http.Request) {
	p1 := person{
		FirstName: "suraj",
	}
	p2 := person{
		FirstName: "divya",
	}

	var persons = []person{p1, p2}

	err := json.NewEncoder(w).Encode(persons)
	if err != nil {
		fmt.Println("Error: ", err)
	}

}

func decoder(w http.ResponseWriter, r *http.Request) {

	var p1 person

	err := json.NewDecoder(r.Body).Decode(&p1)
	if err != nil {
		fmt.Println("Error: ", err)
	}
	log.Println("Person ", p1)
}
