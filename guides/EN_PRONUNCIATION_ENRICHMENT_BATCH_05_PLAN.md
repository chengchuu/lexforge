# 英语音标补全计划: 第五批建筑、地形与常见物品名词

状态: **B05-P01–B05-P30 的具体音标与来源文案已获批准并实施，静态验证完成。** 日期: 2026-11-18。人工听审及真实 Anki 验证尚未进行。

规划阶段仅创建本计划。用户随后明确确认 B05-P01–P30 并授权实施，本轮只写入批准的音标及 Notes 来源行，并使用现有命令重建导出。media、archives、其他既有文档、规范、实现、模板及依赖保持不变，不暂存、提交、推送或操作 Anki。

[第一批](EN_PRONUNCIATION_ENRICHMENT_BATCH_01_PLAN.md)、[第二批](EN_PRONUNCIATION_ENRICHMENT_BATCH_02_PLAN.md)、[第三批](EN_PRONUNCIATION_ENRICHMENT_BATCH_03_PLAN.md)和[第四批](EN_PRONUNCIATION_ENRICHMENT_BATCH_04_PLAN.md)保留为实施历史。第五批独立使用 B05-P01–B05-P30。

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

按当前稳定 ID 盘点全库音标字段；本批仅对下列 30 个缺失项目逐项查证。完整原词条解析值及既有项目文件哈希保存在仓库外，供保全核对。字段盘点不代表已经研究全库每个读音。

| 项目                          | 实际范围                                     |
|:------------------------------|:---------------------------------------------|
| 全库词条                      | 1,320                                        |
| 英语 / 日语                   | 1,077 / 243                                  |
| 英语已有 / 缺失 pronunciation | 397 / 680                                    |
| 日语已有 / 缺失 pronunciation | 243 / 0                                      |
| 本批具体提案                  | 30 个现有名词词条，全部为 add                |
| 口音                          | 30 项 US，0 项 UK                            |
| 新增暂缓项                    | 0                                            |
| 前批保留暂缓项                | 2，不计入本批 30 项                          |
| 允许修改范围                  | 批准后仅 pronunciation 与对应 Notes 来源追加 |
| 新增词条 / 录音 / 关联 / 先修 | 全部为 0                                     |

本批优先选择建筑、地形、常见物品和动物名称。虽然多来自游戏学习材料，当前目标是可用于日常阅读的普通名词，不补写游戏专名的推测读音，也不改动已批准的学习顺序或关系。

规划时英语缺失为 680 项。实施本批后，英语已有音标从 397 增至 427，缺失降至 650。前批两个暂缓项目包含在这 650 项中，不重复相加；其他 648 项未在本批逐项查证。上表保留实施前基线。

## 来源与转写规则

