<div align="center">
    <h1>31) Structs</h1>
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
            ┃ <a href="../../07_pointers/03_project_part_6/README.md"> ⬅&#xFE0E; Previous Topic</a>
            ┃ <a href="../../08_structs/02_methods/README.md"> Next Topic ⮕&#xFE0F;</a>
        </div>
        <img width="100%"  src="../../../images/header.png" />
        <br />
    </div>
</div>

---


## Structs

A struct is a collection of fields grouped together under one type.

It lets you combine different types of data into a single unit.

### How to Define a Struct

Use the `type` keyword followed by the struct name and `struct`:

```go
type City struct {
    Name       string
    Population int
    Area       float64
}
```

Here is the breakdown of the code:
- `type` → keyword to define a new type
- `City` → name of the struct
- `struct` → keyword that tells Go this is a struct
- `Name`, `Population`, `Area` → fields of the struct with their types

### How to Create an Instance

Use the struct name followed by field values:

```go
city := City{
    Name:       "Berlin",
    Population: 3700000,
    Area:       891.7,
}
```

### Accessing Fields

Use dot notation to read a field:

```go
fmt.Println(city.Name)
```

### Modifying Fields

Use dot notation to change a field:

```go
city.Population = 3800000
```

### Example Code

There are some struct declarations in [main.go](main.go) file.

But they are incomplete.

### Fix the Code

Open [main.go](main.go) file and complete the code lines.

Then validate your code by running:
```bash
go run topic.go 31 validate
```

✅ If everything is ok, you should see:
```
OK!
```

## Conclusion

Now you know how to use structs in Go.

Continue to the next topic 👇

[Methods ⮕&#xFE0F;](../../08_structs/02_methods/README.md)
