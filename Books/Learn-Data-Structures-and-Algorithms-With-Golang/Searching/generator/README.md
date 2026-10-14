# 🚀 Generator

## Description
The Generator methods are used to generate arrays and (in the future) lists that can be referenced by test functions.

## Folder Structure

```
generator/
├── generator.go   # Generator type and random-array factory
└── README.md
```

## Files

| File | Contents |
| --- | --- |
| `generator.go` | `Generator` struct, `NewGenerator()` constructor, and `RandomIntegerArrayGenerator`. |

## API

```go
type Generator struct{}

// NewGenerator returns an empty *Generator.
func NewGenerator() *Generator

// RandomIntegerArrayGenerator builds an array of length arrayLength with
// values in the range 0–30. When checkForTarget is true it also reports
// whether target appears in the generated array.
func (g *Generator) RandomIntegerArrayGenerator(
    arrayLength, target int,
    checkForTarget bool,
) (*[]int, bool, error)
```

## Notes

* Returns an error when `arrayLength < 2`.
* Values are generated with `rand.Intn(30)`, so duplicates are possible.
* Used as shared test input by the `1-Linear` and `2-Binary` search tests.
* The README anticipates additional list generators in the future.
