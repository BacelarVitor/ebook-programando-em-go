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
	// testCreateFiles()
	// testRegex()
	//testFibonacci()
	// testStopwatch()
	testServerTimeStamp()
}

func readEntries() []string {
	entries := os.Args[1:]
	return entries
}
