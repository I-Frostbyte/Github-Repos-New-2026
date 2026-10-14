# 🚀 Sorting Exercise #1

-----------------------------------------------------------------------------------------------------------------------------
Sorting Exercise #1
- Create a simple program that implements the different types of Sorting on an array of numbers or an array of an array of numbers..
- For simplicity, we shall only consider numbers.

- Bubble Sort
- Selection Sort
- Insertion Sort
- Shell Sort
- Merge Sort
- Quick Sort
-----------------------------------------------------------------------------------------------------------------------------

## Folder Structure

```
2-Sorting_Exercise_Sorting_Program/
├── cmd/
│   └── main.go          # Entry point (currently empty main())
├── sort/
│   ├── service.go       # ArraySort type + all six sorting algorithms
│   └── service_test.go  # Tests for every algorithm
├── go.mod               # Module: .../Exercises/2-Sorting_Exercise_Sorting_Program (Go 1.26.5)
├── Makefile             # test / run / install-dependencies targets
└── README.md
```

## Implementation (`sort/service.go`)

`ArraySort` is the public type, constructed with `NewArraySort(input []int)`.
It exposes the six required algorithms plus supporting helpers.

| Method | Signature | Notes |
| --- | --- | --- |
| `NewArraySort` | `(input []int) *ArraySort` | Constructor. |
| `BubbleSort` | `(input []int) []int` | Neighbour comparisons; early exit via `isSwapped`. |
| `SelectionSort` | `(input []int) []int` | Tracks the minimum index per pass. |
| `InsertionSort` | `(input []int) []int` | Shifts larger elements to the right. |
| `ShellSort` | `(input []int) []int` | Builds a descending interval sequence. |
| `MergeSort` | `(input []int) []int` | Recursive split + `JoinArrays` merge. |
| `QuickSort` | `(input []int, below, upper int) []int` | Lomuto-style partition via `divideParts`. |

Supporting helpers: `RandomArrayGeneratorForTests`, `SimpleSwap`, `JoinArrays`,
`divideParts`, `swap`, and `power`.

## Tests (`sort/service_test.go`)

Table-driven tests for all six algorithms (`TestBubbleSort`,
`TestSelectionSort`, `TestInsertionSort`, `TestShellSort`, `TestMergeSort`,
`TestQuickSort`). Each case generates a random array and asserts the result
with `slices.IsSorted`.

## Commands

```bash
make install-dependencies   # go mod tidy
make test                   # go test -v ./...
make run                    # go run main.go
```

## Continuous Integration

This exercise is the target of the repo's CI workflow
(`.github/workflows/sorting-ci.yaml`), which installs dependencies and runs the
tests on push/PR.

## Status

* All six algorithms are implemented and tested.
* `cmd/main.go` is still an empty `main()` — there is no interactive program yet.
