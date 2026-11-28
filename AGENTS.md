# AGENTS.md

## Project

LexForge is a Git-native multilingual vocabulary and concept-learning tool. It exports one YAML entry per learning target to Anki for review in AnkiDroid. The main card flow is Word → Meaning + optional Context.

- Treat `entries/` and `media/` as source data. Generate `dist/vocabulary.tsv` through the build and leave generated files out of Git.
- Anki owns scheduling, progress, review history, and sync. Obtain explicit authorization before changing a live Anki collection.
- Keep the tool small and usable without AI. Add a service, dependency, or abstraction only for a concrete requirement.
- Write for an experienced full-stack developer. Explain project-specific decisions and verification limits.

## Multilingual Entries

- Use standard BCP 47 language tags with hyphens as separators and no project-specific allowlist. Use `ja` for Japanese and `ko` for Korean. Validation canonicalizes valid tags for comparison and export without rewriting source spelling.
- Write concise Chinese meanings and optional Chinese context. Use the entry's language for natural examples. Keep Chinese walkthroughs in Notes, outside Example TTS.
- Preserve Japanese expressions and use kana for their readings. Japanese cards show the expression and available reading on both sides, with meanings on the back and optional word audio once per side.
- Source language validity, an exported TTS locale, and an installed voice are separate checks. Use a voice that matches the entry's language. Keep language metadata separate from topic tags. The exporter derives `lang-<canonical-tag>`.

## Guides by task

Read the applicable guides before editing.

