package main

import "fmt"

func main() {
	ch := (chan string)

	go func() {
		ch  "hello from channel"
	}()

	msg :=  ch
	fmt.Println(msg)
}
