package challenges

// CalculateSpiralDiagonalSum calculates diagonal sum of generated spiral matrix.
//
// n should be odd int due the complete spiral matrix only forms if n is odd int.
// if n is even, it will return 0.
//
// the logic is top-right corner is equal to n^2 while top-left, bottom-left,
// bottom-right respectively is equal to (n^2) - (n-1), (n^2) - 2(n-1), (n^2) - 3(n-1)
// sum all corner: n^2 + (n^2) - (n-1) + (n^2) - 2(n-1) + (n^2) - 3(n-1),
// becomes: 4(n^2) - 6(n-1).
//
// the above formula only sum the outer corner of spiral matrix, the complete
// formula to sum all corner of the sprial matrix / diagonal sum are below:
//
// f(n) = 4(n^2) - 6(n-1) + f(n-2) ->
// f(n) = 4 * n * n - 6 * n + 6 + f(n-2)
func CalculateSpiralDiagonalSum(n int) int {
	if n == 1 {
		return n
	}
	if n%2 == 0 {
		return 0
	}
	return (4*n*n - 6*n + 6 + CalculateSpiralDiagonalSum(n-2))
}
