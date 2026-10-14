# 🚀 Binary Searching

## Description
The binary search algorithm compares the input value to the middle element of the sorted collection. If the values are not equal, the half in which the element is not found is eliminated. The search continues on the remaining half of the collection. The time complexity of this algorithm is in the order of O(log n).

## Folder Structure

```
2-Binary/
├── main.go        # BinarySearch implementation + empty main()
├── main_test.go   # Table-driven test (sorts input first)
└── README.md
```

## Files

| File | Contents |
| --- | --- |
| `main.go` | `main()` (no-op entry point) and `BinarySearch(inputArray []int, findElement int) bool`, built on `sort.Search`. |
| `main_test.go` | `TestBinarySearch` with three table cases; each input is sorted with `slices.Sorted` before searching. |

## API

```go
// BinarySearch returns true if findElement exists in the (sorted) inputArray.
func BinarySearch(inputArray []int, findElement int) bool
```

## Notes

* Complexity: **O(log n)** time on a sorted input; **O(1)** space.
* Uses `sort.Search` to locate the index where `findElement` would appear, then
  compares the value at that index to confirm an exact match.
