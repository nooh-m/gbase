package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	// Call date function to print current date and time
	date()
	os.Exit(0)
}

func date() {
	// Define the default date format
	const defaultFormat = "Mon Jan 2 03:04:05 PM MST 2006"
	t := time.Now()                      // Get the current time
	fmt.Println(t.Format(defaultFormat)) // Print time in the specified format
}
