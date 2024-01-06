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
	lines = parsing.StripEmptyLines(lines)

	allSections := parsing.GetSections(lines)

	fmt.Println("Preamble:")
	parsing.PrettyPrint(allSections.Preamble.Content)
	fmt.Println("\nData:")
	parsing.PrettyPrint(allSections.Data.Content)
	fmt.Println("\nBss:")
	parsing.PrettyPrint(allSections.Bss.Content)
	fmt.Println("\nText:")
	parsing.PrettyPrint(allSections.Text.Content)
}
