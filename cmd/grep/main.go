package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

func main() {
	// Open the file specified by the first argument
	file, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Fprintln(os.Stderr, "could not open the file:", err)
		os.Exit(1)
	}
	defer file.Close()

	// Compile the regular expression from the second argument
	re, err := regexp.CompilePOSIX(os.Args[2])
	if err != nil {
		fmt.Fprintln(os.Stderr, "bad regex pattern:", err)
		os.Exit(1)
	}

	// Call grep function to process the file with the regex
	grep(file, re)
	os.Exit(0)
}

// grep searches each line of the file for a match to the regex
func grep(file *os.File, regex *regexp.Regexp) {
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		// Print lines that match the regular expression
		if regex.Match([]byte(line)) {
			fmt.Println(line)
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Error reading the file:", err)
	}
}
