package main

import "fmt"

const version = "1.0.0"

func main() {
	var targetFile string = "guide.md"
	var isVerbose = false

	fmt.Printf("Grolyze v%s\n", version)
	fmt.Printf("Target file: %s\n", targetFile)
	fmt.Printf("Verbose: %t\n", isVerbose)
}
