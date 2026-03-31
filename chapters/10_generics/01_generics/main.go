package main

import "fmt"

func First[T any](items []T) T {
	return items[0]
}

func Contains[T comparable](items []T, target T) bool {
	for _, item := range items {
		if item == target {
			return true
		}
	}
	return false
}

func main() {

	nums := []int{10, 20, 30}
	fmt.Println(First(nums))

	words := []string{"go", "rust", "python"}
	fmt.Println(First(words))

	fmt.Println(Contains([]int{1, 2, 3}, 2))
	fmt.Println(Contains([]string{"go", "rust"}, "python"))
}
