//Package grains determines powers of two!
package grains

import "errors"

//Square calculates 2^n
func Square(n int) (uint64, error) {

	if n < 1 || n > 64 {
		return 0, errors.New("invalid input n")
	}

	return 1 << (uint64(n) - 1), nil
}

//Total calculates sum of (2 ^ n) for 0 - n
func Total() uint64 {
	return (1 << 64) - 1
}
