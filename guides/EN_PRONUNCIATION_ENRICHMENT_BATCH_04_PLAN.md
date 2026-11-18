# 英语音标补全计划: 第四批日常名词与形容词

状态: **B04-P01–B04-P30 的具体音标与来源文案已获批准并实施，静态验证完成。** 日期: 2026-11-18。人工听审及真实 Anki 验证尚未进行。

规划阶段仅创建本计划。用户随后明确确认 B04-P01–P30 并授权实施，本轮只写入批准的音标及 Notes 来源行，并通过现有命令重建导出。media、archives、规范、实现、模板和依赖保持不变。不暂存、提交、推送或操作 Anki。

[第一批](EN_PRONUNCIATION_ENRICHMENT_BATCH_01_PLAN.md)、[第二批](EN_PRONUNCIATION_ENRICHMENT_BATCH_02_PLAN.md)和[第三批](EN_PRONUNCIATION_ENRICHMENT_BATCH_03_PLAN.md)保留为已完成的历史记录。第四批编号独立为 B04-P01–B04-P30，不复用先前批准编号。

## 导航

- [范围与基线](#范围与基线)
- [来源与转写规则](#来源与转写规则)
- [具体音标](#具体音标)
- [逐项核对](#逐项核对)
- [Notes 来源追加文案](#notes-来源追加文案)
- [听审与暂缓边界](#听审与暂缓边界)
- [批准后实施与验证](#批准后实施与验证)
- [规划检查](#规划检查)
- [实施与验证记录](#实施与验证记录)

## 范围与基线

按实际稳定 ID 盘点全库 pronunciation，对本批逐项匹配 Word、词性与 Meaning。完整原字段和文件哈希在仓库外记录，供保全检查。字段盘点不代表全库音标已经逐项研究。

| 项目                          | 实际范围                               |
|:------------------------------|:---------------------------------------|
| 全库词条                      | 1,320                                  |
| 英语 / 日语                   | 1,077 / 243                            |
| 英语已有 / 缺失 pronunciation | 367 / 710                              |
| 日语已有 / 缺失 pronunciation | 243 / 0                                |
| 本批具体提案                  | 30 个已有词条                          |
| 词性分组                      | 15 个名词，15 个形容词                 |
| 口音                          | 30 项 US，0 项 UK                      |
| 新增暂缓项                    | 0                                      |
| 保留的前批暂缓项              | 2，不计入本批 30 项                    |
| 批准后允许修改                | 仅 pronunciation 和各自 Notes 来源追加 |
| 新增词条 / 录音 / 关联 / 先修 | 全部为 0                               |

优先选择身体、物品、自然与食物名词，以及常见状态、品质和感官形容词。tear 与 bass 的读音依赖当前义项，单列核对，避免按拼写或词典首个分区误选。没有为补齐数量而添加专业词、专名、缩写或长句；本批次序不改变学习元数据或 StudyOrder。

规划时英语缺失为 710。实施这 30 项后，英语已有音标增至 397，缺失降至 680。前批 2 个暂缓项包含在这 680 项中，不重复相加；其他 678 个缺失项目不在本批逐项查证范围内。上表保留实施前基线。

## 来源与转写规则

沿用 [AGENTS.md](../AGENTS.md#pronunciation)和[字段指南](ENTRY_SCHEMA.md#pronunciation)的规则。优先美式，可靠美式证据不可得时可接受明确标注的英式；本批 30 项都有可访问的美式文字证据，全部拟写 US 标签，未使用英式替代。

最终证据来自实际打开并阅读的 Collins Dictionary 词条，美式正文署名 Webster’s New World College Dictionary, 5th Digital Edition。下表提供逐项 URL，Notes 行另列分区、口音和访问日 2026-11-18。不是搜索摘要、汉语谐音、拼写猜测或从合成录音反推。

- wheat 的普通词典网址初次可访问，随后定位美式段落时两次超时；最终通过成功读取的 Collins US 网址取得美式正文。来源行引用实际核实的 US 网址，不把超时当作没有美式读音。
- 保留来源记法，例如 /i/、/u/、/r/、/ɛ/、/ɜr/ 及 /oʊ/，不混拼不同口音或词典。root、wheat、durable 等存在多个正常美式变体，本批选择其一，不将其他列出变体判为错误。
- hidden 的所选正文呈现为 "ˈhɪd ə n"，拟值仅合并音素间版面空格为 /ˈhɪdən/，未添加、替换或删除音素。采用当前形容词形式，不写成 hide 的读音。
- tear、bass 和 ring 的分区编号按所选 WNW 词典定位；同页其他词典的编号可能不同，不能只按全页第一个同形词选读音。来源行明确分区，具体解释只留在审阅计划，不授权额外助记改写。
- 未收听词典音频或现有 MP3，未核实所有连读、窄式音值和地域变体。音标与现有录音是否吻合仍需听审；新增 IPA 不授权重录、改语音或调整语言标签和 TTSLanguage。

## 具体音标

B04-P01–B04-P30 均为 add，当前 pronunciation 均省略。Word、词性和 Meaning 引自当前源数据，不属于改写提案。每项批准范围包含完整 pronunciation 值及同编号 Notes 来源行，不涉及其他字段。

| 编号    | 稳定 ID / Word                                                 | 当前词性  | 当前 Meaning                     | 拟写 pronunciation | 词典来源                                                                           |
|:--------|:---------------------------------------------------------------|:----------|:---------------------------------|:-------------------|:-----------------------------------------------------------------------------------|
| B04-P01 | [hand-general](../entries/hand-general.yml) / hand             | noun      | 手。                             | `US /hænd/`        | [Collins: hand](https://www.collinsdictionary.com/dictionary/english/hand)         |
| B04-P02 | [head-general](../entries/head-general.yml) / head             | noun      | 头；头部。                       | `US /hɛd/`         | [Collins: head](https://www.collinsdictionary.com/dictionary/english/head)         |
| B04-P03 | [stone-general](../entries/stone-general.yml) / stone          | noun      | 石头；石料。                     | `US /stoʊn/`       | [Collins: stone](https://www.collinsdictionary.com/dictionary/english/stone)       |
| B04-P04 | [stick-general](../entries/stick-general.yml) / stick          | noun      | 棍；条状或杆状物。               | `US /stɪk/`        | [Collins: stick](https://www.collinsdictionary.com/dictionary/english/stick)       |
| B04-P05 | [root-general](../entries/root-general.yml) / root             | noun      | 根；植物的根。                   | `US /rut/`         | [Collins: root](https://www.collinsdictionary.com/dictionary/english/root)         |
| B04-P06 | [seed-general](../entries/seed-general.yml) / seed             | noun      | 种子。                           | `US /sid/`         | [Collins: seed](https://www.collinsdictionary.com/dictionary/english/seed)         |
| B04-P07 | [fire-general](../entries/fire-general.yml) / fire             | noun      | 火。                             | `US /faɪr/`        | [Collins: fire](https://www.collinsdictionary.com/dictionary/english/fire)         |
| B04-P08 | [flame-general](../entries/flame-general.yml) / flame          | noun      | 火焰。                           | `US /fleɪm/`       | [Collins: flame](https://www.collinsdictionary.com/dictionary/english/flame)       |
| B04-P09 | [ring-general](../entries/ring-general.yml) / ring             | noun      | 环；圈；戒指。                   | `US /rɪŋ/`         | [Collins: ring](https://www.collinsdictionary.com/dictionary/english/ring)         |
| B04-P10 | [mirror-general](../entries/mirror-general.yml) / mirror       | noun      | 镜子。                           | `US /ˈmɪrər/`      | [Collins: mirror](https://www.collinsdictionary.com/dictionary/english/mirror)     |
| B04-P11 | [meat-food](../entries/meat-food.yml) / meat                   | noun      | 肉；供食用的肉类。               | `US /mit/`         | [Collins: meat](https://www.collinsdictionary.com/dictionary/english/meat)         |
| B04-P12 | [oil-food](../entries/oil-food.yml) / oil                      | noun      | 油。                             | `US /ɔɪl/`         | [Collins: oil](https://www.collinsdictionary.com/dictionary/english/oil)           |
| B04-P13 | [wheat-food](../entries/wheat-food.yml) / wheat                | noun      | 小麦。                           | `US /wit/`         | [Collins: wheat](https://www.collinsdictionary.com/us/dictionary/english/wheat)    |
| B04-P14 | [tear-general](../entries/tear-general.yml) / tear             | noun      | 眼泪；泪水。                     | `US /tɪr/`         | [Collins: tear](https://www.collinsdictionary.com/dictionary/english/tear)         |
| B04-P15 | [bass-food](../entries/bass-food.yml) / bass                   | noun      | 鲈鱼。                           | `US /bæs/`         | [Collins: bass](https://www.collinsdictionary.com/dictionary/english/bass)         |
| B04-P16 | [brave-general](../entries/brave-general.yml) / brave          | adjective | 勇敢的                           | `US /breɪv/`       | [Collins: brave](https://www.collinsdictionary.com/dictionary/english/brave)       |
| B04-P17 | [curious-general](../entries/curious-general.yml) / curious    | adjective | 好奇的                           | `US /ˈkjʊriəs/`    | [Collins: curious](https://www.collinsdictionary.com/dictionary/english/curious)   |
| B04-P18 | [loyal-general](../entries/loyal-general.yml) / loyal          | adjective | 忠诚的                           | `US /ˈlɔɪəl/`      | [Collins: loyal](https://www.collinsdictionary.com/dictionary/english/loyal)       |
| B04-P19 | [wise-general](../entries/wise-general.yml) / wise             | adjective | 睿智的                           | `US /waɪz/`        | [Collins: wise](https://www.collinsdictionary.com/dictionary/english/wise)         |
| B04-P20 | [powerful-general](../entries/powerful-general.yml) / powerful | adjective | 强大的                           | `US /ˈpaʊərfəl/`   | [Collins: powerful](https://www.collinsdictionary.com/dictionary/english/powerful) |
| B04-P21 | [sunny-general](../entries/sunny-general.yml) / sunny          | adjective | 阳光明媚的；充满阳光的。         | `US /ˈsʌni/`       | [Collins: sunny](https://www.collinsdictionary.com/dictionary/english/sunny)       |
| B04-P22 | [chilly-general](../entries/chilly-general.yml) / chilly       | adjective | 寒冷的；凉飕飕的。               | `US /ˈtʃɪli/`      | [Collins: chilly](https://www.collinsdictionary.com/dictionary/english/chilly)     |
| B04-P23 | [raw-food](../entries/raw-food.yml) / raw                      | adjective | 生的；未烹饪的。                 | `US /rɔ/`          | [Collins: raw](https://www.collinsdictionary.com/dictionary/english/raw)           |
| B04-P24 | [spicy-general](../entries/spicy-general.yml) / spicy          | adjective | 辛辣的；带香辛料味道的。         | `US /ˈspaɪsi/`     | [Collins: spicy](https://www.collinsdictionary.com/dictionary/english/spicy)       |
| B04-P25 | [durable-general](../entries/durable-general.yml) / durable    | adjective | 耐用的；能经受使用而不易损坏的。 | `US /ˈdʊrəbəl/`    | [Collins: durable](https://www.collinsdictionary.com/dictionary/english/durable)   |
| B04-P26 | [sturdy-general](../entries/sturdy-general.yml) / sturdy       | adjective | 坚固的                           | `US /ˈstɜrdi/`     | [Collins: sturdy](https://www.collinsdictionary.com/dictionary/english/sturdy)     |
| B04-P27 | [rusty-general](../entries/rusty-general.yml) / rusty          | adjective | 生锈的                           | `US /ˈrʌsti/`      | [Collins: rusty](https://www.collinsdictionary.com/dictionary/english/rusty)       |
| B04-P28 | [hidden-general](../entries/hidden-general.yml) / hidden       | adjective | 隐藏的                           | `US /ˈhɪdən/`      | [Collins: hidden](https://www.collinsdictionary.com/dictionary/english/hidden)     |
| B04-P29 | [golden-general](../entries/golden-general.yml) / golden       | adjective | 金色的；金黄色的。               | `US /ˈɡoʊldən/`    | [Collins: golden](https://www.collinsdictionary.com/dictionary/english/golden)     |
| B04-P30 | [tough-general](../entries/tough-general.yml) / tough          | adjective | 坚韧的；结实的；难以对付的。     | `US /tʌf/`         | [Collins: tough](https://www.collinsdictionary.com/dictionary/english/tough)       |

## 逐项核对

以下内容为审阅依据，除下一节列出的来源行外，不授权写入 Notes、Mnemonic 或其他字段。

| 编号    | 核对结果与限制                                                                                                  |
|:--------|:----------------------------------------------------------------------------------------------------------------|
| B04-P01 | 手的名词义，不取 Hand 人名或 hand- 构词成分。                                                                   |
| B04-P02 | 头部义名词，取小写 head 的正文，不以 Head 人名分区作为来源。                                                    |
| B04-P03 | 石头及石料义名词，不采用 Stone 人名或重量单位的释义。                                                           |
| B04-P04 | 棍或条状物义名词，保持当前 noun，不改成粘住的动词义。                                                           |
| B04-P05 | 植物根义，对应 root in American English 1；选来源 /rut/ 变体，来源也列 /rʊt/。                                  |
| B04-P06 | 种子义名词原形，保留来源 /i/ 记法，不自行添 seeds 的词尾。                                                      |
| B04-P07 | 火的名词义，不使用解雇或开火的动词义解释当前条目。                                                              |
| B04-P08 | 火焰义名词，不套用网上争吵的引申义。                                                                            |
| B04-P09 | 环或戒指义名词对应所选词典 ring in American English 2；不要误定位到其鸣响动词分区。                             |
| B04-P10 | 镜子义名词，保留来源的完整卷舌转写。                                                                            |
| B04-P11 | 供食用的肉类义，保留来源 /i/，不采用文章主旨的比喻义。                                                          |
| B04-P12 | 油的名词义，当前食物语境保留，不改为石油专门词条。                                                              |
| B04-P13 | 小麦义名词；取来源所列 /wit/，其 /hwit/ 也是列出的美式变体。普通网址查证超时后改用成功访问的 Collins US 页面。  |
| B04-P14 | 眼泪义名词对应所选词典 tear in American English 2，不能使用撕裂义的 /tɛr/。现有录音未试听，另保留辨义听审待办。 |
| B04-P15 | 鲈鱼义对应 bass in American English 2，不采用低音或低音乐器义的 /beɪs/。现有录音未试听，另保留辨义听审待办。    |
| B04-P16 | 勇敢的形容词原级，不套用 braver、bravest 或动词过去式词尾。                                                     |
| B04-P17 | 好奇义形容词；保留所选来源 /ʊ/，不与同页另一词典 /u/ 记法拼合。                                                 |
| B04-P18 | 忠诚义形容词，取当前词形。                                                                                      |
| B04-P19 | 睿智义形容词对应 wise in American English 1，不采用 -wise 后缀、人名或名词 manner 义。                          |
| B04-P20 | 强大的形容词义，保留来源完整 /aʊər/ 序列，不改写成另一套窄式转写。                                              |
| B04-P21 | 阳光明媚的形容词原级，不套用 sunnier 或 sunniest。                                                              |
| B04-P22 | 寒冷或凉飕飕义形容词，不改成冷淡待人的释义。                                                                    |
| B04-P23 | 未烹饪义形容词，保留当前食物意义，不改为 raw data 术语。                                                        |
| B04-P24 | 辛辣或带香辛料味的形容词，不使用引申的刺激性内容义。                                                            |
| B04-P25 | 耐用义形容词；选择来源不含 /j/ 的美式变体，其所列 /ˈdjʊrəbəl/ 不判为错误。                                      |
| B04-P26 | 坚固义形容词对应 sturdy in American English 1，不选动物疾病的名词义。                                           |
| B04-P27 | 生锈义形容词；不改成技能因缺乏练习而生疏的引申义或 Rusty 人名。                                                 |
| B04-P28 | 隐藏的形容词；来源 WNW 正文显示 ˈhɪd ə n，拟值仅合并音素间的版面空格，不添加音素，不改为 hide。                 |
| B04-P29 | 金色义形容词，保留来源 /oʊ/，不截成 gold 的读音。                                                               |
| B04-P30 | 坚韧、结实或难以对付的形容词，依据完整词典读音，不按拼写拆读 gh。                                               |

## Notes 来源追加文案

保留完整原 Notes。已有 Notes 时在末尾追加对应来源行，原字段省略时才创建；同一来源行已存在就不再追加。若相同 URL 已用于其他解释，先检查是否已记录发音证据和口音，避免重复说明。以下是逐项完整文案，围栏不写入 YAML 值。

### B04-P01 — hand

```plain
Pronunciation source: Collins Dictionary — hand; hand in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/hand (accent: US; accessed 2026-11-18).
```

### B04-P02 — head

```plain
Pronunciation source: Collins Dictionary — head; head in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/head (accent: US; accessed 2026-11-18).
```

### B04-P03 — stone

```plain
Pronunciation source: Collins Dictionary — stone; stone in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/stone (accent: US; accessed 2026-11-18).
```

### B04-P04 — stick

```plain
Pronunciation source: Collins Dictionary — stick; stick in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/stick (accent: US; accessed 2026-11-18).
```

### B04-P05 — root

```plain
Pronunciation source: Collins Dictionary — root; root in American English 1; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/root (accent: US; accessed 2026-11-18).
```

### B04-P06 — seed

```plain
Pronunciation source: Collins Dictionary — seed; seed in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/seed (accent: US; accessed 2026-11-18).
```

### B04-P07 — fire

```plain
Pronunciation source: Collins Dictionary — fire; fire in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/fire (accent: US; accessed 2026-11-18).
```

### B04-P08 — flame

```plain
Pronunciation source: Collins Dictionary — flame; flame in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/flame (accent: US; accessed 2026-11-18).
```

### B04-P09 — ring

```plain
Pronunciation source: Collins Dictionary — ring; ring in American English 2; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/ring (accent: US; accessed 2026-11-18).
```

### B04-P10 — mirror

```plain
Pronunciation source: Collins Dictionary — mirror; mirror in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/mirror (accent: US; accessed 2026-11-18).
```

### B04-P11 — meat

```plain
Pronunciation source: Collins Dictionary — meat; meat in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/meat (accent: US; accessed 2026-11-18).
```

### B04-P12 — oil

```plain
Pronunciation source: Collins Dictionary — oil; oil in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/oil (accent: US; accessed 2026-11-18).
```

### B04-P13 — wheat

```plain
Pronunciation source: Collins Dictionary — wheat; wheat in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/us/dictionary/english/wheat (accent: US; accessed 2026-11-18).
```

### B04-P14 — tear

```plain
Pronunciation source: Collins Dictionary — tear; tear in American English 2; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/tear (accent: US; accessed 2026-11-18).
```

### B04-P15 — bass

```plain
Pronunciation source: Collins Dictionary — bass; bass in American English 2; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/bass (accent: US; accessed 2026-11-18).
```

### B04-P16 — brave

```plain
Pronunciation source: Collins Dictionary — brave; brave in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/brave (accent: US; accessed 2026-11-18).
```

### B04-P17 — curious

```plain
Pronunciation source: Collins Dictionary — curious; curious in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/curious (accent: US; accessed 2026-11-18).
```

### B04-P18 — loyal

```plain
Pronunciation source: Collins Dictionary — loyal; loyal in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/loyal (accent: US; accessed 2026-11-18).
```

### B04-P19 — wise

```plain
Pronunciation source: Collins Dictionary — wise; wise in American English 1; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/wise (accent: US; accessed 2026-11-18).
```

### B04-P20 — powerful

```plain
Pronunciation source: Collins Dictionary — powerful; powerful in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/powerful (accent: US; accessed 2026-11-18).
```

### B04-P21 — sunny

```plain
Pronunciation source: Collins Dictionary — sunny; sunny in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/sunny (accent: US; accessed 2026-11-18).
```

### B04-P22 — chilly

```plain
Pronunciation source: Collins Dictionary — chilly; chilly in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/chilly (accent: US; accessed 2026-11-18).
```

### B04-P23 — raw

```plain
Pronunciation source: Collins Dictionary — raw; raw in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/raw (accent: US; accessed 2026-11-18).
```

### B04-P24 — spicy

```plain
Pronunciation source: Collins Dictionary — spicy; spicy in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/spicy (accent: US; accessed 2026-11-18).
```

### B04-P25 — durable

```plain
Pronunciation source: Collins Dictionary — durable; durable in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/durable (accent: US; accessed 2026-11-18).
```

### B04-P26 — sturdy

```plain
Pronunciation source: Collins Dictionary — sturdy; sturdy in American English 1; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/sturdy (accent: US; accessed 2026-11-18).
```

### B04-P27 — rusty

```plain
Pronunciation source: Collins Dictionary — rusty; rusty in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/rusty (accent: US; accessed 2026-11-18).
```

### B04-P28 — hidden

```plain
Pronunciation source: Collins Dictionary — hidden; hidden in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/hidden (accent: US; accessed 2026-11-18).
```

### B04-P29 — golden

```plain
Pronunciation source: Collins Dictionary — golden; golden in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/golden (accent: US; accessed 2026-11-18).
```

### B04-P30 — tough

```plain
Pronunciation source: Collins Dictionary — tough; tough in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/tough (accent: US; accessed 2026-11-18).
```

## 听审与暂缓边界

- H04-01: [tear-general](../entries/tear-general.yml)当前是眼泪义，拟写 US /tɪr/，不是撕裂义 /tɛr/。原生成器按 Word 朗读，单独词形不足以保证自动选中了当前词义。现有录音未试听，不判定录音错误，保留后续听审待办。
- H04-02: [bass-food](../entries/bass-food.yml)当前是鲈鱼义，拟写 US /bæs/，不是低音义 /beɪs/。现有录音同样未试听，不声称匹配或不匹配；如听审发现问题，再独立报告并取得替换音频的授权。
- 两项听审待办不是新增内容提案，也不阻止已查证的音标在批准后写入。它们不授权新建词条、增加辨音助记、关系、先修或重生成音频。
- 前批暂缓的 [part-of-speech-language](../entries/part-of-speech-language.yml)和[market-value-finance](../entries/market-value-finance.yml)保持原状，沿用[第二批暂缓记录](EN_PRONUNCIATION_ENRICHMENT_BATCH_02_PLAN.md#暂缓项目)。本批没有重新查证这两个整体短语，不自行拼接组件读音。
- 保留前三批全部音标、Notes 来源、关联、学习元数据及全部录音。同页同形词或词族关系不构成批量覆盖和新增关联的授权。

## 批准后实施与验证

1. 重新读取获批计划、当前规范、源码和 Git 状态，按稳定 ID 核对目标及规划以来的变化。若目标已有音标或来源，先协调差异，不覆盖并发编辑；只实施明确获批的编号。
2. 在仓库外快照所有词条原文和解析值、当前 TSV 及现有和归档录音哈希。逐项匹配编号、Word、词性、Meaning、分区、口音和 URL，再只写入批准的 pronunciation 与 Notes 来源行。
3. 保留稳定 ID、词义、词性、语言、Context、Example、旧 Notes、其他辅助字段、关系、学习元数据、标签和音频引用。只在缺少 Notes 的获批条目中创建该字段。不新增词条，不改变 schema、CLI、实现、模板或依赖。
4. 不生成、替换或复制录音。特别是 tear 与 bass，仅报告本轮未听审及后续待办，不从 IPA 自动改语音、重录或改写字段来影响 TTS。
5. 通过现有命令检查并生成导出，不手改 dist。

```bash
go test ./...
go vet ./...
go run scripts/validate/main.go
go run scripts/build/main.go
git diff --check
```

6. 在源不变时再构建一次并比较字节。核对实际基线稳定 ID 集、行数和 17 列 TSV，逐项检查完整 US 标签与 Notes 来源行。只允许获批行的 Pronunciation、Notes 变化，Audio、StudyOrder、RelatedWords、标签及其他列不变。集合若已变化，先核对新基线，不强制历史计数。
7. 核对全部源变更均符合批准字段范围，原有及归档 MP3 哈希不变，无关工作得到保留。另查未跟踪文本格式、本地链接与锚点，更新本计划批准、实施和实际检查结果，不改写前三批历史记录。
8. 不暂存、提交、推送、复制媒体或修改 Anki 集合。分别报告静态导出、人工听审及真实 Anki 导入、显示、播放，不声称执行未做过的验证。

## 规划检查

- 按当前 1,320 个稳定 ID 盘点，英语已有 367 项音标、缺失 710 项。30 个提案 ID 均存在且没有 pronunciation，与前三批已实施目标不重叠；当前词性和 Meaning 与源数据一致，分为 15 个名词、15 个形容词。
- 30 项具体值和完整来源行按编号、Word、URL、分区及 US 标签对应，最终证据均为已读正文。wheat 使用实际成功读取的 US 页面，hidden 的空格处理明确记录；没有新增无来源推测。
- 两个前批暂缓项本轮未重新查证；两个听审待办只标记尚未完成的人工工作，不判定录音错误，不计作新增词条、音标提案或新增暂缓项。
- 本计划本地链接、锚点、表格列数、围栏闭合、行尾空白、制表符和最终换行检查通过，git diff --check 通过；未跟踪计划另做文本格式检查。
- 与规划前快照比较，所有既有项目文件字节不变，包括 1,320 个词条、1,320 个现有 MP3、43 个归档 MP3、前三批计划、既有规范改动、实现、依赖、模板和 dist。唯一新增文件为本计划。
- 本轮未运行 Go 测试、词库验证、构建、音频生成、解码、人工听审或 Anki 操作；以上命令仅为批准后的实施程序。

## 实施与验证记录

- 用户确认 "B04-P01–P30；implement it"，授权完整 pronunciation 值及对应 Notes 来源行。30 项均已写入，全部为 US；没有新增 Notes 字段，原 Notes 全文保留，来源行各追加一次。
- 实施前重新核对稳定 ID、Word、词性、Meaning、缺失音标和来源文案，并在仓库外保存原词条、解析值、导出及既有文件哈希。规划以来没有影响本批目标的变化。
- 全库仍为 1,320 项，其中英语 1,077 项、日语 243 项。英语已有音标 397 项，缺失 680 项；日语保持不变。前批两个暂缓项目未改动。
- 与实施前快照比较，只有本批 30 个词条的 pronunciation 和 Notes 来源发生变化；其他 1,290 个词条字节不变。旧 Notes、稳定 ID、词义、例句、关系、先修、优先级、标签及音频引用均保留；既有未提交工作未改动。
- 1,320 个现有 MP3 和 43 个归档 MP3 的 SHA-256 全部保持不变。没有生成、替换或复制录音。
- go test ./...、go vet ./...、词库验证及 git diff --check 通过。验证覆盖 1,320 个词条；15 条既有同词警告保留，没有新增同词项。未跟踪计划另做格式、本地链接和锚点检查。
- 两次构建字节一致，均为 1,320 个数据行、17 列。相较实施前导出，仅本批 30 行的 Pronunciation、Notes 列变化；其他列及其他行不变，包括 Audio、StudyOrder、RelatedWords 和标签。SHA-256: `43c9aea5e3a63551fa69b7c0c341f0c0406c5fdc86ae2cc90e0c3a725aa75eb8`。
- 未进行录音解码、人工听审或真实 Anki 导入、显示、播放验证。H04-01、H04-02 仍为听审待办，不判定原录音正确或错误。没有暂存、提交、推送或修改 Anki 集合。
