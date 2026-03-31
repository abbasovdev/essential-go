<div align="center">
    <h1>22) Functions</h1>
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
            ┃ <a href="../../05_flow_controls/04_project_part_4/README.md"> ⬅&#xFE0E; Previous Topic</a>
            ┃ <a href="../../06_functions/02_arguments/README.md"> Next Topic ⮕&#xFE0F;</a>
        </div>
        <img width="100%"  src="../../../images/header.png" />
        <br />
    </div>
</div>

---


## Functions

A function is a block of code that performs a specific task.

You have already been using one function: `main`. It is the entry point of every Go program.

You can create your own functions to organize and reuse code.

### How to Declare a Function

Use the `func` keyword to declare a function:

```go
func greet() {
    fmt.Println("Hello")
}
```

Here is the breakdown of the code:
- `func` → keyword to declare a function
- `greet` → name of the function
- `()` → parentheses for parameters (empty here)
- `{ }` → body of the function

### How to Call a Function

To execute a function, call it by its name followed by parentheses:

```go
func main() {
    greet()
}
```

### Functions with Return Values

A function can return a value using the `return` keyword.

You must specify the return type after the parentheses:

```go
func getLanguage() string {
    return "Go"
}
```

Here is the breakdown of the code:
- `string` → the return type
- `return "Go"` → the value returned to the caller

To use the returned value:

```go
language := getLanguage()
fmt.Println(language) // Go
```

### Example Code

There are some function declarations in [main.go](main.go) file.

But they are incomplete.

### Fix the Code

Open [main.go](main.go) file and complete the code lines.

Then validate your code by running:
```bash
go run topic.go 22 validate
```

✅ If everything is ok, you should see:
```
OK!
```

## Conclusion

Now you know how to declare and call functions in Go.

Continue to the next topic 👇

[Arguments ⮕&#xFE0F;](../../06_functions/02_arguments/README.md)
