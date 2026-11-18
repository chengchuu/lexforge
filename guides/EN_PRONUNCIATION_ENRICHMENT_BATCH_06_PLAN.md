# 英语音标补全计划：第六批常用词、技术词与金融短语

状态：**B06-P01–B06-P92 已全部获批并实施，静态与导出验证完成。** 日期：2026-11-18。人工听审及真实 Anki 验证未进行。

规划阶段授权仅整理下一批约 100 项音标供审阅。实际形成 **92 项具体提案（B06-P01–B06-P92）**；没有为凑满 100 项猜读音、拼接短语或套用同形异义词。创建本计划本身不授权实施。用户随后明确确认 B06-P01–B06-P92 并授权实施；本轮仅修改批准的 pronunciation 与 Notes 来源追加，并通过现有命令重建导出。不修改录音、模板、实现或依赖，不暂存、提交、推送或操作 Anki。

## 导航

- [范围与基线](#范围与基线)
- [来源与转写规则](#来源与转写规则)
- [具体音标](#具体音标)
- [逐项核对与来源追加文案](#逐项核对与来源追加文案)
- [本批暂缓项目](#本批暂缓项目)
- [批准后实施与验证](#批准后实施与验证)
- [规划检查](#规划检查)
- [实施与验证记录](#实施与验证记录)

## 范围与基线

按当前 YAML 稳定 ID 盘点 pronunciation；当前词性和释义作为逐项核对基线，不以文件名或历史统计推定身份。原有未提交工作包括前批音标补全和 22 项日文名称助记，均属于本批必须保留的基线。


| 项目 | 实际数量与边界 |
|:---|:---|
| 全库词条 | 1,320；英语 1,077，日语 243 |
| 英语已有 / 缺失 pronunciation | 686 / 391 |
| 日语已有 / 缺失 pronunciation | 243 / 0 |
| 本轮逐项查证目标 | 142 个现有稳定 ID：范围外普通/技术/金融项目 119 个，加《王国之泪》前批暂缓普通词 23 个 |
| 形成具体提案 | 92 项，全部新增 pronunciation：84 项非 TOTK，8 项 TOTK |
| 口音 | 71 项 US（无前缀）；21 项 UK（保留 UK 前缀） |
| 本轮查证后暂缓 | 50 项：35 项非 TOTK，15 项 TOTK；无拟写音标 |
| 其他未逐项查证缺口 | 249 项；不能把字段盘点称为逐项发音查证 |
| 允许实施的字段 | 批准后仅 pronunciation 和对应 Notes 来源追加 |
| 新增词条、录音、关系、先修 | 均为 0 |


如果 92 项全部获批并实施，英语已有音标将从 686 增至 778，缺失从 391 降至 299。299 由本批暂缓 50 项与其他未逐项查证 249 项构成，不重复计数。日语计数不变。

《王国之泪》当前 553 项中已有音标 351 项、缺失 202 项；本批拟解决其中 8 项，届时已有 359 项、仍缺 194 项。本计划不覆盖游戏人物、族群或造词的英文语音核实；此前日文名称助记仅用于联想。

保留[第一批](EN_PRONUNCIATION_ENRICHMENT_BATCH_01_PLAN.md)、[第二批](EN_PRONUNCIATION_ENRICHMENT_BATCH_02_PLAN.md)、[第三批](EN_PRONUNCIATION_ENRICHMENT_BATCH_03_PLAN.md)、[第四批](EN_PRONUNCIATION_ENRICHMENT_BATCH_04_PLAN.md)、[第五批](EN_PRONUNCIATION_ENRICHMENT_BATCH_05_PLAN.md)及 [TOTK 计划](TOTK_PRONUNCIATION_ENRICHMENT_PLAN.md)的历史记录。若本批获批，旧暂缓记录仍作为当时证据边界保留，在本计划记录后续解决情况。

## 来源与转写规则

遵循 [AGENTS.md](../AGENTS.md#pronunciation)与[字段指南](ENTRY_SCHEMA.md#pronunciation)。美式完整 IPA 写 `/…/`，不加 US 前缀；采用英式则写 `UK /…/`。Notes 明确记录所选口音、来源分区、URL 与查阅日期。

- Collins 提供 75 项最终来源，其中 61 项美式、14 项英式。只选实际读取的相应正文分区，不把页面其他词典、派生词或同形词的音标挪过来。Notes 使用 Collins Dictionary 标题，不对未逐项确定的分区版权署名作统一推断。
- Oxford Advanced American Dictionary 提供 dereference 与 leggings 两项明确标为 NAmE 的整体读音。Oxford Advanced Learner’s Dictionary 提供另外 8 项：fiscal year 采用正文明确标为 North American English 的形式；其余 7 项采用该词典首列英式转写，明确标 UK。
- Wiktionary 提供 7 项明确标为 General American 的完整转写，均记录已读页面的固定修订 URL。它是社区词典，不能冒称商业词典、官方术语配音或人工听审；本轮不采纳其词源内容。
- 不声称所有 UK 项都不存在美式读法。它们表示本轮采用可读取、可匹配的英式完整证据。Oxford 通用页及美式专页的访问结果并不一致；已读取的正文随后复查可能返回 403，如 source code 与 fiscal year 的通用页。保留已读证据，不把失败复查写成再次成功。
- Cambridge 本轮尝试的 41 个页面均未取得可读取正文。Collins、Oxford 和 Wiktionary 也有失败或缺少整体 IPA 的页面。搜索摘要、现有 TTS 文件和相邻音频按钮不作为未标明口音 IPA 的替代证据。
- 多变体通常取来源首个完整形式；不把 /r/、/ɹ/、/ɚ/、/i/、/iː/、/ɛ/、/e/ 等改写成另一词典体系。只合并词内版面空格；保留短语词间空格、音节点和可选音。
- 对普通英语词承载的技术义，音标证据仅证明当前英语词形的读音，不声称该词典定义了 React、Git、JavaScript、编译器或金融机制。原有技术释义与来源完整保留。异读词与不同屈折形式另行核对，不由原形推导词尾。
- 不改已有音标，不替换录音，不修改 language、TTSLanguage、学习顺序、关联或既有助记。核实 IPA 不等于核实录音口音、发音质量或真实 Anki 播放。

## 具体音标

每行批准范围为完整 pronunciation 值及下节同编号的完整 Notes 来源行。当前 Word、词性与 Meaning 来自现有条目，不属于重写范围。编号按稳定 ID 排序；分类只便于审阅，不改 tags 或学习顺序。

### 技术、开发与通用词汇

共 50 项。


| 编号 | 稳定 ID / Word | 当前词性 | 当前 Meaning | 拟写 pronunciation | 来源 |
|:---|:---|:---|:---|:---|:---|
| B06-P01 | [array-tech](../entries/array-tech.yml) / array | noun | 数组；按索引组织元素的集合。 | `/əˈreɪ/` | [Collins Dictionary](https://www.collinsdictionary.com/dictionary/english/array) |
| B06-P02 | [backtracking-algorithm](../entries/backtracking-algorithm.yml) / backtracking | noun | 回溯；逐步尝试选择，撤销不合适的选择并继续搜索。 | `/ˈbæk.tɹækɪŋ/` | [Wiktionary](https://en.wiktionary.org/w/index.php?title=backtracking&oldid=92737336) |
| B06-P04 | [bug-tech](../entries/bug-tech.yml) / bug | noun | 程序缺陷；使软件偏离预期行为的问题。 | `/bʌɡ/` | [Collins Dictionary](https://www.collinsdictionary.com/dictionary/english/bug) |
| B06-P08 | [commit-git](../entries/commit-git.yml) / commit | noun; verb | 提交；记录一次项目快照，也可指这次记录。 | `/kəˈmɪt/` | [Collins Dictionary](https://www.collinsdictionary.com/dictionary/english/commit) |
| B06-P09 | [compile-tech](../entries/compile-tech.yml) / compile | verb | 编译；把源代码转换为目标代码或另一种程序表示。 | `/kəmˈpaɪl/` | [Collins Dictionary](https://www.collinsdictionary.com/dictionary/english/compile) |
| B06-P10 | [compiler-tech](../entries/compiler-tech.yml) / compiler | noun | 编译器；执行程序编译的软件。 | `/kəmˈpaɪlər/` | [Collins Dictionary](https://www.collinsdictionary.com/dictionary/english/compiler) |
| B06-P11 | [copy-tech](../entries/copy-tech.yml) / copy | noun | 复制；创建数据副本；也可指副本。 | `/ˈkɑpi/` | [Collins Dictionary](https://www.collinsdictionary.com/dictionary/english/copy) |
| B06-P13 | [counter-tech](../entries/counter-tech.yml) / counter | noun | 计数器；记录某事件发生次数或累计数量的状态。 | `/ˈkaʊntər/` | [Collins Dictionary](https://www.collinsdictionary.com/dictionary/english/counter) |
| B06-P14 | [crash-tech](../entries/crash-tech.yml) / crash | noun; verb | 崩溃；程序因故障意外停止或无法继续正常执行。 | `/kræʃ/` | [Collins Dictionary](https://www.collinsdictionary.com/dictionary/english/crash) |
| B06-P17 | [deduplicate-tech](../entries/deduplicate-tech.yml) / deduplicate | verb | 去重；识别重复项并按规则保留一份。 | `UK /diːˈdjuːplɪˌkeɪt/` | [Collins Dictionary](https://www.collinsdictionary.com/dictionary/english/deduplicate) |
| B06-P20 | [dereference-tech](../entries/dereference-tech.yml) / dereference | verb | 解引用；通过指针或引用访问所指向的值。 | `/ˌdiˈrɛfrəns/` | [Oxford Advanced American Dictionary](https://www.oxfordlearnersdictionaries.com/definition/american_english/dereference) |
| B06-P23 | [draft-workflow](../entries/draft-workflow.yml) / draft | noun | 草稿；尚待修改或审阅的初稿。 | `/dræft/` | [Collins Dictionary](https://www.collinsdictionary.com/dictionary/english/draft) |
| B06-P25 | [effect-react](../entries/effect-react.yml) / effect | noun | React Effect；在组件提交后将组件与外部系统同步的逻辑。 | `/ɛˈfɛkt/` | [Collins Dictionary](https://www.collinsdictionary.com/dictionary/english/effect) |
| B06-P26 | [enum-typescript](../entries/enum-typescript.yml) / enum | noun | 枚举；为一组具名常量定义成员的结构。 | `/ˈinʌm/` | [Wiktionary](https://en.wiktionary.org/w/index.php?title=enum&oldid=94581927) |
| B06-P27 | [enumerable-javascript](../entries/enumerable-javascript.yml) / enumerable | adjective | 可枚举的；属性可被相应枚举操作列举的特性。 | `/ɪˈnuːmərəbəl/` | [Collins Dictionary](https://www.collinsdictionary.com/dictionary/english/enumerable) |
| B06-P28 | [etymology-language](../entries/etymology-language.yml) / etymology | noun | 词源；词语的历史来源与演变。 | `/ˌɛtəˈmɑlədʒi/` | [Collins Dictionary](https://www.collinsdictionary.com/dictionary/english/etymology) |
| B06-P29 | [event-tech](../entries/event-tech.yml) / event | noun | 事件；程序可观察并响应的发生或通知。 | `/iˈvɛnt/` | [Collins Dictionary](https://www.collinsdictionary.com/dictionary/english/event) |
| B06-P40 | [layout-browser-rendering](../entries/layout-browser-rendering.yml) / layout | noun | 布局；根据样式与约束计算页面元素的位置和尺寸。 | `/ˈleɪˌaʊt/` | [Collins Dictionary](https://www.collinsdictionary.com/dictionary/english/layout) |
| B06-P43 | [lowering-compiler](../entries/lowering-compiler.yml) / lowering | noun | 降级转换；把较高层程序表示转换为更低层表示的编译过程。 | `/ˈloʊ.ɚ.ɪŋ/` | [Wiktionary](https://en.wiktionary.org/w/index.php?title=lowering&oldid=88196875) |
| B06-P47 | [memoization-programming](../entries/memoization-programming.yml) / memoization | noun | 记忆化；缓存计算结果，在输入满足复用条件时避免重复计算。 | `/ˌmɛmoʊaɪˈzeɪʃən/` | [Wiktionary](https://en.wiktionary.org/w/index.php?title=memoization&oldid=92747904) |
| B06-P48 | [merge-git](../entries/merge-git.yml) / merge | verb | 合并；把不同开发历史的改动整合到一起。 | `/mɜrdʒ/` | [Collins Dictionary](https://www.collinsdictionary.com/dictionary/english/merge) |
| B06-P49 | [method-tech](../entries/method-tech.yml) / method | noun | 方法；与对象、类或类型关联、用于执行行为的函数成员。 | `/ˈmɛθəd/` | [Collins Dictionary](https://www.collinsdictionary.com/dictionary/english/method) |
| B06-P50 | [middleware-web](../entries/middleware-web.yml) / middleware | noun | 中间件；在请求处理链中组合可复用处理逻辑的函数。 | `UK /ˈmɪdəlˌwɛə/` | [Collins Dictionary](https://www.collinsdictionary.com/dictionary/english/middleware) |
| B06-P51 | [mnemonic-memory](../entries/mnemonic-memory.yml) / mnemonic | noun | 助记法；帮助记忆的提示或方法。 | `/niˈmɑnɪk/` | [Collins Dictionary](https://www.collinsdictionary.com/dictionary/english/mnemonic) |
| B06-P52 | [operand-tech](../entries/operand-tech.yml) / operand | noun | 操作数；运算符所作用的表达式或值。 | `/ˈɑpərˌænd/` | [Collins Dictionary](https://www.collinsdictionary.com/dictionary/english/operand) |
| B06-P54 | [paint-browser-rendering](../entries/paint-browser-rendering.yml) / paint | noun | 绘制；把页面元素的视觉内容记录或绘制出来的过程。 | `/peɪnt/` | [Collins Dictionary](https://www.collinsdictionary.com/dictionary/english/paint) |
| B06-P55 | [parameter-tech](../entries/parameter-tech.yml) / parameter | noun | 参数；函数或泛型定义中用于接收实参的占位名称。 | `/pəˈræmətər/` | [Collins Dictionary](https://www.collinsdictionary.com/dictionary/english/parameter) |
| B06-P57 | [patch-tech](../entries/patch-tech.yml) / patch | noun | 补丁；用于修正或调整程序的一组改动。 | `/pætʃ/` | [Collins Dictionary](https://www.collinsdictionary.com/dictionary/english/patch) |
| B06-P58 | [persist-tech](../entries/persist-tech.yml) / persist | verb | 持久保存；使数据在当前操作或进程结束后仍可恢复使用。 | `/pərˈsɪst/` | [Collins Dictionary](https://www.collinsdictionary.com/dictionary/english/persist) |
| B06-P61 | [primitive-javascript](../entries/primitive-javascript.yml) / primitive | noun; adjective | 原始值；不属于对象、值本身不可变的 JavaScript 数据。 | `/ˈprɪmɪtɪv/` | [Collins Dictionary](https://www.collinsdictionary.com/dictionary/english/primitive) |
| B06-P62 | [property-tech](../entries/property-tech.yml) / property | noun | 属性；以键标识的对象成员。 | `/ˈprɑpərti/` | [Collins Dictionary](https://www.collinsdictionary.com/dictionary/english/property) |
| B06-P64 | [quadratic-tech](../entries/quadratic-tech.yml) / quadratic | adjective | 二次增长的；输入规模增加时，工作量按平方量级增长。 | `/kwɑˈdrætɪk/` | [Collins Dictionary](https://www.collinsdictionary.com/dictionary/english/quadratic) |
| B06-P65 | [query-tech](../entries/query-tech.yml) / query | noun | 查询；向系统请求信息或检索结果。 | `/ˈkwɪri/` | [Collins Dictionary](https://www.collinsdictionary.com/dictionary/english/query) |
| B06-P66 | [quicksort-algorithm](../entries/quicksort-algorithm.yml) / quicksort | noun | 快速排序；按基准值分区，再递归排序各分区。 | `/ˈkwɪkˌsɔɹt/` | [Wiktionary](https://en.wiktionary.org/w/index.php?title=quicksort&oldid=95458003) |
| B06-P67 | [reconciliation-react](../entries/reconciliation-react.yml) / reconciliation | noun | React 协调；比较界面描述并确定更新与状态保留。 | `/ˌrɛkənˌsɪliˈeɪʃən/` | [Collins Dictionary](https://www.collinsdictionary.com/dictionary/english/reconciliation) |
| B06-P68 | [recursion-programming](../entries/recursion-programming.yml) / recursion | noun | 递归；在求解或定义过程中再次处理同类子问题的方式。 | `/rɪˈkɜrʒən/` | [Collins Dictionary](https://www.collinsdictionary.com/dictionary/english/recursion) |
| B06-P69 | [repository-git](../entries/repository-git.yml) / repository | noun | 仓库；保存项目文件及其版本历史的集合。 | `/rɪˈpɑzəˌtɔri/` | [Collins Dictionary](https://www.collinsdictionary.com/dictionary/english/repository) |
| B06-P73 | [shorthand-javascript](../entries/shorthand-javascript.yml) / shorthand | noun | 简写；用更短的语法表达相应的常规写法。 | `/ˈʃɔrtˌhænd/` | [Collins Dictionary](https://www.collinsdictionary.com/dictionary/english/shorthand) |
| B06-P74 | [source-code-tech](../entries/source-code-tech.yml) / source code | noun phrase | 源代码；以编程语言编写的程序文本。 | `UK /ˈsɔːs kəʊd/` | [Oxford Advanced Learner’s Dictionary](https://www.oxfordlearnersdictionaries.com/definition/english/source-code) |
| B06-P75 | [stack-data-structure](../entries/stack-data-structure.yml) / stack | noun | 栈；按后进先出顺序添加和移除元素的数据结构。 | `/stæk/` | [Collins Dictionary](https://www.collinsdictionary.com/dictionary/english/stack) |
| B06-P77 | [syntax-tech](../entries/syntax-tech.yml) / syntax | noun | 语法；规定语言符号如何组成合法结构的规则。 | `/ˈsɪnˌtæks/` | [Collins Dictionary](https://www.collinsdictionary.com/dictionary/english/syntax) |
| B06-P79 | [test-tech](../entries/test-tech.yml) / test | noun | 测试；用规定的检查评估软件行为或属性。 | `/tɛst/` | [Collins Dictionary](https://www.collinsdictionary.com/dictionary/english/test) |
| B06-P80 | [text-to-speech-tech](../entries/text-to-speech-tech.yml) / text-to-speech | noun phrase | 文本转语音；把书面文字合成为可听语音的技术。 | `UK /ˌtekst tə ˈspiːtʃ/` | [Oxford Advanced Learner’s Dictionary](https://www.oxfordlearnersdictionaries.com/definition/english/text-to-speech) |
| B06-P81 | [token-llm](../entries/token-llm.yml) / token | noun | 词元；语言模型将文本编码后使用的基本处理单位。 | `/ˈtoʊkən/` | [Collins Dictionary](https://www.collinsdictionary.com/dictionary/english/token) |
| B06-P83 | [tree-data-structure](../entries/tree-data-structure.yml) / tree | noun | 树；用节点的父子关系组织层级数据的结构。 | `/tri/` | [Collins Dictionary](https://www.collinsdictionary.com/dictionary/english/tree) |
| B06-P84 | [triage-workflow](../entries/triage-workflow.yml) / triage | noun; verb | 分流评估；先按影响、紧急程度和可处理性筛选待办问题。 | `/triˈɑʒ/` | [Collins Dictionary](https://www.collinsdictionary.com/dictionary/english/triage) |
| B06-P85 | [type-tech](../entries/type-tech.yml) / type | noun | 类型；描述值的类别及允许操作或约束的信息。 | `/taɪp/` | [Collins Dictionary](https://www.collinsdictionary.com/dictionary/english/type) |
| B06-P86 | [variance-type-system](../entries/variance-type-system.yml) / variance | noun | 型变；类型参数之间的子类型关系如何影响复合类型之间的关系。 | `/ˈvɛriəns/` | [Collins Dictionary](https://www.collinsdictionary.com/dictionary/english/variance) |
| B06-P90 | [wealth-general](../entries/wealth-general.yml) / wealth | noun | 财富；拥有的金钱、资产和其他有价值资源。 | `/wɛlθ/` | [Collins Dictionary](https://www.collinsdictionary.com/dictionary/english/wealth) |
| B06-P92 | [workaround-tech](../entries/workaround-tech.yml) / workaround | noun | 变通办法；暂时避开问题而不直接修复根因的做法。 | `/ˈwɜrkəˌraʊnd/` | [Collins Dictionary](https://www.collinsdictionary.com/dictionary/english/workaround) |

### 金融与账户词汇

共 30 项。


| 编号 | 稳定 ID / Word | 当前词性 | 当前 Meaning | 拟写 pronunciation | 来源 |
|:---|:---|:---|:---|:---|:---|
| B06-P05 | [buying-power-finance](../entries/buying-power-finance.yml) / buying power | noun | 购买力；按账户和交易规则可用于新增交易的额度。 | `UK /ˈbaɪɪŋ ˈpaʊə/` | [Collins Dictionary](https://www.collinsdictionary.com/dictionary/english/buying-power) |
| B06-P06 | [call-option-finance](../entries/call-option-finance.yml) / call option | noun | 看涨期权；在股票期权语境中给予买方按合约价格买入标的股票的权利。 | `UK /kɔːl ˈɒpʃən/` | [Collins Dictionary](https://www.collinsdictionary.com/dictionary/english/call-option) |
| B06-P07 | [closing-price-finance](../entries/closing-price-finance.yml) / closing price | noun | 收盘价；某一交易时段结束时采用的价格。 | `UK /ˈkləʊzɪŋ praɪs/` | [Collins Dictionary](https://www.collinsdictionary.com/dictionary/english/closing-price) |
| B06-P12 | [cost-finance](../entries/cost-finance.yml) / cost | noun | 成本；取得或持有某项资产所付出的代价。 | `/kɔst/` | [Collins Dictionary](https://www.collinsdictionary.com/dictionary/english/cost) |
| B06-P15 | [crypto-finance](../entries/crypto-finance.yml) / crypto | noun | 加密货币；对 cryptocurrency 或加密资产领域的常用简称。 | `/ˈkɹɪptoʊ/` | [Wiktionary](https://en.wiktionary.org/w/index.php?title=crypto&oldid=94551653) |
| B06-P18 | [deposit-finance](../entries/deposit-finance.yml) / deposit | verb | 存入资金；将钱转入账户。 | `/diˈpɑzɪt/` | [Collins Dictionary](https://www.collinsdictionary.com/dictionary/english/deposit) |
| B06-P21 | [discount-finance](../entries/discount-finance.yml) / discount | noun | 折价；价格低于某个比较基准的情况。 | `/ˈdɪsˌkaʊnt/` | [Collins Dictionary](https://www.collinsdictionary.com/dictionary/english/discount) |
| B06-P22 | [dividend-finance](../entries/dividend-finance.yml) / dividend | noun | 股息；公司向股东分配的收益。 | `/ˈdɪvəˌdɛnd/` | [Collins Dictionary](https://www.collinsdictionary.com/dictionary/english/dividend) |
| B06-P24 | [earnings-finance](../entries/earnings-finance.yml) / earnings | noun | 盈利；公司在一定期间取得的利润。 | `/ˈɜrnɪŋz/` | [Collins Dictionary](https://www.collinsdictionary.com/dictionary/english/earnings) |
| B06-P30 | [exchange-currency](../entries/exchange-currency.yml) / exchange | verb | 兑换；将一种货币换成另一种。 | `/ɛksˈtʃeɪndʒ/` | [Collins Dictionary](https://www.collinsdictionary.com/dictionary/english/exchange) |
| B06-P31 | [financials-finance](../entries/financials-finance.yml) / financials | noun | 财务数据；财务报表等公司财务信息。 | `/faɪˈnænʃəlz/` | [Collins Dictionary](https://www.collinsdictionary.com/dictionary/english/financials) |
| B06-P32 | [fiscal-year-finance](../entries/fiscal-year-finance.yml) / fiscal year | noun | 财年；公司或机构用于财务报告的年度。 | `/ˌfɪskl ˈjɪr/` | [Oxford Advanced Learner’s Dictionary](https://www.oxfordlearnersdictionaries.com/definition/english/fiscal-year) |
| B06-P33 | [futures-finance](../entries/futures-finance.yml) / futures | noun | 期货；按标准化合约约定未来交易条件的交易所交易工具。 | `UK /ˈfjuːtʃəz/` | [Collins Dictionary](https://www.collinsdictionary.com/dictionary/english/futures) |
| B06-P34 | [gainer-finance](../entries/gainer-finance.yml) / gainer | noun | 上涨标的；在指定期间内价格上升的证券或资产。 | `/ˈɡeɪnər/` | [Collins Dictionary](https://www.collinsdictionary.com/dictionary/english/gainer) |
| B06-P37 | [heat-map-data](../entries/heat-map-data.yml) / heat map | noun | 热力图；用颜色或色块表达数值差异的数据图。 | `UK /ˈhiːt mæp/` | [Oxford Advanced Learner’s Dictionary](https://www.oxfordlearnersdictionaries.com/definition/english/heat-map) |
| B06-P39 | [index-finance](../entries/index-finance.yml) / index | noun | 指数；用于概括一组证券或某种市场指标变化的数值。 | `/ˈɪnˌdɛks/` | [Collins Dictionary](https://www.collinsdictionary.com/dictionary/english/index) |
| B06-P42 | [loser-finance](../entries/loser-finance.yml) / loser | noun | 下跌标的；在指定期间内价格下降的证券或资产。 | `/ˈluzər/` | [Collins Dictionary](https://www.collinsdictionary.com/dictionary/english/loser) |
| B06-P44 | [margin-finance](../entries/margin-finance.yml) / margin | noun | 保证金；交易或融资中用于提供担保的资金或资产权益。 | `/ˈmɑrdʒən/` | [Collins Dictionary](https://www.collinsdictionary.com/dictionary/english/margin) |
| B06-P45 | [market-value-finance](../entries/market-value-finance.yml) / market value | noun | 市值；按市场价格衡量的资产价值。 | `UK /ˌmɑːkɪt ˈvæljuː/` | [Oxford Advanced Learner’s Dictionary](https://www.oxfordlearnersdictionaries.com/definition/english/market-value) |
| B06-P53 | [option-finance](../entries/option-finance.yml) / option | noun | 期权；给予买方按合约条件买入或卖出标的权利的合约。 | `/ˈɑpʃən/` | [Collins Dictionary](https://www.collinsdictionary.com/dictionary/english/option) |
| B06-P59 | [portfolio-finance](../entries/portfolio-finance.yml) / portfolio | noun | 投资组合；同一投资者或账户持有的一组投资资产。 | `/pɔrtˈfoʊliˌoʊ/` | [Collins Dictionary](https://www.collinsdictionary.com/dictionary/english/portfolio) |
| B06-P60 | [position-finance](../entries/position-finance.yml) / position | noun | 持仓；账户对某项金融工具持有的数量或形成的风险敞口。 | `/pəˈzɪʃən/` | [Collins Dictionary](https://www.collinsdictionary.com/dictionary/english/position) |
| B06-P63 | [put-option-finance](../entries/put-option-finance.yml) / put option | noun | 看跌期权；在股票期权语境中给予买方按合约价格卖出标的股票的权利。 | `UK /pʊt ˈɒpʃən/` | [Collins Dictionary](https://www.collinsdictionary.com/dictionary/english/put-option) |
| B06-P70 | [return-finance](../entries/return-finance.yml) / return | noun | 投资回报；投入资金产生的收益结果，可用金额或比例表示。 | `/rɪˈtɜrn/` | [Collins Dictionary](https://www.collinsdictionary.com/dictionary/english/return) |
| B06-P72 | [shareholder-finance](../entries/shareholder-finance.yml) / shareholder | noun | 股东；持有公司股份的人或机构。 | `/ˈʃɛrˌhoʊldər/` | [Collins Dictionary](https://www.collinsdictionary.com/dictionary/english/shareholder) |
| B06-P76 | [symbol-market](../entries/symbol-market.yml) / symbol | noun | 证券或交易标的代码；用于标识某个交易对象的字符标记。 | `/ˈsɪmbəl/` | [Collins Dictionary](https://www.collinsdictionary.com/dictionary/english/symbol) |
| B06-P82 | [transaction-finance](../entries/transaction-finance.yml) / transaction | noun | 交易；涉及资金或资产往来的业务行为或记录。 | `/trænˈzækʃən/` | [Collins Dictionary](https://www.collinsdictionary.com/dictionary/english/transaction) |
| B06-P87 | [volatility-finance](../entries/volatility-finance.yml) / volatility | noun | 波动性；价格或收益在一段时间内变化的程度。 | `/ˌvɑləˈtɪləti/` | [Collins Dictionary](https://www.collinsdictionary.com/dictionary/english/volatility) |
| B06-P88 | [volume-trading](../entries/volume-trading.yml) / volume | noun | 成交量；某一时段内完成交易的规模，可按数量或计价金额表示。 | `/ˈvɑljum/` | [Collins Dictionary](https://www.collinsdictionary.com/dictionary/english/volume) |
| B06-P89 | [watchlist-finance](../entries/watchlist-finance.yml) / watchlist | noun | 自选列表；用于持续关注一组证券或其他标的的清单。 | `/ˈwɔt͡ʃˌlɪst/` | [Wiktionary](https://en.wiktionary.org/w/index.php?title=watchlist&oldid=91656417) |

### 语言、学习与饮食词汇

共 4 项。


| 编号 | 稳定 ID / Word | 当前词性 | 当前 Meaning | 拟写 pronunciation | 来源 |
|:---|:---|:---|:---|:---|:---|
| B06-P03 | [braised-food](../entries/braised-food.yml) / braised | adjective | 炖的；焖制的 | `UK /breɪzd/` | [Collins Dictionary](https://www.collinsdictionary.com/dictionary/english/braised) |
| B06-P46 | [mashed-food](../entries/mashed-food.yml) / mashed | adjective | 捣碎的；做成泥的 | `UK /mæʃt/` | [Oxford Advanced Learner’s Dictionary](https://www.oxfordlearnersdictionaries.com/definition/english/mashed) |
| B06-P56 | [part-of-speech-language](../entries/part-of-speech-language.yml) / part of speech | noun phrase | 词性；按语法功能划分的词语类别。 | `UK /ˌpɑːt əv ˈspiːtʃ/` | [Oxford Advanced Learner’s Dictionary](https://www.oxfordlearnersdictionaries.com/definition/english/part-of-speech) |
| B06-P78 | [teriyaki-food](../entries/teriyaki-food.yml) / teriyaki | adjective | 照烧 | `/ˌtɛriˈjɑki/` | [Collins Dictionary](https://www.collinsdictionary.com/dictionary/english/teriyaki) |

### 前批 TOTK 普通词与短语

共 8 项。


| 编号 | 稳定 ID / Word | 当前词性 | 当前 Meaning | 拟写 pronunciation | 来源 |
|:---|:---|:---|:---|:---|:---|
| B06-P16 | [decayed-general](../entries/decayed-general.yml) / decayed | adjective | 腐朽的；腐败或衰损的。 | `UK /dɪˈkeɪd/` | [Collins Dictionary](https://www.collinsdictionary.com/dictionary/english/decayed) |
| B06-P19 | [depths-terrain](../entries/depths-terrain.yml) / depths | noun | 深处；深层区域；常用复数形式。 | `UK /dɛpθs/` | [Collins Dictionary](https://www.collinsdictionary.com/dictionary/english/depths) |
| B06-P35 | [general-store-architecture](../entries/general-store-architecture.yml) / general store | phrase | 杂货店 | `UK /ˌdʒenrəl ˈstɔː(r)/` | [Oxford Advanced Learner’s Dictionary](https://www.oxfordlearnersdictionaries.com/definition/english/general-store) |
| B06-P36 | [greens-food](../entries/greens-food.yml) / greens | noun | 绿叶蔬菜。 | `UK /ɡriːnz/` | [Collins Dictionary](https://www.collinsdictionary.com/dictionary/english/greens) |
| B06-P38 | [highlands-terrain](../entries/highlands-terrain.yml) / highlands | noun | 高地 | `UK /ˈhaɪləndz/` | [Collins Dictionary](https://www.collinsdictionary.com/dictionary/english/highlands) |
| B06-P41 | [leggings-clothing](../entries/leggings-clothing.yml) / leggings | noun | 紧身裤、护腿 | `/ˈlɛɡɪŋz/` | [Oxford Advanced American Dictionary](https://www.oxfordlearnersdictionaries.com/definition/american_english/leggings) |
| B06-P71 | [ruins-architecture](../entries/ruins-architecture.yml) / ruins | noun | 废墟、遗迹 | `UK /ˈruːɪnz/` | [Collins Dictionary](https://www.collinsdictionary.com/dictionary/english/ruins) |
| B06-P91 | [wetlands-terrain](../entries/wetlands-terrain.yml) / wetlands | noun | 湿地 | `UK /ˈwɛtlændz/` | [Collins Dictionary](https://www.collinsdictionary.com/dictionary/english/wetlands) |


## 逐项核对与来源追加文案

以下代码块各为一行待追加到 Notes 的完整文本，围栏和标题不写入 YAML。保留原 Notes 的全部文字和段落；只追加一次，不覆盖或迁移既有来源。核对说明属于审阅记录，不自动添加为词条助记。


### B06-P01 — array

取所列美式分区中的一个完整读音；不混拼其他变体。 普通英语词形的发音证据不代替现有技术释义、机制或产品官方发音证据。

```plain
Pronunciation source: Collins Dictionary — array; array in American English. https://www.collinsdictionary.com/dictionary/english/array (accent: US; accessed 2026-11-18).
```


### B06-P02 — backtracking

社区词典中明确标为 General American 的完整 IPA；保留来源转写，使用已读页面的固定修订链接。 普通英语词形的发音证据不代替现有技术释义、机制或产品官方发音证据。

```plain
Pronunciation source: Wiktionary — backtracking; English — Pronunciation — General American. https://en.wiktionary.org/w/index.php?title=backtracking&oldid=92737336 (accent: US; accessed 2026-11-18).
```


### B06-P03 — braised

采用已读英式分区的完整音标；不自行改写为美式。

```plain
Pronunciation source: Collins Dictionary — braised; braised in British English. https://www.collinsdictionary.com/dictionary/english/braised (accent: UK; accessed 2026-11-18).
```


### B06-P04 — bug

取所列美式分区中的一个完整读音；不混拼其他变体。 普通英语词形的发音证据不代替现有技术释义、机制或产品官方发音证据。

```plain
Pronunciation source: Collins Dictionary — bug; bug in American English 1. https://www.collinsdictionary.com/dictionary/english/bug (accent: US; accessed 2026-11-18).
```


### B06-P05 — buying power

采用已读英式分区的完整音标；不自行改写为美式。

```plain
Pronunciation source: Collins Dictionary — buying power; buying power in British English. https://www.collinsdictionary.com/dictionary/english/buying-power (accent: UK; accessed 2026-11-18).
```


### B06-P06 — call option

采用所选英式整体 /kɔːl ˈɒpʃən/；仅合并版面空格，不由两个词音标拼合。

```plain
Pronunciation source: Collins Dictionary — call option; call option in British English. https://www.collinsdictionary.com/dictionary/english/call-option (accent: UK; accessed 2026-11-18).
```


### B06-P07 — closing price

采用已读英式分区的完整音标；不自行改写为美式。

```plain
Pronunciation source: Collins Dictionary — closing price; closing price in British English. https://www.collinsdictionary.com/dictionary/english/closing-price (accent: UK; accessed 2026-11-18).
```


### B06-P08 — commit

源词条包含名词与动词。本项补普通英语 commit 的 /kəˈmɪt/；不把“托付”等普通词义当作 Git 定义来源，保留现有 Git 释义。 普通英语词形的发音证据不代替现有技术释义、机制或产品官方发音证据。

```plain
Pronunciation source: Collins Dictionary — commit; commit in American English. https://www.collinsdictionary.com/dictionary/english/commit (accent: US; accessed 2026-11-18).
```


### B06-P09 — compile

取所列美式分区中的一个完整读音；不混拼其他变体。 普通英语词形的发音证据不代替现有技术释义、机制或产品官方发音证据。

```plain
Pronunciation source: Collins Dictionary — compile; compile in American English. https://www.collinsdictionary.com/dictionary/english/compile (accent: US; accessed 2026-11-18).
```


### B06-P10 — compiler

取所列美式分区中的一个完整读音；不混拼其他变体。 普通英语词形的发音证据不代替现有技术释义、机制或产品官方发音证据。

```plain
Pronunciation source: Collins Dictionary — compiler; compiler in American English. https://www.collinsdictionary.com/dictionary/english/compiler (accent: US; accessed 2026-11-18).
```


### B06-P11 — copy

取所列美式分区中的一个完整读音；不混拼其他变体。 普通英语词形的发音证据不代替现有技术释义、机制或产品官方发音证据。

```plain
Pronunciation source: Collins Dictionary — copy; copy in American English. https://www.collinsdictionary.com/dictionary/english/copy (accent: US; accessed 2026-11-18).
```


### B06-P12 — cost

取所列美式分区中的一个完整读音；不混拼其他变体。

```plain
Pronunciation source: Collins Dictionary — cost; cost in American English. https://www.collinsdictionary.com/dictionary/english/cost (accent: US; accessed 2026-11-18).
```


### B06-P13 — counter

取 counter in American English 1 的计数者/计数装置名词分区，不取相反方向的其他同形分区。 普通英语词形的发音证据不代替现有技术释义、机制或产品官方发音证据。

```plain
Pronunciation source: Collins Dictionary — counter; counter in American English 1. https://www.collinsdictionary.com/dictionary/english/counter (accent: US; accessed 2026-11-18).
```


### B06-P14 — crash

取所列美式分区中的一个完整读音；不混拼其他变体。 普通英语词形的发音证据不代替现有技术释义、机制或产品官方发音证据。

```plain
Pronunciation source: Collins Dictionary — crash; crash in American English 1. https://www.collinsdictionary.com/dictionary/english/crash (accent: US; accessed 2026-11-18).
```


### B06-P15 — crypto

所选英语分区同时明确列出 cryptocurrency 的简称义与 General American IPA；不引用 Collins 只释为秘密支持者的旧义作为本项意义证据。

```plain
Pronunciation source: Wiktionary — crypto; English — Pronunciation — General American. https://en.wiktionary.org/w/index.php?title=crypto&oldid=94551653 (accent: US; accessed 2026-11-18).
```


### B06-P16 — decayed

采用已读英式分区的完整音标；不自行改写为美式。

```plain
Pronunciation source: Collins Dictionary — decayed; decayed in British English. https://www.collinsdictionary.com/dictionary/english/decayed (accent: UK; accessed 2026-11-18).
```


### B06-P17 — deduplicate

采用已读英式分区的完整音标；不自行改写为美式。 普通英语词形的发音证据不代替现有技术释义、机制或产品官方发音证据。

```plain
Pronunciation source: Collins Dictionary — deduplicate; deduplicate in British English. https://www.collinsdictionary.com/dictionary/english/deduplicate (accent: UK; accessed 2026-11-18).
```


### B06-P18 — deposit

取所列美式分区中的一个完整读音；不混拼其他变体。

```plain
Pronunciation source: Collins Dictionary — deposit; deposit in American English. https://www.collinsdictionary.com/dictionary/english/deposit (accent: US; accessed 2026-11-18).
```


### B06-P19 — depths

采用来源所列复数 depths 整体音标；不由 depth 添加词尾。当前普通“深处”用法与游戏地点专名分开，不声称已核实游戏配音。

```plain
Pronunciation source: Collins Dictionary — depths; depths in British English. https://www.collinsdictionary.com/dictionary/english/depths (accent: UK; accessed 2026-11-18).
```


### B06-P20 — dereference

采用 Oxford Advanced American Dictionary 明列的 NAmE /ˌdiˈrɛfrəns/，保留该词典短 i、ɛ、ɪ 的符号习惯。 普通英语词形的发音证据不代替现有技术释义、机制或产品官方发音证据。

```plain
Pronunciation source: Oxford Advanced American Dictionary — dereference; dereference — NAmE. https://www.oxfordlearnersdictionaries.com/definition/american_english/dereference (accent: US; accessed 2026-11-18).
```


### B06-P21 — discount

当前为折价名词，取来源首列名词 /ˈdɪsˌkaʊnt/，不取后列动词变体 /dɪsˈkaʊnt/。

```plain
Pronunciation source: Collins Dictionary — discount; discount in American English. https://www.collinsdictionary.com/dictionary/english/discount (accent: US; accessed 2026-11-18).
```


### B06-P22 — dividend

取所列美式分区中的一个完整读音；不混拼其他变体。

```plain
Pronunciation source: Collins Dictionary — dividend; dividend in American English. https://www.collinsdictionary.com/dictionary/english/dividend (accent: US; accessed 2026-11-18).
```


### B06-P23 — draft

取所列美式分区中的一个完整读音；不混拼其他变体。

```plain
Pronunciation source: Collins Dictionary — draft; draft in American English. https://www.collinsdictionary.com/dictionary/english/draft (accent: US; accessed 2026-11-18).
```


### B06-P24 — earnings

取所列美式分区中的一个完整读音；不混拼其他变体。

```plain
Pronunciation source: Collins Dictionary — earnings; earnings in American English. https://www.collinsdictionary.com/dictionary/english/earnings (accent: US; accessed 2026-11-18).
```


### B06-P25 — effect

取所列美式分区中的一个完整读音；不混拼其他变体。 普通英语词形的发音证据不代替现有技术释义、机制或产品官方发音证据。

```plain
Pronunciation source: Collins Dictionary — effect; effect in American English. https://www.collinsdictionary.com/dictionary/english/effect (accent: US; accessed 2026-11-18).
```


### B06-P26 — enum

采用来源第一项 General American /ˈinʌm/；同页另列较少见的 /ɪˈnum/。保留所选完整变体，不拼接两种读法，也不声称 TypeScript 官方指定了唯一读法。 普通英语词形的发音证据不代替现有技术释义、机制或产品官方发音证据。

```plain
Pronunciation source: Wiktionary — enum; English — Pronunciation — General American. https://en.wiktionary.org/w/index.php?title=enum&oldid=94581927 (accent: US; accessed 2026-11-18).
```


### B06-P27 — enumerable

取所列美式分区中的一个完整读音；不混拼其他变体。 普通英语词形的发音证据不代替现有技术释义、机制或产品官方发音证据。

```plain
Pronunciation source: Collins Dictionary — enumerable; enumerable in American English. https://www.collinsdictionary.com/dictionary/english/enumerable (accent: US; accessed 2026-11-18).
```


### B06-P28 — etymology

取所列美式分区中的一个完整读音；不混拼其他变体。

```plain
Pronunciation source: Collins Dictionary — etymology; etymology in American English. https://www.collinsdictionary.com/dictionary/english/etymology (accent: US; accessed 2026-11-18).
```


### B06-P29 — event

取所列美式分区中的一个完整读音；不混拼其他变体。 普通英语词形的发音证据不代替现有技术释义、机制或产品官方发音证据。

```plain
Pronunciation source: Collins Dictionary — event; event in American English. https://www.collinsdictionary.com/dictionary/english/event (accent: US; accessed 2026-11-18).
```


### B06-P30 — exchange

取所列美式分区中的一个完整读音；不混拼其他变体。

```plain
Pronunciation source: Collins Dictionary — exchange; exchange in American English. https://www.collinsdictionary.com/dictionary/english/exchange (accent: US; accessed 2026-11-18).
```


### B06-P31 — financials

取所列美式分区中的一个完整读音；不混拼其他变体。

```plain
Pronunciation source: Collins Dictionary — financials; financials in American English. https://www.collinsdictionary.com/dictionary/english/financials (accent: US; accessed 2026-11-18).
```


### B06-P32 — fiscal year

采用 Oxford fiscal year 页面明确标为 North American English 的 /ˌfɪskl ˈjɪr/，不将前列英式 year 的音标拼入。

```plain
Pronunciation source: Oxford Advanced Learner’s Dictionary — fiscal year; fiscal year — NAmE. https://www.oxfordlearnersdictionaries.com/definition/english/fiscal-year (accent: US; accessed 2026-11-18).
```


### B06-P33 — futures

采用已读英式分区的完整音标；不自行改写为美式。

```plain
Pronunciation source: Collins Dictionary — futures; futures in British English. https://www.collinsdictionary.com/dictionary/english/futures (accent: UK; accessed 2026-11-18).
```


### B06-P34 — gainer

取所列美式分区中的一个完整读音；不混拼其他变体。

```plain
Pronunciation source: Collins Dictionary — gainer; gainer in American English. https://www.collinsdictionary.com/dictionary/english/gainer (accent: US; accessed 2026-11-18).
```


### B06-P35 — general store

采用当前词形的完整词典转写；保留原符号及可选音，不拼接组件。

```plain
Pronunciation source: Oxford Advanced Learner’s Dictionary — general store; general store — British English pronunciation. https://www.oxfordlearnersdictionaries.com/definition/english/general-store (accent: UK; accessed 2026-11-18).
```


### B06-P36 — greens

采用已读英式分区的完整音标；不自行改写为美式。

```plain
Pronunciation source: Collins Dictionary — greens; greens in British English. https://www.collinsdictionary.com/dictionary/english/greens (accent: UK; accessed 2026-11-18).
```


### B06-P37 — heat map

采用当前词形的完整词典转写；保留原符号及可选音，不拼接组件。

```plain
Pronunciation source: Oxford Advanced Learner’s Dictionary — heat map; heat map — British English pronunciation. https://www.oxfordlearnersdictionaries.com/definition/english/heat-map (accent: UK; accessed 2026-11-18).
```


### B06-P38 — highlands

所选 Highlands 英式分区除苏格兰专名外，另有 often not capital 的山地国家或地区义；匹配当前普通高地名词，保留原 Word 大小写。

```plain
Pronunciation source: Collins Dictionary — highlands; highlands in British English. https://www.collinsdictionary.com/dictionary/english/highlands (accent: UK; accessed 2026-11-18).
```


### B06-P39 — index

取所列美式分区中的一个完整读音；不混拼其他变体。

```plain
Pronunciation source: Collins Dictionary — index; index in American English. https://www.collinsdictionary.com/dictionary/english/index (accent: US; accessed 2026-11-18).
```


### B06-P40 — layout

取所列美式分区中的一个完整读音；不混拼其他变体。 普通英语词形的发音证据不代替现有技术释义、机制或产品官方发音证据。

```plain
Pronunciation source: Collins Dictionary — layout; layout in American English. https://www.collinsdictionary.com/dictionary/english/layout (accent: US; accessed 2026-11-18).
```


### B06-P41 — leggings

Oxford Advanced American Dictionary 明列 NAmE /ˈlɛɡɪŋz/；采用当前衣物复数整体，不从 legging 追加词尾。

```plain
Pronunciation source: Oxford Advanced American Dictionary — leggings; leggings — NAmE. https://www.oxfordlearnersdictionaries.com/definition/american_english/leggings (accent: US; accessed 2026-11-18).
```


### B06-P42 — loser

取所列美式分区中的一个完整读音；不混拼其他变体。

```plain
Pronunciation source: Collins Dictionary — loser; loser in American English. https://www.collinsdictionary.com/dictionary/english/loser (accent: US; accessed 2026-11-18).
```


### B06-P43 — lowering

取 Etymology 1 中 lower“降低”的名词/动名词读音 /ˈloʊ.ɚ.ɪŋ/；不取 Etymology 2 表示阴沉、威胁的 /ˈlaʊərɪŋ/。来源支持该词形的普通英语读音，不声称词典已经定义编译过程。 普通英语词形的发音证据不代替现有技术释义、机制或产品官方发音证据。

```plain
Pronunciation source: Wiktionary — lowering; English — Etymology 1 — Pronunciation — General American. https://en.wiktionary.org/w/index.php?title=lowering&oldid=88196875 (accent: US; accessed 2026-11-18).
```


### B06-P44 — margin

取所列美式分区中的一个完整读音；不混拼其他变体。

```plain
Pronunciation source: Collins Dictionary — margin; margin account in American English. https://www.collinsdictionary.com/dictionary/english/margin (accent: US; accessed 2026-11-18).
```


### B06-P45 — market value

采用当前词形的完整词典转写；保留原符号及可选音，不拼接组件。

```plain
Pronunciation source: Oxford Advanced Learner’s Dictionary — market value; market value — British English pronunciation. https://www.oxfordlearnersdictionaries.com/definition/english/market-value (accent: UK; accessed 2026-11-18).
```


### B06-P46 — mashed

当前为“捣碎的”形容词，采用 Oxford mashed adjective；不以 Collins 美式页“土豆泥”的口语名词义冒充当前词性证据。

```plain
Pronunciation source: Oxford Advanced Learner’s Dictionary — mashed; mashed — British English pronunciation. https://www.oxfordlearnersdictionaries.com/definition/english/mashed (accent: UK; accessed 2026-11-18).
```


### B06-P47 — memoization

社区词典中明确标为 General American 的完整 IPA；保留来源转写，使用已读页面的固定修订链接。 普通英语词形的发音证据不代替现有技术释义、机制或产品官方发音证据。

```plain
Pronunciation source: Wiktionary — memoization; English — Pronunciation — General American. https://en.wiktionary.org/w/index.php?title=memoization&oldid=92747904 (accent: US; accessed 2026-11-18).
```


### B06-P48 — merge

取所列美式分区中的一个完整读音；不混拼其他变体。 普通英语词形的发音证据不代替现有技术释义、机制或产品官方发音证据。

```plain
Pronunciation source: Collins Dictionary — merge; merge in American English. https://www.collinsdictionary.com/dictionary/english/merge (accent: US; accessed 2026-11-18).
```


### B06-P49 — method

取所列美式分区中的一个完整读音；不混拼其他变体。 普通英语词形的发音证据不代替现有技术释义、机制或产品官方发音证据。

```plain
Pronunciation source: Collins Dictionary — method; method in American English. https://www.collinsdictionary.com/dictionary/english/method (accent: US; accessed 2026-11-18).
```


### B06-P50 — middleware

采用 Collins 英式整体 /ˈmɪdəlˌwɛə/；仅合并 middle 部分的版面空格，不用 middle 与 software 推导读音。 普通英语词形的发音证据不代替现有技术释义、机制或产品官方发音证据。

```plain
Pronunciation source: Collins Dictionary — middleware; middleware in British English. https://www.collinsdictionary.com/dictionary/english/middleware (accent: UK; accessed 2026-11-18).
```


### B06-P51 — mnemonic

取所列美式分区中的一个完整读音；不混拼其他变体。

```plain
Pronunciation source: Collins Dictionary — mnemonic; mnemonic in American English. https://www.collinsdictionary.com/dictionary/english/mnemonic (accent: US; accessed 2026-11-18).
```


### B06-P52 — operand

取所列美式分区中的一个完整读音；不混拼其他变体。 普通英语词形的发音证据不代替现有技术释义、机制或产品官方发音证据。

```plain
Pronunciation source: Collins Dictionary — operand; operand in American English. https://www.collinsdictionary.com/dictionary/english/operand (accent: US; accessed 2026-11-18).
```


### B06-P53 — option

取所列美式分区中的一个完整读音；不混拼其他变体。

```plain
Pronunciation source: Collins Dictionary — option; option in American English. https://www.collinsdictionary.com/dictionary/english/option (accent: US; accessed 2026-11-18).
```


### B06-P54 — paint

取所列美式分区中的一个完整读音；不混拼其他变体。 普通英语词形的发音证据不代替现有技术释义、机制或产品官方发音证据。

```plain
Pronunciation source: Collins Dictionary — paint; paint in American English. https://www.collinsdictionary.com/dictionary/english/paint (accent: US; accessed 2026-11-18).
```


### B06-P55 — parameter

取所列美式分区中的一个完整读音；不混拼其他变体。 普通英语词形的发音证据不代替现有技术释义、机制或产品官方发音证据。

```plain
Pronunciation source: Collins Dictionary — parameter; parameter in American English. https://www.collinsdictionary.com/dictionary/english/parameter (accent: US; accessed 2026-11-18).
```


### B06-P56 — part of speech

采用当前词形的完整词典转写；保留原符号及可选音，不拼接组件。

```plain
Pronunciation source: Oxford Advanced Learner’s Dictionary — part of speech; part of speech — British English pronunciation. https://www.oxfordlearnersdictionaries.com/definition/english/part-of-speech (accent: UK; accessed 2026-11-18).
```


### B06-P57 — patch

取所列美式分区中的一个完整读音；不混拼其他变体。 普通英语词形的发音证据不代替现有技术释义、机制或产品官方发音证据。

```plain
Pronunciation source: Collins Dictionary — patch; patch in American English 1. https://www.collinsdictionary.com/dictionary/english/patch (accent: US; accessed 2026-11-18).
```


### B06-P58 — persist

取所列美式分区中的一个完整读音；不混拼其他变体。 普通英语词形的发音证据不代替现有技术释义、机制或产品官方发音证据。

```plain
Pronunciation source: Collins Dictionary — persist; persist in American English. https://www.collinsdictionary.com/dictionary/english/persist (accent: US; accessed 2026-11-18).
```


### B06-P59 — portfolio

取所列美式分区中的一个完整读音；不混拼其他变体。

```plain
Pronunciation source: Collins Dictionary — portfolio; portfolio in American English. https://www.collinsdictionary.com/dictionary/english/portfolio (accent: US; accessed 2026-11-18).
```


### B06-P60 — position

取所列美式分区中的一个完整读音；不混拼其他变体。

```plain
Pronunciation source: Collins Dictionary — position; position in American English. https://www.collinsdictionary.com/dictionary/english/position (accent: US; accessed 2026-11-18).
```


### B06-P61 — primitive

取所列美式分区中的一个完整读音；不混拼其他变体。 普通英语词形的发音证据不代替现有技术释义、机制或产品官方发音证据。

```plain
Pronunciation source: Collins Dictionary — primitive; primitive in American English. https://www.collinsdictionary.com/dictionary/english/primitive (accent: US; accessed 2026-11-18).
```


### B06-P62 — property

取所列美式分区中的一个完整读音；不混拼其他变体。 普通英语词形的发音证据不代替现有技术释义、机制或产品官方发音证据。

```plain
Pronunciation source: Collins Dictionary — property; property in American English. https://www.collinsdictionary.com/dictionary/english/property (accent: US; accessed 2026-11-18).
```


### B06-P63 — put option

采用所选英式整体 /pʊt ˈɒpʃən/；仅去除来源把 schwa 与相邻辅音分开的版面空格，不拼接 put 与 option。

```plain
Pronunciation source: Collins Dictionary — put option; put option in British English. https://www.collinsdictionary.com/dictionary/english/put-option (accent: UK; accessed 2026-11-18).
```


### B06-P64 — quadratic

取所列美式分区中的一个完整读音；不混拼其他变体。 普通英语词形的发音证据不代替现有技术释义、机制或产品官方发音证据。

```plain
Pronunciation source: Collins Dictionary — quadratic; quadratic in American English. https://www.collinsdictionary.com/dictionary/english/quadratic (accent: US; accessed 2026-11-18).
```


### B06-P65 — query

取所列美式分区中的一个完整读音；不混拼其他变体。 普通英语词形的发音证据不代替现有技术释义、机制或产品官方发音证据。

```plain
Pronunciation source: Collins Dictionary — query; query in American English. https://www.collinsdictionary.com/dictionary/english/query (accent: US; accessed 2026-11-18).
```


### B06-P66 — quicksort

社区词典中明确标为 General American 的完整 IPA；保留来源转写，使用已读页面的固定修订链接。 普通英语词形的发音证据不代替现有技术释义、机制或产品官方发音证据。

```plain
Pronunciation source: Wiktionary — quicksort; English — Pronunciation — General American. https://en.wiktionary.org/w/index.php?title=quicksort&oldid=95458003 (accent: US; accessed 2026-11-18).
```


### B06-P67 — reconciliation

取所列美式分区中的一个完整读音；不混拼其他变体。 普通英语词形的发音证据不代替现有技术释义、机制或产品官方发音证据。

```plain
Pronunciation source: Collins Dictionary — reconciliation; reconciliation in American English. https://www.collinsdictionary.com/dictionary/english/reconciliation (accent: US; accessed 2026-11-18).
```


### B06-P68 — recursion

取所列美式分区中的一个完整读音；不混拼其他变体。 普通英语词形的发音证据不代替现有技术释义、机制或产品官方发音证据。

```plain
Pronunciation source: Collins Dictionary — recursion; recursion in American English. https://www.collinsdictionary.com/dictionary/english/recursion (accent: US; accessed 2026-11-18).
```


### B06-P69 — repository

取所列美式分区中的一个完整读音；不混拼其他变体。 普通英语词形的发音证据不代替现有技术释义、机制或产品官方发音证据。

```plain
Pronunciation source: Collins Dictionary — repository; repository in American English. https://www.collinsdictionary.com/dictionary/english/repository (accent: US; accessed 2026-11-18).
```


### B06-P70 — return

取所列美式分区中的一个完整读音；不混拼其他变体。

```plain
Pronunciation source: Collins Dictionary — return; return in American English. https://www.collinsdictionary.com/dictionary/english/return (accent: US; accessed 2026-11-18).
```


### B06-P71 — ruins

采用已读英式分区的完整音标；不自行改写为美式。

```plain
Pronunciation source: Collins Dictionary — ruins; ruins in British English. https://www.collinsdictionary.com/dictionary/english/ruins (accent: UK; accessed 2026-11-18).
```


### B06-P72 — shareholder

取所列美式分区中的一个完整读音；不混拼其他变体。

```plain
Pronunciation source: Collins Dictionary — shareholder; shareholder in American English. https://www.collinsdictionary.com/dictionary/english/shareholder (accent: US; accessed 2026-11-18).
```


### B06-P73 — shorthand

取所列美式分区中的一个完整读音；不混拼其他变体。 普通英语词形的发音证据不代替现有技术释义、机制或产品官方发音证据。

```plain
Pronunciation source: Collins Dictionary — shorthand; shorthand in American English. https://www.collinsdictionary.com/dictionary/english/shorthand (accent: US; accessed 2026-11-18).
```


### B06-P74 — source code

采用 Oxford 首列完整英式 /ˈsɔːs kəʊd/。同页后列出现 /ˈsɔːrs kəʊd/，本批不将其当作已核实的标准美式整体读音，也不自行把 kəʊd 改成 koʊd。 普通英语词形的发音证据不代替现有技术释义、机制或产品官方发音证据。

```plain
Pronunciation source: Oxford Advanced Learner’s Dictionary — source code; source code — British English pronunciation. https://www.oxfordlearnersdictionaries.com/definition/english/source-code (accent: UK; accessed 2026-11-18).
```


### B06-P75 — stack

取所列美式分区中的一个完整读音；不混拼其他变体。 普通英语词形的发音证据不代替现有技术释义、机制或产品官方发音证据。

```plain
Pronunciation source: Collins Dictionary — stack; stack in American English. https://www.collinsdictionary.com/dictionary/english/stack (accent: US; accessed 2026-11-18).
```


### B06-P76 — symbol

取所列美式分区中的一个完整读音；不混拼其他变体。

```plain
Pronunciation source: Collins Dictionary — symbol; symbol in American English. https://www.collinsdictionary.com/dictionary/english/symbol (accent: US; accessed 2026-11-18).
```


### B06-P77 — syntax

取所列美式分区中的一个完整读音；不混拼其他变体。 普通英语词形的发音证据不代替现有技术释义、机制或产品官方发音证据。

```plain
Pronunciation source: Collins Dictionary — syntax; syntax in American English. https://www.collinsdictionary.com/dictionary/english/syntax (accent: US; accessed 2026-11-18).
```


### B06-P78 — teriyaki

取所列美式分区中的一个完整读音；不混拼其他变体。

```plain
Pronunciation source: Collins Dictionary — teriyaki; teriyaki in American English. https://www.collinsdictionary.com/dictionary/english/teriyaki (accent: US; accessed 2026-11-18).
```


### B06-P79 — test

取所列美式分区中的一个完整读音；不混拼其他变体。 普通英语词形的发音证据不代替现有技术释义、机制或产品官方发音证据。

```plain
Pronunciation source: Collins Dictionary — test; test in American English 1. https://www.collinsdictionary.com/dictionary/english/test (accent: US; accessed 2026-11-18).
```


### B06-P80 — text-to-speech

采用当前词形的完整词典转写；保留原符号及可选音，不拼接组件。 普通英语词形的发音证据不代替现有技术释义、机制或产品官方发音证据。

```plain
Pronunciation source: Oxford Advanced Learner’s Dictionary — text-to-speech; text-to-speech — British English pronunciation. https://www.oxfordlearnersdictionaries.com/definition/english/text-to-speech (accent: UK; accessed 2026-11-18).
```


### B06-P81 — token

取所列美式分区中的一个完整读音；不混拼其他变体。 普通英语词形的发音证据不代替现有技术释义、机制或产品官方发音证据。

```plain
Pronunciation source: Collins Dictionary — token; token in American English. https://www.collinsdictionary.com/dictionary/english/token (accent: US; accessed 2026-11-18).
```


### B06-P82 — transaction

取所列美式分区中的一个完整读音；不混拼其他变体。

```plain
Pronunciation source: Collins Dictionary — transaction; transaction in American English. https://www.collinsdictionary.com/dictionary/english/transaction (accent: US; accessed 2026-11-18).
```


### B06-P83 — tree

取所列美式分区中的一个完整读音；不混拼其他变体。 普通英语词形的发音证据不代替现有技术释义、机制或产品官方发音证据。

```plain
Pronunciation source: Collins Dictionary — tree; tree in American English. https://www.collinsdictionary.com/dictionary/english/tree (accent: US; accessed 2026-11-18).
```


### B06-P84 — triage

取所列美式分区中的一个完整读音；不混拼其他变体。 普通英语词形的发音证据不代替现有技术释义、机制或产品官方发音证据。

```plain
Pronunciation source: Collins Dictionary — triage; triage in American English. https://www.collinsdictionary.com/dictionary/english/triage (accent: US; accessed 2026-11-18).
```


### B06-P85 — type

取所列美式分区中的一个完整读音；不混拼其他变体。 普通英语词形的发音证据不代替现有技术释义、机制或产品官方发音证据。

```plain
Pronunciation source: Collins Dictionary — type; type in American English. https://www.collinsdictionary.com/dictionary/english/type (accent: US; accessed 2026-11-18).
```


### B06-P86 — variance

取所列美式分区中的一个完整读音；不混拼其他变体。 普通英语词形的发音证据不代替现有技术释义、机制或产品官方发音证据。

```plain
Pronunciation source: Collins Dictionary — variance; variance in American English. https://www.collinsdictionary.com/dictionary/english/variance (accent: US; accessed 2026-11-18).
```


### B06-P87 — volatility

取所列美式分区中的一个完整读音；不混拼其他变体。

```plain
Pronunciation source: Collins Dictionary — volatility; volatility in American English. https://www.collinsdictionary.com/dictionary/english/volatility (accent: US; accessed 2026-11-18).
```


### B06-P88 — volume

取所列美式分区中的一个完整读音；不混拼其他变体。

```plain
Pronunciation source: Collins Dictionary — volume; volume in American English. https://www.collinsdictionary.com/dictionary/english/volume (accent: US; accessed 2026-11-18).
```


### B06-P89 — watchlist

社区词典中明确标为 General American 的完整 IPA；保留来源转写，使用已读页面的固定修订链接。

```plain
Pronunciation source: Wiktionary — watchlist; English — Pronunciation — General American. https://en.wiktionary.org/w/index.php?title=watchlist&oldid=91656417 (accent: US; accessed 2026-11-18).
```


### B06-P90 — wealth

取所列美式分区中的一个完整读音；不混拼其他变体。

```plain
Pronunciation source: Collins Dictionary — wealth; wealth in American English. https://www.collinsdictionary.com/dictionary/english/wealth (accent: US; accessed 2026-11-18).
```


### B06-P91 — wetlands

采用已读英式分区的完整音标；不自行改写为美式。

```plain
Pronunciation source: Collins Dictionary — wetlands; wetlands in British English. https://www.collinsdictionary.com/dictionary/english/wetlands (accent: UK; accessed 2026-11-18).
```


### B06-P92 — workaround

来源为 /ˈwɜrkə ˌraʊnd/，仅合并词内版面空格为 /ˈwɜrkəˌraʊnd/，不改变音素或重音。 普通英语词形的发音证据不代替现有技术释义、机制或产品官方发音证据。

```plain
Pronunciation source: Collins Dictionary — workaround; workaround in American English. https://www.collinsdictionary.com/dictionary/english/workaround (accent: US; accessed 2026-11-18).
```


## 本批暂缓项目

以下 50 项没有拟写 pronunciation。它们与 92 项提案互斥；批准 B06-P01–B06-P92 不授权修改暂缓项目。词典没有提供本轮可用证据，不代表该词没有英语读法。


| 稳定 ID / Word | 当前 Meaning | 实际限制 |
|:---|:---|:---|
| [accessor-tech](../entries/accessor-tech.yml) / accessor | 属性访问器 | Collins、Cambridge、Oxford 未取得可采用正文；Wiktionary 的 IPA 出现在拉丁语分区，不能当作英语访问器名词的读音。 |
| [annual-percentage-yield-finance](../entries/annual-percentage-yield-finance.yml) / annual percentage yield | 年化收益率；考虑利息及复利频率后表示的年度收益比例。 | 本轮尝试的对应词典页面不可读取，或没有与当前词形、意义及明确口音匹配的整体 IPA；不猜读、不拼接组件或词尾。 |
| [arowana-general](../entries/arowana-general.yml) / arowana | 龙鱼。 | 本轮尝试的对应词典页面不可读取，或没有与当前词形、意义及明确口音匹配的整体 IPA；不猜读、不拼接组件或词尾。 |
| [ask-price-finance](../entries/ask-price-finance.yml) / ask price | 卖方报价；行情中卖方愿意接受的卖出报价。 | 本轮尝试的对应词典页面不可读取，或没有与当前词形、意义及明确口音匹配的整体 IPA；不猜读、不拼接组件或词尾。 |
| [bid-price-finance](../entries/bid-price-finance.yml) / bid price | 买方报价；行情中买方愿意支付的买入报价。 | Collins 已读取的对应页面未提供可采用的美式整体 IPA；英式查证也未取得匹配当前词形及意义的完整证据。补查来源若仅列释义或组件，不据此拼接。 |
| [binary-search-algorithm](../entries/binary-search-algorithm.yml) / binary search | 二分查找；在有序区间中比较中点并排除一半候选。 | 本轮尝试的对应词典页面不可读取，或没有与当前词形、意义及明确口音匹配的整体 IPA；不猜读、不拼接组件或词尾。 |
| [boots-clothing](../entries/boots-clothing.yml) / boots | 靴子 | Collins boots 的独立英美名词分区表示旅馆擦鞋仆役；不能当作当前“靴子”复数的意义证据。Oxford 对应页面不可读取。 |
| [call-stack-javascript](../entries/call-stack-javascript.yml) / call stack | 调用栈；按后进先出顺序记录函数调用和返回位置。 | 本轮尝试的对应词典页面不可读取，或没有与当前词形、意义及明确口音匹配的整体 IPA；不猜读、不拼接组件或词尾。 |
| [cooking-pot-cooking](../entries/cooking-pot-cooking.yml) / cooking pot | 烹饪锅 | 本轮尝试的对应词典页面不可读取，或没有与当前词形、意义及明确口音匹配的整体 IPA；不猜读、不拼接组件或词尾。 |
| [coordinates-ui](../entries/coordinates-ui.yml) / coordinates | 坐标 | Collins 独立 coordinates 分区表示协调搭配的衣物/物件，不是坐标复数；Oxford 对应页面不可读取。 |
| [corrupted-combat](../entries/corrupted-combat.yml) / corrupted | 被腐化的 | 本轮尝试的对应词典页面不可读取，或没有与当前词形、意义及明确口音匹配的整体 IPA；不猜读、不拼接组件或词尾。 |
| [debounce-javascript](../entries/debounce-javascript.yml) / debounce | 防抖；把短时间内的连续调用合并为较少的执行。 | 本轮尝试的对应词典页面不可读取，或没有与当前词形、意义及明确口音匹配的整体 IPA；不猜读、不拼接组件或词尾。 |
| [deep-copy-javascript](../entries/deep-copy-javascript.yml) / deep copy | 深拷贝；复制嵌套数据，使相应对象引用与原数据隔离。 | 本轮尝试的对应词典页面不可读取，或没有与当前词形、意义及明确口音匹配的整体 IPA；不猜读、不拼接组件或词尾。 |
| [disguised-combat](../entries/disguised-combat.yml) / disguised | 伪装的 | 本轮尝试的对应词典页面不可读取，或没有与当前词形、意义及明确口音匹配的整体 IPA；不猜读、不拼接组件或词尾。 |
| [durability-equipment](../entries/durability-equipment.yml) / durability | 耐用性；耐久度。 | Collins durability 分区没有整体 IPA，后续出现的是 durable 的读音；不把原形音标当作派生词整体。Oxford 对应页面不可读取。 |
| [earnings-per-share-finance](../entries/earnings-per-share-finance.yml) / earnings per share | 每股收益；将相应利润按每股口径表示的财务指标。 | Collins 已读取的对应页面未提供可采用的美式整体 IPA；英式查证也未取得匹配当前词形及意义的完整证据。补查来源若仅列释义或组件，不据此拼接。 |
| [enhancement-general](../entries/enhancement-general.yml) / enhancement | 强化、增强 | 本轮尝试的对应词典页面不可读取，或没有与当前词形、意义及明确口音匹配的整体 IPA；不猜读、不拼接组件或词尾。 |
| [event-loop-javascript](../entries/event-loop-javascript.yml) / event loop | 事件循环；协调 JavaScript 任务执行与宿主活动的机制。 | Wiktionary 有短语整体 IPA，但未明确标口音；不在默认 US 字段中自动采用。Collins、Cambridge 和 Oxford 未取得合适正文。 |
| [exchange-traded-fund-finance](../entries/exchange-traded-fund-finance.yml) / exchange-traded fund | 交易所交易基金；份额可以在交易所按市场价格买卖的基金。 | Collins 已读取的对应页面未提供可采用的美式整体 IPA；英式查证也未取得匹配当前词形及意义的完整证据。补查来源若仅列释义或组件，不据此拼接。 |
| [flexbox-css](../entries/flexbox-css.yml) / flexbox | 弹性布局；主要沿一个维度分配空间和排列项目的布局。 | 本轮尝试的对应词典页面不可读取，或没有与当前词形、意义及明确口音匹配的整体 IPA；不猜读、不拼接组件或词尾。 |
| [garbage-collection-javascript](../entries/garbage-collection-javascript.yml) / garbage collection | 垃圾回收；自动回收不再可达的对象所占内存。 | Collins 已读取的对应页面未提供可采用的美式整体 IPA；英式查证也未取得匹配当前词形及意义的完整证据。补查来源若仅列释义或组件，不据此拼接。 |
| [greaves-clothing](../entries/greaves-clothing.yml) / greaves | 护胫 | Collins greaves 的整体 IPA 对应熬油残渣或人名，不是护胫复数；Oxford 对应页面不可读取。 |
| [hoisting-javascript](../entries/hoisting-javascript.yml) / hoisting | 提升；声明在执行前影响所在作用域的行为。 | Wiktionary 有整体 IPA，但未明确标 US 或 UK；Collins 当前词形与 Oxford 未取得可采用的对应口音整体证据。 |
| [holdings-finance](../entries/holdings-finance.yml) / holdings | 持有的资产；投资组合中的证券等投资项目。 | Collins 已读取的对应页面未提供可采用的美式整体 IPA；英式查证也未取得匹配当前词形及意义的完整证据。补查来源若仅列释义或组件，不据此拼接。 |
| [initializer-tech](../entries/initializer-tech.yml) / initializer | 初始化表达式；在创建或声明值时给出初始值的表达式。 | 本轮尝试的对应词典页面不可读取，或没有与当前词形、意义及明确口音匹配的整体 IPA；不猜读、不拼接组件或词尾。 |
| [iterable-tech](../entries/iterable-tech.yml) / iterable | 可迭代的 | 本轮尝试的对应词典页面不可读取，或没有与当前词形、意义及明确口音匹配的整体 IPA；不猜读、不拼接组件或词尾。 |
| [linked-list-tech](../entries/linked-list-tech.yml) / linked list | 链表；通过节点之间的链接组织元素的线性结构。 | Collins 已读取的对应页面未提供可采用的美式整体 IPA；英式查证也未取得匹配当前词形及意义的完整证据。补查来源若仅列释义或组件，不据此拼接。 |
| [margin-account-finance](../entries/margin-account-finance.yml) / margin account | 保证金账户；允许按规则以账户资产作担保向券商借款交易的账户。 | Collins 已读取的对应页面未提供可采用的美式整体 IPA；英式查证也未取得匹配当前词形及意义的完整证据。补查来源若仅列释义或组件，不据此拼接。 |
| [market-capitalization-finance](../entries/market-capitalization-finance.yml) / market capitalization | 市值；按市场价格与相应数量口径计算的总体价值。 | Collins 已读取的对应页面未提供可采用的美式整体 IPA；英式查证也未取得匹配当前词形及意义的完整证据。补查来源若仅列释义或组件，不据此拼接。 |
| [microtask-javascript](../entries/microtask-javascript.yml) / microtask | 微任务；在规定的检查点处理、通常先于下一个任务的工作。 | 本轮尝试的对应词典页面不可读取，或没有与当前词形、意义及明确口音匹配的整体 IPA；不猜读、不拼接组件或词尾。 |
| [namespace-typescript](../entries/namespace-typescript.yml) / namespace | 命名空间；组织一组命名成员、通过限定名称访问的作用域结构。 | 本轮尝试的对应词典页面不可读取，或没有与当前词形、意义及明确口音匹配的整体 IPA；不猜读、不拼接组件或词尾。 |
| [net-assets-finance](../entries/net-assets-finance.yml) / net assets | 净资产；资产扣除负债后的净值。 | Collins 已读取的对应页面未提供可采用的美式整体 IPA；英式查证也未取得匹配当前词形及意义的完整证据。补查来源若仅列释义或组件，不据此拼接。 |
| [net-income-finance](../entries/net-income-finance.yml) / net income | 净利润；扣除相关成本、费用和税项后的利润。 | Collins 已读取的对应页面未提供可采用的美式整体 IPA；英式查证也未取得匹配当前词形及意义的完整证据。补查来源若仅列释义或组件，不据此拼接。 |
| [noodles-food](../entries/noodles-food.yml) / noodles | 面条 | Wiktionary 有整体英语 IPA 与相邻 US 音频，但 IPA 本身未单独标口音；本轮未取得满足明确口音规则的词典正文，不以音频标签代替核实。 |
| [option-chain-finance](../entries/option-chain-finance.yml) / option chain | 期权链；按到期日、行权价等条件组织的期权合约报价列表。 | 本轮尝试的对应词典页面不可读取，或没有与当前词形、意义及明确口音匹配的整体 IPA；不猜读、不拼接组件或词尾。 |
| [pick-up-general](../entries/pick-up-general.yml) / pick up | 捡起 | 本轮尝试的对应词典页面不可读取，或没有与当前词形、意义及明确口音匹配的整体 IPA；不猜读、不拼接组件或词尾。 |
| [pre-market-finance](../entries/pre-market-finance.yml) / pre-market | 盘前的；发生在常规交易时段开始之前的。 | 本轮尝试的对应词典页面不可读取，或没有与当前词形、意义及明确口音匹配的整体 IPA；不猜读、不拼接组件或词尾。 |
| [profit-and-loss-finance](../entries/profit-and-loss-finance.yml) / profit and loss | 盈亏；投资或经营中的盈利与亏损结果。 | Collins 已读取的对应页面未提供可采用的美式整体 IPA；英式查证也未取得匹配当前词形及意义的完整证据。补查来源若仅列释义或组件，不据此拼接。 |
| [ranged-combat](../entries/ranged-combat.yml) / ranged | 远程 | Collins ranged 美式分区为建筑分层砌筑义，不匹配远程战斗义；不套用或由 range 拼接词尾。 |
| [saute-cooking](../entries/saute-cooking.yml) / sauté | 用少量油快速翻炒或煎炒。 | 本轮尝试的对应词典页面不可读取，或没有与当前词形、意义及明确口音匹配的整体 IPA；不猜读、不拼接组件或词尾。 |
| [shallow-copy-javascript](../entries/shallow-copy-javascript.yml) / shallow copy | 浅拷贝；复制外层结构，嵌套对象仍共享引用。 | 本轮尝试的对应词典页面不可读取，或没有与当前词形、意义及明确口音匹配的整体 IPA；不猜读、不拼接组件或词尾。 |
| [soundness-type-system](../entries/soundness-type-system.yml) / soundness | 健全性；类型检查所接受的程序不会违反该系统承诺的类型安全性质。 | 本轮尝试的对应词典页面不可读取，或没有与当前词形、意义及明确口音匹配的整体 IPA；不猜读、不拼接组件或词尾。 |
| [stack-trace-tech](../entries/stack-trace-tech.yml) / stack trace | 调用栈跟踪；记录一条执行路径上活动调用帧的信息。 | 本轮尝试的对应词典页面不可读取，或没有与当前词形、意义及明确口音匹配的整体 IPA；不猜读、不拼接组件或词尾。 |
| [strike-price-finance](../entries/strike-price-finance.yml) / strike price | 行权价；期权合约约定买卖标的时使用的价格。 | Collins 已读取的对应页面未提供可采用的美式整体 IPA；英式查证也未取得匹配当前词形及意义的完整证据。补查来源若仅列释义或组件，不据此拼接。 |
| [symlink-tech](../entries/symlink-tech.yml) / symlink | 符号链接 | 本轮尝试的对应词典页面不可读取，或没有与当前词形、意义及明确口音匹配的整体 IPA；不猜读、不拼接组件或词尾。 |
| [unfuse-general](../entries/unfuse-general.yml) / unfuse | 解除融合 | 本轮尝试的对应词典页面不可读取，或没有与当前词形、意义及明确口音匹配的整体 IPA；不猜读、不拼接组件或词尾。 |
| [waistguard-clothing](../entries/waistguard-clothing.yml) / waistguard | 护腰 | 本轮尝试的对应词典页面不可读取，或没有与当前词形、意义及明确口音匹配的整体 IPA；不猜读、不拼接组件或词尾。 |
| [warp-general](../entries/warp-general.yml) / warp | 传送；游戏中快速移动到另一地点。 | 当前为传送游戏义；所读 Collins 片段未建立匹配当前意义的证据，不把不同含义直接混同。 |
| [withdrawable-finance](../entries/withdrawable-finance.yml) / withdrawable | 可提取的；在当前规则下允许从账户取出的。 | 本轮尝试的对应词典页面不可读取，或没有与当前词形、意义及明确口音匹配的整体 IPA；不猜读、不拼接组件或词尾。 |
| [year-over-year-finance](../entries/year-over-year-finance.yml) / year-over-year | 同比的；与上一年相同期间相比的。 | 本轮尝试的对应词典页面不可读取，或没有与当前词形、意义及明确口音匹配的整体 IPA；不猜读、不拼接组件或词尾。 |


## 批准后实施与验证

1. 等待明确确认 B06-P01–B06-P92，或确认其中指定编号。每项同时审批其 pronunciation 值与对应 Notes 来源行；本批不包含关系、助记、词源、构词或音频变更。
2. 实施前重读批准内容与当前仓库，检查 Git 状态并核对稳定 ID、Word、词性、Meaning 和 pronunciation 仍符合基线。若目标已被其他工作补齐或内容改变，报告并协调，不覆盖并发编辑。保存词条原始字节、解析值及现有媒体哈希。
3. 仅为批准目标新增对应 pronunciation，并在保留原 Notes 全文的前提下追加来源一次。保留其他字段与源文本、所有录音字节、无关文件和暂缓项，不修改实现、模板、依赖、历史批准文案或生成音频。
4. 运行 `go test ./...`、`go vet ./...`、`go run scripts/validate/main.go`、`go run scripts/build/main.go` 和 `git diff --check`；新文件另查格式，不手工编辑 dist。
5. 无源变化连续构建两次，逐字节比较导出。核对仍为 1,320 个稳定 ID、1,320 行数据和 17 列；仅获批行的 Pronunciation 与 Notes 列改变，Audio、StudyOrder、RelatedWords、Tags 等保持一致。核对所有已存在媒体哈希与非目标词条字节不变。
6. 更新本计划的批准、实施与验证状态，记录实际数量及失败，不把上述预计计数写成已经完成。报告静态和导出验证边界；不声称已完成人工听审或真实 Anki 导入、显示、播放。不暂存、提交、推送、复制音频到 Anki 或修改 Anki 集合。

## 规划检查

- 已通过：92 项稳定 ID、当前词形、词性、释义、缺失状态、拟写音标与来源追加行的一致性检查；142 项查证账目由 92 项提案与 50 项暂缓组成，编号与稳定 ID 无重复。
- 已通过：157 个本地链接及同文件锚点、编号唯一性、新文件换行与尾随空格检查，以及 `git diff --check`。与规划快照比较，仅新增本计划；全部 1,320 个词条、1,363 个既有 MP3、实现、既有文档和导出保持不变。
- 规划阶段未运行构建、音频生成或修改 entries；词典文本核对不包含人工听音或真实 Anki 测试。

## 实施与验证记录

用户明确确认 B06-P01–B06-P92 并授权实施，92 项全部完成，0 项失败。实施日期：2026-11-18。上述范围表、预期计数及规划检查保留为实施前历史基线；批准内容、50 项暂缓清单及其来源文案不变。

| 检查 | 实际结果 |
|:---|:---|
| 修改范围 | 92 个现有词条新增 pronunciation 并追加批准的 Notes 来源行；另更新本计划状态与结果 |
| 口音 | 71 项美式无前缀；21 项英式保留 UK 前缀；已有音标不变 |
| 全库 | 仍为 1,320 个稳定 ID：英语 1,077，日语 243 |
| 英语已有 / 缺失音标 | 778 / 299；缺失由本批暂缓 50 项和其他未逐项查证 249 项组成 |
| TOTK 已有 / 缺失音标 | 359 / 194；本批补全 8 项普通词或短语，不推定专名英文发音 |
| 原始数据保全 | 按实施前快照核对，其他 1,228 个词条字节不变；目标词条解析值仅 pronunciation 与 Notes 改变，原 Notes 内容完整保留，来源行各追加一次 |
| 音频保全 | 全部 1,363 个既有 MP3 哈希保持不变，包括 media 与 archives；未生成、解码、听审或替换录音 |
| Go 检查 | go test ./...、go vet ./... 通过 |
| 词库检查 | go run scripts/validate/main.go 通过；15 条既有同词警告仍保留，无新增警告 |
| 导出检查 | 通过现有构建命令生成 1,320 行数据、17 列；仅 92 个批准行的 Pronunciation、Notes 列改变 |
| 其余导出字段 | ID 集合和顺序、Audio、StudyOrder、RelatedWords、Tags、TTSLanguage 等保持不变 |
| 确定性 | 两次无源变化构建的导出逐字节一致；SHA-256 为 776c778da734d2de6be5d87866a731ac06823834a92c585a9b30288f34a8d47c |

已核对每项实际写入值与批准表格、同编号来源行完全一致。既有未提交工作、其他文档、实现、模板及依赖保持不变。未新增词条、关系、先修或助记。

本记录仅报告源数据、静态检查与导出结果；未进行人工发音听审或真实 Anki 导入、显示、播放验证，未修改 Anki 集合。本计划 158 个本地链接及锚点、新文件格式与 git diff --check 均已通过检查。
