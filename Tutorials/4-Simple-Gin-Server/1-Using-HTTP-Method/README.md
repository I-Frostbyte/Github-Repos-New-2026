# 🚀 Using HTTP Method

Gin provides methods that map directly to HTTP verbs, making it straightforward to build RESTful APIs. Each method registers a route that responds only to the corresponding HTTP request type:

| Method | Typical REST Usage |
| --- | --- |
| GET | Retrieve a resource |
| POST | Create a new resource |
| PUT | Replace an existing resource |
| PATCH | Partially update an existing resource |
| DELETE | Remove a resource |
| HEAD | Same as GET but without a body |
| OPTIONS | Describe communication options |

## Folder Structure

```
1-Using-HTTP-Method/
├── http_methods.go   # HTTPMethods() router + per-verb handlers
└── README.md
```

## Files

| File | Contents |
| --- | --- |
| `http_methods.go` | `HTTPMethods() *gin.Engine` builds a `gin.Default()` router and registers the route handlers `getting`, `posting`, `putting`, `deleting`, `patching`, `head`, `options`. |

## Routes registered by `HTTPMethods()`

| Route | Handler | Response |
| --- | --- | --- |
| `GET /someGet` | `getting` | `{"method":"GET"}` |
| `GET /somePost` | `posting` | `{"method":"POST"}` |
| `GET /somePut` | `putting` | `{"method":"PUT"}` |
| `GET /someDelete` | `deleting` | `{"method":"DELETE"}` |
| `GET /somePatch` | `patching` | `{"method":"PATCH"}` |
| `GET /someHead` | `head` | `200 OK` (no body) |
| `GET /someOptions` | `options` | `200 OK` (no body) |

## API

```go
// HTTPMethods returns a Gin engine with the demonstration routes registered.
func HTTPMethods() *gin.Engine
```

## Notes

* Although the handlers are named after different HTTP verbs, every route is
  registered with `router.GET(...)`. This mirrors the tutorial's illustrative
  intent rather than strict REST semantics (`router.POST`, `router.PUT`, etc.
  would be the idiomatic registrations).
* This router is mounted as the second server (`:5050`) in
  [../main.go](../main.go).
