# Vocabulary Memory Aids Plan

## Status and authorization

Planned; not implemented. This document records the design discussed on 2026-10-30.

The user approved all 21 relationships in the [initial batch](#approved-relationship-batch) on 2026-10-30: E01–E13 and J01–J08. This is content approval for those pairs and authorization to record them here, not approval to implement the feature or edit vocabulary yet.

The user also approved recording the concept-learning design and all five [planning cases](#approved-planning-cases) below. This extends the planned use of existing fields; it does not authorize implementation, entry creation, or additional relationships.

Creating this plan does not authorize schema, vocabulary, exporter, template, or audio changes. Implementation requires explicit approval. Approval to implement the feature does not automatically approve candidate vocabulary relationships or associated mnemonic text; those require content review as described below.

## Goal

Support optional etymology, word-formation explanations, mnemonics, and related vocabulary on card backs. Let Word name a word, expression, term, or concept, and let Meaning explain its core meaning, reasoning, or principle. Make both Context and Example optional when they add no useful information. Keep each card focused and omit empty sections completely.

Use the existing fields for descriptive learning; add no Question, Answer, or Explanation fields. Meaning supports quick review, while Notes provides optional depth. Anki continues to own scheduling and review history.

The relationship mechanism applies to the entire multilingual collection. `closure-tech` and `disclosure-workflow` are one pair in the approved initial batch, not a special case in the implementation.

## Current implementation

- [Entry parsing](../internal/vocabulary/entries.go) requires nonempty `id`, `language`, `word`, `meaning`, `context`, and `example` strings. Unknown fields and empty optional values are rejected.
- [TSV export](../internal/vocabulary/export.go) produces twelve note fields and tags in column 13. `TTSLanguage` and `StudyOrder` are derived fields.
- The [README templates](../README.md#set-up-anki-once) render Context and the Example container unconditionally, and Notes conditionally. Example speech already checks Example and TTSLanguage. The source schema has no dedicated memory-aid or association fields.
- [Learning order](LEARNING_ORDER.md) uses directed prerequisites. Associations must remain separate from this ordering graph.

## Source schema

Add these optional fields:

| Source field      | Type               | Purpose                                                     |
| ----------------- | ------------------ | ----------------------------------------------------------- |
| `etymology`       | String             | Verified historical origin, with uncertainty preserved.     |
| `word_formation`  | String             | Supported morphological analysis or construction patterns.  |
| `mnemonic`        | String             | A concise memory aid, clearly distinguished from fact.      |
| `related_entries` | List of stable IDs | Human-approved direct associations with existing entries.   |

Change `context` and `example` from required to optional. Required fields become `id`, `language`, `word`, and `meaning`.

For `context`, `example`, `etymology`, `word_formation`, and `mnemonic`, treat omission, an empty string, a whitespace-only string, and a YAML null value as absent. Export these as empty strings without rewriting source files. Preserve nonempty text as supplied. Continue rejecting other incorrect types, aliases, and unsupported control characters. This relaxation applies only to these five fields; existing optional-field rules otherwise remain unchanged, including those for `notes`.

Omit unused `related_entries`. If present, require a nonempty list of lowercase kebab-case string IDs, following the existing reference-list conventions. Reject nulls, empty lists, malformed IDs, repeated IDs within one list, self-references, and missing entries.

Keep stable IDs globally unique and unchanged. Add no source fields for derived Anki output. Preserve existing Context and Example values; optionality does not authorize a collection-wide cleanup.

## Descriptive learning and field responsibilities

Use these existing fields for both vocabulary and concept learning:

| Field   | Required | Responsibility                                                          |
|:--------|:---------|:------------------------------------------------------------------------|
| Word    | Yes      | A word, expression, term, or concept name.                              |
| Meaning | Yes      | A self-contained core definition or concise explanation of a principle. |
| Context | No       | Useful background, applications, or conditions.                         |
| Example | No       | A natural example in the entry's language, when useful.                 |
| Notes   | No       | Detailed reasoning, steps, limitations, comparisons, and sources.       |

The front continues to show the name with optional pronunciation and word audio. The back shows the explanation directly, without a question-and-answer format. Do not require an artificial sentence merely to populate Example.

Keep Meaning independently understandable. Notes can be long and remains fully visible near the bottom of the back; readers can skip it during quick review or study it in depth. Do not impose a length limit, truncate useful material, or add collapse controls. Organize longer notes into paragraphs or plain-text numbered steps, and remove repetition only when it adds no value.

Keep the current plain-text storage and HTML escaping. Line breaks can separate explanations, but Markdown tables, rendered lists, syntax highlighting, and rich-text parsing are not part of this plan. Tables below describe field assignments; they are not proposed content to paste into a Notes field.

Keep `language` tied to the Word and any Example. An English concept name can have Chinese Meaning, Context, and Notes. Put Chinese walkthroughs and symbolic traces in Notes rather than sending them through English Example TTS. Example speech remains limited to a present Example with a supported TTSLanguage; do not automatically speak Meaning or Notes. Pronunciation, part of speech, and word audio remain optional.

## Approved planning cases

The user approved these five cases for inclusion in the plan. They illustrate the agreed content structure and are not an entry-import batch. When entry work is authorized, search for matching existing entries, preserve stable IDs, and reuse them where appropriate. In particular, inspect `closure-tech` before creating another closure entry. Do not infer new associations or prerequisites from these cases.

### 1. Closure: vocabulary with a conceptual explanation

- **Word:** closure
- **Meaning:** 闭包：函数及其所引用的词法环境的组合，使函数能够访问定义时所在的外层作用域。
- **Context:** 回调函数、封装内部状态等场景。
- **Example:** A closure keeps access to variables in its outer scope.
- **Notes direction:** Explain how an inner function can retain access to variables after the outer function returns. Distinguish references to bindings from frozen value snapshots, and explain that returning a function is not a requirement for a closure.

This case retains an English example while allowing deeper Chinese explanation in Notes. Source: [MDN Closures](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Guide/Closures).

### 2. Temporal dead zone: a language mechanism without an example

- **Word:** temporal dead zone
- **Meaning:** 暂时性死区：从进入作用域到相关绑定完成初始化之前的阶段；此时访问该绑定会抛出 ReferenceError。
- **Context:** 理解 JavaScript 中 let、const 和 class 的初始化行为。
- **Example:** Omitted.
- **Notes direction:** Distinguish an uninitialized binding from a binding whose value is undefined. Explain how an inner binding shadows an outer variable and why its own initializer can encounter the temporal dead zone.

This case demonstrates that Example can be absent without leaving an empty example container or replay control. Source: [MDN Temporal dead zone](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Statements/let#temporal_dead_zone_tdz).

### 3. Bubble sort: a short principle with a detailed walkthrough

- **Word:** bubble sort
- **Meaning:** 冒泡排序：反复比较相邻元素，顺序不符时交换；按升序排列时，每轮将未排序部分的最大元素移到末尾。
- **Context:** Omitted.
- **Example:** Omitted.
- **Notes walkthrough:**

  > 以升序排列 [3, 2, 1]：
  >
  > 第一轮：比较 3 和 2，交换为 [2, 3, 1]；再比较 3 和 1，交换为 [2, 1, 3]。
  >
  > 第二轮：比较 2 和 1，交换为 [1, 2, 3]。
  >
  > 每轮结束后，末尾已确定位置的元素不再参与后续比较。

Further Notes can cover stopping after a pass without swaps, complexity, and the conditions for stability. Keep the Chinese walkthrough in Notes rather than Example so it is not read by an English TTS voice. Source: [NIST Bubble sort](https://xlinux.nist.gov/dads/HTML/bubblesort.html).

### 4. Riboflavin: an English name for everyday knowledge

- **Word:** riboflavin
- **Meaning:** 核黄素，又称维生素 B2；参与能量代谢，支持细胞的正常生长、发育和功能。
- **Context:** Omitted.
- **Example:** Omitted.
- **Notes direction:** Describe common food sources, such as eggs, milk, and mushrooms, and retain the authoritative source for further reading.

This case uses the English name as the learning anchor without inventing an example sentence. Source: [NIH Riboflavin](https://ods.od.nih.gov/factsheets/Riboflavin-Consumer/).

### 5. Vitamin B6: a separate target within a related subject

- **Word:** vitamin B6
- **Meaning:** 维生素 B6：参与多种代谢相关的酶反应，并参与脑发育和免疫功能。
- **Context:** Omitted.
- **Example:** Omitted.
- **Notes direction:** Expand on its functions and food sources. A comparison with B2 can explain that both participate in metabolism while their roles are not interchangeable. Preserve the source.

Keep this as a separate learning target from riboflavin. A B2/B6 association remains a candidate, not an approved addition to the 21-pair batch. These cases concern nutritional knowledge, not personal supplement recommendations. Source: [NIH Vitamin B6](https://ods.od.nih.gov/factsheets/VitaminB6-Consumer/).

## Content policy

- Use concise Chinese explanations with relevant original-language forms.
- Distinguish historical derivation, modern word formation, and a learner's memory association. Do not present a mnemonic as historical evidence.
- Verify etymological claims against reliable sources. Preserve reconstruction markers and uncertainty. Record source titles and URLs in existing `notes` without replacing existing provenance.
- Omit explanations that cannot be established reliably or do not help the learning target. Do not force a root decomposition onto every word or language.
- Preserve the entry's actual sense. For example, the existing `escape-parity-tech` concerns odd/even escape counts, not feature equivalence.
- Keep Meaning and memory-aid summaries concise; allow detailed Notes for optional deeper reading. Do not automatically create cards for every root, cognate, or related expression.
- Keep example-vocabulary reuse optional and preserve the existing priority of accuracy, naturalness, and practical context.

## General related-entry model

Associations identify concrete entries by stable ID, including their particular meanings. They can support verified word-family relationships, morphological relationships, confusing spellings, synonym/antonym comparisons, usage contrasts, and useful links between words and expressions. A relationship does not itself assert any particular linguistic category; explanations belong in each entry's `mnemonic`.

Export behavior:

1. Treat a declared association between A and B as an undirected relationship for display. Declaring it at either endpoint makes both card backs show the other entry.
2. Generate only direct neighbors. A–B and B–C must not imply A–C.
3. Deduplicate display when both endpoints declare the same pair. Repeated IDs within an individual source list remain validation errors.
4. Allow cycles, unlike the prerequisite graph. Do not change `prerequisites`, priorities, ranks, or Anki scheduling because of an association.
5. Derive each neighbor's current Word and Meaning during the full build; sort neighbors by stable ID for deterministic output. Do not duplicate these values in source entries or expose IDs as the primary learning content.
6. Render neighbors as plain text lines in the derived `RelatedWords` field, using the normal HTML-escaping path. A word such as `<tag>` or a literal media marker must remain text. Cross-language associations are permitted when explicitly approved; include the neighbor's language tag when it differs from the current entry's language.
7. Export an empty value when there are no neighbors. Removing a relationship or changing a neighbor's content requires a complete rebuild and reimport to refresh both sides. If both endpoints declare a pair, remove both declarations to remove the association.

The first version provides inline comparison text. Cross-card navigation, automatic similarity matching, relationship-type schemas, and additional pronunciation or audio controls for related words are outside scope.

## Human confirmation workflow

Codex may inspect existing entries and propose candidates, but must not write unapproved relationships into source data.

Present candidates in a review table containing both stable IDs, words, current meanings, a proposed relationship explanation, and any ambiguity or confusion risk. Include proposed mnemonic edits when relevant, so the user can approve concrete wording.

The user may approve individual pairs or explicitly approve a batch. Apply only approved pairs and approved accompanying content. Approval of a relationship alone does not authorize unrelated mnemonic rewrites, new entries, or changes to meanings.

Unapproved candidates remain in the discussion; do not persist them as active relationships or include them in Anki exports. Human approval is a maintenance workflow, not a new approval-state field in YAML. Validation checks structural correctness, not linguistic truth or whether a person approved a particular edit.

For the approved closure/disclosure pair, distinguish the repository's programming sense of closure from disclosure's reporting sense. Do not label these senses as opposites merely because the words look related. This same review rule applies to all other candidates.

## Approved relationship batch

Approved on 2026-10-30 with the user's instruction: “全部批准”：确认这 21 组关联. The batch contains 13 English pairs and 8 Japanese pairs, connecting 42 distinct existing entries. Each row represents one undirected association; its review code is not a new entry ID.

The words and meanings below are snapshots of the reviewed source entries. During implementation, resolve each endpoint by stable ID and verify that its learning target still matches this record. Export current Word and Meaning values rather than copying this table into generated data.

The relationship explanations and confusion points record the reviewed rationale. They are not final `mnemonic` text and do not authorize automatic mnemonic rewrites, new etymological claims, additional pairs, new entries, or changes to meanings. Proposed mnemonic wording still needs content review.

| Code | A: stable ID · word · reviewed meaning         | B: stable ID · word · reviewed meaning          | Association rationale                    | Confusion risk                                                      |
| ---- | ---------------------------------------------- | ----------------------------------------------- | ---------------------------------------- | ------------------------------------------------------------------- |
| E01  | `closure-tech` · closure · 闭包                | `disclosure-workflow` · disclosure · 披露说明   | 词形相近，适合并列辨认和记忆             | 当前两个技术义项不是反义词，不能理解为“闭包／非闭包”                |
| E02  | `encode-tech` · encode · 编码                  | `decode-tech` · decode · 解码                   | 对照“转换为编码表示”和“按对应规则还原”   | 理解时要保留具体编码规则，不能只记成泛泛的“转换”                    |
| E03  | `register-tech` · register · 注册              | `unregister-tech` · unregister · 取消注册       | 对照加入与移出注册机制                   | 取消注册不等于删除对象或销毁全部资源                                |
| E04  | `mount-tech` · mount · 挂载                    | `unmount-tech` · unmount · 卸载                 | 同一组件生命周期中的对应操作             | 当前语境是组件或根节点，不是软件安装和卸载                          |
| E05  | `staged-git` · staged · 已暂存的               | `unstaged-git` · unstaged · 未暂存的            | 对照更改是否进入 Git 索引                | 已暂存不表示已经提交；判断对象应是具体更改                          |
| E06  | `infer-tech` · infer · 推断                    | `imply-general` · imply · 暗示                  | 对照“根据信息得出结论”和“使表达隐含结论” | 二者方向不同，不应作为同义词替换                                    |
| E07  | `preserve-tech` · preserve · 保留；保持不变    | `retain-tech` · retain · 保留；继续维持         | 中文释义接近，适合比较语境侧重点         | 当前 preserve 强调保持行为或约束，retain 强调仍然保有；不是绝对互斥 |
| E08  | `reproduce-tech` · reproduce · 复现            | `reproducible-tech` · reproducible · 可复现的   | 动作与性质之间的构词关联                 | 一个是动词，一个是形容词；“可复现”不等于已经完成复现                |
| E09  | `artifact-tech` · artifact · 构建产物          | `artifact-workflow` · artifact · 工作成果文件   | 同词不同范围，已有备注明确区分           | 构建产物和报告、草稿等工作成果有交集，但不能完全等同                |
| E10  | `literal-programming` · literal · 字面量       | `literal-text` · literal · 按原样解释的；字面的 | 对照名词术语与形容词用法                 | “源码中的字面量”不等于“按原样处理的任意文本”                        |
| E11  | `resolve-tech` · resolve · 解析；确定所指对象  | `resolve-workflow` · resolve · 解决；处理完毕   | 同一动词在程序分析和工作沟通中的不同义项 | 解析引用不一定是在解决故障                                          |
| E12  | `revision-git` · revision · 代码版本；提交版本 | `revision-general` · revision · 修订            | 对照版本指称与内容修改                   | 文档修订不一定对应一个 Git 提交                                     |
| E13  | `scope-lexical` · scope · 作用域               | `scope-workflow` · scope · 工作范围             | 都涉及边界，但边界对象不同               | 名称可见性规则和任务范围不能混为一谈                                |
| J01  | `kaishain` · 会社員 · 公司职员                 | `shain` · 社員 · 公司成员；职员                 | 现有备注已双向说明区别                   | 会社員侧重职业类别，社員侧重某家公司的所属关系                      |
| J02  | `kyoushi` · 教師 · 教师                        | `sensei` · 先生 · 老师等称呼                    | 对照职业名称与称呼、尊称                 | 先生不限于教师，通常不用于自称                                      |
| J03  | `dare` · 誰 · 谁                               | `donata` · どなた · 哪位                        | 普通说法与更礼貌的对应表达               | 选择取决于交流对象和场合，不只是中文翻译不同                        |
| J04  | `ano-hito` · あの人 · 那个人                   | `ano-kata` · あの方 · 那位；礼貌说法            | 同一指人表达的礼貌程度对照               | 此处「方」读作「かた」，不能套用其他语境的读音                      |
| J05  | `nansai` · 何歳 · 几岁                         | `oikutsu` · おいくつ · 几岁；多少               | 对照询问年龄的普通与礼貌表达             | おいくつ还可用于询问数量，并非任何语境都等于何歳                    |
| J06  | `hanasu` · 話す · 说；讲话                     | `hanasemasen` · 話せません · 不会说；不能说     | 联系字典形和可能形式的礼貌否定           | 話せません强调不能说，不是单纯的“不说”                              |
| J07  | `yaru` · やる · 做；干                         | `yareru` · やれる · 能做；能够做到              | 对照行为与做得到的能力、条件             | やれる不是过去式，也不表示已经完成                                  |
| J08  | `sukuu` · 救う · 拯救；救助                    | `sukueru` · 救える · 能够拯救                   | 字典形与可能形式的关联                   | 能够救助不等于已经救助成功                                          |

E03, E04, E08, J06, J07, and J08 already have corresponding prerequisite links. Preserve those directed learning-order relationships independently of the new bidirectional displays.

After implementation is authorized and the schema supports associations, add each approved pair once in source data; the exporter supplies the reverse display. The content approval for these unchanged pairs is already recorded and need not be requested again. Do not hardcode this batch into the general association logic. Additional candidates require separate confirmation.

## Proposed Anki field order

Replace the README's field-name code block with a Markdown table numbered to match the Anki Fields UI. The sequence number is a documentation aid, not the stable `ID` value.

Preserve the first twelve fields and append four new fields:

| No. | Anki field    | Source or derivation                          |
|:----|:--------------|:----------------------------------------------|
| 1   | ID            | `id`; permanent note identity                 |
| 2   | Language      | Canonicalized `language`                      |
| 3   | Word          | `word`                                        |
| 4   | Pronunciation | Optional `pronunciation`                      |
| 5   | PartOfSpeech  | Optional `part_of_speech`                     |
| 6   | Meaning       | `meaning`                                     |
| 7   | Context       | Optional `context`                            |
| 8   | Example       | Optional `example`                            |
| 9   | Notes         | Optional `notes`                              |
| 10  | Audio         | Sound reference derived from optional `audio` |
| 11  | TTSLanguage   | Derived example-speech locale                 |
| 12  | StudyOrder    | Derived introduction rank                     |
| 13  | Etymology     | Optional `etymology`                          |
| 14  | WordFormation | Optional `word_formation`                     |
| 15  | Mnemonic      | Optional `mnemonic`                           |
| 16  | RelatedWords  | Derived direct associations                   |

The proposed TSV has sixteen note fields and Anki tags in column 17. Update the column headers and tag mapping together. Do not create a Tags note field or change first-field ID matching. Keep `StudyOrder` as the browser sort field.

## Card presentation

- Preserve the current front, including pronunciation and word audio. Keep translations and memory aids on the back.
- On the back, retain Word, pronunciation/audio, part of speech, Meaning, optional Context, and optional Example. Follow these with Etymology, WordFormation, Mnemonic, RelatedWords, then Notes. Show present Notes in full, without truncation or collapse controls.
- Use separate conditional blocks for Context, Example, and each new display field. Put the section label and spacing-bearing wrapper inside the condition. Keep example speech within the Example condition and its existing TTSLanguage condition. Missing content must produce no visible label, replay control, empty container, divider, or reserved space.
- The note type still contains the fields even when a particular card does not display them. Anki's field editor is not dynamically restructured per entry.
- Keep the existing mobile typography, semantic theme tokens, wrapping behavior, and native replay controls. Add no JavaScript merely to hide empty content and no heavy nested card decoration.
- Use a neutral related-vocabulary label. Do not imply every neighbor is a synonym, antonym, or prerequisite.

## Implementation sequence after approval

1. Update Entry parsing and validation for the five optional text fields and `related_entries`. Retain current behavior for all other fields.
2. Validate association references against the complete collection and derive deterministic direct-neighbor displays independently of learning order.
3. Extend TSV output to the proposed seventeen-column layout, preserving text escaping, stable identity, and validation-before-replacement behavior.
4. Update conditional back-template sections and the numbered field table in the README. Keep templates and field order canonical there.
5. Update [AGENTS.md](../AGENTS.md), the [Codex vocabulary guide](ADD_VOCABULARY_WITH_CODEX.md), and the [Anki migration guide](ANKI_MIGRATION_GUIDE.md). Explain concept names, short principles in Meaning, detailed Notes, optional Context and Example, candidate approval, blank-value handling, and the new tag-column position. Update other current mapping references without rewriting historical verification records.
6. Extend audio source-preservation checks and tests to include every new field. Do not change spoken-input selection, voice behavior, filenames, or existing recordings.
7. Add focused regression tests and complete the verification below. Use synthetic fixtures for association tests; do not treat them as approved content for the real collection.
8. Apply the 21 approved pairs from the initial batch once the feature is implemented and verified. Propose any additional associations or memory-aid text separately for human confirmation. Do not perform an automatic whole-collection enrichment or Context/Example deletion pass. The five planning cases do not authorize a separate content import.

## Migration and risks

Existing entries remain valid without the new fields. Existing note types need the four appended fields, updated back template and styling as needed, and a complete reimport. Keep notes and IDs intact rather than deleting or recreating notes. No audio regeneration is required.

Verify that importing an empty Context, Example, or memory-aid field actually clears an old Anki value under the documented update settings. Verify that changed or removed relationships refresh both cards. Do not assume field order alone establishes correct import mapping.

Risks include inaccurate etymology or conceptual claims, misleading associations between different senses, loss of essential conditions in short explanations, incorrect speech language, accidental changes to blank-value rules for unrelated fields, stale reverse displays, and mismapped TSV columns. Mitigate these with source checks, human content approval, self-contained Meaning, structured Notes, conditional sections, strict scoped validation, full builds, and isolated import tests. Notes length itself is not a defect and does not justify removing useful material.

Do not modify a personal Anki collection, perform synchronization, stage, commit, or push as part of creating this plan. Future implementation should document the manual migration and report any live-device verification limits.

## Verification and acceptance

After implementation, run from the repository root:

```bash
go test ./...
go vet ./...
go run scripts/validate/main.go
go run scripts/build/main.go
git diff --check
```

Also verify:

- Omitted, empty, whitespace-only, and null values work for the five designated text fields. Nonempty values survive export; required fields and other optional-field rules retain their existing validation.
- Association validation rejects wrong types, malformed IDs, empty lists, duplicate list members, self-links, and missing targets. Reciprocal declarations and cycles are valid.
- Single-endpoint declarations render both ways, repeated pair declarations deduplicate, chains remain non-transitive, and neighbor order is independent of declaration order.
- Updating a neighbor's Word or Meaning updates the derived display. Deleting or removing references behaves as documented; invalid references leave the existing export untouched.
- Associations alone do not alter introduction ranks or prerequisite validation. HTML-like text and literal audio markers remain inert in all new display fields.
- For the approved initial batch, verify exactly 21 unique pairs and both display directions across the 42 intended entries, without introducing additional associations. Preserve existing prerequisite links, meanings, and recordings.
- Two builds produce identical bytes, all rows have seventeen columns, and tags occupy column 17.
- An isolated Anki collection can migrate from the current twelve-field layout, import twice without duplicates, update and clear optional values, and preserve note/card identities, scheduling, and review history.
- Missing-field combinations leave no empty sections on card backs, including when both Context and Example are absent. Missing Example also removes its speech control. Long multilingual Notes retain their content and line breaks without truncation and remain readable in light and dark modes on narrow screens. Existing word audio and present-example TTS still work; Meaning and Notes are not sent to TTS.
- Use synthetic fixtures reflecting the five planning cases to verify field optionality, detailed Notes, and unchanged plain-text escaping without adding real entries. Keep the existing concept examples and the approved relationship batch as separate content scopes.
- Audio metadata insertion preserves all new source values and unrelated bytes. Existing entries and media remain unchanged outside specifically approved content edits.
- Documentation links and numbered field mappings agree. Report automated, isolated-backend, and live Anki Desktop/AnkiDroid checks separately; do not claim unperformed playback or visual checks passed.

No implementation or runtime verification has been performed by creating this plan.
