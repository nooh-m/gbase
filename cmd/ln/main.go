package main

import (
	"fmt"
	"os"
)

func main() {
	// Check if at least two arguments are provided
	if len(os.Args) == 2 {
		fmt.Fprint(os.Stderr, "At least two arguments are needed to execute")
		os.Exit(1)
	}

	// Call ln function to create a hard link
	ln(os.Args[1], os.Args[2])
	os.Exit(0)
}

func ln(source string, destination string) {
	// Create a hard link from source to destination
	err := os.Link(source, destination)
	if err != nil {
		// If there’s an error, print it and exit
		fmt.Fprint(os.Stderr, err, "\n")
		os.Exit(1)
	}
}
