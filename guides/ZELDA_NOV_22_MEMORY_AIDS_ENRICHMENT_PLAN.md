# 塞尔达词条记忆辅助审阅计划

状态为已批准并实施。规划与实施日期为 2026-11-22。用户明确批准了全部文本和关联提案。下文保留规划基线，实际结果见文末。

## 范围与基线

本轮以当前 YAML 的 `zelda-totk` 标签确定词条范围，并使用稳定 ID 记账。通用词只要带有该主题标签，也纳入审阅。单独标为 `nintendo-switch` 的硬件与连接词条保留在全库关联候选范围内。范围内没有日语词条，本轮不拟写假名或日语词源。

全库使用当前词条的语言、Word、Meaning、Context、Example、Notes、三种辅助字段、先修和关联进行筛查。优先选择能复用的构词规律、能解释当前义项的历史线索，以及需要辨认的易混概念。普通知识已表达清楚、已有辨析足够或证据不足时，本轮留空。筛查不等同于逐个考证历史词源。

以当前实现为准。解析器接受可选 `etymology`、`word_formation` 和 `mnemonic`，关联来自 `related_entries`，导出按稳定 ID 排列直接邻居。跨语言显示会附语言标记。卡背仅在相应字段有内容时显示栏目。参见 [仓库约定](../AGENTS.md)、[模板](../README.md)、[字段定义](ENTRY_SCHEMA.md) 和 [学习顺序](LEARNING_ORDER.md)。

| 项目                       | 计数                                                                                                |
|:----------------------------|:-----------------------------------------------------------------------------------------------------|
| 全库筛查                   | 1833 项，英语 1590 项、日语 243 项                                                                  |
| 塞尔达主题                 | 613 项，全部为英语                                                                                  |
| 现有 ETY / WF / N          | 18 / 54 / 42 个非空值，全部保留                                                                     |
| 现有辅助的唯一词条数       | 109                                                                                                 |
| 现有直接关联               | 32 对涉及主题词条，135 对全库                                                                       |
| 文本提案                   | 8 WF + 2 ETY + 3 N，共 13 个辅助字段，含 1 项移入                                                   |
| Notes 追加与移出           | 10 个词条各追加一条词典出处，其中 1 个词条另移出一句原 Notes。共触及 10 个 Notes 字段 |
| 文本提案涉及词条           | 13                                                                                                  |
| 待批关联                   | 11 对，涉及 20 个端点。在 11 个词条中追加 ID，派生 11 条反向展示，共 22 条新增邻接展示              |
| 文本或关联涉及的唯一词条   | 26 项，其中 24 项在主题内。文本与关联重叠 7 项                                                      |
| 实际拟写入 YAML 的唯一词条 | 19                                                                                                  |
| 主题台账处理               | 提案或反向显示 24 项，暂缓 14 项，保留 136 项，本轮留空 439 项                                      |

若全部批准，将修改 34 个 YAML 字段值，包括 13 个记忆辅助字段、10 个 Notes 字段和 11 个关联列表。一次移出与一次追加可同属一个 Notes 字段。反向显示由导出生成。

## 保留与暂缓

台账的“既有辅助”栏逐条标记需原样保留的 WF、ETY 和 N，“既有邻居”栏包括入站派生的显示。两份已完成的历史记录只提供背景。 [早期增强记录](TOTK_MEMORY_AIDS_ENRICHMENT_PLAN.md) · [后续词汇批次](NINTENDO_SWITCH_TOTK_VOCABULARY_IMPORT_PLAN.md)

保留 spring、pepper、stable、depths 等词条已经写明的辨析。唯一建议移出的 Notes 句子列于 WF08。专名如 Sundelion、Zonaite 和 draconification 保留现有拼法、译名及出处。游戏文本可支持名称和文本键，未提供官方命名动机的证据。名称来源和机制细节在台账中单列暂缓。

## WF 构词提案

以下卡片文本为纯文本。现代构词或词形对应的证据在各条中标出。

### WF01 bottomless

[bottomless-general](../entries/bottomless-general.yml) · bottomless · 无底的；深不见底的

新增 `word_formation`。

```plain
bottom（底部）+ -less（没有……的）。用来形容深坑时，可表示深到看不见底。
```

复用 -less 理解缺少某物的形容词。 只解释当前义项。 证据见 D01。

在完整 Notes 末尾用一个空行分隔并追加以下单独一行。若审批后重读发现相同行或同证据 URL 已存在，保留原出处并跳过重复追加。

```plain
Memory-aid source: Merriam-Webster.com Dictionary — bottomless. https://www.merriam-webster.com/dictionary/bottomless
```

### WF02 reattach

[reattach-general](../entries/reattach-general.yml) · reattach · 重新连接；重新装上

新增 `word_formation`。

```plain
re-（再次）+ attach（连接、附上）。先拆下某个部件，再装回去，就可以用 reattach。
```

用动作顺序记住 re-。 说明普通动词的构成，范围限于连接义。 证据见 D02。

在完整 Notes 末尾用一个空行分隔并追加以下单独一行。若审批后重读发现相同行或同证据 URL 已存在，保留原出处并跳过重复追加。

```plain
Memory-aid source: Merriam-Webster.com Dictionary — reattach. https://www.merriam-webster.com/dictionary/reattach
```

### WF03 updraft

[updraft-flight](../entries/updraft-flight.yml) · updraft · 上升气流

新增 `word_formation`。

```plain
up（向上）+ draft（气流）。这里的 draft 取气流义，整个词表示向上流动的空气。
```

借熟悉的 up 辨认方向，并区分 draft 的草稿义。 现代复合词拆解，气流义以完整词的词典定义为据。 证据见 D03。

在完整 Notes 末尾用一个空行分隔并追加以下单独一行。若审批后重读发现相同行或同证据 URL 已存在，保留原出处并跳过重复追加。

```plain
Memory-aid source: Merriam-Webster.com Dictionary — updraft. https://www.merriam-webster.com/dictionary/updraft
```

### WF04 crystallize

[crystallize-general](../entries/crystallize-general.yml) · crystallize · 结晶；使形成晶体

新增 `word_formation`。

```plain
crystal（晶体）+ -ize（使成为或变成）。crystallize 可以表示使物质结晶，也可以表示物质自身结晶。
```

从名词推到动词，并记住两种句型。 范围限于晶体义。 证据见 D04。

在完整 Notes 末尾用一个空行分隔并追加以下单独一行。若审批后重读发现相同行或同证据 URL 已存在，保留原出处并跳过重复追加。

```plain
Memory-aid source: Merriam-Webster.com Dictionary — crystallize. https://www.merriam-webster.com/dictionary/crystallize
```

### WF05 crystallized

[crystallized-general](../entries/crystallized-general.yml) · crystallized · 结晶的；已形成晶体的

新增 `word_formation`。

```plain
crystallize 的词尾已经有 e，直接加 -d 得到 crystallized。这里用过去分词作形容词，描述已经形成晶体的状态。
```

把动作和结果状态分开，同时掌握拼写。 说明词条当前形容词用法。 证据见 D04。

在完整 Notes 末尾用一个空行分隔并追加以下单独一行。若审批后重读发现相同行或同证据 URL 已存在，保留原出处并跳过重复追加。

```plain
Memory-aid source: Merriam-Webster.com Dictionary — crystallize. https://www.merriam-webster.com/dictionary/crystallize
```

### WF06 disassemble

[disassemble-general](../entries/disassemble-general.yml) · disassemble · 拆卸；拆解

新增 `word_formation`。

```plain
dis- + assemble。这里 dis- 表示逆转 assemble（组装）的动作。同一词族的名词 disassembly 表示拆卸。
```

通过组装与拆卸的反向动作记前缀，再认出名词形式。 只解释此词中 dis- 的作用。既有 Notes 的 decompose 辨析原样保留。 证据见 D05。

在完整 Notes 末尾用一个空行分隔并追加以下单独一行。若审批后重读发现相同行或同证据 URL 已存在，保留原出处并跳过重复追加。

```plain
Memory-aid source: Merriam-Webster.com Dictionary — disassemble. https://www.merriam-webster.com/dictionary/disassemble
```

### WF07 descent

[descent-flight](../entries/descent-flight.yml) · descent · 下降；下降过程

新增 `word_formation`。

```plain
descend 是动词，descent 是表示下降动作或过程的名词。记住结尾的对应，descend 以 -d 结尾，descent 以 -t 结尾。
```

从同词族看一眼区分词性和拼写。 现代词形对照，当前只学习物理下降义。 证据见 D06。

在完整 Notes 末尾用一个空行分隔并追加以下单独一行。若审批后重读发现相同行或同证据 URL 已存在，保留原出处并跳过重复追加。

```plain
Memory-aid source: Merriam-Webster.com Dictionary — descent. https://www.merriam-webster.com/dictionary/descent
```

### WF08 phenomenon

[phenomenon-story](../entries/phenomenon-story.yml) · phenomenon · 现象

移入并扩写至 `word_formation`。

```plain
phenomenon 是单数，phenomena 是常用的不规则复数。Regional Phenomena 的末词就是复数。
```

把游戏任务名与可迁移的单复数语法连起来。 适用于本词和当前任务名。词典另有边缘用法，本卡聚焦常规用法。 证据见 D07。

从现有 Notes 移出的完整一行如下。其前后的来源标题、其余文字和原有空行均保留。

```plain
phenomena 是 phenomenon 的不规则复数。
```

在完整 Notes 末尾用一个空行分隔并追加以下单独一行。若审批后重读发现相同行或同证据 URL 已存在，保留原出处并跳过重复追加。

```plain
Memory-aid source: Merriam-Webster.com Dictionary — phenomenon. https://www.merriam-webster.com/dictionary/phenomenon
```

## ETY 词源提案

历史链以已读取的词源词条正文为依据。

### ETY01 archipelago

[archipelago-terrain](../entries/archipelago-terrain.yml) · archipelago · 群岛

新增 `etymology`。

```plain
英语 archipelago 借自原指爱琴海的意大利语 arcipelago。词义后来扩展到遍布岛屿的海域，再指岛屿本身组成的群岛。关于意大利语形式的形成，有意大利语复合和由中世纪拉丁语 Egeopelagus 演变两种解释。
```

用爱琴海的地理画面记住从海域到群岛的语义变化。 保留两种解释。 证据见 D08。

在完整 Notes 末尾用一个空行分隔并追加以下单独一行。若审批后重读发现相同行或同证据 URL 已存在，保留原出处并跳过重复追加。

```plain
Memory-aid source: Online Etymology Dictionary — archipelago. https://www.etymonline.com/word/archipelago
```

### ETY02 colossal

[colossal-combat](../entries/colossal-combat.yml) · colossal · 极其庞大的

新增 `etymology`。

```plain
英语 colossal 借自法语 colossal，法语词来自 colosse（巨像），可经拉丁语 colossus 追溯到希腊语 kolossos（巨像）。希腊语词的更早来源未明。
```

把强烈的尺寸形容词与巨像联系起来。 只说明英语普通词的历史。 证据见 D09。

在完整 Notes 末尾用一个空行分隔并追加以下单独一行。若审批后重读发现相同行或同证据 URL 已存在，保留原出处并跳过重复追加。

```plain
Memory-aid source: Online Etymology Dictionary — colossal. https://www.etymonline.com/word/colossal
```

## N 助记提案

“记忆画面”是本计划设计的学习辅助。现有词条仅支持措辞中的事实。

### N01 quicksave

[quicksave-game](../entries/quicksave-game.yml) · quicksave · 快速存档；通过快捷操作保存的进度

新增 `mnemonic`。

```plain
记忆画面：游戏界面上有三个标签。QUICK 提醒你用快捷操作，MANUAL 提醒你主动保存，AUTO 提醒你由系统触发。前两项描述的操作可以重叠。
```

用速度或操作方式与触发者两个维度区分三词。 原创提示画面。三种叫法取自现有义项，具体游戏的键位、菜单或存档槽待实测。 证据见 L02。

依据为现有词条字段。Notes 原样保留，无额外来源行。

### N02 Gloom Recovery

[gloom-recovery-zelda](../entries/gloom-recovery-zelda.yml) · Gloom Recovery · 瘴气伤害回复；《王国之泪》中的能力效果名称。

新增 `mnemonic`。

```plain
记忆画面：盾牌标 Resistance，修补贴标 Recovery。先想到防护，再想到损伤后的回复。
```

用防护与回复的顺序区分两个效果名。 原创图像。两个标签仅对应当前释义，图中的道具不代表游戏机制。 证据见 L03。

依据为现有词条字段。Notes 原样保留，无额外来源行。

### N03 Energy Well

[energy-well-zelda](../entries/energy-well-zelda.yml) · Energy Well · 附加电池；提高电池容量的道具

新增 `mnemonic`。

```plain
记忆画面：随身电池旁多出一格空容量。Energy Cell 是供能的电池，Energy Well 是让容量增加的那一格。
```

区分电池本体与额外容量。 原创图像，依据当前两词释义；图中的格子只是记忆提示。 证据见 L04。

