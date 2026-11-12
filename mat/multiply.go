package mat

import (
	"fmt"
	"runtime"
	"sync"

	"github.com/lattots/gonum/number"
)

const (
	blockSize        = 64 // Number of elements in each dimension
	minParBlockCount = 4  // Minimum number of blocks for parallel multiplication
)

func Dot[T number.Num](m1, m2 *Mat[T]) (*Mat[T], error) {
	if m1.N != m2.M {
		return nil, fmt.Errorf(
			"matrix dimensions mismatch: %d != %d",
			m1.N, m2.M,
		)
	}

	resData := make([]T, m1.M*m2.N)
	res := &Mat[T]{
		M:      m1.M,
		N:      m2.N,
		Stride: m2.N,
		Data:   resData,
	}

	err := dotParallel(m1, m2, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// dotSerialAccumulate computes res += m1 * m2.
func dotSerialAccumulate[T number.Num](m1, m2, res *Mat[T]) {
	for i := 0; i < res.M; i++ {
		resRow := res.Data[i*res.Stride : i*res.Stride+res.N]
		m1Row := m1.Data[i*m1.Stride : i*m1.Stride+m1.N]

		for m1i, m1v := range m1Row {
			if m1v == 0 {
				continue
			}

			m2Row := m2.Data[m1i*m2.Stride : m1i*m2.Stride+m2.N]

			for m2i, m2v := range m2Row {
				resRow[m2i] += m1v * m2v
			}
		}
	}
}

func dotSerial[T number.Num](m1, m2, res *Mat[T]) error {
	if m1.N != m2.M {
		return fmt.Errorf("matrix dimensions mismatch: %d != %d", m1.N, m2.M)
	}

	if res.M != m1.M || res.N != m2.N {
		return fmt.Errorf("result matrix has incorrect dimensions")
	}

	for i := 0; i < res.M; i++ {
		clear(res.Data[i*res.Stride : i*res.Stride+res.N])
	}

	dotSerialAccumulate(m1, m2, res)

	return nil
}

func dotParallel[T number.Num](m1, m2, res *Mat[T]) error {
	if m1.N != m2.M {
		return fmt.Errorf("matrix dimensions mismatch: %d != %d", m1.N, m2.M)
	}

	if res.M != m1.M || res.N != m2.N {
		return fmt.Errorf("result matrix has incorrect dimensions")
	}

	// Calculate the number of blocks in each output dimension.
	mBlocks := blockCount(res.M, blockSize)
	nBlocks := blockCount(res.N, blockSize)

	// Not enough output blocks to benefit from parallelization.
	if mBlocks*nBlocks < minParBlockCount {
		return dotSerial(m1, m2, res)
	}

	// Clear the result before accumulating partial products.
	for i := 0; i < res.M; i++ {
		clear(res.Data[i*res.Stride : i*res.Stride+res.N])
	}

	// Calculate the number of blocks in the shared K dimension.
	kBlocks := blockCount(m1.N, blockSize)

	// Limit the number of concurrently executing goroutines.
	workers := make(chan struct{}, runtime.GOMAXPROCS(0))

	var wg sync.WaitGroup

	// Iterate over output row blocks.
	for bi := range mBlocks {
		rowStart := bi*blockSize + 1
		rowEnd := min(rowStart+blockSize-1, res.M)

		// Iterate over output column blocks.
		for bj := range nBlocks {
			colStart := bj*blockSize + 1
			colEnd := min(colStart+blockSize-1, res.N)

			// Acquire a worker slot before starting a goroutine.
			workers <- struct{}{}
			wg.Add(1)

			go func(
				rowStart, rowEnd int,
				colStart, colEnd int,
			) {
				defer wg.Done()
				defer func() {
					<-workers
				}()

				// This goroutine exclusively owns this output tile.
				resBlock := res.Slice(
					rowStart, rowEnd,
					colStart, colEnd,
				)

				// Iterate over the K dimension sequentially.
				for bk := range kBlocks {
					kStart := bk*blockSize + 1
					kEnd := min(kStart+blockSize-1, m1.N)

					// Extract the corresponding input blocks.
					m1Block := m1.Slice(
						rowStart, rowEnd,
						kStart, kEnd,
					)

					m2Block := m2.Slice(
						kStart, kEnd,
						colStart, colEnd,
					)

					// Accumulate this partial product.
					dotSerialAccumulate(m1Block, m2Block, resBlock)
				}
			}(rowStart, rowEnd, colStart, colEnd)
		}
	}

	wg.Wait()

	return nil
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
		M:      m1.M,
		N:      m1.N,
		Stride: m1.Stride,
		Data:   data,
	}, nil
}

// blockCount calculates the number of blocks in each dimension with n number of elements.
func blockCount(n, blockSize int) int {
	return (n + blockSize - 1) / blockSize
}
