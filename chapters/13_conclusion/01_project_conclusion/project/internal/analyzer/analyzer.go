package analyzer

import "strings"

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
