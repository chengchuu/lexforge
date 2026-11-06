# 英语音标补全计划: 第一批普通词与常用短语

状态: **P01–P22 已批准并实施，静态及导出验证通过。** 日期: 2026-11-06。

规划阶段仅授权创建审阅计划。用户随后确认 P01–P22 并明确指示实施，已按批准值更新 22 个词条并重建导出。其他既有工作、录音、归档、实现、模板和依赖保持不变，未暂存、提交、推送或操作 Anki。下文基线和规格作为规划记录保留，实际结果见[实施记录](#实施记录)。

## 导航

- [范围与基线](#范围与基线)
- [音标选择与来源规则](#音标选择与来源规则)
- [已确认的具体音标](#已确认的具体音标)
- [逐项核对与限制](#逐项核对与限制)
- [Notes 来源追加文案](#notes-来源追加文案)
- [暂缓与后续批次](#暂缓与后续批次)
- [批准后实施与验证](#批准后实施与验证)
- [本次规划检查](#本次规划检查)
- [实施记录](#实施记录)

## 范围与基线

以下为实施前的规划基线。以实际 YAML 中的稳定 ID 为准，已读取当时词库并检查词性、释义、上下文、例句、Notes、现有辅助字段、学习元数据和音频引用。全库缺失统计属于字段盘点，不表示已为全部缺失项查证词典。

| 项目                          | 实际数量                               |
|:------------------------------|:---------------------------------------|
| 全库词条                      | 1,320                                  |
| 英语 / 日语                   | 1,077 / 243                            |
| 英语已有 / 缺失 pronunciation | 291 / 786                              |
| 日语已有 / 缺失 pronunciation | 243 / 0                                |
| 本批具体审阅目标              | 23                                     |
| 可批准的音标新增              | 22 (19 个单词、3 个短语)               |
| 本批暂缓目标                  | 1: part-of-speech-language             |
| 预计修改的既有词条            | 22，仅 pronunciation 和 Notes 追加来源 |
| 新增词条 / 关联 / 先修 / 录音 | 均为 0                                 |

选取标准是日常与界面中的高复用性、当前词义明确、可访问来源中存在美式读音，并包含能检查整体发音的常用短语。account、cash 等虽位于金融词条中，本批只核对普通英文词的对应读音，不改写平台机制、释义或金融解释。这个批次不是全部普通词清单，也不按字母顺序或词长决定后续学习次序。

规划时预计全部实施本批 22 项后，英语已有音标从 291 增至 313，英语仍缺失 764 项；暂缓的 part of speech 包含在这 764 项内，不额外相加。日语不受影响。

## 音标选择与来源规则

- 默认采用美式英语。主要证据为实际打开并阅读的 Collins Dictionary 条目中的 American English 分区；其中 20 项采用署名 Webster’s New World College Dictionary, 5th Digital Edition 的读音，2 项采用该页明确标注 American English 的完整短语转写。下表的来源链接均指向对应整体词条。
- 词典可能列多个正常变体。本批每项选择一个与当前词义、词性和词形匹配的形式供审阅，不表示其他列出的美式变体错误。estimate 必须取名词读音；不因为已有录音而推定读音正确。
- 保留所选来源的记号，例如 /i/、/u/、/ɜr/、/r/、/ɛ/、/ɑ/ 和次重音；这些记法可能与现有条目中的 /iː/、/uː/、/ɝ/、/ɚ/ 或其他词典的习惯不同。本批不全库统一转写体系，也不按英式元音改写美式条目。括号转为项目统一的斜线界定；排版空格的唯一特殊处理见 P22。
- 短语必须有整体来源。本批 black pepper 和 mashed potatoes 采用 Collins 页面直接列出的美式整体音标；sweet and sour 采用同页承认的连字符变体，不自行计算短语重音或弱读。不把朗读软件输出、汉语谐音或记忆中的拼写规则作为词典证据。
- 核对的是词典文字和当前词条意义；本轮没有收听词典音频或现有 MP3，也没有检查自然语流中的所有弱读、连读、变体或口音。
- 访问限制: Cambridge 的部分 dictionary/pronunciation 页面以及 Oxford Learner’s Dictionaries 的 account 页面返回 403；Longman 的目标页不可访问。Cambridge account 的搜索结果可见转写，但没有作为本批最终来源，也不把搜索摘要标记为已阅读全文。最终表内 22 项均有实际可读的 Collins 页面证据。

## 已确认的具体音标

用户已于 2026-11-06 确认 P01–P22 的具体音标。按本节原有批准范围，同时确认各项对应 Notes 来源追加文案。用户随后明确指示实施，已按下表写入。P01–P22 的动作均为 add；实施前 pronunciation 均为省略。下表保留当时批准的完整字符串。词语、词性和释义逐字引用规划时仓库，实施中保持不变。批准某项同时涵盖其后逐项列出的 Notes 来源追加文案；不涉及其他字段。

| 编号 | 稳定 ID / 词语                                                                | 当前词性  | 当前 Meaning                                         | 拟写 pronunciation   | 词典来源                                                                                         |
|:-----|:------------------------------------------------------------------------------|:----------|:-----------------------------------------------------|:---------------------|:-------------------------------------------------------------------------------------------------|
| P01  | [account-finance](../entries/account-finance.yml) / account                   | noun      | 账户；用于记录和管理资金、资产及交易的账户。         | `/əˈkaʊnt/`          | [Collins: account](https://www.collinsdictionary.com/dictionary/english/account)                 |
| P02  | [actual-general](../entries/actual-general.yml) / actual                      | adjective | 实际的；实际发生或观察到的。                         | `/ˈæktʃuəl/`         | [Collins: actual](https://www.collinsdictionary.com/dictionary/english/actual)                   |
| P03  | [word-general](../entries/word-general.yml) / word                            | noun      | 词；单词                                             | `/wɜrd/`             | [Collins: word](https://www.collinsdictionary.com/dictionary/english/word)                       |
| P04  | [context-language](../entries/context-language.yml) / context                 | noun      | 语境；上下文；帮助理解含义的背景。                   | `/ˈkɑnˌtɛkst/`       | [Collins: context](https://www.collinsdictionary.com/dictionary/english/context)                 |
| P05  | [example-general](../entries/example-general.yml) / example                   | noun      | 例子；用来说明规则、含义或用法的实例。               | `/ɪɡˈzæmpəl/`        | [Collins: example](https://www.collinsdictionary.com/dictionary/english/example)                 |
| P06  | [audio-general](../entries/audio-general.yml) / audio                         | noun      | 音频；录制或传输的声音。                             | `/ˈɔdiˌoʊ/`          | [Collins: audio](https://www.collinsdictionary.com/dictionary/english/audio)                     |
| P07  | [profit-finance](../entries/profit-finance.yml) / profit                      | noun      | 利润；收入或交易所得超过相关成本的部分。             | `/ˈprɑfɪt/`          | [Collins: profit](https://www.collinsdictionary.com/dictionary/english/profit)                   |
| P08  | [loss-finance](../entries/loss-finance.yml) / loss                            | noun      | 亏损；成本超过所得或资产价值减少造成的损失。         | `/lɔs/`              | [Collins: loss](https://www.collinsdictionary.com/dictionary/english/loss)                       |
| P09  | [income-finance](../entries/income-finance.yml) / income                      | noun      | 收入；所得，是否扣除费用要结合限定语判断。           | `/ˈɪnˌkʌm/`          | [Collins: income](https://www.collinsdictionary.com/dictionary/english/income)                   |
| P10  | [currency-finance](../entries/currency-finance.yml) / currency                | noun      | 货币；用来标价和结算的货币单位或体系。               | `/ˈkɜrənsi/`         | [Collins: currency](https://www.collinsdictionary.com/dictionary/english/currency)               |
| P11  | [cash-finance](../entries/cash-finance.yml) / cash                            | noun      | 现金；账户中以货币金额记录的资金。                   | `/kæʃ/`              | [Collins: cash](https://www.collinsdictionary.com/dictionary/english/cash)                       |
| P12  | [withdraw-finance](../entries/withdraw-finance.yml) / withdraw                | verb      | 提取；将账户资金取出或转出。                         | `/wɪðˈdrɔ/`          | [Collins: withdraw](https://www.collinsdictionary.com/dictionary/english/withdraw)               |
| P13  | [fear-general](../entries/fear-general.yml) / fear                            | noun      | 恐惧；害怕或担忧的情绪。                             | `/fɪr/`              | [Collins: fear](https://www.collinsdictionary.com/dictionary/english/fear)                       |
| P14  | [greed-general](../entries/greed-general.yml) / greed                         | noun      | 贪婪；对金钱、利益等抱有过度欲求。                   | `/ɡrid/`             | [Collins: greed](https://www.collinsdictionary.com/dictionary/english/greed)                     |
| P15  | [quantity-general](../entries/quantity-general.yml) / quantity                | noun      | 数量；按指定单位计数或度量的多少。                   | `/ˈkwɑntəti/`        | [Collins: quantity](https://www.collinsdictionary.com/dictionary/english/quantity)               |
| P16  | [community-general](../entries/community-general.yml) / community             | noun      | 社区；具有共同兴趣、活动或联系的人群及其交流空间。   | `/kəˈmjunəti/`       | [Collins: community](https://www.collinsdictionary.com/dictionary/english/community)             |
| P17  | [discover-general](../entries/discover-general.yml) / discover                | verb      | 发现；找到或了解到原先未知的内容。                   | `/dɪˈskʌvər/`        | [Collins: discover](https://www.collinsdictionary.com/dictionary/english/discover)               |
| P18  | [estimate-general](../entries/estimate-general.yml) / estimate                | noun      | 估计值；在结果确定前依据可用信息作出的预测或近似数。 | `/ˈɛstəmɪt/`         | [Collins: estimate](https://www.collinsdictionary.com/dictionary/english/estimate)               |
| P19  | [revenue-finance](../entries/revenue-finance.yml) / revenue                   | noun      | 营业收入；公司销售商品或提供服务取得的收入。         | `/ˈrɛvəˌnu/`         | [Collins: revenue](https://www.collinsdictionary.com/dictionary/english/revenue)                 |
| P20  | [black-pepper-food](../entries/black-pepper-food.yml) / black pepper          | noun      | 黑胡椒                                               | `/blæk ˈpɛpər/`      | [Collins: black pepper](https://www.collinsdictionary.com/dictionary/english/black-pepper)       |
| P21  | [mashed-potatoes-food](../entries/mashed-potatoes-food.yml) / mashed potatoes | noun      | 土豆泥                                               | `/ˈmæʃt pəˈteɪtoʊz/` | [Collins: mashed potatoes](https://www.collinsdictionary.com/dictionary/english/mashed-potatoes) |
| P22  | [sweet-and-sour-food](../entries/sweet-and-sour-food.yml) / sweet and sour    | adjective | 糖醋的；酸甜的                                       | `/ˈswit ənˈsaʊr/`    | [Collins: sweet and sour](https://www.collinsdictionary.com/dictionary/english/sweet-and-sour)   |

## 逐项核对与限制

以下说明属于计划审阅依据，除下一节明确列出的来源行外，不授权把这些说明追加到 Notes 或其他记忆辅助字段。

| 编号 | 核对结果与限制                                                                                                                          |
|:-----|:----------------------------------------------------------------------------------------------------------------------------------------|
| P01  | 采用账户义名词；该页 American English 的账户译项也明确给出相同读音。                                                                    |
| P02  | 形容词，选取所列美式变体之一；页面还列有其他自然变体，本批不穷举。                                                                      |
| P03  | 语言单位义名词，保留来源的美式卷舌元音记法。                                                                                            |
| P04  | 语境义名词；保留本来源标出的次重音，不与页面下方其他词典的记法混合。                                                                    |
| P05  | 实例义名词；选择来源所列 ɪɡ- 变体，不将另一 ɛɡ- 变体判为错误。                                                                          |
| P06  | 音频义名词；采用 audio 本词条，而不是 audio- 构词成分。                                                                                 |
| P07  | 利润义名词；采用 American English 正文，不采用同页翻译区的另一转写体系。                                                                |
| P08  | 亏损义名词；来源同时列有 /lɑs/，本批选择 /lɔs/，不据此判定其他美式变体错误。                                                            |
| P09  | 收入义名词；保留来源的次重音。                                                                                                          |
| P10  | 货币义名词；不采用 circulation 或 prevalence 义解释当前内容。                                                                           |
| P11  | 采用 cash 1 的资金义名词；不是古代钱币或人名 Cash。                                                                                     |
| P12  | 提取资金义动词原形；来源同时列有 /wɪθˈdrɔ/，选择其一，不套用 withdrew 或 withdrawn。                                                    |
| P13  | 恐惧义名词；不是 Cape Fear 专名。                                                                                                       |
| P14  | 贪婪义名词；保留来源的 /i/ 记法，不改写为另一词典的长度标注。                                                                           |
| P15  | 数量义单数名词；不写成 quantities 的复数读音。                                                                                          |
| P16  | 共同兴趣人群或社区义名词；保留来源的 /u/ 记法。                                                                                         |
| P17  | 发现义动词原形；不按界面标题的大写形式推定为专名。                                                                                      |
| P18  | 当前词性是 noun，采用来源明确标注的名词读音；不能使用其动词 /ˈɛstəˌmeɪt/。                                                              |
| P19  | 收入义名词；选择来源所列不含 /j/ 的美式变体，不增添英式读音。                                                                           |
| P20  | 整体词条的 American English 译项直接给出完整读音，不是拼接 black 与 pepper。                                                            |
| P21  | 整体词条的 American English 译项；保留 potatoes 的自然复数形式及原文两处重音。                                                          |
| P22  | 来源整体条目为 sweet-and-sour，页面同时承认无连字符写法。只合并网页转写中 ə 与 n 之间的版面空格，保留其音素和重音；不是自行拼接三个词。 |

P20、P21 的证据来自同页 "In other languages" 中明确标注 American English 的读音行，不是把英式转写里的元音替换成美式；它们没有署名为该页 Webster’s New World 分区，来源行也分别标明了真实位置。P22 的完整短语转写可直接读取，网页文本显示的 "ə n" 版面空格合并为 "ən"；没有补充原来源未写出的 /d/。

## Notes 来源追加文案

保留各词条完整现有 Notes，只追加对应的一行。若实施时该来源行已经存在，则不重复追加；若只存在相同 URL 的其他用途来源，先核对是否已经记录发音证据，不重复堆叠同一说明。存在旧 Notes 时，以一个换行接在末尾；没有 Notes 时才创建该可选字段。不删除其他来源或修改现有段落。

以下是每项将追加的完整纯文本，实施时不把 Markdown 围栏写入 YAML 值。

### P01 — account

```plain
Pronunciation source: Collins Dictionary — account; account in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/account (accessed 2026-11-06).
```

### P02 — actual

```plain
Pronunciation source: Collins Dictionary — actual; actual in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/actual (accessed 2026-11-06).
```

### P03 — word

```plain
Pronunciation source: Collins Dictionary — word; word in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/word (accessed 2026-11-06).
```

### P04 — context

```plain
Pronunciation source: Collins Dictionary — context; context in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/context (accessed 2026-11-06).
```

### P05 — example

```plain
Pronunciation source: Collins Dictionary — example; example in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/example (accessed 2026-11-06).
```

### P06 — audio

```plain
Pronunciation source: Collins Dictionary — audio; audio in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/audio (accessed 2026-11-06).
```

### P07 — profit

```plain
Pronunciation source: Collins Dictionary — profit; profit in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/profit (accessed 2026-11-06).
```

### P08 — loss

```plain
Pronunciation source: Collins Dictionary — loss; loss in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/loss (accessed 2026-11-06).
```

### P09 — income

```plain
Pronunciation source: Collins Dictionary — income; income in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/income (accessed 2026-11-06).
```

### P10 — currency

```plain
Pronunciation source: Collins Dictionary — currency; currency in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/currency (accessed 2026-11-06).
```

### P11 — cash

```plain
Pronunciation source: Collins Dictionary — cash; cash in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/cash (accessed 2026-11-06).
```

### P12 — withdraw

```plain
Pronunciation source: Collins Dictionary — withdraw; withdraw in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/withdraw (accessed 2026-11-06).
```

### P13 — fear

```plain
Pronunciation source: Collins Dictionary — fear; fear in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/fear (accessed 2026-11-06).
```

### P14 — greed

```plain
Pronunciation source: Collins Dictionary — greed; greed in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/greed (accessed 2026-11-06).
```

### P15 — quantity

```plain
Pronunciation source: Collins Dictionary — quantity; quantity in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/quantity (accessed 2026-11-06).
```

### P16 — community

```plain
Pronunciation source: Collins Dictionary — community; community in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/community (accessed 2026-11-06).
```

### P17 — discover

```plain
Pronunciation source: Collins Dictionary — discover; discover in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/discover (accessed 2026-11-06).
```

### P18 — estimate

```plain
Pronunciation source: Collins Dictionary — estimate; estimate in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/estimate (accessed 2026-11-06).
```

### P19 — revenue

```plain
Pronunciation source: Collins Dictionary — revenue; revenue in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/revenue (accessed 2026-11-06).
```

### P20 — black pepper

```plain
Pronunciation source: Collins Dictionary — black pepper; In other languages / American English. https://www.collinsdictionary.com/dictionary/english/black-pepper (accessed 2026-11-06).
```

### P21 — mashed potatoes

```plain
Pronunciation source: Collins Dictionary — mashed potatoes; In other languages / American English. https://www.collinsdictionary.com/dictionary/english/mashed-potatoes (accessed 2026-11-06).
```

### P22 — sweet and sour

```plain
Pronunciation source: Collins Dictionary — sweet and sour; sweet-and-sour in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/sweet-and-sour (accessed 2026-11-06).
```

## 暂缓与后续批次

D01: [part-of-speech-language](../entries/part-of-speech-language.yml)，Word 为 part of speech，当前 Meaning 为 "词性；按语法功能划分的词语类别。"，当前词性为 noun phrase。实际读取的 [Collins 整体条目](https://www.collinsdictionary.com/dictionary/english/part-of-speech)支持该含义，但未呈现整体音标。因此本批不提供拟写 IPA，也不追加发音来源；后续继续寻找可靠的美式整体读音来源，不能将只有释义的页面冒充发音证据。

其余 763 个缺失英语条目不在本批具体音标审阅范围内，不表示它们不适合补齐。下一批继续选普通词与常用短语，再单列专业术语、缩写、专名、虚构词和长句。对其他已存在的 291 个英语音标不作批量替换或正确性背书。保留全部 243 个日语假名读音。

## 批准后实施与验证

1. 重新读取本计划、AGENTS.md、[字段规范](ENTRY_SCHEMA.md)、[学习顺序](LEARNING_ORDER.md)和当前源码，检查 Git 状态及规划以来的差异。按稳定 ID 定位，不以文件名干代替 ID；若目标已经补有音标，先核对新内容，不能覆盖并发修改。
2. 在仓库外记录完整词条原文、字段值、所有现有及归档 MP3 哈希和基线 TSV。只写批准的 pronunciation 及本计划列出的 Notes 来源行；保留其余源字段和全部来源文字。不得增加词条、修改词性与词义、改学习顺序、关系或音频引用。
3. 不运行音频生成，不替换现有录音；词典文字与 Samantha 录音的实际吻合程度仍需人工听审。如听审发现差异，另列问题，不能借音标补全自动重新录音。
4. 运行下列仓库检查；构建只通过现有命令生成，不手改 dist。

```bash
go test ./...
go vet ./...
go run scripts/validate/main.go
go run scripts/build/main.go
git diff --check
```

5. 以未变化的源再次构建并比较字节，核对基线 ID 集、1,320 行和 17 列。逐项核对批准的音标、Notes 追加文案及 HTML 转义，确认除 Pronunciation 和获批 Notes 外其他导出列不变，尤其是 Audio、StudyOrder 和 RelatedWords。若实施前基线已变化，则先重新核对实际数量，不能强行维持历史数值。
6. 检查未跟踪文件的格式，核对所有源变更均在批准字段允许列表内；原有及归档 MP3 哈希不变。更新本计划的实际批准、实施及检查状态，区分静态验证、人工听审和真实 Anki 导入、显示、播放。最终报告剩余缺失项，不声称已完成全库音标查证。
7. 不暂存、提交、推送、复制媒体或修改 Anki 集合。用户之后可按[迁移指南](ANKI_MIGRATION_GUIDE.md)更新已有笔记；本批没有模板或字段结构迁移。

## 本次规划检查

- 按当前 1,320 个稳定 ID 盘点，23 个目标均存在，22 个提案及 1 个暂缓项不重复，且都没有 pronunciation；当前词语、词性与释义已与源数据比对。
- 22 个音标分别有对应的已读 Collins 词条，具体分区及变体选择在表中说明；来源访问日期为 2026-11-06。访问失败的页面和搜索摘要均未充作最终证据。不声称词典实时内容永远不变。
- 已检查本文本地文件链接、标题锚点、表格列数、围栏闭合、行尾空格、制表符及最终换行，并运行 git diff --check；未跟踪的本文另行检查。
- 与规划前快照比较，除新增本计划外，原有项目文件字节保持不变，覆盖 entries、media、archives、实现、模板、依赖、既有指南及 dist。规划前已有未跟踪工作得到保留。
- 规划阶段未运行 Go 测试、词库验证、构建、音频生成或 Anki 操作；批准后命令的实际结果见实施记录。

已确认: P01–P22 的 pronunciation 值及各自 Notes 来源追加文案。D01 继续暂缓。实施状态: 已按明确实施指令完成。创建或批准这份计划不授权全库其他词条的音标、内容或录音变更。


## 实施记录

2026-11-06，用户明确指示实施。已按 P01–P22 的具体规格新增 22 个 pronunciation 值，并在各自完整原 Notes 后追加批准的来源行；无 Notes 的 2 个条目创建了该字段。包括 19 个普通单词和 3 个常用短语，没有新增词条、关系、先修或录音。D01 的 part-of-speech-language 保持不变。

- 按批准计划逐项核对编号、稳定 ID、词语、词性、Meaning、音标及完整来源行，22 项均匹配，来源行各出现一次。estimate 使用名词音标，其他已批准的记法及变体选择保持不变。
- 全库仍为 1,320 个稳定 ID (英语 1,077 个、日语 243 个)。英语已有音标 313 项，仍缺失 764 项；日语 243 个读音保持不变。
- 除批准的 pronunciation 和 Notes 追加外，全部原有字段值保持不变。其余 1,298 个词条的源字节完全不变。1,320 个现有 MP3 和 43 个归档 MP3 的 SHA-256 全部保持不变，原有未跟踪工作得到保留。
- go test ./...、go vet ./...、词库验证、两次构建及 git diff --check 均通过；原有 15 条同形词警告仍保留，没有新增此类警告。未跟踪文本的格式与计划本地链接另行检查通过。
- 两次无源变更的构建逐字节一致，导出 SHA-256 为 `f14ddbec217cf4b102803a11d01f156277d272c475ad7c95207cc1d366f6e977`。与实施前 TSV 逐行逐列比对，仍为 1,320 行、17 列，仅批准的 22 行的 Pronunciation 与 Notes 发生变化；Audio、StudyOrder、RelatedWords、标签及其他列均未改变。未手改 dist。

本次没有生成或替换录音，没有运行解码检查、人工听审或真实 Anki 导入、显示和播放验证。音标与现有合成录音是否完全吻合仍需人工听审。未暂存、提交、推送、复制媒体或修改 Anki 集合。
