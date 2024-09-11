package main

func getNameCounts(names []string) map[rune]map[string]int {
	nameCounts := map[rune]map[string]int{}
	for _, name := range names {
		rune := []rune(name)[0]
		if _, ok := nameCounts[rune]; !ok {
			nameCounts[rune] = map[string]int{}
		}
		nameCounts[rune][name] += 1
	}
	return nameCounts
}
