# Add vocabulary during work

Capture useful words, expressions, or concepts as you encounter them. Run commands from the repository root.

## Workflow

1. Search `entries/` for the word in the same language and check whether its meaning is already covered.
2. Create one lowercase kebab-case `.yml` file for a new meaning, or update the existing entry without changing its ID.
3. [Validate](#validate-and-export), review the diff, and commit the source file.
4. Build and import a batch when convenient.

For AI-assisted entry creation, use the reusable prompt in [Add vocabulary with Codex](ADD_VOCABULARY_WITH_CODEX.md).

## Choose a learning target

Choose one useful learning target per entry. Prefer dictionary headwords where appropriate, while keeping meaningful fixed expressions, grammatical forms, and natural plural forms. Put ordinary sentences and freely assembled descriptions in `example` or `notes`; create a phrase entry when the phrase itself is the learning target. For example, reuse `beef`, `rice`, and `sauce` for a meal description, but keep a useful expression such as `mashed potatoes` or the pattern `Would you like A or B?`. Do not generate every combination of existing words or split phrases solely by length.

Use separate IDs and descriptive filenames for distinct meanings, such as `commit-git` and `commit-general`. Changing an ID can create another Anki note. Renaming a source file while preserving its ID does not change its identity.

## Write the entry

For example, `entries/idiomatic.yml` can contain:

```yaml
id: idiomatic-tech
language: en
word: idiomatic
pronunciation: /ˌɪdiəˈmætɪk/
part_of_speech: adjective
meaning: 地道的；惯用的
context: 符合某种语言、框架或技术生态惯用方式的。
example: Prefer idiomatic Go code.
tags:
  - tech
  - programming
```

Use the [Entry schema guide](ENTRY_SCHEMA.md) for required and optional fields. Add learning priorities and useful prerequisites using the [Learning order guide](LEARNING_ORDER.md).

Pronunciation audio is optional. Follow the [audio workflow](../README.md#audio) to generate or reuse a recording before adding its audio reference.

## Validate and export

Validate source changes before committing:

```bash
go run scripts/validate/main.go
git diff --check
```

Review changed and new files before committing. When ready to import a batch, build the export:

```bash
go run scripts/build/main.go
```

The build validates the collection and writes `dist/vocabulary.tsv`. Never edit this generated file manually. Follow [Import and update](../README.md#import-and-update) to update Anki, then [apply the learning order](LEARNING_ORDER.md#apply-the-order-in-anki) to new cards when needed. Building and importing can be batched; they are not required after every word you capture.
