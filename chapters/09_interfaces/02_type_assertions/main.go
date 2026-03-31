package main

import "fmt"

func describe(i interface{}) {
	switch v := i.(type) {
	case string:
		fmt.Println("string:", v)
	case int:
		fmt.Println("int:", v)
	default:
		fmt.Println("unknown")
	}
}

func main() {

	describe("hello")
	describe(42)
	describe(3.14)

	var val interface{} = "Go"
	s, ok := val.(string)
	fmt.Println(s, ok)
}
