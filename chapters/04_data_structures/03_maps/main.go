package main

import "fmt"

func main() {

	scores := map[string]int{
		"Alice": 90,
		"Bob":   85,
	}

	scores[""] = 92

	(scores, "Bob")

	score, exists := scores["Alice"]

	fmt.Printf("Scores: %v\n", scores)
	fmt.Printf("Alice's score: %d\n", )
	fmt.Printf("Alice exists: %t\n", exists)
	fmt.Printf("Total entries: %d\n", len(scores))
}
