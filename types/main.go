package main

func main() {
	var _ int8 = 127
	var _ byte = 255 // alias for uint8
	var _ bool = true
	var _ string = "hello"
	var _ float64 = 55.76

	var _ complex128 = 4 - 5i
	var _ rune = '\t' // alias for int32

	var _ interface{}
	var _ any

}
