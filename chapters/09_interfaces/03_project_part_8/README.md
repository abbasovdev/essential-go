<div align="center">
    <h1>37) Let's Build a Project - Part 8</h1>
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
            ┃ <a href="../02_type_assertions/README.md"> ⬅&#xFE0E; Previous Topic</a>
            ┃ <a href="../../10_generics/01_generics/README.md"> Next Topic ⮕&#xFE0F;</a>
        </div>
        <img width="100%"  src="../../../images/header.png" />
        <br />
    </div>
</div>

---


## Let's Build a Project - Part 8

In this topic, you will apply what you learned so far:
- interfaces
- implicit implementation
- type assertions

### Update your Grolyze project

Open the `main.go` file in your `grolyze` project.

Your `AnalysisReport` struct has a `Print()` method that always prints in the same text format. Now you will use an interface to support multiple output formats.

↳ 1) Define a `Formatter` interface

Add this to your `main.go`:

```go
type Formatter interface {
    Format(report AnalysisReport) string
}
```

Any type that has a `Format` method with this signature satisfies the interface.

↳ 2) Implement a `TextFormatter`

```go
type TextFormatter struct{}

func (f TextFormatter) Format(report AnalysisReport) string {
    return fmt.Sprintf("File: %s\nTotal words: %d\nUnique words: %d",
        report.TargetFile, report.TotalWords, report.UniqueWords)
}
```

This produces the same output as the old `Print()` method.

↳ 3) Implement a `JSONFormatter`

```go
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
```

You will need to add `"encoding/json"` to your imports.

Both `TextFormatter` and `JSONFormatter` implement the `Formatter` interface. Go checks this automatically.

↳ 4) Create a `printReport` function that accepts a `Formatter`

Replace the `Print()` method on `AnalysisReport` with a standalone function:

```go
func printReport(f Formatter, report AnalysisReport) {
    fmt.Println(f.Format(report))
}
```

↳ 5) Update `main` to use both formatters

```go
fmt.Println("--- Text ---")
printReport(TextFormatter{}, report)

fmt.Println("--- JSON ---")
printReport(JSONFormatter{}, report)
```

### Run the Program

```bash
go run main.go
```

✅ If everything is ok, you should see:
```
Grolyze v1.0.0
--- Text ---
File: guide.md
Total words: 9
Unique words: 5

--- JSON ---
{"file":"guide.md","total_words":9,"unique_words":5}

Analysis complete.
```

The same `printReport` function produces different output depending on which `Formatter` you pass. That is the power of interfaces: you can add new formats (XML, CSV) without changing any existing code.

## Conclusion

Now your Grolyze project uses an interface to support multiple output formats.

Continue to the next topic 👇

[Generics ⮕&#xFE0F;](../../10_generics/01_generics/README.md)
