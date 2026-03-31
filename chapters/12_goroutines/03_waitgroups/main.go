package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("goroutine done")
	}()

	wg.Wait()
	fmt.Println("all done")
}
