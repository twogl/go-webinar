package magic

import (
	"fmt"
)

func DoMagic(num int) {
	fmt.Println("Hello world!")
	// defer fmt.Println("==================================") //  FILO
	defer fmt.Println("Bye world!")

	if num%2 == 0 {
		fmt.Println("Magic number is even!")
		return
	}
	fmt.Println("Magic number is odd!")
}
