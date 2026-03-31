<div align="center">
    <h1>33) Exported Fields</h1>
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
            ┃ <a href="../../08_structs/02_methods/README.md"> ⬅&#xFE0E; Previous Topic</a>
            ┃ <a href="../../08_structs/04_project_part_7/README.md"> Next Topic ⮕&#xFE0F;</a>
        </div>
        <img width="100%"  src="../../../images/header.png" />
        <br />
    </div>
</div>

---


## Exported Fields

In Go, the first letter of a name controls its visibility.

There are no `public` or `private` keywords.

### Exported vs Unexported

↳ 1) Exported — starts with an uppercase letter

Accessible from other packages.

```go
type User struct {
    Name  string  // exported
    Email string  // exported
}
```

↳ 2) Unexported — starts with a lowercase letter

Only accessible within the same package.

```go
type User struct {
    Name     string  // exported
    password string  // unexported
}
```

### Same Rule for Methods

Methods follow the same rule:

```go
func (u User) GetName() string { ... }  // exported
func (u User) validate() bool { ... }   // unexported
```

### Why It Matters

When you import a struct from another package, you can only access exported fields and methods.

This is how Go controls visibility across packages.

Since all exercises in this guide are single-file `package main` programs, both exported and unexported fields work within the same file. The distinction matters when code spans multiple packages.

### Example Code

There are some exported and unexported fields in [main.go](main.go) file.

But they are incomplete.

### Fix the Code

Open [main.go](main.go) file and complete the code lines.

Then validate your code by running:
```bash
go run topic.go 33 validate
```

✅ If everything is ok, you should see:
```
OK!
```

## Conclusion

Now you know how exported fields work in Go.

Continue to the next topic 👇

[Project ⮕&#xFE0F;](../../08_structs/04_project_part_7/README.md)
