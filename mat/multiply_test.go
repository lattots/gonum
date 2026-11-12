package mat_test

import (
	"testing"

	"github.com/lattots/gonum/internal/util"
	"github.com/lattots/gonum/mat"
)

func TestMatrixDot(t *testing.T) {
	// Test case 1: Valid multiplication
	data1 := [][]int{
		{1, 2, 3, 4, 5},
		{6, 7, 8, 9, 10},
	}

	data2 := [][]int{
		{11, 12},
		{13, 14},
		{15, 16},
		{17, 18},
		{19, 20},
	}

	dataExpected := [][]int{
		{245, 260},
		{620, 660},
	}

	expected, err := mat.New(dataExpected)
	if err != nil {
		t.Errorf("Error creating matrix: %v", err)
	}

	m1, err := mat.New(data1)
	if err != nil {
		t.Errorf("Error creating matrix: %v", err)
	}

	m2, err := mat.New(data2)
	if err != nil {
		t.Errorf("Error creating matrix: %v", err)
	}

	result, err := mat.Dot(m1, m2)
	if err != nil {
		t.Errorf("Error during matrix multiplication: %v", err)
	}

	if !util.EqualMatrix(result, expected) {
		t.Errorf("Wrong result in matrix dot product. Want: %s\nGot: %s", expected, result)
	}

	// Test case 2: Incompatible dimensions
	data3 := [][]float64{
		{1, 2, 3},
		{4, 5, 6},
	}

	data4 := [][]float64{
		{7, 8},
		{9, 10},
	}

	m3, err := mat.New(data3)
	if err != nil {
		t.Errorf("Error creating matrix: %v", err)
	}

	m4, err := mat.New(data4)
	if err != nil {
		t.Errorf("Error creating matrix: %v", err)
	}

	_, err = mat.Dot(m3, m4)
	if err == nil {
		t.Error("Expected error for matrices with incompatible dimensions, but got nil")
	}
}

func TestLargeMatrixMultiply(t *testing.T) {
	const size = 1024

	// Create a 1024x1024 matrix filled with ones
	data1 := make([][]float32, size)
	for i := range data1 {
		data1[i] = make([]float32, size)
		for j := range data1[i] {
			data1[i][j] = 1
		}
	}

	// Create a 1024x1024 matrix with diagonal values as 2, others as 1
	data2 := make([][]float32, size)
	for i := range data2 {
		data2[i] = make([]float32, size)
		for j := range data2[i] {
			if i == j {
				data2[i][j] = 2
			} else {
				data2[i][j] = 1
			}
		}
	}

	m1, err := mat.New(data1)
	if err != nil {
		t.Errorf("Error creating matrix: %v", err)
	}

	m2, err := mat.New(data2)
	if err != nil {
		t.Errorf("Error creating matrix: %v", err)
	}

	// Create the expected result matrix where each element is 1025
	expectedData := make([][]float32, m1.M)
	for i := range expectedData {
		expectedData[i] = make([]float32, m2.N)
		for j := range expectedData[i] {
			expectedData[i][j] = 1025
		}
	}
	expected, err := mat.New(expectedData)
	if err != nil {
		t.Errorf("Error creating expected matrix: %v", err)
	}

	result, err := mat.Dot(m1, m2)
	if err != nil {
		t.Errorf("Error during matrix multiplication: %v", err)
	}

	if !util.EqualMatrix(result, expected) {
		t.Errorf("Wrong result in matrix dot product. Want: %s\nGot: %s", expected, result)
	}
}

func TestMatrixMul(t *testing.T) {
	m1, err := mat.New([][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	})
	if err != nil {
		t.Errorf("Error creating matrix: %v", err)
	}

	m2, err := mat.New([][]int{
		{1, 1, 1},
		{1, 1, 1},
		{1, 1, 1},
	})
	if err != nil {
		t.Errorf("Error creating matrix: %v", err)
	}

	expected, err := mat.New([][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	})
	if err != nil {
		t.Errorf("Error creating matrix: %v", err)
	}

	result, err := mat.Mul(m1, m2)
	if err != nil {
		t.Errorf("Error multiplying matrices element wise: %v", err)
	}

	if !util.EqualMatrix(result, expected) {
		t.Errorf("Wrong result in element wise matrix multiplication. Want: %s\nGot: %s", expected, result)
	}
}
