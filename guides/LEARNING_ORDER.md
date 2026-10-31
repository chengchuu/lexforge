# Learning Order

Use introduction priorities and prerequisites to order new cards, then apply the generated ranks in Anki. Anki owns learning state and review scheduling; applying this order remains a manual step.

## Model

LexForge proposes an introduction order using editorial priorities and selected prerequisites. These describe the content, not your current mastery, and are not CEFR or JLPT levels.

- `1`: foundations and early practical use, including common words, greetings, and useful everyday patterns.
- `2`: standard or applied vocabulary; the default when the field is omitted.
- `3`: specialized concepts, less common forms, and complex quotations.

`learning_priority` is an optional integer from 1 to 3 and defaults to 2.

`prerequisites` lists existing stable IDs. The validator rejects invalid or repeated IDs, missing references, self-references, and cycles. Dependencies describe useful foundations, not all vocabulary in an example. Cross-language references are allowed but should have an explicit learning reason; the initial classification uses only same-language links.

Use prerequisites where learning another existing item first helps explain the target. For example, `mashed-potatoes-food` depends on `potato-food` and `mashed-food`:

```yaml
learning_priority: 2
prerequisites:
  - potato-food
  - mashed-food
```

A short technical term can be advanced; a longer greeting can be foundational. Do not split fixed expressions or link every word in an example. Select dependencies by meaning and grammatical relationship, not shared spelling. Keep example-vocabulary reuse optional.

The build repeatedly selects an entry whose prerequisites have already been placed, preferring the lowest priority number and then stable ID for ties. Prerequisites take precedence over priority. Unrelated entries of equal priority still use ID order; the result is a deterministic starting curriculum, not a precisely measured difficulty scale or a guarantee that related cards are adjacent.

The exporter derives `StudyOrder`: a unique, zero-padded rank shared across all languages. It is not stored in YAML or shown on the cards. Ranks may change when entries or metadata change; IDs remain permanent. Import the complete rebuilt export to refresh ranks consistently. See the [initial classification record](#initial-classification).

## Apply the order in Anki

The current export has sixteen note fields plus tags in column 17. See the [README field mapping](../README.md#set-up-anki-once) for note-type setup.

Import order alone does not rearrange existing new cards. Anki's [Sort Field](https://docs.ankiweb.net/editing.html#customizing-fields) and [Reposition](https://docs.ankiweb.net/browsing.html#cards) controls apply the generated ranks:

1. In **Tools → Manage Note Types → LexForge → Fields**, add `StudyOrder` after `TTSLanguage` if missing. Select `StudyOrder` and enable **Sort by this field in the browser**. Keep `ID` as the first field for import matching; choosing the sort field does not move it.
2. Build and [import the complete TSV](../README.md#import-and-update), updating existing notes. Verify `StudyOrder` maps to column 12 and tags to column 17.
3. In the deck's options, use **Sequential** insertion order, **Ascending position** for **New card gather order**, and **Order gathered** for **New card sort order**. Set insertion order before repositioning: changing it can reorder existing new cards. These are Anki's [display-order settings](https://docs.ankiweb.net/deck-options.html#display-order).
4. Empty any filtered decks containing these cards. In **Browse**, use **Cards** mode and search `note:LexForge is:new`. Show **Sort Field** and **Due** columns; sort **Sort Field** ascending. Confirm the ranks increase from top to bottom.
5. Select all results and choose **Cards → Reposition**. Set the start position to **1**, step to **1**, and turn off **Randomize order** and **Shift position of existing cards**. Check that the Due positions now increase in the same order.
6. Sync Anki Desktop, then AnkiDroid. After later imports or priority changes, repeat the sort and reposition steps for the remaining new cards.

Use a dedicated LexForge deck so unrelated new cards do not share the same positions. Reposition the whole LexForge new-card selection, including suspended new cards, to avoid position collisions from partial selections. To study languages separately, move each language to a dedicated deck, then apply the same workflow to each entire deck using, for example, `note:LexForge deck:LexForge::Japanese is:new`. Review order is governed by the deck you choose to study; parent-deck daily limits can affect which cards are introduced.

Keep the `is:new` filter. Do not reset learned cards or use **Set Due Date** for this task. Repositioning new cards leaves learning/review cards and their history intact. A prerequisite can still be suspended, excluded from your chosen deck, or not yet mastered: these links order introductions and do not unlock or block cards based on learning state. Changing source metadata does not update Anki until you import and reposition.

## Initial classification

This section records the initial classification on 2026-10-30. Counts describe that snapshot.

All 530 entries received explicit priorities:

| Language | Priority 1 | Priority 2 | Priority 3 | Total |
| -------- | ---------- | ---------- | ---------- | ----- |
| English  | 96         | 174        | 35         | 305   |
| Japanese | 75         | 109        | 41         | 225   |
| Total    | 171        | 283        | 76         | 530   |

There are 157 prerequisite links across 93 entries. Examples include:

- `pepper-food` → `black-pepper-food`.
- `potato-food` and `mashed-food` → `mashed-potatoes-food`.
- `diagnostic-tech` → `diagnostic-context-tech` and `false-diagnostic-tech`.
- `hanasu` → `hanasemasen` → `nihongo-ga-hanasemasen`, which also depends on `nihongo`.
- `naru` → `ni-naru-grammar` → `ni-nareru-grammar`, followed by a quotation using that construction.
- `shinu` and `te-miru-grammar` → `shinde-miru` → `ippen-shinde-miru`, which also depends on `ippen`.

Common greetings and the practical `Would you like A or B?` pattern are priority 1 despite their length. Short specialized compiler terms and isolated kanji readings can be priority 3. Distinct senses remain separate; for example, a TypeScript type assertion is not made dependent on a testing assertion merely because they share a word.

The historical form `chisou` follows the more useful modern `gochisou`, rather than forcing all shorter forms to come first. Familiarity, practical use, and the entry's actual meaning guide the order.

This classification can evolve as study priorities change. Edit only the relevant metadata and rebuild the complete export. Do not treat ranks as stable IDs or reset review progress when ranks change.

## Verification

The following results were recorded for the original learning-order implementation on 2026-10-30, before the later memory-aid fields expanded the export. They are historical results, not checks rerun for this documentation move.

Regression tests cover omitted metadata, valid and invalid types and ranges, missing/repeated/self references, cycles, prerequisite precedence over priority, shared dependencies, chained dependencies, stable tie-breaking, rank padding, deterministic builds, source immutability, and preservation of an existing export after validation failure. Audio insertion tests also cover the new source fields.

The full Go tests, vet, vocabulary validation, two byte-identical builds, and diff checks passed. The export contains 530 thirteen-column rows with unique ranks and every prerequisite before its dependent. The five pre-existing duplicate-word warnings remain for distinct senses of artifact, literal, resolve, revision, and scope.

An isolated Anki Python backend collection imported the previous eleven-field layout, migrated to twelve fields, and imported the new export twice without adding duplicate notes or changing card identities, scheduling, or review history. Sorting by the browser's Sort Field and repositioning changed the order of 528 new cards, including a suspended card, while preserving the learning and review cards created for the test.

An audit against the pre-change snapshot confirmed that all original entry text outside the added metadata, all 530 MP3 files, and both card templates and their CSS remained unchanged. No recordings were regenerated and no personal Anki collection was modified. Live Anki Desktop and AnkiDroid UI behavior was not tested.
