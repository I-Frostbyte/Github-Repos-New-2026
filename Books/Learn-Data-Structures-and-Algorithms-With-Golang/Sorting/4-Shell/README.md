# 🚀 Shell Sort

## Description
The shell sort algorithm sorts a pair of elements that are not in sequence in a collection. The distance between the elements to be compared is decreased sequentially. This algorithm performs more operations and has a greater cache miss ratio than the quick sort algorithm.

## Folder Structure

```
4-Shell/
├── main.go        # ShellSort + RandomNumberArrayGenerator + demo main()
├── main_test.go   # TestShellSort (random inputs)
└── README.md
```

## Files

| File | Contents |
| --- | --- |
| `main.go` | `main()` sorts `{11, 4, 8, 6, 19, 21, 71, 13, 15, 2}` and prints a generated array; `ShellSort(numbers []int) []int` builds a descending interval sequence and performs interval-based insertions; `RandomNumberArrayGenerator` is a local helper. |
| `main_test.go` | `TestShellSort` runs three generated arrays (lengths 32, 35, 40) and asserts with `slices.IsSorted`. |

## API

```go
// ShellSort sorts numbers in ascending order and returns the same slice.
func ShellSort(numbers []int) []int

// RandomNumberArrayGenerator generates an array of random ints.
func RandomNumberArrayGenerator(arrayLength int) []int
```

## Notes

* Generalization of insertion sort over a decreasing gap sequence.
* In-place; complexity depends on the chosen interval sequence.
