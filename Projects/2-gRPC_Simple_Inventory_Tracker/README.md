# 🚀 Simple Inventory Tracker Project #1

-----------------------------------------------------------------------------------------------------------------------------
Simple Inventory Tracker #1
- Create an inventory for a business.
- Users can add items to their inventories.
- Inventories are totalled and the price is stored.
- System can keep track of user's favorite items and spend limit.
-----------------------------------------------------------------------------------------------------------------------------

## Overview

A Go **gRPC** backend for an inventory tracker. The service currently exposes a
single `UsersService` RPC (`CreateUser`) defined in Protocol Buffers; the
inventory-specific logic is planned (`inventories/` is an empty placeholder).

## Tech Stack

| Area | Choice |
| --- | --- |
| Language | Go 1.26.5 |
| RPC | gRPC (`google.golang.org/grpc`) + Protocol Buffers |
| Config | `ardanlabs/conf` + `joho/godotenv` |
| Logging | `rs/zerolog` |
| Middleware | `grpc-ecosystem/go-grpc-middleware` (panic recovery) |

## Folder Structure

```
2-gRPC_Simple_Inventory_Tracker/
├── grpc/go/usersgrpc/
│   ├── users.pb.go          # Generated protobuf messages (DO NOT EDIT)
│   └── users_grpc.pb.go     # Generated gRPC client/server stubs (DO NOT EDIT)
├── inventories/             # Empty — inventory feature planned
├── protobuf/
│   ├── users.proto          # Service + message definitions
│   ├── Makefile             # lint / generate-protos targets
│   └── Dockerfile           # (empty)
├── users/
│   ├── cmd/grpc-service/
│   │   └── main.go          # Server bootstrap (config, logging, gRPC, shutdown)
│   ├── public/
│   │   └── model.go         # Env-driven Config / DBConfig loading
│   ├── users/
│   │   ├── service.go       # Impl (CreateUser currently a stub)
│   │   └── service_test.go  # (package declaration only)
│   ├── .env                 # Sample environment configuration
│   └── Makefile             # (empty)
├── go.mod                   # Module: .../Projects/2-gRPC_Simple_Inventory_Tracker
├── go.sum
├── Makefile                 # test / bench-test targets
└── README.md
```

## Service Definition (`protobuf/users.proto`)

* **Enum:** `UserRole` — `USER_ROLE_UNSPECIFIED`, `USER_ROLE_ADMIN`, `USER_ROLE_USER`.
* **Messages:** `User` (id, username, password, phone_number, user_role,
  created_at, updated_at), `SignUpRequest`, `SignUpResponse`.
* **Service:** `UsersService` with `rpc CreateUser(SignUpRequest) returns (SignUpResponse)`.

Generated code lives in `grpc/go/usersgrpc/` and is produced by
`make generate-protos` (uses `protoc`, `protoc-gen-go`, `protoc-gen-go-grpc`).

## Server (`users/cmd/grpc-service/main.go`)

`main()` sets up a cancellable context (SIGINT/SIGTERM), builds a zerolog
logger, then `run()`:

1. Loads configuration (`model.Config.LoadConfig`).
2. Adjusts the log level from config.
3. Creates a gRPC server with a recovery unary interceptor and reflection.
4. Registers `UsersService` (`users.NewUsersService()`).
5. Listens on `config.ListenPort` and serves in a goroutine.
6. Shuts down gracefully on context cancellation via `GracefulStop`.

The file also documents the run lifecycle and server mental model in comments.

## Configuration (`users/public/model.go`)

`Config` is populated from environment variables / `.env`:

| Field | Env var | Notes |
| --- | --- | --- |
| `ListenPort` | `LISTEN_PORT` | Required. |
| `MigrationPath` | `MIGRATION_PATH` | Required. |
| `LogLevel` | `LOG_LEVEL` | Required. |
| `DB.DBUser` | `DB_USER` | Required. |
| `DB.DBPassword` | `DB_PASSWORD` | Required. |
| `DB.DBHost` | `DB_HOST` | Required. |
| `DB.DBPort` | `DB_PORT` | Required. |
| `DB.DBName` | `DB_NAME` | Required. |
| `DB.TLSDisabled` | `DB_TLS_DISABLED` | Optional. |

The included `.env` sets `LISTEN_PORT=40042` and a local PostgreSQL config.

## Commands

```bash
# From the project root
make test          # runs tests in all sub-directories
make bench-test    # runs benchmarks

# From protobuf/
make lint            # clang-format check on *.proto
make generate-protos # regenerate Go stubs into ../grpc/go
```

## Status

| Area | State |
| --- | --- |
| Protobufs & generated stubs | ✅ Present |
| gRPC server bootstrap | ✅ Implemented |
| Config loading | ✅ Implemented |
| Database connection | 🚧 Commented out (`pgxpool` code is disabled) |
| `CreateUser` handler | 🚧 Stub — `panic("unimplemented")` |
| Inventory feature | 🚧 Not started (`inventories/` empty) |
