package main

import "os"

func main() {
	// testQuicksort()
	testStatistics()
}

func readEntries() []string {
	entries := os.Args[1:]
	return entries
}
