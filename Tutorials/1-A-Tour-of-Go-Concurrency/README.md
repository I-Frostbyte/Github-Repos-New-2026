# 🚀 A Tour of Go - Concurrency

-----------------------------------------------------------------------------------------------------------------------------
A Tour of Go:
- Goroutines
- Channels
- Buffered Channels
- Range and Close
- Select
- Default Selection
-----------------------------------------------------------------------------------------------------------------------------

## Folder Structure

```
1-A-Tour-of-Go-Concurrency/
├── 1-Goroutines/          # Goroutines + sync.WaitGroup
├── 2-Channels/            # Channels
├── 3-Buffered_Channels/   # Buffered channels
├── 4-Range_and_Close/     # range over channels, close
├── main.go                # Wires all four topics into one demo
├── go.mod                 # Module: .../Tutorials/1-A-Tour-of-Go-Concurrency
├── go.sum
├── Makefile               # test / bench-test / run
└── README.md
```

## Contents

| Topic | Folder | Package | Key API |
| --- | --- | --- | --- |
| Goroutines | [1-Goroutines](1-Goroutines) | `goroutines` | `RandomArrayGenerator`, `CreateAndSortArray` |
| Channels | [2-Channels](2-Channels) | `channels` | `SumOfArrayElements`, `halfArraySum` |
| Buffered Channels | [3-Buffered_Channels](3-Buffered_Channels) | `bufferedchannels` | `CreateMultipleRandomArrays` |
| Range and Close | [4-Range_and_Close](4-Range_and_Close) | `rangeandclose` | `Fibonacci` |

> `Select` and `Default Selection` are listed above but do not have folders yet —
> only the first four topics are implemented.

## `main.go`

`main()` demonstrates each topic in turn:

1. Creates and sorts a random array on a goroutine (`CreateAndSortArray(12)`).
2. Sums the array across two goroutines using an unbuffered channel.
3. Generates multiple sorted arrays through a buffered channel.
4. Streams Fibonacci numbers through a buffered channel with `range`.

## Notes

* `3-Buffered_Channels/buffered_channels.go` imports `.../Tutorials/1-A-Tour-of-Go/1-Goroutines`
  and `go.mod` requires `.../Tutorials/1-A-Tour-of-Go`. These are stale import
  paths from before this folder was renamed to `1-A-Tour-of-Go-Concurrency`.
* The `_test.go` files in the topic folders contain only the `package`
  declaration (no test cases yet).

## Commands

```bash
make run          # go run main.go
make test         # tests in all sub-directories
make bench-test   # benchmarks in all sub-directories
```
