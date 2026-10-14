# 🚀 Web Service Gin

A small in-memory album REST API built with [Gin](https://gin-gonic.com),
based on the official Gin quickstart.

## Folder Structure

```
1-Web-Service-Gin/
├── main.go        # Album model, routes and handlers
├── main_test.go   # Test scaffolding (TestMain sets Gin test mode)
└── README.md
```

## Files

| File | Contents |
| --- | --- |
| `main.go` | `Album` struct, an in-memory `albums` slice (3 seeded records), `main()` that registers routes and runs the server, and the handlers `getAlbums`, `postAlbums`, `getAlbumByID`. |
| `main_test.go` | `TestMain` sets `gin.SetMode(gin.TestMode)` before running tests; `TestGetAllAlbums` is an empty placeholder. |

## Routes

| Method | Path | Handler | Description |
| --- | --- | --- | --- |
| `GET` | `/albums` | `getAlbums` | Returns all albums as indented JSON. |
| `GET` | `/albums/:id` | `getAlbumByID` | Returns the album matching `:id`, or `404` if not found. |
| `POST` | `/albums` | `postAlbums` | Binds an album from the JSON body and appends it. |

## Running

```bash
go run main.go
```

The server listens on `localhost:8282`.

Example requests:

```bash
curl localhost:8282/albums
curl localhost:8282/albums/1
curl -X POST localhost:8282/albums \
  -H 'Content-Type: application/json' \
  -d '{"id":"4","title":"Kind of Blue","artist":"Miles Davis","price":19.99}'
```

## Status

* The three routes are implemented and functional.
* `main_test.go` only sets up test mode; there are no assertions yet.
