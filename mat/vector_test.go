package mat_test

import (
	"fmt"
	"math"
	"testing"
	"time"

	"github.com/lattots/gonum/internal/util"
	"github.com/lattots/gonum/mat"
)

func TestVectorLength(t *testing.T) {
	start := time.Now()

	// Test case 1: Valid length calculation (3x1 column vector)
	m1, err := mat.New([][]float64{{3}, {4}})
	if err != nil {
		t.Fatalf("Failed to create matrix: %v", err)
	}
	expectedLength := 5.0
	if m1.Length() != expectedLength {
		t.Errorf("Expected length: %f, Got: %f", expectedLength, m1.Length())
	}

	// Test case 2: Panic on non-vector matrix (2x2 matrix)
	m2, _ := mat.New([][]float64{{1, 2}, {3, 4}})
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected length function to panic on a non-vector matrix, but it did not")
		}
	}()
	_ = m2.Length()

	fmt.Printf("Runtime: %v\n", time.Since(start))
}

func TestVectorNormalize(t *testing.T) {
	start := time.Now()

	// Test case 1: Valid normalization (1x2 row vector)
	m, err := mat.New([][]float64{{3, 4}})
	if err != nil {
		t.Fatalf("Failed to create matrix: %v", err)
	}

	expected, _ := mat.New([][]float64{{0.6, 0.8}})
	result := mat.Normalize(m)

	if !util.EqualMatrix(result, expected) {
		t.Errorf("Wrong result in vector normalization. Want: %s\nGot: %s", expected, result)
	}

	// Test case 2: Panic on zero length matrix normalization
	zeroVec, _ := mat.New([][]float64{{0, 0}})
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected normalization to panic on a zero-length vector, but it did not")
		}
	}()
	_ = mat.Normalize(zeroVec)

	fmt.Printf("Runtime: %v\n", time.Since(start))
}

func TestVectorDot(t *testing.T) {
	start := time.Now()

	// Test case 1: Valid vector dot product
	v1, _ := mat.New([][]float64{{1}, {2}, {3}})
	v2, _ := mat.New([][]float64{{4}, {5}, {6}})

	result := mat.VectorDot(v1, v2)
	expected := 32.0

	if result != expected {
		t.Errorf("Expected vector dot product to be %f, Got: %f", expected, result)
	}

	// Test case 2: Panic on mismatched dimensionalities
	v3, _ := mat.New([][]float64{{1}, {2}, {3}})
	v4, _ := mat.New([][]float64{{4}, {5}, {6}, {7}})
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected VectorDot to panic on dimension mismatch, but it did not")
		}
	}()
	_ = mat.VectorDot(v3, v4)

	fmt.Printf("Runtime: %v\n", time.Since(start))
}

func TestVectorCrossProduct(t *testing.T) {
	start := time.Now()

	// Test case 1: Valid cross product
	v1, _ := mat.New([][]float64{{1}, {2}, {3}})
	v2, _ := mat.New([][]float64{{4}, {5}, {6}})

	result := mat.CrossProduct(v1, v2)
	expected, _ := mat.New([][]float64{{-3}, {6}, {-3}})

	if !util.EqualMatrix(result, expected) {
		t.Errorf("Wrong result in cross product. Want: %s\nGot: %s", expected, result)
	}

	// Test case 2: Panic on invalid dimensionalities (must be 3D)
	v3, _ := mat.New([][]float64{{1}, {2}})
	v4, _ := mat.New([][]float64{{4}, {5}})
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected CrossProduct to panic on non-3D vectors, but it did not")
		}
	}()
	_ = mat.CrossProduct(v3, v4)

	fmt.Printf("Runtime: %v\n", time.Since(start))
}

func TestCosineSimilarity(t *testing.T) {
	start := time.Now()

	// Test case 1: Valid cosine similarity
	v1, _ := mat.New([][]float64{{1}, {2}, {3}})
	v2, _ := mat.New([][]float64{{-1}, {0}, {2}})

	expected := (1*-1 + 2*0 + 3*2) / (math.Sqrt(14) * math.Sqrt(5))
	result := mat.CosineSimilarity(v1, v2)

	if result != expected {
		t.Errorf("Expected cosine similarity to be %f, Got: %f", expected, result)
	}

	// Test case 2: Panic on mismatched shapes
	v3, _ := mat.New([][]float64{{1}, {2}, {3}})
	v4, _ := mat.New([][]float64{{-1}, {0}, {2}, {4}})
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected CosineSimilarity to panic on dimension mismatch, but it did not")
		}
	}()
	_ = mat.CosineSimilarity(v3, v4)

	fmt.Printf("Runtime: %v\n", time.Since(start))
}
