package main

import "fmt"

func main() {

	var languages [3]string
	languages[0] = "Go"
	languages[1] = "Python"
	languages[2] = "Java"

	scores := [4]int{85, 92, 78, 95}

	count := len(languages)

	fmt.Printf("Languages: %v\n", languages)
	fmt.Printf("Scores: %v\n", scores)
	fmt.Printf("Number of languages: %d\n", count)
	fmt.Printf("First language: %s\n", languages[0])
	fmt.Printf("Last score: %d\n", scores[3])
}
