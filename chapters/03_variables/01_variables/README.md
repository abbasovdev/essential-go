<div align="center">
    <h1>10) Variables</h1>
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
            ┃ <a href="../../02_basics/04_project_part_1/README.md"> ⬅&#xFE0E; Previous Topic</a>
            ┃ <a href="../../03_variables/02_types/README.md"> Next Topic ⮕&#xFE0F;</a>
        </div>
        <img width="100%"  src="../../../images/header.png" />
        <br />
    </div>
</div>

---


## Variables

Variables are used to store data in a program.

They are like containers that hold values.

You can store any type of data like numbers, text and more in variables.

### How to Declare a Variable

There are couple of ways to declare a variable in Go:

↳ 1) Explicit Typing

You should specify the type of the variable explicitly.

For example:
```go
var age int = 20
```

Here is the breakdown of the code:
- `var` → keyword to declare a variable
- `int` → type of the variable, it used to store integer numbers
- `age` → name of the variable
- `20` → value of the variable

Another example:
```go
var profession string = "Software Engineer"
```

Here is the breakdown of the code:
- `var` → keyword to declare a variable
- `string` → type of the variable, it used to store text
- `profession` → name of the variable
- `Software Engineer` → value of the variable

↳ 2) Implicit Typing

You can omit the type of the variable and let the compiler infer it.

For example:
```go
var age = 20
var profession = "Software Engineer"
```

We did not specify the type of the variable, but the compiler will infer it from the value.

↳ 3) Short Declaration

You can use the `:=` operator to declare and initialize a variable in one step.

For example:

```go
age := 20
profession := "Software Engineer"
```

We just omit `var` keyword and the type of the variable.

↳ 4) Multiple Declarations

You can declare multiple variables in one line.

For example:

```go
var age, profession = 20, "Software Engineer"

book, author := "Harry Potter and the Philosopher's Stone", "J.K. Rowling"
```

The values of the variables are assigned in the order of the variables.

Meaning:

- `age` is assigned the value `20`
- `profession` is assigned the value `Software Engineer`
- `book` is assigned the value `Harry Potter and the Philosopher's Stone`
- `author` is assigned the value `J.K. Rowling`


### Example Code

There are some variables declarations in [main.go](main.go) file.

But they are incomplete.

### Fix the Code

Open [main.go](main.go) file and complete the code lines.

⚠️ Note: We will learn more about variable types in the next topic. For now just use the types we mentioned above.

Then validate your code by running:
```bash
go run topic.go 10 validate
```

✅ If everything is ok, you should see:
```
OK!
```

## Conclusion

Now you know how to use variables in Go.

Continue to the next topic 👇

[Types ⮕&#xFE0F;](../../03_variables/02_types/README.md)