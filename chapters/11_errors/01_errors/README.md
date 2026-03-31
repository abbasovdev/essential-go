<div align="center">
    <h1>41) Errors</h1>
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
            ┃ <a href="../../10_generics/03_project_part_9/README.md"> ⬅&#xFE0E; Previous Topic</a>
            ┃ <a href="../02_error_handling/README.md"> Next Topic ⮕&#xFE0F;</a>
        </div>
        <img width="100%"  src="../../../images/header.png" />
        <br />
    </div>
</div>

---


## Errors

Many languages use exceptions and try/catch blocks to handle errors.

Go does not. In Go, errors are just values.

### The error Type

Go has a built-in type called `error`.

A function that can fail returns an `error` as its last return value.

If nothing went wrong, the error is `nil`.

### Creating Errors

There are two common ways to create an error:

↳ 1) `errors.New`

```go
import "errors"

err := errors.New("something went wrong")
```

This creates a simple error with a message.

↳ 2) `fmt.Errorf`

```go
import "fmt"

name := "config.yaml"
err := fmt.Errorf("file not found: %s", name)
```

This creates an error with formatted text, just like `fmt.Sprintf`.

### Example Code

There are some error declarations in [main.go](main.go) file.

But they are incomplete.

### Fix the Code

Open [main.go](main.go) file and complete the code lines.

Then validate your code by running:
```bash
go run topic.go 41 validate
```

✅ If everything is ok, you should see:
```
OK!
```

## Conclusion

Now you know what errors are in Go.

Continue to the next topic 👇

[Error Handling ⮕&#xFE0F;](../02_error_handling/README.md)
