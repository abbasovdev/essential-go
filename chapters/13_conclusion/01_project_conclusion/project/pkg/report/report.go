package report

import (
	"encoding/json"
	"fmt"
	"grolyze/internal/analyzer"
)

type Formatter interface {
	Format(report analyzer.AnalysisReport) string
}

type TextFormatter struct{}

func (f TextFormatter) Format(report analyzer.AnalysisReport) string {
	return fmt.Sprintf("File: %s\nTotal words: %d\nUnique words: %d",
		report.TargetFile, report.TotalWords, report.UniqueWords)
}

type JSONFormatter struct{}

func (f JSONFormatter) Format(report analyzer.AnalysisReport) string {
	data := map[string]interface{}{
		"file":         report.TargetFile,
		"total_words":  report.TotalWords,
		"unique_words": report.UniqueWords,
	}
	b, _ := json.Marshal(data)
	return string(b)
}

func PrintReport(f Formatter, r analyzer.AnalysisReport) {
	fmt.Println(f.Format(r))
}
