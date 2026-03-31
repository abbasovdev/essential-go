<div align="center">
    <h1>45) Let's Build a Project - Part 10</h1>
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
            ┃ <a href="../04_if_err/README.md"> ⬅&#xFE0E; Previous Topic</a>
            ┃ <a href="../../12_goroutines/01_goroutines/README.md"> Next Topic ⮕&#xFE0F;</a>
        </div>
        <img width="100%"  src="../../../images/header.png" />
        <br />
    </div>
</div>

---


## Let's Build a Project - Part 10

In this topic, you will apply what you learned so far:
- errors
- error handling
- the `if err != nil` pattern
- `os.ReadFile`

### Update your Grolyze project

Open the `main.go` file in your `grolyze` project.

Right now, your `main` function passes a hardcoded text string to `report.Analyze(text)`. Time to read from a real file instead.

↳ 1) Read the target file using `os.ReadFile`

Replace the hardcoded `text` variable with real file reading:

```go
data, err := os.ReadFile(report.TargetFile)
if err != nil {
    fmt.Println("Error:", err)
    return
}
text := string(data)
```

You will need to add `"os"` to your imports.

↳ 2) Pass the file content to `Analyze`

The rest stays the same. `report.Analyze(text)` now works with real file content:

```go
report.Analyze(text)
```

↳ 3) Create a test file

Create a `guide.md` file in your `grolyze` project directory:

```bash
echo "Go is simple and efficient" > guide.md
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
Total words: 3
Unique words: 3

--- JSON ---
{"file":"guide.md","total_words":3,"unique_words":3}

Analysis complete.
```

3 words because "is" and "and" are stop words, leaving "go", "simple", "efficient".

↳ 4) Test the error handling

Delete the `guide.md` file and run the program again:

```bash
rm guide.md
go run main.go
```

✅ You should see an error message like:
```
Grolyze v1.0.0
Error: open guide.md: no such file or directory
```

Your program handles the missing file gracefully instead of crashing.

## Conclusion

Now your Grolyze project reads real files and handles errors. The only change was replacing a hardcoded string with `os.ReadFile`. All the struct, interface, and generic code stays untouched.

Continue to the next topic 👇

[Goroutines ⮕&#xFE0F;](../../12_goroutines/01_goroutines/README.md)
