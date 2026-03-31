<div align="center">
    <h1>35) Interfaces</h1>
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
            ┃ <a href="../../08_structs/04_project_part_7/README.md"> ⬅&#xFE0E; Previous Topic</a>
            ┃ <a href="../02_type_assertions/README.md"> Next Topic ⮕&#xFE0F;</a>
        </div>
        <img width="100%"  src="../../../images/header.png" />
        <br />
    </div>
</div>

---


## Interfaces

An interface is a set of method signatures.

Any type that implements all the methods of an interface satisfies that interface.

There is no `implements` keyword. It is implicit.

### How to Define an Interface

Use the `type` keyword followed by the interface name and `interface`:

```go
type Shape interface {
    Area() float64
}
```

Here is the breakdown of the code:
- `type` → keyword to define a new type
- `Shape` → name of the interface
- `interface` → keyword that tells Go this is an interface
- `Area() float64` → a method signature that any type must implement

### How to Implement an Interface

A type implements an interface by defining all its methods:

```go
type Circle struct {
    Radius float64
}

func (c Circle) Area() float64 {
    return 3.14 * c.Radius * c.Radius
}
```

`Circle` now satisfies the `Shape` interface. No extra declaration needed.

### Using an Interface as a Parameter

A function can accept any type that satisfies an interface:

```go
func printArea(s Shape) {
    fmt.Println(s.Area())
}
```

You can pass a `Circle`, `Rectangle`, or any type that has an `Area() float64` method.

### The Empty Interface

The empty interface `interface{}` has no methods. Every type satisfies it.

Go provides `any` as an alias for `interface{}`:

```go
func printValue(v any) {
    fmt.Println(v)
}
```

### Example Code

There are some interface declarations in [main.go](main.go) file.

But they are incomplete.

### Fix the Code

Open [main.go](main.go) file and complete the code lines.

Then validate your code by running:
```bash
go run topic.go 35 validate
```

✅ If everything is ok, you should see:
```
OK!
```

## Conclusion

Now you know how to use interfaces in Go.

Continue to the next topic 👇

[Type Assertions ⮕&#xFE0F;](../02_type_assertions/README.md)
