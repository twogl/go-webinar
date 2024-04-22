package main

import (
	"errors"
	"log"

	"github.com/twogl/go-webinar/errors_and_panics/store"
)

func main() {
	user, err := store.FindUserById(-5)
	if errors.Is(err, store.ErrConnectionFailed) {
		log.Printf("find user by id: %s", err)
		// handle connection error
		return
	}
	if err != nil {
		log.Fatalf("find user by id: %s", err)
		return
	}

	log.Println("user: ", user)
}
