# Vocabulary Memory Aids Plan

## Status and authorization

Planned; not implemented. This document records the design discussed on 2026-10-30.

Creating this plan does not authorize schema, vocabulary, exporter, template, or audio changes. Implementation requires explicit approval. Approval to implement the feature does not automatically approve candidate vocabulary relationships or associated mnemonic text; those require content review as described below.

## Goal

Support optional etymology, word-formation explanations, mnemonics, and related vocabulary on card backs. Make Context optional when Meaning already explains the learning target adequately. Keep each card focused and omit empty sections completely.

The relationship mechanism applies to the entire multilingual collection. `closure-tech` and `disclosure-workflow` are illustrative candidates, not special cases or an approved relationship batch.

## Current implementation

- [Entry parsing](../internal/vocabulary/entries.go) requires nonempty `id`, `language`, `word`, `meaning`, `context`, and `example` strings. Unknown fields and empty optional values are rejected.
- [TSV export](../internal/vocabulary/export.go) produces twelve note fields and tags in column 13. `TTSLanguage` and `StudyOrder` are derived fields.
- The [README templates](../README.md#set-up-anki-once) render Context unconditionally and Notes conditionally. The source schema has no dedicated memory-aid or association fields.
- [Learning order](LEARNING_ORDER.md) uses directed prerequisites. Associations must remain separate from this ordering graph.

## Source schema

Add these optional fields:

| Source field      | Type               | Purpose                                                     |
| ----------------- | ------------------ | ----------------------------------------------------------- |
| `etymology`       | String             | Verified historical origin, with uncertainty preserved.     |
| `word_formation`  | String             | Supported morphological analysis or construction patterns.  |
| `mnemonic`        | String             | A concise memory aid, clearly distinguished from fact.      |
| `related_entries` | List of stable IDs | Human-approved direct associations with existing entries.   |

Change `context` from required to optional. Required fields become `id`, `language`, `word`, `meaning`, and `example`.

For `context`, `etymology`, `word_formation`, and `mnemonic`, treat omission, an empty string, a whitespace-only string, and a YAML null value as absent. Export these as empty strings without rewriting source files. Preserve nonempty text as supplied. Continue rejecting other incorrect types, aliases, and unsupported control characters. This relaxation applies only to these four fields; existing optional-field rules otherwise remain unchanged.

Omit unused `related_entries`. If present, require a nonempty list of lowercase kebab-case string IDs, following the existing reference-list conventions. Reject nulls, empty lists, malformed IDs, repeated IDs within one list, self-references, and missing entries.

Keep stable IDs globally unique and unchanged. Add no source fields for derived Anki output. Preserve existing Context values; optionality does not authorize a collection-wide cleanup.

## Content policy

- Use concise Chinese explanations with relevant original-language forms.
- Distinguish historical derivation, modern word formation, and a learner's memory association. Do not present a mnemonic as historical evidence.
- Verify etymological claims against reliable sources. Preserve reconstruction markers and uncertainty. Record source titles and URLs in existing `notes` without replacing existing provenance.
- Omit explanations that cannot be established reliably or do not help the learning target. Do not force a root decomposition onto every word or language.
- Preserve the entry's actual sense. For example, the existing `escape-parity-tech` concerns odd/even escape counts, not feature equivalence.
- Keep explanations brief enough for mobile review. Do not automatically create cards for every root, cognate, or related expression.
- Keep example-vocabulary reuse optional and preserve the existing priority of accuracy, naturalness, and practical context.

## General related-entry model

Associations identify concrete entries by stable ID, including their particular meanings. They can support verified word-family relationships, morphological relationships, confusing spellings, synonym/antonym comparisons, usage contrasts, and useful links between words and expressions. A relationship does not itself assert any particular linguistic category; explanations belong in each entry's `mnemonic`.

Export behavior:

1. Treat a declared association between A and B as an undirected relationship for display. Declaring it at either endpoint makes both card backs show the other entry.
2. Generate only direct neighbors. A–B and B–C must not imply A–C.
3. Deduplicate display when both endpoints declare the same pair. Repeated IDs within an individual source list remain validation errors.
4. Allow cycles, unlike the prerequisite graph. Do not change `prerequisites`, priorities, ranks, or Anki scheduling because of an association.
5. Derive each neighbor's current Word and Meaning during the full build; sort neighbors by stable ID for deterministic output. Do not duplicate these values in source entries or expose IDs as the primary learning content.
6. Render neighbors as plain text lines in the derived `RelatedWords` field, using the normal HTML-escaping path. A word such as `<tag>` or a literal media marker must remain text. Cross-language associations are permitted when explicitly approved; include the neighbor's language tag when it differs from the current entry's language.
7. Export an empty value when there are no neighbors. Removing a relationship or changing a neighbor's content requires a complete rebuild and reimport to refresh both sides. If both endpoints declare a pair, remove both declarations to remove the association.

The first version provides inline comparison text. Cross-card navigation, automatic similarity matching, relationship-type schemas, and additional pronunciation or audio controls for related words are outside scope.

## Human confirmation workflow

Codex may inspect existing entries and propose candidates, but must not write unapproved relationships into source data.

Present candidates in a review table containing both stable IDs, words, current meanings, a proposed relationship explanation, and any ambiguity or confusion risk. Include proposed mnemonic edits when relevant, so the user can approve concrete wording.

The user may approve individual pairs or explicitly approve a batch. Apply only approved pairs and approved accompanying content. Approval of a relationship alone does not authorize unrelated mnemonic rewrites, new entries, or changes to meanings.

Unapproved candidates remain in the discussion; do not persist them as active relationships or include them in Anki exports. Human approval is a maintenance workflow, not a new approval-state field in YAML. Validation checks structural correctness, not linguistic truth or whether a person approved a particular edit.

For the illustrative closure/disclosure pair, distinguish the repository's programming sense of closure from disclosure's reporting sense. Do not label these senses as opposites merely because the words look related. This same review rule applies to all other candidates.

## Proposed Anki field order

Replace the README's field-name code block with a Markdown table numbered to match the Anki Fields UI. The sequence number is a documentation aid, not the stable `ID` value.

Preserve the first twelve fields and append four new fields:

| No. | Anki field    | Source or derivation                          |
| --- | ------------- | --------------------------------------------- |
| 1   | ID            | `id`; permanent note identity                 |
| 2   | Language      | Canonicalized `language`                      |
| 3   | Word          | `word`                                        |
| 4   | Pronunciation | Optional `pronunciation`                      |
| 5   | PartOfSpeech  | Optional `part_of_speech`                     |
| 6   | Meaning       | `meaning`                                     |
| 7   | Context       | Optional `context`                            |
| 8   | Example       | `example`                                     |
| 9   | Notes         | Optional `notes`                              |
| 10  | Audio         | Sound reference derived from optional `audio` |
| 11  | TTSLanguage   | Derived example-speech locale                 |
| 12  | StudyOrder    | Derived introduction rank                     |
| 13  | Etymology     | Optional `etymology`                          |
| 14  | WordFormation | Optional `word_formation`                     |
| 15  | Mnemonic      | Optional `mnemonic`                           |
| 16  | RelatedWords  | Derived direct associations                   |

The proposed TSV has sixteen note fields and Anki tags in column 17. Update the column headers and tag mapping together. Do not create a Tags note field or change first-field ID matching. Keep `StudyOrder` as the browser sort field.

## Card presentation

- Preserve the current front, including pronunciation and word audio. Keep translations and memory aids on the back.
- On the back, retain Word, pronunciation/audio, part of speech, Meaning, optional Context, and Example. Follow these with Etymology, WordFormation, Mnemonic, RelatedWords, then Notes.
- Use separate conditional blocks for Context and each new display field. Put the section label and spacing-bearing wrapper inside the condition. Missing content must produce no visible label, empty container, divider, or reserved space.
- The note type still contains the fields even when a particular card does not display them. Anki's field editor is not dynamically restructured per entry.
- Keep the existing mobile typography, semantic theme tokens, wrapping behavior, and native replay controls. Add no JavaScript merely to hide empty content and no heavy nested card decoration.
- Use a neutral related-vocabulary label. Do not imply every neighbor is a synonym, antonym, or prerequisite.

## Implementation sequence after approval

1. Update Entry parsing and validation for the four optional text fields and `related_entries`. Retain current behavior for all other fields.
2. Validate association references against the complete collection and derive deterministic direct-neighbor displays independently of learning order.
3. Extend TSV output to the proposed seventeen-column layout, preserving text escaping, stable identity, and validation-before-replacement behavior.
4. Update conditional back-template sections and the numbered field table in the README. Keep templates and field order canonical there.
5. Update [AGENTS.md](../AGENTS.md), the [Codex vocabulary guide](ADD_VOCABULARY_WITH_CODEX.md), and the [Anki migration guide](ANKI_MIGRATION_GUIDE.md). Explain candidate approval, optional explanations, blank-value handling, and the new tag-column position. Update other current mapping references without rewriting historical verification records.
6. Extend audio source-preservation checks and tests to include every new field. Do not change spoken-input selection, voice behavior, filenames, or existing recordings.
7. Add focused regression tests and complete the verification below. Use synthetic fixtures for association tests; do not treat them as approved content for the real collection.
8. Propose real-entry enrichment and association candidates separately for human confirmation. Do not perform an automatic whole-collection enrichment or Context deletion pass.

## Migration and risks

Existing entries remain valid without the new fields. Existing note types need the four appended fields, updated back template and styling as needed, and a complete reimport. Keep notes and IDs intact rather than deleting or recreating notes. No audio regeneration is required.

Verify that importing an empty Context or memory-aid field actually clears an old Anki value under the documented update settings. Verify that changed or removed relationships refresh both cards. Do not assume field order alone establishes correct import mapping.

Risks include inaccurate etymology, misleading associations between different senses, excessive back-side text, accidental changes to blank-value rules for unrelated fields, stale reverse displays, and mismapped TSV columns. Mitigate these with source checks, human content approval, concise conditional sections, strict scoped validation, full builds, and isolated import tests.

Do not modify a personal Anki collection, perform synchronization, stage, commit, or push as part of creating this plan. Future implementation should document the manual migration and report any live-device verification limits.

## Verification and acceptance

After implementation, run from the repository root:

```bash
go test ./...
go vet ./...
go run scripts/validate/main.go
go run scripts/build/main.go
git diff --check
```

Also verify:

- Omitted, empty, whitespace-only, and null values work for the four designated text fields. Nonempty values survive export; required fields and other optional-field rules retain their existing validation.
- Association validation rejects wrong types, malformed IDs, empty lists, duplicate list members, self-links, and missing targets. Reciprocal declarations and cycles are valid.
- Single-endpoint declarations render both ways, repeated pair declarations deduplicate, chains remain non-transitive, and neighbor order is independent of declaration order.
- Updating a neighbor's Word or Meaning updates the derived display. Deleting or removing references behaves as documented; invalid references leave the existing export untouched.
- Associations alone do not alter introduction ranks or prerequisite validation. HTML-like text and literal audio markers remain inert in all new display fields.
- Two builds produce identical bytes, all rows have seventeen columns, and tags occupy column 17.
- An isolated Anki collection can migrate from the current twelve-field layout, import twice without duplicates, update and clear optional values, and preserve note/card identities, scheduling, and review history.
- Missing-field combinations leave no empty sections on card backs. Long multilingual content remains readable in light and dark modes on narrow screens. Existing word audio and example TTS still work.
- Audio metadata insertion preserves all new source values and unrelated bytes. Existing entries and media remain unchanged outside specifically approved content edits.
- Documentation links and numbered field mappings agree. Report automated, isolated-backend, and live Anki Desktop/AnkiDroid checks separately; do not claim unperformed playback or visual checks passed.

No implementation or runtime verification has been performed by creating this plan.
