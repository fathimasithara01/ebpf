# Problem Statement 3 — Go Concurrency Explanation

## Given Code

```go
package main

import "fmt"

func main() {
    cnp := make(chan func(), 10)

    for i := 0; i < 4; i++ {
        go func() {
            for f := range cnp {
                f()
            }
        }()
    }

    cnp <- func() {
        fmt.Println("HERE1")
    }

    fmt.Println("Hello")
}
```

---

# 1. What is this code attempting to do?

This code demonstrates a simple worker-pool pattern using goroutines and channels in Go.

The program:

* Creates a buffered channel that stores functions.
* Starts 4 worker goroutines.
* Each worker continuously reads functions from the channel.
* When a function is received, the worker executes it.

The line:

```go
cnp <- func() {
    fmt.Println("HERE1")
}
```

pushes a function into the channel, and one of the worker goroutines is expected to execute it.

---

# 2. How do the highlighted constructs work?

## `make(chan func(), 10)`

This creates a buffered channel that can store functions (`func()` type).

* `chan func()` → channel carrying functions
* `10` → buffer capacity

This means up to 10 functions can be queued without blocking the sender.

---

## `go func() { ... }()`

This creates and immediately executes an anonymous goroutine.

Each goroutine acts as a worker thread that processes tasks from the channel.

---

## `for f := range cnp`

This continuously receives functions from the channel until the channel is closed.

For every received function:

```go
f()
```

executes that function.

---

# 3. What is the significance of the loop with 4 iterations?

```go
for i := 0; i < 4; i++
```

This creates 4 worker goroutines.

This pattern is commonly called a worker pool.

Benefits:

* concurrent task execution
* parallel processing
* efficient background job handling

Only one worker will execute a given function, but multiple workers can process multiple queued tasks concurrently.

---

# 4. What is the significance of `make(chan func(), 10)`?

The channel is buffered with capacity 10.

Meaning:

* sender can queue up to 10 functions
* sender does not immediately block
* workers can consume tasks asynchronously

Without buffering:

* sender would block until a worker receives the function.

This improves throughput and decouples producers from consumers.

---

# 5. Why is `HERE1` not getting printed?

The issue is caused by program termination timing.

Sequence:

1. Worker goroutines are started.
2. Function is pushed into the channel.
3. `fmt.Println("Hello")` executes immediately.
4. `main()` exits.
5. All goroutines terminate automatically.

The worker goroutine may not get enough time to execute:

```go
fmt.Println("HERE1")
```

before the process exits.

So usually only:

```text
Hello
```

gets printed.

This is a goroutine scheduling and synchronization issue.

---

# 6. How can this be fixed?

The main goroutine must wait before exiting.

Possible solutions:

* `sync.WaitGroup`
* channel synchronization
* `time.Sleep()` (not ideal but works for demo)

Example using sleep:

```go
package main

import (
    "fmt"
    "time"
)

func main() {
    cnp := make(chan func(), 10)

    for i := 0; i < 4; i++ {
        go func() {
            for f := range cnp {
                f()
            }
        }()
    }

    cnp <- func() {
        fmt.Println("HERE1")
    }

    fmt.Println("Hello")

    time.Sleep(time.Second)
}
```

Now both outputs become visible because the program waits before exiting.

---

# Key Concepts Demonstrated

* Goroutines
* Buffered channels
* Worker pool pattern
* Concurrent task execution
* Scheduler behavior
* Synchronization issues in Go
