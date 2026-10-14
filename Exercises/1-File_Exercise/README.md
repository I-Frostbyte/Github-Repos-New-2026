# 🚀 File Exercise #1

-----------------------------------------------------------------------------------------------------------------------------
File Exercise #1
- Create a simple program that allows for the creation, deletion, rename, editing and copying (single & bulk) of files in between two folders (source & destination).
- For simplicity, files so far should be .txt

- CreateFile(filename, filetype string) error {}
- RenameFile(filename, newFilename string) error {}
- EditFile(filename string, content []bytes) error {}
- DeleteFile(filename string) error {}
- CopyFile(filename, source, destination string) (string, error) {}
-----------------------------------------------------------------------------------------------------------------------------

## Folder Structure

```
1-File_Exercise/
├── main.go                      # FileManager implementation + demo main()
├── go.mod                       # Module: .../Exercises/FileExercise-1 (Go 1.26.5)
├── source/                      # Source .txt files
│   ├── test-1.txt
│   ├── test-2.txt
│   ├── test-3.txt
│   ├── test-4.txt
│   ├── test-5-renamed.txt
│   ├── test-6-renamed.txt
│   └── test-7-renamed.txt
├── destination/                 # Copied .txt files
│   ├── test-7-renamed.txt
│   └── test-8-renamed.txt
└── README.md
```

## Implementation

The demo in `main.go` creates a file, writes to it, reads it back, renames it,
copies it from `source/` to `destination/`, and finally deletes the source copy.

`FileManager` methods actually implemented (note: these differ from the
signatures sketched in the spec above):

| Method | Signature | Purpose |
| --- | --- | --- |
| `NewFileManager` | `(file *os.File) *FileManager` | Constructor. |
| `CreateFile` | `(filename string) (*os.File, error)` | Creates `source/<filename>.txt`. |
| `RenameFile` | `(oldfilePath, newFilePath string) error` | Renames a file via `os.Rename`. |
| `WriteToFile` | `(file *os.File, content []byte) error` | Writes bytes to an open file. |
| `ReadFromFile` | `(file *os.File) (string, error)` | Seeks to start and reads the file. |
| `DeleteFile` | `(filename string) error` | Removes a file via `os.Remove`. |
| `CopyFile` | `(sourcePath, destinationPath string) error` | Streams a copy with `io.Copy`. |
| `RenameFilePath` | `(file *os.File) (string, string, error)` | Helper building the `-renamed` path. |

## Spec vs. Implementation

* The spec proposes an `EditFile(filename, content)` method; the code implements
  the equivalent as `WriteToFile(file, content)`.
* The spec's `CopyFile(filename, source, destination)` differs from the
  implemented `CopyFile(sourcePath, destinationPath)`.
* Bulk copy is described in the spec but not yet implemented.

## Notes

* Only `.txt` files are handled, as per the exercise constraints.
* There is no `Makefile` in this exercise; run it with `go run main.go`.
