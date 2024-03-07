package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {

	characters := flag.Bool("c", false, "character to count")
	words := flag.Bool("w", false, "word to count")
	lines := flag.Bool("l", false, "line to count")

	file := os.Args[2]

	flag.Parse()

	data, err := os.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}

	files := string(data)
	count := 1

	if *characters {
		count = len(data)
	}

	if *lines {
		lines := strings.Split(files, "\n")
		count = len(lines) - 1
	}

	if *words {
		words := strings.Split(files, " ")
		count = len(words)
	}

	fmt.Printf("%8d %s\n", count, file)

}
