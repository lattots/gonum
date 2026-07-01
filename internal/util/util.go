package util

import (
	"math"

	"github.com/lattots/gonum/mat"
	"github.com/lattots/gonum/number"
)

func IsClose[T number.Num](num1, num2 T) bool {
	// Tolerance of +- 0.00001% is accepted.
	const tolerance = 0.0000001

	diff := math.Abs(float64(num1 - num2))
	if diff/float64(num1) > tolerance {
		return false
	}
	return true
}

func EqualMatrix[T number.Num](m1, m2 *mat.Mat[T]) bool {
	if m1.M != m2.M || m1.N != m2.N {
		return false
	}

	for i := range m1.Data {
		if !IsClose(m1.Data[i], m2.Data[i]) {
			return false
		}
	}
	return true
}
