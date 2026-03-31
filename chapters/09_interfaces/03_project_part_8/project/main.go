package main

import (
	"encoding/json"
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

type Formatter interface {
	Format(report AnalysisReport) string
}

type TextFormatter struct{}

func (f TextFormatter) Format(report AnalysisReport) string {
	return fmt.Sprintf("File: %s\nTotal words: %d\nUnique words: %d",
		report.TargetFile, report.TotalWords, report.UniqueWords)
}

type JSONFormatter struct{}

func (f JSONFormatter) Format(report AnalysisReport) string {
	data := map[string]interface{}{
		"file":         report.TargetFile,
		"total_words":  report.TotalWords,
		"unique_words": report.UniqueWords,
	}
	b, _ := json.Marshal(data)
	return string(b)
}

func printReport(f Formatter, report AnalysisReport) {
	fmt.Println(f.Format(report))
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

	fmt.Println("--- Text ---")
	printReport(TextFormatter{}, report)

	fmt.Println("\n--- JSON ---")
	printReport(JSONFormatter{}, report)
}
