package store

import (
	"errors"
	"fmt"
)

type User struct {
	ID   int
	Name string
}

var ErrConnectionFailed error = errors.New("db conenction failed")

func FindUserById(userId int) (*User, error) {
	user, err := getUserFromRepo(userId)
	if err != nil {
		return nil, fmt.Errorf("get user from repo: %w", err)
	}
	return user, nil
}

func getUserFromRepo(userId int) (*User, error) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(fmt.Sprintf("panic recovered in loadUserData: %v", r))
		}
	}()

	if userId < 0 {
		panic("invalid user id")
	}

	return nil, ErrConnectionFailed
}
