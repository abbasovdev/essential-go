package main

import "fmt"

func greet(name ) {
	fmt.Printf("Hello, %s!\n", name)
}

func add(a , b int)  {
	return a + b
}

func main() {
	greet("Gopher")

	result := add(3, 7)
	fmt.Printf("3 + 7 = %d\n", result)
}
