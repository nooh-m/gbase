package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// Start with standard output as the first file
	files := []*os.File{os.Stdout}

	// Create files for each argument, or print an error if file creation fails
	for _, arg := range os.Args[1:] {
		file, err := os.Create(arg)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error creating file: %v\n", err)
			continue
		}
		defer file.Close() // Ensure the file is closed after use
		files = append(files, file)
	}

	// Call tee to handle writing input to all files
	tee(files)
}

func tee(files []*os.File) {
	scanner := bufio.NewScanner(os.Stdin)
	// Read each line from standard input
	for scanner.Scan() {
		line := scanner.Text()
		// Write the line to all output files
		for _, file := range files {
			file.WriteString(line + "\n")
		}
	}
	// Check for errors while reading input
	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Read error: %v\n", err)
		os.Exit(1)
	}
}
