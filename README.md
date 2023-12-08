# Gonum

## What is it?

Gonum is a linear algebra library written in native Go. It provides functionality for vector and matrix manipulation and calculation.

## How to use it?

You can use Gonum by importing your desired package to your own project

```bash
go get github.com/lattots/gonum/pkg/pkg_name
```

```go
package main

import (
	"fmt"
	"github.com/lattots/gonum/pkg/matrix"
	"log"
)

func main() {
	m, err := matrix.NewMatrix([][]float64{
		{1, 2},
		{3, 4},
	})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(m)
}
```
