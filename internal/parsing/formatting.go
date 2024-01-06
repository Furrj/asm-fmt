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
