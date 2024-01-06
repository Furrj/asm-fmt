package parsing

import (
	"fmt"
	"github.com/Furrj/asm-fmt/internal/types"
	"os"
	"strings"
)

func GetAllFileContents(filename string) ([]string, error) {
	contents, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return strings.Split(string(contents[:]), "\n"), nil
}

func GetSection(lines []string, until string, startIndex int) (types.SectionInfo, int) {
	sectionInfo := types.SectionInfo{
		StartLine: startIndex + 1,
	}
	var content []string
	var index int
	until = fmt.Sprintf("section %s", until)

	if strings.ToLower(lines[startIndex]) == strings.ToLower(until) {
		sectionInfo.Content = nil
		sectionInfo.LineCount = 0
		index = startIndex
	} else {
		for index = startIndex; index < len(lines); index++ {
			if strings.Contains(strings.ToLower(lines[index]), until) {
				break
			}
			content = append(content, lines[index])
		}

		sectionInfo.Content = StripEmptyLines(content)
		sectionInfo.LineCount = len(sectionInfo.Content)
	}

	return sectionInfo, index
}

func GetSections(lines []string) types.Sections {
	var allSections types.Sections
	sectionLabels := []string{".data", ".bss", ".text", "EOF"}
	sections := []*types.SectionInfo{
		&allSections.Preamble,
		&allSections.Data,
		&allSections.Bss,
		&allSections.Text,
	}
	currIndex := 0

	for i := range sectionLabels {
		*(sections[i]), currIndex = GetSection(lines, sectionLabels[i], currIndex)
	}

	return allSections
}
