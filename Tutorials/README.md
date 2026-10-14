# 🚀 Tutorials

Guided walk-throughs, mostly following *A Tour of Go* and the Gin documentation.
Each tutorial is a self-contained Go module (`go.mod`) with a `Makefile`.

## Folder Structure

```
Tutorials/
├── 1-A-Tour-of-Go-Concurrency/             # Goroutines, channels, range/close
├── 2-A-Tour-of-Go-Methods_and_Interfaces/  # Methods on types
├── 3-Restful-Api-Golang/                   # REST API with Gin
├── 4-Simple-Gin-Server/                    # Gin routing basics
└── README.md
```

## Contents

| Tutorial | Source | Topics | Files |
| --- | --- | --- | --- |
| [1-A-Tour-of-Go-Concurrency](1-A-Tour-of-Go-Concurrency) | A Tour of Go | Goroutines, Channels, Buffered Channels, Range and Close | `main.go`, four topic sub-packages, `Makefile` |
| [2-A-Tour-of-Go-Methods_and_Interfaces](2-A-Tour-of-Go-Methods_and_Interfaces) | A Tour of Go | Methods | `main.go`, `1-Methods/`, `Makefile` |
| [3-Restful-Api-Golang](3-Restful-Api-Golang) | Gin docs | Building a REST API | `1-Web-Service-Gin/`, `Makefile` |
| [4-Simple-Gin-Server](4-Simple-Gin-Server) | Gin docs | HTTP methods & routing | `main.go`, `1-Using-HTTP-Method/`, `Makefile` |

## Common Commands

```bash
make run          # go run main.go (where a run target exists)
make test         # run tests in sub-directories
make bench-test   # run benchmarks
```
