package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {
	// Convert the first argument to an integer (duration in seconds)
	duration, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Fprint(os.Stderr, "Please enter a valid number\n")
		os.Exit(1)
	}

	// Call sleep with the parsed duration
	sleep(duration)
	os.Exit(0)
}

func sleep(timeDuration int) {
	// Pause execution for the given number of seconds
	time.Sleep(time.Duration(timeDuration) * time.Second)
}
