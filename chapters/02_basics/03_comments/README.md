<div align="center">
    <h1>8) Comments</h1>
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
            ┃ <a href="../../02_basics/02_program_structure/README.md"> ⬅&#xFE0E; Previous Topic</a>
            ┃ <a href="../../02_basics/04_project_part_1/README.md"> Next Topic ⮕&#xFE0F;</a>
        </div>
        <img width="100%"  src="../../../images/header.png" />
        <br />
    </div>
</div>

---


## Comments

Comments are used to explain the code.

They are not executed as code, just ignored.

### Types of Comments

There are two types of comments in Go:

↳ 1) Single-line Comments

They start with `//` and continue until the end of the line. For example:

```go
// This is a single-line comment
```

↳ 2) Multi-line Comments

They start with `/*` and end with `*/`. For example:

```go
/* 
This is a multi-line comment 
with multiple lines
*/
```

### Example Code

There are some comments in [main.go](main.go) file.

Run the code:
```bash
go run topic.go 8
```

✅ If everything is ok, you should see:
```
Author's name is Jeyhun
Hello, World!
```

### Fix the Code

I do not want to print the author's name.

Let's do it by using comments.

How ?

You learned one use case of comments: explaining the code.

Another use case is to disable code.

Open [main.go](main.go) file and comment out the code lines.

Then validate your code by running:
```bash
go run topic.go 8 validate
```

✅ If everything is ok, you should see:
```
OK!
```

## Conclusion

Now you know how to use comments in Go.

Continue to the next topic 👇

[Project ⮕&#xFE0F;](../../02_basics/04_project_part_1/README.md)