依据为现有词条字段。Notes 原样保留，无额外来源行。

## R 关联提案

下表左端是唯一拟写入的词条，按列出的编号顺序在 `related_entries` 追加右端 ID。保留原有列表成员。逐对比较全库双向声明后确认无重复。来自先修的学习顺序保持原样。

跨主题右端有 `pepper-food` 和 `hook-tech`，仅在导出中增加反向显示。这两个词条的 YAML 不需要调整。全库也筛查了其他语言的词条，当前拟议关联均为英语。关联审批与助记审批相互独立。

| 编号 | 声明端 ID、Word、Meaning                                                                                                         | 目标 ID、Word、Meaning                                                                                                        | 学习价值                           | 易混点                                       |
|:------|:----------------------------------------------------------------------------------------------------------------------------------|:-------------------------------------------------------------------------------------------------------------------------------|:------------------------------------|:----------------------------------------------|
| R01  | [manual-save-game](../entries/manual-save-game.yml) · manual save · 手动存档；由玩家主动保存的进度                               | [auto-save-game](../entries/auto-save-game.yml) · auto-save · 自动存档；系统自动保存的进度                                    | 按触发者对照                       | 主动操作与系统触发容易混为一谈。             |
| R02  | [quicksave-game](../entries/quicksave-game.yml) · quicksave · 快速存档；通过快捷操作保存的进度                                   | [manual-save-game](../entries/manual-save-game.yml) · manual save · 手动存档；由玩家主动保存的进度                            | 按快捷方式与主动触发对照           | 快速存档可能也是手动触发。                   |
| R03  | [gloom-recovery-zelda](../entries/gloom-recovery-zelda.yml) · Gloom Recovery · 瘴气伤害回复；《王国之泪》中的能力效果名称。      | [gloom-resistance-zelda](../entries/gloom-resistance-zelda.yml) · Gloom Resistance · 瘴气防护；《王国之泪》中的能力效果名称。 | 区分回复与防护                     | 不能从名字推断具体数值或触发条件。           |
| R04  | [perfect-guard-zelda](../entries/perfect-guard-zelda.yml) · Perfect Guard · 完美防御；在攻击命中前一刻用盾牌弹开攻击             | [perfect-dodge-zelda](../entries/perfect-dodge-zelda.yml) · Perfect Dodge · 完美回避；在恰当时机跳开躲过攻击                  | 同为讲究时机的不同应对动作         | 格挡用盾，应对动作不同于跳避。               |
| R05  | [energy-cell-zelda](../entries/energy-cell-zelda.yml) · Energy Cell · 电池；为使用中的左纳乌装置供能的随身道具                   | [energy-well-zelda](../entries/energy-well-zelda.yml) · Energy Well · 附加电池；提高电池容量的道具                            | 区分供能的随身道具与提升容量的道具 | 两者的中文都含电池，容易当作同物。           |
| R06  | [crystallized-charge-zelda](../entries/crystallized-charge-zelda.yml) · Crystallized Charge · 左纳乌能源结晶；制作附加电池的材料 | [energy-well-zelda](../entries/energy-well-zelda.yml) · Energy Well · 附加电池；提高电池容量的道具                            | 对照材料与用途                     | 当前释义支持制作关系，具体兑换数量无关本次。 |
| R07  | [disassemble-general](../entries/disassemble-general.yml) · disassemble · 拆卸；拆解                                             | [decompose-general](../entries/decompose-general.yml) · decompose · 分解；腐解                                                | 区分拆开组装物与物质分解           | 两词都可译拆解或分解，先看宾语和过程。       |
| R08  | [pepper-vegetable](../entries/pepper-vegetable.yml) · pepper · 辣椒；甜椒；这里不指胡椒粉。                                      | [pepper-food](../entries/pepper-food.yml) · pepper · 胡椒；胡椒粉                                                             | 区分作蔬菜的辣椒或甜椒与调味胡椒   | 同形词，释义和常用搭配不同。                 |
| R09  | [hook-hardware](../entries/hook-hardware.yml) · hook · 挂钩；钩状物                                                              | [hook-tech](../entries/hook-tech.yml) · hook · 钩子；在预定生命周期事件发生时触发配置动作的扩展点。                           | 用实体挂点帮助理解生命周期挂接点   | 比喻只用于记义，技术钩子依赖宿主事件和配置。 |
| R10  | [descent-flight](../entries/descent-flight.yml) · descent · 下降；下降过程                                                       | [descend-general](../entries/descend-general.yml) · descend · 下降                                                            | 对照名词过程与动词动作             | 有先修关系，但 Related Words 需要独立批准。  |
| R11  | [crystallized-general](../entries/crystallized-general.yml) · crystallized · 结晶的；已形成晶体的                                | [crystallize-general](../entries/crystallize-general.yml) · crystallize · 结晶；使形成晶体                                    | 区分结果状态与动词动作             | 过去分词形式不等于本词条学习的动词词头。     |

## 来源登记

D01–D09 在 2026-11-22 通过词典正文查证。学习画面由本计划编写，来源仅支持词义或构词事实。游戏引用沿用已记录的固定版本线索，本轮未重新验证游戏数据或测试机制。

