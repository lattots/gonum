package mat

import (
	"testing"
)

type dotBenchShape struct {
	name string
	m    int
	k    int
	n    int
}

// Keep one value from each result observable so the compiler cannot
// discard the result of a benchmark operation.
var dotBenchSink float64

var dotBenchShapes = []dotBenchShape{
	{"64x64x64", 64, 64, 64},
	{"128x128x128", 128, 128, 128},
	{"256x256x256", 256, 256, 256},
	{"512x512x512", 512, 512, 512},
	{"1024x1024x1024", 1024, 1024, 1024},

	{"256x512x256", 256, 512, 256},
	{"256x256x512", 256, 256, 512},
}

func benchmarkDot(b *testing.B, usePublicDot bool) {
	b.Helper()

	for _, shape := range dotBenchShapes {
		b.Run(shape.name, func(b *testing.B) {
			a, err := Random[float64](shape.m, shape.k)
			if err != nil {
				b.Fatal(err)
			}

			c, err := Random[float64](shape.k, shape.n)
			if err != nil {
				b.Fatal(err)
			}

			b.ReportAllocs()

			for b.Loop() {
				if usePublicDot {
					result, err := Dot(a, c)
					if err != nil {
						b.Fatal(err)
					}

					dotBenchSink = result.Data[0]
				}
			}
		})
	}
}

func BenchmarkDotAuto(b *testing.B) {
	benchmarkDot(b, true)
}
