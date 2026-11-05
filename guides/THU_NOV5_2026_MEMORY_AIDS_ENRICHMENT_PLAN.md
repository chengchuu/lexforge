# November 5, 2026 memory-aid enrichment plan

Status: **Implemented and statically verified on 2026-11-05 — WF01–WF14, ETY01–ETY08, N01–N07, and R01–R26 applied.**

The user approved all items, then explicitly authorized the missing 大学生 entry: "包含新增大学生并完成 R26（不新增音频）". This is the sole addition to the original enrichment scope. The proposal text, inventory, ledger, and planning checks below preserve the planning baseline; the [implementation record](#implementation-record) reports current results and supersedes earlier pending/blocked states. No native Anki rendering or playback verification is claimed.

This batch screens the complete collection and proposes selective additions. Creating or reviewing this document does not authorize entry, audio, schema, tooling, template, dependency, export, Git staging, or Anki changes. Approval codes in other plans do not approve codes in this file.

## Navigation

- [Scope and baseline](#scope-and-baseline)
- [Review method and decisions](#review-method-and-decisions)

- [Approval summary](#approval-summary)
- [WF proposals](#wf-proposals)

- [ETY proposals](#ety-proposals)
- [N proposals](#n-proposals)

- [Relationship proposals](#relationship-proposals)
- [Source register](#source-register)

- [Deferred investigations](#deferred-investigations)
- [Existing aid inventory](#existing-aid-inventory)

- [Existing relationship inventory](#existing-relationship-inventory)
- [Complete review ledger](#complete-review-ledger)

- [Implementation after approval](#implementation-after-approval)
- [Planning verification](#planning-verification)
- [Implementation record](#implementation-record)

## Scope and baseline

Planning date: **2026-11-05 (Asia/Shanghai)**. Git HEAD: `c197af2a4784124eb87e739a15fa17b9ca20c70f`. Initial Git status was clean. The requested plan path did not exist; completed batches in other guides are preserved. During planning, an independent documentation commit advanced HEAD to `e4bbc82c4df9e2e61a008126aa8a994225ac27bf`. Its only changed file, MAINTENANCE_PROMPTS.md, was re-read; the wording clarifications do not change this task’s scope. That work is preserved.

| Baseline                                                                   | Actual value  |
| :------------------------------------------------------------------------- | :------------ |
| Reviewed stable IDs                                                        | 1241          |
| English / Japanese                                                         | 999 / 242     |
| Other languages in current collection                                      | 0             |
| Existing etymology / word_formation / mnemonic values                      | 36 / 138 / 77 |
| Existing text aid values / entries containing them                         | 251 / 238     |
| Existing direct association pairs / declarations / local declaration lists | 89 / 89 / 86  |
| Existing directed neighbor displays                                        | 178           |
| Existing relationship display endpoints                                    | 170           |
| Entries declaring audio / existing MP3 files                               | 1241 / 1241   |
| Entry files / MP3 files included in preservation hashes                    | 1241 / 1241   |
| Baseline files hashed, including ignored generated files                   | 2512          |

Stable IDs, not filenames, determine the scope. The ledger links each ID to its actual source path. Existing audio declarations and recordings were inventoried for preservation, not decoded or listened to. No pronunciation or playback claim follows from a file count.

Baseline SHA-256 fingerprints (for comparison, not a replacement for the implementation snapshot):

| Set                                               | Fingerprint                                                        |
| :------------------------------------------------ | :----------------------------------------------------------------- |
| Sorted IDs, one per line with final newline       | `e3390608a0e89744fa58d2832b45cdd73053da3ff1ef13d7f88f57ce2606a730` |
| Entry path + NUL + file SHA-256 + newline, sorted | `0ac694012524df7cce4faaf5ac2e4051899360faec2ad09df41641ab71866c99` |
| MP3 path + NUL + file SHA-256 + newline, sorted   | `9801d51422b446d22a11d9f8c75c9035e64b48330605cebb7015e54470f8e682` |

Read: [AGENTS.md](../AGENTS.md), [README.md](../README.md), [Entry schema](ENTRY_SCHEMA.md), [Learning order](LEARNING_ORDER.md), and the [maintenance workflow](MAINTENANCE_PROMPTS.md). Inspected [parsing](../internal/vocabulary/entries.go), [export](../internal/vocabulary/export.go), [associations](../internal/vocabulary/related.go), and the README back template. Historical completed import/enrichment records are background only, not execution instructions.

The implemented fields are `etymology`, `word_formation`, `mnemonic`, and `related_entries`. `RelatedWords` is derived. The exporter emits 16 note fields and the tags column (17 total), escapes source text, and turns newlines into `<br>`; it does not render Markdown. Existing conditions hide absent aids, Context, Example/TTS, RelatedWords, and Notes. No template edits are proposed. Associations display both direct neighbors, with language tags for cross-language neighbors, and do not change StudyOrder.

## Review method and decisions

All 1,241 YAML mappings were read into the review inventory, including Word, Meaning, language, Context, Example, complete Notes, existing aids, prerequisites, and local relationships. Screening compared meanings across the collection and checked both association directions. Screening is not a claim that every historical statement, quotation, entity, or source URL was independently reverified.

Selection favors nontransparent spelling families, reusable sentence structure, verified origins that illuminate the current meaning, and practical confusion risks. Transparent vocabulary can remain without an aid. A longer Notes explanation is not moved merely to populate another field. Existing source text is retained unless the exact removal below is approved. No existing aid replacement or factual correction is proposed; retaining an aid is not a fresh certification of all its historical claims.

No-change decisions include useful existing aids, adequate Context/Notes, ordinary transparent vocabulary, and names or fictional terms without verified additional history. The complete ledger accounts for each ID once. Deferred cases are explicit investigations, not entry defects. Proposed additions do not imply new entries, prerequisites, pronunciation, recordings, or automatic associations.

The 251 existing text aids and 89 direct pairs remain intact, including already implemented TOTK enrichment. Existing associations such as compile–compiler, audio–pronunciation, and 子猫–猫ちゃん are not proposed again. Prerequisite links such as stack → call stack are not counted as existing associations; those specific new associations still need approval.

## Approval summary

| Approval group                                                  | Proposed count | Scope                                                                                              |
| :-------------------------------------------------------------- | :------------- | :------------------------------------------------------------------------------------------------- |
| WF01–WF14                                                       | 14             | word_formation: 13 additions; 1 explicit Notes move with a sourced extension                       |
| ETY01–ETY08                                                     | 8              | etymology additions                                                                                |
| N01–N07                                                         | 7              | mnemonic additions; original usage/counting prompts, not historical claims                         |
| Text aid values                                                 | 29             | 27 unique target entries; baseline and 〜てみる each receive two different fields                  |
| Notes updates                                                   | 27             | Exact source-line appends; WF14 also removes only the specified first structure sentence           |
| R01–R25                                                         | 25             | New direct pairs, separately approved; 25 new list members                                         |
| Relationship declaration endpoints                              | 25             | Unique YAML entries that would receive a list member                                               |
| New directed relationship displays                              | 50             | 49 unique display endpoints, counting each entry once                                              |
| Potential source entries changed if all text and pairs approved | 43             | Union of text targets and declaration endpoints; reverse-only displays do not require source edits |
| All entries affected in source or derived relationship display  | 63             | Union, not the sum of unlike counts                                                                |
| Retained existing aid values / direct pairs                     | 251 / 89       | No replacements, removals, or reciprocal redeclarations                                            |
| Additional entries / audio recordings                           | 0 / 0          | Excluded from this enrichment                                                                      |

There are 22 English and 7 Japanese proposed aid values. N07 adds one existing-entry target; no entry is created. R26 is a separately listed blocked candidate and is excluded from the existing-ID pair totals below. The 25 proposed pairs include 23 same-language pairs and 2 cross-language pairs. Of the 1,241 screened entries, 63 have a text or relationship proposal; 1,178 have neither, including 18 explicitly deferred additional investigations. Existing-aid retention overlaps these decisions and must not be added as a separate entry total.

Approval received on 2026-11-05: **WF01–WF14, ETY01–ETY08, N01–N07**, including their exact Notes appends and the WF14 move, and **R01–R26**. R26’s missing-entry prerequisite was separately authorized in the follow-up response quoted above and has been completed. All listed codes are applied. Pair approval alone does not authorize other mnemonic changes, and no unlisted relationship was inferred.

### Exact Notes append contract

For each approved text proposal, preserve the complete existing Notes as the prefix, except the exact WF14 move below. Append one source line per associated source, in source-code order. The exact line is shown in each source-register entry. Add a line only if its URL is not already present in that entry’s Notes; retain any existing source title/URL line unchanged rather than rewriting it. Deduplicate shared sources across approved proposals on the same entry. Insert a newline only when needed to separate the existing text from a new source line. If Notes is absent, create it from the needed lines. No source line is added solely for a relationship proposal. N07 explicitly reuses the already recorded 外国人 dictionary headword under its generic URL rather than appending the equivalent canonical S29 URL.

Every current target aid field is absent. The drafts below are complete new field values, with no Markdown or HTML inside the values. Descriptions of current Word/Meaning are baseline context, not proposed replacements. Source bodies provide facts; modern construction analyses and all mnemonic wording are editorial work.

## WF proposals

### WF01 — allowlist

Target: [allowlist-tech](../entries/allowlist-tech.yml) (`en`). Current Meaning: 允许列表

Decision: **add** into `word_formation` (currently absent). Exact proposed value:

```plain
allow + list: 把允许采用或通过的项目列成名单。这里表示明确准许的集合，不是默认允许所有项目。
```

Learning value: 从常用动词和名词理解复合名词，避免把名单与默认放行混淆。

Evidence: [S01](#s01). Modern construction/word-family analysis applied to the current entry; not an official naming-origin claim.

Limits: 现代构成分析；不声称确认该词的首创者或历史命名过程。

Notes: append the exact source-register line(s) under the append contract; retain all existing content.

### WF02 — baseline

Target: [baseline-tech](../entries/baseline-tech.yml) (`en`). Current Meaning: 基准输出

Decision: **add** into `word_formation` (currently absent). Exact proposed value:

```plain
base + line: 用作比较基础的参照线。测试语境中引申为保存的预期输出，供新结果对照。
```

Learning value: 把一般参照义连接到当前测试输出义。

Evidence: [S02](#s02). Modern construction/word-family analysis applied to the current entry; not an official naming-origin claim.

Limits: 现代构成分析；不是 benchmark 的所有含义，也不是对测试结果正确性的保证。

Notes: append the exact source-register line(s) under the append contract; retain all existing content.

### WF03 — compiler

Target: [compiler-tech](../entries/compiler-tech.yml) (`en`). Current Meaning: 编译器；执行程序编译的软件。

Decision: **add** into `word_formation` (currently absent). Exact proposed value:

```plain
现代词形可按 compile 去掉词尾 e 再接 -er 来记忆，得到 compiler。-er 在这里指执行编译的软件，不限于人。
```

Learning value: 把已有 compile 连接到工具名，避免把施事后缀限定为人。

Evidence: [S03](#s03). Modern construction/word-family analysis applied to the current entry; not an official naming-origin claim.

Limits: 这是现代词形和词族说明；历史上的 compiler 经法语传入，不能把该分析当作英语最初造词过程。

Notes: append the exact source-register line(s) under the append contract; retain all existing content.

### WF04 — identifier

Target: [identifier-tech](../entries/identifier-tech.yml) (`en`). Current Meaning: 标识符

Decision: **add** into `word_formation` (currently absent). Exact proposed value:

```plain
identify → identifier: 结尾 y 变为 i，再接 -er。这里指用来识别代码实体的名称，不是执行识别的人。
```

Learning value: 展示容易拼错的词形变化，并对齐代码标识符义。

Evidence: [S04](#s04). Modern construction/word-family analysis applied to the current entry; not an official naming-origin claim.

Limits: 说明词形与当前技术义，不推断某种语言允许的标识符字符集合。

Notes: append the exact source-register line(s) under the append contract; retain all existing content.

### WF05 — validation

Target: [validation-tech](../entries/validation-tech.yml) (`en`). Current Meaning: 验证

Decision: **add** into `word_formation` (currently absent). Exact proposed value:

```plain
validate → validation: 动词与动作名词对应，结尾 -ate 变为 -ation。这里指检查输入、实现或输出是否符合要求。
```

Learning value: 提供可复用的动词与名词拼写模式。

Evidence: [S05](#s05). Modern construction/word-family analysis applied to the current entry; not an official naming-origin claim.

Limits: 只说明本条用法；不把 validation 与 verification 定成所有行业通用且互斥的分类。

Notes: append the exact source-register line(s) under the append contract; retain all existing content.

### WF06 — verification

Target: [verification-tech](../entries/verification-tech.yml) (`en`). Current Meaning: 核实；验证

Decision: **add** into `word_formation` (currently absent). Exact proposed value:

```plain
verify → verification: 同一词族的动词和名词，注意名词拼作 verification，不能直接把 -tion 加在 verify 后。这里强调核实某个说法或行为是否成立。
```

Learning value: 记住非机械追加的名词形式。

Evidence: [S06](#s06). Modern construction/word-family analysis applied to the current entry; not an official naming-origin claim.

Limits: 只是同族词形说明，不声称这是完整的历史派生过程。

Notes: append the exact source-register line(s) under the append contract; retain all existing content.

### WF07 — invocation

Target: [invocation-workflow](../entries/invocation-workflow.yml) (`en`). Current Meaning: 一次调用

Decision: **add** into `word_formation` (currently absent). Exact proposed value:

```plain
invoke → invocation: 同一词族的动词和名词，注意 invok- 与 invoc- 的拼写对应。本条 invocation 指一次调用及其执行过程。
```

Learning value: 把调用动作与一次调用事件连接起来。

Evidence: [S07](#s07). Modern construction/word-family analysis applied to the current entry; not an official naming-origin claim.

Limits: 不是把 -ation 直接接在完整 invoke 后；不把古代祈求义当成工具调用的定义。

Notes: append the exact source-register line(s) under the append contract; retain all existing content.

### WF08 — resumption

Target: [resumption-workflow](../entries/resumption-workflow.yml) (`en`). Current Meaning: 恢复继续

Decision: **add** into `word_formation` (currently absent). Exact proposed value:

```plain
resume (继续) → resumption: 对应的名词使用 resumpt- 词干，表示中断后的继续，不是直接在 resume 后追加 -tion。
```

Learning value: 解释容易拼错且不完全透明的词族关系。

Evidence: [S08](#s08). Modern construction/word-family analysis applied to the current entry; not an official naming-origin claim.

Limits: 这里的 resume 是继续的动词，不是履历这一名词义。

Notes: append the exact source-register line(s) under the append contract; retain all existing content.

### WF09 — overwrite

Target: [overwrite-tech](../entries/overwrite-tech.yml) (`en`). Current Meaning: 覆盖

Decision: **add** into `word_formation` (currently absent). Exact proposed value:

```plain
over + write: 在原有内容上写入新内容；计算机语境中就是用新信息覆盖旧信息。
```

Learning value: 把空间上的覆盖连接到文件或值替换。

Evidence: [S09](#s09). Modern construction/word-family analysis applied to the current entry; not an official naming-origin claim.

Limits: 现代构成分析；不声称覆盖后一定不可恢复，也不等于追加内容。

Notes: append the exact source-register line(s) under the append contract; retain all existing content.

### WF10 — read-only

Target: [read-only-workflow](../entries/read-only-workflow.yml) (`en`). Current Meaning: 只读的

Decision: **add** into `word_formation` (currently absent). Exact proposed value:

```plain
read + only: 只允许读取。作为修饰语通常用连字符写成 read-only，限制的是修改操作，不是要求只能读一次。
```

Learning value: 解释复合修饰语和容易逐字误解的 only。

Evidence: [S10](#s10). Modern construction/word-family analysis applied to the current entry; not an official naming-origin claim.

Limits: 遵循当前只读工作流义，不扩展到某种特定权限系统的全部行为。

Notes: append the exact source-register line(s) under the append contract; retain all existing content.

### WF11 — 私は外国人です

Target: [watashi-wa-gaikokujin-desu](../entries/watashi-wa-gaikokujin-desu.yml) (`ja`). Current Meaning: 我是外国人

Decision: **add** into `word_formation` (currently absent). Exact proposed value:

```plain
私 + は + 外国人 + です。は 标记话题，作助词时读 わ；外国人です 说明身份。自我介绍中主语清楚时，可以省略 私は。
```

Learning value: 把已有单词组织成可迁移的名词判断句。

Evidence: [S11](#s11). Modern construction/word-family analysis applied to the current entry; not an official naming-origin claim.

Limits: 只是句法说明；保留原 Word 和 YAML 读音，不把发音提示写回 pronunciation。

Notes: append the exact source-register line(s) under the append contract; retain all existing content.

### WF12 — アメリカから来ました

Target: [amerika-kara-kimashita](../entries/amerika-kara-kimashita.yml) (`ja`). Current Meaning: 我来自美国

Decision: **add** into `word_formation` (currently absent). Exact proposed value:

```plain
アメリカ + から + 来ました。から 标记来源地；来ました 是 来る 的礼貌过去式。整句在自我介绍中表达来自美国，不等于声明美国国籍。
```

Learning value: 把固定练习句与可替换地名的来源结构连接起来。

Evidence: [S11](#s11). Modern construction/word-family analysis applied to the current entry; not an official naming-origin claim.

Limits: 根据本条自我介绍语境分析，不推断说话人的国籍或旅行历史。

Notes: append the exact source-register line(s) under the append contract; retain all existing content.

### WF13 — 日本人ではありません

Target: [nihonjin-dewa-arimasen](../entries/nihonjin-dewa-arimasen.yml) (`ja`). Current Meaning: 我不是日本人

Decision: **add** into `word_formation` (currently absent). Exact proposed value:

```plain
日本人 + ではありません。ではありません 是否定名词谓语的礼貌形式；句中没有明说主语，是否指我由上下文决定。
```

Learning value: 把身份名词与礼貌否定形式分开理解。

Evidence: [S12](#s12). Modern construction/word-family analysis applied to the current entry; not an official naming-origin claim.

Limits: 保留当前中文释义；解释省略主语，不另建条目或改写现有例句。

Notes: append the exact source-register line(s) under the append contract; retain all existing content.

### WF14 — 〜てみる

Target: [te-miru-grammar](../entries/te-miru-grammar.yml) (`ja`). Current Meaning: 试着做……

Decision: **move + extend** into `word_formation` (currently absent). Exact proposed value:

```plain
动词て形 + みる；随前接动词的て形变化，也可出现 でみる，如 読んでみる。
```

Learning value: 把已有 Notes 中的结构移到专用字段，并补清 で形变体。

Evidence: [S13](#s13). Modern construction/word-family analysis applied to the current entry; not an official naming-origin claim.

Limits: 只迁移下列明确指定的结构句；占位符说明与作品来源必须保留。

Notes: append the exact source-register line(s) under the append contract; retain all existing content except the explicitly approved move below.

Exact current Notes:

```plain
结构：动词て形 + みる。占位符「〜」不是需要朗读的文字。 学习来源: 《地獄少女》。
```

Remove only this exact clause from Notes:

```plain
结构：动词て形 + みる。
```

Exact Notes after the approved move and source append:

```plain
占位符「〜」不是需要朗读的文字。 学习来源: 《地獄少女》。
Memory-aid source: Tofugu — 〜てみる. https://www.tofugu.com/japanese-grammar/temiru/
```

The placeholder explanation and 地獄少女 provenance remain intact. N05 would use the same source line and must not append it twice. No other Notes move is approved by WF14.

## ETY proposals

### ETY01 — caveat

Target: [caveat-general](../entries/caveat-general.yml) (`en`). Current Meaning: 限定说明；注意事项

Decision: **add** into `etymology` (currently absent). Exact proposed value:

```plain
来自拉丁语 caveat，原意近于"务必留意"，是 cavere (留心、防备) 的一种动词形式。现代英语把它用作提醒限制条件或风险的名词。
```

Learning value: 把注意事项与留心、防备的历史核心连接起来。

Evidence: [S14](#s14). Verified historical dictionary facts, paraphrased in Chinese.

Limits: 拉丁语原句式经过意译；不把现代技术文档的限定说明说成古罗马技术用语。

Notes: append the exact source-register line(s) under the append contract; retain all existing content.

### ETY02 — anonymous

Target: [anonymous-tech](../entries/anonymous-tech.yml) (`en`). Current Meaning: 匿名的

Decision: **add** into `etymology` (currently absent). Exact proposed value:

```plain
经晚期拉丁语 anonymus 追溯到希腊语 anonymos (无名的): an- 表示没有，onyma 是 onoma (名字) 的一种方言形式。代码中的匿名沿用没有显式名称这一核心。
```

Learning value: 用可核实的无名词源连接匿名类或函数。

Evidence: [S15](#s15). Verified historical dictionary facts, paraphrased in Chinese.

Limits: 历史词形说明，不把现代英语拆成任意可独立使用的词根。

Notes: append the exact source-register line(s) under the append contract; retain all existing content.

### ETY03 — adjacent

Target: [adjacent-tech](../entries/adjacent-tech.yml) (`en`). Current Meaning: 相邻的

Decision: **add** into `etymology` (currently absent). Exact proposed value:

```plain
来自拉丁语 adiacens，关联动词 adiacere (位于旁边、靠近)，由 ad (向、在旁) 与 iacere (躺、处于) 构成。历史核心是位置靠近。
```

Learning value: 把相邻与空间关系联系起来。

Evidence: [S16](#s16). Verified historical dictionary facts, paraphrased in Chinese.

Limits: 不由词源断言所有用法都必须接触，也不扩大成对当前技术上下文的纠正。

Notes: append the exact source-register line(s) under the append contract; retain all existing content.

### ETY04 — mutable

Target: [mutable-tech](../entries/mutable-tech.yml) (`en`). Current Meaning: 可修改的

Decision: **add** into `etymology` (currently absent). Exact proposed value:

```plain
来自拉丁语 mutabilis (可改变的)，源于 mutare (改变)。可记住 mut- 与变化有关；不要把它当作现代英语 mutate + -able 的直接拼接。
```

Learning value: 提供可核实的变化核心并阻止伪拆词。

Evidence: [S17](#s17). Verified historical dictionary facts, paraphrased in Chinese.

Limits: mut- 是历史关联提示，不是遇到相似字母就能套用的万能词根。

Notes: append the exact source-register line(s) under the append contract; retain all existing content.

### ETY05 — audit

Target: [audit-tech](../entries/audit-tech.yml) (`en`). Current Meaning: 审查

Decision: **add** into `etymology` (currently absent). Exact proposed value:

```plain
来自拉丁语 auditus (听、听取)，关联 audire (听)。早期账目审查采用口头听取的程序，后来 audit 泛指系统审查；现代代码审查不要求真的听录音。
```

Learning value: 用历史听取账目的背景连接系统审查。

Evidence: [S18](#s18). Verified historical dictionary facts, paraphrased in Chinese.

Limits: 历史程序不是现代审查操作要求，也不把所有 review 都等同于 audit。

Notes: append the exact source-register line(s) under the append contract; retain all existing content.

### ETY06 — audio

Target: [audio-general](../entries/audio-general.yml) (`en`). Current Meaning: 音频；录制或传输的声音。

Decision: **add** into `etymology` (currently absent). Exact proposed value:

```plain
现代英语名词 audio 从构词成分 audio- 独立出来，最终关联拉丁语 audire (听)，保留了声音、听觉这一核心。
```

Learning value: 把音频资料与听的词源联系起来。

Evidence: [S19](#s19). Verified historical dictionary facts, paraphrased in Chinese.

Limits: 不改写已有音频与发音的辨析，也不把 audio 与 audit 视为同义词。

Notes: append the exact source-register line(s) under the append contract; retain all existing content.

### ETY07 — regression

Target: [regression-tech](../entries/regression-tech.yml) (`en`). Current Meaning: 回归问题；行为退化

Decision: **add** into `etymology` (currently absent). Exact proposed value:

```plain
来自拉丁语 regressio (返回)，关联 regredi (退回)。软件中的 regression 用返回或倒退的核心帮助理解: 改动后原本正确的行为发生退化。
```

Learning value: 让软件回归问题与倒退的历史核心建立联系。

Evidence: [S20](#s20). Verified historical dictionary facts, paraphrased in Chinese.

Limits: 软件解释是现代应用；不是统计学回归的完整定义，也不表示每次返回旧状态都是 bug。

Notes: append the exact source-register line(s) under the append contract; retain all existing content.

### ETY08 — inspect

Target: [inspect-tech](../entries/inspect-tech.yml) (`en`). Current Meaning: 检查；查看

Decision: **add** into `etymology` (currently absent). Exact proposed value:

```plain
来自拉丁语 inspicere (察看、检查)，由 in- (向内) 与 specere (看) 构成。它的历史核心与现代仔细查看代码或文件的意思相连。
```

Learning value: 以向内看帮助记住检查动作。

Evidence: [S21](#s21). Verified historical dictionary facts, paraphrased in Chinese.

Limits: 仅说明历史来源；不是把 inspect 定义成必须进入物理内部。

Notes: append the exact source-register line(s) under the append contract; retain all existing content.

## N proposals

### N01 — outstanding

Target: [outstanding-tech](../entries/outstanding-tech.yml) (`en`). Current Meaning: 尚未完成的

Decision: **add** into `mnemonic` (currently absent). Exact proposed value:

```plain
辨义助记: outstanding work 先看还有哪些工作挂着未完成；不要自动读成出色的工作。这里记的是尚待解决，不是优秀。
```

Learning value: 提醒同形多义，不让褒义覆盖当前待办义。

Evidence: [S22](#s22). Original mnemonic/usage prompt; cited sources support factual distinctions, not authorship of this wording.

Limits: 原创用法助记；优秀是确实存在的另义，并非所有 outstanding work 都只能表示未完成。

Notes: append the exact source-register line(s) under the append contract; retain all existing content.

### N02 — overload

Target: [overload-tech](../entries/overload-tech.yml) (`en`). Current Meaning: 重载

Decision: **add** into `mnemonic` (currently absent). Exact proposed value:

```plain
辨义助记: 看到 function overload，先想同一个名字有几种调用签名，不要先想机器负载过高。重载与超载要靠语境区分。
```

Learning value: 区分非字面技术义与日常负载义。

Evidence: [S23](#s23), [S28](#s28). Original mnemonic/usage prompt; cited sources support factual distinctions, not authorship of this wording.

Limits: 原创用法助记；不声称不同语言的重载分派机制完全相同。

Notes: append the exact source-register line(s) under the append contract; retain all existing content.

### N03 — baseline

Target: [baseline-tech](../entries/baseline-tech.yml) (`en`). Current Meaning: 基准输出

Decision: **add** into `mnemonic` (currently absent). Exact proposed value:

```plain
辨义助记: 在这两个词条里，baseline 问"输出和保存的预期一样吗？"，benchmark 问"这个工作负载表现怎样？"。两者都可作比较参照，不能在所有语境中硬分成互斥概念。
```

Learning value: 用两个具体问题区分当前测试输出与性能测试义。

Evidence: [S02](#s02), [S24](#s24). Original mnemonic/usage prompt; cited sources support factual distinctions, not authorship of this wording.

Limits: 原创问题提示，针对当前条目；不构造通用行业术语标准，也不据此添加问答字段。

Notes: append the exact source-register line(s) under the append contract; retain all existing content.

### N04 — byte offset

Target: [byte-offset-tech](../entries/byte-offset-tech.yml) (`en`). Current Meaning: 字节偏移

Decision: **add** into `mnemonic` (currently absent). Exact proposed value:

```plain
计数助记: 对字符串 A😀B，从 0 开始计数时，B 的 UTF-16 码元位置是 3，UTF-8 字节偏移是 5。先问按什么单位计数，再比较位置；不要把码元数直接当作字节数。
```

Learning value: 提供可核算的最小实例，防止单位混淆。

Evidence: [S25](#s25), [S26](#s26). Original mnemonic/usage prompt; cited sources support factual distinctions, not authorship of this wording.

Limits: 原创计数示例，不是字源故事。这里只比较指定编码，不把可见字符、码点和码元一概等同。

Notes: append the exact source-register line(s) under the append contract; retain all existing content.

### N05 — 〜てみる

Target: [te-miru-grammar](../entries/te-miru-grammar.yml) (`ja`). Current Meaning: 试着做……

Decision: **add** into `mnemonic` (currently absent). Exact proposed value:

```plain
用法助记: ～たい 侧重想做，～てみる 侧重试着做、看看如何。想做不等于已经尝试；想试试可以组合成 ～てみたい。
```

Learning value: 把愿望、尝试和想尝试三个层次分开。

Evidence: [S13](#s13). Original mnemonic/usage prompt; cited sources support factual distinctions, not authorship of this wording.

Limits: 原创用法提示，不是出处故事；～てみる 的非过去式本身也不表示动作已经完成。

Notes: append the exact source-register line(s) under the append contract; retain all existing content.

### N06 — あなた

Target: [anata](../entries/anata.yml) (`ja`). Current Meaning: 你

Decision: **add** into `mnemonic` (currently absent). Exact proposed value:

```plain
用法助记: 不要把英语 you 每次都直接替换成 あなた。面对具体的人，先考虑是否省略称呼或使用姓名、身份称呼；あなた 的礼貌程度取决于关系和语境，不是固定的礼貌开关。
```

Learning value: 防止跨语言逐词替换和永远粗鲁的另一种过度规则。

Evidence: [S27](#s27). Original mnemonic/usage prompt; cited sources support factual distinctions, not authorship of this wording.

Limits: 原创使用提示；不是禁用所有第二人称，也不改写现有 Context。

Notes: append the exact source-register line(s) under the append contract; retain all existing content.

### N07 — 外国人: distinguish がいこくじん from だいがくせい

Target: [gaikokujin](../entries/gaikokujin.yml) (`ja`). Current Word: 外国人. Current Meaning: 外国人. Current pronunciation: がいこくじん. Current `mnemonic`: absent. Retain its existing `word_formation`, Notes, pronunciation, and prerequisites unchanged.

Decision: **add** into `mnemonic`. Exact proposed value:

```plain
辨音助记: 外国人读 がいこくじん，大学生读 だいがくせい。先抓结尾: じん 对应这里的 人，せい 对应这里的学生身份；再核对前半段 がいこく (外国) 与 だいがく (大学)。两者分别记成 外国 + 人、大学 + 生，不把相近的声音混成一串。外国人也可以是大学生，两者不是反义词。
```

Learning value: 针对用户明确报告的听辨困难，先用结尾区分，再核对前半段与意义构成。不是为了相似拼写而自动建立关系。

Evidence: [S29](#s29) supports 外国人 and its reading; [S30](#s30) supports 大学生, its reading, and the university-student sense. The contrast and listening strategy are original editorial learning guidance. This is not a claim that either word historically derives from the other.

Limits: 假名只是辨读提示，不是音频、音高重音标注或完整音标转写；本提案不改动 pronunciation 或录音，不新增词源说明。大学生 currently has no entry, so naming it in this aid does not authorize creating a card or writing a dangling association.

Notes: preserve the existing Notes completely. The existing generic 外国人 Kotobank reference already identifies the headword inspected as S29; reuse it and do not append a duplicate source line merely because S29 has a canonical URL with a numeric suffix. Append only the exact S30 source line, once, under the source-append contract. N07 requires text approval independently of R26.

## Relationship proposals

R01–R25 below are absent in both declaration directions at the baseline. R26 is a supplemental blocked candidate described separately; it has no existing target entry. Declare the target ID once on the stated endpoint; append to an existing list without removing, reordering, or replacing its members. Each approved pair produces two directed displays after a full export. Only the listed direct neighbors appear; there is no transitive expansion, navigation link, new prerequisite, or mastery-based unlock.

Both Words and Meanings below are current source values, not proposed rewrites. The rationale is an editorial learning recommendation grounded in these senses. R01 uses S02/S24, R11/R12 retain context-dependent pronoun distinctions (S27 supports the あなた caveat), R17 uses S13, and R25 uses S18/S19. These relationship rationales are plan-only: they are not automatically added to Mnemonic or Notes.

| Code | Endpoint A: ID / Word / current Meaning                                                                                               | Endpoint B: ID / Word / current Meaning                                                                                                 | Value                                                                    | Confusion risk                                                                   | Declare on → target ID                              |
| :--- | :------------------------------------------------------------------------------------------------------------------------------------ | :-------------------------------------------------------------------------------------------------------------------------------------- | :----------------------------------------------------------------------- | :------------------------------------------------------------------------------- | :-------------------------------------------------- |
| R01  | [baseline-tech](../entries/baseline-tech.yml) / baseline / 基准输出                                                                   | [benchmark-tech](../entries/benchmark-tech.yml) / benchmark / 性能基准测试                                                              | 同为比较参照，当前一个保存预期输出，一个衡量性能，适合并看辨义。         | 不是所有 baseline/benchmark 用法都互斥；关系不改变两条现有释义。                 | `baseline-tech` → `benchmark-tech`                  |
| R02  | [validation-tech](../entries/validation-tech.yml) / validation / 验证                                                                 | [verification-tech](../entries/verification-tech.yml) / verification / 核实；验证                                                       | 同译为验证，结合当前 Context 比较符合要求与核实结论的侧重点。            | 不要把局部解释升级成所有学科统一的 V&V 划分；可以重叠。                          | `validation-tech` → `verification-tech`             |
| R03  | [invocation-workflow](../entries/invocation-workflow.yml) / invocation / 一次调用                                                     | [invoke-tech](../entries/invoke-tech.yml) / invoke / 调用                                                                               | 动作名词与动词对应，帮助从执行调用连接到一次调用事件。                   | 不是新增词形卡，也不表示每个 invocation 都是函数调用。                           | `invocation-workflow` → `invoke-tech`               |
| R04  | [resumption-workflow](../entries/resumption-workflow.yml) / resumption / 恢复继续                                                     | [recovery-tech](../entries/recovery-tech.yml) / recovery / 错误后的恢复                                                                 | 中断后继续与错误后恢复都可能出现在任务记录中，适合比较。                 | 继续不必先有错误；恢复也不一定从上次中断点继续。                                 | `resumption-workflow` → `recovery-tech`             |
| R05  | [outstanding-tech](../entries/outstanding-tech.yml) / outstanding / 尚未完成的                                                        | [pending-tech](../entries/pending-tech.yml) / pending / 待处理的；待完成的                                                              | 当前都指待处理工作，帮助识别 outstanding 的非褒义。                      | 不是所有语境都可互换，尤其不能把优秀义 outstanding 代入。                        | `outstanding-tech` → `pending-tech`                 |
| R06  | [bypass-tech](../entries/bypass-tech.yml) / bypass / 绕过                                                                             | [workaround-tech](../entries/workaround-tech.yml) / workaround / 变通办法；暂时避开问题而不直接修复根因的做法。                         | 绕过某个处理层可能用于实施变通办法，连接动作与策略。                     | 不是同义词；bypass 可用于测试隔离，workaround 不必跳过检查。                     | `bypass-tech` → `workaround-tech`                   |
| R07  | [array-hole-javascript](../entries/array-hole-javascript.yml) / array hole / 数组空槽；数组长度范围内某个索引没有对应自有元素的情况。 | [own-property-javascript](../entries/own-property-javascript.yml) / own property / 自有属性；直接属于对象自身、而非从原型链继承的属性。 | 用自有属性是否存在理解数组空槽，与值为 undefined 的元素对照。            | 属性缺失不等于任何访问都必然得到 undefined；还需考虑原型链。                     | `array-hole-javascript` → `own-property-javascript` |
| R08  | [byte-offset-tech](../entries/byte-offset-tech.yml) / byte offset / 字节偏移                                                          | [code-unit-tech](../entries/code-unit-tech.yml) / code unit / 码元                                                                      | 都描述编码相关计量，适合在位置换算时同时辨认单位。                       | 字节偏移是位置，码元是单位；不能假定通用一对一换算。                             | `byte-offset-tech` → `code-unit-tech`               |
| R09  | [stack-data-structure](../entries/stack-data-structure.yml) / stack / 栈；按后进先出顺序添加和移除元素的数据结构。                    | [call-stack-javascript](../entries/call-stack-javascript.yml) / call stack / 调用栈；按后进先出顺序记录函数调用和返回位置。             | 抽象 LIFO 结构与函数调用这一应用形成概念到用途的关联。                   | 不是所有栈都记录调用；不把调用栈说成所有变量的存储位置。                         | `stack-data-structure` → `call-stack-javascript`    |
| R10  | [recursion-programming](../entries/recursion-programming.yml) / recursion / 递归；在求解或定义过程中再次处理同类子问题的方式。        | [backtracking-algorithm](../entries/backtracking-algorithm.yml) / backtracking / 回溯；逐步尝试选择，撤销不合适的选择并继续搜索。       | 回溯常通过递归表达探索与撤销，适合比较手段与搜索策略。                   | 递归不一定是回溯，回溯也不要求只能用递归实现。                                   | `recursion-programming` → `backtracking-algorithm`  |
| R11  | [watashi](../entries/watashi.yml) / 私 / 我                                                                                           | [boku](../entries/boku.yml) / 僕 / 我；男性自称                                                                                         | 同为第一人称，帮助结合现有语境选择自称。                                 | 僕 主要由男性使用，不是绝对性别禁令；不能机械排列成固定礼貌等级。                | `watashi` → `boku`                                  |
| R12  | [anata](../entries/anata.yml) / あなた / 你                                                                                           | [omae](../entries/omae.yml) / お前 / 你；较直接或粗鲁的第二人称                                                                         | 同指对方但关系和语气不同，适合区分日常称呼与影视中强硬表达。             | あなた 不保证礼貌，お前 也不能脱离关系一概翻成辱骂；不鼓励在陌生电话中使用后者。 | `anata` → `omae`                                    |
| R13  | [kami](../entries/kami.yml) / 神 / 神                                                                                                 | [deity-story](../entries/deity-story.yml) / deity / 神；神明。                                                                          | 连接日语和英语中当前神明义，帮助在动漫与游戏学习之间迁移词义。           | 只是词义桥梁，不声称不同文化中的神明概念或具体角色完全对应。                     | `kami` → `deity-story`                              |
| R14  | [neko](../entries/neko.yml) / 猫 / 猫                                                                                                 | [koneko](../entries/koneko.yml) / 子猫 / 小猫                                                                                           | 一般类别与幼年子类有明确差别，可复习 子 的限定作用。                     | 猫 不限定年龄，子猫 不等于所有被亲昵称呼的猫。已有 子猫—猫ちゃん 关系保留。      | `neko` → `koneko`                                   |
| R15  | [ki](../entries/ki.yml) / 木 / 树；木                                                                                                 | [jumoku](../entries/jumoku.yml) / 樹木 / 树木                                                                                           | 当前都可表示树，比较常用单词与较书面集合称呼。                           | 木 也可指木材；该义不能直接与 樹木 互换。                                        | `ki` → `jumoku`                                     |
| R16  | [kenkyuu](../entries/kenkyuu.yml) / 研究 / 研究；深入调查和探究                                                                       | [kenkyuusha](../entries/kenkyuusha.yml) / 研究者 / 研究人员                                                                             | 活动与从事活动的人，配合现有 者 构成说明复习。                           | 关系不是新的先修声明；不是所有做一次调查的人都适合称研究者。                     | `kenkyuu` → `kenkyuusha`                            |
| R17  | [tai-grammar](../entries/tai-grammar.yml) / ～たい / 想做……                                                                           | [te-miru-grammar](../entries/te-miru-grammar.yml) / 〜てみる / 试着做……                                                                 | 区分愿望和尝试，还能理解组合 ～てみたい。                                | ～てみる 非过去式并不宣告动作完成；关系批准不自动批准 N05。                      | `tai-grammar` → `te-miru-grammar`                   |
| R18  | [jibun](../entries/jibun.yml) / 自分 / 自己                                                                                           | [jibun-jishin](../entries/jibun-jishin.yml) / 自分自身 / 自己本身；自己本人                                                             | 比较自己与强调自己本人，连接基础形式与强调形式。                         | 自分 的指代要依语境，不能在所有句子中无差别替换。                                | `jibun` → `jibun-jishin`                            |
| R19  | [namida](../entries/namida.yml) / 涙 / 眼泪                                                                                           | [tear-general](../entries/tear-general.yml) / tear / 眼泪；泪水。                                                                       | 连接当前日语眼泪与英语泪水义；已有 tear 助记可提示另一同形义的风险。     | 不连到撕裂义，也不据关联新增发音或剧情解释。                                     | `namida` → `tear-general`                           |
| R20  | [moyasu](../entries/moyasu.yml) / 燃やす / 燃烧；使……燃烧                                                                             | [moeagaru](../entries/moeagaru.yml) / 燃え上がる / 熊熊燃烧；燃烧起来                                                                   | 通过现有他动词与自动词说明比较使其燃烧和燃烧起来。                       | 不是严格的同词干自动他动一对；不能凭这两个词推导所有燃烧动词的变化。             | `moyasu` → `moeagaru`                               |
| R21  | [call-stack-javascript](../entries/call-stack-javascript.yml) / call stack / 调用栈；按后进先出顺序记录函数调用和返回位置。           | [stack-trace-tech](../entries/stack-trace-tech.yml) / stack trace / 调用栈跟踪；记录一条执行路径上活动调用帧的信息。                    | 运行中的调用结构与记录调用帧的诊断信息适合并看。                         | 栈跟踪是记录或快照，不是调用栈本身，也不单独证明根因。                           | `call-stack-javascript` → `stack-trace-tech`        |
| R22  | [inspect-tech](../entries/inspect-tech.yml) / inspect / 检查；查看                                                                    | [static-inspection-tech](../entries/static-inspection-tech.yml) / static inspection / 静态检查                                          | 检查动作与不执行目标代码的检查方式，帮助把静态限制读清。                 | inspect 不限定静态；静态发现的可疑问题未必已经复现。                             | `inspect-tech` → `static-inspection-tech`           |
| R23  | [gaikoku](../entries/gaikoku.yml) / 外国 / 外国；本国以外的国家                                                                       | [gaikokujin](../entries/gaikokujin.yml) / 外国人 / 外国人                                                                               | 国家范围与人的身份称呼相连，结合现有 ～人 构成复习。                     | 外国 是相对说话立场，外国人 不限定某个国籍。                                     | `gaikoku` → `gaikokujin`                            |
| R24  | [etymology-language](../entries/etymology-language.yml) / etymology / 词源；词语的历史来源与演变。                                    | [formation-general](../entries/formation-general.yml) / formation / 形成；构成                                                          | 区分历史来源与构成方式；两者可互相参考，但不是同一研究问题。             | 现代可见拆分不自动证明词源；不因此新增词条或改写两个条目的字段。                 | `etymology-language` → `formation-general`          |
| R25  | [audit-tech](../entries/audit-tech.yml) / audit / 审查                                                                                | [audio-general](../entries/audio-general.yml) / audio / 音频；录制或传输的声音。                                                        | 经已核实的拉丁语 audire 联系审查的历史听取与现代音频，形成跨义词族记忆。 | 不是现代同义词；代码审查不用听录音，关联也不把审查变成音频功能。                 | `audit-tech` → `audio-general`                      |

Declaration changes, stated explicitly to preserve existing members:

- `anata`: current `related_entries` = `["kimi"]`; append only `["omae"]`.

- `array-hole-javascript`: current `related_entries` = absent; append only `["own-property-javascript"]`.

- `audit-tech`: current `related_entries` = absent; append only `["audio-general"]`.

- `baseline-tech`: current `related_entries` = absent; append only `["benchmark-tech"]`.

- `bypass-tech`: current `related_entries` = absent; append only `["workaround-tech"]`.

- `byte-offset-tech`: current `related_entries` = absent; append only `["code-unit-tech"]`.

- `call-stack-javascript`: current `related_entries` = absent; append only `["stack-trace-tech"]`.

- `etymology-language`: current `related_entries` = absent; append only `["formation-general"]`.

- `gaikoku`: current `related_entries` = absent; append only `["gaikokujin"]`.

- `inspect-tech`: current `related_entries` = absent; append only `["static-inspection-tech"]`.

- `invocation-workflow`: current `related_entries` = absent; append only `["invoke-tech"]`.

- `jibun`: current `related_entries` = absent; append only `["jibun-jishin"]`.

- `kami`: current `related_entries` = absent; append only `["deity-story"]`.

- `kenkyuu`: current `related_entries` = absent; append only `["kenkyuusha"]`.

- `ki`: current `related_entries` = absent; append only `["jumoku"]`.

- `moyasu`: current `related_entries` = absent; append only `["moeagaru"]`.

- `namida`: current `related_entries` = absent; append only `["tear-general"]`.

- `neko`: current `related_entries` = absent; append only `["koneko"]`.

- `outstanding-tech`: current `related_entries` = absent; append only `["pending-tech"]`.

- `recursion-programming`: current `related_entries` = absent; append only `["backtracking-algorithm"]`.

- `resumption-workflow`: current `related_entries` = absent; append only `["recovery-tech"]`.

- `stack-data-structure`: current `related_entries` = absent; append only `["call-stack-javascript"]`.

- `tai-grammar`: current `related_entries` = absent; append only `["te-miru-grammar"]`.

- `validation-tech`: current `related_entries` = absent; append only `["verification-tech"]`.

- `watashi`: current `related_entries` = absent; append only `["boku"]`.

If every pair is approved: **114 unique direct pairs**, **228 directed displays**, and **214 unique display endpoints**. Existing reciprocal-handling behavior remains unchanged. A display on an unchanged reverse endpoint is expected derived output, not an extra YAML edit.

### R26 — 外国人 / 大学生: approved and implemented

Status: **approved and implemented on 2026-11-05**. At planning, this was an additional blocked candidate outside the 25 existing-ID pairs R01–R25. The user separately authorized adding 大学生 without audio, resolving the block. The table below records endpoint availability at the planning baseline; current endpoint B is [daigakusei](../entries/daigakusei.yml). N07 was explicitly approved with the rest of the text proposals.

| Endpoint | Stable ID                  | Word   | Meaning                                                        | Availability                                    |
| :------- | :------------------------- | :----- | :------------------------------------------------------------- | :---------------------------------------------- |
| A        | `gaikokujin`               | 外国人 | 外国人 (current)                                               | Existing entry; reading がいこくじん            |
| B        | `daigakusei` (provisional) | 大学生 | 大学生；在大学就读的学生 (proposed, not a current entry value) | No current entry; verified reading だいがくせい |

Value: 用户明确提出两者发音容易混淆。并列显示有助于对比 がいこく / だいがく 和 じん / せい，结合身份含义分组记忆。来源见 S29、S30。

Confusion risks: 两词不是反义词或同义词，外国人也可以是大学生。じん 和 せい 在这里对应不同构成，不能据此推断所有 人、生 的读音。现有 `gakusei` means 学生 and does not replace the narrower 大学生 learning target. Relationship approval does not approve the explanatory wording in N07, new prerequisites, or audio changes.

Intended declaration after prerequisites are satisfied: append `daigakusei` once to `gaikokujin.related_entries`, preserving any existing members, including R23's reverse display. The current local list on gaikokujin is absent; R23 would be declared on gaikoku and must not be reciprocally redeclared here.

Original prerequisite for R26: a separately authorized 大学生 entry and explicit pair approval. Both are now satisfied. Added `entries/daigakusei.yml` with stable ID `daigakusei` and declared it once on `gaikokujin`. No custom RelatedWords value was written, and neither `gakusei` nor `daigaku` was substituted. This scoped exception authorizes no other new entries or recordings.

Planning accounting: **26 relationship candidates = 25 existing-ID pairs + 1 then-blocked pair**. Those baseline counts and the 1,241-row review ledger remain historical. Final accounting includes R26 and the separately added entry: 1,242 entries and 115 total direct pairs. See the implementation record for current declarations, displays, ranks, and preservation checks.

## Source register

Access date: **2026-11-05**. The listed headword/grammar/documentation sections were opened and inspected; search snippets and automated example aggregations are not evidence. Each source supports only the claim stated below. A historical dictionary entry is evidence for the stated origin, not proof of a present-day software definition or official product naming intent. Current source meanings and Context supply that application.

### S01

[MDN Web Docs — Permissions Policy](https://developer.mozilla.org/en-US/docs/Web/HTTP/Guides/Permissions_Policy)

Supports: Allowlists constrain permitted origins. Supports the modern allow + list analysis, not a historical coinage claim. The compiler-entry sense remains grounded in its existing Context.

Access and limits: Headword or named section inspected on 2026-11-05; not a complete audit of the site.

Exact Notes source line, subject to URL deduplication:

```plain
Memory-aid source: MDN Web Docs — Permissions Policy. https://developer.mozilla.org/en-US/docs/Web/HTTP/Guides/Permissions_Policy
```

### S02

[Merriam-Webster.com Dictionary — baseline](https://www.merriam-webster.com/dictionary/baseline)

Supports: A reference line and comparison data; supports modern base + line analysis. The saved-test-output application comes from the current entry.

Access and limits: Headword or named section inspected on 2026-11-05; not a complete audit of the site.

Exact Notes source line, subject to URL deduplication:

```plain
Memory-aid source: Merriam-Webster.com Dictionary — baseline. https://www.merriam-webster.com/dictionary/baseline
```

### S03

[Online Etymology Dictionary — compiler](https://www.etymonline.com/word/compiler)

Supports: Compiler and compile belong to the same family; the older noun was borrowed through French. WF03 is a modern learner analysis, not a claim that English first coined the historical noun by adding -er.

Access and limits: Headword or named section inspected on 2026-11-05; not a complete audit of the site.

Exact Notes source line, subject to URL deduplication:

```plain
Memory-aid source: Online Etymology Dictionary — compiler. https://www.etymonline.com/word/compiler
```

### S04

[Online Etymology Dictionary — identifier](https://www.etymonline.com/word/identifier)

Supports: Identifier is an agent noun of identify. WF04 describes the visible y-to-i spelling pattern; the programming application comes from the current entry.

Access and limits: Headword or named section inspected on 2026-11-05; not a complete audit of the site.

Exact Notes source line, subject to URL deduplication:

```plain
Memory-aid source: Online Etymology Dictionary — identifier. https://www.etymonline.com/word/identifier
```

### S05

[Online Etymology Dictionary — validation](https://www.etymonline.com/word/validation)

Supports: Validation is the action noun corresponding to validate; supports a modern word-family explanation, not a universal software V&V definition.

Access and limits: Headword or named section inspected on 2026-11-05; not a complete audit of the site.

Exact Notes source line, subject to URL deduplication:

```plain
Memory-aid source: Online Etymology Dictionary — validation. https://www.etymonline.com/word/validation
```

### S06

[Online Etymology Dictionary — verification](https://www.etymonline.com/word/verification)

Supports: Verification and verify share the historical verific- family. Supports the word-family contrast; the entry supplies its contemporary technical scope.

Access and limits: Headword or named section inspected on 2026-11-05; not a complete audit of the site.

Exact Notes source line, subject to URL deduplication:

```plain
Memory-aid source: Online Etymology Dictionary — verification. https://www.etymonline.com/word/verification
```

### S07

[Online Etymology Dictionary — invocation](https://www.etymonline.com/word/invocation)

Supports: Invocation/invoke and their Latin invoc- family. Supports a spelling-family explanation; does not itself document the modern CLI meaning.

Access and limits: Headword or named section inspected on 2026-11-05; not a complete audit of the site.

Exact Notes source line, subject to URL deduplication:

```plain
Memory-aid source: Online Etymology Dictionary — invocation. https://www.etymonline.com/word/invocation
```

### S08

[Online Etymology Dictionary — resumption](https://www.etymonline.com/word/resumption)

Supports: Resumption/resume and the resumpt- stem; distinguishes the spelling relationship from mechanically appending a suffix to resume.

Access and limits: Headword or named section inspected on 2026-11-05; not a complete audit of the site.

Exact Notes source line, subject to URL deduplication:

```plain
Memory-aid source: Online Etymology Dictionary — resumption. https://www.etymonline.com/word/resumption
```

### S09

[Merriam-Webster.com Dictionary — overwrite](https://www.merriam-webster.com/dictionary/overwrite)

Supports: The computer sense replaces old information with new information; supports the modern over + write analysis.

Access and limits: Headword or named section inspected on 2026-11-05; not a complete audit of the site.

Exact Notes source line, subject to URL deduplication:

```plain
Memory-aid source: Merriam-Webster.com Dictionary — overwrite. https://www.merriam-webster.com/dictionary/overwrite
```

### S10

[Merriam-Webster.com Dictionary — read-only](https://www.merriam-webster.com/dictionary/read-only)

Supports: Can be viewed but not changed or deleted; supports read + only.

Access and limits: Headword or named section inspected on 2026-11-05; not a complete audit of the site.

Exact Notes source line, subject to URL deduplication:

```plain
Memory-aid source: Merriam-Webster.com Dictionary — read-only. https://www.merriam-webster.com/dictionary/read-only
```

### S11

[The Japan Foundation — Irodori Starter, Lesson 3: よろしくお願いします](https://www.irodori.jpf.go.jp/assets/data/starter/pdf/X_L03.pdf)

Supports: Grammar notes, printed pages L3-14–15: N1はN2です, topic は pronounced わ, contextual omission, and place + から来ました. Supports WF11/WF12; does not establish citizenship from a place of origin.

Access and limits: Headword or named section inspected on 2026-11-05; not a complete audit of the site.

Exact Notes source line, subject to URL deduplication:

```plain
Memory-aid source: The Japan Foundation — Irodori Starter, Lesson 3: よろしくお願いします. https://www.irodori.jpf.go.jp/assets/data/starter/pdf/X_L03.pdf
```

### S12

[Tokyo University of Foreign Studies Language Modules — Nです／Nではありません](https://www.coelang.tufs.ac.jp/mt/ja/gmod/contents/card/001.html)

Supports: Noun predicate and polite negative noun predicate; WF13 applies the pattern to the existing 日本人 entry.

Access and limits: Headword or named section inspected on 2026-11-05; not a complete audit of the site.

Exact Notes source line, subject to URL deduplication:

```plain
Memory-aid source: Tokyo University of Foreign Studies Language Modules — Nです／Nではありません. https://www.coelang.tufs.ac.jp/mt/ja/gmod/contents/card/001.html
```

### S13

[Tofugu — 〜てみる](https://www.tofugu.com/japanese-grammar/temiru/)

Supports: Grammar equations, conjugation, and 〜てみたい sections: verb て/で form + みる, trying an action, and wanting to try. Supports WF14/N05 and the distinction in R17.

Access and limits: Authored grammar reference inspected directly. The page is not an official anime source; it provides no quotation attribution or naming history. Mnemonic wording is original editorial text, not a quotation from Tofugu.

Exact Notes source line, subject to URL deduplication:

```plain
Memory-aid source: Tofugu — 〜てみる. https://www.tofugu.com/japanese-grammar/temiru/
```

### S14

[Online Etymology Dictionary — caveat](https://www.etymonline.com/word/caveat)

Supports: Latin caveat, a form of cavere meaning beware/take heed.

Access and limits: Headword or named section inspected on 2026-11-05; not a complete audit of the site.

Exact Notes source line, subject to URL deduplication:

```plain
Memory-aid source: Online Etymology Dictionary — caveat. https://www.etymonline.com/word/caveat
```

### S15

[Online Etymology Dictionary — anonymous](https://www.etymonline.com/word/anonymous)

Supports: Late Latin anonymus from Greek anonymos; an- without + onyma, a dialectal form of onoma, name.

Access and limits: Headword or named section inspected on 2026-11-05; not a complete audit of the site.

Exact Notes source line, subject to URL deduplication:

```plain
Memory-aid source: Online Etymology Dictionary — anonymous. https://www.etymonline.com/word/anonymous
```

### S16

[Online Etymology Dictionary — adjacent](https://www.etymonline.com/word/adjacent)

Supports: Latin adiacens/adiacere, ad + iacere: lying near. Only this origin is used; no blanket restriction on persons or compulsory physical contact is proposed.

Access and limits: Headword or named section inspected on 2026-11-05; not a complete audit of the site.

Exact Notes source line, subject to URL deduplication:

```plain
Memory-aid source: Online Etymology Dictionary — adjacent. https://www.etymonline.com/word/adjacent
```

### S17

[Online Etymology Dictionary — mutable](https://www.etymonline.com/word/mutable)

Supports: Latin mutabilis, from mutare, change. No invented direct modern mutate + -able derivation.

Access and limits: Headword or named section inspected on 2026-11-05; not a complete audit of the site.

Exact Notes source line, subject to URL deduplication:

```plain
Memory-aid source: Online Etymology Dictionary — mutable. https://www.etymonline.com/word/mutable
```

### S18

[Online Etymology Dictionary — audit](https://www.etymonline.com/word/audit)

Supports: Latin auditus/audire and historical oral examination of accounts. This history does not make modern code audits necessarily oral.

Access and limits: Headword or named section inspected on 2026-11-05; not a complete audit of the site.

Exact Notes source line, subject to URL deduplication:

```plain
Memory-aid source: Online Etymology Dictionary — audit. https://www.etymonline.com/word/audit
```

### S19

[Online Etymology Dictionary — audio](https://www.etymonline.com/word/audio)

Supports: English noun abstracted from audio-, ultimately Latin audire. Avoids equating the entire English noun directly with a Latin first-person verb form.

Access and limits: Headword or named section inspected on 2026-11-05; not a complete audit of the site.

Exact Notes source line, subject to URL deduplication:

```plain
Memory-aid source: Online Etymology Dictionary — audio. https://www.etymonline.com/word/audio
```

### S20

[Online Etymology Dictionary — regression](https://www.etymonline.com/word/regression)

Supports: Latin regressio/regredi: going back; software interpretation is explicitly a present-day application, not a Latin technical sense.

Access and limits: Headword or named section inspected on 2026-11-05; not a complete audit of the site.

Exact Notes source line, subject to URL deduplication:

```plain
Memory-aid source: Online Etymology Dictionary — regression. https://www.etymonline.com/word/regression
```

### S21

[Online Etymology Dictionary — inspect](https://www.etymonline.com/word/inspect)

Supports: Latin inspectus/inspicere, in- + specere: look into/examine.

Access and limits: Headword or named section inspected on 2026-11-05; not a complete audit of the site.

Exact Notes source line, subject to URL deduplication:

```plain
Memory-aid source: Online Etymology Dictionary — inspect. https://www.etymonline.com/word/inspect
```

### S22

[Merriam-Webster.com Dictionary — outstanding](https://www.merriam-webster.com/dictionary/outstanding)

Supports: Editorial adjective senses include unresolved and distinguished; supports selecting the unfinished-work sense in N01.

Access and limits: Headword or named section inspected on 2026-11-05; not a complete audit of the site.

Exact Notes source line, subject to URL deduplication:

```plain
Memory-aid source: Merriam-Webster.com Dictionary — outstanding. https://www.merriam-webster.com/dictionary/outstanding
```

### S23

[TypeScript Handbook — More on Functions: Function Overloads](https://www.typescriptlang.org/docs/handbook/2/functions.html#function-overloads)

Supports: Different call signatures for a function; not excess runtime load. The handbook illustrates TypeScript, without asserting identical overload dispatch rules for every language.

Access and limits: Headword or named section inspected on 2026-11-05; not a complete audit of the site.

Exact Notes source line, subject to URL deduplication:

```plain
Memory-aid source: TypeScript Handbook — More on Functions: Function Overloads. https://www.typescriptlang.org/docs/handbook/2/functions.html#function-overloads
```

### S24

[Merriam-Webster.com Dictionary — benchmark](https://www.merriam-webster.com/dictionary/benchmark)

Supports: A comparison standard or standardized test, including computer performance. N03 distinguishes only the two current LexForge senses, not every use of baseline/benchmark.

Access and limits: Headword or named section inspected on 2026-11-05; not a complete audit of the site.

Exact Notes source line, subject to URL deduplication:

```plain
Memory-aid source: Merriam-Webster.com Dictionary — benchmark. https://www.merriam-webster.com/dictionary/benchmark
```

### S25

[MDN Web Docs — String: UTF-16 characters, Unicode code points, and grapheme clusters](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/String#utf-16_characters_unicode_code_points_and_grapheme_clusters)

Supports: JavaScript string indices count UTF-16 code units; supplementary code points use surrogate pairs. N04 uses an explicitly counted example, not a universal bytes-to-characters conversion.

Access and limits: Headword or named section inspected on 2026-11-05; not a complete audit of the site.

Exact Notes source line, subject to URL deduplication:

```plain
Memory-aid source: MDN Web Docs — String: UTF-16 characters, Unicode code points, and grapheme clusters. https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/String#utf-16_characters_unicode_code_points_and_grapheme_clusters
```

### S26

[MDN Web Docs — TextEncoder](https://developer.mozilla.org/en-US/docs/Web/API/TextEncoder)

Supports: TextEncoder produces UTF-8 bytes. N04 byte/unit counts were independently checked with local encoding calculations; they are not quoted from MDN.

Access and limits: Headword or named section inspected on 2026-11-05; not a complete audit of the site.

Exact Notes source line, subject to URL deduplication:

```plain
Memory-aid source: MDN Web Docs — TextEncoder. https://developer.mozilla.org/en-US/docs/Web/API/TextEncoder
```

### S27

[Kanae Nakamine, Tofugu — Is あなた Polite or Rude?](https://www.tofugu.com/japanese/anata/)

Supports: Usage varies with context and relationship; あなた can sound distancing when a name or role would be expected. Supports N06 without an always-rude or always-polite rule.

Access and limits: Authored usage analysis inspected directly, including contextual examples. N06 is original editorial guidance; it is not a universal honorific ranking or a quotation.

Exact Notes source line, subject to URL deduplication:

```plain
Memory-aid source: Kanae Nakamine, Tofugu — Is あなた Polite or Rude?. https://www.tofugu.com/japanese/anata/
```

### S28

[Merriam-Webster.com Dictionary — overload](https://www.merriam-webster.com/dictionary/overload)

Supports: The ordinary excessive-load sense contrasted with TypeScript function overloads in N02. The dictionary does not supply the programming mechanism; S23 supports that part.

Access and limits: Editorial noun/verb definitions inspected on 2026-11-05; automatically aggregated examples are not evidence. The comparison wording is original, not a quotation.

Exact Notes source line, subject to URL deduplication:

```plain
Memory-aid source: Merriam-Webster.com Dictionary — overload. https://www.merriam-webster.com/dictionary/overload
```

### S29

[小学館《デジタル大辞泉》— 外国人, via Kotobank](https://kotobank.jp/word/外国人-42384)

Supports: Headword 外国人 and reading がいこくじん. The ordinary sense supplies the contrast with a student's educational status. The page's unrelated legal/history discussions are not part of N07 or R26.

Access and limits: The dictionary headword and definition were opened and inspected on 2026-11-05. The listening sequence is original guidance, not a dictionary quotation or a claim that audio was reviewed.

Exact Notes source line, subject to URL deduplication:

```plain
Memory-aid source: 小学館《デジタル大辞泉》— 外国人 (Kotobank). https://kotobank.jp/word/外国人-42384
```

### S30

[Langoal — 大学生](https://langoal.com/ja/vocabulary/word/daigakusei.html)

Supports: The teacher-facing vocabulary page explicitly gives 大学生, だいがくせい, the university-student meaning, and the component readings 大 (だい), 学 (がく), 生 (せい). N07 uses that reading and meaning; 大学 + 生 is a modern learning analysis, not a historical derivation claim.

Access and limits: Vocabulary heading, reading, learner gloss, and example were opened and inspected on 2026-11-05. The page has a malformed romaji FAQ; that line is not used. Its level classification is not adopted, and no premium recording was downloaded or listened to. The attempted bare Kotobank 大学生 page was inaccessible and supplies no evidence.

Exact Notes source line, subject to URL deduplication:

```plain
Memory-aid source: Langoal — 大学生. https://langoal.com/ja/vocabulary/word/daigakusei.html
```

### Evidence not used

The following attempts are recorded as limitations, not supporting citations: Kotobank bare headwords for ドイツ, イギリス, and バーガー; Online Etymology Dictionary initializer; Merriam-Webster allowlist; and a standalone Tofugu 〜たい page did not yield inspected supporting text. Search-result summaries were not promoted to verified facts. A guessed combined Irodori grammar resource was not used. S01 provides directly inspected allowlist usage; S13 supplies the inspected desire/trying distinction. No unavailable headword is used as the source of a proposed historical claim.

Source links point to inspected web pages, but remote headings and content may change. No claim is made that the entire collection’s historical provenance links were rechecked, that a website’s automated examples are dictionary editorial judgments, or that archived user clarifications are independent entity verification. Existing sources remain in place.

## Deferred investigations

| Code | Stable IDs                                                                                                                                                                                                                                                                   | Decision / uncertainty                                                                                                                                                                                                                                                                          |
| :--- | :--------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | :---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| D01  | [doitsu](../entries/doitsu.yml), [igirisu](../entries/igirisu.yml), [baagaa](../entries/baagaa.yml)                                                                                                                                                                          | Additional loanword etymology deferred. The Japanese dictionary access attempts did not produce inspected headword evidence. Do not infer Dutch, Portuguese, or a particular shortening route from kana resemblance. Existing food/context explanations stay unchanged.                         |
| D02  | [initializer-tech](../entries/initializer-tech.yml)                                                                                                                                                                                                                          | A useful initialize/initializer family is plausible, but the attempted dedicated historical headword was unavailable. Defer a new field until a directly inspected lexical source and the initialization-expression sense are tied together; the current Meaning and Notes already cover usage. |
| D03  | [compile-tech](../entries/compile-tech.yml), [legacy-tech](../entries/legacy-tech.yml)                                                                                                                                                                                       | Additional historical origin not selected: compile has a qualified reconstruction in the dictionary; legacy has a longer historical sense chain. Neither is needed to explain the current technical meaning. WF03 does not resolve or replace those histories.                                  |
| D04  | [ee-kee-shii](../entries/ee-kee-shii.yml), [koube-byouin](../entries/koube-byouin.yml), [sakura-daigaku](../entries/sakura-daigaku.yml), [fuji-daigaku](../entries/fuji-daigaku.yml), [pawaa-denki](../entries/pawaa-denki.yml), [burajiru-eaa](../entries/burajiru-eaa.yml) | Naming origins and real-world identities remain unverified. Retain the user-provided clarifications and any existing component analysis. Do not expand initials, identify a specific organization, or infer official naming intent.                                                             |
| D05  | [morimoto](../entries/morimoto.yml), [edogawa-konan](../entries/edogawa-konan.yml), [tsukino-usagi](../entries/tsukino-usagi.yml)                                                                                                                                            | Do not invent personal-name origins from character meanings or a real-world namesake. Existing readings and work attribution remain; additional naming history would require specific evidence.                                                                                                 |
| D06  | [hermes-agent-ai](../entries/hermes-agent-ai.yml), [muse-ai](../entries/muse-ai.yml), [draconification-zelda](../entries/draconification-zelda.yml)                                                                                                                          | Do not infer official product naming intent or a fictional transformation mechanism from a word or mythological resemblance. Retain existing sourced text; no new historical/mechanistic claim is proposed.                                                                                     |

These **18 IDs** have explicitly deferred additional investigations. An ID may retain an existing aid while its additional history is deferred. No deferred statement is written into entries. No evidence-backed correction to existing text is proposed in this batch.

## Existing aid inventory

Retain all listed values and their existing Notes/provenance byte-for-byte. E = etymology, W = word_formation, N = mnemonic. This is an inventory of **251 populated aid values on 238 entries**, not approval to rewrite them or a claim that every source was reverified. Other entries have all three fields absent; there were no explicit empty/null aid values at the baseline. The ledger also records reverse relationship displays.

| Stable ID                                                                                 | Word                         | Retain populated fields |
| :---------------------------------------------------------------------------------------- | :--------------------------- | :---------------------- |
| [accessor-tech](../entries/accessor-tech.yml)                                             | accessor                     | W                       |
| [actionable-tech](../entries/actionable-tech.yml)                                         | actionable                   | W                       |
| [adapter-tech](../entries/adapter-tech.yml)                                               | adapter                      | W                       |
| [airborne-combat](../entries/airborne-combat.yml)                                         | airborne                     | W                       |
| [aitemu](../entries/aitemu.yml)                                                           | アイテム                     | E                       |
| [anime](../entries/anime.yml)                                                             | アニメ                       | W                       |
| [anime-guzzu](../entries/anime-guzzu.yml)                                                 | アニメグッズ                 | W                       |
| [aquatic-combat](../entries/aquatic-combat.yml)                                           | aquatic                      | E                       |
| [armored-combat](../entries/armored-combat.yml)                                           | armored                      | W                       |
| [artifact-tech](../entries/artifact.yml)                                                  | artifact                     | E, N                    |
| [artifact-workflow](../entries/artifact-workflow.yml)                                     | artifact                     | N                       |
| [asynchronous-tech](../entries/asynchronous-tech.yml)                                     | asynchronous                 | W                       |
| [attack-power-combat](../entries/attack-power-combat.yml)                                 | attack power                 | W                       |
| [autobuild-zelda](../entries/autobuild-zelda.yml)                                         | Autobuild                    | W                       |
| [backtracking-algorithm](../entries/backtracking-algorithm.yml)                           | backtracking                 | W                       |
| [bass-food](../entries/bass-food.yml)                                                     | bass                         | N                       |
| [beam-emitter-zelda](../entries/beam-emitter-zelda.yml)                                   | Beam Emitter                 | W                       |
| [block-formatting-context-css](../entries/block-formatting-context-css.yml)               | block formatting context     | N                       |
| [blocker-workflow](../entries/blocker-workflow.yml)                                       | blocker                      | W                       |
| [boku-wa-shinsekai-no-kami-ni-naru](../entries/boku-wa-shinsekai-no-kami-ni-naru.yml)     | 僕は新世界の神になる         | W                       |
| [breadth-first-tech](../entries/breadth-first-tech.yml)                                   | breadth-first                | W                       |
| [broadsword-weapon](../entries/broadsword-weapon.yml)                                     | broadsword                   | W                       |
| [burajiru-eaa](../entries/burajiru-eaa.yml)                                               | ブラジルエアー               | W                       |
| [cache-tech](../entries/cache-tech.yml)                                                   | cache                        | E                       |
| [callback-tech](../entries/callback-tech.yml)                                             | callback                     | W                       |
| [candidate-tech](../entries/candidate-tech.yml)                                           | candidate                    | E                       |
| [cavern-architecture](../entries/cavern-architecture.yml)                                 | cavern                       | N                       |
| [chasm-terrain](../entries/chasm-terrain.yml)                                             | chasm                        | N                       |
| [chisou](../entries/chisou.yml)                                                           | 馳走                         | E                       |
| [chuugokujin](../entries/chuugokujin.yml)                                                 | 中国人                       | W                       |
| [cleanup-tech](../entries/cleanup-tech.yml)                                               | cleanup                      | W                       |
| [closure-tech](../entries/closure-tech.yml)                                               | closure                      | N                       |
| [construct-head-zelda](../entries/construct-head-zelda.yml)                               | Construct Head               | W                       |
| [cooking-pot-cooking](../entries/cooking-pot-cooking.yml)                                 | cooking pot                  | W                       |
| [copilot-ai](../entries/copilot-ai.yml)                                                   | Copilot                      | W                       |
| [correctness-tech](../entries/correctness-tech.yml)                                       | correctness                  | W                       |
| [courageous-general](../entries/courageous-general.yml)                                   | courageous                   | E, W                    |
| [critical-hit-combat](../entries/critical-hit-combat.yml)                                 | critical hit                 | W                       |
| [daijoubu](../entries/daijoubu.yml)                                                       | 大丈夫                       | N                       |
| [decode-tech](../entries/decode-tech.yml)                                                 | decode                       | N                       |
| [deepseek-ai](../entries/deepseek-ai.yml)                                                 | DeepSeek                     | W                       |
| [denwa-bangou](../entries/denwa-bangou.yml)                                               | 電話番号                     | W                       |
| [deprecated-tech](../entries/deprecated.yml)                                              | deprecated                   | N                       |
| [descendant-general](../entries/descendant-general.yml)                                   | descendant                   | W                       |
| [device-dispenser-zelda](../entries/device-dispenser-zelda.yml)                           | Device Dispenser             | W                       |
| [disclosure-workflow](../entries/disclosure-workflow.yml)                                 | disclosure                   | N                       |
| [dispenser-engineering](../entries/dispenser-engineering.yml)                             | dispenser                    | W                       |
| [dots-ai](../entries/dots-ai.yml)                                                         | dots                         | W                       |
| [durability-equipment](../entries/durability-equipment.yml)                               | durability                   | W                       |
| [eaa](../entries/eaa.yml)                                                                 | エアー                       | E                       |
| [elixir-cooking](../entries/elixir-cooking.yml)                                           | elixir                       | E                       |
| [emitter-engineering](../entries/emitter-engineering.yml)                                 | emitter                      | W                       |
| [encode-tech](../entries/encode-tech.yml)                                                 | encode                       | N                       |
| [enhancement-general](../entries/enhancement-general.yml)                                 | enhancement                  | W                       |
| [enthusiastic-general](../entries/enthusiastic-general.yml)                               | enthusiastic                 | E                       |
| [escape-parity-tech](../entries/escape-parity-tech.yml)                                   | escape parity                | E                       |
| [escape-tech](../entries/escape-tech.yml)                                                 | escape                       | E                       |
| [escaped-backslash-tech](../entries/escaped-backslash-tech.yml)                           | escaped backslash            | N                       |
| [etymology-language](../entries/etymology-language.yml)                                   | etymology                    | E                       |
| [fallback-tech](../entries/fallback.yml)                                                  | fallback                     | W                       |
| [fearless-general](../entries/fearless-general.yml)                                       | fearless                     | W                       |
| [field-presence-tech](../entries/field-presence-tech.yml)                                 | field presence               | N                       |
| [fisshu](../entries/fisshu.yml)                                                           | フィッシュ                   | E                       |
| [fisshu-baagaa](../entries/fisshu-baagaa.yml)                                             | フィッシュバーガー           | W                       |
| [flame-emitter-zelda](../entries/flame-emitter-zelda.yml)                                 | Flame Emitter                | W                       |
| [floating-point-precision-javascript](../entries/floating-point-precision-javascript.yml) | floating-point precision     | N                       |
| [foothill-terrain](../entries/foothill-terrain.yml)                                       | foothill                     | W                       |
| [frost-emitter-zelda](../entries/frost-emitter-zelda.yml)                                 | Frost Emitter                | W                       |
| [fuuin-kaijo-reriizu](../entries/fuuin-kaijo-reriizu.yml)                                 | 封印解除（レリーズ）！       | W, N                    |
| [gaikokujin](../entries/gaikokujin.yml)                                                   | 外国人                       | W                       |
| [gentei-guzzu](../entries/gentei-guzzu.yml)                                               | 限定グッズ                   | W                       |
| [gimu-toka-janai-yo](../entries/gimu-toka-janai-yo.yml)                                   | 義務とかじゃないよ           | W                       |
| [gochisou](../entries/gochisou.yml)                                                       | ごちそう                     | E                       |
| [grill-cooking](../entries/grill-cooking.yml)                                             | grill                        | N                       |
| [guzzu](../entries/guzzu.yml)                                                             | グッズ                       | E                       |
| [hanasemasen](../entries/hanasemasen.yml)                                                 | 話せません                   | W, N                    |
| [hanasu](../entries/hanasu.yml)                                                           | 話す                         | N                       |
| [handoff-workflow](../entries/handoff-workflow.yml)                                       | handoff                      | W                       |
| [headband-clothing](../entries/headband-clothing.yml)                                     | headband                     | W                       |
| [headdress-clothing](../entries/headdress-clothing.yml)                                   | headdress                    | W                       |
| [heir-general](../entries/heir-general.yml)                                               | heir                         | N                       |
| [highlands-terrain](../entries/highlands-terrain.yml)                                     | highlands                    | W                       |
| [hikari-ni-nareru-nda](../entries/hikari-ni-nareru-nda.yml)                               | 光になれるんだ               | W                       |
| [hito-o-norowaba-ana-futatsu](../entries/hito-o-norowaba-ana-futatsu.yml)                 | 人を呪わば穴二つ             | N                       |
| [hokage-ni-naru-made](../entries/hokage-ni-naru-made.yml)                                 | 火影になるまで               | W                       |
| [homing-engineering](../entries/homing-engineering.yml)                                   | homing                       | W                       |
| [hostile-combat](../entries/hostile-combat.yml)                                           | hostile                      | E, N                    |
| [hover-stone-zelda](../entries/hover-stone-zelda.yml)                                     | Hover Stone                  | W                       |
| [http-caching-web](../entries/http-caching-web.yml)                                       | HTTP caching                 | N                       |
| [hydrant-engineering](../entries/hydrant-engineering.yml)                                 | hydrant                      | E                       |
| [idiomatic-tech](../entries/idiomatic.yml)                                                | idiomatic                    | E                       |
| [imply-general](../entries/imply-general.yml)                                             | imply                        | N                       |
| [inaccessible-tech](../entries/inaccessible-tech.yml)                                     | inaccessible                 | W                       |
| [independently-tech](../entries/independently-tech.yml)                                   | independently                | W                       |
| [infer-tech](../entries/infer-tech.yml)                                                   | infer                        | N                       |
| [integrity-general](../entries/integrity-general.yml)                                     | integrity                    | E                       |
| [invaluable-general](../entries/invaluable-general.yml)                                   | invaluable                   | N                       |
| [inventory-ui](../entries/inventory-ui.yml)                                               | inventory                    | N                       |
| [jibun-jishin-no-chikara-de](../entries/jibun-jishin-no-chikara-de.yml)                   | 自分自身の力で               | W                       |
| [jin-nationality-suffix](../entries/jin-nationality-suffix.yml)                           | ～人                         | N                       |
| [kaishain](../entries/kaishain.yml)                                                       | 会社員                       | W, N                    |
| [kenkyuusha](../entries/kenkyuusha.yml)                                                   | 研究者                       | W                       |
| [kimi-ga-yobu-namae-yume-no-tsuzuki](../entries/kimi-ga-yobu-namae-yume-no-tsuzuki.yml)   | きみが呼ぶなまえ～夢のつづき | W                       |
| [kingdom-general](../entries/kingdom-general.yml)                                         | kingdom                      | W                       |
| [koneko](../entries/koneko.yml)                                                           | 子猫                         | W                       |
| [kosumo-saint-seiya](../entries/kosumo-saint-seiya.yml)                                   | 小宇宙                       | N                       |
| [kotoba-wa-magenee](../entries/kotoba-wa-magenee.yml)                                     | 言葉は曲げねェ               | W                       |
| [kyoushi](../entries/kyoushi.yml)                                                         | 教師                         | N                       |
| [least-recently-used-cache](../entries/least-recently-used-cache.yml)                     | least recently used          | N                       |
| [lifecycle-tech](../entries/lifecycle-tech.yml)                                           | lifecycle                    | W                       |
| [lightroot-zelda](../entries/lightroot-zelda.yml)                                         | Lightroot                    | W                       |
| [literal-programming](../entries/literal-programming.yml)                                 | literal                      | N                       |
| [literal-text](../entries/literal-text.yml)                                               | literal                      | N                       |
| [lockfile-tech](../entries/lockfile-tech.yml)                                             | lockfile                     | W                       |
| [maintainable-tech](../entries/maintainable-tech.yml)                                     | maintainable                 | W                       |
| [malevolent-general](../entries/malevolent-general.yml)                                   | malevolent                   | E                       |
| [malformed-tech](../entries/malformed-tech.yml)                                           | malformed                    | W                       |
| [melee-combat](../entries/melee-combat.yml)                                               | melee                        | E                       |
| [microtask-javascript](../entries/microtask-javascript.yml)                               | microtask                    | W                       |
| [miner-general](../entries/miner-general.yml)                                             | miner                        | W                       |
| [misalign-tech](../entries/misalign-tech.yml)                                             | misalign                     | W                       |
| [mnemonic-memory](../entries/mnemonic-memory.yml)                                         | mnemonic                     | E, N                    |
| [mount-tech](../entries/mount-tech.yml)                                                   | mount                        | N                       |
| [multibyte-tech](../entries/multibyte-tech.yml)                                           | multibyte                    | W                       |
| [mysterious-general](../entries/mysterious-general.yml)                                   | mysterious                   | W                       |
| [neko-chan](../entries/neko-chan.yml)                                                     | 猫ちゃん                     | W                       |
| [ni-nareru-grammar](../entries/ni-nareru-grammar.yml)                                     | ～になれる                   | W                       |
| [nihongo](../entries/nihongo.yml)                                                         | 日本語                       | W                       |
| [nihonjin](../entries/nihonjin.yml)                                                       | 日本人                       | W                       |
| [no-ka-grammar](../entries/no-ka-grammar.yml)                                             | ～のか                       | N                       |
| [nonthrowing-tech](../entries/nonthrowing-tech.yml)                                       | nonthrowing                  | W                       |
| [omachigae](../entries/omachigae.yml)                                                     | お間違え                     | W                       |
| [onamae-wa](../entries/onamae-wa.yml)                                                     | お名前は                     | N                       |
| [one-handed-combat](../entries/one-handed-combat.yml)                                     | one-handed                   | W                       |
| [openclaw-ai](../entries/openclaw-ai.yml)                                                 | OpenClaw                     | W                       |
| [ore-wa-ningen-dakara](../entries/ore-wa-ningen-dakara.yml)                               | 俺は人間だから               | W                       |
| [osagashi](../entries/osagashi.yml)                                                       | お探し                       | W                       |
| [own-property-javascript](../entries/own-property-javascript.yml)                         | own property                 | N                       |
| [paella-cooking](../entries/paella-cooking.yml)                                           | paella                       | E                       |
| [parameter-tech](../entries/parameter-tech.yml)                                           | parameter                    | N                       |
| [pawaa](../entries/pawaa.yml)                                                             | パワー                       | E                       |
| [pawaa-denki](../entries/pawaa-denki.yml)                                                 | パワー電気                   | W                       |
| [powerful-general](../entries/powerful-general.yml)                                       | powerful                     | W                       |
| [pre-existing-tech](../entries/pre-existing-tech.yml)                                     | pre-existing                 | W                       |
| [prefilter-tech](../entries/prefilter-tech.yml)                                           | prefilter                    | W                       |
| [preflight-request-web](../entries/preflight-request-web.yml)                             | preflight request            | N                       |
| [preserve-tech](../entries/preserve-tech.yml)                                             | preserve                     | N                       |
| [pristine-general](../entries/pristine-general.yml)                                       | pristine                     | N                       |
| [promise-all-javascript](../entries/promise-all-javascript.yml)                           | Promise.all                  | N                       |
| [pronunciation-general](../entries/pronunciation.yml)                                     | pronunciation                | N                       |
| [quicksort-algorithm](../entries/quicksort-algorithm.yml)                                 | quicksort                    | W                       |
| [re-entry-tech](../entries/re-entry-tech.yml)                                             | re-entry                     | W                       |
| [recharge-general](../entries/recharge-general.yml)                                       | recharge                     | W                       |
| [recipe-cooking](../entries/recipe-cooking.yml)                                           | recipe                       | E                       |
| [reckless-general](../entries/reckless-general.yml)                                       | reckless                     | N                       |
| [refinery-architecture](../entries/refinery-architecture.yml)                             | refinery                     | W                       |
| [regional-general](../entries/regional-general.yml)                                       | regional                     | W                       |
| [register-tech](../entries/register-tech.yml)                                             | register                     | N                       |
| [reproduce-tech](../entries/reproduce-tech.yml)                                           | reproduce                    | N                       |
| [reproducible-tech](../entries/reproducible-tech.yml)                                     | reproducible                 | N                       |
| [resilient-combat](../entries/resilient-combat.yml)                                       | resilient                    | E, N                    |
| [resolve-tech](../entries/resolve-tech.yml)                                               | resolve                      | N                       |
| [resolve-workflow](../entries/resolve-workflow.yml)                                       | resolve                      | N                       |
| [retain-tech](../entries/retain-tech.yml)                                                 | retain                       | N                       |
| [revision-general](../entries/revision-general.yml)                                       | revision                     | N                       |
| [revision-git](../entries/revision-git.yml)                                               | revision                     | N                       |
| [rupee-game](../entries/rupee-game.yml)                                                   | rupee                        | E                       |
| [ruthless-general](../entries/ruthless-general.yml)                                       | ruthless                     | W                       |
| [sa-ending-grammar](../entries/sa-ending-grammar.yml)                                     | ～さ                         | N                       |
| [sacrifice-story](../entries/sacrifice-story.yml)                                         | sacrifice                    | E                       |
| [sakura-no-ki](../entries/sakura-no-ki.yml)                                               | 桜の木                       | W                       |
| [san-name-suffix](../entries/san-name-suffix.yml)                                         | ～さん                       | N                       |
| [saute-cooking](../entries/saute-cooking.yml)                                             | sauté                        | E                       |
| [scope-lexical](../entries/scope-lexical.yml)                                             | scope                        | N                       |
| [scope-workflow](../entries/scope-workflow.yml)                                           | scope                        | N                       |
| [sear-cooking](../entries/sear-cooking.yml)                                               | sear                         | N                       |
| [selfless-general](../entries/selfless-general.yml)                                       | selfless                     | W                       |
| [sensei](../entries/sensei.yml)                                                           | 先生                         | N                       |
| [settlement-architecture](../entries/settlement-architecture.yml)                         | settlement                   | W                       |
| [setup-tech](../entries/setup-tech.yml)                                                   | setup                        | W                       |
| [shain](../entries/shain.yml)                                                             | 社員                         | N                       |
| [shinde-miru](../entries/shinde-miru.yml)                                                 | 死んでみる                   | W                       |
| [shinsekai](../entries/shinsekai.yml)                                                     | 新世界                       | W                       |
| [shock-emitter-zelda](../entries/shock-emitter-zelda.yml)                                 | Shock Emitter                | W                       |
| [simmer-cooking](../entries/simmer-cooking.yml)                                           | simmer                       | N                       |
| [skilled-general](../entries/skilled-general.yml)                                         | skilled                      | W                       |
| [sky-island-zelda](../entries/sky-island-zelda.yml)                                       | sky island                   | W                       |
| [skyview-tower-zelda](../entries/skyview-tower-zelda.yml)                                 | Skyview Tower                | W                       |
| [soundness-type-system](../entries/soundness-type-system.yml)                             | soundness                    | N                       |
| [spring-water](../entries/spring-water.yml)                                               | spring                       | N                       |
| [stabilizer-engineering](../entries/stabilizer-engineering.yml)                           | stabilizer                   | W                       |
| [staged-git](../entries/staged-git.yml)                                                   | staged                       | N                       |
| [staircase-architecture](../entries/staircase-architecture.yml)                           | staircase                    | W                       |
| [stamina-general](../entries/stamina-general.yml)                                         | stamina                      | E, N                    |
| [steam-cooking](../entries/steam-cooking.yml)                                             | steam                        | N                       |
| [steering-stick-zelda](../entries/steering-stick-zelda.yml)                               | Steering Stick               | W                       |
| [sturdy-general](../entries/sturdy-general.yml)                                           | sturdy                       | N                       |
| [sukueru](../entries/sukueru.yml)                                                         | 救える                       | W                       |
| [summit-terrain](../entries/summit-terrain.yml)                                           | summit                       | N                       |
| [susceptible-combat](../entries/susceptible-combat.yml)                                   | susceptible                  | N                       |
| [symlink-tech](../entries/symlink-tech.yml)                                               | symlink                      | W                       |
| [syntax-tech](../entries/syntax-tech.yml)                                                 | syntax                       | E                       |
| [tantei-sa](../entries/tantei-sa.yml)                                                     | 探偵さ！                     | W                       |
| [taoru](../entries/taoru.yml)                                                             | タオル                       | E                       |
| [tear-general](../entries/tear-general.yml)                                               | tear                         | N                       |
| [to-moushimasu-grammar](../entries/to-moushimasu-grammar.yml)                             | ～と申します                 | W                       |
| [to-nanoru-grammar](../entries/to-nanoru-grammar.yml)                                     | ～と名乗る                   | W                       |
| [toolchain-tech](../entries/toolchain-tech.yml)                                           | toolchain                    | W                       |
| [tradeoff-general](../entries/tradeoff-general.yml)                                       | tradeoff                     | W                       |
| [traveler-general](../entries/traveler-general.yml)                                       | traveler                     | W                       |
| [trident-weapon](../entries/trident-weapon.yml)                                           | trident                      | E                       |
| [tsuki-ni-kawatte-oshioki-yo](../entries/tsuki-ni-kawatte-oshioki-yo.yml)                 | 月にかわっておしおきよ！     | W                       |
| [two-handed-combat](../entries/two-handed-combat.yml)                                     | two-handed                   | W                       |
| [typo-general](../entries/typo-general.yml)                                               | typo                         | W                       |
| [ultrahand-zelda](../entries/ultrahand-zelda.yml)                                         | Ultrahand                    | W                       |
| [unclaimed-workflow](../entries/unclaimed-workflow.yml)                                   | unclaimed                    | W                       |
| [undead-combat](../entries/undead-combat.yml)                                             | undead                       | W                       |
| [underground-general](../entries/underground-general.yml)                                 | underground                  | W                       |
| [unfuse-general](../entries/unfuse-general.yml)                                           | unfuse                       | W                       |
| [uninitialized-tech](../entries/uninitialized-tech.yml)                                   | uninitialized                | W                       |
| [unlock-general](../entries/unlock-general.yml)                                           | unlock                       | W                       |
| [unmount-tech](../entries/unmount-tech.yml)                                               | unmount                      | W, N                    |
| [unregister-tech](../entries/unregister-tech.yml)                                         | unregister                   | W, N                    |
| [unstaged-git](../entries/unstaged-git.yml)                                               | unstaged                     | W, N                    |
| [untracked-git](../entries/untracked-git.yml)                                             | untracked                    | W                       |
| [upstream-git](../entries/upstream.yml)                                                   | upstream                     | W                       |
| [virtual-dom-frontend](../entries/virtual-dom-frontend.yml)                               | virtual DOM                  | N                       |
| [vocabulary-general](../entries/vocabulary.yml)                                           | vocabulary                   | E                       |
| [vulnerable-combat](../entries/vulnerable-combat.yml)                                     | vulnerable                   | E, N                    |
| [waistguard-clothing](../entries/waistguard-clothing.yml)                                 | waistguard                   | W                       |
| [waterfall-terrain](../entries/waterfall-terrain.yml)                                     | waterfall                    | W                       |
| [wetlands-terrain](../entries/wetlands-terrain.yml)                                       | wetlands                     | W                       |
| [workload-tech](../entries/workload-tech.yml)                                             | workload                     | W                       |
| [yareru](../entries/yareru.yml)                                                           | やれる                       | W                       |
| [yareru-koto](../entries/yareru-koto.yml)                                                 | やれること                   | W                       |
| [yaritai-dake-da-yo](../entries/yaritai-dake-da-yo.yml)                                   | やりたいだけだよ             | W                       |
| [yomikata](../entries/yomikata.yml)                                                       | 読み方                       | W                       |
| [zettee-shinanee](../entries/zettee-shinanee.yml)                                         | ぜってェ死なねー             | W                       |

## Existing relationship inventory

Preserve these **89 unique direct pairs**. The actual source declaration is shown, so reverse-only display is not mistaken for a missing association. This baseline has 89 declaration members; no reciprocal duplicate declaration is being added or removed.

| Endpoint A                                                                            | Endpoint B                                                                    | Current declaration                                            |
| :------------------------------------------------------------------------------------ | :---------------------------------------------------------------------------- | :------------------------------------------------------------- |
| [accessor-tech](../entries/accessor-tech.yml)                                         | [property-tech](../entries/property-tech.yml)                                 | `property-tech` → `accessor-tech`                              |
| [actual-behavior-workflow](../entries/actual-behavior-workflow.yml)                   | [expected-behavior-workflow](../entries/expected-behavior-workflow.yml)       | `actual-behavior-workflow` → `expected-behavior-workflow`      |
| [aggressive-combat](../entries/aggressive-combat.yml)                                 | [hostile-combat](../entries/hostile-combat.yml)                               | `hostile-combat` → `aggressive-combat`                         |
| [anata](../entries/anata.yml)                                                         | [kimi](../entries/kimi.yml)                                                   | `anata` → `kimi`                                               |
| [ancestor-general](../entries/ancestor-general.yml)                                   | [descendant-general](../entries/descendant-general.yml)                       | `ancestor-general` → `descendant-general`                      |
| [ano-hito](../entries/ano-hito.yml)                                                   | [ano-kata](../entries/ano-kata.yml)                                           | `ano-hito` → `ano-kata`                                        |
| [armor-equipment](../entries/armor-equipment.yml)                                     | [armory-architecture](../entries/armory-architecture.yml)                     | `armor-equipment` → `armory-architecture`                      |
| [artifact-tech](../entries/artifact.yml)                                              | [artifact-workflow](../entries/artifact-workflow.yml)                         | `artifact-tech` → `artifact-workflow`                          |
| [ascend-general](../entries/ascend-general.yml)                                       | [descend-general](../entries/descend-general.yml)                             | `ascend-general` → `descend-general`                           |
| [async-await-javascript](../entries/async-await-javascript.yml)                       | [promise-javascript](../entries/promise-javascript.yml)                       | `async-await-javascript` → `promise-javascript`                |
| [audio-general](../entries/audio-general.yml)                                         | [pronunciation-general](../entries/pronunciation.yml)                         | `audio-general` → `pronunciation-general`                      |
| [boku](../entries/boku.yml)                                                           | [ore](../entries/ore.yml)                                                     | `boku` → `ore`                                                 |
| [brave-general](../entries/brave-general.yml)                                         | [courageous-general](../entries/courageous-general.yml)                       | `courageous-general` → `brave-general`                         |
| [brave-general](../entries/brave-general.yml)                                         | [fearless-general](../entries/fearless-general.yml)                           | `brave-general` → `fearless-general`                           |
| [cache-tech](../entries/cache-tech.yml)                                               | [memoization-programming](../entries/memoization-programming.yml)             | `cache-tech` → `memoization-programming`                       |
| [canyon-terrain](../entries/canyon-terrain.yml)                                       | [chasm-terrain](../entries/chasm-terrain.yml)                                 | `chasm-terrain` → `canyon-terrain`                             |
| [cast-tech](../entries/cast-tech.yml)                                                 | [type-assertion-tech](../entries/type-assertion-tech.yml)                     | `cast-tech` → `type-assertion-tech`                            |
| [cave-architecture](../entries/cave-architecture.yml)                                 | [cavern-architecture](../entries/cavern-architecture.yml)                     | `cave-architecture` → `cavern-architecture`                    |
| [closure-tech](../entries/closure-tech.yml)                                           | [disclosure-workflow](../entries/disclosure-workflow.yml)                     | `closure-tech` → `disclosure-workflow`                         |
| [cold-resistance-combat](../entries/cold-resistance-combat.yml)                       | [heat-resistance-combat](../entries/heat-resistance-combat.yml)               | `cold-resistance-combat` → `heat-resistance-combat`            |
| [compile-tech](../entries/compile-tech.yml)                                           | [compiler-tech](../entries/compiler-tech.yml)                                 | `compile-tech` → `compiler-tech`                               |
| [correctness-tech](../entries/correctness-tech.yml)                                   | [soundness-type-system](../entries/soundness-type-system.yml)                 | `soundness-type-system` → `correctness-tech`                   |
| [courageous-general](../entries/courageous-general.yml)                               | [reckless-general](../entries/reckless-general.yml)                           | `courageous-general` → `reckless-general`                      |
| [cross-origin-resource-sharing-web](../entries/cross-origin-resource-sharing-web.yml) | [same-origin-policy-web](../entries/same-origin-policy-web.yml)               | `cross-origin-resource-sharing-web` → `same-origin-policy-web` |
| [dare](../entries/dare.yml)                                                           | [donata](../entries/donata.yml)                                               | `dare` → `donata`                                              |
| [decayed-general](../entries/decayed-general.yml)                                     | [pristine-general](../entries/pristine-general.yml)                           | `decayed-general` → `pristine-general`                         |
| [decode-tech](../entries/decode-tech.yml)                                             | [encode-tech](../entries/encode-tech.yml)                                     | `encode-tech` → `decode-tech`                                  |
| [deep-copy-javascript](../entries/deep-copy-javascript.yml)                           | [shallow-copy-javascript](../entries/shallow-copy-javascript.yml)             | `deep-copy-javascript` → `shallow-copy-javascript`             |
| [descendant-general](../entries/descendant-general.yml)                               | [heir-general](../entries/heir-general.yml)                                   | `heir-general` → `descendant-general`                          |
| [dewa-arimasen-grammar](../entries/dewa-arimasen-grammar.yml)                         | [janai-grammar](../entries/janai-grammar.yml)                                 | `dewa-arimasen-grammar` → `janai-grammar`                      |
| [durability-equipment](../entries/durability-equipment.yml)                           | [durable-general](../entries/durable-general.yml)                             | `durable-general` → `durability-equipment`                     |
| [durability-equipment](../entries/durability-equipment.yml)                           | [stamina-general](../entries/stamina-general.yml)                             | `stamina-general` → `durability-equipment`                     |
| [emit-tech](../entries/emit-tech.yml)                                                 | [emitter-engineering](../entries/emitter-engineering.yml)                     | `emitter-engineering` → `emit-tech`                            |
| [endure-general](../entries/endure-general.yml)                                       | [energize-general](../entries/energize-general.yml)                           | `endure-general` → `energize-general`                          |
| [event-bubbling-dom](../entries/event-bubbling-dom.yml)                               | [event-capturing-dom](../entries/event-capturing-dom.yml)                     | `event-bubbling-dom` → `event-capturing-dom`                   |
| [event-bubbling-dom](../entries/event-bubbling-dom.yml)                               | [event-delegation-dom](../entries/event-delegation-dom.yml)                   | `event-bubbling-dom` → `event-delegation-dom`                  |
| [fisshu](../entries/fisshu.yml)                                                       | [sakana-general](../entries/sakana-general.yml)                               | `fisshu` → `sakana-general`                                    |
| [flexbox-css](../entries/flexbox-css.yml)                                             | [grid-layout-css](../entries/grid-layout-css.yml)                             | `flexbox-css` → `grid-layout-css`                              |
| [formation-general](../entries/formation-general.yml)                                 | [transformation-tech](../entries/transformation-tech.yml)                     | `formation-general` → `transformation-tech`                    |
| [fuse-general](../entries/fuse-general.yml)                                           | [unfuse-general](../entries/unfuse-general.yml)                               | `fuse-general` → `unfuse-general`                              |
| [gochisousama](../entries/gochisousama.yml)                                           | [gochisousama-deshita](../entries/gochisousama-deshita.yml)                   | `gochisousama` → `gochisousama-deshita`                        |
| [gochisousama](../entries/gochisousama.yml)                                           | [itadakimasu](../entries/itadakimasu.yml)                                     | `gochisousama` → `itadakimasu`                                 |
| [gomen-ne](../entries/gomen-ne.yml)                                                   | [gomennasai](../entries/gomennasai.yml)                                       | `gomen-ne` → `gomennasai`                                      |
| [gomennasai](../entries/gomennasai.yml)                                               | [sumimasen](../entries/sumimasen.yml)                                         | `gomennasai` → `sumimasen`                                     |
| [graceful-degradation-web](../entries/graceful-degradation-web.yml)                   | [progressive-enhancement-web](../entries/progressive-enhancement-web.yml)     | `graceful-degradation-web` → `progressive-enhancement-web`     |
| [grill-cooking](../entries/grill-cooking.yml)                                         | [sear-cooking](../entries/sear-cooking.yml)                                   | `grill-cooking` → `sear-cooking`                               |
| [hanasemasen](../entries/hanasemasen.yml)                                             | [hanasu](../entries/hanasu.yml)                                               | `hanasu` → `hanasemasen`                                       |
| [harness-ai](../entries/harness-ai.yml)                                               | [harness-tech](../entries/harness-tech.yml)                                   | `harness-ai` → `harness-tech`                                  |
| [headband-clothing](../entries/headband-clothing.yml)                                 | [headdress-clothing](../entries/headdress-clothing.yml)                       | `headband-clothing` → `headdress-clothing`                     |
| [helm-clothing](../entries/helm-clothing.yml)                                         | [helmet-clothing](../entries/helmet-clothing.yml)                             | `helm-clothing` → `helmet-clothing`                            |
| [hoisting-javascript](../entries/hoisting-javascript.yml)                             | [temporal-dead-zone-javascript](../entries/temporal-dead-zone-javascript.yml) | `hoisting-javascript` → `temporal-dead-zone-javascript`        |
| [imply-general](../entries/imply-general.yml)                                         | [infer-tech](../entries/infer-tech.yml)                                       | `infer-tech` → `imply-general`                                 |
| [ingredient-cooking](../entries/ingredient-cooking.yml)                               | [recipe-cooking](../entries/recipe-cooking.yml)                               | `ingredient-cooking` → `recipe-cooking`                        |
| [inventory-ui](../entries/inventory-ui.yml)                                           | [pouch-ui](../entries/pouch-ui.yml)                                           | `inventory-ui` → `pouch-ui`                                    |
| [kaishain](../entries/kaishain.yml)                                                   | [shain](../entries/shain.yml)                                                 | `kaishain` → `shain`                                           |
| [kara-kimashita-origin](../entries/kara-kimashita-origin.yml)                         | [kara-reason-grammar](../entries/kara-reason-grammar.yml)                     | `kara-kimashita-origin` → `kara-reason-grammar`                |
| [koneko](../entries/koneko.yml)                                                       | [neko-chan](../entries/neko-chan.yml)                                         | `koneko` → `neko-chan`                                         |
| [kyoushi](../entries/kyoushi.yml)                                                     | [sensei](../entries/sensei.yml)                                               | `kyoushi` → `sensei`                                           |
| [layout-browser-rendering](../entries/layout-browser-rendering.yml)                   | [paint-browser-rendering](../entries/paint-browser-rendering.yml)             | `layout-browser-rendering` → `paint-browser-rendering`         |
| [linear-tech](../entries/linear-tech.yml)                                             | [quadratic-tech](../entries/quadratic-tech.yml)                               | `quadratic-tech` → `linear-tech`                               |
| [literal-programming](../entries/literal-programming.yml)                             | [literal-text](../entries/literal-text.yml)                                   | `literal-programming` → `literal-text`                         |
| [lone-surrogate-tech](../entries/lone-surrogate-tech.yml)                             | [surrogate-pair-tech](../entries/surrogate-pair-tech.yml)                     | `lone-surrogate-tech` → `surrogate-pair-tech`                  |
| [machigaeru](../entries/machigaeru.yml)                                               | [machigai](../entries/machigai.yml)                                           | `machigaeru` → `machigai`                                      |
| [mount-tech](../entries/mount-tech.yml)                                               | [unmount-tech](../entries/unmount-tech.yml)                                   | `mount-tech` → `unmount-tech`                                  |
| [moushimasu](../entries/moushimasu.yml)                                               | [nanoru](../entries/nanoru.yml)                                               | `moushimasu` → `nanoru`                                        |
| [nansai](../entries/nansai.yml)                                                       | [oikutsu](../entries/oikutsu.yml)                                             | `nansai` → `oikutsu`                                           |
| [ni-kawatte-grammar](../entries/ni-kawatte-grammar.yml)                               | [ni-naru-grammar](../entries/ni-naru-grammar.yml)                             | `ni-kawatte-grammar` → `ni-naru-grammar`                       |
| [ni-nareru-grammar](../entries/ni-nareru-grammar.yml)                                 | [ni-naru-grammar](../entries/ni-naru-grammar.yml)                             | `ni-nareru-grammar` → `ni-naru-grammar`                        |
| [nihongo](../entries/nihongo.yml)                                                     | [nihonjin](../entries/nihonjin.yml)                                           | `nihongo` → `nihonjin`                                         |
| [one-way-data-flow-frontend](../entries/one-way-data-flow-frontend.yml)               | [two-way-binding-frontend](../entries/two-way-binding-frontend.yml)           | `one-way-data-flow-frontend` → `two-way-binding-frontend`      |
| [osagashi](../entries/osagashi.yml)                                                   | [sagasu](../entries/sagasu.yml)                                               | `osagashi` → `sagasu`                                          |
| [ownership-token-workflow](../entries/ownership-token-workflow.yml)                   | [token-llm](../entries/token-llm.yml)                                         | `ownership-token-workflow` → `token-llm`                       |
| [parameter-tech](../entries/parameter-tech.yml)                                       | [type-argument-tech](../entries/type-argument-tech.yml)                       | `parameter-tech` → `type-argument-tech`                        |
| [preserve-tech](../entries/preserve-tech.yml)                                         | [retain-tech](../entries/retain-tech.yml)                                     | `preserve-tech` → `retain-tech`                                |
| [recover-story](../entries/recover-story.yml)                                         | [restore-tech](../entries/restore-tech.yml)                                   | `recover-story` → `restore-tech`                               |
| [register-tech](../entries/register-tech.yml)                                         | [unregister-tech](../entries/unregister-tech.yml)                             | `register-tech` → `unregister-tech`                            |
| [reproduce-tech](../entries/reproduce-tech.yml)                                       | [reproducible-tech](../entries/reproducible-tech.yml)                         | `reproduce-tech` → `reproducible-tech`                         |
| [resilient-combat](../entries/resilient-combat.yml)                                   | [sturdy-general](../entries/sturdy-general.yml)                               | `sturdy-general` → `resilient-combat`                          |
| [resolve-tech](../entries/resolve-tech.yml)                                           | [resolve-workflow](../entries/resolve-workflow.yml)                           | `resolve-tech` → `resolve-workflow`                            |
| [revision-general](../entries/revision-general.yml)                                   | [revision-git](../entries/revision-git.yml)                                   | `revision-git` → `revision-general`                            |
| [ridge-terrain](../entries/ridge-terrain.yml)                                         | [summit-terrain](../entries/summit-terrain.yml)                               | `summit-terrain` → `ridge-terrain`                             |
| [scope-lexical](../entries/scope-lexical.yml)                                         | [scope-workflow](../entries/scope-workflow.yml)                               | `scope-lexical` → `scope-workflow`                             |
| [simmer-cooking](../entries/simmer-cooking.yml)                                       | [stew-cooking](../entries/stew-cooking.yml)                                   | `simmer-cooking` → `stew-cooking`                              |
| [spring-mechanical](../entries/spring-mechanical.yml)                                 | [spring-water](../entries/spring-water.yml)                                   | `spring-water` → `spring-mechanical`                           |
| [staged-git](../entries/staged-git.yml)                                               | [unstaged-git](../entries/unstaged-git.yml)                                   | `staged-git` → `unstaged-git`                                  |
| [sukueru](../entries/sukueru.yml)                                                     | [sukuu](../entries/sukuu.yml)                                                 | `sukuu` → `sukueru`                                            |
| [surface-terrain](../entries/surface-terrain.yml)                                     | [the-depths-zelda](../entries/the-depths-zelda.yml)                           | `surface-terrain` → `the-depths-zelda`                         |
| [susceptible-combat](../entries/susceptible-combat.yml)                               | [vulnerable-combat](../entries/vulnerable-combat.yml)                         | `vulnerable-combat` → `susceptible-combat`                     |
| [yareru](../entries/yareru.yml)                                                       | [yaru](../entries/yaru.yml)                                                   | `yaru` → `yareru`                                              |

## Complete review ledger

One row per actual stable ID, sorted within language. Read the linked YAML for the complete Context, Example, Notes, readings, aids, learning metadata, and declarations preserved by this plan. These rows summarize decisions, not truncated replacement content.

Decision labels: **A** = exact aid proposal; **R** = new relationship candidate; **K** = retain existing aid; **D** = additional investigation explicitly deferred; **NC-N** = no added field because existing Context/Notes already carries useful explanation or a boundary; **NC** = no additional aid selected for learning value. A row may contain multiple labels. Blank aids stay absent; no label demands future filling. Existing association counts are local declarations / direct neighbors, including reverse-only neighbors.

The Notes/Context adequacy screen flags explicit structure, register, distinction, limitation, and origin discussions. It does not automatically extract or alter text. Generic NC includes simple vocabulary, already clear concepts, and expressions for which an extra image or decomposition would add noise. Named entities without a D code are also not treated as verified naming histories.

### English (999)

| Stable ID / actual file                                                                                               | Word                                     | Current Meaning                                                            | Existing aids | Relations: local / displayed | Decision             |
| :-------------------------------------------------------------------------------------------------------------------- | :--------------------------------------- | :------------------------------------------------------------------------- | :------------ | :--------------------------- | :------------------- |
| [abandoned-workflow](../entries/abandoned-workflow.yml)                                                               | abandoned                                | 被遗弃的；无人继续处理的                                                   | absent        | 0 / 0                        | NC                   |
| [abstract-syntax-tree-compiler](../entries/abstract-syntax-tree-compiler.yml)                                         | abstract syntax tree                     | 抽象语法树；用树形节点表示程序的语法结构。                                 | absent        | 0 / 0                        | NC-N                 |
| [accented-text-tech](../entries/accented-text-tech.yml)                                                               | accented text                            | 带附加符号的文字                                                           | absent        | 0 / 0                        | NC                   |
| [accessor-tech](../entries/accessor-tech.yml)                                                                         | accessor                                 | 属性访问器                                                                 | W             | 0 / 1                        | K                    |
| [acquire-workflow](../entries/acquire-workflow.yml)                                                                   | acquire                                  | 获取；取得                                                                 | absent        | 0 / 0                        | NC                   |
| [actionable-tech](../entries/actionable-tech.yml)                                                                     | actionable                               | 可据以采取行动的                                                           | W             | 0 / 0                        | K                    |
| [activate-general](../entries/activate-general.yml)                                                                   | activate                                 | 激活、启动                                                                 | absent        | 0 / 0                        | NC                   |
| [actual-behavior-workflow](../entries/actual-behavior-workflow.yml)                                                   | actual behavior                          | 实际行为                                                                   | absent        | 1 / 1                        | NC                   |
| [adapter-tech](../entries/adapter-tech.yml)                                                                           | adapter                                  | 适配器                                                                     | W             | 0 / 0                        | K                    |
| [adjacent-tech](../entries/adjacent-tech.yml)                                                                         | adjacent                                 | 相邻的                                                                     | absent        | 0 / 0                        | A: ETY03             |
| [adventure-general](../entries/adventure-general.yml)                                                                 | adventure                                | 冒险；冒险经历。                                                           | absent        | 0 / 0                        | NC                   |
| [adventure-log-zelda](../entries/adventure-log-zelda.yml)                                                             | Adventure Log                            | 冒险笔记；《王国之泪》中的界面类别或功能名称。                             | absent        | 0 / 0                        | NC                   |
| [aerocuda-zelda](../entries/aerocuda-zelda.yml)                                                                       | Aerocuda                                 | 卡库达；《王国之泪》中的怪物或首领名称。                                   | absent        | 0 / 0                        | NC                   |
| [agent-ai](../entries/agent-ai.yml)                                                                                   | agent                                    | 智能体；围绕目标选择步骤、调用工具，并根据执行结果继续行动的系统。         | absent        | 0 / 0                        | NC-N                 |
| [aggressive-combat](../entries/aggressive-combat.yml)                                                                 | aggressive                               | 攻击性强的                                                                 | absent        | 0 / 1                        | NC                   |
| [aim-general](../entries/aim-general.yml)                                                                             | aim                                      | 瞄准                                                                       | absent        | 0 / 0                        | NC                   |
| [airborne-combat](../entries/airborne-combat.yml)                                                                     | airborne                                 | 空中的                                                                     | W             | 0 / 0                        | K                    |
| [album-ui](../entries/album-ui.yml)                                                                                   | album                                    | 相册；收集照片或图片的册子。                                               | absent        | 0 / 0                        | NC                   |
| [allowlist-tech](../entries/allowlist-tech.yml)                                                                       | allowlist                                | 允许列表                                                                   | absent        | 0 / 0                        | A: WF01              |
| [altar-architecture](../entries/altar-architecture.yml)                                                               | altar                                    | 祭坛                                                                       | absent        | 0 / 0                        | NC                   |
| [alternate-tech](../entries/alternate-tech.yml)                                                                       | alternate                                | 交替变化                                                                   | absent        | 0 / 0                        | NC                   |
| [ambitious-general](../entries/ambitious-general.yml)                                                                 | ambitious                                | 有野心的                                                                   | absent        | 0 / 0                        | NC                   |
| [amend-git](../entries/amend-git.yml)                                                                                 | amend                                    | 修改已有提交                                                               | absent        | 0 / 0                        | NC                   |
| [ancestor-general](../entries/ancestor-general.yml)                                                                   | ancestor                                 | 祖先                                                                       | absent        | 1 / 1                        | NC                   |
| [ancient-arowana-zelda](../entries/ancient-arowana-zelda.yml)                                                         | Ancient Arowana                          | 远昔骨舌鱼；《王国之泪》中的材料名称。                                     | absent        | 0 / 0                        | NC                   |
| [ancient-general](../entries/ancient-general.yml)                                                                     | ancient                                  | 古代的                                                                     | absent        | 0 / 0                        | NC                   |
| [annotation-tech](../entries/annotation-tech.yml)                                                                     | annotation                               | 标注                                                                       | absent        | 0 / 0                        | NC                   |
| [anonymous-tech](../entries/anonymous-tech.yml)                                                                       | anonymous                                | 匿名的                                                                     | absent        | 0 / 0                        | A: ETY02             |
| [apple-food](../entries/apple-food.yml)                                                                               | apple                                    | 苹果。                                                                     | absent        | 0 / 0                        | NC                   |
| [applicable-tech](../entries/applicable-tech.yml)                                                                     | applicable                               | 适用的                                                                     | absent        | 0 / 0                        | NC                   |
| [appreciate-general](../entries/appreciate-general.yml)                                                               | appreciate                               | 感激；认可价值                                                             | absent        | 0 / 0                        | NC                   |
| [aquatic-combat](../entries/aquatic-combat.yml)                                                                       | aquatic                                  | 水生的                                                                     | E             | 0 / 0                        | K                    |
| [archipelago-terrain](../entries/archipelago-terrain.yml)                                                             | archipelago                              | 群岛                                                                       | absent        | 0 / 0                        | NC                   |
| [armor-equipment](../entries/armor-equipment.yml)                                                                     | armor                                    | 护甲                                                                       | absent        | 1 / 1                        | NC                   |
| [armored-combat](../entries/armored-combat.yml)                                                                       | armored                                  | 装甲覆盖的                                                                 | W             | 0 / 0                        | K                    |
| [armory-architecture](../entries/armory-architecture.yml)                                                             | armory                                   | 武器库                                                                     | absent        | 0 / 1                        | NC                   |
| [arowana-general](../entries/arowana-general.yml)                                                                     | arowana                                  | 龙鱼。                                                                     | absent        | 0 / 0                        | NC                   |
| [array-hole-javascript](../entries/array-hole-javascript.yml)                                                         | array hole                               | 数组空槽；数组长度范围内某个索引没有对应自有元素的情况。                   | absent        | 0 / 0                        | R: R07               |
| [array-spread-tech](../entries/array-spread-tech.yml)                                                                 | array spread                             | 数组展开                                                                   | absent        | 0 / 0                        | NC                   |
| [array-tech](../entries/array-tech.yml)                                                                               | array                                    | 数组；按索引组织元素的集合。                                               | absent        | 0 / 0                        | NC                   |
| [arrow-weapon](../entries/arrow-weapon.yml)                                                                           | arrow                                    | 箭                                                                         | absent        | 0 / 0                        | NC                   |
| [artifact-tech](../entries/artifact.yml)                                                                              | artifact                                 | 构建产物                                                                   | E,N           | 1 / 1                        | K                    |
| [artifact-workflow](../entries/artifact-workflow.yml)                                                                 | artifact                                 | 工作成果文件                                                               | N             | 0 / 1                        | K                    |
| [as-proof-of-workflow](../entries/as-proof-of-workflow.yml)                                                           | as proof of                              | 作为……的证明                                                               | absent        | 0 / 0                        | NC                   |
| [ascend-general](../entries/ascend-general.yml)                                                                       | ascend                                   | 上升；向高处移动。                                                         | absent        | 1 / 1                        | NC                   |
| [assemble-general](../entries/assemble-general.yml)                                                                   | assemble                                 | 组装                                                                       | absent        | 0 / 0                        | NC                   |
| [assertion-testing](../entries/assertion-testing.yml)                                                                 | assertion                                | 测试断言                                                                   | absent        | 0 / 0                        | NC                   |
| [assignment-target-tech](../entries/assignment-target-tech.yml)                                                       | assignment target                        | 赋值目标                                                                   | absent        | 0 / 0                        | NC                   |
| [async-await-javascript](../entries/async-await-javascript.yml)                                                       | async/await                              | 异步函数语法；用 await 组织基于 Promise 的控制流。                         | absent        | 1 / 1                        | NC                   |
| [asynchronous-tech](../entries/asynchronous-tech.yml)                                                                 | asynchronous                             | 异步的                                                                     | W             | 0 / 0                        | K                    |
| [at-most-workflow](../entries/at-most-workflow.yml)                                                                   | at most                                  | 最多；至多                                                                 | absent        | 0 / 0                        | NC                   |
| [atomically-workflow](../entries/atomically-workflow.yml)                                                             | atomically                               | 原子地                                                                     | absent        | 0 / 0                        | NC-N                 |
| [attach-general](../entries/attach-general.yml)                                                                       | attach                                   | 连接、附加                                                                 | absent        | 0 / 0                        | NC                   |
| [attack-power-combat](../entries/attack-power-combat.yml)                                                             | attack power                             | 攻击力                                                                     | W             | 0 / 0                        | K                    |
| [attack-up-zelda](../entries/attack-up-zelda.yml)                                                                     | Attack Up                                | 攻击力提升；《王国之泪》中的能力效果名称。                                 | absent        | 0 / 0                        | NC                   |
| [attribute-to-general](../entries/attribute-to-general.yml)                                                           | attribute to                             | 归因于                                                                     | absent        | 0 / 0                        | NC                   |
| [audio-general](../entries/audio-general.yml)                                                                         | audio                                    | 音频；录制或传输的声音。                                                   | absent        | 1 / 1                        | A: ETY06; R: R25     |
| [audit-tech](../entries/audit-tech.yml)                                                                               | audit                                    | 审查                                                                       | absent        | 0 / 0                        | A: ETY05; R: R25     |
| [autobuild-zelda](../entries/autobuild-zelda.yml)                                                                     | Autobuild                                | 蓝图；复现已记录组合设计的能力。                                           | W             | 0 / 0                        | K                    |
| [await-review-workflow](../entries/await-review-workflow.yml)                                                         | await review                             | 等待审查                                                                   | absent        | 0 / 0                        | NC                   |
| [awaken-story](../entries/awaken-story.yml)                                                                           | awaken                                   | 觉醒                                                                       | absent        | 0 / 0                        | NC                   |
| [axe-weapon](../entries/axe-weapon.yml)                                                                               | axe                                      | 斧头                                                                       | absent        | 0 / 0                        | NC                   |
| [backend-tech](../entries/backend-tech.yml)                                                                           | backend                                  | 实现后端                                                                   | absent        | 0 / 0                        | NC                   |
| [backtracking-algorithm](../entries/backtracking-algorithm.yml)                                                       | backtracking                             | 回溯；逐步尝试选择，撤销不合适的选择并继续搜索。                           | W             | 0 / 0                        | R: R10; K            |
| [balloon-engineering](../entries/balloon-engineering.yml)                                                             | balloon                                  | 气球。                                                                     | absent        | 0 / 0                        | NC                   |
| [banana-food](../entries/banana-food.yml)                                                                             | banana                                   | 香蕉。                                                                     | absent        | 0 / 0                        | NC                   |
| [bandanna-clothing](../entries/bandanna-clothing.yml)                                                                 | bandanna                                 | 头巾                                                                       | absent        | 0 / 0                        | NC                   |
| [based-on-workflow](../entries/based-on-workflow.yml)                                                                 | based on                                 | 基于；根据                                                                 | absent        | 0 / 0                        | NC                   |
| [baseline-tech](../entries/baseline-tech.yml)                                                                         | baseline                                 | 基准输出                                                                   | absent        | 0 / 0                        | A: WF02, N03; R: R01 |
| [bass-food](../entries/bass-food.yml)                                                                                 | bass                                     | 鲈鱼。                                                                     | N             | 0 / 0                        | K                    |
| [battery-engineering](../entries/battery-engineering.yml)                                                             | battery                                  | 电池                                                                       | absent        | 0 / 0                        | NC                   |
| [beam-emitter-zelda](../entries/beam-emitter-zelda.yml)                                                               | Beam Emitter                             | 光线头；《王国之泪》中的左纳乌装置。                                       | W             | 0 / 0                        | K                    |
| [beam-engineering](../entries/beam-engineering.yml)                                                                   | beam                                     | 光束；此处指集中的一束光。                                                 | absent        | 0 / 0                        | NC                   |
| [bee-general](../entries/bee-general.yml)                                                                             | bee                                      | 蜂；蜜蜂。                                                                 | absent        | 0 / 0                        | NC                   |
| [beef-food](../entries/beef-food.yml)                                                                                 | beef                                     | 牛肉                                                                       | absent        | 0 / 0                        | NC                   |
| [before-and-after-the-fix-workflow](../entries/before-and-after-the-fix-workflow.yml)                                 | before and after the fix                 | 修复前后                                                                   | absent        | 0 / 0                        | NC                   |
| [before-drawing-conclusions-workflow](../entries/before-drawing-conclusions-workflow.yml)                             | before drawing conclusions               | 在得出结论之前                                                             | absent        | 0 / 0                        | NC                   |
| [benchmark-tech](../entries/benchmark-tech.yml)                                                                       | benchmark                                | 性能基准测试                                                               | absent        | 0 / 0                        | R: R01               |
| [big-battery-zelda](../entries/big-battery-zelda.yml)                                                                 | Big Battery                              | 大容量电池；《王国之泪》中的左纳乌装置。                                   | absent        | 0 / 0                        | NC                   |
| [big-wheel-zelda](../entries/big-wheel-zelda.yml)                                                                     | Big Wheel                                | 大轮胎；《王国之泪》中的左纳乌装置。                                       | absent        | 0 / 0                        | NC                   |
| [bigint-javascript](../entries/bigint-javascript.yml)                                                                 | BigInt                                   | 大整数；用于表示任意精度整数的 JavaScript 原始类型。                       | absent        | 0 / 0                        | NC                   |
| [binary-search-algorithm](../entries/binary-search-algorithm.yml)                                                     | binary search                            | 二分查找；在有序区间中比较中点并排除一半候选。                             | absent        | 0 / 0                        | NC                   |
| [binding-tech](../entries/binding-tech.yml)                                                                           | binding                                  | 绑定                                                                       | absent        | 0 / 0                        | NC                   |
| [black-pepper-food](../entries/black-pepper-food.yml)                                                                 | black pepper                             | 黑胡椒                                                                     | absent        | 0 / 0                        | NC                   |
| [blessing-general](../entries/blessing-general.yml)                                                                   | blessing                                 | 祝福；被视为有益或带来幸福的事物。                                         | absent        | 0 / 0                        | NC                   |
| [block-formatting-context-css](../entries/block-formatting-context-css.yml)                                           | block formatting context                 | 块级格式化上下文；具有独立块布局规则的区域。                               | N             | 0 / 0                        | K                    |
| [block-general](../entries/block-general.yml)                                                                         | block                                    | 格挡；挡住攻击或去路。                                                     | absent        | 0 / 0                        | NC                   |
| [blocker-workflow](../entries/blocker-workflow.yml)                                                                   | blocker                                  | 阻碍事项                                                                   | W             | 0 / 0                        | K                    |
| [bokoblin-zelda](../entries/bokoblin-zelda.yml)                                                                       | Bokoblin                                 | 波克布林；《王国之泪》中的怪物或首领名称。                                 | absent        | 0 / 0                        | NC                   |
| [bomb-weapon](../entries/bomb-weapon.yml)                                                                             | bomb                                     | 炸弹。                                                                     | absent        | 0 / 0                        | NC                   |
| [boomerang-weapon](../entries/boomerang-weapon.yml)                                                                   | boomerang                                | 回旋镖                                                                     | absent        | 0 / 0                        | NC                   |
| [boots-clothing](../entries/boots-clothing.yml)                                                                       | boots                                    | 靴子                                                                       | absent        | 0 / 0                        | NC                   |
| [boss-bokoblin-zelda](../entries/boss-bokoblin-zelda.yml)                                                             | Boss Bokoblin                            | 首领波克布林；《王国之泪》中的怪物或首领名称。                             | absent        | 0 / 0                        | NC                   |
| [bounded-tech](../entries/bounded-tech.yml)                                                                           | bounded                                  | 有界的；范围受限的                                                         | absent        | 0 / 0                        | NC                   |
| [bow-weapon](../entries/bow-weapon.yml)                                                                               | bow                                      | 弓                                                                         | absent        | 0 / 0                        | NC                   |
| [braised-food](../entries/braised-food.yml)                                                                           | braised                                  | 炖的；焖制的                                                               | absent        | 0 / 0                        | NC                   |
| [branch-git](../entries/branch-git.yml)                                                                               | branch                                   | 分支；用于独立推进一条开发历史的引用。                                     | absent        | 0 / 0                        | NC                   |
| [brave-general](../entries/brave-general.yml)                                                                         | brave                                    | 勇敢的                                                                     | absent        | 1 / 2                        | NC                   |
| [breadth-first-tech](../entries/breadth-first-tech.yml)                                                               | breadth-first                            | 广度优先的                                                                 | W             | 0 / 0                        | K                    |
| [bridge-architecture](../entries/bridge-architecture.yml)                                                             | bridge                                   | 桥梁                                                                       | absent        | 0 / 0                        | NC                   |
| [bridge-tech](../entries/bridge-tech.yml)                                                                             | bridge                                   | 桥接层                                                                     | absent        | 0 / 0                        | NC                   |
| [brightcap-zelda](../entries/brightcap-zelda.yml)                                                                     | Brightcap                                | 光亮蘑菇；《王国之泪》中的材料名称。                                       | absent        | 0 / 0                        | NC                   |
| [broadsword-weapon](../entries/broadsword-weapon.yml)                                                                 | broadsword                               | 阔剑                                                                       | W             | 0 / 0                        | K                    |
| [bug-tech](../entries/bug-tech.yml)                                                                                   | bug                                      | 程序缺陷；使软件偏离预期行为的问题。                                       | absent        | 0 / 0                        | NC                   |
| [build-construction](../entries/build-construction.yml)                                                               | build                                    | 建造；搭建建筑、装置等实体。                                               | absent        | 0 / 0                        | NC                   |
| [build-tech](../entries/build-tech.yml)                                                                               | build                                    | 构建；执行编译、打包等步骤生成项目产物。                                   | absent        | 0 / 0                        | NC                   |
| [butter-food](../entries/butter-food.yml)                                                                             | butter                                   | 黄油。                                                                     | absent        | 0 / 0                        | NC                   |
| [bypass-tech](../entries/bypass-tech.yml)                                                                             | bypass                                   | 绕过                                                                       | absent        | 0 / 0                        | R: R06               |
| [byte-offset-tech](../entries/byte-offset-tech.yml)                                                                   | byte offset                              | 字节偏移                                                                   | absent        | 0 / 0                        | A: N04; R: R08       |
| [cache-invalidation-tech](../entries/cache-invalidation-tech.yml)                                                     | cache invalidation                       | 缓存失效处理；在依赖变化后使旧缓存结果不再被当作有效结果使用。             | absent        | 0 / 0                        | NC                   |
| [cache-revalidation-web](../entries/cache-revalidation-web.yml)                                                       | cache revalidation                       | 缓存再验证；向服务器确认已存储响应能否继续使用。                           | absent        | 0 / 0                        | NC                   |
| [cache-tech](../entries/cache-tech.yml)                                                                               | cache                                    | 缓存；暂存数据或计算结果以供后续复用的机制。                               | E             | 1 / 1                        | K                    |
| [calamity-story](../entries/calamity-story.yml)                                                                       | calamity                                 | 灾厄、大灾难                                                               | absent        | 0 / 0                        | NC                   |
| [call-stack-javascript](../entries/call-stack-javascript.yml)                                                         | call stack                               | 调用栈；按后进先出顺序记录函数调用和返回位置。                             | absent        | 0 / 0                        | R: R09, R21          |
| [callback-tech](../entries/callback-tech.yml)                                                                         | callback                                 | 回调函数                                                                   | W             | 0 / 0                        | K                    |
| [cancel-general](../entries/cancel-general.yml)                                                                       | cancel                                   | 取消                                                                       | absent        | 0 / 0                        | NC                   |
| [cancellation-tech](../entries/cancellation-tech.yml)                                                                 | cancellation                             | 取消                                                                       | absent        | 0 / 0                        | NC                   |
| [candidate-tech](../entries/candidate-tech.yml)                                                                       | candidate                                | 候选项                                                                     | E             | 0 / 0                        | K                    |
| [cannon-engineering](../entries/cannon-engineering.yml)                                                               | cannon                                   | 大炮                                                                       | absent        | 0 / 0                        | NC                   |
| [canyon-terrain](../entries/canyon-terrain.yml)                                                                       | canyon                                   | 峡谷                                                                       | absent        | 0 / 1                        | NC                   |
| [captain-construct-zelda](../entries/captain-construct-zelda.yml)                                                     | Captain Construct                        | 队长魔像；《王国之泪》中的怪物或首领名称。                                 | absent        | 0 / 0                        | NC                   |
| [carrot-food](../entries/carrot-food.yml)                                                                             | carrot                                   | 胡萝卜。                                                                   | absent        | 0 / 0                        | NC                   |
| [cart-engineering](../entries/cart-engineering.yml)                                                                   | cart                                     | 小车；运货车。                                                             | absent        | 0 / 0                        | NC                   |
| [cast-tech](../entries/cast-tech.yml)                                                                                 | cast                                     | 类型转换写法；类型断言写法                                                 | absent        | 1 / 1                        | NC                   |
| [castle-architecture](../entries/castle-architecture.yml)                                                             | castle                                   | 城堡                                                                       | absent        | 0 / 0                        | NC                   |
| [catch-handler-tech](../entries/catch-handler-tech.yml)                                                               | catch handler                            | 异常捕获处理器                                                             | absent        | 0 / 0                        | NC                   |
| [cave-architecture](../entries/cave-architecture.yml)                                                                 | cave                                     | 洞穴                                                                       | absent        | 1 / 1                        | NC                   |
| [caveat-general](../entries/caveat-general.yml)                                                                       | caveat                                   | 限定说明；注意事项                                                         | absent        | 0 / 0                        | A: ETY01             |
| [cavern-architecture](../entries/cavern-architecture.yml)                                                             | cavern                                   | 大型洞穴                                                                   | N             | 0 / 1                        | K                    |
| [ceiling-architecture](../entries/ceiling-architecture.yml)                                                           | ceiling                                  | 天花板                                                                     | absent        | 0 / 0                        | NC                   |
| [chamber-architecture](../entries/chamber-architecture.yml)                                                           | chamber                                  | 房间、密室                                                                 | absent        | 0 / 0                        | NC                   |
| [champion-general](../entries/champion-general.yml)                                                                   | champion                                 | 英杰、冠军                                                                 | absent        | 0 / 0                        | NC                   |
| [character-profiles-zelda](../entries/character-profiles-zelda.yml)                                                   | Character Profiles                       | 人物名录；《王国之泪》中的界面类别或功能名称。                             | absent        | 0 / 0                        | NC                   |
| [chasm-terrain](../entries/chasm-terrain.yml)                                                                         | chasm                                    | 深渊；深而宽的裂隙。                                                       | N             | 1 / 1                        | K                    |
| [cheese-food](../entries/cheese-food.yml)                                                                             | cheese                                   | 奶酪。                                                                     | absent        | 0 / 0                        | NC                   |
| [chicken-food](../entries/chicken-food.yml)                                                                           | chicken                                  | 鸡肉                                                                       | absent        | 0 / 0                        | NC-N                 |
| [chilly-general](../entries/chilly-general.yml)                                                                       | chilly                                   | 寒冷的；凉飕飕的。                                                         | absent        | 0 / 0                        | NC                   |
| [chuchu-zelda](../entries/chuchu-zelda.yml)                                                                           | Chuchu                                   | 丘丘；《王国之泪》中的怪物或首领名称。                                     | absent        | 0 / 0                        | NC                   |
| [citadel-architecture](../entries/citadel-architecture.yml)                                                           | citadel                                  | 城塞、堡垒                                                                 | absent        | 0 / 0                        | NC                   |
| [civilization-general](../entries/civilization-general.yml)                                                           | civilization                             | 文明                                                                       | absent        | 0 / 0                        | NC                   |
| [claim-general](../entries/claim-general.yml)                                                                         | claim                                    | 声称；论断                                                                 | absent        | 0 / 0                        | NC                   |
| [clarify-general](../entries/clarify-general.yml)                                                                     | clarify                                  | 澄清；说清楚                                                               | absent        | 0 / 0                        | NC                   |
| [claude-ai](../entries/claude-ai.yml)                                                                                 | Claude                                   | Anthropic 的 AI 模型与助手品牌。                                           | absent        | 0 / 0                        | NC-N                 |
| [clause-tech](../entries/clause-tech.yml)                                                                             | clause                                   | 语法子句                                                                   | absent        | 0 / 0                        | NC-N                 |
| [claw-general](../entries/claw-general.yml)                                                                           | claw                                     | 爪；动物的尖爪或龙虾等动物的螯。                                           | absent        | 0 / 0                        | NC-N                 |
| [claymore-weapon](../entries/claymore-weapon.yml)                                                                     | claymore                                 | 双手大剑                                                                   | absent        | 0 / 0                        | NC                   |
| [cleanup-tech](../entries/cleanup-tech.yml)                                                                           | cleanup                                  | 清理                                                                       | W             | 0 / 0                        | K                    |
| [cliff-terrain](../entries/cliff-terrain.yml)                                                                         | cliff                                    | 悬崖                                                                       | absent        | 0 / 0                        | NC                   |
| [climb-general](../entries/climb-general.yml)                                                                         | climb                                    | 攀爬                                                                       | absent        | 0 / 0                        | NC                   |
| [climbing-gear-zelda](../entries/climbing-gear-zelda.yml)                                                             | Climbing Gear                            | 攀登护手；《王国之泪》中的防具名称。                                       | absent        | 0 / 0                        | NC                   |
| [closed-general](../entries/closed-general.yml)                                                                       | closed                                   | 关闭的；未敞开的。                                                         | absent        | 0 / 0                        | NC                   |
| [closure-tech](../entries/closure-tech.yml)                                                                           | closure                                  | 闭包                                                                       | N             | 1 / 1                        | K                    |
| [club-weapon](../entries/club-weapon.yml)                                                                             | club                                     | 棍棒                                                                       | absent        | 0 / 0                        | NC                   |
| [clue-story](../entries/clue-story.yml)                                                                               | clue                                     | 线索                                                                       | absent        | 0 / 0                        | NC                   |
| [code-generation-tech](../entries/code-generation-tech.yml)                                                           | code generation                          | 代码生成                                                                   | absent        | 0 / 0                        | NC                   |
| [code-unit-tech](../entries/code-unit-tech.yml)                                                                       | code unit                                | 码元                                                                       | absent        | 0 / 0                        | R: R08               |
| [codex-ai](../entries/codex-ai.yml)                                                                                   | Codex                                    | OpenAI 的 AI 编程智能体。                                                  | absent        | 0 / 0                        | NC-N                 |
| [coerce-tech](../entries/coerce-tech.yml)                                                                             | coerce                                   | 转换类型                                                                   | absent        | 0 / 0                        | NC                   |
| [cold-resistance-combat](../entries/cold-resistance-combat.yml)                                                       | cold resistance                          | 耐寒；承受寒冷环境影响的能力。                                             | absent        | 1 / 1                        | NC                   |
| [colgera-zelda](../entries/colgera-zelda.yml)                                                                         | Colgera                                  | 急冻盖拉；《王国之泪》中的怪物或首领名称。                                 | absent        | 0 / 0                        | NC                   |
| [collision-tech](../entries/collision-tech.yml)                                                                       | collision                                | 表示冲突                                                                   | absent        | 0 / 0                        | NC                   |
| [colossal-combat](../entries/colossal-combat.yml)                                                                     | colossal                                 | 极其庞大的                                                                 | absent        | 0 / 0                        | NC                   |
| [combining-mark-tech](../entries/combining-mark-tech.yml)                                                             | combining mark                           | 组合附加符号                                                               | absent        | 0 / 0                        | NC                   |
| [commit-git](../entries/commit-git.yml)                                                                               | commit                                   | 提交；记录一次项目快照，也可指这次记录。                                   | absent        | 0 / 0                        | NC-N                 |
| [compendium-general](../entries/compendium-general.yml)                                                               | compendium                               | 概要汇编；汇集某领域资料的手册。                                           | absent        | 0 / 0                        | NC                   |
| [competing-work-workflow](../entries/competing-work-workflow.yml)                                                     | competing work                           | 相互重叠的在途工作                                                         | absent        | 0 / 0                        | NC                   |
| [compile-tech](../entries/compile-tech.yml)                                                                           | compile                                  | 编译；把源代码转换为目标代码或另一种程序表示。                             | absent        | 1 / 1                        | D: D03               |
| [compiler-tech](../entries/compiler-tech.yml)                                                                         | compiler                                 | 编译器；执行程序编译的软件。                                               | absent        | 0 / 1                        | A: WF03              |
| [compound-assignment-tech](../entries/compound-assignment-tech.yml)                                                   | compound assignment                      | 复合赋值                                                                   | absent        | 0 / 0                        | NC                   |
| [computed-key-tech](../entries/computed-key-tech.yml)                                                                 | computed key                             | 计算属性键                                                                 | absent        | 0 / 0                        | NC                   |
| [concatenation-tech](../entries/concatenation-tech.yml)                                                               | concatenation                            | 连接；拼接                                                                 | absent        | 0 / 0                        | NC                   |
| [concise-general](../entries/concise-general.yml)                                                                     | concise                                  | 简洁的                                                                     | absent        | 0 / 0                        | NC                   |
| [concrete-tech](../entries/concrete-tech.yml)                                                                         | concrete                                 | 具体的；有实际依据的                                                       | absent        | 0 / 0                        | NC                   |
| [confirmed-tech](../entries/confirmed-tech.yml)                                                                       | confirmed                                | 已确认的                                                                   | absent        | 0 / 0                        | NC                   |
| [consistent-tech](../entries/consistent-tech.yml)                                                                     | consistent                               | 一致的                                                                     | absent        | 0 / 0                        | NC                   |
| [constant-propagation-tech](../entries/constant-propagation-tech.yml)                                                 | constant propagation                     | 常量传播                                                                   | absent        | 0 / 0                        | NC                   |
| [construct-being-zelda](../entries/construct-being-zelda.yml)                                                         | Construct                                | 魔像；本作中由左纳乌技术造出的人工生命体或机械实体。                       | absent        | 0 / 0                        | NC                   |
| [construct-factory-zelda](../entries/construct-factory-zelda.yml)                                                     | Construct Factory                        | 魔像制造房；《王国之泪》中的地点或设施。                                   | absent        | 0 / 0                        | NC                   |
| [construct-general](../entries/construct-general.yml)                                                                 | construct                                | 建造、构造                                                                 | absent        | 0 / 0                        | NC                   |
| [construct-head-zelda](../entries/construct-head-zelda.yml)                                                           | Construct Head                           | 魔像的头部；《王国之泪》中的左纳乌装置。                                   | W             | 0 / 0                        | K                    |
| [consume-tech](../entries/consume-tech.yml)                                                                           | consume                                  | 读取并处理                                                                 | absent        | 0 / 0                        | NC                   |
| [context-language](../entries/context-language.yml)                                                                   | context                                  | 语境；上下文；帮助理解含义的背景。                                         | absent        | 0 / 0                        | NC-N                 |
| [contextual-type-tech](../entries/contextual-type-tech.yml)                                                           | contextual type                          | 上下文类型                                                                 | absent        | 0 / 0                        | NC                   |
| [contribution-general](../entries/contribution-general.yml)                                                           | contribution                             | 贡献                                                                       | absent        | 0 / 0                        | NC                   |
| [cook-cooking](../entries/cook-cooking.yml)                                                                           | cook                                     | 烹饪                                                                       | absent        | 0 / 0                        | NC                   |
| [cooking-pot-cooking](../entries/cooking-pot-cooking.yml)                                                             | cooking pot                              | 烹饪锅                                                                     | W             | 0 / 0                        | K                    |
| [coordinates-ui](../entries/coordinates-ui.yml)                                                                       | coordinates                              | 坐标                                                                       | absent        | 0 / 0                        | NC                   |
| [copilot-ai](../entries/copilot-ai.yml)                                                                               | Copilot                                  | 此处指 GitHub Copilot，用于辅助软件开发的 AI 助手。                        | W             | 0 / 0                        | K                    |
| [copy-tech](../entries/copy-tech.yml)                                                                                 | copy                                     | 复制；创建数据副本；也可指副本。                                           | absent        | 0 / 0                        | NC-N                 |
| [correctness-tech](../entries/correctness-tech.yml)                                                                   | correctness                              | 正确性                                                                     | W             | 0 / 1                        | K                    |
| [corresponding-general](../entries/corresponding-general.yml)                                                         | corresponding                            | 对应的                                                                     | absent        | 0 / 0                        | NC                   |
| [corridor-architecture](../entries/corridor-architecture.yml)                                                         | corridor                                 | 走廊                                                                       | absent        | 0 / 0                        | NC                   |
| [corrupt-tech](../entries/corrupt-tech.yml)                                                                           | corrupt                                  | 损坏；错误改变                                                             | absent        | 0 / 0                        | NC                   |
| [corrupted-combat](../entries/corrupted-combat.yml)                                                                   | corrupted                                | 被腐化的                                                                   | absent        | 0 / 0                        | NC                   |
| [counter-tech](../entries/counter-tech.yml)                                                                           | counter                                  | 计数器；记录某事件发生次数或累计数量的状态。                               | absent        | 0 / 0                        | NC                   |
| [courageous-general](../entries/courageous-general.yml)                                                               | courageous                               | 勇敢的                                                                     | E,W           | 2 / 2                        | K                    |
| [courser-bee-honey-zelda](../entries/courser-bee-honey-zelda.yml)                                                     | Courser Bee Honey                        | 精力蜂的蜂蜜；《王国之泪》中的材料名称。                                   | absent        | 0 / 0                        | NC                   |
| [coverage-tech](../entries/coverage-tech.yml)                                                                         | coverage                                 | 覆盖范围                                                                   | absent        | 0 / 0                        | NC                   |
| [crash-tech](../entries/crash-tech.yml)                                                                               | crash                                    | 崩溃；程序因故障意外停止或无法继续正常执行。                               | absent        | 0 / 0                        | NC-N                 |
| [crater-terrain](../entries/crater-terrain.yml)                                                                       | crater                                   | 火山口                                                                     | absent        | 0 / 0                        | NC                   |
| [crisis-story](../entries/crisis-story.yml)                                                                           | crisis                                   | 危机                                                                       | absent        | 0 / 0                        | NC                   |
| [critical-hit-combat](../entries/critical-hit-combat.yml)                                                             | critical hit                             | 暴击                                                                       | W             | 0 / 0                        | K                    |
| [cross-functional-collaboration-workflow](../entries/cross-functional-collaboration-workflow.yml)                     | cross-functional collaboration           | 跨职能协作；围绕共同目标协调不同角色的工作。                               | absent        | 0 / 0                        | NC                   |
| [cross-origin-resource-sharing-web](../entries/cross-origin-resource-sharing-web.yml)                                 | cross-origin resource sharing            | 跨源资源共享；服务器通过响应头授权浏览器进行跨源读取。                     | absent        | 1 / 1                        | NC-N                 |
| [crouch-general](../entries/crouch-general.yml)                                                                       | crouch                                   | 蹲下                                                                       | absent        | 0 / 0                        | NC                   |
| [crystal-general](../entries/crystal-general.yml)                                                                     | crystal                                  | 水晶；晶体。                                                               | absent        | 0 / 0                        | NC                   |
| [crystal-refinery-zelda](../entries/crystal-refinery-zelda.yml)                                                       | Crystal Refinery                         | 电池制造机；《王国之泪》中的地点或设施。                                   | absent        | 0 / 0                        | NC                   |
| [cunning-general](../entries/cunning-general.yml)                                                                     | cunning                                  | 狡猾的                                                                     | absent        | 0 / 0                        | NC                   |
| [curious-general](../entries/curious-general.yml)                                                                     | curious                                  | 好奇的                                                                     | absent        | 0 / 0                        | NC                   |
| [curry-food](../entries/curry-food.yml)                                                                               | curry                                    | 咖喱                                                                       | absent        | 0 / 0                        | NC                   |
| [cursor-ai](../entries/cursor-ai.yml)                                                                                 | Cursor                                   | 用于软件开发的 AI 编辑器与编程智能体产品名称。                             | absent        | 0 / 0                        | NC-N                 |
| [dagger-weapon](../entries/dagger-weapon.yml)                                                                         | dagger                                   | 匕首                                                                       | absent        | 0 / 0                        | NC                   |
| [damage-combat](../entries/damage-combat.yml)                                                                         | damage                                   | 伤害                                                                       | absent        | 0 / 0                        | NC                   |
| [death-mountain-zelda](../entries/death-mountain-zelda.yml)                                                           | Death Mountain                           | 死亡之山；《王国之泪》中的地点或设施。                                     | absent        | 0 / 0                        | NC                   |
| [debounce-javascript](../entries/debounce-javascript.yml)                                                             | debounce                                 | 防抖；把短时间内的连续调用合并为较少的执行。                               | absent        | 0 / 0                        | NC-N                 |
| [decayed-general](../entries/decayed-general.yml)                                                                     | decayed                                  | 腐朽的；腐败或衰损的。                                                     | absent        | 1 / 1                        | NC-N                 |
| [declaration-emit-tech](../entries/declaration-emit-tech.yml)                                                         | declaration emit                         | 声明生成                                                                   | absent        | 0 / 0                        | NC                   |
| [decode-tech](../entries/decode-tech.yml)                                                                             | decode                                   | 解码                                                                       | N             | 0 / 1                        | K                    |
| [decorator-tech](../entries/decorator-tech.yml)                                                                       | decorator                                | 装饰器                                                                     | absent        | 0 / 0                        | NC                   |
| [deduplicate-tech](../entries/deduplicate-tech.yml)                                                                   | deduplicate                              | 去重；识别重复项并按规则保留一份。                                         | absent        | 0 / 0                        | NC                   |
| [deep-copy-javascript](../entries/deep-copy-javascript.yml)                                                           | deep copy                                | 深拷贝；复制嵌套数据，使相应对象引用与原数据隔离。                         | absent        | 1 / 1                        | NC-N                 |
| [deep-general](../entries/deep-general.yml)                                                                           | deep                                     | 深的；深入的。                                                             | absent        | 0 / 0                        | NC-N                 |
| [deepseek-ai](../entries/deepseek-ai.yml)                                                                             | DeepSeek                                 | 深度求索；提供 AI 模型及相关服务的品牌。                                   | W             | 0 / 0                        | K                    |
| [defense-combat](../entries/defense-combat.yml)                                                                       | defense                                  | 防御力                                                                     | absent        | 0 / 0                        | NC                   |
| [defense-up-zelda](../entries/defense-up-zelda.yml)                                                                   | Defense Up                               | 防御提升；《王国之泪》中的能力效果名称。                                   | absent        | 0 / 0                        | NC                   |
| [defer-tech](../entries/defer-tech.yml)                                                                               | defer                                    | 推迟执行                                                                   | absent        | 0 / 0                        | NC                   |
| [deity-story](../entries/deity-story.yml)                                                                             | deity                                    | 神；神明。                                                                 | absent        | 0 / 0                        | R: R13               |
| [demon-dragon-zelda](../entries/demon-dragon-zelda.yml)                                                               | Demon Dragon                             | 黑龙；加侬多夫龙化后的形态。                                               | absent        | 0 / 0                        | NC                   |
| [demon-king-ganondorf-zelda](../entries/demon-king-ganondorf-zelda.yml)                                               | Demon King Ganondorf                     | 魔王加侬多夫；《王国之泪》中的怪物或首领名称。                             | absent        | 0 / 0                        | NC                   |
| [demon-king-zelda](../entries/demon-king-zelda.yml)                                                                   | Demon King                               | 魔王；加侬多夫在本作中的称号。                                             | absent        | 0 / 0                        | NC                   |
| [demon-kings-bow-zelda](../entries/demon-kings-bow-zelda.yml)                                                         | Demon King's Bow                         | 魔王之弓；《王国之泪》中的装备名称。                                       | absent        | 0 / 0                        | NC                   |
| [demon-story](../entries/demon-story.yml)                                                                             | demon                                    | 恶魔；魔鬼。                                                               | absent        | 0 / 0                        | NC                   |
| [deprecated-tech](../entries/deprecated.yml)                                                                          | deprecated                               | 已弃用的；不建议继续使用的                                                 | N             | 0 / 0                        | K                    |
| [depths-terrain](../entries/depths-terrain.yml)                                                                       | depths                                   | 深处；深层区域；常用复数形式。                                             | absent        | 0 / 0                        | NC                   |
| [dereference-tech](../entries/dereference-tech.yml)                                                                   | dereference                              | 解引用；通过指针或引用访问所指向的值。                                     | absent        | 0 / 0                        | NC                   |
| [derive-tech](../entries/derive-tech.yml)                                                                             | derive                                   | 推导；取得                                                                 | absent        | 0 / 0                        | NC                   |
| [descend-general](../entries/descend-general.yml)                                                                     | descend                                  | 下降                                                                       | absent        | 0 / 1                        | NC                   |
| [descendant-general](../entries/descendant-general.yml)                                                               | descendant                               | 后裔                                                                       | W             | 0 / 2                        | K                    |
| [desert-terrain](../entries/desert-terrain.yml)                                                                       | desert                                   | 沙漠                                                                       | absent        | 0 / 0                        | NC                   |
| [destination-ui](../entries/destination-ui.yml)                                                                       | destination                              | 目的地                                                                     | absent        | 0 / 0                        | NC                   |
| [destroy-general](../entries/destroy-general.yml)                                                                     | destroy                                  | 摧毁；毁灭。                                                               | absent        | 0 / 0                        | NC                   |
| [detach-general](../entries/detach-general.yml)                                                                       | detach                                   | 分离、拆卸                                                                 | absent        | 0 / 0                        | NC                   |
| [detach-git](../entries/detach-git.yml)                                                                               | detach                                   | 分离 HEAD                                                                  | absent        | 0 / 0                        | NC                   |
| [determined-general](../entries/determined-general.yml)                                                               | determined                               | 坚定的                                                                     | absent        | 0 / 0                        | NC                   |
| [device-dispenser-zelda](../entries/device-dispenser-zelda.yml)                                                       | Device Dispenser                         | 左纳乌装置制造机；《王国之泪》中的地点或设施。                             | W             | 0 / 0                        | K                    |
| [device-engineering](../entries/device-engineering.yml)                                                               | device                                   | 装置；设备；为某种用途设计的器具。                                         | absent        | 0 / 0                        | NC                   |
| [diagnostic-context-tech](../entries/diagnostic-context-tech.yml)                                                     | diagnostic context                       | 诊断上下文                                                                 | absent        | 0 / 0                        | NC                   |
| [diagnostic-tech](../entries/diagnostic-tech.yml)                                                                     | diagnostic                               | 诊断信息                                                                   | absent        | 0 / 0                        | NC                   |
| [diligent-general](../entries/diligent-general.yml)                                                                   | diligent                                 | 勤勉的                                                                     | absent        | 0 / 0                        | NC                   |
| [directive-tech](../entries/directive-tech.yml)                                                                       | directive                                | 指令                                                                       | absent        | 0 / 0                        | NC                   |
| [discard-workflow](../entries/discard-workflow.yml)                                                                   | discard                                  | 丢弃                                                                       | absent        | 0 / 0                        | NC                   |
| [disclosure-workflow](../entries/disclosure-workflow.yml)                                                             | disclosure                               | 披露说明                                                                   | N             | 0 / 1                        | K                    |
| [disguised-combat](../entries/disguised-combat.yml)                                                                   | disguised                                | 伪装的                                                                     | absent        | 0 / 0                        | NC                   |
| [dispatcher-tech](../entries/dispatcher-tech.yml)                                                                     | dispatcher                               | 分派器                                                                     | absent        | 0 / 0                        | NC                   |
| [dispenser-engineering](../entries/dispenser-engineering.yml)                                                         | dispenser                                | 分配器；按需放出物品的装置。                                               | W             | 0 / 0                        | K                    |
| [distinctive-combat](../entries/distinctive-combat.yml)                                                               | distinctive                              | 独特的、显著的                                                             | absent        | 0 / 0                        | NC                   |
| [distinguish-tech](../entries/distinguish-tech.yml)                                                                   | distinguish                              | 区分；辨别                                                                 | absent        | 0 / 0                        | NC                   |
| [dive-general](../entries/dive-general.yml)                                                                           | dive                                     | 俯冲、潜入                                                                 | absent        | 0 / 0                        | NC                   |
| [dodge-general](../entries/dodge-general.yml)                                                                         | dodge                                    | 闪避                                                                       | absent        | 0 / 0                        | NC                   |
| [door-architecture](../entries/door-architecture.yml)                                                                 | door                                     | 门。                                                                       | absent        | 0 / 0                        | NC                   |
| [dot-general](../entries/dot-general.yml)                                                                             | dot                                      | 点；小圆点。                                                               | absent        | 0 / 0                        | NC-N                 |
| [dots-ai](../entries/dots-ai.yml)                                                                                     | dots                                     | OpenAI 的持续运行型个人智能体产品名称，可围绕用户目标推进任务。            | W             | 0 / 0                        | K                    |
| [draconification-zelda](../entries/draconification-zelda.yml)                                                         | draconification                          | 龙化之法；通过吞下秘石化为不死之龙、同时失去自我的禁忌之术。               | absent        | 0 / 0                        | D: D06               |
| [draft-workflow](../entries/draft-workflow.yml)                                                                       | draft                                    | 草稿；尚待修改或审阅的初稿。                                               | absent        | 0 / 0                        | NC                   |
| [dragon-story](../entries/dragon-story.yml)                                                                           | dragon                                   | 龙。                                                                       | absent        | 0 / 0                        | NC                   |
| [draw-general](../entries/draw-general.yml)                                                                           | draw                                     | 拉开（弓）；拔出（武器）。                                                 | absent        | 0 / 0                        | NC-N                 |
| [drop-general](../entries/drop-general.yml)                                                                           | drop                                     | 丢弃                                                                       | absent        | 0 / 0                        | NC                   |
| [durability-equipment](../entries/durability-equipment.yml)                                                           | durability                               | 耐用性；耐久度。                                                           | W             | 0 / 2                        | K                    |
| [durable-general](../entries/durable-general.yml)                                                                     | durable                                  | 耐用的；能经受使用而不易损坏的。                                           | absent        | 1 / 1                        | NC                   |
| [dynamic-gating-tech](../entries/dynamic-gating-tech.yml)                                                             | dynamic gating                           | 动态条件控制                                                               | absent        | 0 / 0                        | NC                   |
| [eagle-general](../entries/eagle-general.yml)                                                                         | eagle                                    | 鹰。                                                                       | absent        | 0 / 0                        | NC                   |
| [edge-case-tech](../entries/edge-case-tech.yml)                                                                       | edge case                                | 边界情况；特殊情况                                                         | absent        | 0 / 0                        | NC                   |
| [effect-react](../entries/effect-react.yml)                                                                           | effect                                   | React Effect；在组件提交后将组件与外部系统同步的逻辑。                     | absent        | 0 / 0                        | NC                   |
| [effect-replay-tech](../entries/effect-replay-tech.yml)                                                               | effect replay                            | 副作用重放                                                                 | absent        | 0 / 0                        | NC                   |
| [egg-food](../entries/egg-food.yml)                                                                                   | egg                                      | 蛋。                                                                       | absent        | 0 / 0                        | NC                   |
| [elapsed-time-workflow](../entries/elapsed-time-workflow.yml)                                                         | elapsed time                             | 已用时间                                                                   | absent        | 0 / 0                        | NC                   |
| [electro-cooking](../entries/electro-cooking.yml)                                                                     | Electro                                  | 电麻；本作料理效果名前缀，表示提供抗电效果。                               | absent        | 0 / 0                        | NC-N                 |
| [eligible-tech](../entries/eligible-tech.yml)                                                                         | eligible                                 | 符合条件的                                                                 | absent        | 0 / 0                        | NC                   |
| [elixir-cooking](../entries/elixir-cooking.yml)                                                                       | elixir                                   | 药剂                                                                       | E             | 0 / 0                        | K                    |
| [elusive-combat](../entries/elusive-combat.yml)                                                                       | elusive                                  | 难以捉摸的                                                                 | absent        | 0 / 0                        | NC                   |
| [emit-tech](../entries/emit-tech.yml)                                                                                 | emit                                     | 生成；输出                                                                 | absent        | 0 / 1                        | NC                   |
| [emitter-engineering](../entries/emitter-engineering.yml)                                                             | emitter                                  | 发射器、发出装置                                                           | W             | 1 / 1                        | K                    |
| [encode-tech](../entries/encode-tech.yml)                                                                             | encode                                   | 编码                                                                       | N             | 1 / 1                        | K                    |
| [endura-carrot-zelda](../entries/endura-carrot-zelda.yml)                                                             | Endura Carrot                            | 毅力胡萝卜；《王国之泪》中的材料名称。                                     | absent        | 0 / 0                        | NC                   |
| [endura-shroom-zelda](../entries/endura-shroom-zelda.yml)                                                             | Endura Shroom                            | 毅力蘑菇；《王国之泪》中的材料名称。                                       | absent        | 0 / 0                        | NC                   |
| [endure-general](../entries/endure-general.yml)                                                                       | endure                                   | 持久；忍受；持续存在。                                                     | absent        | 1 / 1                        | NC                   |
| [energize-general](../entries/energize-general.yml)                                                                   | energize                                 | 使充满活力；补充能量。                                                     | absent        | 0 / 1                        | NC-N                 |
| [enhancement-general](../entries/enhancement-general.yml)                                                             | enhancement                              | 强化、增强                                                                 | W             | 0 / 0                        | K                    |
| [enormous-combat](../entries/enormous-combat.yml)                                                                     | enormous                                 | 巨大的                                                                     | absent        | 0 / 0                        | NC                   |
| [enthusiastic-general](../entries/enthusiastic-general.yml)                                                           | enthusiastic                             | 热情的                                                                     | E             | 0 / 0                        | K                    |
| [entrance-architecture](../entries/entrance-architecture.yml)                                                         | entrance                                 | 入口                                                                       | absent        | 0 / 0                        | NC                   |
| [enum-typescript](../entries/enum-typescript.yml)                                                                     | enum                                     | 枚举；为一组具名常量定义成员的结构。                                       | absent        | 0 / 0                        | NC                   |
| [enumerable-javascript](../entries/enumerable-javascript.yml)                                                         | enumerable                               | 可枚举的；属性可被相应枚举操作列举的特性。                                 | absent        | 0 / 0                        | NC-N                 |
| [equality-semantics-tech](../entries/equality-semantics-tech.yml)                                                     | equality semantics                       | 相等比较语义                                                               | absent        | 0 / 0                        | NC                   |
| [equip-general](../entries/equip-general.yml)                                                                         | equip                                    | 装备                                                                       | absent        | 0 / 0                        | NC                   |
| [equipment-equipment](../entries/equipment-equipment.yml)                                                             | equipment                                | 装备                                                                       | absent        | 0 / 0                        | NC                   |
| [escape-parity-tech](../entries/escape-parity-tech.yml)                                                               | escape parity                            | 转义符数量的奇偶关系                                                       | E             | 0 / 0                        | K                    |
| [escape-tech](../entries/escape-tech.yml)                                                                             | escape                                   | 转义序列                                                                   | E             | 0 / 0                        | K                    |
| [escaped-backslash-tech](../entries/escaped-backslash-tech.yml)                                                       | escaped backslash                        | 转义后的反斜杠                                                             | N             | 0 / 0                        | K                    |
| [etymology-language](../entries/etymology-language.yml)                                                               | etymology                                | 词源；词语的历史来源与演变。                                               | E             | 0 / 0                        | R: R24; K            |
| [event-bubbling-dom](../entries/event-bubbling-dom.yml)                                                               | event bubbling                           | 事件冒泡；可冒泡事件从目标向祖先传播的阶段。                               | absent        | 2 / 2                        | NC-N                 |
| [event-capturing-dom](../entries/event-capturing-dom.yml)                                                             | event capturing                          | 事件捕获；事件沿传播路径向目标传递的阶段。                                 | absent        | 0 / 1                        | NC                   |
| [event-delegation-dom](../entries/event-delegation-dom.yml)                                                           | event delegation                         | 事件委托；在祖先上监听并处理后代元素的事件。                               | absent        | 0 / 1                        | NC                   |
| [event-loop-javascript](../entries/event-loop-javascript.yml)                                                         | event loop                               | 事件循环；协调 JavaScript 任务执行与宿主活动的机制。                       | absent        | 0 / 0                        | NC                   |
| [event-tech](../entries/event-tech.yml)                                                                               | event                                    | 事件；程序可观察并响应的发生或通知。                                       | absent        | 0 / 0                        | NC                   |
| [evermean-zelda](../entries/evermean-zelda.yml)                                                                       | Evermean                                 | 巨霸柏古；《王国之泪》中的怪物或首领名称。                                 | absent        | 0 / 0                        | NC                   |
| [evidence-tech](../entries/evidence-tech.yml)                                                                         | evidence                                 | 证据；依据                                                                 | absent        | 0 / 0                        | NC                   |
| [examine-general](../entries/examine-general.yml)                                                                     | examine                                  | 检查                                                                       | absent        | 0 / 0                        | NC                   |
| [example-general](../entries/example-general.yml)                                                                     | example                                  | 例子；用来说明规则、含义或用法的实例。                                     | absent        | 0 / 0                        | NC-N                 |
| [exclude-tech](../entries/exclude-tech.yml)                                                                           | exclude                                  | 排除                                                                       | absent        | 0 / 0                        | NC                   |
| [exhaustive-tech](../entries/exhaustive-tech.yml)                                                                     | exhaustive                               | 全面的；穷尽的                                                             | absent        | 0 / 0                        | NC                   |
| [exit-architecture](../entries/exit-architecture.yml)                                                                 | exit                                     | 出口                                                                       | absent        | 0 / 0                        | NC                   |
| [exit-status-tech](../entries/exit-status-tech.yml)                                                                   | exit status                              | 退出状态                                                                   | absent        | 0 / 0                        | NC                   |
| [expected-behavior-workflow](../entries/expected-behavior-workflow.yml)                                               | expected behavior                        | 预期行为                                                                   | absent        | 0 / 1                        | NC                   |
| [explicit-tech](../entries/explicit-tech.yml)                                                                         | explicit                                 | 明确的；显式的                                                             | absent        | 0 / 0                        | NC                   |
| [expose-tech](../entries/expose-tech.yml)                                                                             | expose                                   | 暴露；使显现                                                               | absent        | 0 / 0                        | NC                   |
| [factory-architecture](../entries/factory-architecture.yml)                                                           | factory                                  | 工厂。                                                                     | absent        | 0 / 0                        | NC                   |
| [factual-general](../entries/factual-general.yml)                                                                     | factual                                  | 基于事实的                                                                 | absent        | 0 / 0                        | NC                   |
| [fail-against-the-original-implementation-workflow](../entries/fail-against-the-original-implementation-workflow.yml) | fail against the original implementation | 在原始实现上测试失败                                                       | absent        | 0 / 0                        | NC                   |
| [fallback-tech](../entries/fallback.yml)                                                                              | fallback                                 | 备用方案；回退方案                                                         | W             | 0 / 0                        | K                    |
| [false-diagnostic-tech](../entries/false-diagnostic-tech.yml)                                                         | false diagnostic                         | 误报的诊断信息                                                             | absent        | 0 / 0                        | NC                   |
| [fan-engineering](../entries/fan-engineering.yml)                                                                     | fan                                      | 风扇；扇子。                                                               | absent        | 0 / 0                        | NC                   |
| [fast-path-tech](../entries/fast-path-tech.yml)                                                                       | fast path                                | 快速处理路径                                                               | absent        | 0 / 0                        | NC                   |
| [fearless-general](../entries/fearless-general.yml)                                                                   | fearless                                 | 无畏的                                                                     | W             | 0 / 1                        | K                    |
| [field-presence-tech](../entries/field-presence-tech.yml)                                                             | field presence                           | 字段是否存在                                                               | N             | 0 / 0                        | K                    |
| [fierce-deity-armor-zelda](../entries/fierce-deity-armor-zelda.yml)                                                   | Fierce Deity Armor                       | 鬼神服；《王国之泪》中的防具名称。                                         | absent        | 0 / 0                        | NC                   |
| [fierce-general](../entries/fierce-general.yml)                                                                       | fierce                                   | 凶猛的、猛烈的                                                             | absent        | 0 / 0                        | NC                   |
| [fifth-general](../entries/fifth-general.yml)                                                                         | fifth                                    | 第五。                                                                     | absent        | 0 / 0                        | NC                   |
| [find-general](../entries/find-general.yml)                                                                           | find                                     | 找到；发现。                                                               | absent        | 0 / 0                        | NC                   |
| [finding-tech](../entries/finding-tech.yml)                                                                           | finding                                  | 调查发现                                                                   | absent        | 0 / 0                        | NC                   |
| [fire-general](../entries/fire-general.yml)                                                                           | fire                                     | 火。                                                                       | absent        | 0 / 0                        | NC                   |
| [fire-temple-zelda](../entries/fire-temple-zelda.yml)                                                                 | Fire Temple                              | 火之神殿；《王国之泪》中的地点或设施。                                     | absent        | 0 / 0                        | NC                   |
| [fish-food](../entries/fish-food.yml)                                                                                 | fish                                     | 鱼；鱼肉                                                                   | absent        | 0 / 0                        | NC                   |
| [fixture-tech](../entries/fixture-tech.yml)                                                                           | fixture                                  | 测试或演示样例                                                             | absent        | 0 / 0                        | NC                   |
| [flame-emitter-zelda](../entries/flame-emitter-zelda.yml)                                                             | Flame Emitter                            | 火龙头；《王国之泪》中的左纳乌装置。                                       | W             | 0 / 0                        | K                    |
| [flame-general](../entries/flame-general.yml)                                                                         | flame                                    | 火焰。                                                                     | absent        | 0 / 0                        | NC                   |
| [flamebreaker-armor-zelda](../entries/flamebreaker-armor-zelda.yml)                                                   | Flamebreaker Armor                       | 耐火石铠；《王国之泪》中的防具名称。                                       | absent        | 0 / 0                        | NC                   |
| [flexbox-css](../entries/flexbox-css.yml)                                                                             | flexbox                                  | 弹性布局；主要沿一个维度分配空间和排列项目的布局。                         | absent        | 1 / 1                        | NC                   |
| [floating-point-precision-javascript](../entries/floating-point-precision-javascript.yml)                             | floating-point precision                 | 浮点精度；有限位数表示数值时产生的精度限制。                               | N             | 0 / 0                        | K                    |
| [flow-analysis-tech](../entries/flow-analysis-tech.yml)                                                               | flow analysis                            | 流分析                                                                     | absent        | 0 / 0                        | NC                   |
| [flux-construct-zelda](../entries/flux-construct-zelda.yml)                                                           | Flux Construct                           | 方块魔像；《王国之泪》中的怪物或首领名称。                                 | absent        | 0 / 0                        | NC                   |
| [followed-by-workflow](../entries/followed-by-workflow.yml)                                                           | followed by                              | 后面跟着                                                                   | absent        | 0 / 0                        | NC                   |
| [foothill-terrain](../entries/foothill-terrain.yml)                                                                   | foothill                                 | 山麓                                                                       | W             | 0 / 0                        | K                    |
| [forest-terrain](../entries/forest-terrain.yml)                                                                       | forest                                   | 森林                                                                       | absent        | 0 / 0                        | NC                   |
| [forge-architecture](../entries/forge-architecture.yml)                                                               | forge                                    | 锻造设施                                                                   | absent        | 0 / 0                        | NC                   |
| [forget-general](../entries/forget-general.yml)                                                                       | forget                                   | 忘记；未能想起。                                                           | absent        | 0 / 0                        | NC                   |
| [forgotten-temple-zelda](../entries/forgotten-temple-zelda.yml)                                                       | Forgotten Temple                         | 忘却神殿；《王国之泪》中的地点或设施。                                     | absent        | 0 / 0                        | NC                   |
| [formation-general](../entries/formation-general.yml)                                                                 | formation                                | 形成；构成                                                                 | absent        | 1 / 1                        | R: R24               |
| [fortress-architecture](../entries/fortress-architecture.yml)                                                         | fortress                                 | 要塞、堡垒                                                                 | absent        | 0 / 0                        | NC                   |
| [forwarding-tech](../entries/forwarding-tech.yml)                                                                     | forwarding                               | 转发                                                                       | absent        | 0 / 0                        | NC                   |
| [fountain-architecture](../entries/fountain-architecture.yml)                                                         | fountain                                 | 喷泉、泉水设施                                                             | absent        | 0 / 0                        | NC                   |
| [frog-general](../entries/frog-general.yml)                                                                           | frog                                     | 青蛙。                                                                     | absent        | 0 / 0                        | NC                   |
| [froggy-sleeve-zelda](../entries/froggy-sleeve-zelda.yml)                                                             | Froggy Sleeve                            | 蛙之手甲；《王国之泪》中的防具名称。                                       | absent        | 0 / 0                        | NC                   |
| [frost-emitter-zelda](../entries/frost-emitter-zelda.yml)                                                             | Frost Emitter                            | 冰龙头；《王国之泪》中的左纳乌装置。                                       | W             | 0 / 0                        | K                    |
| [frost-general](../entries/frost-general.yml)                                                                         | frost                                    | 霜；冰霜。                                                                 | absent        | 0 / 0                        | NC                   |
| [frox-zelda](../entries/frox-zelda.yml)                                                                               | Frox                                     | 巨霸迦马；《王国之泪》中的怪物或首领名称。                                 | absent        | 0 / 0                        | NC                   |
| [fruit-food](../entries/fruit-food.yml)                                                                               | fruit                                    | 水果。                                                                     | absent        | 0 / 0                        | NC                   |
| [fry-cooking](../entries/fry-cooking.yml)                                                                             | fry                                      | 油煎；油炸；以热油烹饪。                                                   | absent        | 0 / 0                        | NC                   |
| [fuse-general](../entries/fuse-general.yml)                                                                           | fuse                                     | 融合；结合；使物体连接成为一体。                                           | absent        | 1 / 1                        | NC                   |
| [ganondorf-zelda](../entries/ganondorf-zelda.yml)                                                                     | Ganondorf                                | 加侬多夫；本作的魔王。                                                     | absent        | 0 / 0                        | NC                   |
| [garbage-collection-javascript](../entries/garbage-collection-javascript.yml)                                         | garbage collection                       | 垃圾回收；自动回收不再可达的对象所占内存。                                 | absent        | 0 / 0                        | NC-N                 |
| [gate-architecture](../entries/gate-architecture.yml)                                                                 | gate                                     | 大门                                                                       | absent        | 0 / 0                        | NC                   |
| [gear-equipment](../entries/gear-equipment.yml)                                                                       | gear                                     | 装备；为某种活动准备的器材。                                               | absent        | 0 / 0                        | NC                   |
| [gemini-ai](../entries/gemini-ai.yml)                                                                                 | Gemini                                   | Google 的多模态 AI 模型系列名称。                                          | absent        | 0 / 0                        | NC-N                 |
| [general-store-architecture](../entries/general-store-architecture.yml)                                               | general store                            | 杂货店                                                                     | absent        | 0 / 0                        | NC                   |
| [generalize-general](../entries/generalize-general.yml)                                                               | generalize                               | 推广；概括                                                                 | absent        | 0 / 0                        | NC                   |
| [generated-output-only-workflow](../entries/generated-output-only-workflow.yml)                                       | generated output only                    | 仅限生成结果                                                               | absent        | 0 / 0                        | NC-N                 |
| [genuine-tech](../entries/genuine-tech.yml)                                                                           | genuine                                  | 真正的                                                                     | absent        | 0 / 0                        | NC                   |
| [geoglyph-game](../entries/geoglyph-game.yml)                                                                         | geoglyph                                 | 地上绘；通过地表图案或线条形成的大型图像。                                 | absent        | 0 / 0                        | NC                   |
| [gerudo-desert-zelda](../entries/gerudo-desert-zelda.yml)                                                             | Gerudo Desert                            | 格鲁德沙漠；《王国之泪》中的地点或设施。                                   | absent        | 0 / 0                        | NC                   |
| [gerudo-town-zelda](../entries/gerudo-town-zelda.yml)                                                                 | Gerudo Town                              | 格鲁德小镇；《王国之泪》中的地点或设施。                                   | absent        | 0 / 0                        | NC                   |
| [gerudo-zelda](../entries/gerudo-zelda.yml)                                                                           | Gerudo                                   | 格鲁德族；《王国之泪》中的族群或组织。                                     | absent        | 0 / 0                        | NC                   |
| [getter-tech](../entries/getter-tech.yml)                                                                             | getter                                   | 取值访问器                                                                 | absent        | 0 / 0                        | NC                   |
| [gibdo-zelda](../entries/gibdo-zelda.yml)                                                                             | Gibdo                                    | 吉波得；《王国之泪》中的怪物或首领名称。                                   | absent        | 0 / 0                        | NC                   |
| [gleeok-zelda](../entries/gleeok-zelda.yml)                                                                           | Gleeok                                   | 古栗欧克；《王国之泪》中的怪物或首领名称。                                 | absent        | 0 / 0                        | NC                   |
| [glide-general](../entries/glide-general.yml)                                                                         | glide                                    | 滑翔                                                                       | absent        | 0 / 0                        | NC                   |
| [glide-shirt-zelda](../entries/glide-shirt-zelda.yml)                                                                 | Glide Shirt                              | 滑翔紧身服；《王国之泪》中的防具名称。                                     | absent        | 0 / 0                        | NC                   |
| [global-augmentation-typescript](../entries/global-augmentation-typescript.yml)                                       | global augmentation                      | 全局扩充；在声明中补充已有全局类型或接口的信息。                           | absent        | 0 / 0                        | NC                   |
| [gloom-recovery-zelda](../entries/gloom-recovery-zelda.yml)                                                           | Gloom Recovery                           | 瘴气伤害回复；《王国之泪》中的能力效果名称。                               | absent        | 0 / 0                        | NC                   |
| [gloom-resistance-zelda](../entries/gloom-resistance-zelda.yml)                                                       | Gloom Resistance                         | 瘴气防护；《王国之泪》中的能力效果名称。                                   | absent        | 0 / 0                        | NC                   |
| [gloom-spawn-zelda](../entries/gloom-spawn-zelda.yml)                                                                 | Gloom Spawn                              | 瘴气魔；《王国之泪》中的怪物或首领名称。                                   | absent        | 0 / 0                        | NC                   |
| [gloom-sword-zelda](../entries/gloom-sword-zelda.yml)                                                                 | Gloom Sword                              | 瘴气之剑；《王国之泪》中的装备名称。                                       | absent        | 0 / 0                        | NC                   |
| [gloom-zelda](../entries/gloom-zelda.yml)                                                                             | Gloom                                    | 瘴气；本作中会造成特殊生命损伤的危险现象。                                 | absent        | 0 / 0                        | NC-N                 |
| [golden-apple-zelda](../entries/golden-apple-zelda.yml)                                                               | Golden Apple                             | 金苹果；《王国之泪》中的材料名称。                                         | absent        | 0 / 0                        | NC                   |
| [golden-general](../entries/golden-general.yml)                                                                       | golden                                   | 金色的；金黄色的。                                                         | absent        | 0 / 0                        | NC                   |
| [goron-city-zelda](../entries/goron-city-zelda.yml)                                                                   | Goron City                               | 鼓隆城；《王国之泪》中的地点或设施。                                       | absent        | 0 / 0                        | NC                   |
| [goron-spice-zelda](../entries/goron-spice-zelda.yml)                                                                 | Goron Spice                              | 鼓隆的调味粉；《王国之泪》中的材料名称。                                   | absent        | 0 / 0                        | NC                   |
| [goron-zelda](../entries/goron-zelda.yml)                                                                             | Goron                                    | 鼓隆族；《王国之泪》中的族群或组织。                                       | absent        | 0 / 0                        | NC                   |
| [gourmet-food](../entries/gourmet-food.yml)                                                                           | gourmet                                  | 美食的；讲究精美饮食的。                                                   | absent        | 0 / 0                        | NC                   |
| [graceful-degradation-web](../entries/graceful-degradation-web.yml)                                                   | graceful degradation                     | 优雅降级；在能力受限或部分功能失败时保留核心体验。                         | absent        | 1 / 1                        | NC                   |
| [graceful-shutdown-tech](../entries/graceful-shutdown-tech.yml)                                                       | graceful shutdown                        | 正常有序关闭                                                               | absent        | 0 / 0                        | NC-N                 |
| [gratitude-general](../entries/gratitude-general.yml)                                                                 | gratitude                                | 感谢之情                                                                   | absent        | 0 / 0                        | NC                   |
| [great-abandoned-central-mine-zelda](../entries/great-abandoned-central-mine-zelda.yml)                               | Great Abandoned Central Mine             | 中央大废矿；《王国之泪》中的地点或设施。                                   | absent        | 0 / 0                        | NC                   |
| [great-deku-tree-zelda](../entries/great-deku-tree-zelda.yml)                                                         | Great Deku Tree                          | 德库树长老；克洛格森林中的巨树。                                           | absent        | 0 / 0                        | NC                   |
| [great-eagle-bow-zelda](../entries/great-eagle-bow-zelda.yml)                                                         | Great Eagle Bow                          | 大鹫弓；《王国之泪》中的装备名称。                                         | absent        | 0 / 0                        | NC                   |
| [great-sky-island-zelda](../entries/great-sky-island-zelda.yml)                                                       | Great Sky Island                         | 初始空岛；《王国之泪》中的地点或设施。                                     | absent        | 0 / 0                        | NC                   |
| [greaves-clothing](../entries/greaves-clothing.yml)                                                                   | greaves                                  | 护胫                                                                       | absent        | 0 / 0                        | NC                   |
| [greens-food](../entries/greens-food.yml)                                                                             | greens                                   | 绿叶蔬菜。                                                                 | absent        | 0 / 0                        | NC                   |
| [grid-layout-css](../entries/grid-layout-css.yml)                                                                     | grid layout                              | 网格布局；通过行和列组织项目的二维布局。                                   | absent        | 0 / 1                        | NC                   |
| [grill-cooking](../entries/grill-cooking.yml)                                                                         | grill                                    | 炙烤；在烤架上烹饪。                                                       | N             | 1 / 1                        | K                    |
| [grok-ai](../entries/grok-ai.yml)                                                                                     | Grok                                     | Grok 系列 AI 模型及助手的名称。                                            | absent        | 0 / 0                        | NC-N                 |
| [guard-tech](../entries/guard-tech.yml)                                                                               | guard                                    | 保护性检查或机制                                                           | absent        | 0 / 0                        | NC                   |
| [guardian-general](../entries/guardian-general.yml)                                                                   | guardian                                 | 守护者                                                                     | absent        | 0 / 0                        | NC                   |
| [guidance-story](../entries/guidance-story.yml)                                                                       | guidance                                 | 指引、指导                                                                 | absent        | 0 / 0                        | NC                   |
| [halberd-weapon](../entries/halberd-weapon.yml)                                                                       | halberd                                  | 戟                                                                         | absent        | 0 / 0                        | NC                   |
| [hammer-weapon](../entries/hammer-weapon.yml)                                                                         | hammer                                   | 锤子                                                                       | absent        | 0 / 0                        | NC                   |
| [hand-general](../entries/hand-general.yml)                                                                           | hand                                     | 手。                                                                       | absent        | 0 / 0                        | NC                   |
| [handoff-workflow](../entries/handoff-workflow.yml)                                                                   | handoff                                  | 交接                                                                       | W             | 0 / 0                        | K                    |
| [harness-ai](../entries/harness-ai.yml)                                                                               | harness                                  | 智能体运行框架；连接模型、上下文和工具，并协调执行循环与会话状态的软件层。 | absent        | 1 / 1                        | NC-N                 |
| [harness-tech](../entries/harness-tech.yml)                                                                           | harness                                  | 测试驱动程序                                                               | absent        | 0 / 1                        | NC                   |
| [hashing-tech](../entries/hashing-tech.yml)                                                                           | hashing                                  | 哈希计算                                                                   | absent        | 0 / 0                        | NC                   |
| [hasty-general](../entries/hasty-general.yml)                                                                         | hasty                                    | 匆忙的；仓促的。                                                           | absent        | 0 / 0                        | NC                   |
| [hateno-village-zelda](../entries/hateno-village-zelda.yml)                                                           | Hateno Village                           | 哈特诺村；《王国之泪》中的地点或设施。                                     | absent        | 0 / 0                        | NC                   |
| [head-general](../entries/head-general.yml)                                                                           | head                                     | 头；头部。                                                                 | absent        | 0 / 0                        | NC                   |
| [headband-clothing](../entries/headband-clothing.yml)                                                                 | headband                                 | 头带                                                                       | W             | 1 / 1                        | K                    |
| [headdress-clothing](../entries/headdress-clothing.yml)                                                               | headdress                                | 头饰                                                                       | W             | 0 / 1                        | K                    |
| [heap-memory-javascript](../entries/heap-memory-javascript.yml)                                                       | heap memory                              | 堆内存；运行时管理动态分配对象的内存区域。                                 | absent        | 0 / 0                        | NC                   |
| [hearty-general](../entries/hearty-general.yml)                                                                       | hearty                                   | 丰盛的；有益健康的；充满活力的。                                           | absent        | 0 / 0                        | NC                   |
| [hearty-radish-zelda](../entries/hearty-radish-zelda.yml)                                                             | Hearty Radish                            | 生命小萝卜；《王国之泪》中的材料名称。                                     | absent        | 0 / 0                        | NC                   |
| [hearty-salmon-zelda](../entries/hearty-salmon-zelda.yml)                                                             | Hearty Salmon                            | 生命鲑鱼；《王国之泪》中的材料名称。                                       | absent        | 0 / 0                        | NC                   |
| [hearty-truffle-zelda](../entries/hearty-truffle-zelda.yml)                                                           | Hearty Truffle                           | 生命松露；《王国之泪》中的材料名称。                                       | absent        | 0 / 0                        | NC                   |
| [heat-resistance-combat](../entries/heat-resistance-combat.yml)                                                       | heat resistance                          | 耐热；承受酷暑环境影响的能力。                                             | absent        | 0 / 1                        | NC-N                 |
| [heir-general](../entries/heir-general.yml)                                                                           | heir                                     | 继承人                                                                     | N             | 1 / 1                        | K                    |
| [helm-clothing](../entries/helm-clothing.yml)                                                                         | helm                                     | 头盔                                                                       | absent        | 1 / 1                        | NC                   |
| [helmet-clothing](../entries/helmet-clothing.yml)                                                                     | helmet                                   | 头盔                                                                       | absent        | 0 / 1                        | NC                   |
| [heritage-type-tech](../entries/heritage-type-tech.yml)                                                               | heritage type                            | 继承或实现关系中的类型                                                     | absent        | 0 / 0                        | NC                   |
| [hermes-agent-ai](../entries/hermes-agent-ai.yml)                                                                     | Hermes                                   | 此处指 Nous Research 的 Hermes Agent，开源 AI 智能体项目。                 | absent        | 0 / 0                        | D: D06               |
| [hero-general](../entries/hero-general.yml)                                                                           | hero                                     | 英雄                                                                       | absent        | 0 / 0                        | NC                   |
| [hestu-zelda](../entries/hestu-zelda.yml)                                                                             | Hestu                                    | 伯库林；帮助扩充背包容量的克洛格。                                         | absent        | 0 / 0                        | NC                   |
| [hidden-general](../entries/hidden-general.yml)                                                                       | hidden                                   | 隐藏的                                                                     | absent        | 0 / 0                        | NC                   |
| [hideout-architecture](../entries/hideout-architecture.yml)                                                           | hideout                                  | 藏身处；秘密据点。                                                         | absent        | 0 / 0                        | NC                   |
| [higher-order-function-javascript](../entries/higher-order-function-javascript.yml)                                   | higher-order function                    | 高阶函数；接收函数作为参数或返回函数的函数。                               | absent        | 0 / 0                        | NC                   |
| [highlands-terrain](../entries/highlands-terrain.yml)                                                                 | highlands                                | 高地                                                                       | W             | 0 / 0                        | K                    |
| [hinox-zelda](../entries/hinox-zelda.yml)                                                                             | Hinox                                    | 西诺克斯；《王国之泪》中的怪物或首领名称。                                 | absent        | 0 / 0                        | NC                   |
| [hoisting-javascript](../entries/hoisting-javascript.yml)                                                             | hoisting                                 | 提升；声明在执行前影响所在作用域的行为。                                   | absent        | 1 / 1                        | NC-N                 |
| [hold-general](../entries/hold-general.yml)                                                                           | hold                                     | 手持、持有                                                                 | absent        | 0 / 0                        | NC                   |
| [homing-cart-zelda](../entries/homing-cart-zelda.yml)                                                                 | Homing Cart                              | 追踪平台车；《王国之泪》中的左纳乌装置。                                   | absent        | 0 / 0                        | NC                   |
| [homing-engineering](../entries/homing-engineering.yml)                                                               | homing                                   | 自动导引的；能自动追踪目标的。                                             | W             | 0 / 0                        | K                    |
| [honey-food](../entries/honey-food.yml)                                                                               | honey                                    | 蜂蜜。                                                                     | absent        | 0 / 0                        | NC                   |
| [honorable-general](../entries/honorable-general.yml)                                                                 | honorable                                | 值得尊敬的                                                                 | absent        | 0 / 0                        | NC                   |
| [hood-clothing](../entries/hood-clothing.yml)                                                                         | hood                                     | 兜帽                                                                       | absent        | 0 / 0                        | NC                   |
| [horriblin-zelda](../entries/horriblin-zelda.yml)                                                                     | Horriblin                                | 霍拉布林；《王国之泪》中的怪物或首领名称。                                 | absent        | 0 / 0                        | NC                   |
| [hostile-combat](../entries/hostile-combat.yml)                                                                       | hostile                                  | 敌对的                                                                     | E,N           | 1 / 1                        | K                    |
| [hover-general](../entries/hover-general.yml)                                                                         | hover                                    | 悬停；在空中保持相近位置。                                                 | absent        | 0 / 0                        | NC                   |
| [hover-stone-zelda](../entries/hover-stone-zelda.yml)                                                                 | Hover Stone                              | 浮游石；《王国之泪》中的左纳乌装置。                                       | W             | 0 / 0                        | K                    |
| [http-caching-web](../entries/http-caching-web.yml)                                                                   | HTTP caching                             | HTTP 缓存；按缓存规则复用已存储的响应。                                    | N             | 0 / 0                        | K                    |
| [hudson-zelda](../entries/hudson-zelda.yml)                                                                           | Hudson                                   | 松达；从事建造事业的人物。                                                 | absent        | 0 / 0                        | NC                   |
| [hydrant-engineering](../entries/hydrant-engineering.yml)                                                             | hydrant                                  | 消防栓；连接供水管道的取水装置。                                           | E             | 0 / 0                        | K                    |
| [hylian-hood-zelda](../entries/hylian-hood-zelda.yml)                                                                 | Hylian Hood                              | 海利亚兜帽；《王国之泪》中的防具名称。                                     | absent        | 0 / 0                        | NC                   |
| [hylian-shield-zelda](../entries/hylian-shield-zelda.yml)                                                             | Hylian Shield                            | 海利亚盾；《王国之泪》中的装备名称。                                       | absent        | 0 / 0                        | NC                   |
| [hylian-tunic-zelda](../entries/hylian-tunic-zelda.yml)                                                               | Hylian Tunic                             | 海利亚服；《王国之泪》中的防具名称。                                       | absent        | 0 / 0                        | NC                   |
| [hylian-zelda](../entries/hylian-zelda.yml)                                                                           | Hylian                                   | 海利亚人；《王国之泪》中的族群或组织。                                     | absent        | 0 / 0                        | NC                   |
| [hyrule-bass-zelda](../entries/hyrule-bass-zelda.yml)                                                                 | Hyrule Bass                              | 海拉鲁鲈鱼；《王国之泪》中的材料名称。                                     | absent        | 0 / 0                        | NC                   |
| [hyrule-castle-zelda](../entries/hyrule-castle-zelda.yml)                                                             | Hyrule Castle                            | 海拉鲁城堡；《王国之泪》中的地点或设施。                                   | absent        | 0 / 0                        | NC                   |
| [hyrule-compendium-zelda](../entries/hyrule-compendium-zelda.yml)                                                     | Hyrule Compendium                        | 海拉鲁图鉴；《王国之泪》中的界面类别或功能名称。                           | absent        | 0 / 0                        | NC                   |
| [hyrule-zelda](../entries/hyrule-zelda.yml)                                                                           | Hyrule                                   | 海拉鲁；海拉鲁王国，本作冒险的主要舞台。                                   | absent        | 0 / 0                        | NC                   |
| [identifier-tech](../entries/identifier-tech.yml)                                                                     | identifier                               | 标识符                                                                     | absent        | 0 / 0                        | A: WF04              |
| [idiomatic-tech](../entries/idiomatic.yml)                                                                            | idiomatic                                | 地道的；惯用的                                                             | E             | 0 / 0                        | K                    |
| [impa-zelda](../entries/impa-zelda.yml)                                                                               | Impa                                     | 英帕；希卡族长者。                                                         | absent        | 0 / 0                        | NC                   |
| [imperative-general](../entries/imperative-general.yml)                                                               | imperative                               | 祈使语气的                                                                 | absent        | 0 / 0                        | NC                   |
| [implement-tech](../entries/implement-tech.yml)                                                                       | implement                                | 实现；把设计、规则或接口落实为可执行行为。                                 | absent        | 0 / 0                        | NC                   |
| [imply-general](../entries/imply-general.yml)                                                                         | imply                                    | 暗示                                                                       | N             | 0 / 1                        | K                    |
| [imprison-story](../entries/imprison-story.yml)                                                                       | imprison                                 | 囚禁；监禁。                                                               | absent        | 0 / 0                        | NC-N                 |
| [in-one-pass-workflow](../entries/in-one-pass-workflow.yml)                                                           | in one pass                              | 通过一遍处理                                                               | absent        | 0 / 0                        | NC                   |
| [in-order-workflow](../entries/in-order-workflow.yml)                                                                 | in order                                 | 按顺序                                                                     | absent        | 0 / 0                        | NC                   |
| [inaccessible-tech](../entries/inaccessible-tech.yml)                                                                 | inaccessible                             | 不可访问的                                                                 | W             | 0 / 0                        | K                    |
| [incremental-build-tech](../entries/incremental-build-tech.yml)                                                       | incremental build                        | 增量构建                                                                   | absent        | 0 / 0                        | NC                   |
| [independently-tech](../entries/independently-tech.yml)                                                               | independently                            | 独立地                                                                     | W             | 0 / 0                        | K                    |
| [infer-tech](../entries/infer-tech.yml)                                                                               | infer                                    | 推断                                                                       | N             | 1 / 1                        | K                    |
| [ingredient-cooking](../entries/ingredient-cooking.yml)                                                               | ingredient                               | 食材、原料                                                                 | absent        | 1 / 1                        | NC                   |
| [inherit-story](../entries/inherit-story.yml)                                                                         | inherit                                  | 继承                                                                       | absent        | 0 / 0                        | NC                   |
| [initializer-tech](../entries/initializer-tech.yml)                                                                   | initializer                              | 初始化表达式；在创建或声明值时给出初始值的表达式。                         | absent        | 0 / 0                        | D: D02               |
| [inn-architecture](../entries/inn-architecture.yml)                                                                   | inn                                      | 旅馆                                                                       | absent        | 0 / 0                        | NC                   |
| [inspect-tech](../entries/inspect-tech.yml)                                                                           | inspect                                  | 检查；查看                                                                 | absent        | 0 / 0                        | A: ETY08; R: R22     |
| [intact-tech](../entries/intact-tech.yml)                                                                             | intact                                   | 完整的；未受破坏的                                                         | absent        | 0 / 0                        | NC                   |
| [integration-test-tech](../entries/integration-test-tech.yml)                                                         | integration test                         | 集成测试                                                                   | absent        | 0 / 0                        | NC                   |
| [integrity-general](../entries/integrity-general.yml)                                                                 | integrity                                | 正直；诚信；坚持诚实和道德原则。                                           | E             | 0 / 0                        | K                    |
| [interact-general](../entries/interact-general.yml)                                                                   | interact                                 | 交互                                                                       | absent        | 0 / 0                        | NC                   |
| [internals-tech](../entries/internals-tech.yml)                                                                       | internals                                | 内部实现                                                                   | absent        | 0 / 0                        | NC-N                 |
| [introduce-tech](../entries/introduce-tech.yml)                                                                       | introduce                                | 引入                                                                       | absent        | 0 / 0                        | NC                   |
| [invalidate-tech](../entries/invalidate-tech.yml)                                                                     | invalidate                               | 使失效                                                                     | absent        | 0 / 0                        | NC                   |
| [invaluable-general](../entries/invaluable-general.yml)                                                               | invaluable                               | 极有价值的                                                                 | N             | 0 / 0                        | K                    |
| [inventory-ui](../entries/inventory-ui.yml)                                                                           | inventory                                | 物品清单；库存；游戏中的物品栏。                                           | N             | 1 / 1                        | K                    |
| [investigate-tech](../entries/investigate-tech.yml)                                                                   | investigate                              | 调查；研究问题                                                             | absent        | 0 / 0                        | NC                   |
| [invocation-workflow](../entries/invocation-workflow.yml)                                                             | invocation                               | 一次调用                                                                   | absent        | 0 / 0                        | A: WF07; R: R03      |
| [invoke-tech](../entries/invoke-tech.yml)                                                                             | invoke                                   | 调用                                                                       | absent        | 0 / 0                        | R: R03               |
| [island-terrain](../entries/island-terrain.yml)                                                                       | island                                   | 岛屿                                                                       | absent        | 0 / 0                        | NC                   |
| [isolated-worktree-workflow](../entries/isolated-worktree-workflow.yml)                                               | isolated worktree                        | 隔离工作树                                                                 | absent        | 0 / 0                        | NC                   |
| [iterable-tech](../entries/iterable-tech.yml)                                                                         | iterable                                 | 可迭代的                                                                   | absent        | 0 / 0                        | NC                   |
| [jar-general](../entries/jar-general.yml)                                                                             | jar                                      | 罐；广口瓶。                                                               | absent        | 0 / 0                        | NC                   |
| [joiner-tech](../entries/joiner-tech.yml)                                                                             | joiner                                   | 连接控制字符                                                               | absent        | 0 / 0                        | NC                   |
| [josha-zelda](../entries/josha-zelda.yml)                                                                             | Josha                                    | 乔舒亚；研究地底的调查人员。                                               | absent        | 0 / 0                        | NC                   |
| [justify-general](../entries/justify-general.yml)                                                                     | justify                                  | 证明合理；提供充分理由                                                     | absent        | 0 / 0                        | NC                   |
| [kakariko-village-zelda](../entries/kakariko-village-zelda.yml)                                                       | Kakariko Village                         | 卡卡利科村；《王国之泪》中的地点或设施。                                   | absent        | 0 / 0                        | NC                   |
| [keese-zelda](../entries/keese-zelda.yml)                                                                             | Keese                                    | 蝙蝠；《王国之泪》中的怪物或首领名称。                                     | absent        | 0 / 0                        | NC                   |
| [key-items-zelda](../entries/key-items-zelda.yml)                                                                     | Key Items                                | 重要物品；《王国之泪》中的界面类别或功能名称。                             | absent        | 0 / 0                        | NC                   |
| [king-general](../entries/king-general.yml)                                                                           | king                                     | 国王                                                                       | absent        | 0 / 0                        | NC                   |
| [king-gleeok-zelda](../entries/king-gleeok-zelda.yml)                                                                 | King Gleeok                              | 古栗欧克王；《王国之泪》中的怪物或首领名称。                               | absent        | 0 / 0                        | NC                   |
| [kingdom-general](../entries/kingdom-general.yml)                                                                     | kingdom                                  | 王国；由国王或女王统治的国家。                                             | W             | 0 / 0                        | K                    |
| [knight-general](../entries/knight-general.yml)                                                                       | knight                                   | 骑士。                                                                     | absent        | 0 / 0                        | NC                   |
| [knights-broadsword-zelda](../entries/knights-broadsword-zelda.yml)                                                   | Knight's Broadsword                      | 骑士之剑；《王国之泪》中的装备名称。                                       | absent        | 0 / 0                        | NC                   |
| [korok-forest-zelda](../entries/korok-forest-zelda.yml)                                                               | Korok Forest                             | 克洛格森林；《王国之泪》中的地点或设施。                                   | absent        | 0 / 0                        | NC                   |
| [korok-seed-zelda](../entries/korok-seed-zelda.yml)                                                                   | Korok Seed                               | 克洛格的果实；与克洛格和背包扩充有关的收集物品。                           | absent        | 0 / 0                        | NC                   |
| [korok-zelda](../entries/korok-zelda.yml)                                                                             | Korok                                    | 克洛格族；《王国之泪》中的族群或组织。                                     | absent        | 0 / 0                        | NC                   |
| [lake-terrain](../entries/lake-terrain.yml)                                                                           | lake                                     | 湖泊                                                                       | absent        | 0 / 0                        | NC                   |
| [landing-general](../entries/landing-general.yml)                                                                     | landing                                  | 着陆；登陆；登岸处。                                                       | absent        | 0 / 0                        | NC                   |
| [language-general](../entries/language-general.yml)                                                                   | language                                 | 语言                                                                       | absent        | 0 / 0                        | NC                   |
| [lava-terrain](../entries/lava-terrain.yml)                                                                           | lava                                     | 熔岩                                                                       | absent        | 0 / 0                        | NC                   |
| [layout-browser-rendering](../entries/layout-browser-rendering.yml)                                                   | layout                                   | 布局；根据样式与约束计算页面元素的位置和尺寸。                             | absent        | 1 / 1                        | NC                   |
| [least-recently-used-cache](../entries/least-recently-used-cache.yml)                                                 | least recently used                      | 最近最少使用；容量不足时优先淘汰最久未访问数据的策略。                     | N             | 0 / 0                        | K                    |
| [leave-untouched-workflow](../entries/leave-untouched-workflow.yml)                                                   | leave untouched                          | 保持原样；不修改                                                           | absent        | 0 / 0                        | NC                   |
| [legacy-tech](../entries/legacy-tech.yml)                                                                             | legacy                                   | 旧版的；历史遗留的                                                         | absent        | 0 / 0                        | D: D03               |
| [leggings-clothing](../entries/leggings-clothing.yml)                                                                 | leggings                                 | 紧身裤、护腿                                                               | absent        | 0 / 0                        | NC                   |
| [lifecycle-tech](../entries/lifecycle-tech.yml)                                                                       | lifecycle                                | 生命周期                                                                   | W             | 0 / 0                        | K                    |
| [light-dragon-zelda](../entries/light-dragon-zelda.yml)                                                               | Light Dragon                             | 白龙；本作在空中飞行的一条龙。                                             | absent        | 0 / 0                        | NC                   |
| [light-engineering](../entries/light-engineering.yml)                                                                 | light                                    | 光；光源；照明装置。                                                       | absent        | 0 / 0                        | NC                   |
| [light-of-blessing-zelda](../entries/light-of-blessing-zelda.yml)                                                     | Light of Blessing                        | 祝福之光；通过神庙挑战后获得的重要物品。                                   | absent        | 0 / 0                        | NC                   |
| [lightning-temple-zelda](../entries/lightning-temple-zelda.yml)                                                       | Lightning Temple                         | 雷之神殿；《王国之泪》中的地点或设施。                                     | absent        | 0 / 0                        | NC                   |
| [lightning-terrain](../entries/lightning-terrain.yml)                                                                 | lightning                                | 闪电。                                                                     | absent        | 0 / 0                        | NC                   |
| [lightroot-zelda](../entries/lightroot-zelda.yml)                                                                     | Lightroot                                | 破魔之根；地底可激活并照亮周围的特殊根状地点。                             | W             | 0 / 0                        | K                    |
| [like-like-zelda](../entries/like-like-zelda.yml)                                                                     | Like Like                                | 莱克莱克；《王国之泪》中的怪物或首领名称。                                 | absent        | 0 / 0                        | NC                   |
| [limitation-tech](../entries/limitation-tech.yml)                                                                     | limitation                               | 限制；局限                                                                 | absent        | 0 / 0                        | NC                   |
| [linear-tech](../entries/linear-tech.yml)                                                                             | linear                                   | 线性增长的                                                                 | absent        | 0 / 1                        | NC                   |
| [link-zelda](../entries/link-zelda.yml)                                                                               | Link                                     | 林克；《王国之泪》的主角。                                                 | absent        | 0 / 0                        | NC                   |
| [linked-list-merge-algorithm](../entries/linked-list-merge-algorithm.yml)                                             | linked list merge                        | 链表合并；比较两个有序链表的当前节点，依次连接较小节点。                   | absent        | 0 / 0                        | NC                   |
| [linked-list-tech](../entries/linked-list-tech.yml)                                                                   | linked list                              | 链表；通过节点之间的链接组织元素的线性结构。                               | absent        | 0 / 0                        | NC                   |
| [literal-programming](../entries/literal-programming.yml)                                                             | literal                                  | 字面量                                                                     | N             | 1 / 1                        | K                    |
| [literal-text](../entries/literal-text.yml)                                                                           | literal                                  | 按原样解释的；字面的                                                       | N             | 0 / 1                        | K                    |
| [lizalfos-zelda](../entries/lizalfos-zelda.yml)                                                                       | Lizalfos                                 | 蜥蜴战士；《王国之泪》中的怪物或首领名称。                                 | absent        | 0 / 0                        | NC                   |
| [lockfile-tech](../entries/lockfile-tech.yml)                                                                         | lockfile                                 | 依赖锁定文件                                                               | W             | 0 / 0                        | K                    |
| [lone-surrogate-tech](../entries/lone-surrogate-tech.yml)                                                             | lone surrogate                           | 未配对的代理码元                                                           | absent        | 1 / 1                        | NC                   |
| [lookout-general](../entries/lookout-general.yml)                                                                     | lookout                                  | 瞭望；警戒人员；瞭望处。                                                   | absent        | 0 / 0                        | NC                   |
| [lookout-landing-zelda](../entries/lookout-landing-zelda.yml)                                                         | Lookout Landing                          | 监视堡垒；《王国之泪》中的地点或设施。                                     | absent        | 0 / 0                        | NC                   |
| [lowering-compiler](../entries/lowering-compiler.yml)                                                                 | lowering                                 | 降级转换；把较高层程序表示转换为更低层表示的编译过程。                     | absent        | 0 / 0                        | NC-N                 |
| [loyal-general](../entries/loyal-general.yml)                                                                         | loyal                                    | 忠诚的                                                                     | absent        | 0 / 0                        | NC                   |
| [lurelin-village-zelda](../entries/lurelin-village-zelda.yml)                                                         | Lurelin Village                          | 沃托里村；《王国之泪》中的地点或设施。                                     | absent        | 0 / 0                        | NC                   |
| [lynel-zelda](../entries/lynel-zelda.yml)                                                                             | Lynel                                    | 莱尼尔；《王国之泪》中的怪物或首领名称。                                   | absent        | 0 / 0                        | NC                   |
| [magic-rod-zelda](../entries/magic-rod-zelda.yml)                                                                     | Magic Rod                                | 魔法杖；《王国之泪》中的装备名称。                                         | absent        | 0 / 0                        | NC                   |
| [main-quests-zelda](../entries/main-quests-zelda.yml)                                                                 | Main Quests                              | 主剧情挑战；《王国之泪》中的界面类别或功能名称。                           | absent        | 0 / 0                        | NC                   |
| [maintainable-tech](../entries/maintainable-tech.yml)                                                                 | maintainable                             | 易维护的                                                                   | W             | 0 / 0                        | K                    |
| [majestic-general](../entries/majestic-general.yml)                                                                   | majestic                                 | 雄伟的、威严的                                                             | absent        | 0 / 0                        | NC                   |
| [malevolent-general](../entries/malevolent-general.yml)                                                               | malevolent                               | 邪恶的、怀有恶意的                                                         | E             | 0 / 0                        | K                    |
| [malformed-tech](../entries/malformed-tech.yml)                                                                       | malformed                                | 格式错误的                                                                 | W             | 0 / 0                        | K                    |
| [manifest-tech](../entries/manifest-tech.yml)                                                                         | manifest                                 | 清单文件                                                                   | absent        | 0 / 0                        | NC                   |
| [map-ui](../entries/map-ui.yml)                                                                                       | map                                      | 地图。                                                                     | absent        | 0 / 0                        | NC                   |
| [mapped-type-typescript](../entries/mapped-type-typescript.yml)                                                       | mapped type                              | 映射类型；遍历一组属性键，为各键构造属性类型的 TypeScript 类型写法。       | absent        | 0 / 0                        | NC-N                 |
| [marbled-gohma-zelda](../entries/marbled-gohma-zelda.yml)                                                             | Marbled Gohma                            | 巨砾哥马；《王国之泪》中的怪物或首领名称。                                 | absent        | 0 / 0                        | NC                   |
| [marker-tech](../entries/marker-tech.yml)                                                                             | marker                                   | 标记                                                                       | absent        | 0 / 0                        | NC                   |
| [mashed-food](../entries/mashed-food.yml)                                                                             | mashed                                   | 捣碎的；做成泥的                                                           | absent        | 0 / 0                        | NC                   |
| [mashed-potatoes-food](../entries/mashed-potatoes-food.yml)                                                           | mashed potatoes                          | 土豆泥                                                                     | absent        | 0 / 0                        | NC                   |
| [mask-clothing](../entries/mask-clothing.yml)                                                                         | mask                                     | 面具                                                                       | absent        | 0 / 0                        | NC                   |
| [master-kohga-zelda](../entries/master-kohga-zelda.yml)                                                               | Master Kohga                             | 可盖大人；依盖队的首领。                                                   | absent        | 0 / 0                        | NC                   |
| [master-sword-zelda](../entries/master-sword-zelda.yml)                                                               | Master Sword                             | 大师之剑；《王国之泪》中的装备名称。                                       | absent        | 0 / 0                        | NC                   |
| [material-general](../entries/material-general.yml)                                                                   | material                                 | 材料；用于制作物品的原料。                                                 | absent        | 0 / 0                        | NC                   |
| [material-tech](../entries/material-tech.yml)                                                                         | material                                 | 实质性的；影响判断的                                                       | absent        | 0 / 0                        | NC                   |
| [meal-cooking](../entries/meal-cooking.yml)                                                                           | meal                                     | 一餐、料理                                                                 | absent        | 0 / 0                        | NC                   |
| [meat-food](../entries/meat-food.yml)                                                                                 | meat                                     | 肉；供食用的肉类。                                                         | absent        | 0 / 0                        | NC                   |
| [melee-combat](../entries/melee-combat.yml)                                                                           | melee                                    | 近战                                                                       | E             | 0 / 0                        | K                    |
| [memoization-programming](../entries/memoization-programming.yml)                                                     | memoization                              | 记忆化；缓存计算结果，在输入满足复用条件时避免重复计算。                   | absent        | 0 / 1                        | NC-N                 |
| [memory-story](../entries/memory-story.yml)                                                                           | memory                                   | 记忆                                                                       | absent        | 0 / 0                        | NC                   |
| [menu-ui](../entries/menu-ui.yml)                                                                                     | menu                                     | 菜单；可供选择的项目列表。                                                 | absent        | 0 / 0                        | NC                   |
| [merge-git](../entries/merge-git.yml)                                                                                 | merge                                    | 合并；把不同开发历史的改动整合到一起。                                     | absent        | 0 / 0                        | NC-N                 |
| [merit-general](../entries/merit-general.yml)                                                                         | merit                                    | 值得                                                                       | absent        | 0 / 0                        | NC                   |
| [method-tech](../entries/method-tech.yml)                                                                             | method                                   | 方法；与对象、类或类型关联、用于执行行为的函数成员。                       | absent        | 0 / 0                        | NC                   |
| [microtask-javascript](../entries/microtask-javascript.yml)                                                           | microtask                                | 微任务；在规定的检查点处理、通常先于下一个任务的工作。                     | W             | 0 / 0                        | K                    |
| [middleware-web](../entries/middleware-web.yml)                                                                       | middleware                               | 中间件；在请求处理链中组合可复用处理逻辑的函数。                           | absent        | 0 / 0                        | NC                   |
| [mighty-bananas-zelda](../entries/mighty-bananas-zelda.yml)                                                           | Mighty Bananas                           | 大剑香蕉；《王国之泪》中的材料名称。                                       | absent        | 0 / 0                        | NC                   |
| [mighty-general](../entries/mighty-general.yml)                                                                       | mighty                                   | 强大的                                                                     | absent        | 0 / 0                        | NC                   |
| [milk-food](../entries/milk-food.yml)                                                                                 | milk                                     | 奶；乳汁。                                                                 | absent        | 0 / 0                        | NC                   |
| [mine-architecture](../entries/mine-architecture.yml)                                                                 | mine                                     | 矿场                                                                       | absent        | 0 / 0                        | NC                   |
| [miner-general](../entries/miner-general.yml)                                                                         | miner                                    | 矿工；采矿的人。                                                           | W             | 0 / 0                        | K                    |
| [miners-mask-zelda](../entries/miners-mask-zelda.yml)                                                                 | Miner's Mask                             | 采矿面罩；《王国之泪》中的防具名称。                                       | absent        | 0 / 0                        | NC                   |
| [mineru-zelda](../entries/mineru-zelda.yml)                                                                           | Mineru                                   | 米涅鲁；古代左纳乌族人，魂之贤者。                                         | absent        | 0 / 0                        | NC                   |
| [mirror-general](../entries/mirror-general.yml)                                                                       | mirror                                   | 镜子。                                                                     | absent        | 0 / 0                        | NC                   |
| [misalign-tech](../entries/misalign-tech.yml)                                                                         | misalign                                 | 使错位；未正确对齐                                                         | W             | 0 / 0                        | K                    |
| [mnemonic-memory](../entries/mnemonic-memory.yml)                                                                     | mnemonic                                 | 助记法；帮助记忆的提示或方法。                                             | E,N           | 0 / 0                        | K                    |
| [moblin-zelda](../entries/moblin-zelda.yml)                                                                           | Moblin                                   | 莫力布林；《王国之泪》中的怪物或首领名称。                                 | absent        | 0 / 0                        | NC                   |
| [mock-tech](../entries/mock-tech.yml)                                                                                 | mock                                     | 模拟；替代依赖                                                             | absent        | 0 / 0                        | NC                   |
| [module-specifier-tech](../entries/module-specifier-tech.yml)                                                         | module specifier                         | 模块说明符；导入或导出语句中用来指明模块的字符串。                         | absent        | 0 / 0                        | NC                   |
| [mount-tech](../entries/mount-tech.yml)                                                                               | mount                                    | 挂载                                                                       | N             | 1 / 1                        | K                    |
| [mountain-terrain](../entries/mountain-terrain.yml)                                                                   | mountain                                 | 山；山岳。                                                                 | absent        | 0 / 0                        | NC                   |
| [mucktorok-zelda](../entries/mucktorok-zelda.yml)                                                                     | Mucktorok                                | 八爪怪章；《王国之泪》中的怪物或首领名称。                                 | absent        | 0 / 0                        | NC                   |
| [multibyte-tech](../entries/multibyte-tech.yml)                                                                       | multibyte                                | 多字节的                                                                   | W             | 0 / 0                        | K                    |
| [multiline-range-tech](../entries/multiline-range-tech.yml)                                                           | multiline range                          | 跨行范围                                                                   | absent        | 0 / 0                        | NC                   |
| [muse-ai](../entries/muse-ai.yml)                                                                                     | Muse                                     | 此处指 Meta 的个人 AI 智能体，可围绕用户目标执行多步骤任务。               | absent        | 0 / 0                        | D: D06               |
| [mushroom-food](../entries/mushroom-food.yml)                                                                         | mushroom                                 | 蘑菇                                                                       | absent        | 0 / 0                        | NC                   |
| [mutable-tech](../entries/mutable-tech.yml)                                                                           | mutable                                  | 可修改的                                                                   | absent        | 0 / 0                        | A: ETY04             |
| [mysterious-general](../entries/mysterious-general.yml)                                                               | mysterious                               | 神秘的                                                                     | W             | 0 / 0                        | K                    |
| [mystery-story](../entries/mystery-story.yml)                                                                         | mystery                                  | 谜团、神秘之事                                                             | absent        | 0 / 0                        | NC                   |
| [namespace-typescript](../entries/namespace-typescript.yml)                                                           | namespace                                | 命名空间；组织一组命名成员、通过限定名称访问的作用域结构。                 | absent        | 0 / 0                        | NC                   |
| [native-addon-tech](../entries/native-addon-tech.yml)                                                                 | native addon                             | 原生扩展                                                                   | absent        | 0 / 0                        | NC                   |
| [negative-lookbehind-tech](../entries/negative-lookbehind-tech.yml)                                                   | negative lookbehind                      | 否定后行断言                                                               | absent        | 0 / 0                        | NC                   |
| [nested-tech](../entries/nested-tech.yml)                                                                             | nested                                   | 嵌套的                                                                     | absent        | 0 / 0                        | NC-N                 |
| [noble-general](../entries/noble-general.yml)                                                                         | noble                                    | 高贵的                                                                     | absent        | 0 / 0                        | NC                   |
| [nonthrowing-tech](../entries/nonthrowing-tech.yml)                                                                   | nonthrowing                              | 不会抛出异常的                                                             | W             | 0 / 0                        | K                    |
| [noodles-food](../entries/noodles-food.yml)                                                                           | noodles                                  | 面条                                                                       | absent        | 0 / 0                        | NC                   |
| [note-general](../entries/note-general.yml)                                                                           | note                                     | 笔记；注释；补充说明                                                       | absent        | 0 / 0                        | NC                   |
| [objective-ui](../entries/objective-ui.yml)                                                                           | objective                                | 当前目标                                                                   | absent        | 0 / 0                        | NC                   |
| [observed-result-workflow](../entries/observed-result-workflow.yml)                                                   | observed result                          | 实际观察结果                                                               | absent        | 0 / 0                        | NC                   |
| [occupied-workflow](../entries/occupied-workflow.yml)                                                                 | occupied                                 | 被占用的                                                                   | absent        | 0 / 0                        | NC                   |
| [octorok-zelda](../entries/octorok-zelda.yml)                                                                         | Octorok                                  | 八爪怪；《王国之泪》中的怪物或首领名称。                                   | absent        | 0 / 0                        | NC                   |
| [oil-food](../entries/oil-food.yml)                                                                                   | oil                                      | 油。                                                                       | absent        | 0 / 0                        | NC                   |
| [omelet-cooking](../entries/omelet-cooking.yml)                                                                       | omelet                                   | 煎蛋卷                                                                     | absent        | 0 / 0                        | NC                   |
| [omit-tech](../entries/omit-tech.yml)                                                                                 | omit                                     | 省略                                                                       | absent        | 0 / 0                        | NC                   |
| [on-every-exit-path-workflow](../entries/on-every-exit-path-workflow.yml)                                             | on every exit path                       | 在每条退出路径上                                                           | absent        | 0 / 0                        | NC                   |
| [one-handed-combat](../entries/one-handed-combat.yml)                                                                 | one-handed                               | 单手的                                                                     | W             | 0 / 0                        | K                    |
| [one-way-data-flow-frontend](../entries/one-way-data-flow-frontend.yml)                                               | one-way data flow                        | 单向数据流；让状态沿明确方向传递并由指定入口更新。                         | absent        | 1 / 1                        | NC-N                 |
| [onion-model-koa](../entries/onion-model-koa.yml)                                                                     | onion model                              | 洋葱模型；控制流先进入下游，再返回上游继续处理的中间件模型。               | absent        | 0 / 0                        | NC-N                 |
| [open-general](../entries/open-general.yml)                                                                           | open                                     | 开着的；开放的。                                                           | absent        | 0 / 0                        | NC-N                 |
| [openclaw-ai](../entries/openclaw-ai.yml)                                                                             | OpenClaw                                 | 可自行部署的开源 AI 助手项目。                                             | W             | 0 / 0                        | K                    |
| [operand-tech](../entries/operand-tech.yml)                                                                           | operand                                  | 操作数；运算符所作用的表达式或值。                                         | absent        | 0 / 0                        | NC                   |
| [out-of-order-tech](../entries/out-of-order-tech.yml)                                                                 | out-of-order                             | 乱序的；未按预期顺序的                                                     | absent        | 0 / 0                        | NC                   |
| [outstanding-tech](../entries/outstanding-tech.yml)                                                                   | outstanding                              | 尚未完成的                                                                 | absent        | 0 / 0                        | A: N01; R: R05       |
| [overlap-prevention-workflow](../entries/overlap-prevention-workflow.yml)                                             | overlap prevention                       | 重叠执行防护                                                               | absent        | 0 / 0                        | NC                   |
| [overlap-tech](../entries/overlap-tech.yml)                                                                           | overlap                                  | 重叠                                                                       | absent        | 0 / 0                        | NC                   |
| [overload-tech](../entries/overload-tech.yml)                                                                         | overload                                 | 重载                                                                       | absent        | 0 / 0                        | A: N02               |
| [overwrite-tech](../entries/overwrite-tech.yml)                                                                       | overwrite                                | 覆盖                                                                       | absent        | 0 / 0                        | A: WF09              |
| [own-property-javascript](../entries/own-property-javascript.yml)                                                     | own property                             | 自有属性；直接属于对象自身、而非从原型链继承的属性。                       | N             | 0 / 0                        | R: R07; K            |
| [ownership-token-workflow](../entries/ownership-token-workflow.yml)                                                   | ownership token                          | 所有权标记                                                                 | absent        | 1 / 1                        | NC                   |
| [pad-general](../entries/pad-general.yml)                                                                             | pad                                      | 垫；写字用的便笺簿；此处也用于平板装置名称。                               | absent        | 0 / 0                        | NC                   |
| [paella-cooking](../entries/paella-cooking.yml)                                                                       | paella                                   | 西班牙海鲜饭或烩饭；以米饭配合肉类、海鲜或蔬菜烹制的料理。                 | E             | 0 / 0                        | K                    |
| [paint-browser-rendering](../entries/paint-browser-rendering.yml)                                                     | paint                                    | 绘制；把页面元素的视觉内容记录或绘制出来的过程。                           | absent        | 0 / 1                        | NC                   |
| [paraglider-zelda](../entries/paraglider-zelda.yml)                                                                   | Paraglider                               | 滑翔帆；本作用于空中滑翔的装备。                                           | absent        | 0 / 0                        | NC                   |
| [parameter-tech](../entries/parameter-tech.yml)                                                                       | parameter                                | 参数；函数或泛型定义中用于接收实参的占位名称。                             | N             | 1 / 1                        | K                    |
| [part-of-speech-language](../entries/part-of-speech-language.yml)                                                     | part of speech                           | 词性；按语法功能划分的词语类别。                                           | absent        | 0 / 0                        | NC                   |
| [partial-clone-git](../entries/partial-clone-git.yml)                                                                 | partial clone                            | 部分克隆；最初只获取部分 Git 对象，并按需要补取其余对象的克隆方式。        | absent        | 0 / 0                        | NC                   |
| [pass-with-the-fix-workflow](../entries/pass-with-the-fix-workflow.yml)                                               | pass with the fix                        | 应用修复后通过                                                             | absent        | 0 / 0                        | NC                   |
| [passage-architecture](../entries/passage-architecture.yml)                                                           | passage                                  | 通道                                                                       | absent        | 0 / 0                        | NC                   |
| [pasta-food](../entries/pasta-food.yml)                                                                               | pasta                                    | 意大利面；意面                                                             | absent        | 0 / 0                        | NC-N                 |
| [patch-tech](../entries/patch-tech.yml)                                                                               | patch                                    | 补丁；用于修正或调整程序的一组改动。                                       | absent        | 0 / 0                        | NC-N                 |
| [paya-zelda](../entries/paya-zelda.yml)                                                                               | Paya                                     | 帕雅；卡卡利科村的领导者。                                                 | absent        | 0 / 0                        | NC                   |
| [pending-tech](../entries/pending-tech.yml)                                                                           | pending                                  | 待处理的；待完成的                                                         | absent        | 0 / 0                        | R: R05               |
| [pepper-food](../entries/pepper-food.yml)                                                                             | pepper                                   | 胡椒；胡椒粉                                                               | absent        | 0 / 0                        | NC-N                 |
| [pepper-vegetable](../entries/pepper-vegetable.yml)                                                                   | pepper                                   | 辣椒；甜椒；这里不指胡椒粉。                                               | absent        | 0 / 0                        | NC-N                 |
| [persist-tech](../entries/persist-tech.yml)                                                                           | persist                                  | 持久保存；使数据在当前操作或进程结束后仍可恢复使用。                       | absent        | 0 / 0                        | NC-N                 |
| [phantom-ganon-zelda](../entries/phantom-ganon-zelda.yml)                                                             | Phantom Ganon                            | 幻影加侬；本作中加侬多夫的幻影敌人。                                       | absent        | 0 / 0                        | NC                   |
| [phenomenon-story](../entries/phenomenon-story.yml)                                                                   | phenomenon                               | 现象                                                                       | absent        | 0 / 0                        | NC                   |
| [pick-up-general](../entries/pick-up-general.yml)                                                                     | pick up                                  | 捡起                                                                       | absent        | 0 / 0                        | NC                   |
| [pilot-general](../entries/pilot-general.yml)                                                                         | pilot                                    | 飞行员；驾驶飞机的人。                                                     | absent        | 0 / 0                        | NC-N                 |
| [pinned-checkout-workflow](../entries/pinned-checkout-workflow.yml)                                                   | pinned checkout                          | 固定版本的检出目录                                                         | absent        | 0 / 0                        | NC                   |
| [plateau-terrain](../entries/plateau-terrain.yml)                                                                     | plateau                                  | 高原                                                                       | absent        | 0 / 0                        | NC                   |
| [platform-architecture](../entries/platform-architecture.yml)                                                         | platform                                 | 平台                                                                       | absent        | 0 / 0                        | NC                   |
| [pork-food](../entries/pork-food.yml)                                                                                 | pork                                     | 猪肉                                                                       | absent        | 0 / 0                        | NC                   |
| [porridge-cooking](../entries/porridge-cooking.yml)                                                                   | porridge                                 | 粥                                                                         | absent        | 0 / 0                        | NC                   |
| [portable-pot-zelda](../entries/portable-pot-zelda.yml)                                                               | Portable Pot                             | 便携锅；《王国之泪》中的左纳乌装置。                                       | absent        | 0 / 0                        | NC                   |
| [possessed-combat](../entries/possessed-combat.yml)                                                                   | possessed                                | 被附身的                                                                   | absent        | 0 / 0                        | NC                   |
| [postorder-traversal-tree](../entries/postorder-traversal-tree.yml)                                                   | postorder traversal                      | 后序遍历；先处理子树，再访问根节点的树遍历顺序。                           | absent        | 0 / 0                        | NC                   |
| [pot-general](../entries/pot-general.yml)                                                                             | pot                                      | 锅；罐。                                                                   | absent        | 0 / 0                        | NC                   |
| [potato-food](../entries/potato-food.yml)                                                                             | potato                                   | 土豆；马铃薯                                                               | absent        | 0 / 0                        | NC                   |
| [pouch-ui](../entries/pouch-ui.yml)                                                                                   | pouch                                    | 小袋；随身袋。                                                             | absent        | 0 / 1                        | NC                   |
| [powerful-general](../entries/powerful-general.yml)                                                                   | powerful                                 | 强大的                                                                     | W             | 0 / 0                        | K                    |
| [pre-existing-tech](../entries/pre-existing-tech.yml)                                                                 | pre-existing                             | 原先就存在的                                                               | W             | 0 / 0                        | K                    |
| [precede-tech](../entries/precede-tech.yml)                                                                           | precede                                  | 位于之前                                                                   | absent        | 0 / 0                        | NC                   |
| [prefilter-tech](../entries/prefilter-tech.yml)                                                                       | prefilter                                | 预筛选器                                                                   | W             | 0 / 0                        | K                    |
| [preflight-request-web](../entries/preflight-request-web.yml)                                                         | preflight request                        | 预检请求；浏览器在部分跨源请求前发送的权限检查请求。                       | N             | 0 / 0                        | K                    |
| [preserve-tech](../entries/preserve-tech.yml)                                                                         | preserve                                 | 保留；保持不变                                                             | N             | 1 / 1                        | K                    |
| [prevent-tech](../entries/prevent-tech.yml)                                                                           | prevent                                  | 防止                                                                       | absent        | 0 / 0                        | NC                   |
| [primary-checkout-workflow](../entries/primary-checkout-workflow.yml)                                                 | primary checkout                         | 主检出目录                                                                 | absent        | 0 / 0                        | NC                   |
| [prime-food](../entries/prime-food.yml)                                                                               | prime                                    | 上等的；优质的。                                                           | absent        | 0 / 0                        | NC                   |
| [primitive-javascript](../entries/primitive-javascript.yml)                                                           | primitive                                | 原始值；不属于对象、值本身不可变的 JavaScript 数据。                       | absent        | 0 / 0                        | NC-N                 |
| [prince-general](../entries/prince-general.yml)                                                                       | prince                                   | 王子                                                                       | absent        | 0 / 0                        | NC                   |
| [princess-general](../entries/princess-general.yml)                                                                   | princess                                 | 公主                                                                       | absent        | 0 / 0                        | NC                   |
| [princess-zelda-zelda](../entries/princess-zelda-zelda.yml)                                                           | Princess Zelda                           | 塞尔达公主；海拉鲁的公主。                                                 | absent        | 0 / 0                        | NC                   |
| [pristine-general](../entries/pristine-general.yml)                                                                   | pristine                                 | 完好如初的；未受损或未被改变的。                                           | N             | 0 / 1                        | K                    |
| [probe-tech](../entries/probe-tech.yml)                                                                               | probe                                    | 针对性验证程序；探查                                                       | absent        | 0 / 0                        | NC                   |
| [profile-ui](../entries/profile-ui.yml)                                                                               | profile                                  | 简介；概况。                                                               | absent        | 0 / 0                        | NC                   |
| [profiling-tech](../entries/profiling-tech.yml)                                                                       | profiling                                | 性能分析                                                                   | absent        | 0 / 0                        | NC                   |
| [progressive-enhancement-web](../entries/progressive-enhancement-web.yml)                                             | progressive enhancement                  | 渐进增强；先提供基础可用能力，再为更强环境增加功能。                       | absent        | 0 / 1                        | NC                   |
| [project-onboarding-workflow](../entries/project-onboarding-workflow.yml)                                             | project onboarding                       | 项目熟悉过程；建立对业务、代码与交付流程的工作认识。                       | absent        | 0 / 0                        | NC                   |
| [promise-all-javascript](../entries/promise-all-javascript.yml)                                                       | Promise.all                              | Promise 聚合；全部成功时按输入顺序返回结果，任一拒绝则拒绝。               | N             | 0 / 0                        | K                    |
| [promise-javascript](../entries/promise-javascript.yml)                                                               | Promise                                  | Promise；表示异步操作最终成功或失败结果的对象。                            | absent        | 0 / 1                        | NC                   |
| [pronunciation-general](../entries/pronunciation.yml)                                                                 | pronunciation                            | 发音；读音                                                                 | N             | 0 / 1                        | K                    |
| [propagate-tech](../entries/propagate-tech.yml)                                                                       | propagate                                | 传播；传递                                                                 | absent        | 0 / 0                        | NC-N                 |
| [property-descriptor-javascript](../entries/property-descriptor-javascript.yml)                                       | property descriptor                      | 属性描述符；描述属性的值、访问器及相关特性的信息对象。                     | absent        | 0 / 0                        | NC-N                 |
| [property-tech](../entries/property-tech.yml)                                                                         | property                                 | 属性；以键标识的对象成员。                                                 | absent        | 1 / 1                        | NC                   |
| [prophecy-story](../entries/prophecy-story.yml)                                                                       | prophecy                                 | 预言                                                                       | absent        | 0 / 0                        | NC                   |
| [prototype-chain-javascript](../entries/prototype-chain-javascript.yml)                                               | prototype chain                          | 原型链；沿对象的原型关系逐层查找属性。                                     | absent        | 0 / 0                        | NC-N                 |
| [prototype-setter-javascript](../entries/prototype-setter-javascript.yml)                                             | prototype setter                         | 原型设置语法；此处指对象字面量中用 __proto__: value 设置原型的特殊写法。   | absent        | 0 / 0                        | NC-N                 |
| [prune-tech](../entries/prune-tech.yml)                                                                               | prune                                    | 剪除；删减                                                                 | absent        | 0 / 0                        | NC                   |
| [puffshroom-zelda](../entries/puffshroom-zelda.yml)                                                                   | Puffshroom                               | 烟雾蘑菇；《王国之泪》中的材料名称。                                       | absent        | 0 / 0                        | NC                   |
| [purah-pad-zelda](../entries/purah-pad-zelda.yml)                                                                     | Purah Pad                                | 普尔亚平板；林克使用的多功能装置。                                         | absent        | 0 / 0                        | NC                   |
| [purah-zelda](../entries/purah-zelda.yml)                                                                             | Purah                                    | 普尔亚；监视堡垒的重要研究人员。                                           | absent        | 0 / 0                        | NC                   |
| [quadratic-tech](../entries/quadratic-tech.yml)                                                                       | quadratic                                | 二次增长的；输入规模增加时，工作量按平方量级增长。                         | absent        | 1 / 1                        | NC                   |
| [queen-general](../entries/queen-general.yml)                                                                         | queen                                    | 王后、女王                                                                 | absent        | 0 / 0                        | NC                   |
| [queen-gibdo-zelda](../entries/queen-gibdo-zelda.yml)                                                                 | Queen Gibdo                              | 吉波得女王；《王国之泪》中的怪物或首领名称。                               | absent        | 0 / 0                        | NC                   |
| [query-tech](../entries/query-tech.yml)                                                                               | query                                    | 查询；向系统请求信息或检索结果。                                           | absent        | 0 / 0                        | NC-N                 |
| [quest-game](../entries/quest-game.yml)                                                                               | quest                                    | 任务；追寻或探索。                                                         | absent        | 0 / 0                        | NC                   |
| [quicksort-algorithm](../entries/quicksort-algorithm.yml)                                                             | quicksort                                | 快速排序；按基准值分区，再递归排序各分区。                                 | W             | 0 / 0                        | K                    |
| [radish-food](../entries/radish-food.yml)                                                                             | radish                                   | 萝卜。                                                                     | absent        | 0 / 0                        | NC                   |
| [ranged-combat](../entries/ranged-combat.yml)                                                                         | ranged                                   | 远程                                                                       | absent        | 0 / 0                        | NC                   |
| [rather-than-workflow](../entries/rather-than-workflow.yml)                                                           | rather than                              | 而不是                                                                     | absent        | 0 / 0                        | NC                   |
| [rauru-zelda](../entries/rauru-zelda.yml)                                                                             | Rauru                                    | 劳鲁；古代海拉鲁的国王，左纳乌族人。                                       | absent        | 0 / 0                        | NC                   |
| [raw-food](../entries/raw-food.yml)                                                                                   | raw                                      | 生的；未烹饪的。                                                           | absent        | 0 / 0                        | NC                   |
| [raw-gourmet-meat-zelda](../entries/raw-gourmet-meat-zelda.yml)                                                       | Raw Gourmet Meat                         | 顶级兽肉；《王国之泪》中的材料名称。                                       | absent        | 0 / 0                        | NC                   |
| [raw-prime-meat-zelda](../entries/raw-prime-meat-zelda.yml)                                                           | Raw Prime Meat                           | 高级兽肉；《王国之泪》中的材料名称。                                       | absent        | 0 / 0                        | NC                   |
| [re-entry-tech](../entries/re-entry-tech.yml)                                                                         | re-entry                                 | 再次进入                                                                   | W             | 0 / 0                        | K                    |
| [read-only-workflow](../entries/read-only-workflow.yml)                                                               | read-only                                | 只读的                                                                     | absent        | 0 / 0                        | A: WF10              |
| [ready-for-review-workflow](../entries/ready-for-review-workflow.yml)                                                 | ready for review                         | 已准备好供审查                                                             | absent        | 0 / 0                        | NC                   |
| [recall-general](../entries/recall-general.yml)                                                                       | recall                                   | 回想；召回；叫回。                                                         | absent        | 0 / 0                        | NC-N                 |
| [recharge-general](../entries/recharge-general.yml)                                                                   | recharge                                 | 充能                                                                       | W             | 0 / 0                        | K                    |
| [recipe-cooking](../entries/recipe-cooking.yml)                                                                       | recipe                                   | 食谱                                                                       | E             | 0 / 1                        | K                    |
| [reckless-general](../entries/reckless-general.yml)                                                                   | reckless                                 | 鲁莽的                                                                     | N             | 0 / 1                        | K                    |
| [reconciliation-react](../entries/reconciliation-react.yml)                                                           | reconciliation                           | React 协调；比较界面描述并确定更新与状态保留。                             | absent        | 0 / 0                        | NC-N                 |
| [recover-story](../entries/recover-story.yml)                                                                         | recover                                  | 找回、恢复                                                                 | absent        | 1 / 1                        | NC                   |
| [recovery-tech](../entries/recovery-tech.yml)                                                                         | recovery                                 | 错误后的恢复                                                               | absent        | 0 / 0                        | R: R04               |
| [recursion-programming](../entries/recursion-programming.yml)                                                         | recursion                                | 递归；在求解或定义过程中再次处理同类子问题的方式。                         | absent        | 0 / 0                        | R: R10               |
| [refinery-architecture](../entries/refinery-architecture.yml)                                                         | refinery                                 | 精炼设施                                                                   | W             | 0 / 0                        | K                    |
| [regional-general](../entries/regional-general.yml)                                                                   | regional                                 | 地区性的                                                                   | W             | 0 / 0                        | K                    |
| [register-tech](../entries/register-tech.yml)                                                                         | register                                 | 注册                                                                       | N             | 1 / 1                        | K                    |
| [registry-tech](../entries/registry-tech.yml)                                                                         | registry                                 | 注册表；注册集合                                                           | absent        | 0 / 0                        | NC                   |
| [regression-tech](../entries/regression-tech.yml)                                                                     | regression                               | 回归问题；行为退化                                                         | absent        | 0 / 0                        | A: ETY07             |
| [regression-test-tech](../entries/regression-test-tech.yml)                                                           | regression test                          | 回归测试                                                                   | absent        | 0 / 0                        | NC                   |
| [reject-tech](../entries/reject-tech.yml)                                                                             | reject                                   | 拒绝接受                                                                   | absent        | 0 / 0                        | NC                   |
| [related-words-language](../entries/related-words-language.yml)                                                       | related words                            | 相关词语；在意义、形式或用法上有联系的词。                                 | absent        | 0 / 0                        | NC                   |
| [release-workflow](../entries/release-workflow.yml)                                                                   | release                                  | 释放                                                                       | absent        | 0 / 0                        | NC                   |
| [remain-intact-workflow](../entries/remain-intact-workflow.yml)                                                       | remain intact                            | 保持完整；未受破坏                                                         | absent        | 0 / 0                        | NC                   |
| [remain-unchanged-workflow](../entries/remain-unchanged-workflow.yml)                                                 | remain unchanged                         | 保持不变                                                                   | absent        | 0 / 0                        | NC                   |
| [remote-tracking-branch-git](../entries/remote-tracking-branch-git.yml)                                               | remote-tracking branch                   | 远端跟踪分支                                                               | absent        | 0 / 0                        | NC                   |
| [reported-behavior-workflow](../entries/reported-behavior-workflow.yml)                                               | reported behavior                        | 报告中的行为                                                               | absent        | 0 / 0                        | NC                   |
| [repository-git](../entries/repository-git.yml)                                                                       | repository                               | 仓库；保存项目文件及其版本历史的集合。                                     | absent        | 0 / 0                        | NC-N                 |
| [reproduce-tech](../entries/reproduce-tech.yml)                                                                       | reproduce                                | 复现                                                                       | N             | 1 / 1                        | K                    |
| [reproducible-tech](../entries/reproducible-tech.yml)                                                                 | reproducible                             | 可复现的                                                                   | N             | 0 / 1                        | K                    |
| [reserve-workflow](../entries/reserve-workflow.yml)                                                                   | reserve                                  | 预留量                                                                     | absent        | 0 / 0                        | NC                   |
| [reserved-prefix-tech](../entries/reserved-prefix-tech.yml)                                                           | reserved prefix                          | 保留前缀                                                                   | absent        | 0 / 0                        | NC                   |
| [resilient-combat](../entries/resilient-combat.yml)                                                                   | resilient                                | 有韧性的；遭受打击后能恢复或继续应对的。                                   | E,N           | 0 / 1                        | K                    |
| [resistance-general](../entries/resistance-general.yml)                                                               | resistance                               | 抵抗；抵抗力。                                                             | absent        | 0 / 0                        | NC                   |
| [resolve-tech](../entries/resolve-tech.yml)                                                                           | resolve                                  | 解析；确定所指对象                                                         | N             | 1 / 1                        | K                    |
| [resolve-workflow](../entries/resolve-workflow.yml)                                                                   | resolve                                  | 解决；处理完毕                                                             | N             | 0 / 1                        | K                    |
| [restore-tech](../entries/restore-tech.yml)                                                                           | restore                                  | 恢复                                                                       | absent        | 0 / 1                        | NC                   |
| [restrict-tech](../entries/restrict-tech.yml)                                                                         | restrict                                 | 限制                                                                       | absent        | 0 / 0                        | NC                   |
| [resumption-workflow](../entries/resumption-workflow.yml)                                                             | resumption                               | 恢复继续                                                                   | absent        | 0 / 0                        | A: WF08; R: R04      |
| [retain-tech](../entries/retain-tech.yml)                                                                             | retain                                   | 保留；继续维持                                                             | N             | 0 / 1                        | K                    |
| [return-type-tech](../entries/return-type-tech.yml)                                                                   | return type                              | 返回类型；函数结果在类型系统中的类型。                                     | absent        | 0 / 0                        | NC                   |
| [review-workflow](../entries/review-workflow.yml)                                                                     | review                                   | 审查；检查工作内容并给出判断或反馈。                                       | absent        | 0 / 0                        | NC                   |
| [revision-general](../entries/revision-general.yml)                                                                   | revision                                 | 修订                                                                       | N             | 0 / 1                        | K                    |
| [revision-git](../entries/revision-git.yml)                                                                           | revision                                 | 代码版本；提交版本                                                         | N             | 1 / 1                        | K                    |
| [rice-food](../entries/rice-food.yml)                                                                                 | rice                                     | 米；米饭。                                                                 | absent        | 0 / 0                        | NC                   |
| [ridge-terrain](../entries/ridge-terrain.yml)                                                                         | ridge                                    | 山脊                                                                       | absent        | 0 / 1                        | NC                   |
| [riju-zelda](../entries/riju-zelda.yml)                                                                               | Riju                                     | 露珠；格鲁德族的领导者，雷之贤者。                                         | absent        | 0 / 0                        | NC                   |
| [ring-general](../entries/ring-general.yml)                                                                           | ring                                     | 环；圈；戒指。                                                             | absent        | 0 / 0                        | NC                   |
| [risotto-cooking](../entries/risotto-cooking.yml)                                                                     | risotto                                  | 意大利烩饭                                                                 | absent        | 0 / 0                        | NC                   |
| [rito-village-zelda](../entries/rito-village-zelda.yml)                                                               | Rito Village                             | 利特村；《王国之泪》中的地点或设施。                                       | absent        | 0 / 0                        | NC                   |
| [rito-zelda](../entries/rito-zelda.yml)                                                                               | Rito                                     | 利特族；《王国之泪》中的族群或组织。                                       | absent        | 0 / 0                        | NC                   |
| [river-terrain](../entries/river-terrain.yml)                                                                         | river                                    | 河流                                                                       | absent        | 0 / 0                        | NC                   |
| [roast-food](../entries/roast-food.yml)                                                                               | roast                                    | 烤制的；烘烤的                                                             | absent        | 0 / 0                        | NC                   |
| [robbie-zelda](../entries/robbie-zelda.yml)                                                                           | Robbie                                   | 洛贝利；研究古代技术的研究人员。                                           | absent        | 0 / 0                        | NC                   |
| [robe-clothing](../entries/robe-clothing.yml)                                                                         | robe                                     | 长袍                                                                       | absent        | 0 / 0                        | NC                   |
| [rocket-engineering](../entries/rocket-engineering.yml)                                                               | rocket                                   | 火箭。                                                                     | absent        | 0 / 0                        | NC                   |
| [rod-general](../entries/rod-general.yml)                                                                             | rod                                      | 杆；棒。                                                                   | absent        | 0 / 0                        | NC                   |
| [root-general](../entries/root-general.yml)                                                                           | root                                     | 根；植物的根。                                                             | absent        | 0 / 0                        | NC                   |
| [rotate-general](../entries/rotate-general.yml)                                                                       | rotate                                   | 旋转                                                                       | absent        | 0 / 0                        | NC                   |
| [round-trip-tech](../entries/round-trip-tech.yml)                                                                     | round trip                               | 往返转换                                                                   | absent        | 0 / 0                        | NC                   |
| [royal-broadsword-zelda](../entries/royal-broadsword-zelda.yml)                                                       | Royal Broadsword                         | 王族之剑；《王国之泪》中的装备名称。                                       | absent        | 0 / 0                        | NC                   |
| [royal-general](../entries/royal-general.yml)                                                                         | royal                                    | 王室的                                                                     | absent        | 0 / 0                        | NC                   |
| [ruins-architecture](../entries/ruins-architecture.yml)                                                               | ruins                                    | 废墟、遗迹                                                                 | absent        | 0 / 0                        | NC                   |
| [run-budget-workflow](../entries/run-budget-workflow.yml)                                                             | run budget                               | 任务预算                                                                   | absent        | 0 / 0                        | NC                   |
| [run-in-parallel-workflow](../entries/run-in-parallel-workflow.yml)                                                   | run in parallel                          | 并行运行                                                                   | absent        | 0 / 0                        | NC                   |
| [runtime-evaluation-tech](../entries/runtime-evaluation-tech.yml)                                                     | runtime evaluation                       | 运行时求值；运行验证                                                       | absent        | 0 / 0                        | NC                   |
| [runtime-tech](../entries/runtime-tech.yml)                                                                           | runtime                                  | 运行时；运行支持组件                                                       | absent        | 0 / 0                        | NC                   |
| [rupee-game](../entries/rupee-game.yml)                                                                               | rupee                                    | 卢比；本作的货币单位。                                                     | E             | 0 / 0                        | K                    |
| [rusty-general](../entries/rusty-general.yml)                                                                         | rusty                                    | 生锈的                                                                     | absent        | 0 / 0                        | NC                   |
| [ruthless-general](../entries/ruthless-general.yml)                                                                   | ruthless                                 | 无情的                                                                     | W             | 0 / 0                        | K                    |
| [sacred-general](../entries/sacred-general.yml)                                                                       | sacred                                   | 神圣的                                                                     | absent        | 0 / 0                        | NC                   |
| [sacrifice-story](../entries/sacrifice-story.yml)                                                                     | sacrifice                                | 牺牲                                                                       | E             | 0 / 0                        | K                    |
| [sage-general](../entries/sage-general.yml)                                                                           | sage                                     | 贤者；富有智慧的人。                                                       | absent        | 0 / 0                        | NC                   |
| [salmon-food](../entries/salmon-food.yml)                                                                             | salmon                                   | 三文鱼；鲑鱼。                                                             | absent        | 0 / 0                        | NC                   |
| [salt-food](../entries/salt-food.yml)                                                                                 | salt                                     | 盐。                                                                       | absent        | 0 / 0                        | NC                   |
| [same-origin-policy-web](../entries/same-origin-policy-web.yml)                                                       | same-origin policy                       | 同源策略；限制脚本访问不同源资源的浏览器安全机制。                         | absent        | 0 / 1                        | NC-N                 |
| [sanitize-tech](../entries/sanitize-tech.yml)                                                                         | sanitize                                 | 清理或转换为可接受的输入                                                   | absent        | 0 / 0                        | NC                   |
| [sauce-food](../entries/sauce-food.yml)                                                                               | sauce                                    | 酱汁                                                                       | absent        | 0 / 0                        | NC                   |
| [saute-cooking](../entries/saute-cooking.yml)                                                                         | sauté                                    | 用少量油快速翻炒或煎炒。                                                   | E             | 0 / 0                        | K                    |
| [savage-general](../entries/savage-general.yml)                                                                       | savage                                   | 凶猛的                                                                     | absent        | 0 / 0                        | NC                   |
| [scimitar-of-the-seven-zelda](../entries/scimitar-of-the-seven-zelda.yml)                                             | Scimitar of the Seven                    | 七宝匕首；《王国之泪》中的装备名称。                                       | absent        | 0 / 0                        | NC                   |
| [scimitar-weapon](../entries/scimitar-weapon.yml)                                                                     | scimitar                                 | 弯刀                                                                       | absent        | 0 / 0                        | NC                   |
| [scope-lexical](../entries/scope-lexical.yml)                                                                         | scope                                    | 作用域                                                                     | N             | 1 / 1                        | K                    |
| [scope-workflow](../entries/scope-workflow.yml)                                                                       | scope                                    | 工作范围                                                                   | N             | 0 / 1                        | K                    |
| [seafood-food](../entries/seafood-food.yml)                                                                           | seafood                                  | 海鲜                                                                       | absent        | 0 / 0                        | NC                   |
| [sear-cooking](../entries/sear-cooking.yml)                                                                           | sear                                     | 以高温迅速煎烤食物表面。                                                   | N             | 0 / 1                        | K                    |
| [secret-stone-zelda](../entries/secret-stone-zelda.yml)                                                               | Secret Stone                             | 秘石；本作中可放大持有者力量的重要物品。                                   | absent        | 0 / 0                        | NC                   |
| [secret-story](../entries/secret-story.yml)                                                                           | secret                                   | 秘密                                                                       | absent        | 0 / 0                        | NC                   |
| [seed-general](../entries/seed-general.yml)                                                                           | seed                                     | 种子。                                                                     | absent        | 0 / 0                        | NC                   |
| [seek-general](../entries/seek-general.yml)                                                                           | seek                                     | 寻找；寻求。                                                               | absent        | 0 / 0                        | NC-N                 |
| [seized-construct-zelda](../entries/seized-construct-zelda.yml)                                                       | Seized Construct                         | 被夺之魔像；《王国之泪》中的怪物或首领名称。                               | absent        | 0 / 0                        | NC                   |
| [select-general](../entries/select-general.yml)                                                                       | select                                   | 选择                                                                       | absent        | 0 / 0                        | NC                   |
| [selfless-general](../entries/selfless-general.yml)                                                                   | selfless                                 | 无私的                                                                     | W             | 0 / 0                        | K                    |
| [semantic-versioning-software](../entries/semantic-versioning-software.yml)                                           | semantic versioning                      | 语义化版本；用主版本、次版本和补丁表达公共 API 变化。                      | absent        | 0 / 0                        | NC                   |
| [serialization-tech](../entries/serialization-tech.yml)                                                               | serialization                            | 序列化                                                                     | absent        | 0 / 0                        | NC                   |
| [settlement-architecture](../entries/settlement-architecture.yml)                                                     | settlement                               | 定居点、聚落                                                               | W             | 0 / 0                        | K                    |
| [setup-tech](../entries/setup-tech.yml)                                                                               | setup                                    | 初始化设置                                                                 | W             | 0 / 0                        | K                    |
| [shallow-copy-javascript](../entries/shallow-copy-javascript.yml)                                                     | shallow copy                             | 浅拷贝；复制外层结构，嵌套对象仍共享引用。                                 | absent        | 0 / 1                        | NC-N                 |
| [sheikah-zelda](../entries/sheikah-zelda.yml)                                                                         | Sheikah                                  | 希卡族；《王国之泪》中的族群或组织。                                       | absent        | 0 / 0                        | NC                   |
| [shield-weapon](../entries/shield-weapon.yml)                                                                         | shield                                   | 盾牌                                                                       | absent        | 0 / 0                        | NC                   |
| [shirt-clothing](../entries/shirt-clothing.yml)                                                                       | shirt                                    | 衬衫；上衣。                                                               | absent        | 0 / 0                        | NC                   |
| [shock-emitter-zelda](../entries/shock-emitter-zelda.yml)                                                             | Shock Emitter                            | 雷龙头；《王国之泪》中的左纳乌装置。                                       | W             | 0 / 0                        | K                    |
| [shock-resistance-combat](../entries/shock-resistance-combat.yml)                                                     | shock resistance                         | 抗电；抵抗电击影响的能力。                                                 | absent        | 0 / 0                        | NC                   |
| [shorthand-javascript](../entries/shorthand-javascript.yml)                                                           | shorthand                                | 简写；用更短的语法表达相应的常规写法。                                     | absent        | 0 / 0                        | NC-N                 |
| [shrine-architecture](../entries/shrine-architecture.yml)                                                             | shrine                                   | 神龛；圣祠；供奉或祭祀的场所。                                             | absent        | 0 / 0                        | NC                   |
| [shrine-quests-zelda](../entries/shrine-quests-zelda.yml)                                                             | Shrine Quests                            | 神庙挑战；《王国之泪》中的界面类别或功能名称。                             | absent        | 0 / 0                        | NC                   |
| [side-adventures-zelda](../entries/side-adventures-zelda.yml)                                                         | Side Adventures                          | 情节挑战；《王国之泪》中的界面类别或功能名称。                             | absent        | 0 / 0                        | NC                   |
| [side-quests-zelda](../entries/side-quests-zelda.yml)                                                                 | Side Quests                              | 迷你挑战；《王国之泪》中的界面类别或功能名称。                             | absent        | 0 / 0                        | NC                   |
| [sidon-zelda](../entries/sidon-zelda.yml)                                                                             | Sidon                                    | 希多；卓拉族人，水之贤者。                                                 | absent        | 0 / 0                        | NC                   |
| [signed-zero-tech](../entries/signed-zero-tech.yml)                                                                   | signed zero                              | 带符号的零                                                                 | absent        | 0 / 0                        | NC-N                 |
| [silent-princess-zelda](../entries/silent-princess-zelda.yml)                                                         | Silent Princess                          | 宁静公主；《王国之泪》中的材料名称。                                       | absent        | 0 / 0                        | NC                   |
| [simmer-cooking](../entries/simmer-cooking.yml)                                                                       | simmer                                   | 小火慢煮；使液体保持轻微沸腾。                                             | N             | 1 / 1                        | K                    |
| [skewer-cooking](../entries/skewer-cooking.yml)                                                                       | skewer                                   | 串食物用的签；串在签上烹饪的食物。                                         | absent        | 0 / 0                        | NC                   |
| [skilled-general](../entries/skilled-general.yml)                                                                     | skilled                                  | 熟练的、有技巧的                                                           | W             | 0 / 0                        | K                    |
| [sky-island-zelda](../entries/sky-island-zelda.yml)                                                                   | sky island                               | 空岛；悬浮在空中的岛屿。                                                   | W             | 0 / 0                        | K                    |
| [sky-terrain](../entries/sky-terrain.yml)                                                                             | sky                                      | 天空                                                                       | absent        | 0 / 0                        | NC                   |
| [skyshroom-zelda](../entries/skyshroom-zelda.yml)                                                                     | Skyshroom                                | 天空蘑菇；《王国之泪》中的材料名称。                                       | absent        | 0 / 0                        | NC                   |
| [skyview-tower-zelda](../entries/skyview-tower-zelda.yml)                                                             | Skyview Tower                            | 鸟望台；用于调查区域地图的塔状设施。                                       | W             | 0 / 0                        | K                    |
| [sled-engineering](../entries/sled-engineering.yml)                                                                   | sled                                     | 雪橇；可在雪地等表面滑行的载具。                                           | absent        | 0 / 0                        | NC                   |
| [sleeve-clothing](../entries/sleeve-clothing.yml)                                                                     | sleeve                                   | 袖子。                                                                     | absent        | 0 / 0                        | NC                   |
| [slip-general](../entries/slip-general.yml)                                                                           | slip                                     | 滑动；滑倒；失去抓地力。                                                   | absent        | 0 / 0                        | NC                   |
| [slip-resistance-zelda](../entries/slip-resistance-zelda.yml)                                                         | Slip Resistance                          | 滑落减轻；减少攀爬湿滑表面时滑落的效果。                                   | absent        | 0 / 0                        | NC                   |
| [slope-terrain](../entries/slope-terrain.yml)                                                                         | slope                                    | 斜坡                                                                       | absent        | 0 / 0                        | NC                   |
| [small-wheel-zelda](../entries/small-wheel-zelda.yml)                                                                 | Small Wheel                              | 小轮胎；《王国之泪》中的左纳乌装置。                                       | absent        | 0 / 0                        | NC                   |
| [snapshot-tech](../entries/snapshot-tech.yml)                                                                         | snapshot                                 | 快照                                                                       | absent        | 0 / 0                        | NC                   |
| [sneaky-general](../entries/sneaky-general.yml)                                                                       | sneaky                                   | 偷偷摸摸的；行动隐蔽的。                                                   | absent        | 0 / 0                        | NC                   |
| [snowquill-tunic-zelda](../entries/snowquill-tunic-zelda.yml)                                                         | Snowquill Tunic                          | 利特的羽绒服；《王国之泪》中的防具名称。                                   | absent        | 0 / 0                        | NC                   |
| [soldier-construct-zelda](../entries/soldier-construct-zelda.yml)                                                     | Soldier Construct                        | 兵队魔像；《王国之泪》中的怪物或首领名称。                                 | absent        | 0 / 0                        | NC                   |
| [soldier-general](../entries/soldier-general.yml)                                                                     | soldier                                  | 士兵。                                                                     | absent        | 0 / 0                        | NC                   |
| [soldiers-broadsword-zelda](../entries/soldiers-broadsword-zelda.yml)                                                 | Soldier's Broadsword                     | 士兵之剑；《王国之泪》中的装备名称。                                       | absent        | 0 / 0                        | NC                   |
| [sonia-zelda](../entries/sonia-zelda.yml)                                                                             | Sonia                                    | 索尼娅；古代海拉鲁的王后。                                                 | absent        | 0 / 0                        | NC                   |
| [sort-general](../entries/sort-general.yml)                                                                           | sort                                     | 排序                                                                       | absent        | 0 / 0                        | NC                   |
| [soundness-type-system](../entries/soundness-type-system.yml)                                                         | soundness                                | 健全性；类型检查所接受的程序不会违反该系统承诺的类型安全性质。             | N             | 1 / 1                        | K                    |
| [source-code-tech](../entries/source-code-tech.yml)                                                                   | source code                              | 源代码；以编程语言编写的程序文本。                                         | absent        | 0 / 0                        | NC                   |
| [source-map-tech](../entries/source-map-tech.yml)                                                                     | source map                               | 源码映射                                                                   | absent        | 0 / 0                        | NC                   |
| [spaulder-clothing](../entries/spaulder-clothing.yml)                                                                 | spaulder                                 | 肩甲                                                                       | absent        | 0 / 0                        | NC                   |
| [spear-weapon](../entries/spear-weapon.yml)                                                                           | spear                                    | 长矛                                                                       | absent        | 0 / 0                        | NC                   |
| [special-controls-zelda](../entries/special-controls-zelda.yml)                                                       | Special Controls                         | 特殊动作；《王国之泪》中的界面类别或功能名称。                             | absent        | 0 / 0                        | NC                   |
| [speculative-tech](../entries/speculative-tech.yml)                                                                   | speculative                              | 推测性的                                                                   | absent        | 0 / 0                        | NC                   |
| [speed-up-zelda](../entries/speed-up-zelda.yml)                                                                       | Speed Up                                 | 移动力提升；《王国之泪》中的能力效果名称。                                 | absent        | 0 / 0                        | NC                   |
| [spice-food](../entries/spice-food.yml)                                                                               | spice                                    | 香辛料；调味香料。                                                         | absent        | 0 / 0                        | NC                   |
| [spicy-general](../entries/spicy-general.yml)                                                                         | spicy                                    | 辛辣的；带香辛料味道的。                                                   | absent        | 0 / 0                        | NC                   |
| [spicy-pepper-zelda](../entries/spicy-pepper-zelda.yml)                                                               | Spicy Pepper                             | 暖暖草果；《王国之泪》中的材料名称。                                       | absent        | 0 / 0                        | NC                   |
| [spirit-story](../entries/spirit-story.yml)                                                                           | spirit                                   | 精神；灵魂。                                                               | absent        | 0 / 0                        | NC                   |
| [spirit-temple-zelda](../entries/spirit-temple-zelda.yml)                                                             | Spirit Temple                            | 魂之神殿；《王国之泪》中的地点或设施。                                     | absent        | 0 / 0                        | NC                   |
| [spring-mechanical](../entries/spring-mechanical.yml)                                                                 | spring                                   | 弹簧；受力形变后能恢复原状的弹性元件。                                     | absent        | 0 / 1                        | NC-N                 |
| [spring-water](../entries/spring-water.yml)                                                                           | spring                                   | 泉；泉水；地下水自然涌出之处。                                             | N             | 1 / 1                        | K                    |
| [sprint-general](../entries/sprint-general.yml)                                                                       | sprint                                   | 冲刺                                                                       | absent        | 0 / 0                        | NC                   |
| [stabilizer-engineering](../entries/stabilizer-engineering.yml)                                                       | stabilizer                               | 稳定器                                                                     | W             | 0 / 0                        | K                    |
| [stable-architecture](../entries/stable-architecture.yml)                                                             | stable                                   | 马厩；用于安置马匹的建筑。                                                 | absent        | 0 / 0                        | NC                   |
| [stack-data-structure](../entries/stack-data-structure.yml)                                                           | stack                                    | 栈；按后进先出顺序添加和移除元素的数据结构。                               | absent        | 0 / 0                        | R: R09               |
| [stack-trace-tech](../entries/stack-trace-tech.yml)                                                                   | stack trace                              | 调用栈跟踪；记录一条执行路径上活动调用帧的信息。                           | absent        | 0 / 0                        | R: R21               |
| [staged-git](../entries/staged-git.yml)                                                                               | staged                                   | 已暂存的                                                                   | N             | 1 / 1                        | K                    |
| [staircase-architecture](../entries/staircase-architecture.yml)                                                       | staircase                                | 楼梯                                                                       | W             | 0 / 0                        | K                    |
| [stake-engineering](../entries/stake-engineering.yml)                                                                 | stake                                    | 桩、固定钉                                                                 | absent        | 0 / 0                        | NC                   |
| [stale-tech](../entries/stale-tech.yml)                                                                               | stale                                    | 过时的；未更新的                                                           | absent        | 0 / 0                        | NC                   |
| [stamina-general](../entries/stamina-general.yml)                                                                     | stamina                                  | 耐力；持续活动的体力；游戏中对应精力。                                     | E,N           | 1 / 1                        | K                    |
| [star-method-interview](../entries/star-method-interview.yml)                                                         | STAR method                              | STAR 方法；按情境、任务、行动和结果组织真实经历。                          | absent        | 0 / 0                        | NC-N                 |
| [static-initialization-tech](../entries/static-initialization-tech.yml)                                               | static initialization                    | 静态初始化                                                                 | absent        | 0 / 0                        | NC                   |
| [static-inspection-tech](../entries/static-inspection-tech.yml)                                                       | static inspection                        | 静态检查                                                                   | absent        | 0 / 0                        | R: R22               |
| [statue-architecture](../entries/statue-architecture.yml)                                                             | statue                                   | 雕像                                                                       | absent        | 0 / 0                        | NC                   |
| [stealth-mask-zelda](../entries/stealth-mask-zelda.yml)                                                               | Stealth Mask                             | 潜行面罩；《王国之泪》中的防具名称。                                       | absent        | 0 / 0                        | NC                   |
| [stealth-up-zelda](../entries/stealth-up-zelda.yml)                                                                   | Stealth Up                               | 安静性提升；《王国之泪》中的能力效果名称。                                 | absent        | 0 / 0                        | NC                   |
| [steam-cooking](../entries/steam-cooking.yml)                                                                         | steam                                    | 蒸；利用蒸汽加热食物。                                                     | N             | 0 / 0                        | K                    |
| [steer-general](../entries/steer-general.yml)                                                                         | steer                                    | 控制方向；掌舵；驾驶时使车辆转向。                                         | absent        | 0 / 0                        | NC                   |
| [steering-stick-zelda](../entries/steering-stick-zelda.yml)                                                           | Steering Stick                           | 操纵杆；《王国之泪》中的左纳乌装置。                                       | W             | 0 / 0                        | K                    |
| [stew-cooking](../entries/stew-cooking.yml)                                                                           | stew                                     | 炖菜                                                                       | absent        | 0 / 1                        | NC                   |
| [stick-general](../entries/stick-general.yml)                                                                         | stick                                    | 棍；条状或杆状物。                                                         | absent        | 0 / 0                        | NC                   |
| [stone-general](../entries/stone-general.yml)                                                                         | stone                                    | 石头；石料。                                                               | absent        | 0 / 0                        | NC                   |
| [stone-talus-zelda](../entries/stone-talus-zelda.yml)                                                                 | Stone Talus                              | 岩石巨人；《王国之泪》中的怪物或首领名称。                                 | absent        | 0 / 0                        | NC                   |
| [storage-event-browser](../entries/storage-event-browser.yml)                                                         | storage event                            | 存储事件；其他文档修改共享存储区域时发出的通知。                           | absent        | 0 / 0                        | NC                   |
| [structural-check-tech](../entries/structural-check-tech.yml)                                                         | structural check                         | 结构检查；此处指按约定比较值及其嵌套结构、报告差异的检查。                 | absent        | 0 / 0                        | NC-N                 |
| [sturdy-general](../entries/sturdy-general.yml)                                                                       | sturdy                                   | 坚固的                                                                     | N             | 1 / 1                        | K                    |
| [submission-workflow](../entries/submission-workflow.yml)                                                             | submission                               | 提交审查                                                                   | absent        | 0 / 0                        | NC                   |
| [subsequent-tech](../entries/subsequent-tech.yml)                                                                     | subsequent                               | 后续的；随后的                                                             | absent        | 0 / 0                        | NC                   |
| [sugar-food](../entries/sugar-food.yml)                                                                               | sugar                                    | 糖。                                                                       | absent        | 0 / 0                        | NC                   |
| [summit-terrain](../entries/summit-terrain.yml)                                                                       | summit                                   | 山顶                                                                       | N             | 1 / 1                        | K                    |
| [summon-general](../entries/summon-general.yml)                                                                       | summon                                   | 召唤                                                                       | absent        | 0 / 0                        | NC                   |
| [sundelion-zelda](../entries/sundelion-zelda.yml)                                                                     | Sundelion                                | 向阳草；《王国之泪》中的材料名称。                                         | absent        | 0 / 0                        | NC                   |
| [sunny-general](../entries/sunny-general.yml)                                                                         | sunny                                    | 阳光明媚的；充满阳光的。                                                   | absent        | 0 / 0                        | NC                   |
| [support-general](../entries/support-general.yml)                                                                     | support                                  | 支持；提供依据                                                             | absent        | 0 / 0                        | NC                   |
| [surface-terrain](../entries/surface-terrain.yml)                                                                     | surface                                  | 地表                                                                       | absent        | 1 / 1                        | NC                   |
| [surrogate-pair-tech](../entries/surrogate-pair-tech.yml)                                                             | surrogate pair                           | 代理对                                                                     | absent        | 0 / 1                        | NC                   |
| [susceptible-combat](../entries/susceptible-combat.yml)                                                               | susceptible                              | 易受某种影响的                                                             | N             | 0 / 1                        | K                    |
| [suspected-tech](../entries/suspected-tech.yml)                                                                       | suspected                                | 疑似的；尚未证实的                                                         | absent        | 0 / 0                        | NC                   |
| [swamp-terrain](../entries/swamp-terrain.yml)                                                                         | swamp                                    | 沼泽                                                                       | absent        | 0 / 0                        | NC                   |
| [sweet-and-sour-food](../entries/sweet-and-sour-food.yml)                                                             | sweet and sour                           | 糖醋的；酸甜的                                                             | absent        | 0 / 0                        | NC                   |
| [sword-weapon](../entries/sword-weapon.yml)                                                                           | sword                                    | 剑                                                                         | absent        | 0 / 0                        | NC                   |
| [symbol-javascript](../entries/symbol-javascript.yml)                                                                 | Symbol                                   | Symbol；用于表示独一且不可变标识值的 JavaScript 原始类型。                 | absent        | 0 / 0                        | NC                   |
| [symlink-tech](../entries/symlink-tech.yml)                                                                           | symlink                                  | 符号链接                                                                   | W             | 0 / 0                        | K                    |
| [synchronize-tech](../entries/synchronize-tech.yml)                                                                   | synchronize                              | 同步                                                                       | absent        | 0 / 0                        | NC                   |
| [syntax-tech](../entries/syntax-tech.yml)                                                                             | syntax                                   | 语法；规定语言符号如何组成合法结构的规则。                                 | E             | 0 / 0                        | K                    |
| [synthesize-tech](../entries/synthesize-tech.yml)                                                                     | synthesize                               | 合成；构造                                                                 | absent        | 0 / 0                        | NC-N                 |
| [synthetic-tech](../entries/synthetic-tech.yml)                                                                       | synthetic                                | 人工构造的                                                                 | absent        | 0 / 0                        | NC                   |
| [system-ui](../entries/system-ui.yml)                                                                                 | system                                   | 系统；由相互联系的部分组成的整体。                                         | absent        | 0 / 0                        | NC                   |
| [tarrey-town-zelda](../entries/tarrey-town-zelda.yml)                                                                 | Tarrey Town                              | 一始村；《王国之泪》中的地点或设施。                                       | absent        | 0 / 0                        | NC                   |
| [tear-general](../entries/tear-general.yml)                                                                           | tear                                     | 眼泪；泪水。                                                               | N             | 0 / 0                        | R: R19; K            |
| [teba-zelda](../entries/teba-zelda.yml)                                                                               | Teba                                     | 特巴；利特族战士，丘栗的父亲。                                             | absent        | 0 / 0                        | NC                   |
| [temple-architecture](../entries/temple-architecture.yml)                                                             | temple                                   | 神殿、寺庙                                                                 | absent        | 0 / 0                        | NC                   |
| [temple-of-time-zelda](../entries/temple-of-time-zelda.yml)                                                           | Temple of Time                           | 时之神殿；《王国之泪》中的地点或设施。                                     | absent        | 0 / 0                        | NC                   |
| [temporal-dead-zone-javascript](../entries/temporal-dead-zone-javascript.yml)                                         | temporal dead zone                       | 暂时性死区；词法绑定从作用域开始到初始化前不可访问的阶段。                 | absent        | 0 / 1                        | NC-N                 |
| [teriyaki-food](../entries/teriyaki-food.yml)                                                                         | teriyaki                                 | 照烧                                                                       | absent        | 0 / 0                        | NC                   |
| [test-suite-tech](../entries/test-suite-tech.yml)                                                                     | test suite                               | 测试套件                                                                   | absent        | 0 / 0                        | NC                   |
| [test-tech](../entries/test-tech.yml)                                                                                 | test                                     | 测试；用规定的检查评估软件行为或属性。                                     | absent        | 0 / 0                        | NC                   |
| [text-to-speech-tech](../entries/text-to-speech-tech.yml)                                                             | text-to-speech                           | 文本转语音；把书面文字合成为可听语音的技术。                               | absent        | 0 / 0                        | NC-N                 |
| [the-depths-zelda](../entries/the-depths-zelda.yml)                                                                   | The Depths                               | 地底；本作位于地表下方的广大区域。                                         | absent        | 0 / 1                        | NC                   |
| [throw-general](../entries/throw-general.yml)                                                                         | throw                                    | 投掷                                                                       | absent        | 0 / 0                        | NC                   |
| [tights-clothing](../entries/tights-clothing.yml)                                                                     | tights                                   | 紧身裤；连裤袜。                                                           | absent        | 0 / 0                        | NC                   |
| [time-bomb-zelda](../entries/time-bomb-zelda.yml)                                                                     | Time Bomb                                | 计时炸弹；《王国之泪》中的左纳乌装置。                                     | absent        | 0 / 0                        | NC                   |
| [timid-general](../entries/timid-general.yml)                                                                         | timid                                    | 胆小的                                                                     | absent        | 0 / 0                        | NC                   |
| [toggle-tech](../entries/toggle-tech.yml)                                                                             | toggle                                   | 切换                                                                       | absent        | 0 / 0                        | NC                   |
| [token-llm](../entries/token-llm.yml)                                                                                 | token                                    | 词元；语言模型将文本编码后使用的基本处理单位。                             | absent        | 0 / 1                        | NC-N                 |
| [tomato-food](../entries/tomato-food.yml)                                                                             | tomato                                   | 番茄。                                                                     | absent        | 0 / 0                        | NC                   |
| [toolchain-tech](../entries/toolchain-tech.yml)                                                                       | toolchain                                | 工具链                                                                     | W             | 0 / 0                        | K                    |
| [torch-weapon](../entries/torch-weapon.yml)                                                                           | torch                                    | 火把                                                                       | absent        | 0 / 0                        | NC                   |
| [tough-general](../entries/tough-general.yml)                                                                         | tough                                    | 坚韧的；结实的；难以对付的。                                               | absent        | 0 / 0                        | NC                   |
| [tower-architecture](../entries/tower-architecture.yml)                                                               | tower                                    | 高塔                                                                       | absent        | 0 / 0                        | NC                   |
| [tradeoff-general](../entries/tradeoff-general.yml)                                                                   | tradeoff                                 | 取舍                                                                       | W             | 0 / 0                        | K                    |
| [trail-story](../entries/trail-story.yml)                                                                             | trail                                    | 踪迹、线索                                                                 | absent        | 0 / 0                        | NC                   |
| [transformation-tech](../entries/transformation-tech.yml)                                                             | transformation                           | 转换                                                                       | absent        | 0 / 1                        | NC-N                 |
| [transport-failure-tech](../entries/transport-failure-tech.yml)                                                       | transport failure                        | 传输层故障                                                                 | absent        | 0 / 0                        | NC                   |
| [transport-metadata-tech](../entries/transport-metadata-tech.yml)                                                     | transport metadata                       | 传输元数据                                                                 | absent        | 0 / 0                        | NC                   |
| [traveler-general](../entries/traveler-general.yml)                                                                   | traveler                                 | 旅行者；旅人。                                                             | W             | 0 / 0                        | K                    |
| [travelers-sword-zelda](../entries/travelers-sword-zelda.yml)                                                         | Traveler's Sword                         | 旅人之剑；《王国之泪》中的装备名称。                                       | absent        | 0 / 0                        | NC                   |
| [traversal-tech](../entries/traversal-tech.yml)                                                                       | traversal                                | 遍历                                                                       | absent        | 0 / 0                        | NC-N                 |
| [tree-data-structure](../entries/tree-data-structure.yml)                                                             | tree                                     | 树；用节点的父子关系组织层级数据的结构。                                   | absent        | 0 / 0                        | NC-N                 |
| [tree-shaking-build](../entries/tree-shaking-build.yml)                                                               | tree shaking                             | 摇树优化；构建时识别并移除未使用代码的优化。                               | absent        | 0 / 0                        | NC-N                 |
| [tree-terrain](../entries/tree-terrain.yml)                                                                           | tree                                     | 树；木本植物。                                                             | absent        | 0 / 0                        | NC                   |
| [triage-workflow](../entries/triage-workflow.yml)                                                                     | triage                                   | 分流评估；先按影响、紧急程度和可处理性筛选待办问题。                       | absent        | 0 / 0                        | NC                   |
| [trident-weapon](../entries/trident-weapon.yml)                                                                       | trident                                  | 三叉戟                                                                     | E             | 0 / 0                        | K                    |
| [trigger-tech](../entries/trigger-tech.yml)                                                                           | trigger                                  | 触发条件                                                                   | absent        | 0 / 0                        | NC                   |
| [trousers-clothing](../entries/trousers-clothing.yml)                                                                 | trousers                                 | 长裤                                                                       | absent        | 0 / 0                        | NC                   |
| [truffle-food](../entries/truffle-food.yml)                                                                           | truffle                                  | 松露；生长在地下的食用真菌。                                               | absent        | 0 / 0                        | NC                   |
| [truncated-tech](../entries/truncated-tech.yml)                                                                       | truncated                                | 被截断的                                                                   | absent        | 0 / 0                        | NC                   |
| [tts-language-tech](../entries/tts-language-tech.yml)                                                                 | TTS language                             | 语音合成语言；文本转语音时使用的语言设置。                                 | absent        | 0 / 0                        | NC                   |
| [tulin-zelda](../entries/tulin-zelda.yml)                                                                             | Tulin                                    | 丘栗；利特族青年，风之贤者。                                               | absent        | 0 / 0                        | NC                   |
| [tunic-clothing](../entries/tunic-clothing.yml)                                                                       | tunic                                    | 束腰上衣、短袍                                                             | absent        | 0 / 0                        | NC                   |
| [two-handed-combat](../entries/two-handed-combat.yml)                                                                 | two-handed                               | 双手的                                                                     | W             | 0 / 0                        | K                    |
| [two-way-binding-frontend](../entries/two-way-binding-frontend.yml)                                                   | two-way binding                          | 双向绑定；让状态值与界面输入保持相互同步的绑定方式。                       | absent        | 0 / 1                        | NC-N                 |
| [type-argument-tech](../entries/type-argument-tech.yml)                                                               | type argument                            | 类型实参                                                                   | absent        | 0 / 1                        | NC                   |
| [type-assertion-tech](../entries/type-assertion-tech.yml)                                                             | type assertion                           | 类型断言                                                                   | absent        | 0 / 1                        | NC                   |
| [type-instantiation-tech](../entries/type-instantiation-tech.yml)                                                     | type instantiation                       | 类型实例化；为泛型类型提供具体类型实参并得到相应类型的过程。               | absent        | 0 / 0                        | NC-N                 |
| [type-tech](../entries/type-tech.yml)                                                                                 | type                                     | 类型；描述值的类别及允许操作或约束的信息。                                 | absent        | 0 / 0                        | NC                   |
| [typo-general](../entries/typo-general.yml)                                                                           | typo                                     | 笔误；打字错误                                                             | W             | 0 / 0                        | K                    |
| [ultrahand-zelda](../entries/ultrahand-zelda.yml)                                                                     | Ultrahand                                | 究极手；可抓取、移动和粘接许多物体的能力。                                 | W             | 0 / 0                        | K                    |
| [unclaimed-workflow](../entries/unclaimed-workflow.yml)                                                               | unclaimed                                | 尚未被承接的                                                               | W             | 0 / 0                        | K                    |
| [undead-combat](../entries/undead-combat.yml)                                                                         | undead                                   | 不死的                                                                     | W             | 0 / 0                        | K                    |
| [underground-general](../entries/underground-general.yml)                                                             | underground                              | 地下的                                                                     | W             | 0 / 0                        | K                    |
| [underline-tech](../entries/underline-tech.yml)                                                                       | underline                                | 下划线标记                                                                 | absent        | 0 / 0                        | NC                   |
| [unfuse-general](../entries/unfuse-general.yml)                                                                       | unfuse                                   | 解除融合                                                                   | W             | 0 / 1                        | K                    |
| [uninitialized-tech](../entries/uninitialized-tech.yml)                                                               | uninitialized                            | 未初始化的                                                                 | W             | 0 / 0                        | K                    |
| [union-tech](../entries/union-tech.yml)                                                                               | union                                    | 联合类型                                                                   | absent        | 0 / 0                        | NC                   |
| [unlock-general](../entries/unlock-general.yml)                                                                       | unlock                                   | 解锁                                                                       | W             | 0 / 0                        | K                    |
| [unmount-tech](../entries/unmount-tech.yml)                                                                           | unmount                                  | 卸载                                                                       | W,N           | 0 / 1                        | K                    |
| [unregister-tech](../entries/unregister-tech.yml)                                                                     | unregister                               | 取消注册                                                                   | W,N           | 0 / 1                        | K                    |
| [unstaged-git](../entries/unstaged-git.yml)                                                                           | unstaged                                 | 未暂存的                                                                   | W,N           | 0 / 1                        | K                    |
| [untracked-git](../entries/untracked-git.yml)                                                                         | untracked                                | 未跟踪的                                                                   | W             | 0 / 0                        | K                    |
| [upgrade-general](../entries/upgrade-general.yml)                                                                     | upgrade                                  | 升级、强化                                                                 | absent        | 0 / 0                        | NC                   |
| [upheaval-story](../entries/upheaval-story.yml)                                                                       | upheaval                                 | 剧变；动荡；重大变动。                                                     | absent        | 0 / 0                        | NC                   |
| [upstream-git](../entries/upstream.yml)                                                                               | upstream                                 | 上游的                                                                     | W             | 0 / 0                        | K                    |
| [validation-tech](../entries/validation-tech.yml)                                                                     | validation                               | 验证                                                                       | absent        | 0 / 0                        | A: WF05; R: R02      |
| [valley-terrain](../entries/valley-terrain.yml)                                                                       | valley                                   | 山谷                                                                       | absent        | 0 / 0                        | NC                   |
| [variance-type-system](../entries/variance-type-system.yml)                                                           | variance                                 | 型变；类型参数之间的子类型关系如何影响复合类型之间的关系。                 | absent        | 0 / 0                        | NC-N                 |
| [verification-tech](../entries/verification-tech.yml)                                                                 | verification                             | 核实；验证                                                                 | absent        | 0 / 0                        | A: WF06; R: R02      |
| [version-pin-tech](../entries/version-pin-tech.yml)                                                                   | version pin                              | 版本固定                                                                   | absent        | 0 / 0                        | NC                   |
| [view-general](../entries/view-general.yml)                                                                           | view                                     | 景色；视野；看到的景象。                                                   | absent        | 0 / 0                        | NC                   |
| [village-architecture](../entries/village-architecture.yml)                                                           | village                                  | 村庄                                                                       | absent        | 0 / 0                        | NC                   |
| [virtual-dom-frontend](../entries/virtual-dom-frontend.yml)                                                           | virtual DOM                              | 虚拟 DOM；用内存中的节点结构描述目标界面。                                 | N             | 0 / 0                        | K                    |
| [vocabulary-general](../entries/vocabulary.yml)                                                                       | vocabulary                               | 词汇；词汇量                                                               | E             | 0 / 0                        | K                    |
| [volcanic-combat](../entries/volcanic-combat.yml)                                                                     | volcanic                                 | 火山的                                                                     | absent        | 0 / 0                        | NC                   |
| [vulnerable-combat](../entries/vulnerable-combat.yml)                                                                 | vulnerable                               | 易受攻击的                                                                 | E,N           | 1 / 1                        | K                    |
| [waistguard-clothing](../entries/waistguard-clothing.yml)                                                             | waistguard                               | 护腰                                                                       | W             | 0 / 0                        | K                    |
| [warp-general](../entries/warp-general.yml)                                                                           | warp                                     | 传送；游戏中快速移动到另一地点。                                           | absent        | 0 / 0                        | NC-N                 |
| [warrior-general](../entries/warrior-general.yml)                                                                     | warrior                                  | 战士                                                                       | absent        | 0 / 0                        | NC                   |
| [water-general](../entries/water-general.yml)                                                                         | water                                    | 水。                                                                       | absent        | 0 / 0                        | NC                   |
| [water-temple-zelda](../entries/water-temple-zelda.yml)                                                               | Water Temple                             | 水之神殿；《王国之泪》中的地点或设施。                                     | absent        | 0 / 0                        | NC                   |
| [waterfall-terrain](../entries/waterfall-terrain.yml)                                                                 | waterfall                                | 瀑布                                                                       | W             | 0 / 0                        | K                    |
| [weapon-weapon](../entries/weapon-weapon.yml)                                                                         | weapon                                   | 武器。                                                                     | absent        | 0 / 0                        | NC                   |
| [web-storage-browser](../entries/web-storage-browser.yml)                                                             | web storage                              | Web 存储；浏览器提供的按源隔离的同步键值存储。                             | absent        | 0 / 0                        | NC-N                 |
| [well-architecture](../entries/well-architecture.yml)                                                                 | well                                     | 水井                                                                       | absent        | 0 / 0                        | NC                   |
| [wetlands-terrain](../entries/wetlands-terrain.yml)                                                                   | wetlands                                 | 湿地                                                                       | W             | 0 / 0                        | K                    |
| [wheat-food](../entries/wheat-food.yml)                                                                               | wheat                                    | 小麦。                                                                     | absent        | 0 / 0                        | NC                   |
| [wheel-engineering](../entries/wheel-engineering.yml)                                                                 | wheel                                    | 轮子。                                                                     | absent        | 0 / 0                        | NC                   |
| [when-applicable-workflow](../entries/when-applicable-workflow.yml)                                                   | when applicable                          | 在适用时                                                                   | absent        | 0 / 0                        | NC                   |
| [where-practical-workflow](../entries/where-practical-workflow.yml)                                                   | where practical                          | 在实际可行时                                                               | absent        | 0 / 0                        | NC                   |
| [wield-story](../entries/wield-story.yml)                                                                             | wield                                    | 挥舞、使用武器                                                             | absent        | 0 / 0                        | NC                   |
| [wind-temple-zelda](../entries/wind-temple-zelda.yml)                                                                 | Wind Temple                              | 风之神殿；《王国之泪》中的地点或设施。                                     | absent        | 0 / 0                        | NC                   |
| [wind-terrain](../entries/wind-terrain.yml)                                                                           | wind                                     | 风。                                                                       | absent        | 0 / 0                        | NC                   |
| [wing-engineering](../entries/wing-engineering.yml)                                                                   | wing                                     | 翼；翅膀。                                                                 | absent        | 0 / 0                        | NC                   |
| [wise-general](../entries/wise-general.yml)                                                                           | wise                                     | 睿智的                                                                     | absent        | 0 / 0                        | NC                   |
| [wizzrobe-zelda](../entries/wizzrobe-zelda.yml)                                                                       | Wizzrobe                                 | 长袍魔法师；《王国之泪》中的怪物或首领名称。                               | absent        | 0 / 0                        | NC                   |
| [word-general](../entries/word-general.yml)                                                                           | word                                     | 词；单词                                                                   | absent        | 0 / 0                        | NC                   |
| [wording-general](../entries/wording-general.yml)                                                                     | wording                                  | 措辞                                                                       | absent        | 0 / 0                        | NC                   |
| [workaround-tech](../entries/workaround-tech.yml)                                                                     | workaround                               | 变通办法；暂时避开问题而不直接修复根因的做法。                             | absent        | 0 / 0                        | R: R06               |
| [worker-pool-concurrency](../entries/worker-pool-concurrency.yml)                                                     | worker pool                              | 工作池；用有限工作单元处理排队任务并控制并发。                             | absent        | 0 / 0                        | NC-N                 |
| [workload-tech](../entries/workload-tech.yml)                                                                         | workload                                 | 工作负载                                                                   | W             | 0 / 0                        | K                    |
| [would-you-like-a-or-b](../entries/would-you-like-a-or-b.yml)                                                         | Would you like A or B?                   | 您想要 A 还是 B？                                                          | absent        | 0 / 0                        | NC-N                 |
| [yiga-clan-hideout-zelda](../entries/yiga-clan-hideout-zelda.yml)                                                     | Yiga Clan Hideout                        | 依盖队基地；《王国之泪》中的地点或设施。                                   | absent        | 0 / 0                        | NC                   |
| [yiga-clan-zelda](../entries/yiga-clan-zelda.yml)                                                                     | Yiga Clan                                | 依盖队；《王国之泪》中的族群或组织。                                       | absent        | 0 / 0                        | NC                   |
| [yona-zelda](../entries/yona-zelda.yml)                                                                               | Yona                                     | 优娜；卓拉族人，希多的伴侣。                                               | absent        | 0 / 0                        | NC                   |
| [yunobo-zelda](../entries/yunobo-zelda.yml)                                                                           | Yunobo                                   | 阿沅；鼓隆族人，火之贤者。                                                 | absent        | 0 / 0                        | NC                   |
| [zonai-device-zelda](../entries/zonai-device-zelda.yml)                                                               | Zonai Device                             | 左纳乌装置；使用左纳乌技术的功能装置。                                     | absent        | 0 / 0                        | NC                   |
| [zonai-zelda](../entries/zonai-zelda.yml)                                                                             | Zonai                                    | 左纳乌族；《王国之泪》中的族群或组织。                                     | absent        | 0 / 0                        | NC                   |
| [zonaite-zelda](../entries/zonaite-zelda.yml)                                                                         | Zonaite                                  | 左纳尼乌姆；用于左纳乌技术相关制作和交换的矿石材料。                       | absent        | 0 / 0                        | NC                   |
| [zora-armor-zelda](../entries/zora-armor-zelda.yml)                                                                   | Zora Armor                               | 卓拉铠甲；《王国之泪》中的防具名称。                                       | absent        | 0 / 0                        | NC                   |
| [zora-zelda](../entries/zora-zelda.yml)                                                                               | Zora                                     | 卓拉族；《王国之泪》中的族群或组织。                                       | absent        | 0 / 0                        | NC                   |
| [zoras-domain-zelda](../entries/zoras-domain-zelda.yml)                                                               | Zora's Domain                            | 卓拉领地；《王国之泪》中的地点或设施。                                     | absent        | 0 / 0                        | NC                   |

### Japanese (242)

| Stable ID / actual file                                                                                 | Word                                         | Current Meaning                                     | Existing aids | Relations: local / displayed | Decision             |
| :------------------------------------------------------------------------------------------------------ | :------------------------------------------- | :-------------------------------------------------- | :------------ | :--------------------------- | :------------------- |
| [aitemu](../entries/aitemu.yml)                                                                         | アイテム                                     | 物品；项目；道具                                    | E             | 0 / 0                        | K                    |
| [amerika](../entries/amerika.yml)                                                                       | アメリカ                                     | 美国                                                | absent        | 0 / 0                        | NC                   |
| [amerika-kara-kimashita](../entries/amerika-kara-kimashita.yml)                                         | アメリカから来ました                         | 我来自美国                                          | absent        | 0 / 0                        | A: WF12              |
| [ana](../entries/ana.yml)                                                                               | 穴                                           | 洞；坑                                              | absent        | 0 / 0                        | NC                   |
| [anata](../entries/anata.yml)                                                                           | あなた                                       | 你                                                  | absent        | 1 / 1                        | A: N06; R: R12       |
| [anime](../entries/anime.yml)                                                                           | アニメ                                       | 动画                                                | W             | 0 / 0                        | K                    |
| [anime-guzzu](../entries/anime-guzzu.yml)                                                               | アニメグッズ                                 | 动漫周边                                            | W             | 0 / 0                        | K                    |
| [ano-hito](../entries/ano-hito.yml)                                                                     | あの人                                       | 那个人                                              | absent        | 1 / 1                        | NC-N                 |
| [ano-kata](../entries/ano-kata.yml)                                                                     | あの方                                       | 那位；礼貌说法                                      | absent        | 0 / 1                        | NC-N                 |
| [baagaa](../entries/baagaa.yml)                                                                         | バーガー                                     | 汉堡；汉堡类食品                                    | absent        | 0 / 0                        | D: D01               |
| [bangou](../entries/bangou.yml)                                                                         | 番号                                         | 号码                                                | absent        | 0 / 0                        | NC                   |
| [bangou-o-omachigae-phone](../entries/bangou-o-omachigae-phone.yml)                                     | 番号をお間違えではないでしょうか。           | 您可能打错号码了吧？                                | absent        | 0 / 0                        | NC                   |
| [boku](../entries/boku.yml)                                                                             | 僕                                           | 我；男性自称                                        | absent        | 1 / 1                        | R: R11               |
| [boku-wa-shinsekai-no-kami-ni-naru](../entries/boku-wa-shinsekai-no-kami-ni-naru.yml)                   | 僕は新世界の神になる                         | 我将成为新世界的神                                  | W             | 0 / 0                        | K                    |
| [burajiru](../entries/burajiru.yml)                                                                     | ブラジル                                     | 巴西                                                | absent        | 0 / 0                        | NC                   |
| [burajiru-eaa](../entries/burajiru-eaa.yml)                                                             | ブラジルエアー                               | 用户所给的航空公司名称；ブラジルエアー              | W             | 0 / 0                        | K; D: D04            |
| [byouin](../entries/byouin.yml)                                                                         | 病院                                         | 医院                                                | absent        | 0 / 0                        | NC                   |
| [chan-suffix](../entries/chan-suffix.yml)                                                               | ちゃん                                       | 亲昵称呼后缀                                        | absent        | 0 / 0                        | NC                   |
| [chikara](../entries/chikara.yml)                                                                       | 力                                           | 力量；能力                                          | absent        | 0 / 0                        | NC                   |
| [chikyuu](../entries/chikyuu.yml)                                                                       | 地球                                         | 地球                                                | absent        | 0 / 0                        | NC-N                 |
| [chisou](../entries/chisou.yml)                                                                         | 馳走                                         | 奔走准备；款待的词源形式                            | E             | 0 / 0                        | K                    |
| [chuugoku](../entries/chuugoku.yml)                                                                     | 中国                                         | 中国                                                | absent        | 0 / 0                        | NC                   |
| [chuugokujin](../entries/chuugokujin.yml)                                                               | 中国人                                       | 中国人                                              | W             | 0 / 0                        | K                    |
| [daigaku](../entries/daigaku.yml)                                                                       | 大学                                         | 大学                                                | absent        | 0 / 0                        | NC                   |
| [daijoubu](../entries/daijoubu.yml)                                                                     | 大丈夫                                       | 没问题；可以；不要紧                                | N             | 0 / 0                        | K                    |
| [dake-grammar](../entries/dake-grammar.yml)                                                             | ～だけ                                       | 只是……；仅仅……                                      | absent        | 0 / 0                        | NC                   |
| [dare](../entries/dare.yml)                                                                             | 誰                                           | 谁                                                  | absent        | 1 / 1                        | NC-N                 |
| [denki](../entries/denki.yml)                                                                           | 電気                                         | 电；电能                                            | absent        | 0 / 0                        | NC                   |
| [denwa](../entries/denwa.yml)                                                                           | 電話                                         | 电话                                                | absent        | 0 / 0                        | NC                   |
| [denwa-bangou](../entries/denwa-bangou.yml)                                                             | 電話番号                                     | 电话号码                                            | W             | 0 / 0                        | K                    |
| [denwa-ga-aru](../entries/denwa-ga-aru.yml)                                                             | 電話がある                                   | 有电话；接到电话                                    | absent        | 0 / 0                        | NC                   |
| [deshou-ka-grammar](../entries/deshou-ka-grammar.yml)                                                   | ～でしょうか                                 | 请问是否……；会不会……                                | absent        | 0 / 0                        | NC-N                 |
| [dewa-arimasen-grammar](../entries/dewa-arimasen-grammar.yml)                                           | ～ではありません                             | 不是……                                              | absent        | 1 / 1                        | NC-N                 |
| [dewa-nai-deshou-ka-grammar](../entries/dewa-nai-deshou-ka-grammar.yml)                                 | ～ではないでしょうか                         | 会不会是……；是不是……                                | absent        | 0 / 0                        | NC                   |
| [doitsu](../entries/doitsu.yml)                                                                         | ドイツ                                       | 德国                                                | absent        | 0 / 0                        | D: D01               |
| [donata](../entries/donata.yml)                                                                         | どなた                                       | 哪位                                                | absent        | 0 / 1                        | NC-N                 |
| [donata-o-osagashi-phone](../entries/donata-o-osagashi-phone.yml)                                       | どなたをお探しですか？                       | 请问您找哪位？                                      | absent        | 0 / 0                        | NC-N                 |
| [eaa](../entries/eaa.yml)                                                                               | エアー                                       | 空气；航空                                          | E             | 0 / 0                        | K                    |
| [edogawa-konan](../entries/edogawa-konan.yml)                                                           | 江戸川コナン                                 | 江户川柯南                                          | absent        | 0 / 0                        | D: D05               |
| [ee-kee-shii](../entries/ee-kee-shii.yml)                                                               | エー・ケー・シー                             | 名为 A・K・C 的组织；其名称的日语字母读音           | absent        | 0 / 0                        | D: D04               |
| [eigo](../entries/eigo.yml)                                                                             | 英語                                         | 英语                                                | absent        | 0 / 0                        | NC                   |
| [eigo-de-daijoubu-phone](../entries/eigo-de-daijoubu-phone.yml)                                         | 英語で大丈夫ですか？                         | 可以说英语吗？                                      | absent        | 0 / 0                        | NC                   |
| [fisshu](../entries/fisshu.yml)                                                                         | フィッシュ                                   | fish；鱼类相关借词                                  | E             | 1 / 1                        | K                    |
| [fisshu-baagaa](../entries/fisshu-baagaa.yml)                                                           | フィッシュバーガー                           | 鱼肉汉堡                                            | W             | 0 / 0                        | K                    |
| [fuji-daigaku](../entries/fuji-daigaku.yml)                                                             | 富士大学                                     | 大学名称；富士大学                                  | absent        | 0 / 0                        | D: D04               |
| [fuminijiru](../entries/fuminijiru.yml)                                                                 | 踏みにじる                                   | 践踏；蹂躏                                          | absent        | 0 / 0                        | NC                   |
| [futatsu](../entries/futatsu.yml)                                                                       | 二つ                                         | 两个                                                | absent        | 0 / 0                        | NC                   |
| [fuuin](../entries/fuuin.yml)                                                                           | 封印                                         | 封印                                                | absent        | 0 / 0                        | NC                   |
| [fuuin-kaijo-reriizu](../entries/fuuin-kaijo-reriizu.yml)                                               | 封印解除（レリーズ）！                       | 封印解除（Release）！                               | W,N           | 0 / 0                        | K                    |
| [gaikoku](../entries/gaikoku.yml)                                                                       | 外国                                         | 外国；本国以外的国家                                | absent        | 0 / 0                        | R: R23               |
| [gaikokujin](../entries/gaikokujin.yml)                                                                 | 外国人                                       | 外国人                                              | W             | 0 / 0                        | A: N07; R: R23; K    |
| [gakusei](../entries/gakusei.yml)                                                                       | 学生                                         | 学生                                                | absent        | 0 / 0                        | NC                   |
| [gentei](../entries/gentei.yml)                                                                         | 限定                                         | 限定；限制范围或数量                                | absent        | 0 / 0                        | NC-N                 |
| [gentei-guzzu](../entries/gentei-guzzu.yml)                                                             | 限定グッズ                                   | 限定周边                                            | W             | 0 / 0                        | K                    |
| [gimu](../entries/gimu.yml)                                                                             | 義務                                         | 义务                                                | absent        | 0 / 0                        | NC                   |
| [gimu-toka-janai-yo](../entries/gimu-toka-janai-yo.yml)                                                 | 義務とかじゃないよ                           | 这不是什么义务                                      | W             | 0 / 0                        | K                    |
| [gochisou](../entries/gochisou.yml)                                                                     | ごちそう                                     | 盛宴；好饭菜；款待                                  | E             | 0 / 0                        | K                    |
| [gochisousama](../entries/gochisousama.yml)                                                             | ごちそうさま                                 | 多谢款待                                            | absent        | 2 / 2                        | NC                   |
| [gochisousama-deshita](../entries/gochisousama-deshita.yml)                                             | ごちそうさまでした                           | 多谢款待；谢谢这顿饭                                | absent        | 0 / 1                        | NC-N                 |
| [gomen-ne](../entries/gomen-ne.yml)                                                                     | ごめんね                                     | 抱歉啦；对不起                                      | absent        | 1 / 1                        | NC                   |
| [gomennasai](../entries/gomennasai.yml)                                                                 | ごめんなさい                                 | 对不起；请原谅                                      | absent        | 1 / 2                        | NC                   |
| [guzzu](../entries/guzzu.yml)                                                                           | グッズ                                       | 周边；商品                                          | E             | 0 / 0                        | K                    |
| [hai](../entries/hai.yml)                                                                               | はい                                         | 是；好的                                            | absent        | 0 / 0                        | NC                   |
| [hajimemashite](../entries/hajimemashite.yml)                                                           | はじめまして                                 | 初次见面                                            | absent        | 0 / 0                        | NC                   |
| [hanasemasen](../entries/hanasemasen.yml)                                                               | 話せません                                   | 不会说；不能说                                      | W,N           | 0 / 1                        | K                    |
| [hanasu](../entries/hanasu.yml)                                                                         | 話す                                         | 说；讲话                                            | N             | 1 / 1                        | K                    |
| [hassai](../entries/hassai.yml)                                                                         | 八歳                                         | 八岁                                                | absent        | 0 / 0                        | NC                   |
| [hikari](../entries/hikari.yml)                                                                         | 光                                           | 光                                                  | absent        | 0 / 0                        | NC                   |
| [hikari-ni-nareru-nda](../entries/hikari-ni-nareru-nda.yml)                                             | 光になれるんだ                               | 能够成为光                                          | W             | 0 / 0                        | K                    |
| [hito](../entries/hito.yml)                                                                             | 人                                           | 人                                                  | absent        | 0 / 0                        | NC                   |
| [hito-o-norowaba-ana-futatsu](../entries/hito-o-norowaba-ana-futatsu.yml)                               | 人を呪わば穴二つ                             | 害人终害己                                          | N             | 0 / 0                        | K                    |
| [hitotsu](../entries/hitotsu.yml)                                                                       | 一つ                                         | 一个                                                | absent        | 0 / 0                        | NC                   |
| [hokage-naruto](../entries/hokage-naruto.yml)                                                           | 火影                                         | 火影                                                | absent        | 0 / 0                        | NC                   |
| [hokage-ni-naru-made](../entries/hokage-ni-naru-made.yml)                                               | 火影になるまで                               | 直到成为火影                                        | W             | 0 / 0                        | K                    |
| [honjitsu](../entries/honjitsu.yml)                                                                     | 本日                                         | 今天；本日（正式表达）                              | absent        | 0 / 0                        | NC                   |
| [hontou-ni](../entries/hontou-ni.yml)                                                                   | 本当に                                       | 真的；确实                                          | absent        | 0 / 0                        | NC                   |
| [hontou-ni-gomennasai](../entries/hontou-ni-gomennasai.yml)                                             | 本当にごめんなさい                           | 真的很抱歉                                          | absent        | 0 / 0                        | NC                   |
| [igirisu](../entries/igirisu.yml)                                                                       | イギリス                                     | 英国                                                | absent        | 0 / 0                        | D: D01               |
| [iie](../entries/iie.yml)                                                                               | いいえ                                       | 不；不是                                            | absent        | 0 / 0                        | NC                   |
| [indoneshia](../entries/indoneshia.yml)                                                                 | インドネシア                                 | 印度尼西亚                                          | absent        | 0 / 0                        | NC                   |
| [ippen](../entries/ippen.yml)                                                                           | いっぺん                                     | 一次；一回                                          | absent        | 0 / 0                        | NC                   |
| [ippen-shinde-miru](../entries/ippen-shinde-miru.yml)                                                   | いっぺん、死んでみる？                       | 要不要死一次看看？                                  | absent        | 0 / 0                        | NC-N                 |
| [iryou-houjin-shadan](../entries/iryou-houjin-shadan.yml)                                               | 医療法人社団                                 | 医疗法人社团                                        | absent        | 0 / 0                        | NC                   |
| [isha](../entries/isha.yml)                                                                             | 医者                                         | 医生                                                | absent        | 0 / 0                        | NC                   |
| [issai](../entries/issai.yml)                                                                           | 一歳                                         | 一岁                                                | absent        | 0 / 0                        | NC                   |
| [itadakimasu](../entries/itadakimasu.yml)                                                               | いただきます                                 | 我开动了                                            | absent        | 0 / 1                        | NC                   |
| [itsumo](../entries/itsumo.yml)                                                                         | いつも                                       | 总是；一直                                          | absent        | 0 / 0                        | NC                   |
| [janai-grammar](../entries/janai-grammar.yml)                                                           | ～じゃない                                   | 不是……                                              | absent        | 0 / 1                        | NC                   |
| [jibun](../entries/jibun.yml)                                                                           | 自分                                         | 自己                                                | absent        | 0 / 0                        | R: R18               |
| [jibun-jishin](../entries/jibun-jishin.yml)                                                             | 自分自身                                     | 自己本身；自己本人                                  | absent        | 0 / 0                        | R: R18               |
| [jibun-jishin-no-chikara-de](../entries/jibun-jishin-no-chikara-de.yml)                                 | 自分自身の力で                               | 凭自己的力量                                        | W             | 0 / 0                        | K                    |
| [jin-nationality-suffix](../entries/jin-nationality-suffix.yml)                                         | ～人                                         | ……人；国籍后缀                                      | N             | 0 / 0                        | K                    |
| [jiyuu](../entries/jiyuu.yml)                                                                           | 自由                                         | 自由                                                | absent        | 0 / 0                        | NC                   |
| [jumoku](../entries/jumoku.yml)                                                                         | 樹木                                         | 树木                                                | absent        | 0 / 0                        | R: R15               |
| [kai-counter](../entries/kai-counter.yml)                                                               | 回                                           | 次；回                                              | absent        | 0 / 0                        | NC                   |
| [kaijo](../entries/kaijo.yml)                                                                           | 解除                                         | 解除；取消                                          | absent        | 0 / 0                        | NC                   |
| [kaisha](../entries/kaisha.yml)                                                                         | 会社                                         | 公司                                                | absent        | 0 / 0                        | NC                   |
| [kaishain](../entries/kaishain.yml)                                                                     | 会社員                                       | 公司职员                                            | W,N           | 1 / 1                        | K                    |
| [kami](../entries/kami.yml)                                                                             | 神                                           | 神                                                  | absent        | 0 / 0                        | R: R13               |
| [kangoshi](../entries/kangoshi.yml)                                                                     | 看護師                                       | 护士                                                | absent        | 0 / 0                        | NC-N                 |
| [kankoku](../entries/kankoku.yml)                                                                       | 韓国                                         | 韩国                                                | absent        | 0 / 0                        | NC                   |
| [kara-kimashita-origin](../entries/kara-kimashita-origin.yml)                                           | ～から来ました                               | 来自……                                              | absent        | 1 / 1                        | NC                   |
| [kara-reason-grammar](../entries/kara-reason-grammar.yml)                                               | ～から                                       | 因为……                                              | absent        | 0 / 1                        | NC                   |
| [kata-person](../entries/kata-person.yml)                                                               | 方                                           | 人；那位（礼貌说法）                                | absent        | 0 / 0                        | NC-N                 |
| [kenkyuu](../entries/kenkyuu.yml)                                                                       | 研究                                         | 研究；深入调查和探究                                | absent        | 0 / 0                        | R: R16               |
| [kenkyuusha](../entries/kenkyuusha.yml)                                                                 | 研究者                                       | 研究人员                                            | W             | 0 / 0                        | R: R16; K            |
| [ki](../entries/ki.yml)                                                                                 | 木                                           | 树；木                                              | absent        | 0 / 0                        | R: R15               |
| [kimi](../entries/kimi.yml)                                                                             | 君                                           | 你；用于较亲近、同辈或晚辈之间的称呼。              | absent        | 0 / 1                        | NC-N                 |
| [kimi-ga-yobu-namae-yume-no-tsuzuki](../entries/kimi-ga-yobu-namae-yume-no-tsuzuki.yml)                 | きみが呼ぶなまえ～夢のつづき                 | 你呼唤的名字～梦的延续；《続 夏目友人帳》原声曲名。 | W             | 0 / 0                        | K                    |
| [kochira-wa-san-desu](../entries/kochira-wa-san-desu.yml)                                               | こちらは～さんです                           | 这位是……                                            | absent        | 0 / 0                        | NC-N                 |
| [koneko](../entries/koneko.yml)                                                                         | 子猫                                         | 小猫                                                | W             | 1 / 1                        | R: R14; K            |
| [kono-bangou](../entries/kono-bangou.yml)                                                               | この番号                                     | 这个号码                                            | absent        | 0 / 0                        | NC                   |
| [kono-bangou-wa-sochira-no-byouin-phone](../entries/kono-bangou-wa-sochira-no-byouin-phone.yml)         | この番号はそちらの病院の電話番号でしょうか。 | 请问这个号码是贵院的电话号码吗？                    | absent        | 0 / 0                        | NC                   |
| [kosumo-o-moyase](../entries/kosumo-o-moyase.yml)                                                       | 小宇宙を燃やせ！                             | 燃烧小宇宙！                                        | absent        | 0 / 0                        | NC-N                 |
| [kosumo-saint-seiya](../entries/kosumo-saint-seiya.yml)                                                 | 小宇宙                                       | 小宇宙；作品中的宇宙能量                            | N             | 0 / 0                        | K                    |
| [kotoba](../entries/kotoba.yml)                                                                         | 言葉                                         | 话；语言；言辞                                      | absent        | 0 / 0                        | NC                   |
| [kotoba-wa-magenee](../entries/kotoba-wa-magenee.yml)                                                   | 言葉は曲げねェ                               | 不违背自己说过的话                                  | W             | 0 / 0                        | K                    |
| [koube](../entries/koube.yml)                                                                           | 神戸                                         | 神户                                                | absent        | 0 / 0                        | NC                   |
| [koube-byouin](../entries/koube-byouin.yml)                                                             | 神戸病院                                     | 医院名称；神戸病院                                  | absent        | 0 / 0                        | D: D04               |
| [koukishin](../entries/koukishin.yml)                                                                   | 好奇心                                       | 好奇心                                              | absent        | 0 / 0                        | NC                   |
| [koukishin-ousei-na-koneko](../entries/koukishin-ousei-na-koneko.yml)                                   | 好奇心旺盛な子猫                             | 好奇心旺盛的小猫                                    | absent        | 0 / 0                        | NC                   |
| [kyoushi](../entries/kyoushi.yml)                                                                       | 教師                                         | 教师                                                | N             | 1 / 1                        | K                    |
| [kyouto-fu](../entries/kyouto-fu.yml)                                                                   | 京都府                                       | 京都府                                              | absent        | 0 / 0                        | NC                   |
| [kyouto-shi](../entries/kyouto-shi.yml)                                                                 | 京都市                                       | 京都市                                              | absent        | 0 / 0                        | NC                   |
| [kyuu-number](../entries/kyuu-number.yml)                                                               | 九                                           | 九                                                  | absent        | 0 / 0                        | NC-N                 |
| [machigaeru](../entries/machigaeru.yml)                                                                 | 間違える                                     | 弄错；搞错                                          | absent        | 1 / 1                        | NC                   |
| [machigai](../entries/machigai.yml)                                                                     | 間違い                                       | 错误；弄错                                          | absent        | 0 / 1                        | NC                   |
| [made-grammar](../entries/made-grammar.yml)                                                             | ～まで                                       | 直到……；到……为止                                    | absent        | 0 / 0                        | NC                   |
| [mageru](../entries/mageru.yml)                                                                         | 曲げる                                       | 弯曲；改变；违背                                    | absent        | 0 / 0                        | NC                   |
| [massugu](../entries/massugu.yml)                                                                       | まっすぐ                                     | 笔直地；坦率地；坚定地                              | absent        | 0 / 0                        | NC                   |
| [minna](../entries/minna.yml)                                                                           | みんな                                       | 大家；所有人                                        | absent        | 0 / 0                        | NC                   |
| [moeagare-ore-no-kosumo](../entries/moeagare-ore-no-kosumo.yml)                                         | 燃え上がれ、俺の小宇宙！                     | 燃烧吧，我的小宇宙！                                | absent        | 0 / 0                        | NC-N                 |
| [moeagaru](../entries/moeagaru.yml)                                                                     | 燃え上がる                                   | 熊熊燃烧；燃烧起来                                  | absent        | 0 / 0                        | R: R20               |
| [morimoto](../entries/morimoto.yml)                                                                     | 森本                                         | 森本，日本姓氏                                      | absent        | 0 / 0                        | D: D05               |
| [mottomo-antei-ui](../entries/mottomo-antei-ui.yml)                                                     | 最も安定                                     | 最稳定                                              | absent        | 0 / 0                        | NC-N                 |
| [moushimasu](../entries/moushimasu.yml)                                                                 | 申します                                     | 叫作；自称（谦逊语）                                | absent        | 1 / 1                        | NC-N                 |
| [moushiwake-arimasen](../entries/moushiwake-arimasen.yml)                                               | 申し訳ありません                             | 非常抱歉；深表歉意                                  | absent        | 0 / 0                        | NC                   |
| [moyasu](../entries/moyasu.yml)                                                                         | 燃やす                                       | 燃烧；使……燃烧                                      | absent        | 0 / 0                        | R: R20               |
| [namae](../entries/namae.yml)                                                                           | 名前                                         | 名字                                                | absent        | 0 / 0                        | NC-N                 |
| [namida](../entries/namida.yml)                                                                         | 涙                                           | 眼泪                                                | absent        | 0 / 0                        | R: R19               |
| [nani-ga-da-rhetorical](../entries/nani-ga-da-rhetorical.yml)                                           | 何が～だ                                     | 这算什么……；还谈什么……                              | absent        | 0 / 0                        | NC-N                 |
| [nani-ga-seigi-da](../entries/nani-ga-seigi-da.yml)                                                     | 何が正義だ！                                 | 还谈什么正义！                                      | absent        | 0 / 0                        | NC-N                 |
| [nanoru](../entries/nanoru.yml)                                                                         | 名乗る                                       | 自报姓名；自称                                      | absent        | 0 / 1                        | NC                   |
| [nansai](../entries/nansai.yml)                                                                         | 何歳                                         | 几岁                                                | absent        | 1 / 1                        | NC-N                 |
| [nara](../entries/nara.yml)                                                                             | 奈良                                         | 奈良                                                | absent        | 0 / 0                        | NC                   |
| [naru](../entries/naru.yml)                                                                             | なる                                         | 成为；变成                                          | absent        | 0 / 0                        | NC                   |
| [nda-grammar](../entries/nda-grammar.yml)                                                               | ～んだ                                       | 用于说明、解释或强调                                | absent        | 0 / 0                        | NC                   |
| [neko](../entries/neko.yml)                                                                             | 猫                                           | 猫                                                  | absent        | 0 / 0                        | R: R14               |
| [neko-chan](../entries/neko-chan.yml)                                                                   | 猫ちゃん                                     | 猫猫；对猫的亲昵称呼                                | W             | 0 / 1                        | K                    |
| [ni-kawatte-grammar](../entries/ni-kawatte-grammar.yml)                                                 | ～にかわって                                 | 代替……；代表……                                      | absent        | 1 / 1                        | NC-N                 |
| [ni-nareru-grammar](../entries/ni-nareru-grammar.yml)                                                   | ～になれる                                   | 能够成为……                                          | W             | 1 / 1                        | K                    |
| [ni-naru-grammar](../entries/ni-naru-grammar.yml)                                                       | 〜になる                                     | 成为……；变成……                                      | absent        | 0 / 2                        | NC                   |
| [nihon](../entries/nihon.yml)                                                                           | 日本                                         | 日本                                                | absent        | 0 / 0                        | NC                   |
| [nihongo](../entries/nihongo.yml)                                                                       | 日本語                                       | 日语                                                | W             | 1 / 1                        | K                    |
| [nihongo-ga-hanasemasen](../entries/nihongo-ga-hanasemasen.yml)                                         | 日本語が話せません。                         | 我不会说日语                                        | absent        | 0 / 0                        | NC-N                 |
| [nihonjin](../entries/nihonjin.yml)                                                                     | 日本人                                       | 日本人                                              | W             | 0 / 1                        | K                    |
| [nihonjin-dewa-arimasen](../entries/nihonjin-dewa-arimasen.yml)                                         | 日本人ではありません                         | 我不是日本人                                        | absent        | 0 / 0                        | A: WF13              |
| [nindou-naruto](../entries/nindou-naruto.yml)                                                           | 忍道                                         | 忍者之道；自己的忍者信念                            | absent        | 0 / 0                        | NC                   |
| [ningen](../entries/ningen.yml)                                                                         | 人間                                         | 人；人类                                            | absent        | 0 / 0                        | NC                   |
| [no-ka-grammar](../entries/no-ka-grammar.yml)                                                           | ～のか                                       | 疑问；根据语气可表示强烈反问                        | N             | 0 / 0                        | K                    |
| [norou](../entries/norou.yml)                                                                           | 呪う                                         | 诅咒                                                | absent        | 0 / 0                        | NC                   |
| [ohayou-gozaimasu](../entries/ohayou-gozaimasu.yml)                                                     | おはようございます                           | 早上好                                              | absent        | 0 / 0                        | NC-N                 |
| [oikutsu](../entries/oikutsu.yml)                                                                       | おいくつ                                     | 几岁；多少                                          | absent        | 0 / 1                        | NC-N                 |
| [okane](../entries/okane.yml)                                                                           | お金                                         | 钱                                                  | absent        | 0 / 0                        | NC                   |
| [okurete-gomen-ne](../entries/okurete-gomen-ne.yml)                                                     | 遅れてごめんね                               | 抱歉我迟到了                                        | absent        | 0 / 0                        | NC                   |
| [omachigae](../entries/omachigae.yml)                                                                   | お間違え                                     | 弄错；「間違える」的礼貌表达形式                    | W             | 0 / 0                        | K                    |
| [omae](../entries/omae.yml)                                                                             | お前                                         | 你；较直接或粗鲁的第二人称                          | absent        | 0 / 0                        | R: R12               |
| [omae-no-namida-de-chikyuu-ga-sukueru-no-ka](../entries/omae-no-namida-de-chikyuu-ga-sukueru-no-ka.yml) | お前の涙でこの地球が救えるのか？             | 你的眼泪能拯救这个地球吗？                          | absent        | 0 / 0                        | NC-N                 |
| [omotenashi](../entries/omotenashi.yml)                                                                 | おもてなし                                   | 款待；待客之道                                      | absent        | 0 / 0                        | NC                   |
| [onamae-wa](../entries/onamae-wa.yml)                                                                   | お名前は                                     | 请问您的名字是？                                    | N             | 0 / 0                        | K                    |
| [ore](../entries/ore.yml)                                                                               | 俺                                           | 我；男性较随意或粗犷的自称                          | absent        | 0 / 1                        | NC                   |
| [ore-wa-ningen-dakara](../entries/ore-wa-ningen-dakara.yml)                                             | 俺は人間だから                               | 因为我是人类                                        | W             | 0 / 0                        | K                    |
| [osagashi](../entries/osagashi.yml)                                                                     | お探し                                       | 寻找；「探す」的礼貌表达形式                        | W             | 1 / 1                        | K                    |
| [oshioki](../entries/oshioki.yml)                                                                       | お仕置き                                     | 惩罚；教训                                          | absent        | 0 / 0                        | NC-N                 |
| [osusume](../entries/osusume.yml)                                                                       | おすすめ                                     | 推荐                                                | absent        | 0 / 0                        | NC                   |
| [ousei](../entries/ousei.yml)                                                                           | 旺盛                                         | 旺盛；强烈                                          | absent        | 0 / 0                        | NC                   |
| [pawaa](../entries/pawaa.yml)                                                                           | パワー                                       | 力量；动力                                          | E             | 0 / 0                        | K                    |
| [pawaa-denki](../entries/pawaa-denki.yml)                                                               | パワー電気                                   | 名称；パワー電気                                    | W             | 0 / 0                        | K; D: D04            |
| [rakuwakai-otowa-byouin](../entries/rakuwakai-otowa-byouin.yml)                                         | 洛和会音羽病院                               | 洛和会音羽医院                                      | absent        | 0 / 0                        | NC-N                 |
| [rakuwakai-otowa-byouin-kangoshi-morimoto](../entries/rakuwakai-otowa-byouin-kangoshi-morimoto.yml)     | 洛和会音羽病院の看護師、森本と申します。     | 我是洛和会音羽病院的护士，我姓森本                  | absent        | 0 / 0                        | NC                   |
| [sa-ending-grammar](../entries/sa-ending-grammar.yml)                                                   | ～さ                                         | 句尾表示断定、强调等语气                            | N             | 0 / 0                        | K                    |
| [sagasu](../entries/sagasu.yml)                                                                         | 探す                                         | 寻找                                                | absent        | 0 / 1                        | NC                   |
| [saishou-katto-ui](../entries/saishou-katto-ui.yml)                                                     | 最小カット                                   | 最小裁剪                                            | absent        | 0 / 0                        | NC                   |
| [sakana-general](../entries/sakana-general.yml)                                                         | 魚                                           | 鱼                                                  | absent        | 0 / 1                        | NC                   |
| [sakura](../entries/sakura.yml)                                                                         | 桜                                           | 樱花；樱树                                          | absent        | 0 / 0                        | NC-N                 |
| [sakura-daigaku](../entries/sakura-daigaku.yml)                                                         | さくら大学                                   | 大学名称；さくら大学                                | absent        | 0 / 0                        | D: D04               |
| [sakura-no-ki](../entries/sakura-no-ki.yml)                                                             | 桜の木                                       | 樱花树                                              | W             | 0 / 0                        | K                    |
| [san-name-suffix](../entries/san-name-suffix.yml)                                                       | ～さん                                       | 接在人名后的礼貌称呼                                | N             | 0 / 0                        | K                    |
| [seeraa-muun](../entries/seeraa-muun.yml)                                                               | セーラームーン                               | 水兵月；Sailor Moon                                 | absent        | 0 / 0                        | NC                   |
| [sei-kanji-reading](../entries/sei-kanji-reading.yml)                                                   | 成                                           | “成”的日语音读之一                                  | absent        | 0 / 0                        | NC-N                 |
| [sei-surname](../entries/sei-surname.yml)                                                               | 姓                                           | 姓氏                                                | absent        | 0 / 0                        | NC-N                 |
| [seigi](../entries/seigi.yml)                                                                           | 正義                                         | 正义                                                | absent        | 0 / 0                        | NC                   |
| [sen-kanji-reading](../entries/sen-kanji-reading.yml)                                                   | 銭                                           | 钱；“銭”的读音                                      | absent        | 0 / 0                        | NC                   |
| [sensei](../entries/sensei.yml)                                                                         | 先生                                         | 老师等称呼                                          | N             | 0 / 1                        | K                    |
| [shain](../entries/shain.yml)                                                                           | 社員                                         | 公司成员；职员                                      | N             | 0 / 1                        | K                    |
| [shinde-miru](../entries/shinde-miru.yml)                                                               | 死んでみる                                   | 试着死一次看看                                      | W             | 0 / 0                        | K                    |
| [shinjitsu](../entries/shinjitsu.yml)                                                                   | 真実                                         | 真相；真实                                          | absent        | 0 / 0                        | NC                   |
| [shinjitsu-wa-itsumo-hitotsu](../entries/shinjitsu-wa-itsumo-hitotsu.yml)                               | 真実はいつもひとつ！                         | 真相永远只有一个！                                  | absent        | 0 / 0                        | NC                   |
| [shinsekai](../entries/shinsekai.yml)                                                                   | 新世界                                       | 新世界                                              | W             | 0 / 0                        | K                    |
| [shinu](../entries/shinu.yml)                                                                           | 死ぬ                                         | 死                                                  | absent        | 0 / 0                        | NC-N                 |
| [shitsurei-desu-ga](../entries/shitsurei-desu-ga.yml)                                                   | 失礼ですが                                   | 冒昧问一下                                          | absent        | 0 / 0                        | NC-N                 |
| [shokudou](../entries/shokudou.yml)                                                                     | 食堂                                         | 食堂；餐馆                                          | absent        | 0 / 0                        | NC                   |
| [shouhin](../entries/shouhin.yml)                                                                       | 商品                                         | 商品；产品                                          | absent        | 0 / 0                        | NC                   |
| [sochira](../entries/sochira.yml)                                                                       | そちら                                       | 那边；贵方；你们那里                                | absent        | 0 / 0                        | NC-N                 |
| [sochira-no-byouin](../entries/sochira-no-byouin.yml)                                                   | そちらの病院                                 | 贵院；你们医院                                      | absent        | 0 / 0                        | NC-N                 |
| [sono-kata](../entries/sono-kata.yml)                                                                   | その方                                       | 那位；那个人（礼貌）                                | absent        | 0 / 0                        | NC-N                 |
| [sono-namida-wa-nan-da](../entries/sono-namida-wa-nan-da.yml)                                           | その涙は何だ？                               | 你那眼泪算什么？                                    | absent        | 0 / 0                        | NC-N                 |
| [sore](../entries/sore.yml)                                                                             | それ                                         | 那个；那件事                                        | absent        | 0 / 0                        | NC                   |
| [sore-ga-ore-no-nindou](../entries/sore-ga-ore-no-nindou.yml)                                           | それがオレの忍道だ!!                         | 这就是我的忍道！                                    | absent        | 0 / 0                        | NC                   |
| [sukueru](../entries/sukueru.yml)                                                                       | 救える                                       | 能够拯救                                            | W             | 0 / 1                        | K                    |
| [sukuu](../entries/sukuu.yml)                                                                           | 救う                                         | 拯救；救助                                          | absent        | 1 / 1                        | NC                   |
| [sumimasen](../entries/sumimasen.yml)                                                                   | すみません                                   | 对不起；劳驾；不好意思                              | absent        | 0 / 1                        | NC-N                 |
| [tai-country](../entries/tai-country.yml)                                                               | タイ                                         | 泰国                                                | absent        | 0 / 0                        | NC-N                 |
| [tai-grammar](../entries/tai-grammar.yml)                                                               | ～たい                                       | 想做……                                              | absent        | 0 / 0                        | R: R17               |
| [tantei](../entries/tantei.yml)                                                                         | 探偵                                         | 侦探                                                | absent        | 0 / 0                        | NC                   |
| [tantei-sa](../entries/tantei-sa.yml)                                                                   | 探偵さ！                                     | 是侦探！                                            | W             | 0 / 0                        | K                    |
| [taoru](../entries/taoru.yml)                                                                           | タオル                                       | 毛巾                                                | E             | 0 / 0                        | K                    |
| [te-miru-grammar](../entries/te-miru-grammar.yml)                                                       | 〜てみる                                     | 试着做……                                            | absent        | 0 / 0                        | A: WF14, N05; R: R17 |
| [tei-kanji-reading](../entries/tei-kanji-reading.yml)                                                   | 程                                           | “程”的音读                                          | absent        | 0 / 0                        | NC-N                 |
| [to-moushimasu-grammar](../entries/to-moushimasu-grammar.yml)                                           | ～と申します                                 | 我叫……；我是……                                      | W             | 0 / 0                        | K                    |
| [to-nanoru-grammar](../entries/to-nanoru-grammar.yml)                                                   | ～と名乗る                                   | 自称是……                                            | W             | 0 / 0                        | K                    |
| [toka-grammar](../entries/toka-grammar.yml)                                                             | ～とか                                       | ……之类；诸如……                                      | absent        | 0 / 0                        | NC-N                 |
| [tsuki](../entries/tsuki.yml)                                                                           | 月                                           | 月亮                                                | absent        | 0 / 0                        | NC-N                 |
| [tsuki-ni-kawatte-oshioki-yo](../entries/tsuki-ni-kawatte-oshioki-yo.yml)                               | 月にかわっておしおきよ！                     | 我要代表月亮惩罚你！                                | W             | 0 / 0                        | K                    |
| [tsukino-usagi](../entries/tsukino-usagi.yml)                                                           | 月野うさぎ                                   | 月野兔                                              | absent        | 0 / 0                        | D: D05               |
| [tsuzuki](../entries/tsuzuki.yml)                                                                       | 続き                                         | 后续；接下来的部分；延续。                          | absent        | 0 / 0                        | NC-N                 |
| [watashi](../entries/watashi.yml)                                                                       | 私                                           | 我                                                  | absent        | 0 / 0                        | R: R11               |
| [watashi-wa-gaikokujin-desu](../entries/watashi-wa-gaikokujin-desu.yml)                                 | 私は外国人です                               | 我是外国人                                          | absent        | 0 / 0                        | A: WF11              |
| [watashi-wa-sono-kata-dewa-arimasen](../entries/watashi-wa-sono-kata-dewa-arimasen.yml)                 | 私はその方ではありません。                   | 我不是您要找的那个人                                | absent        | 0 / 0                        | NC-N                 |
| [wazawaza](../entries/wazawaza.yml)                                                                     | わざわざ                                     | 特意；专程                                          | absent        | 0 / 0                        | NC                   |
| [yamashina-ku](../entries/yamashina-ku.yml)                                                             | 山科区                                       | 山科区                                              | absent        | 0 / 0                        | NC                   |
| [yareru](../entries/yareru.yml)                                                                         | やれる                                       | 能做；能够做到                                      | W             | 0 / 1                        | K                    |
| [yareru-koto](../entries/yareru-koto.yml)                                                               | やれること                                   | 能做到的事                                          | W             | 0 / 0                        | K                    |
| [yaritai-dake-da-yo](../entries/yaritai-dake-da-yo.yml)                                                 | やりたいだけだよ                             | 只是想做而已                                        | W             | 0 / 0                        | K                    |
| [yaru](../entries/yaru.yml)                                                                             | やる                                         | 做；干                                              | absent        | 1 / 1                        | NC-N                 |
| [yobu](../entries/yobu.yml)                                                                             | 呼ぶ                                         | 呼唤；叫喊；叫某人的名字。                          | absent        | 0 / 0                        | NC-N                 |
| [yomikata](../entries/yomikata.yml)                                                                     | 読み方                                       | 读法                                                | W             | 0 / 0                        | K                    |
| [yoroshiku-onegai-shimasu](../entries/yoroshiku-onegai-shimasu.yml)                                     | よろしくお願いします                         | 请多关照                                            | absent        | 0 / 0                        | NC-N                 |
| [yume](../entries/yume.yml)                                                                             | 夢                                           | 梦；梦想；希望实现的愿望。                          | absent        | 0 / 0                        | NC-N                 |
| [zettai](../entries/zettai.yml)                                                                         | 絶対                                         | 绝对；一定                                          | absent        | 0 / 0                        | NC-N                 |
| [zettai-daijoubu-da-yo](../entries/zettai-daijoubu-da-yo.yml)                                           | 絶対、だいじょうぶだよ。                     | 绝对没问题的；一定会没事的                          | absent        | 0 / 0                        | NC-N                 |
| [zettee-shinanee](../entries/zettee-shinanee.yml)                                                       | ぜってェ死なねー                             | 绝对不会死                                          | W             | 0 / 0                        | K                    |

## Implementation after approval

This procedure was followed after approval. The separately authorized 大学生 entry is recorded as the sole exception to the original no-new-entry boundary; the enrichment was verified independently before adding it.

1. Re-read this approved plan, current AGENTS/README/schema/order guides, source entries, implementation, and Git status. Record exactly which text and R codes are approved. Reconcile changed source text, ID set, counts, existing aids, pairs, and HEAD against this baseline; stop and resolve conflicting edits instead of overwriting them. Snapshot full entry bytes and existing media hashes, and preserve unrelated work.

2. Before edits, produce an isolated baseline export from the reconciled current sources for comparison. Record its ID set, row count, audio fields, StudyOrder ranks, and all 17 columns. Build again after enrichment; do not assume an old ignored dist file is current. Never hand-edit generated output.

3. Apply only approved exact aid values and Notes source appends. WF14 alone permits the explicitly shown Notes removal and replacement. Keep every existing aid, source/provenance line, stable ID, Word, Meaning, pronunciation, Context, Example, tag, audio reference, learning priority, prerequisite, and unrelated field unchanged unless a separate approved correction names it. Do not rewrite or reformat unrelated YAML.

4. For approved, unblocked R01–R25 codes only, append the stated target IDs on the stated endpoint, preserving list members. Declare each new pair once. R26 requires its separately authorized entry; that condition has now been fulfilled as recorded below. Pair approval changes no mnemonic or prerequisite. Text approval adds no association. Preserve all 89 baseline direct pairs. Apart from the separately authorized `daigakusei` entry, do not add entries, audio, schema fields, tooling, templates, dependencies, or new generated artifacts beyond the documented export. Memory-aid changes do not require audio regeneration.

5. Run these documented commands from the repository root, and inspect new text files separately because Git diff whitespace checking omits untracked files:

```bash
go test ./...
go vet ./...
go run scripts/validate/main.go
go run scripts/build/main.go
git diff --check
```

6. Build twice without source changes and compare export bytes. Verify the reconciled baseline ID set and row count (1,241 at planning), 17 columns on every parsed data row, exact approved aid values after HTML escaping, unchanged Audio values and (for the enrichment-only phase) StudyOrder ranks, and all baseline/new direct bidirectional RelatedWords displays. Adding the separately approved entry changes ranks; measure and report that effect independently. Partial approvals must use their actual subset counts. Check current Word/Meaning text, direct-neighbor sorting, cross-language labels, absence of transitive neighbors, and reciprocal deduplication without inventing extra links.

7. Check empty aid fields remain empty in export and inspect the unchanged conditional README template: absent fields must not leave headings or replay controls. This is static/export verification, not a claim that Anki Desktop or AnkiDroid rendered or spoke the content. No native rendering test is implied.

8. Compare changed entry fields and bytes with the approved allowlist. Aside from approved aids/Notes and pair declarations, all parsed values and unrelated source wording must match the snapshot. Compare existing media hashes exactly. Confirm no recordings or unapproved entries were added/removed, no unapproved source moved, and unrelated work was preserved. Report any partial failure without claiming a transaction or silently broadening approval.

9. Update this plan with actual approvals, implementation date, applied codes, measured counts, preservation checks, command results, and remaining failures. Distinguish static/export checks from human listening and live Anki import/rendering/playback, which require their own actual verification. Do not stage, commit, push, copy media to Anki, or modify an Anki collection.

## Planning verification

Planning checks are read-only except creation of this document. The following results are filled from the final plan audit; they do not stand in for implementation checks:

- Passed: 1,809 local file/heading links resolve under GitHub-style Markdown heading conventions. Source paths are relative to this file. No site-root paths or ambiguous renderer-specific anchors are used. All 30 supporting source URLs map to register entries and inspected pages/sections; remote targets are not guaranteed immutable.
- Passed: exactly 1,241 ledger rows with the complete baseline ID set and actual source paths; 238 retained-aid rows cover all 251 existing aid values; the retained-pair inventory contains exactly the baseline 89 unique pairs. The 29 text proposals target absent fields on 27 existing IDs. N07 covers the user-requested pronunciation contrast; R26 is blocked and not included in the executable-pair totals. All 25 relationship pairs exist as IDs, are unique, and are absent in both declaration directions. Counts reconcile: 25 declaration endpoints, 49 new display endpoints, 43 potential source-file changes, and 63 entries in the source-or-display union. If all pairs are approved, the complete graph has 114 pairs, 228 directed displays, and 214 unique display endpoints.
- Passed: untracked plan text checked separately for balanced fences, table column consistency, trailing whitespace, tabs, final newline, stable review codes, and source-register references. The illustrative UTF-16/UTF-8 offsets in N04 were independently calculated as 3 and 5.
- Preservation: all 1,241 original entry files, all 1,241 existing MP3s, implementation, templates, dependencies, and existing generated exports match their original hashes. Of 2,512 baseline files, 2,511 remain identical. The sole exception is the independently committed MAINTENANCE_PROMPTS.md change above; its original hash matches the parent commit, its current hash matches HEAD, and this task did not edit it. Ignored Finder metadata is outside the source/content snapshot and was not modified by this task.
- Follow-up on 2026-11-05: added N07 and blocked R26 at the user’s request; inspected S29/S30 and verified that no current stable ID covers 大学生. Changes remain planning-only; pair and exact wording approvals remain separate. The follow-up byte comparison confirmed every other source/content file, including the existing plan baseline corpus and recordings, stayed unchanged.
- Final working-tree scope: only this new plan is untracked. `git diff --check` passed for tracked changes; the separate untracked-text audit above covers this file. No other project file was created or edited by this task.

- Not run during planning: Go tests, vet, validation/build commands, deterministic export builds, MP3 decoding, audio generation, human pronunciation review, native Anki import/rendering/playback. No approval or implementation result is claimed.

## Implementation record

Completed: **2026-11-05**. Authorization: the user approved every proposal and explicitly approved creating 大学生 to complete R26, without generating audio. No remaining approved code is blocked. Deferred investigations D01–D06 remain intentionally unimplemented: they contain no approved field values and still lack the evidence described above.

### Applied scope

- **WF01–WF14, ETY01–ETY08, N01–N07**: all 29 exact approved aid values applied to 27 existing entries. Existing aid values were retained. The WF14 move removes only the approved first structure clause from Notes; the placeholder explanation and 地獄少女 provenance remain intact. Source lines were appended without duplicate URLs; N07 reuses the existing 外国人 headword reference.
- **R01–R26**: 26 new pair declarations across 26 existing source entries, preserving the original 89 pairs. Each new pair is declared once. R26 is declared on gaikokujin, targeting daigakusei. R23 remains declared on gaikoku; no reciprocal redeclaration was added.
- **New entry**: [daigakusei](../entries/daigakusei.yml), Japanese 大学生, reading だいがくせい, noun, meaning 大学生；在大学就读的学生, example `私は大学生です。`. Notes records S30. It uses priority 1 as practical introductory vocabulary, with existing daigaku and gakusei as prerequisites to introduce the useful foundations first. No separate component entry or new mnemonic was added for it. It has no audio field or recording, as explicitly requested.
- **Changed files**: 43 existing entry files, the one new entry, and this plan (45 source/documentation files total). The ignored vocabulary.tsv was regenerated exclusively through the build command. Source changes affect only approved aid fields, approved Notes updates, relationship declarations, and the separately authorized entry. No implementation, template, dependency, or audio file changed.

### Measured results

| Check                               | Result                                                                                              |
| :---------------------------------- | :-------------------------------------------------------------------------------------------------- |
| Collection                          | 1,242 entries: 999 English, 243 Japanese                                                            |
| Aid fields applied                  | 29: 14 word_formation, 8 etymology, 7 mnemonic                                                      |
| Existing aid values retained        | All 251; collection now contains 280 populated aid values                                           |
| Notes updates                       | 27 existing entries; one exact approved move, source appends, no provenance removed                 |
| New / total direct pairs            | 26 / 115                                                                                            |
| Derived directed neighbor displays  | 230 across 215 unique endpoints                                                                     |
| Export                              | 1,242 parsed data rows, exactly 17 columns each                                                     |
| Existing stable IDs / source fields | All 1,241 IDs preserved; unrelated parsed fields and top-level source byte spans match the snapshot |
| Enrichment-only StudyOrder          | All 1,241 baseline ranks unchanged before adding 大学生                                             |
| Final StudyOrder                    | New daigakusei rank 0167; 1,075 existing ranks shift by one after insertion; 166 remain unchanged   |
| Prerequisites                       | Every existing and new prerequisite precedes its dependent; no existing learning metadata changed   |
| Existing Audio fields / media       | All 1,241 references and MP3 hashes unchanged                                                       |
| Newly generated / reused audio      | 0 / 0                                                                                               |
| Missing audio declarations          | Only the new daigakusei entry intentionally has no audio                                            |
| Deterministic builds                | Two sequential final builds produced identical bytes                                                |

Final export SHA-256:

```plain
76937ec68ddac9542b2600af633954cde045dbd14d99bb4f2d04e450d95f5543
```

### Verification performed

- `go test ./...`: passed (existing Go tests, cached; no runtime code changed).
- `go vet ./...`: passed.
- `go run scripts/validate/main.go`: passed for 1,242 entries. The 14 pre-existing repeated-word warnings remain; the added entry introduces no new duplicate warning.
- `go run scripts/build/main.go`: passed. Two sequential final builds were byte-identical.
- `git diff --check`: passed. The new plan and new YAML were also checked separately for trailing whitespace, tabs, balanced fences/table formatting where applicable, and final newlines.
- Artifact/source audit: all 29 aid values equal their approved wording in YAML and in HTML-escaped TSV output. Every existing and approved pair matches the complete direct-neighbor graph in both directions, with deterministic ID sorting and the expected cross-language labels. No transitive neighbors are displayed. Empty aid values remain empty in export.
- Preservation audit: the 1,241 original entry files were compared against a pre-edit byte snapshot. Parsed changes match the approved allowlist, and unrelated top-level byte spans are identical. All original media hashes are unchanged. Implementation, README/templates, schema guides, dependencies, and other unrelated files match their pre-edit hashes. No existing entry or recording was removed.
- Static template check: the unchanged README conditionally wraps Etymology, WordFormation, Mnemonic, RelatedWords, Notes, Context, and Example/TTS. This is source/export verification only, not native rendering validation.
- Document check: local file links, heading anchors, stable-ID references, review-code coverage, and counts checked after the completion record was added. The original review ledger is retained as a dated baseline; the new entry is explicitly recorded here rather than retroactively counted as screened.

### Limits and next import

No human pronunciation review, MP3 decoding, native Anki import, rendering, or playback was performed for this enrichment. Existing audio did not change. No audio was generated for 大学生. No files were staged, committed, pushed, or copied into Anki, and no Anki collection was modified.

Reimport the complete rebuilt export to refresh both memory aids and reverse association displays. Because the separately approved entry changes StudyOrder ranks, use the existing [learning-order repositioning workflow](LEARNING_ORDER.md#apply-the-order-in-anki) for new cards if you want Anki to follow the updated order. Existing review history remains Anki’s responsibility.
