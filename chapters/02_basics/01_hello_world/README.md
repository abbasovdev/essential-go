<div align="center">
    <h1>6) Hello World</h1>
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
            ┃ <a href="../../01_intro/04_project_info/README.md"> ⬅&#xFE0E; Previous Topic</a>
            ┃ <a href="../../02_basics/02_program_structure/README.md"> Next Topic ⮕&#xFE0F;</a>
        </div>
        <img width="100%"  src="../../../images/header.png" />
        <br />
    </div>
</div>

---


## Hello World

`Hello World` is a traditional first program for every developer.

Let's write a `Hello World` program in Go.

## Your First Program in Go

Open [main.go](main.go) file.

You should see the following code:
```go
package main

import "fmt"

func main() {
	fmt.Println("Hello!")
}
```

Not worry about the code for now.

We will talk about each line of code in the upcoming topics.

First, let's run the program.

Open a terminal and run the following command:
```bash
go run topic.go 6
```

⚠️ Make sure you are in the root directory, if you got error like `topic.go: no such file or directory`.

⚠️ Why do we use `go run topic.go 6` ? Checkout the [How to Use this Repo](../../01_intro/03_how_to_use/README.md) topic.


✅ If everything is ok, you should see:
```
Hello!
```

Nice, but something feels not right.

We should see "Hello, World!" instead of "Hello!".

Let's fix it.

### Fix the Code

Open [main.go](main.go) file and change the code.

Verify your output again by running:
```bash
go run topic.go 6
```

✅ If everything is ok, you should see:
```
Hello, World!
```

You can also validate your code by running:
```bash
go run topic.go 6 validate
```

✅ If everything is ok, you should see:
```
OK!
```

🎉 Congratulations! You have written your first Go program.

## Conclusion

Now you know how to write your first Go program.

Continue to the next topic 👇

[Program Structure ⮕&#xFE0F;](../../02_basics/02_program_structure/README.md)