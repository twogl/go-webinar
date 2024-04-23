package main

import "fmt"

func main() {

	sl := []int{}
	FillSlice(sl)

	sl = make([]int, 0, 10)
	FillSlice(sl)

	// arr := [2]int{1, 1}
	// sl1 := arr[:]
	// sl1[0] = 99
	// fmt.Println(arr)

}

func FillSlice(sl []int) {
	for n := range 10 {
		fmt.Printf("capacity= %d, length= %d\n", cap(sl), len(sl))
		sl = append(sl, n)
	}
	fmt.Printf("capacity= %d, length= %d\n", cap(sl), len(sl))
	fmt.Println("=====================")

}
