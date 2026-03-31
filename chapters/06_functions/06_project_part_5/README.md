<div align="center">
    <h1>27) Let's Build a Project - Part 5</h1>
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
            ┃ <a href="../../06_functions/05_defer/README.md"> ⬅&#xFE0E; Previous Topic</a>
            ┃ <a href="../../07_pointers/01_pointers/README.md"> Next Topic ⮕&#xFE0F;</a>
        </div>
        <img width="100%"  src="../../../images/header.png" />
        <br />
    </div>
</div>

---


## Let's Build a Project - Part 5

In this topic, you will apply what you learned so far:
- functions
- arguments
- multiple returns
- defer

### Update your Grolyze project

Open the `main.go` file in your `grolyze` project.

Right now, all your logic (iterating words, filtering stop words, counting) lives directly inside `main`. Time to extract it into functions.

↳ 1) Create an `isStopWord` function

Extract the stop word check from your loop into its own function:

```go
func isStopWord(word string, stopWords []string) bool {
    for _, sw := range stopWords {
        if word == sw {
            return true
        }
    }
    return false
}
```

↳ 2) Create a `countWords` function

Extract the counting loop from `main` into a function. It takes the same `words` and `stopWords` slices you already have, and returns the frequency map and total count:

```go
func countWords(words []string, stopWords []string) (map[string]int, int) {
    freq := make(map[string]int)
    total := 0
    for _, word := range words {
        if isStopWord(word, stopWords) {
            continue
        }
        freq[word]++
        total++
    }
    return freq, total
}
```

This is the same logic that was inline in `main`, now wrapped in a function.

↳ 3) Use `defer` to print a completion message

Add a defer statement at the top of `main`:

```go
defer fmt.Println("Analysis complete.")
```

↳ 4) Update `main` to call your new functions

Replace the inline loop and counting logic with a function call:

```go
freq, total := countWords(words, stopWords)

fmt.Printf("Total words (excluding stop words): %d\n", total)
fmt.Println("Word frequency (stop words excluded):")
for w, count := range freq {
    fmt.Printf("  %s: %d\n", w, count)
}
```

You can remove the switch categorizer since it served its purpose in the previous chapter.

### Run the Program

```bash
go run main.go
```

✅ If everything is ok, you should see output like:
```
Grolyze v1.0.0
Target file: guide.md
Verbose: false
Total words (excluding stop words): 9
Word frequency (stop words excluded):
  go: 3
  great: 2
  has: 1
  tooling: 2
  fast: 1
Analysis complete.
```

Notice that "Analysis complete." prints last because of `defer`.

The exact order of the word counts may vary (map iteration order is random in Go).

## Conclusion

Now your Grolyze project uses functions to organize logic. The same code that was inline in `main` is now reusable and testable.

Continue to the next topic 👇

[Pointers ⮕&#xFE0F;](../../07_pointers/01_pointers/README.md)
