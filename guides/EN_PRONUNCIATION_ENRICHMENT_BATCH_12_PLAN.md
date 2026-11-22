# 书面读音全库审查与第十二批音标计划

状态为已实施。规划基线日期为 2026-11-22。用户批准全部提案，已实施 B12-P01 和 B12-P02。[第十一批](EN_PRONUNCIATION_ENRICHMENT_BATCH_11_PLAN.md)保留实施记录。

## 全库盘点

以当前 YAML 的 `id` 和 `language` 为准。空值、空白文本及未提供 `pronunciation` 均计为缺失。已读 [AGENTS.md](../AGENTS.md)、[README.md](../README.md) 和[词条规范](ENTRY_SCHEMA.md#pronunciation)。盘点包含工作区已有修改。

| 语言 | 总数 | 已有读音 | 缺失 |
|:--|--:|--:|--:|
| `en` | 1,590 | 1,038 | 552 |
| `ja` | 243 | 243 | 0 |
| 合计 | 1,833 | 1,281 | 552 |

本轮对 56 个缺失 ID 逐项查证，提出 2 项（美式 2 项、英式 0 项），54 项查证后暂缓。另有 104 项沿用[第十一批](EN_PRONUNCIATION_ENRICHMENT_BATCH_11_PLAN.md)或[王国之泪批次](TOTK_PRONUNCIATION_ENRICHMENT_PLAN.md)的逐项查证记录，本轮未重查。剩余 392 项待查证，其中包含此前只做过内容筛查、尚未逐名追查读音的游戏专名。四类合计 552 项。实施后缺失为 550 项。

### 已填读音的疑点

对英语现值核对斜线、口音前缀和 Notes 中明确标注的 US／UK 发音来源。对日语现值核对是否夹有汉字、拉丁字母或阿拉伯数字。未发现这些类别的异常。现有字段没有拟改项。本轮未逐项重新审定已填音标或假名读法。

## 批次与证据

先复核常用词和日常短语，再检查界面、开发与金融术语及专名。下方清单保留缩写、技术标识符、符号与命令的独立队列。证据采用词典的可读正文，只有音频、重拼法或相邻释义时继续暂缓。

美式 IPA 无前缀。本批没有英式回退。来源保留其完整转写规范。源页访问受限只记录为查证限制。普通名词用作游戏名称时，核对拼写、词性及道具的字面含义；专名或用法不同的技术标识符仍需对应证据。题名、句型及专名保留全名，单独词头的发音不能支持整式。

以下 Notes 来源行是逐字拟追加内容。整批批准仅覆盖这些编号中的 IPA 和来源行。

### 批次 A · 普通词组在游戏中的用法

| 审阅 ID | 稳定 ID / Word | 当前词性 | 当前 Meaning | 拟写 `pronunciation` | 适用性与证据 |
|:--|:--|:--|:--|:--|:--|
| B12-P01 | [time-bomb-zelda](../entries/time-bomb-zelda.yml) / Time Bomb | noun | 计时炸弹；《王国之泪》中的左纳乌装置。 | `/ˈtaɪm ˌbɑːm/` | 使用 Cambridge 美式整词组。对应该名词的实体炸弹义，适用于同一普通词组用作游戏道具名。 [Cambridge Dictionary — time bomb](https://dictionary.cambridge.org/dictionary/english/time-bomb) |

### 批次 B · 软件界面词组

| 审阅 ID | 稳定 ID / Word | 当前词性 | 当前 Meaning | 拟写 `pronunciation` | 适用性与证据 |
|:--|:--|:--|:--|:--|:--|
| B12-P02 | [progress-bar-bootstrap](../entries/progress-bar-bootstrap.yml) / progress bar | noun phrase | 进度条，用条形图表示任务完成进度的组件 | `/ˈprɒg rɛs ˌbɑr/` | American 栏的第一条完整转写，保留其元音、空格和重音标记。释义直接对应软件进度条；同栏第二个读法以缩略形式列出。 [Dictionary.com Unabridged — progress bar](https://www.dictionary.com/browse/progress-bar) |

### 拟追加的 Notes 来源行

- B12-P01 `Pronunciation source: Cambridge Dictionary — time bomb. https://dictionary.cambridge.org/dictionary/english/time-bomb (accent: US; accessed 2026-11-22).`
- B12-P02 `Pronunciation source: Dictionary.com Unabridged — progress bar. https://www.dictionary.com/browse/progress-bar (accent: US; accessed 2026-11-22).`

### 本轮主要证据与限制

| 来源 | 可核对内容与限制 |
|:--|:--|
| [Merriam-Webster — arowana](https://www.merriam-webster.com/dictionary/arowana) | 鱼类词义与重拼发音。 |
| [新综合英俄词典 — greaves](https://slovar-vocab.com/english-russian/new-big-dictionary/greaves-3974437.html) · [English to Chin Hakha Dictionary — G](https://chin-dictionary.com/assets/07-Gg_final.pdf#page=49) | 护胫义的复数读音；未核实 IPA 的默认口音。 |
| [Collins — drop-down](https://www.collinsdictionary.com/dictionary/english/drop-down) · [Cambridge — drop-down menu](https://dictionary.cambridge.org/pronunciation/english/drop-down-menu) | 前者为计算机形容词及其异形，后者为更长的名词短语。两个词头都与当前独立名词有差别。 |
| [Cambridge — popover](https://dictionary.cambridge.org/dictionary/english/popover) | 可读词条的词义是烘焙食品。 |
| [Collins — garbage collection](https://www.collinsdictionary.com/dictionary/english/garbage-collection) | 计算机词义可核对，正文未找到整词组 IPA。 |
| [Study4 股市词表](https://study4.com/posts/2458/tu-vung-tieng-anh-chu-de-chung-khoan/) · [NetEnglish 经济术语词表](https://netenglish.edu.vn/tieng-anh-kinh-te-tu-vung-thuat-ngu-mau-cau/) | 分别可核对 ask price 和 strike price 的词义及完整转写，但未核实口音。 |
| [Zelda Wiki 用户发音草稿](https://zeldawiki.wiki/wiki/User%3ASnorlaxMonster/Pronunciation) | 访问受限。本轮没有把该草稿当作已核实证据。 |

## 缺失 ID 清单

每个缺失稳定 ID 只列一次。队列表示下一轮的主要查证方向。待查证项保留当前拼写和用法，专名优先核实实际所指，标识符和符号优先寻找已出版的口读依据。

| 队列 | 缺失 | 拟补 | 本轮暂缓 | 前批暂缓 | 待查证 |
|:--|--:|--:|--:|--:|--:|
| 01 常用词与短语 | 12 | 0 | 7 | 5 | 0 |
| 02 技术、金融与游戏通用语 | 287 | 1 | 20 | 72 | 194 |
| 03 缩写及技术标识符 | 31 | 0 | 2 | 12 | 17 |
| 04 专名及游戏名称 | 210 | 1 | 24 | 14 | 171 |
| 05 语法模板、符号与命令 | 12 | 0 | 1 | 1 | 10 |

### 01 常用词与短语

| 稳定 ID | 当前 Word | 状态及依据 |
|:--|:--|:--|
| [arowana-general](../entries/arowana-general.yml) | arowana | 本轮查证后暂缓。Merriam-Webster 的鱼类名词词条能确认义项，但可见发音采用重拼法。查到的音频及其他语言同形词不提供可用的英语 IPA。 |
| [cooking-pot-cooking](../entries/cooking-pot-cooking.yml) | cooking pot | 本轮查证后暂缓。双语词典结果未标口音。查得 Longman 图片词典，但整份 PDF 超出读取限制，未核实对应词条的整式音标。 |
| [cross-functional-collaboration-workflow](../entries/cross-functional-collaboration-workflow.yml) | cross-functional collaboration | 前批查证后暂缓，本轮未重查。复查命中工作场景用例，未核实整短语。 |
| [greaves-clothing](../entries/greaves-clothing.yml) | greaves | 本轮查证后暂缓。《新综合英俄词典》和 David Van Bik 的《English to Chin Hakha Dictionary》在历史护胫复数义下均列 gri:vz。两条可见词条都未标口音，本轮未核实适用的语音说明。Cambridge 同形词为专名。 |
| [lex-name-element](../entries/lex-name-element.yml) | lex | 前批查证后暂缓，本轮未重查。检索出现同名产品或独立词典名词，仍缺当前项目命名片段的口读依据。 |
| [mine-cart-engineering](../entries/mine-cart-engineering.yml) | mine cart | 本轮查证后暂缓。本轮检索未找到矿车名词词组的标口音书面整式。相近英文拼写及游戏用例不能提供整式 IPA。 |
| [opposite-shore-terrain](../entries/opposite-shore-terrain.yml) | opposite shore | 前批查证后暂缓，本轮未重查。Cambridge 与 Collins 命中搭配用例，未列全短语音标。 |
| [project-onboarding-workflow](../entries/project-onboarding-workflow.yml) | project onboarding | 前批查证后暂缓，本轮未重查。复查命中项目说明和成分词，未核实整短语。 |
| [ranged-combat](../entries/ranged-combat.yml) | ranged | 本轮查证后暂缓。Cambridge 的 range 词条有动词屈折形式及名词义。未找到覆盖远程攻击形容词义的标口音完整转写。 |
| [related-words-language](../entries/related-words-language.yml) | related words | 前批查证后暂缓，本轮未重查。检索多为单词 related 的 IPA，未核实完整名词短语。 |
| [unfuse-general](../entries/unfuse-general.yml) | unfuse | 本轮查证后暂缓。未在本轮查询的英语词典中找到对应解除组合义的标口音转写。保留词头原样。 |
| [waistguard-clothing](../entries/waistguard-clothing.yml) | waistguard | 本轮查证后暂缓。本轮没有找到对应护腰义、当前合写词形的标口音 IPA。 |

### 02 技术、金融与游戏通用语

| 稳定 ID | 当前 Word | 状态及依据 |
|:--|:--|:--|
| [absolute-positioning-css](../entries/absolute-positioning-css.yml) | absolute positioning | 待查证 |
| [abstract-syntax-tree-compiler](../entries/abstract-syntax-tree-compiler.yml) | abstract syntax tree | 前批查证后暂缓，本轮未重查。计算机词典定向请求超时，未核实完整短语读法。 |
| [accessible-name-html](../entries/accessible-name-html.yml) | accessible name | 前批查证后暂缓，本轮未重查。检索命中无障碍文档与用例，未查得整词组音标。 |
| [accessor-tech](../entries/accessor-tech.yml) | accessor | 前批查证后暂缓，本轮未重查。计算机词典第 4 页的 accessor 指访问数据的人。该页 UK 读音的义项未覆盖当前的属性访问器。 |
| [alternative-text-html](../entries/alternative-text-html.yml) | alternative text | 前批查证后暂缓，本轮未重查。Cambridge 的近邻是缩写 alt text，未取得当前完整词形的读音。 |
| [annual-percentage-yield-finance](../entries/annual-percentage-yield-finance.yml) | annual percentage yield | 本轮查证后暂缓。查到释义、术语使用和音频发音结果。未核实与词头一致的标口音整式。 |
| [array-hole-javascript](../entries/array-hole-javascript.yml) | array hole | 待查证 |
| [ask-price-finance](../entries/ask-price-finance.yml) | ask price | 本轮查证后暂缓。Study4 股市词表能核对卖方报价义及完整转写 /æsk praɪs/，可见词条未标口音。未独立核实带口音的整式。 |
| [assignability-typescript](../entries/assignability-typescript.yml) | assignability | 前批查证后暂缓，本轮未重查。Wiktionary 可读页面未列整词 IPA，仍需完整派生词读法依据。 |
| [at-rule-css](../entries/at-rule-css.yml) | at-rule | 前批查证后暂缓，本轮未重查。Wiktionary 有术语义但未见 IPA。 |
| [attack-power-combat](../entries/attack-power-combat.yml) | attack power | 前批查证后暂缓，本轮未重查。Collins、Cambridge 未取得匹配的整短语发音，检索结果只支持单词或用例。 |
| [attribute-selector-css](../entries/attribute-selector-css.yml) | attribute selector | 前批查证后暂缓，本轮未重查。Collins 与 Wiktionary 未取得准确短语页面。组件读音不足以证实全词组。 |
| [backend-server-nginx](../entries/backend-server-nginx.yml) | backend server | 待查证 |
| [backing-array-go](../entries/backing-array-go.yml) | backing array | 待查证 |
| [base-image-docker](../entries/base-image-docker.yml) | base image | 待查证 |
| [batching-react](../entries/batching-react.yml) | batching | 前批查证后暂缓，本轮未重查。Wiktionary 提供完整词形 IPA，但未注明口音及 React 用义。 |
| [beat-estimates-finance](../entries/beat-estimates-finance.yml) | beat estimates | 待查证 |
| [bind-mount-docker](../entries/bind-mount-docker.yml) | bind mount | 待查证 |
| [block-formatting-context-css](../entries/block-formatting-context-css.yml) | block formatting context | 待查证 |
| [block-level-box-css](../entries/block-level-box-css.yml) | block-level box | 待查证 |
| [boolean-attribute-html](../entries/boolean-attribute-html.yml) | boolean attribute | 待查证 |
| [box-model-css](../entries/box-model-css.yml) | box model | 前批查证后暂缓，本轮未重查。查到同名技术词条以及不同词组 black box model，未确认 CSS 名词短语有标口音书面值。 |
| [buffered-channel-go](../entries/buffered-channel-go.yml) | buffered channel | 待查证 |
| [build-cache-docker](../entries/build-cache-docker.yml) | build cache | 待查证 |
| [build-context-docker](../entries/build-context-docker.yml) | build context | 待查证 |
| [button-group-bootstrap](../entries/button-group-bootstrap.yml) | button group | 待查证 |
| [cache-invalidation-tech](../entries/cache-invalidation-tech.yml) | cache invalidation | 前批查证后暂缓，本轮未重查。沿用第八批清单。本轮查证后暂缓。Cambridge 定向查询未取得匹配的整短语音标。 |
| [cache-revalidation-web](../entries/cache-revalidation-web.yml) | cache revalidation | 待查证 |
| [call-stack-javascript](../entries/call-stack-javascript.yml) | call stack | 前批查证后暂缓，本轮未重查。翻译词典及聚合条目匹配术语，未取得注明口音且可追溯的整短语 IPA。 |
| [cascade-layer-css](../entries/cascade-layer-css.yml) | cascade layer | 待查证 |
| [cash-sweep-finance](../entries/cash-sweep-finance.yml) | cash sweep | 本轮查证后暂缓。Investor.gov 的券商账户义可用于核对概念；本轮未核实该义的美式或英式整词组书面音标。 |
| [character-encoding-html](../entries/character-encoding-html.yml) | character encoding | 前批查证后暂缓，本轮未重查。查得技术释义材料，未核实确切词头的 IPA。 |
| [character-reference-html](../entries/character-reference-html.yml) | character reference | 待查证 |
| [class-selector-css](../entries/class-selector-css.yml) | class selector | 待查证 |
| [cold-resistance-combat](../entries/cold-resistance-combat.yml) | cold resistance | 前批查证后暂缓，本轮未重查。检索未取得与游戏耐寒义匹配的书面整短语 IPA。 |
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
| [context-window-ai](../entries/context-window-ai.yml) | context window | 前批查证后暂缓，本轮未重查。Collins 有新词提议，未提供可用书面音标。 |
| [controlled-component-react](../entries/controlled-component-react.yml) | controlled component | 待查证 |
| [critical-hit-combat](../entries/critical-hit-combat.yml) | critical hit | 本轮查证后暂缓。检索了该完整搭配，未核实游戏暴击义的标口音整式。 |
| [cross-axis-css](../entries/cross-axis-css.yml) | cross axis | 待查证 |
| [cross-origin-resource-sharing-web](../entries/cross-origin-resource-sharing-web.yml) | cross-origin resource sharing | 待查证 |
| [custom-hook-react](../entries/custom-hook-react.yml) | custom Hook | 待查证 |
| [custom-property-css](../entries/custom-property-css.yml) | custom property | 待查证 |
| [data-attribute-bootstrap](../entries/data-attribute-bootstrap.yml) | data attribute | 前批查证后暂缓，本轮未重查。检索命中技术含义或例子，缺整短语 IPA。 |
| [data-cell-html](../entries/data-cell-html.yml) | data cell | 前批查证后暂缓，本轮未重查。检索命中双语聚合词页和其他语种定义，未找到可核实的口音音标及同义词条。 |
| [data-race-go](../entries/data-race-go.yml) | data race | 前批查证后暂缓，本轮未重查。计算机词典定向请求超时，未核实完整短语读法。 |
| [data-table-html](../entries/data-table-html.yml) | data table | 前批查证后暂缓，本轮未重查。检索主要得到两个组成词的读音，暂缺完整转写。 |
| [debounce-javascript](../entries/debounce-javascript.yml) | debounce | 前批查证后暂缓，本轮未重查。Wiktionary 英式 IPA 位于动词分区，尚不用于防抖方式名词。 |
| [declaration-file-typescript](../entries/declaration-file-typescript.yml) | declaration file | 待查证 |
| [decorative-image-html](../entries/decorative-image-html.yml) | decorative image | 待查证 |
| [deep-copy-javascript](../entries/deep-copy-javascript.yml) | deep copy | 前批查证后暂缓，本轮未重查。Wiktionary 词义匹配但未见整短语 IPA，Collins 页面未取得。 |
| [dependency-array-react](../entries/dependency-array-react.yml) | dependency array | 待查证 |
| [derived-state-react](../entries/derived-state-react.yml) | derived state | 待查证 |
| [detached-mode-docker](../entries/detached-mode-docker.yml) | detached mode | 待查证 |
| [disabled-control-html](../entries/disabled-control-html.yml) | disabled control | 待查证 |
| [discriminated-union-typescript](../entries/discriminated-union-typescript.yml) | discriminated union | 待查证 |
| [display-utility-bootstrap](../entries/display-utility-bootstrap.yml) | display utility | 待查证 |
| [dockerfile-docker](../entries/dockerfile-docker.yml) | Dockerfile | 本轮查证后暂缓。查到拼写和技术使用，未取得可核实的书面整词转写。 |
| [doctype-html](../entries/doctype-html.yml) | doctype | 前批查证后暂缓，本轮未重查。Wiktionary 可读页面未列整词 IPA。相似拼写或展开形式不作为替代。 |
| [document-title-html](../entries/document-title-html.yml) | document title | 待查证 |
| [draconification-zelda](../entries/draconification-zelda.yml) | draconification | 待查证 |
| [dropdown-bootstrap](../entries/dropdown-bootstrap.yml) | dropdown | 本轮查证后暂缓。Collins 将 dropdown 列为 drop-down 的异形并提供形容词美式 IPA。Cambridge 列完整 drop-down menu，独立下拉菜单名词的同形读音仍未核实。 |
| [effect-cleanup-react](../entries/effect-cleanup-react.yml) | Effect cleanup | 待查证 |
| [end-tag-html](../entries/end-tag-html.yml) | end tag | 待查证 |
| [error-boundary-react](../entries/error-boundary-react.yml) | error boundary | 待查证 |
| [error-log-nginx](../entries/error-log-nginx.yml) | error log | 前批查证后暂缓，本轮未重查。计算机词典第 128 页有 error logging 等相近词头，未取得 error log 整短语音标。 |
| [error-wrapping-go](../entries/error-wrapping-go.yml) | error wrapping | 待查证 |
| [event-bubbling-dom](../entries/event-bubbling-dom.yml) | event bubbling | 待查证 |
| [event-capturing-dom](../entries/event-capturing-dom.yml) | event capturing | 待查证 |
| [event-delegation-dom](../entries/event-delegation-dom.yml) | event delegation | 待查证 |
| [event-loop-javascript](../entries/event-loop-javascript.yml) | event loop | 前批查证后暂缓，本轮未重查。Wiktionary 确切词形 IPA 仍无口音限定，镜像未补证。 |
| [exact-match-nginx](../entries/exact-match-nginx.yml) | exact match | 前批查证后暂缓，本轮未重查。Collins 与 Cambridge 返回词组或用例，未取得整短语 IPA。 |
| [exhaustiveness-checking-typescript](../entries/exhaustiveness-checking-typescript.yml) | exhaustiveness checking | 待查证 |
| [expected-move-finance](../entries/expected-move-finance.yml) | expected move | 前批查证后暂缓，本轮未重查。定向检索尚无标口音的金融用法完整 IPA。 |
| [fall-damage-game](../entries/fall-damage-game.yml) | fall damage | 前批查证后暂缓，本轮未重查。IPA 定向检索未得到该游戏名词短语的可核实读法。 |
| [fallback-value-css](../entries/fallback-value-css.yml) | fallback value | 待查证 |
| [fear-and-greed-index-finance](../entries/fear-and-greed-index-finance.yml) | fear and greed index | 待查证 |
| [fixed-positioning-css](../entries/fixed-positioning-css.yml) | fixed positioning | 待查证 |
| [flex-container-css](../entries/flex-container-css.yml) | flex container | 待查证 |
| [flex-item-css](../entries/flex-item-css.yml) | flex item | 待查证 |
| [flexbox-css](../entries/flexbox-css.yml) | flexbox | 前批查证后暂缓，本轮未重查。Wiktionary 有相关术语，未取得可核实的整词音标。 |
| [floating-point-precision-javascript](../entries/floating-point-precision-javascript.yml) | floating-point precision | 前批查证后暂缓，本轮未重查。Cambridge floating-point operations per second 是另一词组，Collins 的对应整条未取得。 |
| [fluid-container-bootstrap](../entries/fluid-container-bootstrap.yml) | fluid container | 待查证 |
| [focus-indicator-css](../entries/focus-indicator-css.yml) | focus indicator | 待查证 |
| [font-stack-css](../entries/font-stack-css.yml) | font stack | 待查证 |
| [form-control-bootstrap](../entries/form-control-bootstrap.yml) | form control | 前批查证后暂缓，本轮未重查。TERMIUM Plus 和双语词典命中翻译，未确认当前整短语的书面读音。 |
| [form-submission-html](../entries/form-submission-html.yml) | form submission | 前批查证后暂缓，本轮未重查。未取得可核实的整个 HTML 术语转写。 |
| [fragment-identifier-html](../entries/fragment-identifier-html.yml) | fragment identifier | 待查证 |
| [function-component-react](../entries/function-component-react.yml) | function component | 待查证 |
| [function-signature-typescript](../entries/function-signature-typescript.yml) | function signature | 前批查证后暂缓，本轮未重查。检索未获可核验的带口音整短语 IPA。 |
| [funds-on-hold-finance](../entries/funds-on-hold-finance.yml) | funds on hold | 前批查证后暂缓，本轮未重查。相关银行术语表不提供可用的完整词组音标。 |
| [fuzz-testing-go](../entries/fuzz-testing-go.yml) | fuzz testing | 前批查证后暂缓，本轮未重查。Wiktionary 可核实测试概念，但未发现词组的 IPA。 |
| [garbage-collection-javascript](../entries/garbage-collection-javascript.yml) | garbage collection | 本轮查证后暂缓。Collins 和 Dictionary.com 的可读计算机义正文没有完整音标。其他检索结果的口音或编辑来源暂未核实。 |
| [generic-constraint-typescript](../entries/generic-constraint-typescript.yml) | generic constraint | 待查证 |
| [global-attribute-html](../entries/global-attribute-html.yml) | global attribute | 待查证 |
| [global-augmentation-typescript](../entries/global-augmentation-typescript.yml) | global augmentation | 待查证 |
| [go-module](../entries/go-module.yml) | Go module | 待查证 |
| [goroutine-go](../entries/goroutine-go.yml) | goroutine | 前批查证后暂缓，本轮未重查。Wiktionary 可读页面未列整词 IPA，尚缺 Go 术语读法证据。 |
| [graceful-degradation-web](../entries/graceful-degradation-web.yml) | graceful degradation | 前批查证后暂缓，本轮未重查。Collins 对应页面未取得，Wiktionary 命中 degradation 的单词音标。 |
| [grid-area-css](../entries/grid-area-css.yml) | grid area | 待查证 |
| [grid-container-css](../entries/grid-container-css.yml) | grid container | 待查证 |
| [grid-layout-css](../entries/grid-layout-css.yml) | grid layout | 前批查证后暂缓，本轮未重查。检索返回单独 layout 及搭配列表，未核实完整词组读音。 |
| [grid-system-bootstrap](../entries/grid-system-bootstrap.yml) | grid system | 待查证 |
| [grid-track-css](../entries/grid-track-css.yml) | grid track | 待查证 |
| [happens-before-go](../entries/happens-before-go.yml) | happens-before | 待查证 |
| [header-cell-html](../entries/header-cell-html.yml) | header cell | 待查证 |
| [heading-level-html](../entries/heading-level-html.yml) | heading level | 待查证 |
| [heap-memory-javascript](../entries/heap-memory-javascript.yml) | heap memory | 前批查证后暂缓，本轮未重查。计算机词典定向请求超时，未核实完整短语读法。 |
| [heat-resistance-combat](../entries/heat-resistance-combat.yml) | heat resistance | 本轮查证后暂缓。查到术语表和翻译结果，未核实整个搭配的书面 IPA。 |
| [higher-order-function-javascript](../entries/higher-order-function-javascript.yml) | higher-order function | 前批查证后暂缓，本轮未重查。Wiktionary 收录编程术语，IPA 搜索命中其组成词，未确认整短语读法。 |
| [hoisting-javascript](../entries/hoisting-javascript.yml) | hoisting | 前批查证后暂缓，本轮未重查。Wiktionary 所列整词 IPA 未标口音，无 JavaScript 专义转写。 |
| [image-digest-docker](../entries/image-digest-docker.yml) | image digest | 待查证 |
| [image-layer-docker](../entries/image-layer-docker.yml) | image layer | 待查证 |
| [image-repository-docker](../entries/image-repository-docker.yml) | image repository | 待查证 |
| [image-tag-docker](../entries/image-tag-docker.yml) | image tag | 待查证 |
| [implicit-implementation-go](../entries/implicit-implementation-go.yml) | implicit implementation | 待查证 |
| [indexed-access-type-typescript](../entries/indexed-access-type-typescript.yml) | indexed access type | 待查证 |
| [initializer-tech](../entries/initializer-tech.yml) | initializer | 本轮查证后暂缓。Fluent Python 及软件词汇表可核实技术义和拼写，未取得标口音整词 IPA。 |
| [inline-box-css](../entries/inline-box-css.yml) | inline box | 待查证 |
| [inline-frame-html](../entries/inline-frame-html.yml) | inline frame | 前批查证后暂缓，本轮未重查。英语检索多为 iframe。已定位双语词条，正文请求触发访问检查，尚未核实标口音的完整读音。 |
| [input-group-bootstrap](../entries/input-group-bootstrap.yml) | input group | 待查证 |
| [input-type-html](../entries/input-type-html.yml) | input type | 待查证 |
| [interactive-content-html](../entries/interactive-content-html.yml) | interactive content | 待查证 |
| [internal-redirect-nginx](../entries/internal-redirect-nginx.yml) | internal redirect | 待查证 |
| [intersection-type-typescript](../entries/intersection-type-typescript.yml) | intersection type | 待查证 |
| [invalid-state-bootstrap](../entries/invalid-state-bootstrap.yml) | invalid state | 待查证 |
| [iterable-tech](../entries/iterable-tech.yml) | iterable | 本轮查证后暂缓。Fluent Python 和课程术语表有可迭代对象义，未核实确切词形的标口音书面转写。 |
| [javascript-plugin-bootstrap](../entries/javascript-plugin-bootstrap.yml) | JavaScript plugin | 待查证 |
| [lazy-loading-html](../entries/lazy-loading-html.yml) | lazy loading | 前批查证后暂缓，本轮未重查。Collins 对应页面未取得，Wiktionary lazy load 是不同词形。 |
| [least-recently-used-cache](../entries/least-recently-used-cache.yml) | least recently used | 待查证 |
| [lifting-state-up-react](../entries/lifting-state-up-react.yml) | lifting state up | 待查证 |
| [line-height-css](../entries/line-height-css.yml) | line height | 前批查证后暂缓，本轮未重查。计算机词典定向请求超时，未核实完整短语读法。 |
| [linked-list-merge-algorithm](../entries/linked-list-merge-algorithm.yml) | linked list merge | 待查证 |
| [list-group-bootstrap](../entries/list-group-bootstrap.yml) | list group | 待查证 |
| [list-item-html](../entries/list-item-html.yml) | list item | 待查证 |
| [list-rendering-react](../entries/list-rendering-react.yml) | list rendering | 待查证 |
| [literal-type-typescript](../entries/literal-type-typescript.yml) | literal type | 待查证 |
| [load-balancing-nginx](../entries/load-balancing-nginx.yml) | load balancing | 前批查证后暂缓，本轮未重查。计算机词典的字形复查受访问限制，前批 Collins 释义可确认概念但未给整短语 IPA。 |
| [longest-prefix-nginx](../entries/longest-prefix-nginx.yml) | longest prefix | 待查证 |
| [longhand-property-css](../entries/longhand-property-css.yml) | longhand property | 待查证 |
| [main-axis-css](../entries/main-axis-css.yml) | main axis | 待查证 |
| [main-quest-game](../entries/main-quest-game.yml) | main quest | 前批查证后暂缓，本轮未重查。Collins 与 Wiktionary 的精确短语页面未取得。搜索到的 sidequest 属于另一词形及义项。 |
| [male-to-male-hardware](../entries/male-to-male-hardware.yml) | male-to-male | 前批查证后暂缓，本轮未重查。未取得适用于插头连接器义及连字符形容词的完整 IPA。 |
| [manual-save-game](../entries/manual-save-game.yml) | manual save | 前批查证后暂缓，本轮未重查。未取得自动或手动存档对应词形的可靠整短语 IPA。已确认 auto-save 的资料不套用到 manual save。 |
| [mapped-type-typescript](../entries/mapped-type-typescript.yml) | mapped type | 待查证 |
| [margin-account-finance](../entries/margin-account-finance.yml) | margin account | 本轮查证后暂缓。查到金融词典及发音相关结果，未核实可读的标口音书面整式，Longman PDF 读取受限。 |
| [margin-collapsing-css](../entries/margin-collapsing-css.yml) | margin collapsing | 待查证 |
| [media-query-css](../entries/media-query-css.yml) | media query | 前批查证后暂缓，本轮未重查。定向词典检索未核实完整短语及 CSS 含义。 |
| [method-set-go](../entries/method-set-go.yml) | method set | 待查证 |
| [microtask-javascript](../entries/microtask-javascript.yml) | microtask | 本轮查证后暂缓。检索有技术定义及使用，未核实事件循环义的书面读音。 |
| [module-path-go](../entries/module-path-go.yml) | module path | 待查证 |
| [module-specifier-tech](../entries/module-specifier-tech.yml) | module specifier | 待查证 |
| [mount-point-docker](../entries/mount-point-docker.yml) | mount point | 前批查证后暂缓，本轮未重查。定向检索只取得 mount 的音标或不同写法 mountpoint，未核实当前词形的整短语。 |
| [multiple-return-values-go](../entries/multiple-return-values-go.yml) | multiple return values | 待查证 |
| [named-result-parameter-go](../entries/named-result-parameter-go.yml) | named result parameter | 待查证 |
| [namespace-typescript](../entries/namespace-typescript.yml) | namespace | 本轮查证后暂缓。Wiktionary 的同形技术词给出 /ˈneɪmspeɪs/，未标口音。另一个学习站同时标两种口音，但未核实其转写来源或编辑方法。 |
| [navbar-bootstrap](../entries/navbar-bootstrap.yml) | navbar | 前批查证后暂缓，本轮未重查。Wiktionary 与 Collins 新词提议检查未取得可用整词 IPA。 |
| [normal-flow-css](../entries/normal-flow-css.yml) | normal flow | 待查证 |
| [offcanvas-bootstrap](../entries/offcanvas-bootstrap.yml) | offcanvas | 前批查证后暂缓，本轮未重查。未取得适用整词 IPA，分写 off canvas 或不同连字符形不能直接代用。 |
| [one-way-data-flow-frontend](../entries/one-way-data-flow-frontend.yml) | one-way data flow | 待查证 |
| [onion-model-koa](../entries/onion-model-koa.yml) | onion model | 待查证 |
| [option-chain-finance](../entries/option-chain-finance.yml) | option chain | 待查证 |
| [optional-parameter-typescript](../entries/optional-parameter-typescript.yml) | optional parameter | 待查证 |
| [optional-property-typescript](../entries/optional-property-typescript.yml) | optional property | 待查证 |
| [ordered-list-html](../entries/ordered-list-html.yml) | ordered list | 前批查证后暂缓，本轮未重查。计算机词典定向请求超时，旧 Cambridge 用例仍不足以支持音标。 |
| [own-property-javascript](../entries/own-property-javascript.yml) | own property | 待查证 |
| [partial-clone-git](../entries/partial-clone-git.yml) | partial clone | 待查证 |
| [persistent-data-docker](../entries/persistent-data-docker.yml) | persistent data | 待查证 |
| [pointer-receiver-go](../entries/pointer-receiver-go.yml) | pointer receiver | 待查证 |
| [popover-bootstrap](../entries/popover-bootstrap.yml) | popover | 本轮查证后暂缓。Cambridge 和 Dictionary.com 的可见转写仅覆盖烘焙食品义，未核实界面组件义。 |
| [port-publishing-docker](../entries/port-publishing-docker.yml) | port publishing | 待查证 |
| [postorder-traversal-tree](../entries/postorder-traversal-tree.yml) | postorder traversal | 待查证 |
| [pre-market-finance](../entries/pre-market-finance.yml) | pre-market | 本轮查证后暂缓。核对了 Cambridge 发音页的全文片段，只确认整词音频入口，未找到完整书面 IPA。 |
| [prefix-match-nginx](../entries/prefix-match-nginx.yml) | prefix match | 待查证 |
| [preflight-request-web](../entries/preflight-request-web.yml) | preflight request | 待查证 |
| [progress-bar-bootstrap](../entries/progress-bar-bootstrap.yml) | progress bar | 拟补 B12-P02 |
| [progressive-enhancement-web](../entries/progressive-enhancement-web.yml) | progressive enhancement | 前批查证后暂缓，本轮未重查。Collins 对应页面未取得，Wiktionary 命中 enhancement 的单词音标。 |
| [property-descriptor-javascript](../entries/property-descriptor-javascript.yml) | property descriptor | 待查证 |
| [property-value-css](../entries/property-value-css.yml) | property value | 待查证 |
| [props-react](../entries/props-react.yml) | props | 前批查证后暂缓，本轮未重查。沿用第八批清单。本轮查证后暂缓。Collins 命中其他义项；Wiktionary 给出单数 prop 或其他语言的音标。未据此补 React 义。 |
| [prototype-chain-javascript](../entries/prototype-chain-javascript.yml) | prototype chain | 待查证 |
| [prototype-setter-javascript](../entries/prototype-setter-javascript.yml) | prototype setter | 待查证 |
| [pseudo-class-css](../entries/pseudo-class-css.yml) | pseudo-class | 本轮查证后暂缓。资料可核对 CSS 词头和概念；未取得标口音的完整词头转写。 |
| [pseudo-element-css](../entries/pseudo-element-css.yml) | pseudo-element | 前批查证后暂缓，本轮未重查。Wiktionary 页面未取得，仍需确切词形和义项的证据。 |
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
| [responsive-design-css](../entries/responsive-design-css.yml) | responsive design | 前批查证后暂缓，本轮未重查。Wiktionary 可确定术语义，未见英语整短语 IPA。 |
| [responsive-image-html](../entries/responsive-image-html.yml) | responsive image | 前批查证后暂缓，本轮未重查。搜索结果多是技术文档，未取得可核实的整式音标。 |
| [responsive-utility-bootstrap](../entries/responsive-utility-bootstrap.yml) | responsive utility | 待查证 |
| [rest-parameter-typescript](../entries/rest-parameter-typescript.yml) | rest parameter | 待查证 |
| [return-type-tech](../entries/return-type-tech.yml) | return type | 前批查证后暂缓，本轮未重查。词典检索命中 return 或 type 的单词读音，未找到整短语。 |
| [reverse-proxy-nginx](../entries/reverse-proxy-nginx.yml) | reverse proxy | 本轮查证后暂缓。TurkCouncil 同页列出互不一致的书面结果。Wiktionary 查询结果呈现的是 proxy 单词等邻近结果；未取得可靠的整词组转写。 |
| [row-span-html](../entries/row-span-html.yml) | row span | 待查证 |
| [rules-of-hooks-react](../entries/rules-of-hooks-react.yml) | Rules of Hooks | 待查证 |
| [same-origin-policy-web](../entries/same-origin-policy-web.yml) | same-origin policy | 待查证 |
| [sass-mixin-bootstrap](../entries/sass-mixin-bootstrap.yml) | Sass mixin | 待查证 |
| [sass-variable-bootstrap](../entries/sass-variable-bootstrap.yml) | Sass variable | 待查证 |
| [semantic-versioning-software](../entries/semantic-versioning-software.yml) | semantic versioning | 待查证 |
| [sentinel-error-go](../entries/sentinel-error-go.yml) | sentinel error | 待查证 |
| [server-rendering-react](../entries/server-rendering-react.yml) | server rendering | 待查证 |
| [shallow-copy-javascript](../entries/shallow-copy-javascript.yml) | shallow copy | 待查证 |
| [shock-resistance-combat](../entries/shock-resistance-combat.yml) | shock resistance | 前批查证后暂缓，本轮未重查。未取得说明电击义及整短语读法的词典资料。 |
| [short-variable-declaration-go](../entries/short-variable-declaration-go.yml) | short variable declaration | 待查证 |
| [shorthand-property-css](../entries/shorthand-property-css.yml) | shorthand property | 待查证 |
| [side-hop-combat](../entries/side-hop-combat.yml) | side hop | 前批查证后暂缓，本轮未重查。检索结果包含其他跳跃动作，未核实当前游戏用法的整式读法。 |
| [sky-island-zelda](../entries/sky-island-zelda.yml) | sky island | 待查证 |
| [source-order-css](../entries/source-order-css.yml) | source order | 待查证 |
| [source-set-html](../entries/source-set-html.yml) | source set | 待查证 |
| [spacing-utility-bootstrap](../entries/spacing-utility-bootstrap.yml) | spacing utility | 待查证 |
| [stack-trace-tech](../entries/stack-trace-tech.yml) | stack trace | 前批查证后暂缓，本轮未重查。检索得到 trace 词页和双语材料，未证实整短语的合格 IPA。 |
| [stacking-context-css](../entries/stacking-context-css.yml) | stacking context | 待查证 |
| [start-tag-html](../entries/start-tag-html.yml) | start tag | 待查证 |
| [state-preservation-react](../entries/state-preservation-react.yml) | state preservation | 待查证 |
| [state-setter-react](../entries/state-setter-react.yml) | state setter | 待查证 |
| [state-snapshot-react](../entries/state-snapshot-react.yml) | state snapshot | 待查证 |
| [sticky-positioning-css](../entries/sticky-positioning-css.yml) | sticky positioning | 待查证 |
| [storage-event-browser](../entries/storage-event-browser.yml) | storage event | 待查证 |
| [strict-mode-react](../entries/strict-mode-react.yml) | Strict Mode | 待查证 |
| [strike-price-finance](../entries/strike-price-finance.yml) | strike price | 本轮查证后暂缓。NetEnglish 的期权义词表列出 /straɪk praɪs/，未标口音。Cambridge 的整词组发音入口只核实到音频。先保留缺失。 |
| [structural-check-tech](../entries/structural-check-tech.yml) | structural check | 待查证 |
| [structural-typing-typescript](../entries/structural-typing-typescript.yml) | structural typing | 待查证 |
| [style-rule-css](../entries/style-rule-css.yml) | style rule | 待查证 |
| [subtest-go](../entries/subtest-go.yml) | subtest | 待查证 |
| [symlink-tech](../entries/symlink-tech.yml) | symlink | 前批查证后暂缓，本轮未重查。Wiktionary 有技术定义但无适用的整词 IPA。 |
| [system-prompt-ai](../entries/system-prompt-ai.yml) | system prompt | 前批查证后暂缓，本轮未重查。检索命中芬兰语对应词的 IPA 及双语用例，未证实英语整词组。 |
| [table-driven-test-go](../entries/table-driven-test-go.yml) | table-driven test | 待查证 |
| [table-row-html](../entries/table-row-html.yml) | table row | 待查证 |
| [template-literal-type-typescript](../entries/template-literal-type-typescript.yml) | template literal type | 待查证 |
| [temporal-dead-zone-javascript](../entries/temporal-dead-zone-javascript.yml) | temporal dead zone | 待查证 |
| [text-wrapping-css](../entries/text-wrapping-css.yml) | text wrapping | 待查证 |
| [theming-bootstrap](../entries/theming-bootstrap.yml) | theming | 前批查证后暂缓，本轮未重查。Wiktionary 可读页面未列整词 IPA。 |
| [toggle-trigger-bootstrap](../entries/toggle-trigger-bootstrap.yml) | toggle trigger | 待查证 |
| [tool-call-ai](../entries/tool-call-ai.yml) | tool call | 前批查证后暂缓，本轮未重查。定向词典检索未取得符合 AI 工具调用义的整条 IPA。 |
| [tool-result-ai](../entries/tool-result-ai.yml) | tool result | 前批查证后暂缓，本轮未重查。定向词典检索未取得符合 AI 工具结果义的整条 IPA。 |
| [tree-shaking-build](../entries/tree-shaking-build.yml) | tree shaking | 前批查证后暂缓，本轮未重查。搜索命中其他语言对该术语的释义与 shake 的 IPA，未取得整式。 |
| [two-way-binding-frontend](../entries/two-way-binding-frontend.yml) | two-way binding | 待查证 |
| [type-alias-typescript](../entries/type-alias-typescript.yml) | type alias | 待查证 |
| [type-annotation-typescript](../entries/type-annotation-typescript.yml) | type annotation | 待查证 |
| [type-constraint-go](../entries/type-constraint-go.yml) | type constraint | 待查证 |
| [type-guard-typescript](../entries/type-guard-typescript.yml) | type guard | 待查证 |
| [type-inference-typescript](../entries/type-inference-typescript.yml) | type inference | 前批查证后暂缓，本轮未重查。计算机词典定向请求超时，前批 Wiktionary 术语页未给整短语读法。 |
| [type-instantiation-tech](../entries/type-instantiation-tech.yml) | type instantiation | 待查证 |
| [type-only-import-typescript](../entries/type-only-import-typescript.yml) | type-only import | 待查证 |
| [type-parameter-typescript](../entries/type-parameter-typescript.yml) | type parameter | 前批查证后暂缓，本轮未重查。检索命中其他术语中的用例，未取得确切词头转写。 |
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
| [visually-hidden-bootstrap](../entries/visually-hidden-bootstrap.yml) | visually hidden | 待查证 |
| [void-element-html](../entries/void-element-html.yml) | void element | 待查证 |
| [wait-group-go](../entries/wait-group-go.yml) | wait group | 待查证 |
| [web-storage-browser](../entries/web-storage-browser.yml) | web storage | 待查证 |
| [withdrawable-finance](../entries/withdrawable-finance.yml) | withdrawable | 本轮查证后暂缓。可见来源未提供足够可靠的标口音整式，不能从 withdraw 的读音及后缀拼出字段。 |
| [worker-pool-concurrency](../entries/worker-pool-concurrency.yml) | worker pool | 前批查证后暂缓，本轮未重查。检索出现组成词或其他语种，未核实该并发术语完整读法。 |
| [year-over-year-finance](../entries/year-over-year-finance.yml) | year-over-year | 本轮查证后暂缓。Cambridge 整词组发音页有音频入口。本轮未核实整词组的书面 IPA。 |
| [zero-value-go](../entries/zero-value-go.yml) | zero value | 待查证 |

### 03 缩写及技术标识符

| 稳定 ID | 当前 Word | 状态及依据 |
|:--|:--|:--|
| [bigint-javascript](../entries/bigint-javascript.yml) | BigInt | 待查证 |
| [bool-go](../entries/bool-go.yml) | bool | 前批查证后暂缓，本轮未重查。沿用第八批清单。本轮查证后暂缓。Wiktionary 命中的相邻语种 IPA 不适用。保留 Go 标识符，待查其口读约定。 |
| [comparable-go](../entries/comparable-go.yml) | comparable | 本轮查证后暂缓。Cambridge 有普通形容词的读音及相近义。本轮未核实 Go 预声明约束标识符在该用法下的口读依据。 |
| [compose-yaml-docker](../entries/compose-yaml-docker.yml) | compose.yaml | 待查证 |
| [dom-html](../entries/dom-html.yml) | DOM | 前批查证后暂缓，本轮未重查。技术社区收录当前缩写但词条为 community-consensus，未给 /dɒm/ 的出处；Oxford 的 Dom 是其他义项。 |
| [fallback-ui-react](../entries/fallback-ui-react.yml) | fallback UI | 待查证 |
| [go-mod-go](../entries/go-mod-go.yml) | go.mod | 待查证 |
| [go-sum-go](../entries/go-sum-go.yml) | go.sum | 待查证 |
| [hdmi-cable-hardware](../entries/hdmi-cable-hardware.yml) | HDMI cable | 前批查证后暂缓，本轮未重查。Cambridge 与 Oxford 给出 HDMI 缩写的发音，未核实含 cable 的整短语。 |
| [hdmi-type-a-connector](../entries/hdmi-type-a-connector.yml) | HDMI Type A connector | 前批查证后暂缓，本轮未重查。核实缩写读音资料后仍缺完整连接器名称的连贯音标，未拼接。 |
| [http-caching-web](../entries/http-caching-web.yml) | HTTP caching | 待查证 |
| [id-selector-css](../entries/id-selector-css.yml) | ID selector | 待查证 |
| [int-go](../entries/int-go.yml) | int | 前批查证后暂缓，本轮未重查。沿用第八批清单。本轮查证后暂缓。Wiktionary 命中的相邻语种 IPA 不适用。保留 Go 标识符，待查其口读约定。 |
| [jsx-react](../entries/jsx-react.yml) | JSX | 前批查证后暂缓，本轮未重查。找不到可访问且明确采用完整缩写 IPA 的来源。 |
| [keyof-typescript](../entries/keyof-typescript.yml) | keyof | 待查证 |
| [new-go](../entries/new-go.yml) | new | 前批查证后暂缓，本轮未重查。沿用第八批清单。本轮查证后暂缓。Collins 美式音标为普通英语词。Go 内置函数口读惯例待确认。 |
| [omit-utility-type-typescript](../entries/omit-utility-type-typescript.yml) | Omit | 前批查证后暂缓，本轮未重查。沿用第八批清单。本轮查证后暂缓。Collins 提供普通动词 IPA。TypeScript 标识符的读法需要单独确认。 |
| [partial-utility-type-typescript](../entries/partial-utility-type-typescript.yml) | Partial | 前批查证后暂缓，本轮未重查。沿用第八批清单。本轮查证后暂缓。Collins 美式为形容词。当前为 TypeScript 工具类型名，待核实标识符口读依据。 |
| [pick-utility-type-typescript](../entries/pick-utility-type-typescript.yml) | Pick | 前批查证后暂缓，本轮未重查。沿用第八批清单。本轮查证后暂缓。Collins 提供普通动词 IPA。TypeScript 标识符的读法需要单独确认。 |
| [promise-all-javascript](../entries/promise-all-javascript.yml) | Promise.all | 待查证 |
| [property-css](../entries/property-css.yml) | CSS property | 待查证 |
| [readonly-typescript](../entries/readonly-typescript.yml) | readonly | 本轮查证后暂缓。Cambridge 有形容词 read-only 的美式 IPA；查得语言文档中技术拼写，但未找到 TypeScript 关键词的独立书面口读依据。 |
| [record-utility-type-typescript](../entries/record-utility-type-typescript.yml) | Record | 前批查证后暂缓，本轮未重查。沿用第八批清单。本轮查证后暂缓。Collins 区分 record 的名词和动词重音。当前为 TypeScript 工具类型，待核实口读依据。 |
| [returntype-utility-type-typescript](../entries/returntype-utility-type-typescript.yml) | ReturnType | 待查证 |
| [semantic-html](../entries/semantic-html.yml) | semantic HTML | 待查证 |
| [star-method-interview](../entries/star-method-interview.yml) | STAR method | 待查证 |
| [struct-go](../entries/struct-go.yml) | struct | 前批查证后暂缓，本轮未重查。社区词表可确认常见 struckt 读法，但无 Go 关键词的相关口读来源。 |
| [tts-language-tech](../entries/tts-language-tech.yml) | TTS language | 待查证 |
| [tv-mode-switch](../entries/tv-mode-switch.yml) | TV Mode | 待查证 |
| [utility-api-bootstrap](../entries/utility-api-bootstrap.yml) | utility API | 待查证 |
| [virtual-dom-frontend](../entries/virtual-dom-frontend.yml) | virtual DOM | 待查证 |

### 04 专名及游戏名称

| 稳定 ID | 当前 Word | 状态及依据 |
|:--|:--|:--|
| [acana-pet-food](../entries/acana-pet-food.yml) | ACANA | 本轮查证后暂缓。查到品牌资料与发音相关入口，未核实品牌书面整名 IPA。 |
| [adventure-log-zelda](../entries/adventure-log-zelda.yml) | Adventure Log | 待查证 |
| [aerocuda-zelda](../entries/aerocuda-zelda.yml) | Aerocuda | 待查证 |
| [ancient-arowana-zelda](../entries/ancient-arowana-zelda.yml) | Ancient Arowana | 前批查证后暂缓，本轮未重查。沿用第八批清单。本轮查证后暂缓。与普通鱼名一起筛查。相近英文或芬兰语拼写无法核实该游戏材料的全名读法。 |
| [attack-up-zelda](../entries/attack-up-zelda.yml) | Attack Up | 待查证 |
| [autobuild-zelda](../entries/autobuild-zelda.yml) | Autobuild | 待查证 |
| [beam-emitter-zelda](../entries/beam-emitter-zelda.yml) | Beam Emitter | 待查证 |
| [big-battery-zelda](../entries/big-battery-zelda.yml) | Big Battery | 待查证 |
| [big-wheel-zelda](../entries/big-wheel-zelda.yml) | Big Wheel | 待查证 |
| [bokoblin-zelda](../entries/bokoblin-zelda.yml) | Bokoblin | 待查证 |
| [bootstrap-framework](../entries/bootstrap-framework.yml) | Bootstrap | 前批查证后暂缓，本轮未重查。Collins IPA 对应普通词义，尚缺对应框架专名的证据。 |
| [boss-bokoblin-zelda](../entries/boss-bokoblin-zelda.yml) | Boss Bokoblin | 待查证 |
| [brightbloom-seed-zelda](../entries/brightbloom-seed-zelda.yml) | Brightbloom Seed | 待查证 |
| [brightcap-zelda](../entries/brightcap-zelda.yml) | Brightcap | 待查证 |
| [captain-construct-zelda](../entries/captain-construct-zelda.yml) | Captain Construct | 待查证 |
| [character-profiles-zelda](../entries/character-profiles-zelda.yml) | Character Profiles | 待查证 |
| [chuchu-zelda](../entries/chuchu-zelda.yml) | Chuchu | 待查证 |
| [claude-ai](../entries/claude-ai.yml) | Claude | 本轮查证后暂缓。查到产品资料及普通人名读音，未核实适用于产品的书面转写。 |
| [climbing-gear-zelda](../entries/climbing-gear-zelda.yml) | Climbing Gear | 待查证 |
| [codex-ai](../entries/codex-ai.yml) | Codex | 本轮查证后暂缓。查到产品资料与普通名词释义，未核实适用于产品名的书面转写。 |
| [colgera-zelda](../entries/colgera-zelda.yml) | Colgera | 待查证 |
| [construct-factory-zelda](../entries/construct-factory-zelda.yml) | Construct Factory | 待查证 |
| [construct-head-zelda](../entries/construct-head-zelda.yml) | Construct Head | 待查证 |
| [copilot-ai](../entries/copilot-ai.yml) | Copilot | 前批查证后暂缓，本轮未重查。Wiktionary 读取未取得产品名的转写。需要实际词条所指产品发音依据。 |
| [courser-bee-honey-zelda](../entries/courser-bee-honey-zelda.yml) | Courser Bee Honey | 待查证 |
| [crystal-refinery-zelda](../entries/crystal-refinery-zelda.yml) | Crystal Refinery | 待查证 |
| [crystallized-charge-zelda](../entries/crystallized-charge-zelda.yml) | Crystallized Charge | 待查证 |
| [cursor-ai](../entries/cursor-ai.yml) | Cursor | 前批查证后暂缓，本轮未重查。社区词表标记 consensus，链接只能确认指哪个产品，暂未确认词条书面读音的依据。 |
| [death-mountain-zelda](../entries/death-mountain-zelda.yml) | Death Mountain | 待查证 |
| [deepseek-ai](../entries/deepseek-ai.yml) | DeepSeek | 前批查证后暂缓，本轮未重查。社区词表有 IPA，来源链接仅支持产品身份，未找到正式口读依据。 |
| [defense-up-zelda](../entries/defense-up-zelda.yml) | Defense Up | 待查证 |
| [demon-dragon-zelda](../entries/demon-dragon-zelda.yml) | Demon Dragon | 待查证 |
| [demon-king-ganondorf-zelda](../entries/demon-king-ganondorf-zelda.yml) | Demon King Ganondorf | 待查证 |
| [demon-king-zelda](../entries/demon-king-zelda.yml) | Demon King | 待查证 |
| [demon-kings-bow-zelda](../entries/demon-kings-bow-zelda.yml) | Demon King's Bow | 待查证 |
| [device-dispenser-zelda](../entries/device-dispenser-zelda.yml) | Device Dispenser | 待查证 |
| [docker-compose-docker](../entries/docker-compose-docker.yml) | Docker Compose | 待查证 |
| [docker-platform](../entries/docker-platform.yml) | Docker | 本轮查证后暂缓。查到软件产品资料及普通同形词，未核实适用于 Docker 平台名称的书面转写。 |
| [dots-ai](../entries/dots-ai.yml) | dots | 待查证 |
| [electro-cooking](../entries/electro-cooking.yml) | Electro | 前批查证后暂缓，本轮未重查。沿用王国之泪计划 ZD013。当前 Electro 为游戏效果名称前缀；Collins electro 已读分区为普通名词或缩略形式，不直接套用其音标。需要名称用法的英文读音证据。 |
| [endura-carrot-zelda](../entries/endura-carrot-zelda.yml) | Endura Carrot | 待查证 |
| [endura-shroom-zelda](../entries/endura-shroom-zelda.yml) | Endura Shroom | 待查证 |
| [energy-cell-zelda](../entries/energy-cell-zelda.yml) | Energy Cell | 本轮查证后暂缓。可见结果包含普通翻译及邻近电学术语，未核实对应整个游戏道具名的英语转写。 |
| [energy-well-zelda](../entries/energy-well-zelda.yml) | Energy Well | 待查证 |
| [evermean-zelda](../entries/evermean-zelda.yml) | Evermean | 待查证 |
| [fancy-feast-pet-food](../entries/fancy-feast-pet-food.yml) | Fancy Feast | 本轮查证后暂缓。查到产品名称和发音相关结果，未核实该品牌的英语书面整名转写。保留词组原样。 |
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
| [ganondorf-zelda](../entries/ganondorf-zelda.yml) | Ganondorf | 本轮查证后暂缓。查到游戏官方拼写和社区讨论。Zelda Wiki 相关结果访问受限，未核实整名书面读法。 |
| [gemini-ai](../entries/gemini-ai.yml) | Gemini | 前批查证后暂缓，本轮未重查。社区记录借用同名星座读音，待核实 Google 产品专名的可靠转写。 |
| [gerudo-desert-zelda](../entries/gerudo-desert-zelda.yml) | Gerudo Desert | 待查证 |
| [gerudo-town-zelda](../entries/gerudo-town-zelda.yml) | Gerudo Town | 待查证 |
| [gerudo-zelda](../entries/gerudo-zelda.yml) | Gerudo | 本轮查证后暂缓。检索到官方拼写及社群讨论，未核实该族名的书面英语转写。 |
| [gibdo-zelda](../entries/gibdo-zelda.yml) | Gibdo | 待查证 |
| [gleeok-zelda](../entries/gleeok-zelda.yml) | Gleeok | 待查证 |
| [glide-shirt-zelda](../entries/glide-shirt-zelda.yml) | Glide Shirt | 待查证 |
| [gloom-recovery-zelda](../entries/gloom-recovery-zelda.yml) | Gloom Recovery | 待查证 |
| [gloom-resistance-zelda](../entries/gloom-resistance-zelda.yml) | Gloom Resistance | 待查证 |
| [gloom-spawn-zelda](../entries/gloom-spawn-zelda.yml) | Gloom Spawn | 待查证 |
| [gloom-sword-zelda](../entries/gloom-sword-zelda.yml) | Gloom Sword | 待查证 |
| [golden-apple-zelda](../entries/golden-apple-zelda.yml) | Golden Apple | 前批查证后暂缓，本轮未重查。定向检索未取得游戏道具全名的书面 IPA。 |
| [goron-city-zelda](../entries/goron-city-zelda.yml) | Goron City | 待查证 |
| [goron-spice-zelda](../entries/goron-spice-zelda.yml) | Goron Spice | 待查证 |
| [goron-zelda](../entries/goron-zelda.yml) | Goron | 本轮查证后暂缓。同形或近形查询含其他语言词义。未核实游戏族名的书面英语转写。 |
| [great-abandoned-central-mine-zelda](../entries/great-abandoned-central-mine-zelda.yml) | Great Abandoned Central Mine | 待查证 |
| [great-deku-tree-zelda](../entries/great-deku-tree-zelda.yml) | Great Deku Tree | 待查证 |
| [great-eagle-bow-zelda](../entries/great-eagle-bow-zelda.yml) | Great Eagle Bow | 待查证 |
| [great-sky-island-zelda](../entries/great-sky-island-zelda.yml) | Great Sky Island | 待查证 |
| [grok-ai](../entries/grok-ai.yml) | Grok | 前批查证后暂缓，本轮未重查。社区词表的 IPA 与重拼法在元音上需要进一步澄清，暂缓品牌读法。 |
| [hateno-village-zelda](../entries/hateno-village-zelda.yml) | Hateno Village | 待查证 |
| [heart-container-zelda](../entries/heart-container-zelda.yml) | Heart Container | 待查证 |
| [hearty-radish-zelda](../entries/hearty-radish-zelda.yml) | Hearty Radish | 待查证 |
| [hearty-salmon-zelda](../entries/hearty-salmon-zelda.yml) | Hearty Salmon | 待查证 |
| [hearty-truffle-zelda](../entries/hearty-truffle-zelda.yml) | Hearty Truffle | 待查证 |
| [hermes-agent-ai](../entries/hermes-agent-ai.yml) | Hermes | 本轮查证后暂缓。查到同形神话名称及项目结果，未确定目标产品的书面读法。 |
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
| [hylian-zelda](../entries/hylian-zelda.yml) | Hylian | 本轮查证后暂缓。查到游戏用例及社区讨论，未核实适用于该名的书面英语转写。 |
| [hyrule-bass-zelda](../entries/hyrule-bass-zelda.yml) | Hyrule Bass | 待查证 |
| [hyrule-castle-zelda](../entries/hyrule-castle-zelda.yml) | Hyrule Castle | 待查证 |
| [hyrule-compendium-zelda](../entries/hyrule-compendium-zelda.yml) | Hyrule Compendium | 待查证 |
| [hyrule-zelda](../entries/hyrule-zelda.yml) | Hyrule | 本轮查证后暂缓。检索到游戏名词及社区发音讨论。Zelda Wiki 的相关发音草稿无法访问。未核实整名的书面英语读音。 |
| [impa-zelda](../entries/impa-zelda.yml) | Impa | 待查证 |
| [instinct-pet-food](../entries/instinct-pet-food.yml) | Instinct | 待查证 |
| [josha-zelda](../entries/josha-zelda.yml) | Josha | 待查证 |
| [kakariko-village-zelda](../entries/kakariko-village-zelda.yml) | Kakariko Village | 本轮查证后暂缓。检索到地名资料，未核实整个村名的书面英语转写。 |
| [keese-zelda](../entries/keese-zelda.yml) | Keese | 待查证 |
| [key-items-zelda](../entries/key-items-zelda.yml) | Key Items | 待查证 |
| [king-gleeok-zelda](../entries/king-gleeok-zelda.yml) | King Gleeok | 待查证 |
| [knights-broadsword-zelda](../entries/knights-broadsword-zelda.yml) | Knight's Broadsword | 待查证 |
| [korok-forest-zelda](../entries/korok-forest-zelda.yml) | Korok Forest | 待查证 |
| [korok-seed-zelda](../entries/korok-seed-zelda.yml) | Korok Seed | 待查证 |
| [korok-zelda](../entries/korok-zelda.yml) | Korok | 待查证 |
| [legendsandy-pet-food](../entries/legendsandy-pet-food.yml) | Legendsandy | 待查证 |
| [lexforge-project](../entries/lexforge-project.yml) | LexForge | 前批查证后暂缓，本轮未重查。定向检索未找到用于此项目名的书面读音。保留已构建的名称和录音。 |
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
| [muse-ai](../entries/muse-ai.yml) | Muse | 前批查证后暂缓，本轮未重查。Wiktionary 英语释义和音标指神话、灵感来源等，未核实 Meta 产品名。 |
| [myfoodie-pet-food](../entries/myfoodie-pet-food.yml) | Myfoodie | 待查证 |
| [nutro-pet-food](../entries/nutro-pet-food.yml) | Nutro | 待查证 |
| [octorok-zelda](../entries/octorok-zelda.yml) | Octorok | 待查证 |
| [openclaw-ai](../entries/openclaw-ai.yml) | OpenClaw | 待查证 |
| [orijen-pet-food](../entries/orijen-pet-food.yml) | ORIJEN | 本轮查证后暂缓。品牌经销商网页提到 origin 作为命名来源。其他介绍把品牌读法比作 origin，未核实可用的品牌书面转写或可靠的同音声明。 |
| [paya-zelda](../entries/paya-zelda.yml) | Paya | 待查证 |
| [perfect-dodge-zelda](../entries/perfect-dodge-zelda.yml) | Perfect Dodge | 前批查证后暂缓，本轮未重查。游戏动作名定向检索未找到对应标签的整项书面音标。 |
| [perfect-guard-zelda](../entries/perfect-guard-zelda.yml) | Perfect Guard | 前批查证后暂缓，本轮未重查。游戏动作名定向检索未找到对应标签的整项书面音标。 |
| [phantom-ganon-zelda](../entries/phantom-ganon-zelda.yml) | Phantom Ganon | 待查证 |
| [portable-pot-zelda](../entries/portable-pot-zelda.yml) | Portable Pot | 待查证 |
| [princess-zelda-zelda](../entries/princess-zelda-zelda.yml) | Princess Zelda | 本轮查证后暂缓。查到角色称号和单独人名的读音结果，未核实包含 Princess 的完整称号书面转写。 |
| [puffshroom-zelda](../entries/puffshroom-zelda.yml) | Puffshroom | 待查证 |
| [purah-pad-zelda](../entries/purah-pad-zelda.yml) | Purah Pad | 待查证 |
| [purah-zelda](../entries/purah-zelda.yml) | Purah | 本轮查证后暂缓。游戏人物资料和读音讨论未提供可核实的书面英语转写。 |
| [purina-pro-plan-pet-food](../entries/purina-pro-plan-pet-food.yml) | Purina Pro Plan | 待查证 |
| [queen-gibdo-zelda](../entries/queen-gibdo-zelda.yml) | Queen Gibdo | 待查证 |
| [rauru-zelda](../entries/rauru-zelda.yml) | Rauru | 待查证 |
| [raw-gourmet-meat-zelda](../entries/raw-gourmet-meat-zelda.yml) | Raw Gourmet Meat | 待查证 |
| [raw-prime-meat-zelda](../entries/raw-prime-meat-zelda.yml) | Raw Prime Meat | 待查证 |
| [react-library](../entries/react-library.yml) | React | 前批查证后暂缓，本轮未重查。技术社区表给出非官方的常见读法，工程英语页未取得，留待来源查证。 |
| [regional-phenomena-zelda](../entries/regional-phenomena-zelda.yml) | Regional Phenomena | 待查证 |
| [riju-zelda](../entries/riju-zelda.yml) | Riju | 待查证 |
| [rito-village-zelda](../entries/rito-village-zelda.yml) | Rito Village | 待查证 |
| [rito-zelda](../entries/rito-zelda.yml) | Rito | 本轮查证后暂缓。检索到游戏用法，未核实适用于该族名的书面英语转写。 |
| [robbie-zelda](../entries/robbie-zelda.yml) | Robbie | 待查证 |
| [royal-broadsword-zelda](../entries/royal-broadsword-zelda.yml) | Royal Broadsword | 待查证 |
| [royal-canin-pet-food](../entries/royal-canin-pet-food.yml) | Royal Canin | 本轮查证后暂缓。查到品牌名称或其他语言读音资料；未核实品牌英语整名书面转写。 |
| [scimitar-of-the-seven-zelda](../entries/scimitar-of-the-seven-zelda.yml) | Scimitar of the Seven | 待查证 |
| [secret-stone-zelda](../entries/secret-stone-zelda.yml) | Secret Stone | 待查证 |
| [seized-construct-zelda](../entries/seized-construct-zelda.yml) | Seized Construct | 待查证 |
| [sheba-pet-food](../entries/sheba-pet-food.yml) | Sheba | 本轮查证后暂缓。查到同形名称及品牌用例；未取得可证明普通专名读音也适用于该品牌的书面证据。 |
| [sheikah-zelda](../entries/sheikah-zelda.yml) | Sheikah | 本轮查证后暂缓。查到游戏词形和社区用例，未核实该专名的书面英语转写。 |
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
| [spicy-pepper-zelda](../entries/spicy-pepper-zelda.yml) | Spicy Pepper | 本轮查证后暂缓。检索到名称用例及发音相关结果，未核实游戏道具名的完整书面英语转写。 |
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
| [time-bomb-zelda](../entries/time-bomb-zelda.yml) | Time Bomb | 拟补 B12-P01 |
| [travelers-sword-zelda](../entries/travelers-sword-zelda.yml) | Traveler's Sword | 待查证 |
| [tulin-zelda](../entries/tulin-zelda.yml) | Tulin | 本轮查证后暂缓。检索含其他语言同形词。未核实对应游戏人物的书面英语读音。 |
| [ultrahand-zelda](../entries/ultrahand-zelda.yml) | Ultrahand | 待查证 |
| [water-temple-zelda](../entries/water-temple-zelda.yml) | Water Temple | 待查证 |
| [whiskas-pet-food](../entries/whiskas-pet-food.yml) | Whiskas | 本轮查证后暂缓。Cambridge 正文快照在 w 前单列 h，仍未核实排印是否标为可选。保留缺失，避免自行修改原转写。 |
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
| [zora-zelda](../entries/zora-zelda.yml) | Zora | 本轮查证后暂缓。查到游戏资料及其他同形专名。未核实对应游戏族名的书面英语转写。 |
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
| [transform-css](../entries/transform-css.yml) | transform | 本轮查证后暂缓。Oxford 的发音属于动词，Dictionary.com 的数学名词不直接对应 CSS 属性名。属性的口读未核实。 |
| [try-files-nginx](../entries/try-files-nginx.yml) | try_files | 待查证 |
| [var-function-css](../entries/var-function-css.yml) | var() | 待查证 |
| [would-you-like-a-or-b](../entries/would-you-like-a-or-b.yml) | Would you like A or B? | 前批查证后暂缓，本轮未重查。Cambridge Grammar 确认句型用法，但示例用具体名词，不为带 A、B 占位符的完整模板提供读音。 |
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

已核对全库稳定 ID、语言计数、缺失清单、审阅编号、源词条路径和本地链接。规划前后哈希一致，原有词条、录音、工具和导出均保留。规划阶段只新增本计划。收到明确批准后实施。

## 实施结果

- B12-P01 和 B12-P02 已写入对应词条。每个词条增加一个书面美式 IPA 和一行 Notes 来源，保留原 Notes 和其余 YAML 内容。
- `go run scripts/validate/main.go` 校验 1,833 项通过，报告 38 条既有同词形警告。`go run scripts/build/main.go` 连续两次构建的导出字节一致。`git diff --check` 通过，计划格式和本地链接另行核对。
- 实施前的隔离构建与原导出一致。对比实施前后 1,833 行、18 列的导出，仅两个获批词条的 Pronunciation 和 Notes 有差异。ID、Audio、StudyOrder 和其他列均一致。1,875 个 MP3 的哈希与快照一致。
- 英语 1,040 项已有书面读音，仍缺 550 项。日语 243 项均有读音。人工听审及真实 Anki 导入、显示和播放尚未验证。
