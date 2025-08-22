package matrix

import (
	"fmt"
	"reflect"
	"testing"
	"time"

	"github.com/lattots/gonum/pkg/util"
)

func equalMatrix(m1, m2 *Matrix) bool {
	if m1.M != m2.M || m1.N != m2.N {
		return false
	}

	for i := range m1.Data {
		for j := range m1.Data[i] {
			if !util.EqualFloat(m2.Data[i][j], m1.Data[i][j]) {
				return false
			}
		}
	}
	return true
}

func TestNewMatrix(t *testing.T) {
	start := time.Now()

	// Test case 1: Valid matrix
	validMatrix := [][]float64{
		{1.1, 2.2, 3.3},
		{4.4, 5.5, 6.6},
	}
	expectedM := 2
	expectedN := 3

	matrix, err := NewMatrix(validMatrix)
	if err != nil {
		t.Errorf("Error creating matrix: %v", err)
	}

	if matrix.M != expectedM || matrix.N != expectedN || !reflect.DeepEqual(matrix.Data, validMatrix) {
		t.Errorf("Expected: (%d, %d, %v), Got: (%d, %d, %v)", expectedM, expectedN, validMatrix, matrix.M, matrix.N, matrix.Data)
	}

	// Test case 2: Invalid matrix (nil)
	var invalidMatrix [][]float64
	_, err = NewMatrix(invalidMatrix)
	if err == nil {
		t.Error("Expected error for nil matrix, but got nil")
	}

	// Test case 3: Invalid matrix (different row lengths)
	invalidMatrix2 := [][]float64{
		{1.1, 2.2, 3.3},
		{4.4, 5.5},
	}
	_, err = NewMatrix(invalidMatrix2)
	if err == nil {
		t.Error("Expected error for invalid matrix, but got nil")
	}

	fmt.Printf("Runtime: %v\n", time.Since(start))
}

func TestNewZeroMatrix(t *testing.T) {
	start := time.Now()

	const (
		m = 3
		n = 5
	)

	mat, err := NewZeroMatrix(m, n)
	if err != nil {
		t.Errorf("Error creating matrix: %v", err)
	}

	if mat.M != m || mat.N != n {
		t.Error("Created matrices dimensions are wrong")
	}

	for _, row := range mat.Data {
		for _, val := range row {
			if val != 0 {
				t.Errorf("Found non-zero value in matrix: %f", val)
			}
		}
	}

	fmt.Printf("Runtime: %v\n", time.Since(start))
}

func TestMatrixString(t *testing.T) {
	smallMatrix, err := NewMatrix([][]float64{
		{1.1, 2.2, 3.3},
		{4.4, 5.5, 6.6},
		{7.7, 8.8, 9.9},
	})
	if err != nil {
		t.Errorf("Error creating matrix: %v", err)
	}

	expectedString := "3 x 3\n1.10 2.20 3.30\n4.40 5.50 6.60\n7.70 8.80 9.90\n"

	resultString := smallMatrix.String()
	if resultString != expectedString {
		t.Errorf("Expected:\n%s\nGot:\n%s", expectedString, resultString)
	}

	largeMatrix, err := NewMatrix([][]float64{
		{1, 2, 3, 4, 5, 6, 7},
		{2, 2, 3, 4, 5, 6, 7},
		{3, 2, 3, 4, 5, 6, 7},
		{4, 2, 3, 4, 5, 6, 7},
	})

	expectedString = "4 x 7\n1.00 2.00 ... 7.00\n2.00 2.00 ... 7.00\n...\n4.00 2.00 ... 7.00\n"

	resultString = largeMatrix.String()
	if resultString != expectedString {
		t.Errorf("Expected:\n%s\nGot:\n%s", expectedString, resultString)
	}
}

func TestMatrixSum(t *testing.T) {
	start := time.Now()

	// Test case 1: Valid matrices with the same dimensions
	matrix1 := [][]float64{
		{1.1, 2.2, 3.3},
		{4.4, 5.5, 6.6},
	}
	matrix2 := [][]float64{
		{0.1, 1.2, 2.3},
		{3.4, 4.5, 5.6},
	}
	expectedSum := [][]float64{
		{1.2, 3.4, 5.6},
		{7.8, 10.0, 12.2},
	}

	m1, err := NewMatrix(matrix1)
	if err != nil {
		t.Errorf("Error creating matrix1: %v", err)
	}

	m2, err := NewMatrix(matrix2)
	if err != nil {
		t.Errorf("Error creating matrix2: %v", err)
	}

	result := m1.Sum(m2)
	if err != nil {
		t.Errorf("Error during matrix summation: %v", err)
	}

	// Compare matrices element-wise with tolerance
	for i := range expectedSum {
		for j := range expectedSum[i] {
			if !util.EqualFloat(result.Data[i][j], expectedSum[i][j]) {
				t.Errorf("Expected sum: %v, Got: %v", expectedSum, result.Data)
				break
			}
		}
	}

	fmt.Printf("Runtime: %v\n", time.Since(start))
}

