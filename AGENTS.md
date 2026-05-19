# AGENTS.md
Repository guidance for coding agents working in this project.

## Repository Snapshot
- Framework: **Ard** static site generator (custom, compiled to Go).
- Purpose: Personal blog/site with posts, pages, and RSS.
- Content source: markdown in `posts/` and `pages/`.
- Build output: `dist/` directory.
- SSG source: `site.ard` (Ard) + `ffi.go` (Go FFI externs).

## URL Requirements
- Blog posts: `/blog/:title` (slug derived from filename).
- Static pages: `/about/`, `/projects/`.
- RSS feed: `/feed.xml`.

## Directory Conventions

### Content
- `posts/` — blog post markdown files. Filename format: `YYYY-M-D-slug.md` or `YYYY-MM-DD-slug.markdown`.
- `pages/` — static page markdown (about.md, projects.md). No frontmatter required.

### SSG
- `site.ard` — main SSG program (Ard source). Contains data types, template rendering, build functions, and unit tests.
- `ffi.go` — Go FFI externs: `ParseFrontmatter` (YAML frontmatter → JSON) and `MarkdownToHTML` (markdown → HTML).
- `ard.toml` — Ard project config.
- `templates/layout.html` — full HTML shell with `@@title@@`, `@@content@@`, `@@canonical_url@@`, `@@description@@` placeholders.
- `templates/post-card.html` — post list item partial with `@@url@@`, `@@title@@`, `@@date@@`, `@@excerpt@@` placeholders.

### Static assets
- `css/style.css` — main site styles (layout, typography, responsive). Converted from previous SASS.
- `css/syntax.css` — highlight.js syntax highlighting overrides.
- `dist/` — build output (gitignored).

## Build Commands
Primary development commands (run from repo root):
```bash
ard build site.ard --target go --out site-gen    # compile SSG to native binary
./site-gen                                         # generate dist/
ard test site.ard                                  # run 17 unit tests
```

Package.json scripts:
```bash
bun run build    # builds and runs site-gen
bun run test     # runs ard test
bun run clean    # removes dist/, site-gen, ard-out/
```

Current quality gate:
- `ard test site.ard` must pass (17 tests: slug extraction, date parsing, formatting, templates, XML escaping).
- `ard build site.ard --target go --out site-gen` must succeed.
- `./site-gen` must complete without errors.

## Running a Single Test
Ard's test framework supports filtering:
```bash
ard test site.ard --filter test_extract_slug    # run tests matching "test_extract_slug"
ard test site.ard --filter test_format_date     # run tests matching "test_format_date"
```

## Code Style Guidelines

### General Editing
- Make minimal, focused diffs.
- Keep route behavior stable unless explicitly requested.
- Preserve existing architecture and file placement.
- Avoid unrelated refactors in feature/bugfix changes.

### Ard Source (`site.ard`)
- Import stdlib modules at top: `ard/fs`, `ard/io`, `ard/decode`, `ard/maybe`, `ard/testing`.
- Extern FFI functions use shorthand syntax: `extern fn name(args) RetType = "GoFuncName"`.
- Define `struct` types before functions that use them.
- Functions must be defined before they are called (no hoisting).
- Return types: `Void!Str` for fallible functions, `Str` for pure functions.
- Use `try` for error propagation, `match` on `Result` for error handling.
- End `Void!Str` functions with `Result::ok(())`.
- Use `match` on Bool for conditional returns (not `if` as expression).
- Function calls with string literal args inside `{interpolation}` are now supported (fixed in v0.18.0).
- Use `for item in list` or `for item, index in list` for iteration.
- Use `not` for negation, `or` for logical-or (not `!` or `||`).
- No `return` keyword — last expression is the return value.
- No `continue` keyword — use nested `if` instead.

### Go FFI (`ffi.go`)
- Functions return `(string, error)` for Ard `Str!Str` types.
- Frontmatter parsing returns JSON string that Ard decodes via `ard/decode`.
- Markdown renderer uses Go stdlib only (no external dependencies).
- Keep the extern surface minimal — two functions currently.

### Templates
- Use `@@placeholder@@` syntax (not `{{...}}` — Ard interprets `{}` as interpolation).
- Placeholders: `@@title@@`, `@@content@@`, `@@canonical_url@@`, `@@description@@`.
- Header and footer are inlined in `layout.html`, not separate partials.
- `post-card.html` is a separate partial since it's reused in the index loop.

### Testing
- Tests are co-located in `site.ard` as `test fn` returning `Void!Str`.
- Use `use ard/testing` for `assert(condition, message)`, `pass()`, `fail(message)`.
- Use `try testing::assert(...)` to propagate failures.
- Test pure functions (string manipulation, date formatting, template replacement).
- Integration tests (file I/O, extern calls) are run via the build pipeline.

### CSS
- `css/style.css` is plain CSS (no preprocessor). Converted from prior SASS.
- `css/syntax.css` handles highlight.js code block styles.
- Keep light/dark mode via `@media (prefers-color-scheme: ...)`.
- Use CSS custom properties for theming (`--bg`, `--fg`, `--accent`, etc.).

### Naming Conventions
- Slugs are lowercase hyphenated, derived from post filename.
- CSS class names are lowercase kebab-case.
- Ard function/type names use snake_case.

### Error Handling
- Fallible functions return `Result<T, Str>` and use `try` to propagate.
- Frontmatter parsing errors are converted to `Str` via `decode::flatten(errs)`.
- Build functions surface errors with context ("title: ...", "body: ...").
- The `main` function catches and prints errors: `try build_site() -> err { io::print(err) }`.

## URL and Routing Requirements
- Blog post URLs: `/blog/:title` where `:title` matches the slug from filename.
- About page: `/about/` (rendered from `pages/about.md`).
- Projects page: `/projects/` (rendered from `pages/projects.md`).
- RSS feed: `/feed.xml`.
- Canonical URLs use `https://akonwi.io`.

## Agent Workflow Expectations
- Inspect nearby files for local conventions before editing.
- Run `ard test site.ard` after changes.
- Run `ard build site.ard --target go --out site-gen && ./site-gen` for full verification.
- Call out assumptions and unverified behavior in your handoff.
- Do not commit generated output (`dist/`, `site-gen`, `ard-out/`) unless explicitly requested.
- Updates to FFI (`ffi.go`) require a rebuild of the binary.
- New template placeholders must use `@@name@@` syntax and be added to `apply_layout()` or `apply_post_card()`.
