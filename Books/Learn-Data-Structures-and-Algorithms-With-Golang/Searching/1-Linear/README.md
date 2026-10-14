# 🚀 Linear Searching

## Description
The linear search method finds a given value within a collection by sequentially checking every element in the collection. The time complexity of the linear search algorithm is O(n). The binary search algorithm and hash tables perform better than this search algorithm.

## Folder Structure

```
1-Linear/
├── main.go        # LinearSearch implementation + empty main()
├── main_test.go   # Table-driven test using the shared generator
└── README.md
```

## Files

| File | Contents |
| --- | --- |
| `main.go` | `main()` (no-op entry point) and `LinearSearch(numberArray []int, findElement int) bool`. |
| `main_test.go` | `TestLinearSearch` with three table cases; input arrays come from `Searching/generator.RandomIntegerArrayGenerator`. |

## API

```go
// LinearSearch returns true if findElement is present in numberArray.
func LinearSearch(numberArray []int, findElement int) bool
```

## Notes

* Complexity: **O(n)** time, **O(1)** space.
* Tests generate random arrays and assert the boolean result against the
  generator's `checkForTarget` outcome.
