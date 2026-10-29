# English airline meal vocabulary import

Source: the user-provided approved list for airline in-flight meal conversation practice. The list defines the complete import scope; no additional vocabulary was selected.

## Import result

- Approved items: 62; new entries: 62; existing entries reused or updated: 0.
- Sentence pattern: 1; base vocabulary: 18 (11 meal words and 7 cooking descriptions or ingredients); food compounds: 14; meal-option phrases: 15; complete sentences: 14.
- Skipped approved items: 0. The collection contains 346 entries after this import.
- Base words, compounds, meal descriptions, and complete questions remain separate learning units. All supplied word forms, full sentences, and Chinese meanings are preserved, including the question containing `pork rice`.
- The A/B pattern is one entry with explained placeholders. No separate placeholder entries, unlisted combinations, document headings, commands, or excluded greeting/filler examples were added.

## Pronunciation and audio

Pronunciation is omitted for all 62 new entries because verified IPA was not supplied or available in matching existing entries. No separate pronunciation research was performed. Audio is omitted for the new entries; all existing recordings are preserved. The import requires no audio generation.

Anki collections, review state, AnkiWeb synchronization, and AnkiDroid were not modified.

## Verification

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

## Approved item index

### Sentence pattern

| Approved item          | Entry                                                         |
|:-----------------------|:--------------------------------------------------------------|
| Would you like A or B? | [would-you-like-a-or-b](../entries/would-you-like-a-or-b.yml) |

### Base vocabulary

| Approved item  | Entry                                                     |
|:---------------|:----------------------------------------------------------|
| chicken        | [chicken-food](../entries/chicken-food.yml)               |
| beef           | [beef-food](../entries/beef-food.yml)                     |
| pork           | [pork-food](../entries/pork-food.yml)                     |
| fish           | [fish-food](../entries/fish-food.yml)                     |
| seafood        | [seafood-food](../entries/seafood-food.yml)               |
| rice           | [rice-food](../entries/rice-food.yml)                     |
| noodles        | [noodles-food](../entries/noodles-food.yml)               |
| pasta          | [pasta-food](../entries/pasta-food.yml)                   |
| potatoes       | [potatoes-food](../entries/potatoes-food.yml)             |
| curry          | [curry-food](../entries/curry-food.yml)                   |
| sauce          | [sauce-food](../entries/sauce-food.yml)                   |
| roast          | [roast-food](../entries/roast-food.yml)                   |
| braised        | [braised-food](../entries/braised-food.yml)               |
| teriyaki       | [teriyaki-food](../entries/teriyaki-food.yml)             |
| sweet and sour | [sweet-and-sour-food](../entries/sweet-and-sour-food.yml) |
| mashed         | [mashed-food](../entries/mashed-food.yml)                 |
| black pepper   | [black-pepper-food](../entries/black-pepper-food.yml)     |
| mushroom       | [mushroom-food](../entries/mushroom-food.yml)             |

### Food compounds

| Approved item       | Entry                                                               |
|:--------------------|:--------------------------------------------------------------------|
| mashed potatoes     | [mashed-potatoes-food](../entries/mashed-potatoes-food.yml)         |
| mushroom sauce      | [mushroom-sauce-food](../entries/mushroom-sauce-food.yml)           |
| black pepper sauce  | [black-pepper-sauce-food](../entries/black-pepper-sauce-food.yml)   |
| chicken curry       | [chicken-curry-food](../entries/chicken-curry-food.yml)             |
| teriyaki chicken    | [teriyaki-chicken-food](../entries/teriyaki-chicken-food.yml)       |
| roast chicken       | [roast-chicken-food](../entries/roast-chicken-food.yml)             |
| braised beef        | [braised-beef-food](../entries/braised-beef-food.yml)               |
| braised chicken     | [braised-chicken-food](../entries/braised-chicken-food.yml)         |
| sweet and sour pork | [sweet-and-sour-pork-food](../entries/sweet-and-sour-pork-food.yml) |
| beef noodles        | [beef-noodles-food](../entries/beef-noodles-food.yml)               |
| chicken noodles     | [chicken-noodles-food](../entries/chicken-noodles-food.yml)         |
| chicken pasta       | [chicken-pasta-food](../entries/chicken-pasta-food.yml)             |
| beef pasta          | [beef-pasta-food](../entries/beef-pasta-food.yml)                   |
| seafood pasta       | [seafood-pasta-food](../entries/seafood-pasta-food.yml)             |

### Meal options

