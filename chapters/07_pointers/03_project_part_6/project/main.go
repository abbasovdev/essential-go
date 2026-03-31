package main

import "fmt"

const version = "1.0.0"

func isStopWord(word string, stopWords []string) bool {
	for _, sw := range stopWords {
		if word == sw {
			return true
		}
	}
	return false
}

func countWords(words []string, stopWords []string, total *int) map[string]int {
	freq := make(map[string]int)
	*total = 0
	for _, word := range words {
		if isStopWord(word, stopWords) {
			continue
		}
		freq[word]++
		*total++
	}
	return freq
}

func main() {
	defer fmt.Println("Analysis complete.")

	fmt.Printf("Grolyze v%s\n", version)

	var targetFile string = "guide.md"
	var isVerbose = false
	fmt.Printf("Target file: %s\n", targetFile)
	fmt.Printf("Verbose: %t\n", isVerbose)

	words := []string{"go", "is", "great", "go", "has", "great", "tooling", "go", "tooling", "is", "fast"}
	stopWords := []string{"the", "is", "a", "an", "and", "or", "of", "to", "in"}

	var total int
	freq := countWords(words, stopWords, &total)

	fmt.Printf("Total words (excluding stop words): %d\n", total)
	fmt.Println("Word frequency (stop words excluded):")
	for w, count := range freq {
		fmt.Printf("  %s: %d\n", w, count)
	}
}
