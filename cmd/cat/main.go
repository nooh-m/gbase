package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	args := os.Args

	// Exit if no files are provided
	if len(args) == 1 {
		fmt.Fprint(os.Stderr, "No file where provided")
		os.Exit(1)
	}

	// Process each file argument
	for _, arg := range args[1:] {
		fp, err := os.Open(arg)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		defer fp.Close()

		cat(fp)
	}
	os.Exit(0)
}

// Prints file contents line by line
func cat(file *os.File) {
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
}
