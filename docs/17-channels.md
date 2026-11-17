# Chapter 17 — Channels

A **channel** is a typed conduit for sending and receiving values between goroutines. It is Go's "share memory by communicating" primitive.

## Creating channels

```go
ch := make(chan int)        // unbuffered
ch := make(chan int, 10)    // buffered, capacity 10
var ch chan int             // nil channel — always blocks
```

The zero value is `nil`. Sending to or receiving from a nil channel blocks forever (useful in `select` to disable a case).

## Send, receive, close

```go
ch <- 42           // send
v := <-ch          // receive
v, ok := <-ch      // ok == false when channel is closed and drained
close(ch)          // close: no more values can be sent
```

Closing is the **sender's** responsibility. Receivers can detect closure with the comma-ok form or by ranging:

```go
for v := range ch {
    fmt.Println(v)
}
```

The loop ends automatically when the channel is closed and empty.

## Unbuffered vs buffered

| Channel | Behavior |
|---------|----------|
| Unbuffered `make(chan T)` | send blocks until a receiver is ready (rendezvous) |
| Buffered `make(chan T, n)` | send blocks only when buffer is full |

```go
ch := make(chan string)

go func() {
    ch <- "ping"       // blocks until main receives
}()

msg := <-ch
fmt.Println(msg)       // ping
```

Unbuffered channels provide **synchronization**: the send completes only when the receive happens.

Buffered channels decouple producer and consumer:

```go
ch := make(chan int, 3)
ch <- 1
ch <- 2
ch <- 3
// ch <- 4  // would block until someone receives
```

Use buffered channels to smooth bursts, but avoid using large buffers to hide design problems.

## Closing and the comma-ok idiom

```go
ch := make(chan int, 2)
ch <- 1
ch <- 2
close(ch)

for {
    v, ok := <-ch
    if !ok {
        break
    }
    fmt.Println(v)
}
```

Reading from a **closed** channel immediately returns the zero value (with `ok == false`). Reading from a **nil** channel blocks forever.

**Panic conditions:** sending on a closed channel, or closing a nil/already-closed channel.

## Directional channels

Restrict direction in function signatures for safety:

```go
func producer(out chan<- int) {
    for i := 0; i < 5; i++ {
        out <- i
    }
    close(out)
}

func consumer(in <-chan int) {
    for v := range in {
        fmt.Println(v)
    }
}

ch := make(chan int)
go producer(ch)
consumer(ch)
```

`chan<- T` = send-only, `<-chan T` = receive-only. A bidirectional `chan T` can be passed to either.

## Fan-out, fan-in

**Fan-out** — multiple goroutines read from the same channel:

```go
jobs := make(chan int)
for w := 0; w < 3; w++ {
    go worker(jobs)
}
for j := 0; j < 10; j++ {
    jobs <- j
}
close(jobs)
```

**Fan-in** — multiple goroutines send to one channel:

```go
func merge(cs ...<-chan int) <-chan int {
    out := make(chan int)
    var wg sync.WaitGroup
    for _, c := range cs {
        wg.Add(1)
        go func(c <-chan int) {
            defer wg.Done()
            for v := range c {
                out <- v
            }
        }(c)
    }
    go func() {
        wg.Wait()
        close(out)
    }()
    return out
}
```

## Pipelines

A pipeline is a series of stages connected by channels:

```go
func gen(nums ...int) <-chan int {
    out := make(chan int)
    go func() {
        for _, n := range nums {
            out <- n
        }
        close(out)
    }()
    return out
}

func sq(in <-chan int) <-chan int {
    out := make(chan int)
    go func() {
        for n := range in {
            out <- n * n
        }
        close(out)
    }()
    return out
}

for n := range sq(gen(2, 3, 4)) {
    fmt.Println(n)   // 4 9 16
}
```

## Backpressure

Channels naturally provide backpressure: a slow consumer blocks a fast producer (on unbuffered/full channels), preventing unbounded memory growth. This is a major advantage over unbounded queues.

## Selecting and timeouts

`select` (Chapter 18) lets you wait on multiple channels, add timeouts, and implement cancellation:

```go
select {
case v := <-ch:
    fmt.Println("got", v)
case <-time.After(1 * time.Second):
    fmt.Println("timeout")
}
```

## Common pitfalls

1. **Blocking forever** — sending with no receiver, or receiving with no sender.
2. **Forgetting to close** — `range` over a channel never terminates.
3. **Closing from the receiver** — only the sender should close.
4. **Closing too early** — other senders panic.
5. **Nil channels** — accidentally blocking forever (useful deliberately in `select`).

## Runnable example

See `examples/16-channels/main.go`.

## Chapter summary

- Channels are typed conduits: `ch <- v` sends, `<-ch` receives.
- Unbuffered channels rendezvous; buffered ones queue up to capacity.
- The sender closes; receivers detect closure with comma-ok or `range`.
- Use directional channel types in signatures.
- Channels enable pipelines, fan-out/fan-in, and natural backpressure.

## Exercises

1. Write a generator that emits the Fibonacci sequence on a channel.
2. Implement fan-out: distribute jobs to 3 workers.
3. Write a `merge` function combining two channels.
4. Show what happens when you send on a closed channel (in a `recover`-protected demo).

*(Solutions: `exercises/SOLUTIONS.md`, Chapter 17)*
