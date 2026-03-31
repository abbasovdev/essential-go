<div align="center">
    <h1>18) For Loop</h1>
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
            ┃ <a href="../../04_data_structures/04_project_part_3/README.md"> ⬅&#xFE0E; Previous Topic</a>
            ┃ <a href="../../05_flow_controls/02_if/README.md"> Next Topic ⮕&#xFE0F;</a>
        </div>
        <img width="100%"  src="../../../images/header.png" />
        <br />
    </div>
</div>

---


## For Loop

Go has only one looping construct: the `for` loop.

It covers all use cases that other languages split across `for`, `while`, and `foreach`.

### How to Use a For Loop

There are several forms of the `for` loop in Go:

↳ 1) Basic For Loop

The classic three-component loop:

```go
for i := 0; i < 3; i++ {
    fmt.Println(i)
}
```

- `i := 0` → initialization
- `i < 3` → condition (loop runs while true)
- `i++` → post statement (runs after each iteration)

↳ 2) While-Style For Loop

Omit the initialization and post statement to get a while-style loop:

```go
count := 0
for count < 3 {
    fmt.Println(count)
    count++
}
```

Go does not have a `while` keyword. This form replaces it.

↳ 3) Range Over a Slice

Use `range` to iterate over slices:

```go
languages := []string{"Go", "Rust", "Zig"}
for index, value := range languages {
    fmt.Printf("%d: %s\n", index, value)
}
```

`range` returns two values on each iteration:
- `index` → the position
- `value` → the element at that position

If you do not need the index, use `_` to ignore it:

```go
for _, value := range languages {
    fmt.Println(value)
}
```

### Example Code

There are some for loops in [main.go](main.go) file.

But they are incomplete.

### Fix the Code

Open [main.go](main.go) file and complete the code lines.

Then validate your code by running:
```bash
go run topic.go 18 validate
```

✅ If everything is ok, you should see:
```
OK!
```

## Conclusion

Now you know how to use for loops in Go.

Continue to the next topic 👇

[If ⮕&#xFE0F;](../../05_flow_controls/02_if/README.md)