package main

import "fmt"

func main() {

	const language = "Go"

	const (
		year     = 2009
		creator  string = "Robert Griesemer"
		maxScore = 100
	)

	fmt.Printf("%s was created in %d by %s. Max score: %d\n", language, year, creator, maxScore)
}
