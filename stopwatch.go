package main

// chapter six
import (
	"fmt"
	"time"
)

func testStopwatch() {
	stopwatch(GenerateFibonacci(8))
	stopwatch(GenerateFibonacci(48))
	stopwatch(GenerateFibonacci(88))
}

func stopwatch(function func()) {
	start := time.Now()

	function()

	fmt.Printf("\nExecution time: %s\n\n", time.Since(start))
}

func GenerateFibonacci(n int) func() {
	return func() {
		a, b := 0, 1

		fib := func() int {
			a, b = b, a+b

			return a
		}

		for i := 0; i < n; i++ {
			fmt.Printf("%d ", fib())
		}
	}
}
