# 🚀 Searching

## Description
Search algorithms are used to retrieve information that's stored in a data source or a collection. The algorithm is given the key of the element in question, and the associated value will be found. Search algorithms return a true or a false Boolean value based on the availability of the information. They can be enhanced to display multiple values related to the search criteria. Different types of search algorithms include linear, binary, and interpolation. These algorithms are categorized by the type of search. Search algorithms include brute force and heuristic methods. The algorithms are chosen for their efficiency. Different factors for choosing these algorithms are as follows:
* Input type
* Output type
* Definiteness
* Correctness
* Finiteness
* Effectiveness
* Generality

In the folders, you'll find the different searching algorithms, that is, linear, binary, interpolation.

## Folder Structure

```
Searching/
├── 1-Linear/     # Linear search (O(n))
├── 2-Binary/     # Binary search via sort.Search (O(log n))
├── generator/    # Shared random-array generator used by the tests
├── go.mod        # Go module definition (Go 1.26.5)
├── Makefile      # Test target
└── README.md
```

## Contents

| Path | Type | Description |
| --- | --- | --- |
| [1-Linear](1-Linear) | Module package | Sequential search; returns `true`/`false`. Files: `main.go`, `main_test.go`. |
| [2-Binary](2-Binary) | Module package | Binary search built on `sort.Search`; operates on a sorted slice. Files: `main.go`, `main_test.go`. |
| [generator](generator) | Module package | `Generator.RandomIntegerArrayGenerator` produces test input arrays (values `0`–`30`). File: `generator.go`. |

> The README mentions an *interpolation* search, but no `interpolation` folder exists yet.

## Implementation Notes

* Both search implementations are tested against arrays produced by `generator`.
* `2-Binary`'s test sorts its input with `slices.Sorted(slices.Values(...))` before searching.

## Running the Tests

```bash
make test        # runs `go test -v ./...` across the module
```
