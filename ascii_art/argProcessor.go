package ascii_art_web

import (
	"os"
	"strings"
)

const ART_CHAR_HEIGHT = 8

// ArgProcessor - A function that takes a string as input and
// a template file returns its ascii_art as string.
func ArgProcessor(input, templatePath string) (string, error) {
	output := ""

	input = strings.ReplaceAll(input, `\n`, "\n")
	if strings.Count(input, "\n") == len(input) {
		return input, nil
	}

	data, err := os.ReadFile(templatePath)

	content := strings.ReplaceAll(string(data), "\r", "")
	bannerContent := strings.Split(content[1:], "\n\n")

	artMap := AsciiArtGenerator(bannerContent)
	argWords := strings.Split(input, "\r\n")
	for _, word := range argWords {
		if word == "" {
			output += "\n"
			continue
		}

		for l := 0; l < ART_CHAR_HEIGHT; l++ {
			for _, c := range word {
				output += artMap[c][l]
			}
			output += "\n"
		}
	}
	return output, err
}

func AsciiArtGenerator(bannerContent []string) map[rune][]string {
	characterMatrix := make(map[rune][]string)

	for i, artChar := range bannerContent {
		characterMatrix[rune(32+i)] = strings.Split(artChar, "\n")
	}
	return characterMatrix
}
