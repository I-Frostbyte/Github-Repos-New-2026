# 🚀 Insertion Sort

## Description
Insertion sort is an algorithm that creates a final sorted array one element at a time. The algorithm's performance is of the order O(n2). This algorithm is less efficient on large collections than other algorithms, such as quick, heap, and merge sort. In real life, a good example of insertion sort is the way cards are manually sorted by the players in a game of bridge.

## Folder Structure

```
3-Insertion/
├── main.go        # InsertionSort + RandomNumberArrayGenerator + demo main()
├── main_test.go   # TestInsertionSort (random inputs)
└── README.md
```

## Files

| File | Contents |
| --- | --- |
| `main.go` | `main()` sorts `RandomNumberArrayGenerator(12)`; `InsertionSort(numbers []int) []int` shifts larger elements right one at a time; `RandomNumberArrayGenerator(arrayLength int) []int` fills values with `rand.Intn(150)`. |
| `main_test.go` | `TestInsertionSort` runs three generated arrays (lengths 10, 11, 12) and asserts with `slices.IsSorted`. |

## API

```go
// InsertionSort sorts numbers in ascending order and returns the same slice.
func InsertionSort(numbers []int) []int

// RandomNumberArrayGenerator generates an array of random ints in [0, 150).
func RandomNumberArrayGenerator(arrayLength int) []int
```

## Notes

* Complexity: **O(n²)** time (best case O(n) on nearly-sorted data),
  **O(1)** space.
* In-place and stable.
