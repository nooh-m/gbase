package main

import (
	"fmt"
	"os"
)

func main() {
	path := os.Args[1]

	touch(path)
}

func touch(path string) {
	_, err := os.Stat(path)
	if !os.IsNotExist(err) {
		fmt.Fprintf(os.Stderr, "The file '%s' already exists.\n", path)
		os.Exit(1)
	}

	// Create the file if it does not exist
	file, err := os.OpenFile(path, os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Cannot create file '%s': %s\n", path, err)
		os.Exit(1)
	}
	file.Close()
}
