<div align="center">
    <h1>36) Type Assertions</h1>
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
            ┃ <a href="../01_interfaces/README.md"> ⬅&#xFE0E; Previous Topic</a>
            ┃ <a href="../03_project_part_8/README.md"> Next Topic ⮕&#xFE0F;</a>
        </div>
        <img width="100%"  src="../../../images/header.png" />
        <br />
    </div>
</div>

---


## Type Assertions

A type assertion extracts the concrete type from an interface value.

### Basic Syntax

Use `value.(Type)` to assert a type:

```go
var i interface{} = "hello"

s := i.(string)
fmt.Println(s)
```

If the assertion is wrong, it panics.

### Two-Value Form

Use the two-value form to avoid panics:

```go
s, ok := i.(string)
fmt.Println(s, ok)
```

If `i` holds a `string`, `ok` is `true`. Otherwise, `ok` is `false` and `s` is the zero value.

### Type Switch

A type switch checks the type of an interface value against multiple types:

```go
switch v := i.(type) {
case string:
    fmt.Println("string:", v)
case int:
    fmt.Println("int:", v)
default:
    fmt.Println("unknown")
}
```

Here is the breakdown of the code:
- `i.(type)` → special syntax, only valid inside a switch
- `case string:` → matches if `i` holds a `string`
- `v` → the value, already cast to the matched type

### Example Code

There are some type assertions in [main.go](main.go) file.

But they are incomplete.

### Fix the Code

Open [main.go](main.go) file and complete the code lines.

Then validate your code by running:
```bash
go run topic.go 36 validate
```

✅ If everything is ok, you should see:
```
OK!
```

## Conclusion

Now you know how to use type assertions in Go.

Continue to the next topic 👇

[Project ⮕&#xFE0F;](../03_project_part_8/README.md)
