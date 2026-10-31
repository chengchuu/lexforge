# Learning Order

Status: implemented on 2026-10-30. The source schema, validation, export, and README workflow support introduction priorities and prerequisites. Applying the order to a personal Anki collection remains a manual step.

## Model

`learning_priority` is an optional integer: 1 for foundations and early practical use, 2 for standard or applied vocabulary, and 3 for specialized material. Omission means 2. These are editorial choices for this collection, not standardized proficiency levels or a measurement of the learner's knowledge.

`prerequisites` lists existing stable IDs. The validator rejects invalid or repeated IDs, missing references, self-references, and cycles. Dependencies describe useful foundations, not all vocabulary in an example. Cross-language references are allowed but should have an explicit learning reason; the initial classification uses only same-language links.

The export places prerequisites before their dependents. Among currently eligible entries, lower priority numbers come first, with stable ID as a deterministic tie-breaker. The derived `StudyOrder` field contains a zero-padded rank. It is separate from permanent note identity and can change after any source update.

## Initial classification

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

## Anki application

The export now has sixteen note fields plus tags in column 17. Add `StudyOrder` after `TTSLanguage` and select it as the browser sort field; keep `ID` first for import matching. Import the complete export, then use the [README procedure](../README.md#apply-the-order-in-anki) to sort and reposition new cards.

Anki retains learning state and review scheduling. Prerequisites do not imply mastery, unlock cards, or override suspended cards, deck selection, or daily limits. Previously introduced cards continue their existing schedules.

## Verification

Regression tests cover omitted metadata, valid and invalid types and ranges, missing/repeated/self references, cycles, prerequisite precedence over priority, shared dependencies, chained dependencies, stable tie-breaking, rank padding, deterministic builds, source immutability, and preservation of an existing export after validation failure. Audio insertion tests also cover the new source fields.

The full Go tests, vet, vocabulary validation, two byte-identical builds, and diff checks passed. The export contains 530 thirteen-column rows with unique ranks and every prerequisite before its dependent. The five pre-existing duplicate-word warnings remain for distinct senses of artifact, literal, resolve, revision, and scope.

An isolated Anki Python backend collection imported the previous eleven-field layout, migrated to twelve fields, and imported the new export twice without adding duplicate notes or changing card identities, scheduling, or review history. Sorting by the browser's Sort Field and repositioning changed the order of 528 new cards, including a suspended card, while preserving the learning and review cards created for the test.

An audit against the pre-change snapshot confirmed that all original entry text outside the added metadata, all 530 MP3 files, and both card templates and their CSS remained unchanged. No recordings were regenerated and no personal Anki collection was modified. Live Anki Desktop and AnkiDroid UI behavior was not tested.
