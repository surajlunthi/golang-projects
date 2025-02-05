package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type person struct {
	First string
}

func main() {

	http.HandleFunc("/encode", foo)
	http.HandleFunc("/decode", bar)
	http.ListenAndServe(":8081", nil)
}

func foo(w http.ResponseWriter, r *http.Request) {
	p1 := person{
		First: "jenny",
	}
	err := json.NewEncoder(w).Encode(p1)
	if err != nil {
		log.Println("Encode bad data :", err)
	}
}
func bar(w http.ResponseWriter, r *http.Request) {
	var p1 person
	err := json.NewDecoder(r.Body).Decode(&p1)
	if err != nil {
		log.Println(err)
	}
	log.Println("Person ", p1)
}
