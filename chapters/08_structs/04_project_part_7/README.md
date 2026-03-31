<div align="center">
    <h1>34) Let's Build a Project - Part 7</h1>
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
            ┃ <a href="../../08_structs/03_exported_fields/README.md"> ⬅&#xFE0E; Previous Topic</a>
            ┃ <a href="../../09_interfaces/01_interfaces/README.md"> Next Topic ⮕&#xFE0F;</a>
        </div>
        <img width="100%"  src="../../../images/header.png" />
        <br />
    </div>
</div>

---


## Let's Build a Project - Part 7

In this topic, you will apply what you learned so far:
- structs
- methods
- exported fields

### Update your Grolyze project

Open the `main.go` file in your `grolyze` project.

Right now you have standalone functions (`isStopWord`, `countWords`) and separate variables (`words`, `stopWords`, `targetFile`). Time to wrap everything into a struct.

↳ 1) Define an `AnalysisReport` struct with exported fields

```go
type AnalysisReport struct {
    TargetFile  string
    TotalWords  int
    UniqueWords int
    WordFreq    map[string]int
    StopWords   []string
}
```

All fields start with an uppercase letter so they can be accessed from other packages later.

↳ 2) Turn `isStopWord` into a method

Move the standalone `isStopWord` function to a method on the struct. It can use the struct's `StopWords` field instead of taking it as a parameter:

```go
func (r *AnalysisReport) isStopWord(word string) bool {
    for _, sw := range r.StopWords {
        if word == sw {
            return true
        }
    }
    return false
}
```

↳ 3) Create an `Analyze` method

This replaces the standalone `countWords` function. It takes a text string, splits it into words using `strings.Fields`, filters stop words, and populates the struct fields:

```go
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
```

You will need to add `"strings"` to your imports.

Notice how the pointer receiver (`*AnalysisReport`) lets the method modify struct fields directly. This is the same concept as the `*int` pointer from the previous chapter, but more idiomatic.

↳ 4) Add a method to print the report

Use a value receiver because this method only reads data:

```go
func (r AnalysisReport) Print() {
    fmt.Printf("File: %s\n", r.TargetFile)
    fmt.Printf("Total words: %d\n", r.TotalWords)
    fmt.Printf("Unique words: %d\n", r.UniqueWords)
    fmt.Println("Word frequency:")
    for w, count := range r.WordFreq {
        fmt.Printf("  %s: %d\n", w, count)
    }
}
```

↳ 5) Update the main function

Remove the standalone `isStopWord` and `countWords` functions. Update `main` to use the struct:

```go
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
```

The hardcoded `words` slice from previous chapters is now a `text` string. The `Analyze` method handles splitting it into words internally.

### Run the Program

```bash
go run main.go
```

✅ If everything is ok, you should see:
```
Grolyze v1.0.0
File: guide.md
Total words: 9
Unique words: 5
Word frequency:
  go: 3
  great: 2
  has: 1
  tooling: 2
  fast: 1
Analysis complete.
```

## Conclusion

Now your Grolyze project uses a struct to organize the analysis logic.

Continue to the next topic 👇

[Interfaces ⮕&#xFE0F;](../../09_interfaces/01_interfaces/README.md)
