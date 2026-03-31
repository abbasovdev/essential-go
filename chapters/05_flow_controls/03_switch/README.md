<div align="center">
    <h1>20) Switch</h1>
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
            ┃ <a href="../../05_flow_controls/02_if/README.md"> ⬅&#xFE0E; Previous Topic</a>
            ┃ <a href="../../05_flow_controls/04_project_part_4/README.md"> Next Topic ⮕&#xFE0F;</a>
        </div>
        <img width="100%"  src="../../../images/header.png" />
        <br />
    </div>
</div>

---


## Switch

The `switch` statement selects one of many code blocks to execute.

Go's switch does not need `break` statements. It only runs the matched case.

### How to Use Switch

↳ 1) Basic Switch

```go
day := "Monday"
switch day {
case "Monday":
    fmt.Println("Start of the work week")
case "Friday":
    fmt.Println("Almost weekend")
}
```

↳ 2) Multiple Values in a Case

You can match multiple values in a single case:

```go
day := "Saturday"
switch day {
case "Saturday", "Sunday":
    fmt.Println("Weekend")
case "Monday", "Tuesday", "Wednesday", "Thursday", "Friday":
    fmt.Println("Weekday")
}
```

↳ 3) Default Case

The `default` case runs when no other case matches:

```go
language := "Rust"
switch language {
case "Go":
    fmt.Println("Compiled, garbage collected")
case "Rust":
    fmt.Println("Compiled, no garbage collector")
default:
    fmt.Println("Unknown language")
}
```

↳ 4) Conditionless Switch

You can write a switch without a value. Each case has its own condition:

```go
score := 85
switch {
case score >= 90:
    fmt.Println("A")
case score >= 80:
    fmt.Println("B")
default:
    fmt.Println("C")
}
```

This works like a chain of if-else statements.

### Example Code

There are some switch statements in [main.go](main.go) file.

But they are incomplete.

### Fix the Code

Open [main.go](main.go) file and complete the code lines.

Then validate your code by running:
```bash
go run topic.go 20 validate
```

✅ If everything is ok, you should see:
```
OK!
```

## Conclusion

Now you know how to use switch statements in Go.

Continue to the next topic 👇

[Project ⮕&#xFE0F;](../../05_flow_controls/04_project_part_4/README.md)