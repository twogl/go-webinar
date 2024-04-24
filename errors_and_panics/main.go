package main

import (
	"errors"
	"log"

	"github.com/twogl/go-webinar/errors_and_panics/store"
)

func main() {

	// // defer func() {
	// // 	if r := recover(); r != nil {
	// // 		fmt.Println(fmt.Sprintf("panic recovered: %v", r))
	// // 	}
	// // }()

	// _ = store.LoadConfig()

	user, err := store.FindRandomUser()
	if errors.Is(err, store.ErrConnectionFailed) {
		log.Printf("find user by id: %s", err)
		// handle connection error
		return
	}
	if err != nil {
		log.Fatalf("unexpected error: %s", err)
		return
	}

	log.Println("user: ", user)
}
