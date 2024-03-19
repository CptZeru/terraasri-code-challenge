package challenges

import (
	"CptZeru/terraasri-code-challenge/utils"
	"testing"
)

func TestCalculateSpiralDiagonalSum(t *testing.T) {
	println("Calculate Sprial Diagonal Sum Test Case")
	nInputs := []int{1, 3, 5, 2, 4, 10}
	nExpects := []int{1, 25, 101, 0, 0, 0}

	for index, n := range nInputs {
		res := CalculateSpiralDiagonalSum(n)
		if res != nExpects[index] {
			utils.ThrowTestingErr(t, nExpects[index], res)
		}
	}
}
