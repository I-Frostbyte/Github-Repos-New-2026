Project Rules

Framework:
- Expo SDK 54
- React Native
- TypeScript

Styling:
- NativeWind only
- No StyleSheet unless necessary

Components:
- Functional components
- Hooks only

Structure:
- Reusable components
- No duplicated code

Imports:
- Absolute imports

Naming:
- PascalCase components
- camelCase functions

Code:
- Keep functions under 40 lines where practical.
- Prefer composition over large monolithic components.
- Add comments only when the intent isn't obvious.

<!-- 
This isn't required by Continue or Cline, but it's a simple convention that works well because you can explicitly reference it in prompts ("Follow .ai/project.md"). Put in it:

The project's architecture.
Your preferred libraries (Expo SDK 54, NativeWind, React Query, etc.).
Coding conventions.
Folder layout.
Things the AI should avoid (for example, "Don't use inline styles" or "Don't introduce new dependencies without asking").

Treat it as the onboarding document you'd hand to a new teammate. Whether that teammate is a human or an AI, the result is the same: the generated code stays much more consistent as the project grows.
 -->