func TestMatrixSubtract(t *testing.T) {
	start := time.Now()

	// Test case 1: Valid matrices with the same dimensions
	matrix1 := [][]float64{
		{1.1, 2.2, 3.3},
		{4.4, 5.5, 6.6},
	}
	matrix2 := [][]float64{
		{0.1, 1.2, 2.3},
		{3.4, 4.5, 5.6},
	}
	expectedResult := [][]float64{
		{1.0, 1.0, 1.0},
		{1.0, 1.0, 1.0},
	}

	m1, err := NewMatrix(matrix1)
	if err != nil {
		t.Errorf("Error creating matrix1: %v", err)
	}

	m2, err := NewMatrix(matrix2)
	if err != nil {
		t.Errorf("Error creating matrix2: %v", err)
	}

	result := m1.Subtract(m2)

	// Compare matrices element-wise with tolerance
	for i := range expectedResult {
		for j := range expectedResult[i] {
			if !util.EqualFloat(result.Data[i][j], expectedResult[i][j]) {
				t.Errorf("Expected result: %v, Got: %v", expectedResult, result.Data)
				break
			}
		}
	}

	fmt.Printf("Runtime: %v\n", time.Since(start))
}

func TestMatrixScale(t *testing.T) {
	start := time.Now()

	// Test case 1: Positive scalar
	matrix1 := [][]float64{
		{1.1, 2.2, 3.3},
		{4.4, 5.5, 6.6},
	}
	scalar1 := 2.0
	expectedResult1 := [][]float64{
		{2.2, 4.4, 6.6},
		{8.8, 11.0, 13.2},
	}

	m1, err := NewMatrix(matrix1)
	if err != nil {
		t.Errorf("Error creating matrix1: %v", err)
	}

	m1.Scale(scalar1)

	// Compare matrices element-wise with tolerance
	for i := range expectedResult1 {
		for j := range expectedResult1[i] {
			if !util.EqualFloat(m1.Data[i][j], expectedResult1[i][j]) {
				t.Errorf("Expected result: %v, Got: %v", expectedResult1, m1.Data)
				break
			}
		}
	}

	// Test case 2: Zero scalar
	matrix2 := [][]float64{
		{1.1, 2.2, 3.3},
		{4.4, 5.5, 6.6},
	}
	scalar2 := 0.0
	expectedResult2 := [][]float64{
		{0.0, 0.0, 0.0},
		{0.0, 0.0, 0.0},
	}

	m2, err := NewMatrix(matrix2)
	if err != nil {
		t.Errorf("Error creating matrix2: %v", err)
	}

	m2.Scale(scalar2)

	// Compare matrices element-wise with tolerance
	for i := range expectedResult2 {
		for j := range expectedResult2[i] {
			if !util.EqualFloat(m2.Data[i][j], expectedResult2[i][j]) {
				t.Errorf("Expected result: %v, Got: %v", expectedResult2, m2.Data)
				break
			}
		}
	}

	fmt.Printf("Runtime: %v\n", time.Since(start))
}

func TestMatrixTranspose(t *testing.T) {
	start := time.Now()

	// Test case 1: Valid matrix
	matrix1 := [][]float64{
		{1.1, 2.2, 3.3},
		{4.4, 5.5, 6.6},
	}
	expectedTranspose1 := [][]float64{
		{1.1, 4.4},
		{2.2, 5.5},
		{3.3, 6.6},
	}

	m1, err := NewMatrix(matrix1)
	if err != nil {
		t.Errorf("Error creating matrix1: %v", err)
	}

	transposed := m1.Transpose()

	// Compare matrices element-wise with tolerance
	for i := range expectedTranspose1 {
		for j := range expectedTranspose1[i] {
			if !util.EqualFloat(transposed.Data[i][j], expectedTranspose1[i][j]) {
				t.Errorf("Expected transpose: %v, Got: %v", expectedTranspose1, transposed.Data)
				break
			}
		}
	}

	fmt.Printf("Runtime: %v\n", time.Since(start))
}

