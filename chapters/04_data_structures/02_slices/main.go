package main

import "fmt"

func main() {

	fruits := []string{"apple", "banana", "cherry"}

	fruits = append(fruits, "date")

	total := len(fruits)

	subset := fruits[1:3]

	fmt.Printf("Fruits: %v\n", fruits)
	fmt.Printf("Total: %d\n", total)
	fmt.Printf("Subset [1:3]: %v\n", subset)
	fmt.Printf("First fruit: %s\n", fruits[0])
}
