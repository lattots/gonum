package matrix

import (
	"reflect"
	"testing"
)

func TestNewMatrix(t *testing.T) {
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
}

func TestMatrixString(t *testing.T) {
	// Test case 1: Valid matrix
	validMatrix := [][]float64{
		{1.1, 2.2, 3.3},
		{4.4, 5.5, 6.6},
	}
	expectedString := "2 x 3\n1.10 2.20 3.30\n4.40 5.50 6.60\n"

	matrix, err := NewMatrix(validMatrix)
	if err != nil {
		t.Errorf("Error creating matrix: %v", err)
	}

	resultString := matrix.String()
	if resultString != expectedString {
		t.Errorf("Expected: %s, Got: %s", expectedString, resultString)
	}
}

func TestMatrixSum(t *testing.T) {
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

	result, err := m1.Sum(m2)
	if err != nil {
		t.Errorf("Error during matrix summation: %v", err)
	}

	// Compare matrices element-wise with tolerance
	for i := range expectedSum {
		for j := range expectedSum[i] {
			if !equalFloat(result.Data[i][j], expectedSum[i][j]) {
				t.Errorf("Expected sum: %v, Got: %v", expectedSum, result.Data)
				break
			}
		}
	}

	// Test case 2: Matrices with different dimensions
	matrix3 := [][]float64{
		{1.1, 2.2, 3.3},
		{4.4, 5.5, 6.6},
	}
	matrix4 := [][]float64{
		{0.1, 1.2},
		{3.4, 4.5},
	}

	m3, err := NewMatrix(matrix3)
	if err != nil {
		t.Errorf("Error creating matrix3: %v", err)
	}

	m4, err := NewMatrix(matrix4)
	if err != nil {
		t.Errorf("Error creating matrix4: %v", err)
	}

	_, err = m3.Sum(m4)
	if err == nil {
		t.Error("Expected error for matrices with different dimensions, but got nil")
	}
}
