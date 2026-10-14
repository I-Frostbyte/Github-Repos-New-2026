# 🚀 Methods

-----------------------------------------------------------------------------------------------------------------------------
## Methods
Go does not have classes. However, you can define methods on types.

A method is a function with a special receiver argument.

The receiver appears in its own argument list between the `func` keyword and the method name.

Remember: a method is just a function with a receiver argument.

You can declare a method on non-struct types, too.

You can only declare a method with a receiver whose type is defined in the same package as the method. You cannot declare a method with a receiver whose type is defined in another package (which includes the built-in types such as int).
-----------------------------------------------------------------------------------------------------------------------------

## Folder Structure

```
1-Methods/
├── methods.go       # Types and methods
├── methods_test.go  # (package declaration only)
└── README.md
```

## Files

| File | Contents |
| --- | --- |
| `methods.go` | Types `Vertex`, `User`, `Age`; method `Vertex.Abs`; constructor `NewUser`; methods `User.GetName` and `Age.IsAdult`. |
| `methods_test.go` | Empty test file (package `methods`). |

## API

```go
type Vertex struct{ X, Y float64 }
type User struct { /* unexported fields */ }
type Age int64

// Abs returns the Euclidean length sqrt(X² + Y²).
func (v Vertex) Abs() float64

// NewUser constructs a User.
func NewUser(name, email, phone_number, address string, age int64) User

// GetName returns the given user's name.
func (u User) GetName(user User) string

// IsAdult reports whether the age is >= 18.
func (a Age) IsAdult() bool
```

## Notes

* Demonstrates methods declared on both struct types (`Vertex`, `User`) and a
  named non-struct type (`Age`), matching the tour's point that methods can be
  defined on any type declared in the same package.
