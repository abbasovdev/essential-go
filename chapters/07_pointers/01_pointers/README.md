<div align="center">
    <h1>28) Pointers</h1>
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
            ┃ <a href="../../06_functions/06_project_part_5/README.md"> ⬅&#xFE0E; Previous Topic</a>
            ┃ <a href="../../07_pointers/02_more_pointers/README.md"> Next Topic ⮕&#xFE0F;</a>
        </div>
        <img width="100%"  src="../../../images/header.png" />
        <br />
    </div>
</div>

---


## Pointers

A pointer stores the memory address of a variable.

Instead of holding a value directly, a pointer holds the location where that value lives.

```
    score             p
   +-------+       +-------+
   |  100  |       |&score |---+
   +-------+       +-------+   |
   0xc000  <--------------------+

   &score  → gives 0xc000 (the address)
   *p      → gives 100 (the value at the address)
```

### How Pointers Work

There are two key operators for pointers:

↳ 1) `&` — Address-of Operator

`&` gives you the memory address of a variable.

```go
score := 100
fmt.Println(&score) // prints something like 0xc0000b2008
```

`&score` returns the address where `score` is stored.

↳ 2) `*` — Dereference Operator

`*` reads the value stored at an address.

```go
score := 100
p := &score
fmt.Println(*p) // 100
```

`*p` follows the pointer and returns the value `100`.

### Declaring a Pointer Variable

You can declare a pointer variable using the `*Type` syntax:

```go
var p *int
```

This declares `p` as a pointer to an `int`. Its zero value is `nil`.

You can also use short declaration:

```go
score := 100
p := &score
```

Here `p` is of type `*int` and holds the address of `score`.

### Changing a Value Through a Pointer

You can modify the original variable through a pointer:

```go
score := 100
p := &score
*p = 200
fmt.Println(score) // 200
```

`*p = 200` changes the value at the address `p` points to. Since `p` points to `score`, `score` becomes `200`.

### Example Code

There are some pointer operations in [main.go](main.go) file.

But they are incomplete.

### Fix the Code

Open [main.go](main.go) file and complete the code lines.

Then validate your code by running:
```bash
go run topic.go 28 validate
```

✅ If everything is ok, you should see:
```
OK!
```

## Conclusion

Now you know how pointers work in Go.

Continue to the next topic 👇

[More Pointers ⮕&#xFE0F;](../../07_pointers/02_more_pointers/README.md)
