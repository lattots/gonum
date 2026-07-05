package mat

import (
	"fmt"
	"math/bits"
	"runtime"
	"sync"

	"github.com/lattots/gonum/number"
)

const naiveThreshold = 128

func Dot[T number.Num](m1, m2 *Mat[T]) (*Mat[T], error) {
	if m1.N != m2.M {
		return nil, fmt.Errorf("cannot multiply matrices: Number of columns in the first matrix (%d) must be equal to the number of rows in the second matrix (%d)", m1.N, m2.M)
	}

	// If any of the dimensions are smaller than the threshold, there is likely no benefit
	// to using the strassen dot product algorithm.
	if min(m1.M, m1.N, m2.M, m2.N) < naiveThreshold {
		return dotNaive(m1, m2), nil
	}

	sq1, sq2 := Square(m1, m2)

	paddedRes := dotStrassen(sq1, sq2)

	data := make([]T, m1.M*m2.N)

	for i := 0; i < m1.M; i++ {
		// Row starting index in the padded slice
		paddedStart := i * paddedRes.N
		// Row ending index in the padded slice
		paddedEnd := paddedStart + m2.N

		// Row starting index in the result slice
		resStart := i * m2.N

		copy(data[resStart:resStart+m2.N], paddedRes.Data[paddedStart:paddedEnd])
	}

	return &Mat[T]{
		M:    m1.M,
		N:    m2.N,
		Data: data,
	}, nil
}

func Mul[T number.Num](m1, m2 *Mat[T]) (*Mat[T], error) {
	if m1.M != m2.M {
		return nil, fmt.Errorf("number of rows must be equal in both matrices when performing element wise multiplication. %d != %d", m1.M, m2.M)
	}
	if m1.N != m2.N {
		return nil, fmt.Errorf("number of columns must be equal in both matrices when performing element wise multiplication. %d != %d", m1.N, m2.N)
	}

	data := make([]T, len(m1.Data))
	for i := range m1.Data {
		data[i] = m1.Data[i] * m2.Data[i]
	}

	return &Mat[T]{
		M:    m1.M,
		N:    m1.N,
		Data: data,
	}, nil
}

// dotNaive calculates the dot product of matrices m1 and m2.
// It expects the input matrices to have compatible shapes.
func dotNaive[T number.Num](m1, m2 *Mat[T]) *Mat[T] {
	result, _ := Zeros[T](m1.M, m2.N)

	numWorkers := min(runtime.GOMAXPROCS(0), m1.M)

	var wg sync.WaitGroup
	wg.Add(numWorkers)

	rowsPerWorker := m1.M / numWorkers

	for w := range numWorkers {
		startRow := w * rowsPerWorker
		endRow := (w + 1) * rowsPerWorker

		// Last worker handles all remaining rows
		if w == numWorkers-1 {
			endRow = m1.M
		}

		go func(start, end int) {
			defer wg.Done()

			for i := start; i < end; i++ {
				for j := 0; j < m2.N; j++ {
					var sum T
					for k := 0; k < m1.N; k++ {
						m1Val := m1.Data[i*m1.N+k]
						m2Val := m2.Data[k*m2.N+j]
						sum += m1Val * m2Val
					}
					result.Data[i*result.N+j] = sum
				}
			}
		}(startRow, endRow)
	}

	wg.Wait()
	return result
}

