# 书面读音审查计划：全库盘点与第八批英语音标

状态为已批准并实施。用户批准 B08-P01–B08-P11。下文保留 2026-11-22 规划基线和审阅文案；实施结果见末节。

## 盘点

以当前 YAML 的 `id` 和 `language` 为准。空值、空白文本或缺少 `pronunciation` 均计为缺失。已读取 [AGENTS.md](../AGENTS.md)、[README.md](../README.md) 和[词条规范](ENTRY_SCHEMA.md#pronunciation)。工作区包含既有的未提交修改，盘点使用包含这些修改的实际文件。[第七批](EN_PRONUNCIATION_ENRICHMENT_BATCH_07_PLAN.md)记录保持原样。

| 语言 | 总数 | 已有书面读音 | 缺失 |
|:--|--:|--:|--:|
| `en` | 1,590 | 1,007 | 583 |
| `ja` | 243 | 243 | 0 |
| 合计 | 1,833 | 1,250 | 583 |

本轮逐项查证 62 项，其中拟补 11 项（美式 6、英式 5），本轮查证后暂缓 51 项。另有 6 项沿用第七批的暂缓记录，本轮未重查。其余 515 项已盘点、待查证。四种状态合计 583 项。若本批全部实施且基线保持不变，缺失降至 572 项。

### 已有读音的疑点检查

已填英语项检查斜线及口音前缀格式，日语项检查是否夹有 ASCII 字母。未检出这两类结构异常。此静态检查不复核 IPA 的音段、重音、口音标记准确性或每个日语词条的实际读法。已填内容继续保留；将来发现有来源支持的疑点时另列审阅。

## 审阅范围与证据

A 批优先覆盖日常词形和可独立记忆的短语。B 批覆盖已核实的技术、金融和游戏通用词。后续沿下方队列查证，条数按证据和学习价值决定。专名先核实身份或游戏出处；缩写、符号及标识符先确定实际口读惯例。示例句、占位模板及自由组合不根据单词音标拼装。

采用词典释义、发音正文或原出版方教材的可读内容核对词形与词性。Cambridge 的发音转写来自可读词典正文快照。直接请求的 403 不当作词条不存在的证据。Wiktionary 只采用英语分区内明确标为 US 或 General American、并覆盖目标用法的 IPA。保留所用来源的转写系统。Collins 采用注明口音的分区。英国拼写、页面 URL 中的地区或未标注的音频本身不足以判定某条 IPA 的口音。

下面每项同时给出拟写读音和逐字追加的 Notes 来源行。既有 Notes 完整保留；若已包含同一读音来源则不重复追加。不同口音的录音继续保留。

### 批次 A · 常用词与短语

| 审阅 ID | 稳定 ID / Word | 当前词性 | 当前 Meaning | 拟写 `pronunciation` | 适用性与证据 |
|:--|:--|:--|:--|:--|:--|
| B08-P01 | [saute-cooking](../entries/saute-cooking.yml) / sauté | verb | 用少量油快速翻炒或煎炒。 | `/sɔːˈteɪ/` | 美式完整词形。Cambridge 释义页另列烹饪动词。词条为原形 sauté。 [Cambridge Dictionary — sauté (English pronunciation)](https://dictionary.cambridge.org/pronunciation/english/saute) |
| B08-P02 | [noodles-food](../entries/noodles-food.yml) / noodles | noun | 面条 | `UK /ˈnuːdlz/` | 词典在可读正文中只显示英式复数 IPA。Cambridge 的美式 noodle 是单数，Wiktionary 的复数未标口音。本次未核实美式复数 IPA。 [Cambridge Learner’s Dictionary — noodles](https://dictionary.cambridge.org/dictionary/learner-english/noodles) |
| B08-P03 | [pick-up-general](../entries/pick-up-general.yml) / pick up | verb | 捡起 | `UK /pɪk ˈʌp/` | 英式教材逐词表收录完整短语动词，例句为捡起衣服。Cambridge 只给动词 pick 的 IPA，Wiktionary 美式只有音频，另一个完整转写来自以模型撰写为主的词典。此次未核实可靠的美式整短语 IPA。 [Oxford University Press — English File Pre-intermediate Wordlist, File 11](https://yabancidiller.medeniyet.edu.tr/documents/yabancidiller/2021-2022%20AKADEM%C4%B0K%20YIL/%C3%96%C4%9ERET%C4%B0M%20YILI%20%C4%B0NG%C4%B0L%C4%B0ZCE%20HAZIRLIK%20PROGRAMI/word_list_english-file-pre-intermediate.pdf) |
| B08-P04 | [railway-track-engineering](../entries/railway-track-engineering.yml) / railway track | noun | 铁路轨道；铁轨 | `UK /ˈreɪlˌweɪ træk/` | Collins 明列英式整短语及轨道义。Cambridge 和 Wiktionary 未取得可用的美式整短语 IPA。 [Collins Dictionary — railway track (British English)](https://www.collinsdictionary.com/dictionary/english/railway-track) |
| B08-P05 | [weak-point-combat](../entries/weak-point-combat.yml) / weak point | noun | 弱点；容易受到攻击的部位 | `UK /wiːk pɔɪnt/` | Collins 整短语名词义包含薄弱处，可用于受攻击部位。Cambridge 和 Wiktionary 未取得可用的美式整短语 IPA。 [Collins Dictionary — weak point (British English)](https://www.collinsdictionary.com/dictionary/english/weak-point) |

### 批次 B · 技术、金融及游戏通用词

| 审阅 ID | 稳定 ID / Word | 当前词性 | 当前 Meaning | 拟写 `pronunciation` | 适用性与证据 |
|:--|:--|:--|:--|:--|:--|
| B08-P06 | [holdings-finance](../entries/holdings-finance.yml) / holdings | noun | 持有的资产；投资组合中的证券等投资项目。 | `/ˈhoʊl.dɪŋz/` | 美式发音页列出完整复数。Cambridge 同词释义页列资产所有权义。 [Cambridge Dictionary — holdings (English pronunciation)](https://dictionary.cambridge.org/us/pronunciation/english/holdings) |
| B08-P07 | [hydration-react](../entries/hydration-react.yml) / hydration | noun | 水合，让客户端 React 接管已有服务端 HTML 并使界面可交互 | `/haɪˈdɹeɪʃən/` | English 正文标明 US，且名词义第 5 项是 Web 客户端处理预渲染数据。 [Wiktionary — hydration (English pronunciation)](https://en.wiktionary.org/wiki/hydration) |
| B08-P08 | [orchestration-ai](../entries/orchestration-ai.yml) / orchestration | noun | 协调执行；安排工具、智能体或步骤的执行顺序与结果流转。 | `/ˌɔːr.kəˈstreɪ.ʃən/` | 美国发音页给出完整名词 IPA。词条为协调安排义在代理工作流中的应用。渲染抓取中弱读音节的字形布局空白没有写入字段。 [Cambridge Dictionary — orchestration (English pronunciation)](https://dictionary.cambridge.org/us/pronunciation/english/orchestration) |
| B08-P09 | [soundness-type-system](../entries/soundness-type-system.yml) / soundness | noun | 健全性；类型检查所接受的程序不会违反该系统承诺的类型安全性质。 | `/ˈsaʊnd.nəs/` | 美式完整名词 IPA。同词释义页的语料含逻辑中的 soundness proof。具体类型系统含义沿用当前词条。 [Cambridge Dictionary — soundness (English pronunciation)](https://dictionary.cambridge.org/pronunciation/english/soundness) |
| B08-P10 | [easing-css](../entries/easing-css.yml) / easing | noun | 缓动，控制动画或过渡的变化速度如何随时间改变 | `UK /ˈiːzɪŋ/` | Collins 的确切名词词形为动作逐渐缓和，卡片取动画中的缓动义。Cambridge 可读发音页仅标英式，Wiktionary 的 IPA 未标口音。未采用自动转写工具的美式结果。 [Collins Dictionary — easing (British English)](https://www.collinsdictionary.com/dictionary/english/easing) |
| B08-P11 | [quicksave-game](../entries/quicksave-game.yml) / quicksave | noun | 快速存档；通过快捷操作保存的进度 | `/ˈkwɪkˌseɪv/` | General American 完整词形。该修订的名词第二义是快捷操作保存的进度。 [Wiktionary — quicksave (English pronunciation)](https://en.wiktionary.org/w/index.php?title=quicksave&oldid=93717634) |

### 拟追加的 Notes 来源行

- B08-P01 `Pronunciation source: Cambridge Dictionary — sauté (English pronunciation). https://dictionary.cambridge.org/pronunciation/english/saute (accent: US; accessed 2026-11-22).`
- B08-P02 `Pronunciation source: Cambridge Learner’s Dictionary — noodles. https://dictionary.cambridge.org/dictionary/learner-english/noodles (accent: UK; accessed 2026-11-22).`
- B08-P03 `Pronunciation source: Oxford University Press — English File Pre-intermediate Wordlist, File 11. https://yabancidiller.medeniyet.edu.tr/documents/yabancidiller/2021-2022%20AKADEM%C4%B0K%20YIL/%C3%96%C4%9ERET%C4%B0M%20YILI%20%C4%B0NG%C4%B0L%C4%B0ZCE%20HAZIRLIK%20PROGRAMI/word_list_english-file-pre-intermediate.pdf (accent: UK; accessed 2026-11-22).`
- B08-P04 `Pronunciation source: Collins Dictionary — railway track (British English). https://www.collinsdictionary.com/dictionary/english/railway-track (accent: UK; accessed 2026-11-22).`
- B08-P05 `Pronunciation source: Collins Dictionary — weak point (British English). https://www.collinsdictionary.com/dictionary/english/weak-point (accent: UK; accessed 2026-11-22).`
- B08-P06 `Pronunciation source: Cambridge Dictionary — holdings (English pronunciation). https://dictionary.cambridge.org/us/pronunciation/english/holdings (accent: US; accessed 2026-11-22).`
- B08-P07 `Pronunciation source: Wiktionary — hydration (English pronunciation). https://en.wiktionary.org/wiki/hydration (accent: US; accessed 2026-11-22).`
- B08-P08 `Pronunciation source: Cambridge Dictionary — orchestration (English pronunciation). https://dictionary.cambridge.org/us/pronunciation/english/orchestration (accent: US; accessed 2026-11-22).`
- B08-P09 `Pronunciation source: Cambridge Dictionary — soundness (English pronunciation). https://dictionary.cambridge.org/pronunciation/english/soundness (accent: US; accessed 2026-11-22).`
- B08-P10 `Pronunciation source: Collins Dictionary — easing (British English). https://www.collinsdictionary.com/dictionary/english/easing (accent: UK; accessed 2026-11-22).`
- B08-P11 `Pronunciation source: Wiktionary — quicksave (English pronunciation). https://en.wiktionary.org/w/index.php?title=quicksave&oldid=93717634 (accent: US; accessed 2026-11-22).`

## 缺失 ID 清单

每个缺失稳定 ID 在下表仅有一行。原词形按当前 YAML 呈现。队列表示主要查证方向，同一词条可能同时属于多个学习主题。

| 队列 | 缺失 | 拟补 | 本轮暂缓 | 前批暂缓 | 待查证 |
|:--|--:|--:|--:|--:|--:|
| 01 常用词与短语 | 13 | 5 | 4 | 0 | 4 |
| 02 技术、金融与游戏通用语 | 315 | 6 | 37 | 6 | 266 |
| 03 缩写及技术标识符 | 32 | 0 | 8 | 0 | 24 |
| 04 专名及游戏名称 | 211 | 0 | 2 | 0 | 209 |
| 05 语法模板、符号与命令 | 12 | 0 | 0 | 0 | 12 |

### 01 常用词与短语

| 稳定 ID | 当前 Word | 状态及依据 |
|:--|:--|:--|
| [arowana-general](../entries/arowana-general.yml) | arowana | 本轮查证后暂缓。Collins 词页未取得。Wiktionary 英语 arowana 未见 IPA；检索命中的近似 aravana 为芬兰语。 |
| [cooking-pot-cooking](../entries/cooking-pot-cooking.yml) | cooking pot | 本轮查证后暂缓。Collins 可读正文列组合用例及 cook、pot 的分项音标，Cambridge 与 Wiktionary 未取得整短语音标。 |
| [coordinates-ui](../entries/coordinates-ui.yml) | coordinates | 本轮查证后暂缓。Collins 带准确词形的分区指配套服饰。未核实坐标义同形复数的转写。 |
| [cross-functional-collaboration-workflow](../entries/cross-functional-collaboration-workflow.yml) | cross-functional collaboration | 待查证 |
| [greaves-clothing](../entries/greaves-clothing.yml) | greaves | 本轮查证后暂缓。Collins 明标口音的 greaves 音标属于油渣义。护胫义的 greave 为单数；Wiktionary 的复数 IPA 未标口音。 |
| [lex-name-element](../entries/lex-name-element.yml) | lex | 待查证 |
| [noodles-food](../entries/noodles-food.yml) | noodles | 拟补 B08-P02 |
| [pick-up-general](../entries/pick-up-general.yml) | pick up | 拟补 B08-P03 |
| [project-onboarding-workflow](../entries/project-onboarding-workflow.yml) | project onboarding | 待查证 |
| [railway-track-engineering](../entries/railway-track-engineering.yml) | railway track | 拟补 B08-P04 |
| [related-words-language](../entries/related-words-language.yml) | related words | 待查证 |
| [saute-cooking](../entries/saute-cooking.yml) | sauté | 拟补 B08-P01 |
| [weak-point-combat](../entries/weak-point-combat.yml) | weak point | 拟补 B08-P05 |

### 02 技术、金融与游戏通用语

| 稳定 ID | 当前 Word | 状态及依据 |
|:--|:--|:--|
| [absolute-positioning-css](../entries/absolute-positioning-css.yml) | absolute positioning | 待查证 |
| [abstract-syntax-tree-compiler](../entries/abstract-syntax-tree-compiler.yml) | abstract syntax tree | 待查证 |
| [access-log-nginx](../entries/access-log-nginx.yml) | access log | 待查证 |
| [accessible-name-html](../entries/accessible-name-html.yml) | accessible name | 待查证 |
| [accessor-tech](../entries/accessor-tech.yml) | accessor | 本轮查证后暂缓。补查 Wiktionary 未见英语 IPA。Collins 的访问限制见第七批。 |
| [alternative-text-html](../entries/alternative-text-html.yml) | alternative text | 待查证 |
| [annual-percentage-yield-finance](../entries/annual-percentage-yield-finance.yml) | annual percentage yield | 本轮查证后暂缓。Collins 与 Wiktionary 检查未取得可用的整短语 IPA。 |
| [array-hole-javascript](../entries/array-hole-javascript.yml) | array hole | 待查证 |
| [ask-price-finance](../entries/ask-price-finance.yml) | ask price | 本轮查证后暂缓。Collins 页面未取得，Wiktionary 未见可用整短语 IPA。 |
| [assignability-typescript](../entries/assignability-typescript.yml) | assignability | 待查证 |
| [at-rule-css](../entries/at-rule-css.yml) | at-rule | 待查证 |
| [attack-power-combat](../entries/attack-power-combat.yml) | attack power | 待查证 |
| [attribute-selector-css](../entries/attribute-selector-css.yml) | attribute selector | 待查证 |
| [auto-save-game](../entries/auto-save-game.yml) | auto-save | 本轮查证后暂缓。Cambridge 定向查询未取得当前连字符形与游戏名词义的 IPA。 |
| [backend-server-nginx](../entries/backend-server-nginx.yml) | backend server | 待查证 |
| [backing-array-go](../entries/backing-array-go.yml) | backing array | 待查证 |
| [base-image-docker](../entries/base-image-docker.yml) | base image | 待查证 |
| [batching-react](../entries/batching-react.yml) | batching | 第七批查证后暂缓，本轮未重查。Collins 未取得 batching 独立分区或完整 batching IPA。 |
| [beat-estimates-finance](../entries/beat-estimates-finance.yml) | beat estimates | 待查证 |
| [bid-price-finance](../entries/bid-price-finance.yml) | bid price | 本轮查证后暂缓。Collins 确认短语义，未见标口音整短语 IPA。Cambridge 返回 403。 |
| [binary-search-algorithm](../entries/binary-search-algorithm.yml) | binary search | 本轮查证后暂缓。Collins 未取得；Wiktionary 可读词条未见整短语 IPA。 |
| [bind-mount-docker](../entries/bind-mount-docker.yml) | bind mount | 待查证 |
| [block-formatting-context-css](../entries/block-formatting-context-css.yml) | block formatting context | 待查证 |
| [block-level-box-css](../entries/block-level-box-css.yml) | block-level box | 待查证 |
| [boolean-attribute-html](../entries/boolean-attribute-html.yml) | boolean attribute | 待查证 |
| [box-model-css](../entries/box-model-css.yml) | box model | 待查证 |
| [buffered-channel-go](../entries/buffered-channel-go.yml) | buffered channel | 待查证 |
| [build-cache-docker](../entries/build-cache-docker.yml) | build cache | 待查证 |
| [build-context-docker](../entries/build-context-docker.yml) | build context | 待查证 |
| [button-group-bootstrap](../entries/button-group-bootstrap.yml) | button group | 待查证 |
| [cache-invalidation-tech](../entries/cache-invalidation-tech.yml) | cache invalidation | 本轮查证后暂缓。Cambridge 定向查询未取得匹配的整短语音标。 |
| [cache-revalidation-web](../entries/cache-revalidation-web.yml) | cache revalidation | 待查证 |
| [call-stack-javascript](../entries/call-stack-javascript.yml) | call stack | 待查证 |
| [cascade-layer-css](../entries/cascade-layer-css.yml) | cascade layer | 待查证 |
| [cash-sweep-finance](../entries/cash-sweep-finance.yml) | cash sweep | 待查证 |
| [character-encoding-html](../entries/character-encoding-html.yml) | character encoding | 待查证 |
| [character-reference-html](../entries/character-reference-html.yml) | character reference | 待查证 |
| [class-selector-css](../entries/class-selector-css.yml) | class selector | 待查证 |
| [cold-resistance-combat](../entries/cold-resistance-combat.yml) | cold resistance | 待查证 |
| [color-mode-bootstrap](../entries/color-mode-bootstrap.yml) | color mode | 待查证 |
| [column-offset-bootstrap](../entries/column-offset-bootstrap.yml) | column offset | 待查证 |
| [column-span-html](../entries/column-span-html.yml) | column span | 待查证 |
| [combinator-css](../entries/combinator-css.yml) | combinator | 本轮查证后暂缓。Wiktionary 对英语整词有标口音音标，但其义项是 λ 演算，尚未找到与 CSS 选择器义匹配的发音证据。 |
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
| [context-window-ai](../entries/context-window-ai.yml) | context window | 本轮查证后暂缓。Cambridge 定向查询未取得匹配的整短语音标。 |
| [controlled-component-react](../entries/controlled-component-react.yml) | controlled component | 待查证 |
| [corrupted-combat](../entries/corrupted-combat.yml) | corrupted | 第七批查证后暂缓，本轮未重查。Collins 未取得明确美式整词 IPA。Wiktionary 的 IPA 未明确标口音，另有英格兰南部音频。 |
| [critical-hit-combat](../entries/critical-hit-combat.yml) | critical hit | 待查证 |
| [cross-axis-css](../entries/cross-axis-css.yml) | cross axis | 待查证 |
| [cross-origin-resource-sharing-web](../entries/cross-origin-resource-sharing-web.yml) | cross-origin resource sharing | 待查证 |
| [custom-hook-react](../entries/custom-hook-react.yml) | custom Hook | 待查证 |
| [custom-property-css](../entries/custom-property-css.yml) | custom property | 待查证 |
| [data-attribute-bootstrap](../entries/data-attribute-bootstrap.yml) | data attribute | 待查证 |
| [data-cell-html](../entries/data-cell-html.yml) | data cell | 待查证 |
| [data-race-go](../entries/data-race-go.yml) | data race | 待查证 |
| [data-table-html](../entries/data-table-html.yml) | data table | 待查证 |
| [debounce-javascript](../entries/debounce-javascript.yml) | debounce | 本轮查证后暂缓。Wiktionary 可见的英式音标适用于动词，当前卡片为防抖技术/策略名词。 |
| [declaration-file-typescript](../entries/declaration-file-typescript.yml) | declaration file | 待查证 |
| [decorative-image-html](../entries/decorative-image-html.yml) | decorative image | 待查证 |
| [deep-copy-javascript](../entries/deep-copy-javascript.yml) | deep copy | 待查证 |
| [dependency-array-react](../entries/dependency-array-react.yml) | dependency array | 待查证 |
| [derived-state-react](../entries/derived-state-react.yml) | derived state | 待查证 |
| [detached-mode-docker](../entries/detached-mode-docker.yml) | detached mode | 待查证 |
| [disabled-control-html](../entries/disabled-control-html.yml) | disabled control | 待查证 |
| [discriminated-union-typescript](../entries/discriminated-union-typescript.yml) | discriminated union | 待查证 |
| [display-utility-bootstrap](../entries/display-utility-bootstrap.yml) | display utility | 待查证 |
| [dockerfile-docker](../entries/dockerfile-docker.yml) | Dockerfile | 待查证 |
| [doctype-html](../entries/doctype-html.yml) | doctype | 本轮查证后暂缓。定向检索未找到可核实的 HTML 术语或同义名词转写。 |
| [document-title-html](../entries/document-title-html.yml) | document title | 待查证 |
| [draconification-zelda](../entries/draconification-zelda.yml) | draconification | 待查证 |
| [dropdown-bootstrap](../entries/dropdown-bootstrap.yml) | dropdown | 本轮查证后暂缓。Wiktionary 未见可用英语 IPA。Collins 的访问限制见第七批。 |
| [earnings-per-share-finance](../entries/earnings-per-share-finance.yml) | earnings per share | 本轮查证后暂缓。Collins 有财务释义，无明确口音的整短语 IPA；Cambridge 返回 403。 |
| [easing-css](../entries/easing-css.yml) | easing | 拟补 B08-P10 |
| [effect-cleanup-react](../entries/effect-cleanup-react.yml) | Effect cleanup | 待查证 |
| [end-tag-html](../entries/end-tag-html.yml) | end tag | 待查证 |
| [error-boundary-react](../entries/error-boundary-react.yml) | error boundary | 待查证 |
| [error-log-nginx](../entries/error-log-nginx.yml) | error log | 待查证 |
| [error-wrapping-go](../entries/error-wrapping-go.yml) | error wrapping | 待查证 |
| [event-bubbling-dom](../entries/event-bubbling-dom.yml) | event bubbling | 待查证 |
| [event-capturing-dom](../entries/event-capturing-dom.yml) | event capturing | 待查证 |
| [event-delegation-dom](../entries/event-delegation-dom.yml) | event delegation | 待查证 |
| [event-handler-react](../entries/event-handler-react.yml) | event handler | 待查证 |
| [event-loop-javascript](../entries/event-loop-javascript.yml) | event loop | 待查证 |
| [exact-match-nginx](../entries/exact-match-nginx.yml) | exact match | 待查证 |
| [exchange-traded-fund-finance](../entries/exchange-traded-fund-finance.yml) | exchange-traded fund | 本轮查证后暂缓。Collins 未见整短语 IPA，分项发音不能直接代用。Wiktionary 未见适用 IPA。 |
| [exhaustiveness-checking-typescript](../entries/exhaustiveness-checking-typescript.yml) | exhaustiveness checking | 待查证 |
| [expected-move-finance](../entries/expected-move-finance.yml) | expected move | 待查证 |
| [fall-damage-game](../entries/fall-damage-game.yml) | fall damage | 本轮查证后暂缓。Cambridge 定向查询未取得确切游戏义整短语 IPA。 |
| [fallback-value-css](../entries/fallback-value-css.yml) | fallback value | 待查证 |
| [fear-and-greed-index-finance](../entries/fear-and-greed-index-finance.yml) | fear and greed index | 待查证 |
| [fixed-positioning-css](../entries/fixed-positioning-css.yml) | fixed positioning | 待查证 |
| [flex-container-css](../entries/flex-container-css.yml) | flex container | 待查证 |
| [flex-item-css](../entries/flex-item-css.yml) | flex item | 待查证 |
| [flexbox-css](../entries/flexbox-css.yml) | flexbox | 本轮查证后暂缓。定向检索未找到可核实的、与 CSS 概念匹配的整词转写。 |
| [floating-point-precision-javascript](../entries/floating-point-precision-javascript.yml) | floating-point precision | 待查证 |
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
| [garbage-collection-javascript](../entries/garbage-collection-javascript.yml) | garbage collection | 本轮查证后暂缓。Collins/Wiktionary 有计算机义，未见整短语 IPA；Cambridge 返回 403。 |
| [generic-constraint-typescript](../entries/generic-constraint-typescript.yml) | generic constraint | 待查证 |
| [global-attribute-html](../entries/global-attribute-html.yml) | global attribute | 待查证 |
| [global-augmentation-typescript](../entries/global-augmentation-typescript.yml) | global augmentation | 待查证 |
| [go-module](../entries/go-module.yml) | Go module | 待查证 |
| [goroutine-go](../entries/goroutine-go.yml) | goroutine | 本轮查证后暂缓。定向检索未找到可核实的、标口音的英语整词转写。 |
| [graceful-degradation-web](../entries/graceful-degradation-web.yml) | graceful degradation | 待查证 |
| [grid-area-css](../entries/grid-area-css.yml) | grid area | 待查证 |
| [grid-container-css](../entries/grid-container-css.yml) | grid container | 待查证 |
| [grid-layout-css](../entries/grid-layout-css.yml) | grid layout | 待查证 |
| [grid-system-bootstrap](../entries/grid-system-bootstrap.yml) | grid system | 待查证 |
| [grid-track-css](../entries/grid-track-css.yml) | grid track | 待查证 |
| [happens-before-go](../entries/happens-before-go.yml) | happens-before | 待查证 |
| [header-cell-html](../entries/header-cell-html.yml) | header cell | 待查证 |
| [heading-level-html](../entries/heading-level-html.yml) | heading level | 待查证 |
| [heap-memory-javascript](../entries/heap-memory-javascript.yml) | heap memory | 待查证 |
| [heat-resistance-combat](../entries/heat-resistance-combat.yml) | heat resistance | 待查证 |
| [higher-order-function-javascript](../entries/higher-order-function-javascript.yml) | higher-order function | 待查证 |
| [hoisting-javascript](../entries/hoisting-javascript.yml) | hoisting | 第七批查证后暂缓，本轮未重查。Collins 未取得可匹配的美式整词分区。 |
| [holdings-finance](../entries/holdings-finance.yml) | holdings | 拟补 B08-P06 |
| [hydration-react](../entries/hydration-react.yml) | hydration | 拟补 B08-P07 |
| [image-digest-docker](../entries/image-digest-docker.yml) | image digest | 待查证 |
| [image-layer-docker](../entries/image-layer-docker.yml) | image layer | 待查证 |
| [image-repository-docker](../entries/image-repository-docker.yml) | image repository | 待查证 |
| [image-tag-docker](../entries/image-tag-docker.yml) | image tag | 待查证 |
| [implicit-implementation-go](../entries/implicit-implementation-go.yml) | implicit implementation | 待查证 |
| [indexed-access-type-typescript](../entries/indexed-access-type-typescript.yml) | indexed access type | 待查证 |
| [initializer-tech](../entries/initializer-tech.yml) | initializer | 本轮查证后暂缓。Collins 未取得，Wiktionary 未见整词 IPA，Cambridge 检索未获得可用正文。 |
| [inline-box-css](../entries/inline-box-css.yml) | inline box | 待查证 |
| [inline-frame-html](../entries/inline-frame-html.yml) | inline frame | 待查证 |
| [input-group-bootstrap](../entries/input-group-bootstrap.yml) | input group | 待查证 |
| [input-type-html](../entries/input-type-html.yml) | input type | 待查证 |
| [interactive-content-html](../entries/interactive-content-html.yml) | interactive content | 待查证 |
| [internal-redirect-nginx](../entries/internal-redirect-nginx.yml) | internal redirect | 待查证 |
| [intersection-type-typescript](../entries/intersection-type-typescript.yml) | intersection type | 待查证 |
| [invalid-state-bootstrap](../entries/invalid-state-bootstrap.yml) | invalid state | 待查证 |
| [iterable-tech](../entries/iterable-tech.yml) | iterable | 本轮查证后暂缓。Wiktionary 可见的 IPA 属于西班牙语；Cambridge 检索未获得可用英语正文。 |
| [javascript-plugin-bootstrap](../entries/javascript-plugin-bootstrap.yml) | JavaScript plugin | 待查证 |
| [keyframe-css](../entries/keyframe-css.yml) | keyframe | 第七批查证后暂缓，本轮未重查。Collins 词页未能读取。 |
| [lazy-loading-html](../entries/lazy-loading-html.yml) | lazy loading | 待查证 |
| [least-recently-used-cache](../entries/least-recently-used-cache.yml) | least recently used | 待查证 |
| [lifting-state-up-react](../entries/lifting-state-up-react.yml) | lifting state up | 待查证 |
| [line-height-css](../entries/line-height-css.yml) | line height | 待查证 |
| [linked-list-merge-algorithm](../entries/linked-list-merge-algorithm.yml) | linked list merge | 待查证 |
| [linked-list-tech](../entries/linked-list-tech.yml) | linked list | 本轮查证后暂缓。Collins/Wiktionary 有计算机义，未见整短语 IPA。 |
| [list-group-bootstrap](../entries/list-group-bootstrap.yml) | list group | 待查证 |
| [list-item-html](../entries/list-item-html.yml) | list item | 待查证 |
| [list-rendering-react](../entries/list-rendering-react.yml) | list rendering | 待查证 |
| [literal-type-typescript](../entries/literal-type-typescript.yml) | literal type | 待查证 |
| [load-balancing-nginx](../entries/load-balancing-nginx.yml) | load balancing | 本轮查证后暂缓。定向检索未取得同义、完整的英语短语 IPA。 |
| [longest-prefix-nginx](../entries/longest-prefix-nginx.yml) | longest prefix | 待查证 |
| [longhand-property-css](../entries/longhand-property-css.yml) | longhand property | 待查证 |
| [main-axis-css](../entries/main-axis-css.yml) | main axis | 待查证 |
| [main-quest-game](../entries/main-quest-game.yml) | main quest | 待查证 |
| [male-to-male-hardware](../entries/male-to-male-hardware.yml) | male-to-male | 待查证 |
| [manual-save-game](../entries/manual-save-game.yml) | manual save | 待查证 |
| [mapped-type-typescript](../entries/mapped-type-typescript.yml) | mapped type | 待查证 |
| [margin-account-finance](../entries/margin-account-finance.yml) | margin account | 本轮查证后暂缓。Collins 美式分区有释义，未列整短语 IPA。Cambridge 返回 403。 |
| [margin-collapsing-css](../entries/margin-collapsing-css.yml) | margin collapsing | 待查证 |
| [market-capitalization-finance](../entries/market-capitalization-finance.yml) | market capitalization | 本轮查证后暂缓。Collins 有相应搭配和分项词音标，未获整短语转写。 |
| [media-query-css](../entries/media-query-css.yml) | media query | 待查证 |
| [method-set-go](../entries/method-set-go.yml) | method set | 待查证 |
| [microtask-javascript](../entries/microtask-javascript.yml) | microtask | 本轮查证后暂缓。Wiktionary 未见英语整词 IPA。 |
| [mine-cart-engineering](../entries/mine-cart-engineering.yml) | mine cart | 待查证 |
| [module-path-go](../entries/module-path-go.yml) | module path | 待查证 |
| [module-specifier-tech](../entries/module-specifier-tech.yml) | module specifier | 待查证 |
| [mount-point-docker](../entries/mount-point-docker.yml) | mount point | 待查证 |
| [multiple-return-values-go](../entries/multiple-return-values-go.yml) | multiple return values | 待查证 |
| [mutex-go](../entries/mutex-go.yml) | mutex | 本轮查证后暂缓。可见的 Wiktionary 结果未提供英语 IPA。 |
| [named-result-parameter-go](../entries/named-result-parameter-go.yml) | named result parameter | 待查证 |
| [namespace-typescript](../entries/namespace-typescript.yml) | namespace | 本轮查证后暂缓。Wiktionary 英语 IPA 未标口音；其他注明口音的相邻分区属于葡萄牙语。 |
| [navbar-bootstrap](../entries/navbar-bootstrap.yml) | navbar | 待查证 |
| [net-assets-finance](../entries/net-assets-finance.yml) | net assets | 待查证 |
| [net-income-finance](../entries/net-income-finance.yml) | net income | 本轮查证后暂缓。Collins 有财务释义及另一区的转写，但未为该整短语转写标明口音。Cambridge 返回 403。 |
| [normal-flow-css](../entries/normal-flow-css.yml) | normal flow | 待查证 |
| [offcanvas-bootstrap](../entries/offcanvas-bootstrap.yml) | offcanvas | 待查证 |
| [one-way-data-flow-frontend](../entries/one-way-data-flow-frontend.yml) | one-way data flow | 待查证 |
| [onion-model-koa](../entries/onion-model-koa.yml) | onion model | 待查证 |
| [opposite-shore-terrain](../entries/opposite-shore-terrain.yml) | opposite shore | 待查证 |
| [option-chain-finance](../entries/option-chain-finance.yml) | option chain | 待查证 |
| [optional-parameter-typescript](../entries/optional-parameter-typescript.yml) | optional parameter | 待查证 |
| [optional-property-typescript](../entries/optional-property-typescript.yml) | optional property | 待查证 |
| [orchestration-ai](../entries/orchestration-ai.yml) | orchestration | 拟补 B08-P08 |
| [ordered-list-html](../entries/ordered-list-html.yml) | ordered list | 待查证 |
| [own-property-javascript](../entries/own-property-javascript.yml) | own property | 待查证 |
| [partial-clone-git](../entries/partial-clone-git.yml) | partial clone | 待查证 |
| [persistent-data-docker](../entries/persistent-data-docker.yml) | persistent data | 待查证 |
| [pointer-receiver-go](../entries/pointer-receiver-go.yml) | pointer receiver | 待查证 |
| [popover-bootstrap](../entries/popover-bootstrap.yml) | popover | 第七批查证后暂缓，本轮未重查。Collins 找到食品义的音标，尚未核实 UI 义可用的转写。 |
| [port-publishing-docker](../entries/port-publishing-docker.yml) | port publishing | 待查证 |
| [postorder-traversal-tree](../entries/postorder-traversal-tree.yml) | postorder traversal | 待查证 |
| [pre-market-finance](../entries/pre-market-finance.yml) | pre-market | 本轮查证后暂缓。检索命中 market 押韵页及不同形式，未获精确带连字符形容词 IPA。 |
| [prefix-match-nginx](../entries/prefix-match-nginx.yml) | prefix match | 待查证 |
| [preflight-request-web](../entries/preflight-request-web.yml) | preflight request | 待查证 |
| [profit-and-loss-finance](../entries/profit-and-loss-finance.yml) | profit and loss | 待查证 |
| [progress-bar-bootstrap](../entries/progress-bar-bootstrap.yml) | progress bar | 待查证 |
| [progressive-enhancement-web](../entries/progressive-enhancement-web.yml) | progressive enhancement | 待查证 |
| [property-descriptor-javascript](../entries/property-descriptor-javascript.yml) | property descriptor | 待查证 |
| [property-value-css](../entries/property-value-css.yml) | property value | 待查证 |
| [props-react](../entries/props-react.yml) | props | 本轮查证后暂缓。Collins 命中其他义项；Wiktionary 给出单数 prop 或其他语言的音标。未据此补 React 义。 |
| [prototype-chain-javascript](../entries/prototype-chain-javascript.yml) | prototype chain | 待查证 |
| [prototype-setter-javascript](../entries/prototype-setter-javascript.yml) | prototype setter | 待查证 |
| [pseudo-class-css](../entries/pseudo-class-css.yml) | pseudo-class | 待查证 |
| [pseudo-element-css](../entries/pseudo-element-css.yml) | pseudo-element | 待查证 |
| [pure-component-react](../entries/pure-component-react.yml) | pure component | 待查证 |
| [quicksave-game](../entries/quicksave-game.yml) | quicksave | 拟补 B08-P11 |
| [race-detector-go](../entries/race-detector-go.yml) | race detector | 待查证 |
| [ranged-combat](../entries/ranged-combat.yml) | ranged | 第七批查证后暂缓，本轮未重查。Collins 美式音标所在分区是建筑义，暂不用于“远程”。 |
| [reactive-dependency-react](../entries/reactive-dependency-react.yml) | reactive dependency | 待查证 |
| [read-only-control-html](../entries/read-only-control-html.yml) | read-only control | 待查证 |
| [reduced-motion-css](../entries/reduced-motion-css.yml) | reduced motion | 待查证 |
| [ref-react](../entries/ref-react.yml) | ref | 待查证 |
| [regex-match-nginx](../entries/regex-match-nginx.yml) | regex match | 待查证 |
| [relative-positioning-css](../entries/relative-positioning-css.yml) | relative positioning | 待查证 |
| [relative-unit-css](../entries/relative-unit-css.yml) | relative unit | 待查证 |
| [required-field-html](../entries/required-field-html.yml) | required field | 待查证 |
| [rerender-react](../entries/rerender-react.yml) | re-render | 本轮查证后暂缓。检索命中自动生成转写，未核实带连字符的名词义词典发音。 |
| [responsive-container-bootstrap](../entries/responsive-container-bootstrap.yml) | responsive container | 待查证 |
| [responsive-design-css](../entries/responsive-design-css.yml) | responsive design | 待查证 |
| [responsive-image-html](../entries/responsive-image-html.yml) | responsive image | 待查证 |
| [responsive-utility-bootstrap](../entries/responsive-utility-bootstrap.yml) | responsive utility | 待查证 |
| [rest-parameter-typescript](../entries/rest-parameter-typescript.yml) | rest parameter | 待查证 |
| [return-type-tech](../entries/return-type-tech.yml) | return type | 待查证 |
| [reverse-proxy-nginx](../entries/reverse-proxy-nginx.yml) | reverse proxy | 待查证 |
| [row-span-html](../entries/row-span-html.yml) | row span | 待查证 |
| [rules-of-hooks-react](../entries/rules-of-hooks-react.yml) | Rules of Hooks | 待查证 |
| [same-origin-policy-web](../entries/same-origin-policy-web.yml) | same-origin policy | 待查证 |
| [sass-mixin-bootstrap](../entries/sass-mixin-bootstrap.yml) | Sass mixin | 待查证 |
| [sass-variable-bootstrap](../entries/sass-variable-bootstrap.yml) | Sass variable | 待查证 |
| [semantic-versioning-software](../entries/semantic-versioning-software.yml) | semantic versioning | 待查证 |
| [sentinel-error-go](../entries/sentinel-error-go.yml) | sentinel error | 待查证 |
| [server-rendering-react](../entries/server-rendering-react.yml) | server rendering | 待查证 |
| [shallow-copy-javascript](../entries/shallow-copy-javascript.yml) | shallow copy | 待查证 |
| [shock-resistance-combat](../entries/shock-resistance-combat.yml) | shock resistance | 待查证 |
| [short-variable-declaration-go](../entries/short-variable-declaration-go.yml) | short variable declaration | 待查证 |
| [shorthand-property-css](../entries/shorthand-property-css.yml) | shorthand property | 待查证 |
| [side-hop-combat](../entries/side-hop-combat.yml) | side hop | 待查证 |
| [sky-island-zelda](../entries/sky-island-zelda.yml) | sky island | 待查证 |
| [soundness-type-system](../entries/soundness-type-system.yml) | soundness | 拟补 B08-P09 |
| [source-order-css](../entries/source-order-css.yml) | source order | 待查证 |
| [source-set-html](../entries/source-set-html.yml) | source set | 待查证 |
| [spacing-utility-bootstrap](../entries/spacing-utility-bootstrap.yml) | spacing utility | 待查证 |
| [stack-trace-tech](../entries/stack-trace-tech.yml) | stack trace | 本轮查证后暂缓。Collins 未取得，Wiktionary 未见整短语 IPA。 |
| [stacking-context-css](../entries/stacking-context-css.yml) | stacking context | 待查证 |
| [start-tag-html](../entries/start-tag-html.yml) | start tag | 待查证 |
| [state-preservation-react](../entries/state-preservation-react.yml) | state preservation | 待查证 |
| [state-setter-react](../entries/state-setter-react.yml) | state setter | 待查证 |
| [state-snapshot-react](../entries/state-snapshot-react.yml) | state snapshot | 待查证 |
| [sticky-positioning-css](../entries/sticky-positioning-css.yml) | sticky positioning | 待查证 |
| [storage-event-browser](../entries/storage-event-browser.yml) | storage event | 待查证 |
| [strict-mode-react](../entries/strict-mode-react.yml) | Strict Mode | 待查证 |
| [strike-price-finance](../entries/strike-price-finance.yml) | strike price | 待查证 |
| [structural-check-tech](../entries/structural-check-tech.yml) | structural check | 待查证 |
| [structural-typing-typescript](../entries/structural-typing-typescript.yml) | structural typing | 待查证 |
| [style-rule-css](../entries/style-rule-css.yml) | style rule | 待查证 |
| [stylesheet-css](../entries/stylesheet-css.yml) | stylesheet | 本轮查证后暂缓。Wiktionary IPA 未标口音。Collins 的访问限制见第七批。 |
| [subtest-go](../entries/subtest-go.yml) | subtest | 待查证 |
| [symlink-tech](../entries/symlink-tech.yml) | symlink | 本轮查证后暂缓。Wiktionary 未见标明口音的英语 IPA。 |
| [system-prompt-ai](../entries/system-prompt-ai.yml) | system prompt | 待查证 |
| [table-driven-test-go](../entries/table-driven-test-go.yml) | table-driven test | 待查证 |
| [table-row-html](../entries/table-row-html.yml) | table row | 待查证 |
| [template-literal-type-typescript](../entries/template-literal-type-typescript.yml) | template literal type | 待查证 |
| [temporal-dead-zone-javascript](../entries/temporal-dead-zone-javascript.yml) | temporal dead zone | 待查证 |
| [text-wrapping-css](../entries/text-wrapping-css.yml) | text wrapping | 待查证 |
| [theming-bootstrap](../entries/theming-bootstrap.yml) | theming | 待查证 |
| [toggle-trigger-bootstrap](../entries/toggle-trigger-bootstrap.yml) | toggle trigger | 待查证 |
| [tool-call-ai](../entries/tool-call-ai.yml) | tool call | 待查证 |
| [tool-result-ai](../entries/tool-result-ai.yml) | tool result | 待查证 |
| [tree-shaking-build](../entries/tree-shaking-build.yml) | tree shaking | 待查证 |
| [two-way-binding-frontend](../entries/two-way-binding-frontend.yml) | two-way binding | 待查证 |
| [type-alias-typescript](../entries/type-alias-typescript.yml) | type alias | 待查证 |
| [type-annotation-typescript](../entries/type-annotation-typescript.yml) | type annotation | 待查证 |
| [type-constraint-go](../entries/type-constraint-go.yml) | type constraint | 待查证 |
| [type-guard-typescript](../entries/type-guard-typescript.yml) | type guard | 待查证 |
| [type-inference-typescript](../entries/type-inference-typescript.yml) | type inference | 待查证 |
| [type-instantiation-tech](../entries/type-instantiation-tech.yml) | type instantiation | 待查证 |
| [type-only-import-typescript](../entries/type-only-import-typescript.yml) | type-only import | 待查证 |
| [type-parameter-typescript](../entries/type-parameter-typescript.yml) | type parameter | 待查证 |
| [type-predicate-typescript](../entries/type-predicate-typescript.yml) | type predicate | 待查证 |
| [type-selector-css](../entries/type-selector-css.yml) | type selector | 待查证 |
| [type-switch-go](../entries/type-switch-go.yml) | type switch | 待查证 |
| [typeof-type-operator-typescript](../entries/typeof-type-operator-typescript.yml) | typeof type operator | 待查证 |
| [uncontrolled-component-react](../entries/uncontrolled-component-react.yml) | uncontrolled component | 待查证 |
| [unfuse-general](../entries/unfuse-general.yml) | unfuse | 待查证 |
| [unordered-list-html](../entries/unordered-list-html.yml) | unordered list | 待查证 |
| [updater-function-react](../entries/updater-function-react.yml) | updater function | 待查证 |
| [utility-class-bootstrap](../entries/utility-class-bootstrap.yml) | utility class | 待查证 |
| [utility-type-typescript](../entries/utility-type-typescript.yml) | utility type | 待查证 |
| [valid-state-bootstrap](../entries/valid-state-bootstrap.yml) | valid state | 待查证 |
| [validation-feedback-bootstrap](../entries/validation-feedback-bootstrap.yml) | validation feedback | 待查证 |
| [viewport-css](../entries/viewport-css.yml) | viewport | 本轮查证后暂缓。Wiktionary IPA 未标口音。Collins 的访问限制见第七批。 |
| [visually-hidden-bootstrap](../entries/visually-hidden-bootstrap.yml) | visually hidden | 待查证 |
| [void-element-html](../entries/void-element-html.yml) | void element | 待查证 |
| [waistguard-clothing](../entries/waistguard-clothing.yml) | waistguard | 待查证 |
| [wait-group-go](../entries/wait-group-go.yml) | wait group | 待查证 |
| [web-storage-browser](../entries/web-storage-browser.yml) | web storage | 待查证 |
| [withdrawable-finance](../entries/withdrawable-finance.yml) | withdrawable | 本轮查证后暂缓。Collins 未取得、Wiktionary 未见整词 IPA。现有词根发音不能与后缀相拼。 |
| [worker-pool-concurrency](../entries/worker-pool-concurrency.yml) | worker pool | 待查证 |
| [year-over-year-finance](../entries/year-over-year-finance.yml) | year-over-year | 本轮查证后暂缓。检索命中 year 单词和 YoY 缩写的转写，未获完整形容词短语 IPA。 |
| [zero-value-go](../entries/zero-value-go.yml) | zero value | 待查证 |

### 03 缩写及技术标识符

| 稳定 ID | 当前 Word | 状态及依据 |
|:--|:--|:--|
| [bigint-javascript](../entries/bigint-javascript.yml) | BigInt | 待查证 |
| [bool-go](../entries/bool-go.yml) | bool | 本轮查证后暂缓。Wiktionary 命中的相邻语种 IPA 不适用。保留 Go 标识符，待查其口读约定。 |
| [comparable-go](../entries/comparable-go.yml) | comparable | 本轮查证后暂缓。Collins 有英语形容词读音。Go 预声明约束需要确认标识符的口读约定。 |
| [compose-yaml-docker](../entries/compose-yaml-docker.yml) | compose.yaml | 待查证 |
| [dom-html](../entries/dom-html.yml) | DOM | 待查证 |
| [ebit-finance](../entries/ebit-finance.yml) | EBIT | 待查证 |
| [fallback-ui-react](../entries/fallback-ui-react.yml) | fallback UI | 待查证 |
| [go-mod-go](../entries/go-mod-go.yml) | go.mod | 待查证 |
| [go-sum-go](../entries/go-sum-go.yml) | go.sum | 待查证 |
| [hdmi-cable-hardware](../entries/hdmi-cable-hardware.yml) | HDMI cable | 待查证 |
| [hdmi-type-a-connector](../entries/hdmi-type-a-connector.yml) | HDMI Type A connector | 待查证 |
| [http-caching-web](../entries/http-caching-web.yml) | HTTP caching | 待查证 |
| [id-selector-css](../entries/id-selector-css.yml) | ID selector | 待查证 |
| [int-go](../entries/int-go.yml) | int | 本轮查证后暂缓。Wiktionary 命中的相邻语种 IPA 不适用。保留 Go 标识符，待查其口读约定。 |
| [jsx-react](../entries/jsx-react.yml) | JSX | 待查证 |
| [keyof-typescript](../entries/keyof-typescript.yml) | keyof | 待查证 |
| [new-go](../entries/new-go.yml) | new | 本轮查证后暂缓。Collins 美式音标为普通英语词。Go 内置函数口读惯例待确认。 |
| [omit-utility-type-typescript](../entries/omit-utility-type-typescript.yml) | Omit | 本轮查证后暂缓。Collins 提供普通动词 IPA。TypeScript 标识符的读法需要单独确认。 |
| [partial-utility-type-typescript](../entries/partial-utility-type-typescript.yml) | Partial | 本轮查证后暂缓。Collins 美式为形容词。当前为 TypeScript 工具类型名，待核实标识符口读依据。 |
| [pick-utility-type-typescript](../entries/pick-utility-type-typescript.yml) | Pick | 本轮查证后暂缓。Collins 提供普通动词 IPA。TypeScript 标识符的读法需要单独确认。 |
| [promise-all-javascript](../entries/promise-all-javascript.yml) | Promise.all | 待查证 |
| [property-css](../entries/property-css.yml) | CSS property | 待查证 |
| [readonly-typescript](../entries/readonly-typescript.yml) | readonly | 待查证 |
| [record-utility-type-typescript](../entries/record-utility-type-typescript.yml) | Record | 本轮查证后暂缓。Collins 区分 record 的名词和动词重音。当前为 TypeScript 工具类型，待核实口读依据。 |
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
| [ancient-arowana-zelda](../entries/ancient-arowana-zelda.yml) | Ancient Arowana | 本轮查证后暂缓。与普通鱼名一起筛查。相近英文或芬兰语拼写无法核实该游戏材料的全名读法。 |
| [attack-up-zelda](../entries/attack-up-zelda.yml) | Attack Up | 待查证 |
| [autobuild-zelda](../entries/autobuild-zelda.yml) | Autobuild | 待查证 |
| [beam-emitter-zelda](../entries/beam-emitter-zelda.yml) | Beam Emitter | 待查证 |
| [big-battery-zelda](../entries/big-battery-zelda.yml) | Big Battery | 待查证 |
| [big-wheel-zelda](../entries/big-wheel-zelda.yml) | Big Wheel | 待查证 |
| [bokoblin-zelda](../entries/bokoblin-zelda.yml) | Bokoblin | 待查证 |
| [bootstrap-framework](../entries/bootstrap-framework.yml) | Bootstrap | 本轮查证后暂缓。Collins 有普通名词及计算机动词等读法。该词条为框架专名，待查项目名称读法。 |
| [boss-bokoblin-zelda](../entries/boss-bokoblin-zelda.yml) | Boss Bokoblin | 待查证 |
| [brightbloom-seed-zelda](../entries/brightbloom-seed-zelda.yml) | Brightbloom Seed | 待查证 |
| [brightcap-zelda](../entries/brightcap-zelda.yml) | Brightcap | 待查证 |
| [captain-construct-zelda](../entries/captain-construct-zelda.yml) | Captain Construct | 待查证 |
| [character-profiles-zelda](../entries/character-profiles-zelda.yml) | Character Profiles | 待查证 |
| [chuchu-zelda](../entries/chuchu-zelda.yml) | Chuchu | 待查证 |
| [claude-ai](../entries/claude-ai.yml) | Claude | 待查证 |
| [climbing-gear-zelda](../entries/climbing-gear-zelda.yml) | Climbing Gear | 待查证 |
| [codex-ai](../entries/codex-ai.yml) | Codex | 待查证 |
| [colgera-zelda](../entries/colgera-zelda.yml) | Colgera | 待查证 |
| [construct-factory-zelda](../entries/construct-factory-zelda.yml) | Construct Factory | 待查证 |
| [construct-head-zelda](../entries/construct-head-zelda.yml) | Construct Head | 待查证 |
| [copilot-ai](../entries/copilot-ai.yml) | Copilot | 待查证 |
| [courser-bee-honey-zelda](../entries/courser-bee-honey-zelda.yml) | Courser Bee Honey | 待查证 |
| [crystal-refinery-zelda](../entries/crystal-refinery-zelda.yml) | Crystal Refinery | 待查证 |
| [crystallized-charge-zelda](../entries/crystallized-charge-zelda.yml) | Crystallized Charge | 待查证 |
| [cursor-ai](../entries/cursor-ai.yml) | Cursor | 待查证 |
| [death-mountain-zelda](../entries/death-mountain-zelda.yml) | Death Mountain | 待查证 |
| [deepseek-ai](../entries/deepseek-ai.yml) | DeepSeek | 待查证 |
| [defense-up-zelda](../entries/defense-up-zelda.yml) | Defense Up | 待查证 |
| [demon-dragon-zelda](../entries/demon-dragon-zelda.yml) | Demon Dragon | 待查证 |
| [demon-king-ganondorf-zelda](../entries/demon-king-ganondorf-zelda.yml) | Demon King Ganondorf | 待查证 |
| [demon-king-zelda](../entries/demon-king-zelda.yml) | Demon King | 待查证 |
| [demon-kings-bow-zelda](../entries/demon-kings-bow-zelda.yml) | Demon King's Bow | 待查证 |
| [device-dispenser-zelda](../entries/device-dispenser-zelda.yml) | Device Dispenser | 待查证 |
| [docker-compose-docker](../entries/docker-compose-docker.yml) | Docker Compose | 待查证 |
| [docker-platform](../entries/docker-platform.yml) | Docker | 待查证 |
| [dots-ai](../entries/dots-ai.yml) | dots | 待查证 |
| [electro-cooking](../entries/electro-cooking.yml) | Electro | 待查证 |
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
| [ganondorf-zelda](../entries/ganondorf-zelda.yml) | Ganondorf | 待查证 |
| [gemini-ai](../entries/gemini-ai.yml) | Gemini | 待查证 |
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
| [grok-ai](../entries/grok-ai.yml) | Grok | 待查证 |
| [hateno-village-zelda](../entries/hateno-village-zelda.yml) | Hateno Village | 待查证 |
| [heart-container-zelda](../entries/heart-container-zelda.yml) | Heart Container | 待查证 |
| [hearty-radish-zelda](../entries/hearty-radish-zelda.yml) | Hearty Radish | 待查证 |
| [hearty-salmon-zelda](../entries/hearty-salmon-zelda.yml) | Hearty Salmon | 待查证 |
| [hearty-truffle-zelda](../entries/hearty-truffle-zelda.yml) | Hearty Truffle | 待查证 |
| [hermes-agent-ai](../entries/hermes-agent-ai.yml) | Hermes | 待查证 |
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
| [hyrule-zelda](../entries/hyrule-zelda.yml) | Hyrule | 待查证 |
| [impa-zelda](../entries/impa-zelda.yml) | Impa | 待查证 |
| [instinct-pet-food](../entries/instinct-pet-food.yml) | Instinct | 待查证 |
| [josha-zelda](../entries/josha-zelda.yml) | Josha | 待查证 |
| [kakariko-village-zelda](../entries/kakariko-village-zelda.yml) | Kakariko Village | 待查证 |
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
| [muse-ai](../entries/muse-ai.yml) | Muse | 待查证 |
| [myfoodie-pet-food](../entries/myfoodie-pet-food.yml) | Myfoodie | 待查证 |
| [nginx-platform](../entries/nginx-platform.yml) | NGINX | 待查证 |
| [nutro-pet-food](../entries/nutro-pet-food.yml) | Nutro | 待查证 |
| [octorok-zelda](../entries/octorok-zelda.yml) | Octorok | 待查证 |
| [openclaw-ai](../entries/openclaw-ai.yml) | OpenClaw | 待查证 |
| [orijen-pet-food](../entries/orijen-pet-food.yml) | ORIJEN | 待查证 |
| [paya-zelda](../entries/paya-zelda.yml) | Paya | 待查证 |
| [perfect-dodge-zelda](../entries/perfect-dodge-zelda.yml) | Perfect Dodge | 待查证 |
| [perfect-guard-zelda](../entries/perfect-guard-zelda.yml) | Perfect Guard | 待查证 |
| [phantom-ganon-zelda](../entries/phantom-ganon-zelda.yml) | Phantom Ganon | 待查证 |
| [portable-pot-zelda](../entries/portable-pot-zelda.yml) | Portable Pot | 待查证 |
| [princess-zelda-zelda](../entries/princess-zelda-zelda.yml) | Princess Zelda | 待查证 |
| [puffshroom-zelda](../entries/puffshroom-zelda.yml) | Puffshroom | 待查证 |
| [purah-pad-zelda](../entries/purah-pad-zelda.yml) | Purah Pad | 待查证 |
| [purah-zelda](../entries/purah-zelda.yml) | Purah | 待查证 |
| [purina-pro-plan-pet-food](../entries/purina-pro-plan-pet-food.yml) | Purina Pro Plan | 待查证 |
| [queen-gibdo-zelda](../entries/queen-gibdo-zelda.yml) | Queen Gibdo | 待查证 |
| [rauru-zelda](../entries/rauru-zelda.yml) | Rauru | 待查证 |
| [raw-gourmet-meat-zelda](../entries/raw-gourmet-meat-zelda.yml) | Raw Gourmet Meat | 待查证 |
| [raw-prime-meat-zelda](../entries/raw-prime-meat-zelda.yml) | Raw Prime Meat | 待查证 |
| [react-library](../entries/react-library.yml) | React | 待查证 |
| [regional-phenomena-zelda](../entries/regional-phenomena-zelda.yml) | Regional Phenomena | 待查证 |
| [riju-zelda](../entries/riju-zelda.yml) | Riju | 待查证 |
| [rito-village-zelda](../entries/rito-village-zelda.yml) | Rito Village | 待查证 |
| [rito-zelda](../entries/rito-zelda.yml) | Rito | 待查证 |
| [robbie-zelda](../entries/robbie-zelda.yml) | Robbie | 待查证 |
| [royal-broadsword-zelda](../entries/royal-broadsword-zelda.yml) | Royal Broadsword | 待查证 |
| [royal-canin-pet-food](../entries/royal-canin-pet-food.yml) | Royal Canin | 待查证 |
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
| [whiskas-pet-food](../entries/whiskas-pet-food.yml) | Whiskas | 待查证 |
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

1. 重新查看 Git 状态及已批准的稳定 ID。核对词形、词义、词性、现有读音和 Notes；如有后来改动导致文案失配，保留当前数据并报告。快照原始 YAML、已有导出和媒体文件哈希。
2. 只写已批准的 `pronunciation` 值与来源行。去重并保留原 Notes；保留其余字段、现有录音字节、语种与语音设置。若快照中的导出缺失或落后于 YAML，用实施前同一基线的隔离构建取得可比导出，再核验顺序与字段差异。
3. 从仓库根目录运行以下命令。单独检查新文件的格式和 Git 差异检查未覆盖的内容。

```bash
go run scripts/validate/main.go
go run scripts/build/main.go
git diff --check
```

4. 对比源文件和导出。按稳定 ID 检查核准读音、追加来源、ID 集合、行数、字段顺序、Audio、StudyOrder 及其余无关列；复核媒体哈希。将任何先于本批存在的导出漂移与本批改动分别记录。
5. 复查剩余缺口，回填批准、实施及实际验证结果。人工听审、真实 Anki 导入和设备播放各自记录。将导出核验作为静态结果。

## 规划验证

稳定 ID、语言计数、审阅编号和本地链接均在当前工作区核验。对照规划前哈希检查原有文件；规划期间未运行验证、构建或音频命令。已收到明确实施批准。

## 实施结果

- 按 B08-P01–B08-P11 更新 11 个 YAML 文件，添加 11 项书面 IPA 与 11 行 Notes 来源。保留原 Notes 内容和其余字段。
- `go run scripts/validate/main.go` 通过，校验 1,833 项。报告 38 条同词形警告。两次 `go run scripts/build/main.go` 生成字节一致的导出。`git diff --check` 通过，新计划的格式和本地链接单独核验。
- 对比 1,833 行、18 列导出，确认只有获批条目的 Pronunciation 与 Notes 列有变化。条目 ID、Audio、StudyOrder 和其余列一致。1,875 个 MP3 的哈希均与快照一致。
- 英语已有书面读音 1,018 项，尚缺 572 项。日语 243 项均有读法。未做人工听审或真实 Anki 导入和播放验证。
