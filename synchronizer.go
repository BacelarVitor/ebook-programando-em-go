package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func testSynchronizer() {
	start := time.Now()
	rand.Seed(start.UnixNano())

	var control sync.WaitGroup

	for i := 0; i < 5; i++ {
		control.Add(1)
		go executeWaitGroup(&control)
	}

	control.Wait()

	fmt.Printf("Finishing in %s.\n", time.Since(start))
}

func executeWaitGroup(control *sync.WaitGroup) {
	defer control.Done()

	duration := time.Duration(1+rand.Intn(5)) * time.Second
	fmt.Printf("Sleeping for %s...\n", duration)
	time.Sleep(duration)
}
