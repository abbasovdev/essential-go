<div align="center">
    <h1>23) Arguments</h1>
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
            ┃ <a href="../../06_functions/01_functions/README.md"> ⬅&#xFE0E; Previous Topic</a>
            ┃ <a href="../../06_functions/03_multiple_returns/README.md"> Next Topic ⮕&#xFE0F;</a>
        </div>
        <img width="100%"  src="../../../images/header.png" />
        <br />
    </div>
</div>

---


## Arguments

Functions can accept input values called arguments.

Arguments allow you to pass data into a function when you call it.

### How to Define Parameters

Parameters are defined inside the parentheses of a function:

```go
func greet(name string) {
    fmt.Println("Hello,", name)
}
```

Here is the breakdown of the code:
- `name` → parameter name
- `string` → parameter type

To call the function:

```go
greet("Gopher")  // Hello, Gopher
```

### Multiple Parameters

A function can accept more than one parameter:

```go
func add(a int, b int) int {
    return a + b
}
```

When multiple parameters share the same type, you can shorten the declaration:

```go
func add(a, b int) int {
    return a + b
}
```

### Example Code

There are some function declarations with parameters in [main.go](main.go) file.

But they are incomplete.

### Fix the Code

Open [main.go](main.go) file and complete the code lines.

Then validate your code by running:
```bash
go run topic.go 23 validate
```

✅ If everything is ok, you should see:
```
OK!
```

## Conclusion

Now you know how to pass arguments to functions in Go.

Continue to the next topic 👇

[Multiple Returns ⮕&#xFE0F;](../../06_functions/03_multiple_returns/README.md)
