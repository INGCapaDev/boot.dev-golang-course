package main

func indexOfFirstBadWord(msg []string, badWords []string) int {
	for index, msgWord := range msg {
		for _, badWord := range badWords {
			if msgWord == badWord {
				return index
			}
		}
	}
	return -1
}
