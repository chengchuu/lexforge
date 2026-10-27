# LexForge — Project Goal

## Overview

LexForge is a Git-native personal English vocabulary system designed for continuous learning during everyday software development and technical work.

The project treats vocabulary as structured source data managed in Git, while Anki and AnkiDroid are used for spaced-repetition review.

The primary learning direction is:

> English → Meaning + Context

LexForge should make it fast to capture useful words encountered during work, enrich them with practical context, generate Anki-compatible data, optionally generate pronunciation audio, and keep the vocabulary database easy to maintain through normal Git workflows.

---

## Primary Goals

LexForge should:

1. Provide a simple and maintainable personal vocabulary database.
2. Make adding new vocabulary fast enough to use during everyday work.
3. Store vocabulary in a Git-friendly format.
4. Keep vocabulary source data separate from generated artifacts.
5. Generate Anki-compatible import data.
6. Support repeated imports without creating duplicate notes.
7. Preserve Anki review history when vocabulary content is updated.
8. Support AnkiWeb and AnkiDroid as the review and synchronization layer.
9. Support pronunciation audio generation for vocabulary entries.
10. Make the workflow easy to automate with Codex.

---

## Target User

The primary user is an experienced full-stack developer.

Technical background includes:

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
- Git and GitHub

LexForge should therefore prioritize concise, practical, developer-oriented workflows instead of beginner-oriented abstractions.

Technical vocabulary should use real-world software engineering meanings and examples whenever appropriate.

---

## Core Learning Model

The primary card direction is:

```text
English
↓
Meaning
↓
Context
↓
Example
```

The goal is not only to memorize dictionary definitions.

A vocabulary entry should help answer:

- What does this word mean?
- What does it mean in the context where I encountered it?
- How is it naturally used?
- How is it used in software engineering or technical communication when relevant?

Example:

```text
English:
idiomatic

Meaning:
地道的；惯用的

Context:
符合某种语言、框架或技术生态惯用方式的。

Example:
Prefer idiomatic Go code.
```

---

## Source of Truth

Git is the source of truth.

Vocabulary should be stored as individual source files rather than one large manually maintained TSV file.

Recommended structure:

```text
lexforge/
├── README.md
├── GOAL.md
├── AGENTS.md
├── entries/
│   ├── artifact.yml
│   ├── deprecated.yml
│   ├── fallback.yml
│   ├── idiomatic.yml
│   └── upstream.yml
├── media/
├── scripts/
│   ├── build/
│   │   └── main.go
│   └── validate/
│       └── main.go
└── dist/
    └── vocabulary.tsv
```

Each vocabulary item should normally have its own file.

Example:

```yaml
id: idiomatic-tech
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

## Stable Identifier

Each vocabulary entry must have a stable `id`.

The ID should not depend entirely on the displayed English word because the same word may eventually have multiple meanings or contexts.

Examples:

```text
idiomatic-tech
commit-git
commit-general
address-network
address-general
```

The stable ID should be used for:

- Anki note identity
- generated audio filenames
- update detection
- duplicate prevention
- deterministic exports

---

## Vocabulary Fields

The initial schema should support:

```text
id
word
pronunciation
part_of_speech
meaning
context
example
notes
audio
tags
```

### Required Fields

At minimum:

```text
id
word
meaning
context
example
```

### Optional Fields

```text
pronunciation
part_of_speech
notes
audio
tags
```

The schema should remain simple and extensible.

Do not over-engineer the first version.

---

## Git Workflow

LexForge is intended to be updated frequently.

Typical workflow:

```text
Encounter a useful word during work
        ↓
Ask Codex to create or update the entry
        ↓
Create entries/<word>.yml
        ↓
Validate
        ↓
Build Anki export
        ↓
git add
        ↓
git commit
        ↓
