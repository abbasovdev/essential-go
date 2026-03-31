package main

import (
	"fmt"
	"strconv"
)

func main() {

	value, _ := strconv.Atoi("42")
	fmt.Println("value:", value)

	_, err := strconv.Atoi("abc")
	fmt.Println("error:", err)
}