func dotStrassen[T number.Num](m1, m2 *Mat[T]) *Mat[T] {
	// Base case: fall back to naive standard multiplication
	if m1.M <= naiveThreshold {
		return dotNaive(m1, m2)
	}

	n := m1.M / 2

	a11, a12, a21, a22 := Split(m1)
	b11, b12, b21, b22 := Split(m2)

	var wg sync.WaitGroup
	products := make([]*Mat[T], 7)

	wg.Add(7)

	calculateProduct := func(index int, op1, op2 *Mat[T]) {
		defer wg.Done()
		// Recursively call dotStrassen
		products[index] = dotStrassen(op1, op2)
	}

	go calculateProduct(0, a11, Subtract(b12, b22))           // p1
	go calculateProduct(1, Sum(a11, a12), b22)                // p2
	go calculateProduct(2, Sum(a21, a22), b11)                // p3
	go calculateProduct(3, a22, Subtract(b21, b11))           // p4
	go calculateProduct(4, Sum(a11, a22), Sum(b11, b22))      // p5
	go calculateProduct(5, Subtract(a12, a22), Sum(b21, b22)) // p6
	go calculateProduct(6, Subtract(a11, a21), Sum(b11, b12)) // p7

	wg.Wait()

	// c11 = p5 + p4 - p2 + p6
	c11 := Sum(Subtract(Sum(products[4], products[3]), products[1]), products[5])
	// c12 = p1 + p2
	c12 := Sum(products[0], products[1])
	// c21 = p3 + p4
	c21 := Sum(products[2], products[3])
	// c22 = p5 + p1 - p3 - p7
	c22 := Subtract(Subtract(Sum(products[4], products[0]), products[2]), products[6])

	return Combine(c11, c12, c21, c22, n)
}

func nextPowerOfTwo(n int) int {
	if n <= 1 {
		return 1
	}
	return 1 << bits.Len(uint(n-1))
}

// Square pads two matrices to the next power of 2.
func Square[T number.Num](m1, m2 *Mat[T]) (*Mat[T], *Mat[T]) {
	largest := m1.M
	if m1.N > largest {
		largest = m1.N
	}
	if m2.M > largest {
		largest = m2.M
	}
	if m2.N > largest {
		largest = m2.N
	}

	size := nextPowerOfTwo(largest)

	newM1, _ := Zeros[T](size, size)
	newM2, _ := Zeros[T](size, size)

	for i := 0; i < m1.M; i++ {
		start := i * m1.N
		end := start + m1.N
		newStart := i * size
		copy(newM1.Data[newStart:newStart+m1.N], m1.Data[start:end])
	}

	for i := 0; i < m2.M; i++ {
		start := i * m2.N
		end := start + m2.N
		newStart := i * size
		copy(newM2.Data[newStart:newStart+m2.N], m2.Data[start:end])
	}

	return newM1, newM2
}

// split divides an N x N matrix into four (N/2) x (N/2) sub-matrices.
// It expects the input matrix to be squared to the next power of two.
func Split[T number.Num](m *Mat[T]) (*Mat[T], *Mat[T], *Mat[T], *Mat[T]) {
	n := m.M / 2
	m11, _ := Zeros[T](n, n)
	m12, _ := Zeros[T](n, n)
	m21, _ := Zeros[T](n, n)
	m22, _ := Zeros[T](n, n)

	for i := 0; i < n; i++ {
		rowStart := i * m.N
		nextHalfRowStart := (i + n) * m.N

		// Top half (m11 and m12)
		copy(m11.Data[i*n:(i+1)*n], m.Data[rowStart:rowStart+n])
		copy(m12.Data[i*n:(i+1)*n], m.Data[rowStart+n:rowStart+(2*n)])

		// Bottom half (m21 and m22)
		copy(m21.Data[i*n:(i+1)*n], m.Data[nextHalfRowStart:nextHalfRowStart+n])
		copy(m22.Data[i*n:(i+1)*n], m.Data[nextHalfRowStart+n:nextHalfRowStart+(2*n)])
	}

	return m11, m12, m21, m22
}

// combine merges four (N) x (N) matrices into one (2N) x (2N) matrix.
func Combine[T number.Num](m11, m12, m21, m22 *Mat[T], n int) *Mat[T] {
	result, _ := Zeros[T](2*n, 2*n)

	for i := 0; i < n; i++ {
		resTopRowStart := i * result.N
		resBotRowStart := (i + n) * result.N

		// Top half (m11 and m12)
		copy(result.Data[resTopRowStart:resTopRowStart+n], m11.Data[i*n:(i+1)*n])
		copy(result.Data[resTopRowStart+n:resTopRowStart+(2*n)], m12.Data[i*n:(i+1)*n])

		// Bottom half (m21 and m22)
		copy(result.Data[resBotRowStart:resBotRowStart+n], m21.Data[i*n:(i+1)*n])
		copy(result.Data[resBotRowStart+n:resBotRowStart+(2*n)], m22.Data[i*n:(i+1)*n])
	}

	return result
}
