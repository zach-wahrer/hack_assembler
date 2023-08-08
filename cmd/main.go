package main

import (
	"flag"
	"fmt"
	"hackassembler"
	"os"
	"path"
	"strings"
)

func main() {
	inputFilename := flag.String("i", "", "input file name")
	flag.Parse()

	if inputFilename == nil || *inputFilename == "" {
		fmt.Println("Filename must be specified: '-i'")
		os.Exit(1)
	}

	file, err := os.Open(*inputFilename)
	if err != nil {
		fmt.Printf("opening file: %s", err)
		os.Exit(1)
	}
	defer func() {
		if err := file.Close(); err != nil {
			fmt.Printf("closing read file: %s", err)
		}
	}()

	assembled, err := hackassembler.Assemble(file)
	if err != nil {
		fmt.Printf("assembling file: %s", err)
		os.Exit(1)
	}

	workDir, workFileName := path.Split(*inputFilename)
	fileNameSplit := strings.Split(workFileName, ".")

	if err := os.WriteFile(path.Join(workDir, fmt.Sprintf("%s.hack", fileNameSplit[0])), assembled, 0600); err != nil {
		fmt.Printf("writing assembled file: %s", err)
		os.Exit(1)
	}

	fmt.Println("Successfully wrote assembly file.")
}
