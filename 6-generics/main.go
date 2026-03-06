package main

import "fmt"

func sumInts(m map[string]int64) int64 {
	var s int64
	for _, v := range m {
		s += v
	}
	return s
}

func sumFloats(m map[string]float64) float64 {
	var s float64
	for _, v := range m {
		s += v
	}
	return s
}

func sumIntsAndFloats[K comparable, V int64 | float64](m map[K]V) V {
	var s V
	for _, v := range m {
		s += v
	}
	return s
}

type Number interface {
	int64 | float64
}

func sumNumbers[K comparable, V Number](m map[K]V) V {
	var s V
	for _, v := range m {
		s += v
	}
	return s
}

func main() {
	ints := map[string]int64{
		"first":  12,
		"second": 13,
	}

	floats := map[string]float64{
		"first":  24.52,
		"second": 25.47,
	}

	// Print sums by calling concrete type functions
	fmt.Printf("Non-Generic Sums Are: %v & %.2f\n", sumInts(ints), sumFloats(floats))

	// Print sums by calling generic function with type arguments
	fmt.Printf("Generic Sums Are: %v & %.2f\n",
		sumIntsAndFloats[string, int64](ints),
		sumIntsAndFloats[string, float64](floats),
	)

	// Print sums by calling generic function without type arguments 
	// Go infers types from the argument.
	fmt.Printf("Generic Sums, type parameters inferred: %v & %.2f\n",
		sumIntsAndFloats(ints),
		sumIntsAndFloats(floats),
	)

	// Print sums by calling generic function with type constraint
	fmt.Printf("Generic Sums with type constraint: %v & %.2f\n",
		sumNumbers(ints),
		sumNumbers(floats),
	)
}
