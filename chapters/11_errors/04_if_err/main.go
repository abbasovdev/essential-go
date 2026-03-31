package main

import (
	"fmt"
	"strconv"
)

func main() {

	if value, err := strconv.Atoi("42"); err != nil {
		fmt.Println("error:", err)
	} else {
		fmt.Println("converted:", value)
	}

	if _, err := strconv.Atoi("abc"); err != nil {
		fmt.Println("error:", err)
	}
}
