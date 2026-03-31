package main

import (
	"fmt"
	"strings"
)

const version = "1.0.0"

type AnalysisReport struct {
	TargetFile  string
	TotalWords  int
	UniqueWords int
	WordFreq    map[string]int
	StopWords   []string
}

func (r *AnalysisReport) isStopWord(word string) bool {
	for _, sw := range r.StopWords {
		if word == sw {
			return true
		}
	}
	return false
}

func (r *AnalysisReport) Analyze(text string) {
	words := strings.Fields(text)
	r.WordFreq = make(map[string]int)
	r.TotalWords = 0
	for _, w := range words {
		w = strings.ToLower(w)
		if r.isStopWord(w) {
			continue
		}
		r.WordFreq[w]++
		r.TotalWords++
	}
	r.UniqueWords = len(r.WordFreq)
}

func (r AnalysisReport) Print() {
	fmt.Printf("File: %s\n", r.TargetFile)
	fmt.Printf("Total words: %d\n", r.TotalWords)
	fmt.Printf("Unique words: %d\n", r.UniqueWords)
	fmt.Println("Word frequency:")
	for w, count := range r.WordFreq {
		fmt.Printf("  %s: %d\n", w, count)
	}
}

func main() {
	defer fmt.Println("Analysis complete.")

	fmt.Printf("Grolyze v%s\n", version)

	report := AnalysisReport{
		TargetFile: "guide.md",
		StopWords:  []string{"the", "is", "a", "an", "and", "or", "of", "to", "in"},
	}

	text := "go is great go has great tooling go tooling is fast"
	report.Analyze(text)
	report.Print()
}
