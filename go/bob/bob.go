/*
Package bob the lackadasical teenager who responds to remarks
Bob answers 'Sure.' if you ask him a question.
He answers 'Whoa, chill out!' if you yell at him.
He answers 'Calm down, I know what I'm doing!' if you yell a question at him.
He says 'Fine. Be that way!' if you address him without actually saying anything.
He answers 'Whatever.' to anything else.
*/
package bob

import "strings"

const Alphas = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

//Hey responds to an input string with a teenage response!
func Hey(remark string) string {
	remarkTrim := strings.TrimSpace(remark)
	remarkUpper := strings.ToUpper(remarkTrim)
	isShouting := (remarkTrim == remarkUpper && strings.ContainsAny(remarkUpper, Alphas))
	isQuestion := strings.HasSuffix(remarkTrim, "?")

	switch {
	case remarkTrim == "":
		return "Fine. Be that way!"
	case isQuestion && isShouting:
		return "Calm down, I know what I'm doing!"
	case isQuestion:
		return "Sure."
	case isShouting:
		return "Whoa, chill out!"
	}

	return "Whatever."
}
