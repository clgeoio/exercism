//Package reverse makes rikki happy and implments a string reverse
package reverse

//String reverses a string!
func String(input string) string {
	runes := []rune(input)
	var rev []rune

	for i := len(runes) - 1; i > -1; i-- {
		rev = append(rev, runes[i])
	}
	return string(rev)
}
