package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type person struct {
First string
}

func main(){

	p1 := person{
		First : "jenny",
	}
	p2 := person{
		First:"James",
	}

	xp := []person{p1, p2}

	bs,err := json.Marshal(xp)

	if err!=nil{
		log.Panic(err)
	}
	fmt.Println("byte array",bs)
	fmt.Println(string(bs))


}