<div align="center">
    <h1>38) Generics</h1>
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
            ┃ <a href="../../09_interfaces/03_project_part_8/README.md"> ⬅&#xFE0E; Previous Topic</a>
            ┃ <a href="../02_generic_constraints/README.md"> Next Topic ⮕&#xFE0F;</a>
        </div>
        <img width="100%"  src="../../../images/header.png" />
        <br />
    </div>
</div>

---


## Generics

Generics let you write functions and types that work with multiple types.

Instead of writing the same function for `int`, `string`, and `float64`, you write it once.

### Type Parameters

A generic function has type parameters in square brackets:

```go
func Print[T any](value T) {
    fmt.Println(value)
}
```

Here is the breakdown of the code:
- `[T any]` → `T` is a type parameter, `any` is the constraint
- `any` means any type is accepted
- `value T` → the parameter uses the type `T`

### Calling a Generic Function

Go infers the type from the argument:

```go
Print("hello")
Print(42)
```

### The comparable Constraint

Use `comparable` when you need `==` or `!=`:

```go
func Contains[T comparable](items []T, target T) bool {
    for _, item := range items {
        if item == target {
            return true
        }
    }
    return false
}
```

`comparable` allows types that support equality checks: `int`, `string`, `float64`, structs with comparable fields, etc.

### Example Code

There are some generic functions in [main.go](main.go) file.

But they are incomplete.

### Fix the Code

Open [main.go](main.go) file and complete the code lines.

Then validate your code by running:
```bash
go run topic.go 38 validate
```

✅ If everything is ok, you should see:
```
OK!
```

## Conclusion

Now you know how to use generics in Go.

Continue to the next topic 👇

[Generic Constraints ⮕&#xFE0F;](../02_generic_constraints/README.md)
