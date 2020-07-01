//Package diffsquares provides methods for working with sums of squares and squares of sums.
package diffsquares

//SquareOfSum calculates the square of the sum of every integer from 1 to num.
func SquareOfSum(num int) int {
	var sum int
	for i := 1; i <= num; i++ {
		sum += i
	}
	return sum * sum
}

//SumOfSquares calculates the sum of the squares of every integer from 1 to num.
func SumOfSquares(num int) int {
	var sum int
	for i := 1; i <= num; i++ {
		sum += i * i
	}
	return sum
}

//Difference calculates the difference between SquareOfSum and SumOfSquares.
func Difference(num int) int {
	return SquareOfSum(num) - SumOfSquares(num)
}
