//Package luhn makes rikki happy and calculates Luhn algo
package luhn

var nums = map[rune]int{'0': 0, '1': 1, '2': 2, '3': 3, '4': 4, '5': 5, '6': 6, '7': 7, '8': 8, '9': 9}

//Valid determines if a string is a valid number based on the Luhn algo.
func Valid(in string) bool {

	trim := RemoveSpaces(in)
	if len(trim) < 2 {
		return false
	}

	i := 1
	luhn := 0

	for _, r := range Reverse(trim) {
		if val, ok := nums[r]; ok {
			if i%2 == 0 {
				val = val * 2
				if val > 9 {
					val = val - 9
				}
			}
			luhn = luhn + val
			i++
		} else {
			return false
		}
	}

	return luhn%10 == 0
}

//Reverse reverses a string!
func Reverse(input string) string {
	runes := []rune(input)
	var rev []rune

	for i := len(runes) - 1; i > -1; i-- {
		rev = append(rev, runes[i])
	}
	return string(rev)
}

//RemoveSpaces removes all spaces from a string
func RemoveSpaces(input string) string {
	out := ""
	for _, r := range input {
		if r != ' ' {
			out += string(r)
		}
	}
	return out
}
