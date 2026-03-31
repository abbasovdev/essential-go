<div align="center">
    <h1>Project Conclusion</h1>
    <a class="header-badge" target="_blank" href="https://www.linkedin.com/in/abbasovdev/">
        <img alt="Linkedin Follow" src="https://img.shields.io/badge/style--5eba00.svg?label=LinkedIn&logo=linkedin&style=social">
    </a>
    <a class="header-badge" target="_blank" href="https://x.com/abbcyhn">
        <img alt="X Follow" src="https://img.shields.io/twitter/follow/abbcyhn?style=social">
    </a>
    <h2>Author:
        <a href="https://www.linkedin.com/in/abbasovdev/" target="_blank">Jeyhun Abbasov</a>
    </h2>
    <div>
        <span>Interactive Learning</span>
        <br />
        <div align="left">
            ☰ <a href="../../../README.md">🏠 Home</a>
            ┃ <a href="../../12_goroutines/04_project_part_11/README.md"> ⬅&#xFE0E; Previous Topic</a>
            ┃ <a href="../02_what_is_next/README.md"> Next Topic ⮕&#xFE0F;</a>
        </div>
        <img width="100%"  src="../../../images/header.png" />
        <br />
    </div>
</div>

---


## Let's Build a Project - Conclusion

In this topic, you will apply what you learned so far:
- modules
- packages
- project structure

## Modules

Before Go 1.11, all Go code had to live inside a single workspace called `GOPATH`.

This was limiting. You could not have multiple versions of the same dependency, and your project had to be in a specific directory.

Go modules solved this. A module is a collection of Go packages with a `go.mod` file at the root.

### go.mod

`go.mod` declares the module name, the Go version, and the dependencies.

It is created by running `go mod init <module-name>`.

Here is what a `go.mod` file looks like:

```
module grolyze

go 1.22

require github.com/some/package v1.2.3
```

- `module grolyze` → the name of the module, used as the import path prefix
- `go 1.22` → the minimum Go version required
- `require` → lists external dependencies and their versions

If your project has no external dependencies, the `require` block will not appear.

### go.sum

When you add a dependency, Go also creates a `go.sum` file.

`go.sum` contains the cryptographic checksums of each dependency.

It ensures that the exact same code is downloaded every time, even on a different machine.

You do not edit `go.sum` manually. Go manages it automatically.

Both `go.mod` and `go.sum` should be committed to version control.

### Initialize a Go Module

Your Grolyze project does not have a `go.mod` file yet.

Navigate to your `grolyze` project directory and run:

```bash
go mod init grolyze
```

This creates a `go.mod` file that defines your module name and Go version.

### Restructure the Project

Go projects use a standard directory layout.

Create the following directories:

```bash
mkdir -p internal/analyzer
mkdir -p pkg/report
```

- `internal/` contains packages that are only used within this project. Other projects cannot import from `internal/`.
- `pkg/` contains packages that can be reused by other projects.

### Move Code Into Packages

Right now, everything lives in `main.go`, including the `AnalysisReport` struct, `Formatter` interface, generic functions, and all methods. Time to organize.

↳ 1) Create `internal/analyzer/analyzer.go`

Move the `AnalysisReport` struct, its methods (`isStopWord`, `Analyze`), and the generic utility functions (`Filter`, `Keys`) into this package:

```go
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
```

↳ 2) Create `pkg/report/report.go`

Move the `Formatter` interface, `TextFormatter`, `JSONFormatter`, and `PrintReport` function into this package:

```go
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
```

↳ 3) Update `main.go`

Import and use the new packages. `main.go` becomes a thin orchestrator:

```go
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

	// os.Args[0] is the program name, os.Args[1:] are the arguments passed by the user
	// Example: go run . guide.md notes.md → os.Args[1:] = ["guide.md", "notes.md"]
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
```

### Final Directory Structure

Your project should look like this:

```
grolyze/
├── go.mod
├── main.go
├── internal/
│   └── analyzer/
│       └── analyzer.go
└── pkg/
    └── report/
        └── report.go
```

### Run the Program

Pass the files you want to analyze as arguments:

```bash
go run . guide.md notes.md
```

## Conclusion

Now your Grolyze project is structured as a proper Go module with internal and reusable packages.

Continue to the next topic 👇

[What is Next ⮕&#xFE0F;](../02_what_is_next/README.md)
