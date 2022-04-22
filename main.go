package main

import "os"

func main() {
	// testQuicksort()
	// testStatistics()
	// testStack()
}

func readEntries() []string {
	entries := os.Args[1:]
	return entries
}
