package main

func getCounts(messagedUsers []string, validUsers map[string]int) {
	for _, user := range messagedUsers {
		if _, isValid := validUsers[user]; isValid {
			validUsers[user] += 1
		}
	}
}
