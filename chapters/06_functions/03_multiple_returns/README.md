<div align="center">
    <h1>24) Multiple Returns</h1>
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
            ┃ <a href="../../06_functions/02_arguments/README.md"> ⬅&#xFE0E; Previous Topic</a>
            ┃ <a href="../../06_functions/04_variadic/README.md"> Next Topic ⮕&#xFE0F;</a>
        </div>
        <img width="100%"  src="../../../images/header.png" />
        <br />
    </div>
</div>

---


## Multiple Returns

In Go, a function can return more than one value.

This is different from many other languages where a function returns only one value.

### How to Return Multiple Values

List the return types inside parentheses after the parameters:

```go
func swap(a, b string) (string, string) {
    return b, a
}
```

Here is the breakdown of the code:
- `(string, string)` → two return types
- `return b, a` → two values returned

To receive the returned values:

```go
first, second := swap("hello", "world")
```

### When Is This Useful

A common pattern in Go is returning a result and an error:

```go
func divide(a, b float64) (float64, string) {
    if b == 0 {
        return 0, "cannot divide by zero"
    }
    return a / b, ""
}
```

### Example Code

There are some functions with multiple returns in [main.go](main.go) file.

But they are incomplete.

### Fix the Code

Open [main.go](main.go) file and complete the code lines.

Then validate your code by running:
```bash
go run topic.go 24 validate
```

✅ If everything is ok, you should see:
```
OK!
```

## Conclusion

Now you know how to return multiple values from functions in Go.

Continue to the next topic 👇

[Variadic ⮕&#xFE0F;](../../06_functions/04_variadic/README.md)
