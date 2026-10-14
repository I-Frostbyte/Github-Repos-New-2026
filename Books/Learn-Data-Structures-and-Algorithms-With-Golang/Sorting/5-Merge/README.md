# 🚀 Merge Sort

## Description
The merge sort algorithm is a comparison-based method that was invented by John Von Neumann. Each element in the adjacent list is compared for sorting. The performance of the algorithm is in the order of O(n log n). This algorithm is the best algorithm for sorting a linked list.

## Folder Structure

```
5-Merge/
├── main.go        # MergeSort + JoinArrays + RandomlyGeneratedArray + demo main()
├── main_test.go   # TestMergeSort (random inputs)
└── README.md
```

## Files

| File | Contents |
| --- | --- |
| `main.go` | `main()` sorts `RandomlyGeneratedArray(32)`; `MergeSort(arrayToSort []int) []int` recursively splits then rejoins; `JoinArrays(leftArray, rightArray []int) []int` merges two sorted halves; `RandomlyGeneratedArray` is a local helper. |
| `main_test.go` | `TestMergeSort` runs three generated arrays (lengths 45, 60, 90) and asserts with `slices.IsSorted`. |

## API

```go
// MergeSort sorts arrayToSort in ascending order and returns the result.
func MergeSort(arrayToSort []int) []int

// JoinArrays merges two sorted slices into one sorted slice.
func JoinArrays(leftArray, rightArray []int) []int

// RandomlyGeneratedArray generates an array of random ints.
func RandomlyGeneratedArray(arrayLength int) []int
```

## Notes

* Complexity: **O(n log n)** time in all cases; **O(n)** auxiliary space.
* Stable; a good fit for linked lists (sequential access).
