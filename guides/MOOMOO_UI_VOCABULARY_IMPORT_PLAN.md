# 富途牛牛界面词汇导入计划

状态: **已批准并实施，静态与导出验证通过。** 日期: 2026-11-06。语言: 英语。音频: 本次未请求，未生成。

根据用户提供的 8 张 Moomoo/富途牛牛英文界面截图整理。规划阶段仅授权创建本计划；用户随后明确批准实施，已创建批准的新词条并重建导出。原有词条、media、archives、实现、模板和依赖均保持不变，未操作 Git 提交或 Anki。下文保留规划时的基线与具体规格；实际结果见[实施记录](#实施记录)。后续截图使用新文件保存下一批。

## 导航

- [确认记录](#确认记录)
- [来源截图](#来源截图)
- [筛选与复用](#筛选与复用)
- [候选内容](#候选内容)
- [缩写与组合处理](#缩写与组合处理)
- [可选记忆辅助](#可选记忆辅助)
- [关联候选](#关联候选)
- [待核实与排除](#待核实与排除)
- [参考来源](#参考来源)
- [来源到目标台账](#来源到目标台账)
- [实施与验证](#实施与验证)
- [规划检查](#规划检查)
- [实施记录](#实施记录)

## 确认记录

2026-11-06，用户确认以下内容:

- 账户与资金、行情与观察、财报与预测、期权与报价、通用界面词和复用: C01–C79，共 79 个学习目标，包含 78 个拟新增词条和 1 个复用项。确认范围包括候选表中的 Word、Meaning、优先级和有用的先修安排。
- 例句与 Notes 草案: 按表中具体内容保留；复用 pending-tech 仍不改写已有例句、Notes、标签或录音。
- 缩写与组合处理: 按当前分组、规范化和来源形式记录，不额外增加组合卡。

首次确认覆盖词条、例句、Notes 和缩写处理。随后用户于 2026-11-06 单独确认 W01–W04、N01 的全部 5 项记忆辅助具体文案，以及 R01–R11 的全部 11 对关联。两组独立批准均已记录，当前没有待审阅的拟议内容。Excess Liquidity、Turnover、Smart Money、Short / Long 等未核实项目仍按原限制暂缓，不因内容确认而视为事实已查证。

实施状态: 用户明确批准实施后，已新增 78 个词条、复用 pending-tech，写入 5 项批准的记忆辅助与 11 对关联。未生成音频。

## 来源截图

8 张截图均可读取。没有复制或编辑截图。以下路径位于用户 Downloads，不在 Git 仓库内；这些本地链接只保证在当前机器上可用。只保留有学习价值的界面文字，账户编号、持仓金额、行情数值、收益率和新闻跑马灯内容不写进例句或词条。

### S1

[Screenshot_20261106-173510.webp](/Users/mazey/Downloads/Screenshot_20261106-173510.webp) — 公司财报与预测。

本页辨认范围: Earnings、Financials、Revenue、Net Income、Estimate、Actual、YoY、EPS、EBIT、Shareholders、Dividends 等。

### S2

[Screenshot_20261106-173312.webp](/Users/mazey/Downloads/Screenshot_20261106-173312.webp) — 加密资产市场。

本页辨认范围: Crypto、Market Cap、Volume、Fear & Greed Index、Top Gainers、Top Losers、Turnover、Heat Map 等。

### S3

[Screenshot_20261106-173302.webp](/Users/mazey/Downloads/Screenshot_20261106-173302.webp) — Discover 第二页。

本页辨认范围: Smart Money、Buffett Holdings、1Y Return、Trade、Crypto。

### S4

[Screenshot_20261106-173238.webp](/Users/mazey/Downloads/Screenshot_20261106-173238.webp) — Discover 第一页。

本页辨认范围: Discover、Community、Wealth；其他基础导航单独记作排除项。

### S5

[Screenshot_20261106-173228.webp](/Users/mazey/Downloads/Screenshot_20261106-173228.webp) — 期权报价与期权链。

本页辨认范围: Options、Chain、Call、Put、Strike、Bid、Ask、Closing Price、Pre-Mkt、Futures、At Discount 等。

### S6

[Screenshot_20261106-173207.webp](/Users/mazey/Downloads/Screenshot_20261106-173207.webp) — 账户资产明细。

本页辨认范围: Net Assets、Market Value、Position P/L、Funds on Hold、Excess Liquidity、Withdrawable、Max Buying Power、Cash BP、Cash Sweep、APY、Interest、Portfolio、MV/QTY 等。

### S7

[Screenshot_20261106-173158.webp](/Users/mazey/Downloads/Screenshot_20261106-173158.webp) — 账户列表。

本页辨认范围: Accounts、Assets、Deposit、Cash Transfer、Exchange、Margin Account、Cash Account。

### S8

[Screenshot_20261106-173147.webp](/Users/mazey/Downloads/Screenshot_20261106-173147.webp) — 自选列表。

本页辨认范围: Watchlists、Positions、Index、Short、Long、Price、% Chg；证券名称、代码与金额不作学习目标。

位置引用采用截图编号和界面区块，不为图片编造行号。截断、仅露出一部分的标签不补全。观察到的页面文字与参考资料中的扩展概念分开记录；不把示例截图中的旧结果当作当前行情。

## 筛选与复用

已读取 [AGENTS.md](../AGENTS.md)、[README.md](../README.md)、[Entry schema](ENTRY_SCHEMA.md)、[Learning order](LEARNING_ORDER.md) 和 [Maintenance prompts](MAINTENANCE_PROMPTS.md#collect-vocabulary-from-sources)，并核对现有解析器、17 列导出及双向关联实现。初始 Git 状态干净，现有集合为 **1,242 个词条: 999 个英语、243 个日语**。搜索使用整个集合的稳定 ID、Word、Meaning、Context 和 Notes，不以文件名代替身份，也不只做字符串匹配。

词库目前偏技术及游戏语境，没有覆盖本批主要金融义项。编程中的 call stack、return type、prototype chain 和 Symbol 不替代金融 call、return、option chain 和 symbol。`pending-tech` 的待处理/待完成核心义可以复用，保留原内容和录音。本批不提议重写已有词条。拟新增 symbol-market 与既有 symbol-javascript 同形但词义不同，预计产生一条合理的重复词警告；两者不合并。

学习方向仍为 Word → Meaning。优先级 1 是日常账户和基础财务词，2 是常用应用词或组合，3 是期权、收益指标等更专门的概念。先修关系只说明有帮助的词义基础，不代表学会后自动解锁。以下依赖包含本批拟新增稳定 ID；只有目标实际创建后才可写入 YAML。

| 统计口径                 | 数量或说明                                                     |
| :----------------------- | :------------------------------------------------------------- |
| 截图数                   | 8                                                              |
| 选中并合并的源标签写法   | 95                                                             |
| 唯一学习目标             | 79                                                             |
| 拟新增词条               | 78                                                             |
| 复用既有词条             | 1                                                              |
| 拟修改既有词条           | 0                                                              |
| 额外基础子项             | 5 个，已包含在拟新增数中: profit、loss、withdraw、share、greed |
| 可选记忆字段草案         | 5 项；无词源草案                                               |
| 独立关联候选             | 11                                                             |
| 批准全部内容后的预计集合 | 1320                                                           |
| 当前音频生成授权         | 无                                                             |

这些数字不是截图全部文字的 OCR 统计。源标签写法按选中且完整辨认的不同文本去重；同一写法多次出现不重复计数。一个标签可拆出多个有用目标，一个目标也可吸收多个标签，因此源标签数不与目标数相加。额外基础子项不是额外截图出现次数；待核实与排除表也不计入已选目标。

## 候选内容

以下内容已获用户确认并写入 YAML，保留为本批具体规格。所有新增 ID 已核对全局唯一；新文件为 `entries/<id>.yml`。大多数目标使用 noun，例外在字段草案中注明。没有可靠 IPA 核验，因此本批发音字段暂留空；不凭拼写编 IPA。

### 账户与资金

| 编号 / 决策 | 稳定 ID                           | Word                    | Meaning                                                      | 来源写法 / 位置                                                                 | 优先级 / 先修                       |
| :---------- | :-------------------------------- | :---------------------- | :----------------------------------------------------------- | :------------------------------------------------------------------------------ | :---------------------------------- |
| C01 / new   | `account-finance`                 | account                 | 账户；用于记录和管理资金、资产及交易的账户。                 | Accounts；All Accounts；Account Management / [S7](#s7), [S8](#s8)               | 1 / 无                              |
| C02 / new   | `asset-finance`                   | asset                   | 资产；具有经济价值、由个人或机构持有的资源。                 | Assets；Total Assets / [S6](#s6), [S7](#s7), [S8](#s8)                          | 1 / 无                              |
| C03 / new   | `cash-finance`                    | cash                    | 现金；账户中以货币金额记录的资金。                           | Total Cash；Currency/Cash / [S6](#s6)                                           | 1 / 无                              |
| C04 / new   | `net-finance`                     | net                     | 净额的；扣除相关项目后剩余的。                               | Net Assets；Net Income / [S1](#s1), [S6](#s6), [S7](#s7)                        | 1 / 无                              |
| C05 / new   | `net-assets-finance`              | net assets              | 净资产；资产扣除负债后的净值。                               | Net Assets / [S6](#s6), [S7](#s7)                                               | 2 / net-finance, asset-finance      |
| C06 / new   | `market-value-finance`            | market value            | 市值；按市场价格衡量的资产价值。                             | Market Value；MV/QTY / [S6](#s6)                                                | 2 / asset-finance                   |
| C07 / new   | `margin-finance`                  | margin                  | 保证金；交易或融资中用于提供担保的资金或资产权益。           | Margin Account / [S6](#s6), [S7](#s7)                                           | 2 / 无                              |
| C08 / new   | `margin-account-finance`          | margin account          | 保证金账户；允许按规则以账户资产作担保向券商借款交易的账户。 | Margin Account；US Margin Account / [S6](#s6), [S7](#s7), [S8](#s8)             | 2 / account-finance, margin-finance |
| C09 / new   | `cash-account-finance`            | cash account            | 现金账户；证券买入需要由账户资金全额支付的账户。             | Cash Account / [S7](#s7)                                                        | 2 / account-finance, cash-finance   |
| C10 / new   | `position-finance`                | position                | 持仓；账户对某项金融工具持有的数量或形成的风险敞口。         | Positions；Total Position P/L / [S6](#s6), [S8](#s8)                            | 1 / 无                              |
| C11 / new   | `holdings-finance`                | holdings                | 持有的资产；投资组合中的证券等投资项目。                     | Buffett Holdings / [S3](#s3)                                                    | 2 / 无                              |
| C12 / new   | `profit-finance`                  | profit                  | 利润；收入或交易所得超过相关成本的部分。                     | 补充子项，非独立截图标签 / [S6](#s6), [S7](#s7), [S8](#s8)                      | 1 / 无                              |
| C13 / new   | `loss-finance`                    | loss                    | 亏损；成本超过所得或资产价值减少造成的损失。                 | 补充子项，非独立截图标签 / [S6](#s6), [S7](#s7), [S8](#s8)                      | 1 / 无                              |
| C14 / new   | `profit-and-loss-finance`         | profit and loss         | 盈亏；投资或经营中的盈利与亏损结果。                         | P/L Analysis；Today’s P/L；Total Position P/L / [S6](#s6), [S7](#s7), [S8](#s8) | 2 / profit-finance, loss-finance    |
| C15 / new   | `buying-power-finance`            | buying power            | 购买力；按账户和交易规则可用于新增交易的额度。               | Max Buying Power；Cash BP / [S6](#s6)                                           | 2 / cash-finance                    |
| C16 / new   | `withdraw-finance`                | withdraw                | 提取；将账户资金取出或转出。                                 | 补充子项，非独立截图标签 / [S6](#s6)                                            | 1 / 无                              |
| C17 / new   | `withdrawable-finance`            | withdrawable            | 可提取的；在当前规则下允许从账户取出的。                     | Withdrawable / [S6](#s6)                                                        | 2 / withdraw-finance                |
| C18 / new   | `deposit-finance`                 | deposit                 | 存入资金；将钱转入账户。                                     | Deposit / [S6](#s6), [S7](#s7)                                                  | 1 / 无                              |
| C19 / new   | `transfer-finance`                | transfer                | 划转；将资金从一个账户或位置转到另一个。                     | Cash Transfer / [S7](#s7)                                                       | 1 / 无                              |
| C20 / new   | `currency-finance`                | currency                | 货币；用来标价和结算的货币单位或体系。                       | Currency；Currency/Cash / [S1](#s1), [S6](#s6)                                  | 1 / 无                              |
| C21 / new   | `exchange-currency`               | exchange                | 兑换；将一种货币换成另一种。                                 | Exchange / [S7](#s7)                                                            | 2 / currency-finance                |
| C22 / new   | `transaction-finance`             | transaction             | 交易；涉及资金或资产往来的业务行为或记录。                   | Transaction / [S6](#s6)                                                         | 2 / 无                              |
| C23 / new   | `funds-on-hold-finance`           | funds on hold           | 冻结或被占用的资金；暂时不能自由使用的资金。                 | Funds on Hold / [S6](#s6)                                                       | 2 / cash-finance                    |
| C24 / new   | `cash-sweep-finance`              | cash sweep              | 现金归集；将符合条件的闲置现金转入指定现金管理或生息安排。   | Cash Sweep / [S6](#s6)                                                          | 3 / cash-finance, interest-finance  |
| C25 / new   | `interest-finance`                | interest                | 利息；因存放或借用资金产生的收益或费用。                     | Total Interest；Today’s Interest；Pending Interest / [S6](#s6)                  | 1 / 无                              |
| C26 / new   | `annual-percentage-yield-finance` | annual percentage yield | 年化收益率；考虑利息及复利频率后表示的年度收益比例。         | APY / [S6](#s6)                                                                 | 3 / interest-finance                |
| C27 / new   | `portfolio-finance`               | portfolio               | 投资组合；同一投资者或账户持有的一组投资资产。               | % Portfolio / [S6](#s6)                                                         | 2 / asset-finance                   |
| C28 / new   | `cost-finance`                    | cost                    | 成本；取得或持有某项资产所付出的代价。                       | Price/Cost / [S6](#s6)                                                          | 2 / 无                              |
| C76 / new   | `order-finance`                   | order                   | 委托单；提交给交易系统的买卖指令。                           | Orders；Order History / [S6](#s6)                                               | 1 / 无                              |

### 行情与观察

| 编号 / 决策 | 稳定 ID                         | Word                  | Meaning                                                    | 来源写法 / 位置                                                               | 优先级 / 先修                                  |
| :---------- | :------------------------------ | :-------------------- | :--------------------------------------------------------- | :---------------------------------------------------------------------------- | :--------------------------------------------- |
| C29 / new   | `watchlist-finance`             | watchlist             | 自选列表；用于持续关注一组证券或其他标的的清单。           | Watchlists / [S2](#s2), [S3](#s3), [S4](#s4), [S6](#s6), [S7](#s7), [S8](#s8) | 1 / 无                                         |
| C30 / new   | `trade-finance`                 | trade                 | 交易；买卖证券或其他金融工具的行为。                       | Trade / [S3](#s3), [S5](#s5)                                                  | 1 / 无                                         |
| C31 / new   | `stock-finance`                 | stock                 | 股票；代表公司所有权份额的证券。                           | Crypto Stocks / [S2](#s2)                                                     | 1 / 无                                         |
| C32 / new   | `crypto-finance`                | crypto                | 加密货币；对 cryptocurrency 或加密资产领域的常用简称。     | Crypto；Crypto Market Cap；Crypto Stocks / [S2](#s2), [S3](#s3)               | 2 / 无                                         |
| C33 / new   | `market-capitalization-finance` | market capitalization | 市值；按市场价格与相应数量口径计算的总体价值。             | Market Cap；Crypto Market Cap / [S2](#s2)                                     | 2 / 无                                         |
| C34 / new   | `volume-trading`                | volume                | 成交量；某一时段内完成交易的规模，可按数量或计价金额表示。 | Volume / [S2](#s2), [S5](#s5)                                                 | 2 / 无                                         |
| C35 / new   | `gainer-finance`                | gainer                | 上涨标的；在指定期间内价格上升的证券或资产。               | Top Gainers / [S2](#s2)                                                       | 2 / 无                                         |
| C36 / new   | `loser-finance`                 | loser                 | 下跌标的；在指定期间内价格下降的证券或资产。               | Top Losers / [S2](#s2)                                                        | 2 / 无                                         |
| C37 / new   | `heat-map-data`                 | heat map              | 热力图；用颜色或色块表达数值差异的数据图。                 | Heat Map / [S2](#s2)                                                          | 2 / 无                                         |
| C38 / new   | `index-finance`                 | index                 | 指数；用于概括一组证券或某种市场指标变化的数值。           | Index；Fear & Greed Index / [S2](#s2), [S8](#s8)                              | 2 / 无                                         |
| C39 / new   | `return-finance`                | return                | 投资回报；投入资金产生的收益结果，可用金额或比例表示。     | 1Y Return / [S3](#s3)                                                         | 2 / 无                                         |
| C40 / new   | `pre-market-finance`            | pre-market            | 盘前的；发生在常规交易时段开始之前的。                     | Pre-Mkt / [S1](#s1), [S5](#s5)                                                | 2 / 无                                         |
| C41 / new   | `futures-finance`               | futures               | 期货；按标准化合约约定未来交易条件的交易所交易工具。       | Futures / [S5](#s5)                                                           | 3 / 无                                         |
| C42 / new   | `discount-finance`              | discount              | 折价；价格低于某个比较基准的情况。                         | At Discount / [S5](#s5)                                                       | 2 / 无                                         |
| C43 / new   | `volatility-finance`            | volatility            | 波动性；价格或收益在一段时间内变化的程度。                 | Earnings Move & Volatility / [S1](#s1)                                        | 3 / 无                                         |
| C44 / new   | `expected-move-finance`         | expected move         | 预期波幅；对某一时期内价格可能变动幅度的估计。             | Latest expected move / [S1](#s1)                                              | 3 / estimate-general, volatility-finance       |
| C75 / new   | `exchange-traded-fund-finance`  | exchange-traded fund  | 交易所交易基金；份额可以在交易所按市场价格买卖的基金。     | ETFs / [S1](#s1)                                                              | 2 / share-finance                              |
| C77 / new   | `fear-general`                  | fear                  | 恐惧；害怕或担忧的情绪。                                   | Fear；Fear & Greed Index / [S2](#s2)                                          | 1 / 无                                         |
| C78 / new   | `greed-general`                 | greed                 | 贪婪；对金钱、利益等抱有过度欲求。                         | 补充子项，非独立截图标签 / [S2](#s2)                                          | 2 / 无                                         |
| C79 / new   | `fear-and-greed-index-finance`  | fear and greed index  | 恐惧与贪婪指数；用于概括市场偏恐惧或偏贪婪情绪的指标。     | Fear & Greed Index / [S2](#s2)                                                | 2 / fear-general, greed-general, index-finance |

### 财报与预测

| 编号 / 决策 | 稳定 ID                      | Word               | Meaning                                              | 来源写法 / 位置                                                                | 优先级 / 先修                          |
| :---------- | :--------------------------- | :----------------- | :--------------------------------------------------- | :----------------------------------------------------------------------------- | :------------------------------------- |
| C45 / new   | `financials-finance`         | financials         | 财务数据；财务报表等公司财务信息。                   | Financials (USD) / [S1](#s1)                                                   | 2 / 无                                 |
| C46 / new   | `earnings-finance`           | earnings           | 盈利；公司在一定期间取得的利润。                     | Q2 FY2027 Earnings；Earnings Hub；Earnings Move & Volatility / [S1](#s1)       | 2 / 无                                 |
| C47 / new   | `revenue-finance`            | revenue            | 营业收入；公司销售商品或提供服务取得的收入。         | Revenue / [S1](#s1)                                                            | 1 / 无                                 |
| C48 / new   | `income-finance`             | income             | 收入；所得，是否扣除费用要结合限定语判断。           | Net Income / [S1](#s1)                                                         | 1 / 无                                 |
| C49 / new   | `net-income-finance`         | net income         | 净利润；扣除相关成本、费用和税项后的利润。           | Net Income / [S1](#s1)                                                         | 2 / net-finance, income-finance        |
| C50 / new   | `estimate-general`           | estimate           | 估计值；在结果确定前依据可用信息作出的预测或近似数。 | Estimate (YoY)；Vs Estimate (% Chg)；Financial Estimates；Estimate / [S1](#s1) | 1 / 无                                 |
| C51 / new   | `actual-general`             | actual             | 实际的；实际发生或观察到的。                         | Actual (YoY)；Actual / [S1](#s1)                                               | 1 / 无                                 |
| C52 / new   | `beat-estimates-finance`     | beat estimates     | 超过预期；实际业绩高于先前预测。                     | Beat Est. / [S1](#s1)                                                          | 2 / estimate-general                   |
| C53 / new   | `year-over-year-finance`     | year-over-year     | 同比的；与上一年相同期间相比的。                     | YoY / [S1](#s1)                                                                | 2 / 无                                 |
| C54 / new   | `fiscal-year-finance`        | fiscal year        | 财年；公司或机构用于财务报告的年度。                 | FY2027 / [S1](#s1)                                                             | 2 / 无                                 |
| C55 / new   | `quarter-finance`            | quarter            | 季度；一年中的四分之一，在财报中表示一个报告季度。   | Q2；2026/Q4；2027/Q1；2027/Q2；2027/Q3；2027/Q4 / [S1](#s1)                    | 1 / 无                                 |
| C56 / new   | `earnings-per-share-finance` | earnings per share | 每股收益；将相应利润按每股口径表示的财务指标。       | EPS / [S1](#s1)                                                                | 3 / earnings-finance, share-finance    |
| C57 / new   | `ebit-finance`               | EBIT               | 息税前利润；未扣除利息和税项的利润指标。             | EBIT / [S1](#s1)                                                               | 3 / earnings-finance, interest-finance |
| C58 / new   | `share-finance`              | share              | 股份；公司资本或某项权益中的一份。                   | 补充子项，非独立截图标签 / [S1](#s1)                                           | 1 / 无                                 |
| C59 / new   | `shareholder-finance`        | shareholder        | 股东；持有公司股份的人或机构。                       | Shareholders / [S1](#s1)                                                       | 2 / share-finance                      |
| C60 / new   | `dividend-finance`           | dividend           | 股息；公司向股东分配的收益。                         | Dividends / [S1](#s1)                                                          | 2 / shareholder-finance                |

### 期权与报价

| 编号 / 决策 | 稳定 ID                 | Word          | Meaning                                                          | 来源写法 / 位置                               | 优先级 / 先修                            |
| :---------- | :---------------------- | :------------ | :--------------------------------------------------------------- | :-------------------------------------------- | :--------------------------------------- |
| C61 / new   | `option-finance`        | option        | 期权；给予买方按合约条件买入或卖出标的权利的合约。               | Options；Single Option / [S1](#s1), [S5](#s5) | 2 / 无                                   |
| C62 / new   | `call-option-finance`   | call option   | 看涨期权；在股票期权语境中给予买方按合约价格买入标的股票的权利。 | Call；CALL / [S5](#s5), [S8](#s8)             | 3 / option-finance                       |
| C63 / new   | `put-option-finance`    | put option    | 看跌期权；在股票期权语境中给予买方按合约价格卖出标的股票的权利。 | Put；PUT / [S5](#s5), [S8](#s8)               | 3 / option-finance                       |
| C64 / new   | `option-chain-finance`  | option chain  | 期权链；按到期日、行权价等条件组织的期权合约报价列表。           | Chain / [S5](#s5)                             | 3 / option-finance, strike-price-finance |
| C65 / new   | `strike-price-finance`  | strike price  | 行权价；期权合约约定买卖标的时使用的价格。                       | Strike / [S5](#s5)                            | 3 / option-finance                       |
| C66 / new   | `bid-price-finance`     | bid price     | 买方报价；行情中买方愿意支付的买入报价。                         | Bid / [S5](#s5)                               | 2 / 无                                   |
| C67 / new   | `ask-price-finance`     | ask price     | 卖方报价；行情中卖方愿意接受的卖出报价。                         | Ask / [S5](#s5)                               | 2 / 无                                   |
| C68 / new   | `closing-price-finance` | closing price | 收盘价；某一交易时段结束时采用的价格。                           | Closing Price / [S5](#s5)                     | 2 / 无                                   |
| C69 / new   | `symbol-market`         | symbol        | 证券或交易标的代码；用于标识某个交易对象的字符标记。             | Symbol / [S6](#s6)                            | 1 / 无                                   |
| C70 / new   | `quantity-general`      | quantity      | 数量；按指定单位计数或度量的多少。                               | MV/QTY / [S6](#s6)                            | 1 / 无                                   |

### 通用界面词

| 编号 / 决策 | 稳定 ID             | Word      | Meaning                                            | 来源写法 / 位置                 | 优先级 / 先修 |
| :---------- | :------------------ | :-------- | :------------------------------------------------- | :------------------------------ | :------------ |
| C71 / new   | `discover-general`  | discover  | 发现；找到或了解到原先未知的内容。                 | Discover / [S3](#s3), [S4](#s4) | 1 / 无        |
| C72 / new   | `community-general` | community | 社区；具有共同兴趣、活动或联系的人群及其交流空间。 | Community / [S4](#s4)           | 1 / 无        |
| C73 / new   | `wealth-general`    | wealth    | 财富；拥有的金钱、资产和其他有价值资源。           | Wealth / [S4](#s4)              | 2 / 无        |

### 复用

| 编号 / 决策 | 稳定 ID                                     | Word    | Meaning            | 来源写法 / 位置              | 优先级 / 先修 |
| :---------- | :------------------------------------------ | :------ | :----------------- | :--------------------------- | :------------ |
| C74 / reuse | [pending-tech](../entries/pending-tech.yml) | pending | 待处理的；待完成的 | Pending Interest / [S6](#s6) | 1 / 无        |

### 例句与 Notes 草案

下表的例句均为原创语言用例，不是对用户账户、某家公司或当前市场的事实断言。Notes 中的说明为拟写内容。实施时，对新增词条追加来源行 `Source: 用户提供的 Moomoo 界面截图，2026-11-06，<S 编号与原标签>。`；每条实际使用的参考资料追加 `Reference: <参考来源完整标题>. <URL>`，按 URL 去重。参考链接支持一般术语，不证明页面专属计算公式。截图不支持精确财务机制时，明确保留限制而不是补公式。

| 编号 / Word                   | 词性        | Example                                                    | 拟写 Notes / 限制                                                                                                                                                  | 参考                     |
| :---------------------------- | :---------- | :--------------------------------------------------------- | :----------------------------------------------------------------------------------------------------------------------------------------------------------------- | :----------------------- |
| C01 / account                 | noun        | Open the account details.                                  | 除截图来源与下列参考资料外，不增加重复解释。                                                                                                                       | [F03](#f03)              |
| C02 / asset                   | noun        | The account holds cash and other assets.                   | Assets 通常用复数。截图 Total Assets 是平台汇总标签，不根据一个标签补写总资产或净资产计算公式。                                                                    | [F01](#f01)              |
| C03 / cash                    | noun        | The cash balance changed after the trade.                  | 账户现金不等于纸币，也不必等于可提现金额或证券市值。                                                                                                               | [F03](#f03)              |
| C04 / net                     | adjective   | The report shows the net amount.                           | 扣除哪些项目由术语决定；net assets 与 net income 不是同一个指标，也不是网络的 net。                                                                                | [F01](#f01)              |
| C05 / net assets              | noun        | Check the account’s net assets.                            | 不等于现金余额、持仓市值或可提现金额。平台币种换算与具体估值口径需看字段说明。                                                                                     | [F01](#f01), [F03](#f03) |
| C06 / market value            | noun        | The market value of the position changed.                  | 账户 Market Value 与某家公司或整个加密市场的 Market Cap 要结合计算对象区分；MV 是 market value 的缩写。                                                            | [F03](#f03)              |
| C07 / margin                  | noun        | The broker sets the margin requirement.                    | 此处为担保与融资语境，不是财报中的利润率 margin。                                                                                                                  | [F05](#f05)              |
| C08 / margin account          | noun        | A margin account can involve borrowing.                    | 账户具备融资能力不代表每一笔交易都已经借款，不从截图推断当前负债。                                                                                                 | [F05](#f05), [F12](#f12) |
| C09 / cash account            | noun        | The investor uses a cash account.                          | 一般现金账户不能向券商借钱支付证券买入；结算与资金可用规则另看账户说明。                                                                                           | [F05](#f05)              |
| C10 / position                | noun        | The account has a position in the ETF.                     | 此处不是元素位置或代码位置。不要由持仓一词推断一定是多头。                                                                                                         | [F02](#f02), [F03](#f03) |
| C11 / holdings                | noun        | The fund publishes its holdings.                           | 保留自然复数 holdings。截图中 Buffett 是栏目或组合名称的一部分，不据此确认其真实持仓或披露时点。                                                                   | [F09](#f09)              |
| C12 / profit                  | noun        | The trade produced a profit.                               | 为理解 P/L = profit and loss 而补充的基础子项，截图未把 profit 全拼出来。                                                                                          | [F01](#f01)              |
| C13 / loss                    | noun        | The position shows a loss.                                 | 为理解 P/L = profit and loss 而补充的基础子项，截图未把 loss 全拼出来。                                                                                            | [F01](#f01)              |
| C14 / profit and loss         | noun        | The app shows today’s profit and loss.                     | P/L 是常见缩写，也可见 P&L。保留 Today’s P/L 与 Total Position P/L 作为标签用例，不为每个组合单独建卡；具体期间、费用和已实现/未实现口径不能仅凭截图确定。         | [F03](#f03)              |
| C15 / buying power            | noun        | Buying power is not the same as withdrawable cash.         | BP = buying power。Max Buying Power 与 Cash BP 是不同标签；不把最大购买力当作自有现金，也不从截图规定固定杠杆倍数或实际可买数量。                                  | [F03](#f03), [F05](#f05) |
| C16 / withdraw                | verb        | Withdraw funds from the account.                           | 从 Withdrawable 提出的有独立学习价值的动词子项；截图并未出现独立 Withdraw 按钮。                                                                                   | [F03](#f03)              |
| C17 / withdrawable            | adjective   | Only part of the cash is withdrawable.                     | 可提现额度不必等于现金余额、净资产或购买力；不补写到账时间或结算期限。                                                                                             | [F03](#f03)              |
| C18 / deposit                 | verb        | The customer deposits funds into the account.              | 此条以动词学习；界面也可能把 Deposit 用作入金功能名称。                                                                                                            | [F12](#f12)              |
| C19 / transfer                | verb        | Transfer cash between the accounts.                        | Cash Transfer 作为界面搭配保留，不为每个现金划转组合另外建卡。此处不是文件传输。                                                                                   | [F03](#f03)              |
| C20 / currency                | noun        | Choose the display currency.                               | USD 是美元代码，放在用例中，不新增 USD 专名词条。                                                                                                                  | [F03](#f03)              |
| C21 / exchange                | verb        | Exchange one currency for another.                         | 这里对应账户中的换汇操作；不是交易所这一名词义。                                                                                                                   | [F12](#f12)              |
| C22 / transaction             | noun        | The transaction appears in the account history.            | 不把委托单 order 一律当作已经完成的交易；此条也不是数据库事务。                                                                                                    | [F03](#f03)              |
| C23 / funds on hold           | noun        | Some funds are on hold.                                    | on hold 在这里表示被暂时保留或限制；不是持仓 holdings。具体原因可能来自订单或其他账户限制，不从截图猜测。                                                          | [F03](#f03)              |
| C24 / cash sweep              | noun        | The account offers a cash sweep program.                   | Moomoo 历史客户协议将该计划归为闲置现金管理安排；当前参与银行、归集时间、利率和保障范围未在本次核实，不写入词条。                                                  | [F12](#f12)              |
| C25 / interest                | noun        | The account earns interest on eligible cash.               | 此处不是兴趣。Total、Today’s、Pending 是界面限定语，不另外制造三张利息组合卡；Pending Interest 的具体计提与入账口径仍待平台说明。                                  | [F07](#f07), [F12](#f12) |
| C26 / annual percentage yield | noun        | Compare the annual percentage yields.                      | APY 是全称缩写，不等于某天收到的利息，也不是对未来收益的保证。不保留截图宣传数字为有效利率。                                                                       | [F07](#f07)              |
| C27 / portfolio               | noun        | The portfolio contains several positions.                  | 此处不是设计作品集；% Portfolio 表示持仓在组合中的占比标签，具体分母与估值口径需看平台说明。                                                                       | [F09](#f09)              |
| C28 / cost                    | noun        | The market price is above the recorded cost.               | Price/Cost 中两者是不同数据。未核实该账户选择的是平均成本、摊薄成本还是其他口径，不擅自规范为 cost basis 或填写费用公式。                                          | [F01](#f01)              |
| C29 / watchlist               | noun        | Add the ETF to the watchlist.                              | 关注列表不等于已持有这些资产。                                                                                                                                     | 截图语境与编辑性语言草案 |
| C30 / trade                   | noun        | Review the details of the trade.                           | 此条以名词学习；界面 Trade 也作为交易功能入口。                                                                                                                    | [F02](#f02)              |
| C31 / stock                   | noun        | The portfolio includes stocks and ETFs.                    | Crypto Stocks 是主题分类，不是每一只股票都等同于加密货币。                                                                                                         | [F09](#f09), [F02](#f02) |
| C32 / crypto                  | noun        | The page shows crypto market data.                         | 保留界面常见简称 crypto；Notes 记录 cryptocurrency 全称，不为二者建重复条目。此处不学习加密算法这一技术义。                                                        | [F10](#f10)              |
| C33 / market capitalization   | noun        | The assets are ranked by market capitalization.            | Market Cap 是缩写。股票常按流通在外股数计算，加密资产需区分流通量与完全稀释口径。截图总量范围未核实，不把市值当成真实流入资金。                                    | [F10](#f10)              |
| C34 / volume                  | noun        | Trading volume increased during the session.               | 单位随标的而变。加密市场总览的 Volume 显示美元量纲，表示按金额折算的交易规模，不能把所有 Volume 栏都译成股数。                                                     | [F04](#f04)              |
| C35 / gainer                  | noun        | The stock is one of today’s top gainers.                   | Top Gainers 是涨幅领先的标的列表，不等于每名持有者都已实现盈利。                                                                                                   | 截图语境与编辑性语言草案 |
| C36 / loser                   | noun        | The list shows the market’s biggest losers.                | 行情榜单中的 loser 指下跌标的，不是对人的贬称，也不直接给出投资者实际亏损。                                                                                        | 截图语境与编辑性语言草案 |
| C37 / heat map                | noun        | The heat map highlights price changes.                     | 不把红绿颜色与涨跌方向写成跨平台固定规则。                                                                                                                         | 截图语境与编辑性语言草案 |
| C38 / index                   | noun        | The index tracks a selected group of stocks.               | 股票指数与情绪指数属于不同用途，不给所有指数套同一算法。此处不是数组索引。                                                                                         | 截图语境与编辑性语言草案 |
| C39 / return                  | noun        | The chart shows the portfolio’s one-year return.           | 1Y 是 one year 的界面缩写。这里不是函数返回值，也不由历史回报推断以后表现。                                                                                        | [F05](#f05)              |
| C40 / pre-market              | adjective   | Pre-market prices can change.                              | Mkt 缩写 market。交易时段因市场、标的及服务而异，本条不填写具体开盘时间。                                                                                          | [F04](#f04)              |
| C41 / futures                 | noun        | The screen includes futures quotes.                        | 保留金融词常用复数 futures，不按普通 future 的将来义处理。不判断 NVDA Futures 的具体交易所、结算方式或合约规格。                                                   | [F11](#f11)              |
| C42 / discount                | noun        | The instrument trades at a discount.                       | 参考基准随产品而异。截图未展示此项计算说明，不解释其数值来自哪个报价差，也不等同于期权权利金。                                                                     | [F09](#f09)              |
| C43 / volatility              | noun        | Volatility can increase around an earnings release.        | 波动程度不等于涨跌方向；截图未说明该处采用历史波动率还是隐含波动率，不将二者混写。                                                                                 | [F02](#f02)              |
| C44 / expected move           | noun        | The expected move is an estimate, not a guarantee.         | 截图显示相对幅度，但没有时间窗口、计算方法或置信口径。这里只收录语言与概念，不补写基于期权价格的公式或保证区间。                                                   | 截图语境与编辑性语言草案 |
| C45 / financials              | noun        | The company’s financials include revenue data.             | 保留名词性复数 financials，不机械改成形容词 financial。具体栏目包含哪些报表由平台决定。                                                                            | [F01](#f01)              |
| C46 / earnings                | noun        | The company will report quarterly earnings.                | 财报页面的 Earnings 也可作为业绩发布栏目名称；不把 earnings 与 revenue 视为同义词。                                                                                | [F01](#f01)              |
| C47 / revenue                 | noun        | Revenue increased from the previous year.                  | 还没有扣除相关成本与费用，不等于净利润。                                                                                                                           | [F01](#f01)              |
| C48 / income                  | noun        | The report separates income from expenses.                 | 作为 net income 的有用基础词学习，不把 income 在所有财务语境中都机械等同 revenue。                                                                                 | [F01](#f01)              |
| C49 / net income              | noun        | Net income can change even when revenue is stable.         | 不是营业收入，也不等于账户现金的增加。                                                                                                                             | [F01](#f01)              |
| C50 / estimate                | noun        | The actual result differs from the estimate.               | 本条以名词学习。财报中常指分析师预期；截图没有展示汇总方法，不认定必然是某种一致预期算法。                                                                         | 截图语境与编辑性语言草案 |
| C51 / actual                  | adjective   | Compare the actual revenue with the estimate.              | 图例 Actual 可省略被修饰的 result/value。已有 actual-behavior-workflow 是实际行为这一短语，不代替这个形容词目标。                                                  | 截图语境与编辑性语言草案 |
| C52 / beat estimates          | verb phrase | The company’s revenue beat estimates.                      | Est. 在该财报标记中对应 estimates。比较的是指定指标，不表示所有财报指标都超过预期，也不保证股价上涨。                                                              | 截图语境与编辑性语言草案 |
| C53 / year-over-year          | adjective   | Revenue grew on a year-over-year basis.                    | YoY = year-over-year，不是与上一季度相比。保留缩写作为 Notes 中的界面形式。                                                                                        | 截图语境与编辑性语言草案 |
| C54 / fiscal year             | noun        | The fiscal year does not have to match the calendar year.  | FY = fiscal year。不要把财年编号直接当成自然年日期范围；本条不补 NVIDIA 财年的起止日期。                                                                           | 截图语境与编辑性语言草案 |
| C55 / quarter                 | noun        | The company reports results each quarter.                  | Q 是 quarter 的标签缩写。Q2 FY2027 表示该财年的第二季度，不自动等于自然年第二季度。                                                                                | 截图语境与编辑性语言草案 |
| C56 / earnings per share      | noun        | The report includes earnings per share.                    | EPS = earnings per share。基本与稀释每股收益的分母和调整可能不同；截图未说明口径，不写简单公式冒充精确会计规则。                                                   | [F01](#f01)              |
| C57 / EBIT                    | noun        | EBIT and net income measure different things.              | 全称 earnings before interest and taxes。不能直接等同所有口径的 operating income，也不是每股指标；不新增 tax 词条只为补依赖。                                      | [F08](#f08)              |
| C58 / share                   | noun        | A shareholder can own more than one share.                 | 为 Shareholders 和 EPS 全称理解提出的基础子项；截图未显示独立 Share 标签。不是分享这一动词义。                                                                     | [F01](#f01)              |
| C59 / shareholder             | noun        | Shareholders own shares in the company.                    | 除截图来源与下列参考资料外，不增加重复解释。                                                                                                                       | [F01](#f01)              |
| C60 / dividend                | noun        | The company announced a dividend.                          | 是否分配、金额和时间需看具体决定；不能把所有盈利自动当作会分给股东的股息。                                                                                         | [F01](#f01)              |
| C61 / option                  | noun        | The option has a specified strike price.                   | 此处不是程序选项或偏好设置。Single Option 留作组合用例，不另建一张卡。                                                                                             | [F02](#f02)              |
| C62 / call option             | noun        | A call option gives its buyer a right to buy.              | Call 是期权类型缩写，不是打电话或函数调用。不能把 call 等同于一定看涨的完整持仓策略，卖出 call 的敞口不同。                                                        | [F02](#f02)              |
| C63 / put option              | noun        | A put option gives its buyer a right to sell.              | Put 在这里不是普通放置动作。买入与卖出 put 的风险方向不同，不能只凭名称推断账户整体看法。                                                                          | [F02](#f02)              |
| C64 / option chain            | noun        | The option chain lists calls and puts.                     | 截图位于 Options > Chain；完整术语不是原型链，也不是区块链。到期日来自表上方的日期标签，不推导具体合约编码规则。                                                   | [F02](#f02), [F04](#f04) |
| C65 / strike price            | noun        | The strike price differs from the option’s market price.   | Strike 是列标题的简写；不是期权自身的交易价格。                                                                                                                    | [F02](#f02)              |
| C66 / bid price               | noun        | The bid price is lower than the ask price in this example. | 报价表常显示最高买价。与卖方报价 ask 区分，不把 bid 自动当作最近成交价。                                                                                           | [F06](#f06)              |
| C67 / ask price               | noun        | The ask price is shown beside the bid price.               | 报价表常显示最低卖价；不是普通提问的 ask。报价可能变化，不保证任何订单按显示价成交。                                                                               | [F06](#f06)              |
| C68 / closing price           | noun        | The chart marks the previous closing price.                | 不等于盘前报价；截图未说明该项是否采用常规时段或特殊报价规则，不据此补写收盘计算算法。                                                                             | [F04](#f04)              |
| C69 / symbol                  | noun        | The table sorts positions by symbol.                       | 这里不是 JavaScript Symbol 原始类型。具体符号格式由市场或平台规定，不为 NVDA、BTC 等每个代码建卡。                                                                 | [F04](#f04)              |
| C70 / quantity                | noun        | Check the quantity and the market value.                   | QTY 是 quantity 的界面缩写。持仓数量的单位可能是股、份额或合约，不能统一认定为股票股数。                                                                           | 截图语境与编辑性语言草案 |
| C71 / discover                | verb        | Discover new features in the app.                          | 在本界面作为内容发现入口，不据标签推断具体推荐算法。                                                                                                               | 截图语境与编辑性语言草案 |
| C72 / community               | noun        | The community discusses market news.                       | 是用户交流社区，不是公司这一含义。                                                                                                                                 | 截图语境与编辑性语言草案 |
| C73 / wealth                  | noun        | Wealth can include more than cash.                         | 此处是栏目名称所用普通含义，不推断入口内有哪些具体理财产品。                                                                                                       | 截图语境与编辑性语言草案 |
| C74 / pending                 | noun        | The framework tests are still pending.                     | 复用已有待处理、待完成这一核心含义，保留当前技术例句和录音；不因截图出现就新增 pending-finance。利息究竟待计提、待结算还是待入账需看平台详情，未确认前不改原词条。 | 截图语境与编辑性语言草案 |
| C75 / exchange-traded fund    | noun        | ETF shares trade on an exchange.                           | ETF 是 exchange-traded fund 的缩写；并非所有 ETF 都是被动指数基金。基金份额与单一公司股票的权益不同。                                                              | [F09](#f09)              |
| C76 / order                   | noun        | An order may remain unfilled.                              | 委托不等于成交；Order History 作为用例保留。不能复用编程中的顺序 order。                                                                                           | [F15](#f15)              |
| C77 / fear                    | noun        | Fear can affect an investor’s decisions.                   | 此处学习普通情绪名词，截图中 Fear 是指标的情绪标签，不推断某个人的心理状态。                                                                                       | [F13](#f13)              |
| C78 / greed                   | noun        | Greed can influence decisions.                             | 从 Fear & Greed Index 提出的有用基础子项，不声称有独立 Greed 按钮。指标中的情绪称呼不等于对每位投资者的道德判断。                                                  | [F13](#f13)              |
| C79 / fear and greed index    | noun        | The fear and greed index summarizes market sentiment.      | 只收录指标名称与一般用途。提供方算法、权重、分数阈值和统计范围未核实，不写进字段，也不把某个读数当成未来涨跌的保证。                                               | [F13](#f13)              |

复用项 pending-tech 的例句在上表中是现有例句展示，不授权修改。Pending Interest 只用于证明本批复用判断；利息状态的精确定义未确认，不回写该短语。所有拟写 Context 默认省略: Meaning 和 Notes 已能表达当前学习目标，不为填满字段制造重复句子。

拟用标签: 金融目标使用 `finance` 和相应词性标签；含空格的固定术语加 `phrase`。discover、community、wealth 使用 `general`、`ui` 和相应词性标签；fear、greed、estimate、actual、quantity、heat map 可使用 `general` 加词性标签，并按截图语境选用 `finance`。词性标签用 `noun`、`verb`、`adjective` 等现有格式；不新增大量近义领域标签。复用词条的标签保持原样。

## 缩写与组合处理

| 截图写法                                                              | 规范化或分组                                                                   | 边界                                                                                 |
| :-------------------------------------------------------------------- | :----------------------------------------------------------------------------- | :----------------------------------------------------------------------------------- |
| Accounts / Assets / Positions / Watchlists / Shareholders / Dividends | 按通常词头学习 account / asset / position / watchlist / shareholder / dividend | 保留截图复数在 Notes；holdings、earnings、financials、futures 保留有意义的自然复数。 |
| Max Buying Power / Cash BP                                            | 统一学习 buying power，Notes 对比两个标签；BP = buying power                   | 不制造每个修饰词组合的卡，也不把购买力当成现金。                                     |
| P/L Analysis / Today’s P/L / Total Position P/L                       | profit and loss；另提 profit、loss 两个基础子项                                | 标签是源材料，两个全拼基础词属于补充，不声称截图写出了全称。                         |
| Price/Cost / MV/QTY / % Portfolio                                     | 分别理解 cost、market value、quantity、portfolio；QTY = quantity               | 不由缩写推断统一单位、成本计算方法或组合占比分母。                                   |
| Cash Transfer                                                         | transfer 词条中保留 cash transfer 搭配                                         | 账户划转义，不同于文件传输。                                                         |
| Total / Today’s / Pending Interest                                    | 统一到 interest；pending 复用 pending-tech                                     | 不把三种标签拆成三张利息卡。Pending Interest 不直接译为已经可提取的利息。            |
| Estimate / Financial Estimates / Vs Estimate / Beat Est.              | estimate 与固定财报表达 beat estimates                                         | Beat Est. 的完整写法保留在来源，不把超过预期推成股价结论。                           |
| YoY / FY / Q2 / 1Y                                                    | year-over-year / fiscal year / quarter；1Y 放入 return 的 Notes                | YoY 不等于环比；财年不必等于自然年；不为每个年度编号建卡。                           |
| EPS / EBIT / ETFs                                                     | earnings per share / EBIT / exchange-traded fund                               | EBIT 保留缩写作 Word、全称放 Notes；ETF 与 EPS 在 Notes 保留缩写。                   |
| Pre-Mkt                                                               | pre-market                                                                     | 不补交易时段起止时间。                                                               |
| Call / Put / Chain / Strike                                           | call option / put option / option chain / strike price                         | 基于 Options 页补足完整术语，原短标签保留；与编程术语分开。                          |
| Crypto / Crypto Stocks                                                | crypto 及 stock；crypto 的 Notes 记录 cryptocurrency                           | Crypto Stocks 是主题分类，属于组合用例，不新建任意主题组合。                         |
| Buffett Holdings                                                      | holdings，保留此搭配作来源形式                                                 | 不新增 Buffett 专名，不确认栏目代表哪一披露期的真实持仓。                            |
| Watchlist / Shareholder / Withdrawable                                | 可用构成说明见 W01–W03                                                         | withdraw、share 为有独立用途的额外基础子项；不把每个词缀都建卡。                     |

## 可选记忆辅助

以下 5 项具体文案已于 2026-11-06 单独确认，属于批准的构成或辨义说明；没有提出任何词源。历史来源无法从界面拼写推断。这组文案与关联分别取得批准，现已按原文写入词条。

### W01 — watchlist

目标: `watchlist-finance`；字段: `word_formation`；决策: 新增词条时可选添加。

```plain
watch + list: 把需要持续关注的项目列成清单。自选列表不表示已经买入。
```

依据与限制: Screenshot S8；现代可见复合结构，不是命名史。

### W02 — shareholder

目标: `shareholder-finance`；字段: `word_formation`；决策: 新增词条时可选添加。

```plain
share (股份) + holder (持有者): 持有股份的人或机构。
```

依据与限制: F01 支持股份与股东关系；现代构成分析，不额外建立 holder 条目。

### W03 — withdrawable

目标: `withdrawable-finance`；字段: `word_formation`；决策: 新增词条时可选添加。

```plain
withdraw + -able: 可以提取的。这里修饰资金，表示允许取出的金额，而不是账户所有现金。
```

依据与限制: S6/F03 支持界面用法；现代后缀分析，不是词源说明。

### W04 — pre-market

目标: `pre-market-finance`；字段: `word_formation`；决策: 新增词条时可选添加。

```plain
pre- (在……之前) + market: 在本条中指常规市场交易时段开始之前。
```

依据与限制: S5/F04 支持盘前语境；不把所有 market 的含义都解释成交易时段。

### N01 — bid price

目标: `bid-price-finance`；字段: `mnemonic`；决策: 新增词条时可选添加。

```plain
辨义助记: bid 先想买方出价，ask 先想卖方要价；两者是报价的两侧，不等于最近成交价。
```

依据与限制: F06 支持报价方向；这段提示为原创，不是原文引语。

## 关联候选

R01–R11 全部关系已于 2026-11-06 单独确认，并随相应新增词条写入。表中新增端的 Meaning 是本计划拟写值；R11 的 fearless-general 展示当前 Word 和 Meaning，保留全部原字段。R11 只在新 fear-general 端声明，现有 fearless 只会得到派生反向展示。每一对仅在 A 端声明 B，保留后续可能存在的列表成员；导出器再生成双向直接邻居。批准关系不授权额外助记、先修、示例或词义改写。

| 编号 | A: ID / Word / Meaning                                                                                   | B: ID / Word / Meaning                                                                                         | 关联价值                                                               | 混淆风险 / 声明方向                                                                                                                              |
| :--- | :------------------------------------------------------------------------------------------------------- | :------------------------------------------------------------------------------------------------------------- | :--------------------------------------------------------------------- | :----------------------------------------------------------------------------------------------------------------------------------------------- |
| R01  | `margin-account-finance` / margin account / 保证金账户；允许按规则以账户资产作担保向券商借款交易的账户。 | `cash-account-finance` / cash account / 现金账户；证券买入需要由账户资金全额支付的账户。                       | 对比是否允许以账户资产支持券商借款。                                   | 不要把保证金账户理解成每笔交易都已使用融资。 声明: margin-account-finance → cash-account-finance                                                 |
| R02  | `net-assets-finance` / net assets / 净资产；资产扣除负债后的净值。                                       | `market-value-finance` / market value / 市值；按市场价格衡量的资产价值。                                       | 对比扣除负债后的账户净值与按市场价格衡量的资产价值。                   | 两个字段对象和口径不同，不能互相当作现金或可提现额。 声明: net-assets-finance → market-value-finance                                             |
| R03  | `revenue-finance` / revenue / 营业收入；公司销售商品或提供服务取得的收入。                               | `net-income-finance` / net income / 净利润；扣除相关成本、费用和税项后的利润。                                 | 区别营业收入与扣除相关成本费用后的净利润。                             | 收入增长不保证净利润同时增长。 声明: revenue-finance → net-income-finance                                                                        |
| R04  | `bid-price-finance` / bid price / 买方报价；行情中买方愿意支付的买入报价。                               | `ask-price-finance` / ask price / 卖方报价；行情中卖方愿意接受的卖出报价。                                     | 对比买方与卖方报价方向。                                               | 不是开盘价/收盘价，也不是保证成交价格。 声明: bid-price-finance → ask-price-finance                                                              |
| R05  | `call-option-finance` / call option / 看涨期权；在股票期权语境中给予买方按合约价格买入标的股票的权利。   | `put-option-finance` / put option / 看跌期权；在股票期权语境中给予买方按合约价格卖出标的股票的权利。           | 对比股票期权买方买入标的与卖出标的的权利。                             | 名称不代表账户整体方向，不能忽略买卖方身份。 声明: call-option-finance → put-option-finance                                                      |
| R06  | `profit-finance` / profit / 利润；收入或交易所得超过相关成本的部分。                                     | `loss-finance` / loss / 亏损；成本超过所得或资产价值减少造成的损失。                                           | 盈亏的正负两侧，帮助理解 P/L 缩写。                                    | 不能仅凭上涨或下跌榜就推断账户已经实现盈亏。 声明: profit-finance → loss-finance                                                                 |
| R07  | `buying-power-finance` / buying power / 购买力；按账户和交易规则可用于新增交易的额度。                   | `withdrawable-finance` / withdrawable / 可提取的；在当前规则下允许从账户取出的。                               | 区分交易额度与可提取资金的可用范围。                                   | 词性不同，也不是简单相等或互为反义；融资和结算会影响范围。 声明: buying-power-finance → withdrawable-finance                                     |
| R08  | `holdings-finance` / holdings / 持有的资产；投资组合中的证券等投资项目。                                 | `position-finance` / position / 持仓；账户对某项金融工具持有的数量或形成的风险敞口。                           | 联系组合持有资产与某个标的的持仓。                                     | 不同界面用词可能重叠，不把两者编成一套严格通用分类。 声明: holdings-finance → position-finance                                                   |
| R09  | `option-finance` / option / 期权；给予买方按合约条件买入或卖出标的权利的合约。                           | `futures-finance` / futures / 期货；按标准化合约约定未来交易条件的交易所交易工具。                             | 同属衍生工具，比较期权买方的权利与期货合约约定义务。                   | 不推断具体结算方式、保证金或某一策略的风险大小。 声明: option-finance → futures-finance                                                          |
| R10  | `stock-finance` / stock / 股票；代表公司所有权份额的证券。                                               | `exchange-traded-fund-finance` / exchange-traded fund / 交易所交易基金；份额可以在交易所按市场价格买卖的基金。 | 对比公司股票与可以在交易所买卖的基金份额。                             | ETF 不是单一公司股票，也不必是被动指数基金。 声明: stock-finance → exchange-traded-fund-finance                                                  |
| R11  | `fear-general` / fear / 恐惧；害怕或担忧的情绪。                                                         | [fearless-general](../entries/fearless-general.yml) / fearless / 无畏的                                        | 从恐惧连接现有 fearless 的 fear + -less 构成，帮助记住有无恐惧的词义。 | fearless 是形容词，fear 是名词；不把无畏等同于不考虑风险。已有 fearless 内容保留，只在新 fear 端声明关联。 声明: fear-general → fearless-general |

## 待核实与排除

| 项目                                                                                                                                                                        | 目前可整理的意思                                  | 未确认的内容 / 本批决定                                                                                                        |
| :-------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | :------------------------------------------------ | :----------------------------------------------------------------------------------------------------------------------------- |
| Excess Liquidity (S6)                                                                                                                                                       | 通用保证金语境中常指剩余风险缓冲                  | F14 是另一券商的定义，不能移植到 Moomoo US。暂不新增此词条；需账户字段帮助说明确认本平台定义与口径。                           |
| Turnover (S2)                                                                                                                                                               | 行情中可能指成交额，另有 turnover rate 表示换手率 | 截图只显示排行榜标签，无展开数据；F04 区分两个字段，但未证明该 Crypto 标签的单位、汇总范围和时间窗口。暂不新增。               |
| Smart Money (S3)                                                                                                                                                            | 常见金融表达可指被认为更有经验或信息优势的资金    | 此处也是产品栏目名。未取得可完整检查的词典定义与栏目说明，不把它翻成已被证明聪明、赚钱或机构实时流向，暂不新增。               |
| Short / Long (S8)                                                                                                                                                           | 可能是多空方向，也可能是短长周期或自定义分组      | 截图位于自选分组栏，不能仅凭单词选择金融义项。等待用户说明分组含义，暂不新增。                                                 |
| Fear & Greed Index / Fear (S2)                                                                                                                                              | 恐惧与贪婪情绪指标，以及当前显示的情绪标签        | 已列通用概念候选，并补充 fear、greed 两个基础词；F13 确认名称及历史序列，具体方法、权重和阈值继续待核实，不补入 Notes 为事实。 |
| Latest expected move (S1)                                                                                                                                                   | 预期波幅这一通用概念已列入候选                    | 该词可按一般意思学习；平台窗口与算法仍未知。候选 Notes 明确省略公式和保证范围。                                                |
| Safe / Risk Status (S6)                                                                                                                                                     | 安全状态 / 风险状态标签                           | Safe 只是平台风险标签，不能概括为资产安全或不会亏损。不依据截图新增一张安全机制概念卡。                                        |
| 截图右侧半截标签、下方裁切内容                                                                                                                                              | 未完整辨认                                        | 如公司页右侧仅露出 Pro…、资产页被截断的标签、更多期权字段，不补全原文。                                                        |
| NVDA、TCOM、BTC、ETH、USDT、BNB、USDC、基金与公司全名                                                                                                                       | 证券或资产标识与专名                              | 这次不为每个代码、品牌或产品建卡；symbol 已覆盖代码概念。Berkshire/真实 Buffett 持仓未在图中证实。                             |
| 1Y、24H、USD、M、B、T、% Chg、Chg、Vol、High、Low                                                                                                                           | 时间、货币、量级与常用行情缩写                    | 作为来源阅读提示或后续补充，不扩写成大量基础卡；不把 B/M 的数值与当前行情一起写入 Notes。                                      |
| All、Me、More、Learn、Chats、News、Comments、Company、Settings、Charts/Chart、Strategy、Analysis、Unusual Activity、Market Overview、Rankings、Coins、Closed、Order History | 一般导航、修饰词或可自由组合的界面标签            | 本批优先选账户、资金、财报、报价等高价值目标；这些项目未消失，只是暂不建卡。Order History 由 order 提案覆盖。                  |
| 账户编号、金额、日期、收益率、合约完整编码、新闻标题                                                                                                                        | 截图背景与动态信息                                | 不纳入词条、例句或可推断用户投资行为的解释，不把界面数据当成经核实的当前事实。                                                 |

## 参考来源

访问日期: 2026-11-06。以下仅说明实际检查过的段落与支持范围。金融定义由监管机构、交易所、行业协会、平台文档等第一手来源支持；例句和中文概括由编辑拟写。没有把搜索摘要或无法访问的页面当作完整核查证据。

### F01

[SEC — Beginners’ Guide to Financial Statements](https://www.sec.gov/about/reports-publications/beginners-guide-financial-statements)

已阅读资产、收入、净利润、股东和 EPS 的基础说明。用于一般财务概念，不替代平台账户字段定义。

### F02

[Investor.gov — An Introduction to Options](https://www.investor.gov/introduction-investing/general-resources/news-alerts/alerts-bulletins/investor-bulletins-63)

已阅读上市股票期权、call、put、strike 和到期日说明。草案解释买方权利，不推断截图合约的结算细节。

### F03

[Moomoo API — Get Account Funds](https://openapi.moomoo.com/moomoo-api-doc/en/trade/get-funds.html)

已阅读 power、total_assets、market_val、frozen_cash、avl_withdrawal_cash、现金购买力等字段。API 支持账户术语辨认；不能据字段同名认定每个截图值都有完全相同的计算口径。

### F04

[Moomoo API — Get Real-time Quote](https://openapi.moomoo.com/moomoo-api-doc/en/quote/get-stock-quote.html)

已阅读 code、volume、turnover/turnover_rate、strike_price、contract_size、pre_market 及报价字段。截图中的 Crypto Turnover 统计方法仍未核实。

### F05

[Investor.gov — Understanding Margin Accounts](https://www.investor.gov/introduction-investing/general-resources/news-alerts/alerts-bulletins/investor-bulletins-29)

已阅读现金账户与保证金账户的区别，以及融资、利息和购买力示例。用于一般美国证券账户概念，不写入具体杠杆倍数。

### F06

[Investor.gov — Bid Price/Ask Price](https://www.investor.gov/introduction-investing/investing-basics/glossary/ask-price)

已阅读买方最高买价、卖方最低卖价与报价差的说明。用于报价方向，不能保证订单按当前显示价格成交。

### F07

[CFPB — Annual Percentage Yield Calculation](https://www.consumerfinance.gov/rules-policy/regulations/1030/a/)

已阅读 APY 衡量利息及复利频率的说明。不把截图宣传利率当作当前有效、固定或人人适用的收益。

### F08

[SEC — Non-GAAP Financial Measures, Section 103](https://www.sec.gov/rules-regulations/staff-guidance/corporation-finance-interpretations/non-gaap-financial-measures)

已阅读 EBIT 全称与其不能直接等同 operating income 的说明。区分 EBIT 与每股指标，不凭截图猜 GAAP/non-GAAP 调整口径。

### F09

[Investment Company Institute — ETF Basics and Structure: FAQs](https://www.ici.org/faqs/faqs_etfs)

已阅读 ETF 在交易所按市场价格交易的基金份额定义。不把 ETF 概括成必然被动跟踪指数或保本产品。

### F10

[Coinbase — What is market cap?](https://www.coinbase.com/learn/crypto-basics/what-is-market-cap)

已阅读加密资产市值及 circulating/fully diluted supply 的口径区别。这里只作概念说明，不核实截图总市值。

### F11

[CME Group — Definition of a Futures Contract](https://www.cmegroup.com/education/courses/introduction-to-futures/definition-of-a-futures-contract)

已阅读交易所标准化期货合约定义。不据此认定截图中的 NVDA Futures 属于某一家交易所或某种具体合约。

### F12

[Moomoo Financial Inc. — Customer Agreement, May 28, 2024](https://foss-us-sec.moomoo.com/customer_file_service/prod/disclosure_1717572128614_16222017.pdf)

已阅读 7.13 Margin Account 和 7.14 Cash Sweep 的定义。该文件是历史协议，只支持产品类别的解释；不用于确认当前利率、银行名单、保障额度或具体归集规则。

### F13

[Moomoo Open API — Crypto fear-and-greed history and overview](https://open.moomoo.com/api/quote/crypto-market/crypto-fear-greed-history)

已阅读概览及历史指数序列字段。只确认指标名称与历史/当前读数类别，未验证提供方算法、阈值或预测能力。

### F14

[IBKR Glossary — Current Excess Liquidity](https://www.interactivebrokers.com/campus/glossary-terms/current-excess-liquidity/)

已阅读保证金风险缓冲及不同账户分部的不同公式。仅支持待核实项的通用背景，不将 IBKR 公式套到 Moomoo US。

访问限制: Moomoo US 的部分帮助页和手册在后续读取时返回 403；Cash Sweep 完整条款页面仅显示需要 JavaScript。Cambridge 的 Smart Money/Financials 页面返回 403。上述页面不作为已完整核实的证据；不继续请求受限页面来绕过限制。F12 只作为注明日期的历史定义，F03/F04 不代表界面所有字段与 API 字段一一相同。

### F15

[Investor.gov — Types of Orders](https://www.investor.gov/introduction-investing/investing-basics/how-stock-markets-work/types-orders)

已阅读 market、limit、stop 订单说明及委托与成交价的区别。用于 order 这一买卖指令含义，不新增这些订单类型卡或给出下单策略。

## 来源到目标台账

按完整可辨认且选中的来源写法去重；图中相同词重复出现归入同一项。源标签位置详见候选表，核心来源写法如下。五个额外子项 profit、loss、withdraw、share、greed 独立列在后面，不假装是屏幕上的完整词。

| 来源写法                   | 目标稳定 ID                                               | 处理                                       |
| :------------------------- | :-------------------------------------------------------- | :----------------------------------------- |
| % Portfolio                | portfolio-finance                                         | 规范化 / 复用 / 保留完整术语；具体见候选表 |
| 1Y Return                  | return-finance                                            | 规范化 / 复用 / 保留完整术语；具体见候选表 |
| 2026/Q4                    | quarter-finance                                           | 规范化 / 复用 / 保留完整术语；具体见候选表 |
| 2027/Q1                    | quarter-finance                                           | 规范化 / 复用 / 保留完整术语；具体见候选表 |
| 2027/Q2                    | quarter-finance                                           | 规范化 / 复用 / 保留完整术语；具体见候选表 |
| 2027/Q3                    | quarter-finance                                           | 规范化 / 复用 / 保留完整术语；具体见候选表 |
| 2027/Q4                    | quarter-finance                                           | 规范化 / 复用 / 保留完整术语；具体见候选表 |
| APY                        | annual-percentage-yield-finance                           | 规范化 / 复用 / 保留完整术语；具体见候选表 |
| Account Management         | account-finance                                           | 规范化 / 复用 / 保留完整术语；具体见候选表 |
| Accounts                   | account-finance                                           | 规范化 / 复用 / 保留完整术语；具体见候选表 |
| Actual                     | actual-general                                            | 规范化 / 复用 / 保留完整术语；具体见候选表 |
| Actual (YoY)               | actual-general                                            | 规范化 / 复用 / 保留完整术语；具体见候选表 |
| All Accounts               | account-finance                                           | 规范化 / 复用 / 保留完整术语；具体见候选表 |
| Ask                        | ask-price-finance                                         | 规范化 / 复用 / 保留完整术语；具体见候选表 |
| Assets                     | asset-finance                                             | 规范化 / 复用 / 保留完整术语；具体见候选表 |
| At Discount                | discount-finance                                          | 规范化 / 复用 / 保留完整术语；具体见候选表 |
| Beat Est.                  | beat-estimates-finance                                    | 规范化 / 复用 / 保留完整术语；具体见候选表 |
| Bid                        | bid-price-finance                                         | 规范化 / 复用 / 保留完整术语；具体见候选表 |
| Buffett Holdings           | holdings-finance                                          | 规范化 / 复用 / 保留完整术语；具体见候选表 |
| CALL                       | call-option-finance                                       | 规范化 / 复用 / 保留完整术语；具体见候选表 |
| Call                       | call-option-finance                                       | 规范化 / 复用 / 保留完整术语；具体见候选表 |
| Cash Account               | cash-account-finance                                      | 规范化 / 复用 / 保留完整术语；具体见候选表 |
| Cash BP                    | buying-power-finance                                      | 规范化 / 复用 / 保留完整术语；具体见候选表 |
| Cash Sweep                 | cash-sweep-finance                                        | 规范化 / 复用 / 保留完整术语；具体见候选表 |
| Cash Transfer              | transfer-finance                                          | 规范化 / 复用 / 保留完整术语；具体见候选表 |
| Chain                      | option-chain-finance                                      | 规范化 / 复用 / 保留完整术语；具体见候选表 |
| Closing Price              | closing-price-finance                                     | 规范化 / 复用 / 保留完整术语；具体见候选表 |
| Community                  | community-general                                         | 规范化 / 复用 / 保留完整术语；具体见候选表 |
| Crypto                     | crypto-finance                                            | 规范化 / 复用 / 保留完整术语；具体见候选表 |
| Crypto Market Cap          | crypto-finance, market-capitalization-finance             | 规范化 / 复用 / 保留完整术语；具体见候选表 |
| Crypto Stocks              | crypto-finance, stock-finance                             | 规范化 / 复用 / 保留完整术语；具体见候选表 |
| Currency                   | currency-finance                                          | 规范化 / 复用 / 保留完整术语；具体见候选表 |
| Currency/Cash              | cash-finance, currency-finance                            | 规范化 / 复用 / 保留完整术语；具体见候选表 |
| Deposit                    | deposit-finance                                           | 规范化 / 复用 / 保留完整术语；具体见候选表 |
| Discover                   | discover-general                                          | 规范化 / 复用 / 保留完整术语；具体见候选表 |
| Dividends                  | dividend-finance                                          | 规范化 / 复用 / 保留完整术语；具体见候选表 |
| EBIT                       | ebit-finance                                              | 规范化 / 复用 / 保留完整术语；具体见候选表 |
| EPS                        | earnings-per-share-finance                                | 规范化 / 复用 / 保留完整术语；具体见候选表 |
| ETFs                       | exchange-traded-fund-finance                              | 规范化 / 复用 / 保留完整术语；具体见候选表 |
| Earnings Hub               | earnings-finance                                          | 规范化 / 复用 / 保留完整术语；具体见候选表 |
| Earnings Move & Volatility | earnings-finance, volatility-finance                      | 规范化 / 复用 / 保留完整术语；具体见候选表 |
| Estimate                   | estimate-general                                          | 规范化 / 复用 / 保留完整术语；具体见候选表 |
| Estimate (YoY)             | estimate-general                                          | 规范化 / 复用 / 保留完整术语；具体见候选表 |
| Exchange                   | exchange-currency                                         | 规范化 / 复用 / 保留完整术语；具体见候选表 |
| FY2027                     | fiscal-year-finance                                       | 规范化 / 复用 / 保留完整术语；具体见候选表 |
| Fear                       | fear-general                                              | 规范化 / 复用 / 保留完整术语；具体见候选表 |
| Fear & Greed Index         | fear-and-greed-index-finance, fear-general, index-finance | 规范化 / 复用 / 保留完整术语；具体见候选表 |
| Financial Estimates        | estimate-general                                          | 规范化 / 复用 / 保留完整术语；具体见候选表 |
| Financials (USD)           | financials-finance                                        | 规范化 / 复用 / 保留完整术语；具体见候选表 |
| Funds on Hold              | funds-on-hold-finance                                     | 规范化 / 复用 / 保留完整术语；具体见候选表 |
| Futures                    | futures-finance                                           | 规范化 / 复用 / 保留完整术语；具体见候选表 |
| Heat Map                   | heat-map-data                                             | 规范化 / 复用 / 保留完整术语；具体见候选表 |
| Index                      | index-finance                                             | 规范化 / 复用 / 保留完整术语；具体见候选表 |
| Latest expected move       | expected-move-finance                                     | 规范化 / 复用 / 保留完整术语；具体见候选表 |
| MV/QTY                     | market-value-finance, quantity-general                    | 规范化 / 复用 / 保留完整术语；具体见候选表 |
| Margin Account             | margin-account-finance, margin-finance                    | 规范化 / 复用 / 保留完整术语；具体见候选表 |
| Market Cap                 | market-capitalization-finance                             | 规范化 / 复用 / 保留完整术语；具体见候选表 |
| Market Value               | market-value-finance                                      | 规范化 / 复用 / 保留完整术语；具体见候选表 |
| Max Buying Power           | buying-power-finance                                      | 规范化 / 复用 / 保留完整术语；具体见候选表 |
| Net Assets                 | net-assets-finance, net-finance                           | 规范化 / 复用 / 保留完整术语；具体见候选表 |
| Net Income                 | income-finance, net-finance, net-income-finance           | 规范化 / 复用 / 保留完整术语；具体见候选表 |
| Options                    | option-finance                                            | 规范化 / 复用 / 保留完整术语；具体见候选表 |
| Order History              | order-finance                                             | 规范化 / 复用 / 保留完整术语；具体见候选表 |
| Orders                     | order-finance                                             | 规范化 / 复用 / 保留完整术语；具体见候选表 |
| P/L Analysis               | profit-and-loss-finance                                   | 规范化 / 复用 / 保留完整术语；具体见候选表 |
| PUT                        | put-option-finance                                        | 规范化 / 复用 / 保留完整术语；具体见候选表 |
| Pending Interest           | interest-finance, pending-tech                            | 规范化 / 复用 / 保留完整术语；具体见候选表 |
| Positions                  | position-finance                                          | 规范化 / 复用 / 保留完整术语；具体见候选表 |
| Pre-Mkt                    | pre-market-finance                                        | 规范化 / 复用 / 保留完整术语；具体见候选表 |
| Price/Cost                 | cost-finance                                              | 规范化 / 复用 / 保留完整术语；具体见候选表 |
| Put                        | put-option-finance                                        | 规范化 / 复用 / 保留完整术语；具体见候选表 |
| Q2                         | quarter-finance                                           | 规范化 / 复用 / 保留完整术语；具体见候选表 |
| Q2 FY2027 Earnings         | earnings-finance                                          | 规范化 / 复用 / 保留完整术语；具体见候选表 |
| Revenue                    | revenue-finance                                           | 规范化 / 复用 / 保留完整术语；具体见候选表 |
| Shareholders               | shareholder-finance                                       | 规范化 / 复用 / 保留完整术语；具体见候选表 |
| Single Option              | option-finance                                            | 规范化 / 复用 / 保留完整术语；具体见候选表 |
| Strike                     | strike-price-finance                                      | 规范化 / 复用 / 保留完整术语；具体见候选表 |
| Symbol                     | symbol-market                                             | 规范化 / 复用 / 保留完整术语；具体见候选表 |
| Today’s Interest           | interest-finance                                          | 规范化 / 复用 / 保留完整术语；具体见候选表 |
| Today’s P/L                | profit-and-loss-finance                                   | 规范化 / 复用 / 保留完整术语；具体见候选表 |
| Top Gainers                | gainer-finance                                            | 规范化 / 复用 / 保留完整术语；具体见候选表 |
| Top Losers                 | loser-finance                                             | 规范化 / 复用 / 保留完整术语；具体见候选表 |
| Total Assets               | asset-finance                                             | 规范化 / 复用 / 保留完整术语；具体见候选表 |
| Total Cash                 | cash-finance                                              | 规范化 / 复用 / 保留完整术语；具体见候选表 |
| Total Interest             | interest-finance                                          | 规范化 / 复用 / 保留完整术语；具体见候选表 |
| Total Position P/L         | position-finance, profit-and-loss-finance                 | 规范化 / 复用 / 保留完整术语；具体见候选表 |
| Trade                      | trade-finance                                             | 规范化 / 复用 / 保留完整术语；具体见候选表 |
| Transaction                | transaction-finance                                       | 规范化 / 复用 / 保留完整术语；具体见候选表 |
| US Margin Account          | margin-account-finance                                    | 规范化 / 复用 / 保留完整术语；具体见候选表 |
| Volume                     | volume-trading                                            | 规范化 / 复用 / 保留完整术语；具体见候选表 |
| Vs Estimate (% Chg)        | estimate-general                                          | 规范化 / 复用 / 保留完整术语；具体见候选表 |
| Watchlists                 | watchlist-finance                                         | 规范化 / 复用 / 保留完整术语；具体见候选表 |
| Wealth                     | wealth-general                                            | 规范化 / 复用 / 保留完整术语；具体见候选表 |
| Withdrawable               | withdrawable-finance                                      | 规范化 / 复用 / 保留完整术语；具体见候选表 |
| YoY                        | year-over-year-finance                                    | 规范化 / 复用 / 保留完整术语；具体见候选表 |

| 额外子项                    | 提出依据                                                                                                          | 状态                   |
| :-------------------------- | :---------------------------------------------------------------------------------------------------------------- | :--------------------- |
| profit-finance / profit     | 为理解 P/L = profit and loss 而补充的基础子项，截图未把 profit 全拼出来。                                         | 拟新增；已计入候选总数 |
| loss-finance / loss         | 为理解 P/L = profit and loss 而补充的基础子项，截图未把 loss 全拼出来。                                           | 拟新增；已计入候选总数 |
| withdraw-finance / withdraw | 从 Withdrawable 提出的有独立学习价值的动词子项；截图并未出现独立 Withdraw 按钮。                                  | 拟新增；已计入候选总数 |
| share-finance / share       | 为 Shareholders 和 EPS 全称理解提出的基础子项；截图未显示独立 Share 标签。不是分享这一动词义。                    | 拟新增；已计入候选总数 |
| greed-general / greed       | 从 Fear & Greed Index 提出的有用基础子项，不声称有独立 Greed 按钮。指标中的情绪称呼不等于对每位投资者的道德判断。 | 拟新增；已计入候选总数 |

## 实施与验证

内容候选、例句与 Notes 草案、缩写与组合处理，以及单独列出的 W01–W04、N01 具体文案和 R01–R11 关联均已于 2026-11-06 确认。用户随后明确批准实施，已按下列流程完成。未确认项目继续留空或暂缓；不以批准本计划替代对未知平台公式的核实。当前批次未生成音频。

1. 重新读取本计划、当前仓库规范、源截图与 Git 状态，按稳定 ID 再次排重并核对原有词义。记录新旧源文件与所有现有录音/归档哈希，保留无关工作。若期间已出现某个目标，先复用现有 ID，不创建重复项。

2. 只创建批准的新词条；复用 pending-tech 不改其既有内容。写入批准的 Meaning、Example、Notes、优先级和有用的先修关系，按本计划单独批准的范围添加记忆字段与关系。所有先修和关联 ID 必须已经存在，不能写悬空引用。可选 Context/IPA 不为填充而生成，发音未核实就省略。

3. 如以后另行批准音频，先运行全库验证，再检查 macOS、say、lame、mpg123 与匹配的英语语音。只用现有生成器逐个稳定 ID 生成或复用，不使用 --all/--overwrite，不安装工具，不覆盖已有录音。当前计划本身不授权录音。

```bash
go test ./...
go vet ./...
go run scripts/validate/main.go
go run scripts/build/main.go
git diff --check
```

4. 校验实际新建/复用/更新数量、稳定 ID 全局唯一、同语言词义重复、依赖无环、前置顺序、17 列 TSV 与新条目英文 TTS locale。构建两次比较字节；新增条目会使 StudyOrder 变化，需明确报告。不得手改 dist。

5. 对已批准关系校验两端的直接双向显示，不把现有或拟定先修当作已批准关联。比较原始词条与录音哈希；未授权的既有条目、媒体、模板、依赖和实现保持不变。对未跟踪文本另查格式；报告来源未解问题，不把静态导出检查说成真实 Anki 导入/播放测试。

6. 更新计划批准、实施及验证状态，报告实际文件和命令结果。不暂存、提交、推送、复制媒体到 Anki 或修改 Anki 集合。

## 规划检查

规划阶段完成以下检查，当时未运行 Go 构建、音频生成或 Anki 操作:

- 8 个截图路径可访问；207 个本地文件或标题锚点链接解析通过。15 个参考 URL 对应本次实际检查的页面或段落，受限页面已单独说明，未将访问失败或搜索摘要标为完整验证。
- 79 个候选编号与稳定 ID 唯一，78 个拟新增 ID 不与现有集合冲突；pending-tech 为唯一复用项。95 个源标签写法的映射均指向候选 ID；5 个额外子项已包含在新增数内，不重复计数。
- 所有先修 ID 指向当前集合或本批拟新增目标，合并依赖图无环。11 对关系无重复，当前集合也不存在任一方向的直接声明；R11 包含现有 fearless-general，其余关系两端均在本批候选内。
- 已检查所有候选的 Word、Meaning、Example、Notes 草案，及 5 项可选记忆字段的具体文案。未填写未经可靠核验的 IPA 或历史词源。symbol 与现有 JavaScript Symbol 属于不同义项，重复词警告是预期的人工审阅项。
- Markdown 表格列数、左对齐、代码围栏、行尾空白、制表符和最终换行检查通过。git diff --check 通过；本文件未跟踪，另做了文本格式检查。
- 规划前记录的 2,558 个项目文件哈希均未变化，包括全部词条、现有及归档录音、实现、模板、依赖和已有导出。唯一新增文件为本计划。未暂存、提交或推送。

确认记录检查 (实施前): 仅更新内容确认范围与待办状态，候选具体文案和计数不变；当时的文件保全比较确认未修改其他项目文件。


## 实施记录

2026-11-06，用户明确批准实施。以下为本次实际结果，与上文规划基线分开记录。

- 新增 78 个英语 YAML 词条，复用 pending-tech，未修改任何原有词条。集合从 1,242 增至 1,320 个稳定 ID，其中英语 1,077 个、日语 243 个。新增目标包含计划中的 5 个独立基础子项，没有额外组合卡。
- 逐字段核对新增内容与批准规格一致，包括 Meaning、Example、Notes、来源行、参考资料、学习优先级、37 条先修引用和 5 项记忆辅助 (W01–W04、N01)。没有新增 Context、IPA、历史词源或音频字段。
- 11 对批准关联各声明一次；导出保留原有关系，全库合计 126 对直接关系、252 个方向的邻居展示，涉及 236 个词条。原有词条中仅 fearless-general 的派生 RelatedWords 增加 fear-general；其源文件没有变化。
- 原有 1,242 个词条的源字节与解析值均保持不变；1,242 个现有 MP3 和 43 个归档 MP3 的 SHA-256 全部保持不变。实现、模板、依赖及其他原有文件保持不变，只有本计划与生成的导出属于允许变化的原有文件。
- TSV 含 1,320 行数据、每行 17 列；新增行的 TTSLanguage 为 en_US、Audio 为空。StudyOrder 为唯一的四位补零顺序，全部先修在目标前；原有条目之间的相对顺序不变，1,241 个原有条目的排名因插入新项而变化。
- 全部直接邻居展示、转义、标签及新增字段与源数据一致。原有导出行除 StudyOrder 和上述批准的 RelatedWords 外保持不变。没有新增未批准关联或传递关联。
- go test ./...、go vet ./...、词库验证、两次构建与 git diff --check 均通过。验证共有 15 条同形词警告，比基线新增 1 条: symbol-market 的证券代码义与 symbol-javascript 的 JavaScript Symbol 义不同，按批准方案分别保留。
- 两次无源变更的构建逐字节一致，SHA-256 为 `e3d26ebd36219a5d22e5c01fd04b769b7f47f7098461a4d9498d7e73b5118923`。未手改 dist；新增未跟踪文本另行检查格式和本地链接。

本批未生成或解码音频，未进行人工听审、真实 Anki 导入、显示或播放验证。未暂存、提交、推送或修改 Anki 集合。使用完整导出导入后，如需应用新的引入顺序，应按[学习顺序指南](LEARNING_ORDER.md#apply-the-order-in-anki)重排新卡；导入 StudyOrder 本身不会重排现有新卡。

Excess Liquidity、Turnover、Smart Money、Short / Long 等暂缓项及未核实的平台公式仍未收录为已知事实。
