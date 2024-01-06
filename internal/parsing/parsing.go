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

func GetSection(lines []string, until string, startIndex int) types.SectionInfo {
	sectionInfo := types.SectionInfo{
		StartLine: startIndex + 1,
	}
	var content []string

	for i := startIndex; i < len(lines); i++ {
		if strings.Contains(strings.ToLower(lines[i]), fmt.Sprintf("section %s", until)) {
			break
		}
		content = append(content, lines[i])
	}

	sectionInfo.Content = StripEmptyLines(content)
	sectionInfo.LineCount = len(sectionInfo.Content)

	return sectionInfo
}

//func GetPreamble(lines []string) types.SectionInfo {
//	var preamble types.SectionInfo
//	var content []string
//
//scan:
//	for _, v := range lines {
//		if strings.Contains(strings.ToLower(v), "section .data") {
//			break scan
//		}
//		content = append(content, v)
//	}
//
//	preamble.StartLine = 1
//	preamble.Content = StripEmptyLines(content)
//	preamble.LineCount = len(preamble.Content)
//	return preamble
//}
