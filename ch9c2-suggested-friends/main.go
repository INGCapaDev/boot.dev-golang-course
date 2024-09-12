package main

func findSuggestedFriends(username string, friendships map[string][]string) []string {
	var suggestions []string
	alreadySuggested := map[string]bool{}

	// If the username doesn't exists return an nil slice
	directFriends, ok := friendships[username]
	if !ok {
		return suggestions
	}

	// Avoid suggest itself
	alreadySuggested[username] = true

	for _, name := range directFriends {
		// Avoid suggest already friends
		alreadySuggested[name] = true
		if fof, ok := friendships[name]; ok {
			for _, friend := range fof {
				if _, ok := alreadySuggested[friend]; !ok {
					suggestions = append(suggestions, friend)
					alreadySuggested[friend] = true
				}
			}
		}
	}

	return suggestions
}
