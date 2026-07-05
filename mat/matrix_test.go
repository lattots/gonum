package mat_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/lattots/gonum/internal/util"
	"github.com/lattots/gonum/mat"
)

func TestNew(t *testing.T) {
	start := time.Now()

	// Test case 1: Valid float matrix
	validFloatMatrix := [][]float64{
		{1.1, 2.2, 3.3},
		{4.4, 5.5, 6.6},
	}
	expectedM := 2
	expectedN := 3

	m1, err := mat.New(validFloatMatrix)
	if err != nil {
		t.Errorf("Error creating float matrix: %v", err)
	}

	if m1.M != expectedM || m1.N != expectedN {
		t.Errorf("Invalid dimensions. Want (%v x %v), Got (%v x %v)", expectedM, expectedN, m1.M, m1.N)
	}

	// Test case 2: Invalid matrix (nil)
	var nilMatrix [][]float64
	_, err = mat.New(nilMatrix)
	if err == nil {
		t.Error("Expected error for nil matrix, but got nil")
	}

	// Test case 3: Incomplete matrix (different row lengths)
	incompleteMatrix := [][]float64{
		{1.1, 2.2, 3.3},
		{4.4, 5.5},
	}
	_, err = mat.New(incompleteMatrix)
	if err == nil {
		t.Error("Expected error for an incomplete matrix, but got nil")
	}

	// Test case 4: Valid integer matrix
	validIntMatrix := [][]int{
		{1, 2},
		{3, 4},
	}
	expectedM = 2
	expectedN = 2

	m2, err := mat.New(validIntMatrix)
	if err != nil {
		t.Errorf("Error creating integer matrix: %v", err)
	}

	if m2.M != expectedM || m2.N != expectedN {
		t.Errorf("Invalid dimensions. Want (%v x %v), Got (%v x %v)", expectedM, expectedN, m2.M, m2.N)
	}

	fmt.Printf("Runtime: %v\n", time.Since(start))
}

func TestNewZeroMatrix(t *testing.T) {
	start := time.Now()

	const (
		rows    = 3
		columns = 5
	)

	m, err := mat.Zeros[int](rows, columns)
	if err != nil {
		t.Errorf("Error creating zero integer matrix: %v", err)
	}

	if m.M != rows || m.N != columns {
		t.Error("Created matrices dimensions are wrong")
	}

	for _, val := range m.Data {
		if val != 0 {
			t.Errorf("Found non-zero value in matrix: %v", val)
		}
	}

	fmt.Printf("Runtime: %v\n", time.Since(start))
}

func TestNewOneMatrix(t *testing.T) {
	start := time.Now()

	const (
		rows    = 3
		columns = 5
	)

	m, err := mat.Ones[int](rows, columns)
	if err != nil {
		t.Errorf("Error creating one integer matrix: %v", err)
	}

	if m.M != rows || m.N != columns {
		t.Error("Created matrices dimensions are wrong")
	}

	for _, val := range m.Data {
		if val != 1 {
			t.Errorf("Found non-zero value in matrix: %v", val)
		}
	}

	fmt.Printf("Runtime: %v\n", time.Since(start))
}

func TestAt(t *testing.T) {
	start := time.Now()

	data := [][]int{
		{1, 2},
		{3, 4},
	}

	m, err := mat.New(data)
	if err != nil {
		t.Errorf("Error creating matrix: %v", err)
	}

	if m.At(1, 1) != 1 {
		t.Errorf("Found wrong value at 1, 1. Want: %v\nGot: %v", m.Data[0], m.At(1, 1))
	}

	if m.At(2, 1) != 3 {
		t.Errorf("Found wrong value at 2, 1. Want: %v\nGot: %v", m.Data[2], m.At(2, 1))
	}

	fmt.Printf("Runtime: %v\n", time.Since(start))
}

