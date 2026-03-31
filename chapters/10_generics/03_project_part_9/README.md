<div align="center">
    <h1>40) Let's Build a Project - Part 9</h1>
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
            ┃ <a href="../02_generic_constraints/README.md"> ⬅&#xFE0E; Previous Topic</a>
            ┃ <a href="../../11_errors/01_errors/README.md"> Next Topic ⮕&#xFE0F;</a>
        </div>
        <img width="100%"  src="../../../images/header.png" />
        <br />
    </div>
</div>

---


## Let's Build a Project - Part 9

In this topic, you will apply what you learned so far:
- generics
- type parameters
- constraints

### Update your Grolyze project

Open the `main.go` file in your `grolyze` project.

You will add two generic utility functions and use them to refactor the stop word filtering inside `Analyze`.

↳ 1) Add a generic `Filter` function

```go
func Filter[T any](items []T, keep func(T) bool) []T {
    var result []T
    for _, item := range items {
        if keep(item) {
            result = append(result, item)
        }
    }
    return result
}
```

This works with any slice type. The predicate function decides which items to keep.

↳ 2) Add a generic `Keys` function

```go
func Keys[K comparable, V any](m map[K]V) []K {
    keys := make([]K, 0, len(m))
    for k := range m {
        keys = append(keys, k)
    }
    return keys
}
```

This extracts keys from any map, regardless of key and value types.

↳ 3) Refactor `Analyze` to use `Filter`

Replace the manual loop with stop word checks inside `Analyze`:

```go
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
```

The logic is the same as before, but the filtering step now uses your generic `Filter` function.

↳ 4) Use `Keys` in `main` to demonstrate generic map key extraction

After calling `report.Analyze(text)`, add:

```go
uniqueWords := Keys(report.WordFreq)
fmt.Printf("Unique word list: %v\n", uniqueWords)
```

### Run the Program

```bash
go run main.go
```

✅ If everything is ok, you should see the same analysis output as before, plus the unique word list.

The `Filter` and `Keys` functions are reusable with any type. That is the point of generics: write once, use with many types.

## Conclusion

Now your Grolyze project uses generics for reusable utility functions.

Continue to the next topic 👇

[Errors ⮕&#xFE0F;](../../11_errors/01_errors/README.md)
