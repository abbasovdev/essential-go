package main

import "fmt"

const version = "1.0.0"

func main() {
	var targetFile string = "guide.md"
	var isVerbose = false

	fmt.Printf("Grolyze v%s\n", version)
	fmt.Printf("Target file: %s\n", targetFile)
	fmt.Printf("Verbose: %t\n", isVerbose)

	words := []string{"go", "is", "great", "go", "has", "great", "tooling", "go", "tooling", "is", "fast"}
	stopWords := []string{"the", "is", "a", "an", "and", "or", "of", "to", "in"}
	wordFrequency := map[string]int{
		"go":      3,
		"is":      2,
		"great":   2,
		"has":     1,
		"tooling": 2,
		"fast":    1,
	}

	fmt.Printf("Words: %v\n", words)
	fmt.Printf("Stop words: %v\n", stopWords)
	fmt.Printf("Word frequency: %v\n", wordFrequency)
	fmt.Printf("Number of words: %d\n", len(words))
	fmt.Printf("Number of stop words: %d\n", len(stopWords))
	fmt.Printf("Number of unique words: %d\n", len(wordFrequency))
}
