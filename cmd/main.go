package main

import (
	"fmt"
	"github.com/Furrj/asm-fmt/internal/parsing"
	"os"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Printf("Please enter a filename\n")
		os.Exit(0)
	}
	filename := os.Args[1]

	lines, err := parsing.GetAllFileContents(filename)
	if err != nil {
		fmt.Printf("Error opening file %s: %+v\n", filename, err)
		os.Exit(1)
	}

	preamble := parsing.GetSection(lines, ".data", 0)
	parsing.PrettyPrint(preamble.Content)
}
