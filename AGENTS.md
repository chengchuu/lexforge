# AGENTS.md

## Project

Project name:

```text
LexForge
```

Repository name:

```text
lexforge
```

LexForge is a Git-native personal multilingual vocabulary system.

It turns vocabulary encountered during everyday technical work into structured entries for export to Anki and review in AnkiDroid.

Primary learning direction:

```text
Word → Meaning + Context
```

---

## User Background

Write for an experienced full-stack developer with a background in:

- Go
- Node.js
- PHP
- Bash
- PowerShell
- JavaScript
- TypeScript
- React.js
- Vue.js
- Bootstrap
- jQuery
- HTML
- CSS
- Git
- GitHub

Explain project-specific behavior; avoid beginner explanations unless requested.

Prefer:

- concise solutions
- production-ready implementations
- idiomatic code
- established community conventions
- maintainable designs
- clear trade-offs
- deterministic behavior
- practical tooling

---

## Core Principles

Always follow these principles.

### 1. Git Is the Source of Truth

Vocabulary source files are canonical. Generate exports from those files; do not edit generated output manually.

Conceptually:

```text
entries/*.yml
    ↓
validate
    ↓
build
    ↓
dist/*
```

---

### 2. One Vocabulary Item per Source File

Prefer:

```text
entries/idiomatic.yml
entries/idempotent.yml
entries/race-condition.yml
```

over one large source file.

This reduces merge conflicts and keeps Git history easy to review.

---

### 3. Keep the Project Simple

Do not introduce unnecessary:

- frameworks
- databases
- services
- web applications
- state-management layers
- abstractions
- dependency-injection systems
- plugin systems
- configuration layers

Add complexity only when there is a concrete requirement.

---

### 4. Use Anki for Learning State

LexForge should not implement:

- spaced repetition
- review scheduling
- review history
- cloud synchronization
- mobile review

Anki, AnkiWeb, and AnkiDroid already provide these capabilities.

LexForge manages vocabulary content and generated import data.

---

### 5. AI-Assisted, Not AI-Dependent

Codex may generate and enrich vocabulary entries, but the repository must remain understandable and maintainable without AI.

Do not store essential information only in generated prose or hidden AI state.

---

## Expected Repository Structure

Prefer the following structure unless there is a strong reason to change it:

```text
lexforge/
├── README.md
├── AGENTS.md
├── entries/
├── media/
├── scripts/
│   ├── build/
│   │   └── main.go
│   └── validate/
│       └── main.go
└── dist/
    └── vocabulary.tsv
```

