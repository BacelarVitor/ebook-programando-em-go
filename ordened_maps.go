package main

import (
	"fmt"
	"sort"
)

// chapter four

func generateMap() {
	squares := make(map[int]int, 15)

	for n := 1; n <= 15; n++ {
		squares[n] = n * n
	}

	numbers := make([]int, 0, len(squares))

	for n := range squares {
		numbers = append(numbers, n)
	}
	sort.Ints(numbers)

	for _, number := range numbers {
		square := squares[number]
		fmt.Printf("%dˆ2 =%d\n", number, square)
	}
}
