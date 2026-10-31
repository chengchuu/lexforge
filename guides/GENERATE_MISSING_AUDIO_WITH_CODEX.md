# Generate missing audio with Codex

Open Codex in the repository root and use the following prompt to scan all entries and fill missing pronunciation audio. It preserves existing recordings and uses the local macOS generator. See the [audio workflow](../README.md#audio) for tool and voice requirements.

```text
Scan all LexForge entries and generate missing pronunciation audio.

Read AGENTS.md, the audio workflow in README.md, and the current audio
implementation before making changes. Use the existing generator;
do not change the schema, CLI, exporter, templates, or dependencies.

Inspect Git status and preserve unrelated work. Record the original
entry contents and hashes of existing media so preservation can be
checked afterward. Use stable IDs rather than filename stems.

Inventory every entry and distinguish:
- An audio field referencing an existing nonempty regular file.
- No audio field, but an existing media/<id>.mp3 that may be reusable.
- No audio field and no media/<id>.mp3: a generation candidate.
- A declared audio reference whose file is missing or invalid.

Run vocabulary validation before generation. The current generator loads
and validates the entire collection, even with --id. A missing declared
recording therefore blocks generation; --overwrite does not bypass this.
If validation fails, report the affected IDs, paths, and errors. Do not
delete audio declarations to hide failures or change validation behavior.
Stop generation until the invalid source or missing reference is repaired.
Distinguish this repair from generating audio for entries without audio.

Check that macOS, say, lame, mpg123, and suitable installed voices are
available. Report missing prerequisites without automatically installing
tools or voices or switching to a remote service.

Process candidates individually in stable-ID order:
- Reuse existing unreferenced media/<id>.mp3 through the audio command.
  The command decodes the file before attaching audio metadata.
- For English, generate from word using:
  go run scripts/audio/main.go --id <stable-id>
- For Japanese, generate from the supplied kana pronunciation using:
  go run scripts/audio/main.go --id <stable-id> --text-source pronunciation
  Use the implemented punctuation and placeholder handling. Do not rewrite
  the YAML reading, invent kana, or substitute romaji. If the reading is
  missing or unsuitable, report it and leave that recording pending.
- For other languages, use word only when a matching installed voice and
  the current implementation support it. Never substitute another
  language's voice. Report ambiguous or unavailable voice selection.

Do not use --all for this mixed-language workflow: it defaults to word
and cannot select Japanese pronunciation input. Do not use --overwrite,
replace existing files, delete surplus media, or regenerate valid audio.
Report empty or undecodable existing recordings as repair candidates.
Referenced files are skipped by the generator without decoding, so do
not claim that skipping them verifies their contents or pronunciation.

Let the generator attach audio: <id>.mp3 after successful generation or
reuse. Preserve all other YAML fields, source wording, stable IDs,
learning metadata, relationships, and existing recording bytes.
Generate pronunciation audio only, not Example or Notes recordings.

Record failures per entry. Continue independent candidates when the
collection remains valid and the failure is entry-specific. Stop for
missing shared prerequisites or concurrent source edits. Preserve
completed work and report partial completion; do not claim the operation
is a collection-wide transaction.

After processing, run:
go run scripts/validate/main.go
go run scripts/build/main.go
git diff --check

Inspect new files separately because git diff --check excludes untracked
files. Confirm new and reused recordings decode to audio samples, all
new references match their IDs, and existing media hashes are unchanged.
Compare entry changes against the original contents: only the intended
audio metadata may change. Never edit dist/ manually.

Rescan the collection. Report total entries, already-referenced files,
reused recordings, generated recordings, and every remaining entry
without usable audio, with its reason. Separate file/reference checks,
decoding, human pronunciation review, and live Anki playback. Do not
claim checks that were not performed.

Report changed files and validation/build results. If no changes were
needed, say so. Do not stage, commit, push, copy files into an Anki
profile, or modify the user's Anki collection.
```

After reviewing the generated recordings, follow the [media import instructions](../README.md#audio) and [reimport the TSV](../README.md#import-and-update). Generating repository media does not install it in Anki.
