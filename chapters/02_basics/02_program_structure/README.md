<div align="center">
    <h1>7) Program Structure</h1>
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
            ┃ <a href="../../02_basics/01_hello_world/README.md"> ⬅&#xFE0E; Previous Topic</a>
            ┃ <a href="../../02_basics/03_comments/README.md"> Next Topic ⮕&#xFE0F;</a>
        </div>
        <img width="100%"  src="../../../images/header.png" />
        <br />
    </div>
</div>

---


## Program Structure

You already wrote your first Go program:
```go
package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")
}
```

Let's analyze the each line of code.

### Package

The first statement is `package main`.

A package is a collection of Go files. It is like a folder in a file system.

Some rules for packages:
* each go file must start with a `package` declaration
* each go file can have only one `package`
* each `package` name should be:
    * lowercase
    * a single word, no underscores or mixedCaps

Lastly, `package` named `main` is special.

Will tell you at the end why it is special.

For now, let's move on to the next statement.

### Import

Next is the `import` statement.

`import` is used to import packages into your program.

Think of it as a way to use someone else's code in your program.

We imported the `fmt` (short for "format") package.

`fmt` is a built-in package, used to format text, print messages to the terminal, and read user input.

In our example, we used `fmt.Println` to print "Hello, World!" to the terminal.

## Func main()

Next is the `func main()` statement.

`func` is a keyword used to define a function.

We will talk a lot about functions in Chapter 6.

For now, just know that `func main()` is special function.

Together, `package main` and `func main()` are the entry point of the Go program.

Everything between the curly braces in `func main() {...}` will be executed when the program is run.

## Conclusion

Now you know the basic structure of a Go program.

Continue to the next topic 👇

[Comments ⮕&#xFE0F;](../../02_basics/03_comments/README.md)