package main

import (
	"fmt"
	"os"
)

func main() {
	echo(os.Args[1:])
	os.Exit(0)
}

func echo(set []string) {
	// Loop through arguments, skipping the first one (the program name)
	for i, arg := range set {
		if i > 0 {
			// Print a space between arguments
			fmt.Print(" ")
		}
		// Print the argument without a newline
		fmt.Print(arg)
	}
}
