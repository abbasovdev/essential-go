package main

import "fmt"

func main() {
	fmt.Println("start")

	defer fmt.Println("middle")

	fmt.Println("end")

	fmt.Println("---")

	defer fmt.Println("1")
	defer fmt.Println("2")
	defer fmt.Println("3")
}
