package main

import (
	"fmt"
	"os"
)

// chapter six

func testCreateFiles() {
	tmp := os.TempDir()

	CreateFiles(tmp)
	CreateFiles(tmp, "test1")
	CreateFiles(tmp, "test2", "test3", "test4")

}
func CreateFiles(dirBase string, files ...string) {
	for _, name := range files {
		pathFile := fmt.Sprintf("%s/%s.%s", dirBase, name, "txt")

		file, err := os.Create(pathFile)

		defer file.Close()

		if err != nil {
			fmt.Printf("Error creating the file %s: %v\n", name, err)
			os.Exit(1)
		}

		fmt.Printf("File %s created.\n", file.Name())
	}
}
