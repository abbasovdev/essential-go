package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.

	wg.(1)
	go func() {
		defer wg.()
		fmt.Println("goroutine done")
	}()

	wg.()
	fmt.Println("all done")
}
