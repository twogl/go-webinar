package main

import (
	"errors"
	"fmt"
	"log"

	"github.com/twogl/go-webinar/errors_and_panics/store"
)

func main() {

	defer func() {
		if r := recover(); r != nil {
			fmt.Println(fmt.Sprintf("panic recovered: %v", r))
		}
	}()

	_ = store.LoadConfig()

	user, err := store.FindUserById(5)
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