func TestMatrixString(t *testing.T) {
	smallMatrix, err := mat.New([][]float64{
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

	largeMatrix, err := mat.New([][]float64{
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

func TestMatrixScale(t *testing.T) {
	start := time.Now()

	// Test case 1: Positive scalar
	data1 := [][]float64{
		{1.1, 2.2, 3.3},
		{4.4, 5.5, 6.6},
	}
	scalar1 := 2.0
	expectedData1 := [][]float64{
		{2.2, 4.4, 6.6},
		{8.8, 11.0, 13.2},
	}

	m1, err := mat.New(data1)
	if err != nil {
		t.Errorf("Error creating matrix: %v", err)
	}

	expected1, err := mat.New(expectedData1)
	if err != nil {
		t.Errorf("Error creating matrix: %v", err)
	}

	result1 := mat.Scale(m1, scalar1)

	if !util.EqualMatrix(result1, expected1) {
		t.Errorf("Wrong result in matrix scaling. Want: %s\nGot: %s", expected1, result1)
	}

	// Test case 2: Zero scalar
	data2 := [][]float64{
		{1.1, 2.2, 3.3},
		{4.4, 5.5, 6.6},
	}
	scalar2 := 0.0
	expectedData2 := [][]float64{
		{0.0, 0.0, 0.0},
		{0.0, 0.0, 0.0},
	}

	m2, err := mat.New(data2)
	if err != nil {
		t.Errorf("Error creating matrix: %v", err)
	}

	expected2, err := mat.New(expectedData2)
	if err != nil {
		t.Errorf("Error creating matrix: %v", err)
	}

	result2 := mat.Scale(m2, scalar2)

	if !util.EqualMatrix(result2, expected2) {
		t.Errorf("Wrong result in matrix scaling. Want: %s\nGot: %s", expected2, result2)
	}

	fmt.Printf("Runtime: %v\n", time.Since(start))
}

func TestMatrixAdd(t *testing.T) {
	start := time.Now()

	// Test case 1: Add a positive scalar float
	data1 := [][]float64{
		{1.1, 2.2, 3.3},
		{4.4, 5.5, 6.6},
	}
	scalar1 := 1.5
	expectedData1 := [][]float64{
		{2.6, 3.7, 4.8},
		{5.9, 7.0, 8.1},
	}

	m1, err := mat.New(data1)
	if err != nil {
		t.Errorf("Error creating matrix: %v", err)
	}

	expected1, err := mat.New(expectedData1)
	if err != nil {
		t.Errorf("Error creating matrix: %v", err)
	}

	result1 := mat.Add(m1, scalar1)

	if !util.EqualMatrix(result1, expected1) {
		t.Errorf("Wrong result in matrix scalar addition. Want: %s\nGot: %s", expected1, result1)
	}

	// Test case 2: Add a negative scalar float
	data2 := [][]float64{
		{1.0, 2.0},
		{3.0, 4.0},
	}
	scalar2 := -2.0
	expectedData2 := [][]float64{
		{-1.0, 0.0},
		{1.0, 2.0},
	}

	m2, err := mat.New(data2)
	if err != nil {
		t.Errorf("Error creating matrix: %v", err)
	}

	expected2, err := mat.New(expectedData2)
	if err != nil {
		t.Errorf("Error creating matrix: %v", err)
	}

	result2 := mat.Add(m2, scalar2)

	if !util.EqualMatrix(result2, expected2) {
		t.Errorf("Wrong result in matrix scalar addition. Want: %s\nGot: %s", expected2, result2)
	}

	fmt.Printf("Runtime: %v\n", time.Since(start))
}

func TestMatrixMap(t *testing.T) {
	start := time.Now()

	// Test case 1: Squaring elements using float64
	data1 := [][]float64{
		{1.0, 2.0, 3.0},
		{4.0, 5.0, 6.0},
	}
	expectedData1 := [][]float64{
		{1.0, 4.0, 9.0},
		{16.0, 25.0, 36.0},
	}

	m1, err := mat.New(data1)
	if err != nil {
		t.Errorf("Error creating matrix: %v", err)
	}

	expected1, err := mat.New(expectedData1)
	if err != nil {
		t.Errorf("Error creating matrix: %v", err)
	}

	// Map function to square numbers
	squareFn := func(val float64) float64 {
		return val * val
	}

	result1 := mat.Map(m1, squareFn)

	if !util.EqualMatrix(result1, expected1) {
		t.Errorf("Wrong result in matrix map (float square). Want: %s\nGot: %s", expected1, result1)
	}

	// Test case 2: Absolute value calculation on generic int32 matrix
	data2 := [][]int32{
		{-1, 2, -3},
		{4, -5, 6},
	}
	expectedData2 := [][]int32{
		{1, 2, 3},
		{4, 5, 6},
	}

	m2, err := mat.New(data2)
	if err != nil {
		t.Errorf("Error creating matrix: %v", err)
	}

	expected2, err := mat.New(expectedData2)
	if err != nil {
		t.Errorf("Error creating matrix: %v", err)
	}

	// Map function for absolute values
	absFn := func(val int32) int32 {
		if val < 0 {
			return -val
		}
		return val
	}

	result2 := mat.Map(m2, absFn)

	if !util.EqualMatrix(result2, expected2) {
		t.Errorf("Wrong result in matrix map (int abs). Want: %s\nGot: %s", expected2, result2)
	}

	fmt.Printf("Runtime: %v\n", time.Since(start))
}

func TestMatrixTranspose(t *testing.T) {
	start := time.Now()

	// Test case 1: Valid matrix
	data := [][]float64{
		{1.1, 2.2, 3.3},
		{4.4, 5.5, 6.6},
	}
	expectedData := [][]float64{
		{1.1, 4.4},
		{2.2, 5.5},
		{3.3, 6.6},
	}

	m, err := mat.New(data)
	if err != nil {
		t.Errorf("Error creating matrix: %v", err)
	}

	expected, err := mat.New(expectedData)
	if err != nil {
		t.Errorf("Error creating matrix: %v", err)
	}

	result := mat.T(m)

	if !util.EqualMatrix(result, expected) {
		t.Errorf("Wrong result in matrix transpose. Want: %s\nGot: %s", expected, result)
	}

	fmt.Printf("Runtime: %v\n", time.Since(start))
}

func TestMatrixMin(t *testing.T) {
	start := time.Now()

	// Test case 1: Valid matrix with varied numbers
	m, err := mat.New([][]float64{
		{5.5, 12.1, -3.4},
		{0.0, -10.2, 8.8},
	})
	if err != nil {
		t.Errorf("Error creating matrix: %v", err)
	}

	expected := -10.2
	result := mat.Min(m)

	if result != expected {
		t.Errorf("Wrong result in Matrix Min. Want: %f, Got: %f", expected, result)
	}

	// Test case 2: Empty matrix (should panic)
	emptyMat := &mat.Mat[float64]{
		M:    0,
		N:    0,
		Data: []float64{},
	}

	func() {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("Min did not panic on an empty matrix")
			}
		}()
		mat.Min(emptyMat)
	}()

	fmt.Printf("Runtime: %v\n", time.Since(start))
}

func TestMatrixMax(t *testing.T) {
	start := time.Now()

	// Test case 1: Valid matrix with varied numbers
	m, err := mat.New([][]float64{
		{-5.5, 12.1, -3.4},
		{0.0, 102.4, 8.8},
	})
	if err != nil {
		t.Errorf("Error creating matrix: %v", err)
	}

	expected := 102.4
	result := mat.Max(m)

	if result != expected {
		t.Errorf("Wrong result in Matrix Max. Want: %f, Got: %f", expected, result)
	}

	// Test case 2: Empty matrix (should panic)
	emptyMat := &mat.Mat[float64]{
		M:    0,
		N:    0,
		Data: []float64{},
	}

	func() {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("Max did not panic on an empty matrix")
			}
		}()
		mat.Max(emptyMat)
	}()

	fmt.Printf("Runtime: %v\n", time.Since(start))
}

func TestSliceRows(t *testing.T) {
	start := time.Now()

	// Test case 1: Slice middle rows using float64 matrix
	data1 := [][]float64{
		{1.0, 2.0, 3.0}, // Row 0
		{4.0, 5.0, 6.0}, // Row 1
		{7.0, 8.0, 9.0}, // Row 2
	}
	expectedData1 := [][]float64{
		{4.0, 5.0, 6.0},
		{7.0, 8.0, 9.0},
	}

	m1, err := mat.New(data1)
	if err != nil {
		t.Errorf("Error creating matrix: %v", err)
	}

	expected1, err := mat.New(expectedData1)
	if err != nil {
		t.Errorf("Error creating matrix: %v", err)
	}

	// Slice from row 1 to 3 (exclusive, so rows 1 and 2)
	result1 := mat.SliceRows(m1, 1, 3)

	if !util.EqualMatrix(result1, expected1) {
		t.Errorf("Wrong result in SliceRows (float64 middle slice). Want: %s\nGot: %s", expected1, result1)
	}

	// Test case 2: Slice beginning rows using generic int32 matrix
	data2 := [][]int32{
		{10, 20}, // Row 0
		{30, 40}, // Row 1
		{50, 60}, // Row 2
	}
	expectedData2 := [][]int32{
		{10, 20},
		{30, 40},
	}

	m2, err := mat.New(data2)
	if err != nil {
		t.Errorf("Error creating matrix: %v", err)
	}

	expected2, err := mat.New(expectedData2)
	if err != nil {
		t.Errorf("Error creating matrix: %v", err)
	}

	// Slice from row 0 to 2 (exclusive, so rows 0 and 1)
	result2 := mat.SliceRows(m2, 0, 2)

	if !util.EqualMatrix(result2, expected2) {
		t.Errorf("Wrong result in SliceRows (int32 front slice). Want: %s\nGot: %s", expected2, result2)
	}

	// Test case 3: Verify Deep Copy (modifying result shouldn't affect original matrix)
	if len(result2.Data) > 0 {
		originalVal := result2.Data[0]
		result2.Data[0] = 999 // Mutate the sliced result

		// If it's a deep copy, the parent matrix m2.Data[0] must still be its original value (10)
		if m2.Data[0] == 999 {
			t.Errorf("SliceRows did not perform a deep copy; modifying slice corrupted the original matrix data.")
		}
		result2.Data[0] = originalVal // Restore value just in case
	}

	// Test case 4: Invalid index ranges (should panic when start >= end)
	func() {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("SliceRows did not panic when start equaled end")
			}
		}()
		mat.SliceRows(m1, 2, 2)
	}()

	func() {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("SliceRows did not panic when start was greater than end")
			}
		}()
		mat.SliceRows(m1, 2, 1)
	}()

	fmt.Printf("Runtime: %v\n", time.Since(start))
}
