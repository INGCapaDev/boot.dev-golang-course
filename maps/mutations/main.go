package main

import "fmt"

func deleteIfNecessary(users map[string]user, name string) (deleted bool, err error) {
	user, exists := users[name]
	if !exists {
		return false, fmt.Errorf("not found")
	}

	if user.scheduledForDeletion {
		delete(users, name)
	}
	return user.scheduledForDeletion, nil
}

type user struct {
	name                 string
	number               int
	scheduledForDeletion bool
}