This outline shows the core vocabulary pipeline. Shared tooling and tests live in `internal/vocabulary/`; the audio command lives in `scripts/audio/`. See the [README project layout](README.md#project-layout) for the current command directories. Add directories only when needed.

---

## Vocabulary Schema

The vocabulary schema contains these fields:

```text
id
language
word
pronunciation
part_of_speech
meaning
context
example
notes
audio
learning_priority
prerequisites
tags
```

Example:

```yaml
id: idiomatic-tech
language: en
word: idiomatic
pronunciation: /ˌɪdiəˈmætɪk/
part_of_speech: adjective
meaning: 地道的；惯用的
context: 符合某种语言、框架或技术生态惯用方式的。
example: Prefer idiomatic Go code.
notes: Common usage includes "idiomatic Go".
audio: idiomatic-tech.mp3
tags:
  - tech
  - programming
```

---

## Required Vocabulary Fields

Every entry must contain these nonempty string fields:

```text
id
language
word
meaning
context
example
```

Optional fields are:

```text
pronunciation
part_of_speech
notes
audio
learning_priority
prerequisites
tags
```

Optional fields are strings except for `tags` and `prerequisites` (lists of strings) and `learning_priority` (an integer from 1 to 3, defaulting to 2). Omit unused optional fields; validation rejects empty values.

---

## Learning Order

Use `learning_priority` to propose introduction order: 1 for foundations and early practical use, 2 for standard or applied vocabulary, and 3 for specialized material. Omit it to use 2 when uncertain. Judge usefulness, meaning, and complexity rather than word length; greetings can be foundational and short technical terms advanced.

Add `prerequisites` only for existing stable IDs that help explain the target. Prefer direct dependencies in the same language. Do not infer them from every word in an example, confuse shared spelling with shared meaning, or create extra vocabulary just to complete a graph. Missing references, repeated IDs, self-references, and cycles fail validation.

Builds place prerequisites first, then choose the lowest priority number among eligible entries, using stable ID for ties. `StudyOrder` is a derived export rank, never a source field or note identity. `Load` retains ID ordering for audio batches. Adding or removing entries can change ranks; always rebuild the full export.

These fields describe content, not personal mastery. Anki owns learning state, review scheduling, and history. Use the [README workflow](README.md#apply-the-order-in-anki) to apply ranks to new cards; do not modify a live Anki collection without explicit authorization.

## Stable IDs

Every entry must have a globally unique, stable `id`. It is the entry’s permanent identity.

Examples:

```text
idiomatic-tech
commit-git
commit-general
address-network
address-general
```

Stable IDs support:

- Anki note identity
- filenames
- generated audio
- update detection
- duplicate detection
- exported records

Do not change an existing ID casually.

Changing an ID may cause an existing Anki note to be treated as a new entry.

---

## Filenames

Use lowercase kebab-case.

Examples:

```text
entries/idiomatic.yml
entries/idempotent.yml
entries/race-condition.yml
entries/backward-compatible.yml
entries/de-facto-standard.yml
```

For a word with multiple meanings, use a descriptive filename for each entry.

Example:

```text
entries/commit-git.yml
entries/commit-general.yml
```

Avoid unnecessary abbreviations in filenames.

---

## Vocabulary Content Rules

When creating or updating vocabulary, prioritize:

- common modern usage in the entry’s language
- real-world usage
- technical usage when relevant
- concise meanings
- natural context
- natural example sentences
- useful collocations
- distinctions from confusing terms

Avoid:

- obscure meanings unless relevant
- dictionary dumps
- overly academic definitions
- artificial examples
- excessive synonyms
- redundant explanations
- beginner-level teaching

Choose one useful learning target per entry. Prefer dictionary headwords where appropriate, while retaining meaningful fixed expressions, grammatical forms, and natural plural forms. Search for existing component vocabulary before adding a phrase from a supplied sentence.

Use ordinary sentences and freely assembled descriptions as examples or notes. Create a phrase entry when the phrase itself has independent learning value, such as a fixed expression, technical term, named dish, or reusable grammar pattern. Do not generate every combination of already-covered words or split expressions solely by word count.

When consolidating a supplied list, record how source items map to retained entries or preserved examples. Apply this selection policy to the requested scope; it does not authorize unrelated entry deletion, ID changes, or audio regeneration.

---

## Technical Vocabulary

When a word has a specific software-engineering meaning, explain that meaning clearly.

Example:

```text
word: artifact
```

Good context:

```text
A file or output produced by a build, compilation, CI, or deployment process.
```

Good example:

```text
Upload the build artifacts after the test job completes.
```

Do not prioritize an unrelated dictionary meaning when the entry is clearly technical.

---

## Meaning Language

Use concise Chinese in `meaning` and explain practical usage in `context`. Write `example` in the entry’s language.

Example:

```yaml
meaning: 幂等的
context: 重复执行相同操作不会继续改变最终状态。
example: HTTP PUT is expected to be idempotent.
```

---

## Example Sentence Rules

Examples should be:

- natural
- concise
- grammatically correct
- realistic
- relevant to the target meaning

When writing an example, reuse relevant existing vocabulary in the same language when it fits naturally. Otherwise, use a more suitable example freely. Accuracy, naturalness, and practical context take priority over reuse; do not force vocabulary overlap.

For technical vocabulary, prefer examples from:

- software engineering
- APIs
- Git
- DevOps
- frontend
- backend
- infrastructure
- databases
- networking

when appropriate.

Avoid awkward textbook-style sentences.

---

## Pronunciation

Use a consistent pronunciation format: verified IPA for English, kana readings for Japanese, and the appropriate conventions for other languages.

Example:

```text
/ˌɪdiəˈmætɪk/
```

Do not invent pronunciation.

If pronunciation cannot be determined reliably, leave it unset and report the limitation.

---

## Tags

Use lowercase tags.

Examples:

```text
tech
programming
git
go
javascript
typescript
frontend
backend
devops
networking
database
finance
general
noun
verb
adjective
phrase
```

Do not create many near-duplicate tags.

Prefer:

```text
backend
```

over:

```text
back-end
backend-dev
backend-development
```

unless there is a clear reason.

---

## Adding New Vocabulary

When asked to add a word:

1. Search existing entries for the same word, language, and meaning.
2. Reuse a matching entry. For a new meaning, choose a stable ID and create one YAML file.
3. Add a concise meaning, practical context, and a natural example.
4. Add relevant tags and pronunciation if reliable. Choose a learning priority and selective prerequisites using the rules above.
5. If audio generation is enabled, generate the recording before adding its metadata.
6. Run validation and build the export when appropriate.
7. Report the files changed and the verification results.

Do not create duplicate entries for the same meaning.

---

## Updating Existing Vocabulary

When updating an entry:

- preserve its stable ID
- preserve unrelated fields
- avoid rewriting content unnecessarily
- keep diffs small
- maintain existing style
- validate after modification

If the requested update changes the meaning substantially, consider whether it should become a separate entry instead.

---

## Duplicate Handling

Duplicates may exist at different semantic levels.

These may be valid:

```text
commit-git
commit-general
```

These are likely invalid:

```text
idiomatic-tech
idiomatic-programming
```

if both represent the same meaning.

Duplicate detection should consider:

- stable ID
- word
- meaning
- context

Do not automatically merge semantically different entries.

---

## Validation

Run vocabulary validation from the repository root:

```bash
go run scripts/validate/main.go
```

Validation checks for:

- malformed YAML
- missing required fields
- duplicate IDs
- invalid IDs
- empty required fields
- invalid filenames
- invalid tag format
- missing referenced audio files
- invalid language tags and unsupported fields
- invalid learning priorities, prerequisite references, and dependency cycles

Errors should identify:

- the affected file
- the affected field
- the reason
- how to fix the issue when practical

Example:

```text
entries/idiomatic.yml: missing required field "context"
```

Prefer deterministic validation output.

---

## Build

Build the Anki export from the repository root:

```bash
go run scripts/build/main.go
```

Expected output:

```text
dist/vocabulary.tsv
```

The build must continue to:

- read all valid source entries
- produce deterministic ordering
- generate deterministic output
- escape fields correctly
- generate Anki-compatible fields
- include tags
- include audio references when available

Running the build twice without source changes should produce identical output.

---

## Anki Export

The export contains these columns in order:

```text
ID
Language
Word
Pronunciation
PartOfSpeech
Meaning
Context
Example
Notes
Audio
TTSLanguage
StudyOrder
Tags
```

Keep the stable ID in the first column. The export has twelve note fields followed by Anki tags in column 13. `TTSLanguage` and `StudyOrder` are derived during export, not stored in YAML. `StudyOrder` is the browser sort field; keep `ID` first for import matching.

Example audio field:

```text
[sound:idiomatic-tech.mp3]
```

Do not mix Anki review state into the LexForge source data.

---

## Generated Files

Keep generated artifacts under:

```text
dist/
```

Do not edit files under `dist/` manually.

If a generated file is incorrect, fix:

- source data
- validation logic
- build logic

then regenerate it.

---

## Audio

Use deterministic pronunciation-audio filenames based on the stable ID:

```text
media/<id>.mp3
```

Example:

```text
media/idiomatic-tech.mp3
```

Rules:

- use the stable entry ID
- do not regenerate existing audio unless requested
- prefer neutral, clear pronunciation in the entry’s language
- keep the audio generation implementation replaceable
- do not bind the vocabulary schema to one TTS provider

If a future audio provider requires credentials, read them from environment variables rather than hardcoding them.

---

## Secrets

Never commit:

- API keys
- tokens
- credentials
- passwords
- private configuration

Use environment variables.

Example:

```text
OPENAI_API_KEY
```

If example configuration is needed, provide something like:

```text
.env.example
```

with placeholders only.

---

## Go Guidelines

Prefer Go for project tooling.

Examples:

```text
scripts/build/main.go
scripts/validate/main.go
```

Use idiomatic Go.

Prefer:

- standard library where practical
- small focused packages only when justified
- clear error handling
- deterministic behavior
- simple CLI interfaces
- readable code
- explicit data structures

Avoid:

- unnecessary reflection
- premature interfaces
- unnecessary concurrency
- large dependency trees
- framework-like abstractions

---

## Go File Organization

For small tools, prefer a single complete file when practical.

Example:

```text
scripts/build/main.go
```

Do not split a small utility into multiple packages unless it materially improves maintainability.

If complexity grows enough to justify packages, refactor deliberately.

---

## Dependency Management

Keep dependencies minimal.

Before adding a dependency:

1. check whether the Go standard library is sufficient
2. evaluate maintenance status
3. evaluate dependency size
4. confirm that the dependency solves a real problem

A YAML parser may reasonably require an external dependency.

Do not introduce dependencies only to reduce a few lines of code.

---

## CLI Behavior

CLI tools should:

- return non-zero exit codes on failure
- write useful errors to stderr
- avoid interactive prompts unless required
- behave deterministically
- be automation-friendly

Example:

```bash
go run scripts/validate/main.go
```

should either:

```text
Validation successful: 42 entries
```

or return clear validation errors.

---

## Git Workflow

LexForge is expected to receive frequent small updates.

Optimize for:

- small diffs
- independent files
- easy merges
- meaningful commits

Adding one word should normally modify one entry file.

Avoid broad formatting changes while adding unrelated vocabulary.

---

## Commit-Friendly Changes

Do not:

- reorder unrelated YAML files
- rewrite all entries
- reformat the entire repository
- regenerate unrelated files unnecessarily

unless the task explicitly requires it.

Keep changes scoped.

---

## GitHub

No GitHub Actions workflow is currently configured. If CI is needed, useful checks include:

- validation
- build verification
- generated artifact checks

Do not introduce CI until it provides clear value.

When CI is added, keep it simple.

---

## Package Management Preferences

For local Node.js dependency operations, if Node tooling is ever introduced:

```text
pnpm install
pnpm add
pnpm update
pnpm remove
```

Use npm for local development and lifecycle scripts:

```text
npm run <script>
npm pack
```

For GitHub Actions:

```text
npm install
npm run <script>
```

Do not use:

```text
npm ci
```

Do not add a `packageManager` field unless explicitly requested.

Do not introduce Node.js tooling unless it is actually needed.

---

## Formatting

Use consistent formatting.

For YAML:

- 2-space indentation
- no tabs
- lowercase field names where defined by schema
- stable field ordering when practical

Recommended order:

```text
id
language
word
pronunciation
part_of_speech
meaning
context
example
notes
audio
learning_priority
prerequisites
tags
```

---

## Documentation

Keep documentation concise and practical.

Prefer examples over long explanations.

Update documentation when:

- schema changes
- commands change
- directory structure changes
- workflow changes
- Anki export behavior changes

Do not let documentation drift from implementation.

---

## README Responsibilities

`README.md` should explain:

- what LexForge is
- how to add vocabulary
- how to validate
- how to build
- how to import into Anki
- how audio works

`AGENTS.md` defines how coding agents should work.

Do not duplicate large sections unnecessarily across both files.

---

## Scope Control

Do not turn LexForge into:

- a general dictionary
- a language-learning SaaS
- a flashcard application
- a note-taking application
- a web browser extension
- a replacement for Anki
- a full learning-management system

Prefer integration over reimplementation.

---

## Non-Goals

Unless explicitly requested, do not implement:

- user accounts
- authentication
- cloud databases
- web dashboards
- mobile applications
- spaced-repetition algorithms
- proprietary sync protocols
- gamification
- chat interfaces
- complex AI pipelines

---

## Performance

The vocabulary collection is expected to be relatively small compared with normal software datasets.

Prioritize simplicity and correctness over micro-optimization.

Still avoid obvious inefficiencies such as:

- parsing the same file repeatedly
- unnecessary network calls
- regenerating audio unnecessarily
- non-deterministic iteration
- unnecessary full-repository rewrites

---

## Determinism

Deterministic behavior is important.

Given identical source files, build output should remain identical.

Sort entries explicitly.

Do not depend on:

- map iteration order
- filesystem enumeration order
- timestamps
- random values

unless explicitly required.

---

## Error Handling

Do not silently ignore invalid source data.

Fail clearly when data cannot safely be processed.

Good:

```text
entries/upstream.yml: invalid field "id": must use lowercase kebab-case
```

Bad:

```text
Failed.
```

---

## Backward Compatibility

Once entries exist, avoid breaking schema changes.

If the schema must evolve:

1. Explain the reason.
2. Preserve existing data where possible.
3. Provide migration logic if necessary.
4. Update validation, build logic, and documentation.

Do not silently reinterpret existing fields.

---

## Testing

Add tests when logic becomes non-trivial.

Prioritize tests for:

- validation
- ID handling
- duplicate detection
- deterministic sorting
- TSV escaping
- Anki audio formatting

Do not create excessive test infrastructure for trivial code.

---

## Before Completing a Task

Before considering a task complete:

1. Inspect existing conventions and keep the change scoped.
2. Format changed files.
3. Run validation and relevant tests.
4. Run the build if affected.
5. Inspect the Git diff and confirm that no secrets were added.
6. Report important changes, verification results, and limitations.

---

## Expected Commands

Run these commands from the repository root:

```bash
go run scripts/validate/main.go
```

and:

```bash
go run scripts/build/main.go
```

Prefer commands that can be copied and run directly.

---

## Codex Response Style

When reporting completed work:

- be concise
- describe what changed
- list important files
- mention validation/test results
- mention unresolved limitations
- avoid explaining basic programming concepts

Example:

```text
Implemented the initial vocabulary pipeline.

Changes:
- Added YAML entry schema.
- Added validation for required fields and duplicate IDs.
- Added deterministic TSV generation.
- Added sample vocabulary entries.

Validation:
- `go run scripts/validate/main.go` passed.
- `go run scripts/build/main.go` generated `dist/vocabulary.tsv`.
```

---

## Decision Rule

When multiple implementations are possible, prefer the solution that is:

1. simpler
2. easier to maintain
3. more Git-friendly
4. more deterministic
5. more idiomatic
6. less dependent on external services
7. easier to replace later

---

## Final Principle

LexForge should feel like a small, high-quality developer tool.

The ideal workflow remains:

```text
Find word
→ add entry
→ validate
→ build
→ commit
→ sync to Anki
→ review on AnkiDroid
```

Every implementation decision should make that workflow simpler, safer, or more maintainable.

## Multilingual Entries

Require a standard BCP 47 `language` tag with no project-specific allowlist. Use `ko` for Korean and `ja` for Japanese. Validation uses a standards-aware parser and canonicalizes language tags internally and in exports without rewriting source spelling. Duplicate-word warnings apply within each canonical language tag; stable IDs remain globally unique.

Write meanings and practical context in Chinese, and examples in the vocabulary language. Japanese cards show the original expression and kana reading on both sides, with optional audio once per side. Render each side explicitly and keep meanings on the back. Keep language metadata separate from lowercase vocabulary tags; the exporter derives `lang-<canonical-tag>` automatically.

Use the documented local macOS audio command with language-matching voices. Never substitute an English voice for another language. Voice and rate overrides belong to CLI options, not vocabulary data. A valid language does not require an installed voice for validation or export.
