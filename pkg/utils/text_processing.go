package utils

import (
	"strings"
)

// Capitalize takes a string and returns it with the first character capitalized.
func Capitalize(text string) string {
	if len(text) == 0 {
		return text
	}
	return strings.ToUpper(string(text[0])) + text[1:]
}

// Depluralize takes a string and returns its singular form.
func Depluralize(text string) string {
	if strings.HasSuffix(text, "ies") {
		return text[:len(text)-3] + "y"
	}
	if strings.HasSuffix(text, "s") {
		return text[:len(text)-1]
	}
	return text
}
