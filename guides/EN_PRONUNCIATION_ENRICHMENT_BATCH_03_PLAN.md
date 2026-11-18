# 英语音标补全计划: 第三批日常名词与动作词

状态: **B03-P01–B03-P30 已确认并实施，静态及导出验证通过。** 日期: 2026-11-18。

规划阶段沿用前两批的逐项审阅流程，只创建第三批计划。用户随后确认 B03-P01–B03-P30 并明确授权实施，已按批准内容更新 30 个词条并重建导出。录音、归档、既有规范改动、实现、模板和依赖保持不变。未暂存、提交、推送或操作 Anki，实际结果见[实施记录](#实施记录)。

[第一批](EN_PRONUNCIATION_ENRICHMENT_BATCH_01_PLAN.md)和[第二批](EN_PRONUNCIATION_ENRICHMENT_BATCH_02_PLAN.md)作为已完成的历史记录保留。使用独立编号 B03-P01–B03-P30，不复用前两批批准编号。

## 导航

- [范围与基线](#范围与基线)
- [来源与口音](#来源与口音)
- [具体音标](#具体音标)
- [逐项核对](#逐项核对)
- [Notes 来源追加文案](#notes-来源追加文案)
- [保留与后续范围](#保留与后续范围)
- [批准后实施与验证](#批准后实施与验证)
- [规划检查](#规划检查)
- [实施记录](#实施记录)

## 范围与基线

以下保留 2026-11-18 实施前的规划基线。已按实际稳定 ID 盘点当时全库 pronunciation，并逐项匹配本批 Word、词性与 Meaning。完整原字段已在仓库外快照中记录，供保全核对；没有把字段盘点说成全库发音查证。

| 项目                          | 实际范围                                     |
|:------------------------------|:---------------------------------------------|
| 全库条目                      | 1,320                                        |
| 英语 / 日语                   | 1,077 / 243                                  |
| 英语已有 / 缺失 pronunciation | 337 / 740                                    |
| 日语已有 / 缺失 pronunciation | 243 / 0                                      |
| 本批查证并提案                | 30 个已有词条                                |
| 内容分组                      | 15 个日常名词，15 个常用动作词               |
| 口音                          | 30 项 US，0 项 UK                            |
| 本批新增暂缓项                | 0                                            |
| 前批暂缓项                    | 2，单独保留，不计入本批 30 项                |
| 批准后允许修改                | 仅 30 项 pronunciation 与各自 Notes 来源追加 |
| 新增词条 / 音频 / 关联 / 先修 | 全部为 0                                     |

本批按日常使用价值选择食物、饮水名词及常见动作词，不以补满比例为目标。所选项目均有可读取的美式词典证据，因此本批无需使用英式替代。它们与前两批已实施目标不重叠，词长、字母顺序或本批表格次序不改变 StudyOrder。

规划阶段英语缺失数量为 740。规划时预计实施这 30 项后，英语已有音标增至 367，缺失降至 710；实际结果已在实施记录核对。前批 2 个暂缓项属于这 710 项的一部分，不额外相加。其他 708 项未在本批逐项研究，不代表不适合继续补充。

## 来源与口音

遵循 [AGENTS.md](../AGENTS.md#pronunciation)、[字段指南](ENTRY_SCHEMA.md#pronunciation)和[维护提示词](MAINTENANCE_PROMPTS.md)。优先美式；可靠美式证据暂不可得时可接受英式，但须明确标注和说明原因。本批所有值均标为 US，不混入英式音素或另一词典的转写。

30 项最终证据均来自实际打开并阅读的 Collins Dictionary 对应词条的 American English 分区，署名为 Webster’s New World College Dictionary, 5th Digital Edition。每个目标在下表都有具体 URL，完整 Notes 行另列访问日 2026-11-18 和口音。没有使用搜索摘要、拼写推断、汉语谐音或合成录音反推 IPA。

保留来源的音素和重音，包括 /i/、/u/、/r/、/ɛ/、/ɑ/ 和次重音，不为追求表面统一而改成其他转写体系。egg、mushroom、tomato、water、forget、detach、examine 和 descend 等有多个美式变体，本批选其中一个并说明。所选形式不意味着其他正常变体错误。

本轮没有收听词典音频或现有录音，没有核验连读、窄式音值、所有地域变体或原录音与 IPA 的吻合程度。新增 IPA 不授权修改已有 MP3、语音配置、语言标签或 TTSLanguage。

## 具体音标

用户于 2026-11-18 确认 B03-P01–B03-P30 并明确授权实施。各项实施前 pronunciation 均省略，现已按 add 动作写入下表批准值及同编号来源行。Word、词性和 Meaning 引自当前原文，仅用于匹配，不是改写提案。每项批准范围包含完整 pronunciation 值和同编号的 Notes 来源行，不涉及其他字段。

| 编号    | 稳定 ID / Word                                                 | 当前词性 | 当前 Meaning       | 拟写 pronunciation | 词典来源                                                                           |
|:--------|:---------------------------------------------------------------|:---------|:-------------------|:-------------------|:-----------------------------------------------------------------------------------|
| B03-P01 | [apple-food](../entries/apple-food.yml) / apple                | noun     | 苹果。             | `US /ˈæpəl/`       | [Collins: apple](https://www.collinsdictionary.com/dictionary/english/apple)       |
| B03-P02 | [banana-food](../entries/banana-food.yml) / banana             | noun     | 香蕉。             | `US /bəˈnænə/`     | [Collins: banana](https://www.collinsdictionary.com/dictionary/english/banana)     |
| B03-P03 | [butter-food](../entries/butter-food.yml) / butter             | noun     | 黄油。             | `US /ˈbʌtər/`      | [Collins: butter](https://www.collinsdictionary.com/dictionary/english/butter)     |
| B03-P04 | [carrot-food](../entries/carrot-food.yml) / carrot             | noun     | 胡萝卜。           | `US /ˈkærət/`      | [Collins: carrot](https://www.collinsdictionary.com/dictionary/english/carrot)     |
| B03-P05 | [cheese-food](../entries/cheese-food.yml) / cheese             | noun     | 奶酪。             | `US /tʃiz/`        | [Collins: cheese](https://www.collinsdictionary.com/dictionary/english/cheese)     |
| B03-P06 | [egg-food](../entries/egg-food.yml) / egg                      | noun     | 蛋。               | `US /ɛɡ/`          | [Collins: egg](https://www.collinsdictionary.com/dictionary/english/egg)           |
| B03-P07 | [honey-food](../entries/honey-food.yml) / honey                | noun     | 蜂蜜。             | `US /ˈhʌni/`       | [Collins: honey](https://www.collinsdictionary.com/dictionary/english/honey)       |
| B03-P08 | [milk-food](../entries/milk-food.yml) / milk                   | noun     | 奶；乳汁。         | `US /mɪlk/`        | [Collins: milk](https://www.collinsdictionary.com/dictionary/english/milk)         |
| B03-P09 | [mushroom-food](../entries/mushroom-food.yml) / mushroom       | noun     | 蘑菇               | `US /ˈmʌʃˌrum/`    | [Collins: mushroom](https://www.collinsdictionary.com/dictionary/english/mushroom) |
| B03-P10 | [rice-food](../entries/rice-food.yml) / rice                   | noun     | 米；米饭。         | `US /raɪs/`        | [Collins: rice](https://www.collinsdictionary.com/dictionary/english/rice)         |
| B03-P11 | [salt-food](../entries/salt-food.yml) / salt                   | noun     | 盐。               | `US /sɔlt/`        | [Collins: salt](https://www.collinsdictionary.com/dictionary/english/salt)         |
| B03-P12 | [sugar-food](../entries/sugar-food.yml) / sugar                | noun     | 糖。               | `US /ˈʃʊɡər/`      | [Collins: sugar](https://www.collinsdictionary.com/dictionary/english/sugar)       |
| B03-P13 | [tomato-food](../entries/tomato-food.yml) / tomato             | noun     | 番茄。             | `US /təˈmeɪtoʊ/`   | [Collins: tomato](https://www.collinsdictionary.com/dictionary/english/tomato)     |
| B03-P14 | [water-general](../entries/water-general.yml) / water          | noun     | 水。               | `US /ˈwɔtər/`      | [Collins: water](https://www.collinsdictionary.com/dictionary/english/water)       |
| B03-P15 | [fruit-food](../entries/fruit-food.yml) / fruit                | noun     | 水果。             | `US /frut/`        | [Collins: fruit](https://www.collinsdictionary.com/dictionary/english/fruit)       |
| B03-P16 | [cancel-general](../entries/cancel-general.yml) / cancel       | verb     | 取消               | `US /ˈkænsəl/`     | [Collins: cancel](https://www.collinsdictionary.com/dictionary/english/cancel)     |
| B03-P17 | [find-general](../entries/find-general.yml) / find             | verb     | 找到；发现。       | `US /faɪnd/`       | [Collins: find](https://www.collinsdictionary.com/dictionary/english/find)         |
| B03-P18 | [forget-general](../entries/forget-general.yml) / forget       | verb     | 忘记；未能想起。   | `US /fərˈɡɛt/`     | [Collins: forget](https://www.collinsdictionary.com/dictionary/english/forget)     |
| B03-P19 | [select-general](../entries/select-general.yml) / select       | verb     | 选择               | `US /səˈlɛkt/`     | [Collins: select](https://www.collinsdictionary.com/dictionary/english/select)     |
| B03-P20 | [climb-general](../entries/climb-general.yml) / climb          | verb     | 攀爬               | `US /klaɪm/`       | [Collins: climb](https://www.collinsdictionary.com/dictionary/english/climb)       |
| B03-P21 | [throw-general](../entries/throw-general.yml) / throw          | verb     | 投掷               | `US /θroʊ/`        | [Collins: throw](https://www.collinsdictionary.com/dictionary/english/throw)       |
| B03-P22 | [hold-general](../entries/hold-general.yml) / hold             | verb     | 手持、持有         | `US /hoʊld/`       | [Collins: hold](https://www.collinsdictionary.com/dictionary/english/hold)         |
| B03-P23 | [drop-general](../entries/drop-general.yml) / drop             | verb     | 丢弃               | `US /drɑp/`        | [Collins: drop](https://www.collinsdictionary.com/dictionary/english/drop)         |
| B03-P24 | [attach-general](../entries/attach-general.yml) / attach       | verb     | 连接、附加         | `US /əˈtætʃ/`      | [Collins: attach](https://www.collinsdictionary.com/dictionary/english/attach)     |
| B03-P25 | [detach-general](../entries/detach-general.yml) / detach       | verb     | 分离、拆卸         | `US /dɪˈtætʃ/`     | [Collins: detach](https://www.collinsdictionary.com/dictionary/english/detach)     |
| B03-P26 | [assemble-general](../entries/assemble-general.yml) / assemble | verb     | 组装               | `US /əˈsɛmbəl/`    | [Collins: assemble](https://www.collinsdictionary.com/dictionary/english/assemble) |
| B03-P27 | [examine-general](../entries/examine-general.yml) / examine    | verb     | 检查               | `US /ɪɡˈzæmən/`    | [Collins: examine](https://www.collinsdictionary.com/dictionary/english/examine)   |
| B03-P28 | [destroy-general](../entries/destroy-general.yml) / destroy    | verb     | 摧毁；毁灭。       | `US /dɪˈstrɔɪ/`    | [Collins: destroy](https://www.collinsdictionary.com/dictionary/english/destroy)   |
| B03-P29 | [ascend-general](../entries/ascend-general.yml) / ascend       | verb     | 上升；向高处移动。 | `US /əˈsɛnd/`      | [Collins: ascend](https://www.collinsdictionary.com/dictionary/english/ascend)     |
| B03-P30 | [descend-general](../entries/descend-general.yml) / descend    | verb     | 下降               | `US /dɪˈsɛnd/`     | [Collins: descend](https://www.collinsdictionary.com/dictionary/english/descend)   |

## 逐项核对

以下是读音和词形核对依据，不是额外助记或 Notes 改写授权。只追加下一节的明确来源行。

| 编号    | 核对结果与限制                                                                        |
|:--------|:--------------------------------------------------------------------------------------|
| B03-P01 | 苹果义名词，不推定品牌或专名的读法。                                                  |
| B03-P02 | 香蕉义名词，保持单数，不添加 bananas 的词尾。                                         |
| B03-P03 | 黄油义名词；保留来源 /tər/，不自行改成另一套窄式转写。                                |
| B03-P04 | 胡萝卜义名词，不使用引诱或奖励的比喻义解释当前条目。                                  |
| B03-P05 | 采用 cheese in American English 1 的奶酪义，保留来源 /i/ 记法。                       |
| B03-P06 | 采用 egg in American English 1 的蛋义；来源也列有 /eɪɡ/，所选变体不代表另一读法错误。 |
| B03-P07 | 蜂蜜义名词，不使用昵称或爱称义解释当前内容。                                          |
| B03-P08 | 奶或乳汁义名词，不套用 milked 的动词过去式。                                          |
| B03-P09 | 蘑菇义名词；选择来源列出的 /rum/ 变体并保留次重音，另有 /rʊm/ 变体。                  |
| B03-P10 | 米或米饭义名词，不使用同页 Rice 人名义。                                              |
| B03-P11 | 食盐义名词，保留来源美式元音；不采用 SALT 缩写义。                                    |
| B03-P12 | 糖义名词；采用所选来源 /ʊ/，不混用同页另一词典的 /u/。                                |
| B03-P13 | 番茄义单数名词；选择来源列出的 /eɪ/ 变体，不据此排除其另一美式变体。                  |
| B03-P14 | 水义名词；选择来源列出的 /ɔ/ 变体，同页还列 /ɑ/，不将其判错。                         |
| B03-P15 | 水果义名词，保留来源 /u/，不加复数或派生词尾。                                        |
| B03-P16 | 取消义动词原形，不使用 canceled 或 cancelled 的过去式读音。                           |
| B03-P17 | 找到或发现义动词原形，不使用 found 的读音。                                           |
| B03-P18 | 忘记义动词原形；选择来源所列 /fər-/ 变体，不混入 forgot 或 forgotten。                |
| B03-P19 | 选择义动词；来源同一头词下包含形容词，当前条目仍保留 verb，不改词性。                 |
| B03-P20 | 攀爬义动词原形；依据词典转写，不能按拼写额外读出 b。                                  |
| B03-P21 | 投掷义动词原形，不使用 threw 或 thrown 的读音。                                       |
| B03-P22 | 手持或持有义动词，对应 hold in American English 1，不使用货舱义解释当前内容。         |
| B03-P23 | 丢弃义动词原形；同一头词也有名词义，不将当前词性改为 noun。                           |
| B03-P24 | 连接或附加义动词原形，不改成 attached 的词形。                                        |
| B03-P25 | 分离或拆卸义动词，选择来源所列 /dɪ-/ 变体；不修改另外的 detach-git。                  |
| B03-P26 | 组装义动词，匹配把部件装配起来的义项；不将汇编语言的专门义强加到本条。                |
| B03-P27 | 检查义动词，选择来源所列 /ɪɡ-/ 变体；保留该来源词尾 /ən/，不与其他词典拼接。          |
| B03-P28 | 摧毁或毁灭义动词原形，不使用 destroyed 的过去式形式。                                 |
| B03-P29 | 上升义普通动词；不修改同形的游戏能力专名 Ascend 或其词条。                            |
| B03-P30 | 下降义动词，选择来源所列 /dɪ-/ 变体；不混入 descended 的词尾。                        |

## Notes 来源追加文案

保留原 Notes 全文。已有 Notes 时在末尾追加相应一行，原字段省略时才创建；若实施时已有同一来源行，则不重复追加。相同 URL 若已用于其他解释，应先检查是否已记录音标及口音，再避免重复说明。以下纯文本为逐项完整文案，Markdown 围栏不写入字段。

### B03-P01 — apple

```plain
Pronunciation source: Collins Dictionary — apple; apple in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/apple (accent: US; accessed 2026-11-18).
```

### B03-P02 — banana

```plain
Pronunciation source: Collins Dictionary — banana; banana in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/banana (accent: US; accessed 2026-11-18).
```

### B03-P03 — butter

```plain
Pronunciation source: Collins Dictionary — butter; butter in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/butter (accent: US; accessed 2026-11-18).
```

### B03-P04 — carrot

```plain
Pronunciation source: Collins Dictionary — carrot; carrot in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/carrot (accent: US; accessed 2026-11-18).
```

### B03-P05 — cheese

```plain
Pronunciation source: Collins Dictionary — cheese; cheese in American English 1; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/cheese (accent: US; accessed 2026-11-18).
```

### B03-P06 — egg

```plain
Pronunciation source: Collins Dictionary — egg; egg in American English 1; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/egg (accent: US; accessed 2026-11-18).
```

### B03-P07 — honey

```plain
Pronunciation source: Collins Dictionary — honey; honey in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/honey (accent: US; accessed 2026-11-18).
```

### B03-P08 — milk

```plain
Pronunciation source: Collins Dictionary — milk; milk in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/milk (accent: US; accessed 2026-11-18).
```

### B03-P09 — mushroom

```plain
Pronunciation source: Collins Dictionary — mushroom; mushroom in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/mushroom (accent: US; accessed 2026-11-18).
```

### B03-P10 — rice

```plain
Pronunciation source: Collins Dictionary — rice; rice in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/rice (accent: US; accessed 2026-11-18).
```

### B03-P11 — salt

```plain
Pronunciation source: Collins Dictionary — salt; salt in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/salt (accent: US; accessed 2026-11-18).
```

### B03-P12 — sugar

```plain
Pronunciation source: Collins Dictionary — sugar; sugar in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/sugar (accent: US; accessed 2026-11-18).
```

### B03-P13 — tomato

```plain
Pronunciation source: Collins Dictionary — tomato; tomato in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/tomato (accent: US; accessed 2026-11-18).
```

### B03-P14 — water

```plain
Pronunciation source: Collins Dictionary — water; water in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/water (accent: US; accessed 2026-11-18).
```

### B03-P15 — fruit

```plain
Pronunciation source: Collins Dictionary — fruit; fruit in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/fruit (accent: US; accessed 2026-11-18).
```

### B03-P16 — cancel

```plain
Pronunciation source: Collins Dictionary — cancel; cancel in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/cancel (accent: US; accessed 2026-11-18).
```

### B03-P17 — find

```plain
Pronunciation source: Collins Dictionary — find; find in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/find (accent: US; accessed 2026-11-18).
```

### B03-P18 — forget

```plain
Pronunciation source: Collins Dictionary — forget; forget in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/forget (accent: US; accessed 2026-11-18).
```

### B03-P19 — select

```plain
Pronunciation source: Collins Dictionary — select; select in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/select (accent: US; accessed 2026-11-18).
```

### B03-P20 — climb

```plain
Pronunciation source: Collins Dictionary — climb; climb in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/climb (accent: US; accessed 2026-11-18).
```

### B03-P21 — throw

```plain
Pronunciation source: Collins Dictionary — throw; throw in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/throw (accent: US; accessed 2026-11-18).
```

### B03-P22 — hold

```plain
Pronunciation source: Collins Dictionary — hold; hold in American English 1; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/hold (accent: US; accessed 2026-11-18).
```

### B03-P23 — drop

```plain
Pronunciation source: Collins Dictionary — drop; drop in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/drop (accent: US; accessed 2026-11-18).
```

### B03-P24 — attach

```plain
Pronunciation source: Collins Dictionary — attach; attach in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/attach (accent: US; accessed 2026-11-18).
```

### B03-P25 — detach

```plain
Pronunciation source: Collins Dictionary — detach; detach in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/detach (accent: US; accessed 2026-11-18).
```

### B03-P26 — assemble

```plain
Pronunciation source: Collins Dictionary — assemble; assemble in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/assemble (accent: US; accessed 2026-11-18).
```

### B03-P27 — examine

```plain
Pronunciation source: Collins Dictionary — examine; examine in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/examine (accent: US; accessed 2026-11-18).
```

### B03-P28 — destroy

```plain
Pronunciation source: Collins Dictionary — destroy; destroy in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/destroy (accent: US; accessed 2026-11-18).
```

### B03-P29 — ascend

```plain
Pronunciation source: Collins Dictionary — ascend; ascend in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/ascend (accent: US; accessed 2026-11-18).
```

### B03-P30 — descend

```plain
Pronunciation source: Collins Dictionary — descend; descend in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/descend (accent: US; accessed 2026-11-18).
```

## 保留与后续范围

- 前批暂缓的 [part-of-speech-language](../entries/part-of-speech-language.yml)和[market-value-finance](../entries/market-value-finance.yml)保持不变。沿用[第二批暂缓记录](EN_PRONUNCIATION_ENRICHMENT_BATCH_02_PLAN.md#暂缓项目)，本批没有重新查证这两个整体短语；不能因接受英式就自行拼接各词读音。
- 本批没有提出新的短语、专名、缩写或虚构词转写，不改写先前已有的音标和口音标签。其他 708 个缺失项目留待后续筛选；没有筛除或删除任何词条。
- detach-general 与 detach-git、ascend-general 与游戏能力专名等同形项目按各自稳定 ID 区分。本批只修改提案表中的目标，不将相同拼写作为批量覆盖依据。
- 现有关联、先修、优先级和完整来源文字保留。attach/detach、ascend/descend 等在本批同时出现，仅表示各自值得补音标，不授权增加关联或重写助记。

## 批准后实施与验证

1. 重新读取本计划、当前规范及实现，检查 Git 状态和规划以来的差异。按稳定 ID 核对目标及批准编号；若目标已经有音标、来源或其他并发编辑，先协调差异，不盲目覆盖。
2. 在仓库外快照全库源原文、解析值、当前导出及所有现有和归档录音哈希。只写入批准的 pronunciation 字符串及各自 Notes 来源行，逐项核对编号、Word、词性、口音与 URL，防止来源串用。
3. 保留稳定 ID、Word、Meaning、语言、词性、Context、Example、既有 Notes、辅助字段、关联、学习元数据、标签和音频引用。只在缺少 Notes 的获批目标上创建该字段；不新增条目、schema、工具、模板或依赖。
4. 不生成、替换或复制录音，也不从 IPA 推断实际录音口音。需要改录音时另提问题并取得独立授权。原有及归档 MP3 字节必须不变。
5. 使用现有命令运行验证和构建，不手改 dist。

```bash
go test ./...
go vet ./...
go run scripts/validate/main.go
go run scripts/build/main.go
git diff --check
```

6. 在源不变时再构建一次并比较字节。核对基线稳定 ID 集、数据行数及 17 列 TSV；逐项核对完整 US 标签和 Notes 来源。只允许获批行的 Pronunciation 与 Notes 变化，Audio、StudyOrder、RelatedWords、标签及其他列不变。实施时若集合已变化，先核对新的实际基线，不强行套用历史计数。
7. 检查源差异均在批准字段范围内，录音哈希不变，原有无关工作保留。对未跟踪文本另查格式，复查本地链接、锚点及实际剩余缺失数。更新本计划的批准、实施和检查状态，前两批记录不变。
8. 不暂存、提交、推送或修改 Anki 集合。静态导出验证、人工听审和真实 Anki 导入、显示、播放分别报告，不声称执行未做过的检查。

## 规划检查

以下为实施前完成的规划检查，保留当时基线和操作边界。

- 当前 1,320 个稳定 ID 唯一，英语已有 337 项音标、缺失 740 项；本批 30 个 ID 均存在并缺少 pronunciation，且与已完成前两批目标不重叠。15 个名词和 15 个动词的当前词性及释义与源数据一致。
- 30 个具体值及来源行按编号、Word、URL 和 US 标签逐项对应，来源均为已读正文。前批 2 个暂缓项在本批中仅保留历史状态，不计为本轮重新查证或新增提案。
- 本计划的本地链接、锚点、表格列数、围栏闭合、行尾空白、制表符和最终换行检查通过，git diff --check 通过；未跟踪计划另行检查文本格式。
- 与规划前哈希快照比较，所有既有项目文件不变，包括 1,320 个词条、1,320 个现有 MP3、43 个归档 MP3、先前规范改动、前两批计划、实现、依赖、模板和 dist。唯一新增文件为本计划。
- 规划阶段未运行 Go 测试、词库验证、构建、音频生成、解码、人工听审或 Anki 操作；批准后实际检查结果见实施记录。

已确认并实施: B03-P01–B03-P30 的完整 pronunciation 值及各自 Notes 来源追加文案。前批两个暂缓项保持不变，本次确认不授权重写其他音标或替换录音。


## 实施记录

2026-11-18，按用户确认及明确实施指令完成 B03-P01–B03-P30。15 个日常名词和 15 个常用动作词新增了批准的 US 音标和各自词典来源行，其中 1 个原本没有 Notes 的条目创建了该字段；其余完整保留原 Notes 并追加来源。没有新增词条、录音、关联、先修或其他辅助字段。

- 逐项比对批准计划中的编号、稳定 ID、Word、词性、Meaning、音标、口音及完整来源行，30 项均一致。每条来源行只出现一次，没有发生不同词条之间的来源串用。
- 全库仍为 1,320 个稳定 ID，英语 1,077 项、日语 243 项。英语已有音标从 337 增至 367，缺失从 740 降至 710。前两批全部值和日语假名保持不变，两个暂缓项目仍未补写。
- 源字段变化仅限批准的 pronunciation 新增和 Notes 来源追加。其他 1,290 个词条的源字节完全不变，1,320 个现有 MP3 和 43 个归档 MP3 的 SHA-256 均保持不变；此前已有词条、规范和计划改动得到保留。
- go test ./...、go vet ./...、词库验证、两次构建及 git diff --check 均通过。原有 15 条同形词警告保留，没有新增同形词警告。修改文本及未跟踪计划的格式、本地链接与锚点另行检查通过。
- 两次无源变更的构建逐字节一致，导出 SHA-256 为 `4b656cfc29d23637cba3b02244d91fdef14bc73297f39b9342977ff1c447e143`。与实施前 TSV 逐行逐列比对，仍为 1,320 行、17 列，仅获批 30 行的 Pronunciation 和 Notes 改变。Audio、StudyOrder、RelatedWords、标签和其他列不变，US 标签作为完整普通文本导出。未手改 dist。

本次没有生成、替换或解码录音，没有进行人工听审、真实 Anki 导入、显示或播放验证，不声称已有录音与所选 IPA 经听审完全吻合。未暂存、提交、推送、复制媒体或修改 Anki 集合。
