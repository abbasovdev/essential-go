<div align="center">
    <h1>16) Maps</h1>
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
            ┃ <a href="../../04_data_structures/02_slices/README.md"> ⬅&#xFE0E; Previous Topic</a>
            ┃ <a href="../../04_data_structures/04_project_part_3/README.md"> Next Topic ⮕&#xFE0F;</a>
        </div>
        <img width="100%"  src="../../../images/header.png" />
        <br />
    </div>
</div>

---


## Maps

A map stores key-value pairs.

Each key maps to a value. Keys must be unique.

### How to Declare a Map

↳ 1) Declare with Values

```go
ages := map[string]int{
    "Alice": 30,
    "Bob":   25,
}
```

Here is the breakdown of the code:
- `map` → keyword to declare a map
- `[string]` → type of the key
- `int` → type of the value
- `"Alice": 30` → a key-value pair

↳ 2) Declare with `make()`

```go
ages := make(map[string]int)
```

This creates an empty map. You can add elements later.

### Maps Operations

↳ 1) Add or Update Elements

```go
ages["Charlie"] = 35
ages["Alice"] = 31
```

If the key exists, the value is updated. If not, a new pair is added.

↳ 2) Access a Value

```go
fmt.Println(ages["Alice"]) // 31
```

↳ 3) Delete a Key

Use `delete()`:

```go
delete(ages, "Bob")
```

↳ 4) Check if a Key Exists

Use the comma-ok pattern:

```go
age, ok := ages["Alice"]
```

If the key exists, `ok` is `true` and `age` holds the value. If not, `ok` is `false` and `age` is the zero value.

### Example Code

There are some map operations in [main.go](main.go) file.

But they are incomplete.

### Fix the Code

Open [main.go](main.go) file and complete the code lines.

Then validate your code by running:
```bash
go run topic.go 16 validate
```

✅ If everything is ok, you should see:
```
OK!
```

## Conclusion

Now you know how to use maps in Go.

Continue to the next topic 👇

[Project ⮕&#xFE0F;](../../04_data_structures/04_project_part_3/README.md)
