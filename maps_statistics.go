package main

import "strings"

// chapter 3
func gatherStatistics(words []string) map[string]int {
	statistics := make(map[string]int)

	for _, word := range words {
		initialLetter := strings.ToUpper(string(word[0]))
		count, hasLetter := statistics[initialLetter]

		if hasLetter {
			statistics[initialLetter] = count + 1
			continue
		}
		statistics[initialLetter] = 1
	}

	return statistics
}
