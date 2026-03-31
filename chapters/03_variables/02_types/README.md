<div align="center">
    <h1>11) Types</h1>
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
            ┃ <a href="../../03_variables/01_variables/README.md"> ⬅&#xFE0E; Previous Topic</a>
            ┃ <a href="../../03_variables/03_const/README.md"> Next Topic ⮕&#xFE0F;</a>
        </div>
        <img width="100%"  src="../../../images/header.png" />
        <br />
    </div>
</div>

---


## Types

Every variable in Go has a type.

The type defines what kind of data the variable can hold.

### Basic Types

Go has four commonly used basic types:

↳ 1) `string` - stores text

```go
var city string = "Berlin"
```

↳ 2) `int` - stores whole numbers

```go
var population int = 3700000
```

↳ 3) `float64` - stores decimal numbers

```go
var area float64 = 891.7
```

↳ 4) `bool` - stores `true` or `false`

```go
var isCapital bool = true
```

### Zero Values

If you declare a variable without assigning a value, Go assigns a default value called a **zero value**.

```go
var name string    // ""
var age int        // 0
var score float64  // 0.0
var active bool    // false
```

Each type has its own zero value:
- `string` → `""`
- `int` → `0`
- `float64` → `0.0`
- `bool` → `false`

### Type Conversion

You can convert a value from one type to another.

To convert, use the type name as a function:

```go
year := 2007
yearAsFloat := float64(year)
```

`float64(year)` converts the `int` value `2007` to a `float64` value `2007.0`.

Another example:

```go
var score float64 = 9.8
scoreAsInt := int(score)
```

`int(score)` converts the `float64` value `9.8` to an `int` value `9`. The decimal part is dropped.

### Example Code

There are some type declarations in [main.go](main.go) file.

But they are incomplete.

### Fix the Code

Open [main.go](main.go) file and complete the code lines.

Then validate your code by running:
```bash
go run topic.go 11 validate
```

✅ If everything is ok, you should see:
```
OK!
```

## Conclusion

Now you know how to use types in Go.

Continue to the next topic 👇

[Const ⮕&#xFE0F;](../../03_variables/03_const/README.md)