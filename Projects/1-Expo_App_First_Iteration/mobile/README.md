# 🚀 mobile

A [React Native](https://reactnative.dev) application built with
[Expo](https://expo.dev) (SDK 54) and TypeScript, using
[expo-router](https://docs.expo.dev/router/introduction) for file-based routing.

> ⚠️ Expo SDK 54 has changed significantly. Read the exact versioned docs at
> <https://docs.expo.dev/versions/v54.0.0/> before writing any code.

## Tech Stack

| Area | Choice |
| --- | --- |
| Platform | Expo SDK 54, React Native 0.81, React 19 |
| Language | TypeScript (strict) |
| Routing | expo-router (file-based) |
| Lint | ESLint via `eslint-config-expo` (flat config) |

## Folder Structure

```
mobile/
├── app/                     # expo-router routes (file-based)
│   ├── _layout.tsx          # Root <Stack /> layout
│   └── index.tsx            # Home screen
├── assets/
│   └── images/              # App icons, splash and logo assets
├── .ai/
│   └── project.md           # Project rules/conventions for AI assistants
├── .claude/
│   └── settings.json        # Claude plugin configuration
├── .vscode/
│   ├── extensions.json      # Recommended: expo.vscode-expo-tools
│   └── settings.json        # Editor code-action settings
├── app.json                 # Expo app configuration
├── tsconfig.json            # TypeScript config (path alias @/* -> ./*)
├── eslint.config.js         # ESLint (flat) config
├── package.json             # Dependencies and npm scripts
├── package-lock.json
├── expo-env.d.ts            # Generated Expo typings
├── AGENTS.md                # Agent instructions
├── CLAUDE.md                # Points to AGENTS.md
├── .gitignore
└── README.md
```

### Screens

| File | Route | Description |
| --- | --- | --- |
| `app/_layout.tsx` | — | Root layout; renders the expo-router `<Stack />`. |
| `app/index.tsx` | `/` | Placeholder home screen ("Edit app/index.tsx to edit this screen."). |

### Configuration Notes

* `app.json` — app name/slug `mobile`, scheme `mobile`, new architecture
  enabled, `typedRoutes` and `reactCompiler` experiments on, splash screen and
  adaptive icons configured.
* `tsconfig.json` — extends `expo/tsconfig.base`, `strict: true`, and maps `@/*`
  to the project root.
* Assets include `icon.png`, `favicon.png`, `splash-icon.png`, Android adaptive
  icon layers and React logo PNGs.

## Get started

1. Install dependencies

   ```bash
   npm install
   ```

2. Start the app

   ```bash
   npx expo start
   ```

In the output, you'll find options to open the app in a

- [development build](https://docs.expo.dev/develop/development-builds/introduction/)
- [Android emulator](https://docs.expo.dev/workflow/android-studio-emulator/)
- [iOS simulator](https://docs.expo.dev/workflow/ios-simulator/)
- [Expo Go](https://expo.dev/go), a limited sandbox for trying out app development with Expo

You can start developing by editing the files inside the **app** directory. This project uses [file-based routing](https://docs.expo.dev/router/introduction).

## Available Scripts

| Script | Command |
| --- | --- |
| `npm start` | `expo start` |
| `npm run android` | `expo start --android` |
| `npm run ios` | `expo start --ios` |
| `npm run web` | `expo start --web` |
| `npm run lint` | `expo lint` |
| `npm run reset-project` | Moves starter code to `app-example/` and blanks `app/`. |

## Get a fresh project

When you're ready, run:

```bash
npm run reset-project
```

This command will move the starter code to the **app-example** directory and create a blank **app** directory where you can start developing.

## Learn more

To learn more about developing your project with Expo, look at the following resources:

- [Expo documentation](https://docs.expo.dev/): Learn fundamentals, or go into advanced topics with our [guides](https://docs.expo.dev/guides).
- [Learn Expo tutorial](https://docs.expo.dev/tutorial/introduction/): Follow a step-by-step tutorial where you'll create a project that runs on Android, iOS, and the web.

## Join the community

Join our community of developers creating universal apps.

- [Expo on GitHub](https://github.com/expo/expo): View our open source platform and contribute.
- [Discord community](https://chat.expo.dev): Chat with Expo users and ask questions.
