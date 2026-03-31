package main

import "fmt"

func sayHello() {
	fmt.Println("Hello from a function!")
}

func getYear() int {
	return 2007
}

func main() {
	sayHello()

	year := getYear()
	fmt.Printf("Go development started in %d\n", year)
}
