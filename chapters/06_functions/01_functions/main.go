package main

import "fmt"

 sayHello() {
	fmt.Println("Hello from a function!")
}

func getYear()  {
	 2007
}

func main() {
	sayHello()

	year := getYear()
	fmt.Printf("Go development started in %d\n", year)
}
