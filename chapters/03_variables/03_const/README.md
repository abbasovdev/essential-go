<div align="center">
    <h1>12) Constants</h1>
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
            ┃ <a href="../../03_variables/02_types/README.md"> ⬅&#xFE0E; Previous Topic</a>
            ┃ <a href="../../03_variables/04_project_part_2/README.md"> Next Topic ⮕&#xFE0F;</a>
        </div>
        <img width="100%"  src="../../../images/header.png" />
        <br />
    </div>
</div>

---


## Constants

Constants are like variables, but their values cannot be changed after they are set.

Use `const` instead of `var` to declare a constant.

### How to Declare a Constant

↳ 1) Single Constant

```go
const pi float64 = 3.14
```

Here is the breakdown of the code:
- `const` → keyword to declare a constant
- `pi` → name of the constant
- `float64` → type of the constant
- `3.14` → value of the constant

You can also omit the type:

```go
const pi = 3.14
```

↳ 2) Multiple Constants

You can declare multiple constants using a `const` block:

```go
const (
    language = "Go"
    year     = 2009
)
```

### Variables vs Constants

Variables can be changed after they are set:

```go
var score int = 10
score = 20 // allowed
```

Constants cannot be changed:

```go
const maxScore int = 100
maxScore = 200 // error: cannot assign to maxScore
```

### Example Code

There are some constant declarations in [main.go](main.go) file.

But they are incomplete.

### Fix the Code

Open [main.go](main.go) file and complete the code lines.

Then validate your code by running:
```bash
go run topic.go 12 validate
```

✅ If everything is ok, you should see:
```
OK!
```

## Conclusion

Now you know how to use constants in Go.

Continue to the next topic 👇

[Project ⮕&#xFE0F;](../../03_variables/04_project_part_2/README.md)
