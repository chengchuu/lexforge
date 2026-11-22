# 书面读音全库审查与第九批音标计划

状态为已批准并实施。用户批准 B09-P01–B09-P08。下文保留 2026-11-22 规划基线和审阅文案。实施结果见末节。[第八批](EN_PRONUNCIATION_ENRICHMENT_BATCH_08_PLAN.md)保留实施记录。

## 全库盘点

以当前 YAML 的 `id` 和 `language` 为准。空值、空白文本及未提供 `pronunciation` 均计为缺失。已读 [AGENTS.md](../AGENTS.md)、[README.md](../README.md) 和[词条规范](ENTRY_SCHEMA.md#pronunciation)。盘点包含工作区已有修改。

| 语言 | 总数 | 已有读音 | 缺失 |
|:--|--:|--:|--:|
| `en` | 1,590 | 1,018 | 572 |
| `ja` | 243 | 243 | 0 |
| 合计 | 1,833 | 1,261 | 572 |

本轮对 100 个缺失 ID 逐项查证，提出 8 项（美式 2 项、英式 6 项），92 项查证后暂缓。另有 24 项沿用[第八批](EN_PRONUNCIATION_ENRICHMENT_BATCH_08_PLAN.md)或[王国之泪批次](TOTK_PRONUNCIATION_ENRICHMENT_PLAN.md)的逐项查证记录，本轮未重查。剩余 448 项待查证，其中包含此前只做过内容筛查、尚未逐名追查读音的游戏专名。四类合计 572 项。按现有基线，全部获批实施后缺失将降至 564 项。

### 已填读音的疑点

对英语现值检查斜线和口音前缀，对日语现值检查是否夹有罗马字母。这两类格式检查未发现异常。抽查同形词的不同转写时，保留词性导致的重音差异以及不同词典的记音方式，例如 `construct` 的动词和名词重音。现有字段没有拟改项。此次未逐项重新审定已填 IPA 或日语读音。

## 批次与证据

从日常词形和常用短语入手，再查技术词和金融术语。常用短语的词典命中多为用例、单个词的 IPA 或音频入口。本轮只列出有完整书面转写且可核对用法的条目。另分队列跟踪缩写、技术标识符、专名、符号与命令。

美式 IPA 无前缀。每个英式回退说明美式查证的限制。访问受阻只作为查证限制。Oxford 和 Cambridge 的快照包含词条或发音正文，使用其中的完整内容。会计词典涉及旧字体编码的 IPA，以对应印刷版页面为准，来源 URL 定位 PDF 页码。专名近似、合成语音以及拆开查到的单词读音均不足以支持完整值。

以下 Notes 来源行是逐字拟追加内容。整批批准仅覆盖这些编号中的 IPA 和来源行。

### 批次 A · 常用词形

| 审阅 ID | 稳定 ID / Word | 当前词性 | 当前 Meaning | 拟写 `pronunciation` | 适用性与证据 |
|:--|:--|:--|:--|:--|:--|
| B09-P01 | [corrupted-combat](../entries/corrupted-combat.yml) / corrupted | adjective | 被腐化的 | `UK /kəˈrʌptɪd/` | 英式教材列出完整形容词，涵盖变坏、受腐化义，可用于游戏中的转用。Cambridge 只转写原形，Wiktionary 的完整 IPA 未标口音，本轮未核实美式完整形式。 [Oxford University Press — Wordlist, Unit 7 (corrupted)](https://www.zsmladi.cz/upload/140764132e1c2545c338112b7bae726c883fc1de.pdf) |

### 批次 B · 开发基础词

| 审阅 ID | 稳定 ID / Word | 当前词性 | 当前 Meaning | 拟写 `pronunciation` | 适用性与证据 |
|:--|:--|:--|:--|:--|:--|
| B09-P02 | [viewport-css](../entries/viewport-css.yml) / viewport | noun | 视口，浏览器中供网页布局或显示的区域 | `/ˈvjuːpɔːrt/` | 词典正文收录同形计算机名词，采用并列两套 IPA 中的北美读法；Oxford Advanced American Dictionary 的同一计算机义另用其美式教学转写。网页直开返回 403，使用可读正文快照核对。 [Oxford Advanced Learner’s Dictionary — viewport](https://www.oxfordlearnersdictionaries.com/definition/english/viewport) |
| B09-P03 | [keyframe-css](../entries/keyframe-css.yml) / keyframe | noun | 关键帧，定义动画某一进度点样式的规则 | `/ˈkiˌfɹeɪm/` | 英语名词第一义为动画关键帧，适用于 CSS 动画。采用 General American 的完整词形。 [Wiktionary — keyframe (English pronunciation)](https://en.wiktionary.org/w/index.php?title=keyframe&oldid=87981835) |
| B09-P04 | [mutex-go](../entries/mutex-go.yml) / mutex | noun | 互斥锁，用于协调对共享资源的访问 | `UK /ˈmjuːtɛks/` | 同形计算机名词对应互斥锁，IPA 标为 Received Pronunciation。核对 Cambridge、Collins 和美式定向检索后未获得可靠美式 IPA。 [Wiktionary — mutex (English pronunciation)](https://en.wiktionary.org/w/index.php?title=mutex&oldid=88231891) |

### 批次 C · 财报与报价

| 审阅 ID | 稳定 ID / Word | 当前词性 | 当前 Meaning | 拟写 `pronunciation` | 适用性与证据 |
|:--|:--|:--|:--|:--|:--|
| B09-P05 | [net-assets-finance](../entries/net-assets-finance.yml) / net assets | noun | 净资产；资产扣除负债后的净值。 | `UK /ˌnet ˈæsets/` | 核对印刷版完整复数和从资产扣除负债的义项。Cambridge 可读正文有该义及英美音频入口，但无 IPA。采用英国出版原书的英式转写。 [A & C Black — Dictionary of Accounting, 4th ed., p. 149 (net assets)](https://castudynotes.com/wp-content/uploads/2022/03/14.Dictionary-of-Accounting.pdf#page=156) |
| B09-P06 | [net-income-finance](../entries/net-income-finance.yml) / net income | noun | 净利润；扣除相关成本、费用和税项后的利润。 | `UK /ˌnet ˈɪnkʌm/` | 核对印刷版税费等扣除后的收入义，整短语同一转写系统。Collins 其他分区的转写口音未确认，Cambridge 直开 403，未核实美式整条。 [A & C Black — Dictionary of Accounting, 4th ed., p. 149 (net income)](https://castudynotes.com/wp-content/uploads/2022/03/14.Dictionary-of-Accounting.pdf#page=156) |
| B09-P07 | [bid-price-finance](../entries/bid-price-finance.yml) / bid price | noun | 买方报价；行情中买方愿意支付的买入报价。 | `UK /ˈbɪd praɪs/` | 英国原书从持有人卖给出价方的角度解释，适用于买方报价这一金融术语。Collins 美式分区未给可用整短语 IPA，Cambridge 直开 403。 [A & C Black — Dictionary of Accounting, 4th ed., p. 25 (bid price)](https://castudynotes.com/wp-content/uploads/2022/03/14.Dictionary-of-Accounting.pdf#page=32) |
| B09-P08 | [ebit-finance](../entries/ebit-finance.yml) / EBIT | noun | 息税前利润；未扣除利息和税项的利润指标。 | `UK /ˈebɪt/` | Cambridge 在完整缩写发音正文仅标 UK。A & C Black 的会计词典另录英式 /ˈiːbɪt/。本轮未核实美式整词 IPA，保留 Cambridge 所列的一种读法。 [Cambridge Dictionary — EBIT (English pronunciation)](https://dictionary.cambridge.org/us/pronunciation/english/ebit) |

### 拟追加的 Notes 来源行

- B09-P01 `Pronunciation source: Oxford University Press — Wordlist, Unit 7 (corrupted). https://www.zsmladi.cz/upload/140764132e1c2545c338112b7bae726c883fc1de.pdf (accent: UK; accessed 2026-11-22).`
- B09-P02 `Pronunciation source: Oxford Advanced Learner’s Dictionary — viewport. https://www.oxfordlearnersdictionaries.com/definition/english/viewport (accent: US; accessed 2026-11-22).`
- B09-P03 `Pronunciation source: Wiktionary — keyframe (English pronunciation). https://en.wiktionary.org/w/index.php?title=keyframe&oldid=87981835 (accent: US; accessed 2026-11-22).`
- B09-P04 `Pronunciation source: Wiktionary — mutex (English pronunciation). https://en.wiktionary.org/w/index.php?title=mutex&oldid=88231891 (accent: UK; accessed 2026-11-22).`
- B09-P05 `Pronunciation source: A & C Black — Dictionary of Accounting, 4th ed., p. 149 (net assets). https://castudynotes.com/wp-content/uploads/2022/03/14.Dictionary-of-Accounting.pdf#page=156 (accent: UK; accessed 2026-11-22).`
- B09-P06 `Pronunciation source: A & C Black — Dictionary of Accounting, 4th ed., p. 149 (net income). https://castudynotes.com/wp-content/uploads/2022/03/14.Dictionary-of-Accounting.pdf#page=156 (accent: UK; accessed 2026-11-22).`
- B09-P07 `Pronunciation source: A & C Black — Dictionary of Accounting, 4th ed., p. 25 (bid price). https://castudynotes.com/wp-content/uploads/2022/03/14.Dictionary-of-Accounting.pdf#page=32 (accent: UK; accessed 2026-11-22).`
- B09-P08 `Pronunciation source: Cambridge Dictionary — EBIT (English pronunciation). https://dictionary.cambridge.org/us/pronunciation/english/ebit (accent: UK; accessed 2026-11-22).`

## 缺失 ID 清单

每个缺失稳定 ID 只列一次。队列表示下一轮的主要查证方向。待查证项保留当前拼写和用法，专名优先核实实际所指，标识符和符号优先寻找已出版的口读依据。

| 队列 | 缺失 | 拟补 | 本轮暂缓 | 前批暂缓 | 待查证 |
|:--|--:|--:|--:|--:|--:|
| 01 常用词与短语 | 14 | 1 | 12 | 0 | 1 |
| 02 技术、金融与游戏通用语 | 303 | 6 | 59 | 14 | 224 |
| 03 缩写及技术标识符 | 32 | 1 | 2 | 8 | 21 |
| 04 专名及游戏名称 | 211 | 0 | 19 | 2 | 190 |
| 05 语法模板、符号与命令 | 12 | 0 | 0 | 0 | 12 |

### 01 常用词与短语

| 稳定 ID | 当前 Word | 状态及依据 |
|:--|:--|:--|
| [arowana-general](../entries/arowana-general.yml) | arowana | 本轮查证后暂缓。Collins 页面不可用；Wiktionary 可读英语鱼名无 IPA，近似拼法 Finnish aravana 不适用。 |
| [cooking-pot-cooking](../entries/cooking-pot-cooking.yml) | cooking pot | 本轮查证后暂缓。所查结果为组件词音标、其他语言的词条或音频。尚无可靠标口音的英语整短语读法。 |
| [coordinates-ui](../entries/coordinates-ui.yml) | coordinates | 本轮查证后暂缓。Cambridge 音标为单数 coordinate。Collins 及 Dictionary.com 所转载的带音标定义是服装搭配义，坐标名词复数仍需整词证据。 |
| [corrupted-combat](../entries/corrupted-combat.yml) | corrupted | 拟补 B09-P01 |
| [cross-functional-collaboration-workflow](../entries/cross-functional-collaboration-workflow.yml) | cross-functional collaboration | 本轮查证后暂缓。Collins 页面未取得，未找到完整短语转写。 |
| [greaves-clothing](../entries/greaves-clothing.yml) | greaves | 本轮查证后暂缓。Collins 可读 IPA 属于油渣义，Wiktionary 的护胫复数 IPA 未标口音。原书单数不足以确认复数。 |
| [lex-name-element](../entries/lex-name-element.yml) | lex | 待查证 |
| [mine-cart-engineering](../entries/mine-cart-engineering.yml) | mine cart | 本轮查证后暂缓。Collins、Cambridge 未取得对应页面，Wiktionary 搜索命中连写 minecart，未给适用全词组转写。 |
| [opposite-shore-terrain](../entries/opposite-shore-terrain.yml) | opposite shore | 本轮查证后暂缓。Wiktionary 页面未取得，尚无整短语证据。 |
| [project-onboarding-workflow](../entries/project-onboarding-workflow.yml) | project onboarding | 本轮查证后暂缓。Collins 页面未取得，组件词 onboarding 的读法不能代替整个短语。 |
| [ranged-combat](../entries/ranged-combat.yml) | ranged | 本轮查证后暂缓。Collins 音标分区是建筑义。Wiktionary 未找到远程战斗形容词的标口音转写，ranged weapon 的词组不能代替单独 ranged。 |
| [related-words-language](../entries/related-words-language.yml) | related words | 本轮查证后暂缓。Collins 页面未取得，其他词典检索多为页面标签或词组用例，尚无完整读音。 |
| [unfuse-general](../entries/unfuse-general.yml) | unfuse | 本轮查证后暂缓。Collins 页面未取得。Wiktionary 可见词页未给英语动词的完整音标。 |
| [waistguard-clothing](../entries/waistguard-clothing.yml) | waistguard | 本轮查证后暂缓。Collins 和 Wiktionary 的确切词形页面未取得；未组合组件读法。 |

### 02 技术、金融与游戏通用语

| 稳定 ID | 当前 Word | 状态及依据 |
|:--|:--|:--|
| [absolute-positioning-css](../entries/absolute-positioning-css.yml) | absolute positioning | 待查证 |
| [abstract-syntax-tree-compiler](../entries/abstract-syntax-tree-compiler.yml) | abstract syntax tree | 待查证 |
| [access-log-nginx](../entries/access-log-nginx.yml) | access log | 本轮查证后暂缓。Wiktionary 和搜得的词典资料未能给出有口音标签的全词组 IPA。 |
| [accessible-name-html](../entries/accessible-name-html.yml) | accessible name | 待查证 |
| [accessor-tech](../entries/accessor-tech.yml) | accessor | 前批查证后暂缓，本轮未重查。沿用第八批清单。本轮查证后暂缓。补查 Wiktionary 未见英语 IPA。Collins 的访问限制见第七批。 |
| [alternative-text-html](../entries/alternative-text-html.yml) | alternative text | 本轮查证后暂缓。Wiktionary 与 Cambridge 主要命中缩写 alt text，未取得当前完整词组的 IPA。 |
| [annual-percentage-yield-finance](../entries/annual-percentage-yield-finance.yml) | annual percentage yield | 本轮查证后暂缓。Cambridge 结果将此词列在 annual percentage rate 的正文用例中，后者的读音不能代用。 |
| [array-hole-javascript](../entries/array-hole-javascript.yml) | array hole | 待查证 |
| [ask-price-finance](../entries/ask-price-finance.yml) | ask price | 本轮查证后暂缓。Cambridge 直开 403，会计词典未找到同一词头。其他报价术语的读法未作为替代。 |
| [assignability-typescript](../entries/assignability-typescript.yml) | assignability | 本轮查证后暂缓。Wiktionary 可读页面未列整词 IPA，仍需完整派生词读法依据。 |
| [at-rule-css](../entries/at-rule-css.yml) | at-rule | 本轮查证后暂缓。Wiktionary 有术语义但未见 IPA。 |
| [attack-power-combat](../entries/attack-power-combat.yml) | attack power | 本轮查证后暂缓。Collins、Cambridge 未取得匹配的整短语发音，检索结果只支持单词或用例。 |
| [attribute-selector-css](../entries/attribute-selector-css.yml) | attribute selector | 本轮查证后暂缓。Collins 与 Wiktionary 未取得准确短语页面。组件读音不足以证实全词组。 |
| [auto-save-game](../entries/auto-save-game.yml) | auto-save | 前批查证后暂缓，本轮未重查。沿用第八批清单。本轮查证后暂缓。Cambridge 定向查询未取得当前连字符形与游戏名词义的 IPA。 |
| [backend-server-nginx](../entries/backend-server-nginx.yml) | backend server | 待查证 |
| [backing-array-go](../entries/backing-array-go.yml) | backing array | 待查证 |
| [base-image-docker](../entries/base-image-docker.yml) | base image | 待查证 |
| [batching-react](../entries/batching-react.yml) | batching | 前批查证后暂缓，本轮未重查。沿用第八批清单。第七批查证后暂缓，本轮未重查。Collins 未取得 batching 独立分区或完整 batching IPA。 |
| [beat-estimates-finance](../entries/beat-estimates-finance.yml) | beat estimates | 待查证 |
| [bid-price-finance](../entries/bid-price-finance.yml) | bid price | 拟补 B09-P07 |
| [binary-search-algorithm](../entries/binary-search-algorithm.yml) | binary search | 前批查证后暂缓，本轮未重查。沿用第八批清单。本轮查证后暂缓。Collins 未取得；Wiktionary 可读词条未见整短语 IPA。 |
| [bind-mount-docker](../entries/bind-mount-docker.yml) | bind mount | 待查证 |
| [block-formatting-context-css](../entries/block-formatting-context-css.yml) | block formatting context | 待查证 |
| [block-level-box-css](../entries/block-level-box-css.yml) | block-level box | 待查证 |
| [boolean-attribute-html](../entries/boolean-attribute-html.yml) | boolean attribute | 待查证 |
| [box-model-css](../entries/box-model-css.yml) | box model | 待查证 |
| [buffered-channel-go](../entries/buffered-channel-go.yml) | buffered channel | 待查证 |
| [build-cache-docker](../entries/build-cache-docker.yml) | build cache | 待查证 |
| [build-context-docker](../entries/build-context-docker.yml) | build context | 待查证 |
| [button-group-bootstrap](../entries/button-group-bootstrap.yml) | button group | 待查证 |
| [cache-invalidation-tech](../entries/cache-invalidation-tech.yml) | cache invalidation | 前批查证后暂缓，本轮未重查。沿用第八批清单。本轮查证后暂缓。Cambridge 定向查询未取得匹配的整短语音标。 |
| [cache-revalidation-web](../entries/cache-revalidation-web.yml) | cache revalidation | 待查证 |
| [call-stack-javascript](../entries/call-stack-javascript.yml) | call stack | 本轮查证后暂缓。Wiktionary 可读术语页未提供整短语 IPA。 |
| [cascade-layer-css](../entries/cascade-layer-css.yml) | cascade layer | 待查证 |
| [cash-sweep-finance](../entries/cash-sweep-finance.yml) | cash sweep | 本轮查证后暂缓。Cambridge 与 Collins 词页未取得，暂无整短语 IPA。 |
| [character-encoding-html](../entries/character-encoding-html.yml) | character encoding | 本轮查证后暂缓。Collins、Cambridge 页面未取得。Wiktionary 命中组件与其他语种释义。 |
| [character-reference-html](../entries/character-reference-html.yml) | character reference | 待查证 |
| [class-selector-css](../entries/class-selector-css.yml) | class selector | 待查证 |
| [cold-resistance-combat](../entries/cold-resistance-combat.yml) | cold resistance | 本轮查证后暂缓。检索命中其他语言或单词读法，尚无英语整短语 IPA。 |
| [color-mode-bootstrap](../entries/color-mode-bootstrap.yml) | color mode | 待查证 |
| [column-offset-bootstrap](../entries/column-offset-bootstrap.yml) | column offset | 待查证 |
| [column-span-html](../entries/column-span-html.yml) | column span | 待查证 |
| [combinator-css](../entries/combinator-css.yml) | combinator | 前批查证后暂缓，本轮未重查。沿用第八批清单。本轮查证后暂缓。Wiktionary 对英语整词有标口音音标，但其义项是 λ 演算，尚未找到与 CSS 选择器义匹配的发音证据。 |
| [commit-phase-react](../entries/commit-phase-react.yml) | commit phase | 待查证 |
| [component-identity-react](../entries/component-identity-react.yml) | component identity | 待查证 |
| [conditional-rendering-react](../entries/conditional-rendering-react.yml) | conditional rendering | 待查证 |
| [conditional-type-typescript](../entries/conditional-type-typescript.yml) | conditional type | 待查证 |
| [configuration-test-nginx](../entries/configuration-test-nginx.yml) | configuration test | 待查证 |
| [constraint-validation-html](../entries/constraint-validation-html.yml) | constraint validation | 待查证 |
| [container-query-css](../entries/container-query-css.yml) | container query | 待查证 |
| [containing-block-css](../entries/containing-block-css.yml) | containing block | 待查证 |
| [content-box-css](../entries/content-box-css.yml) | content box | 待查证 |
| [content-model-html](../entries/content-model-html.yml) | content model | 待查证 |
| [context-provider-react](../entries/context-provider-react.yml) | context provider | 待查证 |
| [context-window-ai](../entries/context-window-ai.yml) | context window | 前批查证后暂缓，本轮未重查。沿用第八批清单。本轮查证后暂缓。Cambridge 定向查询未取得匹配的整短语音标。 |
| [controlled-component-react](../entries/controlled-component-react.yml) | controlled component | 待查证 |
| [critical-hit-combat](../entries/critical-hit-combat.yml) | critical hit | 本轮查证后暂缓。Wiktionary 收录游戏义，未显示完整 IPA；Collins 页面未取得。 |
| [cross-axis-css](../entries/cross-axis-css.yml) | cross axis | 待查证 |
| [cross-origin-resource-sharing-web](../entries/cross-origin-resource-sharing-web.yml) | cross-origin resource sharing | 待查证 |
| [custom-hook-react](../entries/custom-hook-react.yml) | custom Hook | 待查证 |
| [custom-property-css](../entries/custom-property-css.yml) | custom property | 待查证 |
| [data-attribute-bootstrap](../entries/data-attribute-bootstrap.yml) | data attribute | 待查证 |
| [data-cell-html](../entries/data-cell-html.yml) | data cell | 待查证 |
| [data-race-go](../entries/data-race-go.yml) | data race | 本轮查证后暂缓。Collins 页面未取得；Wiktionary 搜索仅命中成分词和术语列表。 |
| [data-table-html](../entries/data-table-html.yml) | data table | 待查证 |
| [debounce-javascript](../entries/debounce-javascript.yml) | debounce | 本轮查证后暂缓。Wiktionary 英式 IPA 所在用法为动词。当前是表防抖方式的名词。 |
| [declaration-file-typescript](../entries/declaration-file-typescript.yml) | declaration file | 待查证 |
| [decorative-image-html](../entries/decorative-image-html.yml) | decorative image | 待查证 |
| [deep-copy-javascript](../entries/deep-copy-javascript.yml) | deep copy | 本轮查证后暂缓。Wiktionary 词义匹配但未见整短语 IPA，Collins 页面未取得。 |
| [dependency-array-react](../entries/dependency-array-react.yml) | dependency array | 待查证 |
| [derived-state-react](../entries/derived-state-react.yml) | derived state | 待查证 |
| [detached-mode-docker](../entries/detached-mode-docker.yml) | detached mode | 待查证 |
| [disabled-control-html](../entries/disabled-control-html.yml) | disabled control | 待查证 |
| [discriminated-union-typescript](../entries/discriminated-union-typescript.yml) | discriminated union | 待查证 |
| [display-utility-bootstrap](../entries/display-utility-bootstrap.yml) | display utility | 待查证 |
| [dockerfile-docker](../entries/dockerfile-docker.yml) | Dockerfile | 本轮查证后暂缓。Docker 官方文档确认拼写及概念，未给 IPA。待找同一术语发音资料。 |
| [doctype-html](../entries/doctype-html.yml) | doctype | 本轮查证后暂缓。Wiktionary 可读页面未列整词 IPA。相似拼写或展开形式不作为替代。 |
| [document-title-html](../entries/document-title-html.yml) | document title | 待查证 |
| [draconification-zelda](../entries/draconification-zelda.yml) | draconification | 待查证 |
| [dropdown-bootstrap](../entries/dropdown-bootstrap.yml) | dropdown | 本轮查证后暂缓。Wiktionary 可读页面无英语整词 IPA，旧 Collins 访问限制未解决。 |
| [earnings-per-share-finance](../entries/earnings-per-share-finance.yml) | earnings per share | 本轮查证后暂缓。Cambridge 同义词条有音频入口无 IPA；A & C Black 可读整词条却定义为以股价为基准的股息，与每股收益的卡片义不符。 |
| [effect-cleanup-react](../entries/effect-cleanup-react.yml) | Effect cleanup | 待查证 |
| [end-tag-html](../entries/end-tag-html.yml) | end tag | 待查证 |
| [error-boundary-react](../entries/error-boundary-react.yml) | error boundary | 待查证 |
| [error-log-nginx](../entries/error-log-nginx.yml) | error log | 本轮查证后暂缓。词典检索命中其他术语如 log file；Wiktionary 精确条目未取得。 |
| [error-wrapping-go](../entries/error-wrapping-go.yml) | error wrapping | 待查证 |
| [event-bubbling-dom](../entries/event-bubbling-dom.yml) | event bubbling | 待查证 |
| [event-capturing-dom](../entries/event-capturing-dom.yml) | event capturing | 待查证 |
| [event-delegation-dom](../entries/event-delegation-dom.yml) | event delegation | 待查证 |
| [event-handler-react](../entries/event-handler-react.yml) | event handler | 本轮查证后暂缓。Collins 页面未取得。Wiktionary 检索命中 listener 等相关词，不能替代整短语 IPA。 |
| [event-loop-javascript](../entries/event-loop-javascript.yml) | event loop | 本轮查证后暂缓。Wiktionary 有计算机义及 /ɪˈvɛnt luːp/，但未标口音。Collins、Cambridge 页面未取得。 |
| [exact-match-nginx](../entries/exact-match-nginx.yml) | exact match | 本轮查证后暂缓。Collins 与 Cambridge 返回词组或用例，未取得整短语 IPA。 |
| [exchange-traded-fund-finance](../entries/exchange-traded-fund-finance.yml) | exchange-traded fund | 本轮查证后暂缓。Cambridge 搜索命中 ETF 缩写，缺当前完整词头音标。 |
| [exhaustiveness-checking-typescript](../entries/exhaustiveness-checking-typescript.yml) | exhaustiveness checking | 待查证 |
| [expected-move-finance](../entries/expected-move-finance.yml) | expected move | 待查证 |
| [fall-damage-game](../entries/fall-damage-game.yml) | fall damage | 前批查证后暂缓，本轮未重查。沿用第八批清单。本轮查证后暂缓。Cambridge 定向查询未取得确切游戏义整短语 IPA。 |
| [fallback-value-css](../entries/fallback-value-css.yml) | fallback value | 待查证 |
| [fear-and-greed-index-finance](../entries/fear-and-greed-index-finance.yml) | fear and greed index | 待查证 |
| [fixed-positioning-css](../entries/fixed-positioning-css.yml) | fixed positioning | 待查证 |
| [flex-container-css](../entries/flex-container-css.yml) | flex container | 待查证 |
| [flex-item-css](../entries/flex-item-css.yml) | flex item | 待查证 |
| [flexbox-css](../entries/flexbox-css.yml) | flexbox | 本轮查证后暂缓。Wiktionary 有相关术语，未取得可核实的整词音标。 |
| [floating-point-precision-javascript](../entries/floating-point-precision-javascript.yml) | floating-point precision | 本轮查证后暂缓。Cambridge floating-point operations per second 是另一词组，Collins 的对应整条未取得。 |
| [fluid-container-bootstrap](../entries/fluid-container-bootstrap.yml) | fluid container | 待查证 |
| [focus-indicator-css](../entries/focus-indicator-css.yml) | focus indicator | 待查证 |
| [font-stack-css](../entries/font-stack-css.yml) | font stack | 待查证 |
| [form-control-bootstrap](../entries/form-control-bootstrap.yml) | form control | 待查证 |
| [form-submission-html](../entries/form-submission-html.yml) | form submission | 待查证 |
| [fragment-identifier-html](../entries/fragment-identifier-html.yml) | fragment identifier | 待查证 |
| [function-component-react](../entries/function-component-react.yml) | function component | 待查证 |
| [function-signature-typescript](../entries/function-signature-typescript.yml) | function signature | 待查证 |
| [funds-on-hold-finance](../entries/funds-on-hold-finance.yml) | funds on hold | 待查证 |
| [fuzz-testing-go](../entries/fuzz-testing-go.yml) | fuzz testing | 待查证 |
| [garbage-collection-javascript](../entries/garbage-collection-javascript.yml) | garbage collection | 前批查证后暂缓，本轮未重查。沿用第八批清单。本轮查证后暂缓。Collins/Wiktionary 有计算机义，未见整短语 IPA；Cambridge 返回 403。 |
| [generic-constraint-typescript](../entries/generic-constraint-typescript.yml) | generic constraint | 待查证 |
| [global-attribute-html](../entries/global-attribute-html.yml) | global attribute | 待查证 |
| [global-augmentation-typescript](../entries/global-augmentation-typescript.yml) | global augmentation | 待查证 |
| [go-module](../entries/go-module.yml) | Go module | 待查证 |
| [goroutine-go](../entries/goroutine-go.yml) | goroutine | 本轮查证后暂缓。Wiktionary 可读页面未列整词 IPA，尚缺 Go 术语读法证据。 |
| [graceful-degradation-web](../entries/graceful-degradation-web.yml) | graceful degradation | 本轮查证后暂缓。Collins 对应页面未取得，Wiktionary 命中 degradation 的单词音标。 |
| [grid-area-css](../entries/grid-area-css.yml) | grid area | 待查证 |
| [grid-container-css](../entries/grid-container-css.yml) | grid container | 待查证 |
| [grid-layout-css](../entries/grid-layout-css.yml) | grid layout | 待查证 |
| [grid-system-bootstrap](../entries/grid-system-bootstrap.yml) | grid system | 待查证 |
| [grid-track-css](../entries/grid-track-css.yml) | grid track | 待查证 |
| [happens-before-go](../entries/happens-before-go.yml) | happens-before | 待查证 |
| [header-cell-html](../entries/header-cell-html.yml) | header cell | 待查证 |
| [heading-level-html](../entries/heading-level-html.yml) | heading level | 待查证 |
| [heap-memory-javascript](../entries/heap-memory-javascript.yml) | heap memory | 待查证 |
| [heat-resistance-combat](../entries/heat-resistance-combat.yml) | heat resistance | 本轮查证后暂缓。检索命中相关义及其他语言，不含英语整短语 IPA。前批词典访问限制仍适用。 |
| [higher-order-function-javascript](../entries/higher-order-function-javascript.yml) | higher-order function | 待查证 |
| [hoisting-javascript](../entries/hoisting-javascript.yml) | hoisting | 本轮查证后暂缓。Wiktionary 所列整词 IPA 未标口音，无 JavaScript 专义转写。 |
| [image-digest-docker](../entries/image-digest-docker.yml) | image digest | 待查证 |
| [image-layer-docker](../entries/image-layer-docker.yml) | image layer | 待查证 |
| [image-repository-docker](../entries/image-repository-docker.yml) | image repository | 待查证 |
| [image-tag-docker](../entries/image-tag-docker.yml) | image tag | 待查证 |
| [implicit-implementation-go](../entries/implicit-implementation-go.yml) | implicit implementation | 待查证 |
| [indexed-access-type-typescript](../entries/indexed-access-type-typescript.yml) | indexed access type | 待查证 |
| [initializer-tech](../entries/initializer-tech.yml) | initializer | 本轮查证后暂缓。Wiktionary 可读页面未给整词 IPA；短语或其他词形不能代用。 |
| [inline-box-css](../entries/inline-box-css.yml) | inline box | 待查证 |
| [inline-frame-html](../entries/inline-frame-html.yml) | inline frame | 待查证 |
| [input-group-bootstrap](../entries/input-group-bootstrap.yml) | input group | 待查证 |
| [input-type-html](../entries/input-type-html.yml) | input type | 待查证 |
| [interactive-content-html](../entries/interactive-content-html.yml) | interactive content | 待查证 |
| [internal-redirect-nginx](../entries/internal-redirect-nginx.yml) | internal redirect | 待查证 |
| [intersection-type-typescript](../entries/intersection-type-typescript.yml) | intersection type | 待查证 |
| [invalid-state-bootstrap](../entries/invalid-state-bootstrap.yml) | invalid state | 待查证 |
| [iterable-tech](../entries/iterable-tech.yml) | iterable | 前批查证后暂缓，本轮未重查。沿用第八批清单。本轮查证后暂缓。Wiktionary 可见的 IPA 属于西班牙语；Cambridge 检索未获得可用英语正文。 |
| [javascript-plugin-bootstrap](../entries/javascript-plugin-bootstrap.yml) | JavaScript plugin | 待查证 |
| [keyframe-css](../entries/keyframe-css.yml) | keyframe | 拟补 B09-P03 |
| [lazy-loading-html](../entries/lazy-loading-html.yml) | lazy loading | 本轮查证后暂缓。Collins 对应页面未取得，Wiktionary lazy load 是不同词形。 |
| [least-recently-used-cache](../entries/least-recently-used-cache.yml) | least recently used | 待查证 |
| [lifting-state-up-react](../entries/lifting-state-up-react.yml) | lifting state up | 待查证 |
| [line-height-css](../entries/line-height-css.yml) | line height | 待查证 |
| [linked-list-merge-algorithm](../entries/linked-list-merge-algorithm.yml) | linked list merge | 待查证 |
| [linked-list-tech](../entries/linked-list-tech.yml) | linked list | 前批查证后暂缓，本轮未重查。沿用第八批清单。本轮查证后暂缓。Collins/Wiktionary 有计算机义，未见整短语 IPA。 |
| [list-group-bootstrap](../entries/list-group-bootstrap.yml) | list group | 待查证 |
| [list-item-html](../entries/list-item-html.yml) | list item | 待查证 |
| [list-rendering-react](../entries/list-rendering-react.yml) | list rendering | 待查证 |
| [literal-type-typescript](../entries/literal-type-typescript.yml) | literal type | 待查证 |
| [load-balancing-nginx](../entries/load-balancing-nginx.yml) | load balancing | 本轮查证后暂缓。Collins 有同义释义，未发现整短语 IPA。Wiktionary 和 Cambridge 页面未取得。 |
| [longest-prefix-nginx](../entries/longest-prefix-nginx.yml) | longest prefix | 待查证 |
| [longhand-property-css](../entries/longhand-property-css.yml) | longhand property | 待查证 |
| [main-axis-css](../entries/main-axis-css.yml) | main axis | 待查证 |
| [main-quest-game](../entries/main-quest-game.yml) | main quest | 本轮查证后暂缓。Collins 与 Wiktionary 的精确短语页面未取得。搜索到的 sidequest 属于另一词形及义项。 |
| [male-to-male-hardware](../entries/male-to-male-hardware.yml) | male-to-male | 待查证 |
| [manual-save-game](../entries/manual-save-game.yml) | manual save | 待查证 |
| [mapped-type-typescript](../entries/mapped-type-typescript.yml) | mapped type | 待查证 |
| [margin-account-finance](../entries/margin-account-finance.yml) | margin account | 本轮查证后暂缓。Cambridge 直开 403，会计词典未找到精确词头，旧 Collins 有释义无音标。 |
| [margin-collapsing-css](../entries/margin-collapsing-css.yml) | margin collapsing | 待查证 |
| [market-capitalization-finance](../entries/market-capitalization-finance.yml) | market capitalization | 本轮查证后暂缓。Cambridge 释义和发音正文的口音标签下只有音频入口，无整短语 IPA。 |
| [media-query-css](../entries/media-query-css.yml) | media query | 待查证 |
| [method-set-go](../entries/method-set-go.yml) | method set | 待查证 |
| [microtask-javascript](../entries/microtask-javascript.yml) | microtask | 本轮查证后暂缓。Wiktionary 英语词页未列整词 IPA。自动转写或其他技术词不适用。 |
| [module-path-go](../entries/module-path-go.yml) | module path | 待查证 |
| [module-specifier-tech](../entries/module-specifier-tech.yml) | module specifier | 待查证 |
| [mount-point-docker](../entries/mount-point-docker.yml) | mount point | 本轮查证后暂缓。Collins 页面未取得，检索只得 mount 的 IPA 或其他语种词条。 |
| [multiple-return-values-go](../entries/multiple-return-values-go.yml) | multiple return values | 待查证 |
| [mutex-go](../entries/mutex-go.yml) | mutex | 拟补 B09-P04 |
| [named-result-parameter-go](../entries/named-result-parameter-go.yml) | named result parameter | 待查证 |
| [namespace-typescript](../entries/namespace-typescript.yml) | namespace | 本轮查证后暂缓。Wiktionary 的英语 IPA 未标口音；明标巴西口音的是葡萄牙语。Collins 页面未取得。 |
| [navbar-bootstrap](../entries/navbar-bootstrap.yml) | navbar | 本轮查证后暂缓。Wiktionary 与 Collins 新词提议检查未取得可用整词 IPA。 |
| [net-assets-finance](../entries/net-assets-finance.yml) | net assets | 拟补 B09-P05 |
| [net-income-finance](../entries/net-income-finance.yml) | net income | 拟补 B09-P06 |
| [normal-flow-css](../entries/normal-flow-css.yml) | normal flow | 待查证 |
| [offcanvas-bootstrap](../entries/offcanvas-bootstrap.yml) | offcanvas | 本轮查证后暂缓。未取得适用整词 IPA，分写 off canvas 或不同连字符形不能直接代用。 |
| [one-way-data-flow-frontend](../entries/one-way-data-flow-frontend.yml) | one-way data flow | 待查证 |
| [onion-model-koa](../entries/onion-model-koa.yml) | onion model | 待查证 |
| [option-chain-finance](../entries/option-chain-finance.yml) | option chain | 待查证 |
| [optional-parameter-typescript](../entries/optional-parameter-typescript.yml) | optional parameter | 待查证 |
| [optional-property-typescript](../entries/optional-property-typescript.yml) | optional property | 待查证 |
| [ordered-list-html](../entries/ordered-list-html.yml) | ordered list | 本轮查证后暂缓。Cambridge 有词组用例，Collins 页面未取得，尚无整个词组的 IPA。 |
| [own-property-javascript](../entries/own-property-javascript.yml) | own property | 待查证 |
| [partial-clone-git](../entries/partial-clone-git.yml) | partial clone | 待查证 |
| [persistent-data-docker](../entries/persistent-data-docker.yml) | persistent data | 待查证 |
| [pointer-receiver-go](../entries/pointer-receiver-go.yml) | pointer receiver | 待查证 |
| [popover-bootstrap](../entries/popover-bootstrap.yml) | popover | 本轮查证后暂缓。Wiktionary 未见 UI 义的 IPA，Collins 既有命中为食品义。 |
| [port-publishing-docker](../entries/port-publishing-docker.yml) | port publishing | 待查证 |
| [postorder-traversal-tree](../entries/postorder-traversal-tree.yml) | postorder traversal | 待查证 |
| [pre-market-finance](../entries/pre-market-finance.yml) | pre-market | 本轮查证后暂缓。Cambridge 返回 403。本轮没有取得带连字符形容词、盘前义的转写。 |
| [prefix-match-nginx](../entries/prefix-match-nginx.yml) | prefix match | 待查证 |
| [preflight-request-web](../entries/preflight-request-web.yml) | preflight request | 待查证 |
| [profit-and-loss-finance](../entries/profit-and-loss-finance.yml) | profit and loss | 本轮查证后暂缓。Collins 找到金融释义但未见口音标明的整短语 IPA。Cambridge 的 profit-and-loss account 是不同词头。 |
| [progress-bar-bootstrap](../entries/progress-bar-bootstrap.yml) | progress bar | 本轮查证后暂缓。Wiktionary 有计算机义但未见标口音的整短语 IPA。Collins、Cambridge 未取得页面。 |
| [progressive-enhancement-web](../entries/progressive-enhancement-web.yml) | progressive enhancement | 本轮查证后暂缓。Collins 对应页面未取得，Wiktionary 命中 enhancement 的单词音标。 |
| [property-descriptor-javascript](../entries/property-descriptor-javascript.yml) | property descriptor | 待查证 |
| [property-value-css](../entries/property-value-css.yml) | property value | 待查证 |
| [props-react](../entries/props-react.yml) | props | 前批查证后暂缓，本轮未重查。沿用第八批清单。本轮查证后暂缓。Collins 命中其他义项；Wiktionary 给出单数 prop 或其他语言的音标。未据此补 React 义。 |
| [prototype-chain-javascript](../entries/prototype-chain-javascript.yml) | prototype chain | 待查证 |
| [prototype-setter-javascript](../entries/prototype-setter-javascript.yml) | prototype setter | 待查证 |
| [pseudo-class-css](../entries/pseudo-class-css.yml) | pseudo-class | 本轮查证后暂缓。Wiktionary 页面未取得，仍需确切词形和义项的证据。 |
| [pseudo-element-css](../entries/pseudo-element-css.yml) | pseudo-element | 本轮查证后暂缓。Wiktionary 页面未取得，仍需确切词形和义项的证据。 |
| [pure-component-react](../entries/pure-component-react.yml) | pure component | 待查证 |
| [race-detector-go](../entries/race-detector-go.yml) | race detector | 待查证 |
| [reactive-dependency-react](../entries/reactive-dependency-react.yml) | reactive dependency | 待查证 |
| [read-only-control-html](../entries/read-only-control-html.yml) | read-only control | 待查证 |
| [reduced-motion-css](../entries/reduced-motion-css.yml) | reduced motion | 待查证 |
| [ref-react](../entries/ref-react.yml) | ref | 待查证 |
| [regex-match-nginx](../entries/regex-match-nginx.yml) | regex match | 待查证 |
| [relative-positioning-css](../entries/relative-positioning-css.yml) | relative positioning | 待查证 |
| [relative-unit-css](../entries/relative-unit-css.yml) | relative unit | 待查证 |
| [required-field-html](../entries/required-field-html.yml) | required field | 待查证 |
| [rerender-react](../entries/rerender-react.yml) | re-render | 前批查证后暂缓，本轮未重查。沿用第八批清单。本轮查证后暂缓。检索命中自动生成转写，未核实带连字符的名词义词典发音。 |
| [responsive-container-bootstrap](../entries/responsive-container-bootstrap.yml) | responsive container | 待查证 |
| [responsive-design-css](../entries/responsive-design-css.yml) | responsive design | 待查证 |
| [responsive-image-html](../entries/responsive-image-html.yml) | responsive image | 待查证 |
| [responsive-utility-bootstrap](../entries/responsive-utility-bootstrap.yml) | responsive utility | 待查证 |
| [rest-parameter-typescript](../entries/rest-parameter-typescript.yml) | rest parameter | 待查证 |
| [return-type-tech](../entries/return-type-tech.yml) | return type | 待查证 |
| [reverse-proxy-nginx](../entries/reverse-proxy-nginx.yml) | reverse proxy | 本轮查证后暂缓。Wiktionary 有相应术语义但未显示整短语 IPA，Collins 页面未取得。 |
| [row-span-html](../entries/row-span-html.yml) | row span | 待查证 |
| [rules-of-hooks-react](../entries/rules-of-hooks-react.yml) | Rules of Hooks | 待查证 |
| [same-origin-policy-web](../entries/same-origin-policy-web.yml) | same-origin policy | 待查证 |
| [sass-mixin-bootstrap](../entries/sass-mixin-bootstrap.yml) | Sass mixin | 待查证 |
| [sass-variable-bootstrap](../entries/sass-variable-bootstrap.yml) | Sass variable | 待查证 |
| [semantic-versioning-software](../entries/semantic-versioning-software.yml) | semantic versioning | 待查证 |
| [sentinel-error-go](../entries/sentinel-error-go.yml) | sentinel error | 待查证 |
| [server-rendering-react](../entries/server-rendering-react.yml) | server rendering | 待查证 |
| [shallow-copy-javascript](../entries/shallow-copy-javascript.yml) | shallow copy | 待查证 |
| [shock-resistance-combat](../entries/shock-resistance-combat.yml) | shock resistance | 本轮查证后暂缓。检索未取得电击义的整短语 IPA，继续保留旧暂缓。 |
| [short-variable-declaration-go](../entries/short-variable-declaration-go.yml) | short variable declaration | 待查证 |
| [shorthand-property-css](../entries/shorthand-property-css.yml) | shorthand property | 待查证 |
| [side-hop-combat](../entries/side-hop-combat.yml) | side hop | 待查证 |
| [sky-island-zelda](../entries/sky-island-zelda.yml) | sky island | 待查证 |
| [source-order-css](../entries/source-order-css.yml) | source order | 待查证 |
| [source-set-html](../entries/source-set-html.yml) | source set | 待查证 |
| [spacing-utility-bootstrap](../entries/spacing-utility-bootstrap.yml) | spacing utility | 待查证 |
| [stack-trace-tech](../entries/stack-trace-tech.yml) | stack trace | 前批查证后暂缓，本轮未重查。沿用第八批清单。本轮查证后暂缓。Collins 未取得，Wiktionary 未见整短语 IPA。 |
| [stacking-context-css](../entries/stacking-context-css.yml) | stacking context | 待查证 |
| [start-tag-html](../entries/start-tag-html.yml) | start tag | 待查证 |
| [state-preservation-react](../entries/state-preservation-react.yml) | state preservation | 待查证 |
| [state-setter-react](../entries/state-setter-react.yml) | state setter | 待查证 |
| [state-snapshot-react](../entries/state-snapshot-react.yml) | state snapshot | 待查证 |
| [sticky-positioning-css](../entries/sticky-positioning-css.yml) | sticky positioning | 待查证 |
| [storage-event-browser](../entries/storage-event-browser.yml) | storage event | 待查证 |
| [strict-mode-react](../entries/strict-mode-react.yml) | Strict Mode | 待查证 |
| [strike-price-finance](../entries/strike-price-finance.yml) | strike price | 本轮查证后暂缓。Cambridge 发音正文显示 UK、US 音频入口但无书面 IPA；会计词典未找到该词头。 |
| [structural-check-tech](../entries/structural-check-tech.yml) | structural check | 待查证 |
| [structural-typing-typescript](../entries/structural-typing-typescript.yml) | structural typing | 待查证 |
| [style-rule-css](../entries/style-rule-css.yml) | style rule | 待查证 |
| [stylesheet-css](../entries/stylesheet-css.yml) | stylesheet | 本轮查证后暂缓。Wiktionary 的闭合拼法有 IPA，未标口音。Oxford 收录分写 style sheet，保留当前拼法等待证据。 |
| [subtest-go](../entries/subtest-go.yml) | subtest | 待查证 |
| [symlink-tech](../entries/symlink-tech.yml) | symlink | 本轮查证后暂缓。Wiktionary 有技术定义但无适用的整词 IPA。 |
| [system-prompt-ai](../entries/system-prompt-ai.yml) | system prompt | 待查证 |
| [table-driven-test-go](../entries/table-driven-test-go.yml) | table-driven test | 待查证 |
| [table-row-html](../entries/table-row-html.yml) | table row | 待查证 |
| [template-literal-type-typescript](../entries/template-literal-type-typescript.yml) | template literal type | 待查证 |
| [temporal-dead-zone-javascript](../entries/temporal-dead-zone-javascript.yml) | temporal dead zone | 待查证 |
| [text-wrapping-css](../entries/text-wrapping-css.yml) | text wrapping | 待查证 |
| [theming-bootstrap](../entries/theming-bootstrap.yml) | theming | 本轮查证后暂缓。Wiktionary 可读页面未列整词 IPA。 |
| [toggle-trigger-bootstrap](../entries/toggle-trigger-bootstrap.yml) | toggle trigger | 待查证 |
| [tool-call-ai](../entries/tool-call-ai.yml) | tool call | 待查证 |
| [tool-result-ai](../entries/tool-result-ai.yml) | tool result | 待查证 |
| [tree-shaking-build](../entries/tree-shaking-build.yml) | tree shaking | 待查证 |
| [two-way-binding-frontend](../entries/two-way-binding-frontend.yml) | two-way binding | 待查证 |
| [type-alias-typescript](../entries/type-alias-typescript.yml) | type alias | 待查证 |
| [type-annotation-typescript](../entries/type-annotation-typescript.yml) | type annotation | 待查证 |
| [type-constraint-go](../entries/type-constraint-go.yml) | type constraint | 待查证 |
| [type-guard-typescript](../entries/type-guard-typescript.yml) | type guard | 待查证 |
| [type-inference-typescript](../entries/type-inference-typescript.yml) | type inference | 本轮查证后暂缓。Wiktionary 可读术语页未提供整短语 IPA。 |
| [type-instantiation-tech](../entries/type-instantiation-tech.yml) | type instantiation | 待查证 |
| [type-only-import-typescript](../entries/type-only-import-typescript.yml) | type-only import | 待查证 |
| [type-parameter-typescript](../entries/type-parameter-typescript.yml) | type parameter | 待查证 |
| [type-predicate-typescript](../entries/type-predicate-typescript.yml) | type predicate | 待查证 |
| [type-selector-css](../entries/type-selector-css.yml) | type selector | 待查证 |
| [type-switch-go](../entries/type-switch-go.yml) | type switch | 待查证 |
| [typeof-type-operator-typescript](../entries/typeof-type-operator-typescript.yml) | typeof type operator | 待查证 |
| [uncontrolled-component-react](../entries/uncontrolled-component-react.yml) | uncontrolled component | 待查证 |
| [unordered-list-html](../entries/unordered-list-html.yml) | unordered list | 待查证 |
| [updater-function-react](../entries/updater-function-react.yml) | updater function | 待查证 |
| [utility-class-bootstrap](../entries/utility-class-bootstrap.yml) | utility class | 待查证 |
| [utility-type-typescript](../entries/utility-type-typescript.yml) | utility type | 待查证 |
| [valid-state-bootstrap](../entries/valid-state-bootstrap.yml) | valid state | 待查证 |
| [validation-feedback-bootstrap](../entries/validation-feedback-bootstrap.yml) | validation feedback | 待查证 |
| [viewport-css](../entries/viewport-css.yml) | viewport | 拟补 B09-P02 |
| [visually-hidden-bootstrap](../entries/visually-hidden-bootstrap.yml) | visually hidden | 待查证 |
| [void-element-html](../entries/void-element-html.yml) | void element | 待查证 |
| [wait-group-go](../entries/wait-group-go.yml) | wait group | 待查证 |
| [web-storage-browser](../entries/web-storage-browser.yml) | web storage | 待查证 |
| [withdrawable-finance](../entries/withdrawable-finance.yml) | withdrawable | 本轮查证后暂缓。Collins 页面不可用，Wiktionary 未见适用整词 IPA。 |
| [worker-pool-concurrency](../entries/worker-pool-concurrency.yml) | worker pool | 待查证 |
| [year-over-year-finance](../entries/year-over-year-finance.yml) | year-over-year | 本轮查证后暂缓。Cambridge 返回 403；不同词头 YoY 的字母读法不可代用。 |
| [zero-value-go](../entries/zero-value-go.yml) | zero value | 待查证 |

### 03 缩写及技术标识符

| 稳定 ID | 当前 Word | 状态及依据 |
|:--|:--|:--|
| [bigint-javascript](../entries/bigint-javascript.yml) | BigInt | 待查证 |
| [bool-go](../entries/bool-go.yml) | bool | 前批查证后暂缓，本轮未重查。沿用第八批清单。本轮查证后暂缓。Wiktionary 命中的相邻语种 IPA 不适用。保留 Go 标识符，待查其口读约定。 |
| [comparable-go](../entries/comparable-go.yml) | comparable | 前批查证后暂缓，本轮未重查。沿用第八批清单。本轮查证后暂缓。Collins 有英语形容词读音。Go 预声明约束需要确认标识符的口读约定。 |
| [compose-yaml-docker](../entries/compose-yaml-docker.yml) | compose.yaml | 待查证 |
| [dom-html](../entries/dom-html.yml) | DOM | 本轮查证后暂缓。Wiktionary 的法语或印尼语音标、Cambridge 的税务 non-dom 均不适用文档对象模型。 |
| [ebit-finance](../entries/ebit-finance.yml) | EBIT | 拟补 B09-P08 |
| [fallback-ui-react](../entries/fallback-ui-react.yml) | fallback UI | 待查证 |
| [go-mod-go](../entries/go-mod-go.yml) | go.mod | 待查证 |
| [go-sum-go](../entries/go-sum-go.yml) | go.sum | 待查证 |
| [hdmi-cable-hardware](../entries/hdmi-cable-hardware.yml) | HDMI cable | 待查证 |
| [hdmi-type-a-connector](../entries/hdmi-type-a-connector.yml) | HDMI Type A connector | 待查证 |
| [http-caching-web](../entries/http-caching-web.yml) | HTTP caching | 待查证 |
| [id-selector-css](../entries/id-selector-css.yml) | ID selector | 待查证 |
| [int-go](../entries/int-go.yml) | int | 前批查证后暂缓，本轮未重查。沿用第八批清单。本轮查证后暂缓。Wiktionary 命中的相邻语种 IPA 不适用。保留 Go 标识符，待查其口读约定。 |
| [jsx-react](../entries/jsx-react.yml) | JSX | 本轮查证后暂缓。Wiktionary 有条目，未见该技术缩写的 IPA，需可核验的读法资料。 |
| [keyof-typescript](../entries/keyof-typescript.yml) | keyof | 待查证 |
| [new-go](../entries/new-go.yml) | new | 前批查证后暂缓，本轮未重查。沿用第八批清单。本轮查证后暂缓。Collins 美式音标为普通英语词。Go 内置函数口读惯例待确认。 |
| [omit-utility-type-typescript](../entries/omit-utility-type-typescript.yml) | Omit | 前批查证后暂缓，本轮未重查。沿用第八批清单。本轮查证后暂缓。Collins 提供普通动词 IPA。TypeScript 标识符的读法需要单独确认。 |
| [partial-utility-type-typescript](../entries/partial-utility-type-typescript.yml) | Partial | 前批查证后暂缓，本轮未重查。沿用第八批清单。本轮查证后暂缓。Collins 美式为形容词。当前为 TypeScript 工具类型名，待核实标识符口读依据。 |
| [pick-utility-type-typescript](../entries/pick-utility-type-typescript.yml) | Pick | 前批查证后暂缓，本轮未重查。沿用第八批清单。本轮查证后暂缓。Collins 提供普通动词 IPA。TypeScript 标识符的读法需要单独确认。 |
| [promise-all-javascript](../entries/promise-all-javascript.yml) | Promise.all | 待查证 |
| [property-css](../entries/property-css.yml) | CSS property | 待查证 |
| [readonly-typescript](../entries/readonly-typescript.yml) | readonly | 待查证 |
| [record-utility-type-typescript](../entries/record-utility-type-typescript.yml) | Record | 前批查证后暂缓，本轮未重查。沿用第八批清单。本轮查证后暂缓。Collins 区分 record 的名词和动词重音。当前为 TypeScript 工具类型，待核实口读依据。 |
| [returntype-utility-type-typescript](../entries/returntype-utility-type-typescript.yml) | ReturnType | 待查证 |
| [semantic-html](../entries/semantic-html.yml) | semantic HTML | 待查证 |
| [star-method-interview](../entries/star-method-interview.yml) | STAR method | 待查证 |
| [struct-go](../entries/struct-go.yml) | struct | 待查证 |
| [tts-language-tech](../entries/tts-language-tech.yml) | TTS language | 待查证 |
| [tv-mode-switch](../entries/tv-mode-switch.yml) | TV Mode | 待查证 |
| [utility-api-bootstrap](../entries/utility-api-bootstrap.yml) | utility API | 待查证 |
| [virtual-dom-frontend](../entries/virtual-dom-frontend.yml) | virtual DOM | 待查证 |

### 04 专名及游戏名称

| 稳定 ID | 当前 Word | 状态及依据 |
|:--|:--|:--|
| [acana-pet-food](../entries/acana-pet-food.yml) | ACANA | 待查证 |
| [adventure-log-zelda](../entries/adventure-log-zelda.yml) | Adventure Log | 待查证 |
| [aerocuda-zelda](../entries/aerocuda-zelda.yml) | Aerocuda | 待查证 |
| [ancient-arowana-zelda](../entries/ancient-arowana-zelda.yml) | Ancient Arowana | 前批查证后暂缓，本轮未重查。沿用第八批清单。本轮查证后暂缓。与普通鱼名一起筛查。相近英文或芬兰语拼写无法核实该游戏材料的全名读法。 |
| [attack-up-zelda](../entries/attack-up-zelda.yml) | Attack Up | 待查证 |
| [autobuild-zelda](../entries/autobuild-zelda.yml) | Autobuild | 待查证 |
| [beam-emitter-zelda](../entries/beam-emitter-zelda.yml) | Beam Emitter | 待查证 |
| [big-battery-zelda](../entries/big-battery-zelda.yml) | Big Battery | 待查证 |
| [big-wheel-zelda](../entries/big-wheel-zelda.yml) | Big Wheel | 待查证 |
| [bokoblin-zelda](../entries/bokoblin-zelda.yml) | Bokoblin | 待查证 |
| [bootstrap-framework](../entries/bootstrap-framework.yml) | Bootstrap | 本轮查证后暂缓。Collins IPA 对应普通词义，尚缺对应框架专名的证据。 |
| [boss-bokoblin-zelda](../entries/boss-bokoblin-zelda.yml) | Boss Bokoblin | 待查证 |
| [brightbloom-seed-zelda](../entries/brightbloom-seed-zelda.yml) | Brightbloom Seed | 待查证 |
| [brightcap-zelda](../entries/brightcap-zelda.yml) | Brightcap | 待查证 |
| [captain-construct-zelda](../entries/captain-construct-zelda.yml) | Captain Construct | 待查证 |
| [character-profiles-zelda](../entries/character-profiles-zelda.yml) | Character Profiles | 待查证 |
| [chuchu-zelda](../entries/chuchu-zelda.yml) | Chuchu | 待查证 |
| [claude-ai](../entries/claude-ai.yml) | Claude | 本轮查证后暂缓。Wiktionary 有人名的口音转写，未把该证据直接套用到 Anthropic 产品。 |
| [climbing-gear-zelda](../entries/climbing-gear-zelda.yml) | Climbing Gear | 待查证 |
| [codex-ai](../entries/codex-ai.yml) | Codex | 本轮查证后暂缓。Wiktionary 大写词头访问超时，普通名词 codex 的证据未直接转给产品专名。 |
| [colgera-zelda](../entries/colgera-zelda.yml) | Colgera | 待查证 |
| [construct-factory-zelda](../entries/construct-factory-zelda.yml) | Construct Factory | 待查证 |
| [construct-head-zelda](../entries/construct-head-zelda.yml) | Construct Head | 待查证 |
| [copilot-ai](../entries/copilot-ai.yml) | Copilot | 本轮查证后暂缓。Wiktionary 读取未取得产品名的转写。需要实际词条所指产品发音依据。 |
| [courser-bee-honey-zelda](../entries/courser-bee-honey-zelda.yml) | Courser Bee Honey | 待查证 |
| [crystal-refinery-zelda](../entries/crystal-refinery-zelda.yml) | Crystal Refinery | 待查证 |
| [crystallized-charge-zelda](../entries/crystallized-charge-zelda.yml) | Crystallized Charge | 待查证 |
| [cursor-ai](../entries/cursor-ai.yml) | Cursor | 本轮查证后暂缓。Wiktionary 大写词头的已见转写为德语及拉丁语，未核实 Cursor 产品的读法。 |
| [death-mountain-zelda](../entries/death-mountain-zelda.yml) | Death Mountain | 待查证 |
| [deepseek-ai](../entries/deepseek-ai.yml) | DeepSeek | 待查证 |
| [defense-up-zelda](../entries/defense-up-zelda.yml) | Defense Up | 待查证 |
| [demon-dragon-zelda](../entries/demon-dragon-zelda.yml) | Demon Dragon | 待查证 |
| [demon-king-ganondorf-zelda](../entries/demon-king-ganondorf-zelda.yml) | Demon King Ganondorf | 待查证 |
| [demon-king-zelda](../entries/demon-king-zelda.yml) | Demon King | 待查证 |
| [demon-kings-bow-zelda](../entries/demon-kings-bow-zelda.yml) | Demon King's Bow | 待查证 |
| [device-dispenser-zelda](../entries/device-dispenser-zelda.yml) | Device Dispenser | 待查证 |
| [docker-compose-docker](../entries/docker-compose-docker.yml) | Docker Compose | 待查证 |
| [docker-platform](../entries/docker-platform.yml) | Docker | 本轮查证后暂缓。Wiktionary 大写 Docker 的 IPA 对应地名或姓氏，小写 docker 的 IPA 对应码头工人；未取得软件专名的有据转写。 |
| [dots-ai](../entries/dots-ai.yml) | dots | 待查证 |
| [electro-cooking](../entries/electro-cooking.yml) | Electro | 前批查证后暂缓，本轮未重查。沿用王国之泪计划 ZD013。当前 Electro 为游戏效果名称前缀；Collins electro 已读分区为普通名词或缩略形式，不直接套用其音标。需要名称用法的英文读音证据。 |
| [endura-carrot-zelda](../entries/endura-carrot-zelda.yml) | Endura Carrot | 待查证 |
| [endura-shroom-zelda](../entries/endura-shroom-zelda.yml) | Endura Shroom | 待查证 |
| [energy-cell-zelda](../entries/energy-cell-zelda.yml) | Energy Cell | 待查证 |
| [energy-well-zelda](../entries/energy-well-zelda.yml) | Energy Well | 待查证 |
| [evermean-zelda](../entries/evermean-zelda.yml) | Evermean | 待查证 |
| [fancy-feast-pet-food](../entries/fancy-feast-pet-food.yml) | Fancy Feast | 待查证 |
| [fierce-deity-armor-zelda](../entries/fierce-deity-armor-zelda.yml) | Fierce Deity Armor | 待查证 |
| [fire-temple-zelda](../entries/fire-temple-zelda.yml) | Fire Temple | 待查证 |
| [flame-emitter-zelda](../entries/flame-emitter-zelda.yml) | Flame Emitter | 待查证 |
| [flamebreaker-armor-zelda](../entries/flamebreaker-armor-zelda.yml) | Flamebreaker Armor | 待查证 |
| [flurry-rush-zelda](../entries/flurry-rush-zelda.yml) | Flurry Rush | 待查证 |
| [flux-construct-zelda](../entries/flux-construct-zelda.yml) | Flux Construct | 待查证 |
| [forgotten-temple-zelda](../entries/forgotten-temple-zelda.yml) | Forgotten Temple | 待查证 |
| [fregate-pet-food](../entries/fregate-pet-food.yml) | Fregate | 待查证 |
| [froggy-sleeve-zelda](../entries/froggy-sleeve-zelda.yml) | Froggy Sleeve | 待查证 |
| [frost-emitter-zelda](../entries/frost-emitter-zelda.yml) | Frost Emitter | 待查证 |
| [frox-zelda](../entries/frox-zelda.yml) | Frox | 待查证 |
| [ganondorf-zelda](../entries/ganondorf-zelda.yml) | Ganondorf | 本轮查证后暂缓。定向词典检索和 Wiktionary 查询未取得可核实的英语 IPA。 |
| [gemini-ai](../entries/gemini-ai.yml) | Gemini | 本轮查证后暂缓。Wiktionary 有星座等义项的转写，未核实 Google 产品名及口音。 |
| [gerudo-desert-zelda](../entries/gerudo-desert-zelda.yml) | Gerudo Desert | 待查证 |
| [gerudo-town-zelda](../entries/gerudo-town-zelda.yml) | Gerudo Town | 待查证 |
| [gerudo-zelda](../entries/gerudo-zelda.yml) | Gerudo | 待查证 |
| [gibdo-zelda](../entries/gibdo-zelda.yml) | Gibdo | 待查证 |
| [gleeok-zelda](../entries/gleeok-zelda.yml) | Gleeok | 待查证 |
| [glide-shirt-zelda](../entries/glide-shirt-zelda.yml) | Glide Shirt | 待查证 |
| [gloom-recovery-zelda](../entries/gloom-recovery-zelda.yml) | Gloom Recovery | 待查证 |
| [gloom-resistance-zelda](../entries/gloom-resistance-zelda.yml) | Gloom Resistance | 待查证 |
| [gloom-spawn-zelda](../entries/gloom-spawn-zelda.yml) | Gloom Spawn | 待查证 |
| [gloom-sword-zelda](../entries/gloom-sword-zelda.yml) | Gloom Sword | 待查证 |
| [golden-apple-zelda](../entries/golden-apple-zelda.yml) | Golden Apple | 待查证 |
| [goron-city-zelda](../entries/goron-city-zelda.yml) | Goron City | 待查证 |
| [goron-spice-zelda](../entries/goron-spice-zelda.yml) | Goron Spice | 待查证 |
| [goron-zelda](../entries/goron-zelda.yml) | Goron | 待查证 |
| [great-abandoned-central-mine-zelda](../entries/great-abandoned-central-mine-zelda.yml) | Great Abandoned Central Mine | 待查证 |
| [great-deku-tree-zelda](../entries/great-deku-tree-zelda.yml) | Great Deku Tree | 待查证 |
| [great-eagle-bow-zelda](../entries/great-eagle-bow-zelda.yml) | Great Eagle Bow | 待查证 |
| [great-sky-island-zelda](../entries/great-sky-island-zelda.yml) | Great Sky Island | 待查证 |
| [grok-ai](../entries/grok-ai.yml) | Grok | 本轮查证后暂缓。词典检索未取得该产品义的可用 IPA；旧普通动词义需与品牌证据区分。 |
| [hateno-village-zelda](../entries/hateno-village-zelda.yml) | Hateno Village | 待查证 |
| [heart-container-zelda](../entries/heart-container-zelda.yml) | Heart Container | 待查证 |
| [hearty-radish-zelda](../entries/hearty-radish-zelda.yml) | Hearty Radish | 待查证 |
| [hearty-salmon-zelda](../entries/hearty-salmon-zelda.yml) | Hearty Salmon | 待查证 |
| [hearty-truffle-zelda](../entries/hearty-truffle-zelda.yml) | Hearty Truffle | 待查证 |
| [hermes-agent-ai](../entries/hermes-agent-ai.yml) | Hermes | 本轮查证后暂缓。Wiktionary 的相似词为神祇名或其他语言形式，须核实 Nous Research 的产品称呼。 |
| [hestu-zelda](../entries/hestu-zelda.yml) | Hestu | 待查证 |
| [hills-science-diet-pet-food](../entries/hills-science-diet-pet-food.yml) | Hill’s Science Diet | 待查证 |
| [hinox-zelda](../entries/hinox-zelda.yml) | Hinox | 待查证 |
| [homing-cart-zelda](../entries/homing-cart-zelda.yml) | Homing Cart | 待查证 |
| [horriblin-zelda](../entries/horriblin-zelda.yml) | Horriblin | 待查证 |
| [hover-stone-zelda](../entries/hover-stone-zelda.yml) | Hover Stone | 待查证 |
| [hudson-zelda](../entries/hudson-zelda.yml) | Hudson | 待查证 |
| [hylian-hood-zelda](../entries/hylian-hood-zelda.yml) | Hylian Hood | 待查证 |
| [hylian-shield-zelda](../entries/hylian-shield-zelda.yml) | Hylian Shield | 待查证 |
| [hylian-tunic-zelda](../entries/hylian-tunic-zelda.yml) | Hylian Tunic | 待查证 |
| [hylian-zelda](../entries/hylian-zelda.yml) | Hylian | 待查证 |
| [hyrule-bass-zelda](../entries/hyrule-bass-zelda.yml) | Hyrule Bass | 待查证 |
| [hyrule-castle-zelda](../entries/hyrule-castle-zelda.yml) | Hyrule Castle | 待查证 |
| [hyrule-compendium-zelda](../entries/hyrule-compendium-zelda.yml) | Hyrule Compendium | 待查证 |
| [hyrule-zelda](../entries/hyrule-zelda.yml) | Hyrule | 本轮查证后暂缓。Wiktionary 词头请求返回 404。所查资料无可核实的标口音英语 IPA。 |
| [impa-zelda](../entries/impa-zelda.yml) | Impa | 待查证 |
| [instinct-pet-food](../entries/instinct-pet-food.yml) | Instinct | 待查证 |
| [josha-zelda](../entries/josha-zelda.yml) | Josha | 待查证 |
| [kakariko-village-zelda](../entries/kakariko-village-zelda.yml) | Kakariko Village | 本轮查证后暂缓。全名定向检索未取得词典正文的整名 IPA。Kakariko 单独写法不可替代。 |
| [keese-zelda](../entries/keese-zelda.yml) | Keese | 待查证 |
| [key-items-zelda](../entries/key-items-zelda.yml) | Key Items | 待查证 |
| [king-gleeok-zelda](../entries/king-gleeok-zelda.yml) | King Gleeok | 待查证 |
| [knights-broadsword-zelda](../entries/knights-broadsword-zelda.yml) | Knight's Broadsword | 待查证 |
| [korok-forest-zelda](../entries/korok-forest-zelda.yml) | Korok Forest | 待查证 |
| [korok-seed-zelda](../entries/korok-seed-zelda.yml) | Korok Seed | 待查证 |
| [korok-zelda](../entries/korok-zelda.yml) | Korok | 待查证 |
| [legendsandy-pet-food](../entries/legendsandy-pet-food.yml) | Legendsandy | 待查证 |
| [lexforge-project](../entries/lexforge-project.yml) | LexForge | 待查证 |
| [light-dragon-zelda](../entries/light-dragon-zelda.yml) | Light Dragon | 待查证 |
| [light-of-blessing-zelda](../entries/light-of-blessing-zelda.yml) | Light of Blessing | 待查证 |
| [lightning-temple-zelda](../entries/lightning-temple-zelda.yml) | Lightning Temple | 待查证 |
| [lightroot-zelda](../entries/lightroot-zelda.yml) | Lightroot | 待查证 |
| [like-like-zelda](../entries/like-like-zelda.yml) | Like Like | 待查证 |
| [link-zelda](../entries/link-zelda.yml) | Link | 待查证 |
| [lizalfos-zelda](../entries/lizalfos-zelda.yml) | Lizalfos | 待查证 |
| [lookout-landing-zelda](../entries/lookout-landing-zelda.yml) | Lookout Landing | 待查证 |
| [lurelin-village-zelda](../entries/lurelin-village-zelda.yml) | Lurelin Village | 待查证 |
| [lynel-zelda](../entries/lynel-zelda.yml) | Lynel | 待查证 |
| [magic-rod-zelda](../entries/magic-rod-zelda.yml) | Magic Rod | 待查证 |
| [main-quests-zelda](../entries/main-quests-zelda.yml) | Main Quests | 待查证 |
| [marbled-gohma-zelda](../entries/marbled-gohma-zelda.yml) | Marbled Gohma | 待查证 |
| [master-kohga-zelda](../entries/master-kohga-zelda.yml) | Master Kohga | 待查证 |
| [master-sword-zelda](../entries/master-sword-zelda.yml) | Master Sword | 待查证 |
| [mighty-bananas-zelda](../entries/mighty-bananas-zelda.yml) | Mighty Bananas | 待查证 |
| [miners-mask-zelda](../entries/miners-mask-zelda.yml) | Miner's Mask | 待查证 |
| [mineru-zelda](../entries/mineru-zelda.yml) | Mineru | 待查证 |
| [moblin-zelda](../entries/moblin-zelda.yml) | Moblin | 待查证 |
| [mucktorok-zelda](../entries/mucktorok-zelda.yml) | Mucktorok | 待查证 |
| [muse-ai](../entries/muse-ai.yml) | Muse | 本轮查证后暂缓。Wiktionary 英语释义和音标指神话、灵感来源等，未核实 Meta 产品名。 |
| [myfoodie-pet-food](../entries/myfoodie-pet-food.yml) | Myfoodie | 待查证 |
| [nginx-platform](../entries/nginx-platform.yml) | NGINX | 本轮查证后暂缓。官方博客确认口读 engine X。xAI 的 TTS 文档给出完整 IPA 示例，但未标口音，保留待核实。 |
| [nutro-pet-food](../entries/nutro-pet-food.yml) | Nutro | 待查证 |
| [octorok-zelda](../entries/octorok-zelda.yml) | Octorok | 待查证 |
| [openclaw-ai](../entries/openclaw-ai.yml) | OpenClaw | 待查证 |
| [orijen-pet-food](../entries/orijen-pet-food.yml) | ORIJEN | 本轮查证后暂缓。检索取得发音音频网站，未取得可靠的整名书面音标。 |
| [paya-zelda](../entries/paya-zelda.yml) | Paya | 待查证 |
| [perfect-dodge-zelda](../entries/perfect-dodge-zelda.yml) | Perfect Dodge | 待查证 |
| [perfect-guard-zelda](../entries/perfect-guard-zelda.yml) | Perfect Guard | 待查证 |
| [phantom-ganon-zelda](../entries/phantom-ganon-zelda.yml) | Phantom Ganon | 待查证 |
| [portable-pot-zelda](../entries/portable-pot-zelda.yml) | Portable Pot | 待查证 |
| [princess-zelda-zelda](../entries/princess-zelda-zelda.yml) | Princess Zelda | 本轮查证后暂缓。Wiktionary 只给部分同名形式 Zelda 的 IPA，未找到整名并明确口音的证据。 |
| [puffshroom-zelda](../entries/puffshroom-zelda.yml) | Puffshroom | 待查证 |
| [purah-pad-zelda](../entries/purah-pad-zelda.yml) | Purah Pad | 待查证 |
| [purah-zelda](../entries/purah-zelda.yml) | Purah | 待查证 |
| [purina-pro-plan-pet-food](../entries/purina-pro-plan-pet-food.yml) | Purina Pro Plan | 待查证 |
| [queen-gibdo-zelda](../entries/queen-gibdo-zelda.yml) | Queen Gibdo | 待查证 |
| [rauru-zelda](../entries/rauru-zelda.yml) | Rauru | 待查证 |
| [raw-gourmet-meat-zelda](../entries/raw-gourmet-meat-zelda.yml) | Raw Gourmet Meat | 待查证 |
| [raw-prime-meat-zelda](../entries/raw-prime-meat-zelda.yml) | Raw Prime Meat | 待查证 |
| [react-library](../entries/react-library.yml) | React | 本轮查证后暂缓。Wiktionary 大写词头未取得可用发音，普通动词的读音需另证与库名相符。 |
| [regional-phenomena-zelda](../entries/regional-phenomena-zelda.yml) | Regional Phenomena | 待查证 |
| [riju-zelda](../entries/riju-zelda.yml) | Riju | 待查证 |
| [rito-village-zelda](../entries/rito-village-zelda.yml) | Rito Village | 待查证 |
| [rito-zelda](../entries/rito-zelda.yml) | Rito | 待查证 |
| [robbie-zelda](../entries/robbie-zelda.yml) | Robbie | 待查证 |
| [royal-broadsword-zelda](../entries/royal-broadsword-zelda.yml) | Royal Broadsword | 待查证 |
| [royal-canin-pet-food](../entries/royal-canin-pet-food.yml) | Royal Canin | 本轮查证后暂缓。品牌文件可确认拼写，检索命中发音音频，未取得有口音的整名 IPA。 |
| [scimitar-of-the-seven-zelda](../entries/scimitar-of-the-seven-zelda.yml) | Scimitar of the Seven | 待查证 |
| [secret-stone-zelda](../entries/secret-stone-zelda.yml) | Secret Stone | 待查证 |
| [seized-construct-zelda](../entries/seized-construct-zelda.yml) | Seized Construct | 待查证 |
| [sheba-pet-food](../entries/sheba-pet-food.yml) | Sheba | 待查证 |
| [sheikah-zelda](../entries/sheikah-zelda.yml) | Sheikah | 待查证 |
| [shock-emitter-zelda](../entries/shock-emitter-zelda.yml) | Shock Emitter | 待查证 |
| [shrine-quests-zelda](../entries/shrine-quests-zelda.yml) | Shrine Quests | 待查证 |
| [side-adventures-zelda](../entries/side-adventures-zelda.yml) | Side Adventures | 待查证 |
| [side-quests-zelda](../entries/side-quests-zelda.yml) | Side Quests | 待查证 |
| [sidon-zelda](../entries/sidon-zelda.yml) | Sidon | 待查证 |
| [silent-princess-zelda](../entries/silent-princess-zelda.yml) | Silent Princess | 待查证 |
| [skyshroom-zelda](../entries/skyshroom-zelda.yml) | Skyshroom | 待查证 |
| [skyview-tower-zelda](../entries/skyview-tower-zelda.yml) | Skyview Tower | 待查证 |
| [slip-resistance-zelda](../entries/slip-resistance-zelda.yml) | Slip Resistance | 待查证 |
| [small-wheel-zelda](../entries/small-wheel-zelda.yml) | Small Wheel | 待查证 |
| [snowquill-tunic-zelda](../entries/snowquill-tunic-zelda.yml) | Snowquill Tunic | 待查证 |
| [soldier-construct-zelda](../entries/soldier-construct-zelda.yml) | Soldier Construct | 待查证 |
| [soldiers-broadsword-zelda](../entries/soldiers-broadsword-zelda.yml) | Soldier's Broadsword | 待查证 |
| [sonia-zelda](../entries/sonia-zelda.yml) | Sonia | 待查证 |
| [special-controls-zelda](../entries/special-controls-zelda.yml) | Special Controls | 待查证 |
| [speed-up-zelda](../entries/speed-up-zelda.yml) | Speed Up | 待查证 |
| [spicy-pepper-zelda](../entries/spicy-pepper-zelda.yml) | Spicy Pepper | 待查证 |
| [spirit-temple-zelda](../entries/spirit-temple-zelda.yml) | Spirit Temple | 待查证 |
| [stealth-mask-zelda](../entries/stealth-mask-zelda.yml) | Stealth Mask | 待查证 |
| [stealth-up-zelda](../entries/stealth-up-zelda.yml) | Stealth Up | 待查证 |
| [steering-stick-zelda](../entries/steering-stick-zelda.yml) | Steering Stick | 待查证 |
| [stone-talus-zelda](../entries/stone-talus-zelda.yml) | Stone Talus | 待查证 |
| [sundelion-zelda](../entries/sundelion-zelda.yml) | Sundelion | 待查证 |
| [tarrey-town-zelda](../entries/tarrey-town-zelda.yml) | Tarrey Town | 待查证 |
| [teba-zelda](../entries/teba-zelda.yml) | Teba | 待查证 |
| [temple-of-time-zelda](../entries/temple-of-time-zelda.yml) | Temple of Time | 待查证 |
| [the-depths-zelda](../entries/the-depths-zelda.yml) | The Depths | 待查证 |
| [time-bomb-zelda](../entries/time-bomb-zelda.yml) | Time Bomb | 待查证 |
| [travelers-sword-zelda](../entries/travelers-sword-zelda.yml) | Traveler's Sword | 待查证 |
| [tulin-zelda](../entries/tulin-zelda.yml) | Tulin | 待查证 |
| [ultrahand-zelda](../entries/ultrahand-zelda.yml) | Ultrahand | 待查证 |
| [water-temple-zelda](../entries/water-temple-zelda.yml) | Water Temple | 待查证 |
| [whiskas-pet-food](../entries/whiskas-pet-food.yml) | Whiskas | 本轮查证后暂缓。Cambridge 可读快照给出口音及音标，但 h 周围的展示格式未能从 403 的直开页面核实，暂不定稿。 |
| [wind-temple-zelda](../entries/wind-temple-zelda.yml) | Wind Temple | 待查证 |
| [wizzrobe-zelda](../entries/wizzrobe-zelda.yml) | Wizzrobe | 待查证 |
| [yiga-clan-hideout-zelda](../entries/yiga-clan-hideout-zelda.yml) | Yiga Clan Hideout | 待查证 |
| [yiga-clan-zelda](../entries/yiga-clan-zelda.yml) | Yiga Clan | 待查证 |
| [yona-zelda](../entries/yona-zelda.yml) | Yona | 待查证 |
| [yunobo-zelda](../entries/yunobo-zelda.yml) | Yunobo | 待查证 |
| [zonai-device-zelda](../entries/zonai-device-zelda.yml) | Zonai Device | 待查证 |
| [zonai-zelda](../entries/zonai-zelda.yml) | Zonai | 待查证 |
| [zonaite-zelda](../entries/zonaite-zelda.yml) | Zonaite | 待查证 |
| [zora-armor-zelda](../entries/zora-armor-zelda.yml) | Zora Armor | 待查证 |
| [zora-zelda](../entries/zora-zelda.yml) | Zora | 待查证 |
| [zoras-domain-zelda](../entries/zoras-domain-zelda.yml) | Zora's Domain | 待查证 |

### 05 语法模板、符号与命令

| 稳定 ID | 当前 Word | 状态及依据 |
|:--|:--|:--|
| [async-await-javascript](../entries/async-await-javascript.yml) | async/await | 待查证 |
| [box-sizing-css](../entries/box-sizing-css.yml) | box-sizing | 待查证 |
| [fr-css](../entries/fr-css.yml) | fr | 待查证 |
| [important-css](../entries/important-css.yml) | !important | 待查证 |
| [prefix-priority-nginx](../entries/prefix-priority-nginx.yml) | ^~ | 待查证 |
| [proxy-pass-nginx](../entries/proxy-pass-nginx.yml) | proxy_pass | 待查证 |
| [rem-css](../entries/rem-css.yml) | rem | 待查证 |
| [transform-css](../entries/transform-css.yml) | transform | 待查证 |
| [try-files-nginx](../entries/try-files-nginx.yml) | try_files | 待查证 |
| [var-function-css](../entries/var-function-css.yml) | var() | 待查证 |
| [would-you-like-a-or-b](../entries/would-you-like-a-or-b.yml) | Would you like A or B? | 待查证 |
| [z-index-css](../entries/z-index-css.yml) | z-index | 待查证 |

## 批准后的实施与验证

1. 检查当前 Git 状态，重新核对获批稳定 ID 的词形、义项、词性和现值。快照原始 YAML、导出和媒体哈希。若导出已落后于源文件，先用实施前 YAML 在隔离路径生成可比版本。
2. 只添加获批的 `pronunciation` 和 Notes 来源行。完整保留原 Notes，跳过已有的同一来源。保留现有音频字节、引用和语言设置，无论书面音标采用什么口音。
3. 从仓库根目录运行以下命令。对未跟踪文件另查格式及链接。

```bash
go run scripts/validate/main.go
go run scripts/build/main.go
git diff --check
```

4. 按稳定 ID 对比原始 YAML、导出及媒体哈希。检查 18 列的字段顺序、ID 与行数、获批音标和 Notes、其余列、Audio 及 StudyOrder。重复构建并对比字节，再统计剩余缺口。更新计划，记载实际执行结果和任何偏差。
5. 单独记录人工听审及真实 Anki 验证。静态导出检查仅验证字段和导入数据。

## 规划验证

已核对全库稳定 ID、语言计数、缺失清单、审阅编号和本地链接。按规划前哈希验证既有文件。规划期间只新增本计划，未运行构建或音频命令。收到明确批准后实施。

## 实施结果

- 按 B09-P01–B09-P08 更新 8 个 YAML 文件，添加 8 项书面 IPA 和 8 行 Notes 来源。原有 Notes 内容和其余 YAML 字段均保留。
- `go run scripts/validate/main.go` 校验 1,833 项通过，报告 38 条同词形警告。两次 `go run scripts/build/main.go` 生成字节一致的导出。`git diff --check` 通过，未跟踪计划的格式和本地链接另行核对。
- 对比实施前后的 1,833 行、18 列导出。获批条目只有 Pronunciation 与 Notes 列有改动。所有 ID、Audio、StudyOrder 和其他列一致。快照中 1,875 个 MP3 的哈希一致。
- 英语 1,026 项已有书面读音，564 项仍缺失。日语 243 项均有读法。本轮未进行人工听审、真实 Anki 导入或播放验证。
