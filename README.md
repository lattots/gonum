# Gonum

## What is it?

Gonum is a linear algebra library written in native Go. It provides functionality for vector and matrix manipulation and calculation.

## How to use it?

You can use Gonum by importing your desired package to your own project

```bash
go get github.com/lattots/gonum
```

```go
package main

import (
	"fmt"
	"github.com/lattots/gonum/mat"
	"log"
)

func main() {
	m, err := mat.New([][]float64{
		{1, 2},
		{3, 4},
	})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(m)
}
```
