<div align="center">
    <h1>25) Variadic Functions</h1>
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
            ┃ <a href="../../06_functions/03_multiple_returns/README.md"> ⬅&#xFE0E; Previous Topic</a>
            ┃ <a href="../../06_functions/05_defer/README.md"> Next Topic ⮕&#xFE0F;</a>
        </div>
        <img width="100%"  src="../../../images/header.png" />
        <br />
    </div>
</div>

---


## Variadic Functions

A variadic function can accept any number of arguments.

You have already used one: `fmt.Println`.

### How to Declare a Variadic Function

Use `...` before the type of the last parameter:

```go
func sum(numbers ...int) int {
    total := 0
    for _, n := range numbers {
        total += n
    }
    return total
}
```

Here is the breakdown of the code:
- `...int` → accepts zero or more `int` arguments
- Inside the function, `numbers` is a slice of `int`

### How to Call a Variadic Function

You can pass any number of arguments:

```go
sum(1, 2)       // 3
sum(1, 2, 3, 4) // 10
sum()            // 0
```

### Passing a Slice to a Variadic Function

To pass an existing slice, use `...` after the slice:

```go
nums := []int{1, 2, 3}
sum(nums...)
```

### Example Code

There are some variadic function declarations in [main.go](main.go) file.

But they are incomplete.

### Fix the Code

Open [main.go](main.go) file and complete the code lines.

Then validate your code by running:
```bash
go run topic.go 25 validate
```

✅ If everything is ok, you should see:
```
OK!
```

## Conclusion

Now you know how to use variadic functions in Go.

Continue to the next topic 👇

[Defer ⮕&#xFE0F;](../../06_functions/05_defer/README.md)
