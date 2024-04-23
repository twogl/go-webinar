package main

import "fmt"

func main() {
	sl := []int{}
	for n := range 10 {
		fmt.Printf("capacity= %d, length= %d\n", cap(sl), len(sl))
		sl = append(sl, n)
	}

	fmt.Println("====================")

	sl2 := make([]int, 0, 10)
	for n := range 10 {
		fmt.Printf("capacity= %d, length= %d\n", cap(sl2), len(sl2))
		sl = append(sl2, n)
	}

	fmt.Println("====================")

	// arr := [2]int{1, 2}
	// arr2 := arr
	// arr[0] = 111 // <-- change 1st element

	// fmt.Println(arr)
	// fmt.Println(arr2)

}
