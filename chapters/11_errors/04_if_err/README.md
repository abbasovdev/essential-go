<div align="center">
    <h1>44) If Err</h1>
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
            ┃ <a href="../03_skip/README.md"> ⬅&#xFE0E; Previous Topic</a>
            ┃ <a href="../05_project_part_10/README.md"> Next Topic ⮕&#xFE0F;</a>
        </div>
        <img width="100%"  src="../../../images/header.png" />
        <br />
    </div>
</div>

---


## If Err

Go has a common pattern for handling errors right where they happen.

You will see this pattern in almost every Go codebase.

### The Pattern

```go
if err != nil {
    // handle the error
}
```

This is the most common error handling pattern in Go.

### Short Variable Declaration in If

Go lets you declare a variable and check it in the same `if` statement:

```go
if value, err := strconv.Atoi("42"); err != nil {
    fmt.Println("error:", err)
} else {
    fmt.Println("converted:", value)
}
```

The variable declaration comes before the semicolon `;`.

The condition comes after the semicolon.

Both `value` and `err` are only available inside the `if` and `else` blocks.

### Reading Files

The `os.ReadFile` function reads a file and returns its content as bytes and an error:

```go
data, err := os.ReadFile("guide.md")
if err != nil {
    fmt.Println("error:", err)
    return
}
fmt.Println(string(data))
```

If the file does not exist, `err` will contain the error message.

If the file exists, `data` will contain the file content.

### Example Code

There is some code in [main.go](main.go) file.

But it is incomplete.

### Fix the Code

Open [main.go](main.go) file and complete the code lines.

Then validate your code by running:
```bash
go run topic.go 44 validate
```

✅ If everything is ok, you should see:
```
OK!
```

## Conclusion

Now you know the idiomatic way to handle errors in Go.

Continue to the next topic 👇

[Project ⮕&#xFE0F;](../05_project_part_10/README.md)
