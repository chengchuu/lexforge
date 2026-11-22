# 英语音标补全计划：第七批常用词与开发界面词

状态为已批准并实施。2026-11-22 核对。下文保留规划基线，最后一节记录实施结果。

## 范围与基线

当前词库有 1,833 项，英语 1,590 项、日语 243 项。英语缺少 pronunciation 的 608 项作为候选池。日语 243 项均有读法。此批先查常见词形和开发中常见的词，逐项查看 43 个稳定 ID，形成 25 项文案。另 18 项查证后暂缓，565 项尚未逐项查证。25 项均实施后，缺口将减至 583 项。

沿用[词条规范](ENTRY_SCHEMA.md#pronunciation)及[第六批记录](EN_PRONUNCIATION_ENRICHMENT_BATCH_06_PLAN.md)，保留已有音标和已实施的记忆辅助。浏览词典正文时核对词形、词性和适用义项。技术义仍由当前词条和原出处说明。游戏专名、品牌、缩写、命令符号及需要整短语证据的条目留到专项核对。

## 具体提案

23 项采用美式 IPA，无前缀。2 项采用明确的英式整词 IPA，标为 UK。Wiktionary 的 5 项使用正文标明的 US 或 General American 完整词形及固定修订地址。其余提案以 Collins 的相应口音分区为准。

| 编号 | 稳定 ID / Word | 当前词性 | 当前 Meaning | 拟写 pronunciation | 证据 |
|:--|:--|:--|:--|:--|:--|
| B07-P01 | [alert-bootstrap](../entries/alert-bootstrap.yml) / alert | noun | 提示框，用于在页面中显示操作结果或其他上下文反馈的组件 | `/əˈlɜrt/` | [Collins Dictionary — alert; American English](https://www.collinsdictionary.com/dictionary/english/alert) |
| B07-P02 | [argument-tech](../entries/argument-tech.yml) / argument | noun | 实参，调用时传给函数的值 | `/ˈɑrɡjumənt/` | [Collins Dictionary — argument; American English](https://www.collinsdictionary.com/dictionary/english/argument) |
| B07-P03 | [boots-clothing](../entries/boots-clothing.yml) / boots | noun | 靴子 | `/buts/` | [Wiktionary — boots; English Pronunciation, US](https://en.wiktionary.org/w/index.php?title=boots&oldid=93681493) |
| B07-P04 | [breadcrumb-bootstrap](../entries/breadcrumb-bootstrap.yml) / breadcrumb | noun | 面包屑导航，显示当前页面所在层级路径的导航组件 | `/ˈbredˌkrʌm/` | [Collins Dictionary — breadcrumb; American English](https://www.collinsdictionary.com/dictionary/english/breadcrumb) |
| B07-P05 | [breakpoint-css](../entries/breakpoint-css.yml) / breakpoint | noun | 断点，在特定条件下切换布局或样式的阈值 | `/ˈbreikˌpɔint/` | [Collins Dictionary — breakpoint; American English](https://www.collinsdictionary.com/dictionary/english/breakpoint) |
| B07-P06 | [card-bootstrap](../entries/card-bootstrap.yml) / card | noun | 卡片，用于组合标题、正文、图片或其他内容的柔性容器组件 | `/kɑrd/` | [Collins Dictionary — card; American English](https://www.collinsdictionary.com/dictionary/english/card) |
| B07-P07 | [collapse-bootstrap](../entries/collapse-bootstrap.yml) / collapse | noun | 折叠，用于展开和收起来切换内容可见性的组件行为 | `/kəˈlæps/` | [Collins Dictionary — collapse; American English](https://www.collinsdictionary.com/dictionary/english/collapse) |
| B07-P08 | [column-bootstrap](../entries/column-bootstrap.yml) / column | noun | 列，位于栅格行内并可占据若干栅格单位的布局项 | `/ˈkɑləm/` | [Collins Dictionary — column; American English](https://www.collinsdictionary.com/dictionary/english/column) |
| B07-P09 | [deadlock-go](../entries/deadlock-go.yml) / deadlock | noun | 死锁，相关执行单元一直等待而无法继续 | `/ˈdɛdˌlɑk/` | [Collins Dictionary — deadlock; American English](https://www.collinsdictionary.com/dictionary/english/deadlock) |
| B07-P10 | [declaration-css](../entries/declaration-css.yml) / declaration | noun | 声明，给某个 CSS 属性指定一个值 | `/ˌdɛkləˈreɪʃən/` | [Collins Dictionary — declaration; American English](https://www.collinsdictionary.com/dictionary/english/declaration) |
| B07-P11 | [disguised-combat](../entries/disguised-combat.yml) / disguised | adjective | 伪装的 | `/dɪsˈɡaɪzd/` | [Wiktionary — disguised; English Pronunciation, General American](https://en.wiktionary.org/w/index.php?title=disguised&oldid=83217429) |
| B07-P12 | [durability-equipment](../entries/durability-equipment.yml) / durability | noun | 耐用性；耐久度。 | `/dəɹəˈbɪləti/` | [Wiktionary — durability; English Pronunciation, US](https://en.wiktionary.org/w/index.php?title=durability&oldid=94553680) |
| B07-P13 | [embedding-go](../entries/embedding-go.yml) / embedding | noun | 嵌入，在结构体中声明嵌入字段或在接口中包含其他接口 | `/emˈbedɪŋ/` | [Collins Dictionary — embedding; American English](https://www.collinsdictionary.com/dictionary/english/embedding) |
| B07-P14 | [enhancement-general](../entries/enhancement-general.yml) / enhancement | noun | 强化、增强 | `/ɪnˈhæns.mənt/` | [Wiktionary — enhancement; English Pronunciation, General American](https://en.wiktionary.org/w/index.php?title=enhancement&oldid=90606147) |
| B07-P15 | [generic-typescript](../entries/generic-typescript.yml) / generic | noun | 泛型，使用类型参数描述可复用的函数或类型 | `/dʒəˈnɛrɪk/` | [Collins Dictionary — generic; American English](https://www.collinsdictionary.com/dictionary/english/generic) |
| B07-P16 | [interface-go](../entries/interface-go.yml) / interface | keyword | 接口，通过类型集规定哪些类型符合要求，常用一组方法描述行为 | `/ˈɪntərˌfeɪs/` | [Collins Dictionary — interface; American English](https://www.collinsdictionary.com/dictionary/english/interface) |
| B07-P17 | [modal-bootstrap](../entries/modal-bootstrap.yml) / modal | noun | 模态对话框，显示在页面内容上方并限制用户与背景交互的组件 | `/ˈmoʊdəl/` | [Collins Dictionary — modal; American English](https://www.collinsdictionary.com/dictionary/english/modal) |
| B07-P18 | [narrowing-typescript](../entries/narrowing-typescript.yml) / narrowing | noun | 类型收窄，根据检查和控制流缩小可能的类型 | `UK /ˈnærəʊɪŋ/` | [Collins Dictionary — narrowing; British English](https://www.collinsdictionary.com/dictionary/english/narrowing) |
| B07-P19 | [nesting-html](../entries/nesting-html.yml) / nesting | noun | 嵌套，把一个元素放在另一个元素内部形成层级结构 | `UK /ˈnɛstɪŋ/` | [Collins Dictionary — nesting; British English](https://www.collinsdictionary.com/dictionary/english/nesting) |
| B07-P20 | [pagination-bootstrap](../entries/pagination-bootstrap.yml) / pagination | noun | 分页，用于在多页内容之间导航的组件 | `/ˌpædʒəˈneɪʃən/` | [Collins Dictionary — pagination; American English](https://www.collinsdictionary.com/dictionary/english/pagination) |
| B07-P21 | [promise-javascript](../entries/promise-javascript.yml) / Promise | noun | Promise；表示异步操作最终成功或失败结果的对象。 | `/ˈprɑmɪs/` | [Collins Dictionary — promise; American English](https://www.collinsdictionary.com/dictionary/english/promise) |
| B07-P22 | [specificity-css](../entries/specificity-css.yml) / specificity | noun | 选择器权重，层叠中用于比较匹配规则的一个因素 | `/ˌspɛsəˈfɪsəti/` | [Collins Dictionary — specificity; American English](https://www.collinsdictionary.com/dictionary/english/specificity) |
| B07-P23 | [symbol-javascript](../entries/symbol-javascript.yml) / Symbol | noun | Symbol；用于表示独一且不可变标识值的 JavaScript 原始类型。 | `/ˈsɪmbəl/` | [Collins Dictionary — symbol; American English](https://www.collinsdictionary.com/dictionary/english/symbol) |
| B07-P24 | [tuple-typescript](../entries/tuple-typescript.yml) / tuple | noun | 元组，按位置定义元素类型的数组类型 | `/ˈtuːpəl/` | [Wiktionary — tuple; English Pronunciation, US](https://en.wiktionary.org/w/index.php?title=tuple&oldid=89359660) |
| B07-P25 | [warp-general](../entries/warp-general.yml) / warp | verb | 传送；游戏中快速移动到另一地点。 | `/wɔrp/` | [Collins Dictionary — warp; American English](https://www.collinsdictionary.com/dictionary/english/warp) |

### 英式回退依据

- `narrowing-typescript`：Collins 英式分区有完整词形。Cambridge 拒绝访问，Oxford 未能读取，所读 Wiktionary 未提供该词形 IPA。
- `nesting-html`：Collins 英式分区提供完整 IPA。Cambridge 拒绝访问，Wiktionary 的完整 IPA 未明确标口音。


访问失败只表示未取得证据。现有录音可以与英式音标并存，标签描述转写口音。录音的实际读音及播放情况仍需另行验证。

### Notes 拟追加原文

只追加同编号的一行，与已有 Notes 使用空行分隔，保留原文。若实施前发现相同来源行或同口音的等价证据已记录，复用原出处。

- **B07-P01** `Pronunciation source: Collins Dictionary — alert; American English. https://www.collinsdictionary.com/dictionary/english/alert (accent: US; accessed 2026-11-22).`
- **B07-P02** `Pronunciation source: Collins Dictionary — argument; American English. https://www.collinsdictionary.com/dictionary/english/argument (accent: US; accessed 2026-11-22).`
- **B07-P03** `Pronunciation source: Wiktionary — boots; English Pronunciation, US. https://en.wiktionary.org/w/index.php?title=boots&oldid=93681493 (accent: US; accessed 2026-11-22).`
- **B07-P04** `Pronunciation source: Collins Dictionary — breadcrumb; American English. https://www.collinsdictionary.com/dictionary/english/breadcrumb (accent: US; accessed 2026-11-22).`
- **B07-P05** `Pronunciation source: Collins Dictionary — breakpoint; American English. https://www.collinsdictionary.com/dictionary/english/breakpoint (accent: US; accessed 2026-11-22).`
- **B07-P06** `Pronunciation source: Collins Dictionary — card; American English. https://www.collinsdictionary.com/dictionary/english/card (accent: US; accessed 2026-11-22).`
- **B07-P07** `Pronunciation source: Collins Dictionary — collapse; American English. https://www.collinsdictionary.com/dictionary/english/collapse (accent: US; accessed 2026-11-22).`
- **B07-P08** `Pronunciation source: Collins Dictionary — column; American English. https://www.collinsdictionary.com/dictionary/english/column (accent: US; accessed 2026-11-22).`
- **B07-P09** `Pronunciation source: Collins Dictionary — deadlock; American English. https://www.collinsdictionary.com/dictionary/english/deadlock (accent: US; accessed 2026-11-22).`
- **B07-P10** `Pronunciation source: Collins Dictionary — declaration; American English. https://www.collinsdictionary.com/dictionary/english/declaration (accent: US; accessed 2026-11-22).`
- **B07-P11** `Pronunciation source: Wiktionary — disguised; English Pronunciation, General American. https://en.wiktionary.org/w/index.php?title=disguised&oldid=83217429 (accent: US; accessed 2026-11-22).`
- **B07-P12** `Pronunciation source: Wiktionary — durability; English Pronunciation, US. https://en.wiktionary.org/w/index.php?title=durability&oldid=94553680 (accent: US; accessed 2026-11-22).`
- **B07-P13** `Pronunciation source: Collins Dictionary — embedding; American English. https://www.collinsdictionary.com/dictionary/english/embedding (accent: US; accessed 2026-11-22).`
- **B07-P14** `Pronunciation source: Wiktionary — enhancement; English Pronunciation, General American. https://en.wiktionary.org/w/index.php?title=enhancement&oldid=90606147 (accent: US; accessed 2026-11-22).`
- **B07-P15** `Pronunciation source: Collins Dictionary — generic; American English. https://www.collinsdictionary.com/dictionary/english/generic (accent: US; accessed 2026-11-22).`
- **B07-P16** `Pronunciation source: Collins Dictionary — interface; American English. https://www.collinsdictionary.com/dictionary/english/interface (accent: US; accessed 2026-11-22).`
- **B07-P17** `Pronunciation source: Collins Dictionary — modal; American English. https://www.collinsdictionary.com/dictionary/english/modal (accent: US; accessed 2026-11-22).`
- **B07-P18** `Pronunciation source: Collins Dictionary — narrowing; British English. https://www.collinsdictionary.com/dictionary/english/narrowing (accent: UK; accessed 2026-11-22).`
- **B07-P19** `Pronunciation source: Collins Dictionary — nesting; British English. https://www.collinsdictionary.com/dictionary/english/nesting (accent: UK; accessed 2026-11-22).`
- **B07-P20** `Pronunciation source: Collins Dictionary — pagination; American English. https://www.collinsdictionary.com/dictionary/english/pagination (accent: US; accessed 2026-11-22).`
- **B07-P21** `Pronunciation source: Collins Dictionary — promise; American English. https://www.collinsdictionary.com/dictionary/english/promise (accent: US; accessed 2026-11-22).`
- **B07-P22** `Pronunciation source: Collins Dictionary — specificity; American English. https://www.collinsdictionary.com/dictionary/english/specificity (accent: US; accessed 2026-11-22).`
- **B07-P23** `Pronunciation source: Collins Dictionary — symbol; American English. https://www.collinsdictionary.com/dictionary/english/symbol (accent: US; accessed 2026-11-22).`
- **B07-P24** `Pronunciation source: Wiktionary — tuple; English Pronunciation, US. https://en.wiktionary.org/w/index.php?title=tuple&oldid=89359660 (accent: US; accessed 2026-11-22).`
- **B07-P25** `Pronunciation source: Collins Dictionary — warp; American English. https://www.collinsdictionary.com/dictionary/english/warp (accent: US; accessed 2026-11-22).`

## 本批暂缓

| 稳定 ID | 原因 |
|:--|:--|
| [accessor-tech](../entries/accessor-tech.yml) | Collins 词页未能读取。 |
| [batching-react](../entries/batching-react.yml) | Collins 未取得 batching 独立分区或完整 batching IPA。 |
| [coordinates-ui](../entries/coordinates-ui.yml) | Collins 带读音的分区是配套服饰义。暂缺坐标义的完整复数证据。 |
| [corrupted-combat](../entries/corrupted-combat.yml) | Collins 未取得明确美式整词 IPA。Wiktionary 的 IPA 未明确标口音，另有英格兰南部音频。 |
| [dropdown-bootstrap](../entries/dropdown-bootstrap.yml) | Collins 词页未能读取。 |
| [hoisting-javascript](../entries/hoisting-javascript.yml) | Collins 未取得可匹配的美式整词分区。 |
| [holdings-finance](../entries/holdings-finance.yml) | Collins 提供金融用例，但未取得完整 holdings IPA。 |
| [hydration-react](../entries/hydration-react.yml) | Collins 未取得明确的美式整词分区。 |
| [iterable-tech](../entries/iterable-tech.yml) | Collins 词页未能读取。 |
| [keyframe-css](../entries/keyframe-css.yml) | Collins 词页未能读取。 |
| [namespace-typescript](../entries/namespace-typescript.yml) | Collins 词页未能读取。 |
| [noodles-food](../entries/noodles-food.yml) | Collins 与 Cambridge 词页未能读取。Wiktionary 完整 IPA 未明确标口音。 |
| [orchestration-ai](../entries/orchestration-ai.yml) | Collins 音标未明确标口音。Cambridge 和 Oxford 请求失败。Wiktionary 匹配到的 IPA 属于法语。 |
| [popover-bootstrap](../entries/popover-bootstrap.yml) | Collins 找到食品义的音标，尚未核实 UI 义可用的转写。 |
| [ranged-combat](../entries/ranged-combat.yml) | Collins 美式音标所在分区是建筑义，暂不用于“远程”。 |
| [soundness-type-system](../entries/soundness-type-system.yml) | Collins 词页未能读取。 |
| [stylesheet-css](../entries/stylesheet-css.yml) | Collins 的该拼法词页未能读取。 |
| [viewport-css](../entries/viewport-css.yml) | Collins 词页未能读取。 |


复数和分词均要求整词证据。prefix-priority-nginx 留在符号专项，当前也没有录音。

## 批准后实施与检查

1. 重读批准范围和当前 YAML，核对稳定 ID、词形、释义、已有音标和 Notes。保留前批改动，快照词条、MP3 哈希与当前导出。
2. 只写入批准的 pronunciation 和对应 Notes 来源行。美式无前缀，英式保留 UK。保留其他字段和现有录音。
3. 执行 go run scripts/validate/main.go、go run scripts/build/main.go 和 git diff --check。通过构建器生成 dist。核对 18 列导出的发音栏及原 StudyOrder、Audio、CodeExample 和关联。更新计划状态、实际数量和验证边界。
4. 实施范围限于本地词条和导出。提交、推送、生成录音与操作 Anki 须另行授权。

## 规划验证

核对 25 个提案、18 个暂缓 ID、当前空字段、分组计数、本地链接和格式。全库原有文件的哈希与规划基线匹配。构建、音频生成和真机验证留待各自获批后进行。

## 实施记录

已实施用户批准的 B07-P01–B07-P25。25 个既有英语词条新增 pronunciation 与对应出处，其中 23 项为美式、2 项明确标注 UK。英语音标缺口从 608 降至 583，包含本批暂缓的 18 项。243 个日语读法保持原样。

词库验证、两次导出构建及 Git 差异检查均通过。两次导出字节一致，共 1833 行、18 列。与实施前相比，导出仅在获批 ID 的 Pronunciation 和 Notes 列发生变化。词条检查确认保留原 Notes 内容和其他字段，全部 1875 个既有 MP3 哈希匹配。未进行人工听审或实际 Anki 播放验证。
