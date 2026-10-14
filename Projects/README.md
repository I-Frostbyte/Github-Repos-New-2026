# 🚀 Projects

A sub-folder containing larger, multi-file applications.

## Folder Structure

```
Projects/
├── 1-Expo_App_First_Iteration/             # Expo / React Native mobile app
├── 2-gRPC_Simple_Inventory_Tracker/        # Go gRPC inventory service
└── README.md
```

## Contents

| Project | Type | Tech | Description |
| --- | --- | --- | --- |
| [1-Expo_App_First_Iteration](1-Expo_App_First_Iteration) | Mobile app | Expo SDK 54, React Native, TypeScript | First iteration of an Expo app using file-based routing (`expo-router`). See [mobile/README.md](1-Expo_App_First_Iteration/mobile/README.md). |
| [2-gRPC_Simple_Inventory_Tracker](2-gRPC_Simple_Inventory_Tracker) | Backend service | Go, gRPC, Protocol Buffers | An inventory tracker exposing a `UsersService` over gRPC, with protobuf definitions, config loading and graceful shutdown. |

Each project is self-contained: the mobile app is an npm/Expo project, and the
inventory tracker is a Go module with a `Makefile`.
