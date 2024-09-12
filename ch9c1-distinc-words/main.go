package main

import "strings"

func countDistinctWords(messages []string) int {
	wMap := map[string]bool{}

	for _, msg := range messages {
		words := strings.Fields(msg)
		for _, word := range words {
			wLower := strings.ToLower(word)
			if _, ok := wMap[wLower]; !ok {
				wMap[wLower] = true
			}
		}

	}
	return len(wMap)
}
