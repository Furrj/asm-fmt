package parsing

import (
	"github.com/Furrj/asm-fmt/internal/types"
	"os"
	"strings"
)

func GetContents(filename string) ([]string, error) {
	contents, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return strings.Split(string(contents[:]), "\n"), nil
}

func GetPreamble(lines []string) types.SectionInfo {
	var preamble types.SectionInfo
	var content []string

scan:
	for _, v := range lines {
		if strings.Contains(strings.ToLower(v), "section .data") {
			break scan
		}
		content = append(content, v)
	}

	preamble.StartLine = 1
	preamble.Content = StripEmptyLines(content)
	preamble.LineCount = len(preamble.Content)
	return preamble
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
