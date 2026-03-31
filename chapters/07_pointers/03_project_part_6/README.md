<div align="center">
    <h1>30) Let's Build a Project - Part 6</h1>
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
            ┃ <a href="../../07_pointers/02_more_pointers/README.md"> ⬅&#xFE0E; Previous Topic</a>
            ┃ <a href="../../08_structs/01_structs/README.md"> Next Topic ⮕&#xFE0F;</a>
        </div>
        <img width="100%"  src="../../../images/header.png" />
        <br />
    </div>
</div>

---


## Let's Build a Project - Part 6

In this topic, you will apply what you learned so far:
- pointers
- pass by value vs pass by pointer

### Update the Grolyze project

Open your `grolyze` project.

Your `countWords` function currently returns the total as a second return value:

```go
func countWords(words []string, stopWords []string) (map[string]int, int) {
```

You will refactor it to write the total through a pointer instead.

↳ 1) Change the `countWords` function signature to accept a pointer to the total

Before:
```go
func countWords(words []string, stopWords []string) (map[string]int, int) {
```

After:
```go
func countWords(words []string, stopWords []string, total *int) map[string]int {
```

↳ 2) Inside `countWords`, write to the total through the pointer

Replace `total := 0` and `return freq, total` with:

```go
*total = 0
```

And increment it through the pointer:

```go
*total++
```

The function now returns only the frequency map. The full function should look like:

```go
func countWords(words []string, stopWords []string, total *int) map[string]int {
    freq := make(map[string]int)
    *total = 0
    for _, word := range words {
        if isStopWord(word, stopWords) {
            continue
        }
        freq[word]++
        *total++
    }
    return freq
}
```

↳ 3) Update `main` to pass a pointer for the total

Before:
```go
freq, total := countWords(words, stopWords)
```

After:
```go
var total int
freq := countWords(words, stopWords, &total)
```

### Run the Program

```bash
go run main.go
```

✅ Your output should include the word frequency map and the total word count, same as before. The difference is that the data is now modified through pointers.

### Why Pointers Here ?

Returning an `int` as a second return value (like in Part 5) works fine. Using a `*int` pointer here is not strictly better.

But the goal is to practice using pointers. In the next chapter, you will use structs with pointer receivers, where pointers become much more useful.

## Conclusion

Now you know how to use pointers in a real project.

Continue to the next topic 👇

[Structs ⮕&#xFE0F;](../../08_structs/01_structs/README.md)
