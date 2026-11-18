# 英语音标补全计划: 第二批普通词与常用短语

状态: **B02-P01–B02-P24 已批准并实施，静态及导出验证通过。** 日期: 2026-11-18。

规划阶段仅授权创建第二批计划并更新维护规范。用户随后确认 B02-P01–B02-P24，并明确指示实施已确认项目。现已更新 24 个词条并重建导出；录音、归档、实现、模板、依赖及其他原有工作保持不变。未暂存、提交、推送或操作 Anki，实际结果见[实施记录](#实施记录)。

[第一批计划](EN_PRONUNCIATION_ENRICHMENT_BATCH_01_PLAN.md)保留为已完成的历史记录，不补写新规则或重写已实施值。第二批使用 B02-P01–B02-P24，避免与第一批 P01–P22 混淆。

## 导航

- [范围与基线](#范围与基线)
- [来源和口音规则](#来源和口音规则)
- [具体音标提案](#具体音标提案)
- [逐项证据与限制](#逐项证据与限制)
- [Notes 来源追加文案](#notes-来源追加文案)
- [暂缓项目](#暂缓项目)
- [批准后的实施与验证](#批准后的实施与验证)
- [规划检查](#规划检查)
- [实施记录](#实施记录)

## 范围与基线

以下保留 2026-11-18 实施前的规划基线。以实际稳定 ID 盘点，已读取当时所有词条。对本批目标核对了 Word、词性、Meaning、Context、Example、Notes、现有辅助字段、学习元数据及音频引用。全库盘点不表示已经查证全部缺失音标。

| 项目                          | 数量或范围                                  |
|:------------------------------|:--------------------------------------------|
| 全库词条                      | 1,320                                       |
| 英语 / 日语                   | 1,077 / 243                                 |
| 英语已有 / 缺失 pronunciation | 313 / 764                                   |
| 日语已有 / 缺失 pronunciation | 243 / 0                                     |
| 本批已检查目标                | 26 个稳定 ID                                |
| 本批可审阅提案                | 24 项，23 个单词和 1 个整体短语             |
| 口音                          | 23 项 US、1 项 UK                           |
| 暂缓                          | 2 项整体短语                                |
| 拟修改字段                    | 24 个词条的 pronunciation 及 Notes 来源追加 |
| 新增词条 / 录音 / 关系 / 先修 | 全部为 0                                    |

选择常见食物、语言、基础资金词及日常开发中使用的普通词，按学习价值分批处理。branch、build、cache 的读音有可访问证据，因此纳入本批；词典的普通词释义不替代现有技术释义。词长和字母顺序不决定是否补充，也不影响学习顺序。

规划时预计全部实施本批 24 项后，英语已有音标增至 337 项，仍缺失 740 项。暂缓的 2 项包含在这 740 项内，不重复计算。其他 738 项未逐项研究，不表示应继续留空或不适合补充。

## 来源和口音规则

按已确认的规则，优先查证美式，合理检索后可靠美式证据暂不可得时接受英式。新增值明确标注 US 或 UK，标签只描述显示的转写。用户接受英式 IPA 与原有美式合成录音并存，不能因此自动替换录音。

- 本批 23 个单词采用已打开并阅读的 Collins Dictionary 美式正文。具体分区署名为 Webster’s New World College Dictionary, 5th Digital Edition。保留来源音素、重音及记法，不混入另一口音的转写。
- cash account 的英式分区明确列出整体音标；同页美式分区没有呈现整体 IPA，Cambridge 美式页面返回 403。因此提出 UK 标注的可靠整体读音，不将访问失败说成美式读音不存在，也不自行拼接 cash 与 account。
- 本批只有 cash account 选取 UK 值；这是本轮可得证据的选择，不是断言它与美式读音一定不同。其他正常变体可存在，表内只选一个适合当前词形和词性的形式。
- 所有表内最终来源均为实际访问的词典正文，不以搜索摘要代替。访问日为 2026-11-18；页面内容可能变化，分区名称用于再次定位，网页行号不作为永久引用。
- 没有收听词典或现有 MP3，没有声称录音与所选 IPA 完全吻合。没有改写来源未提供的音素、弱读或连读。

规范已更新至 [AGENTS.md](../AGENTS.md#pronunciation)、[字段指南](ENTRY_SCHEMA.md#pronunciation)及[维护提示词](MAINTENANCE_PROMPTS.md)。pronunciation 仍为可选的普通字符串，不新增口音字段、解析规则或主题脚本。原有未标口音的音标、第一批值和日语假名保持不变。

## 具体音标提案

2026-11-18，用户确认 B02-P01–B02-P24。按下文逐项批准范围，确认内容包括完整 pronunciation 值及对应 Notes 来源追加文案。随后已按明确实施指令写入词条，B02-D01 和 B02-D02 继续暂缓。

下表 Word、词性和 Meaning 保留当前原文，仅用于匹配，不授权改写。24 项 pronunciation 在实施前均省略，已按 add 动作写入。每项批准范围包含该值及下文同编号的完整 Notes 来源行；不授权修改其他字段。

| 编号    | 稳定 ID / Word                                                             | 当前词性   | 当前 Meaning                                       | 拟写 pronunciation | 来源                                                                                       |
|:--------|:---------------------------------------------------------------------------|:-----------|:---------------------------------------------------|:-------------------|:-------------------------------------------------------------------------------------------|
| B02-P01 | [beef-food](../entries/beef-food.yml) / beef                               | noun       | 牛肉                                               | `US /bif/`         | [Collins: beef](https://www.collinsdictionary.com/dictionary/english/beef)                 |
| B02-P02 | [chicken-food](../entries/chicken-food.yml) / chicken                      | noun       | 鸡肉                                               | `US /ˈtʃɪkən/`     | [Collins: chicken](https://www.collinsdictionary.com/dictionary/english/chicken)           |
| B02-P03 | [fish-food](../entries/fish-food.yml) / fish                               | noun       | 鱼；鱼肉                                           | `US /fɪʃ/`         | [Collins: fish](https://www.collinsdictionary.com/dictionary/english/fish)                 |
| B02-P04 | [pork-food](../entries/pork-food.yml) / pork                               | noun       | 猪肉                                               | `US /pɔrk/`        | [Collins: pork](https://www.collinsdictionary.com/dictionary/english/pork)                 |
| B02-P05 | [potato-food](../entries/potato-food.yml) / potato                         | noun       | 土豆；马铃薯                                       | `US /pəˈteɪtoʊ/`   | [Collins: potato](https://www.collinsdictionary.com/dictionary/english/potato)             |
| B02-P06 | [sauce-food](../entries/sauce-food.yml) / sauce                            | noun       | 酱汁                                               | `US /sɔs/`         | [Collins: sauce](https://www.collinsdictionary.com/dictionary/english/sauce)               |
| B02-P07 | [pepper-food](../entries/pepper-food.yml) / pepper                         | noun       | 胡椒；胡椒粉                                       | `US /ˈpɛpər/`      | [Collins: pepper](https://www.collinsdictionary.com/dictionary/english/pepper)             |
| B02-P08 | [pasta-food](../entries/pasta-food.yml) / pasta                            | noun       | 意大利面；意面                                     | `US /ˈpɑstə/`      | [Collins: pasta](https://www.collinsdictionary.com/dictionary/english/pasta)               |
| B02-P09 | [language-general](../entries/language-general.yml) / language             | noun       | 语言                                               | `US /ˈlæŋɡwɪdʒ/`   | [Collins: language](https://www.collinsdictionary.com/dictionary/english/language)         |
| B02-P10 | [asset-finance](../entries/asset-finance.yml) / asset                      | noun       | 资产；具有经济价值、由个人或机构持有的资源。       | `US /ˈæsɛt/`       | [Collins: asset](https://www.collinsdictionary.com/dictionary/english/asset)               |
| B02-P11 | [interest-finance](../entries/interest-finance.yml) / interest             | noun       | 利息；因存放或借用资金产生的收益或费用。           | `US /ˈɪntrɪst/`    | [Collins: interest](https://www.collinsdictionary.com/dictionary/english/interest)         |
| B02-P12 | [net-finance](../entries/net-finance.yml) / net                            | adjective  | 净额的；扣除相关项目后剩余的。                     | `US /nɛt/`         | [Collins: net](https://www.collinsdictionary.com/dictionary/english/net)                   |
| B02-P13 | [order-finance](../entries/order-finance.yml) / order                      | noun       | 委托单；提交给交易系统的买卖指令。                 | `US /ˈɔrdər/`      | [Collins: order](https://www.collinsdictionary.com/dictionary/english/order)               |
| B02-P14 | [quarter-finance](../entries/quarter-finance.yml) / quarter                | noun       | 季度；一年中的四分之一，在财报中表示一个报告季度。 | `US /ˈkwɔrtər/`    | [Collins: quarter](https://www.collinsdictionary.com/dictionary/english/quarter)           |
| B02-P15 | [share-finance](../entries/share-finance.yml) / share                      | noun       | 股份；公司资本或某项权益中的一份。                 | `US /ʃɛr/`         | [Collins: share](https://www.collinsdictionary.com/dictionary/english/share)               |
| B02-P16 | [stock-finance](../entries/stock-finance.yml) / stock                      | noun       | 股票；代表公司所有权份额的证券。                   | `US /stɑk/`        | [Collins: stock](https://www.collinsdictionary.com/dictionary/english/stock)               |
| B02-P17 | [trade-finance](../entries/trade-finance.yml) / trade                      | noun       | 交易；买卖证券或其他金融工具的行为。               | `US /treɪd/`       | [Collins: trade](https://www.collinsdictionary.com/dictionary/english/trade)               |
| B02-P18 | [transfer-finance](../entries/transfer-finance.yml) / transfer             | verb       | 划转；将资金从一个账户或位置转到另一个。           | `US /trænsˈfɜr/`   | [Collins: transfer](https://www.collinsdictionary.com/dictionary/english/transfer)         |
| B02-P19 | [review-workflow](../entries/review-workflow.yml) / review                 | noun; verb | 审查；检查工作内容并给出判断或反馈。               | `US /rɪˈvju/`      | [Collins: review](https://www.collinsdictionary.com/dictionary/english/review)             |
| B02-P20 | [implement-tech](../entries/implement-tech.yml) / implement                | verb       | 实现；把设计、规则或接口落实为可执行行为。         | `US /ˈɪmpləˌmɛnt/` | [Collins: implement](https://www.collinsdictionary.com/dictionary/english/implement)       |
| B02-P21 | [branch-git](../entries/branch-git.yml) / branch                           | noun       | 分支；用于独立推进一条开发历史的引用。             | `US /bræntʃ/`      | [Collins: branch](https://www.collinsdictionary.com/dictionary/english/branch)             |
| B02-P22 | [build-tech](../entries/build-tech.yml) / build                            | noun; verb | 构建；执行编译、打包等步骤生成项目产物。           | `US /bɪld/`        | [Collins: build](https://www.collinsdictionary.com/dictionary/english/build)               |
| B02-P23 | [cache-tech](../entries/cache-tech.yml) / cache                            | noun       | 缓存；暂存数据或计算结果以供后续复用的机制。       | `US /kæʃ/`         | [Collins: cache](https://www.collinsdictionary.com/dictionary/english/cache)               |
| B02-P24 | [cash-account-finance](../entries/cash-account-finance.yml) / cash account | noun       | 现金账户；证券买入需要由账户资金全额支付的账户。   | `UK /kæʃ əˈkaʊnt/` | [Collins: cash account](https://www.collinsdictionary.com/dictionary/english/cash-account) |

## 逐项证据与限制

这些说明是审阅依据，不是另加助记或改写 Notes 的授权。只有下一节的完整来源行会在实施时追加。

| 编号    | 读音选择及意义匹配                                                                                                                                    |
|:--------|:------------------------------------------------------------------------------------------------------------------------------------------------------|
| B02-P01 | 牛肉义名词，不采用俚语或投诉义；保留来源 /i/。                                                                                                        |
| B02-P02 | 鸡肉义名词，不采用胆小鬼义。                                                                                                                          |
| B02-P03 | 鱼及鱼肉义名词；不把 fishes 的词尾加到当前形式。                                                                                                      |
| B02-P04 | 猪肉义名词，不采用政治拨款的引申义。                                                                                                                  |
| B02-P05 | 单数名词；选择所列完整词尾变体，不使用 potatoes 的复数转写。                                                                                          |
| B02-P06 | 酱汁义名词；保留来源美式元音，不替换为英式转写。                                                                                                      |
| B02-P07 | 胡椒调味料义名词；该来源也包含其他义项，本批不修改释义或关联。                                                                                        |
| B02-P08 | 意面义名词，使用来源 American English 分区，不从英式形式推导。                                                                                        |
| B02-P09 | 语言义名词。                                                                                                                                          |
| B02-P10 | 有交换价值的资产义名词，保留单数形式。                                                                                                                |
| B02-P11 | 利息义名词；选择所列双音节变体，其他列出的美式变体并非错误。                                                                                          |
| B02-P12 | 对应 net 2 的净额形容词，不使用 .net 域名读法或网状物释义。                                                                                           |
| B02-P13 | 委托指令义名词；本批只补普通词的对应读音，不补交易机制。                                                                                              |
| B02-P14 | 一年中四分之一的季度义名词，不改写成复数。                                                                                                            |
| B02-P15 | 股份义名词，对应 share 1，不采用 share 2 的犁铧释义。                                                                                                 |
| B02-P16 | 股票义名词；采用美式正文，不混入英式 /ɒ/。                                                                                                            |
| B02-P17 | 交易义名词，保留原形。                                                                                                                                |
| B02-P18 | 当前为动词，选择来源明确列出的后重音变体；不把前重音变体一概判为错误。                                                                                |
| B02-P19 | 采用审查或评价用法的常规读音；来源对特定再次查看用法另列变体，不套用于本条。                                                                          |
| B02-P20 | 当前为动词，使用来源 for v. 指定的形式，不使用工具义名词音标。                                                                                        |
| B02-P21 | 选择所列美式变体之一；词典核实普通词读音，不将其计算机跳转释义当作 Git 分支定义。                                                                     |
| B02-P22 | 原词性 noun; verb，采用同形词读音；词典来源不替代仓库的软件构建释义。                                                                                 |
| B02-P23 | 缓存义名词，来源包含 computing 用法；不按字母逐个读。                                                                                                 |
| B02-P24 | 采用来源 British English 分区的完整短语读音。Collins 美式分区未呈现整体 IPA，Cambridge 美式页面返回 403；记录为本轮证据不可得，不声称美式读法不存在。 |

## Notes 来源追加文案

保留完整原 Notes。已有 Notes 时，在末尾追加对应来源行；缺少 Notes 时才创建。来源行已存在就不再追加；相同 URL 若已用于其他用途，应先核对是否已记录口音和发音证据，避免重复说明。以下纯文本是逐项拟追加的完整文案，围栏不写入字段。

### B02-P01 — beef

```plain
Pronunciation source: Collins Dictionary — beef; beef in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/beef (accent: US; accessed 2026-11-18).
```

### B02-P02 — chicken

```plain
Pronunciation source: Collins Dictionary — chicken; chicken in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/chicken (accent: US; accessed 2026-11-18).
```

### B02-P03 — fish

```plain
Pronunciation source: Collins Dictionary — fish; fish in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/fish (accent: US; accessed 2026-11-18).
```

### B02-P04 — pork

```plain
Pronunciation source: Collins Dictionary — pork; pork in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/pork (accent: US; accessed 2026-11-18).
```

### B02-P05 — potato

```plain
Pronunciation source: Collins Dictionary — potato; potato in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/potato (accent: US; accessed 2026-11-18).
```

### B02-P06 — sauce

```plain
Pronunciation source: Collins Dictionary — sauce; sauce in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/sauce (accent: US; accessed 2026-11-18).
```

### B02-P07 — pepper

```plain
Pronunciation source: Collins Dictionary — pepper; pepper in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/pepper (accent: US; accessed 2026-11-18).
```

### B02-P08 — pasta

```plain
Pronunciation source: Collins Dictionary — pasta; pasta in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/pasta (accent: US; accessed 2026-11-18).
```

### B02-P09 — language

```plain
Pronunciation source: Collins Dictionary — language; language in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/language (accent: US; accessed 2026-11-18).
```

### B02-P10 — asset

```plain
Pronunciation source: Collins Dictionary — asset; asset in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/asset (accent: US; accessed 2026-11-18).
```

### B02-P11 — interest

```plain
Pronunciation source: Collins Dictionary — interest; interest in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/interest (accent: US; accessed 2026-11-18).
```

### B02-P12 — net

```plain
Pronunciation source: Collins Dictionary — net; net in American English 2; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/net (accent: US; accessed 2026-11-18).
```

### B02-P13 — order

```plain
Pronunciation source: Collins Dictionary — order; order in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/order (accent: US; accessed 2026-11-18).
```

### B02-P14 — quarter

```plain
Pronunciation source: Collins Dictionary — quarter; quarter in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/quarter (accent: US; accessed 2026-11-18).
```

### B02-P15 — share

```plain
Pronunciation source: Collins Dictionary — share; share in American English 1; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/share (accent: US; accessed 2026-11-18).
```

### B02-P16 — stock

```plain
Pronunciation source: Collins Dictionary — stock; stock in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/stock (accent: US; accessed 2026-11-18).
```

### B02-P17 — trade

```plain
Pronunciation source: Collins Dictionary — trade; trade in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/trade (accent: US; accessed 2026-11-18).
```

### B02-P18 — transfer

```plain
Pronunciation source: Collins Dictionary — transfer; transfer in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/transfer (accent: US; accessed 2026-11-18).
```

### B02-P19 — review

```plain
Pronunciation source: Collins Dictionary — review; review in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/review (accent: US; accessed 2026-11-18).
```

### B02-P20 — implement

```plain
Pronunciation source: Collins Dictionary — implement; implement in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/implement (accent: US; accessed 2026-11-18).
```

### B02-P21 — branch

```plain
Pronunciation source: Collins Dictionary — branch; branch in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/branch (accent: US; accessed 2026-11-18).
```

### B02-P22 — build

```plain
Pronunciation source: Collins Dictionary — build; build in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/build (accent: US; accessed 2026-11-18).
```

### B02-P23 — cache

```plain
Pronunciation source: Collins Dictionary — cache; cache in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/cache (accent: US; accessed 2026-11-18).
```

### B02-P24 — cash account

```plain
Pronunciation source: Collins Dictionary — cash account; cash account in British English; Collins English Dictionary. https://www.collinsdictionary.com/dictionary/english/cash-account (accent: UK; accessed 2026-11-18).
```

## 暂缓项目

- B02-D01: [part-of-speech-language](../entries/part-of-speech-language.yml)，part of speech，noun phrase，当前释义为 "词性；按语法功能划分的词语类别。"。延续第一批 D01。本次重查 [Collins 整体条目](https://www.collinsdictionary.com/dictionary/english/part-of-speech)，英美分区均未呈现整体 IPA；[Cambridge 美式页](https://dictionary.cambridge.org/us/dictionary/english/part-of-speech)返回 403，Collins pronunciation 页不可访问。接受英式并不能填补证据空缺，不拟写音标或来源行。
- B02-D02: [market-value-finance](../entries/market-value-finance.yml)，market value，noun，当前释义为 "市值；按市场价格衡量的资产价值。"。[Collins 整体条目](https://www.collinsdictionary.com/dictionary/english/market-value)支持含义，但英美分区均未呈现整体 IPA；[Cambridge 美式页](https://dictionary.cambridge.org/us/dictionary/english/market-value)返回 403。继续查找整体读音，不把两词的词典读音拼成短语。

这些结论只描述本轮访问结果，不代表其他可靠词典没有整体读音。未核实的专名、缩写、虚构词和其他复杂项目留待后续批次，不自动填值。未批准任何新增关联或助记。

## 批准后的实施与验证

1. 重新读取批准计划及当前规范、Git 状态和相关实现，核对稳定 ID、当前字段与规划差异。若目标已有新音标或来源，先协调差异，不能覆盖并发编辑。只实施明确批准的编号；暂缓项继续留空。
2. 在仓库外保存所有词条原文、原解析值、当前导出以及现有和归档录音哈希。按目标 ID 定位文件，保留原 Notes；只新增获批 pronunciation 值和逐项来源行，核对编号、词语、口音、URL 一一对应。
3. 不新增词条，不改变 Word、Meaning、词性、其他辅助字段、关联、先修、优先级、标签或音频引用。US/UK 作为原字符串的一部分导出，无需改变 schema、CLI、解析器、导出器、模板或依赖；不修改语言标签或 TTSLanguage。
4. 不生成或替换音频。后续若人工发现词典转写与原录音差异，另行报告，不能以本次音标批准作为重新录音的授权。
5. 运行仓库检查，生成输出只通过现有命令重建。

```bash
go test ./...
go vet ./...
go run scripts/validate/main.go
go run scripts/build/main.go
git diff --check
```

6. 在源不变时再构建一次并比较字节。核对当前基线稳定 ID 集、行数及 17 列导出，逐项检查带 US/UK 标签的完整显示文本和获批 Notes。与基线相比只允许批准行的 Pronunciation、Notes 改变；Audio、StudyOrder、RelatedWords 和其他列均不变。若实施前集合已变，先更新实际基线，不强制历史计数。
7. 检查全部源变化符合批准字段范围，原有及归档录音哈希不变；对未跟踪文本单独检查格式，复查本地链接和锚点。更新本计划实际结果，不手改 dist，不改写第一批历史记录。
8. 不暂存、提交、推送、复制媒体或修改 Anki 集合。区分静态导出验证、人工听审与真实 Anki 导入、显示、播放，不能用静态成功代替设备验证。

## 规划检查

- 当前集合与缺失数来自实际稳定 ID；26 个审阅目标唯一且均存在。24 个提案均缺少 pronunciation，与第一批 22 项没有重叠；2 个暂缓项未列为可写入提案。
- 每个拟写值都有相应实际阅读的词典正文，来源行按编号及 Word 对应。全部 Notes 文案明确口音和访问日，US/UK 没有混拼。没有从合成录音推导音标。
- 本地链接、锚点、表格列数、围栏闭合、行尾空白、制表符和最终换行已检查，git diff --check 通过；新增未跟踪计划另行检查文本格式。
- 与规划前文件哈希比较，仅 AGENTS.md、ENTRY_SCHEMA.md、MAINTENANCE_PROMPTS.md 的规范文字及本计划发生变化。entries、media、archives、实现、模板、依赖、第一批计划及 dist 均保持不变。
- 规划阶段未运行 Go 测试、词库验证、构建、音频生成、解码、人工听审或 Anki 操作；批准后运行的检查结果见实施记录。

已确认: B02-P01–B02-P24 的完整 pronunciation 值及同编号 Notes 来源追加文案。实施状态: 已按用户明确指令完成。B02-D01 和 B02-D02 继续暂缓；本次确认不授权批量改写旧音标或替换录音。


## 实施记录

2026-11-18，按用户明确实施指令完成 B02-P01–B02-P24。24 个词条新增了批准的 pronunciation 值和各自来源行，其中 23 项标注 US，cash-account-finance 标注 UK。6 个原本没有 Notes 的词条创建了该字段；其余词条完整保留原 Notes 并追加来源。B02-D01 的 part-of-speech-language 和 B02-D02 的 market-value-finance 均未修改。

- 逐项比对批准计划中的编号、稳定 ID、Word、词性、Meaning、口音、音标和完整来源行，24 项全部一致；各自来源行只出现一次。transfer 保留批准的动词后重音变体，implement 使用批准的动词读音。
- 全库仍为 1,320 项，英语 1,077 项、日语 243 项。英语已有音标从 313 增至 337，剩余缺失从 764 降至 740。日语读音及第一批全部音标保持不变。
- 源字段差异仅有批准的 pronunciation 新增和 Notes 来源追加。其他 1,296 个词条的源字节完全不变；1,320 个现有 MP3 和 43 个归档 MP3 的 SHA-256 全部保持不变。此前已有的规范修改和其他无关工作保持原状。
- go test ./...、go vet ./...、词库验证、两次构建及 git diff --check 均通过。原有 15 条同形词警告保留，没有新增同形词警告。修改文本和未跟踪计划的格式、本地链接及锚点另行检查通过。
- 两次无源变更的构建逐字节一致，导出 SHA-256 为 `c17bf2292989cb273bced103f8c58526074779a5accd18efa61d9261976cc5e6`。与实施前 TSV 逐行逐列比对，仍为 1,320 行、17 列，只有批准的 24 行的 Pronunciation 和 Notes 改变；Audio、StudyOrder、RelatedWords、标签和其他列保持不变。US/UK 标签按完整普通文本导出，未手改 dist。

本次没有生成、替换或解码录音，没有进行人工听审、真实 Anki 导入、显示或播放验证。cash account 的 UK 音标继续搭配既有录音，不声称两者经人工核实完全一致。未暂存、提交、推送、复制媒体或修改 Anki 集合。
