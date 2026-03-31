package main

import (
	"fmt"
	"strconv"
)

func main() {

	if value, err := strconv.Atoi("42");  != nil {
		fmt.Println("error:", err)
	} else {
		fmt.Println("converted:", )
	}

	if ,  := strconv.Atoi("abc"); err != nil {
		fmt.Println("error:", err)
	}
}
