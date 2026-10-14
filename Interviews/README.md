# 🚀 Interviews

A sub-folder containing software-engineering interview-style problems and
take-home simulations.

## Folder Structure

```
Interviews/
├── SWE/        # Software-engineering interview problem
└── README.md
```

| Folder | Description | Tech |
| --- | --- | --- |
| [SWE](SWE) | An in-memory file-storage service (`Upload`, `GetMetadata`, `Download`, `Delete`) built as a take-home-style problem. | Go |

Each interview problem lives in its own module (`go.mod`) and is numbered by
role or track where applicable.

## SWE

```
SWE/
├── main.go    # FileStorage service + demo main()
├── go.mod     # Module: github.com/I-Frostbyte/SWE (Go 1.26.5)
├── go.sum
└── README.md  # (not present)
```

`main.go` implements an in-memory `FileStorage`:

| Method | Signature | Purpose |
| --- | --- | --- |
| `NewFileStorage` | `(uploadedFiles []file) *FileStorage` | Constructor. |
| `Upload` | `(ctx, filePath) (string, error)` | Reads a file and stores it, returning a UUID. |
| `GetMetadata` | `(ctx, fileID) (file, error)` | Returns the stored file record for an ID. |
| `Download` | `(ctx, fileID, destinationPath) error` | Recreates the file at a destination path. |
| `Delete` | `(ctx, fileID) error` | Removes a file; returns `ErrNotFound` if missing. |

Dependencies: `github.com/google/uuid`. There are no tests yet.
