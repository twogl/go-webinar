package main

import "fmt"

func main() {
	var firstName string = "Penko"
	lastName := "Penkov"

	if firstName == "Penko" {
		fmt.Println(firstName, lastName)
	}

	// same as while(true)
	for {
		fmt.Println("endless loop")
	}

	var condition bool
	for condition {
		fmt.Println("condition is true")
	}

	numbers := []int{1, 2, 3, 4, 5, 6, 7}
	for i := 0; i < len(numbers); i++ {
		fmt.Println("numbers[i]=", numbers[i])
	}

	pplAge := map[string]int{
		"Alex": 30,
		"Bob":  20,
		"John": 25,
	}
	for key, val := range pplAge {
		fmt.Printf("key= %s, value= %d\n", key, val)
	}

	_, ok := pplAge["Max"]
	if !ok {
		fmt.Println("Adding `Max` to the map")
		pplAge["Max"] = 35
	}
}
