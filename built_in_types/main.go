package main

func main() {
	var _ int8 = 127
	var _ byte = 255 // alias for uint8
	var _ bool = true
	var _ string = "hello"
	var _ float64 = 55.76

	var _ complex128 = 4 - 5i
	var _ rune = 'ь' // alias for int32
	var _ int32 = 'ж'

	var _ interface{}
	var _ any // alias for interface{}

	var _ map[string]int
	var _ [5]string // array
	var _ []string  // slice

	var _ chan int
	var _ func(bool) int

	// var str string = "世界世界"
	// fmt.Println("string length=", len(str))

	// var runes []rune = []rune(str)
	// fmt.Println("runes array length=", len(runes))
	// for _, v := range runes {
	// 	fmt.Printf("%s ", string(v))
	// }
	// fmt.Println()
}
