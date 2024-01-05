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

	contents, err := getContents(filename)
	if err != nil {
		fmt.Printf("Error opening file %s: %+v\n", filename, err)
		os.Exit(1)
	}
	prettyPrint(contents)
}

func getContents(filename string) ([]string, error) {
	contents, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return strings.Split(string(contents[:]), "\n"), nil
}

func prettyPrint(lines []string) {
	for i := range arr {
		arr[i] = strings.TrimSpace(arr[i])
		fmt.Printf("%d: %s\n", i, arr[i])
		fmt.Println("---------------------")
	}
}

func getSections(lines []string) {
	fmt.Println("Sections")
}
