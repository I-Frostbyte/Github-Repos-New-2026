# 🚀 Simple Gin Server

A minimal [Gin](https://gin-gonic.com) server demonstrating routing basics. This
module runs two HTTP servers concurrently using `errgroup`.

## Folder Structure

```
4-Simple-Gin-Server/
├── 1-Using-HTTP-Method/   # Per-verb route handlers
├── main.go                # Runs two servers (:8080 and :5050)
├── go.mod                 # Module: .../Tutorials/4-Simple-Gin-Server
├── go.sum
├── Makefile               # run
└── README.md
```

## Contents

| Tutorial | Folder | Description |
| --- | --- | --- |
| Using HTTP Method | [1-Using-HTTP-Method](1-Using-HTTP-Method) | Registers routes for GET/POST/PUT/PATCH/DELETE/HEAD/OPTIONS. |

## `main.go`

* Server 1 (`:8080`) — a `gin.Default()` router with a `GET /ping` endpoint
  returning `{"message":"pong"}`.
* Server 2 (`:5050`) — the router returned by `http_methods.HTTPMethods()`.

Both are started as goroutines inside a `golang.org/x/sync/errgroup.Group` and
the program waits for either to finish or error.

## Running

```bash
make run          # go run main.go
```

```bash
curl localhost:8080/ping
curl localhost:5050/someGet
```

## Notes

* This folder previously contained a second topic (`2-Parameters-in-path`);
  it is not present in the current revision — only `1-Using-HTTP-Method` exists.
