package main

import "fmt"

func testBuffers() {
	c := make(chan int, 3)
	go toProcuce(c)

	for value := range c {
		fmt.Println(value)
	}
}

func toProcuce(c chan int) {
	c <- 1
	c <- 2
	c <- 3

	close(c)
}