| Approved item                         | Entry                                                                                                         |
|:--------------------------------------|:--------------------------------------------------------------------------------------------------------------|
| chicken with rice                     | [chicken-with-rice-airline](../entries/chicken-with-rice-airline.yml)                                         |
| beef with noodles                     | [beef-with-noodles-airline](../entries/beef-with-noodles-airline.yml)                                         |
| chicken curry with rice               | [chicken-curry-with-rice-airline](../entries/chicken-curry-with-rice-airline.yml)                             |
| teriyaki chicken with rice            | [teriyaki-chicken-with-rice-airline](../entries/teriyaki-chicken-with-rice-airline.yml)                       |
| pork with noodles                     | [pork-with-noodles-airline](../entries/pork-with-noodles-airline.yml)                                         |
| braised beef with rice                | [braised-beef-with-rice-airline](../entries/braised-beef-with-rice-airline.yml)                               |
| roast chicken with potatoes           | [roast-chicken-with-potatoes-airline](../entries/roast-chicken-with-potatoes-airline.yml)                     |
| beef with rice                        | [beef-with-rice-airline](../entries/beef-with-rice-airline.yml)                                               |
| pork with rice                        | [pork-with-rice-airline](../entries/pork-with-rice-airline.yml)                                               |
| chicken with mashed potatoes          | [chicken-with-mashed-potatoes-airline](../entries/chicken-with-mashed-potatoes-airline.yml)                   |
| sweet and sour pork with rice         | [sweet-and-sour-pork-with-rice-airline](../entries/sweet-and-sour-pork-with-rice-airline.yml)                 |
| beef with black pepper sauce and rice | [beef-with-black-pepper-sauce-and-rice-airline](../entries/beef-with-black-pepper-sauce-and-rice-airline.yml) |
| fish with potatoes                    | [fish-with-potatoes-airline](../entries/fish-with-potatoes-airline.yml)                                       |
| chicken with mushroom sauce           | [chicken-with-mushroom-sauce-airline](../entries/chicken-with-mushroom-sauce-airline.yml)                     |
| fish with rice                        | [fish-with-rice-airline](../entries/fish-with-rice-airline.yml)                                               |

### Complete sentences

| Approved item                                                               | Entry                                                                                                                                 |
|:----------------------------------------------------------------------------|:--------------------------------------------------------------------------------------------------------------------------------------|
| Would you like chicken with rice or beef with noodles?                      | [airline-chicken-rice-or-beef-noodles](../entries/airline-chicken-rice-or-beef-noodles.yml)                                           |
| Would you like chicken curry with rice or beef noodles?                     | [airline-chicken-curry-rice-or-beef-noodles](../entries/airline-chicken-curry-rice-or-beef-noodles.yml)                               |
| Would you like teriyaki chicken with rice or pork with noodles?             | [airline-teriyaki-chicken-rice-or-pork-noodles](../entries/airline-teriyaki-chicken-rice-or-pork-noodles.yml)                         |
| Would you like braised beef with rice or chicken pasta?                     | [airline-braised-beef-rice-or-chicken-pasta](../entries/airline-braised-beef-rice-or-chicken-pasta.yml)                               |
| Would you like roast chicken with potatoes or beef with rice?               | [airline-roast-chicken-potatoes-or-beef-rice](../entries/airline-roast-chicken-potatoes-or-beef-rice.yml)                             |
| Would you like pork with rice or seafood pasta?                             | [airline-pork-rice-or-seafood-pasta](../entries/airline-pork-rice-or-seafood-pasta.yml)                                               |
| Would you like chicken with mashed potatoes or beef noodles?                | [airline-chicken-mashed-potatoes-or-beef-noodles](../entries/airline-chicken-mashed-potatoes-or-beef-noodles.yml)                     |
| Would you like sweet and sour pork with rice or chicken noodles?            | [airline-sweet-and-sour-pork-rice-or-chicken-noodles](../entries/airline-sweet-and-sour-pork-rice-or-chicken-noodles.yml)             |
| Would you like beef with black pepper sauce and rice or fish with potatoes? | [airline-beef-black-pepper-sauce-and-rice-or-fish-potatoes](../entries/airline-beef-black-pepper-sauce-and-rice-or-fish-potatoes.yml) |
| Would you like chicken with mushroom sauce or beef pasta?                   | [airline-chicken-mushroom-sauce-or-beef-pasta](../entries/airline-chicken-mushroom-sauce-or-beef-pasta.yml)                           |
| Would you like pork rice or chicken curry?                                  | [airline-pork-rice-or-chicken-curry](../entries/airline-pork-rice-or-chicken-curry.yml)                                               |
| Would you like fish with rice or beef with noodles?                         | [airline-fish-rice-or-beef-noodles](../entries/airline-fish-rice-or-beef-noodles.yml)                                                 |
| Would you like chicken pasta or pork with rice?                             | [airline-chicken-pasta-or-pork-rice](../entries/airline-chicken-pasta-or-pork-rice.yml)                                               |
| Would you like braised chicken with rice or beef pasta?                     | [airline-braised-chicken-rice-or-beef-pasta](../entries/airline-braised-chicken-rice-or-beef-pasta.yml)                               |
