package main

import "fmt"

func main() {

	_ = [3]string{"one", "two", "three"}

	sl := []int{}
	FillSlice(sl)

	// sl := make([]int, 0, 10)
	// FillSlice(sl)

	// important!
	// arr := [3]int{1, 1, 1}
	// sl := arr[1:3]
	// sl[0] = 99
	// fmt.Println(arr)
	// fmt.Println(sl)

}

func FillSlice(sl []int) {
	for n := range 10 {
		fmt.Printf("capacity= %d, length= %d\n", cap(sl), len(sl))
		sl = append(sl, n)
	}
	fmt.Printf("capacity= %d, length= %d\n", cap(sl), len(sl))
	fmt.Println("=====================")

}
