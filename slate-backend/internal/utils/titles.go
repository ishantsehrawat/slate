package utils

import (
	"strings"
)

func GenerateHeadingFromContent(content string) string {
	lines := strings.Split(content, "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line != "" {
			return ToTitleCase(line)
		}
	}
	return "Untitled Journal"
}

func ToTitleCase(text string) string {
	words := strings.Fields(text)
	for i, word := range words {
		words[i] = strings.Title(strings.ToLower(word))
	}
	return strings.Join(words, " ")
}
