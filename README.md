# LexForge

LexForge turns vocabulary and concepts encountered during work and learning into Anki cards. Each entry has a meaning, with optional context, examples, and memory aids. It supports English, Japanese, Korean, and other languages. Git tracks learning content; Anki manages review history and scheduling.

- [Quick start](#quick-start)
- [Add vocabulary during work](#add-vocabulary-during-work)
  - [Add vocabulary with Codex](#add-vocabulary-with-codex)
  - [Entry schema](#entry-schema)
- [Learning order](#learning-order)
- [Set up Anki once](#set-up-anki-once)
  - [Example speech](#example-speech)
  - [Larger AnkiDroid answer buttons](#larger-ankidroid-answer-buttons)
  - [Import and update](#import-and-update)
  - [Audio](#audio)
- [Verify changes](#verify-changes)
- [Project layout](#project-layout)

## Quick start

Run these commands from the repository root:

```bash
go run scripts/validate/main.go
go run scripts/build/main.go
```

The first run downloads uncached Go dependencies. With the toolchain and dependencies installed, both commands work offline.

The build validates all entries and writes a deterministic `dist/vocabulary.tsv` in [learning order](#learning-order). Failed validation preserves the existing export.

`dist/` is Git-ignored. Rebuild before importing a batch into Anki; never edit the TSV manually.

## Add vocabulary during work

Use the [Add vocabulary during work guide](guides/ADD_VOCABULARY_DURING_WORK.md) for the capture workflow, learning-target selection, a YAML example, and validation and export steps.

### Add vocabulary with Codex

Use the [Add vocabulary with Codex guide](guides/ADD_VOCABULARY_WITH_CODEX.md) for the reusable prompt. It covers entry reuse, natural examples, learning priorities, optional audio, and validation.

For existing entries, use the [Generate missing audio with Codex guide](guides/GENERATE_MISSING_AUDIO_WITH_CODEX.md) to scan the collection, reuse available recordings, and generate missing pronunciation audio.

### Entry schema

Use the [Entry schema guide](guides/ENTRY_SCHEMA.md) for field definitions and validation rules. It covers required and optional fields, language tags, memory aids, related entries, and source formatting.

## Learning order

Use the [Learning order guide](guides/LEARNING_ORDER.md) to define introduction priorities and prerequisites and apply the generated order in Anki. It covers StudyOrder, new-card repositioning, and review-history preservation.

## Set up Anki once

Use a current Anki Desktop release. These instructions require Anki 2.1.58 or later for the CSV importer and its match-scope setting. The export contains seventeen columns: sixteen note fields and Anki tags. See the [Anki release notes](https://changes.ankiweb.net/changes/2.1.50-59.html#changes-in-2158) for the import changes.

If you already have LexForge notes, use the [Anki migration guide](guides/ANKI_MIGRATION_GUIDE.md) to update the existing note type while preserving review history.

1. In **Tools → Manage Note Types**, clone the **Basic** note type and name it `LexForge`.
2. In **Fields**, rename `Front` to `ID` and `Back` to `Language`. Add the remaining fields in the recommended order below, which matches the TSV export:

   | No. | Field         |
   |:----|:--------------|
   | 1   | ID            |
   | 2   | Language      |
   | 3   | Word          |
   | 4   | Pronunciation |
   | 5   | PartOfSpeech  |
   | 6   | Meaning       |
   | 7   | Context       |
   | 8   | Example       |
   | 9   | Notes         |
   | 10  | Audio         |
   | 11  | TTSLanguage   |
   | 12  | StudyOrder    |
   | 13  | Etymology     |
   | 14  | WordFormation |
   | 15  | Mnemonic      |
   | 16  | RelatedWords  |

3. Keep `ID` first. Select `StudyOrder` as the browser sort field. Do not create a custom `Tags` field or map `ID` to Anki's internal GUID.
4. In **Cards**, keep one forward card and replace its front, back, and styling with the templates below.
5. Create a deck named `LexForge`, or choose another deck for this collection.

The recommended field order keeps setup and import checks consistent. You can reorder fields after `ID` for editing convenience, provided each TSV column maps to the correct destination field when importing. LexForge always exports columns in the order shown above. The front and back templates determine card display order independently of the field list.

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
{{#Context}}<div class="context">{{Context}}</div>{{/Context}}
{{#Example}}<div class="example">{{Example}}
{{#TTSLanguage}}<span class="example-speech">[anki:tts lang={{TTSLanguage}}]{{Example}}[/anki:tts]</span>{{/TTSLanguage}}
</div>{{/Example}}
{{#Etymology}}<div class="memory-aid"><strong>词源</strong><br>{{Etymology}}</div>{{/Etymology}}
{{#WordFormation}}<div class="memory-aid"><strong>构词</strong><br>{{WordFormation}}</div>{{/WordFormation}}
{{#Mnemonic}}<div class="memory-aid"><strong>助记</strong><br>{{Mnemonic}}</div>{{/Mnemonic}}
{{#RelatedWords}}<div class="memory-aid"><strong>相关词条</strong><br>{{RelatedWords}}</div>{{/RelatedWords}}
{{#Notes}}<div class="notes">{{Notes}}</div>{{/Notes}}
{{#Tags}}<div class="tags">{{Tags}}</div>{{/Tags}}
<script>
(function () {
  var card = document.querySelector(".lexforge-card");
  var tags = card.querySelector(".tags");
  if (tags && !tags.hasAttribute("data-tags-formatted")) {
    tags.textContent = tags.textContent.trim().split(/\s+/).join(", ");
    tags.setAttribute("data-tags-formatted", "true");
  }
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
.notes,
.tags {
  font-size: 19px;
  font-size: clamp(19px, 4.5vw, 20px);
}

/* Small metadata uses the body color in light mode for contrast. */
.nightMode .pronunciation,
.nightMode .part-of-speech,
.nightMode .notes,
.nightMode .tags {
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
.memory-aid,
.notes,
.tags {
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

The front shows the word, with pronunciation and audio when available. The back adds part of speech, meaning, optional context and example, memory aids, related entries, notes, and all note tags at the bottom. Absent sections leave no labels, spacing, or replay controls. Meaning and Notes are not read by example TTS. Japanese pronunciation uses kana.

Tags use Anki's native `{{Tags}}` special field, including derived language tags such as `lang-en`; do not add a custom field. The back script separates displayed tags with a comma and a space, leaving actual tags unchanged. To apply the separator update, replace **Back Template** in your installed note type and sync. If your cards do not yet display tags, also copy **Styling**. No TSV reimport or audio regeneration is needed.

Both sides render pronunciation and word audio explicitly, once per side. They share a row: long readings wrap while the native replay icon stays beside the reading. Missing pronunciation or audio fields do not leave an empty row.

Word and example replay controls use 24px icons inside 44px touch targets. On renderers supporting `:focus-visible`, ordinary taps do not leave an outline; keyboard focus remains visible. Older renderers retain the focus outline. To apply this styling-only update, replace **Styling** in the existing note type and sync; no TSV reimport or audio regeneration is needed.

### Example speech

For an existing note type, first follow the [Anki migration guide](guides/ANKI_MIGRATION_GUIDE.md).

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
3. Verify the preview: tab separator, HTML enabled, columns 1–16 mapped to the matching fields by name, and column 17 mapped to Anki tags. Check the mapping even if you use the recommended field order. The export headers supply defaults, but the note type must already exist.
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

The [PR vocabulary import record](guides/PR_VOCABULARY_IMPORT.md) maps the reviewed source vocabulary to entries, grouped forms, and exclusions.

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
