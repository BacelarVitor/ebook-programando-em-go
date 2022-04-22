package main

import "os"

func main() {
	// testQuicksort()
	// testStatistics()
	// testStack()
	testInfiniteLoops()
}

func readEntries() []string {
	entries := os.Args[1:]
	return entries
}