git push
```

On another machine:

```bash
git pull
```

The repository should be designed to minimize merge conflicts.

This is one reason to prefer:

```text
one vocabulary item = one source file
```

instead of maintaining one large source TSV file.

---

## Build Workflow

Source vocabulary files should be converted into generated Anki-compatible data.

Example:

```bash
go run scripts/build/main.go
```

Expected output:

```text
dist/vocabulary.tsv
```

The generated file must not be manually edited.

Conceptually:

```text
entries/*.yml
    ↓
validate
    ↓
build
    ↓
dist/vocabulary.tsv
```

---

## Validation

Provide a validation command:

```bash
go run scripts/validate/main.go
```

Validation should detect at least:

- missing required fields
- duplicate IDs
- invalid IDs
- duplicate vocabulary entries where appropriate
- invalid tags
- missing referenced audio files
- malformed YAML
- empty values
- inconsistent source data

Validation errors should be clear and actionable.

---

## Anki Integration

LexForge should generate data that can be repeatedly imported into Anki.

The import design should allow existing notes to be updated instead of duplicated.

The stable `id` should be the primary identity field used for generated Anki data.

Recommended generated fields:

```text
ID
English
Pronunciation
PartOfSpeech
Meaning
Context
Example
Notes
Audio
Tags
```

Example:

```text
idiomatic-tech
idiomatic
/ˌɪdiəˈmætɪk/
adjective
地道的；惯用的
符合某种技术生态惯用方式的
Prefer idiomatic Go code.
Common usage includes "idiomatic Go".
[sound:idiomatic-tech.mp3]
tech programming
```

Anki is responsible for:

- note storage
- cards
- scheduling
- spaced repetition
- review history

LexForge is responsible for:

- vocabulary source data
- vocabulary enrichment
- export generation
- optional media generation

---

## AnkiDroid and AnkiWeb

The expected review flow is:

```text
LexForge
    ↓
Anki Desktop import
    ↓
AnkiWeb sync
    ↓
AnkiDroid
```

AnkiDroid is the primary mobile review client.

LexForge should not attempt to replace Anki's synchronization or spaced-repetition system.

---

## Pronunciation Audio

LexForge should optionally support pronunciation audio.

Recommended convention:

```text
media/<id>.mp3
```

Example:

```text
media/idiomatic-tech.mp3
```

Generated Anki audio field:

```text
[sound:idiomatic-tech.mp3]
```

Audio generation should be deterministic.

Rules:

- use the entry ID as the filename
- do not regenerate existing audio unless explicitly requested
- prefer clear, neutral English pronunciation
- keep the TTS implementation isolated from vocabulary data
- make it possible to replace the TTS provider later

The first implementation may support either:

- a remote TTS API
- a local macOS TTS workflow

The vocabulary schema must not depend on a specific TTS provider.

---

## Codex Workflow

Codex will be used as a primary maintenance assistant.

Typical request:

```text
Add this word to LexForge:

idempotent
```

Codex should:

1. inspect the existing repository conventions
2. determine whether the entry already exists
3. create or update the vocabulary source file
4. provide a concise and accurate meaning
5. add practical context
6. prefer technical context when relevant
7. provide a natural example sentence
8. add appropriate tags
9. optionally generate pronunciation audio
10. run validation
11. rebuild generated output when appropriate

Codex should avoid creating unnecessary files or abstractions.

---

## Vocabulary Entry Guidelines

Vocabulary entries should prioritize:

- common modern usage
- practical meanings
- meanings encountered in real work
- software engineering terminology when relevant
- concise explanations
- natural example sentences
- useful collocations
- distinctions from commonly confused terms

Avoid:

- obscure dictionary meanings unless relevant
- overly academic definitions
- unnecessarily long explanations
- beginner-level explanations
- artificial example sentences
- duplicate entries without a meaningful distinction

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

Prefer tags over creating many separate Anki decks.

The initial project should use one primary vocabulary collection and rely on tags for filtering.

---

## Naming Conventions

Project name:

```text
LexForge
```

Repository name:

```text
lexforge
```

Vocabulary filenames should normally use lowercase kebab-case:

```text
entries/idempotent.yml
entries/race-condition.yml
entries/backward-compatible.yml
entries/de-facto-standard.yml
```

Generated audio should use the stable entry ID:

```text
media/idempotent-tech.mp3
```

---

## Implementation Preferences

Prefer Go for project scripts.

Examples:

```text
scripts/build/main.go
scripts/validate/main.go
```

Implementation should favor:

- standard library where practical
- simple dependencies
- deterministic output
- clear CLI behavior
- useful error messages
- easy local execution
- cross-platform behavior where practical
- maintainable code over unnecessary abstraction

Avoid premature complexity.

---

## Non-Goals

The first version of LexForge is not intended to:

- replace Anki
- implement its own spaced-repetition algorithm
- implement its own cloud synchronization service
- become a full language-learning platform
- become a general note-taking application
- reproduce projects such as Onigiri or Synapse
- provide complex gamification
- provide account management
- provide a web application
- provide a mobile application
- automatically collect every unknown word

LexForge should remain focused on maintaining high-quality personal vocabulary data.

---

## Design Principles

### Git-Native

All important vocabulary source data should be easy to:

```text
diff
review
commit
push
pull
merge
```

### Source First

Source files are canonical.

Generated files are disposable and reproducible.

### Small Changes

Adding one word should normally create or modify only one vocabulary entry file.

### Deterministic

The same source data should produce the same generated output.

### Simple

Prefer a small, understandable system over a feature-heavy platform.

### Extensible

The schema and build system should allow future additions without requiring a rewrite.

### Anki-Compatible

Use Anki for what Anki already does well.

### AI-Assisted, Not AI-Dependent

Codex should make maintaining LexForge faster, but the repository should remain understandable and usable without AI.

---

## Initial Milestone

The first usable version should support:

1. `entries/*.yml`
2. a documented vocabulary schema
3. at least several example entries
4. YAML validation
5. duplicate ID detection
6. deterministic TSV generation
7. Anki-compatible audio references
8. tags
9. a simple README
10. commands that work locally

Expected commands:

```bash
go run scripts/validate/main.go
go run scripts/build/main.go
```

After building:

```text
dist/vocabulary.tsv
```

should be ready for import into Anki.

---

## Future Possibilities

Potential future improvements include:

- automatic pronunciation audio generation
- AnkiConnect integration
- incremental Anki updates
- automatic dictionary lookup
- automatic IPA lookup
- example sentence generation
- duplicate detection by semantic similarity
- vocabulary statistics
- GitHub Actions validation
- GitHub Actions artifact generation
- CLI commands such as:

```bash
lexforge add idiomatic
lexforge validate
lexforge build
lexforge audio
```

These should only be added when they provide clear practical value.

---

## Success Criteria

LexForge is successful when adding a useful word during everyday work feels lightweight.

The ideal workflow should be close to:

```text
Find word
→ add entry
→ validate
→ commit
→ sync to Anki
→ review on AnkiDroid
```

The project should make vocabulary accumulation feel like maintaining a small, well-structured software repository rather than maintaining a traditional word list.

---

## One-Sentence Goal

> LexForge is a Git-native personal English vocabulary system that turns words encountered during everyday technical work into structured, contextual, Anki-ready learning material.