func TestMatrixMultiply(t *testing.T) {
	start := time.Now()

	// Test case 1: Valid multiplication
	// Matrix 1: a 2x5 matrix
	matrix1 := [][]float64{
		{1, 2, 3, 4, 5},
		{6, 7, 8, 9, 10},
	}

	// Matrix 2: a 5x2 matrix
	matrix2 := [][]float64{
		{11, 12},
		{13, 14},
		{15, 16},
		{17, 18},
		{19, 20},
	}

	// Expected Result: a 2x2 matrix
	expectedResult := [][]float64{
		{245, 260},
		{620, 660},
	}

	expectedMatrix, err := NewMatrix(expectedResult)
	if err != nil {
		t.Errorf("Error creating expectedMatrix: %v", err)
	}

	m1, err := NewMatrix(matrix1)
	if err != nil {
		t.Errorf("Error creating matrix1: %v", err)
	}

	m2, err := NewMatrix(matrix2)
	if err != nil {
		t.Errorf("Error creating matrix2: %v", err)
	}

	result, err := m1.Multiply(m2)
	if err != nil {
		t.Errorf("Error during matrix multiplication: %v", err)
	}

	if !equalMatrix(result, expectedMatrix) {
		t.Errorf("Expected result: %v, Got: %v", expectedMatrix, result)
	}

	// Test case 2: Incompatible dimensions
	matrix3 := [][]float64{
		{1, 2, 3},
		{4, 5, 6},
	}

	matrix4 := [][]float64{
		{7, 8},
		{9, 10},
	}

	m3, err := NewMatrix(matrix3)
	if err != nil {
		t.Errorf("Error creating matrix3: %v", err)
	}

	m4, err := NewMatrix(matrix4)
	if err != nil {
		t.Errorf("Error creating matrix4: %v", err)
	}

	_, err = m3.Multiply(m4)
	if err == nil {
		t.Error("Expected error for matrices with incompatible dimensions, but got nil")
	}

	fmt.Printf("Runtime: %v\n", time.Since(start))
}

func TestLargeMatrixMultiply(t *testing.T) {
	// Create a 100x100 matrix filled with ones
	matrix1 := make([][]float64, 1000)
	for i := range matrix1 {
		matrix1[i] = make([]float64, 1000)
		for j := range matrix1[i] {
			matrix1[i][j] = 1.0
		}
	}

	// Create a 100x100 matrix with diagonal values as 2, others as 1
	matrix2 := make([][]float64, 1000)
	for i := range matrix2 {
		matrix2[i] = make([]float64, 1000)
		for j := range matrix2[i] {
			if i == j {
				matrix2[i][j] = 2.0
			} else {
				matrix2[i][j] = 1.0
			}
		}
	}

	m1, err := NewMatrix(matrix1)
	if err != nil {
		t.Errorf("Error creating matrix1: %v", err)
	}

	m2, err := NewMatrix(matrix2)
	if err != nil {
		t.Errorf("Error creating matrix2: %v", err)
	}

	// Create the expected result matrix where each element is 101
	expectedData := make([][]float64, m1.M)
	for i := range expectedData {
		expectedData[i] = make([]float64, m2.N)
		for j := range expectedData[i] {
			expectedData[i][j] = 1001.0
		}
	}
	expectedResult, err := NewMatrix(expectedData)
	if err != nil {
		t.Errorf("Error creating expected matrix: %v", err)
	}

	start := time.Now()

	result, err := m1.Multiply(m2)
	if err != nil {
		t.Errorf("Error during matrix multiplication: %v", err)
	}

	fmt.Printf("Strassen multiply took: %v\n", time.Since(start))

	start = time.Now()

	result = m1.multiplyStandard(m2)
	if err != nil {
		t.Errorf("Error during matrix multiplication: %v", err)
	}

	fmt.Printf("Naive multiply took: %v\n", time.Since(start))
	if !equalMatrix(result, expectedResult) {
		t.Errorf("Expected result:\n\n%s\nGot: %s\n", expectedResult.String(), result.String())
	}
}

func TestMatrixMultiplyElements(t *testing.T) {
	start := time.Now()

	m1, err := NewMatrix([][]float64{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	})
	if err != nil {
		t.Error(err)
	}

	m2, err := NewMatrix([][]float64{
		{1, 1, 1},
		{1, 1, 1},
		{1, 1, 1},
	})
	if err != nil {
		t.Error(err)
	}

	expectedResult, err := NewMatrix([][]float64{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	})

	result, err := m1.MultiplyElements(m2)
	if err != nil {
		t.Error(err)
	}

	fmt.Printf("Runtime: %v\n", time.Since(start))

	for i := range expectedResult.Data {
		for j := range expectedResult.Data[i] {
			if !util.EqualFloat(result.Data[i][j], expectedResult.Data[i][j]) {
				t.Errorf("Expected result: %v, Got: %v", expectedResult.Data, result.Data)
				break
			}
		}
	}
}

