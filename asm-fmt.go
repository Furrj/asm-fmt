package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Printf("Please enter a filename\n")
		os.Exit(0)
	}
	filename := os.Args[1]

	lines, err := getContents(filename)
	if err != nil {
		fmt.Printf("Error opening file %s: %+v\n", filename, err)
		os.Exit(1)
	}
	preamble := getPreamble(lines)
	fmt.Println(preamble.LineCount)
	prettyPrint(preamble.Content)
}

type section struct {
	LineCount int
	StartLine int
	Content   []string
}

type sections struct {
	Preamble section
	Data     []string
	Bss      []string
	Text     []string
}

func getContents(filename string) ([]string, error) {
	contents, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return strings.Split(string(contents[:]), "\n"), nil
}

func prettyPrint(lines []string) {
	for i := range lines {
		lines[i] = strings.TrimSpace(lines[i])
		fmt.Printf("%d: %s\n", i, lines[i])
		fmt.Println("---------------------")
	}
}

func getPreamble(lines []string) section {
	var preamble section
	var content []string

scan:
	for i, v := range lines {
		if strings.Contains(strings.ToLower(v), "section .data") {
			preamble.LineCount = i
			break scan
		}
		content = append(content, v)
	}

	preamble.StartLine = 1
	preamble.Content = content
	return preamble
}

func stripEmptyLines(lines []string) []string {
	length := len(lines)

	if lines[length-1] == "" {
		lines = lines[:length]
	}

	return lines
}
