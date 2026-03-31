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

	wordFrequency := make(map[string]int)

	for _, word := range words {
		isStop := false
		for _, sw := range stopWords {
			if word == sw {
				isStop = true
				break
			}
		}
		if isStop {
			continue
		}
		wordFrequency[word]++

		switch {
		case len(word) <= 3:
			fmt.Printf("  [short]  %s\n", word)
		case len(word) <= 6:
			fmt.Printf("  [medium] %s\n", word)
		default:
			fmt.Printf("  [long]   %s\n", word)
		}
	}

	fmt.Println("Word frequency (stop words excluded):")
	for w, count := range wordFrequency {
		fmt.Printf("  %s: %d\n", w, count)
	}
}
