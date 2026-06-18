package functions

import (
	"strings"
)

func AsciiConverter(input string, banner string) string {
	banner = strings.ReplaceAll(banner, "\r", "")
	lines := strings.Split(banner, "\n")

	input = strings.ReplaceAll(input, "\\n", "\n")
	inputLines := strings.Split(input, "\n")

	var output string

	for _, inputLine := range inputLines {
		for j := 0; j < 8; j++ {
			for _, char := range inputLine {
				n := int(char)
				output += lines[((n-32)*9)+j+1]
			}
			output += "\n"
		}
	}
	return output
}
