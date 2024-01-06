package parsing

import (
	"fmt"
	"strings"
)

func PrettyPrint(lines []string) {
	for i := range lines {
		lines[i] = strings.TrimSpace(lines[i])
		fmt.Printf("%d: %s\n", i, lines[i])
		fmt.Println("---------------------")
	}
}

func StripEmptyLines(lines []string) []string {
	currIndex := len(lines) - 1
	scanning := true

	for scanning {
		if lines[currIndex] == "" && currIndex != 0 {
			lines = lines[:currIndex]
			currIndex--
		} else {
			scanning = false
		}
	}

	return lines
}
