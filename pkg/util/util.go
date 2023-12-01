package util

import "math"

func equalFloat(float1 float64, float2 float64) bool {
	// Tolerance of +- 0.00001% is accepted.
	tolerance := 0.0000001

	diff := math.Abs(float1 - float2)
	if diff/float1 > tolerance {
		return false
	}
	return true
}
