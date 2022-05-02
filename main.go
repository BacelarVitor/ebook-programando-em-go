package main

import "os"

func main() {
	// testQuicksort()
	// testStatistics()
	// testStack()
	// testInfiniteLoops()
	// testNamedLoop()
	testConvertions()
}

func readEntries() []string {
	entries := os.Args[1:]
	return entries
}
