<div align="center">
    <h1>32) Methods</h1>
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
            ┃ <a href="../../08_structs/01_structs/README.md"> ⬅&#xFE0E; Previous Topic</a>
            ┃ <a href="../../08_structs/03_exported_fields/README.md"> Next Topic ⮕&#xFE0F;</a>
        </div>
        <img width="100%"  src="../../../images/header.png" />
        <br />
    </div>
</div>

---


## Methods

A method is a function that belongs to a type.

It has a special parameter called a receiver that ties it to that type.

### Value Receiver

A value receiver gets a copy of the struct. It can read fields but cannot modify the original.

```go
func (c City) Summary() string {
    return fmt.Sprintf("%s: %d people", c.Name, c.Population)
}
```

Here is the breakdown of the code:
- `(c City)` → the receiver, `c` is the variable name and `City` is the type
- `Summary()` → the method name
- `string` → the return type

### Calling a Method

Call a method using dot notation on an instance:

```go
fmt.Println(city.Summary())
```

### Pointer Receiver

A pointer receiver gets a reference to the original struct. It can modify the fields.

```go
func (c *City) Grow(amount int) {
    c.Population += amount
}
```

`*City` means the receiver is a pointer. This is the same `*` concept from the Pointers chapter.

### When to Use Which

Use a value receiver when the method only reads data.

Use a pointer receiver when the method needs to modify the struct.

```
  Value Receiver (c City)          Pointer Receiver (c *City)

  city{Pop: 1000}                  city{Pop: 1000}
       |                                |
       | copy                           | reference
       v                                v
  c{Pop: 1000}                    c ──> city{Pop: 1000}
  c.Pop = 2000                    c.Pop = 2000
       |                                |
  city{Pop: 1000} (unchanged)     city{Pop: 2000} (modified!)
```

### Example Code

There are some method declarations in [main.go](main.go) file.

But they are incomplete.

### Fix the Code

Open [main.go](main.go) file and complete the code lines.

Then validate your code by running:
```bash
go run topic.go 32 validate
```

✅ If everything is ok, you should see:
```
OK!
```

## Conclusion

Now you know how to use methods in Go.

Continue to the next topic 👇

[Exported Fields ⮕&#xFE0F;](../../08_structs/03_exported_fields/README.md)
