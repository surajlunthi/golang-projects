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
	fmt.Println("print json ",string(bs))

	v := []person{};

	err = json.Unmarshal(bs,&v);

	if err!=nil{
		log.Panic("error in Unmarshal json",err)
	}
	fmt.Print(v);
}