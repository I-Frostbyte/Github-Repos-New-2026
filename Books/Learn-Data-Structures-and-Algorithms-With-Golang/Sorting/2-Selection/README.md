# 🚀 Selection Sort

## Description
Selection sort is an algorithm that divides the input collection into two fragments. This sublist of elements is sorted by swapping the smallest or largest element from the left of the list to the right. The algorithm is of the order O(n2). This algorithm is inefficient for large collections, and it performs worse than the insertion sort algorithm.

## Folder Structure

```
2-Selection/
├── main.go        # SelectionSort + swap helper + demo main()
├── main_test.go   # TestSelectionSort (table-driven)
└── README.md
```

## Files

| File | Contents |
| --- | --- |
| `main.go` | `main()` demonstrates the sort on `{11, 4, 8, 6, 19, 21, 71, 13, 15, 2}`; `SelectionSort(elements []int) []int` tracks the minimum index per pass and swaps via `swap`. |
| `main_test.go` | `TestSelectionSort` runs three fixed inputs and asserts with `slices.IsSorted`. |

## API

```go
// SelectionSort sorts elements in ascending order and returns the same slice.
func SelectionSort(elements []int) []int
```

## Notes

* Complexity: **O(n²)** time, **O(1)** space.
* In-place; not stable (swaps can reorder equal elements).
