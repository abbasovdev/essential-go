package main

import (
	"encoding/json"
	"fmt"
	"os"
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

func Filter[T any](items []T, keep func(T) bool) []T {
	var result []T
	for _, item := range items {
		if keep(item) {
			result = append(result, item)
		}
	}
	return result
}

func Keys[K comparable, V any](m map[K]V) []K {
	keys := make([]K, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

func (r *AnalysisReport) Analyze(text string) {
	allWords := strings.Fields(text)

	filtered := Filter(allWords, func(w string) bool {
		return !r.isStopWord(strings.ToLower(w))
	})

	r.WordFreq = make(map[string]int)
	for _, w := range filtered {
		r.WordFreq[strings.ToLower(w)]++
	}
	r.TotalWords = len(filtered)
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

func analyzeFile(path string, stopWords []string, results chan<- AnalysisReport) {
	report := AnalysisReport{
		TargetFile: path,
		StopWords:  stopWords,
	}
	data, err := os.ReadFile(path)
	if err != nil {
		fmt.Printf("Error reading %s: %v\n", path, err)
		results <- report
		return
	}
	report.Analyze(string(data))
	results <- report
}

func main() {
	defer fmt.Println("Analysis complete.")

	fmt.Printf("Grolyze v%s\n", version)

	files := []string{"guide.md", "notes.md"}
	stopWords := []string{"the", "is", "a", "an", "and", "or", "of", "to", "in"}

	results := make(chan AnalysisReport, len(files))

	for _, f := range files {
		go analyzeFile(f, stopWords, results)
	}

	for range files {
		report := <-results
		fmt.Println("--- Text ---")
		printReport(TextFormatter{}, report)
	}
}
