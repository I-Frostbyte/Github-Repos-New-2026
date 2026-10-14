# 🚀 Range and Close

-----------------------------------------------------------------------------------------------------------------------------
## Range and Close
A sender can close a channel to indicate that no more values will be sent. Receivers can test whether a channel has been closed by assigning a second parameter to the receive expression: after

```go
v, ok := <-ch
```

`ok` is `false` if there are no more values to receive and the channel is closed.

The loop for `i := range c` receives values from the channel repeatedly until it is closed.

Note: Only the sender should close a channel, never the receiver. Sending on a closed channel will cause a panic.

Another note: Channels aren't like files; you don't usually need to close them. Closing is only necessary when the receiver must be told there are no more values coming, such as to terminate a `range` loop.
-----------------------------------------------------------------------------------------------------------------------------

## Folder Structure

```
4-Range_and_Close/
├── range_and_close.go       # Fibonacci generator
├── range_and_close_test.go  # (package declaration only)
└── README.md
```

## Files

| File | Contents |
| --- | --- |
| `range_and_close.go` | `Fibonacci(n int, c chan int)` sends the first `n` Fibonacci numbers into `c`, then calls `close(c)`. |
| `range_and_close_test.go` | Empty test file (package `rangeandclose`). |

## API

```go
// Fibonacci sends the first n Fibonacci numbers on c and then closes c.
func Fibonacci(n int, c chan int)
```

## Notes

* Demonstrates the sender-only-close contract and `range` terminating when the
  channel closes.
* In `main.go` it is used with a buffered channel: `c := make(chan int, 10)`,
  then `go range_and_close.Fibonacci(cap(c), c)`.
