package main

import (
	"fmt"
	"math/rand"
	"time"
)

// chapter four

func testInfiniteLoops() {
	generateSeed()
	infiniteLoop()
}

func generateSeed() {
	rand.Seed(time.Now().UnixNano())
}

func infiniteLoop() {
	n := 0
	for {
		n++
		i := rand.Intn(4200)
		fmt.Println(i)

		if i%42 == 0 {
			break
		}
	}

	fmt.Printf("Exit after %d iterations.\n", n)
}
