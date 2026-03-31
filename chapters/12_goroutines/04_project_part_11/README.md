<div align="center">
    <h1>49) Let's Build a Project - Part 11</h1>
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
            ┃ <a href="../03_waitgroups/README.md"> ⬅&#xFE0E; Previous Topic</a>
            ┃ <a href="../../13_conclusion/01_project_conclusion/README.md"> Next Topic ⮕&#xFE0F;</a>
        </div>
        <img width="100%"  src="../../../images/header.png" />
        <br />
    </div>
</div>

---


## Let's Build a Project - Part 11

In this topic, you will apply what you learned so far:
- goroutines
- channels

### Update your Grolyze project

Open the `main.go` file in your `grolyze` project.

Right now, your program reads and analyzes a single file. You will extend it to process multiple files concurrently.

↳ 1) Create an `analyzeFile` function

This function wraps the per-file logic into a goroutine-friendly function. It creates an `AnalysisReport`, reads the file, runs the analysis, and sends the report through a channel:

```go
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
```

`chan<- AnalysisReport` means the channel is send-only inside this function. The channel carries `AnalysisReport` values, reusing the struct you already have.

↳ 2) Update `main` to process multiple files

Replace the single-file reading logic with a list of files and concurrent processing:

```go
files := []string{"guide.md", "notes.md"}
stopWords := []string{"the", "is", "a", "an", "and", "or", "of", "to", "in"}

results := make(chan AnalysisReport, len(files))

for _, f := range files {
    go analyzeFile(f, stopWords, results)
}
```

A buffered channel with capacity `len(files)` prevents goroutines from blocking on send.

↳ 3) Collect and print results

```go
for range files {
    report := <-results
    fmt.Println("--- Text ---")
    printReport(TextFormatter{}, report)
}
```

↳ 4) Create test files

```bash
echo "Go is simple and efficient" > guide.md
echo "Concurrency is powerful" > notes.md
```

### Run the Program

```bash
go run main.go
```

✅ If everything is ok, you should see output like:
```
Grolyze v1.0.0
--- Text ---
File: guide.md
Total words: 3
Unique words: 3

--- Text ---
File: notes.md
Total words: 2
Unique words: 2

Analysis complete.
```

The order of the two files may vary. Goroutines run concurrently.

## Conclusion

Now your Grolyze project processes multiple files concurrently using goroutines and channels. The `AnalysisReport` struct, `Formatter` interface, and generic functions all work together seamlessly.

Continue to the next topic 👇

[Project Conclusion ⮕&#xFE0F;](../../13_conclusion/01_project_conclusion/README.md)
