package main

import (
	"fmt"
	"strings"
)

// chapter three
func testStatistics() {
	words := readEntries()
	statistics := gatherStatistics(words)
	showStatistics(statistics)
}

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

func showStatistics(statistics map[string]int) {
	fmt.Println("Word count starting with each letter:")

	for initial, count := range statistics {
		fmt.Printf("%s = %d\n", initial, count)
	}
}
