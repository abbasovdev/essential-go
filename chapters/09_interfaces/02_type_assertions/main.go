package main

import "fmt"

func describe(i ) {
	switch v := i.() {
	case string:
		fmt.Println("string:", v)
	case :
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
	s,  := val.(string)
	fmt.Println(s, ok)
}
