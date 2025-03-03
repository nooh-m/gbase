package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	// Ensure there are exactly 2 arguments (string and suffix)
	if len(os.Args) != 3 {
		fmt.Printf("usage: %s string [suffix]\n", os.Args[0])
		os.Exit(1)
	}

	// Call basename function and print result
	fmt.Println(basename(os.Args[1], os.Args[2]))

	os.Exit(0)
}

func basename(str string, suffix string) string {

	// Trim trailing slashes
	str = strings.TrimRight(str, "/")

	// Extract basename (after last slash)
	if idx := strings.LastIndex(str, "/"); idx != -1 {
		str = str[idx+1:]
	}

	// Remove suffix if it matches
	if strings.HasSuffix(str, suffix) {
		str = strings.TrimSuffix(str, suffix)
	}
	return str
}
