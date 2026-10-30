# English airline meal vocabulary import

## Status

Vocabulary-focused cleanup implemented on 2026-10-29 after user approval. The original 62 cards are consolidated into 21 learning items: 18 base items (with `potatoes` normalized to `potato`), one additional `pepper` entry, `mashed potatoes`, and the reusable question pattern. The collection now contains 530 entries (305 English and 225 Japanese).

All 62 original items remain accounted for below. The original import results and checks are historical; the cleanup and its verification are recorded separately.

Source: the user-provided approved list for airline in-flight meal conversation practice. The list defines the complete import scope; no additional vocabulary was selected.

## Original import result

- Approved items: 62; new entries: 62; existing entries reused or updated: 0.
- Sentence pattern: 1; base vocabulary: 18 (11 meal words and 7 cooking descriptions or ingredients); food compounds: 14; meal-option phrases: 15; complete sentences: 14.
- Skipped approved items: 0. The collection contains 346 entries after this import.
- Base words, compounds, meal descriptions, and complete questions remain separate learning units. All supplied word forms, full sentences, and Chinese meanings are preserved, including the question containing `pork rice`.
- The A/B pattern is one entry with explained placeholders. No separate placeholder entries, unlisted combinations, document headings, commands, or excluded greeting/filler examples were added.

## Original pronunciation and audio

Pronunciation is omitted for all 62 new entries because verified IPA was not supplied or available in matching existing entries. No separate pronunciation research was performed. Audio was omitted during the original import, which preserved all existing recordings. The later audio follow-up generated these 62 recordings from their English words using Samantha; all passed decoding.

Anki collections, review state, AnkiWeb synchronization, and AnkiDroid were not modified.

## Original verification

- `go test ./...`: passed.
- `go vet ./...`: passed.
- `go run scripts/validate/main.go`: passed for 346 entries.
- `go run scripts/build/main.go`: passed; two builds produced identical bytes.
- `git diff --check`: passed; new files were also checked for trailing whitespace and final newlines.
- All 62 approved items have one corresponding entry. IDs are globally unique; no unapproved entries were added.
- All 346 export rows have eleven columns. The 62 new rows have `en`, the expected content and derived `lang-en` tag, and blank pronunciation and audio columns.
- All 284 existing entry files and 284 MP3 files remain byte-identical.
- The five existing duplicate-word warnings concern distinct senses of artifact, literal, resolve, revision, and scope. This import adds no duplicate-word warnings.
- No Go code, schema, dependency, or card-template changes were required. No Anki import or device testing was performed.

## Vocabulary-focused cleanup

- Retain one entry per useful learning target. Keep `mashed potatoes` as a named side dish whose preparation is part of its meaning. Keep `black pepper` as a specific seasoning, `sweet and sour` as a flavor expression, and `Would you like A or B?` as the reusable grammar pattern.
- Consolidate the other 13 compounds into their component vocabulary. Cooking-method, sauce, meat-and-starch, and flavor combinations remain in examples or notes; this batch does not need an independent card for every menu combination. The compound table records each decision.
- Consolidate all 15 meal options into retained entries' examples. Preserve all 14 complete questions and their Chinese meanings in the source index below, linked to the shared grammar pattern. The pattern keeps one representative question as its example.
- Rename `potatoes-food` to `potato-food`, including its source filename and recording. The plural remains in the example and notes. This is an explicitly approved identity change before the first release; all other retained IDs are unchanged. Keep `noodles` as the usual food form and `braised`/`mashed` as cooking adjectives.
- Add `pepper-food` for the seasoning sense, distinguished from vegetable peppers and from the narrower `black-pepper-food` entry. IPA remains unset because this cleanup did not research pronunciation sources.
- Remove 42 retired entry files and their recordings, plus the old plural-ID source and recording. Generate new word recordings for `pepper-food` and `potato-food`; preserve all 528 surviving original recordings.
- Update the README prompt and AGENTS.md selection guidance. No schema, exporter, template, or audio-tool changes are needed. No other vocabulary batch is consolidated.

Removing files from this repository does not remove previously imported Anki notes or media. No live Anki collection is modified by this cleanup.

### Cleanup verification

- `go test ./...`, `go vet ./...`, vocabulary validation, the export build, and `git diff --check` passed. The five existing duplicate-word warnings remain unchanged.
- Two export builds produced identical bytes: 530 twelve-column data rows, including 305 English and 225 Japanese entries, with matching audio references, TTS locales, and derived language tags.
- All 62 original items and their Chinese meanings are accounted for in the tables below, mapping to 21 current learning items. All documentation file links resolve. YAML and Markdown passed LF, final-newline, and trailing-whitespace checks.
- Both new recordings use Samantha at 150 words per minute and LAME at 96 kbps mono, generated from `word` through the existing single-ID command. Both decoded to nonempty PCM samples. No existing recording was overwritten.
- All 509 entries outside this batch and all 528 surviving original recordings retain their original bytes. All 530 entries reference existing MP3s, and no unreferenced MP3 files remain in `media/`.
- Audio was not independently auditioned. No live Anki import, playback, synchronization, or review-state changes were performed.

