<div align="center">
    <h1>29) More Pointers</h1>
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
            ┃ <a href="../../07_pointers/01_pointers/README.md"> ⬅&#xFE0E; Previous Topic</a>
            ┃ <a href="../../07_pointers/03_project_part_6/README.md"> Next Topic ⮕&#xFE0F;</a>
        </div>
        <img width="100%"  src="../../../images/header.png" />
        <br />
    </div>
</div>

---


## More Pointers

In Go, function arguments are passed by value by default.

This means the function gets a copy of the value, not the original.

```
  Pass by Value                Pass by Pointer

  main:                        main:
  score = 10                   score = 10
     |                            |
     | copy                       | address
     v                            v
  double:                      double:
  n = 10  (separate copy)     n = &score (points to original)
  n = 20  (copy changed)      *n = 20   (original changed)
     |                            |
  main:                        main:
  score = 10 (unchanged)      score = 20 (modified!)
```

### Pass by Value

```go
func double(n int) {
    n = n * 2
}

score := 10
double(score)
fmt.Println(score) // 10
```

`score` is still `10`. The function modified a copy, not the original.

### Pass by Pointer

To modify the original variable, pass a pointer:

```go
func double(n *int) {
    *n = *n * 2
}

score := 10
double(&score)
fmt.Println(score) // 20
```

`score` is now `20`. The function received the address and modified the value at that address.

Notice two things:
- The parameter type is `*int` (pointer to int)
- The caller passes `&score` (address of score)

### When to Use Pointers

↳ 1) When you need a function to modify the original value

↳ 2) When you want to avoid copying large data structures

### Example Code

There are some pointer operations in [main.go](main.go) file.

But they are incomplete.

### Fix the Code

Open [main.go](main.go) file and complete the code lines.

Then validate your code by running:
```bash
go run topic.go 29 validate
```

✅ If everything is ok, you should see:
```
OK!
```

## Conclusion

Now you know how to pass pointers to functions in Go.

Continue to the next topic 👇

[Project ⮕&#xFE0F;](../../07_pointers/03_project_part_6/README.md)
