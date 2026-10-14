# 🚀 Expo App - First Iteration

## Description

First iteration of a mobile application built with **Expo SDK 54**, **React
Native** and **TypeScript**. The full app lives in the [`mobile/`](mobile)
directory; see [mobile/README.md](mobile/README.md) for setup and run
instructions.

## Folder Structure

```
1-Expo_App_First_Iteration/
├── mobile/      # The Expo / React Native application
├── Makefile     # (empty)
└── README.md
```

## Contents

| Path | Description |
| --- | --- |
| [mobile/](mobile) | Expo application with file-based routing (`expo-router`), TypeScript config, ESLint config, assets and AI/agent guidance files. Configuration: `app.json`, `tsconfig.json` (`@/*` path alias), `eslint.config.js`, `package.json`. |
| `Makefile` | Empty placeholder; the project is driven by npm scripts instead. |

## Getting Started

```bash
cd mobile
npm install
npx expo start
```

See [mobile/README.md](mobile/README.md) for the full command reference
(`npm run android`, `npm run ios`, `npm run web`, `npm run lint`, etc.).
