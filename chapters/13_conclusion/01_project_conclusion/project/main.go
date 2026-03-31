package main

import (
	"fmt"
	"grolyze/internal/analyzer"
	"grolyze/pkg/report"
	"os"
)

const version = "1.0.0"

func analyzeFile(path string, stopWords []string, results chan<- analyzer.AnalysisReport) {
	r := analyzer.AnalysisReport{
		TargetFile: path,
		StopWords:  stopWords,
	}
	data, err := os.ReadFile(path)
	if err != nil {
		fmt.Printf("Error reading %s: %v\n", path, err)
		results <- r
		return
	}
	r.Analyze(string(data))
	results <- r
}

func main() {
	defer fmt.Println("Analysis complete.")

	fmt.Printf("Grolyze v%s\n", version)

	files := os.Args[1:]

	if len(files) == 0 {
		fmt.Println("Usage: grolyze <file1> <file2> ...")
		return
	}

	stopWords := []string{"the", "is", "a", "an", "and", "or", "of", "to", "in"}
	results := make(chan analyzer.AnalysisReport, len(files))

	for _, f := range files {
		go analyzeFile(f, stopWords, results)
	}

	for range files {
		r := <-results
		report.PrintReport(report.TextFormatter{}, r)
	}
}
