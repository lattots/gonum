package matrix

import (
	"fmt"
	"strconv"
)

type Matrix struct {
	M    int
	N    int
	Data [][]float64
}

func NewMatrix(matrix [][]float64) (*Matrix, error) {

	if matrix == nil || len(matrix) == 0 || len(matrix[0]) == 0 {
		return nil, fmt.Errorf("CANNOT INITIALIZE MATRIX WITH NIL ARRAY")
	}

	if !isValid(matrix) {
		return nil, fmt.Errorf("INVALID MATRIX")
	}

	m := Matrix{
		M:    len(matrix),
		N:    len(matrix[0]),
		Data: matrix,
	}

	return &m, nil
}

func NewZeroMatrix(m, n int) (*Matrix, error) {
	if m <= 0 || n <= 0 {
		return nil, fmt.Errorf("dimensions of matrices must be above zero")
	}
	data := make([][]float64, m)
	for i := range data {
		rowData := make([]float64, n)
		for i := range rowData {
			rowData[i] = 0
		}
		data[i] = rowData
	}
	mat, err := NewMatrix(data)
	if err != nil {
		return nil, err
	}
	return mat, nil
}

func isValid(m [][]float64) bool {

	n := len(m[0])
	for i := range m {
		if len(m[i]) != n {
			return false
		}
	}
	return true
}

func (m *Matrix) String() string {

	matString := fmt.Sprintf("%d x %d\n", m.M, m.N)

	for _, row := range m.Data {
		matString += stringRow(row) + "\n"
	}

	return matString
}

func stringRow(r []float64) string {

	var row string

	if len(r) <= 3 {
		row = fStr(r[0], 2)

		r = r[1:]

		for i := range r {
			row += " " + fStr(r[i], 2)
		}
	} else {
		row = fStr(r[0], 2)

		r = r[1:]

		for i := 0; i < 3; i++ {
			row += " " + fStr(r[i], 2)
		}

		row += "..." + fStr(r[len(r)-1], 2)
	}

	return row
}

func fStr(f float64, prec int) string {
	return strconv.FormatFloat(f, 'f', prec, 64)
}

func (m *Matrix) Sum(other *Matrix) (*Matrix, error) {

	if m.M != other.M || m.N != other.N {
		return nil, fmt.Errorf("CANNOT SUM MATRICES WITH DIFFERENT DIMENSIONALITY")
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
		return nil, err
	}

	return resultMatrix, nil
}

func (m *Matrix) Subtract(other *Matrix) (*Matrix, error) {
	duplicate, err := NewMatrix(other.Data)
	if err != nil {
		return nil, err
	}

	duplicate.Scale(-1)

	resultMatrix, err := m.Sum(duplicate)
	if err != nil {
		return nil, err
	}

	return resultMatrix, nil
}

func (m *Matrix) Scale(scalar float64) {
	for i := 0; i < m.M; i++ {
		for j := 0; j < m.N; j++ {
			m.Data[i][j] = m.Data[i][j] * scalar
		}
	}
}

func (m *Matrix) Transpose() *Matrix {
	tData := make([][]float64, m.N)
	for i := range tData {
		tData[i] = make([]float64, m.M)
	}

	for i := 0; i < m.M; i++ {
		for j := 0; j < m.N; j++ {
			tData[j][i] = m.Data[i][j]
		}
	}

	t, _ := NewMatrix(tData) // The new matrice's data is known to be correct, so error can be ignored

	return t
}

func (m *Matrix) Multiply(other *Matrix) (*Matrix, error) {
	if m.N != other.M {
		return nil, fmt.Errorf("cannot multiply matrices: Number of columns in the first matrix (%d) must be equal to the number of rows in the second matrix (%d)", m.N, other.M)
	} else if m.M == 0 {
		return nil, fmt.Errorf("cannot multiply with nil matrix")
	}

	resultData := make([][]float64, m.M)
	for i := 0; i < m.M; i++ {
		resultData[i] = make([]float64, other.N)
	}

	for i := 0; i < m.M; i++ {
		for j := 0; j < other.N; j++ {
			var r float64
			for k := 0; k < m.N; k++ {
				r += m.Data[i][k] * other.Data[k][j]
			}
			resultData[i][j] = r
		}
	}

	resultMatrix, err := NewMatrix(resultData)
	if err != nil {
		return nil, err
	}

	return resultMatrix, nil
}

func (m *Matrix) MultiplyElements(other *Matrix) (*Matrix, error) {
	if m.M != other.M || m.N != other.N {
		return nil, fmt.Errorf("cannot multiply matrices: matrices must have the same shape")
	} else if m.M == 0 {
		return nil, fmt.Errorf("cannot multiply with nil matrix")
	}

	resultData := make([][]float64, m.M)
	for i := range m.Data {
		resultData[i] = make([]float64, m.N)
		for j := range m.Data[i] {
			resultData[i][j] = m.Data[i][j] * other.Data[i][j]
		}
	}

	resultMatrix, err := NewMatrix(resultData)
	if err != nil {
		return nil, err
	}

	return resultMatrix, nil
}
