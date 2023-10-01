package main

import "fmt"

type Number interface {
	int64 | float64
}

func SumInts(m map[string]int64) int64 {
	var s int64
	for _, v := range m {
		s += v
	}
	return s
}

func SumFloats(m map[string]float64) float64 {
	var f float64
	for _, v := range m {
		f += v
	}
	return f
}

func SumIntsOrFloats[K comparable, V Number](m map[K]V) V {
	var s V
	for _, v := range m {
		s += v
	}

	return s
}

func main() {
	ints := map[string]int64{
		"first":  32,
		"second": 43,
	}

	floats := map[string]float64{
		"first":  36.49,
		"second": 40.123,
	}

	fmt.Printf("Non-generics sum: %v and %v", SumInts(ints), SumFloats(floats))
	fmt.Println()
	fmt.Printf("Generics sum: %v and %v", SumIntsOrFloats(ints), SumIntsOrFloats(floats))
}
