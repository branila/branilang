package main

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func main() {
	// Checks if the user provided an input file
	if len(os.Args) == 1 {
		log.Fatal("No input file provided")
	}

	sourceName := os.Args[1]

	// Checks if the file is a branilang file
	if !strings.HasSuffix(sourceName, ".br") {
		log.Fatal("The file is not a branilang file")
	}

	// Reads the source
	source, err := os.Open(sourceName)
	if err != nil {
		log.Fatal("An error occurred while reading the file:", err)
	}

	defer source.Close()

	// Creates the output file
	output, err := os.Create(sourceName[:len(sourceName)-3] + ".js")
	if err != nil {
		log.Fatal("An error occurred while creating the output file:", err)
	}

	defer output.Close()

	scanner := bufio.NewScanner(source)

	// Reads the file line by line
	for scanner.Scan() {
		line := scanner.Text()

		// Writes the line to the output file
		output.WriteString(line)
	}
}
