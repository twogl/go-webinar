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

func FindRandomUser() (*User, error) {
	user, err := getRandomUserFromRepo()
	if err != nil {
		return nil, fmt.Errorf("get user from repo: %w", err)
	}
	return user, nil
}

func getRandomUserFromRepo() (*User, error) {
	// always return conn error
	return nil, ErrConnectionFailed
}

func LoadConfig() string {
	panic("no config file found")
}
