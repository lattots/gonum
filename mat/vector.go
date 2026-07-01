package mat

import (
	"math"

	"github.com/lattots/gonum/number"
)

func (m *Mat[T]) IsVector() bool {
	return m.M == 1 || m.N == 1
}

// Length calculates the Euclidean norm (length) of a column or row vector
func (m *Mat[T]) Length() float64 {
	if !m.IsVector() {
		panic("matrix math error: cannot calculate length of a non-vector matrix")
	}

	var sum float64
	for _, val := range m.Data {
		v := float64(val)
		sum += v * v
	}
	return math.Sqrt(sum)
}

// Normalize scales a vector matrix to a length of 1.
func Normalize[T number.Num](v *Mat[T]) *Mat[T] {
	length := v.Length()
	if length == 0 {
		panic("matrix math error: cannot normalize a vector of length 0")
	}

	return Map(v, func(val T) T {
		return T(float64(val) / length)
	})
}

// VectorDot computes the vector dot product (returning a scalar value).
func VectorDot[T number.Num](v1, v2 *Mat[T]) T {
	if !v1.IsVector() || !v2.IsVector() {
		panic("matrix math error: inputs must be vectors for a vector dot product")
	}
	if len(v1.Data) != len(v2.Data) {
		panic("matrix math error: vector dimensions must match for a dot product")
	}

	var sum T
	for i := range v1.Data {
		sum += v1.Data[i] * v2.Data[i]
	}
	return sum
}

// CrossProduct calculates the 3D cross product of two 3-element vectors.
func CrossProduct[T number.Num](v1, v2 *Mat[T]) *Mat[T] {
	if len(v1.Data) != 3 || len(v2.Data) != 3 {
		panic("matrix math error: cross product is only defined for 3-dimensional vectors")
	}

	d := make([]T, 3)
	d[0] = v1.Data[1]*v2.Data[2] - v1.Data[2]*v2.Data[1]
	d[1] = v1.Data[2]*v2.Data[0] - v1.Data[0]*v2.Data[2]
	d[2] = v1.Data[0]*v2.Data[1] - v1.Data[1]*v2.Data[0]

	return &Mat[T]{
		M:    v1.M,
		N:    v1.N,
		Data: d,
	}
}

func CosineSimilarity[T number.Num](v1, v2 *Mat[T]) float64 {
	if !v1.IsVector() || !v2.IsVector() {
		panic("matrix math error: inputs must be vector-shaped matrices for cosine similarity")
	}
	if len(v1.Data) != len(v2.Data) {
		panic("matrix math error: vector dimensions must match to calculate cosine similarity")
	}

	len1 := v1.Length()
	len2 := v2.Length()

	if len1 == 0 || len2 == 0 {
		panic("matrix math error: cannot calculate cosine similarity for a vector of length 0")
	}

	dotProduct := float64(VectorDot(v1, v2))

	return dotProduct / (len1 * len2)
}
