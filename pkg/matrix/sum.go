package matrix

import "fmt"

func (m *Matrix) Sum(other *Matrix) *Matrix {
	if m.M != other.M || m.N != other.N {
		panic(fmt.Errorf("can't sum matrices with different dimensionality:\n\n%s\n%s\n", m, other))
	}

	resultData := make([][]float64, m.M)
	for i := 0; i < m.M; i++ {
		resultRow := make([]float64, m.N)
		for j := 0; j < m.N; j++ {
			resultRow[j] = m.Data[i][j] + other.Data[i][j]
		}
		resultData[i] = resultRow
	}

	resultMatrix, err := NewMatrix(resultData)
	if err != nil {
		panic(fmt.Errorf("error creating result matrix: %s", err))
	}

	return resultMatrix
}

func (m *Matrix) Subtract(other *Matrix) *Matrix {
	duplicate, err := NewMatrix(other.Data)
	if err != nil {
		panic(fmt.Errorf("error creating duplicate matrix: %s", err))
	}

	duplicate.Scale(-1)

	resultMatrix := m.Sum(duplicate)

	return resultMatrix
}

// SumRows sums up the values in all rows and returns a 1xn matrix.
// If the matrix m doesn't have rows, method returns a nil pointer
func (m *Matrix) SumRows() *Matrix {
	if m.M <= 0 {
		return nil
	}

	res, _ := NewZeroMatrix(1, m.N)

	for i := range m.Data {
		for j := range m.Data[i] {
			res.Data[0][j] += m.Data[i][j]
		}
	}

	return res
}

// SumColumns sums up the values in all columns and returns an mx1 matrix.
// If the matrix m doesn't have columns, method returns a nil pointer
func (m *Matrix) SumColumns() *Matrix {
	if m.N <= 0 {
		return nil
	}

	res, _ := NewZeroMatrix(m.M, 1)

	for j := range m.Data[0] {
		for i := range m.Data {
			res.Data[i][0] += m.Data[i][j]
		}
	}

	return res
}
