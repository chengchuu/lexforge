# Add vocabulary with Codex

Open Codex in the repository root and use the following prompt. Replace the placeholders with your vocabulary and language; add context if useful, and specify whether to generate audio.

```text
Add the following vocabulary to LexForge.

Words, expressions, or concepts:
- <word or expression>

Language: <English, Japanese, Korean, or another language>
Encountered in: <sentence, technical context, or intended meaning>
Generate audio: <yes or no; default: no>

Read AGENTS.md and inspect the current schema, existing entries,
and available commands before making changes.

Search for existing entries covering the same language and meaning.
Reuse or update matching entries without changing their stable IDs.
Create a separate entry only for a distinct meaning.

For supplied sentences or descriptions, identify the useful learning
targets and check existing component vocabulary first. Keep ordinary
usage in examples or notes instead of creating a card for each sentence
or combination. Preserve useful fixed expressions and grammar patterns;
do not split them mechanically or enforce a word-count limit.
Prefer dictionary headwords where appropriate, retaining meaningful
inflections and natural plural forms. Report any grouping decisions.

For compounds and meaningful phrases, identify useful independent
components before introducing the whole. Reuse matching entries and add
missing components within the approved scope. Use prerequisites to put
components first while retaining the whole term's distinct meaning.
Choose component priorities by their own usefulness and complexity;
prerequisites enforce ordering even when priority numbers are equal.
Keep nonliteral technical meanings and fixed expressions intact.

For names with limited context, check for recognizable components that
have independent learning value. Verify their ordinary meanings and
readings, reuse existing entries first, and add missing useful components
within the requested scope. Put their stable IDs in the whole name's
prerequisites so components are introduced first. Preserve the whole
name's ID, spelling, and recording. Explain composition in Notes without
claiming a verified naming origin, real-world identity, or business scope.
Do not guess ambiguous kanji, expand unknown initials, or split names
mechanically. Prerequisites order introductions; they do not unlock
cards based on Anki mastery. Associations require separate approval.

For each entry:
- Use a globally unique stable ID and lowercase kebab-case filename.
- Add a self-contained Chinese meaning or short explanation of a principle.
  Include context and a natural example in the entry's language only when useful.
  Use Notes for detailed steps, reasoning, limitations, and sources; do not
  truncate useful notes. Put Chinese walkthroughs there, outside Example TTS.
- Add verified etymology or word formation only when useful. Distinguish
  historical facts from mnemonics and preserve source titles and URLs in Notes.
- Propose related_entries candidates with both stable IDs, words, meanings,
  rationale, and confusion risks. Write only human-approved pairs; the exporter
  displays them both ways. Get approval for accompanying mnemonic edits.
  Associations do not imply prerequisites.
- Reuse relevant existing vocabulary in the same language when it fits
  naturally in the example. Otherwise, use a more suitable example freely.
  Accuracy, naturalness, and practical context take priority over reuse;
  do not force vocabulary overlap.
- Prefer software-engineering usage when it matches the supplied context.
- Choose learning_priority: 1 for foundations and early practical use,
  2 for standard or applied vocabulary, or 3 for specialized material.
  If uncertain, omit it to use 2. Judge the learning target, not word length.
- Add prerequisites only for existing stable IDs that help explain the
  target. Do not infer dependencies from every word in its example or
  create extra entries solely to satisfy a dependency. Omit when unused.
- Follow existing tag conventions and omit unused optional fields.
- Include pronunciation only when reliable. For Japanese, use kana
  rather than romaji.
- Preserve unrelated content and existing audio.

Use only fields and language capabilities supported by the current
implementation. A proposed plan is not an implemented feature.
If the requested language is unsupported, explain the limitation
before adding entries; do not change the schema or tooling.

If audio is requested, use the existing generator only when it
supports the entry's language and a suitable voice is available.
Preserve existing recordings unless replacement is explicitly requested.
Report unavailable audio without substituting another language's voice.

Run validation and rebuild the Anki export using the documented commands.
Never edit generated output manually.

Review the diff and report added or updated entries, validation/build
results, and any pronunciation or audio limitations.
Do not commit, push, or modify the user's Anki collection.
```

See the [entry schema](ENTRY_SCHEMA.md), [learning-order rules](LEARNING_ORDER.md), and [audio workflow](../README.md#audio) for the current requirements and commands. After reviewing the changes, [import the generated TSV into Anki](../README.md#import-and-update).
