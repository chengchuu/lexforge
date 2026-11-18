# 王国之泪词条音标补全计划

状态: **ZP001–ZP259 的具体音标与来源文案已获批准并实施，静态验证完成；ZD001–ZD202 仍待补足证据。** 日期: 2026-11-18。人工听审及真实 Anki 验证尚未进行。

本计划承接英语音标补全任务，按塞尔达类别推进，不再以 30 项作为批次上限。规划阶段仅创建本计划；用户随后明确确认 ZP001–ZP259 并授权实施，本轮已写入批准的 pronunciation 和对应发音来源，并通过现有命令重建导出。词义、例句、助记、关联、学习顺序及音频保持不变。

## 导航

- [范围与基线](#范围与基线)
- [来源与转写规则](#来源与转写规则)
- [按类别审阅具体音标](#按类别审阅具体音标)
- [逐项来源追加文案](#逐项来源追加文案)
- [待核实项目](#待核实项目)
- [已有音标保留清单](#已有音标保留清单)
- [听审与确认边界](#听审与确认边界)
- [批准后的实施与验证](#批准后的实施与验证)
- [规划检查](#规划检查)
- [实施与验证记录](#实施与验证记录)

## 范围与基线

当前词库中塞尔达主题均来自 The Legend of Zelda: Tears of the Kingdom，以实际 `zelda-totk` 标签确定范围。另检索全库 Word、Meaning、Context、Example 和 Notes 中的 Zelda、Tears of the Kingdom、王国之泪及塞尔达字样，没有发现需要补纳的未标记词条。按稳定 ID 识别条目，不用文件名词干代替 ID。

下表保留实施前源文件的基线。全类别盘点指每个范围内 ID 都有保留、具体提案或待核实决定；不表示每个专名已完成官方语音研究，也不表示已有音标经过重新审定。实施后计数见末尾验证记录。

| 项目                    | 数量与范围                                                    |
|:------------------------|:--------------------------------------------------------------|
| 全库词条                | 1,320，英语 1,077，日语 243                                   |
| 全库英语已有 / 缺失音标 | 427 / 650                                                     |
| 塞尔达范围              | 553 项，全部 language: en                                     |
| 范围内已有音标，保留    | 92 项                                                         |
| 范围内缺失音标          | 461 项                                                        |
| 缺失项内容分组          | 285 项普通词、屈折形式或短语；176 项游戏名称或造词            |
| 本轮具体提案            | 259 项: 256 项普通词及构式，另有 3 项同形普通名词用于游戏名称 |
| 提案口音                | 258 项美式，1 项英式                                          |
| 最终发音来源网址        | 共 257 个；重复网址对应不同 ID 或词性，不重复计算来源页面     |
| 待核实                  | 202 项: 29 项普通词形、短语或游戏用法，173 项游戏名称整体     |
| 范围外缺失音标          | 189 项，本轮不研究或修改                                      |
| 规划阶段修改            | 仅本计划；不运行构建，不修改词条或录音                        |

全部 259 项已获批准并实施。范围内已有音标增至 351 项，仍缺 202 项；全库英语已有音标增至 686 项，仍缺 391 项，其中包括范围外 189 项。259 是本轮证据支持的实际数量，不是配额；不把暂缓项写成已经完成。

保留[第一批](EN_PRONUNCIATION_ENRICHMENT_BATCH_01_PLAN.md)、[第二批](EN_PRONUNCIATION_ENRICHMENT_BATCH_02_PLAN.md)、[第三批](EN_PRONUNCIATION_ENRICHMENT_BATCH_03_PLAN.md)、[第四批](EN_PRONUNCIATION_ENRICHMENT_BATCH_04_PLAN.md)和[第五批](EN_PRONUNCIATION_ENRICHMENT_BATCH_05_PLAN.md)的历史批准及验证记录。本轮编号独立为 ZP，不重写过去的批准文本或计数。

## 来源与转写规则

遵循 [AGENTS.md](../AGENTS.md#pronunciation)和[字段指南](ENTRY_SCHEMA.md#pronunciation)的最新约定。美式音标写 `/…/`，不加 `US `；英式保留 `UK /…/`。Notes 仍明确记录 accent、来源分区和访问日期。口音标签只描述文字，不改变现有录音、language、TTSLanguage 或语音设置。

- 主要来源是实际打开并阅读的 Collins Dictionary 美式正文。通常使用 Webster’s New World College Dictionary, 5th Digital Edition 分区；one-handed、spaulder 和 ruthless 明确改用同页署名 Penguin Random House LLC 和 HarperCollins Publishers Ltd 的美式参考内容，来源行不冒称 WNW。
- hideout 的 Collins 美式正文采用 ai/au 记法；另查 Oxford Advanced Learner's Dictionary 的通用页面可读，但其美式专页复查返回 403。为不自行将字典记法改写成另一体系，本项采用已读 Collins 英式分区完整 IPA，标为 UK。不是断言没有美式发音，也不替换美式录音。
- geoglyph 的 Collins 页面不可访问，采用已读 Wiktionary 英语分区明确标为 General American 的完整 IPA。页面显示修订号 92260556，来源行使用该修订的固定链接。Wiktionary 是社区维护词典；其词源部分引用 OED，不代表本轮读过 OED，且本轮不新增词源。其他在线转写器、搜索摘要或镜像未作为最终发音证据。
- 285 个普通目标对应 284 个不同词形，spring 有两个现有义项 ID。普通目标均尝试对应 Collins 页面；访问失败或读音、词形不匹配的项目另列。针对其中特定缺口还尝试 Cambridge 发音页，多数返回 403 或不可访问。搜索摘要显示音标但正文不能读取时，仍暂缓，不假称已检查正文。
- 保留所选来源的一套完整转写和正常变体，不将美式、英式或不同词典的符号拼合。若来源在音素之间带版面空格，仅去除空格，例如 fountain、guidance、settlement 和 trident；不添加音素，不从组件推导读音。多变体通常取同一来源一个完整形式，逐项例外在表中说明。
- 按当前词性和意义处理 desert、construct、wind、entrance、upgrade 等异读项目。当前为复数、分词或派生词时，不把原形 IPA 当作完整证据，不自行拼接词尾。
- Construct、Gloom、Paraglider 三项单独标注为普通英语名词读音提案。现有游戏意义及原文保持不变，Notes 来源行明确说明不是已核实的游戏配音转写；不把字典普通释义当作游戏机制。其他造词、人物、族群、地点及长名称不按拼写猜读，不从日语读音或组件拼接整体 IPA。
- 所有具体值都是新增，不替换现有音标。原 Notes 保留全文，新增发音来源行各一次。259 项中没有现有 Notes 已含相同最终来源 URL 或完整待追加行的情况。卡片内容是纯文本，不加入 HTML 或 Markdown。

## 按类别审阅具体音标

每行批准范围为完整 pronunciation 值及下一节同编号 Notes 行。Word、词性和 Meaning 来自当前源，不属于改写提案。分类只便于审阅，不修改 tags、learning_priority、prerequisites 或 StudyOrder；ZP 编号按稳定 ID 分配，类别中编号不一定连续。

### 建筑与设施

共 24 项。

| 编号  | 稳定 ID / Word                                                                 | 当前词性 | 当前 Meaning                   | 拟写 pronunciation | 来源与分区                                                                                        | 核对说明                                                                                                                                    |
|:------|:-------------------------------------------------------------------------------|:---------|:-------------------------------|:-------------------|:--------------------------------------------------------------------------------------------------|:--------------------------------------------------------------------------------------------------------------------------------------------|
| ZP007 | [altar-architecture](../entries/altar-architecture.yml) / altar                | noun     | 祭坛                           | `/ˈɔltər/`         | [altar in American English](https://www.collinsdictionary.com/dictionary/english/altar)           | 匹配当前词形与词义，保留所选来源完整变体。                                                                                                  |
| ZP015 | [armory-architecture](../entries/armory-architecture.yml) / armory             | noun     | 武器库                         | `/ˈɑrməri/`        | [armory in American English](https://www.collinsdictionary.com/dictionary/english/armory)         | 匹配当前词形与词义，保留所选来源完整变体。                                                                                                  |
| ZP030 | [cavern-architecture](../entries/cavern-architecture.yml) / cavern             | noun     | 大型洞穴                       | `/ˈkævərn/`        | [cavern in American English](https://www.collinsdictionary.com/dictionary/english/cavern)         | 匹配当前词形与词义，保留所选来源完整变体。                                                                                                  |
| ZP031 | [chamber-architecture](../entries/chamber-architecture.yml) / chamber          | noun     | 房间、密室                     | `/ˈtʃeɪmbər/`      | [chamber in American English](https://www.collinsdictionary.com/dictionary/english/chamber)       | 匹配当前词形与词义，保留所选来源完整变体。                                                                                                  |
| ZP034 | [citadel-architecture](../entries/citadel-architecture.yml) / citadel          | noun     | 城塞、堡垒                     | `/ˈsɪtəˌdɛl/`      | [citadel in American English](https://www.collinsdictionary.com/dictionary/english/citadel)       | 匹配当前词形与词义，保留所选来源完整变体。                                                                                                  |
| ZP045 | [corridor-architecture](../entries/corridor-architecture.yml) / corridor       | noun     | 走廊                           | `/ˈkɔrədər/`       | [corridor in American English](https://www.collinsdictionary.com/dictionary/english/corridor)     | 匹配当前词形与词义，保留所选来源完整变体。                                                                                                  |
| ZP077 | [entrance-architecture](../entries/entrance-architecture.yml) / entrance       | noun     | 入口                           | `/ˈɛntrəns/`       | [entrance in American English 1](https://www.collinsdictionary.com/dictionary/english/entrance)   | 匹配当前词形与词义，保留所选来源完整变体。                                                                                                  |
| ZP080 | [exit-architecture](../entries/exit-architecture.yml) / exit                   | noun     | 出口                           | `/ˈɛksɪt/`         | [exit in American English](https://www.collinsdictionary.com/dictionary/english/exit)             | 匹配当前词形与词义，保留所选来源完整变体。                                                                                                  |
| ZP085 | [forge-architecture](../entries/forge-architecture.yml) / forge                | noun     | 锻造设施                       | `/fɔrdʒ/`          | [forge in American English 1](https://www.collinsdictionary.com/dictionary/english/forge)         | 匹配当前词形与词义，保留所选来源完整变体。                                                                                                  |
| ZP086 | [fortress-architecture](../entries/fortress-architecture.yml) / fortress       | noun     | 要塞、堡垒                     | `/ˈfɔrtrɪs/`       | [fortress in American English](https://www.collinsdictionary.com/dictionary/english/fortress)     | 匹配当前词形与词义，保留所选来源完整变体。                                                                                                  |
| ZP087 | [fountain-architecture](../entries/fountain-architecture.yml) / fountain       | noun     | 喷泉、泉水设施                 | `/ˈfaʊntən/`       | [fountain in American English](https://www.collinsdictionary.com/dictionary/english/fountain)     | 匹配当前词形与词义，保留所选来源完整变体。                                                                                                  |
| ZP109 | [hideout-architecture](../entries/hideout-architecture.yml) / hideout          | noun     | 藏身处；秘密据点。             | `UK /ˈhaɪdˌaʊt/`   | [hideout in British English](https://www.collinsdictionary.com/dictionary/english/hideout)        | Collins 美式分区采用 ai/au 记法；Oxford 美式专页复查返回 403。为不自行转换来源记法，本项使用已读 Collins 英式分区完整 IPA，并保留 UK 标签。 |
| ZP119 | [inn-architecture](../entries/inn-architecture.yml) / inn                      | noun     | 旅馆                           | `/ɪn/`             | [inn in American English](https://www.collinsdictionary.com/dictionary/english/inn)               | 匹配当前词形与词义，保留所选来源完整变体。                                                                                                  |
| ZP139 | [mine-architecture](../entries/mine-architecture.yml) / mine                   | noun     | 矿场                           | `/maɪn/`           | [mine in American English 2](https://www.collinsdictionary.com/dictionary/english/mine)           | 取 WNW mine in American English 2 的矿场义，不把同音的所有格代词当作当前来源。                                                              |
| ZP150 | [passage-architecture](../entries/passage-architecture.yml) / passage          | noun     | 通道                           | `/ˈpæsɪdʒ/`        | [passage in American English](https://www.collinsdictionary.com/dictionary/english/passage)       | 匹配当前词形与词义，保留所选来源完整变体。                                                                                                  |
| ZP154 | [platform-architecture](../entries/platform-architecture.yml) / platform       | noun     | 平台                           | `/ˈplætˌfɔrm/`     | [platform in American English](https://www.collinsdictionary.com/dictionary/english/platform)     | 匹配当前词形与词义，保留所选来源完整变体。                                                                                                  |
| ZP172 | [refinery-architecture](../entries/refinery-architecture.yml) / refinery       | noun     | 精炼设施                       | `/rɪˈfaɪnəri/`     | [refinery in American English](https://www.collinsdictionary.com/dictionary/english/refinery)     | 匹配当前词形与词义，保留所选来源完整变体。                                                                                                  |
| ZP196 | [settlement-architecture](../entries/settlement-architecture.yml) / settlement | noun     | 定居点、聚落                   | `/ˈsɛtəlmənt/`     | [settlement in American English](https://www.collinsdictionary.com/dictionary/english/settlement) | 匹配当前词形与词义，保留所选来源完整变体。                                                                                                  |
| ZP198 | [shrine-architecture](../entries/shrine-architecture.yml) / shrine             | noun     | 神龛；圣祠；供奉或祭祀的场所。 | `/ʃraɪn/`          | [shrine in American English](https://www.collinsdictionary.com/dictionary/english/shrine)         | 匹配当前词形与词义，保留所选来源完整变体。                                                                                                  |
| ZP218 | [stable-architecture](../entries/stable-architecture.yml) / stable             | noun     | 马厩；用于安置马匹的建筑。     | `/ˈsteɪbəl/`       | [stable in American English 2](https://www.collinsdictionary.com/dictionary/english/stable)       | 取 WNW stable in American English 2 的马厩名词，不改为稳定的形容词。                                                                        |
| ZP219 | [staircase-architecture](../entries/staircase-architecture.yml) / staircase    | noun     | 楼梯                           | `/ˈstɛrˌkeɪs/`     | [staircase in American English](https://www.collinsdictionary.com/dictionary/english/staircase)   | 匹配当前词形与词义，保留所选来源完整变体。                                                                                                  |
| ZP222 | [statue-architecture](../entries/statue-architecture.yml) / statue             | noun     | 雕像                           | `/ˈstætʃu/`        | [statue in American English](https://www.collinsdictionary.com/dictionary/english/statue)         | 匹配当前词形与词义，保留所选来源完整变体。                                                                                                  |
| ZP232 | [temple-architecture](../entries/temple-architecture.yml) / temple             | noun     | 神殿、寺庙                     | `/ˈtɛmpəl/`        | [temple in American English 1](https://www.collinsdictionary.com/dictionary/english/temple)       | 匹配当前词形与词义，保留所选来源完整变体。                                                                                                  |
| ZP255 | [well-architecture](../entries/well-architecture.yml) / well                   | noun     | 水井                           | `/wɛl/`            | [well in American English 1](https://www.collinsdictionary.com/dictionary/english/well)           | 匹配当前词形与词义，保留所选来源完整变体。                                                                                                  |

### 地形与自然

共 19 项。

| 编号  | 稳定 ID / Word                                                          | 当前词性 | 当前 Meaning                   | 拟写 pronunciation | 来源与分区                                                                                          | 核对说明                                                                           |
|:------|:------------------------------------------------------------------------|:---------|:-------------------------------|:-------------------|:----------------------------------------------------------------------------------------------------|:-----------------------------------------------------------------------------------|
| ZP012 | [archipelago-terrain](../entries/archipelago-terrain.yml) / archipelago | noun     | 群岛                           | `/ˌɑrkəˈpɛləˌɡoʊ/` | [archipelago in American English](https://www.collinsdictionary.com/dictionary/english/archipelago) | 匹配当前词形与词义，保留所选来源完整变体。                                         |
| ZP028 | [canyon-terrain](../entries/canyon-terrain.yml) / canyon                | noun     | 峡谷                           | `/ˈkænjən/`        | [canyon in American English](https://www.collinsdictionary.com/dictionary/english/canyon)           | 匹配当前词形与词义，保留所选来源完整变体。                                         |
| ZP033 | [chasm-terrain](../entries/chasm-terrain.yml) / chasm                   | noun     | 深渊；深而宽的裂隙。           | `/ˈkæzəm/`         | [chasm in American English](https://www.collinsdictionary.com/dictionary/english/chasm)             | 匹配当前词形与词义，保留所选来源完整变体。                                         |
| ZP047 | [crater-terrain](../entries/crater-terrain.yml) / crater                | noun     | 火山口                         | `/ˈkreɪtər/`       | [crater in American English](https://www.collinsdictionary.com/dictionary/english/crater)           | 匹配当前词形与词义，保留所选来源完整变体。                                         |
| ZP059 | [desert-terrain](../entries/desert-terrain.yml) / desert                | noun     | 沙漠                           | `/ˈdɛzərt/`        | [desert in American English 2](https://www.collinsdictionary.com/dictionary/english/desert)         | 取 WNW desert in American English 2 的沙漠名词 /ˈdɛzərt/，不是遗弃动词 /dɪˈzɜrt/。 |
| ZP084 | [foothill-terrain](../entries/foothill-terrain.yml) / foothill          | noun     | 山麓                           | `/ˈfʊtˌhɪl/`       | [foothill in American English](https://www.collinsdictionary.com/dictionary/english/foothill)       | 匹配当前词形与词义，保留所选来源完整变体。                                         |
| ZP127 | [lava-terrain](../entries/lava-terrain.yml) / lava                      | noun     | 熔岩                           | `/ˈlɑvə/`          | [lava in American English](https://www.collinsdictionary.com/dictionary/english/lava)               | 匹配当前词形与词义，保留所选来源完整变体。                                         |
| ZP129 | [lightning-terrain](../entries/lightning-terrain.yml) / lightning       | noun     | 闪电。                         | `/ˈlaɪtnɪŋ/`       | [lightning in American English](https://www.collinsdictionary.com/dictionary/english/lightning)     | 匹配当前词形与词义，保留所选来源完整变体。                                         |
| ZP153 | [plateau-terrain](../entries/plateau-terrain.yml) / plateau             | noun     | 高原                           | `/plæˈtoʊ/`        | [plateau in American English](https://www.collinsdictionary.com/dictionary/english/plateau)         | 匹配当前词形与词义，保留所选来源完整变体。                                         |
| ZP176 | [ridge-terrain](../entries/ridge-terrain.yml) / ridge                   | noun     | 山脊                           | `/rɪdʒ/`           | [ridge in American English](https://www.collinsdictionary.com/dictionary/english/ridge)             | 匹配当前词形与词义，保留所选来源完整变体。                                         |
| ZP202 | [sky-terrain](../entries/sky-terrain.yml) / sky                         | noun     | 天空                           | `/skaɪ/`           | [sky in American English](https://www.collinsdictionary.com/dictionary/english/sky)                 | 匹配当前词形与词义，保留所选来源完整变体。                                         |
| ZP206 | [slope-terrain](../entries/slope-terrain.yml) / slope                   | noun     | 斜坡                           | `/sloʊp/`          | [slope in American English](https://www.collinsdictionary.com/dictionary/english/slope)             | 匹配当前词形与词义，保留所选来源完整变体。                                         |
| ZP215 | [spring-water](../entries/spring-water.yml) / spring                    | noun     | 泉；泉水；地下水自然涌出之处。 | `/sprɪŋ/`          | [spring in American English](https://www.collinsdictionary.com/dictionary/english/spring)           | WNW 名词义 18 为自然水源；与弹簧同形同音，现有两个 ID 的释义和关系保持不变。       |
| ZP226 | [summit-terrain](../entries/summit-terrain.yml) / summit                | noun     | 山顶                           | `/ˈsʌmɪt/`         | [summit in American English](https://www.collinsdictionary.com/dictionary/english/summit)           | 匹配当前词形与词义，保留所选来源完整变体。                                         |
| ZP228 | [surface-terrain](../entries/surface-terrain.yml) / surface             | noun     | 地表                           | `/ˈsɜrfɪs/`        | [surface in American English](https://www.collinsdictionary.com/dictionary/english/surface)         | 匹配当前词形与词义，保留所选来源完整变体。                                         |
| ZP230 | [swamp-terrain](../entries/swamp-terrain.yml) / swamp                   | noun     | 沼泽                           | `/swɑmp/`          | [swamp in American English](https://www.collinsdictionary.com/dictionary/english/swamp)             | 匹配当前词形与词义，保留所选来源完整变体。                                         |
| ZP238 | [tree-terrain](../entries/tree-terrain.yml) / tree                      | noun     | 树；木本植物。                 | `/tri/`            | [tree in American English](https://www.collinsdictionary.com/dictionary/english/tree)               | 匹配当前词形与词义，保留所选来源完整变体。                                         |
| ZP253 | [waterfall-terrain](../entries/waterfall-terrain.yml) / waterfall       | noun     | 瀑布                           | `/ˈwɔtərˌfɔl/`     | [waterfall in American English](https://www.collinsdictionary.com/dictionary/english/waterfall)     | 匹配当前词形与词义，保留所选来源完整变体。                                         |
| ZP258 | [wind-terrain](../entries/wind-terrain.yml) / wind                      | noun     | 风。                           | `/wɪnd/`           | [wind in American English 2](https://www.collinsdictionary.com/dictionary/english/wind)             | 取 WNW wind in American English 2 的风之名词 /wɪnd/，不是卷绕动词 /waɪnd/。        |

### 武器

共 13 项。

| 编号  | 稳定 ID / Word                                                     | 当前词性 | 当前 Meaning | 拟写 pronunciation | 来源与分区                                                                                        | 核对说明                                   |
|:------|:-------------------------------------------------------------------|:---------|:-------------|:-------------------|:--------------------------------------------------------------------------------------------------|:-------------------------------------------|
| ZP022 | [bomb-weapon](../entries/bomb-weapon.yml) / bomb                   | noun     | 炸弹。       | `/bɑm/`            | [bomb in American English](https://www.collinsdictionary.com/dictionary/english/bomb)             | 匹配当前词形与词义，保留所选来源完整变体。 |
| ZP023 | [boomerang-weapon](../entries/boomerang-weapon.yml) / boomerang    | noun     | 回旋镖       | `/ˈbumərˌæŋ/`      | [boomerang in American English](https://www.collinsdictionary.com/dictionary/english/boomerang)   | 匹配当前词形与词义，保留所选来源完整变体。 |
| ZP024 | [broadsword-weapon](../entries/broadsword-weapon.yml) / broadsword | noun     | 阔剑         | `/ˈbrɔdˌsɔrd/`     | [broadsword in American English](https://www.collinsdictionary.com/dictionary/english/broadsword) | 匹配当前词形与词义，保留所选来源完整变体。 |
| ZP036 | [claymore-weapon](../entries/claymore-weapon.yml) / claymore       | noun     | 双手大剑     | `/ˈkleɪˌmɔr/`      | [claymore in American English](https://www.collinsdictionary.com/dictionary/english/claymore)     | 匹配当前词形与词义，保留所选来源完整变体。 |
| ZP038 | [club-weapon](../entries/club-weapon.yml) / club                   | noun     | 棍棒         | `/klʌb/`           | [club in American English](https://www.collinsdictionary.com/dictionary/english/club)             | 匹配当前词形与词义，保留所选来源完整变体。 |
| ZP053 | [dagger-weapon](../entries/dagger-weapon.yml) / dagger             | noun     | 匕首         | `/ˈdæɡər/`         | [dagger in American English](https://www.collinsdictionary.com/dictionary/english/dagger)         | 匹配当前词形与词义，保留所选来源完整变体。 |
| ZP099 | [halberd-weapon](../entries/halberd-weapon.yml) / halberd          | noun     | 戟           | `/ˈhælbərd/`       | [halberd in American English](https://www.collinsdictionary.com/dictionary/english/halberd)       | 匹配当前词形与词义，保留所选来源完整变体。 |
| ZP100 | [hammer-weapon](../entries/hammer-weapon.yml) / hammer             | noun     | 锤子         | `/ˈhæmər/`         | [hammer in American English](https://www.collinsdictionary.com/dictionary/english/hammer)         | 匹配当前词形与词义，保留所选来源完整变体。 |
| ZP191 | [scimitar-weapon](../entries/scimitar-weapon.yml) / scimitar       | noun     | 弯刀         | `/ˈsɪmətər/`       | [scimitar in American English](https://www.collinsdictionary.com/dictionary/english/scimitar)     | 匹配当前词形与词义，保留所选来源完整变体。 |
| ZP211 | [spear-weapon](../entries/spear-weapon.yml) / spear                | noun     | 长矛         | `/spɪr/`           | [spear in American English](https://www.collinsdictionary.com/dictionary/english/spear)           | 匹配当前词形与词义，保留所选来源完整变体。 |
| ZP235 | [torch-weapon](../entries/torch-weapon.yml) / torch                | noun     | 火把         | `/tɔrtʃ/`          | [torch in American English](https://www.collinsdictionary.com/dictionary/english/torch)           | 匹配当前词形与词义，保留所选来源完整变体。 |
| ZP239 | [trident-weapon](../entries/trident-weapon.yml) / trident          | noun     | 三叉戟       | `/ˈtraɪdənt/`      | [trident in American English](https://www.collinsdictionary.com/dictionary/english/trident)       | 匹配当前词形与词义，保留所选来源完整变体。 |
| ZP254 | [weapon-weapon](../entries/weapon-weapon.yml) / weapon             | noun     | 武器。       | `/ˈwɛpən/`         | [weapon in American English](https://www.collinsdictionary.com/dictionary/english/weapon)         | 匹配当前词形与词义，保留所选来源完整变体。 |

### 衣物与防具

共 14 项。

| 编号  | 稳定 ID / Word                                                      | 当前词性 | 当前 Meaning     | 拟写 pronunciation | 来源与分区                                                                                      | 核对说明                                                                                    |
|:------|:--------------------------------------------------------------------|:---------|:-----------------|:-------------------|:------------------------------------------------------------------------------------------------|:--------------------------------------------------------------------------------------------|
| ZP017 | [bandanna-clothing](../entries/bandanna-clothing.yml) / bandanna    | noun     | 头巾             | `/bænˈdænə/`       | [bandanna in American English](https://www.collinsdictionary.com/dictionary/english/bandanna)   | 匹配当前词形与词义，保留所选来源完整变体。                                                  |
| ZP102 | [headband-clothing](../entries/headband-clothing.yml) / headband    | noun     | 头带             | `/ˈhɛdˌbænd/`      | [headband in American English](https://www.collinsdictionary.com/dictionary/english/headband)   | 匹配当前词形与词义，保留所选来源完整变体。                                                  |
| ZP103 | [headdress-clothing](../entries/headdress-clothing.yml) / headdress | noun     | 头饰             | `/ˈhɛdˌdrɛs/`      | [headdress in American English](https://www.collinsdictionary.com/dictionary/english/headdress) | 匹配当前词形与词义，保留所选来源完整变体。                                                  |
| ZP106 | [helm-clothing](../entries/helm-clothing.yml) / helm                | noun     | 头盔             | `/hɛlm/`           | [helm in American English 1](https://www.collinsdictionary.com/dictionary/english/helm)         | 匹配当前词形与词义，保留所选来源完整变体。                                                  |
| ZP107 | [helmet-clothing](../entries/helmet-clothing.yml) / helmet          | noun     | 头盔             | `/ˈhɛlmət/`        | [helmet in American English](https://www.collinsdictionary.com/dictionary/english/helmet)       | 匹配当前词形与词义，保留所选来源完整变体。                                                  |
| ZP112 | [hood-clothing](../entries/hood-clothing.yml) / hood                | noun     | 兜帽             | `/hud/`            | [hood in American English 1](https://www.collinsdictionary.com/dictionary/english/hood)         | 取 WNW hood in American English 1 的兜帽义，已另读该分区；不取 hoodlum 缩略词或 Hood 人名。 |
| ZP133 | [mask-clothing](../entries/mask-clothing.yml) / mask                | noun     | 面具             | `/mæsk/`           | [mask in American English](https://www.collinsdictionary.com/dictionary/english/mask)           | 匹配当前词形与词义，保留所选来源完整变体。                                                  |
| ZP179 | [robe-clothing](../entries/robe-clothing.yml) / robe                | noun     | 长袍             | `/roʊb/`           | [robe in American English](https://www.collinsdictionary.com/dictionary/english/robe)           | 匹配当前词形与词义，保留所选来源完整变体。                                                  |
| ZP197 | [shirt-clothing](../entries/shirt-clothing.yml) / shirt             | noun     | 衬衫；上衣。     | `/ʃɜrt/`           | [shirt in American English](https://www.collinsdictionary.com/dictionary/english/shirt)         | 匹配当前词形与词义，保留所选来源完整变体。                                                  |
| ZP204 | [sleeve-clothing](../entries/sleeve-clothing.yml) / sleeve          | noun     | 袖子。           | `/sliv/`           | [sleeve in American English](https://www.collinsdictionary.com/dictionary/english/sleeve)       | 匹配当前词形与词义，保留所选来源完整变体。                                                  |
| ZP210 | [spaulder-clothing](../entries/spaulder-clothing.yml) / spaulder    | noun     | 肩甲             | `/ˈspɔldər/`       | [spaulder in American English](https://www.collinsdictionary.com/dictionary/english/spaulder)   | 匹配当前词形与词义，保留所选来源完整变体。                                                  |
| ZP233 | [tights-clothing](../entries/tights-clothing.yml) / tights          | noun     | 紧身裤；连裤袜。 | `/taɪts/`          | [tights in American English](https://www.collinsdictionary.com/dictionary/english/tights)       | 匹配当前词形与词义，保留所选来源完整变体。                                                  |
| ZP240 | [trousers-clothing](../entries/trousers-clothing.yml) / trousers    | noun     | 长裤             | `/ˈtraʊzərz/`      | [trousers in American English](https://www.collinsdictionary.com/dictionary/english/trousers)   | 匹配当前词形与词义，保留所选来源完整变体。                                                  |
| ZP242 | [tunic-clothing](../entries/tunic-clothing.yml) / tunic             | noun     | 束腰上衣、短袍   | `/ˈtunɪk/`         | [tunic in American English](https://www.collinsdictionary.com/dictionary/english/tunic)         | 匹配当前词形与词义，保留所选来源完整变体。                                                  |

### 装置与工程

共 17 项。

| 编号  | 稳定 ID / Word                                                               | 当前词性  | 当前 Meaning                           | 拟写 pronunciation | 来源与分区                                                                                        | 核对说明                                                                  |
|:------|:-----------------------------------------------------------------------------|:----------|:---------------------------------------|:-------------------|:--------------------------------------------------------------------------------------------------|:--------------------------------------------------------------------------|
| ZP018 | [battery-engineering](../entries/battery-engineering.yml) / battery          | noun      | 电池                                   | `/ˈbætəri/`        | [battery in American English](https://www.collinsdictionary.com/dictionary/english/battery)       | 匹配当前词形与词义，保留所选来源完整变体。                                |
| ZP019 | [beam-engineering](../entries/beam-engineering.yml) / beam                   | noun      | 光束；此处指集中的一束光。             | `/bim/`            | [beam in American English](https://www.collinsdictionary.com/dictionary/english/beam)             | 匹配当前词形与词义，保留所选来源完整变体。                                |
| ZP027 | [cannon-engineering](../entries/cannon-engineering.yml) / cannon             | noun      | 大炮                                   | `/ˈkænən/`         | [cannon in American English](https://www.collinsdictionary.com/dictionary/english/cannon)         | 匹配当前词形与词义，保留所选来源完整变体。                                |
| ZP029 | [cart-engineering](../entries/cart-engineering.yml) / cart                   | noun      | 小车；运货车。                         | `/kɑrt/`           | [cart in American English](https://www.collinsdictionary.com/dictionary/english/cart)             | 匹配当前词形与词义，保留所选来源完整变体。                                |
| ZP062 | [device-engineering](../entries/device-engineering.yml) / device             | noun      | 装置；设备；为某种用途设计的器具。     | `/dɪˈvaɪs/`        | [device in American English](https://www.collinsdictionary.com/dictionary/english/device)         | 匹配当前词形与词义，保留所选来源完整变体。                                |
| ZP064 | [dispenser-engineering](../entries/dispenser-engineering.yml) / dispenser    | noun      | 分配器；按需放出物品的装置。           | `/dɪˈspɛnsər/`     | [dispenser in American English](https://www.collinsdictionary.com/dictionary/english/dispenser)   | 匹配当前词形与词义，保留所选来源完整变体。                                |
| ZP072 | [emitter-engineering](../entries/emitter-engineering.yml) / emitter          | noun      | 发射器、发出装置                       | `/iˈmɪtər/`        | [emitter in American English](https://www.collinsdictionary.com/dictionary/english/emitter)       | 匹配当前词形与词义，保留所选来源完整变体。                                |
| ZP110 | [homing-engineering](../entries/homing-engineering.yml) / homing             | adjective | 自动导引的；能自动追踪目标的。         | `/ˈhoʊmɪŋ/`        | [homing in American English](https://www.collinsdictionary.com/dictionary/english/homing)         | 匹配当前词形与词义，保留所选来源完整变体。                                |
| ZP115 | [hydrant-engineering](../entries/hydrant-engineering.yml) / hydrant          | noun      | 消防栓；连接供水管道的取水装置。       | `/ˈhaɪdrənt/`      | [hydrant in American English](https://www.collinsdictionary.com/dictionary/english/hydrant)       | 匹配当前词形与词义，保留所选来源完整变体。                                |
| ZP128 | [light-engineering](../entries/light-engineering.yml) / light                | noun      | 光；光源；照明装置。                   | `/laɪt/`           | [light in American English 1](https://www.collinsdictionary.com/dictionary/english/light)         | 匹配当前词形与词义，保留所选来源完整变体。                                |
| ZP180 | [rocket-engineering](../entries/rocket-engineering.yml) / rocket             | noun      | 火箭。                                 | `/ˈrɑkɪt/`         | [rocket in American English 1](https://www.collinsdictionary.com/dictionary/english/rocket)       | 取 WNW rocket in American English 1，已另读火箭装置定义；不取植物名分区。 |
| ZP203 | [sled-engineering](../entries/sled-engineering.yml) / sled                   | noun      | 雪橇；可在雪地等表面滑行的载具。       | `/slɛd/`           | [sled in American English](https://www.collinsdictionary.com/dictionary/english/sled)             | 匹配当前词形与词义，保留所选来源完整变体。                                |
| ZP214 | [spring-mechanical](../entries/spring-mechanical.yml) / spring               | noun      | 弹簧；受力形变后能恢复原状的弹性元件。 | `/sprɪŋ/`          | [spring in American English](https://www.collinsdictionary.com/dictionary/english/spring)         | WNW 名词义 17 为恢复原状的弹性装置；匹配当前弹簧义，不取春季义。          |
| ZP217 | [stabilizer-engineering](../entries/stabilizer-engineering.yml) / stabilizer | noun      | 稳定器                                 | `/ˈsteɪbəˌlaɪzər/` | [stabilizer in American English](https://www.collinsdictionary.com/dictionary/english/stabilizer) | 匹配当前词形与词义，保留所选来源完整变体。                                |
| ZP220 | [stake-engineering](../entries/stake-engineering.yml) / stake                | noun      | 桩、固定钉                             | `/steɪk/`          | [stake in American English](https://www.collinsdictionary.com/dictionary/english/stake)           | 匹配当前词形与词义，保留所选来源完整变体。                                |
| ZP256 | [wheel-engineering](../entries/wheel-engineering.yml) / wheel                | noun      | 轮子。                                 | `/hwil/`           | [wheel in American English](https://www.collinsdictionary.com/dictionary/english/wheel)           | 匹配当前词形与词义，保留所选来源完整变体。                                |
| ZP259 | [wing-engineering](../entries/wing-engineering.yml) / wing                   | noun      | 翼；翅膀。                             | `/wɪŋ/`            | [wing in American English](https://www.collinsdictionary.com/dictionary/english/wing)             | 匹配当前词形与词义，保留所选来源完整变体。                                |

### 装备与属性

共 3 项。

| 编号  | 稳定 ID / Word                                                        | 当前词性 | 当前 Meaning                 | 拟写 pronunciation | 来源与分区                                                                                      | 核对说明                                   |
|:------|:----------------------------------------------------------------------|:---------|:-----------------------------|:-------------------|:------------------------------------------------------------------------------------------------|:-------------------------------------------|
| ZP013 | [armor-equipment](../entries/armor-equipment.yml) / armor             | noun     | 护甲                         | `/ˈɑrmər/`         | [armor in American English](https://www.collinsdictionary.com/dictionary/english/armor)         | 匹配当前词形与词义，保留所选来源完整变体。 |
| ZP079 | [equipment-equipment](../entries/equipment-equipment.yml) / equipment | noun     | 装备                         | `/iˈkwɪpmənt/`     | [equipment in American English](https://www.collinsdictionary.com/dictionary/english/equipment) | 匹配当前词形与词义，保留所选来源完整变体。 |
| ZP091 | [gear-equipment](../entries/gear-equipment.yml) / gear                | noun     | 装备；为某种活动准备的器材。 | `/ɡɪr/`            | [gear in American English](https://www.collinsdictionary.com/dictionary/english/gear)           | 匹配当前词形与词义，保留所选来源完整变体。 |

### 烹饪

共 16 项。

| 编号  | 稳定 ID / Word                                                       | 当前词性 | 当前 Meaning                                               | 拟写 pronunciation | 来源与分区                                                                                        | 核对说明                                                                |
|:------|:---------------------------------------------------------------------|:---------|:-----------------------------------------------------------|:-------------------|:--------------------------------------------------------------------------------------------------|:------------------------------------------------------------------------|
| ZP044 | [cook-cooking](../entries/cook-cooking.yml) / cook                   | verb     | 烹饪                                                       | `/kʊk/`            | [cook in American English](https://www.collinsdictionary.com/dictionary/english/cook)             | 匹配当前词形与词义，保留所选来源完整变体。                              |
| ZP070 | [elixir-cooking](../entries/elixir-cooking.yml) / elixir             | noun     | 药剂                                                       | `/ɪˈlɪksər/`       | [elixir in American English](https://www.collinsdictionary.com/dictionary/english/elixir)         | 匹配当前词形与词义，保留所选来源完整变体。                              |
| ZP089 | [fry-cooking](../entries/fry-cooking.yml) / fry                      | verb     | 油煎；油炸；以热油烹饪。                                   | `/fraɪ/`           | [fry in American English 1](https://www.collinsdictionary.com/dictionary/english/fry)             | 匹配当前词形与词义，保留所选来源完整变体。                              |
| ZP096 | [grill-cooking](../entries/grill-cooking.yml) / grill                | verb     | 炙烤；在烤架上烹饪。                                       | `/ɡrɪl/`           | [grill in American English 1](https://www.collinsdictionary.com/dictionary/english/grill)         | 匹配当前词形与词义，保留所选来源完整变体。                              |
| ZP117 | [ingredient-cooking](../entries/ingredient-cooking.yml) / ingredient | noun     | 食材、原料                                                 | `/ɪnˈɡridiənt/`    | [ingredient in American English](https://www.collinsdictionary.com/dictionary/english/ingredient) | 匹配当前词形与词义，保留所选来源完整变体。                              |
| ZP135 | [meal-cooking](../entries/meal-cooking.yml) / meal                   | noun     | 一餐、料理                                                 | `/mil/`            | [meal in American English 1](https://www.collinsdictionary.com/dictionary/english/meal)           | 匹配当前词形与词义，保留所选来源完整变体。                              |
| ZP145 | [omelet-cooking](../entries/omelet-cooking.yml) / omelet             | noun     | 煎蛋卷                                                     | `/ˈɑmlɪt/`         | [omelet in American English](https://www.collinsdictionary.com/dictionary/english/omelet)         | 匹配当前词形与词义，保留所选来源完整变体。                              |
| ZP148 | [paella-cooking](../entries/paella-cooking.yml) / paella             | noun     | 西班牙海鲜饭或烩饭；以米饭配合肉类、海鲜或蔬菜烹制的料理。 | `/pɑˈjɛlə/`        | [paella in American English](https://www.collinsdictionary.com/dictionary/english/paella)         | 选来源列出的英语 /pɑˈjɛlə/，不将同一行标为 Spanish 的读音视为美式英语。 |
| ZP155 | [porridge-cooking](../entries/porridge-cooking.yml) / porridge       | noun     | 粥                                                         | `/ˈpɔrɪdʒ/`        | [porridge in American English](https://www.collinsdictionary.com/dictionary/english/porridge)     | 匹配当前词形与词义，保留所选来源完整变体。                              |
| ZP169 | [recipe-cooking](../entries/recipe-cooking.yml) / recipe             | noun     | 食谱                                                       | `/ˈrɛsəˌpi/`       | [recipe in American English](https://www.collinsdictionary.com/dictionary/english/recipe)         | 匹配当前词形与词义，保留所选来源完整变体。                              |
| ZP177 | [risotto-cooking](../entries/risotto-cooking.yml) / risotto          | noun     | 意大利烩饭                                                 | `/rɪˈzoʊtoʊ/`      | [risotto in American English](https://www.collinsdictionary.com/dictionary/english/risotto)       | 匹配当前词形与词义，保留所选来源完整变体。                              |
| ZP193 | [sear-cooking](../entries/sear-cooking.yml) / sear                   | verb     | 以高温迅速煎烤食物表面。                                   | `/sɪr/`            | [sear in American English 1](https://www.collinsdictionary.com/dictionary/english/sear)           | 匹配当前词形与词义，保留所选来源完整变体。                              |
| ZP199 | [simmer-cooking](../entries/simmer-cooking.yml) / simmer             | verb     | 小火慢煮；使液体保持轻微沸腾。                             | `/ˈsɪmər/`         | [simmer in American English](https://www.collinsdictionary.com/dictionary/english/simmer)         | 匹配当前词形与词义，保留所选来源完整变体。                              |
| ZP200 | [skewer-cooking](../entries/skewer-cooking.yml) / skewer             | noun     | 串食物用的签；串在签上烹饪的食物。                         | `/ˈskjuər/`        | [skewer in American English](https://www.collinsdictionary.com/dictionary/english/skewer)         | 匹配当前词形与词义，保留所选来源完整变体。                              |
| ZP223 | [steam-cooking](../entries/steam-cooking.yml) / steam                | verb     | 蒸；利用蒸汽加热食物。                                     | `/stim/`           | [steam in American English](https://www.collinsdictionary.com/dictionary/english/steam)           | 匹配当前词形与词义，保留所选来源完整变体。                              |
| ZP225 | [stew-cooking](../entries/stew-cooking.yml) / stew                   | noun     | 炖菜                                                       | `/stu/`            | [stew in American English](https://www.collinsdictionary.com/dictionary/english/stew)             | 匹配当前词形与词义，保留所选来源完整变体。                              |

### 食物与材料

共 9 项。

| 编号  | 稳定 ID / Word                                        | 当前词性  | 当前 Meaning                 | 拟写 pronunciation | 来源与分区                                                                                  | 核对说明                                                                      |
|:------|:------------------------------------------------------|:----------|:-----------------------------|:-------------------|:--------------------------------------------------------------------------------------------|:------------------------------------------------------------------------------|
| ZP052 | [curry-food](../entries/curry-food.yml) / curry       | noun      | 咖喱                         | `/ˈkɜri/`          | [curry in American English 2](https://www.collinsdictionary.com/dictionary/english/curry)   | 取 WNW curry in American English 2 的咖喱料理义，不取处理皮革或梳刷的动词义。 |
| ZP095 | [gourmet-food](../entries/gourmet-food.yml) / gourmet | adjective | 美食的；讲究精美饮食的。     | `/ɡʊrˈmeɪ/`        | [gourmet in American English](https://www.collinsdictionary.com/dictionary/english/gourmet) | 匹配当前词形与词义，保留所选来源完整变体。                                    |
| ZP158 | [prime-food](../entries/prime-food.yml) / prime       | adjective | 上等的；优质的。             | `/praɪm/`          | [prime in American English](https://www.collinsdictionary.com/dictionary/english/prime)     | 匹配当前词形与词义，保留所选来源完整变体。                                    |
| ZP166 | [radish-food](../entries/radish-food.yml) / radish    | noun      | 萝卜。                       | `/ˈrædɪʃ/`         | [radish in American English](https://www.collinsdictionary.com/dictionary/english/radish)   | 匹配当前词形与词义，保留所选来源完整变体。                                    |
| ZP178 | [roast-food](../entries/roast-food.yml) / roast       | adjective | 烤制的；烘烤的               | `/roʊst/`          | [roast in American English](https://www.collinsdictionary.com/dictionary/english/roast)     | WNW 义 16 是表示 roasted 的形容词；匹配当前烤制义，不将 Word 改写为 roasted。 |
| ZP189 | [salmon-food](../entries/salmon-food.yml) / salmon    | noun      | 三文鱼；鲑鱼。               | `/ˈsæmən/`         | [salmon in American English](https://www.collinsdictionary.com/dictionary/english/salmon)   | 匹配当前词形与词义，保留所选来源完整变体。                                    |
| ZP192 | [seafood-food](../entries/seafood-food.yml) / seafood | noun      | 海鲜                         | `/ˈsiˌfud/`        | [seafood in American English](https://www.collinsdictionary.com/dictionary/english/seafood) | 匹配当前词形与词义，保留所选来源完整变体。                                    |
| ZP212 | [spice-food](../entries/spice-food.yml) / spice       | noun      | 香辛料；调味香料。           | `/spaɪs/`          | [spice in American English](https://www.collinsdictionary.com/dictionary/english/spice)     | 匹配当前词形与词义，保留所选来源完整变体。                                    |
| ZP241 | [truffle-food](../entries/truffle-food.yml) / truffle | noun      | 松露；生长在地下的食用真菌。 | `/ˈtrʌfəl/`        | [truffle in American English](https://www.collinsdictionary.com/dictionary/english/truffle) | 匹配当前词形与词义，保留所选来源完整变体。                                    |

### 战斗与状态

共 20 项。

| 编号  | 稳定 ID / Word                                                        | 当前词性  | 当前 Meaning                             | 拟写 pronunciation | 来源与分区                                                                                          | 核对说明                                                                               |
|:------|:----------------------------------------------------------------------|:----------|:-----------------------------------------|:-------------------|:----------------------------------------------------------------------------------------------------|:---------------------------------------------------------------------------------------|
| ZP003 | [aggressive-combat](../entries/aggressive-combat.yml) / aggressive    | adjective | 攻击性强的                               | `/əˈɡrɛsɪv/`       | [aggressive in American English](https://www.collinsdictionary.com/dictionary/english/aggressive)   | 匹配当前词形与词义，保留所选来源完整变体。                                             |
| ZP005 | [airborne-combat](../entries/airborne-combat.yml) / airborne          | adjective | 空中的                                   | `/ˈɛrˌbɔrn/`       | [airborne in American English](https://www.collinsdictionary.com/dictionary/english/airborne)       | 匹配当前词形与词义，保留所选来源完整变体。                                             |
| ZP011 | [aquatic-combat](../entries/aquatic-combat.yml) / aquatic             | adjective | 水生的                                   | `/əˈkwɑtɪk/`       | [aquatic in American English](https://www.collinsdictionary.com/dictionary/english/aquatic)         | 匹配当前词形与词义，保留所选来源完整变体。                                             |
| ZP014 | [armored-combat](../entries/armored-combat.yml) / armored             | adjective | 装甲覆盖的                               | `/ˈɑrmərd/`        | [armored in American English](https://www.collinsdictionary.com/dictionary/english/armored)         | 匹配当前词形与词义，保留所选来源完整变体。                                             |
| ZP040 | [colossal-combat](../entries/colossal-combat.yml) / colossal          | adjective | 极其庞大的                               | `/kəˈlɑsəl/`       | [colossal in American English](https://www.collinsdictionary.com/dictionary/english/colossal)       | 匹配当前词形与词义，保留所选来源完整变体。                                             |
| ZP054 | [damage-combat](../entries/damage-combat.yml) / damage                | noun      | 伤害                                     | `/ˈdæmɪdʒ/`        | [damage in American English](https://www.collinsdictionary.com/dictionary/english/damage)           | 匹配当前词形与词义，保留所选来源完整变体。                                             |
| ZP055 | [defense-combat](../entries/defense-combat.yml) / defense             | noun      | 防御力                                   | `/dɪˈfɛns/`        | [defense in American English](https://www.collinsdictionary.com/dictionary/english/defense)         | 取来源列出的 /dɪˈfɛns/，用于当前一般防御义；不混入体育团队集合义常见的首音节重音读法。 |
| ZP065 | [distinctive-combat](../entries/distinctive-combat.yml) / distinctive | adjective | 独特的、显著的                           | `/dɪˈstɪŋktɪv/`    | [distinctive in American English](https://www.collinsdictionary.com/dictionary/english/distinctive) | 匹配当前词形与词义，保留所选来源完整变体。                                             |
| ZP071 | [elusive-combat](../entries/elusive-combat.yml) / elusive             | adjective | 难以捉摸的                               | `/iˈlusɪv/`        | [elusive in American English](https://www.collinsdictionary.com/dictionary/english/elusive)         | 匹配当前词形与词义，保留所选来源完整变体。                                             |
| ZP075 | [enormous-combat](../entries/enormous-combat.yml) / enormous          | adjective | 巨大的                                   | `/iˈnɔrməs/`       | [enormous in American English](https://www.collinsdictionary.com/dictionary/english/enormous)       | 匹配当前词形与词义，保留所选来源完整变体。                                             |
| ZP113 | [hostile-combat](../entries/hostile-combat.yml) / hostile             | adjective | 敌对的                                   | `/ˈhɑstəl/`        | [hostile in American English](https://www.collinsdictionary.com/dictionary/english/hostile)         | 选来源通常美式 /ˈhɑstəl/，不选同一行明确标为 chiefly British 的其他变体。              |
| ZP136 | [melee-combat](../entries/melee-combat.yml) / melee                   | noun      | 近战                                     | `/ˈmeɪˌleɪ/`       | [melee in American English](https://www.collinsdictionary.com/dictionary/english/melee)             | 匹配当前词形与词义，保留所选来源完整变体。                                             |
| ZP146 | [one-handed-combat](../entries/one-handed-combat.yml) / one-handed    | adjective | 单手的                                   | `/ˈwʌnˈhændɪd/`    | [one-handed in American English](https://www.collinsdictionary.com/dictionary/english/one-handed)   | 来源完整给出 one-handed 形容词的 /ˈwʌnˈhændɪd/，保留该来源重音记法，不由组件拼读。     |
| ZP156 | [possessed-combat](../entries/possessed-combat.yml) / possessed       | adjective | 被附身的                                 | `/pəˈzɛst/`        | [possessed in American English](https://www.collinsdictionary.com/dictionary/english/possessed)     | 匹配当前词形与词义，保留所选来源完整变体。                                             |
| ZP174 | [resilient-combat](../entries/resilient-combat.yml) / resilient       | adjective | 有韧性的；遭受打击后能恢复或继续应对的。 | `/rɪˈzɪljənt/`     | [resilient in American English](https://www.collinsdictionary.com/dictionary/english/resilient)     | 匹配当前词形与词义，保留所选来源完整变体。                                             |
| ZP229 | [susceptible-combat](../entries/susceptible-combat.yml) / susceptible | adjective | 易受某种影响的                           | `/səˈsɛptəbəl/`    | [susceptible in American English](https://www.collinsdictionary.com/dictionary/english/susceptible) | 匹配当前词形与词义，保留所选来源完整变体。                                             |
| ZP243 | [two-handed-combat](../entries/two-handed-combat.yml) / two-handed    | adjective | 双手的                                   | `/ˌtuˈhændɪd/`     | [two-handed in American English](https://www.collinsdictionary.com/dictionary/english/two-handed)   | 匹配当前词形与词义，保留所选来源完整变体。                                             |
| ZP244 | [undead-combat](../entries/undead-combat.yml) / undead                | adjective | 不死的                                   | `/ʌnˈdɛd/`         | [undead in American English](https://www.collinsdictionary.com/dictionary/english/undead)           | 匹配当前词形与词义，保留所选来源完整变体。                                             |
| ZP250 | [volcanic-combat](../entries/volcanic-combat.yml) / volcanic          | adjective | 火山的                                   | `/vɑlˈkænɪk/`      | [volcanic in American English](https://www.collinsdictionary.com/dictionary/english/volcanic)       | 匹配当前词形与词义，保留所选来源完整变体。                                             |
| ZP251 | [vulnerable-combat](../entries/vulnerable-combat.yml) / vulnerable    | adjective | 易受攻击的                               | `/ˈvʌlnərəbəl/`    | [vulnerable in American English](https://www.collinsdictionary.com/dictionary/english/vulnerable)   | 匹配当前词形与词义，保留所选来源完整变体。                                             |

### 界面与操作

共 7 项。

| 编号  | 稳定 ID / Word                                                | 当前词性 | 当前 Meaning                       | 拟写 pronunciation | 来源与分区                                                                                          | 核对说明                                           |
|:------|:--------------------------------------------------------------|:---------|:-----------------------------------|:-------------------|:----------------------------------------------------------------------------------------------------|:---------------------------------------------------|
| ZP006 | [album-ui](../entries/album-ui.yml) / album                   | noun     | 相册；收集照片或图片的册子。       | `/ˈælbəm/`         | [album in American English](https://www.collinsdictionary.com/dictionary/english/album)             | 匹配当前词形与词义，保留所选来源完整变体。         |
| ZP060 | [destination-ui](../entries/destination-ui.yml) / destination | noun     | 目的地                             | `/ˌdɛstəˈneɪʃən/`  | [destination in American English](https://www.collinsdictionary.com/dictionary/english/destination) | 匹配当前词形与词义，保留所选来源完整变体。         |
| ZP122 | [inventory-ui](../entries/inventory-ui.yml) / inventory       | noun     | 物品清单；库存；游戏中的物品栏。   | `/ˈɪnvənˌtɔri/`    | [inventory in American English](https://www.collinsdictionary.com/dictionary/english/inventory)     | 匹配当前词形与词义，保留所选来源完整变体。         |
| ZP144 | [objective-ui](../entries/objective-ui.yml) / objective       | noun     | 当前目标                           | `/əbˈdʒɛktɪv/`     | [objective in American English](https://www.collinsdictionary.com/dictionary/english/objective)     | WNW 名词义 10 为目标，不改成客观的形容词学习目标。 |
| ZP157 | [pouch-ui](../entries/pouch-ui.yml) / pouch                   | noun     | 小袋；随身袋。                     | `/paʊtʃ/`          | [pouch in American English](https://www.collinsdictionary.com/dictionary/english/pouch)             | 匹配当前词形与词义，保留所选来源完整变体。         |
| ZP162 | [profile-ui](../entries/profile-ui.yml) / profile             | noun     | 简介；概况。                       | `/ˈproʊˌfaɪl/`     | [profile in American English](https://www.collinsdictionary.com/dictionary/english/profile)         | 匹配当前词形与词义，保留所选来源完整变体。         |
| ZP231 | [system-ui](../entries/system-ui.yml) / system                | noun     | 系统；由相互联系的部分组成的整体。 | `/ˈsɪstəm/`        | [system in American English](https://www.collinsdictionary.com/dictionary/english/system)           | 匹配当前词形与词义，保留所选来源完整变体。         |

### 剧情与人物称谓

共 21 项。

| 编号  | 稳定 ID / Word                                                   | 当前词性 | 当前 Meaning           | 拟写 pronunciation | 来源与分区                                                                                        | 核对说明                                                 |
|:------|:-----------------------------------------------------------------|:---------|:-----------------------|:-------------------|:--------------------------------------------------------------------------------------------------|:---------------------------------------------------------|
| ZP016 | [awaken-story](../entries/awaken-story.yml) / awaken             | verb     | 觉醒                   | `/əˈweɪkən/`       | [awaken in American English](https://www.collinsdictionary.com/dictionary/english/awaken)         | 匹配当前词形与词义，保留所选来源完整变体。               |
| ZP026 | [calamity-story](../entries/calamity-story.yml) / calamity       | noun     | 灾厄、大灾难           | `/kəˈlæməti/`      | [calamity in American English](https://www.collinsdictionary.com/dictionary/english/calamity)     | 匹配当前词形与词义，保留所选来源完整变体。               |
| ZP039 | [clue-story](../entries/clue-story.yml) / clue                   | noun     | 线索                   | `/klu/`            | [clue in American English](https://www.collinsdictionary.com/dictionary/english/clue)             | 匹配当前词形与词义，保留所选来源完整变体。               |
| ZP048 | [crisis-story](../entries/crisis-story.yml) / crisis             | noun     | 危机                   | `/ˈkraɪsɪs/`       | [crisis in American English](https://www.collinsdictionary.com/dictionary/english/crisis)         | 匹配当前词形与词义，保留所选来源完整变体。               |
| ZP056 | [deity-story](../entries/deity-story.yml) / deity                | noun     | 神；神明。             | `/ˈdiəti/`         | [deity in American English](https://www.collinsdictionary.com/dictionary/english/deity)           | 匹配当前词形与词义，保留所选来源完整变体。               |
| ZP057 | [demon-story](../entries/demon-story.yml) / demon                | noun     | 恶魔；魔鬼。           | `/ˈdimən/`         | [demon in American English](https://www.collinsdictionary.com/dictionary/english/demon)           | 匹配当前词形与词义，保留所选来源完整变体。               |
| ZP068 | [dragon-story](../entries/dragon-story.yml) / dragon             | noun     | 龙。                   | `/ˈdræɡən/`        | [dragon in American English](https://www.collinsdictionary.com/dictionary/english/dragon)         | 匹配当前词形与词义，保留所选来源完整变体。               |
| ZP098 | [guidance-story](../entries/guidance-story.yml) / guidance       | noun     | 指引、指导             | `/ˈɡaɪdəns/`       | [guidance in American English](https://www.collinsdictionary.com/dictionary/english/guidance)     | 匹配当前词形与词义，保留所选来源完整变体。               |
| ZP116 | [imprison-story](../entries/imprison-story.yml) / imprison       | verb     | 囚禁；监禁。           | `/ɪmˈprɪzən/`      | [imprison in American English](https://www.collinsdictionary.com/dictionary/english/imprison)     | 匹配当前词形与词义，保留所选来源完整变体。               |
| ZP118 | [inherit-story](../entries/inherit-story.yml) / inherit          | verb     | 继承                   | `/ɪnˈhɛrɪt/`       | [inherit in American English](https://www.collinsdictionary.com/dictionary/english/inherit)       | 匹配当前词形与词义，保留所选来源完整变体。               |
| ZP137 | [memory-story](../entries/memory-story.yml) / memory             | noun     | 记忆                   | `/ˈmɛməri/`        | [memory in American English](https://www.collinsdictionary.com/dictionary/english/memory)         | 匹配当前词形与词义，保留所选来源完整变体。               |
| ZP142 | [mystery-story](../entries/mystery-story.yml) / mystery          | noun     | 谜团、神秘之事         | `/ˈmɪstəri/`       | [mystery in American English 1](https://www.collinsdictionary.com/dictionary/english/mystery)     | 匹配当前词形与词义，保留所选来源完整变体。               |
| ZP152 | [phenomenon-story](../entries/phenomenon-story.yml) / phenomenon | noun     | 现象                   | `/fəˈnɑmənən/`     | [phenomenon in American English](https://www.collinsdictionary.com/dictionary/english/phenomenon) | 匹配当前词形与词义，保留所选来源完整变体。               |
| ZP163 | [prophecy-story](../entries/prophecy-story.yml) / prophecy       | noun     | 预言                   | `/ˈprɑfəsi/`       | [prophecy in American English](https://www.collinsdictionary.com/dictionary/english/prophecy)     | 匹配当前词形与词义，保留所选来源完整变体。               |
| ZP171 | [recover-story](../entries/recover-story.yml) / recover          | verb     | 找回、恢复             | `/rɪˈkʌvər/`       | [recover in American English](https://www.collinsdictionary.com/dictionary/english/recover)       | 匹配当前词形与词义，保留所选来源完整变体。               |
| ZP187 | [sacrifice-story](../entries/sacrifice-story.yml) / sacrifice    | noun     | 牺牲                   | `/ˈsækrəˌfaɪs/`    | [sacrifice in American English](https://www.collinsdictionary.com/dictionary/english/sacrifice)   | 匹配当前词形与词义，保留所选来源完整变体。               |
| ZP194 | [secret-story](../entries/secret-story.yml) / secret             | noun     | 秘密                   | `/ˈsikrɪt/`        | [secret in American English](https://www.collinsdictionary.com/dictionary/english/secret)         | WNW 名词义 7 为秘密；不因条头先列形容词而更改当前 noun。 |
| ZP213 | [spirit-story](../entries/spirit-story.yml) / spirit             | noun     | 精神；灵魂。           | `/ˈspɪrɪt/`        | [spirit in American English](https://www.collinsdictionary.com/dictionary/english/spirit)         | 匹配当前词形与词义，保留所选来源完整变体。               |
| ZP236 | [trail-story](../entries/trail-story.yml) / trail                | noun     | 踪迹、线索             | `/treɪl/`          | [trail in American English](https://www.collinsdictionary.com/dictionary/english/trail)           | 匹配当前词形与词义，保留所选来源完整变体。               |
| ZP248 | [upheaval-story](../entries/upheaval-story.yml) / upheaval       | noun     | 剧变；动荡；重大变动。 | `/ʌpˈhivəl/`       | [upheaval in American English](https://www.collinsdictionary.com/dictionary/english/upheaval)     | 匹配当前词形与词义，保留所选来源完整变体。               |
| ZP257 | [wield-story](../entries/wield-story.yml) / wield                | verb     | 挥舞、使用武器         | `/wild/`           | [wield in American English](https://www.collinsdictionary.com/dictionary/english/wield)           | 匹配当前词形与词义，保留所选来源完整变体。               |

### 通用词与描述

共 88 项。

| 编号  | 稳定 ID / Word                                                             | 当前词性       | 当前 Meaning                                 | 拟写 pronunciation | 来源与分区                                                                                            | 核对说明                                                                                    |
|:------|:---------------------------------------------------------------------------|:---------------|:---------------------------------------------|:-------------------|:------------------------------------------------------------------------------------------------------|:--------------------------------------------------------------------------------------------|
| ZP001 | [activate-general](../entries/activate-general.yml) / activate             | verb           | 激活、启动                                   | `/ˈæktəˌveɪt/`     | [activate in American English](https://www.collinsdictionary.com/dictionary/english/activate)         | 匹配当前词形与词义，保留所选来源完整变体。                                                  |
| ZP002 | [adventure-general](../entries/adventure-general.yml) / adventure          | noun           | 冒险；冒险经历。                             | `/ædˈvɛntʃər/`     | [adventure in American English](https://www.collinsdictionary.com/dictionary/english/adventure)       | 匹配当前词形与词义，保留所选来源完整变体。                                                  |
| ZP004 | [aim-general](../entries/aim-general.yml) / aim                            | verb           | 瞄准                                         | `/eɪm/`            | [aim in American English](https://www.collinsdictionary.com/dictionary/english/aim)                   | 匹配当前词形与词义，保留所选来源完整变体。                                                  |
| ZP008 | [ambitious-general](../entries/ambitious-general.yml) / ambitious          | adjective      | 有野心的                                     | `/æmˈbɪʃəs/`       | [ambitious in American English](https://www.collinsdictionary.com/dictionary/english/ambitious)       | 匹配当前词形与词义，保留所选来源完整变体。                                                  |
| ZP009 | [ancestor-general](../entries/ancestor-general.yml) / ancestor             | noun           | 祖先                                         | `/ˈænˌsɛstər/`     | [ancestor in American English](https://www.collinsdictionary.com/dictionary/english/ancestor)         | 匹配当前词形与词义，保留所选来源完整变体。                                                  |
| ZP010 | [ancient-general](../entries/ancient-general.yml) / ancient                | adjective      | 古代的                                       | `/ˈeɪntʃənt/`      | [ancient in American English 1](https://www.collinsdictionary.com/dictionary/english/ancient)         | 匹配当前词形与词义，保留所选来源完整变体。                                                  |
| ZP020 | [blessing-general](../entries/blessing-general.yml) / blessing             | noun           | 祝福；被视为有益或带来幸福的事物。           | `/ˈblɛsɪŋ/`        | [blessing in American English](https://www.collinsdictionary.com/dictionary/english/blessing)         | 匹配当前词形与词义，保留所选来源完整变体。                                                  |
| ZP021 | [block-general](../entries/block-general.yml) / block                      | verb           | 格挡；挡住攻击或去路。                       | `/blɑk/`           | [block in American English](https://www.collinsdictionary.com/dictionary/english/block)               | 匹配当前词形与词义，保留所选来源完整变体。                                                  |
| ZP032 | [champion-general](../entries/champion-general.yml) / champion             | noun           | 英杰、冠军                                   | `/ˈtʃæmpiən/`      | [champion in American English](https://www.collinsdictionary.com/dictionary/english/champion)         | 匹配当前词形与词义，保留所选来源完整变体。                                                  |
| ZP035 | [civilization-general](../entries/civilization-general.yml) / civilization | noun           | 文明                                         | `/ˌsɪvələˈzeɪʃən/` | [civilization in American English](https://www.collinsdictionary.com/dictionary/english/civilization) | 匹配当前词形与词义，保留所选来源完整变体。                                                  |
| ZP037 | [closed-general](../entries/closed-general.yml) / closed                   | adjective      | 关闭的；未敞开的。                           | `/kloʊzd/`         | [closed in American English](https://www.collinsdictionary.com/dictionary/english/closed)             | 匹配当前词形与词义，保留所选来源完整变体。                                                  |
| ZP041 | [compendium-general](../entries/compendium-general.yml) / compendium       | noun           | 概要汇编；汇集某领域资料的手册。             | `/kəmˈpɛndiəm/`    | [compendium in American English](https://www.collinsdictionary.com/dictionary/english/compendium)     | 匹配当前词形与词义，保留所选来源完整变体。                                                  |
| ZP043 | [construct-general](../entries/construct-general.yml) / construct          | verb           | 建造、构造                                   | `/kənˈstrʌkt/`     | [construct in American English](https://www.collinsdictionary.com/dictionary/english/construct)       | 取当前动词 /kənˈstrʌkt/；来源另列名词 /ˈkɑnˌstrʌkt/，不可混用。                             |
| ZP046 | [courageous-general](../entries/courageous-general.yml) / courageous       | adjective      | 勇敢的                                       | `/kəˈreɪdʒəs/`     | [courageous in American English](https://www.collinsdictionary.com/dictionary/english/courageous)     | 匹配当前词形与词义，保留所选来源完整变体。                                                  |
| ZP049 | [crouch-general](../entries/crouch-general.yml) / crouch                   | verb           | 蹲下                                         | `/kraʊtʃ/`         | [crouch in American English](https://www.collinsdictionary.com/dictionary/english/crouch)             | 匹配当前词形与词义，保留所选来源完整变体。                                                  |
| ZP050 | [crystal-general](../entries/crystal-general.yml) / crystal                | noun           | 水晶；晶体。                                 | `/ˈkrɪstəl/`       | [crystal in American English](https://www.collinsdictionary.com/dictionary/english/crystal)           | 匹配当前词形与词义，保留所选来源完整变体。                                                  |
| ZP051 | [cunning-general](../entries/cunning-general.yml) / cunning                | adjective      | 狡猾的                                       | `/ˈkʌnɪŋ/`         | [cunning in American English](https://www.collinsdictionary.com/dictionary/english/cunning)           | 匹配当前词形与词义，保留所选来源完整变体。                                                  |
| ZP058 | [descendant-general](../entries/descendant-general.yml) / descendant       | noun           | 后裔                                         | `/diˈsɛndənt/`     | [descendant in American English](https://www.collinsdictionary.com/dictionary/english/descendant)     | 匹配当前词形与词义，保留所选来源完整变体。                                                  |
| ZP061 | [determined-general](../entries/determined-general.yml) / determined       | adjective      | 坚定的                                       | `/diˈtɜrmənd/`     | [determined in American English](https://www.collinsdictionary.com/dictionary/english/determined)     | 匹配当前词形与词义，保留所选来源完整变体。                                                  |
| ZP063 | [diligent-general](../entries/diligent-general.yml) / diligent             | adjective      | 勤勉的                                       | `/ˈdɪlədʒənt/`     | [diligent in American English](https://www.collinsdictionary.com/dictionary/english/diligent)         | 匹配当前词形与词义，保留所选来源完整变体。                                                  |
| ZP066 | [dive-general](../entries/dive-general.yml) / dive                         | verb           | 俯冲、潜入                                   | `/daɪv/`           | [dive in American English](https://www.collinsdictionary.com/dictionary/english/dive)                 | 匹配当前词形与词义，保留所选来源完整变体。                                                  |
| ZP067 | [dodge-general](../entries/dodge-general.yml) / dodge                      | verb           | 闪避                                         | `/dɑdʒ/`           | [dodge in American English](https://www.collinsdictionary.com/dictionary/english/dodge)               | 匹配当前词形与词义，保留所选来源完整变体。                                                  |
| ZP069 | [draw-general](../entries/draw-general.yml) / draw                         | verb           | 拉开（弓）；拔出（武器）。                   | `/drɔ/`            | [draw in American English](https://www.collinsdictionary.com/dictionary/english/draw)                 | WNW 动词义 1 明确涵盖拉弓弦，保留原词义及词形。                                             |
| ZP073 | [endure-general](../entries/endure-general.yml) / endure                   | verb           | 持久；忍受；持续存在。                       | `/ɛnˈdʊr/`         | [endure in American English](https://www.collinsdictionary.com/dictionary/english/endure)             | 匹配当前词形与词义，保留所选来源完整变体。                                                  |
| ZP074 | [energize-general](../entries/energize-general.yml) / energize             | verb           | 使充满活力；补充能量。                       | `/ˈɛnərˌdʒaɪz/`    | [energize in American English](https://www.collinsdictionary.com/dictionary/english/energize)         | 匹配当前词形与词义，保留所选来源完整变体。                                                  |
| ZP076 | [enthusiastic-general](../entries/enthusiastic-general.yml) / enthusiastic | adjective      | 热情的                                       | `/ɛnˌθuziˈæstɪk/`  | [enthusiastic in American English](https://www.collinsdictionary.com/dictionary/english/enthusiastic) | 匹配当前词形与词义，保留所选来源完整变体。                                                  |
| ZP078 | [equip-general](../entries/equip-general.yml) / equip                      | verb           | 装备                                         | `/iˈkwɪp/`         | [equip in American English 1](https://www.collinsdictionary.com/dictionary/english/equip)             | 匹配当前词形与词义，保留所选来源完整变体。                                                  |
| ZP081 | [fearless-general](../entries/fearless-general.yml) / fearless             | adjective      | 无畏的                                       | `/ˈfɪrlɪs/`        | [fearless in American English](https://www.collinsdictionary.com/dictionary/english/fearless)         | 匹配当前词形与词义，保留所选来源完整变体。                                                  |
| ZP082 | [fierce-general](../entries/fierce-general.yml) / fierce                   | adjective      | 凶猛的、猛烈的                               | `/fɪrs/`           | [fierce in American English](https://www.collinsdictionary.com/dictionary/english/fierce)             | 匹配当前词形与词义，保留所选来源完整变体。                                                  |
| ZP083 | [fifth-general](../entries/fifth-general.yml) / fifth                      | ordinal number | 第五。                                       | `/fɪfθ/`           | [fifth in American English](https://www.collinsdictionary.com/dictionary/english/fifth)               | 匹配当前词形与词义，保留所选来源完整变体。                                                  |
| ZP088 | [frost-general](../entries/frost-general.yml) / frost                      | noun           | 霜；冰霜。                                   | `/frɔst/`          | [frost in American English](https://www.collinsdictionary.com/dictionary/english/frost)               | 匹配当前词形与词义，保留所选来源完整变体。                                                  |
| ZP090 | [fuse-general](../entries/fuse-general.yml) / fuse                         | verb           | 融合；结合；使物体连接成为一体。             | `/fjuz/`           | [fuse in American English 1](https://www.collinsdictionary.com/dictionary/english/fuse)               | 匹配当前词形与词义，保留所选来源完整变体。                                                  |
| ZP093 | [glide-general](../entries/glide-general.yml) / glide                      | verb           | 滑翔                                         | `/ɡlaɪd/`          | [glide in American English](https://www.collinsdictionary.com/dictionary/english/glide)               | 匹配当前词形与词义，保留所选来源完整变体。                                                  |
| ZP097 | [guardian-general](../entries/guardian-general.yml) / guardian             | noun           | 守护者                                       | `/ˈɡɑrdiən/`       | [guardian in American English](https://www.collinsdictionary.com/dictionary/english/guardian)         | 匹配当前词形与词义，保留所选来源完整变体。                                                  |
| ZP101 | [hasty-general](../entries/hasty-general.yml) / hasty                      | adjective      | 匆忙的；仓促的。                             | `/ˈheɪsti/`        | [hasty in American English](https://www.collinsdictionary.com/dictionary/english/hasty)               | 匹配当前词形与词义，保留所选来源完整变体。                                                  |
| ZP104 | [hearty-general](../entries/hearty-general.yml) / hearty                   | adjective      | 丰盛的；有益健康的；充满活力的。             | `/ˈhɑrti/`         | [hearty in American English](https://www.collinsdictionary.com/dictionary/english/hearty)             | 匹配当前词形与词义，保留所选来源完整变体。                                                  |
| ZP105 | [heir-general](../entries/heir-general.yml) / heir                         | noun           | 继承人                                       | `/ɛr/`             | [heir in American English](https://www.collinsdictionary.com/dictionary/english/heir)                 | 匹配当前词形与词义，保留所选来源完整变体。                                                  |
| ZP108 | [hero-general](../entries/hero-general.yml) / hero                         | noun           | 英雄                                         | `/ˈhɪroʊ/`         | [hero in American English](https://www.collinsdictionary.com/dictionary/english/hero)                 | 匹配当前词形与词义，保留所选来源完整变体。                                                  |
| ZP111 | [honorable-general](../entries/honorable-general.yml) / honorable          | adjective      | 值得尊敬的                                   | `/ˈɑnərəbəl/`      | [honorable in American English](https://www.collinsdictionary.com/dictionary/english/honorable)       | 匹配当前词形与词义，保留所选来源完整变体。                                                  |
| ZP114 | [hover-general](../entries/hover-general.yml) / hover                      | verb           | 悬停；在空中保持相近位置。                   | `/ˈhʌvər/`         | [hover in American English](https://www.collinsdictionary.com/dictionary/english/hover)               | 匹配当前词形与词义，保留所选来源完整变体。                                                  |
| ZP120 | [integrity-general](../entries/integrity-general.yml) / integrity          | noun           | 正直；诚信；坚持诚实和道德原则。             | `/ɪnˈtɛɡrəti/`     | [integrity in American English](https://www.collinsdictionary.com/dictionary/english/integrity)       | 匹配当前词形与词义，保留所选来源完整变体。                                                  |
| ZP121 | [interact-general](../entries/interact-general.yml) / interact             | verb           | 交互                                         | `/ˌɪntərˈækt/`     | [interact in American English](https://www.collinsdictionary.com/dictionary/english/interact)         | 匹配当前词形与词义，保留所选来源完整变体。                                                  |
| ZP123 | [king-general](../entries/king-general.yml) / king                         | noun           | 国王                                         | `/kɪŋ/`            | [king in American English](https://www.collinsdictionary.com/dictionary/english/king)                 | 匹配当前词形与词义，保留所选来源完整变体。                                                  |
| ZP124 | [kingdom-general](../entries/kingdom-general.yml) / kingdom                | noun           | 王国；由国王或女王统治的国家。               | `/ˈkɪŋdəm/`        | [kingdom in American English](https://www.collinsdictionary.com/dictionary/english/kingdom)           | 匹配当前词形与词义，保留所选来源完整变体。                                                  |
| ZP125 | [knight-general](../entries/knight-general.yml) / knight                   | noun           | 骑士。                                       | `/naɪt/`           | [knight in American English](https://www.collinsdictionary.com/dictionary/english/knight)             | 匹配当前词形与词义，保留所选来源完整变体。                                                  |
| ZP126 | [landing-general](../entries/landing-general.yml) / landing                | noun           | 着陆；登陆；登岸处。                         | `/ˈlændɪŋ/`        | [landing in American English](https://www.collinsdictionary.com/dictionary/english/landing)           | 匹配当前词形与词义，保留所选来源完整变体。                                                  |
| ZP130 | [lookout-general](../entries/lookout-general.yml) / lookout                | noun           | 瞭望；警戒人员；瞭望处。                     | `/ˈlʊkˌaʊt/`       | [lookout in American English](https://www.collinsdictionary.com/dictionary/english/lookout)           | 匹配当前词形与词义，保留所选来源完整变体。                                                  |
| ZP131 | [majestic-general](../entries/majestic-general.yml) / majestic             | adjective      | 雄伟的、威严的                               | `/məˈdʒɛstɪk/`     | [majestic in American English](https://www.collinsdictionary.com/dictionary/english/majestic)         | 匹配当前词形与词义，保留所选来源完整变体。                                                  |
| ZP132 | [malevolent-general](../entries/malevolent-general.yml) / malevolent       | adjective      | 邪恶的、怀有恶意的                           | `/məˈlɛvələnt/`    | [malevolent in American English](https://www.collinsdictionary.com/dictionary/english/malevolent)     | 匹配当前词形与词义，保留所选来源完整变体。                                                  |
| ZP134 | [material-general](../entries/material-general.yml) / material             | noun           | 材料；用于制作物品的原料。                   | `/məˈtɪriəl/`      | [material in American English](https://www.collinsdictionary.com/dictionary/english/material)         | WNW 名词义 6 为材料和原料；不因条头先列形容词而更改当前词性。                               |
| ZP138 | [mighty-general](../entries/mighty-general.yml) / mighty                   | adjective      | 强大的                                       | `/ˈmaɪti/`         | [mighty in American English](https://www.collinsdictionary.com/dictionary/english/mighty)             | 匹配当前词形与词义，保留所选来源完整变体。                                                  |
| ZP140 | [miner-general](../entries/miner-general.yml) / miner                      | noun           | 矿工；采矿的人。                             | `/ˈmaɪnər/`        | [miner in American English](https://www.collinsdictionary.com/dictionary/english/miner)               | 匹配当前词形与词义，保留所选来源完整变体。                                                  |
| ZP141 | [mysterious-general](../entries/mysterious-general.yml) / mysterious       | adjective      | 神秘的                                       | `/mɪsˈtɪriəs/`     | [mysterious in American English](https://www.collinsdictionary.com/dictionary/english/mysterious)     | 匹配当前词形与词义，保留所选来源完整变体。                                                  |
| ZP143 | [noble-general](../entries/noble-general.yml) / noble                      | adjective      | 高贵的                                       | `/ˈnoʊbəl/`        | [noble in American English](https://www.collinsdictionary.com/dictionary/english/noble)               | 匹配当前词形与词义，保留所选来源完整变体。                                                  |
| ZP147 | [pad-general](../entries/pad-general.yml) / pad                            | noun           | 垫；写字用的便笺簿；此处也用于平板装置名称。 | `/pæd/`            | [pad in American English 2](https://www.collinsdictionary.com/dictionary/english/pad)                 | 取 WNW pad in American English 2，包含软垫和书写便笺簿义，不取脚步声或步行分区。            |
| ZP159 | [prince-general](../entries/prince-general.yml) / prince                   | noun           | 王子                                         | `/prɪns/`          | [prince in American English](https://www.collinsdictionary.com/dictionary/english/prince)             | 匹配当前词形与词义，保留所选来源完整变体。                                                  |
| ZP160 | [princess-general](../entries/princess-general.yml) / princess             | noun           | 公主                                         | `/ˈprɪnsɪs/`       | [princess in American English 1](https://www.collinsdictionary.com/dictionary/english/princess)       | 选来源首个美式变体 /ˈprɪnsɪs/；同一行标为 British 的变体不拼入拟值。                        |
| ZP161 | [pristine-general](../entries/pristine-general.yml) / pristine             | adjective      | 完好如初的；未受损或未被改变的。             | `/ˈprɪsˌtin/`      | [pristine in American English](https://www.collinsdictionary.com/dictionary/english/pristine)         | 选来源列出的 /ˈprɪsˌtin/，不选标为 chiefly British 的 /aɪ/ 变体。                           |
| ZP164 | [queen-general](../entries/queen-general.yml) / queen                      | noun           | 王后、女王                                   | `/kwin/`           | [queen in American English](https://www.collinsdictionary.com/dictionary/english/queen)               | 匹配当前词形与词义，保留所选来源完整变体。                                                  |
| ZP167 | [recall-general](../entries/recall-general.yml) / recall                   | verb           | 回想；召回；叫回。                           | `/rɪˈkɔl/`         | [recall in American English](https://www.collinsdictionary.com/dictionary/english/recall)             | 匹配当前词形与词义，保留所选来源完整变体。                                                  |
| ZP168 | [recharge-general](../entries/recharge-general.yml) / recharge             | verb           | 充能                                         | `/riˈtʃɑrdʒ/`      | [recharge in American English](https://www.collinsdictionary.com/dictionary/english/recharge)         | 匹配当前词形与词义，保留所选来源完整变体。                                                  |
| ZP170 | [reckless-general](../entries/reckless-general.yml) / reckless             | adjective      | 鲁莽的                                       | `/ˈrɛklɪs/`        | [reckless in American English](https://www.collinsdictionary.com/dictionary/english/reckless)         | 匹配当前词形与词义，保留所选来源完整变体。                                                  |
| ZP173 | [regional-general](../entries/regional-general.yml) / regional             | adjective      | 地区性的                                     | `/ˈridʒənəl/`      | [regional in American English](https://www.collinsdictionary.com/dictionary/english/regional)         | 匹配当前词形与词义，保留所选来源完整变体。                                                  |
| ZP175 | [resistance-general](../entries/resistance-general.yml) / resistance       | noun           | 抵抗；抵抗力。                               | `/rɪˈzɪstəns/`     | [resistance in American English](https://www.collinsdictionary.com/dictionary/english/resistance)     | 匹配当前词形与词义，保留所选来源完整变体。                                                  |
| ZP181 | [rod-general](../entries/rod-general.yml) / rod                            | noun           | 杆；棒。                                     | `/rɑd/`            | [rod in American English](https://www.collinsdictionary.com/dictionary/english/rod)                   | 匹配当前词形与词义，保留所选来源完整变体。                                                  |
| ZP182 | [rotate-general](../entries/rotate-general.yml) / rotate                   | verb           | 旋转                                         | `/ˈroʊˌteɪt/`      | [rotate in American English](https://www.collinsdictionary.com/dictionary/english/rotate)             | 匹配当前词形与词义，保留所选来源完整变体。                                                  |
| ZP183 | [royal-general](../entries/royal-general.yml) / royal                      | adjective      | 王室的                                       | `/ˈrɔɪəl/`         | [royal in American English](https://www.collinsdictionary.com/dictionary/english/royal)               | 匹配当前词形与词义，保留所选来源完整变体。                                                  |
| ZP185 | [ruthless-general](../entries/ruthless-general.yml) / ruthless             | adjective      | 无情的                                       | `/ˈruːθlɪs/`       | [ruthless in American English](https://www.collinsdictionary.com/dictionary/english/ruthless)         | WNW 分区显示异常的双重重音形式，本项改选同页另一美式分区完整的 /ˈruːθlɪs/，不自行修补 WNW。 |
| ZP186 | [sacred-general](../entries/sacred-general.yml) / sacred                   | adjective      | 神圣的                                       | `/ˈseɪkrɪd/`       | [sacred in American English](https://www.collinsdictionary.com/dictionary/english/sacred)             | 匹配当前词形与词义，保留所选来源完整变体。                                                  |
| ZP188 | [sage-general](../entries/sage-general.yml) / sage                         | noun           | 贤者；富有智慧的人。                         | `/seɪdʒ/`          | [sage in American English 1](https://www.collinsdictionary.com/dictionary/english/sage)               | 匹配当前词形与词义，保留所选来源完整变体。                                                  |
| ZP190 | [savage-general](../entries/savage-general.yml) / savage                   | adjective      | 凶猛的                                       | `/ˈsævɪdʒ/`        | [savage in American English](https://www.collinsdictionary.com/dictionary/english/savage)             | 匹配当前词形与词义，保留所选来源完整变体。                                                  |
| ZP195 | [selfless-general](../entries/selfless-general.yml) / selfless             | adjective      | 无私的                                       | `/ˈsɛlflɪs/`       | [selfless in American English](https://www.collinsdictionary.com/dictionary/english/selfless)         | 匹配当前词形与词义，保留所选来源完整变体。                                                  |
| ZP201 | [skilled-general](../entries/skilled-general.yml) / skilled                | adjective      | 熟练的、有技巧的                             | `/skɪld/`          | [skilled in American English](https://www.collinsdictionary.com/dictionary/english/skilled)           | 匹配当前词形与词义，保留所选来源完整变体。                                                  |
| ZP205 | [slip-general](../entries/slip-general.yml) / slip                         | verb           | 滑动；滑倒；失去抓地力。                     | `/slɪp/`           | [slip in American English 1](https://www.collinsdictionary.com/dictionary/english/slip)               | 匹配当前词形与词义，保留所选来源完整变体。                                                  |
| ZP207 | [sneaky-general](../entries/sneaky-general.yml) / sneaky                   | adjective      | 偷偷摸摸的；行动隐蔽的。                     | `/ˈsniki/`         | [sneaky in American English](https://www.collinsdictionary.com/dictionary/english/sneaky)             | 匹配当前词形与词义，保留所选来源完整变体。                                                  |
| ZP208 | [soldier-general](../entries/soldier-general.yml) / soldier                | noun           | 士兵。                                       | `/ˈsoʊldʒər/`      | [soldier in American English](https://www.collinsdictionary.com/dictionary/english/soldier)           | 匹配当前词形与词义，保留所选来源完整变体。                                                  |
| ZP209 | [sort-general](../entries/sort-general.yml) / sort                         | verb           | 排序                                         | `/sɔrt/`           | [sort in American English](https://www.collinsdictionary.com/dictionary/english/sort)                 | WNW 动词义 6、9 涵盖安排和排序，保留当前 verb。                                             |
| ZP216 | [sprint-general](../entries/sprint-general.yml) / sprint                   | verb           | 冲刺                                         | `/sprɪnt/`         | [sprint in American English](https://www.collinsdictionary.com/dictionary/english/sprint)             | 匹配当前词形与词义，保留所选来源完整变体。                                                  |
| ZP221 | [stamina-general](../entries/stamina-general.yml) / stamina                | noun           | 耐力；持续活动的体力；游戏中对应精力。       | `/ˈstæmənə/`       | [stamina in American English 1](https://www.collinsdictionary.com/dictionary/english/stamina)         | 匹配当前词形与词义，保留所选来源完整变体。                                                  |
| ZP224 | [steer-general](../entries/steer-general.yml) / steer                      | verb           | 控制方向；掌舵；驾驶时使车辆转向。           | `/stɪr/`           | [steer in American English 1](https://www.collinsdictionary.com/dictionary/english/steer)             | 匹配当前词形与词义，保留所选来源完整变体。                                                  |
| ZP227 | [summon-general](../entries/summon-general.yml) / summon                   | verb           | 召唤                                         | `/ˈsʌmən/`         | [summon in American English](https://www.collinsdictionary.com/dictionary/english/summon)             | 匹配当前词形与词义，保留所选来源完整变体。                                                  |
| ZP234 | [timid-general](../entries/timid-general.yml) / timid                      | adjective      | 胆小的                                       | `/ˈtɪmɪd/`         | [timid in American English](https://www.collinsdictionary.com/dictionary/english/timid)               | 匹配当前词形与词义，保留所选来源完整变体。                                                  |
| ZP237 | [traveler-general](../entries/traveler-general.yml) / traveler             | noun           | 旅行者；旅人。                               | `/ˈtrævələr/`      | [traveler in American English](https://www.collinsdictionary.com/dictionary/english/traveler)         | 匹配当前词形与词义，保留所选来源完整变体。                                                  |
| ZP245 | [underground-general](../entries/underground-general.yml) / underground    | adjective      | 地下的                                       | `/ˈʌndərˌɡraʊnd/`  | [underground in American English](https://www.collinsdictionary.com/dictionary/english/underground)   | 匹配当前词形与词义，保留所选来源完整变体。                                                  |
| ZP246 | [unlock-general](../entries/unlock-general.yml) / unlock                   | verb           | 解锁                                         | `/ʌnˈlɑk/`         | [unlock in American English](https://www.collinsdictionary.com/dictionary/english/unlock)             | 匹配当前词形与词义，保留所选来源完整变体。                                                  |
| ZP247 | [upgrade-general](../entries/upgrade-general.yml) / upgrade                | verb           | 升级、强化                                   | `/ʌpˈɡreɪd/`       | [upgrade in American English](https://www.collinsdictionary.com/dictionary/english/upgrade)           | 取来源明确标注 for v. 的 /ʌpˈɡreɪd/，不直接取先列名词读音。                                 |
| ZP249 | [view-general](../entries/view-general.yml) / view                         | noun           | 景色；视野；看到的景象。                     | `/vju/`            | [view in American English](https://www.collinsdictionary.com/dictionary/english/view)                 | 匹配当前词形与词义，保留所选来源完整变体。                                                  |
| ZP252 | [warrior-general](../entries/warrior-general.yml) / warrior                | noun           | 战士                                         | `/ˈwɔrjər/`        | [warrior in American English](https://www.collinsdictionary.com/dictionary/english/warrior)           | 匹配当前词形与词义，保留所选来源完整变体。                                                  |

### 普通英语与游戏用法

共 8 项。

| 编号  | 稳定 ID / Word                                                            | 当前词性 | 当前 Meaning                                         | 拟写 pronunciation | 来源与分区                                                                                                      | 核对说明                                                                                                                             |
|:------|:--------------------------------------------------------------------------|:---------|:-----------------------------------------------------|:-------------------|:----------------------------------------------------------------------------------------------------------------|:-------------------------------------------------------------------------------------------------------------------------------------|
| ZP025 | [build-construction](../entries/build-construction.yml) / build           | verb     | 建造；搭建建筑、装置等实体。                         | `/bɪld/`           | [build in American English](https://www.collinsdictionary.com/dictionary/english/build)                         | 匹配当前词形与词义，保留所选来源完整变体。                                                                                           |
| ZP042 | [construct-being-zelda](../entries/construct-being-zelda.yml) / Construct | noun     | 魔像；本作中由左纳乌技术造出的人工生命体或机械实体。 | `/ˈkɑnˌstrʌkt/`    | [construct in American English](https://www.collinsdictionary.com/dictionary/english/construct)                 | 将普通英语 construct 的完整名词读音用于同形游戏名称学习；保留游戏意义及 Notes，不声称已听审官方配音。名词重音不同于 construct 动词。 |
| ZP092 | [geoglyph-game](../entries/geoglyph-game.yml) / geoglyph                  | noun     | 地上绘；通过地表图案或线条形成的大型图像。           | `/ˈd͡ʒiːəɡlɪf/`     | [English, Pronunciation, General American](https://en.wiktionary.org/w/index.php?title=geoglyph&oldid=92260556) | 取 Wiktionary 明标 General American 的 /ˈd͡ʒiːəɡlɪf/，名词义为地表大型图案；社区维护来源，未把其 OED 引用冒充已核实的 OED 音标。      |
| ZP094 | [gloom-zelda](../entries/gloom-zelda.yml) / Gloom                         | noun     | 瘴气；本作中会造成特殊生命损伤的危险现象。           | `/ɡlum/`           | [gloom in American English](https://www.collinsdictionary.com/dictionary/english/gloom)                         | 将普通英语 gloom 的完整名词读音用于同形游戏名称学习；保留游戏意义及 Notes，不声称已听审官方配音。                                    |
| ZP149 | [paraglider-zelda](../entries/paraglider-zelda.yml) / Paraglider          | noun     | 滑翔帆；本作用于空中滑翔的装备。                     | `/ˈpærəˌglaɪdər/`  | [paraglider in American English](https://www.collinsdictionary.com/dictionary/english/paraglider)               | 将普通英语 paraglider 的完整名词读音用于同形游戏名称学习；保留游戏意义及 Notes，不声称已听审官方配音。                               |
| ZP151 | [pepper-vegetable](../entries/pepper-vegetable.yml) / pepper              | noun     | 辣椒；甜椒；这里不指胡椒粉。                         | `/ˈpɛpər/`         | [pepper in American English](https://www.collinsdictionary.com/dictionary/english/pepper)                       | 匹配当前词形与词义，保留所选来源完整变体。                                                                                           |
| ZP165 | [quest-game](../entries/quest-game.yml) / quest                           | noun     | 任务；追寻或探索。                                   | `/kwɛst/`          | [quest in American English](https://www.collinsdictionary.com/dictionary/english/quest)                         | 匹配当前词形与词义，保留所选来源完整变体。                                                                                           |
| ZP184 | [rupee-game](../entries/rupee-game.yml) / rupee                           | noun     | 卢比；本作的货币单位。                               | `/ˈrupi/`          | [rupee in American English](https://www.collinsdictionary.com/dictionary/english/rupee)                         | 词典提供普通英语货币名称 rupee 的读音，拟用于当前同形货币名学习；不据此声称已核实游戏内官方配音或币材质。                            |

## 逐项来源追加文案

以下每个代码块是一行完整、可审阅的 Notes 追加值。实施时只追加对应行一次，保留原 Notes 所有内容与段落；来源中的 accent: US 仍保留，尽管 pronunciation 不显示 US 前缀。围栏及标题不写入 YAML。

### ZP001 — activate

```plain
Pronunciation source: Collins Dictionary — activate; activate in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/activate (accent: US; accessed 2026-11-18).
```

### ZP002 — adventure

```plain
Pronunciation source: Collins Dictionary — adventure; adventure in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/adventure (accent: US; accessed 2026-11-18).
```

### ZP003 — aggressive

```plain
Pronunciation source: Collins Dictionary — aggressive; aggressive in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/aggressive (accent: US; accessed 2026-11-18).
```

### ZP004 — aim

```plain
Pronunciation source: Collins Dictionary — aim; aim in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/aim (accent: US; accessed 2026-11-18).
```

### ZP005 — airborne

```plain
Pronunciation source: Collins Dictionary — airborne; airborne in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/airborne (accent: US; accessed 2026-11-18).
```

### ZP006 — album

```plain
Pronunciation source: Collins Dictionary — album; album in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/album (accent: US; accessed 2026-11-18).
```

### ZP007 — altar

```plain
Pronunciation source: Collins Dictionary — altar; altar in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/altar (accent: US; accessed 2026-11-18).
```

### ZP008 — ambitious

```plain
Pronunciation source: Collins Dictionary — ambitious; ambitious in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/ambitious (accent: US; accessed 2026-11-18).
```

### ZP009 — ancestor

```plain
Pronunciation source: Collins Dictionary — ancestor; ancestor in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/ancestor (accent: US; accessed 2026-11-18).
```

### ZP010 — ancient

```plain
Pronunciation source: Collins Dictionary — ancient; ancient in American English 1; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/ancient (accent: US; accessed 2026-11-18).
```

### ZP011 — aquatic

```plain
Pronunciation source: Collins Dictionary — aquatic; aquatic in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/aquatic (accent: US; accessed 2026-11-18).
```

### ZP012 — archipelago

```plain
Pronunciation source: Collins Dictionary — archipelago; archipelago in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/archipelago (accent: US; accessed 2026-11-18).
```

### ZP013 — armor

```plain
Pronunciation source: Collins Dictionary — armor; armor in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/armor (accent: US; accessed 2026-11-18).
```

### ZP014 — armored

```plain
Pronunciation source: Collins Dictionary — armored; armored in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/armored (accent: US; accessed 2026-11-18).
```

### ZP015 — armory

```plain
Pronunciation source: Collins Dictionary — armory; armory in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/armory (accent: US; accessed 2026-11-18).
```

### ZP016 — awaken

```plain
Pronunciation source: Collins Dictionary — awaken; awaken in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/awaken (accent: US; accessed 2026-11-18).
```

### ZP017 — bandanna

```plain
Pronunciation source: Collins Dictionary — bandanna; bandanna in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/bandanna (accent: US; accessed 2026-11-18).
```

### ZP018 — battery

```plain
Pronunciation source: Collins Dictionary — battery; battery in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/battery (accent: US; accessed 2026-11-18).
```

### ZP019 — beam

```plain
Pronunciation source: Collins Dictionary — beam; beam in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/beam (accent: US; accessed 2026-11-18).
```

### ZP020 — blessing

```plain
Pronunciation source: Collins Dictionary — blessing; blessing in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/blessing (accent: US; accessed 2026-11-18).
```

### ZP021 — block

```plain
Pronunciation source: Collins Dictionary — block; block in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/block (accent: US; accessed 2026-11-18).
```

### ZP022 — bomb

```plain
Pronunciation source: Collins Dictionary — bomb; bomb in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/bomb (accent: US; accessed 2026-11-18).
```

### ZP023 — boomerang

```plain
Pronunciation source: Collins Dictionary — boomerang; boomerang in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/boomerang (accent: US; accessed 2026-11-18).
```

### ZP024 — broadsword

```plain
Pronunciation source: Collins Dictionary — broadsword; broadsword in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/broadsword (accent: US; accessed 2026-11-18).
```

### ZP025 — build

```plain
Pronunciation source: Collins Dictionary — build; build in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/build (accent: US; accessed 2026-11-18).
```

### ZP026 — calamity

```plain
Pronunciation source: Collins Dictionary — calamity; calamity in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/calamity (accent: US; accessed 2026-11-18).
```

### ZP027 — cannon

```plain
Pronunciation source: Collins Dictionary — cannon; cannon in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/cannon (accent: US; accessed 2026-11-18).
```

### ZP028 — canyon

```plain
Pronunciation source: Collins Dictionary — canyon; canyon in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/canyon (accent: US; accessed 2026-11-18).
```

### ZP029 — cart

```plain
Pronunciation source: Collins Dictionary — cart; cart in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/cart (accent: US; accessed 2026-11-18).
```

### ZP030 — cavern

```plain
Pronunciation source: Collins Dictionary — cavern; cavern in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/cavern (accent: US; accessed 2026-11-18).
```

### ZP031 — chamber

```plain
Pronunciation source: Collins Dictionary — chamber; chamber in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/chamber (accent: US; accessed 2026-11-18).
```

### ZP032 — champion

```plain
Pronunciation source: Collins Dictionary — champion; champion in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/champion (accent: US; accessed 2026-11-18).
```

### ZP033 — chasm

```plain
Pronunciation source: Collins Dictionary — chasm; chasm in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/chasm (accent: US; accessed 2026-11-18).
```

### ZP034 — citadel

```plain
Pronunciation source: Collins Dictionary — citadel; citadel in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/citadel (accent: US; accessed 2026-11-18).
```

### ZP035 — civilization

```plain
Pronunciation source: Collins Dictionary — civilization; civilization in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/civilization (accent: US; accessed 2026-11-18).
```

### ZP036 — claymore

```plain
Pronunciation source: Collins Dictionary — claymore; claymore in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/claymore (accent: US; accessed 2026-11-18).
```

### ZP037 — closed

```plain
Pronunciation source: Collins Dictionary — closed; closed in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/closed (accent: US; accessed 2026-11-18).
```

### ZP038 — club

```plain
Pronunciation source: Collins Dictionary — club; club in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/club (accent: US; accessed 2026-11-18).
```

### ZP039 — clue

```plain
Pronunciation source: Collins Dictionary — clue; clue in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/clue (accent: US; accessed 2026-11-18).
```

### ZP040 — colossal

```plain
Pronunciation source: Collins Dictionary — colossal; colossal in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/colossal (accent: US; accessed 2026-11-18).
```

### ZP041 — compendium

```plain
Pronunciation source: Collins Dictionary — compendium; compendium in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/compendium (accent: US; accessed 2026-11-18).
```

### ZP042 — Construct

```plain
Pronunciation source: Collins Dictionary — Construct; construct in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/construct (accent: US; accessed 2026-11-18; ordinary English noun pronunciation, not a verified game-voice transcription).
```

### ZP043 — construct

```plain
Pronunciation source: Collins Dictionary — construct; construct in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/construct (accent: US; accessed 2026-11-18).
```

### ZP044 — cook

```plain
Pronunciation source: Collins Dictionary — cook; cook in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/cook (accent: US; accessed 2026-11-18).
```

### ZP045 — corridor

```plain
Pronunciation source: Collins Dictionary — corridor; corridor in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/corridor (accent: US; accessed 2026-11-18).
```

### ZP046 — courageous

```plain
Pronunciation source: Collins Dictionary — courageous; courageous in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/courageous (accent: US; accessed 2026-11-18).
```

### ZP047 — crater

```plain
Pronunciation source: Collins Dictionary — crater; crater in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/crater (accent: US; accessed 2026-11-18).
```

### ZP048 — crisis

```plain
Pronunciation source: Collins Dictionary — crisis; crisis in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/crisis (accent: US; accessed 2026-11-18).
```

### ZP049 — crouch

```plain
Pronunciation source: Collins Dictionary — crouch; crouch in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/crouch (accent: US; accessed 2026-11-18).
```

### ZP050 — crystal

```plain
Pronunciation source: Collins Dictionary — crystal; crystal in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/crystal (accent: US; accessed 2026-11-18).
```

### ZP051 — cunning

```plain
Pronunciation source: Collins Dictionary — cunning; cunning in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/cunning (accent: US; accessed 2026-11-18).
```

### ZP052 — curry

```plain
Pronunciation source: Collins Dictionary — curry; curry in American English 2; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/curry (accent: US; accessed 2026-11-18).
```

### ZP053 — dagger

```plain
Pronunciation source: Collins Dictionary — dagger; dagger in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/dagger (accent: US; accessed 2026-11-18).
```

### ZP054 — damage

```plain
Pronunciation source: Collins Dictionary — damage; damage in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/damage (accent: US; accessed 2026-11-18).
```

### ZP055 — defense

```plain
Pronunciation source: Collins Dictionary — defense; defense in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/defense (accent: US; accessed 2026-11-18).
```

### ZP056 — deity

```plain
Pronunciation source: Collins Dictionary — deity; deity in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/deity (accent: US; accessed 2026-11-18).
```

### ZP057 — demon

```plain
Pronunciation source: Collins Dictionary — demon; demon in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/demon (accent: US; accessed 2026-11-18).
```

### ZP058 — descendant

```plain
Pronunciation source: Collins Dictionary — descendant; descendant in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/descendant (accent: US; accessed 2026-11-18).
```

### ZP059 — desert

```plain
Pronunciation source: Collins Dictionary — desert; desert in American English 2; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/desert (accent: US; accessed 2026-11-18).
```

### ZP060 — destination

```plain
Pronunciation source: Collins Dictionary — destination; destination in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/destination (accent: US; accessed 2026-11-18).
```

### ZP061 — determined

```plain
Pronunciation source: Collins Dictionary — determined; determined in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/determined (accent: US; accessed 2026-11-18).
```

### ZP062 — device

```plain
Pronunciation source: Collins Dictionary — device; device in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/device (accent: US; accessed 2026-11-18).
```

### ZP063 — diligent

```plain
Pronunciation source: Collins Dictionary — diligent; diligent in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/diligent (accent: US; accessed 2026-11-18).
```

### ZP064 — dispenser

```plain
Pronunciation source: Collins Dictionary — dispenser; dispenser in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/dispenser (accent: US; accessed 2026-11-18).
```

### ZP065 — distinctive

```plain
Pronunciation source: Collins Dictionary — distinctive; distinctive in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/distinctive (accent: US; accessed 2026-11-18).
```

### ZP066 — dive

```plain
Pronunciation source: Collins Dictionary — dive; dive in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/dive (accent: US; accessed 2026-11-18).
```

### ZP067 — dodge

```plain
Pronunciation source: Collins Dictionary — dodge; dodge in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/dodge (accent: US; accessed 2026-11-18).
```

### ZP068 — dragon

```plain
Pronunciation source: Collins Dictionary — dragon; dragon in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/dragon (accent: US; accessed 2026-11-18).
```

### ZP069 — draw

```plain
Pronunciation source: Collins Dictionary — draw; draw in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/draw (accent: US; accessed 2026-11-18).
```

### ZP070 — elixir

```plain
Pronunciation source: Collins Dictionary — elixir; elixir in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/elixir (accent: US; accessed 2026-11-18).
```

### ZP071 — elusive

```plain
Pronunciation source: Collins Dictionary — elusive; elusive in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/elusive (accent: US; accessed 2026-11-18).
```

### ZP072 — emitter

```plain
Pronunciation source: Collins Dictionary — emitter; emitter in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/emitter (accent: US; accessed 2026-11-18).
```

### ZP073 — endure

```plain
Pronunciation source: Collins Dictionary — endure; endure in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/endure (accent: US; accessed 2026-11-18).
```

### ZP074 — energize

```plain
Pronunciation source: Collins Dictionary — energize; energize in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/energize (accent: US; accessed 2026-11-18).
```

### ZP075 — enormous

```plain
Pronunciation source: Collins Dictionary — enormous; enormous in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/enormous (accent: US; accessed 2026-11-18).
```

### ZP076 — enthusiastic

```plain
Pronunciation source: Collins Dictionary — enthusiastic; enthusiastic in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/enthusiastic (accent: US; accessed 2026-11-18).
```

### ZP077 — entrance

```plain
Pronunciation source: Collins Dictionary — entrance; entrance in American English 1; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/entrance (accent: US; accessed 2026-11-18).
```

### ZP078 — equip

```plain
Pronunciation source: Collins Dictionary — equip; equip in American English 1; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/equip (accent: US; accessed 2026-11-18).
```

### ZP079 — equipment

```plain
Pronunciation source: Collins Dictionary — equipment; equipment in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/equipment (accent: US; accessed 2026-11-18).
```

### ZP080 — exit

```plain
Pronunciation source: Collins Dictionary — exit; exit in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/exit (accent: US; accessed 2026-11-18).
```

### ZP081 — fearless

```plain
Pronunciation source: Collins Dictionary — fearless; fearless in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/fearless (accent: US; accessed 2026-11-18).
```

### ZP082 — fierce

```plain
Pronunciation source: Collins Dictionary — fierce; fierce in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/fierce (accent: US; accessed 2026-11-18).
```

### ZP083 — fifth

```plain
Pronunciation source: Collins Dictionary — fifth; fifth in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/fifth (accent: US; accessed 2026-11-18).
```

### ZP084 — foothill

```plain
Pronunciation source: Collins Dictionary — foothill; foothill in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/foothill (accent: US; accessed 2026-11-18).
```

### ZP085 — forge

```plain
Pronunciation source: Collins Dictionary — forge; forge in American English 1; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/forge (accent: US; accessed 2026-11-18).
```

### ZP086 — fortress

```plain
Pronunciation source: Collins Dictionary — fortress; fortress in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/fortress (accent: US; accessed 2026-11-18).
```

### ZP087 — fountain

```plain
Pronunciation source: Collins Dictionary — fountain; fountain in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/fountain (accent: US; accessed 2026-11-18).
```

### ZP088 — frost

```plain
Pronunciation source: Collins Dictionary — frost; frost in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/frost (accent: US; accessed 2026-11-18).
```

### ZP089 — fry

```plain
Pronunciation source: Collins Dictionary — fry; fry in American English 1; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/fry (accent: US; accessed 2026-11-18).
```

### ZP090 — fuse

```plain
Pronunciation source: Collins Dictionary — fuse; fuse in American English 1; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/fuse (accent: US; accessed 2026-11-18).
```

### ZP091 — gear

```plain
Pronunciation source: Collins Dictionary — gear; gear in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/gear (accent: US; accessed 2026-11-18).
```

### ZP092 — geoglyph

```plain
Pronunciation source: Wiktionary — geoglyph; English, Pronunciation, General American; Wiktionary. https://en.wiktionary.org/w/index.php?title=geoglyph&oldid=92260556 (accent: US; accessed 2026-11-18).
```

### ZP093 — glide

```plain
Pronunciation source: Collins Dictionary — glide; glide in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/glide (accent: US; accessed 2026-11-18).
```

### ZP094 — Gloom

```plain
Pronunciation source: Collins Dictionary — Gloom; gloom in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/gloom (accent: US; accessed 2026-11-18; ordinary English noun pronunciation, not a verified game-voice transcription).
```

### ZP095 — gourmet

```plain
Pronunciation source: Collins Dictionary — gourmet; gourmet in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/gourmet (accent: US; accessed 2026-11-18).
```

### ZP096 — grill

```plain
Pronunciation source: Collins Dictionary — grill; grill in American English 1; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/grill (accent: US; accessed 2026-11-18).
```

### ZP097 — guardian

```plain
Pronunciation source: Collins Dictionary — guardian; guardian in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/guardian (accent: US; accessed 2026-11-18).
```

### ZP098 — guidance

```plain
Pronunciation source: Collins Dictionary — guidance; guidance in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/guidance (accent: US; accessed 2026-11-18).
```

### ZP099 — halberd

```plain
Pronunciation source: Collins Dictionary — halberd; halberd in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/halberd (accent: US; accessed 2026-11-18).
```

### ZP100 — hammer

```plain
Pronunciation source: Collins Dictionary — hammer; hammer in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/hammer (accent: US; accessed 2026-11-18).
```

### ZP101 — hasty

```plain
Pronunciation source: Collins Dictionary — hasty; hasty in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/hasty (accent: US; accessed 2026-11-18).
```

### ZP102 — headband

```plain
Pronunciation source: Collins Dictionary — headband; headband in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/headband (accent: US; accessed 2026-11-18).
```

### ZP103 — headdress

```plain
Pronunciation source: Collins Dictionary — headdress; headdress in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/headdress (accent: US; accessed 2026-11-18).
```

### ZP104 — hearty

```plain
Pronunciation source: Collins Dictionary — hearty; hearty in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/hearty (accent: US; accessed 2026-11-18).
```

### ZP105 — heir

```plain
Pronunciation source: Collins Dictionary — heir; heir in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/heir (accent: US; accessed 2026-11-18).
```

### ZP106 — helm

```plain
Pronunciation source: Collins Dictionary — helm; helm in American English 1; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/helm (accent: US; accessed 2026-11-18).
```

### ZP107 — helmet

```plain
Pronunciation source: Collins Dictionary — helmet; helmet in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/helmet (accent: US; accessed 2026-11-18).
```

### ZP108 — hero

```plain
Pronunciation source: Collins Dictionary — hero; hero in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/hero (accent: US; accessed 2026-11-18).
```

### ZP109 — hideout

```plain
Pronunciation source: Collins Dictionary — hideout; hideout in British English; Collins English Dictionary. https://www.collinsdictionary.com/dictionary/english/hideout (accent: UK; accessed 2026-11-18).
```

### ZP110 — homing

```plain
Pronunciation source: Collins Dictionary — homing; homing in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/homing (accent: US; accessed 2026-11-18).
```

### ZP111 — honorable

```plain
Pronunciation source: Collins Dictionary — honorable; honorable in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/honorable (accent: US; accessed 2026-11-18).
```

### ZP112 — hood

```plain
Pronunciation source: Collins Dictionary — hood; hood in American English 1; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/hood (accent: US; accessed 2026-11-18).
```

### ZP113 — hostile

```plain
Pronunciation source: Collins Dictionary — hostile; hostile in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/hostile (accent: US; accessed 2026-11-18).
```

### ZP114 — hover

```plain
Pronunciation source: Collins Dictionary — hover; hover in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/hover (accent: US; accessed 2026-11-18).
```

### ZP115 — hydrant

```plain
Pronunciation source: Collins Dictionary — hydrant; hydrant in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/hydrant (accent: US; accessed 2026-11-18).
```

### ZP116 — imprison

```plain
Pronunciation source: Collins Dictionary — imprison; imprison in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/imprison (accent: US; accessed 2026-11-18).
```

### ZP117 — ingredient

```plain
Pronunciation source: Collins Dictionary — ingredient; ingredient in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/ingredient (accent: US; accessed 2026-11-18).
```

### ZP118 — inherit

```plain
Pronunciation source: Collins Dictionary — inherit; inherit in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/inherit (accent: US; accessed 2026-11-18).
```

### ZP119 — inn

```plain
Pronunciation source: Collins Dictionary — inn; inn in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/inn (accent: US; accessed 2026-11-18).
```

### ZP120 — integrity

```plain
Pronunciation source: Collins Dictionary — integrity; integrity in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/integrity (accent: US; accessed 2026-11-18).
```

### ZP121 — interact

```plain
Pronunciation source: Collins Dictionary — interact; interact in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/interact (accent: US; accessed 2026-11-18).
```

### ZP122 — inventory

```plain
Pronunciation source: Collins Dictionary — inventory; inventory in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/inventory (accent: US; accessed 2026-11-18).
```

### ZP123 — king

```plain
Pronunciation source: Collins Dictionary — king; king in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/king (accent: US; accessed 2026-11-18).
```

### ZP124 — kingdom

```plain
Pronunciation source: Collins Dictionary — kingdom; kingdom in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/kingdom (accent: US; accessed 2026-11-18).
```

### ZP125 — knight

```plain
Pronunciation source: Collins Dictionary — knight; knight in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/knight (accent: US; accessed 2026-11-18).
```

### ZP126 — landing

```plain
Pronunciation source: Collins Dictionary — landing; landing in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/landing (accent: US; accessed 2026-11-18).
```

### ZP127 — lava

```plain
Pronunciation source: Collins Dictionary — lava; lava in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/lava (accent: US; accessed 2026-11-18).
```

### ZP128 — light

```plain
Pronunciation source: Collins Dictionary — light; light in American English 1; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/light (accent: US; accessed 2026-11-18).
```

### ZP129 — lightning

```plain
Pronunciation source: Collins Dictionary — lightning; lightning in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/lightning (accent: US; accessed 2026-11-18).
```

### ZP130 — lookout

```plain
Pronunciation source: Collins Dictionary — lookout; lookout in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/lookout (accent: US; accessed 2026-11-18).
```

### ZP131 — majestic

```plain
Pronunciation source: Collins Dictionary — majestic; majestic in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/majestic (accent: US; accessed 2026-11-18).
```

### ZP132 — malevolent

```plain
Pronunciation source: Collins Dictionary — malevolent; malevolent in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/malevolent (accent: US; accessed 2026-11-18).
```

### ZP133 — mask

```plain
Pronunciation source: Collins Dictionary — mask; mask in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/mask (accent: US; accessed 2026-11-18).
```

### ZP134 — material

```plain
Pronunciation source: Collins Dictionary — material; material in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/material (accent: US; accessed 2026-11-18).
```

### ZP135 — meal

```plain
Pronunciation source: Collins Dictionary — meal; meal in American English 1; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/meal (accent: US; accessed 2026-11-18).
```

### ZP136 — melee

```plain
Pronunciation source: Collins Dictionary — melee; melee in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/melee (accent: US; accessed 2026-11-18).
```

### ZP137 — memory

```plain
Pronunciation source: Collins Dictionary — memory; memory in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/memory (accent: US; accessed 2026-11-18).
```

### ZP138 — mighty

```plain
Pronunciation source: Collins Dictionary — mighty; mighty in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/mighty (accent: US; accessed 2026-11-18).
```

### ZP139 — mine

```plain
Pronunciation source: Collins Dictionary — mine; mine in American English 2; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/mine (accent: US; accessed 2026-11-18).
```

### ZP140 — miner

```plain
Pronunciation source: Collins Dictionary — miner; miner in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/miner (accent: US; accessed 2026-11-18).
```

### ZP141 — mysterious

```plain
Pronunciation source: Collins Dictionary — mysterious; mysterious in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/mysterious (accent: US; accessed 2026-11-18).
```

### ZP142 — mystery

```plain
Pronunciation source: Collins Dictionary — mystery; mystery in American English 1; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/mystery (accent: US; accessed 2026-11-18).
```

### ZP143 — noble

```plain
Pronunciation source: Collins Dictionary — noble; noble in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/noble (accent: US; accessed 2026-11-18).
```

### ZP144 — objective

```plain
Pronunciation source: Collins Dictionary — objective; objective in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/objective (accent: US; accessed 2026-11-18).
```

### ZP145 — omelet

```plain
Pronunciation source: Collins Dictionary — omelet; omelet in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/omelet (accent: US; accessed 2026-11-18).
```

### ZP146 — one-handed

```plain
Pronunciation source: Collins Dictionary — one-handed; one-handed in American English; American English reference material credited to Penguin Random House LLC and HarperCollins Publishers Ltd. https://www.collinsdictionary.com/dictionary/english/one-handed (accent: US; accessed 2026-11-18).
```

### ZP147 — pad

```plain
Pronunciation source: Collins Dictionary — pad; pad in American English 2; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/pad (accent: US; accessed 2026-11-18).
```

### ZP148 — paella

```plain
Pronunciation source: Collins Dictionary — paella; paella in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/paella (accent: US; accessed 2026-11-18).
```

### ZP149 — Paraglider

```plain
Pronunciation source: Collins Dictionary — Paraglider; paraglider in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/paraglider (accent: US; accessed 2026-11-18; ordinary English noun pronunciation, not a verified game-voice transcription).
```

### ZP150 — passage

```plain
Pronunciation source: Collins Dictionary — passage; passage in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/passage (accent: US; accessed 2026-11-18).
```

### ZP151 — pepper

```plain
Pronunciation source: Collins Dictionary — pepper; pepper in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/pepper (accent: US; accessed 2026-11-18).
```

### ZP152 — phenomenon

```plain
Pronunciation source: Collins Dictionary — phenomenon; phenomenon in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/phenomenon (accent: US; accessed 2026-11-18).
```

### ZP153 — plateau

```plain
Pronunciation source: Collins Dictionary — plateau; plateau in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/plateau (accent: US; accessed 2026-11-18).
```

### ZP154 — platform

```plain
Pronunciation source: Collins Dictionary — platform; platform in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/platform (accent: US; accessed 2026-11-18).
```

### ZP155 — porridge

```plain
Pronunciation source: Collins Dictionary — porridge; porridge in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/porridge (accent: US; accessed 2026-11-18).
```

### ZP156 — possessed

```plain
Pronunciation source: Collins Dictionary — possessed; possessed in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/possessed (accent: US; accessed 2026-11-18).
```

### ZP157 — pouch

```plain
Pronunciation source: Collins Dictionary — pouch; pouch in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/pouch (accent: US; accessed 2026-11-18).
```

### ZP158 — prime

```plain
Pronunciation source: Collins Dictionary — prime; prime in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/prime (accent: US; accessed 2026-11-18).
```

### ZP159 — prince

```plain
Pronunciation source: Collins Dictionary — prince; prince in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/prince (accent: US; accessed 2026-11-18).
```

### ZP160 — princess

```plain
Pronunciation source: Collins Dictionary — princess; princess in American English 1; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/princess (accent: US; accessed 2026-11-18).
```

### ZP161 — pristine

```plain
Pronunciation source: Collins Dictionary — pristine; pristine in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/pristine (accent: US; accessed 2026-11-18).
```

### ZP162 — profile

```plain
Pronunciation source: Collins Dictionary — profile; profile in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/profile (accent: US; accessed 2026-11-18).
```

### ZP163 — prophecy

```plain
Pronunciation source: Collins Dictionary — prophecy; prophecy in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/prophecy (accent: US; accessed 2026-11-18).
```

### ZP164 — queen

```plain
Pronunciation source: Collins Dictionary — queen; queen in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/queen (accent: US; accessed 2026-11-18).
```

### ZP165 — quest

```plain
Pronunciation source: Collins Dictionary — quest; quest in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/quest (accent: US; accessed 2026-11-18).
```

### ZP166 — radish

```plain
Pronunciation source: Collins Dictionary — radish; radish in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/radish (accent: US; accessed 2026-11-18).
```

### ZP167 — recall

```plain
Pronunciation source: Collins Dictionary — recall; recall in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/recall (accent: US; accessed 2026-11-18).
```

### ZP168 — recharge

```plain
Pronunciation source: Collins Dictionary — recharge; recharge in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/recharge (accent: US; accessed 2026-11-18).
```

### ZP169 — recipe

```plain
Pronunciation source: Collins Dictionary — recipe; recipe in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/recipe (accent: US; accessed 2026-11-18).
```

### ZP170 — reckless

```plain
Pronunciation source: Collins Dictionary — reckless; reckless in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/reckless (accent: US; accessed 2026-11-18).
```

### ZP171 — recover

```plain
Pronunciation source: Collins Dictionary — recover; recover in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/recover (accent: US; accessed 2026-11-18).
```

### ZP172 — refinery

```plain
Pronunciation source: Collins Dictionary — refinery; refinery in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/refinery (accent: US; accessed 2026-11-18).
```

### ZP173 — regional

```plain
Pronunciation source: Collins Dictionary — regional; regional in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/regional (accent: US; accessed 2026-11-18).
```

### ZP174 — resilient

```plain
Pronunciation source: Collins Dictionary — resilient; resilient in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/resilient (accent: US; accessed 2026-11-18).
```

### ZP175 — resistance

```plain
Pronunciation source: Collins Dictionary — resistance; resistance in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/resistance (accent: US; accessed 2026-11-18).
```

### ZP176 — ridge

```plain
Pronunciation source: Collins Dictionary — ridge; ridge in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/ridge (accent: US; accessed 2026-11-18).
```

### ZP177 — risotto

```plain
Pronunciation source: Collins Dictionary — risotto; risotto in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/risotto (accent: US; accessed 2026-11-18).
```

### ZP178 — roast

```plain
Pronunciation source: Collins Dictionary — roast; roast in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/roast (accent: US; accessed 2026-11-18).
```

### ZP179 — robe

```plain
Pronunciation source: Collins Dictionary — robe; robe in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/robe (accent: US; accessed 2026-11-18).
```

### ZP180 — rocket

```plain
Pronunciation source: Collins Dictionary — rocket; rocket in American English 1; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/rocket (accent: US; accessed 2026-11-18).
```

### ZP181 — rod

```plain
Pronunciation source: Collins Dictionary — rod; rod in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/rod (accent: US; accessed 2026-11-18).
```

### ZP182 — rotate

```plain
Pronunciation source: Collins Dictionary — rotate; rotate in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/rotate (accent: US; accessed 2026-11-18).
```

### ZP183 — royal

```plain
Pronunciation source: Collins Dictionary — royal; royal in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/royal (accent: US; accessed 2026-11-18).
```

### ZP184 — rupee

```plain
Pronunciation source: Collins Dictionary — rupee; rupee in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/rupee (accent: US; accessed 2026-11-18).
```

### ZP185 — ruthless

```plain
Pronunciation source: Collins Dictionary — ruthless; ruthless in American English; American English reference material credited to Penguin Random House LLC and HarperCollins Publishers Ltd. https://www.collinsdictionary.com/dictionary/english/ruthless (accent: US; accessed 2026-11-18).
```

### ZP186 — sacred

```plain
Pronunciation source: Collins Dictionary — sacred; sacred in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/sacred (accent: US; accessed 2026-11-18).
```

### ZP187 — sacrifice

```plain
Pronunciation source: Collins Dictionary — sacrifice; sacrifice in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/sacrifice (accent: US; accessed 2026-11-18).
```

### ZP188 — sage

```plain
Pronunciation source: Collins Dictionary — sage; sage in American English 1; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/sage (accent: US; accessed 2026-11-18).
```

### ZP189 — salmon

```plain
Pronunciation source: Collins Dictionary — salmon; salmon in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/salmon (accent: US; accessed 2026-11-18).
```

### ZP190 — savage

```plain
Pronunciation source: Collins Dictionary — savage; savage in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/savage (accent: US; accessed 2026-11-18).
```

### ZP191 — scimitar

```plain
Pronunciation source: Collins Dictionary — scimitar; scimitar in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/scimitar (accent: US; accessed 2026-11-18).
```

### ZP192 — seafood

```plain
Pronunciation source: Collins Dictionary — seafood; seafood in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/seafood (accent: US; accessed 2026-11-18).
```

### ZP193 — sear

```plain
Pronunciation source: Collins Dictionary — sear; sear in American English 1; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/sear (accent: US; accessed 2026-11-18).
```

### ZP194 — secret

```plain
Pronunciation source: Collins Dictionary — secret; secret in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/secret (accent: US; accessed 2026-11-18).
```

### ZP195 — selfless

```plain
Pronunciation source: Collins Dictionary — selfless; selfless in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/selfless (accent: US; accessed 2026-11-18).
```

### ZP196 — settlement

```plain
Pronunciation source: Collins Dictionary — settlement; settlement in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/settlement (accent: US; accessed 2026-11-18).
```

### ZP197 — shirt

```plain
Pronunciation source: Collins Dictionary — shirt; shirt in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/shirt (accent: US; accessed 2026-11-18).
```

### ZP198 — shrine

```plain
Pronunciation source: Collins Dictionary — shrine; shrine in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/shrine (accent: US; accessed 2026-11-18).
```

### ZP199 — simmer

```plain
Pronunciation source: Collins Dictionary — simmer; simmer in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/simmer (accent: US; accessed 2026-11-18).
```

### ZP200 — skewer

```plain
Pronunciation source: Collins Dictionary — skewer; skewer in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/skewer (accent: US; accessed 2026-11-18).
```

### ZP201 — skilled

```plain
Pronunciation source: Collins Dictionary — skilled; skilled in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/skilled (accent: US; accessed 2026-11-18).
```

### ZP202 — sky

```plain
Pronunciation source: Collins Dictionary — sky; sky in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/sky (accent: US; accessed 2026-11-18).
```

### ZP203 — sled

```plain
Pronunciation source: Collins Dictionary — sled; sled in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/sled (accent: US; accessed 2026-11-18).
```

### ZP204 — sleeve

```plain
Pronunciation source: Collins Dictionary — sleeve; sleeve in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/sleeve (accent: US; accessed 2026-11-18).
```

### ZP205 — slip

```plain
Pronunciation source: Collins Dictionary — slip; slip in American English 1; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/slip (accent: US; accessed 2026-11-18).
```

### ZP206 — slope

```plain
Pronunciation source: Collins Dictionary — slope; slope in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/slope (accent: US; accessed 2026-11-18).
```

### ZP207 — sneaky

```plain
Pronunciation source: Collins Dictionary — sneaky; sneaky in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/sneaky (accent: US; accessed 2026-11-18).
```

### ZP208 — soldier

```plain
Pronunciation source: Collins Dictionary — soldier; soldier in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/soldier (accent: US; accessed 2026-11-18).
```

### ZP209 — sort

```plain
Pronunciation source: Collins Dictionary — sort; sort in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/sort (accent: US; accessed 2026-11-18).
```

### ZP210 — spaulder

```plain
Pronunciation source: Collins Dictionary — spaulder; spaulder in American English; American English reference material credited to Penguin Random House LLC and HarperCollins Publishers Ltd. https://www.collinsdictionary.com/dictionary/english/spaulder (accent: US; accessed 2026-11-18).
```

### ZP211 — spear

```plain
Pronunciation source: Collins Dictionary — spear; spear in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/spear (accent: US; accessed 2026-11-18).
```

### ZP212 — spice

```plain
Pronunciation source: Collins Dictionary — spice; spice in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/spice (accent: US; accessed 2026-11-18).
```

### ZP213 — spirit

```plain
Pronunciation source: Collins Dictionary — spirit; spirit in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/spirit (accent: US; accessed 2026-11-18).
```

### ZP214 — spring

```plain
Pronunciation source: Collins Dictionary — spring; spring in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/spring (accent: US; accessed 2026-11-18).
```

### ZP215 — spring

```plain
Pronunciation source: Collins Dictionary — spring; spring in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/spring (accent: US; accessed 2026-11-18).
```

### ZP216 — sprint

```plain
Pronunciation source: Collins Dictionary — sprint; sprint in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/sprint (accent: US; accessed 2026-11-18).
```

### ZP217 — stabilizer

```plain
Pronunciation source: Collins Dictionary — stabilizer; stabilizer in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/stabilizer (accent: US; accessed 2026-11-18).
```

### ZP218 — stable

```plain
Pronunciation source: Collins Dictionary — stable; stable in American English 2; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/stable (accent: US; accessed 2026-11-18).
```

### ZP219 — staircase

```plain
Pronunciation source: Collins Dictionary — staircase; staircase in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/staircase (accent: US; accessed 2026-11-18).
```

### ZP220 — stake

```plain
Pronunciation source: Collins Dictionary — stake; stake in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/stake (accent: US; accessed 2026-11-18).
```

### ZP221 — stamina

```plain
Pronunciation source: Collins Dictionary — stamina; stamina in American English 1; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/stamina (accent: US; accessed 2026-11-18).
```

### ZP222 — statue

```plain
Pronunciation source: Collins Dictionary — statue; statue in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/statue (accent: US; accessed 2026-11-18).
```

### ZP223 — steam

```plain
Pronunciation source: Collins Dictionary — steam; steam in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/steam (accent: US; accessed 2026-11-18).
```

### ZP224 — steer

```plain
Pronunciation source: Collins Dictionary — steer; steer in American English 1; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/steer (accent: US; accessed 2026-11-18).
```

### ZP225 — stew

```plain
Pronunciation source: Collins Dictionary — stew; stew in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/stew (accent: US; accessed 2026-11-18).
```

### ZP226 — summit

```plain
Pronunciation source: Collins Dictionary — summit; summit in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/summit (accent: US; accessed 2026-11-18).
```

### ZP227 — summon

```plain
Pronunciation source: Collins Dictionary — summon; summon in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/summon (accent: US; accessed 2026-11-18).
```

### ZP228 — surface

```plain
Pronunciation source: Collins Dictionary — surface; surface in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/surface (accent: US; accessed 2026-11-18).
```

### ZP229 — susceptible

```plain
Pronunciation source: Collins Dictionary — susceptible; susceptible in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/susceptible (accent: US; accessed 2026-11-18).
```

### ZP230 — swamp

```plain
Pronunciation source: Collins Dictionary — swamp; swamp in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/swamp (accent: US; accessed 2026-11-18).
```

### ZP231 — system

```plain
Pronunciation source: Collins Dictionary — system; system in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/system (accent: US; accessed 2026-11-18).
```

### ZP232 — temple

```plain
Pronunciation source: Collins Dictionary — temple; temple in American English 1; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/temple (accent: US; accessed 2026-11-18).
```

### ZP233 — tights

```plain
Pronunciation source: Collins Dictionary — tights; tights in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/tights (accent: US; accessed 2026-11-18).
```

### ZP234 — timid

```plain
Pronunciation source: Collins Dictionary — timid; timid in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/timid (accent: US; accessed 2026-11-18).
```

### ZP235 — torch

```plain
Pronunciation source: Collins Dictionary — torch; torch in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/torch (accent: US; accessed 2026-11-18).
```

### ZP236 — trail

```plain
Pronunciation source: Collins Dictionary — trail; trail in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/trail (accent: US; accessed 2026-11-18).
```

### ZP237 — traveler

```plain
Pronunciation source: Collins Dictionary — traveler; traveler in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/traveler (accent: US; accessed 2026-11-18).
```

### ZP238 — tree

```plain
Pronunciation source: Collins Dictionary — tree; tree in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/tree (accent: US; accessed 2026-11-18).
```

### ZP239 — trident

```plain
Pronunciation source: Collins Dictionary — trident; trident in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/trident (accent: US; accessed 2026-11-18).
```

### ZP240 — trousers

```plain
Pronunciation source: Collins Dictionary — trousers; trousers in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/trousers (accent: US; accessed 2026-11-18).
```

### ZP241 — truffle

```plain
Pronunciation source: Collins Dictionary — truffle; truffle in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/truffle (accent: US; accessed 2026-11-18).
```

### ZP242 — tunic

```plain
Pronunciation source: Collins Dictionary — tunic; tunic in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/tunic (accent: US; accessed 2026-11-18).
```

### ZP243 — two-handed

```plain
Pronunciation source: Collins Dictionary — two-handed; two-handed in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/two-handed (accent: US; accessed 2026-11-18).
```

### ZP244 — undead

```plain
Pronunciation source: Collins Dictionary — undead; undead in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/undead (accent: US; accessed 2026-11-18).
```

### ZP245 — underground

```plain
Pronunciation source: Collins Dictionary — underground; underground in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/underground (accent: US; accessed 2026-11-18).
```

### ZP246 — unlock

```plain
Pronunciation source: Collins Dictionary — unlock; unlock in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/unlock (accent: US; accessed 2026-11-18).
```

### ZP247 — upgrade

```plain
Pronunciation source: Collins Dictionary — upgrade; upgrade in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/upgrade (accent: US; accessed 2026-11-18).
```

### ZP248 — upheaval

```plain
Pronunciation source: Collins Dictionary — upheaval; upheaval in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/upheaval (accent: US; accessed 2026-11-18).
```

### ZP249 — view

```plain
Pronunciation source: Collins Dictionary — view; view in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/view (accent: US; accessed 2026-11-18).
```

### ZP250 — volcanic

```plain
Pronunciation source: Collins Dictionary — volcanic; volcanic in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/volcanic (accent: US; accessed 2026-11-18).
```

### ZP251 — vulnerable

```plain
Pronunciation source: Collins Dictionary — vulnerable; vulnerable in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/vulnerable (accent: US; accessed 2026-11-18).
```

### ZP252 — warrior

```plain
Pronunciation source: Collins Dictionary — warrior; warrior in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/warrior (accent: US; accessed 2026-11-18).
```

### ZP253 — waterfall

```plain
Pronunciation source: Collins Dictionary — waterfall; waterfall in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/waterfall (accent: US; accessed 2026-11-18).
```

### ZP254 — weapon

```plain
Pronunciation source: Collins Dictionary — weapon; weapon in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/weapon (accent: US; accessed 2026-11-18).
```

### ZP255 — well

```plain
Pronunciation source: Collins Dictionary — well; well in American English 1; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/well (accent: US; accessed 2026-11-18).
```

### ZP256 — wheel

```plain
Pronunciation source: Collins Dictionary — wheel; wheel in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/wheel (accent: US; accessed 2026-11-18).
```

### ZP257 — wield

```plain
Pronunciation source: Collins Dictionary — wield; wield in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/wield (accent: US; accessed 2026-11-18).
```

### ZP258 — wind

```plain
Pronunciation source: Collins Dictionary — wind; wind in American English 2; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/wind (accent: US; accessed 2026-11-18).
```

### ZP259 — wing

```plain
Pronunciation source: Collins Dictionary — wing; wing in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/wing (accent: US; accessed 2026-11-18).
```

## 待核实项目

以下没有待写音标值，批准具体提案也不授权猜写这些项目。普通词的限制列出实际查证结果；游戏名称组完成内容筛查和逐项列账，但没有逐个追查、收听或转录官方英文语音。这里的暂缓表示本轮证据不够，不表示英语中没有读法。

### 普通词形、短语与游戏用法

29 项。后续可补充可读取的可靠词典页面，或另行批准有证据的整词转写。即使已经有组件词的音标，也不直接拼出短语、屈折词或派生词的完整转写。

| 编号  | 稳定 ID / Word                                                                          | 当前 Meaning                                 | 原因         | 实际限制及下一步                                                                                                                                           |
|:------|:----------------------------------------------------------------------------------------|:---------------------------------------------|:-------------|:-----------------------------------------------------------------------------------------------------------------------------------------------------------|
| ZD001 | [arowana-general](../entries/arowana-general.yml) / arowana                             | 龙鱼。                                       | 来源访问失败 | Collins arowana 两次读取均不可访问；Cambridge 发音页不可访问。搜索未建立可采用的完整发音证据，不按拼写猜测。                                               |
| ZD002 | [attack-power-combat](../entries/attack-power-combat.yml) / attack power                | 攻击力                                       | 短语整体     | 所查 Collins 对应页面没有可采用的整体音标，或页面不可访问；备用 Cambridge 发音页也未取得可读整体音标。短语释义或组件已有读音不等于整句发音证据，暂不拼接。 |
| ZD003 | [boots-clothing](../entries/boots-clothing.yml) / boots                                 | 靴子                                         | 同形不同义   | Collins boots 美式页给出的独立名词义是擦鞋仆役，不能作为当前靴子义的词形证据；备用 Cambridge 发音页返回 403。保留复数，不由 boot 自行补词尾。              |
| ZD004 | [cold-resistance-combat](../entries/cold-resistance-combat.yml) / cold resistance       | 耐寒；承受寒冷环境影响的能力。               | 短语整体     | 所查 Collins 对应页面没有可采用的整体音标，或页面不可访问；备用 Cambridge 发音页也未取得可读整体音标。短语释义或组件已有读音不等于整句发音证据，暂不拼接。 |
| ZD005 | [cooking-pot-cooking](../entries/cooking-pot-cooking.yml) / cooking pot                 | 烹饪锅                                       | 短语整体     | 所查 Collins 对应页面没有可采用的整体音标，或页面不可访问；备用 Cambridge 发音页也未取得可读整体音标。短语释义或组件已有读音不等于整句发音证据，暂不拼接。 |
| ZD006 | [coordinates-ui](../entries/coordinates-ui.yml) / coordinates                           | 坐标                                         | 词形与重音   | Collins coordinates 美式片段为协调搭配的衣物等义，未提供坐标复数整体音标。需要坐标名词义而非动词的证据；Cambridge 发音页返回 403。                         |
| ZD007 | [corrupted-combat](../entries/corrupted-combat.yml) / corrupted                         | 被腐化的                                     | 当前词形     | 所查 Collins 当前词形页面未取得可采用的整体美式音标；备用 Cambridge 发音页未能读取。不以原形读音替代，也不自行添加屈折词尾。                               |
| ZD008 | [critical-hit-combat](../entries/critical-hit-combat.yml) / critical hit                | 暴击                                         | 短语整体     | 所查 Collins 对应页面没有可采用的整体音标，或页面不可访问；备用 Cambridge 发音页也未取得可读整体音标。短语释义或组件已有读音不等于整句发音证据，暂不拼接。 |
| ZD009 | [decayed-general](../entries/decayed-general.yml) / decayed                             | 腐朽的；腐败或衰损的。                       | 当前词形     | 所查 Collins 当前词形页面未取得可采用的整体美式音标；备用 Cambridge 发音页未能读取。不以原形读音替代，也不自行添加屈折词尾。                               |
| ZD010 | [depths-terrain](../entries/depths-terrain.yml) / depths                                | 深处；深层区域；常用复数形式。               | 当前词形     | 所查 Collins 当前词形页面未取得可采用的整体美式音标；备用 Cambridge 发音页未能读取。不以原形读音替代，也不自行添加屈折词尾。                               |
| ZD011 | [disguised-combat](../entries/disguised-combat.yml) / disguised                         | 伪装的                                       | 当前词形     | 所查 Collins 当前词形页面未取得可采用的整体美式音标；备用 Cambridge 发音页未能读取。不以原形读音替代，也不自行添加屈折词尾。                               |
| ZD012 | [durability-equipment](../entries/durability-equipment.yml) / durability                | 耐用性；耐久度。                             | 派生词整体   | Collins durability 返回正文未取得可采用的美式整体音标；Cambridge 词条及发音页返回 403。不从 durable 推导重音。                                             |
| ZD013 | [electro-cooking](../entries/electro-cooking.yml) / Electro                             | 电麻；本作料理效果名前缀，表示提供抗电效果。 | 游戏前缀     | 当前 Electro 为游戏效果名称前缀；Collins electro 已读分区为普通名词或缩略形式，不直接套用其音标。需要名称用法的英文读音证据。                              |
| ZD014 | [enhancement-general](../entries/enhancement-general.yml) / enhancement                 | 强化、增强                                   | 派生词整体   | Collins enhancement 返回正文未取得可采用的美式整体音标；Cambridge 词条及发音页返回 403。不从 enhance 拼接派生词读音。                                      |
| ZD015 | [general-store-architecture](../entries/general-store-architecture.yml) / general store | 杂货店                                       | 短语整体     | 所查 Collins 对应页面没有可采用的整体音标，或页面不可访问；备用 Cambridge 发音页也未取得可读整体音标。短语释义或组件已有读音不等于整句发音证据，暂不拼接。 |
| ZD016 | [greaves-clothing](../entries/greaves-clothing.yml) / greaves                           | 护胫                                         | 同形不同义   | Collins greaves 美式页为熬油后残渣义，不是护胫；备用 Cambridge 发音页返回 403。需要当前护胫复数词形的可靠音标。                                            |
| ZD017 | [greens-food](../entries/greens-food.yml) / greens                                      | 绿叶蔬菜。                                   | 复数词形     | Collins greens 美式正文涵盖绿叶蔬菜义，但没有整体读音。Cambridge 发音页返回 403；不自行拼接 green 与复数词尾。                                             |
| ZD018 | [heat-resistance-combat](../entries/heat-resistance-combat.yml) / heat resistance       | 耐热；承受酷暑环境影响的能力。               | 短语整体     | 所查 Collins 对应页面没有可采用的整体音标，或页面不可访问；备用 Cambridge 发音页也未取得可读整体音标。短语释义或组件已有读音不等于整句发音证据，暂不拼接。 |
| ZD019 | [highlands-terrain](../entries/highlands-terrain.yml) / highlands                       | 高地                                         | 复数词形     | Collins highlands 美式正文有高地释义但没有对应美式整体 IPA；当前未改用翻译列表中明确标为英式的读音。Cambridge 发音页返回 403。                             |
| ZD020 | [leggings-clothing](../entries/leggings-clothing.yml) / leggings                        | 紧身裤、护腿                                 | 复数词形     | Collins leggings 美式正文有当前衣物义但没有整体 IPA；Cambridge 发音页返回 403。不自行从 legging 构造复数音标。                                             |
| ZD021 | [pick-up-general](../entries/pick-up-general.yml) / pick up                             | 捡起                                         | 短语整体     | 所查 Collins 对应页面没有可采用的整体音标，或页面不可访问；备用 Cambridge 发音页也未取得可读整体音标。短语释义或组件已有读音不等于整句发音证据，暂不拼接。 |
| ZD022 | [ranged-combat](../entries/ranged-combat.yml) / ranged                                  | 远程                                         | 同形不同义   | Collins ranged 美式页为建筑分层砌筑义，不能当作当前远程战斗义的来源；备用 Cambridge 发音页不可访问。                                                       |
| ZD023 | [ruins-architecture](../entries/ruins-architecture.yml) / ruins                         | 废墟、遗迹                                   | 复数词形     | Collins ruins 美式正文有遗迹义但没有整体音标；Cambridge 发音页返回 403。不自行从 ruin 追加词尾。                                                           |
| ZD024 | [saute-cooking](../entries/saute-cooking.yml) / sauté                                   | 用少量油快速翻炒或煎炒。                     | 来源访问失败 | 当前拼写为 sauté；Collins saute 及重音字符网址未取得可读的对应美式正文。Cambridge 搜索结果有音标，但实际发音页返回 403，未将摘要当作已读正文。             |
| ZD025 | [shock-resistance-combat](../entries/shock-resistance-combat.yml) / shock resistance    | 抗电；抵抗电击影响的能力。                   | 短语整体     | 所查 Collins 对应页面没有可采用的整体音标，或页面不可访问；备用 Cambridge 发音页也未取得可读整体音标。短语释义或组件已有读音不等于整句发音证据，暂不拼接。 |
| ZD026 | [unfuse-general](../entries/unfuse-general.yml) / unfuse                                | 解除融合                                     | 来源访问失败 | Collins 与 Cambridge 返回正文均未取得可采用的整词音标，不从 un- 和 fuse 拼接。                                                                             |
| ZD027 | [waistguard-clothing](../entries/waistguard-clothing.yml) / waistguard                  | 护腰                                         | 复合词整体   | Collins 与 Cambridge 返回正文均未取得可采用的整词音标，不从 waist 和 guard 拼接。                                                                          |
| ZD028 | [warp-general](../entries/warp-general.yml) / warp                                      | 传送；游戏中快速移动到另一地点。             | 游戏义待核   | 当前是游戏传送义；所读普通词典片段以扭曲、牵引等意义为主，尚未核对传送义对应的读音来源。暂不自动套用。                                                     |
| ZD029 | [wetlands-terrain](../entries/wetlands-terrain.yml) / wetlands                          | 湿地                                         | 复数词形     | Collins wetlands 返回正文未取得对应整体美式 IPA；Cambridge 发音页返回 403。保留当前复数形式。                                                              |

### 游戏名称与造词整体

173 项。下表逐项保留 Word、ID 和当前 Meaning，按名称内容分类；没有把英语名称拆成已有词音标来填值，也没有用当前 TTS 录音验证自身读音。Notes 中已有的固定版本游戏文本链接支持拼写、身份、译名或机制，通常不直接包含 IPA。

这些项目需要可定位的英文配音或可靠的专名整体发音来源。仅确认中文名、日文名或世界观机制，不能解决英语重音、元音和连读问题。人物同名于现实英语姓名，也不直接套用现实人物的读音。普通词组成的完整名称还需要整体重音证据。

**人物: 21 项。**

| 编号  | 稳定 ID / Word                                                                  | 当前 Meaning                         | 待补证据                                                   |
|:------|:--------------------------------------------------------------------------------|:-------------------------------------|:-----------------------------------------------------------|
| ZD070 | [ganondorf-zelda](../entries/ganondorf-zelda.yml) / Ganondorf                   | 加侬多夫；本作的魔王。               | 当前名称的英文整体读音来源；未逐项查证官方语音或整体 IPA。 |
| ZD086 | [great-deku-tree-zelda](../entries/great-deku-tree-zelda.yml) / Great Deku Tree | 德库树长老；克洛格森林中的巨树。     | 当前名称的英文整体读音来源；未逐项查证官方语音或整体 IPA。 |
| ZD093 | [hestu-zelda](../entries/hestu-zelda.yml) / Hestu                               | 伯库林；帮助扩充背包容量的克洛格。   | 当前名称的英文整体读音来源；未逐项查证官方语音或整体 IPA。 |
| ZD098 | [hudson-zelda](../entries/hudson-zelda.yml) / Hudson                            | 松达；从事建造事业的人物。           | 当前名称的英文整体读音来源；未逐项查证官方语音或整体 IPA。 |
| ZD107 | [impa-zelda](../entries/impa-zelda.yml) / Impa                                  | 英帕；希卡族长者。                   | 当前名称的英文整体读音来源；未逐项查证官方语音或整体 IPA。 |
| ZD108 | [josha-zelda](../entries/josha-zelda.yml) / Josha                               | 乔舒亚；研究地底的调查人员。         | 当前名称的英文整体读音来源；未逐项查证官方语音或整体 IPA。 |
| ZD122 | [link-zelda](../entries/link-zelda.yml) / Link                                  | 林克；《王国之泪》的主角。           | 当前名称的英文整体读音来源；未逐项查证官方语音或整体 IPA。 |
| ZD130 | [master-kohga-zelda](../entries/master-kohga-zelda.yml) / Master Kohga          | 可盖大人；依盖队的首领。             | 当前名称的英文整体读音来源；未逐项查证官方语音或整体 IPA。 |
| ZD134 | [mineru-zelda](../entries/mineru-zelda.yml) / Mineru                            | 米涅鲁；古代左纳乌族人，魂之贤者。   | 当前名称的英文整体读音来源；未逐项查证官方语音或整体 IPA。 |
| ZD138 | [paya-zelda](../entries/paya-zelda.yml) / Paya                                  | 帕雅；卡卡利科村的领导者。           | 当前名称的英文整体读音来源；未逐项查证官方语音或整体 IPA。 |
| ZD141 | [princess-zelda-zelda](../entries/princess-zelda-zelda.yml) / Princess Zelda    | 塞尔达公主；海拉鲁的公主。           | 当前名称的英文整体读音来源；未逐项查证官方语音或整体 IPA。 |
| ZD144 | [purah-zelda](../entries/purah-zelda.yml) / Purah                               | 普尔亚；监视堡垒的重要研究人员。     | 当前名称的英文整体读音来源；未逐项查证官方语音或整体 IPA。 |
| ZD146 | [rauru-zelda](../entries/rauru-zelda.yml) / Rauru                               | 劳鲁；古代海拉鲁的国王，左纳乌族人。 | 当前名称的英文整体读音来源；未逐项查证官方语音或整体 IPA。 |
| ZD149 | [riju-zelda](../entries/riju-zelda.yml) / Riju                                  | 露珠；格鲁德族的领导者，雷之贤者。   | 当前名称的英文整体读音来源；未逐项查证官方语音或整体 IPA。 |
| ZD152 | [robbie-zelda](../entries/robbie-zelda.yml) / Robbie                            | 洛贝利；研究古代技术的研究人员。     | 当前名称的英文整体读音来源；未逐项查证官方语音或整体 IPA。 |
| ZD162 | [sidon-zelda](../entries/sidon-zelda.yml) / Sidon                               | 希多；卓拉族人，水之贤者。           | 当前名称的英文整体读音来源；未逐项查证官方语音或整体 IPA。 |
| ZD172 | [sonia-zelda](../entries/sonia-zelda.yml) / Sonia                               | 索尼娅；古代海拉鲁的王后。           | 当前名称的英文整体读音来源；未逐项查证官方语音或整体 IPA。 |
| ZD183 | [teba-zelda](../entries/teba-zelda.yml) / Teba                                  | 特巴；利特族战士，丘栗的父亲。       | 当前名称的英文整体读音来源；未逐项查证官方语音或整体 IPA。 |
| ZD188 | [tulin-zelda](../entries/tulin-zelda.yml) / Tulin                               | 丘栗；利特族青年，风之贤者。         | 当前名称的英文整体读音来源；未逐项查证官方语音或整体 IPA。 |
| ZD195 | [yona-zelda](../entries/yona-zelda.yml) / Yona                                  | 优娜；卓拉族人，希多的伴侣。         | 当前名称的英文整体读音来源；未逐项查证官方语音或整体 IPA。 |
| ZD196 | [yunobo-zelda](../entries/yunobo-zelda.yml) / Yunobo                            | 阿沅；鼓隆族人，火之贤者。           | 当前名称的英文整体读音来源；未逐项查证官方语音或整体 IPA。 |

**族群与组织: 9 项。**

| 编号  | 稳定 ID / Word                                                | 当前 Meaning                           | 待补证据                                                   |
|:------|:--------------------------------------------------------------|:---------------------------------------|:-----------------------------------------------------------|
| ZD073 | [gerudo-zelda](../entries/gerudo-zelda.yml) / Gerudo          | 格鲁德族；《王国之泪》中的族群或组织。 | 当前名称的英文整体读音来源；未逐项查证官方语音或整体 IPA。 |
| ZD084 | [goron-zelda](../entries/goron-zelda.yml) / Goron             | 鼓隆族；《王国之泪》中的族群或组织。   | 当前名称的英文整体读音来源；未逐项查证官方语音或整体 IPA。 |
| ZD102 | [hylian-zelda](../entries/hylian-zelda.yml) / Hylian          | 海利亚人；《王国之泪》中的族群或组织。 | 当前名称的英文整体读音来源；未逐项查证官方语音或整体 IPA。 |
| ZD116 | [korok-zelda](../entries/korok-zelda.yml) / Korok             | 克洛格族；《王国之泪》中的族群或组织。 | 当前名称的英文整体读音来源；未逐项查证官方语音或整体 IPA。 |
| ZD151 | [rito-zelda](../entries/rito-zelda.yml) / Rito                | 利特族；《王国之泪》中的族群或组织。   | 当前名称的英文整体读音来源；未逐项查证官方语音或整体 IPA。 |
| ZD157 | [sheikah-zelda](../entries/sheikah-zelda.yml) / Sheikah       | 希卡族；《王国之泪》中的族群或组织。   | 当前名称的英文整体读音来源；未逐项查证官方语音或整体 IPA。 |
| ZD194 | [yiga-clan-zelda](../entries/yiga-clan-zelda.yml) / Yiga Clan | 依盖队；《王国之泪》中的族群或组织。   | 当前名称的英文整体读音来源；未逐项查证官方语音或整体 IPA。 |
| ZD198 | [zonai-zelda](../entries/zonai-zelda.yml) / Zonai             | 左纳乌族；《王国之泪》中的族群或组织。 | 当前名称的英文整体读音来源；未逐项查证官方语音或整体 IPA。 |
| ZD201 | [zora-zelda](../entries/zora-zelda.yml) / Zora                | 卓拉族；《王国之泪》中的族群或组织。   | 当前名称的英文整体读音来源；未逐项查证官方语音或整体 IPA。 |

**怪物与首领: 29 项。**

| 编号  | 稳定 ID / Word                                                                                 | 当前 Meaning                                   | 待补证据                                                   |
|:------|:-----------------------------------------------------------------------------------------------|:-----------------------------------------------|:-----------------------------------------------------------|
| ZD031 | [aerocuda-zelda](../entries/aerocuda-zelda.yml) / Aerocuda                                     | 卡库达；《王国之泪》中的怪物或首领名称。       | 当前名称的英文整体读音来源；未逐项查证官方语音或整体 IPA。 |
| ZD038 | [bokoblin-zelda](../entries/bokoblin-zelda.yml) / Bokoblin                                     | 波克布林；《王国之泪》中的怪物或首领名称。     | 当前名称的英文整体读音来源；未逐项查证官方语音或整体 IPA。 |
| ZD039 | [boss-bokoblin-zelda](../entries/boss-bokoblin-zelda.yml) / Boss Bokoblin                      | 首领波克布林；《王国之泪》中的怪物或首领名称。 | 当前名称的英文整体读音来源；未逐项查证官方语音或整体 IPA。 |
| ZD041 | [captain-construct-zelda](../entries/captain-construct-zelda.yml) / Captain Construct          | 队长魔像；《王国之泪》中的怪物或首领名称。     | 当前名称的英文整体读音来源；未逐项查证官方语音或整体 IPA。 |
| ZD043 | [chuchu-zelda](../entries/chuchu-zelda.yml) / Chuchu                                           | 丘丘；《王国之泪》中的怪物或首领名称。         | 当前名称的英文整体读音来源；未逐项查证官方语音或整体 IPA。 |
| ZD045 | [colgera-zelda](../entries/colgera-zelda.yml) / Colgera                                        | 急冻盖拉；《王国之泪》中的怪物或首领名称。     | 当前名称的英文整体读音来源；未逐项查证官方语音或整体 IPA。 |
| ZD053 | [demon-king-ganondorf-zelda](../entries/demon-king-ganondorf-zelda.yml) / Demon King Ganondorf | 魔王加侬多夫；《王国之泪》中的怪物或首领名称。 | 当前名称的英文整体读音来源；未逐项查证官方语音或整体 IPA。 |
| ZD060 | [evermean-zelda](../entries/evermean-zelda.yml) / Evermean                                     | 巨霸柏古；《王国之泪》中的怪物或首领名称。     | 当前名称的英文整体读音来源；未逐项查证官方语音或整体 IPA。 |
| ZD065 | [flux-construct-zelda](../entries/flux-construct-zelda.yml) / Flux Construct                   | 方块魔像；《王国之泪》中的怪物或首领名称。     | 当前名称的英文整体读音来源；未逐项查证官方语音或整体 IPA。 |
| ZD069 | [frox-zelda](../entries/frox-zelda.yml) / Frox                                                 | 巨霸迦马；《王国之泪》中的怪物或首领名称。     | 当前名称的英文整体读音来源；未逐项查证官方语音或整体 IPA。 |
| ZD074 | [gibdo-zelda](../entries/gibdo-zelda.yml) / Gibdo                                              | 吉波得；《王国之泪》中的怪物或首领名称。       | 当前名称的英文整体读音来源；未逐项查证官方语音或整体 IPA。 |
| ZD075 | [gleeok-zelda](../entries/gleeok-zelda.yml) / Gleeok                                           | 古栗欧克；《王国之泪》中的怪物或首领名称。     | 当前名称的英文整体读音来源；未逐项查证官方语音或整体 IPA。 |
| ZD079 | [gloom-spawn-zelda](../entries/gloom-spawn-zelda.yml) / Gloom Spawn                            | 瘴气魔；《王国之泪》中的怪物或首领名称。       | 当前名称的英文整体读音来源；未逐项查证官方语音或整体 IPA。 |
| ZD094 | [hinox-zelda](../entries/hinox-zelda.yml) / Hinox                                              | 西诺克斯；《王国之泪》中的怪物或首领名称。     | 当前名称的英文整体读音来源；未逐项查证官方语音或整体 IPA。 |
| ZD096 | [horriblin-zelda](../entries/horriblin-zelda.yml) / Horriblin                                  | 霍拉布林；《王国之泪》中的怪物或首领名称。     | 当前名称的英文整体读音来源；未逐项查证官方语音或整体 IPA。 |
| ZD110 | [keese-zelda](../entries/keese-zelda.yml) / Keese                                              | 蝙蝠；《王国之泪》中的怪物或首领名称。         | 当前名称的英文整体读音来源；未逐项查证官方语音或整体 IPA。 |
| ZD112 | [king-gleeok-zelda](../entries/king-gleeok-zelda.yml) / King Gleeok                            | 古栗欧克王；《王国之泪》中的怪物或首领名称。   | 当前名称的英文整体读音来源；未逐项查证官方语音或整体 IPA。 |
| ZD121 | [like-like-zelda](../entries/like-like-zelda.yml) / Like Like                                  | 莱克莱克；《王国之泪》中的怪物或首领名称。     | 当前名称的英文整体读音来源；未逐项查证官方语音或整体 IPA。 |
| ZD123 | [lizalfos-zelda](../entries/lizalfos-zelda.yml) / Lizalfos                                     | 蜥蜴战士；《王国之泪》中的怪物或首领名称。     | 当前名称的英文整体读音来源；未逐项查证官方语音或整体 IPA。 |
| ZD126 | [lynel-zelda](../entries/lynel-zelda.yml) / Lynel                                              | 莱尼尔；《王国之泪》中的怪物或首领名称。       | 当前名称的英文整体读音来源；未逐项查证官方语音或整体 IPA。 |
| ZD129 | [marbled-gohma-zelda](../entries/marbled-gohma-zelda.yml) / Marbled Gohma                      | 巨砾哥马；《王国之泪》中的怪物或首领名称。     | 当前名称的英文整体读音来源；未逐项查证官方语音或整体 IPA。 |
| ZD135 | [moblin-zelda](../entries/moblin-zelda.yml) / Moblin                                           | 莫力布林；《王国之泪》中的怪物或首领名称。     | 当前名称的英文整体读音来源；未逐项查证官方语音或整体 IPA。 |
| ZD136 | [mucktorok-zelda](../entries/mucktorok-zelda.yml) / Mucktorok                                  | 八爪怪章；《王国之泪》中的怪物或首领名称。     | 当前名称的英文整体读音来源；未逐项查证官方语音或整体 IPA。 |
| ZD137 | [octorok-zelda](../entries/octorok-zelda.yml) / Octorok                                        | 八爪怪；《王国之泪》中的怪物或首领名称。       | 当前名称的英文整体读音来源；未逐项查证官方语音或整体 IPA。 |
| ZD145 | [queen-gibdo-zelda](../entries/queen-gibdo-zelda.yml) / Queen Gibdo                            | 吉波得女王；《王国之泪》中的怪物或首领名称。   | 当前名称的英文整体读音来源；未逐项查证官方语音或整体 IPA。 |
| ZD156 | [seized-construct-zelda](../entries/seized-construct-zelda.yml) / Seized Construct             | 被夺之魔像；《王国之泪》中的怪物或首领名称。   | 当前名称的英文整体读音来源；未逐项查证官方语音或整体 IPA。 |
| ZD170 | [soldier-construct-zelda](../entries/soldier-construct-zelda.yml) / Soldier Construct          | 兵队魔像；《王国之泪》中的怪物或首领名称。     | 当前名称的英文整体读音来源；未逐项查证官方语音或整体 IPA。 |
| ZD180 | [stone-talus-zelda](../entries/stone-talus-zelda.yml) / Stone Talus                            | 岩石巨人；《王国之泪》中的怪物或首领名称。     | 当前名称的英文整体读音来源；未逐项查证官方语音或整体 IPA。 |
| ZD192 | [wizzrobe-zelda](../entries/wizzrobe-zelda.yml) / Wizzrobe                                     | 长袍魔法师；《王国之泪》中的怪物或首领名称。   | 当前名称的英文整体读音来源；未逐项查证官方语音或整体 IPA。 |

**地点与设施: 27 项。**

| 编号  | 稳定 ID / Word                                                                                                         | 当前 Meaning                                   | 待补证据                                                   |
|:------|:-----------------------------------------------------------------------------------------------------------------------|:-----------------------------------------------|:-----------------------------------------------------------|
| ZD046 | [construct-factory-zelda](../entries/construct-factory-zelda.yml) / Construct Factory                                  | 魔像制造房；《王国之泪》中的地点或设施。       | 当前名称的英文整体读音来源；未逐项查证官方语音或整体 IPA。 |
| ZD049 | [crystal-refinery-zelda](../entries/crystal-refinery-zelda.yml) / Crystal Refinery                                     | 电池制造机；《王国之泪》中的地点或设施。       | 当前名称的英文整体读音来源；未逐项查证官方语音或整体 IPA。 |
| ZD050 | [death-mountain-zelda](../entries/death-mountain-zelda.yml) / Death Mountain                                           | 死亡之山；《王国之泪》中的地点或设施。         | 当前名称的英文整体读音来源；未逐项查证官方语音或整体 IPA。 |
| ZD056 | [device-dispenser-zelda](../entries/device-dispenser-zelda.yml) / Device Dispenser                                     | 左纳乌装置制造机；《王国之泪》中的地点或设施。 | 当前名称的英文整体读音来源；未逐项查证官方语音或整体 IPA。 |
| ZD062 | [fire-temple-zelda](../entries/fire-temple-zelda.yml) / Fire Temple                                                    | 火之神殿；《王国之泪》中的地点或设施。         | 当前名称的英文整体读音来源；未逐项查证官方语音或整体 IPA。 |
| ZD066 | [forgotten-temple-zelda](../entries/forgotten-temple-zelda.yml) / Forgotten Temple                                     | 忘却神殿；《王国之泪》中的地点或设施。         | 当前名称的英文整体读音来源；未逐项查证官方语音或整体 IPA。 |
| ZD071 | [gerudo-desert-zelda](../entries/gerudo-desert-zelda.yml) / Gerudo Desert                                              | 格鲁德沙漠；《王国之泪》中的地点或设施。       | 当前名称的英文整体读音来源；未逐项查证官方语音或整体 IPA。 |
| ZD072 | [gerudo-town-zelda](../entries/gerudo-town-zelda.yml) / Gerudo Town                                                    | 格鲁德小镇；《王国之泪》中的地点或设施。       | 当前名称的英文整体读音来源；未逐项查证官方语音或整体 IPA。 |
| ZD082 | [goron-city-zelda](../entries/goron-city-zelda.yml) / Goron City                                                       | 鼓隆城；《王国之泪》中的地点或设施。           | 当前名称的英文整体读音来源；未逐项查证官方语音或整体 IPA。 |
| ZD085 | [great-abandoned-central-mine-zelda](../entries/great-abandoned-central-mine-zelda.yml) / Great Abandoned Central Mine | 中央大废矿；《王国之泪》中的地点或设施。       | 当前名称的英文整体读音来源；未逐项查证官方语音或整体 IPA。 |
| ZD088 | [great-sky-island-zelda](../entries/great-sky-island-zelda.yml) / Great Sky Island                                     | 初始空岛；《王国之泪》中的地点或设施。         | 当前名称的英文整体读音来源；未逐项查证官方语音或整体 IPA。 |
| ZD089 | [hateno-village-zelda](../entries/hateno-village-zelda.yml) / Hateno Village                                           | 哈特诺村；《王国之泪》中的地点或设施。         | 当前名称的英文整体读音来源；未逐项查证官方语音或整体 IPA。 |
| ZD104 | [hyrule-castle-zelda](../entries/hyrule-castle-zelda.yml) / Hyrule Castle                                              | 海拉鲁城堡；《王国之泪》中的地点或设施。       | 当前名称的英文整体读音来源；未逐项查证官方语音或整体 IPA。 |
| ZD106 | [hyrule-zelda](../entries/hyrule-zelda.yml) / Hyrule                                                                   | 海拉鲁；海拉鲁王国，本作冒险的主要舞台。       | 当前名称的英文整体读音来源；未逐项查证官方语音或整体 IPA。 |
| ZD109 | [kakariko-village-zelda](../entries/kakariko-village-zelda.yml) / Kakariko Village                                     | 卡卡利科村；《王国之泪》中的地点或设施。       | 当前名称的英文整体读音来源；未逐项查证官方语音或整体 IPA。 |
| ZD114 | [korok-forest-zelda](../entries/korok-forest-zelda.yml) / Korok Forest                                                 | 克洛格森林；《王国之泪》中的地点或设施。       | 当前名称的英文整体读音来源；未逐项查证官方语音或整体 IPA。 |
| ZD119 | [lightning-temple-zelda](../entries/lightning-temple-zelda.yml) / Lightning Temple                                     | 雷之神殿；《王国之泪》中的地点或设施。         | 当前名称的英文整体读音来源；未逐项查证官方语音或整体 IPA。 |
| ZD124 | [lookout-landing-zelda](../entries/lookout-landing-zelda.yml) / Lookout Landing                                        | 监视堡垒；《王国之泪》中的地点或设施。         | 当前名称的英文整体读音来源；未逐项查证官方语音或整体 IPA。 |
| ZD125 | [lurelin-village-zelda](../entries/lurelin-village-zelda.yml) / Lurelin Village                                        | 沃托里村；《王国之泪》中的地点或设施。         | 当前名称的英文整体读音来源；未逐项查证官方语音或整体 IPA。 |
| ZD150 | [rito-village-zelda](../entries/rito-village-zelda.yml) / Rito Village                                                 | 利特村；《王国之泪》中的地点或设施。           | 当前名称的英文整体读音来源；未逐项查证官方语音或整体 IPA。 |
| ZD176 | [spirit-temple-zelda](../entries/spirit-temple-zelda.yml) / Spirit Temple                                              | 魂之神殿；《王国之泪》中的地点或设施。         | 当前名称的英文整体读音来源；未逐项查证官方语音或整体 IPA。 |
| ZD182 | [tarrey-town-zelda](../entries/tarrey-town-zelda.yml) / Tarrey Town                                                    | 一始村；《王国之泪》中的地点或设施。           | 当前名称的英文整体读音来源；未逐项查证官方语音或整体 IPA。 |
| ZD184 | [temple-of-time-zelda](../entries/temple-of-time-zelda.yml) / Temple of Time                                           | 时之神殿；《王国之泪》中的地点或设施。         | 当前名称的英文整体读音来源；未逐项查证官方语音或整体 IPA。 |
| ZD190 | [water-temple-zelda](../entries/water-temple-zelda.yml) / Water Temple                                                 | 水之神殿；《王国之泪》中的地点或设施。         | 当前名称的英文整体读音来源；未逐项查证官方语音或整体 IPA。 |
| ZD191 | [wind-temple-zelda](../entries/wind-temple-zelda.yml) / Wind Temple                                                    | 风之神殿；《王国之泪》中的地点或设施。         | 当前名称的英文整体读音来源；未逐项查证官方语音或整体 IPA。 |
| ZD193 | [yiga-clan-hideout-zelda](../entries/yiga-clan-hideout-zelda.yml) / Yiga Clan Hideout                                  | 依盖队基地；《王国之泪》中的地点或设施。       | 当前名称的英文整体读音来源；未逐项查证官方语音或整体 IPA。 |
| ZD202 | [zoras-domain-zelda](../entries/zoras-domain-zelda.yml) / Zora's Domain                                                | 卓拉领地；《王国之泪》中的地点或设施。         | 当前名称的英文整体读音来源；未逐项查证官方语音或整体 IPA。 |

**防具名称: 11 项。**

| 编号  | 稳定 ID / Word                                                                           | 当前 Meaning                             | 待补证据                                                   |
|:------|:-----------------------------------------------------------------------------------------|:-----------------------------------------|:-----------------------------------------------------------|
| ZD044 | [climbing-gear-zelda](../entries/climbing-gear-zelda.yml) / Climbing Gear                | 攀登护手；《王国之泪》中的防具名称。     | 当前名称的英文整体读音来源；未逐项查证官方语音或整体 IPA。 |
| ZD061 | [fierce-deity-armor-zelda](../entries/fierce-deity-armor-zelda.yml) / Fierce Deity Armor | 鬼神服；《王国之泪》中的防具名称。       | 当前名称的英文整体读音来源；未逐项查证官方语音或整体 IPA。 |
| ZD064 | [flamebreaker-armor-zelda](../entries/flamebreaker-armor-zelda.yml) / Flamebreaker Armor | 耐火石铠；《王国之泪》中的防具名称。     | 当前名称的英文整体读音来源；未逐项查证官方语音或整体 IPA。 |
| ZD067 | [froggy-sleeve-zelda](../entries/froggy-sleeve-zelda.yml) / Froggy Sleeve                | 蛙之手甲；《王国之泪》中的防具名称。     | 当前名称的英文整体读音来源；未逐项查证官方语音或整体 IPA。 |
| ZD076 | [glide-shirt-zelda](../entries/glide-shirt-zelda.yml) / Glide Shirt                      | 滑翔紧身服；《王国之泪》中的防具名称。   | 当前名称的英文整体读音来源；未逐项查证官方语音或整体 IPA。 |
| ZD099 | [hylian-hood-zelda](../entries/hylian-hood-zelda.yml) / Hylian Hood                      | 海利亚兜帽；《王国之泪》中的防具名称。   | 当前名称的英文整体读音来源；未逐项查证官方语音或整体 IPA。 |
| ZD101 | [hylian-tunic-zelda](../entries/hylian-tunic-zelda.yml) / Hylian Tunic                   | 海利亚服；《王国之泪》中的防具名称。     | 当前名称的英文整体读音来源；未逐项查证官方语音或整体 IPA。 |
| ZD133 | [miners-mask-zelda](../entries/miners-mask-zelda.yml) / Miner's Mask                     | 采矿面罩；《王国之泪》中的防具名称。     | 当前名称的英文整体读音来源；未逐项查证官方语音或整体 IPA。 |
| ZD169 | [snowquill-tunic-zelda](../entries/snowquill-tunic-zelda.yml) / Snowquill Tunic          | 利特的羽绒服；《王国之泪》中的防具名称。 | 当前名称的英文整体读音来源；未逐项查证官方语音或整体 IPA。 |
| ZD177 | [stealth-mask-zelda](../entries/stealth-mask-zelda.yml) / Stealth Mask                   | 潜行面罩；《王国之泪》中的防具名称。     | 当前名称的英文整体读音来源；未逐项查证官方语音或整体 IPA。 |
| ZD200 | [zora-armor-zelda](../entries/zora-armor-zelda.yml) / Zora Armor                         | 卓拉铠甲；《王国之泪》中的防具名称。     | 当前名称的英文整体读音来源；未逐项查证官方语音或整体 IPA。 |

**装备名称: 11 项。**

| 编号  | 稳定 ID / Word                                                                                    | 当前 Meaning                         | 待补证据                                                   |
|:------|:--------------------------------------------------------------------------------------------------|:-------------------------------------|:-----------------------------------------------------------|
| ZD055 | [demon-kings-bow-zelda](../entries/demon-kings-bow-zelda.yml) / Demon King's Bow                  | 魔王之弓；《王国之泪》中的装备名称。 | 当前名称的英文整体读音来源；未逐项查证官方语音或整体 IPA。 |
| ZD080 | [gloom-sword-zelda](../entries/gloom-sword-zelda.yml) / Gloom Sword                               | 瘴气之剑；《王国之泪》中的装备名称。 | 当前名称的英文整体读音来源；未逐项查证官方语音或整体 IPA。 |
| ZD087 | [great-eagle-bow-zelda](../entries/great-eagle-bow-zelda.yml) / Great Eagle Bow                   | 大鹫弓；《王国之泪》中的装备名称。   | 当前名称的英文整体读音来源；未逐项查证官方语音或整体 IPA。 |
| ZD100 | [hylian-shield-zelda](../entries/hylian-shield-zelda.yml) / Hylian Shield                         | 海利亚盾；《王国之泪》中的装备名称。 | 当前名称的英文整体读音来源；未逐项查证官方语音或整体 IPA。 |
| ZD113 | [knights-broadsword-zelda](../entries/knights-broadsword-zelda.yml) / Knight's Broadsword         | 骑士之剑；《王国之泪》中的装备名称。 | 当前名称的英文整体读音来源；未逐项查证官方语音或整体 IPA。 |
| ZD127 | [magic-rod-zelda](../entries/magic-rod-zelda.yml) / Magic Rod                                     | 魔法杖；《王国之泪》中的装备名称。   | 当前名称的英文整体读音来源；未逐项查证官方语音或整体 IPA。 |
| ZD131 | [master-sword-zelda](../entries/master-sword-zelda.yml) / Master Sword                            | 大师之剑；《王国之泪》中的装备名称。 | 当前名称的英文整体读音来源；未逐项查证官方语音或整体 IPA。 |
| ZD153 | [royal-broadsword-zelda](../entries/royal-broadsword-zelda.yml) / Royal Broadsword                | 王族之剑；《王国之泪》中的装备名称。 | 当前名称的英文整体读音来源；未逐项查证官方语音或整体 IPA。 |
| ZD154 | [scimitar-of-the-seven-zelda](../entries/scimitar-of-the-seven-zelda.yml) / Scimitar of the Seven | 七宝匕首；《王国之泪》中的装备名称。 | 当前名称的英文整体读音来源；未逐项查证官方语音或整体 IPA。 |
| ZD171 | [soldiers-broadsword-zelda](../entries/soldiers-broadsword-zelda.yml) / Soldier's Broadsword      | 士兵之剑；《王国之泪》中的装备名称。 | 当前名称的英文整体读音来源；未逐项查证官方语音或整体 IPA。 |
| ZD187 | [travelers-sword-zelda](../entries/travelers-sword-zelda.yml) / Traveler's Sword                  | 旅人之剑；《王国之泪》中的装备名称。 | 当前名称的英文整体读音来源；未逐项查证官方语音或整体 IPA。 |

**左纳乌装置: 13 项。**

| 编号  | 稳定 ID / Word                                                               | 当前 Meaning                             | 待补证据                                                   |
|:------|:-----------------------------------------------------------------------------|:-----------------------------------------|:-----------------------------------------------------------|
| ZD035 | [beam-emitter-zelda](../entries/beam-emitter-zelda.yml) / Beam Emitter       | 光线头；《王国之泪》中的左纳乌装置。     | 当前名称的英文整体读音来源；未逐项查证官方语音或整体 IPA。 |
| ZD036 | [big-battery-zelda](../entries/big-battery-zelda.yml) / Big Battery          | 大容量电池；《王国之泪》中的左纳乌装置。 | 当前名称的英文整体读音来源；未逐项查证官方语音或整体 IPA。 |
| ZD037 | [big-wheel-zelda](../entries/big-wheel-zelda.yml) / Big Wheel                | 大轮胎；《王国之泪》中的左纳乌装置。     | 当前名称的英文整体读音来源；未逐项查证官方语音或整体 IPA。 |
| ZD047 | [construct-head-zelda](../entries/construct-head-zelda.yml) / Construct Head | 魔像的头部；《王国之泪》中的左纳乌装置。 | 当前名称的英文整体读音来源；未逐项查证官方语音或整体 IPA。 |
| ZD063 | [flame-emitter-zelda](../entries/flame-emitter-zelda.yml) / Flame Emitter    | 火龙头；《王国之泪》中的左纳乌装置。     | 当前名称的英文整体读音来源；未逐项查证官方语音或整体 IPA。 |
| ZD068 | [frost-emitter-zelda](../entries/frost-emitter-zelda.yml) / Frost Emitter    | 冰龙头；《王国之泪》中的左纳乌装置。     | 当前名称的英文整体读音来源；未逐项查证官方语音或整体 IPA。 |
| ZD095 | [homing-cart-zelda](../entries/homing-cart-zelda.yml) / Homing Cart          | 追踪平台车；《王国之泪》中的左纳乌装置。 | 当前名称的英文整体读音来源；未逐项查证官方语音或整体 IPA。 |
| ZD097 | [hover-stone-zelda](../entries/hover-stone-zelda.yml) / Hover Stone          | 浮游石；《王国之泪》中的左纳乌装置。     | 当前名称的英文整体读音来源；未逐项查证官方语音或整体 IPA。 |
| ZD140 | [portable-pot-zelda](../entries/portable-pot-zelda.yml) / Portable Pot       | 便携锅；《王国之泪》中的左纳乌装置。     | 当前名称的英文整体读音来源；未逐项查证官方语音或整体 IPA。 |
| ZD158 | [shock-emitter-zelda](../entries/shock-emitter-zelda.yml) / Shock Emitter    | 雷龙头；《王国之泪》中的左纳乌装置。     | 当前名称的英文整体读音来源；未逐项查证官方语音或整体 IPA。 |
| ZD168 | [small-wheel-zelda](../entries/small-wheel-zelda.yml) / Small Wheel          | 小轮胎；《王国之泪》中的左纳乌装置。     | 当前名称的英文整体读音来源；未逐项查证官方语音或整体 IPA。 |
| ZD179 | [steering-stick-zelda](../entries/steering-stick-zelda.yml) / Steering Stick | 操纵杆；《王国之泪》中的左纳乌装置。     | 当前名称的英文整体读音来源；未逐项查证官方语音或整体 IPA。 |
| ZD186 | [time-bomb-zelda](../entries/time-bomb-zelda.yml) / Time Bomb                | 计时炸弹；《王国之泪》中的左纳乌装置。   | 当前名称的英文整体读音来源；未逐项查证官方语音或整体 IPA。 |

**界面与功能名称: 9 项。**

| 编号  | 稳定 ID / Word                                                                           | 当前 Meaning                                     | 待补证据                                                   |
|:------|:-----------------------------------------------------------------------------------------|:-------------------------------------------------|:-----------------------------------------------------------|
| ZD030 | [adventure-log-zelda](../entries/adventure-log-zelda.yml) / Adventure Log                | 冒险笔记；《王国之泪》中的界面类别或功能名称。   | 当前名称的英文整体读音来源；未逐项查证官方语音或整体 IPA。 |
| ZD042 | [character-profiles-zelda](../entries/character-profiles-zelda.yml) / Character Profiles | 人物名录；《王国之泪》中的界面类别或功能名称。   | 当前名称的英文整体读音来源；未逐项查证官方语音或整体 IPA。 |
| ZD105 | [hyrule-compendium-zelda](../entries/hyrule-compendium-zelda.yml) / Hyrule Compendium    | 海拉鲁图鉴；《王国之泪》中的界面类别或功能名称。 | 当前名称的英文整体读音来源；未逐项查证官方语音或整体 IPA。 |
| ZD111 | [key-items-zelda](../entries/key-items-zelda.yml) / Key Items                            | 重要物品；《王国之泪》中的界面类别或功能名称。   | 当前名称的英文整体读音来源；未逐项查证官方语音或整体 IPA。 |
| ZD128 | [main-quests-zelda](../entries/main-quests-zelda.yml) / Main Quests                      | 主剧情挑战；《王国之泪》中的界面类别或功能名称。 | 当前名称的英文整体读音来源；未逐项查证官方语音或整体 IPA。 |
| ZD159 | [shrine-quests-zelda](../entries/shrine-quests-zelda.yml) / Shrine Quests                | 神庙挑战；《王国之泪》中的界面类别或功能名称。   | 当前名称的英文整体读音来源；未逐项查证官方语音或整体 IPA。 |
| ZD160 | [side-adventures-zelda](../entries/side-adventures-zelda.yml) / Side Adventures          | 情节挑战；《王国之泪》中的界面类别或功能名称。   | 当前名称的英文整体读音来源；未逐项查证官方语音或整体 IPA。 |
| ZD161 | [side-quests-zelda](../entries/side-quests-zelda.yml) / Side Quests                      | 迷你挑战；《王国之泪》中的界面类别或功能名称。   | 当前名称的英文整体读音来源；未逐项查证官方语音或整体 IPA。 |
| ZD173 | [special-controls-zelda](../entries/special-controls-zelda.yml) / Special Controls       | 特殊动作；《王国之泪》中的界面类别或功能名称。   | 当前名称的英文整体读音来源；未逐项查证官方语音或整体 IPA。 |

**能力效果名称: 7 项。**

| 编号  | 稳定 ID / Word                                                                     | 当前 Meaning                                 | 待补证据                                                   |
|:------|:-----------------------------------------------------------------------------------|:---------------------------------------------|:-----------------------------------------------------------|
| ZD033 | [attack-up-zelda](../entries/attack-up-zelda.yml) / Attack Up                      | 攻击力提升；《王国之泪》中的能力效果名称。   | 当前名称的英文整体读音来源；未逐项查证官方语音或整体 IPA。 |
| ZD051 | [defense-up-zelda](../entries/defense-up-zelda.yml) / Defense Up                   | 防御提升；《王国之泪》中的能力效果名称。     | 当前名称的英文整体读音来源；未逐项查证官方语音或整体 IPA。 |
| ZD077 | [gloom-recovery-zelda](../entries/gloom-recovery-zelda.yml) / Gloom Recovery       | 瘴气伤害回复；《王国之泪》中的能力效果名称。 | 当前名称的英文整体读音来源；未逐项查证官方语音或整体 IPA。 |
| ZD078 | [gloom-resistance-zelda](../entries/gloom-resistance-zelda.yml) / Gloom Resistance | 瘴气防护；《王国之泪》中的能力效果名称。     | 当前名称的英文整体读音来源；未逐项查证官方语音或整体 IPA。 |
| ZD167 | [slip-resistance-zelda](../entries/slip-resistance-zelda.yml) / Slip Resistance    | 滑落减轻；减少攀爬湿滑表面时滑落的效果。     | 当前名称的英文整体读音来源；未逐项查证官方语音或整体 IPA。 |
| ZD174 | [speed-up-zelda](../entries/speed-up-zelda.yml) / Speed Up                         | 移动力提升；《王国之泪》中的能力效果名称。   | 当前名称的英文整体读音来源；未逐项查证官方语音或整体 IPA。 |
| ZD178 | [stealth-up-zelda](../entries/stealth-up-zelda.yml) / Stealth Up                   | 安静性提升；《王国之泪》中的能力效果名称。   | 当前名称的英文整体读音来源；未逐项查证官方语音或整体 IPA。 |

**材料名称: 19 项。**

| 编号  | 稳定 ID / Word                                                                        | 当前 Meaning                             | 待补证据                                                   |
|:------|:--------------------------------------------------------------------------------------|:-----------------------------------------|:-----------------------------------------------------------|
| ZD032 | [ancient-arowana-zelda](../entries/ancient-arowana-zelda.yml) / Ancient Arowana       | 远昔骨舌鱼；《王国之泪》中的材料名称。   | 当前名称的英文整体读音来源；未逐项查证官方语音或整体 IPA。 |
| ZD040 | [brightcap-zelda](../entries/brightcap-zelda.yml) / Brightcap                         | 光亮蘑菇；《王国之泪》中的材料名称。     | 当前名称的英文整体读音来源；未逐项查证官方语音或整体 IPA。 |
| ZD048 | [courser-bee-honey-zelda](../entries/courser-bee-honey-zelda.yml) / Courser Bee Honey | 精力蜂的蜂蜜；《王国之泪》中的材料名称。 | 当前名称的英文整体读音来源；未逐项查证官方语音或整体 IPA。 |
| ZD058 | [endura-carrot-zelda](../entries/endura-carrot-zelda.yml) / Endura Carrot             | 毅力胡萝卜；《王国之泪》中的材料名称。   | 当前名称的英文整体读音来源；未逐项查证官方语音或整体 IPA。 |
| ZD059 | [endura-shroom-zelda](../entries/endura-shroom-zelda.yml) / Endura Shroom             | 毅力蘑菇；《王国之泪》中的材料名称。     | 当前名称的英文整体读音来源；未逐项查证官方语音或整体 IPA。 |
| ZD081 | [golden-apple-zelda](../entries/golden-apple-zelda.yml) / Golden Apple                | 金苹果；《王国之泪》中的材料名称。       | 当前名称的英文整体读音来源；未逐项查证官方语音或整体 IPA。 |
| ZD083 | [goron-spice-zelda](../entries/goron-spice-zelda.yml) / Goron Spice                   | 鼓隆的调味粉；《王国之泪》中的材料名称。 | 当前名称的英文整体读音来源；未逐项查证官方语音或整体 IPA。 |
| ZD090 | [hearty-radish-zelda](../entries/hearty-radish-zelda.yml) / Hearty Radish             | 生命小萝卜；《王国之泪》中的材料名称。   | 当前名称的英文整体读音来源；未逐项查证官方语音或整体 IPA。 |
| ZD091 | [hearty-salmon-zelda](../entries/hearty-salmon-zelda.yml) / Hearty Salmon             | 生命鲑鱼；《王国之泪》中的材料名称。     | 当前名称的英文整体读音来源；未逐项查证官方语音或整体 IPA。 |
| ZD092 | [hearty-truffle-zelda](../entries/hearty-truffle-zelda.yml) / Hearty Truffle          | 生命松露；《王国之泪》中的材料名称。     | 当前名称的英文整体读音来源；未逐项查证官方语音或整体 IPA。 |
| ZD103 | [hyrule-bass-zelda](../entries/hyrule-bass-zelda.yml) / Hyrule Bass                   | 海拉鲁鲈鱼；《王国之泪》中的材料名称。   | 当前名称的英文整体读音来源；未逐项查证官方语音或整体 IPA。 |
| ZD132 | [mighty-bananas-zelda](../entries/mighty-bananas-zelda.yml) / Mighty Bananas          | 大剑香蕉；《王国之泪》中的材料名称。     | 当前名称的英文整体读音来源；未逐项查证官方语音或整体 IPA。 |
| ZD142 | [puffshroom-zelda](../entries/puffshroom-zelda.yml) / Puffshroom                      | 烟雾蘑菇；《王国之泪》中的材料名称。     | 当前名称的英文整体读音来源；未逐项查证官方语音或整体 IPA。 |
| ZD147 | [raw-gourmet-meat-zelda](../entries/raw-gourmet-meat-zelda.yml) / Raw Gourmet Meat    | 顶级兽肉；《王国之泪》中的材料名称。     | 当前名称的英文整体读音来源；未逐项查证官方语音或整体 IPA。 |
| ZD148 | [raw-prime-meat-zelda](../entries/raw-prime-meat-zelda.yml) / Raw Prime Meat          | 高级兽肉；《王国之泪》中的材料名称。     | 当前名称的英文整体读音来源；未逐项查证官方语音或整体 IPA。 |
| ZD163 | [silent-princess-zelda](../entries/silent-princess-zelda.yml) / Silent Princess       | 宁静公主；《王国之泪》中的材料名称。     | 当前名称的英文整体读音来源；未逐项查证官方语音或整体 IPA。 |
| ZD165 | [skyshroom-zelda](../entries/skyshroom-zelda.yml) / Skyshroom                         | 天空蘑菇；《王国之泪》中的材料名称。     | 当前名称的英文整体读音来源；未逐项查证官方语音或整体 IPA。 |
| ZD175 | [spicy-pepper-zelda](../entries/spicy-pepper-zelda.yml) / Spicy Pepper                | 暖暖草果；《王国之泪》中的材料名称。     | 当前名称的英文整体读音来源；未逐项查证官方语音或整体 IPA。 |
| ZD181 | [sundelion-zelda](../entries/sundelion-zelda.yml) / Sundelion                         | 向阳草；《王国之泪》中的材料名称。       | 当前名称的英文整体读音来源；未逐项查证官方语音或整体 IPA。 |

**其他游戏名称: 17 项。**

| 编号  | 稳定 ID / Word                                                                        | 当前 Meaning                                                 | 待补证据                                                   |
|:------|:--------------------------------------------------------------------------------------|:-------------------------------------------------------------|:-----------------------------------------------------------|
| ZD034 | [autobuild-zelda](../entries/autobuild-zelda.yml) / Autobuild                         | 蓝图；复现已记录组合设计的能力。                             | 当前名称的英文整体读音来源；未逐项查证官方语音或整体 IPA。 |
| ZD052 | [demon-dragon-zelda](../entries/demon-dragon-zelda.yml) / Demon Dragon                | 黑龙；加侬多夫龙化后的形态。                                 | 当前名称的英文整体读音来源；未逐项查证官方语音或整体 IPA。 |
| ZD054 | [demon-king-zelda](../entries/demon-king-zelda.yml) / Demon King                      | 魔王；加侬多夫在本作中的称号。                               | 当前名称的英文整体读音来源；未逐项查证官方语音或整体 IPA。 |
| ZD057 | [draconification-zelda](../entries/draconification-zelda.yml) / draconification       | 龙化之法；通过吞下秘石化为不死之龙、同时失去自我的禁忌之术。 | 当前名称的英文整体读音来源；未逐项查证官方语音或整体 IPA。 |
| ZD115 | [korok-seed-zelda](../entries/korok-seed-zelda.yml) / Korok Seed                      | 克洛格的果实；与克洛格和背包扩充有关的收集物品。             | 当前名称的英文整体读音来源；未逐项查证官方语音或整体 IPA。 |
| ZD117 | [light-dragon-zelda](../entries/light-dragon-zelda.yml) / Light Dragon                | 白龙；本作在空中飞行的一条龙。                               | 当前名称的英文整体读音来源；未逐项查证官方语音或整体 IPA。 |
| ZD118 | [light-of-blessing-zelda](../entries/light-of-blessing-zelda.yml) / Light of Blessing | 祝福之光；通过神庙挑战后获得的重要物品。                     | 当前名称的英文整体读音来源；未逐项查证官方语音或整体 IPA。 |
| ZD120 | [lightroot-zelda](../entries/lightroot-zelda.yml) / Lightroot                         | 破魔之根；地底可激活并照亮周围的特殊根状地点。               | 当前名称的英文整体读音来源；未逐项查证官方语音或整体 IPA。 |
| ZD139 | [phantom-ganon-zelda](../entries/phantom-ganon-zelda.yml) / Phantom Ganon             | 幻影加侬；本作中加侬多夫的幻影敌人。                         | 当前名称的英文整体读音来源；未逐项查证官方语音或整体 IPA。 |
| ZD143 | [purah-pad-zelda](../entries/purah-pad-zelda.yml) / Purah Pad                         | 普尔亚平板；林克使用的多功能装置。                           | 当前名称的英文整体读音来源；未逐项查证官方语音或整体 IPA。 |
| ZD155 | [secret-stone-zelda](../entries/secret-stone-zelda.yml) / Secret Stone                | 秘石；本作中可放大持有者力量的重要物品。                     | 当前名称的英文整体读音来源；未逐项查证官方语音或整体 IPA。 |
| ZD164 | [sky-island-zelda](../entries/sky-island-zelda.yml) / sky island                      | 空岛；悬浮在空中的岛屿。                                     | 当前名称的英文整体读音来源；未逐项查证官方语音或整体 IPA。 |
| ZD166 | [skyview-tower-zelda](../entries/skyview-tower-zelda.yml) / Skyview Tower             | 鸟望台；用于调查区域地图的塔状设施。                         | 当前名称的英文整体读音来源；未逐项查证官方语音或整体 IPA。 |
| ZD185 | [the-depths-zelda](../entries/the-depths-zelda.yml) / The Depths                      | 地底；本作位于地表下方的广大区域。                           | 当前名称的英文整体读音来源；未逐项查证官方语音或整体 IPA。 |
| ZD189 | [ultrahand-zelda](../entries/ultrahand-zelda.yml) / Ultrahand                         | 究极手；可抓取、移动和粘接许多物体的能力。                   | 当前名称的英文整体读音来源；未逐项查证官方语音或整体 IPA。 |
| ZD197 | [zonai-device-zelda](../entries/zonai-device-zelda.yml) / Zonai Device                | 左纳乌装置；使用左纳乌技术的功能装置。                       | 当前名称的英文整体读音来源；未逐项查证官方语音或整体 IPA。 |
| ZD199 | [zonaite-zelda](../entries/zonaite-zelda.yml) / Zonaite                               | 左纳尼乌姆；用于左纳乌技术相关制作和交换的矿石材料。         | 当前名称的英文整体读音来源；未逐项查证官方语音或整体 IPA。 |

## 已有音标保留清单

规划时已有的 92 项按原值保留。这里只盘点现有字段，不代表重新查证每个来源或确认录音与文字一致；原词条字节及历史验证记录保持不变。

| 稳定 ID / Word                                                        | 当前 Meaning                     | 保留 pronunciation |
|:----------------------------------------------------------------------|:---------------------------------|:-------------------|
| [abandoned-workflow](../entries/abandoned-workflow.yml) / abandoned   | 被遗弃的；无人继续处理的         | `/əˈbændənd/`      |
| [apple-food](../entries/apple-food.yml) / apple                       | 苹果。                           | `/ˈæpəl/`          |
| [arrow-weapon](../entries/arrow-weapon.yml) / arrow                   | 箭                               | `/ˈæroʊ/`          |
| [ascend-general](../entries/ascend-general.yml) / ascend              | 上升；向高处移动。               | `/əˈsɛnd/`         |
| [assemble-general](../entries/assemble-general.yml) / assemble        | 组装                             | `/əˈsɛmbəl/`       |
| [attach-general](../entries/attach-general.yml) / attach              | 连接、附加                       | `/əˈtætʃ/`         |
| [axe-weapon](../entries/axe-weapon.yml) / axe                         | 斧头                             | `/æks/`            |
| [balloon-engineering](../entries/balloon-engineering.yml) / balloon   | 气球。                           | `/bəˈlun/`         |
| [banana-food](../entries/banana-food.yml) / banana                    | 香蕉。                           | `/bəˈnænə/`        |
| [bass-food](../entries/bass-food.yml) / bass                          | 鲈鱼。                           | `/bæs/`            |
| [bee-general](../entries/bee-general.yml) / bee                       | 蜂；蜜蜂。                       | `/bi/`             |
| [bow-weapon](../entries/bow-weapon.yml) / bow                         | 弓                               | `/boʊ/`            |
| [brave-general](../entries/brave-general.yml) / brave                 | 勇敢的                           | `/breɪv/`          |
| [bridge-architecture](../entries/bridge-architecture.yml) / bridge    | 桥梁                             | `/brɪdʒ/`          |
| [butter-food](../entries/butter-food.yml) / butter                    | 黄油。                           | `/ˈbʌtər/`         |
| [cancel-general](../entries/cancel-general.yml) / cancel              | 取消                             | `/ˈkænsəl/`        |
| [carrot-food](../entries/carrot-food.yml) / carrot                    | 胡萝卜。                         | `/ˈkærət/`         |
| [castle-architecture](../entries/castle-architecture.yml) / castle    | 城堡                             | `/ˈkæsəl/`         |
| [cave-architecture](../entries/cave-architecture.yml) / cave          | 洞穴                             | `/keɪv/`           |
| [ceiling-architecture](../entries/ceiling-architecture.yml) / ceiling | 天花板                           | `/ˈsilɪŋ/`         |
| [cheese-food](../entries/cheese-food.yml) / cheese                    | 奶酪。                           | `/tʃiz/`           |
| [chilly-general](../entries/chilly-general.yml) / chilly              | 寒冷的；凉飕飕的。               | `/ˈtʃɪli/`         |
| [cliff-terrain](../entries/cliff-terrain.yml) / cliff                 | 悬崖                             | `/klɪf/`           |
| [climb-general](../entries/climb-general.yml) / climb                 | 攀爬                             | `/klaɪm/`          |
| [curious-general](../entries/curious-general.yml) / curious           | 好奇的                           | `/ˈkjʊriəs/`       |
| [descend-general](../entries/descend-general.yml) / descend           | 下降                             | `/dɪˈsɛnd/`        |
| [destroy-general](../entries/destroy-general.yml) / destroy           | 摧毁；毁灭。                     | `/dɪˈstrɔɪ/`       |
| [detach-general](../entries/detach-general.yml) / detach              | 分离、拆卸                       | `/dɪˈtætʃ/`        |
| [door-architecture](../entries/door-architecture.yml) / door          | 门。                             | `/dɔr/`            |
| [drop-general](../entries/drop-general.yml) / drop                    | 丢弃                             | `/drɑp/`           |
| [durable-general](../entries/durable-general.yml) / durable           | 耐用的；能经受使用而不易损坏的。 | `/ˈdʊrəbəl/`       |
| [eagle-general](../entries/eagle-general.yml) / eagle                 | 鹰。                             | `/ˈiɡəl/`          |
| [egg-food](../entries/egg-food.yml) / egg                             | 蛋。                             | `/ɛɡ/`             |
| [examine-general](../entries/examine-general.yml) / examine           | 检查                             | `/ɪɡˈzæmən/`       |
| [factory-architecture](../entries/factory-architecture.yml) / factory | 工厂。                           | `/ˈfæktəri/`       |
| [fan-engineering](../entries/fan-engineering.yml) / fan               | 风扇；扇子。                     | `/fæn/`            |
| [find-general](../entries/find-general.yml) / find                    | 找到；发现。                     | `/faɪnd/`          |
| [fire-general](../entries/fire-general.yml) / fire                    | 火。                             | `/faɪr/`           |
| [flame-general](../entries/flame-general.yml) / flame                 | 火焰。                           | `/fleɪm/`          |
| [forest-terrain](../entries/forest-terrain.yml) / forest              | 森林                             | `/ˈfɔrɪst/`        |
| [forget-general](../entries/forget-general.yml) / forget              | 忘记；未能想起。                 | `/fərˈɡɛt/`        |
| [frog-general](../entries/frog-general.yml) / frog                    | 青蛙。                           | `/frɔɡ/`           |
| [fruit-food](../entries/fruit-food.yml) / fruit                       | 水果。                           | `/frut/`           |
| [gate-architecture](../entries/gate-architecture.yml) / gate          | 大门                             | `/ɡeɪt/`           |
| [golden-general](../entries/golden-general.yml) / golden              | 金色的；金黄色的。               | `/ˈɡoʊldən/`       |
| [hand-general](../entries/hand-general.yml) / hand                    | 手。                             | `/hænd/`           |
| [head-general](../entries/head-general.yml) / head                    | 头；头部。                       | `/hɛd/`            |
| [hidden-general](../entries/hidden-general.yml) / hidden              | 隐藏的                           | `/ˈhɪdən/`         |
| [hold-general](../entries/hold-general.yml) / hold                    | 手持、持有                       | `/hoʊld/`          |
| [honey-food](../entries/honey-food.yml) / honey                       | 蜂蜜。                           | `/ˈhʌni/`          |
| [island-terrain](../entries/island-terrain.yml) / island              | 岛屿                             | `/ˈaɪlənd/`        |
| [jar-general](../entries/jar-general.yml) / jar                       | 罐；广口瓶。                     | `/dʒɑr/`           |
| [lake-terrain](../entries/lake-terrain.yml) / lake                    | 湖泊                             | `/leɪk/`           |
| [loyal-general](../entries/loyal-general.yml) / loyal                 | 忠诚的                           | `/ˈlɔɪəl/`         |
| [map-ui](../entries/map-ui.yml) / map                                 | 地图。                           | `/mæp/`            |
| [meat-food](../entries/meat-food.yml) / meat                          | 肉；供食用的肉类。               | `/mit/`            |
| [menu-ui](../entries/menu-ui.yml) / menu                              | 菜单；可供选择的项目列表。       | `/ˈmɛnju/`         |
| [milk-food](../entries/milk-food.yml) / milk                          | 奶；乳汁。                       | `/mɪlk/`           |
| [mirror-general](../entries/mirror-general.yml) / mirror              | 镜子。                           | `/ˈmɪrər/`         |
| [mountain-terrain](../entries/mountain-terrain.yml) / mountain        | 山；山岳。                       | `/ˈmaʊntən/`       |
| [mushroom-food](../entries/mushroom-food.yml) / mushroom              | 蘑菇                             | `/ˈmʌʃˌrum/`       |
| [oil-food](../entries/oil-food.yml) / oil                             | 油。                             | `/ɔɪl/`            |
| [pot-general](../entries/pot-general.yml) / pot                       | 锅；罐。                         | `/pɑt/`            |
| [powerful-general](../entries/powerful-general.yml) / powerful        | 强大的                           | `/ˈpaʊərfəl/`      |
| [raw-food](../entries/raw-food.yml) / raw                             | 生的；未烹饪的。                 | `/rɔ/`             |
| [restore-tech](../entries/restore-tech.yml) / restore                 | 恢复                             | `/rɪˈstɔr/`        |
| [rice-food](../entries/rice-food.yml) / rice                          | 米；米饭。                       | `/raɪs/`           |
| [ring-general](../entries/ring-general.yml) / ring                    | 环；圈；戒指。                   | `/rɪŋ/`            |
| [river-terrain](../entries/river-terrain.yml) / river                 | 河流                             | `/ˈrɪvər/`         |
| [root-general](../entries/root-general.yml) / root                    | 根；植物的根。                   | `/rut/`            |
| [rusty-general](../entries/rusty-general.yml) / rusty                 | 生锈的                           | `/ˈrʌsti/`         |
| [salt-food](../entries/salt-food.yml) / salt                          | 盐。                             | `/sɔlt/`           |
| [seed-general](../entries/seed-general.yml) / seed                    | 种子。                           | `/sid/`            |
| [select-general](../entries/select-general.yml) / select              | 选择                             | `/səˈlɛkt/`        |
| [shield-weapon](../entries/shield-weapon.yml) / shield                | 盾牌                             | `/ʃild/`           |
| [spicy-general](../entries/spicy-general.yml) / spicy                 | 辛辣的；带香辛料味道的。         | `/ˈspaɪsi/`        |
| [stick-general](../entries/stick-general.yml) / stick                 | 棍；条状或杆状物。               | `/stɪk/`           |
| [stone-general](../entries/stone-general.yml) / stone                 | 石头；石料。                     | `/stoʊn/`          |
| [sturdy-general](../entries/sturdy-general.yml) / sturdy              | 坚固的                           | `/ˈstɜrdi/`        |
| [sugar-food](../entries/sugar-food.yml) / sugar                       | 糖。                             | `/ˈʃʊɡər/`         |
| [sunny-general](../entries/sunny-general.yml) / sunny                 | 阳光明媚的；充满阳光的。         | `/ˈsʌni/`          |
| [sword-weapon](../entries/sword-weapon.yml) / sword                   | 剑                               | `/sɔrd/`           |
| [tear-general](../entries/tear-general.yml) / tear                    | 眼泪；泪水。                     | `/tɪr/`            |
| [throw-general](../entries/throw-general.yml) / throw                 | 投掷                             | `/θroʊ/`           |
| [tomato-food](../entries/tomato-food.yml) / tomato                    | 番茄。                           | `/təˈmeɪtoʊ/`      |
| [tough-general](../entries/tough-general.yml) / tough                 | 坚韧的；结实的；难以对付的。     | `/tʌf/`            |
| [tower-architecture](../entries/tower-architecture.yml) / tower       | 高塔                             | `/ˈtaʊər/`         |
| [valley-terrain](../entries/valley-terrain.yml) / valley              | 山谷                             | `/ˈvæli/`          |
| [village-architecture](../entries/village-architecture.yml) / village | 村庄                             | `/ˈvɪlɪdʒ/`        |
| [water-general](../entries/water-general.yml) / water                 | 水。                             | `/ˈwɔtər/`         |
| [wheat-food](../entries/wheat-food.yml) / wheat                       | 小麦。                           | `/wit/`            |
| [wise-general](../entries/wise-general.yml) / wise                    | 睿智的                           | `/waɪz/`           |

## 听审与确认边界

- 本轮没有听审词典、游戏官方配音或现有录音，没有进行 MP3 解码或 Anki 操作。特别是 desert、wind、entrance、construct 动词与名词、upgrade 等可能由 TTS 选错义项或重音的词，人工听审仍需单独完成；未发现或断言任何已确认录音错误。
- 前批 tear、bass 和 bow 的听审待办保留，见[第四批](EN_PRONUNCIATION_ENRICHMENT_BATCH_04_PLAN.md#听审与暂缓边界)与[第五批](EN_PRONUNCIATION_ENRICHMENT_BATCH_05_PLAN.md#听审与暂缓边界)。新增音标不授权自动重录。
- 用户已明确批准 ZP001–ZP259 全部具体音标及对应 Notes 来源行；本轮严格按批准文案写入，不扩大其他项目的改写范围。
- ZD001–ZD202 是独立待核实清单，没有可直接批准写入的音标。提供读音证据后，应先形成具体 IPA 与来源文案，再确认实施。
- 本计划没有提出新关联或助记，不修改现有 RelatedWords 派生关系，不新增词条。音频、语言、词义、模板和实现改动均不在批准范围内。

## 批准后的实施与验证

1. 重新读取获批计划、当前 AGENTS.md、字段规范、相关实现及 Git 状态。按稳定 ID 检查规划以来的并发变动，核对批准编号、Word、词性、Meaning、口音、完整 IPA、来源分区与 Notes 文案。已有不同音标或来源时先协调，不覆盖新内容。
2. 在仓库外保存所有词条原文及解析值、当前完整 TSV、现有和归档 MP3 哈希。只写入批准的 pronunciation 值和 Notes 追加行；保留全部旧 Notes，完整来源行只追加一次。
3. 保留 ID 集、Word、词义、语言、词性、例句、辅助字段、关系、优先级、先修、标签和音频引用。不新增词条、媒体，不改 schema、CLI、模板、依赖或实现，不替换录音。
4. 使用现有命令进行检查和构建，不手改 dist。

```bash
go test ./...
go vet ./...
go run scripts/validate/main.go
go run scripts/build/main.go
git diff --check
```

5. 源未变化时再构建一次，比较字节。验证实际实施基线 ID 集、数据行数及 17 列 TSV；只允许获批行的 Pronunciation、Notes 列变化。Audio、StudyOrder、RelatedWords、标签和其余列不变，未批准行的导出不变。若全库已变化，使用重新核实的基线，不强套本计划历史计数。
6. 逐项比较批准的确切值，检查默认美式无 US 前缀、英式保留 UK、Notes 不重复。未改词条保留原字节；所有既有及归档媒体哈希保持不变。检查本地链接、锚点、格式和未跟踪文本，更新本计划的批准、实施与实际验证状态。
7. 分别报告静态和导出检查、MP3 解码、人工听审、真实 Anki 导入、显示及播放，不声称执行未做过的验证。保留完成部分和无关工作，报告失败及未解决项目。
8. 不暂存、提交、推送、复制媒体到 Anki 或修改 Anki 集合。Anki 继续管理复习调度和历史。

## 规划检查

- 553 个范围内稳定 ID 恰好分为 92 个保留、259 个具体提案和 202 个待核实，没有重叠或遗漏。259 项包含 258 个美式和 1 个英式值，均为现有缺失字段的新增提案。当前源中的其他字段与实际路径均逐项核对。
- 普通目标的 284 个不同词形已尝试对应 Collins 读取；三个同形普通名词的游戏名称另按字典核对，geoglyph 采用已读 Wiktionary。待核实项如实区分未提供 IPA、义项或词形不符、来源不可访问，以及尚未逐项研究官方语音。没有把页面访问失败当作不存在读音的证据。
- 本计划的本地链接、锚点、编号、表格、围栏、空白及最终换行检查通过，git diff --check 通过。未跟踪计划另做文本检查。最终来源均记录 URL、分区、口音及日期；外部页面未来是否可用不在本次静态链接检查范围内。
- 与规划前快照相比，所有既有项目文件字节保持不变，包括全部 1,320 个词条、1,320 个现有 MP3、43 个归档 MP3、前批计划、当前未提交工作、规范、README 模板、实现、依赖及 dist。唯一新增文件为本计划。
- 本轮没有运行 Go 测试、词库验证、构建、音频生成、解码、人工听审或真实 Anki 操作。上述命令仅为获得实施批准后的程序。

## 实施与验证记录

- 用户确认 "ZP001–ZP259；implement it"，授权全部 259 项具体音标和对应 Notes 来源行。实施前重新核对 Git 状态、当前规范、实现、稳定 ID、Word、词性、Meaning、IPA、口音及来源；规划以来的既有文件字节均未变化。
- 在仓库外保存全部原词条及解析值、当前完整 TSV、批准计划原文和所有既有项目文件哈希，再按批准编号实施。259 项来源行各追加一次；其中 3 项原无 Notes，因此创建该字段，其余旧 Notes 全文保留。
- 本批 258 项美式值均无 US 前缀；hideout 的英式值保留 UK。既有 cash account 的 UK 音标及其他已有音标不变。Construct、Gloom、Paraglider 保留批准的普通名词读音和来源限定，不声称核实游戏配音。
- 全库仍为 1,320 项，英语 1,077 项、日语 243 项。英语已有音标 686 项、缺失 391 项；553 个塞尔达词条中已有音标 351 项、缺失 202 项。ZD001–ZD202 全部保持原状，范围外缺失仍为 189 项。
- 与实施前快照相比，只有本批 259 个词条的 pronunciation 和对应 Notes 来源变化。其他 1,061 个词条字节不变；所有稳定 ID、语言、Word、Meaning、词性、旧 Notes、Context、Example、辅助字段、关系、学习优先级、先修、标签及音频引用保留。规范、实现、模板、依赖、前批计划和无关未提交工作均未改动。
- 1,320 个现有 MP3 和 43 个归档 MP3 的 SHA-256 全部保持不变，没有生成、替换或复制录音。
- go test ./...、go vet ./...、词库验证和 git diff --check 通过。验证覆盖 1,320 个词条；15 条既有同词警告保留，没有新增同词项。未跟踪计划另查格式、本地链接、锚点、编号、来源文案和待核实清单。
- 两次构建字节完全一致，均为 1,320 个数据行、17 列。相较实施前，仅获批 259 行的 Pronunciation、Notes 列变化；所有其他行和列不变，包括 Audio、StudyOrder、RelatedWords、标签和 TTSLanguage。导出 SHA-256: `ea1f2890d38917d6d87a387a7ad66d40c383ac2b47cf61385769e1cd40482371`。
- 未进行 MP3 解码、人工听审、游戏官方配音转录或真实 Anki 导入、显示、播放验证。已有听审待办继续保留，没有暂存、提交、推送或修改 Anki 集合。
