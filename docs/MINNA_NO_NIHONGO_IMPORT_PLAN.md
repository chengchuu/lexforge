# Minna no Nihongo Vocabulary Import Plan

## Status and source

Implemented on 2026-10-29 after user approval. All 49 items are accounted for: 44 new entries and recordings, plus five reused entries with their recordings preserved. The collection contained 571 entries (346 English and 225 Japanese) immediately after this batch, before the later English cleanup. See the [import record](JP_VOCABULARY_IMPORT.md#2026-10-29-minna-no-nihongo-study-notes) for mappings, passed checks, and preservation results.

Listening verification remains outstanding, especially for names, age readings, suffixes, and incomplete grammar patterns. No Anki import or synchronization was performed. The sections below retain the approved implementation scope.

The user recorded these items while studying 《大家的日语》 (みんなの日本語). The supplied AnkiDroid screenshot, `Screenshot_20261029_012025_AnkiDroid.png`, shows 49 cards. Treat these as the user's study notes, not as a transcription independently checked against a particular textbook edition or lesson.

The user selected normalized Japanese forms with original spellings retained, requested audio, and confirmed these corrections:

- `いちさい` → `一歳`, reading `いっさい`.
- `おいくし` → `おいくつ`.
- `きゅう` → `九`, reading `きゅう`.

## Intended result

Account for all 49 screenshot items: expect 44 new entries and five reused entries. Recheck current repository state before implementation; preserve existing stable IDs and avoid semantic duplicates. Do not perform another vocabulary-selection pass or add words merely because they occur in examples.

Reuse these existing entries:

- `yoroshiku-onegai-shimasu`: retain its primary form and add the screenshot's optional variants to notes.
- `donata`, `hajimemashite`, `byouin`, and `watashi`: retain their content and IDs; record the screenshot mapping in the import record.

All five currently have audio. Preserve those recordings. Do not regenerate them to read alternate spellings or longer variants.

## Approved item mapping

Remove display line wrapping from the source transcription. Preserve source spacing and optional brackets in the import record, even when the primary entry uses normalized spelling. Meanings below identify the intended learning sense; existing matching entries retain their current valid meanings.

| Screenshot text                  | Primary word         | Pronunciation          | Chinese meaning      | Action |
|:---------------------------------|:---------------------|:-----------------------|:---------------------|:-------|
| [どうぞ]よろしく[おねがいします] | よろしくお願いします | よろしくおねがいします | 请多关照             | Reuse  |
| ～からきました                   | ～から来ました       | ～からきました         | 来自……               | New    |
| ～じん                           | ～人                 | ～じん                 | ……人；国籍后缀       | New    |
| あなた                           | あなた               | あなた                 | 你                   | New    |
| あのかた                         | あの方               | あのかた               | 那位；礼貌说法       | New    |
| あのひと                         | あの人               | あのひと               | 那个人               | New    |
| いいえ                           | いいえ               | いいえ                 | 不；不是             | New    |
| いしゃ                           | 医者                 | いしゃ                 | 医生                 | New    |
| いちさい                         | 一歳                 | いっさい               | 一岁                 | New    |
| お なまえ は                     | お名前は             | おなまえは             | 请问您的名字是？     | New    |
| おいくし                         | おいくつ             | おいくつ               | 几岁；多少           | New    |
| おはようございます               | おはようございます   | おはようございます     | 早上好               | New    |
| かいしゃいん                     | 会社員               | かいしゃいん           | 公司职员             | New    |
| かんこく                         | 韓国                 | かんこく               | 韩国                 | New    |
| がくせい                         | 学生                 | がくせい               | 学生                 | New    |
| きゅう                           | 九                   | きゅう                 | 九                   | New    |
| きょうし                         | 教師                 | きょうし               | 教师                 | New    |
| けんきゅうしゃ                   | 研究者               | けんきゅうしゃ         | 研究人员             | New    |
| こうべ                           | 神戸                 | こうべ                 | 神户                 | New    |
| こうべ びょういん                | こうべ病院           | こうべびょういん       | 医院名称；こうべ病院 | New    |
| こちらは～さんです               | こちらは～さんです   | こちらは～さんです     | 这位是……             | New    |
| さくらだいがく                   | さくら大学           | さくらだいがく         | 大学名称；さくら大学 | New    |
| さん                             | ～さん               | ～さん                 | 接在人名后的礼貌称呼 | New    |
| しつれい ですが                  | 失礼ですが           | しつれいですが         | 冒昧问一下           | New    |
| しゃいん                         | 社員                 | しゃいん               | 公司成员；职员       | New    |
| せんせい                         | 先生                 | せんせい               | 老师等称呼           | New    |
| だいがく                         | 大学                 | だいがく               | 大学                 | New    |
| だれ                             | 誰                   | だれ                   | 谁                   | New    |
| ちゅうごく                       | 中国                 | ちゅうごく             | 中国                 | New    |
| ちゅうごくじん                   | 中国人               | ちゅうごくじん         | 中国人               | New    |
| どなた                           | どなた               | どなた                 | 哪位；礼貌说法       | Reuse  |
| なんさい                         | 何歳                 | なんさい               | 几岁                 | New    |
| にほん                           | 日本                 | にほん                 | 日本                 | New    |
| はい                             | はい                 | はい                   | 是；好的             | New    |
| はじめまして                     | はじめまして         | はじめまして           | 初次见面             | Reuse  |
| はっさい                         | 八歳                 | はっさい               | 八岁                 | New    |
| びょういん                       | 病院                 | びょういん             | 医院                 | Reuse  |
| ふじ だいがく                    | ふじ大学             | ふじだいがく           | 大学名称；ふじ大学   | New    |
| わたし                           | 私                   | わたし                 | 我                   | Reuse  |
| アメリカ                         | アメリカ             | アメリカ               | 美国                 | New    |
| アメリカからきました             | アメリカから来ました | アメリカからきました   | 我来自美国           | New    |
| イギリス                         | イギリス             | イギリス               | 英国                 | New    |
| インドネシア                     | インドネシア         | インドネシア           | 印度尼西亚           | New    |
| エー・ケー・シー                 | エー・ケー・シー     | エー・ケー・シー       | A・K・C 的读音；名称 | New    |
| タイ                             | タイ                 | タイ                   | 泰国                 | New    |
| タオル                           | タオル               | タオル                 | 毛巾                 | New    |
| ドイツ                           | ドイツ               | ドイツ                 | 德国                 | New    |
| パワー でんき                    | パワー電気           | パワーでんき           | 名称；パワー電気     | New    |
| ブラジル エアー                  | ブラジルエアー       | ブラジルエアー         | 名称；ブラジルエアー | New    |

## Entry content and distinctions

- Read `AGENTS.md` and the current `README.md`. Use the existing schema and field order, `language: ja`, supplied kana readings, concise Chinese meanings and context, and natural Japanese examples.
- Choose globally unique ASCII Hepburn-style lowercase kebab-case IDs using the existing Japanese conventions. Match filenames to IDs; add sense suffixes where needed to avoid collisions.
- Preserve original screenshot forms through the mapping in the import record. Also explain the two corrected spellings in the affected entries' notes without teaching them as valid readings.
- Keep `教師` and `先生` separate: occupation versus a title or form of address. Keep `会社員` and `社員` separate and explain how `社員` relates to membership in a company.
- Keep base vocabulary, suffixes, sentence patterns, and complete expressions distinct. The origin pattern `～から来ました` differs from the existing causal `～から` entry.
- Interpret screenshot `さん` as the name suffix `～さん`, not the numeral three. Treat `～人` as a nationality suffix. Preserve `一歳` and `八歳` as independent age expressions with their supplied readings.
- Explain that brackets in `[どうぞ]よろしく[おねがいします]` indicate optional wording. Keep `よろしくお願いします` as the primary entry and include `どうぞよろしくお願いします` and the shorter `よろしく` in notes, with register distinctions.
- Preserve institution and organization names conservatively: `こうべ病院`, `さくら大学`, `ふじ大学`, `エー・ケー・シー`, `パワー電気`, and `ブラジルエアー`. Do not infer unconfirmed kanji, real-world identities, acronym expansions, or factual details. Keep `タオル` in its ordinary towel sense; do not infer a person's name from the surrounding list.
- Reuse existing grammatical and semantic tags such as `noun`, `pronoun`, `phrase`, `suffix`, `counter`, `proper-noun`, `grammar`, `daily`, `polite`, and `travel` where appropriate. Do not add every tag to every entry or manually add `lang-ja`.

## Audio

The latest read-only check found Kyoko, macOS `say`, LAME CLI, and `mpg123` available. Recheck before generation; do not install missing prerequisites automatically.

- Generate audio only for the 44 new entries, using the existing single-ID workflow with `--text-source pronunciation`. Do not use `--all` or `--overwrite`.
- Use Kyoko for Japanese, its existing language-default rate, and the existing 96 kbps mono MP3 encoding. Derive filenames from stable IDs and let the existing command attach audio metadata after successful decoding.
- Speak the kana pronunciation rather than asking TTS to infer readings from kanji or proper names.
- The existing command removes `～` and `〜` from speech input without changing YAML. For example, `～からきました` speaks `からきました`. Pattern recordings remain fragments; do not silently fill placeholders with names or replace them with example sentences. In particular, audition the incomplete pattern `こちらは～さんです` and report any unnatural phrasing.
- Preserve all existing recordings, including the five reused entries. Reuse valid unreferenced audio only through the existing command's checks.
- Decode-check every new MP3. Audition representative age readings, names, suffixes, and sentence patterns where listening is available. Successful decoding does not verify accent or natural intonation; explicitly report unavailable listening verification or pronunciation concerns.

## Import record and verification

Append a separate dated batch section to `docs/JP_VOCABULARY_IMPORT.md`, preserving the historical import record. Include the source screenshot filename, all 49 source-to-ID mappings, corrections, reuse decisions, new entry count, audio results, warnings, and any unresolved limitations. Do not invent a textbook chapter or edition.

1. Capture existing entry contents and media hashes before editing. Limit existing-entry changes to the necessary variant notes for `yoroshiku-onegai-shimasu`.
2. Verify all 49 items are represented once as new or reused learning items. Confirm globally unique IDs, correct Japanese readings, Chinese meanings, natural examples, and no unapproved additions.
3. Run `go test ./...`, `go vet ./...`, `go run scripts/validate/main.go`, `go run scripts/build/main.go`, and `git diff --check`. Check untracked files for formatting as well.
4. Review warnings without removing legitimate distinctions. Verify original recordings remain byte-identical and every new audio reference resolves to a successfully decoded MP3.
5. Build twice and compare bytes. Use the current twelve-column layout: eleven note fields, including derived `TTSLanguage`, followed by Anki tags. Verify Japanese rows carry `ja`, `ja_JP`, derived `lang-ja`, the intended readings, and the correct audio references.
6. Inspect the final diff and report actual new/reused counts, audio generated/reused/skipped counts, validation results, and verification limits. If the collection changes before implementation, report actual totals rather than forcing an outdated baseline.

## Boundaries

Modify only the scoped vocabulary, import record, new audio, and generated export. Do not change the schema, tooling, dependencies, card templates, or styling. Do not edit generated TSV manually.

Do not import into Anki, delete notes, change review state, synchronize AnkiWeb, or modify AnkiDroid. Do not stage, commit, or push. The approved repository import and audio generation are complete; these boundaries still apply to any follow-up work.
