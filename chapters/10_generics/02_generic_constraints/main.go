package main

import "fmt"

type Number interface {
	int | float64
}

func Sum[T Number](items []T) T {
	var total T
	for _, item := range items {
		total += item
	}
	return total
}

func main() {

	ints := []int{1, 2, 3, 4, 5}
	fmt.Println(Sum(ints))

	floats := []float64{1.5, 2.5, 3.0}
	fmt.Println(Sum(floats))
}
