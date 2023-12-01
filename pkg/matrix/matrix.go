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
