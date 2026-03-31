<div align="center">
    <h1>39) Generic Constraints</h1>
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
            ┃ <a href="../01_generics/README.md"> ⬅&#xFE0E; Previous Topic</a>
            ┃ <a href="../03_project_part_9/README.md"> Next Topic ⮕&#xFE0F;</a>
        </div>
        <img width="100%"  src="../../../images/header.png" />
        <br />
    </div>
</div>

---


## Generic Constraints

A constraint limits which types can be used with a generic function.

You already know `any` and `comparable`. You can also define your own.

### Custom Constraints

Define a constraint as an interface with a type union:

```go
type Number interface {
    int | float64
}
```

This means `Number` accepts only `int` or `float64`.

### Using a Custom Constraint

```go
func Sum[T Number](items []T) T {
    var total T
    for _, item := range items {
        total += item
    }
    return total
}
```

`Sum` works with `[]int` and `[]float64`, but not `[]string`.

### The ~ Operator

The `~` operator matches the underlying type:

```go
type Celsius float64

type Number interface {
    ~int | ~float64
}
```

Without `~`, `Celsius` would not satisfy `Number` even though its underlying type is `float64`. With `~`, it does.

### Example Code

There are some generic constraints in [main.go](main.go) file.

But they are incomplete.

### Fix the Code

Open [main.go](main.go) file and complete the code lines.

Then validate your code by running:
```bash
go run topic.go 39 validate
```

✅ If everything is ok, you should see:
```
OK!
```

## Conclusion

Now you know how to use generic constraints in Go.

Continue to the next topic 👇

[Project ⮕&#xFE0F;](../03_project_part_9/README.md)
