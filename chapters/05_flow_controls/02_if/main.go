package main

import "fmt"

func main() {

	year := 2009
	if year > 2000 {
		fmt.Println("Go was created in the 21st century")
	}

	version := 1.23
	if version >= 1.21 {
		fmt.Println("Generics supported")
	} else {
		fmt.Println("No generics")
	}

	if length := len("Gopher"); length <= 6 {
		fmt.Printf("Gopher has %d letters (short word)\n", length)
	}
}
