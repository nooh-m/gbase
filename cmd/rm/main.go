package main

import (
	"fmt"
	"os"
)

func main() {
	// Check if file arguments are provided
	if len(os.Args) <= 1 {
		fmt.Fprint(os.Stderr, "missing operand:\n")
		os.Exit(1)
	}

	// Call rm to remove the specified files
	rm(os.Args[1:])
	os.Exit(0)
}

func rm(files []string) {
	// Iterate through the provided file names
	for _, file := range files {
		// Try to remove the file
		err := os.Remove(file)
		if err != nil {
			// Print error if file cannot be removed and exit
			fmt.Fprint(os.Stderr, err, "\n")
			os.Exit(1)
		}
	}
}
