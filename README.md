# Gonum

A lightweight, generic linear algebra library written in Go. Gonum provides a
simple matrix type with basic matrix and vector operations.

## Installation

```sh
go get github.com/lattots/gonum
```

Import the matrix package:

```go
import "github.com/lattots/gonum/mat"
```

## Quick start

Create matrices from two-dimensional slices with `mat.New`:

```go
package main

import (
    "fmt"
    "log"

    "github.com/lattots/gonum/mat"
)

func main() {
    a, err := mat.New([][]float64{
        {1, 2},
        {3, 4},
    })
    if err != nil {
        log.Fatal(err)
    }

    b, err := mat.New([][]float64{
        {5, 6},
        {7, 8},
    })
    if err != nil {
        log.Fatal(err)
    }

    product, err := mat.Dot(a, b)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println(product)
}
```

`mat.Dot` performs matrix multiplication. The example produces:

```text
2 x 2
19.00 22.00
43.00 50.00
```

## The `Mat` type

The central type is `mat.Mat[T]`, a generic matrix whose element type can be
`int`, a signed integer type, `float32`, or `float64`.

```go
type Mat[T number.Num] struct {
    M, N   int    // rows and columns
    Stride int    // distance between the starts of consecutive rows
    Data   []T    // backing data
}
```

Matrices created with `mat.New` are stored in a flat, row-major slice. For a
regular matrix, `Stride` is equal to `N`. The stride matters when working with
submatrix views, because a view can retain the row spacing of its parent matrix.

Use the constructor functions to create matrices:

```go
m, _ := mat.New([][]float64{{1, 2}, {3, 4}})
zeros, _ := mat.Zeros[float64](2, 3)
ones, _ := mat.Ones[float64](2, 3)
random, _ := mat.Random[float64](2, 3)
```

Read and update elements with `At` and `Set`:

```go
value := m.At(1, 2) // 2
m.Set(2, 1, 10)
```

**Indexing note:** `At`, `Set`, and `Slice` use **1-based** row and column
indices. `Slice` also uses inclusive end indices.

## Basic operations

### Transpose

`mat.Transpose` returns a new matrix containing the transpose:

```go
t := mat.Transpose(m)
// [[1, 3],
//  [2, 4]]
```

### Matrix multiplication

Use `mat.Dot` when the number of columns in the left matrix matches the number
of rows in the right matrix:

```go
product, err := mat.Dot(a, b)
if err != nil {
    log.Fatal(err)
}
```

### Element-wise multiplication

`mat.Mul` multiplies corresponding elements. Both matrices must have the same
dimensions:

```go
elementwise, err := mat.Mul(a, b)
if err != nil {
    log.Fatal(err)
}
// [[5, 12],
//  [21, 32]]
```

### Vectors

A vector is represented by a `Mat` with either one row or one column. The
package provides vector operations such as length and dot product:

```go
v1, _ := mat.New([][]float64{{1, 2, 3}})
v2, _ := mat.New([][]float64{{4, 5, 6}})

fmt.Println(v1.Length())          // Euclidean length
fmt.Println(mat.VectorDot(v1, v2)) // 32
```

## Working with matrix data

`Mat` exposes its backing `Data` slice, so you can work with the values directly
when needed. For ordinary matrices, elements are laid out row by row. When
iterating over rows of a matrix that may be a view, use `Stride` rather than
assuming each row starts at `i*N`.

`Slice` creates a **view** without copying the underlying data. Changes through
the view are visible in the original matrix:

```go
view := m.Slice(1, 2, 2, 2) // second column
view.Set(1, 1, 99)
// m.At(1, 2) is now 99
```
