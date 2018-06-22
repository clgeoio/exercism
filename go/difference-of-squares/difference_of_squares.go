//Package diffsquares makes rikki happy and calcs difference bettwen sumOfSq and sqOfSum
package diffsquares

//SquareOfSums returns the square of sums (1 + 2 + ...)^2
func SquareOfSums(n int) int {
	tot := 0
	for i := 1; i <= n; i++ {
		tot += i
	}
	return (tot * tot)
}

//SumOfSquares return the sum of squares (1^2 + 2^2 + ...)
func SumOfSquares(n int) int {
	if n == 1 {
		return 1
	}

	return (n*n + SumOfSquares(n-1))
}

func Difference(n int) int {
	return (SquareOfSums(n) - SumOfSquares(n))
}
