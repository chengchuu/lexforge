# LexForge

LexForge turns English vocabulary from everyday technical work into contextual Anki cards. Git owns the vocabulary; Anki owns review history and scheduling.

## Quick start

Install Go 1.25 or later. Run these commands from the repository root:

```bash
go run scripts/validate/main.go
go run scripts/build/main.go
```

The first run downloads the YAML parser. Subsequent runs can work offline once dependencies are cached. The build validates all entries and writes `dist/vocabulary.tsv`, sorted by stable ID. Identical sources produce identical output. Failed validation leaves an existing export untouched.

`dist/` is ignored by Git. Rebuild it when you want to import a batch into Anki; never edit it manually.

## Add vocabulary during work

1. Search `entries/` for the word and check whether its meaning is already covered.
2. Create one lowercase kebab-case `.yml` file for a new meaning, or update the existing entry without changing its ID.
3. Validate, review the diff, and commit the source file.
4. Build and import a batch when convenient.

For example, `entries/idiomatic.yml` can contain:

```yaml
id: idiomatic-tech
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

### Entry schema

Required string fields are `id`, `word`, `meaning`, `context`, and `example`.

Optional string fields are `pronunciation`, `part_of_speech`, `notes`, and `audio`. `tags` is an optional list of strings. Omit unused optional fields instead of adding empty strings, nulls, or empty lists.

- Use lowercase letters, digits, and single separating hyphens for IDs, filename stems, and tags.
- Keep one YAML mapping in each file directly under `entries/`. Nested directories, symlinks, aliases, merge keys, unknown fields, and duplicate keys are rejected. Dot-prefixed files are ignored.
- Quote text that YAML would interpret as a number or boolean, and text containing `: ` or other YAML syntax.
- Tabs and line breaks are supported in text; other control characters are rejected.
- Use concise Chinese meanings and practical context; examples normally use English. Keep text as plain text, not HTML or Markdown formatting.
- Add IPA only when verified or supplied reliably. The sample IPA comes from the project specification; other samples omit it.
- Use `audio: <id>.mp3` only after adding the corresponding file to `media/`.

Validation rejects missing or blank required fields, incorrect types, invalid names or tags, duplicate IDs or tags, and missing or empty referenced audio files. Errors go to stderr and return a nonzero exit status.

Repeated words produce warnings for human review, including when their capitalization or whitespace differs. Distinct meanings remain valid. This check does not infer semantic equivalence or detect synonyms. Warnings do not block a build.

## Set up Anki once

Use a current Anki Desktop release. These instructions require Anki 2.1.58 or later for the CSV importer and its match-scope setting. The export contains ten columns: nine note fields and Anki tags. See the [Anki release notes](https://changes.ankiweb.net/changes/2.1.50-59.html#changes-in-2158) for the import changes.

1. In **Tools → Manage Note Types**, clone the **Basic** note type and name it `LexForge`.
2. In **Fields**, rename `Front` to `ID` and `Back` to `English`. Add the remaining fields in this exact order:

   ```text
   ID
   English
   Pronunciation
   PartOfSpeech
   Meaning
   Context
   Example
   Notes
   Audio
   ```

3. Keep `ID` first. Do not create a custom `Tags` field or map `ID` to Anki's internal GUID.
4. In **Cards**, keep one forward card and replace its front, back, and styling with the templates below.
5. Create a deck named `LexForge`, or choose another deck for this collection.

Front template:

```html
<div class="english">{{English}}</div>
```

Back template:

```html
{{FrontSide}}
<hr id="answer">
{{#Pronunciation}}<div class="pronunciation">{{Pronunciation}}</div>{{/Pronunciation}}
{{#PartOfSpeech}}<div class="part-of-speech">{{PartOfSpeech}}</div>{{/PartOfSpeech}}
<div class="meaning">{{Meaning}}</div>
<div class="context">{{Context}}</div>
<div class="example">{{Example}}</div>
{{#Notes}}<div class="notes">{{Notes}}</div>{{/Notes}}
{{#Audio}}<div class="audio">{{Audio}}</div>{{/Audio}}
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

.english {
  color: var(--color-heading);
  font-size: 48px;
  font-size: clamp(48px, 12vw, 56px);
  font-weight: bold;
  line-height: 1.15;
}

#answer {
  max-width: 42rem;
  margin: 24px 0 16px;
  border: 0;
  border-top: 1px solid var(--color-divider);
}

.pronunciation {
  font-size: 20px;
  font-size: clamp(20px, 5vw, 22px);
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

.audio {
  margin-top: 20px;
}

@media (min-width: 700px) {
  .card,
  .card.nightMode {
    padding: 28px 32px 40px;
  }
}
```

The styling uses Anki's [night-mode class](https://docs.ankiweb.net/templates/styling.html#night-mode), also supported by AnkiDroid. Font sizes adapt to the viewport and stop growing at desktop widths.

For an existing LexForge note type, replace its **Back Template** and **Styling** in **Cards**, then sync Anki Desktop and AnkiDroid. These presentation changes apply to existing cards without rebuilding or reimporting the TSV.

If multiple meanings make the English-only prompt ambiguous, optionally add `{{Example}}` below `{{English}}` on the front. Keep translations on the back.

### Import and update

1. Build `dist/vocabulary.tsv`.
2. In Anki Desktop, import the TSV and select the `LexForge` note type and your deck.
3. Verify the preview: tab separator, HTML enabled, columns 1–9 mapped to the matching fields, and column 10 mapped to Anki tags. The export headers supply these defaults, but the note type must already exist.
4. Choose to **update** existing notes. Use **note type** as the match scope, rather than note type and deck. Keep these settings for later imports.
5. Import, inspect a card, and sync with AnkiWeb. Then sync AnkiDroid to review on mobile.

Anki matches the first field within the selected note type and preserves scheduling when updating notes in place. See the [official import documentation](https://docs.ankiweb.net/manual/importing/text-files) for settings and behavior.

Make vocabulary corrections in YAML. Reimporting can overwrite edits made only in Anki. Keep tags in YAML as well: importing the mapped tags column can replace tags added only in Anki. Removing a source entry excludes it from future exports; it does not delete the existing Anki note. Handle deletions explicitly in Anki.

The exporter escapes plain text for HTML display and converts line breaks to `<br>`. Technical examples such as `<div>` display literally. Tabs are preserved in the TSV using quoting. Audio markup is generated only from the `audio` field.

### Audio

Add a pronunciation MP3 at `media/<id>.mp3` and set the entry's `audio` field to that filename. The build emits `[sound:<id>.mp3]` without modifying the media file.

Before importing, copy referenced MP3 files directly into your Anki profile's `collection.media` folder, without subdirectories. Anki's **Tools → Check Media** command helps identify missing files. See [Anki file locations](https://docs.ankiweb.net/files.html) to find your profile. Sync media through AnkiWeb to AnkiDroid.

The validator checks the filename and that the file is a nonempty regular file; it does not decode MP3 data. The build does not synthesize audio or require a TTS service. Keep media in Git and reuse existing files; replace audio only when you intend to update it.

The five sample pronunciations were generated locally with macOS's Samantha voice (US English) at 150 words per minute and encoded as 96 kbps mono MP3s using LAME through `lameenc` 1.9.0. These tools are not project dependencies.

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
dist/                  Generated Anki TSV; ignored by Git
```

[GOAL.md](GOAL.md) defines project direction. [AGENTS.md](AGENTS.md) defines coding-agent conventions.
