# 🚀 Github Repos New (2026)

Find all coding challenges, exercises, projects, tutorials and interview challenges here.

A monorepo of hands-on software-engineering practice — primarily **Go** (algorithms,
concurrency, REST/gRPC services) and **TypeScript / React Native** (Expo). Each
top-level folder below is a self-contained category; every item links to its own
folder and notes what it contains and how to run it.

---

## 📚 Table of Contents

- [🗂️ Repository Layout](#️-repository-layout)
- [📚 Books](#-books)
- [🧩 Coding Challenges](#-coding-challenges)
- [✏️ Exercises](#️-exercises)
- [💼 Interviews](#-interviews)
- [🚧 Projects](#-projects)
- [🎓 Tutorials](#-tutorials)
- [⚙️ Continuous Integration](#️-continuous-integration)
- [🛠️ Tooling & Conventions](#️-tooling--conventions)

---

## 🗂️ Repository Layout

| Folder | Purpose |
| --- | --- |
| `Books/` | Implementations that follow along with programming books. |
| `Coding-Challenges/` | Solutions and stubs from online challenge platforms. |
| `Exercises/` | Standalone practice exercises (files, sorting, maps). |
| `Interviews/` | Software-engineering interview-style problems. |
| `Projects/` | Larger, multi-file applications (mobile app, gRPC service). |
| `Tutorials/` | Guided walk-throughs, mostly from *A Tour of Go* and Gin. |
| `.github/` | CI workflows. |

---

## 📚 Books

Implementations and notes that follow along with a published book.

### [Learn Data Structures and Algorithms with Golang](Books/Learn-Data-Structures-and-Algorithms-With-Golang)
- **Author:** Bhagvan Kommadi. **Language:** Go.
- **Searching/** ([README](Books/Learn-Data-Structures-and-Algorithms-With-Golang/Searching/README.md)) — search algorithm implementations.
  - [1-Linear](Books/Learn-Data-Structures-and-Algorithms-With-Golang/Searching/1-Linear) — linear search over a collection; `O(n)`. `main.go` + `main_test.go`.
  - [2-Binary](Books/Learn-Data-Structures-and-Algorithms-With-Golang/Searching/2-Binary) — binary search built on `sort.Search`; `O(log n)`. `main.go` + `main_test.go`.
  - [generator](Books/Learn-Data-Structures-and-Algorithms-With-Golang/Searching/generator) — `Generator.RandomIntegerArrayGenerator` for producing test arrays.
  - Run tests: `make test` ([Makefile](Books/Learn-Data-Structures-and-Algorithms-With-Golang/Searching/Makefile)).
- **Sorting/** ([README](Books/Learn-Data-Structures-and-Algorithms-With-Golang/Sorting/README.md)) — six classic sorting algorithms, each with `main.go` + `main_test.go`.
  - [1-Bubble](Books/Learn-Data-Structures-and-Algorithms-With-Golang/Sorting/1-Bubble) — neighbour-pair swaps; `O(n²)`.
  - [2-Selection](Books/Learn-Data-Structures-and-Algorithms-With-Golang/Sorting/2-Selection) — repeatedly selects the minimum; `O(n²)`.
  - [3-Insertion](Books/Learn-Data-Structures-and-Algorithms-With-Golang/Sorting/3-Insertion) — builds a sorted region one element at a time; `O(n²)`.
  - [4-Shell](Books/Learn-Data-Structures-and-Algorithms-With-Golang/Sorting/4-Shell) — interval/gap-based insertion sort.
  - [5-Merge](Books/Learn-Data-Structures-and-Algorithms-With-Golang/Sorting/5-Merge) — recursive divide-and-conquer; `O(n log n)`.
  - [6-Quick](Books/Learn-Data-Structures-and-Algorithms-With-Golang/Sorting/6-Quick) — pivot partition sort; `O(n log n)`; includes an extra high-level explanation in its README.
  - Run tests: `make test` ([Makefile](Books/Learn-Data-Structures-and-Algorithms-With-Golang/Sorting/Makefile)).

---

## 🧩 Coding Challenges

Solutions sourced from online coding-challenge platforms.

### [Coding-Challenges/](Coding-Challenges/README.md)
- **[Exercism/](Coding-Challenges/Exercism)** — challenges from Exercism.
  - [1-Hello-World](Coding-Challenges/Exercism/1-Hello-World) — placeholder (README only).
- **[Leet-Code/](Coding-Challenges/Leet-Code)** — challenges from LeetCode.
  - [1-Two-Sum](Coding-Challenges/Leet-Code/1-Two-Sum) — the classic *Two Sum* prompt with examples and constraints; `main.go` currently stubs `TwoSum` (`main_test.go` present).

> Note: top-level [Coding-Challenges/README.md](Coding-Challenges/README.md) still lists a `Code-Wars` folder that does not yet exist.

---

## ✏️ Exercises

Self-contained programming exercises with a spec in each README.

- **[1-File_Exercise](Exercises/1-File_Exercise/README.md)** — file management CLI.
  - Implements a `FileManager` with create, rename, write, read, delete and copy.
  - Demo moves files between `source/` and `destination/` folders; works on `.txt`.
  - Files: `main.go`, `go.mod`, sample `source/` and `destination/` files.
- **[2-Sorting_Exercise_Sorting_Program](Exercises/2-Sorting_Exercise_Sorting_Program/README.md)** — sorting library.
  - `sort/service.go` implements `ArraySort` with Bubble, Selection, Insertion, Shell, Merge and Quick sort (plus helpers).
  - Structure: `cmd/main.go` (entry), `sort/` (logic + tests, 190-line test suite).
  - Run: `make test` / `make run` / `make install-dependencies` ([Makefile](Exercises/2-Sorting_Exercise_Sorting_Program/Makefile)).
  - Covered by CI (see below).
- **[3-Simple_Map_Exercise](Exercises/3-Simple_Map_Exercise/README.md)** — classroom attendance map.
  - `simple_map/service.go` implements `StudentClass` with `CreateClass`, `TakeAttendance`, `CheckAttendance`.
  - Structure: `cmd/main.go` (entry), `simple_map/` (logic + tests).
  - Run: `make test` / `make bench-test` ([Makefile](Exercises/3-Simple_Map_Exercise/Makefile)).

---

## 💼 Interviews

Interview-style problems and take-home simulations.

- **[SWE/](Interviews/SWE)** — in-memory file storage service.
  - `main.go` defines `FileStorage` with `Upload`, `GetMetadata`, `Download` and `Delete`.
  - Generates UUIDs (`github.com/google/uuid`), reads file metadata and contents.
  - Module: `go.mod` / `go.sum`.

> [Interviews/README.md](Interviews/README.md) is currently an empty placeholder.

---

## 🚧 Projects

Larger multi-file applications.

### [1-Expo_App_First_Iteration](Projects/1-Expo_App_First_Iteration)
- **Type:** Mobile app (Expo). Root [README](Projects/1-Expo_App_First_Iteration/README.md) is a placeholder; see the app README.
- **[mobile/](Projects/1-Expo_App_First_Iteration/mobile/README.md)** — Expo SDK 54 + React Native + TypeScript app.
  - File-based routing via **expo-router** (`app/_layout.tsx`, `app/index.tsx`).
  - Config: `app.json`, `tsconfig.json` (path alias `@/*`), `eslint.config.js`, `package.json`.
  - AI/agent guidance: `.ai/project.md`, `AGENTS.md`, `CLAUDE.md`, `.claude/settings.json`, `.vscode/`.
  - Assets in `assets/images/`; generated `app-example/` starter kept for reference.
  - Run: `npm install` then `npx expo start` (or `npm run android|ios|web`); lint with `npm run lint`.

### [2-gRPC_Simple_Inventory_Tracker](Projects/2-gRPC_Simple_Inventory_Tracker/README.md)
- **Type:** Go gRPC backend for an inventory tracker.
- **protobuf/** — `users.proto` (User, SignUpRequest/Response, `UsersService.CreateUser`), a
  `Makefile` to lint (`clang-format`) and generate stubs, and a `Dockerfile`.
- **grpc/go/usersgrpc/** — generated stubs: `users.pb.go`, `users_grpc.pb.go`.
- **users/** — service implementation.
  - `cmd/grpc-service/main.go` — server bootstrap: config loading, zerolog logging,
    panic-recovery interceptor, reflection, graceful shutdown (signal handling + `errgroup`).
  - `users/service.go` — `Impl` (embeds `UnimplementedUsersServiceServer`); `CreateUser` is a stub.
  - `public/model.go` — env-driven `Config` via `ardanlabs/conf` + `godotenv`.
  - `.env` — sample config (listen port, PostgreSQL connection, migration path).
- **inventories/** — currently empty (planned feature).
- Run: `make test` / `make bench-test` ([Makefile](Projects/2-gRPC_Simple_Inventory_Tracker/Makefile)).

> [Projects/README.md](Projects/README.md) is currently an empty placeholder.

---

## 🎓 Tutorials

Guided walk-throughs, mostly from the official *A Tour of Go* and Gin docs.

- **[1-A-Tour-of-Go-Concurrency](Tutorials/1-A-Tour-of-Go-Concurrency/README.md)** — Go concurrency tour.
  - [1-Goroutines](Tutorials/1-A-Tour-of-Go-Concurrency/1-Goroutines) — spawning goroutines, `sync.WaitGroup`, sorting in a goroutine.
  - [2-Channels](Tutorials/1-A-Tour-of-Go-Concurrency/2-Channels) — typed channels; parallel array-sum.
  - [3-Buffered_Channels](Tutorials/1-A-Tour-of-Go-Concurrency/3-Buffered_Channels) — buffered channels with capacity.
  - [4-Range_and_Close](Tutorials/1-A-Tour-of-Go-Concurrency/4-Range_and_Close) — closing channels and `range` over them (Fibonacci generator).
  - `main.go` wires all examples together. Run: `make run` / `make test` / `make bench-test` ([Makefile](Tutorials/1-A-Tour-of-Go-Concurrency/Makefile)).
- **[2-A-Tour-of-Go-Methods_and_Interfaces](Tutorials/2-A-Tour-of-Go-Methods_and_Interfaces/README.md)** — methods & interfaces tour.
  - [1-Methods](Tutorials/2-A-Tour-of-Go-Methods_and_Interfaces/1-Methods) — value receivers, custom types (`Vertex`, `User`, `Age`).
  - Run: `make run` / `make test` ([Makefile](Tutorials/2-A-Tour-of-Go-Methods_and_Interfaces/Makefile)).
- **[3-Restful-Api-Golang](Tutorials/3-Restful-Api-Golang/README.md)** — REST API with Go.
  - [1-Web-Service-Gin](Tutorials/3-Restful-Api-Golang/1-Web-Service-Gin) — Gin album service (`GET /albums`, `GET /albums/:id`, `POST /albums`) on port `:8282`; `main.go` + `main_test.go`.
  - Run: `make test` ([Makefile](Tutorials/3-Restful-Api-Golang/Makefile)).
- **[4-Simple-Gin-Server](Tutorials/4-Simple-Gin-Server/README.md)** — Gin server basics.
  - [1-Using-HTTP-Method](Tutorials/4-Simple-Gin-Server/1-Using-HTTP-Method) — route handlers for GET/POST/PUT/PATCH/DELETE/HEAD/OPTIONS.
  - [2-Parameters-in-path](Tutorials/4-Simple-Gin-Server/2-Parameters-in-path) — `:name` and `*action` path parameters.
  - `main.go` runs three servers concurrently (`:8080`, `:5050`, `:5151`) via `errgroup`. Run: `make run` ([Makefile](Tutorials/4-Simple-Gin-Server/Makefile)).

> [Tutorials/README.md](Tutorials/README.md) is currently an empty placeholder.

---

## ⚙️ Continuous Integration

- **[.github/workflows/sorting-ci.yaml](.github/workflows/sorting-ci.yaml)** — *Sorting Program CI*.
  - Triggers on `push` and `pull_request` affecting `Exercises/2-Sorting_Exercise_Sorting_Program/**`.
  - Uses Go `1.26.5`, runs `make install-dependencies` then `make test`.

---

## 🛠️ Tooling & Conventions

- **Languages:** Go (searching/sorting, concurrency, Gin, gRPC) and TypeScript / React Native (Expo).
- **Go modules:** Each Go project/exercise is its own module (`go.mod`), using the module
  path prefix `github.com/I-Frostbyte/Github-Repos-New-2026/...`.
- **Task runner:** Directories include `Makefile`s with `test`, `bench-test`, and/or `run` targets.
  - `test` — `go test -v ./...` (simple) or a recursive `find ... -name "*_test.go"` (multi-package).
  - `bench-test` — runs benchmarks.
  - `run` — `go run main.go`.
- **README style:** Most folders carry a `# 🚀 Title` heading with `## Description` and, where
  relevant, a `## Folder Structure` list.
- **Note:** Several READMEs (`Interviews/`, `Projects/`, `Tutorials/`, and a few sub-folders) are
  intentionally empty placeholders awaiting content.
