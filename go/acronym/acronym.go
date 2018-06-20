// Package acronym provides a function to convert a string in to an acronym
package acronym

import (
	"strings"
	"unicode"
)

// Abbreviate turns an input string in to an abbreviation
func Abbreviate(s string) string {
	noHyphen := strings.Replace(s, "-", " ", -1) //replace all hyphens with spaces
	split := strings.Split(noHyphen, " ")
	var b strings.Builder

	for _, element := range split {
		if element != "" && 'A' <= element[0] && element[0] <= 'z' {
			b.WriteRune(unicode.ToUpper(rune(element[0])))
		}
	}

	return b.String()
}