## Source item accounting

Original IDs below identify the imported records, including retired ones. Links point only to current entries. The original wording and Chinese meanings are retained; `pork rice` remains source wording rather than a newly normalized expression.

### Sentence pattern

| Original item          | Meaning           | Original ID             | Current coverage                                                        |
|:-----------------------|:------------------|:------------------------|:------------------------------------------------------------------------|
| Would you like A or B? | 您想要 A 还是 B？ | `would-you-like-a-or-b` | Retained: [would-you-like-a-or-b](../entries/would-you-like-a-or-b.yml) |

### Base vocabulary

| Original item  | Meaning          | Original ID           | Current coverage                                                    |
|:---------------|:-----------------|:----------------------|:--------------------------------------------------------------------|
| chicken        | 鸡肉             | `chicken-food`        | Retained: [chicken-food](../entries/chicken-food.yml)               |
| beef           | 牛肉             | `beef-food`           | Retained: [beef-food](../entries/beef-food.yml)                     |
| pork           | 猪肉             | `pork-food`           | Retained: [pork-food](../entries/pork-food.yml)                     |
| fish           | 鱼；鱼肉         | `fish-food`           | Retained: [fish-food](../entries/fish-food.yml)                     |
| seafood        | 海鲜             | `seafood-food`        | Retained: [seafood-food](../entries/seafood-food.yml)               |
| rice           | 米饭             | `rice-food`           | Retained: [rice-food](../entries/rice-food.yml)                     |
| noodles        | 面条             | `noodles-food`        | Retained: [noodles-food](../entries/noodles-food.yml)               |
| pasta          | 意大利面；意面   | `pasta-food`          | Retained: [pasta-food](../entries/pasta-food.yml)                   |
| potatoes       | 土豆；马铃薯     | `potatoes-food`       | Normalized: [potato-food](../entries/potato-food.yml)               |
| curry          | 咖喱             | `curry-food`          | Retained: [curry-food](../entries/curry-food.yml)                   |
| sauce          | 酱汁             | `sauce-food`          | Retained: [sauce-food](../entries/sauce-food.yml)                   |
| roast          | 烤制的；烘烤的   | `roast-food`          | Retained: [roast-food](../entries/roast-food.yml)                   |
| braised        | 炖的；焖制的     | `braised-food`        | Retained: [braised-food](../entries/braised-food.yml)               |
| teriyaki       | 照烧             | `teriyaki-food`       | Retained: [teriyaki-food](../entries/teriyaki-food.yml)             |
| sweet and sour | 糖醋的；酸甜的   | `sweet-and-sour-food` | Retained: [sweet-and-sour-food](../entries/sweet-and-sour-food.yml) |
| mashed         | 捣碎的；做成泥的 | `mashed-food`         | Retained: [mashed-food](../entries/mashed-food.yml)                 |
| black pepper   | 黑胡椒           | `black-pepper-food`   | Retained: [black-pepper-food](../entries/black-pepper-food.yml)     |
| mushroom       | 蘑菇             | `mushroom-food`       | Retained: [mushroom-food](../entries/mushroom-food.yml)             |

### Food compounds

