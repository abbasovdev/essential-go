<div align="center">
    <h1>26) Defer</h1>
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
            ┃ <a href="../../06_functions/04_variadic/README.md"> ⬅&#xFE0E; Previous Topic</a>
            ┃ <a href="../../06_functions/06_project_part_5/README.md"> Next Topic ⮕&#xFE0F;</a>
        </div>
        <img width="100%"  src="../../../images/header.png" />
        <br />
    </div>
</div>

---


## Defer

The `defer` keyword delays the execution of a code block until the function returns.

### How Defer Works

A deferred function call runs after the function it is in finishes:

```go
func main() {
    defer fmt.Println("world")
    fmt.Println("hello")
}
```

Output:
```
hello
world
```

Even though `defer fmt.Println("world")` appears first, it runs last.

### Multiple Defers (LIFO)

When you use multiple `defer` statements, they execute in **Last In, First Out** (LIFO) order:

```go
func main() {
    defer fmt.Println("first")
    defer fmt.Println("second")
    defer fmt.Println("third")
}
```

Output:
```
third
second
first
```

The last deferred call runs first.

### When Is Defer Useful

Defer is commonly used to clean up resources:

```go
func readFile() {
    file := openFile("data.txt")
    defer closeFile(file)
    // process file...
}
```

The file will be closed when the function returns, no matter what happens in between.

### Example Code

There are some defer statements in [main.go](main.go) file.

But they are incomplete.

### Fix the Code

Open [main.go](main.go) file and complete the code lines.

Then validate your code by running:
```bash
go run topic.go 26 validate
```

✅ If everything is ok, you should see:
```
OK!
```

## Conclusion

Now you know how to use defer in Go.

Continue to the next topic 👇

[Project ⮕&#xFE0F;](../../06_functions/06_project_part_5/README.md)
