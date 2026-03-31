<div align="center">
    <h1>46) Goroutines</h1>
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
            ┃ <a href="../../11_errors/05_project_part_10/README.md"> ⬅&#xFE0E; Previous Topic</a>
            ┃ <a href="../02_channels/README.md"> Next Topic ⮕&#xFE0F;</a>
        </div>
        <img width="100%"  src="../../../images/header.png" />
        <br />
    </div>
</div>

---


## Goroutines

A goroutine is a lightweight thread managed by the Go runtime.

```
  main()            goroutine 1       goroutine 2
    |
    |--go f1()--------->|
    |                   |
    |--go f2()------------------------->|
    |                   |               |
    |  (continues)      |  f1()         |  f2()
    |                   |  running      |  running
    |                   |               |
    v                   v               v
  main returns → all goroutines stop
```

You start one by putting the `go` keyword before a function call.

### Launching a Goroutine

```go
func sayHello() {
    fmt.Println("hello from goroutine")
}

func main() {
    go sayHello()
}
```

The `go` keyword runs `sayHello` concurrently with `main`.

### Why Nothing Prints

When `main` returns, all goroutines are stopped. If `main` finishes before the goroutine runs, you see nothing.

`time.Sleep` is a temporary way to wait:

```go
func main() {
    go sayHello()
    time.Sleep(100 * time.Millisecond)
}
```

This is not the proper way to wait. Channels and WaitGroups are covered in the next topics.

### Anonymous Goroutines

You can launch a goroutine with an anonymous function:

```go
go func() {
    fmt.Println("anonymous goroutine")
}()
```

The `()` at the end calls the function immediately as a goroutine.

### Example Code

There is some code in [main.go](main.go) file.

But it is incomplete.

### Fix the Code

Open [main.go](main.go) file and complete the code lines.

Then validate your code by running:
```bash
go run topic.go 46 validate
```

✅ If everything is ok, you should see:
```
OK!
```

## Conclusion

Now you know how to launch goroutines in Go.

Continue to the next topic 👇

[Channels ⮕&#xFE0F;](../02_channels/README.md)
