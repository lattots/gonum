package mat_test

import (
	"testing"

	"github.com/lattots/gonum/internal/util"
	"github.com/lattots/gonum/mat"
)

func mustTestMatrix(t *testing.T, data [][]float64) *mat.Mat[float64] {
	t.Helper()

	m, err := mat.New(data)
	if err != nil {
		t.Fatalf("failed to create matrix: %v", err)
	}

	return m
}

func makeDotTestData(rows, cols, seed int) [][]float64 {
	data := make([][]float64, rows)

	for i := range data {
		data[i] = make([]float64, cols)

		for j := range data[i] {
			data[i][j] = float64(((i+1)*7+(j+1)*3+seed*5)%13 - 6)
		}
	}

	return data
}

// referenceDot is a simple, unoptimized matrix multiplication implementation.
func referenceDot(a, b [][]float64) [][]float64 {
	m := len(a)
	k := len(a[0])
	n := len(b[0])

	result := make([][]float64, m)

	for i := range m {
		result[i] = make([]float64, n)

		for j := range n {
			for l := range k {
				result[i][j] += a[i][l] * b[l][j]
			}
		}
	}

	return result
}

func assertMatrixEqual(t *testing.T, got *mat.Mat[float64], want [][]float64) {
	t.Helper()

	if got == nil {
		t.Fatal("got nil matrix")
	}

	if got.M != len(want) {
		t.Fatalf("wrong number of rows: got %d, want %d", got.M, len(want))
	}

	if len(want) == 0 {
		return
	}

	if got.N != len(want[0]) {
		t.Fatalf("wrong number of columns: got %d, want %d", got.N, len(want[0]))
	}

	for i := range want {
		for j := range want[i] {
			actual := got.At(i+1, j+1)
			expected := want[i][j]

			if actual != expected {
				t.Errorf(
					"wrong value at (%d, %d): got %v, want %v",
					i+1, j+1, actual, expected,
				)
			}
		}
	}
}

func TestDot(t *testing.T) {
	tests := []struct {
		name    string
		m, k, n int
	}{
		{
			name: "small rectangular",
			m:    2,
			k:    5,
			n:    3,
		},
		{
			name: "single row",
			m:    1,
			k:    7,
			n:    4,
		},
		{
			name: "single column",
			m:    4,
			k:    7,
			n:    1,
		},
		{
			name: "below block size",
			m:    63,
			k:    65,
			n:    63,
		},
		{
			name: "exactly block size",
			m:    64,
			k:    64,
			n:    64,
		},
		{
			name: "just above block size",
			m:    65,
			k:    65,
			n:    65,
		},
		{
			name: "parallel tall output at threshold",
			m:    256,
			k:    70,
			n:    40,
		},
		{
			name: "parallel wide output",
			m:    70,
			k:    65,
			n:    256,
		},
		{
			name: "multiple partial blocks",
			m:    130,
			k:    129,
			n:    131,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			aData := makeDotTestData(tt.m, tt.k, 0)
			bData := makeDotTestData(tt.k, tt.n, 1)

			a := mustTestMatrix(t, aData)
			b := mustTestMatrix(t, bData)

			want := referenceDot(aData, bData)

			got, err := mat.Dot(a, b)
			if err != nil {
				t.Fatalf("Dot returned an unexpected error: %v", err)
			}

			assertMatrixEqual(t, got, want)
		})
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
