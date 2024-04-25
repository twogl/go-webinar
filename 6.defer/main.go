package main

import "github.com/twogl/go-webinar/6.defer/magic"

func main() {
	for n := range 5 {
		magic.DoMagic(n)
	}

	// bytes, err := ignore.ReadGitignore()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(string(bytes))
}
