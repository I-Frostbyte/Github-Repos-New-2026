# 🚀 Quick Sort

## Description
Quick sort is an algorithm for sorting the elements of a collection in an organized way. Parallelized quick sort is two to three times faster than merge sort and heap sort. The algorithm's performance is of the order O(n log n).  This algorithm is a space-optimized version of the binary tree sort algorithm.

## More Explanation
High-level idea

Unlike Merge Sort, which:

1. Splits the array in half first.
2. Sorts each half.
3. Merges them back together.

Quick Sort instead:

1. Picks one element called the pivot.
2. Rearranges the array so that:
    * everything smaller than the pivot is to its left,
    * everything larger than the pivot is to its right.
3. The pivot is now in its final sorted position.
4. Repeat the same process on the left side and the right side.

Notice something important:

Merge Sort splits first, then sorts.

Quick Sort sorts while splitting.

## Folder Structure

```
6-Quick/
├── main.go        # QuickSort + divideParts + demo main()
├── main_test.go   # TestQuickSort (random inputs)
└── README.md
```

## Files

| File | Contents |
| --- | --- |
| `main.go` | `main()` sorts `{11, 4, 8, 6, 19, 21, 71, 13, 15, 2}` by calling `QuickSort(newArray, 0, 9)`; `QuickSort(numbersArray []int, below, upper int) []int` partitions then recurses; `divideParts` performs the Lomuto-style partition around the last element. |
| `main_test.go` | `TestQuickSort` runs three generated arrays (lengths 15, 25, 50) and asserts with `slices.IsSorted`. |

## API

```go
// QuickSort sorts numbersArray[below..upper] in ascending order.
func QuickSort(numbersArray []int, below, upper int) []int
```

## Notes

* Complexity: **O(n log n)** average, **O(n²)** worst case; **O(log n)** stack
  space on average.
* In-place; not stable.
