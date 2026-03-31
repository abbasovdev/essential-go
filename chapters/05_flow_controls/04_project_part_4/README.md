<div align="center">
    <h1>21) Let's Build a Project - Part 4</h1>
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
            ┃ <a href="../../05_flow_controls/03_switch/README.md"> ⬅&#xFE0E; Previous Topic</a>
            ┃ <a href="../../06_functions/01_functions/README.md"> Next Topic ⮕&#xFE0F;</a>
        </div>
        <img width="100%"  src="../../../images/header.png" />
        <br />
    </div>
</div>

---


## Let's Build a Project - Part 4

In this topic, you will apply what you learned so far:
- for loops
- if statements
- switch statements

### Update your Grolyze project

Open the `main.go` file in your `grolyze` project.

You should already have a `words` slice, a `stopWords` slice, and a `wordFrequency` map from the previous chapter.

Now you will replace the hardcoded `wordFrequency` map with a dynamically computed one.

↳ 1) Replace the hardcoded `wordFrequency` map with an empty one

Remove the hardcoded values and replace with:

```go
wordFrequency := make(map[string]int)
```

↳ 2) Use a for loop to iterate over the raw words

```go
for _, word := range words {
```

↳ 3) Use an if statement to skip stop words

Inside the loop, check if the current word is a stop word. If it is, skip it.

```go
    isStop := false
    for _, sw := range stopWords {
        if word == sw {
            isStop = true
        }
    }
    if isStop {
        continue
    }
```

↳ 4) Count the word if it is not a stop word

```go
    wordFrequency[word]++
```

↳ 5) Use a switch to categorize words by length

After counting, categorize the word:

```go
    switch {
    case len(word) <= 3:
        fmt.Printf("  [short]  %s\n", word)
    case len(word) <= 6:
        fmt.Printf("  [medium] %s\n", word)
    default:
        fmt.Printf("  [long]   %s\n", word)
    }
}
```

↳ 6) Print the word frequency map

After the loop, print the filtered word frequencies:

```go
fmt.Println("Word frequency (stop words excluded):")
for w, count := range wordFrequency {
    fmt.Printf("  %s: %d\n", w, count)
}
```

### Run the Program

```bash
go run main.go
```

Your exact output depends on the words you defined in previous chapters.

## Conclusion

Now your Grolyze project can iterate over words, filter out stop words, and categorize words by length.

Continue to the next topic 👇

[Functions ⮕&#xFE0F;](../../06_functions/01_functions/README.md)