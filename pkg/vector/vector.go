package vector

import (
	"fmt"
	"math"
	"strconv"
)

type Vector struct {
	Vec []float64
	Dim int
	Len float64
}

// NewVector creates a new instance of Vector.
func NewVector(vector []float64) (*Vector, error) {
	if vector == nil {
		return nil, fmt.Errorf("CANNOT INITIALIZE VECTOR WITH NIL ARRAY")
	}
	v := Vector{
		Vec: vector,
		Dim: len(vector),
	}
	v.calculateLength()

	return &v, nil
}

// String returns the vector object in an easy-to-read manner.
func (v *Vector) String() string {
	var vecString string
	if v.Dim > 3 {
		vecString = fmt.Sprintf("[%s %s %s ... %s]", fStr(v.Vec[0], 1), fStr(v.Vec[1], 1), fStr(v.Vec[2], 1), fStr(v.Vec[v.Dim-1], 1))
	} else {
		vecString += "[" + fStr(v.Vec[0], 1)
		for i := range v.Vec[1:] {
			vecString += " " + fStr(v.Vec[i+1], 1)
		}
		vecString += "]"
	}
	return fmt.Sprintf("Dim: %d\nLen: %s\nVec: %s", v.Dim, fStr(v.Len, 2), vecString)
}

func fStr(f float64, prec int) string {
	return strconv.FormatFloat(f, 'f', prec, 64)
}

// calculateLength calculates the length of a vector.
func (v *Vector) calculateLength() {
	var sum float64
	for _, c := range v.Vec {
		sum += math.Pow(c, 2)
	}
	v.Len = math.Sqrt(sum)
}

// Sum calculates the sum for vectors.
func (v *Vector) Sum(other *Vector) (*Vector, error) {
	sum := make([]float64, 0)

	for i := 0; i < v.Dim && i < other.Dim; i++ {
		sum = append(sum, v.Vec[i]+other.Vec[i])
	}

	sumVec, err := NewVector(sum)
	if err != nil {
		return nil, err
	}

	return sumVec, nil
}

// Subtract subtracts the other vector from vector.
func (v *Vector) Subtract(other *Vector) (*Vector, error) {
	duplicate, err := NewVector(other.Vec)
	if err != nil {
		return nil, err
	}
	duplicate.Scale(-1)
	subVec, err := v.Sum(duplicate)
	if err != nil {
		return nil, err
	}
	return subVec, nil
}

// Scale scales vector with the specified scalar.
func (v *Vector) Scale(scalar float64) {
	for i := range v.Vec {
		v.Vec[i] = v.Vec[i] * scalar
	}
	v.calculateLength()
}

// CalculateDotProduct calculates dot product for vectors.
func (v *Vector) CalculateDotProduct(other *Vector) (float64, error) {
	if v.Dim != other.Dim {
		return 0, fmt.Errorf("CANNOT CALCULATE DOT PRODUCT FOR VECTORS WITH DIFFERENT DIMENSIONALITY")
	}

	var sum float64
	for i := range v.Vec {
		sum += v.Vec[i] * other.Vec[i]
	}

	return sum, nil
}

// CalculateCrossProduct calculates cross product for vectors.
func (v *Vector) CalculateCrossProduct(other *Vector) (*Vector, error) {
	if v.Dim != 3 || other.Dim != 3 {
		return nil, fmt.Errorf("CANNOT CALCULATE CROSS PRODUCT FOR VECTORS WITH OTHER THAN 3 DIMENSIONS")
	}

	cross := make([]float64, 3)
	cross[0] = v.Vec[1]*other.Vec[2] - v.Vec[2]*other.Vec[1]
	cross[1] = v.Vec[2]*other.Vec[0] - v.Vec[0]*other.Vec[2]
	cross[2] = v.Vec[0]*other.Vec[1] - v.Vec[1]*other.Vec[0]

	result, err := NewVector(cross)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// NormalizeVector scales vector a length of 1.
func (v *Vector) NormalizeVector() {
	v.Scale(1 / v.Len)
	v.calculateLength()
}

// CosineSimilarity calculates the cosine similarity of two vectors.
func (v *Vector) CosineSimilarity(other *Vector) (float64, error) {
	if v.Dim != other.Dim {
		return 0, fmt.Errorf("CANNOT CALCULATE SIMILARITY FOR VECTORS WITH DIFFERENT DIMENSIONALITY")
	}

	var product float64
	for i, c := range v.Vec {
		product += c * other.Vec[i]
	}

	similarity := product / (v.Len * other.Len)

	// Function returns the similarity score of the vectors.
	return similarity, nil
}
