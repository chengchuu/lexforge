# LexForge

LexForge turns vocabulary encountered during everyday work into Anki cards with meanings, context, and examples. It supports English, Japanese, Korean, and other languages. Git tracks vocabulary content; Anki manages review history and scheduling.

- [Quick start](#quick-start)
- [Add vocabulary during work](#add-vocabulary-during-work)
  - [Add vocabulary with Codex](#add-vocabulary-with-codex)
  - [Entry schema](#entry-schema)
- [Set up Anki once](#set-up-anki-once)
  - [Enable example speech and update an existing note type](#enable-example-speech-and-update-an-existing-note-type)
  - [Larger AnkiDroid answer buttons](#larger-ankidroid-answer-buttons)
  - [Import and update](#import-and-update)
  - [Audio](#audio)
- [Verify changes](#verify-changes)
- [Project layout](#project-layout)

## Quick start

Install Go 1.25 or later. Run these commands from the repository root:

```bash
go run scripts/validate/main.go
go run scripts/build/main.go
```

The first run downloads any uncached Go dependencies, including the YAML and language-tag parsers. Once the toolchain and dependencies are available locally, these commands can run offline.

The build validates all entries and writes `dist/vocabulary.tsv`, sorted by stable ID. Identical sources produce identical output. Failed validation leaves an existing export untouched.

`dist/` is ignored by Git. Rebuild it when you want to import a batch into Anki; never edit it manually.

## Add vocabulary during work

1. Search `entries/` for the word and check whether its meaning is already covered.
2. Create one lowercase kebab-case `.yml` file for a new meaning, or update the existing entry without changing its ID.
3. Validate, review the diff, and commit the source file.
4. Build and import a batch when convenient.

Choose one useful learning target per entry. Prefer dictionary headwords where appropriate, while keeping meaningful fixed expressions, grammatical forms, and natural plural forms. Put ordinary sentences and freely assembled descriptions in `example` or `notes`; create a phrase entry when the phrase itself is the learning target. For example, reuse `beef`, `rice`, and `sauce` for a meal description, but keep a useful expression such as `mashed potatoes` or the pattern `Would you like A or B?`. Do not generate every combination of existing words or split phrases solely by length.

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

Use separate IDs and descriptive filenames for distinct meanings, such as `commit-git` and `commit-general`. Changing an ID can create another Anki note. Renaming a source file while preserving its ID does not change its identity.

### Add vocabulary with Codex

Open Codex in the repository root and use the following prompt. Replace the placeholders with your vocabulary and language; add context if useful, and specify whether to generate audio.

```text
Add the following vocabulary to LexForge.

Words or expressions:
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

For each entry:
- Use a globally unique stable ID and lowercase kebab-case filename.
- Add a concise Chinese meaning, practical context, and a natural
  example in the vocabulary's language.
- Prefer software-engineering usage when it matches the supplied context.
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

### Entry schema

`language` is a BCP 47 tag, parsed with `golang.org/x/text/language` v0.38.0. Use `en`, `ja`, `ko`, `fr`, or a more specific tag such as `en-US` or `zh-Hant`. Underscores are rejected. Parser errors, including unknown subtags, fail validation. Case and deprecated aliases are canonicalized internally and in exports, without rewriting source spelling. Valid private-use tags and extensions are accepted according to the parser; they do not guarantee an available speech voice.

The exporter adds a lowercase `lang-<canonical-tag>` Anki tag and removes duplicates from the combined tag list. IDs remain globally unique across languages. Use descriptive Latin kebab-case filenames for Japanese and Korean entries too.

For Japanese, `word: 確認` and `pronunciation: かくにん` show 確認 and its kana reading on both sides. Use a Japanese example, such as 設定を確認してください。 Meanings remain concise Chinese. Romanized Japanese readings, furigana markup, and pitch-accent fields are not required.

Required string fields are `id`, `language`, `word`, `meaning`, `context`, and `example`.

Optional string fields are `pronunciation`, `part_of_speech`, `notes`, and `audio`. The optional `tags` field is a list of strings. Omit unused optional fields; empty strings, nulls, and empty lists fail validation.

- Use lowercase letters, digits, and single separating hyphens for IDs, filename stems, and tags.
- Keep one YAML mapping in each file directly under `entries/`. Nested directories, symlinks, aliases, merge keys, unknown fields, and duplicate keys are rejected. Dot-prefixed files are ignored.
- Quote text that YAML would interpret as a number or boolean, and text containing `: ` or other YAML syntax.
- Tabs and line breaks are supported in text; other control characters are rejected.
- Use concise Chinese meanings and practical context; examples use the entry’s language. Keep text as plain text, not HTML or Markdown formatting.
- Use verified IPA for English and kana readings for Japanese; omit uncertain pronunciation. The sample IPA comes from the project specification; other samples omit it.
- Use `audio: <id>.mp3` only after adding the corresponding file to `media/`.

Validation rejects missing or blank required fields, incorrect types, invalid names or tags, duplicate IDs or tags, and missing or empty referenced audio files. Errors go to stderr and return a nonzero exit status.

Repeated words within the same canonical language tag produce warnings for human review, including when their capitalization or whitespace differs. Distinct meanings remain valid. This check does not infer semantic equivalence or detect synonyms. Warnings do not block a build.

## Set up Anki once

Use a current Anki Desktop release. These instructions require Anki 2.1.58 or later for the CSV importer and its match-scope setting. The export contains twelve columns: eleven note fields and Anki tags. See the [Anki release notes](https://changes.ankiweb.net/changes/2.1.50-59.html#changes-in-2158) for the import changes.

1. In **Tools → Manage Note Types**, clone the **Basic** note type and name it `LexForge`.
2. In **Fields**, rename `Front` to `ID` and `Back` to `Language`. Add the remaining fields in this exact order:

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
   TTSLanguage
   ```

3. Keep `ID` first. Do not create a custom `Tags` field or map `ID` to Anki's internal GUID.
4. In **Cards**, keep one forward card and replace its front, back, and styling with the templates below.
5. Create a deck named `LexForge`, or choose another deck for this collection.

If you already have LexForge notes, update the existing note type to use this field order and the current templates. Follow [Update an existing note type](#enable-example-speech-and-update-an-existing-note-type). Deleting and recreating notes discards their review history.

Front template:

```html
<div class="lexforge-card">
<div class="word">{{Word}}</div>
<hr class="answer-divider">
<div class="pronunciation-row">
{{#Pronunciation}}<div class="pronunciation">{{Pronunciation}}</div>{{/Pronunciation}}
{{#Audio}}<div class="audio">{{Audio}}</div>{{/Audio}}
</div>
<script>
(function () {
  var card = document.querySelector(".lexforge-card");
  var example = card.querySelector(".example-speech .replay-button, .example-speech .replaybutton");
  if (example) {
    example.setAttribute("aria-label", "Read example");
  }
  var word = card.querySelector(".audio .replay-button, .audio .replaybutton");
  if (word && !word.hasAttribute("data-autoplayed")) {
    word.setAttribute("data-autoplayed", "true");
    setTimeout(function () {
      if (word.isConnected) word.click();
    }, 0);
  }
})();
</script>
</div>
```

Back template:

```html
<div class="lexforge-card">
<div class="word">{{Word}}</div>
<hr class="answer-divider">
<div class="pronunciation-row">
{{#Pronunciation}}<div class="pronunciation">{{Pronunciation}}</div>{{/Pronunciation}}
{{#Audio}}<div class="audio">{{Audio}}</div>{{/Audio}}
</div>
{{#PartOfSpeech}}<div class="part-of-speech">{{PartOfSpeech}}</div>{{/PartOfSpeech}}
<div class="meaning">{{Meaning}}</div>
<div class="context">{{Context}}</div>
<div class="example">{{Example}}
{{#Example}}{{#TTSLanguage}}<span class="example-speech">[anki:tts lang={{TTSLanguage}}]{{Example}}[/anki:tts]</span>{{/TTSLanguage}}{{/Example}}
</div>
{{#Notes}}<div class="notes">{{Notes}}</div>{{/Notes}}
<script>
(function () {
  var card = document.querySelector(".lexforge-card");
  var example = card.querySelector(".example-speech .replay-button, .example-speech .replaybutton");
  if (example) {
    example.setAttribute("aria-label", "Read example");
  }
  var word = card.querySelector(".audio .replay-button, .audio .replaybutton");
  if (word && !word.hasAttribute("data-autoplayed")) {
    word.setAttribute("data-autoplayed", "true");
    setTimeout(function () {
      if (word.isConnected) word.click();
    }, 0);
  }
})();
</script>
</div>
```

Styling:

```css
:root {
  color-scheme: light;
  --color-primary: #4d8ffb;
  --color-on-primary: #141414;
  --color-surface: #ffffff;
  --color-heading: #2d2d2d;
  --color-body: #626262;
  --color-muted: #828282;
  --color-divider: #f6f6f6;
  --color-field: #ebebeb;
  --color-focus: #d9d9d9;
  --color-fieldset: #c0c0c0;
  --color-code-bg: #eeeeee;
  --color-code: #e83e8c;
  --color-highlight: #fff9c0;
}

.card.nightMode {
  color-scheme: dark;
  --color-primary: #5089e8;
  --color-surface: #141414;
  --color-heading: #d6d6d6;
  --color-body: #a5a5a5;
  --color-muted: #878787;
  --color-divider: #1d1d1d;
  --color-field: #272727;
  --color-focus: #373737;
  --color-fieldset: #4e4e4e;
  --color-code-bg: #242424;
  --color-code: #d44386;
  --color-highlight: #413f2b;
}

.card,
.card.nightMode {
  background: var(--color-surface);
  color: var(--color-body);
  margin: 0;
  padding: 20px 16px 32px;
  font-family: sans-serif;
  font-size: 22px;
  font-size: clamp(22px, 5.5vw, 26px);
  line-height: 1.5;
  text-align: left;
  overflow-wrap: break-word;
  overflow-wrap: anywhere;
}

.word {
  color: var(--color-heading);
  font-size: 48px;
  font-size: clamp(48px, 12vw, 56px);
  font-weight: bold;
  line-height: 1.15;
}

.answer-divider {
  max-width: 42rem;
  margin: 24px 0 16px;
  border: 0;
  border-top: 1px solid var(--color-divider);
}

.pronunciation {
  font-size: 20px;
  font-size: clamp(20px, 5vw, 22px);
}

.pronunciation-row {
  display: flex;
  align-items: center;
}

.pronunciation-row .pronunciation {
  min-width: 0;
  flex: 0 1 auto;
}

.pronunciation-row .audio {
  flex: 0 0 auto;
  display: flex;
  align-items: center;
}

.pronunciation-row .audio .replay-button,
.pronunciation-row .audio .replaybutton,
.example-speech .replay-button,
.example-speech .replaybutton {
  box-sizing: border-box;
  display: inline-flex;
  align-items: center;
  justify-content: center;
  width: 44px;
  height: 44px;
  min-width: 44px;
  min-height: 44px;
  margin: 0;
  padding: 0;
  line-height: 1;
}

.pronunciation-row .audio .replay-button span,
.pronunciation-row .audio .replaybutton span,
.example-speech .replay-button span,
.example-speech .replaybutton span {
  display: flex;
  align-items: center;
  justify-content: center;
  margin: 0;
  padding: 0;
  line-height: 1;
}

.pronunciation-row .audio .replay-button svg,
.pronunciation-row .audio .replay-button img,
.pronunciation-row .audio .replaybutton svg,
.pronunciation-row .audio .replaybutton img,
.example-speech .replay-button svg,
.example-speech .replay-button img,
.example-speech .replaybutton svg,
.example-speech .replaybutton img {
  display: block;
  width: 24px;
  height: 24px;
  min-width: 24px;
  min-height: 24px;
  max-width: 24px;
  max-height: 24px;
}

.pronunciation-row .audio .replay-button:focus,
.pronunciation-row .audio .replaybutton:focus,
.example-speech .replay-button:focus,
.example-speech .replaybutton:focus {
  outline: 2px solid var(--color-heading);
  outline-offset: 2px;
}

/* Keep the focus outline as a fallback for older renderers. */
.pronunciation-row .audio .replay-button:focus:not(:focus-visible),
.pronunciation-row .audio .replaybutton:focus:not(:focus-visible),
.example-speech .replay-button:focus:not(:focus-visible),
.example-speech .replaybutton:focus:not(:focus-visible) {
  outline: none;
}

.pronunciation + .audio {
  margin-left: 8px;
}

.part-of-speech,
.notes {
  font-size: 19px;
  font-size: clamp(19px, 4.5vw, 20px);
}

/* Small metadata uses the body color in light mode for contrast. */
.nightMode .pronunciation,
.nightMode .part-of-speech,
.nightMode .notes {
  color: var(--color-muted);
}

.meaning {
  margin-top: 20px;
  color: var(--color-heading);
  font-size: 26px;
  font-size: clamp(26px, 7vw, 30px);
  font-weight: bold;
  line-height: 1.4;
}

.context,
.example,
.notes {
  box-sizing: border-box;
  max-width: 42rem;
  margin-top: 16px;
}

.example {
  padding-left: 14px;
  border-left: 3px solid var(--color-fieldset);
  font-size: 21px;
  font-size: clamp(21px, 5.25vw, 24px);
}

.example-speech {
  display: inline-flex;
  vertical-align: middle;
  margin-left: 0.4em;
}

@media (min-width: 700px) {
  .card,
  .card.nightMode {
    padding: 28px 32px 40px;
  }
}
```

The styling uses Anki's [night-mode class](https://docs.ankiweb.net/templates/styling.html#night-mode), also supported by AnkiDroid. Font sizes adapt to the viewport and stop growing at desktop widths.

The front shows the word, with pronunciation and audio when available. The back adds part of speech, meaning, context, the example, and notes. Japanese pronunciation uses kana.

Both sides render pronunciation and word audio explicitly, once per side. They share a row: long readings wrap while the native replay icon stays beside the reading. Missing pronunciation or audio fields do not leave an empty row.

Word and example replay controls use 24px icons inside 44px touch targets. On renderers supporting `:focus-visible`, ordinary taps do not leave an outline; keyboard focus remains visible. Older renderers retain the focus outline. To apply this styling-only update, replace **Styling** in the existing note type and sync; no TSV reimport or audio regeneration is needed.

### Enable example speech and update an existing note type

Update your existing note type in Anki; editing this README does not update installed templates. Keep existing notes and their `ID` values intact.

1. In **Tools → Manage Note Types → LexForge → Fields**, add `TTSLanguage` after `Audio` if it is missing. Confirm the field order shown in [Set up Anki once](#set-up-anki-once).
2. Replace the front, back, and styling with the templates in that section.
3. Rebuild the TSV and reimport with **update existing notes** and **note type** match scope. Verify all eleven fields and the tags column before importing.
4. Sync Anki Desktop and AnkiDroid.

In the LexForge deck's options, enable **Don't play audio automatically** and keep replay buttons visible in both clients. Use a separate deck preset if other decks should retain native autoplay.

The template script plays only the word recording automatically, once on each side. Tap the native replay icon beside the example to hear the full example; its accessible label is **Read example**. Leaving native autoplay enabled can play the example automatically and play the word twice. The app's general Replay Audio action may replay all audio on that side, including the example. Use the individual controls to choose one.

Example speech requires AnkiDroid 2.17 or later and a matching installed TTS voice. Anki Desktop on macOS uses system voices; AnkiDroid uses its configured Android speech engine. Install the desired language's voice in the device's text-to-speech settings; `{{tts-voices:}}` in a temporary template lists Anki's available voices. Missing voices produce Anki's native playback error. Japanese names and fictional readings in examples may be mispronounced. No example MP3s are generated or synced.

`TTSLanguage` is derived during export, not stored in YAML. Bare `en`, `ja`, `ko`, and `fr` map to `en_US`, `ja_JP`, `ko_KR`, and `fr_FR`. Ordinary language-region tags such as `en-GB` become `en_GB`. Other valid tags remain accepted, but produce an empty TTS locale and a build warning until an explicit mapping is added in `internal/vocabulary/export.go`. Cards with an empty locale show the example without a speech control. A locale mapping does not guarantee an installed voice.

After setup, check English and Japanese cards on Anki Desktop and each Android device:

- Confirm that word audio plays once on the front and once on the back.
- Confirm that examples stay silent until tapped and each control plays only its own content.
- Check rapid card transitions, repeated taps, missing word audio, and light and dark modes.

Verify playback on the actual devices; template checks do not establish device behavior. To roll back the playback behavior, remove the example speech span and both scripts, then restore native autoplay. The extra note field can remain.

Later styling-only edits do not require a rebuild or reimport. See Anki's [native TTS documentation](https://docs.ankiweb.net/templates/fields.html#text-to-speech-for-individual-fields) and [audio options](https://docs.ankiweb.net/deck-options.html#audio).

If multiple meanings make the prompt ambiguous, optionally add `{{Example}}` below `{{Word}}` on the front. Keep translations on the back.

### Larger AnkiDroid answer buttons

To enlarge **Again**, **Hard**, **Good**, and **Easy**, open AnkiDroid's **Settings → Accessibility** and increase **Answer buttons height** (called **Answer button size** in some versions). Try **150%**, then adjust toward **200%** if needed. Current [AnkiDroid accessibility preferences](https://github.com/ankidroid/Anki-Android/blob/main/AnkiDroid/src/main/res/xml/preferences_accessibility.xml) provide this control; older layouts place **Answer button size** under **Reviewing → Display**, as described in the [AnkiDroid manual](https://docs.ankidroid.org/#_display).

These are AnkiDroid controls outside the card content. Change their size in the app; the card template's CSS controls the content above them. No TSV rebuild or reimport is needed.

### Import and update

1. Build `dist/vocabulary.tsv`.
2. In Anki Desktop, import the TSV and select the `LexForge` note type and your deck.
3. Verify the preview: tab separator, HTML enabled, columns 1–11 mapped to the matching fields, and column 12 mapped to Anki tags. The export headers supply these defaults, but the note type must already exist.
4. Choose to **update** existing notes. Use **note type** as the match scope, rather than note type and deck. Keep these settings for later imports.
5. Import, inspect a card, and sync with AnkiWeb. Then sync AnkiDroid to review on mobile.

Anki matches the first field within the selected note type and preserves scheduling when updating notes in place. See the [official import documentation](https://docs.ankiweb.net/manual/importing/text-files) for settings and behavior.

Make vocabulary corrections in YAML. Reimporting can overwrite edits made only in Anki. Keep tags in YAML as well: importing the mapped tags column can replace tags added only in Anki. Removing a source entry excludes it from future exports; it does not delete the existing Anki note. Handle deletions explicitly in Anki.

The exporter escapes plain text for HTML display and converts line breaks to `<br>`. Technical examples such as `<div>` display literally. Tabs are preserved in the TSV using quoting. Audio markup is generated only from the `audio` field.

### Audio

Generate pronunciation audio locally on macOS with an installed language-matching voice, LAME CLI, and `mpg123`. Install the command-line tools yourself if needed:

```bash
brew install lame mpg123
```

The command checks these prerequisites and voices needed for new recordings before processing entries; it does not install anything or call a remote service. Run it from the repository root:

```bash
go run scripts/audio/main.go --id preserve-tech
go run scripts/audio/main.go --all
```

Select exactly one of `--id` or `--all`. The command validates the collection first. By default, it reads each selected entry's `word` and generates `media/<id>.mp3` using the selected voice and LAME at 96 kbps mono.

Before adding the `audio` field, the command verifies that the recording produces decoded audio samples. A successful decoder exit without samples is rejected. Other entry content and formatting are preserved. Omit unverified IPA even when audio is available.

Bare `en`, `ja`, and `ko` prefer Samantha, Kyoko, and Yuna respectively when installed. English defaults to 150 words per minute; other languages use the selected voice's native rate. Japanese and Korean native rates are starting points that need listening review, not a claim of pronunciation quality. Explicit regional tags require a matching region. Other languages use a unique matching installed voice; ambiguous or unavailable matches produce an actionable error.

Inspect installed voices with `say -v '?'`. For a single entry, choose `--voice "Voice Name"` or `--rate 160`. An explicit voice may select a different region of the same language, but cannot substitute another language. `--rate 0` uses the language default. These overrides require `--id`.

For a Japanese entry, add `--text-source pronunciation` to a single-entry command to speak its supplied kana reading instead of `word`. This accepts ordinary sentence punctuation and removes the grammar placeholders `～` and `〜` from speech input only; the YAML reading stays unchanged. It rejects missing readings, readings without kana, unsupported characters, and non-Japanese entries. It does not guarantee correct pitch accent. No voice selection is needed to validate or export vocabulary.

Existing referenced audio is skipped. A correctly named MP3 without an `audio` reference is decode-checked and reused. Replace a recording only when intended:

```bash
go run scripts/audio/main.go --id preserve-tech --overwrite
```

`--all --overwrite` replaces every recording. Without `--overwrite`, reruns preserve existing recordings. An empty or undecodable unreferenced MP3 causes an error instead of being silently replaced. Validation errors, including missing referenced files, must be resolved before generation.

Batch processing follows stable-ID order and stops at the first failure with a nonzero exit status and the affected entry ID. Completed entries remain valid. Without `--overwrite`, rerun the same command to resume while preserving completed recordings.

Avoid editing selected entries while the command runs. It checks for source edits before attaching audio, cleans up temporary files on handled failures, and reports generated, reused, and skipped counts. Filenames and settings are stable, but audio bytes can differ across macOS or encoder versions.

Audio generation does not build the export or perform Git operations. After a successful batch, run:

```bash
go run scripts/validate/main.go
go run scripts/build/main.go
```

Generation requires macOS; validation, export, and the resulting MP3s remain portable. Decoding checks confirm that files are readable, not that every technical term is pronounced correctly. Audition unfamiliar terms before importing them into Anki.

Alternatively, add a pronunciation MP3 at `media/<id>.mp3` and set the entry's `audio` field to that filename. The build emits `[sound:<id>.mp3]` without modifying the media file.

Before importing, copy referenced MP3 files directly into your Anki profile's `collection.media` folder, without subdirectories. Anki's **Tools → Check Media** command helps identify missing files. See [Anki file locations](https://docs.ankiweb.net/files.html) to find your profile. Sync media through AnkiWeb to AnkiDroid.

The validator checks the filename and that the file is a nonempty regular file; it does not decode MP3 data. The build does not synthesize audio or require a TTS service. Keep media in Git and reuse existing files; replace audio only when you intend to update it.

The five original sample pronunciations were generated locally with macOS's Samantha voice (US English) at 150 words per minute and encoded as 96 kbps mono MP3s using LAME through `lameenc` 1.9.0. The `vocabulary` recording and subsequent recordings use LAME CLI. Existing recordings do not need to be regenerated.

The [PR vocabulary import record](docs/PR_VOCABULARY_IMPORT.md) maps the reviewed source vocabulary to entries, grouped forms, and exclusions.

## Verify changes

```bash
go test ./...
go vet ./...
go run scripts/validate/main.go
go run scripts/build/main.go
git diff --check
```

Tests cover strict YAML validation, duplicate handling, deterministic ordering and output, escaping, audio references, and preservation of an existing export after validation fails.

Before using a new note type with your main collection, check it in a separate Anki profile: import twice without increasing the note count, review a card, update its YAML example, rebuild, and reimport. Confirm that the example changes while note identity and scheduling remain intact. Check audio playback and AnkiDroid sync separately when you add media.

## Project layout

```text
entries/               Canonical vocabulary, one meaning per YAML file
media/                 Optional pronunciation MP3 files
internal/vocabulary/   Shared validation and export logic, with tests
scripts/validate/      Validation command
scripts/build/         Build command
scripts/audio/         Local macOS pronunciation command
dist/                  Generated Anki TSV; ignored by Git
```

[AGENTS.md](AGENTS.md) defines coding-agent conventions.
