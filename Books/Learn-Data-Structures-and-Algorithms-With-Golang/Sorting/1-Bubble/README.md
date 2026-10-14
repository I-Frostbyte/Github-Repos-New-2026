# 🚀 Bubble Sort

## Description
The bubble sort algorithm is a sorting algorithm that compares a pair of neighboring elements and swaps them if they are in the wrong order. The algorithm has a complexity of O(n2), where n is the number of elements to be sorted. The smallest or greatest value bubbles up to the top of the collection, or the smallest or greatest sinks to the bottom (depending on whether you're sorting into ascending or descending order).

## Folder Structure

```
1-Bubble/
├── main.go        # BubbleSort + demo main()
├── main_test.go   # TestBubbleSort (table-driven)
└── README.md
```

## Files

| File | Contents |
| --- | --- |
| `main.go` | `main()` demonstrates the sort on `{31, 13, 12, 4, 18, 16, 7, 2, 3, 0, 10}`; `BubbleSort(numbers []int) []int` uses a `isSwapped` flag to stop early once sorted. |
| `main_test.go` | `TestBubbleSort` runs three fixed inputs and asserts the result with `slices.IsSorted`. Includes a legacy Go ≤ 1.22 note. |

## API

```go
// BubbleSort sorts numbers in ascending order and returns the same slice.
func BubbleSort(numbers []int) []int
```

## Notes

* Complexity: **O(n²)** time (best case O(n) with the early-exit flag),
  **O(1)** space.
* In-place and stable.