func TestMatrixSumRows(t *testing.T) {
	m1, err := NewMatrix([][]float64{
		{1, 1, 1},
		{2, 2, 2},
		{3, 3, 3},
		{4, 4, 4},
	})
	if err != nil {
		t.Error(err)
	}

	expectedResult, err := NewMatrix([][]float64{{10, 10, 10}})

	result := m1.SumRows()

	if !equalMatrix(result, expectedResult) {
		t.Errorf("result matrix != expected result:\n%s\n!=\n\n%s", result.String(), expectedResult.String())
	}
}

func TestMatrixSumColumns(t *testing.T) {
	m1, err := NewMatrix([][]float64{
		{1, 1, 1},
		{2, 2, 2},
		{3, 3, 3},
		{4, 4, 4},
	})
	if err != nil {
		t.Error(err)
	}

	expectedResult, err := NewMatrix([][]float64{
		{3},
		{6},
		{9},
		{12},
	})

	result := m1.SumColumns()

	if !equalMatrix(result, expectedResult) {
		t.Errorf("result matrix != expected result:\n%s\n!=\n\n%s", result.String(), expectedResult.String())
	}
}

func TestMatrixSquare(t *testing.T) {
	m11, _ := NewMatrix([][]float64{
		{3},
		{6},
		{9},
	})
	m12, _ := NewMatrix([][]float64{
		{3},
	})

	expectedM11, _ := NewMatrix([][]float64{
		{3, 0, 0, 0},
		{6, 0, 0, 0},
		{9, 0, 0, 0},
		{0, 0, 0, 0},
	})

	expectedM12, _ := NewMatrix([][]float64{
		{3, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
	})

	resultM11, resultM12 := square(m11, m12)

	if !equalMatrix(resultM11, expectedM11) {
		t.Errorf("result matrix != expected result:\n%s\n!=\n\n%s", resultM11.String(), expectedM11.String())
	}
	if !equalMatrix(resultM12, expectedM12) {
		t.Errorf("result matrix != expected result:\n%s\n!=\n\n%s", resultM12.String(), expectedM12.String())
	}
}

func TestMatrixSplit(t *testing.T) {
	m, _ := NewMatrix([][]float64{
		{1, 1, 2, 2},
		{1, 1, 2, 2},
		{3, 3, 4, 4},
		{3, 3, 4, 4},
	})

	expectedM11, _ := NewMatrix([][]float64{
		{1, 1},
		{1, 1},
	})
	expectedM12, _ := NewMatrix([][]float64{
		{2, 2},
		{2, 2},
	})
	expectedM21, _ := NewMatrix([][]float64{
		{3, 3},
		{3, 3},
	})
	expectedM22, _ := NewMatrix([][]float64{
		{4, 4},
		{4, 4},
	})

	m11, m12, m21, m22, err := split(m)
	if err != nil {
		t.Errorf("error splitting matrix: %s", err)
	}

	if !equalMatrix(m11, expectedM11) {
		t.Errorf("result matrix != expected result:\n%s\n!=\n\n%s", m11.String(), expectedM11.String())
	}
	if !equalMatrix(m12, expectedM12) {
		t.Errorf("result matrix != expected result:\n%s\n!=\n\n%s", m12.String(), expectedM12.String())
	}
	if !equalMatrix(m21, expectedM21) {
		t.Errorf("result matrix != expected result:\n%s\n!=\n\n%s", m21.String(), expectedM21.String())
	}
	if !equalMatrix(m22, expectedM22) {
		t.Errorf("result matrix != expected result:\n%s\n!=\n\n%s", m22.String(), expectedM22.String())
	}
}

func TestMatrixCombine(t *testing.T) {
	m11, _ := NewMatrix([][]float64{
		{1, 1},
		{1, 1},
	})
	m12, _ := NewMatrix([][]float64{
		{2, 2},
		{2, 2},
	})
	m21, _ := NewMatrix([][]float64{
		{3, 3},
		{3, 3},
	})
	m22, _ := NewMatrix([][]float64{
		{4, 4},
		{4, 4},
	})

	expectedResult, _ := NewMatrix([][]float64{
		{1, 1, 2, 2},
		{1, 1, 2, 2},
		{3, 3, 4, 4},
		{3, 3, 4, 4},
	})

	result, err := combine(m11, m12, m21, m22, m11.M)
	if err != nil {
		t.Errorf("error combinin matrices: %s", err)
	}

	if !equalMatrix(result, expectedResult) {
		t.Errorf("result matrix != expected result:\n%s\n!=\n\n%s", result.String(), expectedResult.String())
	}
}
