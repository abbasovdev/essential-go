package main

import (
	"fmt"
	"time"
)

func sayHello() {
	fmt.Println("hello from goroutine")
}

func main() {
	go sayHello()

	time.Sleep(100 * time.Millisecond)

	fmt.Println("main done")
}
