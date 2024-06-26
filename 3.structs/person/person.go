package person

import "fmt"

type Person struct {
	Name       string // <-- public
	Age        int
	passportID string // <-- private

	Address Address
	// Address // <-- struct embedding
}

type Address struct {
	City        string
	StreetName  string
	HouseNumber string
}

// This function is a method for the Person type
func (pers Person) Intro() string {
	return fmt.Sprintf("This is %s! He is %d yo and lives in %s!", pers.Name, pers.Age, pers.Address.City)
}
