package main

import "fmt"

func swap(a, b string) (string, string) {
	return b, a
}

func minMax(numbers []int) (int, int) {
	min := numbers[0]
	max := numbers[0]
	for _, n := range numbers {
		if n < min {
			min = n
		}
		if n > max {
			max = n
		}
	}
	return min, max
}

func main() {
	first, second := swap("hello", "world")
	fmt.Printf("first: %s, second: %s\n", first, second)

	low, high := minMax([]int{3, 1, 4, 1, 5, 9})
	fmt.Printf("min: %d, max: %d\n", low, high)
}
