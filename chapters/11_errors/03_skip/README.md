<div align="center">
    <h1>43) Skip</h1>
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
            ┃ <a href="../02_error_handling/README.md"> ⬅&#xFE0E; Previous Topic</a>
            ┃ <a href="../04_if_err/README.md"> Next Topic ⮕&#xFE0F;</a>
        </div>
        <img width="100%"  src="../../../images/header.png" />
        <br />
    </div>
</div>

---


## Skip

Sometimes you do not need the error value.

Go lets you skip it using the blank identifier `_`.

### The Blank Identifier

The `_` tells Go: "I know this value exists, but I do not need it."

```go
value, _ := strconv.Atoi("42")
```

Here we convert a string to an integer. The second return value is an error, but we skip it with `_`.

### Why You Should Be Careful

Skipping errors can hide bugs.

If the conversion fails, you get a zero value with no warning:

```go
value, _ := strconv.Atoi("abc")
fmt.Println(value) // 0, no error shown
```

The program keeps running as if nothing happened.

### When Is It Acceptable

Skipping is acceptable when the error truly does not matter.

For example, `fmt.Println` returns an error, but almost no one checks it:

```go
fmt.Println("hello") // error is silently skipped
```

In most other cases, handle the error.

### Example Code

There is some code in [main.go](main.go) file.

But it is incomplete.

### Fix the Code

Open [main.go](main.go) file and complete the code lines.

Then validate your code by running:
```bash
go run topic.go 43 validate
```

✅ If everything is ok, you should see:
```
OK!
```

## Conclusion

Now you know how to skip errors in Go and when it is appropriate.

Continue to the next topic 👇

[If Err ⮕&#xFE0F;](../04_if_err/README.md)
