# 🚀 RESTFUL API Golang

-----------------------------------------------------------------------------------------------------------------------------
RESTFUL API Golang
- Web Service Gin
-----------------------------------------------------------------------------------------------------------------------------

## Folder Structure

```
3-Restful-Api-Golang/
├── 1-Web-Service-Gin/   # Gin album web service
├── go.mod               # Module: .../Tutorials/2-RESTFUL-API-Golang
├── go.sum
├── Makefile             # test / bench-test
└── README.md
```

## Contents

| Tutorial | Folder | Description |
| --- | --- | --- |
| Web Service Gin | [1-Web-Service-Gin](1-Web-Service-Gin) | A small in-memory album REST API built with Gin (`GET`/`POST /albums`). |

## Notes

* The module path in `go.mod` is `.../Tutorials/2-RESTFUL-API-Golang`, which does
  not match the current folder name `3-Restful-Api-Golang` (a leftover from a
  rename).

## Commands

```bash
make test         # tests in all sub-directories
make bench-test   # benchmarks in all sub-directories
```
