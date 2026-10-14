# 🚀 A Tour of Go - Methods and Interfaces

-----------------------------------------------------------------------------------------------------------------------------
A Tour of Go:
- Methods
- Pointer Receivers
- Methods and Pointer indirection
- Interfaces
- Interface Values
- The Empty Interface
- Type Assertions
- Stringers
- Errors
- Readers
- Images
-----------------------------------------------------------------------------------------------------------------------------

## Folder Structure

```
2-A-Tour-of-Go-Methods_and_Interfaces/
├── 1-Methods/     # Methods on types
├── main.go        # Runs the Methods demos
├── go.mod         # Module: .../Tutorials/2-A-Tour-of-Go-Methods_and_Interfaces
├── Makefile       # test / bench-test / run
└── README.md
```

## Contents

| Topic | Folder | Package | Status |
| --- | --- | --- | --- |
| Methods | [1-Methods](1-Methods) | `methods` | Implemented |

> The remaining topics listed above (pointer receivers, interfaces, type
> assertions, stringers, errors, readers, images) do **not** have folders yet —
> only `Methods` is implemented in this module.

## `main.go`

Demonstrates the `1-Methods` package:

1. `Vertex{X: 3, Y: 4}.Abs()` → `5`.
2. `NewUser(...)` then `GetName(...)`.
3. `Age(23).IsAdult()` → `true`.

## Commands

```bash
make run          # go run main.go
make test         # tests in all sub-directories
make bench-test   # benchmarks in all sub-directories
```
