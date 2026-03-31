<div align="center">
    <h1>19) If</h1>
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
            ┃ <a href="../../05_flow_controls/01_for/README.md"> ⬅&#xFE0E; Previous Topic</a>
            ┃ <a href="../../05_flow_controls/03_switch/README.md"> Next Topic ⮕&#xFE0F;</a>
        </div>
        <img width="100%"  src="../../../images/header.png" />
        <br />
    </div>
</div>

---


## If

The `if` statement executes a block of code when a condition is true.

Go does not require parentheses around conditions, but braces are always required.

### How to Use If

↳ 1) Basic If

```go
age := 20
if age >= 18 {
    fmt.Println("adult")
}
```

↳ 2) If-Else

```go
score := 40
if score >= 50 {
    fmt.Println("pass")
} else {
    fmt.Println("fail")
}
```

↳ 3) Else-If Chain

```go
score := 85
if score >= 90 {
    fmt.Println("A")
} else if score >= 80 {
    fmt.Println("B")
} else if score >= 70 {
    fmt.Println("C")
} else {
    fmt.Println("F")
}
```

↳ 4) Short Statement If

You can declare a variable inside the `if` statement. The variable is scoped to the `if`/`else` block.

```go
if length := len("Go"); length < 5 {
    fmt.Printf("%d is short\n", length)
}
```

The variable `length` is only available inside the `if` and `else` blocks.

### Example Code

There are some if statements in [main.go](main.go) file.

But they are incomplete.

### Fix the Code

Open [main.go](main.go) file and complete the code lines.

Then validate your code by running:
```bash
go run topic.go 19 validate
```

✅ If everything is ok, you should see:
```
OK!
```

## Conclusion

Now you know how to use if statements in Go.

Continue to the next topic 👇

[Switch ⮕&#xFE0F;](../../05_flow_controls/03_switch/README.md)