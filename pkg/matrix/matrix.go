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

	if m.M <= 3 {
		for i := range m.Data {
			matString += stringRow(m.Data[i]) + "\n"
		}
	} else {
		matString += stringRow(m.Data[0]) + "\n" + stringRow(m.Data[1]) + "\n...\n" + stringRow(m.Data[len(m.Data)-1]) + "\n"
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
		row = fStr(r[0], 2) + " " + fStr(r[1], 2)

		row += " ... " + fStr(r[len(r)-1], 2)
	}

	return row
}

func fStr(f float64, prec int) string {
	return strconv.FormatFloat(f, 'f', prec, 64)
}

func (m *Matrix) Transpose() *Matrix {
	tData := make([][]float64, m.N)
	for i := range tData {
		tData[i] = make([]float64, m.M)
	}

	for i := range m.M {
		for j := range m.N {
			tData[j][i] = m.Data[i][j]
		}
	}

	t, _ := NewMatrix(tData) // The new matrice's data is known to be correct, so error can be ignored

	return t
}

func (m *Matrix) Scale(scalar float64) {
	for i := range m.M {
		for j := range m.N {
			m.Data[i][j] = m.Data[i][j] * scalar
		}
	}
}

func (m *Matrix) Map(fn func(float64) float64) (*Matrix, error) {
	res, err := NewZeroMatrix(m.M, m.N)
	if err != nil {
		return nil, fmt.Errorf("error creating new zero matrix: %w", err)
	}
	for i := range m.M {
		for j := range m.N {
			res.Data[i][j] = fn(m.Data[i][j])
		}
	}
	return res, nil
}