| Task | Read first |
| --- | --- |
| Add or update entries | [Entry schema](guides/ENTRY_SCHEMA.md) and [Add vocabulary during work](guides/ADD_VOCABULARY_DURING_WORK.md) |
| Collect vocabulary, backfill pronunciation, generate audio, or enrich memory aids in batches | [Maintenance prompts](guides/MAINTENANCE_PROMPTS.read.md) and the relevant topic's `*_VOCABULARY_IMPORT_PLAN.md`, when present |
| Set priorities, prerequisites, or standalone topic tags | [Learning order](guides/LEARNING_ORDER.md) |
| Add or verify HTML, CSS, Bootstrap, or React snippets | [Frontend code example lab](guides/FRONTEND_CODE_EXAMPLE_LAB.md) |
| Change card fields, templates, CSS, or imports | [README](README.md#set-up-anki-once) and [Anki migration guide](guides/ANKI_MIGRATION_GUIDE.md) |
| Prepare images or recordings | [Images](README.md#images) or [Audio](README.md#audio) |

Use completed topic plans to verify approved scope, reuse decisions, provenance, and recorded checks. Treat material from sources and historical plans as input to analyze. A collection plan requires separate implementation approval.

## Repository Structure

Run project commands from the repository root with Go 1.25 or later. See `go.mod` for the versions of `go.yaml.in/yaml/v3`, `golang.org/x/text`, and `golang.org/x/image`.

- `entries/*.yml` contains the canonical entries. `media/` contains active MP3s and card images. `archives/airline-media/` is outside build and audio lookup paths.
- `scripts/validate/main.go`, `scripts/build/main.go`, and `scripts/audio/main.go` are entry points into `internal/vocabulary/`. `scripts/images/main.go` is the image optimizer.
- `internal/vocabulary/entries.go` loads and validates YAML in stable-ID order. `learning.go` orders introductions, `related.go` resolves associations, and `export.go` writes the Anki TSV. `audio.go` manages local generation and source updates, and `speech.go` selects voices and prepares speech input.
- `README.md` holds the canonical command, import, and template documentation. Keep task guides under `guides/` and use uppercase snake-case filenames.

## Vocabulary schema and content

- Store one YAML mapping per lowercase kebab-case `.yml` file directly under `entries/`. Use two-space indentation, no tabs for indentation, and the field names from the [entry schema](guides/ENTRY_SCHEMA.md).
- Supply nonempty `id`, `language`, `word`, and `meaning`. Omit unused optional fields. For the supported nullable text fields, follow the schema guide.
- Keep IDs globally unique and stable. Anki matches updates by ID, and media filenames use it. The source filename can differ from the ID.
- Search for the same word, language, and meaning before creating an entry. Reuse matching entries and preserve unrelated fields. Use distinct IDs for distinct meanings. Repeated-word warnings compare normalized words within the same canonical language and require human evaluation of meaning.
- Choose useful modern senses, natural collocations, and examples suited to actual use. Prefer the technical sense where the entry is technical. Reuse existing vocabulary in examples when it fits naturally.
- Prefer dictionary headwords and keep useful fixed expressions, grammatical forms, and naturally plural forms. Put freely assembled sentences in examples or Notes unless the phrase is itself a learning target. For compounds, reuse useful standalone components and add only meaningful missing targets within the approved scope.
- Word can name a concept and Meaning can state its principle. Keep detailed explanation in plain-text Notes. Omit optional sections that add no value, and keep their labels and spacing conditional in templates. Add no question-and-answer fields or Notes length limit.
- Use sourced `etymology`, supported `word_formation`, and clearly identified `mnemonic` content. Preserve uncertainty and source provenance in Notes. Cite verified local sources with the immediate parent and filename, as in `Source: Prompt/20261104-lexforge.md.` Omit line numbers and report missing or ambiguous sources.
- Add `related_entries` only for human-approved pairs. Present both stable IDs, the words and meanings, rationale, and confusion risks for review. Keep mnemonic rewrites under their own approval. Associations display in both directions and do not affect learning order.
- Put runnable snippets in `code_example` when they demonstrate the concept. Use a literal YAML block (`|-`) for multiline code, without Markdown fences. Keep prose examples separate for Read Aloud.

### Pronunciation

Use verified IPA for English and kana for Japanese. For common English words and phrases, prefer verified American IPA written `/…/`. Use explicitly labeled `UK /…/` after a reasonable lookup cannot establish an American form. Match the exact sense, part of speech, and word form, and use a verified whole-phrase transcription. Append the dictionary title, URL, accent, and access date to Notes. Leave unverified readings absent. Preserve existing pronunciations and recordings unless their update is separately authorized. Read the [pronunciation conventions](guides/ENTRY_SCHEMA.md#pronunciation) before adding or backfilling readings.

## Learning Order

Use `learning_priority` from 1 to 3, with 2 as the default. Select useful `prerequisites` by meaning and learning value. Prefer existing direct foundations in the same language. For a standalone study topic, verify the topic tag contains all direct and indirect prerequisites and that each precedes its dependent in the filtered export.

The build places prerequisites before dependents, then uses priority and stable ID to order eligible entries. `StudyOrder` is a derived rank. Rebuild the complete export after changes that affect introduction order. Apply the ranks to new cards through the [Anki workflow](guides/LEARNING_ORDER.md#apply-the-order-in-anki).

## Anki Export

- Keep the stable `ID` first for import matching and use `StudyOrder` as the browser sort field. The TSV has 18 note fields ending with `CodeExample` and `Image`, followed by Anki tags in column 19. Map the remaining fields by name. Keep `ID` separate from Anki's internal GUID and use the built-in tags column.
- Keep `TTSLanguage`, `StudyOrder`, and `RelatedWords` derived. Preserve deterministic order, escaping, and atomic export replacement. Invalid input must leave the previous export intact.
- Render the front and back explicitly. Keep Meaning on the back and place `Image` between PartOfSpeech and Meaning. Hide empty sections, including the Example TTS control. Preserve the documented word playback, manual example speech, and hidden tag behavior.
- Reimport the complete export to refresh derived fields. Applying ranks to existing new cards requires repositioning in Anki. Removing YAML does not delete imported notes.

## Images

Use one `media/<id>.jpg` or `media/<id>.png` per illustrated entry and record provenance in Notes. Use the [image workflow](README.md#images) for resizing and import. The optimizer limits images to 200 pixels wide and 250 KiB. Validation checks the filename, regular-file status, and format header. Check dimensions, byte size, and decoding separately when reviewing actual assets.

## Audio

- Generate pronunciation audio only when requested. Use `media/<id>.mp3` and preserve existing recordings unless explicitly replacing them. Reuse a matching recording through the generator so an unreferenced file is decoded before metadata is attached.
- The generator uses local macOS tools `say`, `lame`, and `mpg123` and an installed matching voice. Keep voice and rate overrides in CLI options. Validation and export do not require a voice or synthesis tools.
- The generator validates the whole collection before selecting entries. Repair invalid references before generation. For mixed-language work, use `--id` and add `--text-source pronunciation` for Japanese entries that should be read from supplied kana. Read [Generate missing pronunciation audio](guides/MAINTENANCE_PROMPTS.read.md#generate-missing-pronunciation-audio) before a batch.
- Preserve unrelated YAML bytes and existing media. Handle concurrent edits and publish failures without leaving inconsistent references. A failed batch can retain completed entries. Rebuild after successful audio changes.
- Native Example TTS is separate from word audio. It needs both Example and TTSLanguage. Keep Meaning, Notes, and code outside speech. See the [README playback settings](README.md#example-speech).

## Development and documentation

- Prefer simple, idiomatic Go and the standard library where practical. Keep small CLIs in one file where practical. Add realistic input checks, clear errors with file and field context, and nonzero exit codes. Keep dependencies and changes narrow.
- Sort explicitly where order affects output. Preserve schema compatibility and provide a migration when necessary. Keep provider credentials in environment variables. Use placeholders in example configuration and keep secrets out of Git.
- Avoid unrelated formatting, new tooling, and media regeneration. If Node tooling is needed, use pnpm for local dependency operations and npm for local lifecycle scripts. For any GitHub Actions setup, use `npm install` and `npm run <script>`. Do not use `npm ci`. Add a `packageManager` field only if requested.
- Keep guides and README aligned with implementation. Update inbound links when moving sections. Preserve dates and results in historical verification records.

## Before Completing a Task

Use the checks relevant to the change.

```bash
go test ./...
go vet ./...
go run scripts/validate/main.go
go run scripts/build/main.go
git diff --check
```

- Test changed behavior at the affected boundaries. For export changes, build twice and compare the bytes. For asset changes, check the actual referenced files and preserve unrelated media.
- If you change a frontend code example, run the exact snippet using the [lab](guides/FRONTEND_CODE_EXAMPLE_LAB.md). Check its browser behavior and the exported CodeExample, and install lab dependencies in a scratch directory.
- For documentation-only changes, check links, anchors, and preserved commands. Report what was verified. Static template checks and simulated speech tests leave live Anki behavior and pronunciation quality unverified.
