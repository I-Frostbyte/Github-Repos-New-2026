# 🚀 Simple Map Exercise #1

-----------------------------------------------------------------------------------------------------------------------------
Simple Exercise #1
- Create a simple map that maps a boolean to a string entry.
- The strings are names and the booleans represent their presence in a set array.
- The array represents the names of student present in class on one particular day.
-----------------------------------------------------------------------------------------------------------------------------

## Folder Structure

```
3-Simple_Map_Exercise/
├── cmd/
│   └── main.go               # Entry point (currently empty main())
├── simple_map/
│   ├── service.go            # StudentClass attendance logic
│   └── service_test.go       # Tests for class creation & attendance
├── go.mod                    # Module: .../Exercises/3-Simple_Map_Exercise (Go 1.26.5)
├── Makefile                  # test / bench-test targets
└── README.md
```

## Implementation (`simple_map/service.go`)

The attendance map is keyed by student name (`map[string]bool`): a student maps
to `false` on enrolment and `true` once marked present.

| Item | Signature | Notes |
| --- | --- | --- |
| `StudentClass` | struct | Holds `className string` and `attendance map[string]bool`. |
| `NewStudentClass` | `() StudentClass` | Returns an empty class. |
| `CreateClass` | `(className string, students []string) StudentClass` | Enrols students with `false` attendance. |
| `TakeAttendance` | `(class StudentClass, students []string)` | Marks the given students present (`true`). |
| `CheckAttendance` | `(class StudentClass) []string` | Returns the names of absentees. |
| `EqualIgnortOrderSort` | `(a, b []string) bool` | Order-insensitive slice equality via sorting. |

## Tests (`simple_map/service_test.go`)

* `TestCreateClass` — verifies class name and that everyone starts absent.
* `TestAttendance` — builds Mathematics/Physics/Biology classes, takes
  attendance, and compares absentees using `EqualIgnortOrderSort`.

## Commands

```bash
make test         # runs tests in all subdirectories
make bench-test   # runs benchmarks
```

## Status

All required behaviour is implemented and tested; `cmd/main.go` is still an
empty `main()`.
