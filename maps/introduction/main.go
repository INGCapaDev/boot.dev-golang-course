package main

import "fmt"

func getUserMap(names []string, phoneNumbers []int) (map[string]user, error) {
	if len(names) != len(phoneNumbers) {
		return nil, fmt.Errorf("invalid sizes")
	}

	users := map[string]user{}
	for idx, name := range names {
		users[name] = user{
			name:        name,
			phoneNumber: phoneNumbers[idx],
		}
	}

	return users, nil
}

type user struct {
	name        string
	phoneNumber int
}