| Original item       | Meaning        | Original ID                | Current coverage                                                                                               |
|:--------------------|:---------------|:---------------------------|:---------------------------------------------------------------------------------------------------------------|
| mashed potatoes     | 土豆泥         | `mashed-potatoes-food`     | Retained: [mashed-potatoes-food](../entries/mashed-potatoes-food.yml)                                          |
| mushroom sauce      | 蘑菇酱汁       | `mushroom-sauce-food`      | Consolidated: [mushroom-food](../entries/mushroom-food.yml), [sauce-food](../entries/sauce-food.yml)           |
| black pepper sauce  | 黑胡椒酱汁     | `black-pepper-sauce-food`  | Consolidated: [black-pepper-food](../entries/black-pepper-food.yml), [sauce-food](../entries/sauce-food.yml)   |
| chicken curry       | 咖喱鸡         | `chicken-curry-food`       | Consolidated: [curry-food](../entries/curry-food.yml), [chicken-food](../entries/chicken-food.yml)             |
| teriyaki chicken    | 照烧鸡肉       | `teriyaki-chicken-food`    | Consolidated: [teriyaki-food](../entries/teriyaki-food.yml), [chicken-food](../entries/chicken-food.yml)       |
| roast chicken       | 烤鸡           | `roast-chicken-food`       | Consolidated: [roast-food](../entries/roast-food.yml), [chicken-food](../entries/chicken-food.yml)             |
| braised beef        | 炖牛肉；焖牛肉 | `braised-beef-food`        | Consolidated: [braised-food](../entries/braised-food.yml), [beef-food](../entries/beef-food.yml)               |
| braised chicken     | 炖鸡；焖鸡     | `braised-chicken-food`     | Consolidated: [braised-food](../entries/braised-food.yml), [chicken-food](../entries/chicken-food.yml)         |
| sweet and sour pork | 糖醋猪肉       | `sweet-and-sour-pork-food` | Consolidated: [sweet-and-sour-food](../entries/sweet-and-sour-food.yml), [pork-food](../entries/pork-food.yml) |
| beef noodles        | 牛肉面         | `beef-noodles-food`        | Consolidated: [noodles-food](../entries/noodles-food.yml), [beef-food](../entries/beef-food.yml)               |
| chicken noodles     | 鸡肉面         | `chicken-noodles-food`     | Consolidated: [noodles-food](../entries/noodles-food.yml), [chicken-food](../entries/chicken-food.yml)         |
| chicken pasta       | 鸡肉意面       | `chicken-pasta-food`       | Consolidated: [pasta-food](../entries/pasta-food.yml), [chicken-food](../entries/chicken-food.yml)             |
| beef pasta          | 牛肉意面       | `beef-pasta-food`          | Consolidated: [pasta-food](../entries/pasta-food.yml), [beef-food](../entries/beef-food.yml)                   |
| seafood pasta       | 海鲜意面       | `seafood-pasta-food`       | Consolidated: [seafood-food](../entries/seafood-food.yml), [pasta-food](../entries/pasta-food.yml)             |

### Meal options

| Original item                         | Meaning            | Original ID                                     | Current coverage                                                     |
|:--------------------------------------|:-------------------|:------------------------------------------------|:---------------------------------------------------------------------|
| chicken with rice                     | 鸡肉配米饭         | `chicken-with-rice-airline`                     | Example: [chicken-food](../entries/chicken-food.yml)                 |
| beef with noodles                     | 牛肉配面条         | `beef-with-noodles-airline`                     | Example: [noodles-food](../entries/noodles-food.yml)                 |
| chicken curry with rice               | 咖喱鸡配米饭       | `chicken-curry-with-rice-airline`               | Example: [curry-food](../entries/curry-food.yml)                     |
| teriyaki chicken with rice            | 照烧鸡肉配米饭     | `teriyaki-chicken-with-rice-airline`            | Example: [teriyaki-food](../entries/teriyaki-food.yml)               |
| pork with noodles                     | 猪肉配面条         | `pork-with-noodles-airline`                     | Example: [pork-food](../entries/pork-food.yml)                       |
| braised beef with rice                | 炖牛肉配米饭       | `braised-beef-with-rice-airline`                | Example: [braised-food](../entries/braised-food.yml)                 |
| roast chicken with potatoes           | 烤鸡配土豆         | `roast-chicken-with-potatoes-airline`           | Example: [roast-food](../entries/roast-food.yml)                     |
| beef with rice                        | 牛肉配米饭         | `beef-with-rice-airline`                        | Example: [beef-food](../entries/beef-food.yml)                       |
| pork with rice                        | 猪肉配米饭         | `pork-with-rice-airline`                        | Example: [rice-food](../entries/rice-food.yml)                       |
| chicken with mashed potatoes          | 鸡肉配土豆泥       | `chicken-with-mashed-potatoes-airline`          | Example: [mashed-potatoes-food](../entries/mashed-potatoes-food.yml) |
| sweet and sour pork with rice         | 糖醋猪肉配米饭     | `sweet-and-sour-pork-with-rice-airline`         | Example: [sweet-and-sour-food](../entries/sweet-and-sour-food.yml)   |
| beef with black pepper sauce and rice | 黑胡椒汁牛肉配米饭 | `beef-with-black-pepper-sauce-and-rice-airline` | Example: [black-pepper-food](../entries/black-pepper-food.yml)       |
| fish with potatoes                    | 鱼肉配土豆         | `fish-with-potatoes-airline`                    | Example: [potato-food](../entries/potato-food.yml)                   |
| chicken with mushroom sauce           | 蘑菇汁鸡肉         | `chicken-with-mushroom-sauce-airline`           | Example: [mushroom-food](../entries/mushroom-food.yml)               |
| fish with rice                        | 鱼肉配米饭         | `fish-with-rice-airline`                        | Example: [fish-food](../entries/fish-food.yml)                       |

