// cap two
package main

import (
	"fmt"
	"os"
	"strconv"
)

func testQuicksort() {
	entries := readEntries()
	numbers := convertEntries(entries)
	orderedNumbers := quicksort(numbers)
	fmt.Println(orderedNumbers)
}

func convertEntries(entries []string) []int {
	numbers := make([]int, len(entries))

	for i, n := range entries {
		number, err := strconv.Atoi(n)
		if err != nil {
			fmt.Printf("%s is not a valid number!\n", n)
			os.Exit(-1)
		}
		numbers[i] = number
	}
	return numbers
}

func quicksort(numbers []int) []int {
	if len(numbers) <= 1 {
		return numbers
	}

	n := make([]int, len(numbers))
	copy(n, numbers)

	pivotIndex := len(n) / 2
	pivot := n[pivotIndex]

	// remove pivot, get the start of the slice until pivot (excluding pivot), get everything after the pivot
	// the ... indicates that all elements of the second argument (n[pivotIndex=1:]) should be added
	n = append(n[:pivotIndex], n[pivotIndex+1:]...)
	lower, bigger := split(n, pivot)
	return append(
		append(quicksort(lower), pivot),
		quicksort(bigger)...,
	)
}

func split(numbers []int, pivot int) (lower []int, bigger []int) {
	for _, n := range numbers {
		if n <= pivot {
			lower = append(lower, n)
			continue
		}
		bigger = append(bigger, n)
	}

	return lower, bigger
}
