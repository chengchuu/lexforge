# Multilingual Vocabulary Plan

Implementation requires explicit approval. Creating this plan does not authorize schema, vocabulary, template, or audio changes.

## Goal

Support independent English, Japanese, Korean, and future-language vocabulary in one Git-native workflow:

```text
Capture → YAML → validate → build → Anki Desktop → AnkiWeb → AnkiDroid
```

Use one shared schema, deterministic TSV export, and Anki note type. Keep learning state in Anki. The project is new, and recreating the Anki field setup is acceptable.

## Current Constraints

- Source entries already accept Unicode text, but the exporter and documented templates call the target word `English`.
- The schema has no language field. Duplicate-word warnings currently compare words across the entire collection.
- The local audio generator reads `word` using Samantha at 150 words per minute, then encodes with LAME and verifies decoded samples with `mpg123`.
- IDs, filenames, and vocabulary tags use lowercase kebab-case. Audio filenames use stable IDs.
- Source data and media are canonical; `dist/` is generated and must not be edited manually.

## Shared Schema and Language Validation

Add a required `language` string. Preserve the existing fields and add `Language` to the Go entry model.

- Use standard BCP 47 tags without a project-specific allowlist: `en`, `ja`, `ko`, `fr`, and more specific tags such as `en-US` or `zh-Hant`.
- Korean uses `ko`; `kr` identifies a different language.
- Use a standards-aware parser. Evaluate `golang.org/x/text/language`, verify its documented parsing and canonicalization behavior, and select a release compatible with Go 1.25.
- Reject parser errors, including unknown subtags reported by the parser. Do not accept a partially recovered tag after an error.
- Use canonical tags internally and in exports without automatically rewriting source spelling. Verify how the parser treats deprecated aliases, case differences, private-use tags, and extensions before documenting its acceptance policy.
- Do not infer language from the word, filename, ID, or script.
- Keep vocabulary-tag validation separate from language-tag validation.
- Scope repeated-word warnings by canonical language tag and normalized word. Keep IDs globally unique across languages.
- Add `language: en` to existing English entries while preserving their IDs, content, and audio files.

Accepting a language tag does not imply that an audio voice is installed or that specialized card behavior exists for that language.

## Anki Export and Presentation

Keep the note type named `LexForge`. Export these note fields in order, followed by Anki tags as column 11:

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
```

- Keep `ID` first and preserve first-field update matching. Do not use Anki's internal GUID as the vocabulary ID.
- Rename the exported `English` field to `Word`; update headers, field mapping, templates, and documentation together.
- Retain `dist/vocabulary.tsv` as the single export. Preserve deterministic ordering, escaping, audio references, and safe replacement after validation.
- Derive one lowercase language tag for filtering, using `lang-` plus the lowercase canonical language tag, such as `lang-en` or `lang-en-us`. Merge it with source tags, deduplicate, and sort without modifying source tags.
- Show `Word`, optional pronunciation, and optional audio on the front. Render each side explicitly rather than using `FrontSide`, so these fields appear once per side. The back adds part of speech, meaning, context, example, and notes.
- Preserve optional-field conditionals, mobile typography, light/dark styling, and top alignment. Rename English-specific CSS classes and selectors consistently.

For Japanese, show the original Japanese expression and its kana reading on both sides. For example, `確認` and `かくにん` appear on both sides, with the Chinese meaning on the back. Use Japanese example sentences. Do not use romaji as the Japanese pronunciation convention.

Use language-appropriate pronunciation conventions for other languages. Keep meanings in concise Chinese and explanations practical. Furigana, pitch-accent notation, reverse cards, and additional pronunciation fields are outside this implementation.

Document recreating the Anki field setup and importing the new export. Explain that deleting and recreating notes discards their existing review history; do not automatically delete or migrate the user's Anki collection.

## Local macOS Audio

Extend the existing generator rather than adding another TTS provider or generation framework.

| Language | Initial default voice | Voice locale |
|---|---|---|
| English | Samantha | `en_US` |
| Japanese | Kyoko | `ja_JP` |
| Korean | Yuna | `ko_KR` |

These voices were listed on the development Mac during analysis. Check availability at execution time on every machine.

- Keep the existing `--id`, `--all`, and `--overwrite` workflow.
- Prefer the documented defaults for bare `en`, `ja`, and `ko`. For other tags, discover matching installed voices; require an explicit selection when the match is ambiguous or a different region would be used. Do not silently substitute an English voice.
- Provide an explicit voice override for a single-entry run. Keep provider-specific voice names out of vocabulary source files.
- Speak `word` by default. Provide an explicit single-entry option to speak verified Japanese kana from `pronunciation` when the written expression has an ambiguous reading. Reject this option for non-Japanese entries or missing readings rather than speaking English IPA.
- Keep English's existing rate. Audition Japanese and Korean before selecting their documented defaults; do not assume the English rate transfers appropriately. Provide a single-entry rate override.
- Define exact CLI option names and voice-selection tests before implementing these additions; the override behavior above is the intended interface.
- Preserve `media/<id>.mp3`, encoding and decode checks, reuse behavior, explicit overwrite, source-edit detection, cleanup, and stable-ID batch ordering.
- Check prerequisites for the selected work before generating new files. Preserve the existing documented stop-on-failure behavior for runtime batch errors.
- Keep unsupported voice availability separate from entry validity: valid vocabulary must still validate and build without generation support.

Kana can resolve a reading but does not guarantee correct pitch accent. Listen to generated Japanese audio before importing it. File decoding verifies readable samples, not linguistic accuracy. Audio bytes may vary across macOS or encoder versions.

## Implementation Sequence

1. Verify and pin the language parser; settle its acceptance policy and the audio CLI options.
2. Extend schema parsing, validation, language normalization, duplicate warnings, and source-update safeguards.
3. Add language metadata to existing entries without changing IDs or regenerating media.
4. Update TSV headers, field positions, derived tags, and the Anki templates and setup instructions.
5. Extend local audio voice selection and explicit pronunciation/rate controls.
6. Update README and AGENTS.md language conventions together. Retain focused Git-native tooling and avoid unrelated refactoring.
7. Run regression checks and report actual renderer, device, and audio coverage separately.

## Verification and Acceptance

- Test missing and malformed language values, unknown subtags, equivalent casing, canonical aliases, and valid regional/script tags according to the selected parser's documented policy.
- Test Japanese and Korean Unicode content, kana readings, same-word entries across languages, same-language duplicate warnings, and global duplicate-ID rejection.
- Verify the 11-column export, note-field positions, derived-tag deduplication, stable ordering, escaping, and identical output across repeated builds.
- Verify that invalid input leaves an existing export untouched and that audio metadata updates preserve language and unrelated source content.
- Test installed/missing voices, locale matching, ambiguous choices, explicit overrides, Japanese reading input, rate selection, reuse, overwrite, decode failures, and cleanup paths.
- In an isolated Anki collection, import twice and confirm no duplicate notes; update content and confirm stable identity and scheduling. Confirm optional fields, repeated Japanese words on the back, and audio references.
- Check mobile widths, long content, light/dark mode, and actual AnkiDroid playback where a device is available. Clearly identify simulated or unavailable checks.
- Run `go test ./...`, `go vet ./...`, `go run scripts/validate/main.go`, `go run scripts/build/main.go`, and `git diff --check` after implementation.

Completion means another language can be added as source entries without changing the schema or exporter. Audio generation can require an installed voice and explicit selection without blocking validation or export.
