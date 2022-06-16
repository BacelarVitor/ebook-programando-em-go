package main

import "os"

func main() {
	// testQuicksort()
	// testStatistics()
	// testStack()
	// testInfiniteLoops()
	// testNamedLoop()
	// testConvertions()
	// testGenerics()
	// testFunctions()
	testCreateFiles()
}

func readEntries() []string {
	entries := os.Args[1:]
	return entries
}
