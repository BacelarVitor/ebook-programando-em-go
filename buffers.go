package main

import "fmt"

func testBuffers() {
	/*
		c := make(chan int, 3)
		go toProcuce(c)

		for value := range c {
			fmt.Println(value)
		}
	*/

	i, p := make(chan int), make(chan int)
	isReady := make(chan bool)

	nums := []int{1, 23, 42, 5, 8, 6, 7, 4, 99, 100}
	go pullApart(nums, i, p, isReady)

	var odds, evens []int
	finished := false

	for !finished {
		select {
		case n := <-i:
			odds = append(odds, n)
		case n := <-p:
			evens = append(evens, n)
		case finished = <-isReady:
		}
	}

	fmt.Printf("Odds: %v | Evens: %v\n", odds, evens)
}

// c chan<- just receive
// c <-chan int read only
func toProcuce(c chan<- int) {
	c <- 1
	c <- 2
	c <- 3

	close(c)
}

func pullApart(nums []int, i, p chan<- int, isReady chan<- bool) {
	for _, n := range nums {
		if n%2 == 0 {
			p <- n
		} else {
			i <- n
		}
	}

	isReady <- true
}