| 编号 | 来源                                                                                                                                                                                | 支持范围                                                                         |
|:------|:-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|:----------------------------------------------------------------------------------|
| D01  | [Merriam-Webster.com Dictionary — bottomless](https://www.merriam-webster.com/dictionary/bottomless)                                                                                | 支持无底、极深两种相关用法。                                                     |
| D02  | [Merriam-Webster.com Dictionary — reattach](https://www.merriam-webster.com/dictionary/reattach)                                                                                    | 支持再次连接的动词义。                                                           |
| D03  | [Merriam-Webster.com Dictionary — updraft](https://www.merriam-webster.com/dictionary/updraft)                                                                                      | 支持向上的气流义。                                                               |
| D04  | [Merriam-Webster.com Dictionary — crystallize](https://www.merriam-webster.com/dictionary/crystallize)                                                                              | 支持 crystal + -ize、词形 crystallized 以及形成晶体的义项。                      |
| D05  | [Merriam-Webster.com Dictionary — disassemble](https://www.merriam-webster.com/dictionary/disassemble)                                                                              | 支持拆开的动词义与名词 disassembly。                                             |
| D06  | [Merriam-Webster.com Dictionary — descent](https://www.merriam-webster.com/dictionary/descent)                                                                                      | 支持下降过程的名词义。                                                           |
| D07  | [Merriam-Webster.com Dictionary — phenomenon](https://www.merriam-webster.com/dictionary/phenomenon)                                                                                | 支持 phenomenon 的常用复数 phenomena。                                           |
| D08  | [Online Etymology Dictionary — archipelago](https://www.etymonline.com/word/archipelago)                                                                                            | 支持意大利语、爱琴海和群岛的语义演变，列出构词与早期来源的不同解释。             |
| D09  | [Online Etymology Dictionary — colossal](https://www.etymonline.com/word/colossal)                                                                                                  | 支持法语、拉丁语和希腊语的传承链，希腊语词的更早来源未明。                       |
| L01  | 当前 YAML 全库及全库入站与出站关联                                                                                                                                                  | 支持当前义项、配对理由、已有先修和关联排除。相关 ID 已在各提案中链接。           |
| L02  | [quicksave-game](../entries/quicksave-game.yml), [manual-save-game](../entries/manual-save-game.yml), [auto-save-game](../entries/auto-save-game.yml)                               | 三个存档义项的操作方式和触发来源，未验证具体游戏的可用模式。                     |
| L03  | [gloom-resistance-zelda](../entries/gloom-resistance-zelda.yml), [gloom-recovery-zelda](../entries/gloom-recovery-zelda.yml)                                                        | 当前两个效果名的释义。Notes 引用 1.2.1 MiasmaGuard 与 LifeRepair；结论限于释义。 |
| L04  | [energy-cell-zelda](../entries/energy-cell-zelda.yml), [energy-well-zelda](../entries/energy-well-zelda.yml), [crystallized-charge-zelda](../entries/crystallized-charge-zelda.yml) | 各 YAML 的道具含义与材料用途，未核实格数、兑换数量或触发方式。                   |

## 完整审阅台账

以下每个稳定 ID 出现一次。辅助栏目列出现有非空字段，任何新提案都以编号表示。既有辅助与关联可以同时存在，均保留。“本轮留空”表示未选中补充，不能据此推断词源或未来学习价值。

| 稳定 ID                                                                                 | Word                         | 既有辅助 | 既有邻居 | 决定                                   |
|:-----------------------------------------------------------------------------------------|:------------------------------|:----------|:----------|:----------------------------------------|
| [abandoned-workflow](../entries/abandoned-workflow.yml)                                 | abandoned                    | —        | 0        | 本轮留空                               |
| [activate-general](../entries/activate-general.yml)                                     | activate                     | —        | 0        | 本轮留空                               |
| [adventure-general](../entries/adventure-general.yml)                                   | adventure                    | —        | 0        | 本轮留空                               |
| [adventure-log-zelda](../entries/adventure-log-zelda.yml)                               | Adventure Log                | —        | 0        | 本轮留空                               |
| [aerocuda-zelda](../entries/aerocuda-zelda.yml)                                         | Aerocuda                     | —        | 0        | 暂缓。命名史待证                       |
| [aggressive-combat](../entries/aggressive-combat.yml)                                   | aggressive                   | —        | 1        | 保留已有辅助                           |
| [aim-general](../entries/aim-general.yml)                                               | aim                          | —        | 0        | 本轮留空                               |
| [airborne-combat](../entries/airborne-combat.yml)                                       | airborne                     | WF       | 0        | 保留已有辅助                           |
| [album-ui](../entries/album-ui.yml)                                                     | album                        | —        | 0        | 本轮留空                               |
| [altar-architecture](../entries/altar-architecture.yml)                                 | altar                        | —        | 0        | 本轮留空                               |
| [ambitious-general](../entries/ambitious-general.yml)                                   | ambitious                    | —        | 0        | 本轮留空                               |
| [ancestor-general](../entries/ancestor-general.yml)                                     | ancestor                     | —        | 1        | 保留已有辅助                           |
| [ancient-arowana-zelda](../entries/ancient-arowana-zelda.yml)                           | Ancient Arowana              | —        | 0        | 本轮留空                               |
| [ancient-general](../entries/ancient-general.yml)                                       | ancient                      | —        | 0        | 本轮留空                               |
| [apple-food](../entries/apple-food.yml)                                                 | apple                        | —        | 0        | 本轮留空                               |
| [aquatic-combat](../entries/aquatic-combat.yml)                                         | aquatic                      | ETY      | 0        | 保留已有辅助                           |
| [archipelago-terrain](../entries/archipelago-terrain.yml)                               | archipelago                  | —        | 0        | ETY01                                  |
| [armor-equipment](../entries/armor-equipment.yml)                                       | armor                        | —        | 1        | 保留已有辅助                           |
| [armored-combat](../entries/armored-combat.yml)                                         | armored                      | WF       | 0        | 保留已有辅助                           |
| [armory-architecture](../entries/armory-architecture.yml)                               | armory                       | —        | 1        | 保留已有辅助                           |
| [arowana-general](../entries/arowana-general.yml)                                       | arowana                      | —        | 0        | 本轮留空                               |
| [arrow-weapon](../entries/arrow-weapon.yml)                                             | arrow                        | —        | 0        | 本轮留空                               |
| [ascend-general](../entries/ascend-general.yml)                                         | ascend                       | —        | 1        | 保留已有辅助                           |
| [assemble-general](../entries/assemble-general.yml)                                     | assemble                     | —        | 0        | 本轮留空                               |
| [attach-general](../entries/attach-general.yml)                                         | attach                       | —        | 0        | 本轮留空                               |
| [attack-power-combat](../entries/attack-power-combat.yml)                               | attack power                 | WF       | 0        | 保留已有辅助                           |
| [attack-up-zelda](../entries/attack-up-zelda.yml)                                       | Attack Up                    | —        | 0        | 本轮留空                               |
| [auto-save-game](../entries/auto-save-game.yml)                                         | auto-save                    | —        | 0        | R01（反向）                            |
| [autobuild-zelda](../entries/autobuild-zelda.yml)                                       | Autobuild                    | WF       | 0        | 保留已有辅助                           |
| [awaken-story](../entries/awaken-story.yml)                                             | awaken                       | —        | 0        | 本轮留空                               |
| [axe-weapon](../entries/axe-weapon.yml)                                                 | axe                          | —        | 0        | 本轮留空                               |
| [backflip-movement](../entries/backflip-movement.yml)                                   | backflip                     | —        | 0        | 本轮留空                               |
| [balloon-engineering](../entries/balloon-engineering.yml)                               | balloon                      | —        | 0        | 本轮留空                               |
| [banana-food](../entries/banana-food.yml)                                               | banana                       | —        | 0        | 本轮留空                               |
| [bandanna-clothing](../entries/bandanna-clothing.yml)                                   | bandanna                     | —        | 0        | 本轮留空                               |
| [bass-food](../entries/bass-food.yml)                                                   | bass                         | N        | 0        | 保留已有辅助                           |
| [battery-engineering](../entries/battery-engineering.yml)                               | battery                      | —        | 0        | 本轮留空                               |
| [beam-emitter-zelda](../entries/beam-emitter-zelda.yml)                                 | Beam Emitter                 | WF       | 0        | 保留已有辅助                           |
| [beam-engineering](../entries/beam-engineering.yml)                                     | beam                         | —        | 0        | 本轮留空                               |
| [bee-general](../entries/bee-general.yml)                                               | bee                          | —        | 0        | 本轮留空                               |
| [big-battery-zelda](../entries/big-battery-zelda.yml)                                   | Big Battery                  | —        | 0        | 本轮留空                               |
| [big-wheel-zelda](../entries/big-wheel-zelda.yml)                                       | Big Wheel                    | —        | 0        | 本轮留空                               |
| [blessing-general](../entries/blessing-general.yml)                                     | blessing                     | —        | 0        | 本轮留空                               |
| [block-general](../entries/block-general.yml)                                           | block                        | —        | 0        | 本轮留空                               |
| [bokoblin-zelda](../entries/bokoblin-zelda.yml)                                         | Bokoblin                     | N        | 0        | 保留已有辅助                           |
| [bomb-weapon](../entries/bomb-weapon.yml)                                               | bomb                         | —        | 0        | 本轮留空                               |
| [boomerang-weapon](../entries/boomerang-weapon.yml)                                     | boomerang                    | —        | 0        | 本轮留空                               |
| [boots-clothing](../entries/boots-clothing.yml)                                         | boots                        | —        | 0        | 本轮留空                               |
| [boss-bokoblin-zelda](../entries/boss-bokoblin-zelda.yml)                               | Boss Bokoblin                | —        | 0        | 本轮留空                               |
| [bottomless-general](../entries/bottomless-general.yml)                                 | bottomless                   | —        | 0        | WF01                                   |
| [bow-weapon](../entries/bow-weapon.yml)                                                 | bow                          | —        | 0        | 本轮留空                               |
| [brave-general](../entries/brave-general.yml)                                           | brave                        | —        | 2        | 保留已有辅助                           |
| [bridge-architecture](../entries/bridge-architecture.yml)                               | bridge                       | —        | 0        | 本轮留空                               |
| [brightbloom-seed-zelda](../entries/brightbloom-seed-zelda.yml)                         | Brightbloom Seed             | —        | 0        | 本轮留空                               |
| [brightcap-zelda](../entries/brightcap-zelda.yml)                                       | Brightcap                    | —        | 0        | 本轮留空                               |
| [broadsword-weapon](../entries/broadsword-weapon.yml)                                   | broadsword                   | WF       | 0        | 保留已有辅助                           |
| [build-construction](../entries/build-construction.yml)                                 | build                        | —        | 0        | 本轮留空                               |
| [butter-food](../entries/butter-food.yml)                                               | butter                       | —        | 0        | 本轮留空                               |
| [calamity-story](../entries/calamity-story.yml)                                         | calamity                     | —        | 0        | 本轮留空                               |
| [campfire-general](../entries/campfire-general.yml)                                     | campfire                     | —        | 0        | 本轮留空                               |
| [cancel-general](../entries/cancel-general.yml)                                         | cancel                       | —        | 0        | 本轮留空                               |
| [cannon-engineering](../entries/cannon-engineering.yml)                                 | cannon                       | —        | 0        | 本轮留空                               |
| [canyon-terrain](../entries/canyon-terrain.yml)                                         | canyon                       | —        | 1        | 保留已有辅助                           |
| [capsule-general](../entries/capsule-general.yml)                                       | capsule                      | —        | 0        | 本轮留空                               |
| [captain-construct-zelda](../entries/captain-construct-zelda.yml)                       | Captain Construct            | —        | 0        | 本轮留空                               |
| [carrot-food](../entries/carrot-food.yml)                                               | carrot                       | —        | 0        | 本轮留空                               |
| [cart-engineering](../entries/cart-engineering.yml)                                     | cart                         | —        | 0        | 本轮留空                               |
| [castle-architecture](../entries/castle-architecture.yml)                               | castle                       | —        | 0        | 本轮留空                               |
| [cave-architecture](../entries/cave-architecture.yml)                                   | cave                         | —        | 1        | 保留已有辅助                           |
| [cavern-architecture](../entries/cavern-architecture.yml)                               | cavern                       | N        | 1        | 保留已有辅助                           |
| [ceiling-architecture](../entries/ceiling-architecture.yml)                             | ceiling                      | —        | 0        | 本轮留空                               |
| [chamber-architecture](../entries/chamber-architecture.yml)                             | chamber                      | —        | 0        | 本轮留空                               |
| [champion-general](../entries/champion-general.yml)                                     | champion                     | —        | 0        | 本轮留空                               |
| [character-profiles-zelda](../entries/character-profiles-zelda.yml)                     | Character Profiles           | —        | 0        | 本轮留空                               |
| [charge-electricity](../entries/charge-electricity.yml)                                 | charge                       | —        | 0        | 本轮留空                               |
| [chasm-terrain](../entries/chasm-terrain.yml)                                           | chasm                        | N        | 1        | 保留已有辅助                           |
| [cheese-food](../entries/cheese-food.yml)                                               | cheese                       | —        | 0        | 本轮留空                               |
| [chilly-general](../entries/chilly-general.yml)                                         | chilly                       | —        | 0        | 本轮留空                               |
| [chuchu-zelda](../entries/chuchu-zelda.yml)                                             | Chuchu                       | —        | 0        | 本轮留空                               |
| [citadel-architecture](../entries/citadel-architecture.yml)                             | citadel                      | —        | 0        | 本轮留空                               |
| [civilization-general](../entries/civilization-general.yml)                             | civilization                 | —        | 0        | 本轮留空                               |
| [claymore-weapon](../entries/claymore-weapon.yml)                                       | claymore                     | —        | 0        | 本轮留空                               |
| [cliff-terrain](../entries/cliff-terrain.yml)                                           | cliff                        | —        | 0        | 本轮留空                               |
| [climb-general](../entries/climb-general.yml)                                           | climb                        | —        | 0        | 本轮留空                               |
| [climbing-gear-zelda](../entries/climbing-gear-zelda.yml)                               | Climbing Gear                | —        | 0        | 本轮留空                               |
| [closed-general](../entries/closed-general.yml)                                         | closed                       | —        | 0        | 本轮留空                               |
| [club-weapon](../entries/club-weapon.yml)                                               | club                         | —        | 0        | 本轮留空                               |
| [clue-story](../entries/clue-story.yml)                                                 | clue                         | —        | 0        | 本轮留空                               |
| [cold-resistance-combat](../entries/cold-resistance-combat.yml)                         | cold resistance              | —        | 1        | 保留已有辅助                           |
| [colgera-zelda](../entries/colgera-zelda.yml)                                           | Colgera                      | —        | 0        | 暂缓。命名史待证                       |
| [collectible-game](../entries/collectible-game.yml)                                     | collectible                  | —        | 0        | 本轮留空                               |
| [colossal-combat](../entries/colossal-combat.yml)                                       | colossal                     | —        | 0        | ETY02                                  |
| [compendium-general](../entries/compendium-general.yml)                                 | compendium                   | —        | 0        | 本轮留空                               |
| [construct-being-zelda](../entries/construct-being-zelda.yml)                           | Construct                    | —        | 0        | 本轮留空                               |
| [construct-factory-zelda](../entries/construct-factory-zelda.yml)                       | Construct Factory            | —        | 0        | 本轮留空                               |
| [construct-general](../entries/construct-general.yml)                                   | construct                    | —        | 0        | 本轮留空                               |
| [construct-head-zelda](../entries/construct-head-zelda.yml)                             | Construct Head               | WF       | 0        | 保留已有辅助                           |
| [cook-cooking](../entries/cook-cooking.yml)                                             | cook                         | —        | 0        | 本轮留空                               |
| [cooking-pot-cooking](../entries/cooking-pot-cooking.yml)                               | cooking pot                  | WF       | 0        | 保留已有辅助                           |
| [coordinates-ui](../entries/coordinates-ui.yml)                                         | coordinates                  | —        | 0        | 本轮留空                               |
| [core-general](../entries/core-general.yml)                                             | core                         | —        | 0        | 本轮留空                               |
| [corridor-architecture](../entries/corridor-architecture.yml)                           | corridor                     | —        | 0        | 本轮留空                               |
| [corrupted-combat](../entries/corrupted-combat.yml)                                     | corrupted                    | —        | 0        | 本轮留空                               |
| [courageous-general](../entries/courageous-general.yml)                                 | courageous                   | ETY、WF  | 2        | 保留已有辅助                           |
| [courser-bee-honey-zelda](../entries/courser-bee-honey-zelda.yml)                       | Courser Bee Honey            | —        | 0        | 本轮留空                               |
| [crater-terrain](../entries/crater-terrain.yml)                                         | crater                       | —        | 0        | 本轮留空                               |
| [crisis-story](../entries/crisis-story.yml)                                             | crisis                       | —        | 0        | 本轮留空                               |
| [critical-hit-combat](../entries/critical-hit-combat.yml)                               | critical hit                 | WF       | 0        | 保留已有辅助                           |
| [crouch-general](../entries/crouch-general.yml)                                         | crouch                       | —        | 0        | 本轮留空                               |
| [crystal-general](../entries/crystal-general.yml)                                       | crystal                      | —        | 0        | 本轮留空                               |
| [crystal-refinery-zelda](../entries/crystal-refinery-zelda.yml)                         | Crystal Refinery             | —        | 0        | 本轮留空                               |
| [crystallize-general](../entries/crystallize-general.yml)                               | crystallize                  | —        | 0        | WF04、R11（反向）                      |
| [crystallized-charge-zelda](../entries/crystallized-charge-zelda.yml)                   | Crystallized Charge          | —        | 0        | R06                                    |
| [crystallized-general](../entries/crystallized-general.yml)                             | crystallized                 | —        | 0        | WF05、R11                              |
| [cunning-general](../entries/cunning-general.yml)                                       | cunning                      | —        | 0        | 本轮留空                               |
| [curious-general](../entries/curious-general.yml)                                       | curious                      | —        | 0        | 本轮留空                               |
| [curry-food](../entries/curry-food.yml)                                                 | curry                        | —        | 0        | 本轮留空                               |
| [dagger-weapon](../entries/dagger-weapon.yml)                                           | dagger                       | —        | 0        | 本轮留空                               |
| [damage-combat](../entries/damage-combat.yml)                                           | damage                       | —        | 0        | 本轮留空                               |
| [death-mountain-zelda](../entries/death-mountain-zelda.yml)                             | Death Mountain               | —        | 0        | 本轮留空                               |
| [decayed-general](../entries/decayed-general.yml)                                       | decayed                      | —        | 1        | 保留已有辅助                           |
| [decompose-general](../entries/decompose-general.yml)                                   | decompose                    | —        | 0        | R07（反向）                            |
| [defense-combat](../entries/defense-combat.yml)                                         | defense                      | —        | 0        | 本轮留空                               |
| [defense-up-zelda](../entries/defense-up-zelda.yml)                                     | Defense Up                   | —        | 0        | 本轮留空                               |
| [deity-story](../entries/deity-story.yml)                                               | deity                        | —        | 1        | 保留已有辅助                           |
| [demon-dragon-zelda](../entries/demon-dragon-zelda.yml)                                 | Demon Dragon                 | —        | 0        | 本轮留空                               |
| [demon-king-ganondorf-zelda](../entries/demon-king-ganondorf-zelda.yml)                 | Demon King Ganondorf         | —        | 0        | 本轮留空                               |
| [demon-king-zelda](../entries/demon-king-zelda.yml)                                     | Demon King                   | —        | 0        | 本轮留空                               |
| [demon-kings-bow-zelda](../entries/demon-kings-bow-zelda.yml)                           | Demon King's Bow             | —        | 0        | 本轮留空                               |
| [demon-story](../entries/demon-story.yml)                                               | demon                        | —        | 0        | 本轮留空                               |
| [depths-terrain](../entries/depths-terrain.yml)                                         | depths                       | —        | 0        | 本轮留空                               |
| [descend-general](../entries/descend-general.yml)                                       | descend                      | —        | 1        | R10（反向）                            |
| [descendant-general](../entries/descendant-general.yml)                                 | descendant                   | WF       | 2        | 保留已有辅助                           |
| [descent-flight](../entries/descent-flight.yml)                                         | descent                      | —        | 0        | WF07、R10                              |
| [desert-terrain](../entries/desert-terrain.yml)                                         | desert                       | —        | 0        | 本轮留空                               |
| [destination-ui](../entries/destination-ui.yml)                                         | destination                  | —        | 0        | 本轮留空                               |
| [destroy-general](../entries/destroy-general.yml)                                       | destroy                      | —        | 0        | 本轮留空                               |
| [detach-general](../entries/detach-general.yml)                                         | detach                       | —        | 0        | 本轮留空                               |
| [determined-general](../entries/determined-general.yml)                                 | determined                   | —        | 0        | 本轮留空                               |
| [detour-general](../entries/detour-general.yml)                                         | detour                       | —        | 0        | 本轮留空                               |
| [device-dispenser-zelda](../entries/device-dispenser-zelda.yml)                         | Device Dispenser             | WF       | 0        | 保留已有辅助                           |
| [device-engineering](../entries/device-engineering.yml)                                 | device                       | —        | 0        | 本轮留空                               |
| [diligent-general](../entries/diligent-general.yml)                                     | diligent                     | —        | 0        | 本轮留空                               |
| [disassemble-general](../entries/disassemble-general.yml)                               | disassemble                  | —        | 0        | WF06、R07                              |
| [disguised-combat](../entries/disguised-combat.yml)                                     | disguised                    | —        | 0        | 本轮留空                               |
| [dispense-general](../entries/dispense-general.yml)                                     | dispense                     | —        | 0        | 本轮留空                               |
| [dispenser-engineering](../entries/dispenser-engineering.yml)                           | dispenser                    | WF       | 0        | 保留已有辅助                           |
| [distinctive-combat](../entries/distinctive-combat.yml)                                 | distinctive                  | —        | 0        | 本轮留空                               |
| [dive-general](../entries/dive-general.yml)                                             | dive                         | —        | 0        | 本轮留空                               |
| [dodge-general](../entries/dodge-general.yml)                                           | dodge                        | —        | 0        | 本轮留空                               |
| [door-architecture](../entries/door-architecture.yml)                                   | door                         | —        | 0        | 本轮留空                               |
| [draconification-zelda](../entries/draconification-zelda.yml)                           | draconification              | —        | 0        | 暂缓。命名史待证                       |
| [dragon-story](../entries/dragon-story.yml)                                             | dragon                       | —        | 0        | 本轮留空                               |
| [draw-general](../entries/draw-general.yml)                                             | draw                         | —        | 0        | 本轮留空                               |
| [drop-general](../entries/drop-general.yml)                                             | drop                         | —        | 0        | 本轮留空                               |
| [durability-equipment](../entries/durability-equipment.yml)                             | durability                   | WF       | 2        | 保留已有辅助                           |
| [durable-general](../entries/durable-general.yml)                                       | durable                      | —        | 1        | 保留已有辅助                           |
| [eagle-general](../entries/eagle-general.yml)                                           | eagle                        | —        | 0        | 本轮留空                               |
| [egg-food](../entries/egg-food.yml)                                                     | egg                          | —        | 0        | 本轮留空                               |
| [electro-cooking](../entries/electro-cooking.yml)                                       | Electro                      | —        | 0        | 本轮留空                               |
| [elixir-cooking](../entries/elixir-cooking.yml)                                         | elixir                       | ETY      | 0        | 保留已有辅助                           |
| [elusive-combat](../entries/elusive-combat.yml)                                         | elusive                      | —        | 0        | 本轮留空                               |
| [emit-physical](../entries/emit-physical.yml)                                           | emit                         | —        | 0        | 本轮留空                               |
| [emitter-engineering](../entries/emitter-engineering.yml)                               | emitter                      | WF       | 1        | 保留已有辅助                           |
| [endura-carrot-zelda](../entries/endura-carrot-zelda.yml)                               | Endura Carrot                | —        | 0        | 本轮留空                               |
| [endura-shroom-zelda](../entries/endura-shroom-zelda.yml)                               | Endura Shroom                | —        | 0        | 本轮留空                               |
| [endure-general](../entries/endure-general.yml)                                         | endure                       | —        | 1        | 保留已有辅助                           |
| [energize-general](../entries/energize-general.yml)                                     | energize                     | —        | 1        | 保留已有辅助                           |
| [energy-cell-zelda](../entries/energy-cell-zelda.yml)                                   | Energy Cell                  | —        | 0        | R05                                    |
| [energy-well-zelda](../entries/energy-well-zelda.yml)                                   | Energy Well                  | —        | 0        | N03、R05（反向）、R06（反向）          |
| [enhancement-general](../entries/enhancement-general.yml)                               | enhancement                  | WF       | 0        | 保留已有辅助                           |
| [enormous-combat](../entries/enormous-combat.yml)                                       | enormous                     | —        | 0        | 本轮留空                               |
| [enthusiastic-general](../entries/enthusiastic-general.yml)                             | enthusiastic                 | ETY      | 0        | 保留已有辅助                           |
| [entrance-architecture](../entries/entrance-architecture.yml)                           | entrance                     | —        | 0        | 本轮留空                               |
| [equip-general](../entries/equip-general.yml)                                           | equip                        | —        | 0        | 本轮留空                               |
| [equipment-equipment](../entries/equipment-equipment.yml)                               | equipment                    | —        | 0        | 本轮留空                               |
| [evermean-zelda](../entries/evermean-zelda.yml)                                         | Evermean                     | —        | 0        | 本轮留空                               |
| [examine-general](../entries/examine-general.yml)                                       | examine                      | —        | 0        | 本轮留空                               |
| [exit-architecture](../entries/exit-architecture.yml)                                   | exit                         | —        | 0        | 本轮留空                               |
| [factory-architecture](../entries/factory-architecture.yml)                             | factory                      | —        | 0        | 本轮留空                               |
| [fall-damage-game](../entries/fall-damage-game.yml)                                     | fall damage                  | —        | 0        | 本轮留空                               |
| [fall-general](../entries/fall-general.yml)                                             | fall                         | —        | 0        | 本轮留空                               |
| [fan-engineering](../entries/fan-engineering.yml)                                       | fan                          | —        | 0        | 本轮留空                               |
| [fast-travel-game](../entries/fast-travel-game.yml)                                     | fast travel                  | —        | 0        | 本轮留空                               |
| [fearless-general](../entries/fearless-general.yml)                                     | fearless                     | WF       | 2        | 保留已有辅助                           |
| [fierce-deity-armor-zelda](../entries/fierce-deity-armor-zelda.yml)                     | Fierce Deity Armor           | —        | 0        | 本轮留空                               |
| [fierce-general](../entries/fierce-general.yml)                                         | fierce                       | —        | 0        | 本轮留空                               |
| [fifth-general](../entries/fifth-general.yml)                                           | fifth                        | —        | 0        | 本轮留空                               |
| [find-general](../entries/find-general.yml)                                             | find                         | —        | 0        | 本轮留空                               |
| [fire-general](../entries/fire-general.yml)                                             | fire                         | —        | 0        | 本轮留空                               |
| [fire-temple-zelda](../entries/fire-temple-zelda.yml)                                   | Fire Temple                  | —        | 0        | 本轮留空                               |
| [flame-emitter-zelda](../entries/flame-emitter-zelda.yml)                               | Flame Emitter                | WF       | 0        | 保留已有辅助                           |
| [flame-general](../entries/flame-general.yml)                                           | flame                        | —        | 0        | 本轮留空                               |
| [flamebreaker-armor-zelda](../entries/flamebreaker-armor-zelda.yml)                     | Flamebreaker Armor           | —        | 0        | 本轮留空                               |
| [flurry-rush-zelda](../entries/flurry-rush-zelda.yml)                                   | Flurry Rush                  | —        | 0        | 暂缓。现有译名保留，数量与触发机制待证 |
| [flux-construct-zelda](../entries/flux-construct-zelda.yml)                             | Flux Construct               | —        | 0        | 本轮留空                               |
| [foothill-terrain](../entries/foothill-terrain.yml)                                     | foothill                     | WF       | 0        | 保留已有辅助                           |
| [forest-terrain](../entries/forest-terrain.yml)                                         | forest                       | —        | 0        | 本轮留空                               |
| [forge-architecture](../entries/forge-architecture.yml)                                 | forge                        | —        | 0        | 本轮留空                               |
| [forget-general](../entries/forget-general.yml)                                         | forget                       | —        | 0        | 本轮留空                               |
| [forgotten-temple-zelda](../entries/forgotten-temple-zelda.yml)                         | Forgotten Temple             | —        | 0        | 本轮留空                               |
| [fortress-architecture](../entries/fortress-architecture.yml)                           | fortress                     | —        | 0        | 本轮留空                               |
| [fountain-architecture](../entries/fountain-architecture.yml)                           | fountain                     | —        | 0        | 本轮留空                               |
| [free-fall-flight](../entries/free-fall-flight.yml)                                     | free fall                    | —        | 0        | 本轮留空                               |
| [frog-general](../entries/frog-general.yml)                                             | frog                         | —        | 0        | 本轮留空                               |
| [froggy-sleeve-zelda](../entries/froggy-sleeve-zelda.yml)                               | Froggy Sleeve                | —        | 0        | 本轮留空                               |
| [frost-emitter-zelda](../entries/frost-emitter-zelda.yml)                               | Frost Emitter                | WF       | 0        | 保留已有辅助                           |
| [frost-general](../entries/frost-general.yml)                                           | frost                        | —        | 0        | 本轮留空                               |
| [frox-zelda](../entries/frox-zelda.yml)                                                 | Frox                         | —        | 0        | 暂缓。命名史待证                       |
| [fruit-food](../entries/fruit-food.yml)                                                 | fruit                        | —        | 0        | 本轮留空                               |
| [fry-cooking](../entries/fry-cooking.yml)                                               | fry                          | —        | 0        | 本轮留空                               |
| [fuse-general](../entries/fuse-general.yml)                                             | fuse                         | —        | 1        | 保留已有辅助                           |
| [ganondorf-zelda](../entries/ganondorf-zelda.yml)                                       | Ganondorf                    | N        | 0        | 保留已有辅助                           |
| [gate-architecture](../entries/gate-architecture.yml)                                   | gate                         | —        | 0        | 本轮留空                               |
| [gear-equipment](../entries/gear-equipment.yml)                                         | gear                         | —        | 0        | 本轮留空                               |
| [general-store-architecture](../entries/general-store-architecture.yml)                 | general store                | —        | 0        | 本轮留空                               |
| [geoglyph-game](../entries/geoglyph-game.yml)                                           | geoglyph                     | —        | 0        | 本轮留空                               |
| [gerudo-desert-zelda](../entries/gerudo-desert-zelda.yml)                               | Gerudo Desert                | —        | 0        | 本轮留空                               |
| [gerudo-town-zelda](../entries/gerudo-town-zelda.yml)                                   | Gerudo Town                  | —        | 0        | 本轮留空                               |
| [gerudo-zelda](../entries/gerudo-zelda.yml)                                             | Gerudo                       | N        | 0        | 保留已有辅助                           |
| [gibdo-zelda](../entries/gibdo-zelda.yml)                                               | Gibdo                        | —        | 0        | 本轮留空                               |
| [gleeok-zelda](../entries/gleeok-zelda.yml)                                             | Gleeok                       | —        | 0        | 暂缓。命名史待证                       |
| [glide-general](../entries/glide-general.yml)                                           | glide                        | —        | 0        | 本轮留空                               |
| [glide-shirt-zelda](../entries/glide-shirt-zelda.yml)                                   | Glide Shirt                  | —        | 0        | 本轮留空                               |
| [gloom-recovery-zelda](../entries/gloom-recovery-zelda.yml)                             | Gloom Recovery               | —        | 0        | N02、R03                               |
| [gloom-resistance-zelda](../entries/gloom-resistance-zelda.yml)                         | Gloom Resistance             | —        | 0        | R03（反向）                            |
| [gloom-spawn-zelda](../entries/gloom-spawn-zelda.yml)                                   | Gloom Spawn                  | —        | 0        | 本轮留空                               |
| [gloom-sword-zelda](../entries/gloom-sword-zelda.yml)                                   | Gloom Sword                  | —        | 0        | 本轮留空                               |
| [gloom-zelda](../entries/gloom-zelda.yml)                                               | Gloom                        | —        | 0        | 本轮留空                               |
| [golden-apple-zelda](../entries/golden-apple-zelda.yml)                                 | Golden Apple                 | —        | 0        | 本轮留空                               |
| [golden-general](../entries/golden-general.yml)                                         | golden                       | —        | 0        | 本轮留空                               |
| [goron-city-zelda](../entries/goron-city-zelda.yml)                                     | Goron City                   | —        | 0        | 本轮留空                               |
| [goron-spice-zelda](../entries/goron-spice-zelda.yml)                                   | Goron Spice                  | —        | 0        | 本轮留空                               |
| [goron-zelda](../entries/goron-zelda.yml)                                               | Goron                        | —        | 0        | 本轮留空                               |
| [gourmet-food](../entries/gourmet-food.yml)                                             | gourmet                      | —        | 0        | 本轮留空                               |
| [great-abandoned-central-mine-zelda](../entries/great-abandoned-central-mine-zelda.yml) | Great Abandoned Central Mine | —        | 0        | 本轮留空                               |
| [great-deku-tree-zelda](../entries/great-deku-tree-zelda.yml)                           | Great Deku Tree              | —        | 0        | 本轮留空                               |
| [great-eagle-bow-zelda](../entries/great-eagle-bow-zelda.yml)                           | Great Eagle Bow              | —        | 0        | 本轮留空                               |
| [great-sky-island-zelda](../entries/great-sky-island-zelda.yml)                         | Great Sky Island             | —        | 0        | 本轮留空                               |
| [greaves-clothing](../entries/greaves-clothing.yml)                                     | greaves                      | —        | 0        | 本轮留空                               |
| [greens-food](../entries/greens-food.yml)                                               | greens                       | —        | 0        | 本轮留空                               |
| [grill-cooking](../entries/grill-cooking.yml)                                           | grill                        | N        | 1        | 保留已有辅助                           |
| [guard-combat](../entries/guard-combat.yml)                                             | guard                        | —        | 0        | 本轮留空                               |
| [guardian-general](../entries/guardian-general.yml)                                     | guardian                     | —        | 0        | 本轮留空                               |
| [guidance-story](../entries/guidance-story.yml)                                         | guidance                     | —        | 0        | 本轮留空                               |
| [gust-weather](../entries/gust-weather.yml)                                             | gust                         | —        | 0        | 本轮留空                               |
| [halberd-weapon](../entries/halberd-weapon.yml)                                         | halberd                      | —        | 0        | 本轮留空                               |
| [hammer-weapon](../entries/hammer-weapon.yml)                                           | hammer                       | —        | 0        | 本轮留空                               |
| [hand-general](../entries/hand-general.yml)                                             | hand                         | —        | 0        | 本轮留空                               |
| [hasty-general](../entries/hasty-general.yml)                                           | hasty                        | —        | 0        | 本轮留空                               |
| [hateno-village-zelda](../entries/hateno-village-zelda.yml)                             | Hateno Village               | —        | 0        | 本轮留空                               |
| [head-general](../entries/head-general.yml)                                             | head                         | —        | 0        | 本轮留空                               |
| [headband-clothing](../entries/headband-clothing.yml)                                   | headband                     | WF       | 1        | 保留已有辅助                           |
| [headdress-clothing](../entries/headdress-clothing.yml)                                 | headdress                    | WF       | 1        | 保留已有辅助                           |
| [heart-container-zelda](../entries/heart-container-zelda.yml)                           | Heart Container              | —        | 0        | 本轮留空                               |
| [hearty-general](../entries/hearty-general.yml)                                         | hearty                       | —        | 0        | 本轮留空                               |
| [hearty-radish-zelda](../entries/hearty-radish-zelda.yml)                               | Hearty Radish                | —        | 0        | 本轮留空                               |
| [hearty-salmon-zelda](../entries/hearty-salmon-zelda.yml)                               | Hearty Salmon                | —        | 0        | 本轮留空                               |
| [hearty-truffle-zelda](../entries/hearty-truffle-zelda.yml)                             | Hearty Truffle               | —        | 0        | 本轮留空                               |
| [heat-resistance-combat](../entries/heat-resistance-combat.yml)                         | heat resistance              | —        | 1        | 保留已有辅助                           |
| [heir-general](../entries/heir-general.yml)                                             | heir                         | N        | 1        | 保留已有辅助                           |
| [helm-clothing](../entries/helm-clothing.yml)                                           | helm                         | —        | 1        | 保留已有辅助                           |
| [helmet-clothing](../entries/helmet-clothing.yml)                                       | helmet                       | —        | 1        | 保留已有辅助                           |
| [hero-general](../entries/hero-general.yml)                                             | hero                         | —        | 0        | 本轮留空                               |
| [hestu-zelda](../entries/hestu-zelda.yml)                                               | Hestu                        | N        | 0        | 保留已有辅助                           |
| [hidden-general](../entries/hidden-general.yml)                                         | hidden                       | —        | 0        | 本轮留空                               |
| [hideout-architecture](../entries/hideout-architecture.yml)                             | hideout                      | —        | 0        | 本轮留空                               |
| [highlands-terrain](../entries/highlands-terrain.yml)                                   | highlands                    | WF       | 0        | 保留已有辅助                           |
| [hinox-zelda](../entries/hinox-zelda.yml)                                               | Hinox                        | N        | 0        | 保留已有辅助                           |
| [hold-general](../entries/hold-general.yml)                                             | hold                         | —        | 0        | 本轮留空                               |
| [homing-cart-zelda](../entries/homing-cart-zelda.yml)                                   | Homing Cart                  | —        | 0        | 本轮留空                               |
| [homing-engineering](../entries/homing-engineering.yml)                                 | homing                       | WF       | 0        | 保留已有辅助                           |
| [honey-food](../entries/honey-food.yml)                                                 | honey                        | —        | 0        | 本轮留空                               |
| [honorable-general](../entries/honorable-general.yml)                                   | honorable                    | —        | 0        | 本轮留空                               |
| [hood-clothing](../entries/hood-clothing.yml)                                           | hood                         | —        | 0        | 本轮留空                               |
| [hook-hardware](../entries/hook-hardware.yml)                                           | hook                         | —        | 0        | R09                                    |
| [horriblin-zelda](../entries/horriblin-zelda.yml)                                       | Horriblin                    | —        | 0        | 暂缓。命名史待证                       |
| [hostile-combat](../entries/hostile-combat.yml)                                         | hostile                      | ETY、N   | 1        | 保留已有辅助                           |
| [hover-general](../entries/hover-general.yml)                                           | hover                        | —        | 0        | 本轮留空                               |
| [hover-stone-zelda](../entries/hover-stone-zelda.yml)                                   | Hover Stone                  | WF       | 0        | 保留已有辅助                           |
| [hudson-zelda](../entries/hudson-zelda.yml)                                             | Hudson                       | N        | 0        | 保留已有辅助                           |
| [hydrant-engineering](../entries/hydrant-engineering.yml)                               | hydrant                      | ETY      | 0        | 保留已有辅助                           |
| [hylian-hood-zelda](../entries/hylian-hood-zelda.yml)                                   | Hylian Hood                  | —        | 0        | 本轮留空                               |
| [hylian-shield-zelda](../entries/hylian-shield-zelda.yml)                               | Hylian Shield                | —        | 0        | 本轮留空                               |
| [hylian-tunic-zelda](../entries/hylian-tunic-zelda.yml)                                 | Hylian Tunic                 | —        | 0        | 本轮留空                               |
| [hylian-zelda](../entries/hylian-zelda.yml)                                             | Hylian                       | —        | 0        | 本轮留空                               |
| [hyrule-bass-zelda](../entries/hyrule-bass-zelda.yml)                                   | Hyrule Bass                  | —        | 0        | 本轮留空                               |
| [hyrule-castle-zelda](../entries/hyrule-castle-zelda.yml)                               | Hyrule Castle                | —        | 0        | 本轮留空                               |
| [hyrule-compendium-zelda](../entries/hyrule-compendium-zelda.yml)                       | Hyrule Compendium            | —        | 0        | 本轮留空                               |
| [hyrule-zelda](../entries/hyrule-zelda.yml)                                             | Hyrule                       | —        | 0        | 本轮留空                               |
| [illuminate-general](../entries/illuminate-general.yml)                                 | illuminate                   | —        | 0        | 本轮留空                               |
| [impa-zelda](../entries/impa-zelda.yml)                                                 | Impa                         | —        | 0        | 本轮留空                               |
| [imprison-story](../entries/imprison-story.yml)                                         | imprison                     | —        | 0        | 本轮留空                               |
| [ingredient-cooking](../entries/ingredient-cooking.yml)                                 | ingredient                   | —        | 1        | 保留已有辅助                           |
| [inherit-story](../entries/inherit-story.yml)                                           | inherit                      | —        | 0        | 本轮留空                               |
| [inn-architecture](../entries/inn-architecture.yml)                                     | inn                          | —        | 0        | 本轮留空                               |
| [intact-tech](../entries/intact-tech.yml)                                               | intact                       | —        | 0        | 本轮留空                               |
| [integrity-general](../entries/integrity-general.yml)                                   | integrity                    | ETY      | 0        | 保留已有辅助                           |
| [interact-general](../entries/interact-general.yml)                                     | interact                     | —        | 0        | 本轮留空                               |
| [inventory-ui](../entries/inventory-ui.yml)                                             | inventory                    | N        | 1        | 保留已有辅助                           |
| [island-terrain](../entries/island-terrain.yml)                                         | island                       | —        | 0        | 本轮留空                               |
| [jar-general](../entries/jar-general.yml)                                               | jar                          | —        | 0        | 本轮留空                               |
| [josha-zelda](../entries/josha-zelda.yml)                                               | Josha                        | N        | 0        | 保留已有辅助                           |
| [kakariko-village-zelda](../entries/kakariko-village-zelda.yml)                         | Kakariko Village             | —        | 0        | 本轮留空                               |
| [keese-zelda](../entries/keese-zelda.yml)                                               | Keese                        | —        | 0        | 本轮留空                               |
| [key-items-zelda](../entries/key-items-zelda.yml)                                       | Key Items                    | —        | 0        | 本轮留空                               |
| [king-general](../entries/king-general.yml)                                             | king                         | —        | 0        | 本轮留空                               |
| [king-gleeok-zelda](../entries/king-gleeok-zelda.yml)                                   | King Gleeok                  | —        | 0        | 本轮留空                               |
| [kingdom-general](../entries/kingdom-general.yml)                                       | kingdom                      | WF       | 0        | 保留已有辅助                           |
| [knight-general](../entries/knight-general.yml)                                         | knight                       | —        | 0        | 本轮留空                               |
| [knights-broadsword-zelda](../entries/knights-broadsword-zelda.yml)                     | Knight's Broadsword          | —        | 0        | 本轮留空                               |
| [korok-forest-zelda](../entries/korok-forest-zelda.yml)                                 | Korok Forest                 | —        | 0        | 本轮留空                               |
| [korok-seed-zelda](../entries/korok-seed-zelda.yml)                                     | Korok Seed                   | —        | 0        | 本轮留空                               |
| [korok-zelda](../entries/korok-zelda.yml)                                               | Korok                        | —        | 0        | 本轮留空                               |
| [lake-terrain](../entries/lake-terrain.yml)                                             | lake                         | —        | 0        | 本轮留空                               |
| [land-motion](../entries/land-motion.yml)                                               | land                         | —        | 0        | 本轮留空                               |
| [landing-general](../entries/landing-general.yml)                                       | landing                      | —        | 0        | 本轮留空                               |
| [launch-motion](../entries/launch-motion.yml)                                           | launch                       | —        | 0        | 本轮留空                               |
| [lava-terrain](../entries/lava-terrain.yml)                                             | lava                         | —        | 0        | 本轮留空                               |
| [leggings-clothing](../entries/leggings-clothing.yml)                                   | leggings                     | —        | 0        | 本轮留空                               |
| [light-dragon-zelda](../entries/light-dragon-zelda.yml)                                 | Light Dragon                 | —        | 0        | 本轮留空                               |
| [light-engineering](../entries/light-engineering.yml)                                   | light                        | —        | 0        | 本轮留空                               |
| [light-of-blessing-zelda](../entries/light-of-blessing-zelda.yml)                       | Light of Blessing            | —        | 0        | 本轮留空                               |
| [lightning-temple-zelda](../entries/lightning-temple-zelda.yml)                         | Lightning Temple             | —        | 0        | 本轮留空                               |
| [lightning-terrain](../entries/lightning-terrain.yml)                                   | lightning                    | —        | 0        | 本轮留空                               |
| [lightroot-zelda](../entries/lightroot-zelda.yml)                                       | Lightroot                    | WF       | 0        | 保留已有辅助                           |
| [like-like-zelda](../entries/like-like-zelda.yml)                                       | Like Like                    | —        | 0        | 本轮留空                               |
| [link-zelda](../entries/link-zelda.yml)                                                 | Link                         | N        | 0        | 保留已有辅助                           |
| [lizalfos-zelda](../entries/lizalfos-zelda.yml)                                         | Lizalfos                     | —        | 0        | 本轮留空                               |
| [lookout-general](../entries/lookout-general.yml)                                       | lookout                      | —        | 0        | 本轮留空                               |
| [lookout-landing-zelda](../entries/lookout-landing-zelda.yml)                           | Lookout Landing              | —        | 0        | 本轮留空                               |
| [loyal-general](../entries/loyal-general.yml)                                           | loyal                        | —        | 0        | 本轮留空                               |
| [lurelin-village-zelda](../entries/lurelin-village-zelda.yml)                           | Lurelin Village              | —        | 0        | 本轮留空                               |
| [lynel-zelda](../entries/lynel-zelda.yml)                                               | Lynel                        | —        | 0        | 暂缓。命名史待证                       |
| [magic-rod-zelda](../entries/magic-rod-zelda.yml)                                       | Magic Rod                    | —        | 0        | 本轮留空                               |
| [main-quest-game](../entries/main-quest-game.yml)                                       | main quest                   | —        | 0        | 本轮留空                               |
| [main-quests-zelda](../entries/main-quests-zelda.yml)                                   | Main Quests                  | —        | 0        | 本轮留空                               |
| [majestic-general](../entries/majestic-general.yml)                                     | majestic                     | —        | 0        | 本轮留空                               |
| [malevolent-general](../entries/malevolent-general.yml)                                 | malevolent                   | ETY      | 0        | 保留已有辅助                           |
| [mandatory-general](../entries/mandatory-general.yml)                                   | mandatory                    | —        | 0        | 本轮留空                               |
| [manual-save-game](../entries/manual-save-game.yml)                                     | manual save                  | —        | 0        | R01、R02（反向）                       |
| [map-ui](../entries/map-ui.yml)                                                         | map                          | —        | 0        | 本轮留空                               |
| [marbled-gohma-zelda](../entries/marbled-gohma-zelda.yml)                               | Marbled Gohma                | —        | 0        | 本轮留空                               |
| [mask-clothing](../entries/mask-clothing.yml)                                           | mask                         | —        | 0        | 本轮留空                               |
| [master-kohga-zelda](../entries/master-kohga-zelda.yml)                                 | Master Kohga                 | N        | 0        | 保留已有辅助                           |
| [master-sword-zelda](../entries/master-sword-zelda.yml)                                 | Master Sword                 | —        | 0        | 本轮留空                               |
| [material-general](../entries/material-general.yml)                                     | material                     | —        | 0        | 本轮留空                               |
| [meal-cooking](../entries/meal-cooking.yml)                                             | meal                         | —        | 0        | 本轮留空                               |
| [meat-food](../entries/meat-food.yml)                                                   | meat                         | —        | 0        | 本轮留空                               |
| [melee-combat](../entries/melee-combat.yml)                                             | melee                        | ETY      | 0        | 保留已有辅助                           |
| [memory-story](../entries/memory-story.yml)                                             | memory                       | —        | 0        | 本轮留空                               |
| [menu-ui](../entries/menu-ui.yml)                                                       | menu                         | —        | 0        | 本轮留空                               |
| [mighty-bananas-zelda](../entries/mighty-bananas-zelda.yml)                             | Mighty Bananas               | —        | 0        | 本轮留空                               |
| [mighty-general](../entries/mighty-general.yml)                                         | mighty                       | —        | 0        | 本轮留空                               |
| [milk-food](../entries/milk-food.yml)                                                   | milk                         | —        | 0        | 本轮留空                               |
| [mine-architecture](../entries/mine-architecture.yml)                                   | mine                         | —        | 0        | 本轮留空                               |
| [mine-cart-engineering](../entries/mine-cart-engineering.yml)                           | mine cart                    | —        | 0        | 本轮留空                               |
| [miner-general](../entries/miner-general.yml)                                           | miner                        | WF       | 0        | 保留已有辅助                           |
| [miners-mask-zelda](../entries/miners-mask-zelda.yml)                                   | Miner's Mask                 | —        | 0        | 本轮留空                               |
| [mineru-zelda](../entries/mineru-zelda.yml)                                             | Mineru                       | N        | 0        | 保留已有辅助                           |
| [mirror-general](../entries/mirror-general.yml)                                         | mirror                       | —        | 0        | 本轮留空                               |
| [moblin-zelda](../entries/moblin-zelda.yml)                                             | Moblin                       | —        | 0        | 本轮留空                               |
| [mountain-terrain](../entries/mountain-terrain.yml)                                     | mountain                     | —        | 0        | 本轮留空                               |
| [mucktorok-zelda](../entries/mucktorok-zelda.yml)                                       | Mucktorok                    | —        | 0        | 暂缓。命名史待证                       |
| [mushroom-food](../entries/mushroom-food.yml)                                           | mushroom                     | —        | 0        | 本轮留空                               |
| [mysterious-general](../entries/mysterious-general.yml)                                 | mysterious                   | WF       | 0        | 保留已有辅助                           |
| [mystery-story](../entries/mystery-story.yml)                                           | mystery                      | —        | 0        | 本轮留空                               |
| [noble-general](../entries/noble-general.yml)                                           | noble                        | —        | 0        | 本轮留空                               |
| [objective-ui](../entries/objective-ui.yml)                                             | objective                    | —        | 0        | 本轮留空                               |
| [observation-deck-general](../entries/observation-deck-general.yml)                     | observation deck             | —        | 0        | 本轮留空                               |
| [octorok-zelda](../entries/octorok-zelda.yml)                                           | Octorok                      | —        | 0        | 暂缓。命名史待证                       |
| [oil-food](../entries/oil-food.yml)                                                     | oil                          | —        | 0        | 本轮留空                               |
| [omelet-cooking](../entries/omelet-cooking.yml)                                         | omelet                       | —        | 0        | 本轮留空                               |
| [one-handed-combat](../entries/one-handed-combat.yml)                                   | one-handed                   | WF       | 0        | 保留已有辅助                           |
| [opposite-shore-terrain](../entries/opposite-shore-terrain.yml)                         | opposite shore               | —        | 0        | 本轮留空                               |
| [optional-general](../entries/optional-general.yml)                                     | optional                     | —        | 0        | 本轮留空                               |
| [pad-general](../entries/pad-general.yml)                                               | pad                          | —        | 0        | 本轮留空                               |
| [paella-cooking](../entries/paella-cooking.yml)                                         | paella                       | ETY      | 0        | 保留已有辅助                           |
| [paraglider-zelda](../entries/paraglider-zelda.yml)                                     | Paraglider                   | —        | 0        | 本轮留空                               |
| [passage-architecture](../entries/passage-architecture.yml)                             | passage                      | —        | 0        | 本轮留空                               |
| [paya-zelda](../entries/paya-zelda.yml)                                                 | Paya                         | N        | 0        | 保留已有辅助                           |
| [pepper-vegetable](../entries/pepper-vegetable.yml)                                     | pepper                       | —        | 0        | R08                                    |
| [perfect-dodge-zelda](../entries/perfect-dodge-zelda.yml)                               | Perfect Dodge                | —        | 0        | R04（反向）                            |
| [perfect-guard-zelda](../entries/perfect-guard-zelda.yml)                               | Perfect Guard                | —        | 0        | R04                                    |
| [phantom-ganon-zelda](../entries/phantom-ganon-zelda.yml)                               | Phantom Ganon                | —        | 0        | 本轮留空                               |
| [phenomenon-story](../entries/phenomenon-story.yml)                                     | phenomenon                   | —        | 0        | WF08                                   |
| [pick-up-general](../entries/pick-up-general.yml)                                       | pick up                      | —        | 0        | 本轮留空                               |
| [pit-terrain](../entries/pit-terrain.yml)                                               | pit                          | —        | 0        | 本轮留空                               |
| [plateau-terrain](../entries/plateau-terrain.yml)                                       | plateau                      | —        | 0        | 本轮留空                               |
| [platform-architecture](../entries/platform-architecture.yml)                           | platform                     | —        | 0        | 本轮留空                               |
| [porridge-cooking](../entries/porridge-cooking.yml)                                     | porridge                     | —        | 0        | 本轮留空                               |
| [portable-pot-zelda](../entries/portable-pot-zelda.yml)                                 | Portable Pot                 | —        | 0        | 本轮留空                               |
| [possessed-combat](../entries/possessed-combat.yml)                                     | possessed                    | —        | 0        | 本轮留空                               |
| [pot-general](../entries/pot-general.yml)                                               | pot                          | —        | 0        | 本轮留空                               |
| [pouch-ui](../entries/pouch-ui.yml)                                                     | pouch                        | —        | 1        | 保留已有辅助                           |
| [powerful-general](../entries/powerful-general.yml)                                     | powerful                     | WF       | 0        | 保留已有辅助                           |
| [prime-food](../entries/prime-food.yml)                                                 | prime                        | —        | 0        | 本轮留空                               |
| [prince-general](../entries/prince-general.yml)                                         | prince                       | —        | 0        | 本轮留空                               |
| [princess-general](../entries/princess-general.yml)                                     | princess                     | —        | 0        | 本轮留空                               |
| [princess-zelda-zelda](../entries/princess-zelda-zelda.yml)                             | Princess Zelda               | N        | 0        | 保留已有辅助                           |
| [pristine-general](../entries/pristine-general.yml)                                     | pristine                     | N        | 1        | 保留已有辅助                           |
| [profile-ui](../entries/profile-ui.yml)                                                 | profile                      | —        | 0        | 本轮留空                               |
| [progression-game](../entries/progression-game.yml)                                     | progression                  | —        | 0        | 本轮留空                               |
| [prophecy-story](../entries/prophecy-story.yml)                                         | prophecy                     | —        | 0        | 本轮留空                               |
| [propulsion-engineering](../entries/propulsion-engineering.yml)                         | propulsion                   | —        | 0        | 本轮留空                               |
| [puffshroom-zelda](../entries/puffshroom-zelda.yml)                                     | Puffshroom                   | —        | 0        | 本轮留空                               |
| [purah-pad-zelda](../entries/purah-pad-zelda.yml)                                       | Purah Pad                    | —        | 0        | 本轮留空                               |
| [purah-zelda](../entries/purah-zelda.yml)                                               | Purah                        | N        | 0        | 保留已有辅助                           |
| [queen-general](../entries/queen-general.yml)                                           | queen                        | —        | 0        | 本轮留空                               |
| [queen-gibdo-zelda](../entries/queen-gibdo-zelda.yml)                                   | Queen Gibdo                  | —        | 0        | 本轮留空                               |
| [quest-game](../entries/quest-game.yml)                                                 | quest                        | —        | 0        | 本轮留空                               |
| [quicksave-game](../entries/quicksave-game.yml)                                         | quicksave                    | —        | 0        | N01、R02                               |
| [radish-food](../entries/radish-food.yml)                                               | radish                       | —        | 0        | 本轮留空                               |
| [raft-watercraft](../entries/raft-watercraft.yml)                                       | raft                         | —        | 0        | 本轮留空                               |
| [railway-track-engineering](../entries/railway-track-engineering.yml)                   | railway track                | —        | 0        | 本轮留空                               |
| [ranged-combat](../entries/ranged-combat.yml)                                           | ranged                       | —        | 0        | 本轮留空                               |
| [rauru-zelda](../entries/rauru-zelda.yml)                                               | Rauru                        | N        | 0        | 保留已有辅助                           |
| [raw-food](../entries/raw-food.yml)                                                     | raw                          | —        | 0        | 本轮留空                               |
| [raw-gourmet-meat-zelda](../entries/raw-gourmet-meat-zelda.yml)                         | Raw Gourmet Meat             | —        | 0        | 本轮留空                               |
| [raw-prime-meat-zelda](../entries/raw-prime-meat-zelda.yml)                             | Raw Prime Meat               | —        | 0        | 本轮留空                               |
| [reattach-general](../entries/reattach-general.yml)                                     | reattach                     | —        | 0        | WF02                                   |
| [recall-general](../entries/recall-general.yml)                                         | recall                       | —        | 0        | 本轮留空                               |
| [recharge-general](../entries/recharge-general.yml)                                     | recharge                     | WF       | 0        | 保留已有辅助                           |
| [recipe-cooking](../entries/recipe-cooking.yml)                                         | recipe                       | ETY      | 1        | 保留已有辅助                           |
| [reckless-general](../entries/reckless-general.yml)                                     | reckless                     | N        | 1        | 保留已有辅助                           |
| [recover-story](../entries/recover-story.yml)                                           | recover                      | —        | 1        | 保留已有辅助                           |
| [refinery-architecture](../entries/refinery-architecture.yml)                           | refinery                     | WF       | 0        | 保留已有辅助                           |
| [regional-general](../entries/regional-general.yml)                                     | regional                     | WF       | 0        | 保留已有辅助                           |
| [regional-phenomena-zelda](../entries/regional-phenomena-zelda.yml)                     | Regional Phenomena           | —        | 0        | 暂缓。现有译名保留，数量与触发机制待证 |
| [resilient-combat](../entries/resilient-combat.yml)                                     | resilient                    | ETY、N   | 1        | 保留已有辅助                           |
| [resistance-general](../entries/resistance-general.yml)                                 | resistance                   | —        | 0        | 本轮留空                               |
| [restore-tech](../entries/restore-tech.yml)                                             | restore                      | —        | 1        | 保留已有辅助                           |
| [reuse-general](../entries/reuse-general.yml)                                           | reuse                        | —        | 0        | 本轮留空                               |
| [rice-food](../entries/rice-food.yml)                                                   | rice                         | —        | 0        | 本轮留空                               |
| [ridge-terrain](../entries/ridge-terrain.yml)                                           | ridge                        | —        | 1        | 保留已有辅助                           |
| [riju-zelda](../entries/riju-zelda.yml)                                                 | Riju                         | N        | 0        | 保留已有辅助                           |
| [ring-general](../entries/ring-general.yml)                                             | ring                         | —        | 0        | 本轮留空                               |
| [risotto-cooking](../entries/risotto-cooking.yml)                                       | risotto                      | —        | 0        | 本轮留空                               |
| [rito-village-zelda](../entries/rito-village-zelda.yml)                                 | Rito Village                 | —        | 0        | 本轮留空                               |
| [rito-zelda](../entries/rito-zelda.yml)                                                 | Rito                         | —        | 0        | 本轮留空                               |
| [river-terrain](../entries/river-terrain.yml)                                           | river                        | —        | 0        | 本轮留空                               |
| [roast-food](../entries/roast-food.yml)                                                 | roast                        | —        | 0        | 本轮留空                               |
| [robbie-zelda](../entries/robbie-zelda.yml)                                             | Robbie                       | N        | 0        | 保留已有辅助                           |
| [robe-clothing](../entries/robe-clothing.yml)                                           | robe                         | —        | 0        | 本轮留空                               |
| [rocket-engineering](../entries/rocket-engineering.yml)                                 | rocket                       | —        | 0        | 本轮留空                               |
| [rod-general](../entries/rod-general.yml)                                               | rod                          | —        | 0        | 本轮留空                               |
| [root-general](../entries/root-general.yml)                                             | root                         | —        | 0        | 本轮留空                               |
| [rotate-general](../entries/rotate-general.yml)                                         | rotate                       | —        | 0        | 本轮留空                               |
| [royal-broadsword-zelda](../entries/royal-broadsword-zelda.yml)                         | Royal Broadsword             | —        | 0        | 本轮留空                               |
| [royal-general](../entries/royal-general.yml)                                           | royal                        | —        | 0        | 本轮留空                               |
| [ruins-architecture](../entries/ruins-architecture.yml)                                 | ruins                        | —        | 0        | 本轮留空                               |
| [runway-flight](../entries/runway-flight.yml)                                           | runway                       | —        | 0        | 本轮留空                               |
| [rupee-game](../entries/rupee-game.yml)                                                 | rupee                        | ETY      | 0        | 保留已有辅助                           |
| [rusty-general](../entries/rusty-general.yml)                                           | rusty                        | —        | 0        | 本轮留空                               |
| [ruthless-general](../entries/ruthless-general.yml)                                     | ruthless                     | WF       | 0        | 保留已有辅助                           |
| [sacred-general](../entries/sacred-general.yml)                                         | sacred                       | —        | 0        | 本轮留空                               |
| [sacrifice-story](../entries/sacrifice-story.yml)                                       | sacrifice                    | ETY      | 0        | 保留已有辅助                           |
| [sage-general](../entries/sage-general.yml)                                             | sage                         | —        | 0        | 本轮留空                               |
| [sail-watercraft](../entries/sail-watercraft.yml)                                       | sail                         | —        | 0        | 本轮留空                               |
| [salmon-food](../entries/salmon-food.yml)                                               | salmon                       | —        | 0        | 本轮留空                               |
| [salt-food](../entries/salt-food.yml)                                                   | salt                         | —        | 0        | 本轮留空                               |
| [saute-cooking](../entries/saute-cooking.yml)                                           | sauté                        | ETY      | 0        | 保留已有辅助                           |
| [savage-general](../entries/savage-general.yml)                                         | savage                       | —        | 0        | 本轮留空                               |
| [save-game](../entries/save-game.yml)                                                   | save                         | —        | 0        | 本轮留空                               |
| [scimitar-of-the-seven-zelda](../entries/scimitar-of-the-seven-zelda.yml)               | Scimitar of the Seven        | —        | 0        | 本轮留空                               |
| [scimitar-weapon](../entries/scimitar-weapon.yml)                                       | scimitar                     | —        | 0        | 本轮留空                               |
| [seafood-food](../entries/seafood-food.yml)                                             | seafood                      | —        | 0        | 本轮留空                               |
| [sear-cooking](../entries/sear-cooking.yml)                                             | sear                         | N        | 1        | 保留已有辅助                           |
| [secret-stone-zelda](../entries/secret-stone-zelda.yml)                                 | Secret Stone                 | —        | 0        | 本轮留空                               |
| [secret-story](../entries/secret-story.yml)                                             | secret                       | —        | 0        | 本轮留空                               |
| [seed-general](../entries/seed-general.yml)                                             | seed                         | —        | 0        | 本轮留空                               |
| [seized-construct-zelda](../entries/seized-construct-zelda.yml)                         | Seized Construct             | —        | 0        | 本轮留空                               |
| [select-general](../entries/select-general.yml)                                         | select                       | —        | 0        | 本轮留空                               |
| [selfless-general](../entries/selfless-general.yml)                                     | selfless                     | WF       | 0        | 保留已有辅助                           |
| [settlement-architecture](../entries/settlement-architecture.yml)                       | settlement                   | WF       | 0        | 保留已有辅助                           |
| [sheikah-zelda](../entries/sheikah-zelda.yml)                                           | Sheikah                      | —        | 0        | 本轮留空                               |
| [shield-weapon](../entries/shield-weapon.yml)                                           | shield                       | —        | 0        | 本轮留空                               |
| [shirt-clothing](../entries/shirt-clothing.yml)                                         | shirt                        | —        | 0        | 本轮留空                               |
| [shock-emitter-zelda](../entries/shock-emitter-zelda.yml)                               | Shock Emitter                | WF       | 0        | 保留已有辅助                           |
| [shock-resistance-combat](../entries/shock-resistance-combat.yml)                       | shock resistance             | —        | 0        | 本轮留空                               |
| [shrine-architecture](../entries/shrine-architecture.yml)                               | shrine                       | —        | 0        | 本轮留空                               |
| [shrine-quests-zelda](../entries/shrine-quests-zelda.yml)                               | Shrine Quests                | —        | 0        | 本轮留空                               |
| [side-adventures-zelda](../entries/side-adventures-zelda.yml)                           | Side Adventures              | —        | 0        | 本轮留空                               |
| [side-hop-combat](../entries/side-hop-combat.yml)                                       | side hop                     | —        | 0        | 本轮留空                               |
| [side-quests-zelda](../entries/side-quests-zelda.yml)                                   | Side Quests                  | —        | 0        | 本轮留空                               |
| [sidon-zelda](../entries/sidon-zelda.yml)                                               | Sidon                        | N        | 0        | 保留已有辅助                           |
| [silent-princess-zelda](../entries/silent-princess-zelda.yml)                           | Silent Princess              | —        | 0        | 本轮留空                               |
| [simmer-cooking](../entries/simmer-cooking.yml)                                         | simmer                       | N        | 1        | 保留已有辅助                           |
| [skewer-cooking](../entries/skewer-cooking.yml)                                         | skewer                       | —        | 0        | 本轮留空                               |
| [skilled-general](../entries/skilled-general.yml)                                       | skilled                      | WF       | 0        | 保留已有辅助                           |
| [sky-island-zelda](../entries/sky-island-zelda.yml)                                     | sky island                   | WF       | 0        | 保留已有辅助                           |
| [sky-terrain](../entries/sky-terrain.yml)                                               | sky                          | —        | 0        | 本轮留空                               |
| [skyshroom-zelda](../entries/skyshroom-zelda.yml)                                       | Skyshroom                    | —        | 0        | 本轮留空                               |
| [skyview-tower-zelda](../entries/skyview-tower-zelda.yml)                               | Skyview Tower                | WF       | 0        | 保留已有辅助                           |
| [sled-engineering](../entries/sled-engineering.yml)                                     | sled                         | —        | 0        | 本轮留空                               |
| [sleeve-clothing](../entries/sleeve-clothing.yml)                                       | sleeve                       | —        | 0        | 本轮留空                               |
| [slip-general](../entries/slip-general.yml)                                             | slip                         | —        | 0        | 本轮留空                               |
| [slip-resistance-zelda](../entries/slip-resistance-zelda.yml)                           | Slip Resistance              | —        | 0        | 本轮留空                               |
| [slope-terrain](../entries/slope-terrain.yml)                                           | slope                        | —        | 0        | 本轮留空                               |
| [small-wheel-zelda](../entries/small-wheel-zelda.yml)                                   | Small Wheel                  | —        | 0        | 本轮留空                               |
| [sneaky-general](../entries/sneaky-general.yml)                                         | sneaky                       | —        | 0        | 本轮留空                               |
| [snowquill-tunic-zelda](../entries/snowquill-tunic-zelda.yml)                           | Snowquill Tunic              | —        | 0        | 本轮留空                               |
| [soldier-construct-zelda](../entries/soldier-construct-zelda.yml)                       | Soldier Construct            | —        | 0        | 本轮留空                               |
| [soldier-general](../entries/soldier-general.yml)                                       | soldier                      | —        | 0        | 本轮留空                               |
| [soldiers-broadsword-zelda](../entries/soldiers-broadsword-zelda.yml)                   | Soldier's Broadsword         | —        | 0        | 本轮留空                               |
| [sonia-zelda](../entries/sonia-zelda.yml)                                               | Sonia                        | N        | 0        | 保留已有辅助                           |
| [sort-general](../entries/sort-general.yml)                                             | sort                         | —        | 0        | 本轮留空                               |
| [spaulder-clothing](../entries/spaulder-clothing.yml)                                   | spaulder                     | —        | 0        | 本轮留空                               |
| [spear-weapon](../entries/spear-weapon.yml)                                             | spear                        | —        | 0        | 本轮留空                               |
| [special-controls-zelda](../entries/special-controls-zelda.yml)                         | Special Controls             | —        | 0        | 本轮留空                               |
| [speed-up-zelda](../entries/speed-up-zelda.yml)                                         | Speed Up                     | —        | 0        | 本轮留空                               |
| [spice-food](../entries/spice-food.yml)                                                 | spice                        | —        | 0        | 本轮留空                               |
| [spicy-general](../entries/spicy-general.yml)                                           | spicy                        | —        | 0        | 本轮留空                               |
| [spicy-pepper-zelda](../entries/spicy-pepper-zelda.yml)                                 | Spicy Pepper                 | —        | 0        | 本轮留空                               |
| [spirit-story](../entries/spirit-story.yml)                                             | spirit                       | —        | 0        | 本轮留空                               |
| [spirit-temple-zelda](../entries/spirit-temple-zelda.yml)                               | Spirit Temple                | —        | 0        | 本轮留空                               |
| [spring-mechanical](../entries/spring-mechanical.yml)                                   | spring                       | —        | 1        | 保留已有辅助                           |
| [spring-water](../entries/spring-water.yml)                                             | spring                       | N        | 1        | 保留已有辅助                           |
| [sprint-general](../entries/sprint-general.yml)                                         | sprint                       | —        | 0        | 本轮留空                               |
| [stabilizer-engineering](../entries/stabilizer-engineering.yml)                         | stabilizer                   | WF       | 0        | 保留已有辅助                           |
| [stable-architecture](../entries/stable-architecture.yml)                               | stable                       | —        | 0        | 本轮留空                               |
| [staircase-architecture](../entries/staircase-architecture.yml)                         | staircase                    | WF       | 0        | 保留已有辅助                           |
| [stake-engineering](../entries/stake-engineering.yml)                                   | stake                        | —        | 0        | 本轮留空                               |
| [stamina-general](../entries/stamina-general.yml)                                       | stamina                      | ETY、N   | 1        | 保留已有辅助                           |
| [statue-architecture](../entries/statue-architecture.yml)                               | statue                       | —        | 0        | 本轮留空                               |
| [stealth-mask-zelda](../entries/stealth-mask-zelda.yml)                                 | Stealth Mask                 | —        | 0        | 本轮留空                               |
| [stealth-up-zelda](../entries/stealth-up-zelda.yml)                                     | Stealth Up                   | —        | 0        | 本轮留空                               |
| [steam-cooking](../entries/steam-cooking.yml)                                           | steam                        | N        | 0        | 保留已有辅助                           |
| [steer-general](../entries/steer-general.yml)                                           | steer                        | —        | 0        | 本轮留空                               |
| [steering-stick-zelda](../entries/steering-stick-zelda.yml)                             | Steering Stick               | WF       | 0        | 保留已有辅助                           |
| [stew-cooking](../entries/stew-cooking.yml)                                             | stew                         | —        | 1        | 保留已有辅助                           |
| [stick-general](../entries/stick-general.yml)                                           | stick                        | —        | 0        | 本轮留空                               |
| [stone-general](../entries/stone-general.yml)                                           | stone                        | —        | 0        | 本轮留空                               |
| [stone-talus-zelda](../entries/stone-talus-zelda.yml)                                   | Stone Talus                  | —        | 0        | 本轮留空                               |
| [sturdy-general](../entries/sturdy-general.yml)                                         | sturdy                       | N        | 1        | 保留已有辅助                           |
| [sugar-food](../entries/sugar-food.yml)                                                 | sugar                        | —        | 0        | 本轮留空                               |
| [summit-terrain](../entries/summit-terrain.yml)                                         | summit                       | N        | 1        | 保留已有辅助                           |
| [summon-general](../entries/summon-general.yml)                                         | summon                       | —        | 0        | 本轮留空                               |
| [sundelion-zelda](../entries/sundelion-zelda.yml)                                       | Sundelion                    | —        | 0        | 暂缓。命名史待证                       |
| [sunny-general](../entries/sunny-general.yml)                                           | sunny                        | —        | 0        | 本轮留空                               |
| [surface-terrain](../entries/surface-terrain.yml)                                       | surface                      | —        | 1        | 保留已有辅助                           |
| [susceptible-combat](../entries/susceptible-combat.yml)                                 | susceptible                  | N        | 1        | 保留已有辅助                           |
| [suspend-hang-general](../entries/suspend-hang-general.yml)                             | suspend                      | —        | 0        | 本轮留空                               |
| [swamp-terrain](../entries/swamp-terrain.yml)                                           | swamp                        | —        | 0        | 本轮留空                               |
| [sword-weapon](../entries/sword-weapon.yml)                                             | sword                        | —        | 0        | 本轮留空                               |
| [system-ui](../entries/system-ui.yml)                                                   | system                       | —        | 0        | 本轮留空                               |
| [tarrey-town-zelda](../entries/tarrey-town-zelda.yml)                                   | Tarrey Town                  | —        | 0        | 本轮留空                               |
| [tear-general](../entries/tear-general.yml)                                             | tear                         | N        | 1        | 保留已有辅助                           |
| [teba-zelda](../entries/teba-zelda.yml)                                                 | Teba                         | N        | 0        | 保留已有辅助                           |
| [temple-architecture](../entries/temple-architecture.yml)                               | temple                       | —        | 0        | 本轮留空                               |
| [temple-of-time-zelda](../entries/temple-of-time-zelda.yml)                             | Temple of Time               | —        | 0        | 本轮留空                               |
| [the-depths-zelda](../entries/the-depths-zelda.yml)                                     | The Depths                   | —        | 1        | 保留已有辅助                           |
| [throw-general](../entries/throw-general.yml)                                           | throw                        | —        | 0        | 本轮留空                               |
| [tights-clothing](../entries/tights-clothing.yml)                                       | tights                       | —        | 0        | 本轮留空                               |
| [time-bomb-zelda](../entries/time-bomb-zelda.yml)                                       | Time Bomb                    | —        | 0        | 本轮留空                               |
| [timid-general](../entries/timid-general.yml)                                           | timid                        | —        | 0        | 本轮留空                               |
| [tomato-food](../entries/tomato-food.yml)                                               | tomato                       | —        | 0        | 本轮留空                               |
| [torch-weapon](../entries/torch-weapon.yml)                                             | torch                        | —        | 0        | 本轮留空                               |
| [tough-general](../entries/tough-general.yml)                                           | tough                        | —        | 0        | 本轮留空                               |
| [tower-architecture](../entries/tower-architecture.yml)                                 | tower                        | —        | 0        | 本轮留空                               |
| [trail-story](../entries/trail-story.yml)                                               | trail                        | —        | 0        | 本轮留空                               |
| [traveler-general](../entries/traveler-general.yml)                                     | traveler                     | WF       | 0        | 保留已有辅助                           |
| [travelers-sword-zelda](../entries/travelers-sword-zelda.yml)                           | Traveler's Sword             | —        | 0        | 本轮留空                               |
| [tree-terrain](../entries/tree-terrain.yml)                                             | tree                         | —        | 0        | 本轮留空                               |
| [trident-weapon](../entries/trident-weapon.yml)                                         | trident                      | ETY      | 0        | 保留已有辅助                           |
| [trousers-clothing](../entries/trousers-clothing.yml)                                   | trousers                     | —        | 0        | 本轮留空                               |
| [truffle-food](../entries/truffle-food.yml)                                             | truffle                      | —        | 0        | 本轮留空                               |
| [tulin-zelda](../entries/tulin-zelda.yml)                                               | Tulin                        | N        | 0        | 保留已有辅助                           |
| [tunic-clothing](../entries/tunic-clothing.yml)                                         | tunic                        | —        | 0        | 本轮留空                               |
| [two-handed-combat](../entries/two-handed-combat.yml)                                   | two-handed                   | WF       | 0        | 保留已有辅助                           |
| [ultrahand-zelda](../entries/ultrahand-zelda.yml)                                       | Ultrahand                    | WF       | 0        | 保留已有辅助                           |
| [undead-combat](../entries/undead-combat.yml)                                           | undead                       | WF       | 0        | 保留已有辅助                           |
| [underground-general](../entries/underground-general.yml)                               | underground                  | WF       | 0        | 保留已有辅助                           |
| [unfuse-general](../entries/unfuse-general.yml)                                         | unfuse                       | WF       | 1        | 保留已有辅助                           |
| [unlock-general](../entries/unlock-general.yml)                                         | unlock                       | WF       | 0        | 保留已有辅助                           |
| [updraft-flight](../entries/updraft-flight.yml)                                         | updraft                      | —        | 0        | WF03                                   |
| [upgrade-general](../entries/upgrade-general.yml)                                       | upgrade                      | —        | 0        | 本轮留空                               |
| [upheaval-story](../entries/upheaval-story.yml)                                         | upheaval                     | —        | 0        | 本轮留空                               |
| [valley-terrain](../entries/valley-terrain.yml)                                         | valley                       | —        | 0        | 本轮留空                               |
| [view-general](../entries/view-general.yml)                                             | view                         | —        | 0        | 本轮留空                               |
| [village-architecture](../entries/village-architecture.yml)                             | village                      | —        | 0        | 本轮留空                               |
| [vitality-general](../entries/vitality-general.yml)                                     | vitality                     | —        | 0        | 本轮留空                               |
| [volcanic-combat](../entries/volcanic-combat.yml)                                       | volcanic                     | —        | 0        | 本轮留空                               |
| [vulnerable-combat](../entries/vulnerable-combat.yml)                                   | vulnerable                   | ETY、N   | 1        | 保留已有辅助                           |
| [waistguard-clothing](../entries/waistguard-clothing.yml)                               | waistguard                   | WF       | 0        | 保留已有辅助                           |
| [warp-general](../entries/warp-general.yml)                                             | warp                         | —        | 0        | 本轮留空                               |
| [warrior-general](../entries/warrior-general.yml)                                       | warrior                      | —        | 0        | 本轮留空                               |
| [water-general](../entries/water-general.yml)                                           | water                        | —        | 0        | 本轮留空                               |
| [water-temple-zelda](../entries/water-temple-zelda.yml)                                 | Water Temple                 | —        | 0        | 本轮留空                               |
| [waterfall-terrain](../entries/waterfall-terrain.yml)                                   | waterfall                    | WF       | 0        | 保留已有辅助                           |
| [weak-point-combat](../entries/weak-point-combat.yml)                                   | weak point                   | —        | 0        | 本轮留空                               |
| [weapon-weapon](../entries/weapon-weapon.yml)                                           | weapon                       | —        | 0        | 本轮留空                               |
| [well-architecture](../entries/well-architecture.yml)                                   | well                         | —        | 0        | 本轮留空                               |
| [wetlands-terrain](../entries/wetlands-terrain.yml)                                     | wetlands                     | WF       | 0        | 保留已有辅助                           |
| [wheat-food](../entries/wheat-food.yml)                                                 | wheat                        | —        | 0        | 本轮留空                               |
| [wheel-engineering](../entries/wheel-engineering.yml)                                   | wheel                        | —        | 0        | 本轮留空                               |
| [wield-story](../entries/wield-story.yml)                                               | wield                        | —        | 0        | 本轮留空                               |
| [wiggle-general](../entries/wiggle-general.yml)                                         | wiggle                       | —        | 0        | 本轮留空                               |
| [wind-temple-zelda](../entries/wind-temple-zelda.yml)                                   | Wind Temple                  | —        | 0        | 本轮留空                               |
| [wind-terrain](../entries/wind-terrain.yml)                                             | wind                         | —        | 0        | 本轮留空                               |
| [wing-engineering](../entries/wing-engineering.yml)                                     | wing                         | —        | 0        | 本轮留空                               |
| [wise-general](../entries/wise-general.yml)                                             | wise                         | —        | 0        | 本轮留空                               |
| [wizzrobe-zelda](../entries/wizzrobe-zelda.yml)                                         | Wizzrobe                     | —        | 0        | 暂缓。命名史待证                       |
| [yiga-clan-hideout-zelda](../entries/yiga-clan-hideout-zelda.yml)                       | Yiga Clan Hideout            | —        | 0        | 本轮留空                               |
| [yiga-clan-zelda](../entries/yiga-clan-zelda.yml)                                       | Yiga Clan                    | —        | 0        | 本轮留空                               |
| [yona-zelda](../entries/yona-zelda.yml)                                                 | Yona                         | —        | 0        | 本轮留空                               |
| [yunobo-zelda](../entries/yunobo-zelda.yml)                                             | Yunobo                       | N        | 0        | 保留已有辅助                           |
| [zonai-device-zelda](../entries/zonai-device-zelda.yml)                                 | Zonai Device                 | —        | 0        | 本轮留空                               |
| [zonai-zelda](../entries/zonai-zelda.yml)                                               | Zonai                        | N        | 0        | 保留已有辅助                           |
| [zonaite-zelda](../entries/zonaite-zelda.yml)                                           | Zonaite                      | —        | 0        | 暂缓。命名史待证                       |
| [zora-armor-zelda](../entries/zora-armor-zelda.yml)                                     | Zora Armor                   | —        | 0        | 本轮留空                               |
| [zora-zelda](../entries/zora-zelda.yml)                                                 | Zora                         | —        | 0        | 本轮留空                               |
| [zoras-domain-zelda](../entries/zoras-domain-zelda.yml)                                 | Zora's Domain                | —        | 0        | 本轮留空                               |

## 待批与实施验证

- 待批文本为 WF01–WF08、ETY01–ETY02、N01–N03。批准相应编号才允许写入该字段、出处或明确的 Notes 移出。
- 待批关联为 R01–R11。仅批准对应有向声明，显示仍为双向直接邻居。

1. 重读获批文本及当前实现。检查 Git 状态，协调规划后发生的变更，再快照词条原文、导出基线及已有媒体哈希。
2. 按批准编号应用文本、来源行、唯一移出和关联声明。保留稳定 ID、Word、Meaning、发音、例句、标签、学习数据、已有辅助、非授权字段和录音字节。本批范围仅含现有词条里的获批辅助，不涉及新增词条、录音、实现、依赖或模板。
3. 从仓库根目录执行以下检查，单独检查未跟踪文本。导出由现有构建器写入。

```bash
go test ./...
go vet ./...
go run scripts/validate/main.go
go run scripts/build/main.go
git diff --check
```

4. 源文件保持相同后再次构建，比较两次导出的字节。核对实施基线 ID 集合、行数、18 列顺序、原 Audio 和 StudyOrder、逐项批准文本及原有和新增直接双向关联。检验空字段和当前卡背条件。列顺序如下。

```plain
ID Language Word Pronunciation PartOfSpeech Meaning Context Example Notes Audio TTSLanguage StudyOrder Etymology WordFormation Mnemonic RelatedWords CodeExample Tags
```

5. 比较字段白名单和既有媒体哈希，检查批准范围外的改动。用实际结果更新本计划的批准、实施和验证记录，报告部分失败和保留待核实项。静态和导出检查与真人发音、实际 Anki 导入、渲染和播放结果分开记录。
6. 保持工作树改动可供用户审查。暂存、提交、推送、复制媒体和操作 Anki 不在授权范围内。

## 规划检查

已核对所有本地文件链接、稳定 ID、提案编号、字段、计数和 Markdown 表格。正文无片段链接，外部证据按来源登记的核实边界记录。原词条、媒体、实现、模板和导出均与规划快照匹配。本轮仅执行静态规划检查，未运行构建或生成音频，未验证实时 Anki。

## 实施记录

2026-11-22 已实施 WF01–WF08、ETY01–ETY02、N01–N03 和 R01–R11。

- 19 个词条文件实际修改 34 个字段，包括 13 个辅助字段、10 个 Notes 字段和 11 个关联列表。10 条出处已追加，WF08 的原句按批准范围移出，其他 Notes 内容保留。
- 11 对关联涉及 20 个端点。导出新增 22 条直接邻接显示。逐行比对确认原有 135 对关系均保留。跨主题的两个端点仅更新反向导出显示。
- Go 测试、vet、词库验证、两次导出构建和 `git diff --check` 均通过。两次构建生成完全相同的字节。
- 核对导出的 1833 行、18 列顺序、稳定 ID 集合、原 Audio 和 StudyOrder、批准文本、空辅助字段与完整的直接双向关联。卡背条件已静态核对。
- 源文件的实际字段变化符合批准范围。1875 个已有 MP3 的哈希与实施前一致。实际 Anki 导入、渲染、播放与人工听审未核验。
