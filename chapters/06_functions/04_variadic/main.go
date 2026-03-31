package main

import "fmt"

func sum(numbers ...int) int {
	total := 0
	for _, n := range numbers {
		total += n
	}
	return total
}

func joinWords(sep string, words ...string) string {
	result := ""
	for i, w := range words {
		if i > 0 {
			result += sep
		}
		result += w
	}
	return result
}

func main() {
	fmt.Printf("sum: %d\n", sum(1, 2, 3, 4, 5))

	nums := []int{10, 20, 30}
	fmt.Printf("sum of slice: %d\n", sum(nums...))

	fmt.Printf("joined: %s\n", joinWords(" ", "Go", "is", "fun"))
}
