package main

import "fmt"

// chapter four

func testNamedLoop() {
	generateNamedLoop()
}

func generateNamedLoop() {
	var i int

loop:
	for i = 0; i < 10; i++ {
		fmt.Printf("for i = %d\n", i)
		switch i {
		case 2, 3:
			fmt.Printf("Break switch, i == %d.\n", i)
			break
		case 5:
			fmt.Printf("Break loop, i == 5.\n")
			break loop
		}
	}
	fmt.Println("End.")
}
