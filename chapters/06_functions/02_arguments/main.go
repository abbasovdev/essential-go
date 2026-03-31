package main

import "fmt"

func greet(name string) {
	fmt.Printf("Hello, %s!\n", name)
}

func add(a int, b int) int {
	return a + b
}

func main() {
	greet("Gopher")

	result := add(3, 7)
	fmt.Printf("3 + 7 = %d\n", result)
}
