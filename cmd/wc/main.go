package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
)

func main() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Fprint(os.Stderr, "error opening: ", os.Args[1], "\n")
		os.Exit(1)
	}
	defer file.Close()
	wc(file, os.Args[1])

	os.Exit(0)
}

func wc(file *os.File, str string) {
	word := false
	nc, nl, nw := 0, 0, 0

	reader := bufio.NewReader(file)
	for {
		char, _, err := reader.ReadRune() // Reads a full UTF-8 encoded rune
		if err != nil {
			if err != io.EOF {
				fmt.Fprint(os.Stderr, "Error Reading Characters", err)
				os.Exit(1)
			}
			break
		}

		nc++
		if char == '\n' {
			nl++
		}
		if !unicode.IsSpace(char) {
			word = true
		} else if word {
			word = false
			nw++
		}
	}

	fmt.Println("lines:", nl, "words:", nw, "chars:", nc, str)
}
