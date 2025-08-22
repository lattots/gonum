package matrix

import (
	"fmt"
)

const baseCaseSize = 64

func (m *Matrix) Multiply(other *Matrix) (*Matrix, error) {
	if m.N != other.M {
		return nil, fmt.Errorf("cannot multiply matrices: Number of columns in the first matrix (%d) must be equal to the number of rows in the second matrix (%d)", m.N, other.M)
	}

	mSq, otherSq := square(m, other)
	res := mSq.multiplyStrassenRecursive(otherSq)

	data := make([][]float64, m.M)
	for i := 0; i < m.M; i++ {
		data[i] = res.Data[i][:other.N]
	}
	result, _ := NewMatrix(data)

	return result, nil
}

func (m *Matrix) multiplyStrassenRecursive(other *Matrix) *Matrix {
	if m.M <= baseCaseSize {
		res := m.multiplyStandard(other)
		return res
	}

	n := m.M / 2

	a11, a12, a21, a22, _ := split(m)
	b11, b12, b21, b22, _ := split(other)

	p1 := a11.multiplyStrassenRecursive(b12.Subtract(b22))
	p2 := (a11.Sum(a12)).multiplyStrassenRecursive(b22)
	p3 := (a21.Sum(a22)).multiplyStrassenRecursive(b11)
	p4 := a22.multiplyStrassenRecursive(b21.Subtract(b11))
	p5 := (a11.Sum(a22)).multiplyStrassenRecursive(b11.Sum(b22))
	p6 := (a12.Subtract(a22)).multiplyStrassenRecursive(b21.Sum(b22))
	p7 := (a11.Subtract(a21)).multiplyStrassenRecursive(b11.Sum(b12))

	c11 := p5.Sum(p4).Subtract(p2).Sum(p6)
	c12 := p1.Sum(p2)
	c21 := p3.Sum(p4)
	c22 := p5.Sum(p1).Subtract(p3).Subtract(p7)

	result, _ := combine(c11, c12, c21, c22, n)

	return result
}

func (m *Matrix) multiplyStandard(other *Matrix) *Matrix {
	result, _ := NewZeroMatrix(m.M, other.N)

	for i := 0; i < m.M; i++ {
		for j := 0; j < other.N; j++ {
			for k := 0; k < m.N; k++ {
				result.Data[i][j] += m.Data[i][k] * other.Data[k][j]
			}
		}
	}

	return result
}

func nextPowerOfTwo(n int) int {
	power := 1
	for power < n {
		power *= 2
	}
	return power
}

func square(m1, m2 *Matrix) (*Matrix, *Matrix) {
	largest := max(m1.M, m1.N, m2.M, m2.N)
	size := nextPowerOfTwo(largest)

	newM1, _ := NewZeroMatrix(size, size)
	newM2, _ := NewZeroMatrix(size, size)

	for i := range m1.Data {
		for j := range m1.Data[i] {
			newM1.Data[i][j] = m1.Data[i][j]
		}
	}
	for i := range m2.Data {
		for j := range m2.Data[i] {
			newM2.Data[i][j] = m2.Data[i][j]
		}
	}
	return newM1, newM2
}

func split(m *Matrix) (*Matrix, *Matrix, *Matrix, *Matrix, error) {
	n := m.M / 2
	m11, _ := NewMatrix(slice(m.Data, 0, n, 0, n))
	m12, _ := NewMatrix(slice(m.Data, 0, n, n, m.N))
	m21, _ := NewMatrix(slice(m.Data, n, m.M, 0, n))
	m22, _ := NewMatrix(slice(m.Data, n, m.M, n, m.N))

	return m11, m12, m21, m22, nil
}

func combine(m11, m12, m21, m22 *Matrix, n int) (*Matrix, error) {
	result, _ := NewZeroMatrix(2*n, 2*n)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			result.Data[i][j] = m11.Data[i][j]
			result.Data[i][j+n] = m12.Data[i][j]
			result.Data[i+n][j] = m21.Data[i][j]
			result.Data[i+n][j+n] = m22.Data[i][j]
		}
	}
	return result, nil
}

func slice(matrix [][]float64, rs, re, cs, ce int) [][]float64 {
	height := re - rs
	width := ce - cs

	sliced := make([][]float64, height)
	for i := 0; i < height; i++ {
		sliced[i] = make([]float64, width)
		for j := 0; j < width; j++ {
			sliced[i][j] = matrix[rs+i][cs+j]
		}
	}
	return sliced
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
