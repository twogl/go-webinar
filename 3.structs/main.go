package main

import (
	"fmt"

	"github.com/twogl/go-webinar/3.structs/person"
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

	fmt.Println("House number:", p.Address.HouseNumber)

	// fmt.Println(p.passportID)
}
