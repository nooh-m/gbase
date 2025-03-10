package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main(){
	// Ensure exactly one argument is provided
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "usage: %s string\n", os.Args[0])
		os.Exit(1)
	}

	// Call dirname to print the directory name of the given path
	dirname(os.Args[1])
	os.Exit(0)
}

func dirname(path string){
	// Print the directory part of the provided path
	fmt.Fprint(os.Stdout, filepath.Dir(path), "\n")
}
