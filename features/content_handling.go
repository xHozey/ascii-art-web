package web

import (
	"strings"
)

var Result string

// Convert content array to a character matrix mapping ASCII characters to their line representations
func ConvertToCharacterMatrix(content []string) map[rune][]string {
	characterMatrix := map[rune][]string{}
	for i, val := range content {
		characterMatrix[rune(32+i)] = strings.Split(val, "\n")
	}
	return characterMatrix
}

// Render the ASCII art based on the character matrix and the input lines
func DrawASCIIArt(characterMatrix map[rune][]string, input string) string {
	Result := ""
	input = strings.ReplaceAll(input, `\n`, "\n")
	if strings.Count(input, "\n") == len(input) {
		return input
	}

	splittedInput := strings.Split(input, "\r\n")

	for _, val := range splittedInput {
		for j := 0; j < 8; j++ {
			for _, k := range val {
				Result += characterMatrix[k][j]
			}
			Result += "\n"
		}
	}
	return Result
}
