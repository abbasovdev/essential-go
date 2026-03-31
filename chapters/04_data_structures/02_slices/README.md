<div align="center">
    <h1>15) Slices</h1>
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
            ┃ <a href="../../04_data_structures/01_arrays/README.md"> ⬅&#xFE0E; Previous Topic</a>
            ┃ <a href="../../04_data_structures/03_maps/README.md"> Next Topic ⮕&#xFE0F;</a>
        </div>
        <img width="100%"  src="../../../images/header.png" />
        <br />
    </div>
</div>

---


## Slices

A slice is like an array, but its size can grow and shrink.

Slices are used much more often than arrays in Go.

```
  Array: [3]string — fixed size, cannot grow

  +--------+--------+----------+
  |"Berlin"| "Tokyo"|"Istanbul"|
  +--------+--------+----------+
    [0]       [1]       [2]       len = 3 (always 3)


  Slice: []string — can grow with append()

  +--------+--------+----------+---+---
  |"Berlin"| "Tokyo"|"Istanbul"|   | ...
  +--------+--------+----------+---+---
    [0]       [1]       [2]       len = 3, but can grow
```

### How to Declare a Slice

↳ 1) Declare with Values

No size inside the brackets:

```go
colors := []string{"red", "green", "blue"}
```

Notice the difference from arrays: `[]string` instead of `[3]string`.

↳ 2) Declare an Empty Slice

```go
var colors []string
```

This creates an empty slice with length `0`.

### Slices Operations

↳ 1) Access Elements by Index

Same as arrays:

```go
fmt.Println(colors[0]) // red
```

↳ 2) Add Elements with `append()`

Use `append()` to add elements to a slice:

```go
colors := []string{"red", "green"}
colors = append(colors, "blue")
```

`append()` returns a new slice with the added element. You must assign it back.

↳ 3) Get the Length

Use `len()`:

```go
fmt.Println(len(colors)) // 3
```

↳ 4) Slice a Slice

You can create a sub-slice using the `[start:end]` syntax:

```go
colors := []string{"red", "green", "blue", "yellow"}
subset := colors[1:3]
fmt.Println(subset) // [green blue]
```

`colors[1:3]` takes elements from index `1` up to (but not including) index `3`.

### Example Code

There are some slice operations in [main.go](main.go) file.

But they are incomplete.

### Fix the Code

Open [main.go](main.go) file and complete the code lines.

Then validate your code by running:
```bash
go run topic.go 15 validate
```

✅ If everything is ok, you should see:
```
OK!
```

## Conclusion

Now you know how to use slices in Go.

Continue to the next topic 👇

[Maps ⮕&#xFE0F;](../../04_data_structures/03_maps/README.md)
