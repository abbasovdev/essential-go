package main

import (
	"errors"
	"fmt"
)

func divide(a, b float64) (float64, ) {
	if b == 0 {
		return 0, errors.New("cannot divide by zero")
	}
	return a / b,
}

func main() {

	result, err := divide(10, 0)
	if err !=  {
		fmt.Println("error:", err)
	}

	result, err = divide(10, 2)
	if err ==  {
		fmt.Println("result:", result)
	}
}
