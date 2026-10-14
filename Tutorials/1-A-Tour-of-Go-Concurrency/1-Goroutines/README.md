# 🚀 Goroutines

-----------------------------------------------------------------------------------------------------------------------------
## Goroutines
A goroutine is a lightweight thread managed by the Go runtime.

```go
go f(x, y, z)
```
starts a new goroutine running

```go
f(x, y, z)
```
The evaluation of f, x, y, and z happens in the current goroutine and the execution of f happens in the new goroutine.

Goroutines run in the same address space, so access to shared memory must be synchronized. The sync package provides useful primitives, although you won't need them much in Go as there are other primitives.
-----------------------------------------------------------------------------------------------------------------------------

## Folder Structure

```
1-Goroutines/
├── goroutines.go       # RandomArrayGenerator + CreateAndSortArray
├── goroutines_test.go  # (package declaration only)
└── README.md
```

## Files

| File | Contents |
| --- | --- |
| `goroutines.go` | `RandomArrayGenerator(arrayLength int) []int` (values `0`–`74`) and `CreateAndSortArray(arrayLength int) []int`, which sorts on a goroutine and waits with `sync.WaitGroup`. |
| `goroutines_test.go` | Empty test file (package `goroutines`). |

## API

```go
// RandomArrayGenerator returns a slice of random ints in [0, 75).
func RandomArrayGenerator(arrayLength int) []int

// CreateAndSortArray generates, sorts (in a goroutine) and returns an array.
func CreateAndSortArray(arrayLength int) []int
```

> Note: `CreateAndSortArray` currently ignores its `arrayLength` argument and
> always generates a length-12 array.