### Complete sentences

| Original item                                                               | Meaning                            | Original ID                                                 | Current coverage                                                                                 |
|:----------------------------------------------------------------------------|:-----------------------------------|:------------------------------------------------------------|:-------------------------------------------------------------------------------------------------|
| Would you like chicken with rice or beef with noodles?                      | 您想要鸡肉饭还是牛肉面？           | `airline-chicken-rice-or-beef-noodles`                      | Source practice sentence; pattern: [would-you-like-a-or-b](../entries/would-you-like-a-or-b.yml) |
| Would you like chicken curry with rice or beef noodles?                     | 您想要咖喱鸡饭还是牛肉面？         | `airline-chicken-curry-rice-or-beef-noodles`                | Source practice sentence; pattern: [would-you-like-a-or-b](../entries/would-you-like-a-or-b.yml) |
| Would you like teriyaki chicken with rice or pork with noodles?             | 您想要照烧鸡饭还是猪肉面？         | `airline-teriyaki-chicken-rice-or-pork-noodles`             | Source practice sentence; pattern: [would-you-like-a-or-b](../entries/would-you-like-a-or-b.yml) |
| Would you like braised beef with rice or chicken pasta?                     | 您想要炖牛肉饭还是鸡肉意面？       | `airline-braised-beef-rice-or-chicken-pasta`                | Source practice sentence; pattern: [would-you-like-a-or-b](../entries/would-you-like-a-or-b.yml) |
| Would you like roast chicken with potatoes or beef with rice?               | 您想要烤鸡配土豆还是牛肉饭？       | `airline-roast-chicken-potatoes-or-beef-rice`               | Source practice sentence; pattern: [would-you-like-a-or-b](../entries/would-you-like-a-or-b.yml) |
| Would you like pork with rice or seafood pasta?                             | 您想要猪肉饭还是海鲜意面？         | `airline-pork-rice-or-seafood-pasta`                        | Source practice sentence; pattern: [would-you-like-a-or-b](../entries/would-you-like-a-or-b.yml) |
| Would you like chicken with mashed potatoes or beef noodles?                | 您想要鸡肉配土豆泥还是牛肉面？     | `airline-chicken-mashed-potatoes-or-beef-noodles`           | Source practice sentence; pattern: [would-you-like-a-or-b](../entries/would-you-like-a-or-b.yml) |
| Would you like sweet and sour pork with rice or chicken noodles?            | 您想要糖醋猪肉饭还是鸡肉面？       | `airline-sweet-and-sour-pork-rice-or-chicken-noodles`       | Source practice sentence; pattern: [would-you-like-a-or-b](../entries/would-you-like-a-or-b.yml) |
| Would you like beef with black pepper sauce and rice or fish with potatoes? | 您想要黑胡椒牛肉饭还是鱼肉配土豆？ | `airline-beef-black-pepper-sauce-and-rice-or-fish-potatoes` | Source practice sentence; pattern: [would-you-like-a-or-b](../entries/would-you-like-a-or-b.yml) |
| Would you like chicken with mushroom sauce or beef pasta?                   | 您想要蘑菇汁鸡肉还是牛肉意面？     | `airline-chicken-mushroom-sauce-or-beef-pasta`              | Source practice sentence; pattern: [would-you-like-a-or-b](../entries/would-you-like-a-or-b.yml) |
| Would you like pork rice or chicken curry?                                  | 您想要猪肉饭还是咖喱鸡？           | `airline-pork-rice-or-chicken-curry`                        | Source practice sentence; pattern: [would-you-like-a-or-b](../entries/would-you-like-a-or-b.yml) |
| Would you like fish with rice or beef with noodles?                         | 您想要鱼肉饭还是牛肉面？           | `airline-fish-rice-or-beef-noodles`                         | Source practice sentence; pattern: [would-you-like-a-or-b](../entries/would-you-like-a-or-b.yml) |
| Would you like chicken pasta or pork with rice?                             | 您想要鸡肉意面还是猪肉饭？         | `airline-chicken-pasta-or-pork-rice`                        | Source practice sentence; pattern: [would-you-like-a-or-b](../entries/would-you-like-a-or-b.yml) |
| Would you like braised chicken with rice or beef pasta?                     | 您想要炖鸡饭还是牛肉意面？         | `airline-braised-chicken-rice-or-beef-pasta`                | Source practice sentence; pattern: [would-you-like-a-or-b](../entries/would-you-like-a-or-b.yml) |

Additional learning item: [pepper-food](../entries/pepper-food.yml) — pepper（胡椒；胡椒粉）, added during cleanup.
