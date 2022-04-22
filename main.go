package main

import "os"

func main() {
	testQuicksort()
}

func readEntries() []string {
	entries := os.Args[1:]
	return entries
}
