package mat

import (
	"fmt"

	"github.com/lattots/gonum/number"
)

// Sum adds m2 to m1 element-wise. Panics if dimensions mismatch.
func Sum[T number.Num](m1, m2 *Mat[T]) *Mat[T] {
	if m1.M != m2.M || m1.N != m2.N {
		panic(fmt.Sprintf("matrix math error: cannot sum matrices with different dimensions (%dx%d and %dx%d)", m1.M, m1.N, m2.M, m2.N))
	}

	data := make([]T, len(m1.Data))
	for i := range m1.Data {
		data[i] = m1.Data[i] + m2.Data[i]
	}

	return &Mat[T]{
		M:    m1.M,
		N:    m1.N,
		Data: data,
	}
}

// Subtract subtracts m2 from m1 element-wise. Panics if dimensions mismatch.
func Subtract[T number.Num](m1, m2 *Mat[T]) *Mat[T] {
	if m1.M != m2.M || m1.N != m2.N {
		panic(fmt.Sprintf("matrix math error: cannot subtract matrices with different dimensions (%dx%d and %dx%d)", m1.M, m1.N, m2.M, m2.N))
	}

	data := make([]T, len(m1.Data))
	for i := range m1.Data {
		data[i] = m1.Data[i] - m2.Data[i]
	}

	return &Mat[T]{
		M:    m1.M,
		N:    m1.N,
		Data: data,
	}
}

// SumRows collapses all rows into a single 1xN row vector.
func SumRows[T number.Num](m *Mat[T]) *Mat[T] {
	if m.M <= 0 {
		panic("matrix math error: matrix must have at least one row to sum")
	}

	data := make([]T, m.N)
	for r := 0; r < m.M; r++ {
		for c := 0; c < m.N; c++ {
			data[c] += m.Data[r*m.N+c]
		}
	}

	return &Mat[T]{
		M:    1,
		N:    m.N,
		Data: data,
	}
}

// SumColumns collapses all columns into a single Mx1 column vector.
func SumColumns[T number.Num](m *Mat[T]) *Mat[T] {
	if m.N <= 0 {
		panic("matrix math error: matrix must have at least one column to sum")
	}

	data := make([]T, m.M)
	for r := 0; r < m.M; r++ {
		for c := 0; c < m.N; c++ {
			data[r] += m.Data[r*m.N+c]
		}
	}

	return &Mat[T]{
		M:    m.M,
		N:    1,
		Data: data,
	}
}
