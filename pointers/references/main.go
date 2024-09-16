package main

import "strings"

func removeProfanity(message *string) {
	profanityWords := []string{
		"fubb", "shiz", "witch",
	}
	for _, word := range profanityWords {
		*message = strings.ReplaceAll(*message, word, strings.Repeat("*", len(word)))
	}

}