遵循 [AGENTS.md](../AGENTS.md#pronunciation)和[字段指南](ENTRY_SCHEMA.md#pronunciation)。优先美式；无法合理核实美式时可以采用明确标注的可靠英式。本批均取得美式正文证据，拟写完整 US 标签，不改已有无标签音标。

- 最终证据均为实际打开并阅读的 Collins Dictionary 美式正文，署名 Webster’s New World College Dictionary, 5th Digital Edition。每项 Notes 行包含所选分区、URL、口音和访问日 2026-11-18。未将搜索摘要作为已读正文。
- axe 页面美式分区交叉引用 ax；另已打开 ax 页面，其 WNW 的 ax in American English 1 明列 "or axe" 及斧头、武器义。因此 B05-P26 引用实际支持当前拼写与意义的 ax 正文，不更改当前 Word、ID 或文件名。其他 29 项来源网址与当前单词一致。
- 保留所选词典的 /i/、/u/、/r/、/ɛ/、/ɔ/ 和 /oʊ/ 等记法，不混拼同页其他词典、口音或不同转写系统。castle、forest、frog、arrow 和 factory 等存在多个正常美式变体，本批各选一个，不将其他来源变体判错。
- mountain 的 WNW 正文呈现为 "ˈmaʊnt ə n"，拟值仅合并版面空格为 /ˈmaʊntən/，没有增删或替换音素。不套用同页另一词典的无 schwa 记法。
- 对 gate、bridge、tower、lake、river、bee、jar、pot、bow、fan 标明所选 WNW 分区编号。同页不同词典的编号可能不同，不能只凭全页第一个同形词匹配。
- 当前 schema 允许带口音标签的纯文本 pronunciation；README 两面均条件显示该字段，不需要 schema、模板或导出实现变更。IPA 来源核实不代表现有合成录音已听审，不授权更换 voice、language 或音频文件。

## 具体音标

以下 30 项当前 pronunciation 均省略。当前 Word、词性和 Meaning 逐项引用源数据，不属于改写范围。每项批准内容为完整 pronunciation 值及同编号 Notes 来源行。

| 编号    | 稳定 ID / Word                                                        | 当前词性 | 当前 Meaning               | 拟写 pronunciation | 词典来源                                                                           |
|:--------|:----------------------------------------------------------------------|:---------|:---------------------------|:-------------------|:-----------------------------------------------------------------------------------|
| B05-P01 | [door-architecture](../entries/door-architecture.yml) / door          | noun     | 门。                       | `US /dɔr/`         | [Collins: door](https://www.collinsdictionary.com/dictionary/english/door)         |
| B05-P02 | [gate-architecture](../entries/gate-architecture.yml) / gate          | noun     | 大门                       | `US /ɡeɪt/`        | [Collins: gate](https://www.collinsdictionary.com/dictionary/english/gate)         |
| B05-P03 | [bridge-architecture](../entries/bridge-architecture.yml) / bridge    | noun     | 桥梁                       | `US /brɪdʒ/`       | [Collins: bridge](https://www.collinsdictionary.com/dictionary/english/bridge)     |
| B05-P04 | [cave-architecture](../entries/cave-architecture.yml) / cave          | noun     | 洞穴                       | `US /keɪv/`        | [Collins: cave](https://www.collinsdictionary.com/dictionary/english/cave)         |
| B05-P05 | [ceiling-architecture](../entries/ceiling-architecture.yml) / ceiling | noun     | 天花板                     | `US /ˈsilɪŋ/`      | [Collins: ceiling](https://www.collinsdictionary.com/dictionary/english/ceiling)   |
| B05-P06 | [castle-architecture](../entries/castle-architecture.yml) / castle    | noun     | 城堡                       | `US /ˈkæsəl/`      | [Collins: castle](https://www.collinsdictionary.com/dictionary/english/castle)     |
| B05-P07 | [tower-architecture](../entries/tower-architecture.yml) / tower       | noun     | 高塔                       | `US /ˈtaʊər/`      | [Collins: tower](https://www.collinsdictionary.com/dictionary/english/tower)       |
| B05-P08 | [village-architecture](../entries/village-architecture.yml) / village | noun     | 村庄                       | `US /ˈvɪlɪdʒ/`     | [Collins: village](https://www.collinsdictionary.com/dictionary/english/village)   |
| B05-P09 | [lake-terrain](../entries/lake-terrain.yml) / lake                    | noun     | 湖泊                       | `US /leɪk/`        | [Collins: lake](https://www.collinsdictionary.com/dictionary/english/lake)         |
| B05-P10 | [river-terrain](../entries/river-terrain.yml) / river                 | noun     | 河流                       | `US /ˈrɪvər/`      | [Collins: river](https://www.collinsdictionary.com/dictionary/english/river)       |
| B05-P11 | [island-terrain](../entries/island-terrain.yml) / island              | noun     | 岛屿                       | `US /ˈaɪlənd/`     | [Collins: island](https://www.collinsdictionary.com/dictionary/english/island)     |
| B05-P12 | [mountain-terrain](../entries/mountain-terrain.yml) / mountain        | noun     | 山；山岳。                 | `US /ˈmaʊntən/`    | [Collins: mountain](https://www.collinsdictionary.com/dictionary/english/mountain) |
| B05-P13 | [forest-terrain](../entries/forest-terrain.yml) / forest              | noun     | 森林                       | `US /ˈfɔrɪst/`     | [Collins: forest](https://www.collinsdictionary.com/dictionary/english/forest)     |
| B05-P14 | [cliff-terrain](../entries/cliff-terrain.yml) / cliff                 | noun     | 悬崖                       | `US /klɪf/`        | [Collins: cliff](https://www.collinsdictionary.com/dictionary/english/cliff)       |
| B05-P15 | [valley-terrain](../entries/valley-terrain.yml) / valley              | noun     | 山谷                       | `US /ˈvæli/`       | [Collins: valley](https://www.collinsdictionary.com/dictionary/english/valley)     |
| B05-P16 | [bee-general](../entries/bee-general.yml) / bee                       | noun     | 蜂；蜜蜂。                 | `US /bi/`          | [Collins: bee](https://www.collinsdictionary.com/dictionary/english/bee)           |
| B05-P17 | [frog-general](../entries/frog-general.yml) / frog                    | noun     | 青蛙。                     | `US /frɔɡ/`        | [Collins: frog](https://www.collinsdictionary.com/dictionary/english/frog)         |
| B05-P18 | [eagle-general](../entries/eagle-general.yml) / eagle                 | noun     | 鹰。                       | `US /ˈiɡəl/`       | [Collins: eagle](https://www.collinsdictionary.com/dictionary/english/eagle)       |
| B05-P19 | [balloon-engineering](../entries/balloon-engineering.yml) / balloon   | noun     | 气球。                     | `US /bəˈlun/`      | [Collins: balloon](https://www.collinsdictionary.com/dictionary/english/balloon)   |
| B05-P20 | [jar-general](../entries/jar-general.yml) / jar                       | noun     | 罐；广口瓶。               | `US /dʒɑr/`        | [Collins: jar](https://www.collinsdictionary.com/dictionary/english/jar)           |
| B05-P21 | [pot-general](../entries/pot-general.yml) / pot                       | noun     | 锅；罐。                   | `US /pɑt/`         | [Collins: pot](https://www.collinsdictionary.com/dictionary/english/pot)           |
| B05-P22 | [map-ui](../entries/map-ui.yml) / map                                 | noun     | 地图。                     | `US /mæp/`         | [Collins: map](https://www.collinsdictionary.com/dictionary/english/map)           |
| B05-P23 | [menu-ui](../entries/menu-ui.yml) / menu                              | noun     | 菜单；可供选择的项目列表。 | `US /ˈmɛnju/`      | [Collins: menu](https://www.collinsdictionary.com/dictionary/english/menu)         |
| B05-P24 | [arrow-weapon](../entries/arrow-weapon.yml) / arrow                   | noun     | 箭                         | `US /ˈæroʊ/`       | [Collins: arrow](https://www.collinsdictionary.com/dictionary/english/arrow)       |
| B05-P25 | [bow-weapon](../entries/bow-weapon.yml) / bow                         | noun     | 弓                         | `US /boʊ/`         | [Collins: bow](https://www.collinsdictionary.com/dictionary/english/bow)           |
| B05-P26 | [axe-weapon](../entries/axe-weapon.yml) / axe                         | noun     | 斧头                       | `US /æks/`         | [Collins: ax (also axe)](https://www.collinsdictionary.com/dictionary/english/ax)  |
| B05-P27 | [shield-weapon](../entries/shield-weapon.yml) / shield                | noun     | 盾牌                       | `US /ʃild/`        | [Collins: shield](https://www.collinsdictionary.com/dictionary/english/shield)     |
| B05-P28 | [sword-weapon](../entries/sword-weapon.yml) / sword                   | noun     | 剑                         | `US /sɔrd/`        | [Collins: sword](https://www.collinsdictionary.com/dictionary/english/sword)       |
| B05-P29 | [fan-engineering](../entries/fan-engineering.yml) / fan               | noun     | 风扇；扇子。               | `US /fæn/`         | [Collins: fan](https://www.collinsdictionary.com/dictionary/english/fan)           |
| B05-P30 | [factory-architecture](../entries/factory-architecture.yml) / factory | noun     | 工厂。                     | `US /ˈfæktəri/`    | [Collins: factory](https://www.collinsdictionary.com/dictionary/english/factory)   |

## 逐项核对

下表是审阅依据，不授权额外写入助记、词源、关联或其他 Notes 说明。

| 编号    | 核对结果与限制                                                                                                                            |
|:--------|:------------------------------------------------------------------------------------------------------------------------------------------|
| B05-P01 | 门的名词义，不改成 doorway 等其他词形。                                                                                                   |
| B05-P02 | 大门义对应 WNW gate in American English 1；不取浇铸通道、方言道路义或 -gate 构词成分。                                                    |
| B05-P03 | 桥梁义对应 WNW bridge in American English 1；保留本项建筑意义，不改软件 bridge-tech，也不取桥牌义。                                       |
| B05-P04 | 洞穴义名词原形，不从 cave in 短语或动词屈折形式取值。                                                                                     |
| B05-P05 | 天花板义名词，匹配室内顶部表面，不以价格上限义改写词义。保留来源 /i/。                                                                    |
| B05-P06 | 城堡义名词，取来源首个美式变体 /ˈkæsəl/；所列 /ˈkɑsəl/ 也不判错。不选作品专名 Castle, The。                                               |
| B05-P07 | 高塔义对应 WNW tower in American English 1，不用 tower 2 表示拖曳者的 /ˈtoʊər/。                                                          |
| B05-P08 | 村庄义普通名词，不取 Village (the) 专名分区。                                                                                             |
| B05-P09 | 湖泊义对应 WNW lake in American English 1；不取颜料 lake 2 或 Lake 人名分区。                                                             |
| B05-P10 | 河流义对应 WNW river in American English 1，不用 river 2 表示劈裂者的 /ˈraɪvər/。                                                         |
| B05-P11 | 岛屿义名词，依据完整来源读音，不按拼写补入 /s/。                                                                                          |
| B05-P12 | 山岳义名词，WNW 正文为 ˈmaʊnt ə n；仅合并版面空格得到 /ˈmaʊntən/，不添加、替换或删除音素。                                                |
| B05-P13 | 森林义名词，选 WNW /ˈfɔrɪst/；其 /ˈfɑrɪst/ 也是正常美式变体，不将两者拼合。                                                               |
| B05-P14 | 悬崖义小写普通名词，不以 Cliff 人名分区为发音依据。                                                                                       |
| B05-P15 | 山谷义名词单数，保留来源 /i/，不附加 valleys 的复数词尾。                                                                                 |
| B05-P16 | 蜜蜂义对应 WNW bee in American English 1；不取拼字比赛、船用构件或字母名称分区。                                                          |
| B05-P17 | 青蛙义名词，选 WNW 首个美式变体 /frɔɡ/；其 /frɑɡ/ 也不判错。不采用衣扣、铁路或马蹄部位的释义。                                            |
| B05-P18 | 鹰的鸟类义名词原形，不改高尔夫术语、货币或徽记等学习目标。                                                                                |
| B05-P19 | 气球义名词，重音在后一音节，保留来源 /u/；游戏装置说明和原录音均不改。                                                                    |
| B05-P20 | 罐或广口瓶义对应 WNW jar in American English 2；同页其他词典分区编号不同，不误取震动的动词义。                                            |
| B05-P21 | 锅或罐义对应 WNW pot in American English 1；不取同形俚语、缩略词或其他分区。                                                              |
| B05-P22 | 地图义名词，不取 Map 人名、MAP 缩写或编程映射的不同学习目标。                                                                             |
| B05-P23 | 菜单及屏幕可选项列表义；选 WNW 通常美式 /ˈmɛnju/，不选其标为偶见的 /ˈmeɪnju/。                                                            |
| B05-P24 | 箭的武器义名词，选 WNW 首个美式变体 /ˈæroʊ/；其 /ˈɛroʊ/ 也不判错。不以 Arrow 人名分区为来源。                                             |
| B05-P25 | 弓义对应 WNW bow in American English 2，不是鞠躬或船首义的 /baʊ/。现有录音未听审，不据此判断正误或自动重录。                              |
| B05-P26 | WNW ax in American English 1 明列 or axe，涵盖斧头工具及武器义。来源 URL 指向 ax，保留当前 axe 拼写和稳定 ID，不写复数 axes 的 /ˈæksɪz/。 |
| B05-P27 | 盾牌义名词，保留 WNW /i/，不使用动词屈折词尾。                                                                                            |
| B05-P28 | 剑的名词义，按完整词典读音，不因拼写自行添加 /w/。                                                                                        |
| B05-P29 | 风扇或扇子义对应 WNW fan in American English 1，不取爱好者 fan 2 或 Fan 专名。游戏装置说明保留。                                          |
| B05-P30 | 工厂义名词，选来源完整美式 /ˈfæktəri/；其标为常见的 /ˈfæktri/ 缩读也不判错。不改为软件工厂模式或历史商站义。                              |

## Notes 来源追加文案

实施时保留完整原 Notes，在末尾追加对应来源行一次；原字段不存在时才创建。若已经有同一发音来源，先核对并发变化，不重复追加。若相同 URL 用于其他解释，检查是否已包含本次分区、口音和发音来源说明，保留原有出处。以下围栏不写入 YAML 值。

### B05-P01 — door

```plain
Pronunciation source: Collins Dictionary — door; door in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/door (accent: US; accessed 2026-11-18).
```

### B05-P02 — gate

```plain
Pronunciation source: Collins Dictionary — gate; gate in American English 1; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/gate (accent: US; accessed 2026-11-18).
```

### B05-P03 — bridge

```plain
Pronunciation source: Collins Dictionary — bridge; bridge in American English 1; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/bridge (accent: US; accessed 2026-11-18).
```

### B05-P04 — cave

```plain
Pronunciation source: Collins Dictionary — cave; cave in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/cave (accent: US; accessed 2026-11-18).
```

### B05-P05 — ceiling

```plain
Pronunciation source: Collins Dictionary — ceiling; ceiling in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/ceiling (accent: US; accessed 2026-11-18).
```

### B05-P06 — castle

```plain
Pronunciation source: Collins Dictionary — castle; castle in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/castle (accent: US; accessed 2026-11-18).
```

### B05-P07 — tower

```plain
Pronunciation source: Collins Dictionary — tower; tower in American English 1; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/tower (accent: US; accessed 2026-11-18).
```

### B05-P08 — village

```plain
Pronunciation source: Collins Dictionary — village; village in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/village (accent: US; accessed 2026-11-18).
```

### B05-P09 — lake

```plain
Pronunciation source: Collins Dictionary — lake; lake in American English 1; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/lake (accent: US; accessed 2026-11-18).
```

### B05-P10 — river

```plain
Pronunciation source: Collins Dictionary — river; river in American English 1; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/river (accent: US; accessed 2026-11-18).
```

### B05-P11 — island

```plain
Pronunciation source: Collins Dictionary — island; island in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/island (accent: US; accessed 2026-11-18).
```

### B05-P12 — mountain

```plain
Pronunciation source: Collins Dictionary — mountain; mountain in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/mountain (accent: US; accessed 2026-11-18).
```

### B05-P13 — forest

```plain
Pronunciation source: Collins Dictionary — forest; forest in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/forest (accent: US; accessed 2026-11-18).
```

### B05-P14 — cliff

```plain
Pronunciation source: Collins Dictionary — cliff; cliff in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/cliff (accent: US; accessed 2026-11-18).
```

### B05-P15 — valley

```plain
Pronunciation source: Collins Dictionary — valley; valley in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/valley (accent: US; accessed 2026-11-18).
```

### B05-P16 — bee

```plain
Pronunciation source: Collins Dictionary — bee; bee in American English 1; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/bee (accent: US; accessed 2026-11-18).
```

### B05-P17 — frog

```plain
Pronunciation source: Collins Dictionary — frog; frog in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/frog (accent: US; accessed 2026-11-18).
```

### B05-P18 — eagle

```plain
Pronunciation source: Collins Dictionary — eagle; eagle in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/eagle (accent: US; accessed 2026-11-18).
```

### B05-P19 — balloon

```plain
Pronunciation source: Collins Dictionary — balloon; balloon in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/balloon (accent: US; accessed 2026-11-18).
```

### B05-P20 — jar

```plain
Pronunciation source: Collins Dictionary — jar; jar in American English 2; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/jar (accent: US; accessed 2026-11-18).
```

### B05-P21 — pot

```plain
Pronunciation source: Collins Dictionary — pot; pot in American English 1; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/pot (accent: US; accessed 2026-11-18).
```

### B05-P22 — map

```plain
Pronunciation source: Collins Dictionary — map; map in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/map (accent: US; accessed 2026-11-18).
```

### B05-P23 — menu

```plain
Pronunciation source: Collins Dictionary — menu; menu in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/menu (accent: US; accessed 2026-11-18).
```

### B05-P24 — arrow

```plain
Pronunciation source: Collins Dictionary — arrow; arrow in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/arrow (accent: US; accessed 2026-11-18).
```

### B05-P25 — bow

```plain
Pronunciation source: Collins Dictionary — bow; bow in American English 2; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/bow (accent: US; accessed 2026-11-18).
```

### B05-P26 — axe

```plain
Pronunciation source: Collins Dictionary — axe; ax in American English 1 (also axe); Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/ax (accent: US; accessed 2026-11-18).
```

### B05-P27 — shield

```plain
Pronunciation source: Collins Dictionary — shield; shield in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/shield (accent: US; accessed 2026-11-18).
```

### B05-P28 — sword

```plain
Pronunciation source: Collins Dictionary — sword; sword in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/sword (accent: US; accessed 2026-11-18).
```

### B05-P29 — fan

```plain
Pronunciation source: Collins Dictionary — fan; fan in American English 1; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/fan (accent: US; accessed 2026-11-18).
```

### B05-P30 — factory

```plain
Pronunciation source: Collins Dictionary — factory; factory in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/factory (accent: US; accessed 2026-11-18).
```

## 听审与暂缓边界

- H05-01: [bow-weapon](../entries/bow-weapon.yml)当前为弓义，拟写 US /boʊ/；鞠躬和船首义另读 /baʊ/。原录音由 Word 合成，当前未试听，不据此判定其读音错误。后续可独立听审；若确有问题，再取得替换录音的授权。
- tower 和 river 也已按当前名词意义排除其他读音分区。所有现有录音均未听审，不能由音标查证或文件存在推断录音正确。前批 H04-01、H04-02 保留在[第四批听审记录](EN_PRONUNCIATION_ENRICHMENT_BATCH_04_PLAN.md#听审与暂缓边界)，本批不重写其结果。
- 前批 [part-of-speech-language](../entries/part-of-speech-language.yml)和[market-value-finance](../entries/market-value-finance.yml)保持原状，沿用[第二批暂缓记录](EN_PRONUNCIATION_ENRICHMENT_BATCH_02_PLAN.md#暂缓项目)。本轮没有重新研究这两个整体短语，不从组件读音拼接音标，不计作本批新增暂缓项。
- H05-01 仅为人工听审待办，不是音标暂缓项，不阻止用户批准后写入已核实的 IPA。与其他词的关联、辨音助记、新词条、先修和录音替换均不在批准范围。

## 批准后实施与验证

1. 重新读取获批计划、当前规范、实现及 Git 状态，按稳定 ID 核对规划以来的变更。逐项确认当前 Word、词性、Meaning、缺失音标、分区、URL 和口音；目标如有并发修改，先协调，不覆盖。
2. 在仓库外保存全部词条原文和解析值、当前 TSV、现有及归档 MP3 哈希。只写入明确批准的 pronunciation 值及 Notes 来源追加行，旧 Notes 全文保留，来源不重复。
3. 保留语言、稳定 ID、Word、Meaning、词性、例句、其他辅助字段、关系、先修、学习优先级、标签及音频引用。不新增词条、录音，不改 schema、CLI、模板、实现或依赖。
4. 使用现有命令验证并重建导出，不手工编辑 dist。

```bash
go test ./...
go vet ./...
go run scripts/validate/main.go
go run scripts/build/main.go
git diff --check
```

5. 在源不变时再构建一次，比较字节。核对实施时真实稳定 ID 集、行数和 17 列 TSV；只允许获批行的 Pronunciation、Notes 列变化。Audio、StudyOrder、RelatedWords、标签及其余列保持不变，不强制已经过时的历史计数。
6. 对源数据执行字段允许范围检查，其他词条保留原字节，全部既有及归档录音哈希保持不变。检查未跟踪文本格式、本地链接和锚点，更新本计划批准、实施与检查状态，保留前批历史。
7. 分别报告静态和导出检查、解码、人工听审、真实 Anki 导入、显示、播放。仅补 IPA 不生成或替换音频，不假称已完成未执行的检查。
8. 不暂存、提交、推送、复制媒体进 Anki 或修改 Anki 集合。保留无关工作，遇到部分失败如实报告。

## 规划检查

- 按当前 1,320 个稳定 ID 核对，英语已有音标 397 项、缺失 680 项；日语 243 项均有读音。本批 30 个 ID 均存在、没有 pronunciation，Word、词性和 Meaning 与当前源一致，与前四批已实施目标不重叠。
- 30 个完整 US 值与同编号 Notes 文案、Word、实际来源 URL 及分区一一对应。axe 的 ax 变体来源与 mountain 的版面空格处理均已明确记录。未新增推测音标或把同页其他转写系统拼入拟值。
- 本计划本地链接、锚点、表格列数与对齐、围栏、行尾空白、制表符、最终换行和 git diff --check 检查通过；未跟踪计划另做文本检查。外部来源访问支持上文证据记录，不代表所有 URL 将持续可用。
- 与规划前快照比较，所有既有项目文件字节保持不变，包括 1,320 个词条、1,320 个现有 MP3、43 个归档 MP3、前批计划、现有规范和未提交工作、实现、模板、依赖及 dist。唯一新增文件为本计划。
- 本轮未运行 Go 测试、词库验证、构建、音频生成、解码、人工听审或 Anki 操作。上述命令为批准后实施程序，不是本轮已执行检查。

## 实施与验证记录

- 用户确认 "B05-P01–P30；implement it"，授权完整 pronunciation 值及对应 Notes 来源行。30 项均已写入，全部为 US；各项原 Notes 全文保留，来源行各追加一次，没有新增 Notes 字段。
- 实施前重新核对计划、规范、实现、Git 状态、稳定 ID、Word、词性、Meaning 和来源文案。与规划快照比较，既有项目文件均未变化；在仓库外另存全部原词条、解析值、导出和既有文件哈希。
- 全库仍为 1,320 项，其中英语 1,077 项、日语 243 项。英语已有音标 427 项、缺失 650 项；日语读音和两个前批暂缓项目保持不变。
- 与实施前快照比较，仅本批 30 个词条的 pronunciation 和 Notes 来源变化；其余 1,290 个词条字节不变。稳定 ID、Word、词义、词性、旧 Notes、例句、记忆辅助、关系、先修、学习优先级、标签及音频引用均保留，既有未提交工作未改动。
- 1,320 个现有 MP3 和 43 个归档 MP3 的 SHA-256 全部不变，没有生成、替换或复制录音。
- go test ./...、go vet ./...、词库验证和 git diff --check 通过。验证覆盖 1,320 个词条；15 条既有同词警告保留，没有新增同词项。未跟踪计划另做格式、本地链接和锚点检查。
- 两次构建字节一致，均为 1,320 个数据行、17 列。与实施前导出相比，仅本批 30 行的 Pronunciation、Notes 列变化；其他行和所有其他列保持不变，包括 Audio、StudyOrder、RelatedWords 和标签。SHA-256: `faf790e417b2aebc8c88e590b83b4e8f3e99702e61b61f716b356d0560cd77e5`。
- 未进行音频解码、人工听审或真实 Anki 导入、显示、播放验证。H05-01 和前批听审待办仍未完成，不判定现有录音正确或错误。没有暂存、提交、推送或修改 Anki 集合。
