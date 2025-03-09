package main

import (
	"fmt"
	"os"
)

func main() {
	// Print the current working directory to stdout
	fmt.Fprint(os.Stdout, pwd(), "\n")
	os.Exit(0)
}

func pwd() string {
	// Get the current working directory
	cwd, err := os.Getwd()
	if err != nil {
		// If there's an error, print it and exit
		fmt.Fprintln(os.Stderr, "Could not get the working directory:", err)
		os.Exit(1)
	}
	return cwd
}
