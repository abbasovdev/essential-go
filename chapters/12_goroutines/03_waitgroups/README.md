<div align="center">
    <h1>48) WaitGroups</h1>
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
            ┃ <a href="../02_channels/README.md"> ⬅&#xFE0E; Previous Topic</a>
            ┃ <a href="../04_project_part_11/README.md"> Next Topic ⮕&#xFE0F;</a>
        </div>
        <img width="100%"  src="../../../images/header.png" />
        <br />
    </div>
</div>

---


## WaitGroups

A `sync.WaitGroup` waits for a collection of goroutines to finish.

It does not pass data like a channel. It just waits.

### How It Works

Three methods: `Add`, `Done`, `Wait`.

```go
var wg sync.WaitGroup

wg.Add(1)
go func() {
    defer wg.Done()
    fmt.Println("task done")
}()

wg.Wait()
```

- `Add(1)` → one goroutine is starting
- `Done()` → the goroutine finished
- `Wait()` → blocks until all goroutines call `Done`

`defer wg.Done()` ensures `Done` is called even if the goroutine panics.

### Multiple Goroutines

```go
var wg sync.WaitGroup

for i := 1; i <= 3; i++ {
    wg.Add(1)
    go func(n int) {
        defer wg.Done()
        fmt.Printf("task %d done\n", n)
    }(i)
}

wg.Wait()
```

The order of output is not guaranteed. Goroutines run concurrently.

### Example Code

There is some code in [main.go](main.go) file.

But it is incomplete.

### Fix the Code

Open [main.go](main.go) file and complete the code lines.

Then validate your code by running:
```bash
go run topic.go 48 validate
```

✅ If everything is ok, you should see:
```
OK!
```

## Conclusion

Now you know how to use WaitGroups to wait for goroutines.

Continue to the next topic 👇

[Project ⮕&#xFE0F;](../04_project_part_11/README.md)
