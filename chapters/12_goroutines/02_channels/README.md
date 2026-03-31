<div align="center">
    <h1>47) Channels</h1>
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
            ┃ <a href="../01_goroutines/README.md"> ⬅&#xFE0E; Previous Topic</a>
            ┃ <a href="../03_waitgroups/README.md"> Next Topic ⮕&#xFE0F;</a>
        </div>
        <img width="100%"  src="../../../images/header.png" />
        <br />
    </div>
</div>

---


## Channels

Channels let goroutines communicate. You send a value into a channel and receive it somewhere else.

### Creating a Channel

```go
ch := make(chan string)
```

`chan string` means the channel carries string values.

### Sending and Receiving

```go
ch <- "hello"   // send
msg := <-ch     // receive
```

Sending blocks until another goroutine receives. Receiving blocks until a value is available.

### Using Channels with Goroutines

```go
ch := make(chan string)

go func() {
    ch <- "hello from goroutine"
}()

msg := <-ch
fmt.Println(msg)
```

No `time.Sleep` needed. The receive `<-ch` blocks until the goroutine sends.

### Example Code

There is some code in [main.go](main.go) file.

But it is incomplete.

### Fix the Code

Open [main.go](main.go) file and complete the code lines.

Then validate your code by running:
```bash
go run topic.go 47 validate
```

✅ If everything is ok, you should see:
```
OK!
```

## Conclusion

Now you know how to use channels to communicate between goroutines.

Continue to the next topic 👇

[WaitGroups ⮕&#xFE0F;](../03_waitgroups/README.md)
