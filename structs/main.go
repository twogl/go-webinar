package main

import (
	"fmt"

	"github.com/twogl/go-webinar/structs/person"
)

func main() {
	p := person.Person{
		Name: "Kem Chan Int",
		Age:  45,
		Address: person.Address{
			City:        "Korea",
			StreetName:  "Happiness str.",
			HouseNumber: "7a",
		},
	}

	fmt.Println(p.Intro())

	fmt.Println(p.Address.HouseNumber)
	// fmt.Println(p.HouseNumber)

	//fmt.Println(p.passportID)
}
