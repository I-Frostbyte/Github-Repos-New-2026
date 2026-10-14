# 🚀 Channels

-----------------------------------------------------------------------------------------------------------------------------
## Channels
Channels are a typed conduit through which you can send and receive values with the channel operator, <-.

```go
ch <- v    // Send v to channel ch.
v := <-ch  // Receive from ch, and assign value to v.
```
(The data flows in the direction of the arrow.)

Like maps and slices, channels must be created before use:

```go
ch := make(chan int)
```
By default, sends and receives block until the other side is ready. This allows goroutines to synchronize without explicit locks or condition variables.
-----------------------------------------------------------------------------------------------------------------------------

## Folder Structure

```
2-Channels/
├── channels.go       # SumOfArrayElements + halfArraySum
├── channels_test.go  # (package declaration only)
└── README.md
```

## Files

| File | Contents |
| --- | --- |
| `channels.go` | `SumOfArrayElements(array []int) int` splits the slice in half, sums each half on its own goroutine, and combines the results received from an unbuffered channel; `halfArraySum` is the worker. |
| `channels_test.go` | Empty test file (package `channels`). |

## API

```go
// SumOfArrayElements sums an array by splitting the work across two goroutines.
func SumOfArrayElements(array []int) int

// halfArraySum sums halfArray and sends the result on sumChannel.
func halfArraySum(halfArray []int, sumChannel chan int)
```

## Notes

* Demonstrates using an unbuffered channel to synchronize two goroutines
  without explicit locks.
* The receive `rightHalfSum, leftHalfSum := <-sumChannel, <-sumChannel` collects
  the two partial sums (order is non-deterministic, but addition is commutative).
