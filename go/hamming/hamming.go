//Package hamming allows calculation of hamming distance between two DNA strands
package hamming

import "errors"

//Distance calculates hamming distance between two DNA strands
func Distance(a, b string) (int, error) {
	ham := 0

	if len(a) != len(b) {
		return -1, errors.New("Strand lengths to not match")
	}

	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			ham++
		}
	}

	return ham, nil
}
