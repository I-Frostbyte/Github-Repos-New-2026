# 🚀 Sorting

## Description
Sorting algorithms arrange the elements in a collection in ascending or descending order. Lexicographical order can be applied to a collection of characters and strings. The efficiency of these algorithms is in the performance of sorting the input data into a sorted collection. The best sorting algorithm time complexity is O(n log n). Sorting algorithms are classified by the following criteria:
* Computational complexity
* Memory usage
* Stability
* Type of sorting: serial/parallel
* Adaptability
* Method of sorting

In the folders, you'll find the different sorting algorithms, that is, bubble, selection, insertion, shell, merge, and quick.

## Folder Structure

```
Sorting/
├── 1-Bubble/      # Bubble sort (O(n²))
├── 2-Selection/   # Selection sort (O(n²))
├── 3-Insertion/   # Insertion sort (O(n²))
├── 4-Shell/       # Shell sort (interval-based)
├── 5-Merge/       # Merge sort (O(n log n))
├── 6-Quick/       # Quick sort (O(n log n))
├── go.mod         # Go module definition (Go 1.26.5)
├── Makefile       # Test target
└── README.md
```

## Contents

| Path | Algorithm | Complexity | Files |
| --- | --- | --- | --- |
| [1-Bubble](1-Bubble) | Bubble Sort | O(n²) | `main.go`, `main_test.go` |
| [2-Selection](2-Selection) | Selection Sort | O(n²) | `main.go`, `main_test.go` |
| [3-Insertion](3-Insertion) | Insertion Sort | O(n²) | `main.go`, `main_test.go` |
| [4-Shell](4-Shell) | Shell Sort | interval-based | `main.go`, `main_test.go` |
| [5-Merge](5-Merge) | Merge Sort | O(n log n) | `main.go`, `main_test.go` |
| [6-Quick](6-Quick) | Quick Sort | O(n log n) | `main.go`, `main_test.go` |

## Implementation Notes

* Each algorithm lives in its own package (`package main`) with a no-op
  `main()` entry point and an exported sort function.
* Every implementation is verified with `slices.IsSorted` against a
  `slices.Sorted` reference in its `main_test.go`.
* `3-Insertion`, `4-Shell`, `5-Merge` and `6-Quick` generate random input via a
  local `RandomNumberArrayGenerator` / `RandomlyGeneratedArray` helper.

## Running the Tests

```bash
make test        # runs `go test -v ./...` across the module
```
