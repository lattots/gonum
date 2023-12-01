package vector

import (
	"fmt"
	"github.com/lattots/gonum/pkg/util"
	"math"
	"reflect"
	"testing"
	"time"
)

func TestNewVector(t *testing.T) {
	fmt.Println("Testing NewVector...")
	start := time.Now()

	// Test case 1: Valid vector creation
	validInput := []float64{0, 1, 2, 3}
	expectedOutput := &Vector{
		Vec: validInput,
		Dim: 4,
		Len: math.Sqrt(0 + 1 + 4 + 9),
	}

	v, err := NewVector(validInput)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if !reflect.DeepEqual(v, expectedOutput) {
		t.Errorf("Expected %+v\nGot: %+v", expectedOutput, v)
	}

	// Test case 2: Invalid vector creation
	invalidInput := []float64(nil)
	expectedError := "CANNOT INITIALIZE VECTOR WITH NIL ARRAY"

	_, err = NewVector(invalidInput)
	if err == nil || err.Error() != expectedError {
		t.Errorf("Expected error: %s\n Got error: %v", expectedError, err)
	}

	fmt.Printf("Runtime: %v\n", time.Since(start))
}

func TestCosineSimilarity(t *testing.T) {
	fmt.Println("Testing CosineSimilarity...")
	start := time.Now()

	// Test case 1: Valid cosine similarity calculation
	vector1, _ := NewVector([]float64{1, 2, 3})
	vector2, _ := NewVector([]float64{-1, 0, 2})
	expectedSimilarity := (1*-1 + 2*0 + 3*2) / (math.Sqrt(14) * math.Sqrt(5))

	similarity, err := vector1.CosineSimilarity(vector2)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if similarity != expectedSimilarity {
		t.Errorf("Expected similarity: %f\nGot: %f", expectedSimilarity, similarity)
	}

	// Test case 2: Invalid dimensionality
	vector3, _ := NewVector([]float64{1, 2, 3})
	vector4, _ := NewVector([]float64{-1, 0, 2, 4})
	expectedError := "CANNOT CALCULATE SIMILARITY FOR VECTORS WITH DIFFERENT DIMENSIONALITY"

	_, err = vector3.CosineSimilarity(vector4)
	if err == nil || err.Error() != expectedError {
		t.Errorf("Expected error: %s\nGot error: %v", expectedError, err)
	}

	fmt.Printf("Runtime: %v\n", time.Since(start))
}

func TestCalculateLength(t *testing.T) {
	fmt.Println("Testing CalculateLength...")
	start := time.Now()

	// Test case 1: Valid length calculation
	vector, _ := NewVector([]float64{3, 4})
	expectedLength := 5.0
	if vector.Len != expectedLength {
		t.Errorf("Expected length: %f\nGot: %f", expectedLength, vector.Len)
	}

	// Test case 2: Empty vector
	emptyVector, _ := NewVector([]float64{})
	if emptyVector.Len != 0 {
		t.Errorf("Expected length: 0\nGot: %f", emptyVector.Len)
	}

	fmt.Printf("Runtime: %v\n", time.Since(start))
}

func TestNormalizeVector(t *testing.T) {
	fmt.Println("Testing NormalizeVector...")
	start := time.Now()

	// Test case 1: Valid normalization
	vector, _ := NewVector([]float64{3, 4})
	vector.NormalizeVector()
	expectedVector, _ := NewVector([]float64{0.6, 0.8})
	for i := range vector.Vec {
		if !util.EqualFloat(vector.Vec[i], expectedVector.Vec[i]) {
			t.Errorf("Expected %+v\nGot: %+v", expectedVector, vector)
		}
	}

	// Test case 2: Empty vector
	emptyVector, _ := NewVector([]float64{})
	emptyVector.NormalizeVector()
	if emptyVector.Len != 0 {
		t.Errorf("Expected length: 0\nGot: %f", emptyVector.Len)
	}

	fmt.Printf("Runtime: %v\n", time.Since(start))
}

func TestScale(t *testing.T) {
	fmt.Println("Testing Scale...")
	start := time.Now()

	// Test case 1: Valid scaling
	vector, _ := NewVector([]float64{3, 4})
	vector.Scale(2)
	expectedVector, _ := NewVector([]float64{6, 8})
	if !reflect.DeepEqual(vector, expectedVector) {
		t.Errorf("Expected %+v\nGot: %+v", expectedVector, vector)
	}

	// Test case 2: Scaling by zero
	vector.Scale(0)
	expectedZeroVector, _ := NewVector([]float64{0, 0})
	if !reflect.DeepEqual(vector, expectedZeroVector) {
		t.Errorf("Expected %+v\nGot: %+v", expectedZeroVector, vector)
	}

	fmt.Printf("Runtime: %v\n", time.Since(start))
}

