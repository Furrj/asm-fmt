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
