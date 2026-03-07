# AGENTS.md
Repository guidance for coding agents working in this project.

## Repository Snapshot
- Framework: Astro (static output).
- Purpose: Personal blog/site with posts, pages, and RSS.
- Content source: markdown in `_posts/` and `src/content/pages/`.
- URL requirement: keep blog post routes as `/blog/:title`.
- Legacy stack has been removed from active runtime.

## Rules Files Check
- Cursor rules directory `.cursor/rules/`: not present.
- Cursor root file `.cursorrules`: not present.
- Copilot instructions `.github/copilot-instructions.md`: not present.

If these files are added later, incorporate their rules into this document.

## Directory Conventions
- `src/layouts/`: Astro page shell/layout components.
- `src/components/`: shared UI components (`Header`, `Footer`).
- `src/lib/`: data-loading/parsing utilities.
- `src/pages/`: route files (`/`, `/about/`, `/projects/`, `/blog/[title]`, `/feed.xml`).
- `src/styles/`: Sass used by Astro build.
- `src/content/pages/`: markdown content for static pages.
- `_posts/`: canonical blog markdown content used for post generation.
- `css/`: asset CSS used by site styling.

## Setup Commands
Run from repo root:
```bash
bun install
```

## Build, Lint, and Test Commands
Primary development commands:
```bash
bun run dev
bun run build
bun run preview
```

Current quality gate:
- `bun run build` must pass.

Linting and tests:
- No dedicated linter configured yet.
- No automated unit/integration test suite configured yet.

Recommended verification flow before handoff:
```bash
bun run build
bun run preview
```

## Running a Single Test
Not currently applicable (no test framework configured).

If a test framework is added later, document framework-native single-test commands, for example:
- Vitest: `bunx vitest run path/to/file.test.ts -t "test name"`
- Bun test runner: `bun test path/to/file.test.ts --test-name-pattern "test name"`

## Code Style Guidelines

### General Editing
- Make minimal, focused diffs.
- Keep route behavior stable unless explicitly requested.
- Preserve existing architecture and file placement.
- Avoid unrelated refactors in feature/bugfix changes.
- Use ASCII unless a file already requires Unicode content.

### Imports and Modules
- Project uses ESM (`"type": "module"` in `package.json`).
- Prefer explicit named imports over large default utility imports.
- Group imports by source type (node built-ins, packages, local files).
- Keep import paths stable and relative within `src/`.

### Formatting
- Match existing formatting in neighboring files.
- Keep Astro frontmatter concise and top-loaded.
- Use readable multiline formatting for arrays/objects in route generation.
- Do not reformat entire files unless required by the change.

### Types
- Codebase currently uses JavaScript, not TypeScript.
- Do not introduce TypeScript in unrelated changes.
- If TypeScript is introduced later, scope it to a dedicated migration.

### Naming Conventions
- Route/page files should map clearly to URL intent.
- CSS class names should remain lowercase kebab-case.
- Content slugs should remain lowercase and hyphenated.
- Keep `/blog/:title` compatibility by deriving title from legacy post filename slug.

### Error Handling
- Prefer failing fast at build time over silently skipping malformed content.
- For content parsing failures, surface actionable errors with file context.
- Do not swallow exceptions in build-path utilities unless there is a safe fallback.

### Astro and Content Rules
- Keep page metadata (`title`, canonical, description) consistent across pages.
- Reuse layout/components instead of duplicating shared markup.
- Keep RSS generation in sync with published post list.
- Preserve existing post ordering by date (newest first).
- Honor `published: false` behavior for unpublished posts.

### Markdown and Frontmatter
- Posts/pages must contain valid YAML frontmatter between `---` lines.
- Keep frontmatter keys stable (`title`, `date`, `categories`, `published`).
- Do not mass-normalize historical post filenames unless asked.
- Avoid editorial rewriting of post content during technical changes.

### Sass and CSS
- Keep `src/styles/styles.sass` as the primary Astro style entry.
- Reuse existing Sass variables/mixins before introducing new patterns.
- Avoid introducing UI frameworks for minor style updates.

## URL and Routing Requirements
- Blog post URLs must remain `/blog/:title`.
- `:title` must match legacy slug derived from `_posts` filename.
- Keep `about` and `projects` pages at `/about/` and `/projects/`.
- Keep RSS route at `/feed.xml`.

## Legacy Tooling Notes
- Jekyll and Grunt runtime files have been removed.
- `_posts/` remains as the source dataset for blog content and slug compatibility.
- Preserve existing slug behavior when moving content storage in future migrations.

## Agent Workflow Expectations
- Inspect nearby files for local conventions before editing.
- Run the smallest relevant verification commands after edits.
- Call out assumptions and unverified behavior in your handoff.
- Do not commit generated output (`dist/`) unless explicitly requested.
- Do not invent repository policies not present in source or this file.
