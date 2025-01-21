package matrix

import (
	"fmt"
	"reflect"
	"testing"
	"time"

	"github.com/lattots/gonum/pkg/util"
)

func TestNewMatrix(t *testing.T) {
	fmt.Println("Testing NewMatrix...")
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
	fmt.Println("Testing NewMatrix...")
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
	fmt.Println("Testing String...")
	start := time.Now()

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

	fmt.Printf("Runtime: %v\n", time.Since(start))
}

func TestMatrixSum(t *testing.T) {
	fmt.Println("Testing Sum...")
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

	result, err := m1.Sum(m2)
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

	fmt.Printf("Runtime: %v\n", time.Since(start))
}

func TestMatrixSubtract(t *testing.T) {
	fmt.Println("Testing Subtract...")
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

	result, err := m1.Subtract(m2)
	if err != nil {
		t.Errorf("Error during matrix subtraction: %v", err)
	}

	// Compare matrices element-wise with tolerance
	for i := range expectedResult {
		for j := range expectedResult[i] {
			if !util.EqualFloat(result.Data[i][j], expectedResult[i][j]) {
				t.Errorf("Expected result: %v, Got: %v", expectedResult, result.Data)
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

	_, err = m3.Subtract(m4)
	if err == nil {
		t.Error("Expected error for matrices with different dimensions, but got nil")
	}

	fmt.Printf("Runtime: %v\n", time.Since(start))
}

func TestMatrixScale(t *testing.T) {
	fmt.Println("Testing Scale...")
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
	fmt.Println("Testing Transpose...")
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

	transposed, err := m1.Transpose()
	if err != nil {
		t.Errorf("Error during matrix transpose: %v", err)
	}

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
	fmt.Println("Testing Multiply...")
	start := time.Now()

	// Test case 1: Valid multiplication
	matrix1 := [][]float64{
		{1, 2},
		{3, 4},
	}

	matrix2 := [][]float64{
		{5, 6},
		{7, 8},
	}

	expectedResult := [][]float64{
		{19, 22},
		{43, 50},
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

	// Compare matrices element-wise with tolerance
	for i := range expectedResult {
		for j := range expectedResult[i] {
			if !util.EqualFloat(result.Data[i][j], expectedResult[i][j]) {
				t.Errorf("Expected result: %v, Got: %v", expectedResult, result.Data)
				break
			}
		}
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
	fmt.Println("Testing Large Matrix Multiply...")

	// Create a 100x100 matrix filled with ones
	matrix1 := make([][]float64, 100)
	for i := range matrix1 {
		matrix1[i] = make([]float64, 100)
		for j := range matrix1[i] {
			matrix1[i][j] = 1.0
		}
	}

	// Create a 100x100 matrix with diagonal values as 2, others as 1
	matrix2 := make([][]float64, 100)
	for i := range matrix2 {
		matrix2[i] = make([]float64, 100)
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
			expectedData[i][j] = 101.0
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

	fmt.Printf("Runtime: %v\n", time.Since(start))

	// Compare matrices element-wise with tolerance
	for i := range expectedResult.Data {
		for j := range expectedResult.Data[i] {
			if !util.EqualFloat(result.Data[i][j], expectedResult.Data[i][j]) {
				t.Errorf("Expected result: %v, Got: %v", expectedResult.Data, result.Data)
				break
			}
		}
	}

}
