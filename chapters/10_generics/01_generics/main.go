package main

import "fmt"

func First[ any](items []T) T {
	return items[0]
}

func Contains[T ](items []T, target T) bool {
	for _, item := range items {
		if item ==  {
			return true
		}
	}
	return
}

func main() {

	nums := []int{10, 20, 30}
	fmt.Println(First())

	words := []string{"go", "rust", "python"}
	fmt.Println((words))

	fmt.Println(Contains([]int{1, 2, 3}, 2))
	fmt.Println(Contains([]string{"go", "rust"}, ""))
}
