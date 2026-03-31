<div align="center">
    <h1>42) Error Handling</h1>
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
            ┃ <a href="../01_errors/README.md"> ⬅&#xFE0E; Previous Topic</a>
            ┃ <a href="../03_skip/README.md"> Next Topic ⮕&#xFE0F;</a>
        </div>
        <img width="100%"  src="../../../images/header.png" />
        <br />
    </div>
</div>

---


## Error Handling

In Go, functions that can fail return an `error` as the last return value.

The caller is responsible for checking the error.

### Returning Errors

A function returns a result and an error:

```go
func divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, errors.New("cannot divide by zero")
    }
    return a / b, nil
}
```

If the operation succeeds, the error is `nil`.

If it fails, the error contains a message.

### Checking Errors

After calling a function, you check the error:

```go
result, err := divide(10, 0)
if err != nil {
    fmt.Println("error:", err)
}
```

`err != nil` means something went wrong.

`err == nil` means everything is fine.

### Example Code

There is a `divide` function in [main.go](main.go) file.

But it is incomplete.

### Fix the Code

Open [main.go](main.go) file and complete the code lines.

Then validate your code by running:
```bash
go run topic.go 42 validate
```

✅ If everything is ok, you should see:
```
OK!
```

## Conclusion

Now you know how to handle errors in Go.

Continue to the next topic 👇

[Skip ⮕&#xFE0F;](../03_skip/README.md)
