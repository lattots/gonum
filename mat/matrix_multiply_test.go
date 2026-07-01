package mat_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/lattots/gonum/internal/util"
	"github.com/lattots/gonum/mat"
)

func TestMatrixDot(t *testing.T) {
	start := time.Now()

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

	fmt.Printf("Runtime: %v\n", time.Since(start))
}

func TestLargeMatrixMultiply(t *testing.T) {
	// Create a 1000x1000 matrix filled with ones
	data1 := make([][]int, 1000)
	for i := range data1 {
		data1[i] = make([]int, 1000)
		for j := range data1[i] {
			data1[i][j] = 1.0
		}
	}

	// Create a 1000x1000 matrix with diagonal values as 2, others as 1
	data2 := make([][]int, 1000)
	for i := range data2 {
		data2[i] = make([]int, 1000)
		for j := range data2[i] {
			if i == j {
				data2[i][j] = 2.0
			} else {
				data2[i][j] = 1.0
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

	// Create the expected result matrix where each element is 1001
	expectedData := make([][]int, m1.M)
	for i := range expectedData {
		expectedData[i] = make([]int, m2.N)
		for j := range expectedData[i] {
			expectedData[i][j] = 1001.0
		}
	}
	expected, err := mat.New(expectedData)
	if err != nil {
		t.Errorf("Error creating expected matrix: %v", err)
	}

	start := time.Now()

	result, err := mat.Dot(m1, m2)
	if err != nil {
		t.Errorf("Error during matrix multiplication: %v", err)
	}

	fmt.Printf("Strassen multiply took: %v\n", time.Since(start))

	if !util.EqualMatrix(result, expected) {
		t.Errorf("Wrong result in matrix dot product. Want: %s\nGot: %s", expected, result)
	}
}

func TestMatrixMul(t *testing.T) {
	start := time.Now()

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

	fmt.Printf("Runtime: %v\n", time.Since(start))
}

func TestMatrixSquare(t *testing.T) {
	m1, _ := mat.New([][]int{
		{3},
		{6},
		{9},
	})
	m2, _ := mat.New([][]int{
		{3},
	})

	expected1, _ := mat.New([][]int{
		{3, 0, 0, 0},
		{6, 0, 0, 0},
		{9, 0, 0, 0},
		{0, 0, 0, 0},
	})

	expected2, _ := mat.New([][]int{
		{3, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
	})

	result1, result2 := mat.Square(m1, m2)

	if !util.EqualMatrix(result1, expected1) {
		t.Errorf("Wrong result squaring matrices. Want: %s\nGot: %s", expected1, result1)
	}
	if !util.EqualMatrix(result2, expected2) {
		t.Errorf("Wrong result squaring matrices. Want: %s\nGot: %s", expected2, result2)
	}
}

func TestMatrixSplit(t *testing.T) {
	m, _ := mat.New([][]int{
		{1, 1, 2, 2},
		{1, 1, 2, 2},
		{3, 3, 4, 4},
		{3, 3, 4, 4},
	})

	expected11, _ := mat.New([][]int{
		{1, 1},
		{1, 1},
	})
	expected12, _ := mat.New([][]int{
		{2, 2},
		{2, 2},
	})
	expected21, _ := mat.New([][]int{
		{3, 3},
		{3, 3},
	})
	expected22, _ := mat.New([][]int{
		{4, 4},
		{4, 4},
	})

	m11, m12, m21, m22 := mat.Split(m)

	if !util.EqualMatrix(m11, expected11) {
		t.Errorf("Wrong result splitting a matrix. Want: %s\nGot: %s", expected11, m11)
	}
	if !util.EqualMatrix(m12, expected12) {
		t.Errorf("Wrong result splitting a matrix. Want: %s\nGot: %s", expected12, m12)
	}
	if !util.EqualMatrix(m21, expected21) {
		t.Errorf("Wrong result splitting a matrix. Want: %s\nGot: %s", expected21, m21)
	}
	if !util.EqualMatrix(m22, expected22) {
		t.Errorf("Wrong result splitting a matrix. Want: %s\nGot: %s", expected22, m22)
	}
}

func TestMatrixCombine(t *testing.T) {
	m11, _ := mat.New([][]int{
		{1, 1},
		{1, 1},
	})
	m12, _ := mat.New([][]int{
		{2, 2},
		{2, 2},
	})
	m21, _ := mat.New([][]int{
		{3, 3},
		{3, 3},
	})
	m22, _ := mat.New([][]int{
		{4, 4},
		{4, 4},
	})

	expected, _ := mat.New([][]int{
		{1, 1, 2, 2},
		{1, 1, 2, 2},
		{3, 3, 4, 4},
		{3, 3, 4, 4},
	})

	result := mat.Combine(m11, m12, m21, m22, m11.M)

	if !util.EqualMatrix(result, expected) {
		t.Errorf("Wrong result combining matrices. Want: %s\nGot: %s", expected, result)
	}
}
