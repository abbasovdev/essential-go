<div align="center">
    <h1>13) Let's Build a Project - Part 2</h1>
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
            ┃ <a href="../../03_variables/03_const/README.md"> ⬅&#xFE0E; Previous Topic</a>
            ┃ <a href="../../04_data_structures/01_arrays/README.md"> Next Topic ⮕&#xFE0F;</a>
        </div>
        <img width="100%"  src="../../../images/header.png" />
        <br />
    </div>
</div>

---


## Let's Build a Project - Part 2

In this topic, you will apply what you learned so far:
- variables
- types
- constants

### Update your Grolyze project

Open the `main.go` file in your `grolyze` project.

↳ 1) Define a constant for the tool version

Use `const` to define a version string:

```go
const version = "1.0.0"
```

↳ 2) Create a variable to hold the target file path

Use explicit typing:

```go
var targetFile string = "guide.md"
```

↳ 3) Create a variable to hold the verbose flag

Use implicit typing:

```go
var isVerbose = false
```

↳ 4) Print the tool info

```go
fmt.Printf("Grolyze v%s\n", version)
fmt.Printf("Target file: %s\n", targetFile)
fmt.Printf("Verbose: %t\n", isVerbose)
```

### Run the Program

```bash
go run main.go
```

✅ If everything is ok, you should see:
```
Grolyze v1.0.0
Target file: guide.md
Verbose: false
```

## Conclusion

Now your Grolyze project has a version constant and configuration variables.

Continue to the next topic 👇

[Arrays ⮕&#xFE0F;](../../04_data_structures/01_arrays/README.md)
