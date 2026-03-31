package main

import "fmt"

func main() {

	language := "Go"

	ptr = language

	fmt.Printf("Value: %s\n", language)
	fmt.Printf("Dereferenced: %s\n", ptr)

	ptr = "Golang"

	fmt.Printf("Updated value: %s\n", language)
}
