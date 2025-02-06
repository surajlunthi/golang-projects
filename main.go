package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

type person struct {
	First string
}

func hashPassword(password string) ([]byte, error) {
	bs, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		return nil, err
	}
	return bs, nil
}

func comparePassword(password string, hashPassword []byte) error {
	err := bcrypt.CompareHashAndPassword(hashPassword, []byte(password))
	if err != nil {
		return err
	}
	return nil

}

func main() {

	pass := "123456789suraj"

	v, err := hashPassword(pass)
	if err != nil {
		fmt.Println("Error is ", err)
	}

	err = comparePassword(pass, v)
	if err != nil {
		fmt.Println("Password comparison error:", err)
	} else {
		fmt.Println("Password matched successfully!")
	}

	// http.HandleFunc("/encode", foo)
	// http.HandleFunc("/decode", bar)
	// http.ListenAndServe(":8081", nil)
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