func TestCalculateDotProduct(t *testing.T) {
	fmt.Println("Testing CalculateDotProduct...")
	start := time.Now()

	// Test case 1: Valid dot product calculation
	vector1, _ := NewVector([]float64{1, 2, 3})
	vector2, _ := NewVector([]float64{4, 5, 6})
	dotProduct, err := vector1.CalculateDotProduct(vector2)
	expectedDotProduct := 32.0
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if dotProduct != expectedDotProduct {
		t.Errorf("Expected dot product: %f\nGot: %f", expectedDotProduct, dotProduct)
	}

	// Test case 2: Invalid dimensionality
	vector3, _ := NewVector([]float64{1, 2, 3})
	vector4, _ := NewVector([]float64{4, 5, 6, 7})
	_, err = vector3.CalculateDotProduct(vector4)
	expectedError := "CANNOT CALCULATE DOT PRODUCT FOR VECTORS WITH DIFFERENT DIMENSIONALITY"
	if err == nil || err.Error() != expectedError {
		t.Errorf("Expected error: %s\nGot error: %v", expectedError, err)
	}

	fmt.Printf("Runtime: %v\n", time.Since(start))
}

func TestSum(t *testing.T) {
	fmt.Println("Testing Sum...")
	start := time.Now()

	// Test case 1: Valid vector sum
	vector1, _ := NewVector([]float64{1, 2, 3})
	vector2, _ := NewVector([]float64{4, 5, 6})
	sumVector, err := vector1.Sum(vector2)
	expectedSum, _ := NewVector([]float64{5, 7, 9})
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if !reflect.DeepEqual(sumVector, expectedSum) {
		t.Errorf("Expected %+v\nGot: %+v", expectedSum, sumVector)
	}

	fmt.Printf("Runtime: %v\n", time.Since(start))
}

func TestSubtract(t *testing.T) {
	fmt.Println("Testing Subtract...")
	start := time.Now()

	// Test case 1: Valid vector subtraction
	vector1, _ := NewVector([]float64{4, 5, 6})
	vector2, _ := NewVector([]float64{1, 2, 3})
	diffVector, err := vector1.Subtract(vector2)
	expectedDiff, _ := NewVector([]float64{3, 3, 3})
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if !reflect.DeepEqual(diffVector, expectedDiff) {
		t.Errorf("Expected %+v\nGot: %+v", expectedDiff, diffVector)
	}

	fmt.Printf("Runtime: %v\n", time.Since(start))
}

func TestCalculateCrossProduct(t *testing.T) {
	fmt.Println("Testing CalculateCrossProduct...")
	start := time.Now()

	// Test case 1: Valid cross product calculation
	vector1, _ := NewVector([]float64{1, 2, 3})
	vector2, _ := NewVector([]float64{4, 5, 6})
	crossProduct, err := vector1.CalculateCrossProduct(vector2)
	expectedCrossProduct, _ := NewVector([]float64{-3, 6, -3})
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if !reflect.DeepEqual(crossProduct, expectedCrossProduct) {
		t.Errorf("Expected %+v\nGot: %+v", expectedCrossProduct, crossProduct)
	}

	// Test case 2: Invalid dimensionality
	vector3, _ := NewVector([]float64{1, 2, 3})
	vector4, _ := NewVector([]float64{4, 5})
	_, err = vector3.CalculateCrossProduct(vector4)
	expectedError := "CANNOT CALCULATE CROSS PRODUCT FOR VECTORS WITH OTHER THAN 3 DIMENSIONS"
	if err == nil || err.Error() != expectedError {
		t.Errorf("Expected error: %s\nGot error: %v", expectedError, err)
	}

	fmt.Printf("Runtime: %v\n", time.Since(start))
}

func TestString(t *testing.T) {
	fmt.Println("Testing String...")
	start := time.Now()

	// Test case 1: Vector with dimension greater than 3
	vector1 := &Vector{Vec: []float64{1, 2, 3, 4, 5}, Dim: 5, Len: 7.071}
	stringRepresentation1 := vector1.String()
	expectedString1 := "Dim: 5\nLen: 7.07\nVec: [1.0 2.0 3.0 ... 5.0]"
	if stringRepresentation1 != expectedString1 {
		t.Errorf("Expected: %s\nGot: %s", expectedString1, stringRepresentation1)
	}

	// Test case 2: Vector with dimension less than or equal to 3
	vector2 := &Vector{Vec: []float64{1, 2, 3}, Dim: 3, Len: 3.741}
	stringRepresentation2 := vector2.String()
	expectedString2 := "Dim: 3\nLen: 3.74\nVec: [1.0 2.0 3.0]"
	if stringRepresentation2 != expectedString2 {
		t.Errorf("Expected: %s\nGot: %s", expectedString2, stringRepresentation2)
	}

	fmt.Printf("Runtime: %v\n", time.Since(start))
}
