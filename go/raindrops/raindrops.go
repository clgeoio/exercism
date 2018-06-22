//Package raindrops implements the classic fizzbuzz style code
package raindrops

import (
	"fmt"
	"strings"
)

//Convert turns a number in to raindrop speak!
func Convert(num int) string {
	var sb strings.Builder
	isDiv3 := (num % 3) == 0
	isDiv5 := (num % 5) == 0
	isDiv7 := (num % 7) == 0

	if isDiv3 {
		sb.WriteString("Pling")
	}
	if isDiv5 {
		sb.WriteString("Plang")
	}
	if isDiv7 {
		sb.WriteString("Plong")
	}

	if isDiv3 || isDiv5 || isDiv7 {
		return sb.String()
	}

	return fmt.Sprintf("%d", num)
}
