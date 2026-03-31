<div align="center">
    <h1>14) Arrays</h1>
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
            ┃ <a href="../../03_variables/04_project_part_2/README.md"> ⬅&#xFE0E; Previous Topic</a>
            ┃ <a href="../../04_data_structures/02_slices/README.md"> Next Topic ⮕&#xFE0F;</a>
        </div>
        <img width="100%"  src="../../../images/header.png" />
        <br />
    </div>
</div>

---


## Arrays

An array is a fixed-size collection of elements of the same type.

Once you declare an array with a size, the size cannot change.

### How to Declare an Array

There are couple of ways to declare an array in Go:

↳ 1) Declare with a Size

You specify the size inside square brackets:

```go
var cities [3]string
```

Here is the breakdown of the code:
- `var` → keyword to declare a variable
- `cities` → name of the array
- `[3]` → size of the array (3 elements)
- `string` → type of each element

The array starts with zero values. For `string`, that means three empty strings `""`.

↳ 2) Declare with Values

You can assign values at declaration:

```go
cities := [3]string{"Berlin", "Tokyo", "Istanbul"}
```

#### Fixed Size

Arrays have a fixed size. You cannot add or remove elements.

The size is part of the type. `[3]string` and `[4]string` are different types.

### Array Operations

↳ 1) Access Elements by Index

Array indices start at `0`:

```go
cities := [3]string{"Berlin", "Tokyo", "Istanbul"}
fmt.Println(cities[0]) // Berlin
fmt.Println(cities[1]) // Tokyo
fmt.Println(cities[2]) // Istanbul
```

↳ 2) Change an Element

You can update an element by its index:

```go
cities[1] = "Paris"
```

↳ 3) Get the Length

Use `len()` to get the number of elements:

```go
fmt.Println(len(cities)) // 3
```

### Example Code

There are some array declarations in [main.go](main.go) file.

But they are incomplete.

### Fix the Code

Open [main.go](main.go) file and complete the code lines.

Then validate your code by running:
```bash
go run topic.go 14 validate
```

✅ If everything is ok, you should see:
```
OK!
```

## Conclusion

Now you know how to use arrays in Go.

Continue to the next topic 👇

[Slices ⮕&#xFE0F;](../../04_data_structures/02_slices/README.md)
