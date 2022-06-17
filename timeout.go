package main

import (
	"fmt"
	"time"
)

func testTimeout() {
	c := make(chan bool, 1)
	go execute(c)

	fmt.Println("Waiting...")

	finished := false
	for !finished {
		select {
		case finished = <-c:
			fmt.Println("End!")
		case <-time.After(2 * time.Second):
			fmt.Println("Timeout!")
			finished = true
		}
	}
}

func execute(c chan<- bool) {
	time.Sleep(5 * time.Second)
	c <- true
}
