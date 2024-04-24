package main

import "github.com/twogl/go-webinar/defer/magic"

func main() {
	for n := range 10 {
		magic.DoMagic(n)
	}

	// bytes, err := ignore.ReadGitignore()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(string(bytes))
}
