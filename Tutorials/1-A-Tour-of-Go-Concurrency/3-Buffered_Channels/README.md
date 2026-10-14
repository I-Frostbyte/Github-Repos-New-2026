# 🚀 Buffered Channels

-----------------------------------------------------------------------------------------------------------------------------
## Buffered Channels
Channels can be buffered. Provide the buffer length as the second argument to make to initialize a buffered channel:

```go
ch := make(chan int, 100)
```
Sends to a buffered channel block only when the buffer is full. Receives block when the buffer is empty.
-----------------------------------------------------------------------------------------------------------------------------

## Folder Structure

```
3-Buffered_Channels/
├── buffered_channels.go       # CreateMultipleRandomArrays
├── buffered_channels_test.go  # (package declaration only)
└── README.md
```

## Files

| File | Contents |
| --- | --- |
| `buffered_channels.go` | `CreateMultipleRandomArrays(arrayLengths []int) [][]int` creates a buffered channel sized to the number of arrays, pushes each sorted array into it, then drains it into a result slice. |
| `buffered_channels_test.go` | Empty test file (package `bufferedchannels`). |

## API

```go
// CreateMultipleRandomArrays builds (and sorts) one random array per requested length.
func CreateMultipleRandomArrays(arrayLengths []int) [][]int
```

## Notes

* The buffer is sized with `len(arrayLengths)` so all sends complete without a
  concurrent receiver.
* Depends on `1-Goroutines.CreateAndSortArray`. The import path in the source
  (`.../Tutorials/1-A-Tour-of-Go/1-Goroutines`) is stale relative to the current
  folder name `1-A-Tour-of-Go-Concurrency`.
