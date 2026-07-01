package mat

import (
	"fmt"
	"strings"

	"github.com/lattots/gonum/number"
)

type Mat[T number.Num] struct {
	M    int
	N    int
	Data []T
}

func New[T number.Num](data [][]T) (*Mat[T], error) {
	if len(data) == 0 || len(data[0]) == 0 {
		return nil, fmt.Errorf("can't initialize a matrix with no data")
	}

	if !isComplete(data) {
		return nil, fmt.Errorf("data must have equal number of elements in each row")
	}

	flat := flatten(data)
	if flat == nil {
		return nil, fmt.Errorf("failed to flatten the original data")
	}

	m := Mat[T]{
		M:    len(data),
		N:    len(data[0]),
		Data: flatten(data),
	}

	return &m, nil
}

func Zeros[T number.Num](m, n int) (*Mat[T], error) {
	if m <= 0 || n <= 0 {
		return nil, fmt.Errorf("dimensions of matrices must be above zero")
	}

	flatData := make([]T, m*n)

	return &Mat[T]{
		M:    m,
		N:    n,
		Data: flatData,
	}, nil
}

func Transpose[T number.Num](m *Mat[T]) *Mat[T] {
	newData := make([]T, len(m.Data))

	// Map old indices to transposed indices in the new matrix
	for r := range m.M {
		for c := range m.N {
			oldIdx := r*m.N + c
			newIdx := c*m.M + r
			newData[newIdx] = m.Data[oldIdx]
		}
	}

	return &Mat[T]{
		M:    m.N,
		N:    m.M,
		Data: newData,
	}
}

func T[T number.Num](m *Mat[T]) *Mat[T] {
	return Transpose(m)
}

func (m *Mat[T]) String() string {
	var sb strings.Builder

	sb.WriteString(fmt.Sprintf("%d x %d\n", m.M, m.N))

	// Indices of the printed rows
	var rowsToPrint []int
	if m.M <= 3 {
		for i := 0; i < m.M; i++ {
			rowsToPrint = append(rowsToPrint, i)
		}
	} else {
		// -1 represents a marker
		rowsToPrint = []int{0, 1, -1, m.M - 1}
	}

	for _, r := range rowsToPrint {
		// Handle marker
		if r == -1 {
			sb.WriteString("...\n")
			continue
		}

		var rowElements []string

		if m.N <= 3 {
			for c := 0; c < m.N; c++ {
				idx := r*m.N + c
				rowElements = append(rowElements, formatElement(m.Data[idx]))
			}
			sb.WriteString(strings.Join(rowElements, " ") + "\n")
		} else {
			idx0 := r*m.N + 0
			idx1 := r*m.N + 1
			idxLast := r*m.N + (m.N - 1)

			rowElements = append(
				rowElements,
				formatElement(m.Data[idx0]),
				formatElement(m.Data[idx1]),
				"...",
				formatElement(m.Data[idxLast]),
			)
			sb.WriteString(strings.Join(rowElements, " ") + "\n")
		}
	}

	return sb.String()
}

func Scale[T number.Num](m *Mat[T], scalar T) *Mat[T] {
	data := make([]T, len(m.Data))
	for i := range m.Data {
		data[i] = m.Data[i] * scalar
	}
	return &Mat[T]{
		M:    m.M,
		N:    m.N,
		Data: data,
	}
}

func Add[T number.Num](m *Mat[T], scalar T) *Mat[T] {
	data := make([]T, len(m.Data))
	for i := range m.Data {
		data[i] = m.Data[i] + scalar
	}
	return &Mat[T]{
		M:    m.M,
		N:    m.N,
		Data: data,
	}
}

func Map[T number.Num](m *Mat[T], fn func(T) T) *Mat[T] {
	data := make([]T, len(m.Data))
	for i := range m.Data {
		data[i] = fn(m.Data[i])
	}

	return &Mat[T]{
		M:    m.M,
		N:    m.N,
		Data: data,
	}
}

// isComplete checks if all rows in the matrix have the same number of elements
func isComplete[T number.Num](data [][]T) bool {
	n := len(data[0])
	for i := range data {
		if len(data[i]) != n {
			return false
		}
	}
	return true
}

func flatten[T number.Num](data [][]T) []T {
	if len(data) == 0 {
		return nil
	}

	totalElements := 0
	for i := 0; i < len(data); i++ {
		totalElements += len(data[i])
	}

	flat := make([]T, totalElements)

	position := 0
	for i := 0; i < len(data); i++ {
		copy(flat[position:], data[i])
		position += len(data[i])
	}

	return flat
}

func formatElement[T number.Num](val T) string {
	switch v := any(val).(type) {
	case float32:
		return fmt.Sprintf("%.2f", v)
	case float64:
		return fmt.Sprintf("%.2f", v)
	default:
		return fmt.Sprintf("%v", v)
	}
}
