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
	prettyPrint(preamble.Content)
}

type sectionInfo struct {
	LineCount int
	StartLine int
	Content   []string
}

type sections struct {
	Preamble sectionInfo
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

func getPreamble(lines []string) sectionInfo {
	var preamble sectionInfo
	var content []string

scan:
	for _, v := range lines {
		if strings.Contains(strings.ToLower(v), "section .data") {
			break scan
		}
		content = append(content, v)
	}

	preamble.StartLine = 1
	preamble.Content = stripEmptyLines(content)
	preamble.LineCount = len(preamble.Content)
	return preamble
}

func stripEmptyLines(lines []string) []string {
	currLine := len(lines) - 1
	scanning := true

	for scanning {
		if lines[currLine] == "" {
			lines = lines[:currLine]
			currLine--
		} else {
			scanning = false
		}
	}

	return lines
}
