<div align="center">
    <h1>17) Let's Build a Project - Part 3</h1>
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
            ┃ <a href="../../04_data_structures/03_maps/README.md"> ⬅&#xFE0E; Previous Topic</a>
            ┃ <a href="../../05_flow_controls/01_for/README.md"> Next Topic ⮕&#xFE0F;</a>
        </div>
        <img width="100%"  src="../../../images/header.png" />
        <br />
    </div>
</div>

---


## Let's Build a Project - Part 3

In this topic, you will apply what you learned so far:
- arrays
- slices
- maps

### Update your Grolyze project

Open the `main.go` file in your `grolyze` project.

↳ 1) Create a slice of words

```go
words := []string{"go", "is", "great", "go", "has", "great", "tooling", "go", "tooling", "is", "fast"}
```

↳ 2) Define a slice of stop words

Stop words are common words that are usually filtered out in text analysis.

```go
stopWords := []string{"the", "is", "a", "an", "and", "or", "of", "to", "in"}
```

↳ 3) Create a map to store word frequencies

```go
wordFrequency := map[string]int{
    "go":      3,
    "is":      2,
    "great":   2,
    "has":     1,
    "tooling": 2,
    "fast":    1,
}
```

These are hardcoded values for now. In a later chapter, you will count words from a real file.

↳ 4) Print the stop words and word frequencies

```go
fmt.Printf("Words: %v\n", words)
fmt.Printf("Stop words: %v\n", stopWords)
fmt.Printf("Word frequency: %v\n", wordFrequency)
fmt.Printf("Number of words: %d\n", len(words))
fmt.Printf("Number of stop words: %d\n", len(stopWords))
fmt.Printf("Number of unique words: %d\n", len(wordFrequency))
```

### Run the Program

```bash
go run main.go
```

✅ If everything is ok, you should see:
```
Grolyze v1.0.0
Target file: guide.md
Verbose: false
Words: [go is great go has great tooling go tooling is fast]
Stop words: [the is a an and or of to in]
Word frequency: map[fast:1 go:3 great:2 has:1 is:2 tooling:2]
Number of words: 11
Number of stop words: 9
Number of unique words: 6
```

## Conclusion

Now your Grolyze project can store stop words and word frequencies.

Continue to the next topic 👇

[For ⮕&#xFE0F;](../../05_flow_controls/01_for/README.md)
