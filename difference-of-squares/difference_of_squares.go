//Package diffsquares provides methods for working with sums of squares and squares of sums.
package diffsquares

//SquareOfSum calculates the square of the sum of every integer from 1 to num.
func SquareOfSum(num int) int {
	sum := (num * (num + 1)) / 2
	return sum * sum
}

//SumOfSquares calculates the sum of the squares of every integer from 1 to num.
func SumOfSquares(num int) int {
	sum := num * (num + 1) * (2*num + 1) / 6
	return sum
}

//Difference calculates the difference between SquareOfSum and SumOfSquares.
func Difference(num int) int {
	return SquareOfSum(num) - SumOfSquares(num)
}
