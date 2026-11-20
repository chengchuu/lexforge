# Entry schema

`language` is a BCP 47 tag, parsed with `golang.org/x/text/language` v0.38.0. Use `en`, `ja`, `ko`, `fr`, or a more specific tag such as `en-US` or `zh-Hant`. Underscores are rejected. Parser errors, including unknown subtags, fail validation. Case and deprecated aliases are canonicalized internally and in exports, without rewriting source spelling. Valid private-use tags and extensions are accepted according to the parser; they do not guarantee an available speech voice.

The exporter adds a lowercase `lang-<canonical-tag>` Anki tag and removes duplicates from the combined tag list. IDs remain globally unique across languages. Use descriptive Latin kebab-case filenames for Japanese and Korean entries too.

For Japanese, `word: 確認` and `pronunciation: かくにん` show 確認 and its kana reading on both sides. Use a Japanese example, such as 設定を確認してください。 Meanings remain concise Chinese. Romanized Japanese readings, furigana markup, and pitch-accent fields are not required.

Required string fields are `id`, `language`, `word`, and `meaning`. Word can name a word, expression, term, or concept. Meaning gives a self-contained Chinese definition or short explanation of its principle; no question-and-answer fields are needed.

Optional string fields are `context`, `example`, `code_example`, `etymology`, `word_formation`, `mnemonic`, `pronunciation`, `part_of_speech`, `notes`, and `audio`. For the first six fields, omission, empty or whitespace-only strings, and YAML null mean absent. Other optional fields must be nonempty when supplied. The optional `tags`, `prerequisites`, and `related_entries` fields are nonempty lists of strings. `learning_priority` is an optional integer from 1 to 3 (default: 2).

Use `etymology` for verified historical origins, `word_formation` for supported construction patterns, and `mnemonic` for memory aids clearly distinguished from fact. Preserve uncertainty and put source titles and URLs in Notes. Notes can hold detailed explanations, steps, comparisons, and sources; it appears in full at the bottom of the back, without a length limit or collapse control. Keep the core explanation in Meaning. Context and Example can be omitted when they add no useful information.

Use `related_entries` for human-approved associations with existing stable IDs. Declaring a pair once displays both entries' current Word and Meaning on each other's backs. Only direct neighbors appear, sorted by stable ID; cross-language neighbors include their language tag. Reciprocal declarations deduplicate, and cycles are allowed. Missing targets, self-links, duplicate list members, and malformed IDs fail validation. Associations do not change prerequisites or learning order. If both endpoints declare a pair, remove both declarations to remove it, then rebuild and reimport the complete collection. Codex can propose candidates, but only approved pairs enter source data; explanatory mnemonic edits need content approval too.

For example, `related_entries: [disclosure-workflow]` on `closure-tech` supplies both display directions. These particular senses are not opposites. Related entries display as text, without cross-card navigation.

- Use lowercase letters, digits, and single separating hyphens for IDs, filename stems, and tags.
- Keep one YAML mapping in each file directly under `entries/`. Nested directories, symlinks, aliases, merge keys, unknown fields, and duplicate keys are rejected. Dot-prefixed files are ignored.
- Quote text that YAML would interpret as a number or boolean, and text containing `: ` or other YAML syntax.
- Tabs and line breaks are supported in text; other control characters are rejected.
- Use concise Chinese meanings and optional practical context; examples, when present, use the entry’s language. Put Chinese walkthroughs or symbolic traces in Notes rather than sending them to Example TTS. Keep text as plain text, not HTML or Markdown formatting.
- Put runnable snippets or configuration examples in `code_example`. Use a literal YAML block (`|-`) for multiline code and omit Markdown fences. The exporter escapes the snippet as text. The back template preserves its indentation and displays it outside Read Aloud. Use `example` for a prose sentence.
- Use verified IPA for English and kana readings for Japanese. Look up pronunciation for common English words and everyday phrases; follow the [pronunciation conventions](#pronunciation). Existing samples retain their original values.
- Use `audio: <id>.mp3` only after adding the corresponding file to `media/`.

Validation rejects missing or blank required fields, incorrect types, invalid names or tags, duplicate IDs or tags, and missing or empty referenced audio files. Errors go to stderr and return a nonzero exit status.

Learning metadata validation also rejects invalid priorities, malformed or repeated prerequisite IDs, missing entries, self-references, and dependency cycles. Cycle errors include the involved IDs.

Repeated words within the same canonical language tag produce warnings for human review, including when their capitalization or whitespace differs. Distinct meanings remain valid. This check does not infer semantic equivalence or detect synonyms. Warnings do not block a build.

## Pronunciation

`pronunciation` remains an optional plain-text string. For common English words and everyday phrases, seek reliable IPA rather than omitting it merely because the field is optional. Prefer American English; after a reasonable lookup, use verified British English if a reliable American pronunciation cannot be established. Record the reason for a British selection. A failed page request means the page was inaccessible, not that the pronunciation does not exist.

Unprefixed English IPA defaults to American English: write `/…/` without a `US ` prefix. Keep explicit prefixes for other accents, such as `UK /…/`. Match the current meaning, part of speech, and word form. Preserve the selected dictionary's transcription conventions and source a phrase as a whole; do not mix accents or splice individual word readings into an unsupported phrase transcription. Preserve existing pronunciation values unless a separate change is approved. Historical plans retain their originally approved wording; use this convention for new work.

Append the dictionary title, specific URL, accent, and access date to Notes, preserving existing text and avoiding duplicate source lines. If neither American nor British pronunciation can be reliably established, leave the field absent and report the specific limitation. Proper names, abbreviations, and fictional terms may need separate evidence; spelling and synthetic speech do not verify IPA.

English IPA labels describe the displayed transcription, not the recording's voice. A `UK /…/` value may accompany existing American audio. Keep the recording and its metadata unchanged unless replacement is separately authorized. These labels do not change `language`, `TTSLanguage`, or the audio generator's voice selection, and require no schema or template changes. Japanese continues to use kana without these English accent labels.
