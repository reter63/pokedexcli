package main

import (
	"strings"
)

func cleanInput(text string) []string {
	formattedText := text
	for strings.Contains(formattedText, "  ") {
		formattedText = strings.ReplaceAll(formattedText, "  ", " ")
	}
	formattedText = strings.Trim(formattedText, " ")
	formattedText = strings.ToLower(formattedText)
	if len(formattedText) == 0 {
		var emptySlice []string
		return emptySlice
	}
	return strings.Split(formattedText, " ")
}
