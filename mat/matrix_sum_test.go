package mat_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/lattots/gonum/internal/util"
	"github.com/lattots/gonum/mat"
)

func TestMatrixSum(t *testing.T) {
	start := time.Now()

	// Test case 1: Valid matrices with the same dimensions
	data1 := [][]float64{
		{1.1, 2.2, 3.3},
		{4.4, 5.5, 6.6},
	}
	data2 := [][]float64{
		{0.1, 1.2, 2.3},
		{3.4, 4.5, 5.6},
	}
	dataExpected := [][]float64{
		{1.2, 3.4, 5.6},
		{7.8, 10.0, 12.2},
	}

	m1, err := mat.New(data1)
	if err != nil {
		t.Errorf("Error creating matrix: %v", err)
	}

	m2, err := mat.New(data2)
	if err != nil {
		t.Errorf("Error creating matrix: %v", err)
	}

	expected, err := mat.New(dataExpected)
	if err != nil {
		t.Errorf("Error creating matrix: %v", err)
	}

	result := mat.Sum(m1, m2)

	if !util.EqualMatrix(result, expected) {
		t.Errorf("Wrong result in matrix addition. Want: %s\nGot: %s", expected, result)
	}

	fmt.Printf("Runtime: %v\n", time.Since(start))
}

func TestMatrixSubtract(t *testing.T) {
	start := time.Now()

	// Test case 1: Valid matrices with the same dimensions
	data1 := [][]float64{
		{1.1, 2.2, 3.3},
		{4.4, 5.5, 6.6},
	}
	data2 := [][]float64{
		{0.1, 1.2, 2.3},
		{3.4, 4.5, 5.6},
	}
	dataExpected := [][]float64{
		{1.0, 1.0, 1.0},
		{1.0, 1.0, 1.0},
	}

	m1, err := mat.New(data1)
	if err != nil {
		t.Errorf("Error creating matrix: %v", err)
	}

	m2, err := mat.New(data2)
	if err != nil {
		t.Errorf("Error creating matrix: %v", err)
	}

	expected, err := mat.New(dataExpected)
	if err != nil {
		t.Errorf("Error creating matrix: %v", err)
	}

	result := mat.Subtract(m1, m2)

	if !util.EqualMatrix(result, expected) {
		t.Errorf("Wrong result in matrix subtraction. Want: %s\nGot: %s", expected, result)
	}

	fmt.Printf("Runtime: %v\n", time.Since(start))
}

func TestMatrixSumRows(t *testing.T) {
	start := time.Now()

	m, err := mat.New([][]float64{
		{1, 1, 1},
		{2, 2, 2},
		{3, 3, 3},
		{4, 4, 4},
	})
	if err != nil {
		t.Errorf("Error creating matrix: %v", err)
	}

	expected, err := mat.New([][]float64{{10, 10, 10}})
	if err != nil {
		t.Errorf("Error creating matrix: %v", err)
	}

	result := mat.SumRows(m)

	if !util.EqualMatrix(result, expected) {
		t.Errorf("Wrong result in column wise addition. Want: %s\nGot: %s", expected, result)
	}

	fmt.Printf("Runtime: %v\n", time.Since(start))
}

func TestMatrixSumColumns(t *testing.T) {
	start := time.Now()

	m, err := mat.New([][]float64{
		{1, 1, 1},
		{2, 2, 2},
		{3, 3, 3},
		{4, 4, 4},
	})
	if err != nil {
		t.Errorf("Error creating matrix: %v", err)
	}

	expected, err := mat.New([][]float64{{3}, {6}, {9}, {12}})
	if err != nil {
		t.Errorf("Error creating matrix: %v", err)
	}

	result := mat.SumColumns(m)

	if !util.EqualMatrix(result, expected) {
		t.Errorf("Wrong result in row wise addition. Want: %s\nGot: %s", expected, result)
	}

	fmt.Printf("Runtime: %v\n", time.Since(start))
}

func TestMatrixAddRowVector(t *testing.T) {
	start := time.Now()

	// Test case 1: Valid dimensions (2x3 matrix and 1x3 row vector)
	m, err := mat.New([][]float64{
		{1.0, 2.0, 3.0},
		{4.0, 5.0, 6.0},
	})
	if err != nil {
		t.Errorf("Error creating matrix: %v", err)
	}

	row, err := mat.New([][]float64{
		{0.5, 1.0, 1.5},
	})
	if err != nil {
		t.Errorf("Error creating matrix: %v", err)
	}

	expected, err := mat.New([][]float64{
		{1.5, 3.0, 4.5},
		{4.5, 6.0, 7.5},
	})
	if err != nil {
		t.Errorf("Error creating matrix: %v", err)
	}

	result := mat.AddRowVector(m, row)

	if !util.EqualMatrix(result, expected) {
		t.Errorf("Wrong result in AddRowVector. Want: %s\nGot: %s", expected, result)
	}

	// Test case 2: Dimension mismatch (should panic)
	invalidRow, err := mat.New([][]float64{
		{0.5, 1.0}, // Wrong number of columns (2 instead of 3)
	})
	if err != nil {
		t.Errorf("Error creating matrix: %v", err)
	}

	func() {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("AddRowVector did not panic on dimension mismatch")
			}
		}()
		mat.AddRowVector(m, invalidRow)
	}()

	fmt.Printf("Runtime: %v\n", time.Since(start))
}

func TestMatrixAddColVector(t *testing.T) {
	start := time.Now()

	// Test case 1: Valid dimensions (2x3 matrix and 2x1 column vector)
	m, err := mat.New([][]float64{
		{1.0, 2.0, 3.0},
		{4.0, 5.0, 6.0},
	})
	if err != nil {
		t.Errorf("Error creating matrix: %v", err)
	}

	col, err := mat.New([][]float64{
		{10.0},
		{20.0},
	})
	if err != nil {
		t.Errorf("Error creating matrix: %v", err)
	}

	expected, err := mat.New([][]float64{
		{11.0, 12.0, 13.0},
		{24.0, 25.0, 26.0},
	})
	if err != nil {
		t.Errorf("Error creating matrix: %v", err)
	}

	result := mat.AddColVector(m, col)

	if !util.EqualMatrix(result, expected) {
		t.Errorf("Wrong result in AddColVector. Want: %s\nGot: %s", expected, result)
	}

	// Test case 2: Dimension mismatch (should panic)
	invalidCol, err := mat.New([][]float64{
		{10.0},
		{20.0},
		{30.0}, // Wrong number of rows (3 instead of 2)
	})
	if err != nil {
		t.Errorf("Error creating matrix: %v", err)
	}

	func() {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("AddColVector did not panic on dimension mismatch")
			}
		}()
		mat.AddColVector(m, invalidCol)
	}()

	fmt.Printf("Runtime: %v\n", time.Since(start))
}
