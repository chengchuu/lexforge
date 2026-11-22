# LexForge 全库记忆辅助审阅计划

状态为已批准并实施。规划日期为 2026-11-22。用户明确批准两轮全部文本与关联。下文保留规划基线和审阅时的文案，实施结果记录在文末。

## 范围与基线

扫描当前全部 YAML 的稳定 ID。逐项读取语言、Word、Meaning、Context、Example、Notes、现有辅助、先修和关系，并按现有释义、语境和用法线索筛选。上次塞尔达批次已实施，现有修改作为本轮基线保留。参见 [塞尔达实施记录](ZELDA_NOV_22_MEMORY_AIDS_ENRICHMENT_PLAN.md)。

保留已清楚表达的内容。优先选择可迁移的词形规律、能帮助理解当前义项的来源，以及确实可能混淆的用法。一项可以继续全空。逐项筛查不等同于逐个查证历史来源。

已核对 [仓库规则](../AGENTS.md)、[卡片模板](../README.md)、[字段规范](ENTRY_SCHEMA.md) 与 [学习顺序](LEARNING_ORDER.md)。现有解析器接受可选的三种文本辅助。RelatedWords 由稳定 ID 邻接派生，去重后双向显示直接邻居，按 ID 排序。跨语言邻居显示语言标记。先修不产生关联。空的辅助字段由卡背条件隐藏。

| 项目           | 当前基线或拟议数量                                                  |
|:---------------|:--------------------------------------------------------------------|
| 审阅词条       | 1833 个稳定 ID，其中英语 1590 个，日语 243 个                       |
| 现有非空辅助   | ETY 46，WF 164，N 110。合计 320 个值，覆盖 304 个 ID                |
| 现有直接关系   | 146 对，全部保留                                                    |
| 拟修改辅助字段 | 12 个，WF 5、ETY 3、N 4。其中 11 个新增、1 个移入，无替换           |
| 拟修改 Notes   | 7 个字段，7 条来源行，其中 WF01 同时移出原句末短语                  |
| 新关系         | 10 对，20 个端点，10 个声明端，新增 10 条反向显示，共 20 条直接展示 |
| 受影响 ID      | 文本涉及 12 个 ID，关系涉及 20 个 ID，重叠 7 个。合并为 25 个 ID    |
| 拟改 YAML      | 17 个文件和 29 个字段值。派生的右端显示不单独计入 YAML              |

各语言台账计数如下。“保留”归入本轮无提案且已有辅助或直接邻居的条目。所有提案或暂缓条目也保留其已有未获替换批准的内容。

| 语言 | 提案或反向显示 | 暂缓 | 保留 | 无提案且无现有辅助或关联 |
|:-----|:---------------|:-----|:-----|:-------------------------|
| en   | 18             | 6    | 360  | 1206                     |
| ja   | 7              | 4    | 106  | 126                      |

当前非空辅助的字段分布与原样保留对象见台账。现有语义对照如 `permission-tech` / `approval-workflow`、`dare` / `donata`、`machigaeru` / `machigai`、`gaikokujin` / `daigakusei` 和 `checkpoint-tech` / `snapshot-tech` 已能双向显示，因而不重复声明。现有技术说明，如正反例、代码和机制边界仍留在原字段。全部提案均使用现有字段。

## WF 构词

### WF01 invalidate

[invalidate-tech](../entries/invalidate-tech.yml) · invalidate · 使失效

移入并扩写至 `word_formation`。

```plain
invalidate 表示使失效，invalidation 是对应的名词。把 -ate 换成 -ation，动作就变成过程或结果的名称。
```

从代码里的动作推到日志或文档里的名词。 只解释当前词族和失效义。 证据见 D01。

移出前的完整 Notes 如下。只删除末尾的 “ Related form: invalidation.”，原来源句完整保留。

```plain
Source: typescript/typescript-64936.md. Related form: invalidation.
```

追加下方出处前的完整 Notes 如下。

```plain
Source: typescript/typescript-64936.md.
```

下列单行以一个空行与现有 Notes 分隔后追加。实施时若同一 URL 或同一出处行已存在，保留原位置并跳过追加。

```plain
Memory-aid source: Merriam-Webster.com Dictionary — invalidate. https://www.merriam-webster.com/dictionary/invalidate
```

### WF02 cache revalidation

[cache-revalidation-web](../entries/cache-revalidation-web.yml) · cache revalidation · 缓存再验证；向服务器确认已存储响应能否继续使用。

新增 `word_formation`。

```plain
cache（缓存）+ revalidation（再次验证）。把 re- 读成“再”，就能认出整个短语是在重新确认缓存能否继续使用。
```

用前缀把再次确认与失效处理分开。 说明现代构成，保留现有 Notes 里的协议细节。 证据见 D02。

Notes 已含同一来源 URL。保留全文，直接复用现有出处。

### WF03 大学生

[daigakusei](../entries/daigakusei.yml) · 大学生 · 大学生；在大学就读的学生

新增 `word_formation`。

```plain
大学（だいがく，大学）+ 生（せい，此处表示学生身份），合起来是大学生。生 在这里读 せい。
```

用大学和身份成分记住词序及读法。 只拆当前词形。现有外国人辨音助记及双向关联原样保留。 证据见 D03。

Notes 已含同一来源 URL。保留全文，直接复用现有出处。

### WF04 ～まで

[made-grammar](../entries/made-grammar.yml) · ～まで · 直到……；到……为止

新增 `word_formation`。

```plain
地点名词 + まで 表示到该处为止，如 駅まで。动词普通形 + まで 表示持续到该事件，如 火影になるまで。
```

把同一终点概念用于地点和事件两种句型。 只覆盖当前卡片的终点用法。 证据见 D04。

下列单行以一个空行与现有 Notes 分隔后追加。实施时若同一 URL 或同一出处行已存在，保留原位置并跳过追加。

```plain
Memory-aid source: Tofugu — Particle まで: Endpoint Marker. https://www.tofugu.com/japanese-grammar/particle-made/
```

### WF05 reuse

[reuse-general](../entries/reuse-general.yml) · reuse · 重复使用；再次利用

新增 `word_formation`。

```plain
re-（再次）+ use（使用）。同一件东西再用一次，就可以用 reuse 描述。
```

复用 re- 理解再次使用。 只说明当前词义及现代构成。 证据见 D05。

下列单行以一个空行与现有 Notes 分隔后追加。实施时若同一 URL 或同一出处行已存在，保留原位置并跳过追加。

```plain
Memory-aid source: Merriam-Webster.com Dictionary — reuse. https://www.merriam-webster.com/dictionary/reuse
```

## ETY 词源

### ETY01 portfolio

[portfolio-finance](../entries/portfolio-finance.yml) · portfolio · 投资组合；同一投资者或账户持有的一组投资资产。

新增 `etymology`。

```plain
portfolio 借自意大利语 portafoglio，原指携带散页的文件夹。portare 是“携带”，foglio 是“纸页”。可以用一只夹子收纳多页的画面记住现代投资组合中的“一组资产”。
```

用原义与现代集合义建立记忆联系。 最后一句是学习联想，词源资料未在此证明金融用法的具体演变路径。 证据见 D06。

下列单行以一个空行与现有 Notes 分隔后追加。实施时若同一 URL 或同一出处行已存在，保留原位置并跳过追加。

```plain
Memory-aid source: Online Etymology Dictionary — portfolio. https://www.etymonline.com/word/portfolio
```

### ETY02 dividend

[dividend-finance](../entries/dividend-finance.yml) · dividend · 股息；公司向股东分配的收益。

新增 `etymology`。

```plain
dividend 来自拉丁语 dividendum，意思近于“有待分配之物”，与表示分开、分配的 dividere 同源。用“可分配的一份”记住当前的股息义。
```

把近似 divide 的拼写与股东分配概念连起来。 最后一句是联想。公司分配的具体条件沿用现有 Notes。 证据见 D07。

下列单行以一个空行与现有 Notes 分隔后追加。实施时若同一 URL 或同一出处行已存在，保留原位置并跳过追加。

```plain
Memory-aid source: Online Etymology Dictionary — dividend. https://www.etymonline.com/word/dividend
```

### ETY03 bug

[bug-tech](../entries/bug-tech.yml) · bug · 程序缺陷；使软件偏离预期行为的问题。

新增 `etymology`。

```plain
bug 的机器故障义在十九世纪已出现，后来也用于软件缺陷。早期机械义的确切造词过程仍不确定。
```

把早期机械故障与当前代码缺陷联系起来。 历史资料列出早期用例和可能来源。暂不采用特定人物首创或飞蛾起源的说法。 证据见 D08。

下列单行以一个空行与现有 Notes 分隔后追加。实施时若同一 URL 或同一出处行已存在，保留原位置并跳过追加。

```plain
Memory-aid source: Online Etymology Dictionary — bug. https://www.etymonline.com/word/bug
```

## N 助记

### N01 cache invalidation

[cache-invalidation-tech](../entries/cache-invalidation-tech.yml) · cache invalidation · 缓存失效处理；在依赖变化后使旧缓存结果不再被当作有效结果使用。

新增 `mnemonic`。

```plain
记忆画面：缓存柜上有两张标签。失效时把旧结果标成“先别用”，再验证时拿着旧凭据问“还能用吗”。
```

在读技术文档时区分使旧结果失效与重新确认可用性。 标签是原创图像，具体 HTTP 行为见相关词条与 MDN。 证据见 D02。

下列单行以一个空行与现有 Notes 分隔后追加。实施时若同一 URL 或同一出处行已存在，保留原位置并跳过追加。

```plain
Memory-aid source: MDN — HTTP caching. https://developer.mozilla.org/en-US/docs/Web/HTTP/Guides/Caching
```

### N02 bind mount

[bind-mount-docker](../entries/bind-mount-docker.yml) · bind mount · 绑定挂载；把宿主机中的具体文件或目录接入容器。

新增 `mnemonic`。

```plain
记忆画面：一根线从容器里连到宿主机上你指定的文件夹，这是 bind mount。另一根线连到由 Docker 管理的存储，这是 volume。
```

先辨认挂载源由谁指定和管理。 原创图像。当前卡片对照常规绑定挂载与 Docker 管理的卷，配置细节依环境而异。 证据见 D09。

Notes 已含同一来源 URL。保留全文，直接复用现有出处。

### N03 position

[position-finance](../entries/position-finance.yml) · position · 持仓；账户对某项金融工具持有的数量或形成的风险敞口。

新增 `mnemonic`。

```plain
记忆画面：表格的一行写一项持仓，整张表放进投资组合文件夹。看一项用 position，看这一组用 portfolio。
```

用单项敞口与集合两个层次理解界面标签。 原创图像，依据当前词义。表格布局和估值口径依界面而异。 证据见 L02。

依据为现有词条，Notes 原样保留。

### N04 間違い

[machigai](../entries/machigai.yml) · 間違い · 错误；弄错

新增 `mnemonic`。

```plain
用法助记：电话里说自己拨错了，可以用 番号を間違えました，关注做错这件事。检查资料时说名字有误，可以用 名前に間違いがあります，关注发现的错误。
```

把名词和动词放进常见电话和资料场景对照。 依据两张现有卡片的句型。只对比词性和搭配。 证据见 L03。

依据为现有词条，Notes 原样保留。

## R 关联单独审阅

只在每行左端的 `related_entries` 声明右端，按编号顺序追加，保留旧列表成员。通过入站与出站联合去重。所有候选的语义依据为表中当前词义及链接词条的语境、例句或 Notes。WF、ETY 或 N 获批不成立关系批准。本批暂无跨语言候选。

| 编号 | 声明端 ID · Word · Meaning                                                                                                                              | 目标 ID · Word · Meaning                                                                                                                | 联系                                 | 混淆风险                                                        |
|:-----|:--------------------------------------------------------------------------------------------------------------------------------------------------------|:----------------------------------------------------------------------------------------------------------------------------------------|:-------------------------------------|:----------------------------------------------------------------|
| R01  | [cache-invalidation-tech](../entries/cache-invalidation-tech.yml) · cache invalidation · 缓存失效处理；在依赖变化后使旧缓存结果不再被当作有效结果使用。 | [cache-revalidation-web](../entries/cache-revalidation-web.yml) · cache revalidation · 缓存再验证；向服务器确认已存储响应能否继续使用。 | 对照淘汰旧结果与确认可复用性         | 两词近形，失效处理与 HTTP 再验证的目的不同。                    |
| R02  | [bind-mount-docker](../entries/bind-mount-docker.yml) · bind mount · 绑定挂载；把宿主机中的具体文件或目录接入容器。                                     | [volume-docker](../entries/volume-docker.yml) · volume · 卷；由 Docker 管理、可独立于容器生命周期存在的存储。                           | 按挂载源区分容器存储                 | 两者都能在容器里呈现为目录。                                    |
| R03  | [image-docker](../entries/image-docker.yml) · image · 镜像；用于创建容器的只读模板。                                                                    | [container-docker](../entries/container-docker.yml) · container · 容器；从镜像创建的隔离进程运行环境。                                  | 对照创建用模板与运行环境             | 镜像与由它创建的实例是两个层次。                                |
| R04  | [alias-nginx](../entries/alias-nginx.yml) · alias · 别名指令；在此前缀位置中将匹配到的 URI 前缀替换为指定路径。                                         | [root-nginx](../entries/root-nginx.yml) · root · 文档根目录指令；把完整 URI 路径接到根目录后面得到文件路径。                            | 对照同一实验中的 URI 到文件路径映射  | 默认记完整 URI 追加与位置前缀替换的区别。范围限于现有前缀实验。 |
| R05  | [context-window-ai](../entries/context-window-ai.yml) · context window · 上下文窗口；模型一次处理可容纳的上下文范围或容量，通常以词元计量。             | [memory-ai](../entries/memory-ai.yml) · memory · 记忆；智能体保存并可在后续交互中重新使用的信息。                                       | 区分一次输入能容纳什么与应用保存什么 | 保存的信息需要被选入当前输入才会占窗口。                        |
| R06  | [daigakusei](../entries/daigakusei.yml) · 大学生 · 大学生；在大学就读的学生                                                                             | [gakusei](../entries/gakusei.yml) · 学生 · 学生                                                                                         | 对照较窄的大学生身份与一般学生身份   | 大学生是学生的一种。既有外国人对照继续保留。                    |
| R07  | [donata](../entries/donata.yml) · どなた · 哪位                                                                                                         | [kata-person](../entries/kata-person.yml) · 方 · 人；那位（礼貌说法）                                                                   | 对照礼貌问“哪位”与礼貌指代某人       | どなた 是疑问词，方 在当前义项可用于指代。                      |
| R08  | [made-grammar](../entries/made-grammar.yml) · ～まで · 直到……；到……为止                                                                                 | [hokage-ni-naru-made](../entries/hokage-ni-naru-made.yml) · 火影になるまで · 直到成为火影                                               | 把终点句型连到完整例句               | 词条一是语法，另一端是一个具体事件。                            |
| R09  | [portfolio-finance](../entries/portfolio-finance.yml) · portfolio · 投资组合；同一投资者或账户持有的一组投资资产。                                      | [position-finance](../entries/position-finance.yml) · position · 持仓；账户对某项金融工具持有的数量或形成的风险敞口。                   | 区分单项持仓与组合                   | 当前口径下持仓可以是组合的一部分。                              |
| R10  | [resume-workflow](../entries/resume-workflow.yml) · resume · 恢复继续；在中断或暂停后继续进行。                                                         | [resumption-workflow](../entries/resumption-workflow.yml) · resumption · 恢复继续                                                       | 对照继续执行的动词与过程名词         | 注意词尾形态以及 résumé 的不同义项。                            |

## 来源与证据边界

D01–D10 的相关正文于 2026-11-22 查阅。词源栏陈述所引用的历史证据，WF 标注现代构词或语法，N 的记忆画面为本计划原创，不表示事件或命名由来。

| 编号 | 来源                                                                                                        | 支持内容                                                                                    |
|:-----|:------------------------------------------------------------------------------------------------------------|:--------------------------------------------------------------------------------------------|
| D01  | [Merriam-Webster.com Dictionary — invalidate](https://www.merriam-webster.com/dictionary/invalidate)        | 支持使某物失效的动词义及名词 invalidation。                                                 |
| D02  | [MDN — HTTP caching](https://developer.mozilla.org/en-US/docs/Web/HTTP/Guides/Caching)                      | 支持使用条件请求重新验证已保存的响应，以及 304 与复用的关系。                               |
| D03  | [Langoal — 大学生](https://langoal.com/ja/vocabulary/word/daigakusei.html)                                  | 支持一般词义及现有读法。大学 + 生 属于本计划的现代词形分析。                                |
| D04  | [Tofugu — Particle まで: Endpoint Marker](https://www.tofugu.com/japanese-grammar/particle-made/)           | 支持地点或时间终点，以及动词普通形 + まで。                                                 |
| D05  | [Merriam-Webster.com Dictionary — reuse](https://www.merriam-webster.com/dictionary/reuse)                  | 支持再次使用义。re- + use 是现代构词提示。                                                  |
| D06  | [Online Etymology Dictionary — portfolio](https://www.etymonline.com/word/portfolio)                        | 支持英语旧义、意大利语形式和成分。现代金融义由当前词条支持，不从词源推定金融义出现年代。    |
| D07  | [Online Etymology Dictionary — dividend](https://www.etymonline.com/word/dividend)                          | 支持拉丁语形式及分配义。当前金融义由词条支持。                                              |
| D08  | [Online Etymology Dictionary — bug](https://www.etymonline.com/word/bug)                                    | 支持机械故障义早于现代软件环境。早期用法起源仍有不确定性。                                  |
| D09  | [Docker Docs — Bind mounts](https://docs.docker.com/engine/storage/bind-mounts/)                            | 支持宿主机路径与 Docker 管理的卷的区别。                                                    |
| D10  | [Docker Docs — Volumes](https://docs.docker.com/engine/storage/volumes/)                                    | 支持 Docker 管理卷的事实，用于核对存储对照。                                                |
| L01  | 当前词条及双向声明清单                                                                                      | 支持所有 R 项的现有释义、重复排除及指定的声明方向。Docker 与 Nginx 的术语释义限于当前词条。 |
| L02  | [portfolio-finance](../entries/portfolio-finance.yml) · [position-finance](../entries/position-finance.yml) | 支持投资组合与单项持仓两个层次。                                                            |
| L03  | [machigai](../entries/machigai.yml) · [machigaeru](../entries/machigaeru.yml)                               | 支持名词和动词的词性及例句搭配。                                                            |

访问限制包括 Smithsonian 的飞蛾资料返回 500，Merriam-Webster 的 mutex 词页和尝试的 Tofugu ～んだ 候选页未能读取。Docker overview 的旧地址返回 404，镜像与容器配对采用本地条目的现有义项。上述页面均不作为已查证材料。bug 只使用词源词条中关于机械故障义的证据。游戏和组织专名的来源说明见台账中的暂缓项。日语候选集中在构词与句型。假名沿用现有读法。

## 全库审阅台账

每个稳定 ID 只列一次。既有字段缩写指非空值。既有邻居数含反向派生。带提案编号的行也保留其旧辅助和关系。

### 英语，1590 项

| 稳定 ID                                                                                                               | Word                                     | 既有字段 | 既有邻居 | 决定                                                 |
|:----------------------------------------------------------------------------------------------------------------------|:-----------------------------------------|:---------|:---------|:-----------------------------------------------------|
| [abandoned-workflow](../entries/abandoned-workflow.yml)                                                               | abandoned                                | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [absolute-positioning-css](../entries/absolute-positioning-css.yml)                                                   | absolute positioning                     | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [abstract-syntax-tree-compiler](../entries/abstract-syntax-tree-compiler.yml)                                         | abstract syntax tree                     | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [acana-pet-food](../entries/acana-pet-food.yml)                                                                       | ACANA                                    | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [accented-text-tech](../entries/accented-text-tech.yml)                                                               | accented text                            | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [access-log-nginx](../entries/access-log-nginx.yml)                                                                   | access log                               | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [accessible-name-html](../entries/accessible-name-html.yml)                                                           | accessible name                          | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [accessor-tech](../entries/accessor-tech.yml)                                                                         | accessor                                 | WF       | 1        | 保留现有辅助与关联                                   |
| [accordion-bootstrap](../entries/accordion-bootstrap.yml)                                                             | accordion                                | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [account-finance](../entries/account-finance.yml)                                                                     | account                                  | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [achieve-general](../entries/achieve-general.yml)                                                                     | achieve                                  | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [acquire-workflow](../entries/acquire-workflow.yml)                                                                   | acquire                                  | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [actionable-tech](../entries/actionable-tech.yml)                                                                     | actionable                               | WF       | 0        | 保留现有辅助与关联                                   |
| [activate-general](../entries/activate-general.yml)                                                                   | activate                                 | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [actual-behavior-workflow](../entries/actual-behavior-workflow.yml)                                                   | actual behavior                          | —        | 1        | 保留现有辅助与关联                                   |
| [actual-general](../entries/actual-general.yml)                                                                       | actual                                   | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [adapter-tech](../entries/adapter-tech.yml)                                                                           | adapter                                  | WF       | 0        | 保留现有辅助与关联                                   |
| [adjacent-tech](../entries/adjacent-tech.yml)                                                                         | adjacent                                 | ETY      | 0        | 保留现有辅助与关联                                   |
| [adventure-general](../entries/adventure-general.yml)                                                                 | adventure                                | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [adventure-log-zelda](../entries/adventure-log-zelda.yml)                                                             | Adventure Log                            | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [aerocuda-zelda](../entries/aerocuda-zelda.yml)                                                                       | Aerocuda                                 | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [agent-ai](../entries/agent-ai.yml)                                                                                   | agent                                    | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [aggressive-combat](../entries/aggressive-combat.yml)                                                                 | aggressive                               | —        | 1        | 保留现有辅助与关联                                   |
| [aim-general](../entries/aim-general.yml)                                                                             | aim                                      | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [airborne-combat](../entries/airborne-combat.yml)                                                                     | airborne                                 | WF       | 0        | 保留现有辅助与关联                                   |
| [alberta-province](../entries/alberta-province.yml)                                                                   | Alberta                                  | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [album-ui](../entries/album-ui.yml)                                                                                   | album                                    | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [alert-bootstrap](../entries/alert-bootstrap.yml)                                                                     | alert                                    | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [alias-nginx](../entries/alias-nginx.yml)                                                                             | alias                                    | —        | 0        | R04                                                  |
| [alignment-css](../entries/alignment-css.yml)                                                                         | alignment                                | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [allowlist-tech](../entries/allowlist-tech.yml)                                                                       | allowlist                                | WF       | 0        | 保留现有辅助与关联                                   |
| [altar-architecture](../entries/altar-architecture.yml)                                                               | altar                                    | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [alternate-tech](../entries/alternate-tech.yml)                                                                       | alternate                                | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [alternative-text-html](../entries/alternative-text-html.yml)                                                         | alternative text                         | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [ambitious-general](../entries/ambitious-general.yml)                                                                 | ambitious                                | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [amend-git](../entries/amend-git.yml)                                                                                 | amend                                    | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [ancestor-general](../entries/ancestor-general.yml)                                                                   | ancestor                                 | —        | 1        | 保留现有辅助与关联                                   |
| [ancient-arowana-zelda](../entries/ancient-arowana-zelda.yml)                                                         | Ancient Arowana                          | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [ancient-general](../entries/ancient-general.yml)                                                                     | ancient                                  | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [animation-css](../entries/animation-css.yml)                                                                         | animation                                | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [annotation-tech](../entries/annotation-tech.yml)                                                                     | annotation                               | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [annual-percentage-yield-finance](../entries/annual-percentage-yield-finance.yml)                                     | annual percentage yield                  | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [anonymous-tech](../entries/anonymous-tech.yml)                                                                       | anonymous                                | ETY      | 0        | 保留现有辅助与关联                                   |
| [any-typescript](../entries/any-typescript.yml)                                                                       | any                                      | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [append-go](../entries/append-go.yml)                                                                                 | append                                   | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [apple-food](../entries/apple-food.yml)                                                                               | apple                                    | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [applicable-tech](../entries/applicable-tech.yml)                                                                     | applicable                               | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [appreciate-general](../entries/appreciate-general.yml)                                                               | appreciate                               | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [approval-workflow](../entries/approval-workflow.yml)                                                                 | approval                                 | —        | 1        | 保留现有辅助与关联                                   |
| [aquatic-combat](../entries/aquatic-combat.yml)                                                                       | aquatic                                  | ETY      | 0        | 保留现有辅助与关联                                   |
| [archipelago-terrain](../entries/archipelago-terrain.yml)                                                             | archipelago                              | ETY      | 0        | 保留现有辅助与关联                                   |
| [argument-tech](../entries/argument-tech.yml)                                                                         | argument                                 | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [armor-equipment](../entries/armor-equipment.yml)                                                                     | armor                                    | —        | 1        | 保留现有辅助与关联                                   |
| [armored-combat](../entries/armored-combat.yml)                                                                       | armored                                  | WF       | 0        | 保留现有辅助与关联                                   |
| [armory-architecture](../entries/armory-architecture.yml)                                                             | armory                                   | —        | 1        | 保留现有辅助与关联                                   |
| [arowana-general](../entries/arowana-general.yml)                                                                     | arowana                                  | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [array-hole-javascript](../entries/array-hole-javascript.yml)                                                         | array hole                               | —        | 1        | 保留现有辅助与关联                                   |
| [array-spread-tech](../entries/array-spread-tech.yml)                                                                 | array spread                             | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [array-tech](../entries/array-tech.yml)                                                                               | array                                    | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [arrow-weapon](../entries/arrow-weapon.yml)                                                                           | arrow                                    | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [artifact-tech](../entries/artifact.yml)                                                                              | artifact                                 | ETY / N  | 1        | 保留现有辅助与关联                                   |
| [artifact-workflow](../entries/artifact-workflow.yml)                                                                 | artifact                                 | N        | 1        | 保留现有辅助与关联                                   |
| [as-proof-of-workflow](../entries/as-proof-of-workflow.yml)                                                           | as proof of                              | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [ascend-general](../entries/ascend-general.yml)                                                                       | ascend                                   | —        | 1        | 保留现有辅助与关联                                   |
| [ask-price-finance](../entries/ask-price-finance.yml)                                                                 | ask price                                | —        | 1        | 保留现有辅助与关联                                   |
| [assemble-general](../entries/assemble-general.yml)                                                                   | assemble                                 | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [assertion-testing](../entries/assertion-testing.yml)                                                                 | assertion                                | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [asset-finance](../entries/asset-finance.yml)                                                                         | asset                                    | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [assignability-typescript](../entries/assignability-typescript.yml)                                                   | assignability                            | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [assignment-target-tech](../entries/assignment-target-tech.yml)                                                       | assignment target                        | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [async-await-javascript](../entries/async-await-javascript.yml)                                                       | async/await                              | —        | 1        | 保留现有辅助与关联                                   |
| [asynchronous-tech](../entries/asynchronous-tech.yml)                                                                 | asynchronous                             | WF       | 0        | 保留现有辅助与关联                                   |
| [at-most-workflow](../entries/at-most-workflow.yml)                                                                   | at most                                  | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [at-rule-css](../entries/at-rule-css.yml)                                                                             | at-rule                                  | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [atomically-workflow](../entries/atomically-workflow.yml)                                                             | atomically                               | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [attach-general](../entries/attach-general.yml)                                                                       | attach                                   | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [attack-power-combat](../entries/attack-power-combat.yml)                                                             | attack power                             | WF       | 0        | 保留现有辅助与关联                                   |
| [attack-up-zelda](../entries/attack-up-zelda.yml)                                                                     | Attack Up                                | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [attribute-html](../entries/attribute-html.yml)                                                                       | attribute                                | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [attribute-selector-css](../entries/attribute-selector-css.yml)                                                       | attribute selector                       | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [attribute-to-general](../entries/attribute-to-general.yml)                                                           | attribute to                             | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [audio-general](../entries/audio-general.yml)                                                                         | audio                                    | ETY      | 2        | 保留现有辅助与关联                                   |
| [audit-tech](../entries/audit-tech.yml)                                                                               | audit                                    | ETY      | 1        | 保留现有辅助与关联                                   |
| [auto-save-game](../entries/auto-save-game.yml)                                                                       | auto-save                                | —        | 1        | 保留现有辅助与关联                                   |
| [autobuild-zelda](../entries/autobuild-zelda.yml)                                                                     | Autobuild                                | WF       | 0        | 保留现有辅助与关联                                   |
| [autocomplete-html](../entries/autocomplete-html.yml)                                                                 | autocomplete                             | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [await-review-workflow](../entries/await-review-workflow.yml)                                                         | await review                             | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [awaken-story](../entries/awaken-story.yml)                                                                           | awaken                                   | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [axe-weapon](../entries/axe-weapon.yml)                                                                               | axe                                      | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [backdrop-bootstrap](../entries/backdrop-bootstrap.yml)                                                               | backdrop                                 | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [backend-server-nginx](../entries/backend-server-nginx.yml)                                                           | backend server                           | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [backend-tech](../entries/backend-tech.yml)                                                                           | backend                                  | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [backflip-movement](../entries/backflip-movement.yml)                                                                 | backflip                                 | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [backing-array-go](../entries/backing-array-go.yml)                                                                   | backing array                            | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [backtracking-algorithm](../entries/backtracking-algorithm.yml)                                                       | backtracking                             | WF       | 1        | 保留现有辅助与关联                                   |
| [badge-bootstrap](../entries/badge-bootstrap.yml)                                                                     | badge                                    | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [balloon-engineering](../entries/balloon-engineering.yml)                                                             | balloon                                  | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [banana-food](../entries/banana-food.yml)                                                                             | banana                                   | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [bandanna-clothing](../entries/bandanna-clothing.yml)                                                                 | bandanna                                 | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [base-image-docker](../entries/base-image-docker.yml)                                                                 | base image                               | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [based-on-workflow](../entries/based-on-workflow.yml)                                                                 | based on                                 | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [baseline-tech](../entries/baseline-tech.yml)                                                                         | baseline                                 | WF / N   | 1        | 保留现有辅助与关联                                   |
| [bass-food](../entries/bass-food.yml)                                                                                 | bass                                     | N        | 0        | 保留现有辅助与关联                                   |
| [batching-react](../entries/batching-react.yml)                                                                       | batching                                 | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [battery-engineering](../entries/battery-engineering.yml)                                                             | battery                                  | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [beam-emitter-zelda](../entries/beam-emitter-zelda.yml)                                                               | Beam Emitter                             | WF       | 0        | 保留现有辅助与关联                                   |
| [beam-engineering](../entries/beam-engineering.yml)                                                                   | beam                                     | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [beat-estimates-finance](../entries/beat-estimates-finance.yml)                                                       | beat estimates                           | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [bee-general](../entries/bee-general.yml)                                                                             | bee                                      | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [beef-food](../entries/beef-food.yml)                                                                                 | beef                                     | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [before-and-after-the-fix-workflow](../entries/before-and-after-the-fix-workflow.yml)                                 | before and after the fix                 | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [before-drawing-conclusions-workflow](../entries/before-drawing-conclusions-workflow.yml)                             | before drawing conclusions               | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [benchmark-tech](../entries/benchmark-tech.yml)                                                                       | benchmark                                | —        | 1        | 保留现有辅助与关联                                   |
| [bid-price-finance](../entries/bid-price-finance.yml)                                                                 | bid price                                | N        | 1        | 保留现有辅助与关联                                   |
| [big-battery-zelda](../entries/big-battery-zelda.yml)                                                                 | Big Battery                              | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [big-wheel-zelda](../entries/big-wheel-zelda.yml)                                                                     | Big Wheel                                | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [bigint-javascript](../entries/bigint-javascript.yml)                                                                 | BigInt                                   | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [binary-search-algorithm](../entries/binary-search-algorithm.yml)                                                     | binary search                            | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [bind-mount-docker](../entries/bind-mount-docker.yml)                                                                 | bind mount                               | —        | 0        | R02、N02                                             |
| [binding-tech](../entries/binding-tech.yml)                                                                           | binding                                  | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [black-pepper-food](../entries/black-pepper-food.yml)                                                                 | black pepper                             | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [blessing-general](../entries/blessing-general.yml)                                                                   | blessing                                 | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [block-formatting-context-css](../entries/block-formatting-context-css.yml)                                           | block formatting context                 | N        | 0        | 保留现有辅助与关联                                   |
| [block-general](../entries/block-general.yml)                                                                         | block                                    | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [block-level-box-css](../entries/block-level-box-css.yml)                                                             | block-level box                          | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [blocker-workflow](../entries/blocker-workflow.yml)                                                                   | blocker                                  | WF       | 0        | 保留现有辅助与关联                                   |
| [bokoblin-zelda](../entries/bokoblin-zelda.yml)                                                                       | Bokoblin                                 | N        | 0        | 保留现有辅助与关联                                   |
| [bomb-weapon](../entries/bomb-weapon.yml)                                                                             | bomb                                     | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [bool-go](../entries/bool-go.yml)                                                                                     | bool                                     | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [boolean-attribute-html](../entries/boolean-attribute-html.yml)                                                       | boolean attribute                        | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [boolean-typescript](../entries/boolean-typescript.yml)                                                               | boolean                                  | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [boomerang-weapon](../entries/boomerang-weapon.yml)                                                                   | boomerang                                | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [boots-clothing](../entries/boots-clothing.yml)                                                                       | boots                                    | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [bootstrap-framework](../entries/bootstrap-framework.yml)                                                             | Bootstrap                                | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [border-css](../entries/border-css.yml)                                                                               | border                                   | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [boss-bokoblin-zelda](../entries/boss-bokoblin-zelda.yml)                                                             | Boss Bokoblin                            | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [bottomless-general](../entries/bottomless-general.yml)                                                               | bottomless                               | WF       | 0        | 保留现有辅助与关联                                   |
| [bounded-tech](../entries/bounded-tech.yml)                                                                           | bounded                                  | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [bow-weapon](../entries/bow-weapon.yml)                                                                               | bow                                      | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [box-model-css](../entries/box-model-css.yml)                                                                         | box model                                | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [box-sizing-css](../entries/box-sizing-css.yml)                                                                       | box-sizing                               | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [braised-food](../entries/braised-food.yml)                                                                           | braised                                  | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [branch-git](../entries/branch-git.yml)                                                                               | branch                                   | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [brave-general](../entries/brave-general.yml)                                                                         | brave                                    | —        | 2        | 保留现有辅助与关联                                   |
| [breadcrumb-bootstrap](../entries/breadcrumb-bootstrap.yml)                                                           | breadcrumb                               | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [breadth-first-tech](../entries/breadth-first-tech.yml)                                                               | breadth-first                            | WF       | 0        | 保留现有辅助与关联                                   |
| [breakpoint-css](../entries/breakpoint-css.yml)                                                                       | breakpoint                               | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [bridge-architecture](../entries/bridge-architecture.yml)                                                             | bridge                                   | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [bridge-tech](../entries/bridge-tech.yml)                                                                             | bridge                                   | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [brightbloom-seed-zelda](../entries/brightbloom-seed-zelda.yml)                                                       | Brightbloom Seed                         | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [brightcap-zelda](../entries/brightcap-zelda.yml)                                                                     | Brightcap                                | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [broadsword-weapon](../entries/broadsword-weapon.yml)                                                                 | broadsword                               | WF       | 0        | 保留现有辅助与关联                                   |
| [buffered-channel-go](../entries/buffered-channel-go.yml)                                                             | buffered channel                         | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [bug-tech](../entries/bug-tech.yml)                                                                                   | bug                                      | —        | 0        | ETY03                                                |
| [build-cache-docker](../entries/build-cache-docker.yml)                                                               | build cache                              | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [build-construction](../entries/build-construction.yml)                                                               | build                                    | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [build-context-docker](../entries/build-context-docker.yml)                                                           | build context                            | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [build-tech](../entries/build-tech.yml)                                                                               | build                                    | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [butter-food](../entries/butter-food.yml)                                                                             | butter                                   | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [button-group-bootstrap](../entries/button-group-bootstrap.yml)                                                       | button group                             | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [buying-power-finance](../entries/buying-power-finance.yml)                                                           | buying power                             | —        | 1        | 保留现有辅助与关联                                   |
| [bypass-tech](../entries/bypass-tech.yml)                                                                             | bypass                                   | —        | 1        | 保留现有辅助与关联                                   |
| [byte-go](../entries/byte-go.yml)                                                                                     | byte                                     | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [byte-offset-tech](../entries/byte-offset-tech.yml)                                                                   | byte offset                              | N        | 1        | 保留现有辅助与关联                                   |
| [cable-hardware](../entries/cable-hardware.yml)                                                                       | cable                                    | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [cache-invalidation-tech](../entries/cache-invalidation-tech.yml)                                                     | cache invalidation                       | —        | 0        | R01、N01                                             |
| [cache-revalidation-web](../entries/cache-revalidation-web.yml)                                                       | cache revalidation                       | —        | 0        | R01 反向、WF02                                       |
| [cache-tech](../entries/cache-tech.yml)                                                                               | cache                                    | ETY      | 1        | 保留现有辅助与关联                                   |
| [calamity-story](../entries/calamity-story.yml)                                                                       | calamity                                 | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [call-option-finance](../entries/call-option-finance.yml)                                                             | call option                              | —        | 1        | 保留现有辅助与关联                                   |
| [call-stack-javascript](../entries/call-stack-javascript.yml)                                                         | call stack                               | —        | 2        | 保留现有辅助与关联                                   |
| [callback-tech](../entries/callback-tech.yml)                                                                         | callback                                 | WF       | 0        | 保留现有辅助与关联                                   |
| [campfire-general](../entries/campfire-general.yml)                                                                   | campfire                                 | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [canada-country](../entries/canada-country.yml)                                                                       | Canada                                   | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [cancel-general](../entries/cancel-general.yml)                                                                       | cancel                                   | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [cancellation-tech](../entries/cancellation-tech.yml)                                                                 | cancellation                             | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [candidate-tech](../entries/candidate-tech.yml)                                                                       | candidate                                | ETY      | 0        | 保留现有辅助与关联                                   |
| [canine-adjective](../entries/canine-adjective.yml)                                                                   | canine                                   | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [cannon-engineering](../entries/cannon-engineering.yml)                                                               | cannon                                   | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [canyon-terrain](../entries/canyon-terrain.yml)                                                                       | canyon                                   | —        | 1        | 保留现有辅助与关联                                   |
| [capacity-go](../entries/capacity-go.yml)                                                                             | capacity                                 | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [capsule-general](../entries/capsule-general.yml)                                                                     | capsule                                  | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [captain-construct-zelda](../entries/captain-construct-zelda.yml)                                                     | Captain Construct                        | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [caption-html](../entries/caption-html.yml)                                                                           | caption                                  | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [card-bootstrap](../entries/card-bootstrap.yml)                                                                       | card                                     | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [carousel-bootstrap](../entries/carousel-bootstrap.yml)                                                               | carousel                                 | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [carrot-food](../entries/carrot-food.yml)                                                                             | carrot                                   | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [cart-engineering](../entries/cart-engineering.yml)                                                                   | cart                                     | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [cascade-css](../entries/cascade-css.yml)                                                                             | cascade                                  | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [cascade-layer-css](../entries/cascade-layer-css.yml)                                                                 | cascade layer                            | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [cash-account-finance](../entries/cash-account-finance.yml)                                                           | cash account                             | —        | 1        | 保留现有辅助与关联                                   |
| [cash-finance](../entries/cash-finance.yml)                                                                           | cash                                     | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [cash-sweep-finance](../entries/cash-sweep-finance.yml)                                                               | cash sweep                               | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [cast-tech](../entries/cast-tech.yml)                                                                                 | cast                                     | —        | 1        | 保留现有辅助与关联                                   |
| [castle-architecture](../entries/castle-architecture.yml)                                                             | castle                                   | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [catch-handler-tech](../entries/catch-handler-tech.yml)                                                               | catch handler                            | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [cave-architecture](../entries/cave-architecture.yml)                                                                 | cave                                     | —        | 1        | 保留现有辅助与关联                                   |
| [caveat-general](../entries/caveat-general.yml)                                                                       | caveat                                   | ETY      | 0        | 保留现有辅助与关联                                   |
| [cavern-architecture](../entries/cavern-architecture.yml)                                                             | cavern                                   | N        | 1        | 保留现有辅助与关联                                   |
| [ceiling-architecture](../entries/ceiling-architecture.yml)                                                           | ceiling                                  | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [chamber-architecture](../entries/chamber-architecture.yml)                                                           | chamber                                  | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [champion-general](../entries/champion-general.yml)                                                                   | champion                                 | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [channel-go](../entries/channel-go.yml)                                                                               | channel                                  | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [character-encoding-html](../entries/character-encoding-html.yml)                                                     | character encoding                       | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [character-profiles-zelda](../entries/character-profiles-zelda.yml)                                                   | Character Profiles                       | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [character-reference-html](../entries/character-reference-html.yml)                                                   | character reference                      | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [charge-electricity](../entries/charge-electricity.yml)                                                               | charge                                   | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [chasm-terrain](../entries/chasm-terrain.yml)                                                                         | chasm                                    | N        | 1        | 保留现有辅助与关联                                   |
| [checkpoint-tech](../entries/checkpoint-tech.yml)                                                                     | checkpoint                               | —        | 1        | 保留现有辅助与关联                                   |
| [cheese-food](../entries/cheese-food.yml)                                                                             | cheese                                   | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [chicken-food](../entries/chicken-food.yml)                                                                           | chicken                                  | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [children-react](../entries/children-react.yml)                                                                       | children                                 | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [chilly-general](../entries/chilly-general.yml)                                                                       | chilly                                   | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [chuchu-zelda](../entries/chuchu-zelda.yml)                                                                           | Chuchu                                   | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [citadel-architecture](../entries/citadel-architecture.yml)                                                           | citadel                                  | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [civilization-general](../entries/civilization-general.yml)                                                           | civilization                             | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [claim-general](../entries/claim-general.yml)                                                                         | claim                                    | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [clarify-general](../entries/clarify-general.yml)                                                                     | clarify                                  | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [class-selector-css](../entries/class-selector-css.yml)                                                               | class selector                           | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [claude-ai](../entries/claude-ai.yml)                                                                                 | Claude                                   | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [clause-tech](../entries/clause-tech.yml)                                                                             | clause                                   | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [claw-general](../entries/claw-general.yml)                                                                           | claw                                     | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [claymore-weapon](../entries/claymore-weapon.yml)                                                                     | claymore                                 | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [cleanup-tech](../entries/cleanup-tech.yml)                                                                           | cleanup                                  | WF       | 0        | 保留现有辅助与关联                                   |
| [cliff-terrain](../entries/cliff-terrain.yml)                                                                         | cliff                                    | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [climb-general](../entries/climb-general.yml)                                                                         | climb                                    | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [climbing-gear-zelda](../entries/climbing-gear-zelda.yml)                                                             | Climbing Gear                            | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [closed-general](../entries/closed-general.yml)                                                                       | closed                                   | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [closing-price-finance](../entries/closing-price-finance.yml)                                                         | closing price                            | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [closure-tech](../entries/closure-tech.yml)                                                                           | closure                                  | N        | 1        | 保留现有辅助与关联                                   |
| [club-weapon](../entries/club-weapon.yml)                                                                             | club                                     | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [clue-story](../entries/clue-story.yml)                                                                               | clue                                     | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [code-generation-tech](../entries/code-generation-tech.yml)                                                           | code generation                          | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [code-unit-tech](../entries/code-unit-tech.yml)                                                                       | code unit                                | —        | 1        | 保留现有辅助与关联                                   |
| [codex-ai](../entries/codex-ai.yml)                                                                                   | Codex                                    | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [coerce-tech](../entries/coerce-tech.yml)                                                                             | coerce                                   | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [cold-resistance-combat](../entries/cold-resistance-combat.yml)                                                       | cold resistance                          | —        | 1        | 保留现有辅助与关联                                   |
| [colgera-zelda](../entries/colgera-zelda.yml)                                                                         | Colgera                                  | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [collapse-bootstrap](../entries/collapse-bootstrap.yml)                                                               | collapse                                 | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [collectible-game](../entries/collectible-game.yml)                                                                   | collectible                              | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [collision-tech](../entries/collision-tech.yml)                                                                       | collision                                | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [color-mode-bootstrap](../entries/color-mode-bootstrap.yml)                                                           | color mode                               | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [colossal-combat](../entries/colossal-combat.yml)                                                                     | colossal                                 | ETY      | 0        | 保留现有辅助与关联                                   |
| [column-bootstrap](../entries/column-bootstrap.yml)                                                                   | column                                   | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [column-offset-bootstrap](../entries/column-offset-bootstrap.yml)                                                     | column offset                            | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [column-span-html](../entries/column-span-html.yml)                                                                   | column span                              | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [combinator-css](../entries/combinator-css.yml)                                                                       | combinator                               | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [combining-mark-tech](../entries/combining-mark-tech.yml)                                                             | combining mark                           | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [command-cli](../entries/command-cli.yml)                                                                             | command                                  | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [commit-git](../entries/commit-git.yml)                                                                               | commit                                   | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [commit-phase-react](../entries/commit-phase-react.yml)                                                               | commit phase                             | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [community-general](../entries/community-general.yml)                                                                 | community                                | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [compact-context-tech](../entries/compact-context-tech.yml)                                                           | compact                                  | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [comparable-go](../entries/comparable-go.yml)                                                                         | comparable                               | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [compendium-general](../entries/compendium-general.yml)                                                               | compendium                               | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [competing-work-workflow](../entries/competing-work-workflow.yml)                                                     | competing work                           | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [compile-tech](../entries/compile-tech.yml)                                                                           | compile                                  | —        | 1        | 保留现有辅助与关联                                   |
| [compiler-tech](../entries/compiler-tech.yml)                                                                         | compiler                                 | WF       | 1        | 保留现有辅助与关联                                   |
| [complete-workflow](../entries/complete-workflow.yml)                                                                 | complete                                 | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [component-identity-react](../entries/component-identity-react.yml)                                                   | component identity                       | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [component-react](../entries/component-react.yml)                                                                     | component                                | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [compose-yaml-docker](../entries/compose-yaml-docker.yml)                                                             | compose.yaml                             | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [compound-assignment-tech](../entries/compound-assignment-tech.yml)                                                   | compound assignment                      | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [computed-key-tech](../entries/computed-key-tech.yml)                                                                 | computed key                             | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [concatenation-tech](../entries/concatenation-tech.yml)                                                               | concatenation                            | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [concise-general](../entries/concise-general.yml)                                                                     | concise                                  | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [concrete-tech](../entries/concrete-tech.yml)                                                                         | concrete                                 | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [conditional-rendering-react](../entries/conditional-rendering-react.yml)                                             | conditional rendering                    | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [conditional-type-typescript](../entries/conditional-type-typescript.yml)                                             | conditional type                         | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [configuration-tech](../entries/configuration-tech.yml)                                                               | configuration                            | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [configuration-test-nginx](../entries/configuration-test-nginx.yml)                                                   | configuration test                       | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [confirmed-tech](../entries/confirmed-tech.yml)                                                                       | confirmed                                | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [connect-general](../entries/connect-general.yml)                                                                     | connect                                  | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [connector-hardware](../entries/connector-hardware.yml)                                                               | connector                                | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [consistent-tech](../entries/consistent-tech.yml)                                                                     | consistent                               | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [constant-go](../entries/constant-go.yml)                                                                             | constant                                 | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [constant-propagation-tech](../entries/constant-propagation-tech.yml)                                                 | constant propagation                     | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [constraint-validation-html](../entries/constraint-validation-html.yml)                                               | constraint validation                    | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [construct-being-zelda](../entries/construct-being-zelda.yml)                                                         | Construct                                | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [construct-factory-zelda](../entries/construct-factory-zelda.yml)                                                     | Construct Factory                        | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [construct-general](../entries/construct-general.yml)                                                                 | construct                                | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [construct-head-zelda](../entries/construct-head-zelda.yml)                                                           | Construct Head                           | WF       | 0        | 保留现有辅助与关联                                   |
| [consume-tech](../entries/consume-tech.yml)                                                                           | consume                                  | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [container-bootstrap](../entries/container-bootstrap.yml)                                                             | container                                | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [container-docker](../entries/container-docker.yml)                                                                   | container                                | —        | 0        | R03 反向                                             |
| [container-query-css](../entries/container-query-css.yml)                                                             | container query                          | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [containing-block-css](../entries/containing-block-css.yml)                                                           | containing block                         | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [content-box-css](../entries/content-box-css.yml)                                                                     | content box                              | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [content-model-html](../entries/content-model-html.yml)                                                               | content model                            | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [context-language](../entries/context-language.yml)                                                                   | context                                  | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [context-provider-react](../entries/context-provider-react.yml)                                                       | context provider                         | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [context-react](../entries/context-react.yml)                                                                         | context                                  | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [context-window-ai](../entries/context-window-ai.yml)                                                                 | context window                           | —        | 0        | R05                                                  |
| [contextual-type-tech](../entries/contextual-type-tech.yml)                                                           | contextual type                          | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [contribution-general](../entries/contribution-general.yml)                                                           | contribution                             | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [controlled-component-react](../entries/controlled-component-react.yml)                                               | controlled component                     | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [cook-cooking](../entries/cook-cooking.yml)                                                                           | cook                                     | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [cooking-pot-cooking](../entries/cooking-pot-cooking.yml)                                                             | cooking pot                              | WF       | 0        | 保留现有辅助与关联                                   |
| [coordinates-ui](../entries/coordinates-ui.yml)                                                                       | coordinates                              | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [copilot-ai](../entries/copilot-ai.yml)                                                                               | Copilot                                  | WF       | 0        | 保留现有辅助与关联                                   |
| [copy-tech](../entries/copy-tech.yml)                                                                                 | copy                                     | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [core-general](../entries/core-general.yml)                                                                           | core                                     | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [correctness-tech](../entries/correctness-tech.yml)                                                                   | correctness                              | WF       | 1        | 保留现有辅助与关联                                   |
| [corresponding-general](../entries/corresponding-general.yml)                                                         | corresponding                            | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [corridor-architecture](../entries/corridor-architecture.yml)                                                         | corridor                                 | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [corrupt-tech](../entries/corrupt-tech.yml)                                                                           | corrupt                                  | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [corrupted-combat](../entries/corrupted-combat.yml)                                                                   | corrupted                                | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [cost-finance](../entries/cost-finance.yml)                                                                           | cost                                     | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [counter-tech](../entries/counter-tech.yml)                                                                           | counter                                  | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [courageous-general](../entries/courageous-general.yml)                                                               | courageous                               | WF / ETY | 2        | 保留现有辅助与关联                                   |
| [courser-bee-honey-zelda](../entries/courser-bee-honey-zelda.yml)                                                     | Courser Bee Honey                        | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [coverage-tech](../entries/coverage-tech.yml)                                                                         | coverage                                 | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [crash-tech](../entries/crash-tech.yml)                                                                               | crash                                    | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [crater-terrain](../entries/crater-terrain.yml)                                                                       | crater                                   | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [crisis-story](../entries/crisis-story.yml)                                                                           | crisis                                   | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [critical-hit-combat](../entries/critical-hit-combat.yml)                                                             | critical hit                             | WF       | 0        | 保留现有辅助与关联                                   |
| [cross-axis-css](../entries/cross-axis-css.yml)                                                                       | cross axis                               | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [cross-functional-collaboration-workflow](../entries/cross-functional-collaboration-workflow.yml)                     | cross-functional collaboration           | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [cross-origin-resource-sharing-web](../entries/cross-origin-resource-sharing-web.yml)                                 | cross-origin resource sharing            | —        | 1        | 保留现有辅助与关联                                   |
| [crouch-general](../entries/crouch-general.yml)                                                                       | crouch                                   | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [crypto-finance](../entries/crypto-finance.yml)                                                                       | crypto                                   | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [crystal-general](../entries/crystal-general.yml)                                                                     | crystal                                  | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [crystal-refinery-zelda](../entries/crystal-refinery-zelda.yml)                                                       | Crystal Refinery                         | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [crystallize-general](../entries/crystallize-general.yml)                                                             | crystallize                              | WF       | 1        | 保留现有辅助与关联                                   |
| [crystallized-charge-zelda](../entries/crystallized-charge-zelda.yml)                                                 | Crystallized Charge                      | —        | 1        | 保留现有辅助与关联                                   |
| [crystallized-general](../entries/crystallized-general.yml)                                                           | crystallized                             | WF       | 1        | 保留现有辅助与关联                                   |
| [css-language](../entries/css-language.yml)                                                                           | CSS                                      | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [cunning-general](../entries/cunning-general.yml)                                                                     | cunning                                  | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [curious-general](../entries/curious-general.yml)                                                                     | curious                                  | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [currency-finance](../entries/currency-finance.yml)                                                                   | currency                                 | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [curry-food](../entries/curry-food.yml)                                                                               | curry                                    | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [cursor-ai](../entries/cursor-ai.yml)                                                                                 | Cursor                                   | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [custom-hook-react](../entries/custom-hook-react.yml)                                                                 | custom Hook                              | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [custom-property-css](../entries/custom-property-css.yml)                                                             | custom property                          | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [dagger-weapon](../entries/dagger-weapon.yml)                                                                         | dagger                                   | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [damage-combat](../entries/damage-combat.yml)                                                                         | damage                                   | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [data-attribute-bootstrap](../entries/data-attribute-bootstrap.yml)                                                   | data attribute                           | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [data-cell-html](../entries/data-cell-html.yml)                                                                       | data cell                                | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [data-race-go](../entries/data-race-go.yml)                                                                           | data race                                | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [data-table-html](../entries/data-table-html.yml)                                                                     | data table                               | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [deadline-go](../entries/deadline-go.yml)                                                                             | deadline                                 | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [deadlock-go](../entries/deadlock-go.yml)                                                                             | deadlock                                 | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [death-mountain-zelda](../entries/death-mountain-zelda.yml)                                                           | Death Mountain                           | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [debounce-javascript](../entries/debounce-javascript.yml)                                                             | debounce                                 | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [decayed-general](../entries/decayed-general.yml)                                                                     | decayed                                  | —        | 1        | 保留现有辅助与关联                                   |
| [declaration-css](../entries/declaration-css.yml)                                                                     | declaration                              | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [declaration-emit-tech](../entries/declaration-emit-tech.yml)                                                         | declaration emit                         | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [declaration-file-typescript](../entries/declaration-file-typescript.yml)                                             | declaration file                         | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [decode-tech](../entries/decode-tech.yml)                                                                             | decode                                   | N        | 1        | 保留现有辅助与关联                                   |
| [decompose-general](../entries/decompose-general.yml)                                                                 | decompose                                | —        | 1        | 保留现有辅助与关联                                   |
| [decorative-image-html](../entries/decorative-image-html.yml)                                                         | decorative image                         | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [decorator-tech](../entries/decorator-tech.yml)                                                                       | decorator                                | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [deduplicate-tech](../entries/deduplicate-tech.yml)                                                                   | deduplicate                              | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [deep-copy-javascript](../entries/deep-copy-javascript.yml)                                                           | deep copy                                | —        | 1        | 保留现有辅助与关联                                   |
| [deep-general](../entries/deep-general.yml)                                                                           | deep                                     | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [deepseek-ai](../entries/deepseek-ai.yml)                                                                             | DeepSeek                                 | WF       | 0        | 保留现有辅助与关联                                   |
| [defense-combat](../entries/defense-combat.yml)                                                                       | defense                                  | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [defense-up-zelda](../entries/defense-up-zelda.yml)                                                                   | Defense Up                               | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [defer-go](../entries/defer-go.yml)                                                                                   | defer                                    | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [defer-tech](../entries/defer-tech.yml)                                                                               | defer                                    | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [deity-story](../entries/deity-story.yml)                                                                             | deity                                    | —        | 1        | 保留现有辅助与关联                                   |
| [delegation-workflow](../entries/delegation-workflow.yml)                                                             | delegation                               | —        | 1        | 保留现有辅助与关联                                   |
| [demon-dragon-zelda](../entries/demon-dragon-zelda.yml)                                                               | Demon Dragon                             | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [demon-king-ganondorf-zelda](../entries/demon-king-ganondorf-zelda.yml)                                               | Demon King Ganondorf                     | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [demon-king-zelda](../entries/demon-king-zelda.yml)                                                                   | Demon King                               | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [demon-kings-bow-zelda](../entries/demon-kings-bow-zelda.yml)                                                         | Demon King's Bow                         | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [demon-story](../entries/demon-story.yml)                                                                             | demon                                    | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [dependency-array-react](../entries/dependency-array-react.yml)                                                       | dependency array                         | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [dependency-go](../entries/dependency-go.yml)                                                                         | dependency                               | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [deposit-finance](../entries/deposit-finance.yml)                                                                     | deposit                                  | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [deprecated-tech](../entries/deprecated.yml)                                                                          | deprecated                               | N        | 0        | 保留现有辅助与关联                                   |
| [depths-terrain](../entries/depths-terrain.yml)                                                                       | depths                                   | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [dereference-tech](../entries/dereference-tech.yml)                                                                   | dereference                              | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [derive-tech](../entries/derive-tech.yml)                                                                             | derive                                   | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [derived-state-react](../entries/derived-state-react.yml)                                                             | derived state                            | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [descend-general](../entries/descend-general.yml)                                                                     | descend                                  | —        | 2        | 保留现有辅助与关联                                   |
| [descendant-general](../entries/descendant-general.yml)                                                               | descendant                               | WF       | 2        | 保留现有辅助与关联                                   |
| [descent-flight](../entries/descent-flight.yml)                                                                       | descent                                  | WF       | 1        | 保留现有辅助与关联                                   |
| [desert-terrain](../entries/desert-terrain.yml)                                                                       | desert                                   | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [destination-ui](../entries/destination-ui.yml)                                                                       | destination                              | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [destroy-general](../entries/destroy-general.yml)                                                                     | destroy                                  | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [detach-general](../entries/detach-general.yml)                                                                       | detach                                   | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [detach-git](../entries/detach-git.yml)                                                                               | detach                                   | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [detached-mode-docker](../entries/detached-mode-docker.yml)                                                           | detached mode                            | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [determined-general](../entries/determined-general.yml)                                                               | determined                               | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [detour-general](../entries/detour-general.yml)                                                                       | detour                                   | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [device-dispenser-zelda](../entries/device-dispenser-zelda.yml)                                                       | Device Dispenser                         | WF       | 0        | 保留现有辅助与关联                                   |
| [device-engineering](../entries/device-engineering.yml)                                                               | device                                   | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [diagnostic-context-tech](../entries/diagnostic-context-tech.yml)                                                     | diagnostic context                       | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [diagnostic-tech](../entries/diagnostic-tech.yml)                                                                     | diagnostic                               | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [diet-food](../entries/diet-food.yml)                                                                                 | diet                                     | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [diligent-general](../entries/diligent-general.yml)                                                                   | diligent                                 | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [directive-nginx](../entries/directive-nginx.yml)                                                                     | directive                                | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [directive-tech](../entries/directive-tech.yml)                                                                       | directive                                | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [disabled-control-html](../entries/disabled-control-html.yml)                                                         | disabled control                         | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [disassemble-general](../entries/disassemble-general.yml)                                                             | disassemble                              | WF       | 1        | 保留现有辅助与关联                                   |
| [discard-workflow](../entries/discard-workflow.yml)                                                                   | discard                                  | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [disclosure-workflow](../entries/disclosure-workflow.yml)                                                             | disclosure                               | N        | 1        | 保留现有辅助与关联                                   |
| [discount-finance](../entries/discount-finance.yml)                                                                   | discount                                 | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [discover-general](../entries/discover-general.yml)                                                                   | discover                                 | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [discriminated-union-typescript](../entries/discriminated-union-typescript.yml)                                       | discriminated union                      | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [disguised-combat](../entries/disguised-combat.yml)                                                                   | disguised                                | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [dispatcher-tech](../entries/dispatcher-tech.yml)                                                                     | dispatcher                               | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [dispense-general](../entries/dispense-general.yml)                                                                   | dispense                                 | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [dispenser-engineering](../entries/dispenser-engineering.yml)                                                         | dispenser                                | WF       | 0        | 保留现有辅助与关联                                   |
| [display-css](../entries/display-css.yml)                                                                             | display                                  | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [display-utility-bootstrap](../entries/display-utility-bootstrap.yml)                                                 | display utility                          | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [distinctive-combat](../entries/distinctive-combat.yml)                                                               | distinctive                              | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [distinguish-tech](../entries/distinguish-tech.yml)                                                                   | distinguish                              | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [dive-general](../entries/dive-general.yml)                                                                           | dive                                     | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [dividend-finance](../entries/dividend-finance.yml)                                                                   | dividend                                 | —        | 0        | ETY02                                                |
| [dock-hardware](../entries/dock-hardware.yml)                                                                         | dock                                     | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [docker-compose-docker](../entries/docker-compose-docker.yml)                                                         | Docker Compose                           | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [docker-platform](../entries/docker-platform.yml)                                                                     | Docker                                   | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [dockerfile-docker](../entries/dockerfile-docker.yml)                                                                 | Dockerfile                               | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [doctype-html](../entries/doctype-html.yml)                                                                           | doctype                                  | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [document-html](../entries/document-html.yml)                                                                         | document                                 | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [document-title-html](../entries/document-title-html.yml)                                                             | document title                           | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [dodge-general](../entries/dodge-general.yml)                                                                         | dodge                                    | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [dom-html](../entries/dom-html.yml)                                                                                   | DOM                                      | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [door-architecture](../entries/door-architecture.yml)                                                                 | door                                     | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [dot-general](../entries/dot-general.yml)                                                                             | dot                                      | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [dots-ai](../entries/dots-ai.yml)                                                                                     | dots                                     | WF       | 0        | 保留现有辅助与关联                                   |
| [draconification-zelda](../entries/draconification-zelda.yml)                                                         | draconification                          | —        | 0        | 暂缓，保留已核实的拼法，命名史暂缓                   |
| [draft-workflow](../entries/draft-workflow.yml)                                                                       | draft                                    | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [dragon-story](../entries/dragon-story.yml)                                                                           | dragon                                   | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [draw-general](../entries/draw-general.yml)                                                                           | draw                                     | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [drop-connection](../entries/drop-connection.yml)                                                                     | drop                                     | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [drop-fall-general](../entries/drop-fall-general.yml)                                                                 | drop                                     | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [drop-general](../entries/drop-general.yml)                                                                           | drop                                     | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [dropdown-bootstrap](../entries/dropdown-bootstrap.yml)                                                               | dropdown                                 | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [durability-equipment](../entries/durability-equipment.yml)                                                           | durability                               | WF       | 2        | 保留现有辅助与关联                                   |
| [durable-general](../entries/durable-general.yml)                                                                     | durable                                  | —        | 1        | 保留现有辅助与关联                                   |
| [dynamic-gating-tech](../entries/dynamic-gating-tech.yml)                                                             | dynamic gating                           | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [eagle-general](../entries/eagle-general.yml)                                                                         | eagle                                    | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [earnings-finance](../entries/earnings-finance.yml)                                                                   | earnings                                 | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [earnings-per-share-finance](../entries/earnings-per-share-finance.yml)                                               | earnings per share                       | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [easing-css](../entries/easing-css.yml)                                                                               | easing                                   | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [ebit-finance](../entries/ebit-finance.yml)                                                                           | EBIT                                     | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [edge-case-tech](../entries/edge-case-tech.yml)                                                                       | edge case                                | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [effect-cleanup-react](../entries/effect-cleanup-react.yml)                                                           | Effect cleanup                           | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [effect-react](../entries/effect-react.yml)                                                                           | effect                                   | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [effect-replay-tech](../entries/effect-replay-tech.yml)                                                               | effect replay                            | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [egg-food](../entries/egg-food.yml)                                                                                   | egg                                      | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [elapsed-time-workflow](../entries/elapsed-time-workflow.yml)                                                         | elapsed time                             | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [electro-cooking](../entries/electro-cooking.yml)                                                                     | Electro                                  | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [element-html](../entries/element-html.yml)                                                                           | element                                  | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [eligible-tech](../entries/eligible-tech.yml)                                                                         | eligible                                 | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [elixir-cooking](../entries/elixir-cooking.yml)                                                                       | elixir                                   | ETY      | 0        | 保留现有辅助与关联                                   |
| [elusive-combat](../entries/elusive-combat.yml)                                                                       | elusive                                  | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [embedding-go](../entries/embedding-go.yml)                                                                           | embedding                                | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [emit-physical](../entries/emit-physical.yml)                                                                         | emit                                     | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [emit-tech](../entries/emit-tech.yml)                                                                                 | emit                                     | —        | 1        | 保留现有辅助与关联                                   |
| [emitter-engineering](../entries/emitter-engineering.yml)                                                             | emitter                                  | WF       | 1        | 保留现有辅助与关联                                   |
| [encode-tech](../entries/encode-tech.yml)                                                                             | encode                                   | N        | 1        | 保留现有辅助与关联                                   |
| [end-tag-html](../entries/end-tag-html.yml)                                                                           | end tag                                  | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [endura-carrot-zelda](../entries/endura-carrot-zelda.yml)                                                             | Endura Carrot                            | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [endura-shroom-zelda](../entries/endura-shroom-zelda.yml)                                                             | Endura Shroom                            | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [endure-general](../entries/endure-general.yml)                                                                       | endure                                   | —        | 1        | 保留现有辅助与关联                                   |
| [energize-general](../entries/energize-general.yml)                                                                   | energize                                 | —        | 1        | 保留现有辅助与关联                                   |
| [energy-cell-zelda](../entries/energy-cell-zelda.yml)                                                                 | Energy Cell                              | —        | 1        | 保留现有辅助与关联                                   |
| [energy-well-zelda](../entries/energy-well-zelda.yml)                                                                 | Energy Well                              | N        | 2        | 保留现有辅助与关联                                   |
| [enhancement-general](../entries/enhancement-general.yml)                                                             | enhancement                              | WF       | 0        | 保留现有辅助与关联                                   |
| [enormous-combat](../entries/enormous-combat.yml)                                                                     | enormous                                 | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [enthusiastic-general](../entries/enthusiastic-general.yml)                                                           | enthusiastic                             | ETY      | 0        | 保留现有辅助与关联                                   |
| [entrance-architecture](../entries/entrance-architecture.yml)                                                         | entrance                                 | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [enum-typescript](../entries/enum-typescript.yml)                                                                     | enum                                     | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [enumerable-javascript](../entries/enumerable-javascript.yml)                                                         | enumerable                               | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [equality-semantics-tech](../entries/equality-semantics-tech.yml)                                                     | equality semantics                       | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [equip-general](../entries/equip-general.yml)                                                                         | equip                                    | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [equipment-equipment](../entries/equipment-equipment.yml)                                                             | equipment                                | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [error-boundary-react](../entries/error-boundary-react.yml)                                                           | error boundary                           | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [error-go](../entries/error-go.yml)                                                                                   | error                                    | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [error-log-nginx](../entries/error-log-nginx.yml)                                                                     | error log                                | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [error-wrapping-go](../entries/error-wrapping-go.yml)                                                                 | error wrapping                           | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [escape-parity-tech](../entries/escape-parity-tech.yml)                                                               | escape parity                            | ETY      | 0        | 保留现有辅助与关联                                   |
| [escape-tech](../entries/escape-tech.yml)                                                                             | escape                                   | ETY      | 0        | 保留现有辅助与关联                                   |
| [escaped-backslash-tech](../entries/escaped-backslash-tech.yml)                                                       | escaped backslash                        | N        | 0        | 保留现有辅助与关联                                   |
| [estimate-general](../entries/estimate-general.yml)                                                                   | estimate                                 | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [etymology-language](../entries/etymology-language.yml)                                                               | etymology                                | ETY      | 1        | 保留现有辅助与关联                                   |
| [event-bubbling-dom](../entries/event-bubbling-dom.yml)                                                               | event bubbling                           | —        | 2        | 保留现有辅助与关联                                   |
| [event-capturing-dom](../entries/event-capturing-dom.yml)                                                             | event capturing                          | —        | 1        | 保留现有辅助与关联                                   |
| [event-delegation-dom](../entries/event-delegation-dom.yml)                                                           | event delegation                         | —        | 1        | 保留现有辅助与关联                                   |
| [event-handler-react](../entries/event-handler-react.yml)                                                             | event handler                            | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [event-loop-javascript](../entries/event-loop-javascript.yml)                                                         | event loop                               | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [event-tech](../entries/event-tech.yml)                                                                               | event                                    | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [evermean-zelda](../entries/evermean-zelda.yml)                                                                       | Evermean                                 | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [evidence-tech](../entries/evidence-tech.yml)                                                                         | evidence                                 | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [exact-match-nginx](../entries/exact-match-nginx.yml)                                                                 | exact match                              | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [examine-general](../entries/examine-general.yml)                                                                     | examine                                  | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [example-general](../entries/example-general.yml)                                                                     | example                                  | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [exchange-currency](../entries/exchange-currency.yml)                                                                 | exchange                                 | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [exchange-traded-fund-finance](../entries/exchange-traded-fund-finance.yml)                                           | exchange-traded fund                     | —        | 1        | 保留现有辅助与关联                                   |
| [exclude-tech](../entries/exclude-tech.yml)                                                                           | exclude                                  | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [exhaustive-tech](../entries/exhaustive-tech.yml)                                                                     | exhaustive                               | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [exhaustiveness-checking-typescript](../entries/exhaustiveness-checking-typescript.yml)                               | exhaustiveness checking                  | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [exit-architecture](../entries/exit-architecture.yml)                                                                 | exit                                     | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [exit-status-tech](../entries/exit-status-tech.yml)                                                                   | exit status                              | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [expected-behavior-workflow](../entries/expected-behavior-workflow.yml)                                               | expected behavior                        | —        | 1        | 保留现有辅助与关联                                   |
| [expected-move-finance](../entries/expected-move-finance.yml)                                                         | expected move                            | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [explicit-tech](../entries/explicit-tech.yml)                                                                         | explicit                                 | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [expose-tech](../entries/expose-tech.yml)                                                                             | expose                                   | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [extension-tech](../entries/extension-tech.yml)                                                                       | extension                                | —        | 1        | 保留现有辅助与关联                                   |
| [factory-architecture](../entries/factory-architecture.yml)                                                           | factory                                  | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [factual-general](../entries/factual-general.yml)                                                                     | factual                                  | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [fail-against-the-original-implementation-workflow](../entries/fail-against-the-original-implementation-workflow.yml) | fail against the original implementation | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [fall-damage-game](../entries/fall-damage-game.yml)                                                                   | fall damage                              | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [fall-general](../entries/fall-general.yml)                                                                           | fall                                     | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [fallback-tech](../entries/fallback.yml)                                                                              | fallback                                 | WF       | 0        | 保留现有辅助与关联                                   |
| [fallback-ui-react](../entries/fallback-ui-react.yml)                                                                 | fallback UI                              | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [fallback-value-css](../entries/fallback-value-css.yml)                                                               | fallback value                           | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [false-diagnostic-tech](../entries/false-diagnostic-tech.yml)                                                         | false diagnostic                         | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [fan-engineering](../entries/fan-engineering.yml)                                                                     | fan                                      | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [fancy-adjective](../entries/fancy-adjective.yml)                                                                     | fancy                                    | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [fancy-feast-pet-food](../entries/fancy-feast-pet-food.yml)                                                           | Fancy Feast                              | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [fast-path-tech](../entries/fast-path-tech.yml)                                                                       | fast path                                | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [fast-travel-game](../entries/fast-travel-game.yml)                                                                   | fast travel                              | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [fear-and-greed-index-finance](../entries/fear-and-greed-index-finance.yml)                                           | fear and greed index                     | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [fear-general](../entries/fear-general.yml)                                                                           | fear                                     | —        | 1        | 保留现有辅助与关联                                   |
| [fearless-general](../entries/fearless-general.yml)                                                                   | fearless                                 | WF       | 2        | 保留现有辅助与关联                                   |
| [feast-food](../entries/feast-food.yml)                                                                               | feast                                    | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [field-presence-tech](../entries/field-presence-tech.yml)                                                             | field presence                           | N        | 0        | 保留现有辅助与关联                                   |
| [fierce-deity-armor-zelda](../entries/fierce-deity-armor-zelda.yml)                                                   | Fierce Deity Armor                       | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [fierce-general](../entries/fierce-general.yml)                                                                       | fierce                                   | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [fifth-general](../entries/fifth-general.yml)                                                                         | fifth                                    | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [financials-finance](../entries/financials-finance.yml)                                                               | financials                               | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [find-general](../entries/find-general.yml)                                                                           | find                                     | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [finding-tech](../entries/finding-tech.yml)                                                                           | finding                                  | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [fire-general](../entries/fire-general.yml)                                                                           | fire                                     | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [fire-temple-zelda](../entries/fire-temple-zelda.yml)                                                                 | Fire Temple                              | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [fiscal-year-finance](../entries/fiscal-year-finance.yml)                                                             | fiscal year                              | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [fish-food](../entries/fish-food.yml)                                                                                 | fish                                     | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [fixed-positioning-css](../entries/fixed-positioning-css.yml)                                                         | fixed positioning                        | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [fixture-tech](../entries/fixture-tech.yml)                                                                           | fixture                                  | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [flame-emitter-zelda](../entries/flame-emitter-zelda.yml)                                                             | Flame Emitter                            | WF       | 0        | 保留现有辅助与关联                                   |
| [flame-general](../entries/flame-general.yml)                                                                         | flame                                    | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [flamebreaker-armor-zelda](../entries/flamebreaker-armor-zelda.yml)                                                   | Flamebreaker Armor                       | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [flex-container-css](../entries/flex-container-css.yml)                                                               | flex container                           | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [flex-item-css](../entries/flex-item-css.yml)                                                                         | flex item                                | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [flexbox-css](../entries/flexbox-css.yml)                                                                             | flexbox                                  | —        | 1        | 保留现有辅助与关联                                   |
| [floating-point-precision-javascript](../entries/floating-point-precision-javascript.yml)                             | floating-point precision                 | N        | 0        | 保留现有辅助与关联                                   |
| [flow-analysis-tech](../entries/flow-analysis-tech.yml)                                                               | flow analysis                            | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [fluid-container-bootstrap](../entries/fluid-container-bootstrap.yml)                                                 | fluid container                          | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [flurry-rush-zelda](../entries/flurry-rush-zelda.yml)                                                                 | Flurry Rush                              | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [flux-construct-zelda](../entries/flux-construct-zelda.yml)                                                           | Flux Construct                           | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [focus-indicator-css](../entries/focus-indicator-css.yml)                                                             | focus indicator                          | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [followed-by-workflow](../entries/followed-by-workflow.yml)                                                           | followed by                              | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [font-stack-css](../entries/font-stack-css.yml)                                                                       | font stack                               | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [foodie-general](../entries/foodie-general.yml)                                                                       | foodie                                   | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [foothill-terrain](../entries/foothill-terrain.yml)                                                                   | foothill                                 | WF       | 0        | 保留现有辅助与关联                                   |
| [forest-terrain](../entries/forest-terrain.yml)                                                                       | forest                                   | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [forge-architecture](../entries/forge-architecture.yml)                                                               | forge                                    | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [forget-general](../entries/forget-general.yml)                                                                       | forget                                   | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [forgotten-temple-zelda](../entries/forgotten-temple-zelda.yml)                                                       | Forgotten Temple                         | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [fork-session-tech](../entries/fork-session-tech.yml)                                                                 | fork                                     | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [form-control-bootstrap](../entries/form-control-bootstrap.yml)                                                       | form control                             | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [form-html](../entries/form-html.yml)                                                                                 | form                                     | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [form-submission-html](../entries/form-submission-html.yml)                                                           | form submission                          | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [formation-general](../entries/formation-general.yml)                                                                 | formation                                | —        | 2        | 保留现有辅助与关联                                   |
| [fortress-architecture](../entries/fortress-architecture.yml)                                                         | fortress                                 | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [forwarding-tech](../entries/forwarding-tech.yml)                                                                     | forwarding                               | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [fountain-architecture](../entries/fountain-architecture.yml)                                                         | fountain                                 | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [fr-css](../entries/fr-css.yml)                                                                                       | fr                                       | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [fragment-identifier-html](../entries/fragment-identifier-html.yml)                                                   | fragment identifier                      | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [fragment-react](../entries/fragment-react.yml)                                                                       | Fragment                                 | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [free-fall-flight](../entries/free-fall-flight.yml)                                                                   | free fall                                | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [fregate-pet-food](../entries/fregate-pet-food.yml)                                                                   | Fregate                                  | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [frog-general](../entries/frog-general.yml)                                                                           | frog                                     | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [froggy-sleeve-zelda](../entries/froggy-sleeve-zelda.yml)                                                             | Froggy Sleeve                            | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [frost-emitter-zelda](../entries/frost-emitter-zelda.yml)                                                             | Frost Emitter                            | WF       | 0        | 保留现有辅助与关联                                   |
| [frost-general](../entries/frost-general.yml)                                                                         | frost                                    | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [frox-zelda](../entries/frox-zelda.yml)                                                                               | Frox                                     | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [fruit-food](../entries/fruit-food.yml)                                                                               | fruit                                    | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [fry-cooking](../entries/fry-cooking.yml)                                                                             | fry                                      | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [function-component-react](../entries/function-component-react.yml)                                                   | function component                       | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [function-signature-typescript](../entries/function-signature-typescript.yml)                                         | function signature                       | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [function-tech](../entries/function-tech.yml)                                                                         | function                                 | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [funds-on-hold-finance](../entries/funds-on-hold-finance.yml)                                                         | funds on hold                            | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [fuse-general](../entries/fuse-general.yml)                                                                           | fuse                                     | —        | 1        | 保留现有辅助与关联                                   |
| [futures-finance](../entries/futures-finance.yml)                                                                     | futures                                  | —        | 1        | 保留现有辅助与关联                                   |
| [fuzz-testing-go](../entries/fuzz-testing-go.yml)                                                                     | fuzz testing                             | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [gainer-finance](../entries/gainer-finance.yml)                                                                       | gainer                                   | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [ganondorf-zelda](../entries/ganondorf-zelda.yml)                                                                     | Ganondorf                                | N        | 0        | 保留现有辅助与关联                                   |
| [gap-css](../entries/gap-css.yml)                                                                                     | gap                                      | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [garbage-collection-javascript](../entries/garbage-collection-javascript.yml)                                         | garbage collection                       | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [gate-architecture](../entries/gate-architecture.yml)                                                                 | gate                                     | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [gear-equipment](../entries/gear-equipment.yml)                                                                       | gear                                     | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [gemini-ai](../entries/gemini-ai.yml)                                                                                 | Gemini                                   | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [general-store-architecture](../entries/general-store-architecture.yml)                                               | general store                            | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [generalize-general](../entries/generalize-general.yml)                                                               | generalize                               | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [generated-output-only-workflow](../entries/generated-output-only-workflow.yml)                                       | generated output only                    | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [generic-constraint-typescript](../entries/generic-constraint-typescript.yml)                                         | generic constraint                       | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [generic-typescript](../entries/generic-typescript.yml)                                                               | generic                                  | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [genuine-tech](../entries/genuine-tech.yml)                                                                           | genuine                                  | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [geoglyph-game](../entries/geoglyph-game.yml)                                                                         | geoglyph                                 | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [gerudo-desert-zelda](../entries/gerudo-desert-zelda.yml)                                                             | Gerudo Desert                            | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [gerudo-town-zelda](../entries/gerudo-town-zelda.yml)                                                                 | Gerudo Town                              | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [gerudo-zelda](../entries/gerudo-zelda.yml)                                                                           | Gerudo                                   | N        | 0        | 保留现有辅助与关联                                   |
| [getter-tech](../entries/getter-tech.yml)                                                                             | getter                                   | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [gibdo-zelda](../entries/gibdo-zelda.yml)                                                                             | Gibdo                                    | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [gleeok-zelda](../entries/gleeok-zelda.yml)                                                                           | Gleeok                                   | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [glide-general](../entries/glide-general.yml)                                                                         | glide                                    | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [glide-shirt-zelda](../entries/glide-shirt-zelda.yml)                                                                 | Glide Shirt                              | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [global-attribute-html](../entries/global-attribute-html.yml)                                                         | global attribute                         | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [global-augmentation-typescript](../entries/global-augmentation-typescript.yml)                                       | global augmentation                      | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [gloom-recovery-zelda](../entries/gloom-recovery-zelda.yml)                                                           | Gloom Recovery                           | N        | 1        | 保留现有辅助与关联                                   |
| [gloom-resistance-zelda](../entries/gloom-resistance-zelda.yml)                                                       | Gloom Resistance                         | —        | 1        | 保留现有辅助与关联                                   |
| [gloom-spawn-zelda](../entries/gloom-spawn-zelda.yml)                                                                 | Gloom Spawn                              | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [gloom-sword-zelda](../entries/gloom-sword-zelda.yml)                                                                 | Gloom Sword                              | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [gloom-zelda](../entries/gloom-zelda.yml)                                                                             | Gloom                                    | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [go-mod-go](../entries/go-mod-go.yml)                                                                                 | go.mod                                   | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [go-module](../entries/go-module.yml)                                                                                 | Go module                                | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [go-sum-go](../entries/go-sum-go.yml)                                                                                 | go.sum                                   | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [goal-general](../entries/goal-general.yml)                                                                           | goal                                     | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [golden-apple-zelda](../entries/golden-apple-zelda.yml)                                                               | Golden Apple                             | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [golden-general](../entries/golden-general.yml)                                                                       | golden                                   | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [goron-city-zelda](../entries/goron-city-zelda.yml)                                                                   | Goron City                               | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [goron-spice-zelda](../entries/goron-spice-zelda.yml)                                                                 | Goron Spice                              | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [goron-zelda](../entries/goron-zelda.yml)                                                                             | Goron                                    | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [goroutine-go](../entries/goroutine-go.yml)                                                                           | goroutine                                | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [gourmet-food](../entries/gourmet-food.yml)                                                                           | gourmet                                  | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [graceful-degradation-web](../entries/graceful-degradation-web.yml)                                                   | graceful degradation                     | —        | 1        | 保留现有辅助与关联                                   |
| [graceful-shutdown-tech](../entries/graceful-shutdown-tech.yml)                                                       | graceful shutdown                        | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [gratitude-general](../entries/gratitude-general.yml)                                                                 | gratitude                                | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [great-abandoned-central-mine-zelda](../entries/great-abandoned-central-mine-zelda.yml)                               | Great Abandoned Central Mine             | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [great-deku-tree-zelda](../entries/great-deku-tree-zelda.yml)                                                         | Great Deku Tree                          | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [great-eagle-bow-zelda](../entries/great-eagle-bow-zelda.yml)                                                         | Great Eagle Bow                          | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [great-sky-island-zelda](../entries/great-sky-island-zelda.yml)                                                       | Great Sky Island                         | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [greaves-clothing](../entries/greaves-clothing.yml)                                                                   | greaves                                  | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [greed-general](../entries/greed-general.yml)                                                                         | greed                                    | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [greens-food](../entries/greens-food.yml)                                                                             | greens                                   | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [grid-area-css](../entries/grid-area-css.yml)                                                                         | grid area                                | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [grid-container-css](../entries/grid-container-css.yml)                                                               | grid container                           | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [grid-layout-css](../entries/grid-layout-css.yml)                                                                     | grid layout                              | —        | 1        | 保留现有辅助与关联                                   |
| [grid-system-bootstrap](../entries/grid-system-bootstrap.yml)                                                         | grid system                              | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [grid-track-css](../entries/grid-track-css.yml)                                                                       | grid track                               | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [grill-cooking](../entries/grill-cooking.yml)                                                                         | grill                                    | N        | 1        | 保留现有辅助与关联                                   |
| [grok-ai](../entries/grok-ai.yml)                                                                                     | Grok                                     | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [guard-combat](../entries/guard-combat.yml)                                                                           | guard                                    | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [guard-tech](../entries/guard-tech.yml)                                                                               | guard                                    | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [guardian-general](../entries/guardian-general.yml)                                                                   | guardian                                 | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [guidance-story](../entries/guidance-story.yml)                                                                       | guidance                                 | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [gust-weather](../entries/gust-weather.yml)                                                                           | gust                                     | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [gutter-bootstrap](../entries/gutter-bootstrap.yml)                                                                   | gutter                                   | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [halberd-weapon](../entries/halberd-weapon.yml)                                                                       | halberd                                  | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [hammer-weapon](../entries/hammer-weapon.yml)                                                                         | hammer                                   | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [hand-general](../entries/hand-general.yml)                                                                           | hand                                     | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [handoff-workflow](../entries/handoff-workflow.yml)                                                                   | handoff                                  | WF       | 1        | 保留现有辅助与关联                                   |
| [happens-before-go](../entries/happens-before-go.yml)                                                                 | happens-before                           | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [harness-ai](../entries/harness-ai.yml)                                                                               | harness                                  | —        | 2        | 保留现有辅助与关联                                   |
| [harness-tech](../entries/harness-tech.yml)                                                                           | harness                                  | —        | 1        | 保留现有辅助与关联                                   |
| [hashing-tech](../entries/hashing-tech.yml)                                                                           | hashing                                  | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [hasty-general](../entries/hasty-general.yml)                                                                         | hasty                                    | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [hateno-village-zelda](../entries/hateno-village-zelda.yml)                                                           | Hateno Village                           | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [hdmi-cable-hardware](../entries/hdmi-cable-hardware.yml)                                                             | HDMI cable                               | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [hdmi-hardware](../entries/hdmi-hardware.yml)                                                                         | HDMI                                     | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [hdmi-type-a-connector](../entries/hdmi-type-a-connector.yml)                                                         | HDMI Type A connector                    | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [head-general](../entries/head-general.yml)                                                                           | head                                     | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [headband-clothing](../entries/headband-clothing.yml)                                                                 | headband                                 | WF       | 1        | 保留现有辅助与关联                                   |
| [headdress-clothing](../entries/headdress-clothing.yml)                                                               | headdress                                | WF       | 1        | 保留现有辅助与关联                                   |
| [header-cell-html](../entries/header-cell-html.yml)                                                                   | header cell                              | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [heading-html](../entries/heading-html.yml)                                                                           | heading                                  | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [heading-level-html](../entries/heading-level-html.yml)                                                               | heading level                            | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [headless-tech](../entries/headless-tech.yml)                                                                         | headless                                 | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [heap-memory-javascript](../entries/heap-memory-javascript.yml)                                                       | heap memory                              | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [heart-container-zelda](../entries/heart-container-zelda.yml)                                                         | Heart Container                          | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [hearty-general](../entries/hearty-general.yml)                                                                       | hearty                                   | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [hearty-radish-zelda](../entries/hearty-radish-zelda.yml)                                                             | Hearty Radish                            | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [hearty-salmon-zelda](../entries/hearty-salmon-zelda.yml)                                                             | Hearty Salmon                            | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [hearty-truffle-zelda](../entries/hearty-truffle-zelda.yml)                                                           | Hearty Truffle                           | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [heat-map-data](../entries/heat-map-data.yml)                                                                         | heat map                                 | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [heat-resistance-combat](../entries/heat-resistance-combat.yml)                                                       | heat resistance                          | —        | 1        | 保留现有辅助与关联                                   |
| [heir-general](../entries/heir-general.yml)                                                                           | heir                                     | N        | 1        | 保留现有辅助与关联                                   |
| [helm-clothing](../entries/helm-clothing.yml)                                                                         | helm                                     | —        | 1        | 保留现有辅助与关联                                   |
| [helmet-clothing](../entries/helmet-clothing.yml)                                                                     | helmet                                   | —        | 1        | 保留现有辅助与关联                                   |
| [heritage-type-tech](../entries/heritage-type-tech.yml)                                                               | heritage type                            | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [hermes-agent-ai](../entries/hermes-agent-ai.yml)                                                                     | Hermes                                   | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [hero-general](../entries/hero-general.yml)                                                                           | hero                                     | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [hestu-zelda](../entries/hestu-zelda.yml)                                                                             | Hestu                                    | N        | 0        | 保留现有辅助与关联                                   |
| [hidden-general](../entries/hidden-general.yml)                                                                       | hidden                                   | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [hideout-architecture](../entries/hideout-architecture.yml)                                                           | hideout                                  | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [higher-order-function-javascript](../entries/higher-order-function-javascript.yml)                                   | higher-order function                    | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [highlands-terrain](../entries/highlands-terrain.yml)                                                                 | highlands                                | WF       | 0        | 保留现有辅助与关联                                   |
| [hills-science-diet-pet-food](../entries/hills-science-diet-pet-food.yml)                                             | Hill’s Science Diet                      | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [hinox-zelda](../entries/hinox-zelda.yml)                                                                             | Hinox                                    | N        | 0        | 保留现有辅助与关联                                   |
| [hoisting-javascript](../entries/hoisting-javascript.yml)                                                             | hoisting                                 | —        | 1        | 保留现有辅助与关联                                   |
| [hold-general](../entries/hold-general.yml)                                                                           | hold                                     | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [holdings-finance](../entries/holdings-finance.yml)                                                                   | holdings                                 | —        | 1        | 保留现有辅助与关联                                   |
| [homing-cart-zelda](../entries/homing-cart-zelda.yml)                                                                 | Homing Cart                              | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [homing-engineering](../entries/homing-engineering.yml)                                                               | homing                                   | WF       | 0        | 保留现有辅助与关联                                   |
| [honey-food](../entries/honey-food.yml)                                                                               | honey                                    | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [honorable-general](../entries/honorable-general.yml)                                                                 | honorable                                | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [hood-clothing](../entries/hood-clothing.yml)                                                                         | hood                                     | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [hook-hardware](../entries/hook-hardware.yml)                                                                         | hook                                     | —        | 1        | 保留现有辅助与关联                                   |
| [hook-react](../entries/hook-react.yml)                                                                               | Hook                                     | —        | 1        | 保留现有辅助与关联                                   |
| [hook-tech](../entries/hook-tech.yml)                                                                                 | hook                                     | —        | 2        | 保留现有辅助与关联                                   |
| [horriblin-zelda](../entries/horriblin-zelda.yml)                                                                     | Horriblin                                | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [host-docker](../entries/host-docker.yml)                                                                             | host                                     | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [hostile-combat](../entries/hostile-combat.yml)                                                                       | hostile                                  | ETY / N  | 1        | 保留现有辅助与关联                                   |
| [hover-general](../entries/hover-general.yml)                                                                         | hover                                    | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [hover-stone-zelda](../entries/hover-stone-zelda.yml)                                                                 | Hover Stone                              | WF       | 0        | 保留现有辅助与关联                                   |
| [html-language](../entries/html-language.yml)                                                                         | HTML                                     | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [http-caching-web](../entries/http-caching-web.yml)                                                                   | HTTP caching                             | N        | 0        | 保留现有辅助与关联                                   |
| [hudson-zelda](../entries/hudson-zelda.yml)                                                                           | Hudson                                   | N        | 0        | 保留现有辅助与关联                                   |
| [hydrant-engineering](../entries/hydrant-engineering.yml)                                                             | hydrant                                  | ETY      | 0        | 保留现有辅助与关联                                   |
| [hydration-react](../entries/hydration-react.yml)                                                                     | hydration                                | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [hylian-hood-zelda](../entries/hylian-hood-zelda.yml)                                                                 | Hylian Hood                              | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [hylian-shield-zelda](../entries/hylian-shield-zelda.yml)                                                             | Hylian Shield                            | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [hylian-tunic-zelda](../entries/hylian-tunic-zelda.yml)                                                               | Hylian Tunic                             | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [hylian-zelda](../entries/hylian-zelda.yml)                                                                           | Hylian                                   | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [hyperlink-html](../entries/hyperlink-html.yml)                                                                       | hyperlink                                | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [hyrule-bass-zelda](../entries/hyrule-bass-zelda.yml)                                                                 | Hyrule Bass                              | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [hyrule-castle-zelda](../entries/hyrule-castle-zelda.yml)                                                             | Hyrule Castle                            | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [hyrule-compendium-zelda](../entries/hyrule-compendium-zelda.yml)                                                     | Hyrule Compendium                        | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [hyrule-zelda](../entries/hyrule-zelda.yml)                                                                           | Hyrule                                   | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [id-selector-css](../entries/id-selector-css.yml)                                                                     | ID selector                              | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [identifier-tech](../entries/identifier-tech.yml)                                                                     | identifier                               | WF       | 0        | 保留现有辅助与关联                                   |
| [idiomatic-tech](../entries/idiomatic.yml)                                                                            | idiomatic                                | ETY      | 0        | 保留现有辅助与关联                                   |
| [illuminate-general](../entries/illuminate-general.yml)                                                               | illuminate                               | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [image-digest-docker](../entries/image-digest-docker.yml)                                                             | image digest                             | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [image-docker](../entries/image-docker.yml)                                                                           | image                                    | —        | 0        | R03                                                  |
| [image-layer-docker](../entries/image-layer-docker.yml)                                                               | image layer                              | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [image-repository-docker](../entries/image-repository-docker.yml)                                                     | image repository                         | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [image-tag-docker](../entries/image-tag-docker.yml)                                                                   | image tag                                | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [impa-zelda](../entries/impa-zelda.yml)                                                                               | Impa                                     | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [imperative-general](../entries/imperative-general.yml)                                                               | imperative                               | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [implement-tech](../entries/implement-tech.yml)                                                                       | implement                                | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [implicit-implementation-go](../entries/implicit-implementation-go.yml)                                               | implicit implementation                  | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [imply-general](../entries/imply-general.yml)                                                                         | imply                                    | N        | 1        | 保留现有辅助与关联                                   |
| [important-css](../entries/important-css.yml)                                                                         | !important                               | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [imprison-story](../entries/imprison-story.yml)                                                                       | imprison                                 | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [in-one-pass-workflow](../entries/in-one-pass-workflow.yml)                                                           | in one pass                              | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [in-order-workflow](../entries/in-order-workflow.yml)                                                                 | in order                                 | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [inaccessible-tech](../entries/inaccessible-tech.yml)                                                                 | inaccessible                             | WF       | 0        | 保留现有辅助与关联                                   |
| [income-finance](../entries/income-finance.yml)                                                                       | income                                   | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [incremental-build-tech](../entries/incremental-build-tech.yml)                                                       | incremental build                        | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [independently-tech](../entries/independently-tech.yml)                                                               | independently                            | WF       | 0        | 保留现有辅助与关联                                   |
| [index-finance](../entries/index-finance.yml)                                                                         | index                                    | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [indexed-access-type-typescript](../entries/indexed-access-type-typescript.yml)                                       | indexed access type                      | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [infer-tech](../entries/infer-tech.yml)                                                                               | infer                                    | N        | 1        | 保留现有辅助与关联                                   |
| [ingredient-cooking](../entries/ingredient-cooking.yml)                                                               | ingredient                               | —        | 1        | 保留现有辅助与关联                                   |
| [inherit-story](../entries/inherit-story.yml)                                                                         | inherit                                  | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [inheritance-css](../entries/inheritance-css.yml)                                                                     | inheritance                              | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [initializer-tech](../entries/initializer-tech.yml)                                                                   | initializer                              | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [inline-box-css](../entries/inline-box-css.yml)                                                                       | inline box                               | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [inline-frame-html](../entries/inline-frame-html.yml)                                                                 | inline frame                             | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [inn-architecture](../entries/inn-architecture.yml)                                                                   | inn                                      | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [input-group-bootstrap](../entries/input-group-bootstrap.yml)                                                         | input group                              | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [input-type-html](../entries/input-type-html.yml)                                                                     | input type                               | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [inspect-tech](../entries/inspect-tech.yml)                                                                           | inspect                                  | ETY      | 1        | 保留现有辅助与关联                                   |
| [instinct-general](../entries/instinct-general.yml)                                                                   | instinct                                 | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [instinct-pet-food](../entries/instinct-pet-food.yml)                                                                 | Instinct                                 | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [instruction-tech](../entries/instruction-tech.yml)                                                                   | instruction                              | —        | 1        | 保留现有辅助与关联                                   |
| [int-go](../entries/int-go.yml)                                                                                       | int                                      | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [intact-tech](../entries/intact-tech.yml)                                                                             | intact                                   | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [integration-test-tech](../entries/integration-test-tech.yml)                                                         | integration test                         | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [integrity-general](../entries/integrity-general.yml)                                                                 | integrity                                | ETY      | 0        | 保留现有辅助与关联                                   |
| [interact-general](../entries/interact-general.yml)                                                                   | interact                                 | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [interactive-content-html](../entries/interactive-content-html.yml)                                                   | interactive content                      | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [interest-finance](../entries/interest-finance.yml)                                                                   | interest                                 | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [interface-go](../entries/interface-go.yml)                                                                           | interface                                | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [interface-typescript](../entries/interface-typescript.yml)                                                           | interface                                | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [internal-redirect-nginx](../entries/internal-redirect-nginx.yml)                                                     | internal redirect                        | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [internals-tech](../entries/internals-tech.yml)                                                                       | internals                                | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [interrupt-workflow](../entries/interrupt-workflow.yml)                                                               | interrupt                                | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [intersection-type-typescript](../entries/intersection-type-typescript.yml)                                           | intersection type                        | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [introduce-tech](../entries/introduce-tech.yml)                                                                       | introduce                                | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [invalid-state-bootstrap](../entries/invalid-state-bootstrap.yml)                                                     | invalid state                            | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [invalidate-tech](../entries/invalidate-tech.yml)                                                                     | invalidate                               | —        | 0        | WF01                                                 |
| [invaluable-general](../entries/invaluable-general.yml)                                                               | invaluable                               | N        | 0        | 保留现有辅助与关联                                   |
| [inventory-ui](../entries/inventory-ui.yml)                                                                           | inventory                                | N        | 1        | 保留现有辅助与关联                                   |
| [investigate-tech](../entries/investigate-tech.yml)                                                                   | investigate                              | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [invocation-workflow](../entries/invocation-workflow.yml)                                                             | invocation                               | WF       | 1        | 保留现有辅助与关联                                   |
| [invoke-tech](../entries/invoke-tech.yml)                                                                             | invoke                                   | —        | 1        | 保留现有辅助与关联                                   |
| [island-terrain](../entries/island-terrain.yml)                                                                       | island                                   | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [isolated-worktree-workflow](../entries/isolated-worktree-workflow.yml)                                               | isolated worktree                        | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [iterable-tech](../entries/iterable-tech.yml)                                                                         | iterable                                 | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [jar-general](../entries/jar-general.yml)                                                                             | jar                                      | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [javascript-plugin-bootstrap](../entries/javascript-plugin-bootstrap.yml)                                             | JavaScript plugin                        | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [joiner-tech](../entries/joiner-tech.yml)                                                                             | joiner                                   | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [josha-zelda](../entries/josha-zelda.yml)                                                                             | Josha                                    | N        | 0        | 保留现有辅助与关联                                   |
| [jsx-react](../entries/jsx-react.yml)                                                                                 | JSX                                      | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [justify-general](../entries/justify-general.yml)                                                                     | justify                                  | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [kakariko-village-zelda](../entries/kakariko-village-zelda.yml)                                                       | Kakariko Village                         | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [keese-zelda](../entries/keese-zelda.yml)                                                                             | Keese                                    | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [key-items-zelda](../entries/key-items-zelda.yml)                                                                     | Key Items                                | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [key-react](../entries/key-react.yml)                                                                                 | key                                      | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [keyframe-css](../entries/keyframe-css.yml)                                                                           | keyframe                                 | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [keyof-typescript](../entries/keyof-typescript.yml)                                                                   | keyof                                    | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [king-general](../entries/king-general.yml)                                                                           | king                                     | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [king-gleeok-zelda](../entries/king-gleeok-zelda.yml)                                                                 | King Gleeok                              | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [kingdom-general](../entries/kingdom-general.yml)                                                                     | kingdom                                  | WF       | 0        | 保留现有辅助与关联                                   |
| [knight-general](../entries/knight-general.yml)                                                                       | knight                                   | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [knights-broadsword-zelda](../entries/knights-broadsword-zelda.yml)                                                   | Knight's Broadsword                      | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [korok-forest-zelda](../entries/korok-forest-zelda.yml)                                                               | Korok Forest                             | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [korok-seed-zelda](../entries/korok-seed-zelda.yml)                                                                   | Korok Seed                               | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [korok-zelda](../entries/korok-zelda.yml)                                                                             | Korok                                    | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [label-html](../entries/label-html.yml)                                                                               | label                                    | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [lake-terrain](../entries/lake-terrain.yml)                                                                           | lake                                     | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [land-motion](../entries/land-motion.yml)                                                                             | land                                     | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [landing-general](../entries/landing-general.yml)                                                                     | landing                                  | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [landmark-html](../entries/landmark-html.yml)                                                                         | landmark                                 | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [language-general](../entries/language-general.yml)                                                                   | language                                 | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [launch-motion](../entries/launch-motion.yml)                                                                         | launch                                   | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [lava-terrain](../entries/lava-terrain.yml)                                                                           | lava                                     | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [layout-browser-rendering](../entries/layout-browser-rendering.yml)                                                   | layout                                   | —        | 1        | 保留现有辅助与关联                                   |
| [lazy-loading-html](../entries/lazy-loading-html.yml)                                                                 | lazy loading                             | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [least-recently-used-cache](../entries/least-recently-used-cache.yml)                                                 | least recently used                      | N        | 0        | 保留现有辅助与关联                                   |
| [leave-untouched-workflow](../entries/leave-untouched-workflow.yml)                                                   | leave untouched                          | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [legacy-tech](../entries/legacy-tech.yml)                                                                             | legacy                                   | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [legendsandy-pet-food](../entries/legendsandy-pet-food.yml)                                                           | Legendsandy                              | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [leggings-clothing](../entries/leggings-clothing.yml)                                                                 | leggings                                 | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [length-go](../entries/length-go.yml)                                                                                 | length                                   | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [lex-name-element](../entries/lex-name-element.yml)                                                                   | lex                                      | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [lexforge-project](../entries/lexforge-project.yml)                                                                   | LexForge                                 | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [lifecycle-tech](../entries/lifecycle-tech.yml)                                                                       | lifecycle                                | WF       | 0        | 保留现有辅助与关联                                   |
| [lifting-state-up-react](../entries/lifting-state-up-react.yml)                                                       | lifting state up                         | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [light-dragon-zelda](../entries/light-dragon-zelda.yml)                                                               | Light Dragon                             | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [light-engineering](../entries/light-engineering.yml)                                                                 | light                                    | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [light-of-blessing-zelda](../entries/light-of-blessing-zelda.yml)                                                     | Light of Blessing                        | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [lightning-temple-zelda](../entries/lightning-temple-zelda.yml)                                                       | Lightning Temple                         | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [lightning-terrain](../entries/lightning-terrain.yml)                                                                 | lightning                                | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [lightroot-zelda](../entries/lightroot-zelda.yml)                                                                     | Lightroot                                | WF       | 0        | 保留现有辅助与关联                                   |
| [like-like-zelda](../entries/like-like-zelda.yml)                                                                     | Like Like                                | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [limitation-tech](../entries/limitation-tech.yml)                                                                     | limitation                               | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [line-height-css](../entries/line-height-css.yml)                                                                     | line height                              | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [linear-tech](../entries/linear-tech.yml)                                                                             | linear                                   | —        | 1        | 保留现有辅助与关联                                   |
| [link-zelda](../entries/link-zelda.yml)                                                                               | Link                                     | N        | 0        | 保留现有辅助与关联                                   |
| [linked-list-merge-algorithm](../entries/linked-list-merge-algorithm.yml)                                             | linked list merge                        | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [linked-list-tech](../entries/linked-list-tech.yml)                                                                   | linked list                              | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [list-group-bootstrap](../entries/list-group-bootstrap.yml)                                                           | list group                               | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [list-item-html](../entries/list-item-html.yml)                                                                       | list item                                | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [list-rendering-react](../entries/list-rendering-react.yml)                                                           | list rendering                           | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [listen-nginx](../entries/listen-nginx.yml)                                                                           | listen                                   | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [literal-programming](../entries/literal-programming.yml)                                                             | literal                                  | N        | 1        | 保留现有辅助与关联                                   |
| [literal-text](../entries/literal-text.yml)                                                                           | literal                                  | N        | 1        | 保留现有辅助与关联                                   |
| [literal-type-typescript](../entries/literal-type-typescript.yml)                                                     | literal type                             | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [lizalfos-zelda](../entries/lizalfos-zelda.yml)                                                                       | Lizalfos                                 | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [load-balancing-nginx](../entries/load-balancing-nginx.yml)                                                           | load balancing                           | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [location-nginx](../entries/location-nginx.yml)                                                                       | location                                 | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [lockfile-tech](../entries/lockfile-tech.yml)                                                                         | lockfile                                 | WF       | 0        | 保留现有辅助与关联                                   |
| [lone-surrogate-tech](../entries/lone-surrogate-tech.yml)                                                             | lone surrogate                           | —        | 1        | 保留现有辅助与关联                                   |
| [longest-prefix-nginx](../entries/longest-prefix-nginx.yml)                                                           | longest prefix                           | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [longhand-property-css](../entries/longhand-property-css.yml)                                                         | longhand property                        | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [lookout-general](../entries/lookout-general.yml)                                                                     | lookout                                  | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [lookout-landing-zelda](../entries/lookout-landing-zelda.yml)                                                         | Lookout Landing                          | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [lose-general](../entries/lose-general.yml)                                                                           | lose                                     | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [loser-finance](../entries/loser-finance.yml)                                                                         | loser                                    | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [loss-finance](../entries/loss-finance.yml)                                                                           | loss                                     | —        | 1        | 保留现有辅助与关联                                   |
| [lowering-compiler](../entries/lowering-compiler.yml)                                                                 | lowering                                 | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [loyal-general](../entries/loyal-general.yml)                                                                         | loyal                                    | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [lurelin-village-zelda](../entries/lurelin-village-zelda.yml)                                                         | Lurelin Village                          | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [lynel-zelda](../entries/lynel-zelda.yml)                                                                             | Lynel                                    | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [magic-rod-zelda](../entries/magic-rod-zelda.yml)                                                                     | Magic Rod                                | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [main-axis-css](../entries/main-axis-css.yml)                                                                         | main axis                                | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [main-quest-game](../entries/main-quest-game.yml)                                                                     | main quest                               | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [main-quests-zelda](../entries/main-quests-zelda.yml)                                                                 | Main Quests                              | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [maintainable-tech](../entries/maintainable-tech.yml)                                                                 | maintainable                             | WF       | 0        | 保留现有辅助与关联                                   |
| [majestic-general](../entries/majestic-general.yml)                                                                   | majestic                                 | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [make-go](../entries/make-go.yml)                                                                                     | make                                     | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [male-to-male-hardware](../entries/male-to-male-hardware.yml)                                                         | male-to-male                             | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [malevolent-general](../entries/malevolent-general.yml)                                                               | malevolent                               | ETY      | 0        | 保留现有辅助与关联                                   |
| [malformed-tech](../entries/malformed-tech.yml)                                                                       | malformed                                | WF       | 0        | 保留现有辅助与关联                                   |
| [mandatory-general](../entries/mandatory-general.yml)                                                                 | mandatory                                | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [manifest-tech](../entries/manifest-tech.yml)                                                                         | manifest                                 | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [manual-save-game](../entries/manual-save-game.yml)                                                                   | manual save                              | —        | 2        | 保留现有辅助与关联                                   |
| [map-go](../entries/map-go.yml)                                                                                       | map                                      | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [map-ui](../entries/map-ui.yml)                                                                                       | map                                      | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [mapped-type-typescript](../entries/mapped-type-typescript.yml)                                                       | mapped type                              | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [marbled-gohma-zelda](../entries/marbled-gohma-zelda.yml)                                                             | Marbled Gohma                            | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [margin-account-finance](../entries/margin-account-finance.yml)                                                       | margin account                           | —        | 1        | 保留现有辅助与关联                                   |
| [margin-collapsing-css](../entries/margin-collapsing-css.yml)                                                         | margin collapsing                        | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [margin-css](../entries/margin-css.yml)                                                                               | margin                                   | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [margin-finance](../entries/margin-finance.yml)                                                                       | margin                                   | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [marker-tech](../entries/marker-tech.yml)                                                                             | marker                                   | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [market-capitalization-finance](../entries/market-capitalization-finance.yml)                                         | market capitalization                    | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [market-value-finance](../entries/market-value-finance.yml)                                                           | market value                             | —        | 1        | 保留现有辅助与关联                                   |
| [markup-html](../entries/markup-html.yml)                                                                             | markup                                   | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [mashed-food](../entries/mashed-food.yml)                                                                             | mashed                                   | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [mashed-potatoes-food](../entries/mashed-potatoes-food.yml)                                                           | mashed potatoes                          | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [mask-clothing](../entries/mask-clothing.yml)                                                                         | mask                                     | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [master-kohga-zelda](../entries/master-kohga-zelda.yml)                                                               | Master Kohga                             | N        | 0        | 保留现有辅助与关联                                   |
| [master-sword-zelda](../entries/master-sword-zelda.yml)                                                               | Master Sword                             | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [material-general](../entries/material-general.yml)                                                                   | material                                 | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [material-tech](../entries/material-tech.yml)                                                                         | material                                 | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [meal-cooking](../entries/meal-cooking.yml)                                                                           | meal                                     | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [meat-food](../entries/meat-food.yml)                                                                                 | meat                                     | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [media-query-css](../entries/media-query-css.yml)                                                                     | media query                              | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [melee-combat](../entries/melee-combat.yml)                                                                           | melee                                    | ETY      | 0        | 保留现有辅助与关联                                   |
| [memoization-programming](../entries/memoization-programming.yml)                                                     | memoization                              | —        | 1        | 保留现有辅助与关联                                   |
| [memory-ai](../entries/memory-ai.yml)                                                                                 | memory                                   | —        | 0        | R05 反向                                             |
| [memory-story](../entries/memory-story.yml)                                                                           | memory                                   | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [menu-ui](../entries/menu-ui.yml)                                                                                     | menu                                     | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [merge-git](../entries/merge-git.yml)                                                                                 | merge                                    | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [merit-general](../entries/merit-general.yml)                                                                         | merit                                    | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [metadata-html](../entries/metadata-html.yml)                                                                         | metadata                                 | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [method-set-go](../entries/method-set-go.yml)                                                                         | method set                               | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [method-tech](../entries/method-tech.yml)                                                                             | method                                   | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [microtask-javascript](../entries/microtask-javascript.yml)                                                           | microtask                                | WF       | 0        | 保留现有辅助与关联                                   |
| [middleware-web](../entries/middleware-web.yml)                                                                       | middleware                               | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [mighty-bananas-zelda](../entries/mighty-bananas-zelda.yml)                                                           | Mighty Bananas                           | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [mighty-general](../entries/mighty-general.yml)                                                                       | mighty                                   | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [milk-food](../entries/milk-food.yml)                                                                                 | milk                                     | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [mine-architecture](../entries/mine-architecture.yml)                                                                 | mine                                     | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [mine-cart-engineering](../entries/mine-cart-engineering.yml)                                                         | mine cart                                | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [miner-general](../entries/miner-general.yml)                                                                         | miner                                    | WF       | 0        | 保留现有辅助与关联                                   |
| [miners-mask-zelda](../entries/miners-mask-zelda.yml)                                                                 | Miner's Mask                             | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [mineru-zelda](../entries/mineru-zelda.yml)                                                                           | Mineru                                   | N        | 0        | 保留现有辅助与关联                                   |
| [mirror-general](../entries/mirror-general.yml)                                                                       | mirror                                   | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [misalign-tech](../entries/misalign-tech.yml)                                                                         | misalign                                 | WF       | 0        | 保留现有辅助与关联                                   |
| [mnemonic-memory](../entries/mnemonic-memory.yml)                                                                     | mnemonic                                 | ETY / N  | 0        | 保留现有辅助与关联                                   |
| [moblin-zelda](../entries/moblin-zelda.yml)                                                                           | Moblin                                   | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [mock-tech](../entries/mock-tech.yml)                                                                                 | mock                                     | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [modal-bootstrap](../entries/modal-bootstrap.yml)                                                                     | modal                                    | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [model-ai](../entries/model-ai.yml)                                                                                   | model                                    | —        | 1        | 暂缓，模型的广义词源尚未评估其对当前义项的学习价值   |
| [module-path-go](../entries/module-path-go.yml)                                                                       | module path                              | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [module-specifier-tech](../entries/module-specifier-tech.yml)                                                         | module specifier                         | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [module-typescript](../entries/module-typescript.yml)                                                                 | module                                   | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [monitor-hardware](../entries/monitor-hardware.yml)                                                                   | monitor                                  | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [mount-point-docker](../entries/mount-point-docker.yml)                                                               | mount point                              | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [mount-tech](../entries/mount-tech.yml)                                                                               | mount                                    | N        | 1        | 保留现有辅助与关联                                   |
| [mountain-terrain](../entries/mountain-terrain.yml)                                                                   | mountain                                 | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [mucktorok-zelda](../entries/mucktorok-zelda.yml)                                                                     | Mucktorok                                | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [multibyte-tech](../entries/multibyte-tech.yml)                                                                       | multibyte                                | WF       | 0        | 保留现有辅助与关联                                   |
| [multiline-range-tech](../entries/multiline-range-tech.yml)                                                           | multiline range                          | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [multiple-return-values-go](../entries/multiple-return-values-go.yml)                                                 | multiple return values                   | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [muse-ai](../entries/muse-ai.yml)                                                                                     | Muse                                     | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [mushroom-food](../entries/mushroom-food.yml)                                                                         | mushroom                                 | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [mutable-tech](../entries/mutable-tech.yml)                                                                           | mutable                                  | ETY      | 0        | 保留现有辅助与关联                                   |
| [mutex-go](../entries/mutex-go.yml)                                                                                   | mutex                                    | —        | 0        | 暂缓，尝试的词典页无法打开，缩写展开暂不写入         |
| [myfoodie-pet-food](../entries/myfoodie-pet-food.yml)                                                                 | Myfoodie                                 | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [mysterious-general](../entries/mysterious-general.yml)                                                               | mysterious                               | WF       | 0        | 保留现有辅助与关联                                   |
| [mystery-story](../entries/mystery-story.yml)                                                                         | mystery                                  | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [named-result-parameter-go](../entries/named-result-parameter-go.yml)                                                 | named result parameter                   | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [namespace-typescript](../entries/namespace-typescript.yml)                                                           | namespace                                | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [narrowing-typescript](../entries/narrowing-typescript.yml)                                                           | narrowing                                | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [native-addon-tech](../entries/native-addon-tech.yml)                                                                 | native addon                             | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [navbar-bootstrap](../entries/navbar-bootstrap.yml)                                                                   | navbar                                   | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [negative-lookbehind-tech](../entries/negative-lookbehind-tech.yml)                                                   | negative lookbehind                      | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [nested-tech](../entries/nested-tech.yml)                                                                             | nested                                   | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [nesting-html](../entries/nesting-html.yml)                                                                           | nesting                                  | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [net-assets-finance](../entries/net-assets-finance.yml)                                                               | net assets                               | —        | 1        | 保留现有辅助与关联                                   |
| [net-finance](../entries/net-finance.yml)                                                                             | net                                      | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [net-income-finance](../entries/net-income-finance.yml)                                                               | net income                               | —        | 1        | 保留现有辅助与关联                                   |
| [never-typescript](../entries/never-typescript.yml)                                                                   | never                                    | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [new-go](../entries/new-go.yml)                                                                                       | new                                      | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [nginx-platform](../entries/nginx-platform.yml)                                                                       | NGINX                                    | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [nil-go](../entries/nil-go.yml)                                                                                       | nil                                      | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [noble-general](../entries/noble-general.yml)                                                                         | noble                                    | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [nonthrowing-tech](../entries/nonthrowing-tech.yml)                                                                   | nonthrowing                              | WF       | 0        | 保留现有辅助与关联                                   |
| [noodles-food](../entries/noodles-food.yml)                                                                           | noodles                                  | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [normal-flow-css](../entries/normal-flow-css.yml)                                                                     | normal flow                              | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [note-general](../entries/note-general.yml)                                                                           | note                                     | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [null-typescript](../entries/null-typescript.yml)                                                                     | null                                     | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [number-typescript](../entries/number-typescript.yml)                                                                 | number                                   | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [nutro-pet-food](../entries/nutro-pet-food.yml)                                                                       | Nutro                                    | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [object-typescript](../entries/object-typescript.yml)                                                                 | object                                   | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [objective-ui](../entries/objective-ui.yml)                                                                           | objective                                | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [observation-deck-general](../entries/observation-deck-general.yml)                                                   | observation deck                         | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [observed-result-workflow](../entries/observed-result-workflow.yml)                                                   | observed result                          | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [occupied-workflow](../entries/occupied-workflow.yml)                                                                 | occupied                                 | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [octorok-zelda](../entries/octorok-zelda.yml)                                                                         | Octorok                                  | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [offcanvas-bootstrap](../entries/offcanvas-bootstrap.yml)                                                             | offcanvas                                | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [oil-food](../entries/oil-food.yml)                                                                                   | oil                                      | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [omelet-cooking](../entries/omelet-cooking.yml)                                                                       | omelet                                   | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [omit-tech](../entries/omit-tech.yml)                                                                                 | omit                                     | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [omit-utility-type-typescript](../entries/omit-utility-type-typescript.yml)                                           | Omit                                     | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [on-every-exit-path-workflow](../entries/on-every-exit-path-workflow.yml)                                             | on every exit path                       | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [one-handed-combat](../entries/one-handed-combat.yml)                                                                 | one-handed                               | WF       | 0        | 保留现有辅助与关联                                   |
| [one-way-data-flow-frontend](../entries/one-way-data-flow-frontend.yml)                                               | one-way data flow                        | —        | 1        | 保留现有辅助与关联                                   |
| [onion-model-koa](../entries/onion-model-koa.yml)                                                                     | onion model                              | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [open-general](../entries/open-general.yml)                                                                           | open                                     | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [openclaw-ai](../entries/openclaw-ai.yml)                                                                             | OpenClaw                                 | WF       | 0        | 保留现有辅助与关联                                   |
| [operand-tech](../entries/operand-tech.yml)                                                                           | operand                                  | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [opposite-shore-terrain](../entries/opposite-shore-terrain.yml)                                                       | opposite shore                           | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [option-chain-finance](../entries/option-chain-finance.yml)                                                           | option chain                             | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [option-finance](../entries/option-finance.yml)                                                                       | option                                   | —        | 1        | 保留现有辅助与关联                                   |
| [optional-general](../entries/optional-general.yml)                                                                   | optional                                 | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [optional-parameter-typescript](../entries/optional-parameter-typescript.yml)                                         | optional parameter                       | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [optional-property-typescript](../entries/optional-property-typescript.yml)                                           | optional property                        | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [orchestration-ai](../entries/orchestration-ai.yml)                                                                   | orchestration                            | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [order-finance](../entries/order-finance.yml)                                                                         | order                                    | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [ordered-list-html](../entries/ordered-list-html.yml)                                                                 | ordered list                             | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [origin-general](../entries/origin-general.yml)                                                                       | origin                                   | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [orijen-pet-food](../entries/orijen-pet-food.yml)                                                                     | ORIJEN                                   | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [out-of-order-tech](../entries/out-of-order-tech.yml)                                                                 | out-of-order                             | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [outstanding-tech](../entries/outstanding-tech.yml)                                                                   | outstanding                              | N        | 1        | 保留现有辅助与关联                                   |
| [overflow-css](../entries/overflow-css.yml)                                                                           | overflow                                 | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [overlap-prevention-workflow](../entries/overlap-prevention-workflow.yml)                                             | overlap prevention                       | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [overlap-tech](../entries/overlap-tech.yml)                                                                           | overlap                                  | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [overload-tech](../entries/overload-tech.yml)                                                                         | overload                                 | N        | 0        | 保留现有辅助与关联                                   |
| [overwrite-tech](../entries/overwrite-tech.yml)                                                                       | overwrite                                | WF       | 0        | 保留现有辅助与关联                                   |
| [own-property-javascript](../entries/own-property-javascript.yml)                                                     | own property                             | N        | 1        | 保留现有辅助与关联                                   |
| [ownership-token-workflow](../entries/ownership-token-workflow.yml)                                                   | ownership token                          | —        | 1        | 保留现有辅助与关联                                   |
| [package-go](../entries/package-go.yml)                                                                               | package                                  | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [pad-general](../entries/pad-general.yml)                                                                             | pad                                      | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [padding-css](../entries/padding-css.yml)                                                                             | padding                                  | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [paella-cooking](../entries/paella-cooking.yml)                                                                       | paella                                   | ETY      | 0        | 保留现有辅助与关联                                   |
| [pagination-bootstrap](../entries/pagination-bootstrap.yml)                                                           | pagination                               | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [paint-browser-rendering](../entries/paint-browser-rendering.yml)                                                     | paint                                    | —        | 1        | 保留现有辅助与关联                                   |
| [panic-go](../entries/panic-go.yml)                                                                                   | panic                                    | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [paraglider-zelda](../entries/paraglider-zelda.yml)                                                                   | Paraglider                               | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [paragraph-html](../entries/paragraph-html.yml)                                                                       | paragraph                                | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [parameter-tech](../entries/parameter-tech.yml)                                                                       | parameter                                | N        | 1        | 保留现有辅助与关联                                   |
| [part-of-speech-language](../entries/part-of-speech-language.yml)                                                     | part of speech                           | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [partial-clone-git](../entries/partial-clone-git.yml)                                                                 | partial clone                            | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [partial-utility-type-typescript](../entries/partial-utility-type-typescript.yml)                                     | Partial                                  | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [pass-with-the-fix-workflow](../entries/pass-with-the-fix-workflow.yml)                                               | pass with the fix                        | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [passage-architecture](../entries/passage-architecture.yml)                                                           | passage                                  | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [pasta-food](../entries/pasta-food.yml)                                                                               | pasta                                    | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [patch-tech](../entries/patch-tech.yml)                                                                               | patch                                    | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [pause-workflow](../entries/pause-workflow.yml)                                                                       | pause                                    | —        | 1        | 保留现有辅助与关联                                   |
| [paya-zelda](../entries/paya-zelda.yml)                                                                               | Paya                                     | N        | 0        | 保留现有辅助与关联                                   |
| [pending-tech](../entries/pending-tech.yml)                                                                           | pending                                  | —        | 1        | 保留现有辅助与关联                                   |
| [pepper-food](../entries/pepper-food.yml)                                                                             | pepper                                   | —        | 1        | 保留现有辅助与关联                                   |
| [pepper-vegetable](../entries/pepper-vegetable.yml)                                                                   | pepper                                   | —        | 1        | 保留现有辅助与关联                                   |
| [perfect-dodge-zelda](../entries/perfect-dodge-zelda.yml)                                                             | Perfect Dodge                            | —        | 1        | 保留现有辅助与关联                                   |
| [perfect-guard-zelda](../entries/perfect-guard-zelda.yml)                                                             | Perfect Guard                            | —        | 1        | 保留现有辅助与关联                                   |
| [permission-tech](../entries/permission-tech.yml)                                                                     | permission                               | —        | 1        | 保留现有辅助与关联                                   |
| [persist-tech](../entries/persist-tech.yml)                                                                           | persist                                  | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [persistent-data-docker](../entries/persistent-data-docker.yml)                                                       | persistent data                          | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [phantom-ganon-zelda](../entries/phantom-ganon-zelda.yml)                                                             | Phantom Ganon                            | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [phenomenon-story](../entries/phenomenon-story.yml)                                                                   | phenomenon                               | WF       | 0        | 保留现有辅助与关联                                   |
| [pick-up-general](../entries/pick-up-general.yml)                                                                     | pick up                                  | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [pick-utility-type-typescript](../entries/pick-utility-type-typescript.yml)                                           | Pick                                     | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [pilot-general](../entries/pilot-general.yml)                                                                         | pilot                                    | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [pinned-checkout-workflow](../entries/pinned-checkout-workflow.yml)                                                   | pinned checkout                          | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [pit-terrain](../entries/pit-terrain.yml)                                                                             | pit                                      | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [plan-general](../entries/plan-general.yml)                                                                           | plan                                     | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [plateau-terrain](../entries/plateau-terrain.yml)                                                                     | plateau                                  | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [platform-architecture](../entries/platform-architecture.yml)                                                         | platform                                 | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [plugin-tech](../entries/plugin-tech.yml)                                                                             | plugin                                   | —        | 1        | 保留现有辅助与关联                                   |
| [pointer-go](../entries/pointer-go.yml)                                                                               | pointer                                  | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [pointer-receiver-go](../entries/pointer-receiver-go.yml)                                                             | pointer receiver                         | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [policy-tech](../entries/policy-tech.yml)                                                                             | policy                                   | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [popover-bootstrap](../entries/popover-bootstrap.yml)                                                                 | popover                                  | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [pork-food](../entries/pork-food.yml)                                                                                 | pork                                     | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [porridge-cooking](../entries/porridge-cooking.yml)                                                                   | porridge                                 | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [port-networking](../entries/port-networking.yml)                                                                     | port                                     | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [port-publishing-docker](../entries/port-publishing-docker.yml)                                                       | port publishing                          | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [portable-pot-zelda](../entries/portable-pot-zelda.yml)                                                               | Portable Pot                             | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [portal-react](../entries/portal-react.yml)                                                                           | portal                                   | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [portfolio-finance](../entries/portfolio-finance.yml)                                                                 | portfolio                                | —        | 0        | R09、ETY01                                           |
| [position-finance](../entries/position-finance.yml)                                                                   | position                                 | —        | 1        | R09 反向、N03                                        |
| [possessed-combat](../entries/possessed-combat.yml)                                                                   | possessed                                | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [postorder-traversal-tree](../entries/postorder-traversal-tree.yml)                                                   | postorder traversal                      | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [pot-general](../entries/pot-general.yml)                                                                             | pot                                      | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [potato-food](../entries/potato-food.yml)                                                                             | potato                                   | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [pouch-ui](../entries/pouch-ui.yml)                                                                                   | pouch                                    | —        | 1        | 保留现有辅助与关联                                   |
| [powerful-general](../entries/powerful-general.yml)                                                                   | powerful                                 | WF       | 0        | 保留现有辅助与关联                                   |
| [pre-existing-tech](../entries/pre-existing-tech.yml)                                                                 | pre-existing                             | WF       | 0        | 保留现有辅助与关联                                   |
| [pre-market-finance](../entries/pre-market-finance.yml)                                                               | pre-market                               | WF       | 0        | 保留现有辅助与关联                                   |
| [precede-tech](../entries/precede-tech.yml)                                                                           | precede                                  | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [prefilter-tech](../entries/prefilter-tech.yml)                                                                       | prefilter                                | WF       | 0        | 保留现有辅助与关联                                   |
| [prefix-match-nginx](../entries/prefix-match-nginx.yml)                                                               | prefix match                             | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [prefix-priority-nginx](../entries/prefix-priority-nginx.yml)                                                         | ^~                                       | —        | 0        | 暂缓，符号读法需单独核实，现有词条和音频状态维持基线 |
| [preflight-request-web](../entries/preflight-request-web.yml)                                                         | preflight request                        | N        | 0        | 保留现有辅助与关联                                   |
| [preserve-tech](../entries/preserve-tech.yml)                                                                         | preserve                                 | N        | 1        | 保留现有辅助与关联                                   |
| [prevent-tech](../entries/prevent-tech.yml)                                                                           | prevent                                  | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [primary-checkout-workflow](../entries/primary-checkout-workflow.yml)                                                 | primary checkout                         | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [prime-food](../entries/prime-food.yml)                                                                               | prime                                    | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [primitive-javascript](../entries/primitive-javascript.yml)                                                           | primitive                                | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [prince-general](../entries/prince-general.yml)                                                                       | prince                                   | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [princess-general](../entries/princess-general.yml)                                                                   | princess                                 | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [princess-zelda-zelda](../entries/princess-zelda-zelda.yml)                                                           | Princess Zelda                           | N        | 0        | 保留现有辅助与关联                                   |
| [pristine-general](../entries/pristine-general.yml)                                                                   | pristine                                 | N        | 1        | 保留现有辅助与关联                                   |
| [pro-professional](../entries/pro-professional.yml)                                                                   | pro                                      | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [probe-tech](../entries/probe-tech.yml)                                                                               | probe                                    | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [profile-ui](../entries/profile-ui.yml)                                                                               | profile                                  | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [profiling-tech](../entries/profiling-tech.yml)                                                                       | profiling                                | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [profit-and-loss-finance](../entries/profit-and-loss-finance.yml)                                                     | profit and loss                          | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [profit-finance](../entries/profit-finance.yml)                                                                       | profit                                   | —        | 1        | 保留现有辅助与关联                                   |
| [progress-bar-bootstrap](../entries/progress-bar-bootstrap.yml)                                                       | progress bar                             | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [progression-game](../entries/progression-game.yml)                                                                   | progression                              | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [progressive-enhancement-web](../entries/progressive-enhancement-web.yml)                                             | progressive enhancement                  | —        | 1        | 保留现有辅助与关联                                   |
| [project-onboarding-workflow](../entries/project-onboarding-workflow.yml)                                             | project onboarding                       | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [promise-all-javascript](../entries/promise-all-javascript.yml)                                                       | Promise.all                              | N        | 0        | 保留现有辅助与关联                                   |
| [promise-javascript](../entries/promise-javascript.yml)                                                               | Promise                                  | —        | 1        | 保留现有辅助与关联                                   |
| [prompt-ai](../entries/prompt-ai.yml)                                                                                 | prompt                                   | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [pronunciation-general](../entries/pronunciation.yml)                                                                 | pronunciation                            | N        | 1        | 保留现有辅助与关联                                   |
| [propagate-tech](../entries/propagate-tech.yml)                                                                       | propagate                                | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [property-css](../entries/property-css.yml)                                                                           | CSS property                             | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [property-descriptor-javascript](../entries/property-descriptor-javascript.yml)                                       | property descriptor                      | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [property-tech](../entries/property-tech.yml)                                                                         | property                                 | —        | 1        | 保留现有辅助与关联                                   |
| [property-value-css](../entries/property-value-css.yml)                                                               | property value                           | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [prophecy-story](../entries/prophecy-story.yml)                                                                       | prophecy                                 | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [props-react](../entries/props-react.yml)                                                                             | props                                    | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [propulsion-engineering](../entries/propulsion-engineering.yml)                                                       | propulsion                               | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [prototype-chain-javascript](../entries/prototype-chain-javascript.yml)                                               | prototype chain                          | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [prototype-setter-javascript](../entries/prototype-setter-javascript.yml)                                             | prototype setter                         | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [provider-ai](../entries/provider-ai.yml)                                                                             | provider                                 | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [proxy-pass-nginx](../entries/proxy-pass-nginx.yml)                                                                   | proxy_pass                               | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [prune-tech](../entries/prune-tech.yml)                                                                               | prune                                    | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [pseudo-class-css](../entries/pseudo-class-css.yml)                                                                   | pseudo-class                             | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [pseudo-element-css](../entries/pseudo-element-css.yml)                                                               | pseudo-element                           | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [puffshroom-zelda](../entries/puffshroom-zelda.yml)                                                                   | Puffshroom                               | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [purah-pad-zelda](../entries/purah-pad-zelda.yml)                                                                     | Purah Pad                                | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [purah-zelda](../entries/purah-zelda.yml)                                                                             | Purah                                    | N        | 0        | 保留现有辅助与关联                                   |
| [pure-component-react](../entries/pure-component-react.yml)                                                           | pure component                           | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [purina-pro-plan-pet-food](../entries/purina-pro-plan-pet-food.yml)                                                   | Purina Pro Plan                          | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [put-option-finance](../entries/put-option-finance.yml)                                                               | put option                               | —        | 1        | 保留现有辅助与关联                                   |
| [quadratic-tech](../entries/quadratic-tech.yml)                                                                       | quadratic                                | —        | 1        | 保留现有辅助与关联                                   |
| [quantity-general](../entries/quantity-general.yml)                                                                   | quantity                                 | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [quarter-finance](../entries/quarter-finance.yml)                                                                     | quarter                                  | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [queen-general](../entries/queen-general.yml)                                                                         | queen                                    | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [queen-gibdo-zelda](../entries/queen-gibdo-zelda.yml)                                                                 | Queen Gibdo                              | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [query-tech](../entries/query-tech.yml)                                                                               | query                                    | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [quest-game](../entries/quest-game.yml)                                                                               | quest                                    | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [quicksave-game](../entries/quicksave-game.yml)                                                                       | quicksave                                | N        | 1        | 保留现有辅助与关联                                   |
| [quicksort-algorithm](../entries/quicksort-algorithm.yml)                                                             | quicksort                                | WF       | 0        | 保留现有辅助与关联                                   |
| [race-detector-go](../entries/race-detector-go.yml)                                                                   | race detector                            | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [radish-food](../entries/radish-food.yml)                                                                             | radish                                   | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [raft-watercraft](../entries/raft-watercraft.yml)                                                                     | raft                                     | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [railway-track-engineering](../entries/railway-track-engineering.yml)                                                 | railway track                            | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [range-go](../entries/range-go.yml)                                                                                   | range                                    | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [ranged-combat](../entries/ranged-combat.yml)                                                                         | ranged                                   | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [rather-than-workflow](../entries/rather-than-workflow.yml)                                                           | rather than                              | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [rauru-zelda](../entries/rauru-zelda.yml)                                                                             | Rauru                                    | N        | 0        | 保留现有辅助与关联                                   |
| [raw-food](../entries/raw-food.yml)                                                                                   | raw                                      | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [raw-gourmet-meat-zelda](../entries/raw-gourmet-meat-zelda.yml)                                                       | Raw Gourmet Meat                         | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [raw-prime-meat-zelda](../entries/raw-prime-meat-zelda.yml)                                                           | Raw Prime Meat                           | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [re-entry-tech](../entries/re-entry-tech.yml)                                                                         | re-entry                                 | WF       | 0        | 保留现有辅助与关联                                   |
| [react-library](../entries/react-library.yml)                                                                         | React                                    | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [reactive-dependency-react](../entries/reactive-dependency-react.yml)                                                 | reactive dependency                      | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [read-only-control-html](../entries/read-only-control-html.yml)                                                       | read-only control                        | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [read-only-workflow](../entries/read-only-workflow.yml)                                                               | read-only                                | WF       | 0        | 保留现有辅助与关联                                   |
| [readonly-typescript](../entries/readonly-typescript.yml)                                                             | readonly                                 | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [ready-for-review-workflow](../entries/ready-for-review-workflow.yml)                                                 | ready for review                         | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [reasoning-ai](../entries/reasoning-ai.yml)                                                                           | reasoning                                | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [reattach-general](../entries/reattach-general.yml)                                                                   | reattach                                 | WF       | 0        | 保留现有辅助与关联                                   |
| [recall-general](../entries/recall-general.yml)                                                                       | recall                                   | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [recap-workflow](../entries/recap-workflow.yml)                                                                       | recap                                    | —        | 1        | 保留现有辅助与关联                                   |
| [receiver-go](../entries/receiver-go.yml)                                                                             | receiver                                 | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [recharge-general](../entries/recharge-general.yml)                                                                   | recharge                                 | WF       | 0        | 保留现有辅助与关联                                   |
| [recipe-cooking](../entries/recipe-cooking.yml)                                                                       | recipe                                   | ETY      | 1        | 保留现有辅助与关联                                   |
| [reckless-general](../entries/reckless-general.yml)                                                                   | reckless                                 | N        | 1        | 保留现有辅助与关联                                   |
| [reconciliation-react](../entries/reconciliation-react.yml)                                                           | reconciliation                           | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [record-utility-type-typescript](../entries/record-utility-type-typescript.yml)                                       | Record                                   | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [recover-go](../entries/recover-go.yml)                                                                               | recover                                  | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [recover-story](../entries/recover-story.yml)                                                                         | recover                                  | —        | 1        | 保留现有辅助与关联                                   |
| [recovery-tech](../entries/recovery-tech.yml)                                                                         | recovery                                 | —        | 1        | 保留现有辅助与关联                                   |
| [recursion-programming](../entries/recursion-programming.yml)                                                         | recursion                                | —        | 1        | 保留现有辅助与关联                                   |
| [reduced-motion-css](../entries/reduced-motion-css.yml)                                                               | reduced motion                           | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [reducer-react](../entries/reducer-react.yml)                                                                         | reducer                                  | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [ref-react](../entries/ref-react.yml)                                                                                 | ref                                      | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [refinery-architecture](../entries/refinery-architecture.yml)                                                         | refinery                                 | WF       | 0        | 保留现有辅助与关联                                   |
| [regex-match-nginx](../entries/regex-match-nginx.yml)                                                                 | regex match                              | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [regional-general](../entries/regional-general.yml)                                                                   | regional                                 | WF       | 0        | 保留现有辅助与关联                                   |
| [regional-phenomena-zelda](../entries/regional-phenomena-zelda.yml)                                                   | Regional Phenomena                       | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [register-tech](../entries/register-tech.yml)                                                                         | register                                 | N        | 1        | 保留现有辅助与关联                                   |
| [registry-docker](../entries/registry-docker.yml)                                                                     | registry                                 | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [registry-tech](../entries/registry-tech.yml)                                                                         | registry                                 | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [regression-tech](../entries/regression-tech.yml)                                                                     | regression                               | ETY      | 0        | 保留现有辅助与关联                                   |
| [regression-test-tech](../entries/regression-test-tech.yml)                                                           | regression test                          | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [reject-tech](../entries/reject-tech.yml)                                                                             | reject                                   | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [related-words-language](../entries/related-words-language.yml)                                                       | related words                            | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [relative-positioning-css](../entries/relative-positioning-css.yml)                                                   | relative positioning                     | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [relative-unit-css](../entries/relative-unit-css.yml)                                                                 | relative unit                            | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [release-workflow](../entries/release-workflow.yml)                                                                   | release                                  | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [reload-nginx](../entries/reload-nginx.yml)                                                                           | reload                                   | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [rem-css](../entries/rem-css.yml)                                                                                     | rem                                      | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [remain-intact-workflow](../entries/remain-intact-workflow.yml)                                                       | remain intact                            | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [remain-unchanged-workflow](../entries/remain-unchanged-workflow.yml)                                                 | remain unchanged                         | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [remote-tracking-branch-git](../entries/remote-tracking-branch-git.yml)                                               | remote-tracking branch                   | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [render-react](../entries/render-react.yml)                                                                           | render                                   | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [reported-behavior-workflow](../entries/reported-behavior-workflow.yml)                                               | reported behavior                        | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [repository-git](../entries/repository-git.yml)                                                                       | repository                               | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [reproduce-tech](../entries/reproduce-tech.yml)                                                                       | reproduce                                | N        | 1        | 保留现有辅助与关联                                   |
| [reproducible-tech](../entries/reproducible-tech.yml)                                                                 | reproducible                             | N        | 1        | 保留现有辅助与关联                                   |
| [required-field-html](../entries/required-field-html.yml)                                                             | required field                           | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [rerender-react](../entries/rerender-react.yml)                                                                       | re-render                                | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [reserve-workflow](../entries/reserve-workflow.yml)                                                                   | reserve                                  | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [reserved-prefix-tech](../entries/reserved-prefix-tech.yml)                                                           | reserved prefix                          | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [resilient-combat](../entries/resilient-combat.yml)                                                                   | resilient                                | ETY / N  | 1        | 保留现有辅助与关联                                   |
| [resistance-general](../entries/resistance-general.yml)                                                               | resistance                               | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [resolve-tech](../entries/resolve-tech.yml)                                                                           | resolve                                  | N        | 1        | 保留现有辅助与关联                                   |
| [resolve-workflow](../entries/resolve-workflow.yml)                                                                   | resolve                                  | N        | 1        | 保留现有辅助与关联                                   |
| [responsive-container-bootstrap](../entries/responsive-container-bootstrap.yml)                                       | responsive container                     | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [responsive-design-css](../entries/responsive-design-css.yml)                                                         | responsive design                        | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [responsive-image-html](../entries/responsive-image-html.yml)                                                         | responsive image                         | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [responsive-utility-bootstrap](../entries/responsive-utility-bootstrap.yml)                                           | responsive utility                       | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [rest-parameter-typescript](../entries/rest-parameter-typescript.yml)                                                 | rest parameter                           | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [restore-tech](../entries/restore-tech.yml)                                                                           | restore                                  | —        | 1        | 保留现有辅助与关联                                   |
| [restrict-tech](../entries/restrict-tech.yml)                                                                         | restrict                                 | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [result-general](../entries/result-general.yml)                                                                       | result                                   | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [resume-workflow](../entries/resume-workflow.yml)                                                                     | resume                                   | —        | 0        | R10                                                  |
| [resumption-workflow](../entries/resumption-workflow.yml)                                                             | resumption                               | WF       | 1        | R10 反向                                             |
| [retain-tech](../entries/retain-tech.yml)                                                                             | retain                                   | N        | 1        | 保留现有辅助与关联                                   |
| [retry-workflow](../entries/retry-workflow.yml)                                                                       | retry                                    | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [return-finance](../entries/return-finance.yml)                                                                       | return                                   | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [return-type-tech](../entries/return-type-tech.yml)                                                                   | return type                              | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [returntype-utility-type-typescript](../entries/returntype-utility-type-typescript.yml)                               | ReturnType                               | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [reuse-general](../entries/reuse-general.yml)                                                                         | reuse                                    | —        | 0        | WF05                                                 |
| [revenue-finance](../entries/revenue-finance.yml)                                                                     | revenue                                  | —        | 1        | 保留现有辅助与关联                                   |
| [reverse-proxy-nginx](../entries/reverse-proxy-nginx.yml)                                                             | reverse proxy                            | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [review-workflow](../entries/review-workflow.yml)                                                                     | review                                   | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [revision-general](../entries/revision-general.yml)                                                                   | revision                                 | N        | 1        | 保留现有辅助与关联                                   |
| [revision-git](../entries/revision-git.yml)                                                                           | revision                                 | N        | 1        | 保留现有辅助与关联                                   |
| [rice-food](../entries/rice-food.yml)                                                                                 | rice                                     | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [ridge-terrain](../entries/ridge-terrain.yml)                                                                         | ridge                                    | —        | 1        | 保留现有辅助与关联                                   |
| [riju-zelda](../entries/riju-zelda.yml)                                                                               | Riju                                     | N        | 0        | 保留现有辅助与关联                                   |
| [ring-general](../entries/ring-general.yml)                                                                           | ring                                     | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [risotto-cooking](../entries/risotto-cooking.yml)                                                                     | risotto                                  | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [rito-village-zelda](../entries/rito-village-zelda.yml)                                                               | Rito Village                             | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [rito-zelda](../entries/rito-zelda.yml)                                                                               | Rito                                     | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [river-terrain](../entries/river-terrain.yml)                                                                         | river                                    | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [roast-food](../entries/roast-food.yml)                                                                               | roast                                    | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [robbie-zelda](../entries/robbie-zelda.yml)                                                                           | Robbie                                   | N        | 0        | 保留现有辅助与关联                                   |
| [robe-clothing](../entries/robe-clothing.yml)                                                                         | robe                                     | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [rocket-engineering](../entries/rocket-engineering.yml)                                                               | rocket                                   | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [rod-general](../entries/rod-general.yml)                                                                             | rod                                      | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [root-general](../entries/root-general.yml)                                                                           | root                                     | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [root-nginx](../entries/root-nginx.yml)                                                                               | root                                     | —        | 0        | R04 反向                                             |
| [rotate-general](../entries/rotate-general.yml)                                                                       | rotate                                   | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [round-trip-tech](../entries/round-trip-tech.yml)                                                                     | round trip                               | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [row-bootstrap](../entries/row-bootstrap.yml)                                                                         | row                                      | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [row-span-html](../entries/row-span-html.yml)                                                                         | row span                                 | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [royal-broadsword-zelda](../entries/royal-broadsword-zelda.yml)                                                       | Royal Broadsword                         | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [royal-canin-pet-food](../entries/royal-canin-pet-food.yml)                                                           | Royal Canin                              | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [royal-general](../entries/royal-general.yml)                                                                         | royal                                    | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [ruins-architecture](../entries/ruins-architecture.yml)                                                               | ruins                                    | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [rules-of-hooks-react](../entries/rules-of-hooks-react.yml)                                                           | Rules of Hooks                           | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [run-budget-workflow](../entries/run-budget-workflow.yml)                                                             | run budget                               | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [run-in-parallel-workflow](../entries/run-in-parallel-workflow.yml)                                                   | run in parallel                          | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [rune-go](../entries/rune-go.yml)                                                                                     | rune                                     | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [runtime-evaluation-tech](../entries/runtime-evaluation-tech.yml)                                                     | runtime evaluation                       | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [runtime-tech](../entries/runtime-tech.yml)                                                                           | runtime                                  | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [runway-flight](../entries/runway-flight.yml)                                                                         | runway                                   | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [rupee-game](../entries/rupee-game.yml)                                                                               | rupee                                    | ETY      | 0        | 保留现有辅助与关联                                   |
| [rusty-general](../entries/rusty-general.yml)                                                                         | rusty                                    | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [ruthless-general](../entries/ruthless-general.yml)                                                                   | ruthless                                 | WF       | 0        | 保留现有辅助与关联                                   |
| [sacred-general](../entries/sacred-general.yml)                                                                       | sacred                                   | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [sacrifice-story](../entries/sacrifice-story.yml)                                                                     | sacrifice                                | ETY      | 0        | 保留现有辅助与关联                                   |
| [sage-general](../entries/sage-general.yml)                                                                           | sage                                     | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [sail-watercraft](../entries/sail-watercraft.yml)                                                                     | sail                                     | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [salmon-food](../entries/salmon-food.yml)                                                                             | salmon                                   | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [salt-food](../entries/salt-food.yml)                                                                                 | salt                                     | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [same-origin-policy-web](../entries/same-origin-policy-web.yml)                                                       | same-origin policy                       | —        | 1        | 保留现有辅助与关联                                   |
| [sandbox-tech](../entries/sandbox-tech.yml)                                                                           | sandbox                                  | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [sanitize-tech](../entries/sanitize-tech.yml)                                                                         | sanitize                                 | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [sass-mixin-bootstrap](../entries/sass-mixin-bootstrap.yml)                                                           | Sass mixin                               | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [sass-variable-bootstrap](../entries/sass-variable-bootstrap.yml)                                                     | Sass variable                            | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [sauce-food](../entries/sauce-food.yml)                                                                               | sauce                                    | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [saute-cooking](../entries/saute-cooking.yml)                                                                         | sauté                                    | ETY      | 0        | 保留现有辅助与关联                                   |
| [savage-general](../entries/savage-general.yml)                                                                       | savage                                   | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [save-game](../entries/save-game.yml)                                                                                 | save                                     | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [science-general](../entries/science-general.yml)                                                                     | science                                  | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [scimitar-of-the-seven-zelda](../entries/scimitar-of-the-seven-zelda.yml)                                             | Scimitar of the Seven                    | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [scimitar-weapon](../entries/scimitar-weapon.yml)                                                                     | scimitar                                 | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [scope-lexical](../entries/scope-lexical.yml)                                                                         | scope                                    | N        | 1        | 保留现有辅助与关联                                   |
| [scope-workflow](../entries/scope-workflow.yml)                                                                       | scope                                    | N        | 1        | 保留现有辅助与关联                                   |
| [screen-display](../entries/screen-display.yml)                                                                       | screen                                   | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [script-html](../entries/script-html.yml)                                                                             | script                                   | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [seafood-food](../entries/seafood-food.yml)                                                                           | seafood                                  | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [sear-cooking](../entries/sear-cooking.yml)                                                                           | sear                                     | N        | 1        | 保留现有辅助与关联                                   |
| [secret-stone-zelda](../entries/secret-stone-zelda.yml)                                                               | Secret Stone                             | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [secret-story](../entries/secret-story.yml)                                                                           | secret                                   | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [seed-general](../entries/seed-general.yml)                                                                           | seed                                     | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [seek-general](../entries/seek-general.yml)                                                                           | seek                                     | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [seized-construct-zelda](../entries/seized-construct-zelda.yml)                                                       | Seized Construct                         | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [select-general](../entries/select-general.yml)                                                                       | select                                   | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [select-go](../entries/select-go.yml)                                                                                 | select                                   | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [selector-css](../entries/selector-css.yml)                                                                           | selector                                 | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [selfless-general](../entries/selfless-general.yml)                                                                   | selfless                                 | WF       | 0        | 保留现有辅助与关联                                   |
| [semantic-html](../entries/semantic-html.yml)                                                                         | semantic HTML                            | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [semantic-versioning-software](../entries/semantic-versioning-software.yml)                                           | semantic versioning                      | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [sentinel-error-go](../entries/sentinel-error-go.yml)                                                                 | sentinel error                           | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [serialization-tech](../entries/serialization-tech.yml)                                                               | serialization                            | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [server-block-nginx](../entries/server-block-nginx.yml)                                                               | server                                   | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [server-rendering-react](../entries/server-rendering-react.yml)                                                       | server rendering                         | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [service-compose](../entries/service-compose.yml)                                                                     | service                                  | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [session-tech](../entries/session-tech.yml)                                                                           | session                                  | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [settlement-architecture](../entries/settlement-architecture.yml)                                                     | settlement                               | WF       | 0        | 保留现有辅助与关联                                   |
| [setup-tech](../entries/setup-tech.yml)                                                                               | setup                                    | WF       | 0        | 保留现有辅助与关联                                   |
| [shallow-copy-javascript](../entries/shallow-copy-javascript.yml)                                                     | shallow copy                             | —        | 1        | 保留现有辅助与关联                                   |
| [share-finance](../entries/share-finance.yml)                                                                         | share                                    | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [shareholder-finance](../entries/shareholder-finance.yml)                                                             | shareholder                              | WF       | 0        | 保留现有辅助与关联                                   |
| [sheba-pet-food](../entries/sheba-pet-food.yml)                                                                       | Sheba                                    | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [sheikah-zelda](../entries/sheikah-zelda.yml)                                                                         | Sheikah                                  | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [shield-weapon](../entries/shield-weapon.yml)                                                                         | shield                                   | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [shirt-clothing](../entries/shirt-clothing.yml)                                                                       | shirt                                    | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [shock-emitter-zelda](../entries/shock-emitter-zelda.yml)                                                             | Shock Emitter                            | WF       | 0        | 保留现有辅助与关联                                   |
| [shock-resistance-combat](../entries/shock-resistance-combat.yml)                                                     | shock resistance                         | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [short-variable-declaration-go](../entries/short-variable-declaration-go.yml)                                         | short variable declaration               | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [shorthand-javascript](../entries/shorthand-javascript.yml)                                                           | shorthand                                | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [shorthand-property-css](../entries/shorthand-property-css.yml)                                                       | shorthand property                       | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [shrine-architecture](../entries/shrine-architecture.yml)                                                             | shrine                                   | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [shrine-quests-zelda](../entries/shrine-quests-zelda.yml)                                                             | Shrine Quests                            | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [side-adventures-zelda](../entries/side-adventures-zelda.yml)                                                         | Side Adventures                          | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [side-hop-combat](../entries/side-hop-combat.yml)                                                                     | side hop                                 | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [side-quests-zelda](../entries/side-quests-zelda.yml)                                                                 | Side Quests                              | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [sidon-zelda](../entries/sidon-zelda.yml)                                                                             | Sidon                                    | N        | 0        | 保留现有辅助与关联                                   |
| [signed-zero-tech](../entries/signed-zero-tech.yml)                                                                   | signed zero                              | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [silent-princess-zelda](../entries/silent-princess-zelda.yml)                                                         | Silent Princess                          | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [simmer-cooking](../entries/simmer-cooking.yml)                                                                       | simmer                                   | N        | 1        | 保留现有辅助与关联                                   |
| [skewer-cooking](../entries/skewer-cooking.yml)                                                                       | skewer                                   | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [skill-ai](../entries/skill-ai.yml)                                                                                   | skill                                    | —        | 1        | 保留现有辅助与关联                                   |
| [skilled-general](../entries/skilled-general.yml)                                                                     | skilled                                  | WF       | 0        | 保留现有辅助与关联                                   |
| [sky-island-zelda](../entries/sky-island-zelda.yml)                                                                   | sky island                               | WF       | 0        | 保留现有辅助与关联                                   |
| [sky-terrain](../entries/sky-terrain.yml)                                                                             | sky                                      | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [skyshroom-zelda](../entries/skyshroom-zelda.yml)                                                                     | Skyshroom                                | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [skyview-tower-zelda](../entries/skyview-tower-zelda.yml)                                                             | Skyview Tower                            | WF       | 0        | 保留现有辅助与关联                                   |
| [sled-engineering](../entries/sled-engineering.yml)                                                                   | sled                                     | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [sleeve-clothing](../entries/sleeve-clothing.yml)                                                                     | sleeve                                   | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [slice-go](../entries/slice-go.yml)                                                                                   | slice                                    | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [slip-general](../entries/slip-general.yml)                                                                           | slip                                     | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [slip-resistance-zelda](../entries/slip-resistance-zelda.yml)                                                         | Slip Resistance                          | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [slope-terrain](../entries/slope-terrain.yml)                                                                         | slope                                    | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [small-wheel-zelda](../entries/small-wheel-zelda.yml)                                                                 | Small Wheel                              | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [snapshot-tech](../entries/snapshot-tech.yml)                                                                         | snapshot                                 | —        | 1        | 保留现有辅助与关联                                   |
| [sneaky-general](../entries/sneaky-general.yml)                                                                       | sneaky                                   | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [snowquill-tunic-zelda](../entries/snowquill-tunic-zelda.yml)                                                         | Snowquill Tunic                          | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [soldier-construct-zelda](../entries/soldier-construct-zelda.yml)                                                     | Soldier Construct                        | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [soldier-general](../entries/soldier-general.yml)                                                                     | soldier                                  | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [soldiers-broadsword-zelda](../entries/soldiers-broadsword-zelda.yml)                                                 | Soldier's Broadsword                     | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [sonia-zelda](../entries/sonia-zelda.yml)                                                                             | Sonia                                    | N        | 0        | 保留现有辅助与关联                                   |
| [sort-general](../entries/sort-general.yml)                                                                           | sort                                     | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [soundness-type-system](../entries/soundness-type-system.yml)                                                         | soundness                                | N        | 1        | 保留现有辅助与关联                                   |
| [source-code-tech](../entries/source-code-tech.yml)                                                                   | source code                              | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [source-map-tech](../entries/source-map-tech.yml)                                                                     | source map                               | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [source-order-css](../entries/source-order-css.yml)                                                                   | source order                             | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [source-set-html](../entries/source-set-html.yml)                                                                     | source set                               | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [spacing-utility-bootstrap](../entries/spacing-utility-bootstrap.yml)                                                 | spacing utility                          | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [spaulder-clothing](../entries/spaulder-clothing.yml)                                                                 | spaulder                                 | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [spear-weapon](../entries/spear-weapon.yml)                                                                           | spear                                    | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [special-controls-zelda](../entries/special-controls-zelda.yml)                                                       | Special Controls                         | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [specificity-css](../entries/specificity-css.yml)                                                                     | specificity                              | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [speculative-tech](../entries/speculative-tech.yml)                                                                   | speculative                              | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [speed-up-zelda](../entries/speed-up-zelda.yml)                                                                       | Speed Up                                 | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [spice-food](../entries/spice-food.yml)                                                                               | spice                                    | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [spicy-general](../entries/spicy-general.yml)                                                                         | spicy                                    | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [spicy-pepper-zelda](../entries/spicy-pepper-zelda.yml)                                                               | Spicy Pepper                             | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [spinner-bootstrap](../entries/spinner-bootstrap.yml)                                                                 | spinner                                  | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [spirit-story](../entries/spirit-story.yml)                                                                           | spirit                                   | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [spirit-temple-zelda](../entries/spirit-temple-zelda.yml)                                                             | Spirit Temple                            | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [spring-mechanical](../entries/spring-mechanical.yml)                                                                 | spring                                   | —        | 1        | 保留现有辅助与关联                                   |
| [spring-water](../entries/spring-water.yml)                                                                           | spring                                   | N        | 1        | 保留现有辅助与关联                                   |
| [sprint-general](../entries/sprint-general.yml)                                                                       | sprint                                   | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [stabilizer-engineering](../entries/stabilizer-engineering.yml)                                                       | stabilizer                               | WF       | 0        | 保留现有辅助与关联                                   |
| [stable-architecture](../entries/stable-architecture.yml)                                                             | stable                                   | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [stack-data-structure](../entries/stack-data-structure.yml)                                                           | stack                                    | —        | 1        | 保留现有辅助与关联                                   |
| [stack-trace-tech](../entries/stack-trace-tech.yml)                                                                   | stack trace                              | —        | 1        | 保留现有辅助与关联                                   |
| [stacking-context-css](../entries/stacking-context-css.yml)                                                           | stacking context                         | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [staged-git](../entries/staged-git.yml)                                                                               | staged                                   | N        | 1        | 保留现有辅助与关联                                   |
| [staircase-architecture](../entries/staircase-architecture.yml)                                                       | staircase                                | WF       | 0        | 保留现有辅助与关联                                   |
| [stake-engineering](../entries/stake-engineering.yml)                                                                 | stake                                    | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [stale-tech](../entries/stale-tech.yml)                                                                               | stale                                    | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [stall-workflow](../entries/stall-workflow.yml)                                                                       | stall                                    | —        | 1        | 保留现有辅助与关联                                   |
| [stamina-general](../entries/stamina-general.yml)                                                                     | stamina                                  | ETY / N  | 1        | 保留现有辅助与关联                                   |
| [star-method-interview](../entries/star-method-interview.yml)                                                         | STAR method                              | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [start-tag-html](../entries/start-tag-html.yml)                                                                       | start tag                                | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [state-preservation-react](../entries/state-preservation-react.yml)                                                   | state preservation                       | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [state-react](../entries/state-react.yml)                                                                             | state                                    | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [state-setter-react](../entries/state-setter-react.yml)                                                               | state setter                             | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [state-snapshot-react](../entries/state-snapshot-react.yml)                                                           | state snapshot                           | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [static-initialization-tech](../entries/static-initialization-tech.yml)                                               | static initialization                    | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [static-inspection-tech](../entries/static-inspection-tech.yml)                                                       | static inspection                        | —        | 1        | 保留现有辅助与关联                                   |
| [statue-architecture](../entries/statue-architecture.yml)                                                             | statue                                   | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [status-workflow](../entries/status-workflow.yml)                                                                     | status                                   | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [stealth-mask-zelda](../entries/stealth-mask-zelda.yml)                                                               | Stealth Mask                             | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [stealth-up-zelda](../entries/stealth-up-zelda.yml)                                                                   | Stealth Up                               | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [steam-cooking](../entries/steam-cooking.yml)                                                                         | steam                                    | N        | 0        | 保留现有辅助与关联                                   |
| [steer-general](../entries/steer-general.yml)                                                                         | steer                                    | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [steering-stick-zelda](../entries/steering-stick-zelda.yml)                                                           | Steering Stick                           | WF       | 0        | 保留现有辅助与关联                                   |
| [stew-cooking](../entries/stew-cooking.yml)                                                                           | stew                                     | —        | 1        | 保留现有辅助与关联                                   |
| [stick-general](../entries/stick-general.yml)                                                                         | stick                                    | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [sticky-positioning-css](../entries/sticky-positioning-css.yml)                                                       | sticky positioning                       | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [stock-finance](../entries/stock-finance.yml)                                                                         | stock                                    | —        | 1        | 保留现有辅助与关联                                   |
| [stone-general](../entries/stone-general.yml)                                                                         | stone                                    | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [stone-talus-zelda](../entries/stone-talus-zelda.yml)                                                                 | Stone Talus                              | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [storage-event-browser](../entries/storage-event-browser.yml)                                                         | storage event                            | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [streaming-tech](../entries/streaming-tech.yml)                                                                       | streaming                                | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [strict-mode-react](../entries/strict-mode-react.yml)                                                                 | Strict Mode                              | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [strike-price-finance](../entries/strike-price-finance.yml)                                                           | strike price                             | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [string-typescript](../entries/string-typescript.yml)                                                                 | string                                   | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [struct-go](../entries/struct-go.yml)                                                                                 | struct                                   | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [structural-check-tech](../entries/structural-check-tech.yml)                                                         | structural check                         | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [structural-typing-typescript](../entries/structural-typing-typescript.yml)                                           | structural typing                        | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [sturdy-general](../entries/sturdy-general.yml)                                                                       | sturdy                                   | N        | 1        | 保留现有辅助与关联                                   |
| [style-rule-css](../entries/style-rule-css.yml)                                                                       | style rule                               | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [stylesheet-css](../entries/stylesheet-css.yml)                                                                       | stylesheet                               | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [subagent-ai](../entries/subagent-ai.yml)                                                                             | subagent                                 | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [submission-workflow](../entries/submission-workflow.yml)                                                             | submission                               | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [subsequent-tech](../entries/subsequent-tech.yml)                                                                     | subsequent                               | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [subtest-go](../entries/subtest-go.yml)                                                                               | subtest                                  | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [sugar-food](../entries/sugar-food.yml)                                                                               | sugar                                    | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [summary-language](../entries/summary-language.yml)                                                                   | summary                                  | —        | 1        | 保留现有辅助与关联                                   |
| [summit-terrain](../entries/summit-terrain.yml)                                                                       | summit                                   | N        | 1        | 保留现有辅助与关联                                   |
| [summon-general](../entries/summon-general.yml)                                                                       | summon                                   | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [sundelion-zelda](../entries/sundelion-zelda.yml)                                                                     | Sundelion                                | —        | 0        | 暂缓，游戏名称的命名动机尚无可靠来源                 |
| [sunny-general](../entries/sunny-general.yml)                                                                         | sunny                                    | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [support-general](../entries/support-general.yml)                                                                     | support                                  | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [surface-terrain](../entries/surface-terrain.yml)                                                                     | surface                                  | —        | 1        | 保留现有辅助与关联                                   |
| [surrogate-pair-tech](../entries/surrogate-pair-tech.yml)                                                             | surrogate pair                           | —        | 1        | 保留现有辅助与关联                                   |
| [susceptible-combat](../entries/susceptible-combat.yml)                                                               | susceptible                              | N        | 1        | 保留现有辅助与关联                                   |
| [suspected-tech](../entries/suspected-tech.yml)                                                                       | suspected                                | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [suspend-hang-general](../entries/suspend-hang-general.yml)                                                           | suspend                                  | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [suspense-react](../entries/suspense-react.yml)                                                                       | Suspense                                 | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [swamp-terrain](../entries/swamp-terrain.yml)                                                                         | swamp                                    | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [sweet-and-sour-food](../entries/sweet-and-sour-food.yml)                                                             | sweet and sour                           | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [sword-weapon](../entries/sword-weapon.yml)                                                                           | sword                                    | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [symbol-javascript](../entries/symbol-javascript.yml)                                                                 | Symbol                                   | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [symbol-market](../entries/symbol-market.yml)                                                                         | symbol                                   | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [symlink-tech](../entries/symlink-tech.yml)                                                                           | symlink                                  | WF       | 0        | 保留现有辅助与关联                                   |
| [synchronize-tech](../entries/synchronize-tech.yml)                                                                   | synchronize                              | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [syntax-tech](../entries/syntax-tech.yml)                                                                             | syntax                                   | ETY      | 0        | 保留现有辅助与关联                                   |
| [synthesize-tech](../entries/synthesize-tech.yml)                                                                     | synthesize                               | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [synthetic-tech](../entries/synthetic-tech.yml)                                                                       | synthetic                                | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [system-prompt-ai](../entries/system-prompt-ai.yml)                                                                   | system prompt                            | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [system-ui](../entries/system-ui.yml)                                                                                 | system                                   | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [table-driven-test-go](../entries/table-driven-test-go.yml)                                                           | table-driven test                        | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [table-row-html](../entries/table-row-html.yml)                                                                       | table row                                | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [tag-html](../entries/tag-html.yml)                                                                                   | tag                                      | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [tarrey-town-zelda](../entries/tarrey-town-zelda.yml)                                                                 | Tarrey Town                              | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [tear-general](../entries/tear-general.yml)                                                                           | tear                                     | N        | 1        | 保留现有辅助与关联                                   |
| [teba-zelda](../entries/teba-zelda.yml)                                                                               | Teba                                     | N        | 0        | 保留现有辅助与关联                                   |
| [template-literal-type-typescript](../entries/template-literal-type-typescript.yml)                                   | template literal type                    | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [temple-architecture](../entries/temple-architecture.yml)                                                             | temple                                   | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [temple-of-time-zelda](../entries/temple-of-time-zelda.yml)                                                           | Temple of Time                           | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [temporal-dead-zone-javascript](../entries/temporal-dead-zone-javascript.yml)                                         | temporal dead zone                       | —        | 1        | 保留现有辅助与关联                                   |
| [teriyaki-food](../entries/teriyaki-food.yml)                                                                         | teriyaki                                 | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [test-suite-tech](../entries/test-suite-tech.yml)                                                                     | test suite                               | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [test-tech](../entries/test-tech.yml)                                                                                 | test                                     | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [text-to-speech-tech](../entries/text-to-speech-tech.yml)                                                             | text-to-speech                           | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [text-wrapping-css](../entries/text-wrapping-css.yml)                                                                 | text wrapping                            | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [the-depths-zelda](../entries/the-depths-zelda.yml)                                                                   | The Depths                               | —        | 1        | 保留现有辅助与关联                                   |
| [theming-bootstrap](../entries/theming-bootstrap.yml)                                                                 | theming                                  | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [throw-general](../entries/throw-general.yml)                                                                         | throw                                    | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [tights-clothing](../entries/tights-clothing.yml)                                                                     | tights                                   | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [time-bomb-zelda](../entries/time-bomb-zelda.yml)                                                                     | Time Bomb                                | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [timid-general](../entries/timid-general.yml)                                                                         | timid                                    | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [toast-bootstrap](../entries/toast-bootstrap.yml)                                                                     | toast                                    | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [toggle-tech](../entries/toggle-tech.yml)                                                                             | toggle                                   | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [toggle-trigger-bootstrap](../entries/toggle-trigger-bootstrap.yml)                                                   | toggle trigger                           | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [token-llm](../entries/token-llm.yml)                                                                                 | token                                    | —        | 1        | 保留现有辅助与关联                                   |
| [tomato-food](../entries/tomato-food.yml)                                                                             | tomato                                   | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [tool-call-ai](../entries/tool-call-ai.yml)                                                                           | tool call                                | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [tool-result-ai](../entries/tool-result-ai.yml)                                                                       | tool result                              | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [tool-tech](../entries/tool-tech.yml)                                                                                 | tool                                     | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [toolchain-tech](../entries/toolchain-tech.yml)                                                                       | toolchain                                | WF       | 0        | 保留现有辅助与关联                                   |
| [tooltip-bootstrap](../entries/tooltip-bootstrap.yml)                                                                 | tooltip                                  | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [torch-weapon](../entries/torch-weapon.yml)                                                                           | torch                                    | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [tough-general](../entries/tough-general.yml)                                                                         | tough                                    | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [tower-architecture](../entries/tower-architecture.yml)                                                               | tower                                    | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [trade-finance](../entries/trade-finance.yml)                                                                         | trade                                    | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [tradeoff-general](../entries/tradeoff-general.yml)                                                                   | tradeoff                                 | WF       | 0        | 保留现有辅助与关联                                   |
| [trail-story](../entries/trail-story.yml)                                                                             | trail                                    | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [transaction-finance](../entries/transaction-finance.yml)                                                             | transaction                              | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [transcript-tech](../entries/transcript-tech.yml)                                                                     | transcript                               | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [transfer-finance](../entries/transfer-finance.yml)                                                                   | transfer                                 | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [transform-css](../entries/transform-css.yml)                                                                         | transform                                | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [transformation-tech](../entries/transformation-tech.yml)                                                             | transformation                           | —        | 1        | 保留现有辅助与关联                                   |
| [transition-css](../entries/transition-css.yml)                                                                       | transition                               | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [transition-react](../entries/transition-react.yml)                                                                   | transition                               | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [transport-failure-tech](../entries/transport-failure-tech.yml)                                                       | transport failure                        | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [transport-metadata-tech](../entries/transport-metadata-tech.yml)                                                     | transport metadata                       | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [traveler-general](../entries/traveler-general.yml)                                                                   | traveler                                 | WF       | 0        | 保留现有辅助与关联                                   |
| [travelers-sword-zelda](../entries/travelers-sword-zelda.yml)                                                         | Traveler's Sword                         | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [traversal-tech](../entries/traversal-tech.yml)                                                                       | traversal                                | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [tree-data-structure](../entries/tree-data-structure.yml)                                                             | tree                                     | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [tree-shaking-build](../entries/tree-shaking-build.yml)                                                               | tree shaking                             | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [tree-terrain](../entries/tree-terrain.yml)                                                                           | tree                                     | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [triage-workflow](../entries/triage-workflow.yml)                                                                     | triage                                   | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [trident-weapon](../entries/trident-weapon.yml)                                                                       | trident                                  | ETY      | 0        | 保留现有辅助与关联                                   |
| [trigger-tech](../entries/trigger-tech.yml)                                                                           | trigger                                  | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [trousers-clothing](../entries/trousers-clothing.yml)                                                                 | trousers                                 | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [truffle-food](../entries/truffle-food.yml)                                                                           | truffle                                  | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [truncated-tech](../entries/truncated-tech.yml)                                                                       | truncated                                | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [try-files-nginx](../entries/try-files-nginx.yml)                                                                     | try_files                                | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [tts-language-tech](../entries/tts-language-tech.yml)                                                                 | TTS language                             | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [tulin-zelda](../entries/tulin-zelda.yml)                                                                             | Tulin                                    | N        | 0        | 保留现有辅助与关联                                   |
| [tunic-clothing](../entries/tunic-clothing.yml)                                                                       | tunic                                    | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [tuple-typescript](../entries/tuple-typescript.yml)                                                                   | tuple                                    | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [tv-mode-switch](../entries/tv-mode-switch.yml)                                                                       | TV Mode                                  | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [two-handed-combat](../entries/two-handed-combat.yml)                                                                 | two-handed                               | WF       | 0        | 保留现有辅助与关联                                   |
| [two-way-binding-frontend](../entries/two-way-binding-frontend.yml)                                                   | two-way binding                          | —        | 1        | 保留现有辅助与关联                                   |
| [type-alias-typescript](../entries/type-alias-typescript.yml)                                                         | type alias                               | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [type-annotation-typescript](../entries/type-annotation-typescript.yml)                                               | type annotation                          | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [type-argument-tech](../entries/type-argument-tech.yml)                                                               | type argument                            | —        | 1        | 保留现有辅助与关联                                   |
| [type-assertion-tech](../entries/type-assertion-tech.yml)                                                             | type assertion                           | —        | 1        | 保留现有辅助与关联                                   |
| [type-constraint-go](../entries/type-constraint-go.yml)                                                               | type constraint                          | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [type-guard-typescript](../entries/type-guard-typescript.yml)                                                         | type guard                               | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [type-inference-typescript](../entries/type-inference-typescript.yml)                                                 | type inference                           | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [type-instantiation-tech](../entries/type-instantiation-tech.yml)                                                     | type instantiation                       | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [type-only-import-typescript](../entries/type-only-import-typescript.yml)                                             | type-only import                         | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [type-parameter-typescript](../entries/type-parameter-typescript.yml)                                                 | type parameter                           | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [type-predicate-typescript](../entries/type-predicate-typescript.yml)                                                 | type predicate                           | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [type-selector-css](../entries/type-selector-css.yml)                                                                 | type selector                            | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [type-switch-go](../entries/type-switch-go.yml)                                                                       | type switch                              | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [type-tech](../entries/type-tech.yml)                                                                                 | type                                     | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [typeof-type-operator-typescript](../entries/typeof-type-operator-typescript.yml)                                     | typeof type operator                     | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [typo-general](../entries/typo-general.yml)                                                                           | typo                                     | WF       | 0        | 保留现有辅助与关联                                   |
| [ultrahand-zelda](../entries/ultrahand-zelda.yml)                                                                     | Ultrahand                                | WF       | 0        | 保留现有辅助与关联                                   |
| [unclaimed-workflow](../entries/unclaimed-workflow.yml)                                                               | unclaimed                                | WF       | 0        | 保留现有辅助与关联                                   |
| [uncontrolled-component-react](../entries/uncontrolled-component-react.yml)                                           | uncontrolled component                   | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [undead-combat](../entries/undead-combat.yml)                                                                         | undead                                   | WF       | 0        | 保留现有辅助与关联                                   |
| [undefined-typescript](../entries/undefined-typescript.yml)                                                           | undefined                                | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [underground-general](../entries/underground-general.yml)                                                             | underground                              | WF       | 0        | 保留现有辅助与关联                                   |
| [underline-tech](../entries/underline-tech.yml)                                                                       | underline                                | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [unfuse-general](../entries/unfuse-general.yml)                                                                       | unfuse                                   | WF       | 1        | 保留现有辅助与关联                                   |
| [uninitialized-tech](../entries/uninitialized-tech.yml)                                                               | uninitialized                            | WF       | 0        | 保留现有辅助与关联                                   |
| [union-tech](../entries/union-tech.yml)                                                                               | union                                    | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [unknown-typescript](../entries/unknown-typescript.yml)                                                               | unknown                                  | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [unlock-general](../entries/unlock-general.yml)                                                                       | unlock                                   | WF       | 0        | 保留现有辅助与关联                                   |
| [unmount-tech](../entries/unmount-tech.yml)                                                                           | unmount                                  | WF / N   | 1        | 保留现有辅助与关联                                   |
| [unordered-list-html](../entries/unordered-list-html.yml)                                                             | unordered list                           | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [unregister-tech](../entries/unregister-tech.yml)                                                                     | unregister                               | WF / N   | 1        | 保留现有辅助与关联                                   |
| [unstaged-git](../entries/unstaged-git.yml)                                                                           | unstaged                                 | WF / N   | 1        | 保留现有辅助与关联                                   |
| [untracked-git](../entries/untracked-git.yml)                                                                         | untracked                                | WF       | 0        | 保留现有辅助与关联                                   |
| [updater-function-react](../entries/updater-function-react.yml)                                                       | updater function                         | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [updraft-flight](../entries/updraft-flight.yml)                                                                       | updraft                                  | WF       | 0        | 保留现有辅助与关联                                   |
| [upgrade-general](../entries/upgrade-general.yml)                                                                     | upgrade                                  | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [upheaval-story](../entries/upheaval-story.yml)                                                                       | upheaval                                 | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [upstream-git](../entries/upstream.yml)                                                                               | upstream                                 | WF       | 0        | 保留现有辅助与关联                                   |
| [upstream-nginx](../entries/upstream-nginx.yml)                                                                       | upstream                                 | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [utility-api-bootstrap](../entries/utility-api-bootstrap.yml)                                                         | utility API                              | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [utility-class-bootstrap](../entries/utility-class-bootstrap.yml)                                                     | utility class                            | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [utility-type-typescript](../entries/utility-type-typescript.yml)                                                     | utility type                             | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [valid-state-bootstrap](../entries/valid-state-bootstrap.yml)                                                         | valid state                              | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [validation-feedback-bootstrap](../entries/validation-feedback-bootstrap.yml)                                         | validation feedback                      | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [validation-tech](../entries/validation-tech.yml)                                                                     | validation                               | WF       | 1        | 保留现有辅助与关联                                   |
| [valley-terrain](../entries/valley-terrain.yml)                                                                       | valley                                   | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [var-function-css](../entries/var-function-css.yml)                                                                   | var()                                    | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [variable-go](../entries/variable-go.yml)                                                                             | variable                                 | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [variance-type-system](../entries/variance-type-system.yml)                                                           | variance                                 | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [verification-tech](../entries/verification-tech.yml)                                                                 | verification                             | WF       | 1        | 保留现有辅助与关联                                   |
| [version-pin-tech](../entries/version-pin-tech.yml)                                                                   | version pin                              | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [view-general](../entries/view-general.yml)                                                                           | view                                     | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [viewport-css](../entries/viewport-css.yml)                                                                           | viewport                                 | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [village-architecture](../entries/village-architecture.yml)                                                           | village                                  | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [virtual-dom-frontend](../entries/virtual-dom-frontend.yml)                                                           | virtual DOM                              | N        | 0        | 保留现有辅助与关联                                   |
| [visually-hidden-bootstrap](../entries/visually-hidden-bootstrap.yml)                                                 | visually hidden                          | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [vitality-general](../entries/vitality-general.yml)                                                                   | vitality                                 | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [vocabulary-general](../entries/vocabulary.yml)                                                                       | vocabulary                               | ETY      | 0        | 保留现有辅助与关联                                   |
| [void-element-html](../entries/void-element-html.yml)                                                                 | void element                             | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [void-typescript](../entries/void-typescript.yml)                                                                     | void                                     | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [volatility-finance](../entries/volatility-finance.yml)                                                               | volatility                               | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [volcanic-combat](../entries/volcanic-combat.yml)                                                                     | volcanic                                 | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [volume-docker](../entries/volume-docker.yml)                                                                         | volume                                   | —        | 0        | R02 反向                                             |
| [volume-trading](../entries/volume-trading.yml)                                                                       | volume                                   | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [vulnerable-combat](../entries/vulnerable-combat.yml)                                                                 | vulnerable                               | ETY / N  | 1        | 保留现有辅助与关联                                   |
| [waistguard-clothing](../entries/waistguard-clothing.yml)                                                             | waistguard                               | WF       | 0        | 保留现有辅助与关联                                   |
| [wait-group-go](../entries/wait-group-go.yml)                                                                         | wait group                               | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [warp-general](../entries/warp-general.yml)                                                                           | warp                                     | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [warrior-general](../entries/warrior-general.yml)                                                                     | warrior                                  | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [watchlist-finance](../entries/watchlist-finance.yml)                                                                 | watchlist                                | WF       | 0        | 保留现有辅助与关联                                   |
| [water-general](../entries/water-general.yml)                                                                         | water                                    | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [water-temple-zelda](../entries/water-temple-zelda.yml)                                                               | Water Temple                             | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [waterfall-terrain](../entries/waterfall-terrain.yml)                                                                 | waterfall                                | WF       | 0        | 保留现有辅助与关联                                   |
| [weak-point-combat](../entries/weak-point-combat.yml)                                                                 | weak point                               | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [wealth-general](../entries/wealth-general.yml)                                                                       | wealth                                   | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [weapon-weapon](../entries/weapon-weapon.yml)                                                                         | weapon                                   | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [web-server-tech](../entries/web-server-tech.yml)                                                                     | web server                               | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [web-storage-browser](../entries/web-storage-browser.yml)                                                             | web storage                              | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [well-architecture](../entries/well-architecture.yml)                                                                 | well                                     | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [wetlands-terrain](../entries/wetlands-terrain.yml)                                                                   | wetlands                                 | WF       | 0        | 保留现有辅助与关联                                   |
| [wheat-food](../entries/wheat-food.yml)                                                                               | wheat                                    | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [wheel-engineering](../entries/wheel-engineering.yml)                                                                 | wheel                                    | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [when-applicable-workflow](../entries/when-applicable-workflow.yml)                                                   | when applicable                          | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [where-practical-workflow](../entries/where-practical-workflow.yml)                                                   | where practical                          | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [whiskas-pet-food](../entries/whiskas-pet-food.yml)                                                                   | Whiskas                                  | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [wield-story](../entries/wield-story.yml)                                                                             | wield                                    | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [wiggle-general](../entries/wiggle-general.yml)                                                                       | wiggle                                   | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [wind-temple-zelda](../entries/wind-temple-zelda.yml)                                                                 | Wind Temple                              | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [wind-terrain](../entries/wind-terrain.yml)                                                                           | wind                                     | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [wing-engineering](../entries/wing-engineering.yml)                                                                   | wing                                     | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [wise-general](../entries/wise-general.yml)                                                                           | wise                                     | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [withdraw-finance](../entries/withdraw-finance.yml)                                                                   | withdraw                                 | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [withdrawable-finance](../entries/withdrawable-finance.yml)                                                           | withdrawable                             | WF       | 1        | 保留现有辅助与关联                                   |
| [wizzrobe-zelda](../entries/wizzrobe-zelda.yml)                                                                       | Wizzrobe                                 | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [word-general](../entries/word-general.yml)                                                                           | word                                     | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [wording-general](../entries/wording-general.yml)                                                                     | wording                                  | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [work-general](../entries/work-general.yml)                                                                           | work                                     | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [workaround-tech](../entries/workaround-tech.yml)                                                                     | workaround                               | —        | 1        | 保留现有辅助与关联                                   |
| [worker-pool-concurrency](../entries/worker-pool-concurrency.yml)                                                     | worker pool                              | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [workload-tech](../entries/workload-tech.yml)                                                                         | workload                                 | WF       | 0        | 保留现有辅助与关联                                   |
| [would-you-like-a-or-b](../entries/would-you-like-a-or-b.yml)                                                         | Would you like A or B?                   | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [year-over-year-finance](../entries/year-over-year-finance.yml)                                                       | year-over-year                           | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [yiga-clan-hideout-zelda](../entries/yiga-clan-hideout-zelda.yml)                                                     | Yiga Clan Hideout                        | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [yiga-clan-zelda](../entries/yiga-clan-zelda.yml)                                                                     | Yiga Clan                                | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [yona-zelda](../entries/yona-zelda.yml)                                                                               | Yona                                     | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [yunobo-zelda](../entries/yunobo-zelda.yml)                                                                           | Yunobo                                   | N        | 0        | 保留现有辅助与关联                                   |
| [z-index-css](../entries/z-index-css.yml)                                                                             | z-index                                  | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [zero-value-go](../entries/zero-value-go.yml)                                                                         | zero value                               | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [zonai-device-zelda](../entries/zonai-device-zelda.yml)                                                               | Zonai Device                             | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [zonai-zelda](../entries/zonai-zelda.yml)                                                                             | Zonai                                    | N        | 0        | 保留现有辅助与关联                                   |
| [zonaite-zelda](../entries/zonaite-zelda.yml)                                                                         | Zonaite                                  | —        | 0        | 暂缓，游戏名称的命名动机尚无可靠来源                 |
| [zora-armor-zelda](../entries/zora-armor-zelda.yml)                                                                   | Zora Armor                               | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [zora-zelda](../entries/zora-zelda.yml)                                                                               | Zora                                     | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [zoras-domain-zelda](../entries/zoras-domain-zelda.yml)                                                               | Zora's Domain                            | —        | 0        | 本轮无提案，空字段保持省略或原值                     |

### 日语，243 项

| 稳定 ID                                                                                                 | Word                                         | 既有字段 | 既有邻居 | 决定                                                 |
|:--------------------------------------------------------------------------------------------------------|:---------------------------------------------|:---------|:---------|:-----------------------------------------------------|
| [aitemu](../entries/aitemu.yml)                                                                         | アイテム                                     | ETY      | 0        | 保留现有辅助与关联                                   |
| [amerika](../entries/amerika.yml)                                                                       | アメリカ                                     | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [amerika-kara-kimashita](../entries/amerika-kara-kimashita.yml)                                         | アメリカから来ました                         | WF       | 0        | 保留现有辅助与关联                                   |
| [ana](../entries/ana.yml)                                                                               | 穴                                           | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [anata](../entries/anata.yml)                                                                           | あなた                                       | N        | 2        | 保留现有辅助与关联                                   |
| [anime](../entries/anime.yml)                                                                           | アニメ                                       | WF       | 0        | 保留现有辅助与关联                                   |
| [anime-guzzu](../entries/anime-guzzu.yml)                                                               | アニメグッズ                                 | WF       | 0        | 保留现有辅助与关联                                   |
| [ano-hito](../entries/ano-hito.yml)                                                                     | あの人                                       | —        | 1        | 保留现有辅助与关联                                   |
| [ano-kata](../entries/ano-kata.yml)                                                                     | あの方                                       | —        | 1        | 保留现有辅助与关联                                   |
| [baagaa](../entries/baagaa.yml)                                                                         | バーガー                                     | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [bangou](../entries/bangou.yml)                                                                         | 番号                                         | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [bangou-o-omachigae-phone](../entries/bangou-o-omachigae-phone.yml)                                     | 番号をお間違えではないでしょうか。           | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [boku](../entries/boku.yml)                                                                             | 僕                                           | —        | 2        | 保留现有辅助与关联                                   |
| [boku-wa-shinsekai-no-kami-ni-naru](../entries/boku-wa-shinsekai-no-kami-ni-naru.yml)                   | 僕は新世界の神になる                         | WF       | 0        | 保留现有辅助与关联                                   |
| [burajiru](../entries/burajiru.yml)                                                                     | ブラジル                                     | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [burajiru-eaa](../entries/burajiru-eaa.yml)                                                             | ブラジルエアー                               | WF       | 0        | 保留现有辅助与关联                                   |
| [byouin](../entries/byouin.yml)                                                                         | 病院                                         | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [chan-suffix](../entries/chan-suffix.yml)                                                               | ちゃん                                       | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [chikara](../entries/chikara.yml)                                                                       | 力                                           | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [chikyuu](../entries/chikyuu.yml)                                                                       | 地球                                         | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [chisou](../entries/chisou.yml)                                                                         | 馳走                                         | ETY      | 0        | 保留现有辅助与关联                                   |
| [chuugoku](../entries/chuugoku.yml)                                                                     | 中国                                         | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [chuugokujin](../entries/chuugokujin.yml)                                                               | 中国人                                       | WF       | 0        | 保留现有辅助与关联                                   |
| [daigaku](../entries/daigaku.yml)                                                                       | 大学                                         | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [daigakusei](../entries/daigakusei.yml)                                                                 | 大学生                                       | —        | 1        | R06、WF03                                            |
| [daijoubu](../entries/daijoubu.yml)                                                                     | 大丈夫                                       | N        | 0        | 保留现有辅助与关联                                   |
| [dake-grammar](../entries/dake-grammar.yml)                                                             | ～だけ                                       | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [dare](../entries/dare.yml)                                                                             | 誰                                           | —        | 1        | 保留现有辅助与关联                                   |
| [denki](../entries/denki.yml)                                                                           | 電気                                         | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [denwa](../entries/denwa.yml)                                                                           | 電話                                         | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [denwa-bangou](../entries/denwa-bangou.yml)                                                             | 電話番号                                     | WF       | 0        | 保留现有辅助与关联                                   |
| [denwa-ga-aru](../entries/denwa-ga-aru.yml)                                                             | 電話がある                                   | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [deshou-ka-grammar](../entries/deshou-ka-grammar.yml)                                                   | ～でしょうか                                 | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [dewa-arimasen-grammar](../entries/dewa-arimasen-grammar.yml)                                           | ～ではありません                             | —        | 1        | 保留现有辅助与关联                                   |
| [dewa-nai-deshou-ka-grammar](../entries/dewa-nai-deshou-ka-grammar.yml)                                 | ～ではないでしょうか                         | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [doitsu](../entries/doitsu.yml)                                                                         | ドイツ                                       | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [donata](../entries/donata.yml)                                                                         | どなた                                       | —        | 1        | R07                                                  |
| [donata-o-osagashi-phone](../entries/donata-o-osagashi-phone.yml)                                       | どなたをお探しですか？                       | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [eaa](../entries/eaa.yml)                                                                               | エアー                                       | ETY      | 0        | 保留现有辅助与关联                                   |
| [edogawa-konan](../entries/edogawa-konan.yml)                                                           | 江戸川コナン                                 | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [ee-kee-shii](../entries/ee-kee-shii.yml)                                                               | エー・ケー・シー                             | —        | 0        | 暂缓，组织名称的缩写展开尚未确认                     |
| [eigo](../entries/eigo.yml)                                                                             | 英語                                         | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [eigo-de-daijoubu-phone](../entries/eigo-de-daijoubu-phone.yml)                                         | 英語で大丈夫ですか？                         | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [fisshu](../entries/fisshu.yml)                                                                         | フィッシュ                                   | ETY      | 1        | 保留现有辅助与关联                                   |
| [fisshu-baagaa](../entries/fisshu-baagaa.yml)                                                           | フィッシュバーガー                           | WF       | 0        | 保留现有辅助与关联                                   |
| [fuji-daigaku](../entries/fuji-daigaku.yml)                                                             | 富士大学                                     | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [fuminijiru](../entries/fuminijiru.yml)                                                                 | 踏みにじる                                   | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [futatsu](../entries/futatsu.yml)                                                                       | 二つ                                         | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [fuuin](../entries/fuuin.yml)                                                                           | 封印                                         | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [fuuin-kaijo-reriizu](../entries/fuuin-kaijo-reriizu.yml)                                               | 封印解除（レリーズ）！                       | WF / N   | 0        | 保留现有辅助与关联                                   |
| [gaikoku](../entries/gaikoku.yml)                                                                       | 外国                                         | —        | 1        | 保留现有辅助与关联                                   |
| [gaikokujin](../entries/gaikokujin.yml)                                                                 | 外国人                                       | WF / N   | 2        | 保留现有辅助与关联                                   |
| [gakusei](../entries/gakusei.yml)                                                                       | 学生                                         | —        | 0        | R06 反向                                             |
| [gentei](../entries/gentei.yml)                                                                         | 限定                                         | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [gentei-guzzu](../entries/gentei-guzzu.yml)                                                             | 限定グッズ                                   | WF       | 0        | 保留现有辅助与关联                                   |
| [gimu](../entries/gimu.yml)                                                                             | 義務                                         | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [gimu-toka-janai-yo](../entries/gimu-toka-janai-yo.yml)                                                 | 義務とかじゃないよ                           | WF       | 0        | 保留现有辅助与关联                                   |
| [gochisou](../entries/gochisou.yml)                                                                     | ごちそう                                     | ETY      | 0        | 保留现有辅助与关联                                   |
| [gochisousama](../entries/gochisousama.yml)                                                             | ごちそうさま                                 | —        | 2        | 保留现有辅助与关联                                   |
| [gochisousama-deshita](../entries/gochisousama-deshita.yml)                                             | ごちそうさまでした                           | —        | 1        | 保留现有辅助与关联                                   |
| [gomen-ne](../entries/gomen-ne.yml)                                                                     | ごめんね                                     | —        | 1        | 保留现有辅助与关联                                   |
| [gomennasai](../entries/gomennasai.yml)                                                                 | ごめんなさい                                 | —        | 2        | 保留现有辅助与关联                                   |
| [guzzu](../entries/guzzu.yml)                                                                           | グッズ                                       | ETY      | 0        | 保留现有辅助与关联                                   |
| [hai](../entries/hai.yml)                                                                               | はい                                         | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [hajimemashite](../entries/hajimemashite.yml)                                                           | はじめまして                                 | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [hanasemasen](../entries/hanasemasen.yml)                                                               | 話せません                                   | WF / N   | 1        | 保留现有辅助与关联                                   |
| [hanasu](../entries/hanasu.yml)                                                                         | 話す                                         | N        | 1        | 保留现有辅助与关联                                   |
| [hassai](../entries/hassai.yml)                                                                         | 八歳                                         | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [hikari](../entries/hikari.yml)                                                                         | 光                                           | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [hikari-ni-nareru-nda](../entries/hikari-ni-nareru-nda.yml)                                             | 光になれるんだ                               | WF       | 0        | 保留现有辅助与关联                                   |
| [hito](../entries/hito.yml)                                                                             | 人                                           | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [hito-o-norowaba-ana-futatsu](../entries/hito-o-norowaba-ana-futatsu.yml)                               | 人を呪わば穴二つ                             | N        | 0        | 保留现有辅助与关联                                   |
| [hitotsu](../entries/hitotsu.yml)                                                                       | 一つ                                         | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [hokage-naruto](../entries/hokage-naruto.yml)                                                           | 火影                                         | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [hokage-ni-naru-made](../entries/hokage-ni-naru-made.yml)                                               | 火影になるまで                               | WF       | 0        | R08 反向                                             |
| [honjitsu](../entries/honjitsu.yml)                                                                     | 本日                                         | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [hontou-ni](../entries/hontou-ni.yml)                                                                   | 本当に                                       | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [hontou-ni-gomennasai](../entries/hontou-ni-gomennasai.yml)                                             | 本当にごめんなさい                           | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [igirisu](../entries/igirisu.yml)                                                                       | イギリス                                     | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [iie](../entries/iie.yml)                                                                               | いいえ                                       | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [indoneshia](../entries/indoneshia.yml)                                                                 | インドネシア                                 | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [ippen](../entries/ippen.yml)                                                                           | いっぺん                                     | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [ippen-shinde-miru](../entries/ippen-shinde-miru.yml)                                                   | いっぺん、死んでみる？                       | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [iryou-houjin-shadan](../entries/iryou-houjin-shadan.yml)                                               | 医療法人社団                                 | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [isha](../entries/isha.yml)                                                                             | 医者                                         | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [issai](../entries/issai.yml)                                                                           | 一歳                                         | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [itadakimasu](../entries/itadakimasu.yml)                                                               | いただきます                                 | —        | 1        | 保留现有辅助与关联                                   |
| [itsumo](../entries/itsumo.yml)                                                                         | いつも                                       | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [janai-grammar](../entries/janai-grammar.yml)                                                           | ～じゃない                                   | —        | 1        | 保留现有辅助与关联                                   |
| [jibun](../entries/jibun.yml)                                                                           | 自分                                         | —        | 1        | 保留现有辅助与关联                                   |
| [jibun-jishin](../entries/jibun-jishin.yml)                                                             | 自分自身                                     | —        | 1        | 保留现有辅助与关联                                   |
| [jibun-jishin-no-chikara-de](../entries/jibun-jishin-no-chikara-de.yml)                                 | 自分自身の力で                               | WF       | 0        | 保留现有辅助与关联                                   |
| [jin-nationality-suffix](../entries/jin-nationality-suffix.yml)                                         | ～人                                         | N        | 0        | 保留现有辅助与关联                                   |
| [jiyuu](../entries/jiyuu.yml)                                                                           | 自由                                         | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [jumoku](../entries/jumoku.yml)                                                                         | 樹木                                         | —        | 1        | 保留现有辅助与关联                                   |
| [kai-counter](../entries/kai-counter.yml)                                                               | 回                                           | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [kaijo](../entries/kaijo.yml)                                                                           | 解除                                         | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [kaisha](../entries/kaisha.yml)                                                                         | 会社                                         | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [kaishain](../entries/kaishain.yml)                                                                     | 会社員                                       | WF / N   | 1        | 保留现有辅助与关联                                   |
| [kami](../entries/kami.yml)                                                                             | 神                                           | —        | 1        | 保留现有辅助与关联                                   |
| [kangoshi](../entries/kangoshi.yml)                                                                     | 看護師                                       | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [kankoku](../entries/kankoku.yml)                                                                       | 韓国                                         | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [kara-kimashita-origin](../entries/kara-kimashita-origin.yml)                                           | ～から来ました                               | —        | 1        | 保留现有辅助与关联                                   |
| [kara-reason-grammar](../entries/kara-reason-grammar.yml)                                               | ～から                                       | —        | 1        | 保留现有辅助与关联                                   |
| [kata-person](../entries/kata-person.yml)                                                               | 方                                           | —        | 0        | R07 反向                                             |
| [kenkyuu](../entries/kenkyuu.yml)                                                                       | 研究                                         | —        | 1        | 保留现有辅助与关联                                   |
| [kenkyuusha](../entries/kenkyuusha.yml)                                                                 | 研究者                                       | WF       | 1        | 保留现有辅助与关联                                   |
| [ki](../entries/ki.yml)                                                                                 | 木                                           | —        | 1        | 保留现有辅助与关联                                   |
| [kimi](../entries/kimi.yml)                                                                             | 君                                           | —        | 1        | 保留现有辅助与关联                                   |
| [kimi-ga-yobu-namae-yume-no-tsuzuki](../entries/kimi-ga-yobu-namae-yume-no-tsuzuki.yml)                 | きみが呼ぶなまえ～夢のつづき                 | WF       | 0        | 保留现有辅助与关联                                   |
| [kochira-wa-san-desu](../entries/kochira-wa-san-desu.yml)                                               | こちらは～さんです                           | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [koneko](../entries/koneko.yml)                                                                         | 子猫                                         | WF       | 2        | 保留现有辅助与关联                                   |
| [kono-bangou](../entries/kono-bangou.yml)                                                               | この番号                                     | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [kono-bangou-wa-sochira-no-byouin-phone](../entries/kono-bangou-wa-sochira-no-byouin-phone.yml)         | この番号はそちらの病院の電話番号でしょうか。 | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [kosumo-o-moyase](../entries/kosumo-o-moyase.yml)                                                       | 小宇宙を燃やせ！                             | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [kosumo-saint-seiya](../entries/kosumo-saint-seiya.yml)                                                 | 小宇宙                                       | N        | 0        | 保留现有辅助与关联                                   |
| [kotoba](../entries/kotoba.yml)                                                                         | 言葉                                         | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [kotoba-wa-magenee](../entries/kotoba-wa-magenee.yml)                                                   | 言葉は曲げねェ                               | WF       | 0        | 保留现有辅助与关联                                   |
| [koube](../entries/koube.yml)                                                                           | 神戸                                         | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [koube-byouin](../entries/koube-byouin.yml)                                                             | 神戸病院                                     | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [koukishin](../entries/koukishin.yml)                                                                   | 好奇心                                       | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [koukishin-ousei-na-koneko](../entries/koukishin-ousei-na-koneko.yml)                                   | 好奇心旺盛な子猫                             | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [kyoushi](../entries/kyoushi.yml)                                                                       | 教師                                         | N        | 1        | 保留现有辅助与关联                                   |
| [kyouto-fu](../entries/kyouto-fu.yml)                                                                   | 京都府                                       | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [kyouto-shi](../entries/kyouto-shi.yml)                                                                 | 京都市                                       | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [kyuu-number](../entries/kyuu-number.yml)                                                               | 九                                           | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [machigaeru](../entries/machigaeru.yml)                                                                 | 間違える                                     | —        | 1        | 保留现有辅助与关联                                   |
| [machigai](../entries/machigai.yml)                                                                     | 間違い                                       | —        | 1        | N04                                                  |
| [made-grammar](../entries/made-grammar.yml)                                                             | ～まで                                       | —        | 0        | R08、WF04                                            |
| [mageru](../entries/mageru.yml)                                                                         | 曲げる                                       | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [massugu](../entries/massugu.yml)                                                                       | まっすぐ                                     | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [minna](../entries/minna.yml)                                                                           | みんな                                       | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [moeagare-ore-no-kosumo](../entries/moeagare-ore-no-kosumo.yml)                                         | 燃え上がれ、俺の小宇宙！                     | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [moeagaru](../entries/moeagaru.yml)                                                                     | 燃え上がる                                   | —        | 1        | 保留现有辅助与关联                                   |
| [morimoto](../entries/morimoto.yml)                                                                     | 森本                                         | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [mottomo-antei-ui](../entries/mottomo-antei-ui.yml)                                                     | 最も安定                                     | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [moushimasu](../entries/moushimasu.yml)                                                                 | 申します                                     | —        | 1        | 保留现有辅助与关联                                   |
| [moushiwake-arimasen](../entries/moushiwake-arimasen.yml)                                               | 申し訳ありません                             | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [moyasu](../entries/moyasu.yml)                                                                         | 燃やす                                       | —        | 1        | 保留现有辅助与关联                                   |
| [namae](../entries/namae.yml)                                                                           | 名前                                         | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [namida](../entries/namida.yml)                                                                         | 涙                                           | —        | 1        | 保留现有辅助与关联                                   |
| [nani-ga-da-rhetorical](../entries/nani-ga-da-rhetorical.yml)                                           | 何が～だ                                     | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [nani-ga-seigi-da](../entries/nani-ga-seigi-da.yml)                                                     | 何が正義だ！                                 | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [nanoru](../entries/nanoru.yml)                                                                         | 名乗る                                       | —        | 1        | 保留现有辅助与关联                                   |
| [nansai](../entries/nansai.yml)                                                                         | 何歳                                         | —        | 1        | 保留现有辅助与关联                                   |
| [nara](../entries/nara.yml)                                                                             | 奈良                                         | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [naru](../entries/naru.yml)                                                                             | なる                                         | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [nda-grammar](../entries/nda-grammar.yml)                                                               | ～んだ                                       | —        | 0        | 暂缓，尝试的语法页面无法打开，口语缩约的适用条件待核 |
| [neko](../entries/neko.yml)                                                                             | 猫                                           | —        | 1        | 保留现有辅助与关联                                   |
| [neko-chan](../entries/neko-chan.yml)                                                                   | 猫ちゃん                                     | WF       | 1        | 保留现有辅助与关联                                   |
| [ni-kawatte-grammar](../entries/ni-kawatte-grammar.yml)                                                 | ～にかわって                                 | —        | 1        | 保留现有辅助与关联                                   |
| [ni-nareru-grammar](../entries/ni-nareru-grammar.yml)                                                   | ～になれる                                   | WF       | 1        | 保留现有辅助与关联                                   |
| [ni-naru-grammar](../entries/ni-naru-grammar.yml)                                                       | 〜になる                                     | —        | 2        | 保留现有辅助与关联                                   |
| [nihon](../entries/nihon.yml)                                                                           | 日本                                         | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [nihongo](../entries/nihongo.yml)                                                                       | 日本語                                       | WF       | 1        | 保留现有辅助与关联                                   |
| [nihongo-ga-hanasemasen](../entries/nihongo-ga-hanasemasen.yml)                                         | 日本語が話せません。                         | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [nihonjin](../entries/nihonjin.yml)                                                                     | 日本人                                       | WF       | 1        | 保留现有辅助与关联                                   |
| [nihonjin-dewa-arimasen](../entries/nihonjin-dewa-arimasen.yml)                                         | 日本人ではありません                         | WF       | 0        | 保留现有辅助与关联                                   |
| [nindou-naruto](../entries/nindou-naruto.yml)                                                           | 忍道                                         | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [ningen](../entries/ningen.yml)                                                                         | 人間                                         | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [no-ka-grammar](../entries/no-ka-grammar.yml)                                                           | ～のか                                       | N        | 0        | 保留现有辅助与关联                                   |
| [norou](../entries/norou.yml)                                                                           | 呪う                                         | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [ohayou-gozaimasu](../entries/ohayou-gozaimasu.yml)                                                     | おはようございます                           | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [oikutsu](../entries/oikutsu.yml)                                                                       | おいくつ                                     | —        | 1        | 保留现有辅助与关联                                   |
| [okane](../entries/okane.yml)                                                                           | お金                                         | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [okurete-gomen-ne](../entries/okurete-gomen-ne.yml)                                                     | 遅れてごめんね                               | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [omachigae](../entries/omachigae.yml)                                                                   | お間違え                                     | WF       | 0        | 保留现有辅助与关联                                   |
| [omae](../entries/omae.yml)                                                                             | お前                                         | —        | 1        | 保留现有辅助与关联                                   |
| [omae-no-namida-de-chikyuu-ga-sukueru-no-ka](../entries/omae-no-namida-de-chikyuu-ga-sukueru-no-ka.yml) | お前の涙でこの地球が救えるのか？             | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [omotenashi](../entries/omotenashi.yml)                                                                 | おもてなし                                   | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [onamae-wa](../entries/onamae-wa.yml)                                                                   | お名前は                                     | N        | 0        | 保留现有辅助与关联                                   |
| [ore](../entries/ore.yml)                                                                               | 俺                                           | —        | 1        | 保留现有辅助与关联                                   |
| [ore-wa-ningen-dakara](../entries/ore-wa-ningen-dakara.yml)                                             | 俺は人間だから                               | WF       | 0        | 保留现有辅助与关联                                   |
| [osagashi](../entries/osagashi.yml)                                                                     | お探し                                       | WF       | 1        | 保留现有辅助与关联                                   |
| [oshioki](../entries/oshioki.yml)                                                                       | お仕置き                                     | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [osusume](../entries/osusume.yml)                                                                       | おすすめ                                     | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [ousei](../entries/ousei.yml)                                                                           | 旺盛                                         | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [pawaa](../entries/pawaa.yml)                                                                           | パワー                                       | ETY      | 0        | 保留现有辅助与关联                                   |
| [pawaa-denki](../entries/pawaa-denki.yml)                                                               | パワー電気                                   | WF       | 0        | 保留现有辅助与关联                                   |
| [rakuwakai-otowa-byouin](../entries/rakuwakai-otowa-byouin.yml)                                         | 洛和会音羽病院                               | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [rakuwakai-otowa-byouin-kangoshi-morimoto](../entries/rakuwakai-otowa-byouin-kangoshi-morimoto.yml)     | 洛和会音羽病院の看護師、森本と申します。     | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [sa-ending-grammar](../entries/sa-ending-grammar.yml)                                                   | ～さ                                         | N        | 0        | 保留现有辅助与关联                                   |
| [sagasu](../entries/sagasu.yml)                                                                         | 探す                                         | —        | 1        | 保留现有辅助与关联                                   |
| [saishou-katto-ui](../entries/saishou-katto-ui.yml)                                                     | 最小カット                                   | —        | 0        | 暂缓，缺少原始界面与カット在此处的确切机制           |
| [sakana-general](../entries/sakana-general.yml)                                                         | 魚                                           | —        | 1        | 保留现有辅助与关联                                   |
| [sakura](../entries/sakura.yml)                                                                         | 桜                                           | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [sakura-daigaku](../entries/sakura-daigaku.yml)                                                         | さくら大学                                   | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [sakura-no-ki](../entries/sakura-no-ki.yml)                                                             | 桜の木                                       | WF       | 0        | 保留现有辅助与关联                                   |
| [san-name-suffix](../entries/san-name-suffix.yml)                                                       | ～さん                                       | N        | 0        | 保留现有辅助与关联                                   |
| [seeraa-muun](../entries/seeraa-muun.yml)                                                               | セーラームーン                               | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [sei-kanji-reading](../entries/sei-kanji-reading.yml)                                                   | 成                                           | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [sei-surname](../entries/sei-surname.yml)                                                               | 姓                                           | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [seigi](../entries/seigi.yml)                                                                           | 正義                                         | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [sen-kanji-reading](../entries/sen-kanji-reading.yml)                                                   | 銭                                           | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [sensei](../entries/sensei.yml)                                                                         | 先生                                         | N        | 1        | 保留现有辅助与关联                                   |
| [shain](../entries/shain.yml)                                                                           | 社員                                         | N        | 1        | 保留现有辅助与关联                                   |
| [shinde-miru](../entries/shinde-miru.yml)                                                               | 死んでみる                                   | WF       | 0        | 保留现有辅助与关联                                   |
| [shinjitsu](../entries/shinjitsu.yml)                                                                   | 真実                                         | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [shinjitsu-wa-itsumo-hitotsu](../entries/shinjitsu-wa-itsumo-hitotsu.yml)                               | 真実はいつもひとつ！                         | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [shinsekai](../entries/shinsekai.yml)                                                                   | 新世界                                       | WF       | 0        | 保留现有辅助与关联                                   |
| [shinu](../entries/shinu.yml)                                                                           | 死ぬ                                         | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [shitsurei-desu-ga](../entries/shitsurei-desu-ga.yml)                                                   | 失礼ですが                                   | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [shokudou](../entries/shokudou.yml)                                                                     | 食堂                                         | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [shouhin](../entries/shouhin.yml)                                                                       | 商品                                         | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [sochira](../entries/sochira.yml)                                                                       | そちら                                       | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [sochira-no-byouin](../entries/sochira-no-byouin.yml)                                                   | そちらの病院                                 | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [sono-kata](../entries/sono-kata.yml)                                                                   | その方                                       | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [sono-namida-wa-nan-da](../entries/sono-namida-wa-nan-da.yml)                                           | その涙は何だ？                               | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [sore](../entries/sore.yml)                                                                             | それ                                         | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [sore-ga-ore-no-nindou](../entries/sore-ga-ore-no-nindou.yml)                                           | それがオレの忍道だ!!                         | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [sukueru](../entries/sukueru.yml)                                                                       | 救える                                       | WF       | 1        | 保留现有辅助与关联                                   |
| [sukuu](../entries/sukuu.yml)                                                                           | 救う                                         | —        | 1        | 保留现有辅助与关联                                   |
| [sumimasen](../entries/sumimasen.yml)                                                                   | すみません                                   | —        | 1        | 保留现有辅助与关联                                   |
| [tai-country](../entries/tai-country.yml)                                                               | タイ                                         | —        | 0        | 暂缓，外来地名的命名链待专门来源核实                 |
| [tai-grammar](../entries/tai-grammar.yml)                                                               | ～たい                                       | —        | 1        | 保留现有辅助与关联                                   |
| [tantei](../entries/tantei.yml)                                                                         | 探偵                                         | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [tantei-sa](../entries/tantei-sa.yml)                                                                   | 探偵さ！                                     | WF       | 0        | 保留现有辅助与关联                                   |
| [taoru](../entries/taoru.yml)                                                                           | タオル                                       | ETY      | 0        | 保留现有辅助与关联                                   |
| [te-miru-grammar](../entries/te-miru-grammar.yml)                                                       | 〜てみる                                     | WF / N   | 1        | 保留现有辅助与关联                                   |
| [tei-kanji-reading](../entries/tei-kanji-reading.yml)                                                   | 程                                           | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [to-moushimasu-grammar](../entries/to-moushimasu-grammar.yml)                                           | ～と申します                                 | WF       | 0        | 保留现有辅助与关联                                   |
| [to-nanoru-grammar](../entries/to-nanoru-grammar.yml)                                                   | ～と名乗る                                   | WF       | 0        | 保留现有辅助与关联                                   |
| [toka-grammar](../entries/toka-grammar.yml)                                                             | ～とか                                       | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [tsuki](../entries/tsuki.yml)                                                                           | 月                                           | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [tsuki-ni-kawatte-oshioki-yo](../entries/tsuki-ni-kawatte-oshioki-yo.yml)                               | 月にかわっておしおきよ！                     | WF       | 0        | 保留现有辅助与关联                                   |
| [tsukino-usagi](../entries/tsukino-usagi.yml)                                                           | 月野うさぎ                                   | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [tsuzuki](../entries/tsuzuki.yml)                                                                       | 続き                                         | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [watashi](../entries/watashi.yml)                                                                       | 私                                           | —        | 1        | 保留现有辅助与关联                                   |
| [watashi-wa-gaikokujin-desu](../entries/watashi-wa-gaikokujin-desu.yml)                                 | 私は外国人です                               | WF       | 0        | 保留现有辅助与关联                                   |
| [watashi-wa-sono-kata-dewa-arimasen](../entries/watashi-wa-sono-kata-dewa-arimasen.yml)                 | 私はその方ではありません。                   | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [wazawaza](../entries/wazawaza.yml)                                                                     | わざわざ                                     | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [yamashina-ku](../entries/yamashina-ku.yml)                                                             | 山科区                                       | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [yareru](../entries/yareru.yml)                                                                         | やれる                                       | WF       | 1        | 保留现有辅助与关联                                   |
| [yareru-koto](../entries/yareru-koto.yml)                                                               | やれること                                   | WF       | 0        | 保留现有辅助与关联                                   |
| [yaritai-dake-da-yo](../entries/yaritai-dake-da-yo.yml)                                                 | やりたいだけだよ                             | WF       | 0        | 保留现有辅助与关联                                   |
| [yaru](../entries/yaru.yml)                                                                             | やる                                         | —        | 1        | 保留现有辅助与关联                                   |
| [yobu](../entries/yobu.yml)                                                                             | 呼ぶ                                         | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [yomikata](../entries/yomikata.yml)                                                                     | 読み方                                       | WF       | 0        | 保留现有辅助与关联                                   |
| [yoroshiku-onegai-shimasu](../entries/yoroshiku-onegai-shimasu.yml)                                     | よろしくお願いします                         | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [yume](../entries/yume.yml)                                                                             | 夢                                           | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [zettai](../entries/zettai.yml)                                                                         | 絶対                                         | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [zettai-daijoubu-da-yo](../entries/zettai-daijoubu-da-yo.yml)                                           | 絶対、だいじょうぶだよ。                     | —        | 0        | 本轮无提案，空字段保持省略或原值                     |
| [zettee-shinanee](../entries/zettee-shinanee.yml)                                                       | ぜってェ死なねー                             | WF       | 0        | 保留现有辅助与关联                                   |

## 待批及实施程序

- 文本待批为 WF01–WF05、ETY01–ETY03 与 N01–N04。批准只涵盖所列原文、来源行及明确的 WF01 移出。
- 关联待批为 R01–R10。只在明确批准的左端追加右端 ID。

1. 重新读取获批计划及当前仓库并检查 Git 状态。协调规划后的工作，快照原词条、原导出和已有媒体哈希。
2. 只按编号写入批准的字段、来源行、精确移出及关系声明。保留稳定 ID、Word、Meaning、发音、例句、标签、学习数据、旧辅助、未授权字段与录音字节。此批只涉及现有词条的内容字段。现有词条发音无需重新生成，既有 schema、工具、模板和依赖保持不变。
3. 运行以下命令，另检查未跟踪文本。通过现有构建器生成 dist，不手工写导出。

```bash
go test ./...
go vet ./...
go run scripts/validate/main.go
go run scripts/build/main.go
git diff --check
```

4. 源文件保持一致后再构建一次并比较字节。核对实施前 ID 集合和行数、18 列和顺序、原 Audio 和 StudyOrder、批准辅助的准确值、原有与新增双向直接关系、空辅助字段及既有条件模板。当前导出列如下。

```plain
ID Language Word Pronunciation PartOfSpeech Meaning Context Example Notes Audio TTSLanguage StudyOrder Etymology WordFormation Mnemonic RelatedWords CodeExample Tags
```

5. 比较字段白名单与已有媒体哈希，保留无关工作并报告实际失败。用实际数据更新本计划的批准、实施和验证状态。静态与导出校验、人工发音听审及真实 Anki 导入、渲染、播放分别报告。
6. 授权范围止于本地词条变更与现有验证流程。暂存、提交、推送、复制媒体与操作 Anki 均不在范围内。

## 规划验证

已核对 1874 处本地链接、全部台账 ID、计数、提案字段和不重复的关系声明。正文无片段链接。核对了表格、文本格式、词典出处对应关系及证据限制。3761 个原有文件的 SHA-256 全部匹配，范围涵盖词条、录音、实现、模板及原导出。本轮未运行构建或生成音频，未执行 Anki 实机验证。

## 第二轮扩充审阅

原草案及其 1,833 项台账保留。下文从先前无提案的词、已记录的同形义项和日语例句重新筛选。优先采用能串起一组常用动作、能帮助辨义，或能迁移到同类句子的补充。各组独立审阅。词源另列，并为还需要历史证据的条目保留待核标记。

第二轮只新增审阅编号。台账补充表覆盖本轮涉及的 ID，其余条目沿用上文决定。文案采用当前稳定 ID、释义、语境和例句。词形和句型分析不推定历史来源，联想画面只用于学习。

| 范围                     | 本轮提案或计数                                                         |
|:--------------------------|:------------------------------------------------------------------------|
| 常用词族                 | 14 项文本，WF 14                                                       |
| 易混义项                 | 14 项文本，N 14                                                        |
| 日语词形与句型           | 14 项文本，N 6、WF 8                                                   |
| 词源单独审阅             | 2 项文本，ETY 2                                                        |
| 第二轮全部文本           | 44 个辅助字段，覆盖 42 个 ID，附 4 条拟追加来源行                      |
| 第二轮独立关系           | 16 对、30 个显示端点、14 个声明端，增加 16 条反向展示及 16 条正向展示  |
| 第二轮拟改文件           | 42 个 YAML，62 个字段值。关系声明端与文本目标重叠 14 个                |
| 先前无提案且无辅助或关联 | 从上文 1,332 个 ID 中新增 44 个文本或关联参与者，其中 34 个有文本文案  |
| 两轮合计                 | 56 个文本字段、11 个附来源行的 Notes 字段、26 对关系，涉及 24 个声明端 |
| 两轮拟改 YAML            | 59 个 YAML，91 个字段值                                                |

审核时可查看各组的编号范围。原 WF01–WF05、ETY01–ETY03、N01–N04 和 R01–R10 仍待批。批准新组不会自动批准旧组。

### 常用词族

#### WF06 detach

[detach-general](../entries/detach-general.yml) · detach · 分离、拆卸

新增 `word_formation`。让三个同场景动词组成动作链。 现代用法联记，词源另查。

```plain
三步记动作：attach 是装上，detach 是把已连接的部件拆开，reattach 是再装回去。
```

对照现有释义、语境或例句。[attach-general](../entries/attach-general.yml) · [reattach-general](../entries/reattach-general.yml)

Notes 原样保留。

#### WF07 equipment

[equipment-equipment](../entries/equipment-equipment.yml) · equipment · 装备

新增 `word_formation`。用句子位置区分动作和器材。 当前是通用装备义，不限定某个游戏的装备规则。

```plain
equip 是动词，如 equip a shield。equipment 是名词，指所需或所用的装备。equipment 的名称里保留 equip，再接 -ment。
```

对照现有释义、语境或例句。[equip-general](../entries/equip-general.yml)

下列来源行与现有 Notes 用一个空行分隔后追加。若同一 URL 或来源行已在 Notes 中，保留原出处。

```plain
Memory-aid source: Merriam-Webster.com Dictionary — equipment. https://www.merriam-webster.com/dictionary/equipment
```

证据见 X01。

#### WF08 dispense

[dispense-general](../entries/dispense-general.yml) · dispense · 分发；发放

新增 `word_formation`。从机器动作认出装置名。 聚焦发放义，不展开 dispense with 等其他搭配。

```plain
装置执行 dispense（发放）的动作，执行这个动作的装置可以叫 dispenser。dispense → dispenser 的拼写是在末尾加 r。
```

对照现有释义、语境或例句。[dispenser-engineering](../entries/dispenser-engineering.yml)

Notes 原样保留。

#### WF09 disguised

[disguised-combat](../entries/disguised-combat.yml) · disguised · 伪装的

新增 `word_formation`。同时记过去分词词形和 as 搭配。 本条例句是被动形式。

```plain
disguise（伪装）已经以 e 结尾，加 -d 得到 disguised。be disguised as 后可接伪装成的身份，如 disguised as a traveler。
```

本项按当前词条的释义、例句或 Notes 分析。

下列来源行与现有 Notes 用一个空行分隔后追加。若同一 URL 或来源行已在 Notes 中，保留原出处。

```plain
Memory-aid source: Merriam-Webster.com Dictionary — disguise. https://www.merriam-webster.com/dictionary/disguise
```

证据见 X02。

#### WF10 landing

[landing-general](../entries/landing-general.yml) · landing · 着陆；登陆；登岸处。

新增 `word_formation`。用动词和动作名词联记。 仅解释共有的着陆义，词条其他地点义保留。

```plain
land 在这里是“着陆”的动词。land + -ing 得到 landing，可指着陆这一动作。a landing 也可指一次着陆。
```

对照现有释义、语境或例句。[land-motion](../entries/land-motion.yml)

Notes 原样保留。

#### WF11 multiline range

[multiline-range-tech](../entries/multiline-range-tech.yml) · multiline range · 跨行范围

新增 `word_formation`。用 multi- 读懂格式化器的范围说明。 具体端点和边界由工具定义。

```plain
multi-（多）+ line（行）构成 multiline，再加 range（范围）。指跨越不止一行的范围。
```

本项按当前词条的释义、例句或 Notes 分析。

Notes 原样保留。

#### WF12 re-render

[rerender-react](../entries/rerender-react.yml) · re-render · 重新渲染，再次运行组件以计算下一次界面输出

新增 `word_formation`。复用 re- 来记框架常用动词。 只覆盖渲染阶段，不从词形推断 DOM 一定变更。

```plain
re-（再次）+ render（渲染）。re-render 是再次运行渲染以计算界面输出。
```

对照现有释义、语境或例句。[render-react](../entries/render-react.yml)

Notes 原样保留。

#### WF13 subagent

[subagent-ai](../entries/subagent-ai.yml) · subagent · 子智能体；受委派处理部分任务并返回结果的智能体。

新增 `word_formation`。由已有 agent 识别子任务执行者。 当前是智能体工具语境，权限和上下文由具体宿主决定。

```plain
sub- + agent 在此表示从属于某个任务分工的智能体。主智能体可把其中一部分工作交给 subagent。
```

对照现有释义、语境或例句。[agent-ai](../entries/agent-ai.yml)

Notes 原样保留。

#### WF14 unordered list

[unordered-list-html](../entries/unordered-list-html.yml) · unordered list · 无序列表，条目的先后顺序不影响内容含义的列表

新增 `word_formation`。用前缀区分有序与无序列表。 学习列表语义。页面上仍有实际排列。

```plain
un- + ordered 构成 unordered。这里指条目的先后对列表含义无关紧要，例如一组支持的文件格式。
```

对照现有释义、语境或例句。[ordered-list-html](../entries/ordered-list-html.yml)

Notes 原样保留。

#### WF15 interact

[interact-general](../entries/interact-general.yml) · interact · 交互

新增 `word_formation`。从前缀和搭配同时记“交互”。 现代拆字只用于学习；当前例句里的对象是装置。

```plain
inter-（彼此、相互）+ act（行动）。常用 interact with，把交互对象放在 with 后面。
```

本项按当前词条的释义、例句或 Notes 分析。

Notes 原样保留。

#### WF16 decayed

[decayed-general](../entries/decayed-general.yml) · decayed · 腐朽的；腐败或衰损的。

新增 `word_formation`。把原来的变化与变化后的状态连起来。 沿用现有武器语境，腐朽与生锈需按原 Context 分别理解。

```plain
decay（腐朽、衰败）+ -ed 得到 decayed。作形容词时描述已经出现腐朽或衰损的状态。
```

本项按当前词条的释义、例句或 Notes 分析。

Notes 原样保留。

#### WF17 regression test

[regression-test-tech](../entries/regression-test-tech.yml) · regression test · 回归测试

新增 `word_formation`。看出“回归”是测试目的。 沿用软件义，不涉及统计术语。

```plain
regression（行为退化、回归问题）+ test（测试）。regression test 是用来检查已有行为在变更后是否退化的测试。
```

对照现有释义、语境或例句。[regression-tech](../entries/regression-tech.yml)

Notes 原样保留。

#### WF18 retry

[retry-workflow](../entries/retry-workflow.yml) · retry · 重试；在先前尝试失败后再次进行尝试。

新增 `word_formation`。遇到重试提示时能从词形读出下一步。 原 Notes 对副作用及回滚的说明继续保留。

```plain
re-（再次）+ try（尝试）。失败后再试一次就是 retry。
```

本项按当前词条的释义、例句或 Notes 分析。

Notes 原样保留。

#### WF19 deduplicate

[deduplicate-tech](../entries/deduplicate-tech.yml) · deduplicate · 去重；识别重复项并按规则保留一份。

新增 `word_formation`。用基础词看懂工具动词。 具体相等规则由任务决定，现有名词说明留在 Notes。

```plain
deduplicate 里的 duplicate 是重复项，de- 表示移除或消减。把重复项识别后按规则保留所需数量，就是去重。
```

本项按当前词条的释义、例句或 Notes 分析。

Notes 原样保留。

### 易混义项

#### N05 attach

[attach-general](../entries/attach-general.yml) · attach · 连接、附加

新增 `mnemonic`。给简单词补上宾语和介词位置。 画面取自当前例句的普通场景，不增加游戏机制。

```plain
记忆画面：把风扇装到平台上。记 attach A to B，前面放装上去的 A，to 后面放接到的 B。
```

对照现有释义、语境或例句。[detach-general](../entries/detach-general.yml) · [reattach-general](../entries/reattach-general.yml)

Notes 原样保留。

#### N06 drop

[drop-fall-general](../entries/drop-fall-general.yml) · drop · 掉落；使落下

新增 `mnemonic`。用意图和主语分开同形义项。 为三张现有卡的用法对照；掉落也可以是有意动作，具体看句子。

```plain
用法助记：手滑掉了接头，用 drop 描述落下。主动丢下物品，用 drop 描述丢弃。说 the connection drops 时，drop 以连接为主语，意思是掉线。
```

对照现有释义、语境或例句。[drop-general](../entries/drop-general.yml) · [drop-connection](../entries/drop-connection.yml)

Notes 原样保留。

#### N07 landing

[landing-general](../entries/landing-general.yml) · landing · 着陆；登陆；登岸处。

新增 `mnemonic`。分清仍在下降与落地过程。 原创简化画面，适用于当前飞行义。

```plain
记忆画面：飞行器先往下降，再碰到地面。descent 是下降过程，landing 是最终的着陆阶段。
```

对照现有释义、语境或例句。[descent-flight](../entries/descent-flight.yml) · [land-motion](../entries/land-motion.yml)

Notes 原样保留。

#### N08 hover

[hover-general](../entries/hover-general.yml) · hover · 悬停；在空中保持相近位置。

新增 `mnemonic`。用运动方向和位置快速选词。 原创画面，适用于当前空间动作义。

```plain
记忆画面：给空中的角色画三个箭头。hover 原地悬停，glide 朝远处滑翔，dive 明显向下俯冲或潜入。
```

对照现有释义、语境或例句。[glide-general](../entries/glide-general.yml) · [dive-general](../entries/dive-general.yml)

Notes 原样保留。

#### N09 ridge

[ridge-terrain](../entries/ridge-terrain.yml) · ridge · 山脊

新增 `mnemonic`。把山脊、山谷、斜坡放进同一幅地图。 原创地形示意，地貌的实际形状可以变化。

```plain
记忆画面：山顶拉一条长线是 ridge，山之间低下去是 valley，两者之间向下倾斜的山面是 slope。
```

对照现有释义、语境或例句。[valley-terrain](../entries/valley-terrain.yml) · [slope-terrain](../entries/slope-terrain.yml)

Notes 原样保留。

#### N10 canyon

[canyon-terrain](../entries/canyon-terrain.yml) · canyon · 峡谷

新增 `mnemonic`。遇到地名或地图提示时区分地貌。 原创示意，只作释义对照。

```plain
记忆画面：沿一条很深的谷底走，两侧是陡壁。canyon 把注意力放在这道深峡谷，cliff 指陡壁，valley 是更宽泛的谷地概念。
```

对照现有释义、语境或例句。[valley-terrain](../entries/valley-terrain.yml) · [cliff-terrain](../entries/cliff-terrain.yml)

Notes 原样保留。

#### N11 draw

[draw-general](../entries/draw-general.yml) · draw · 拉开（弓）；拔出（武器）。

新增 `mnemonic`。同一个动词按装备选择译法。 当前只学习武器搭配。

```plain
用法助记：看到 bow，draw 可以是拉弓；看到 sword，draw 可以是拔剑。先看宾语再决定动作。
```

对照现有释义、语境或例句。[bow-weapon](../entries/bow-weapon.yml) · [sword-weapon](../entries/sword-weapon.yml)

Notes 原样保留。

#### N12 raw

[raw-food](../entries/raw-food.yml) · raw · 生的；未烹饪的。

新增 `mnemonic`。把食材状态与烹饪手段分开。 用当前普通义辨析，游戏物品名按实际本地化。

```plain
记忆画面：先看食材有没有熟，再看怎么做。raw 说未烹饪，roast 说烤制，fry 说用热油煎或炸。
```

对照现有释义、语境或例句。[roast-food](../entries/roast-food.yml) · [fry-cooking](../entries/fry-cooking.yml)

Notes 原样保留。

#### N13 gear

[gear-equipment](../entries/gear-equipment.yml) · gear · 装备；为某种活动准备的器材。

新增 `mnemonic`。从总类到常见功能分类。 原创示意，具体装备分类由语境决定。

```plain
记忆画面：出发前先看整套 gear，再单独看防护用的 armor 和攻击用的 weapon。
```

对照现有释义、语境或例句。[armor-equipment](../entries/armor-equipment.yml) · [weapon-weapon](../entries/weapon-weapon.yml)

Notes 原样保留。

#### N14 defer

[defer-go](../entries/defer-go.yml) · defer · 延迟调用，在外围函数返回前按后进先出顺序运行已安排的调用

新增 `mnemonic`。避免把关键词理解为任意稍后执行。 使用现有 Go 卡的规则，只说明当前调用语境。

```plain
用法助记：普通 defer 可泛指推迟。读 Go 代码时，看到 defer 就去找外围函数何时结束，多个已安排的调用按后进先出执行。
```

对照现有释义、语境或例句。[defer-tech](../entries/defer-tech.yml)

Notes 原样保留。

#### N15 registry

[registry-docker](../entries/registry-docker.yml) · registry · 镜像注册中心；提供镜像存储、分发和拉取的服务。

新增 `mnemonic`。看到“仓库”译法时先辨认存放的对象。 限于这两个义项，Docker 也可使用 repository 一词。

```plain
记忆画面：一边放镜像供拉取，一边放源码和版本历史。前者对应当前 Docker registry 卡，后者对应 Git repository 卡。
```

对照现有释义、语境或例句。[repository-git](../entries/repository-git.yml)

Notes 原样保留。

#### N16 landmark

[landmark-html](../entries/landmark-html.yml) · landmark · 地标区域，可供辅助技术识别和导航的页面区域

新增 `mnemonic`。用真实空间的定位概念理解页面导航。 原创比喻。具体哪些元素生成地标仍以现有 Context 或文档为准。

```plain
用法助记：纸质地图上先找显眼标记；在 HTML 里可把 landmark 想成帮助定位的页面区域。
```

对照现有释义、语境或例句。[land-motion](../entries/land-motion.yml)

Notes 原样保留。

#### N17 sail

[sail-watercraft](../entries/sail-watercraft.yml) · sail · 帆；利用风力推动船只的布面

新增 `mnemonic`。在水上装置场景区分物件与操控。 sail 的动词航行义仍保留在 Notes。

```plain
记忆画面：船上有吃风的布面和控制方向的人。名词 sail 指那面帆，steer 则描述控制方向的动作。
```

对照现有释义、语境或例句。[steer-general](../entries/steer-general.yml)

Notes 原样保留。

#### N18 deadline

[deadline-go](../entries/deadline-go.yml) · deadline · 截止时间，指定操作最晚应结束的时刻

新增 `mnemonic`。用两个排查问题区分常见相似词。 只使用当前 Go 场景的释义。

```plain
用法助记：看到 deadline，问“最迟几点结束”；看到 deadlock，问“谁在等谁”。一个关心时间界限，一个关心互相卡住的执行。
```

对照现有释义、语境或例句。[deadlock-go](../entries/deadlock-go.yml)

Notes 原样保留。

### 日语词形与句型

#### WF20 電話がある

[denwa-ga-aru](../entries/denwa-ga-aru.yml) · 電話がある · 有电话；接到电话

新增 `word_formation`。从名词记到来电表达。 范围限于当前来电义，不概括所有 ある 句型。

```plain
電話（电话）+ が + ある（有）。说已经接到的电话时，可把 ある 换成礼貌过去式 ありました，如 病院から電話がありました。
```

对照现有释义、语境或例句。[denwa](../entries/denwa.yml)

Notes 原样保留。

#### WF21 どなたをお探しですか？

[donata-o-osagashi-phone](../entries/donata-o-osagashi-phone.yml) · どなたをお探しですか？ · 请问您找哪位？

新增 `word_formation`。把接待句拆成疑问对象和礼貌谓语。 只分析现成句型。

```plain
どなた（哪位）+ を（标记寻找的对象）+ お探しですか（礼貌地问是否在找）。整句是在问对方找哪位。
```

对照现有释义、语境或例句。[donata](../entries/donata.yml) · [osagashi](../entries/osagashi.yml)

Notes 原样保留。

#### WF22 ～じゃない

[janai-grammar](../entries/janai-grammar.yml) · ～じゃない · 不是……

新增 `word_formation`。把口语缩约放进常见否定句型。 限于当前名词与な形容词场景，保留已有双向关联。

```plain
名词或な形容词词干 + じゃない，用于口语否定。じゃない 对应较完整的 ではない；需要礼貌表达时，本库另有 ではありません。
```

对照现有释义、语境或例句。[dewa-arimasen-grammar](../entries/dewa-arimasen-grammar.yml)

Notes 原样保留。

#### WF23 好奇心旺盛な子猫

[koukishin-ousei-na-koneko](../entries/koukishin-ousei-na-koneko.yml) · 好奇心旺盛な子猫 · 好奇心旺盛的小猫

新增 `word_formation`。用 な 找名词短语的修饰边界。 分析现代用法，不推断汉字来源。

```plain
好奇心旺盛な 修饰后面的 子猫。读到 な 时先向后找被修饰的名词，整句这里是“好奇心旺盛的小猫”。
```

对照现有释义、语境或例句。[koukishin](../entries/koukishin.yml) · [ousei](../entries/ousei.yml) · [koneko](../entries/koneko.yml)

Notes 原样保留。

#### WF24 英語で大丈夫ですか？

[eigo-de-daijoubu-phone](../entries/eigo-de-daijoubu-phone.yml) · 英語で大丈夫ですか？ · 可以说英语吗？

新增 `word_formation`。遇到替换语言的询问时能辨认槽位。 限于当前电话句型。

```plain
英語 + で + 大丈夫ですか。这里 で 引入所选的交流语言，整句问使用英语可不可以。
```

对照现有释义、语境或例句。[eigo](../entries/eigo.yml) · [daijoubu](../entries/daijoubu.yml)

Notes 原样保留。

#### WF25 日本語が話せません。

[nihongo-ga-hanasemasen](../entries/nihongo-ga-hanasemasen.yml) · 日本語が話せません。 · 我不会说日语

新增 `word_formation`。从否定能力形式读懂整句。 保留已有假名，不推广为所有可能表达只能用 が。

```plain
日本語 + が + 話せません。話せません 是表示不能说的礼貌形式；日本語 是这里不会说的语言。
```

对照现有释义、语境或例句。[nihongo](../entries/nihongo.yml) · [hanasemasen](../entries/hanasemasen.yml)

Notes 原样保留。

#### WF26 失礼ですが

[shitsurei-desu-ga](../entries/shitsurei-desu-ga.yml) · 失礼ですが · 冒昧问一下

新增 `word_formation`。在接待用语中辨认缓和语气的开头。 限于当前句型，不概括 が 的所有用法。

```plain
失礼です + が。放在提问前先表示冒昧，が 在这句话里帮助接到后面的礼貌请求。
```

对照现有释义、语境或例句。[onamae-wa](../entries/onamae-wa.yml)

Notes 原样保留。

#### WF27 何が～だ

[nani-ga-da-rhetorical](../entries/nani-ga-da-rhetorical.yml) · 何が～だ · 这算什么……；还谈什么……

新增 `word_formation`。用槽位读懂作品台词的反问。 情绪来自当前反问语境，不能只凭 何が 判断语气。

```plain
何が + 被质疑的词 + だ。在 何が正義だ！ 中，正義 是被质疑的说法，整体是在反驳对方。
```

对照现有释义、语境或例句。[nani-ga-seigi-da](../entries/nani-ga-seigi-da.yml)

Notes 原样保留。

#### N19 一歳

[issai](../entries/issai.yml) · 一歳 · 一岁

新增 `mnemonic`。把两个容易漏掉的促音放在同一次听辨里。 使用原有词条的假名与年龄义，不改变读音和录音。

```plain
辨音助记：一岁读 いっさい，八岁读 はっさい。两词都保留小 っ 的促音节奏，靠开头 い 和 は 分辨。
```

对照现有释义、语境或例句。[hassai](../entries/hassai.yml) · [nansai](../entries/nansai.yml)

Notes 原样保留。

#### N20 二つ

[futatsu](../entries/futatsu.yml) · 二つ · 两个

新增 `mnemonic`。用购物场景快速提取两种数量。 使用当前假名，不推广到所有物品的计数方式。

```plain
用法助记：点一个说 一つ，点两个说 二つ。结尾都是 つ，先用开头的 ひと 和 ふた 区分数量。
```

对照现有释义、语境或例句。[hitotsu](../entries/hitotsu.yml)

Notes 原样保留。

#### N21 ～から

[kara-reason-grammar](../entries/kara-reason-grammar.yml) · ～から · 因为……

新增 `mnemonic`。将同形助词放进两个可立即辨认的句子。 保留已有关联，适用范围为当前两张卡的句型。

```plain
用法助记：地名から来ました 回答“从哪儿来”；雨が降っているから 回答“为什么”。看到 から 先看前面是地点还是理由句。
```

对照现有释义、语境或例句。[kara-kimashita-origin](../entries/kara-kimashita-origin.yml)

Notes 原样保留。

#### N22 燃え上がる

[moeagaru](../entries/moeagaru.yml) · 燃え上がる · 熊熊燃烧；燃烧起来

新增 `mnemonic`。通过 が 和 を 区分自然发生与使其发生。 使用现成词条的自动词与他动词区别。

```plain
用法助记：画面里火焰自己升起来，用 炎が燃え上がる。有人使某物燃烧时，用“对象 + を + 燃やす”，を 前面是被点燃的对象。
```

对照现有释义、语境或例句。[moyasu](../entries/moyasu.yml)

Notes 原样保留。

#### N23 その方

[sono-kata](../entries/sono-kata.yml) · その方 · 那位；那个人（礼貌）

新增 `mnemonic`。同时区分指代距离和礼貌程度。 适用范围按当前 Context，具体选择由谈话关系决定。

```plain
用法助记：刚听对方提到某人时，用 その方 接着指那位。双方都能认出远处或共同想起的人时，可用 あの方；普通说法可用 あの人。
```

对照现有释义、语境或例句。[ano-kata](../entries/ano-kata.yml) · [ano-hito](../entries/ano-hito.yml)

Notes 原样保留。

#### N24 申します

[moushimasu](../entries/moushimasu.yml) · 申します · 叫作；自称（谦逊语）

新增 `mnemonic`。把自我介绍的固定表达与描述动作分开。 原有配对继续保留。

```plain
用法助记：礼貌自报姓名时先想“姓名 + と申します”。叙述某人报上姓名这件事时，可用普通动词 名乗る。
```

对照现有释义、语境或例句。[nanoru](../entries/nanoru.yml) · [to-moushimasu-grammar](../entries/to-moushimasu-grammar.yml)

Notes 原样保留。

### 词源单独审阅

#### ETY04 ingredient

[ingredient-cooking](../entries/ingredient-cooking.yml) · ingredient · 食材、原料

新增 `etymology`。用“放进去的一份”记忆食材的含义。 词典支持历史链；最后一句用于把历史意思联系到现代义。

```plain
ingredient 可追溯到拉丁语 ingredientem（进入某物的东西），关联 ingredi（进入）。放进混合物里的成分就是原料，料理语境中就是加入菜里的食材。
```

本项按当前词条的释义、例句或 Notes 分析。

下列来源行与现有 Notes 用一个空行分隔后追加。若同一 URL 或来源行已在 Notes 中，保留原出处。

```plain
Memory-aid source: Online Etymology Dictionary — ingredient. https://www.etymonline.com/word/ingredient
```

证据见 X03。

#### ETY05 canyon

[canyon-terrain](../entries/canyon-terrain.yml) · canyon · 峡谷

新增 `etymology`。用长窄通道联想峡谷的形状。 保留较早拼法提示的不同解释，不把联想当成确定的造词过程。

```plain
英语 canyon 借自墨西哥西班牙语 cañón。词源词典把峡谷义与“管道、深坑”等含义联系起来，也指出较早的拼法可能提示与 calle（街道）有关。
```

本项按当前词条的释义、例句或 Notes 分析。

下列来源行与现有 Notes 用一个空行分隔后追加。若同一 URL 或来源行已在 Notes 中，保留原出处。

```plain
Memory-aid source: Online Etymology Dictionary — canyon. https://www.etymonline.com/word/canyon
```

证据见 X04。

### 第二轮关联独立审阅

仅由左端追加右端稳定 ID。对照已有正反向声明及本计划原有待批提案后去重。表内对照使用本库当前义项和例句。文本与关联独立审批，已有列表成员全部保留。

| 编号 | 声明端 ID · Word · Meaning                                                                                                | 目标 ID · Word · Meaning                                                                                    | 联系                                     | 混淆风险                                                   |
|:------|:---------------------------------------------------------------------------------------------------------------------------|:-------------------------------------------------------------------------------------------------------------|:------------------------------------------|:------------------------------------------------------------|
| R11  | [detach-general](../entries/detach-general.yml) · detach · 分离、拆卸                                                     | [attach-general](../entries/attach-general.yml) · attach · 连接、附加                                       | 组装动作的两个方向                       | 观察部件从已连接变成分离，还是从分离变成连接。             |
| R12  | [detach-general](../entries/detach-general.yml) · detach · 分离、拆卸                                                     | [reattach-general](../entries/reattach-general.yml) · reattach · 重新连接；重新装上                         | 把再次装回的动作接到拆卸后               | reattach 强调先前分开的部件又接上了。                    |
| R13  | [equipment-equipment](../entries/equipment-equipment.yml) · equipment · 装备                                              | [equip-general](../entries/equip-general.yml) · equip · 装备                                                | 器材名词与使用装备的动作                 | 中文可同译“装备”，英语语法位置不同。                       |
| R14  | [dispense-general](../entries/dispense-general.yml) · dispense · 分发；发放                                               | [dispenser-engineering](../entries/dispenser-engineering.yml) · dispenser · 分配器；按需放出物品的装置。    | 把发放动作与执行装置相连                 | dispenser 是装置，dispense 是动作。                        |
| R15  | [landing-general](../entries/landing-general.yml) · landing · 着陆；登陆；登岸处。                                        | [descent-flight](../entries/descent-flight.yml) · descent · 下降；下降过程                                  | 按飞行阶段分清下降与着陆                 | 下降过程可能发生在着陆之前。                               |
| R16  | [hover-general](../entries/hover-general.yml) · hover · 悬停；在空中保持相近位置。                                        | [glide-general](../entries/glide-general.yml) · glide · 滑翔                                                | 区分在大致原位悬停与朝目的地方向滑翔     | 两种动作都可能发生在空中。                                 |
| R17  | [drop-fall-general](../entries/drop-fall-general.yml) · drop · 掉落；使落下                                               | [drop-connection](../entries/drop-connection.yml) · drop · （连接）中断；掉线                               | 用实物落下与连接中断区分三个义项中的两个 | 连接在此没有发生物理落下。                                 |
| R18  | [drop-fall-general](../entries/drop-fall-general.yml) · drop · 掉落；使落下                                               | [drop-general](../entries/drop-general.yml) · drop · 丢弃                                                   | 分清物体下落与主动丢弃                   | 单看动词不能断定本意，需看语境。                           |
| R19  | [defer-go](../entries/defer-go.yml) · defer · 延迟调用，在外围函数返回前按后进先出顺序运行已安排的调用                    | [defer-tech](../entries/defer-tech.yml) · defer · 推迟执行                                                  | 把普通推迟义与 Go 调用时机对照           | Go 中已安排的调用有明确的外围函数和执行顺序。              |
| R20  | [registry-docker](../entries/registry-docker.yml) · registry · 镜像注册中心；提供镜像存储、分发和拉取的服务。             | [repository-git](../entries/repository-git.yml) · repository · 仓库；保存项目文件及其版本历史的集合。       | 对照镜像分发服务与版本控制仓库           | 两者都可能被口头叫“仓库”。                                 |
| R21  | [denwa-ga-aru](../entries/denwa-ga-aru.yml) · 電話がある · 有电话；接到电话                                               | [denwa](../entries/denwa.yml) · 電話 · 电话                                                                 | 把来电句型连回基础名词                   | 電話 也能在 電話する 中表示打电话；電話がある 表示有电话。 |
| R22  | [issai](../entries/issai.yml) · 一歳 · 一岁                                                                               | [hassai](../entries/hassai.yml) · 八歳 · 八岁                                                               | 比较两个含促音的年龄读法                 | いっさい 与 はっさい 的起始音不同。                        |
| R23  | [futatsu](../entries/futatsu.yml) · 二つ · 两个                                                                           | [hitotsu](../entries/hitotsu.yml) · 一つ · 一个                                                             | 对照共用 つ 结尾的两个数量               | 分别对应用一个或两个来数，不能用汉字单独推断读法。         |
| R24  | [sono-kata](../entries/sono-kata.yml) · その方 · 那位；那个人（礼貌）                                                     | [ano-kata](../entries/ano-kata.yml) · あの方 · 那位；礼貌说法                                               | 比较礼貌指代时的 その 与 あの            | 选择时要看前文、对象距离与谈话双方是否共同知道。           |
| R25  | [nihongo-ga-hanasemasen](../entries/nihongo-ga-hanasemasen.yml) · 日本語が話せません。 · 我不会说日语                     | [hanasemasen](../entries/hanasemasen.yml) · 話せません · 不会说；不能说                                     | 从短词形定位完整日语能力句的谓语         | 日本語 填语言，話せません 表示不能说。                     |
| R26  | [unordered-list-html](../entries/unordered-list-html.yml) · unordered list · 无序列表，条目的先后顺序不影响内容含义的列表 | [ordered-list-html](../entries/ordered-list-html.yml) · ordered list · 有序列表，条目的先后顺序有意义的列表 | 对照顺序有意义与无意义的语义选择         | 项目符号和编号只是常见外观，判断依据是内容顺序。           |

### 第二轮来源与待核内容

以下外部正文已查阅。未另列外部来源的现代构词、句型或辨析使用目标条目与列出的对照条目。

| 编号 | 出处                                                                                               | 支持范围                                                          |
|:------|:----------------------------------------------------------------------------------------------------|:-------------------------------------------------------------------|
| X01  | [Merriam-Webster.com Dictionary — equipment](https://www.merriam-webster.com/dictionary/equipment) | 支持名词指供配备的器材。equip + -ment 是可见的现代词形对照。      |
| X02  | [Merriam-Webster.com Dictionary — disguise](https://www.merriam-webster.com/dictionary/disguise)   | 支持通过改变外表等方式掩饰身份或本质；-d 和 as 见当前词形和例句。 |
| X03  | [Online Etymology Dictionary — ingredient](https://www.etymonline.com/word/ingredient)             | 支持 ingredientem、ingredi 和混合物成分的联系。                   |
| X04  | [Online Etymology Dictionary — canyon](https://www.etymonline.com/word/canyon)                     | 支持墨西哥西班牙语形式、管道等义及早期拼法的另一种解释。          |

待核词源独立于以上文本。诸如 deadline 的早期来源、stable 各义项的历史分化，以及游戏专名的命名动机，需要查证完整历史词条或原始材料再提出精确文案。本轮仅为 deadline 的当前义项提供用法辨析。

### 台账补充

以下为第二轮新增提案的完整影响清单。关系右端的展示由导出派生。台账外的 ID 延续原台账处理。

| 稳定 ID                                                               | Word                   | 原台账无提案且无辅助或关联 | 新增提案       |
|:-----------------------------------------------------------------------|:------------------------|:----------------------------|:----------------|
| [ano-kata](../entries/ano-kata.yml)                                   | あの方                 | 否                         | R24 反向       |
| [attach-general](../entries/attach-general.yml)                       | attach                 | 是                         | N05、R11 反向  |
| [canyon-terrain](../entries/canyon-terrain.yml)                       | canyon                 | 否                         | N10、ETY05     |
| [deadline-go](../entries/deadline-go.yml)                             | deadline               | 是                         | N18            |
| [decayed-general](../entries/decayed-general.yml)                     | decayed                | 否                         | WF16           |
| [deduplicate-tech](../entries/deduplicate-tech.yml)                   | deduplicate            | 是                         | WF19           |
| [defer-go](../entries/defer-go.yml)                                   | defer                  | 是                         | N14、R19       |
| [defer-tech](../entries/defer-tech.yml)                               | defer                  | 是                         | R19 反向       |
| [denwa](../entries/denwa.yml)                                         | 電話                   | 是                         | R21 反向       |
| [denwa-ga-aru](../entries/denwa-ga-aru.yml)                           | 電話がある             | 是                         | WF20、R21      |
| [descent-flight](../entries/descent-flight.yml)                       | descent                | 否                         | R15 反向       |
| [detach-general](../entries/detach-general.yml)                       | detach                 | 是                         | WF06、R11、R12 |
| [disguised-combat](../entries/disguised-combat.yml)                   | disguised              | 是                         | WF09           |
| [dispense-general](../entries/dispense-general.yml)                   | dispense               | 是                         | WF08、R14      |
| [dispenser-engineering](../entries/dispenser-engineering.yml)         | dispenser              | 否                         | R14 反向       |
| [donata-o-osagashi-phone](../entries/donata-o-osagashi-phone.yml)     | どなたをお探しですか？ | 是                         | WF21           |
| [draw-general](../entries/draw-general.yml)                           | draw                   | 是                         | N11            |
| [drop-connection](../entries/drop-connection.yml)                     | drop                   | 是                         | R17 反向       |
| [drop-fall-general](../entries/drop-fall-general.yml)                 | drop                   | 是                         | N06、R17、R18  |
| [drop-general](../entries/drop-general.yml)                           | drop                   | 是                         | R18 反向       |
| [eigo-de-daijoubu-phone](../entries/eigo-de-daijoubu-phone.yml)       | 英語で大丈夫ですか？   | 是                         | WF24           |
| [equip-general](../entries/equip-general.yml)                         | equip                  | 是                         | R13 反向       |
| [equipment-equipment](../entries/equipment-equipment.yml)             | equipment              | 是                         | WF07、R13      |
| [futatsu](../entries/futatsu.yml)                                     | 二つ                   | 是                         | N20、R23       |
| [gear-equipment](../entries/gear-equipment.yml)                       | gear                   | 是                         | N13            |
| [glide-general](../entries/glide-general.yml)                         | glide                  | 是                         | R16 反向       |
| [hanasemasen](../entries/hanasemasen.yml)                             | 話せません             | 否                         | R25 反向       |
| [hassai](../entries/hassai.yml)                                       | 八歳                   | 是                         | R22 反向       |
| [hitotsu](../entries/hitotsu.yml)                                     | 一つ                   | 是                         | R23 反向       |
| [hover-general](../entries/hover-general.yml)                         | hover                  | 是                         | N08、R16       |
| [ingredient-cooking](../entries/ingredient-cooking.yml)               | ingredient             | 否                         | ETY04          |
| [interact-general](../entries/interact-general.yml)                   | interact               | 是                         | WF15           |
| [issai](../entries/issai.yml)                                         | 一歳                   | 是                         | N19、R22       |
| [janai-grammar](../entries/janai-grammar.yml)                         | ～じゃない             | 否                         | WF22           |
| [kara-reason-grammar](../entries/kara-reason-grammar.yml)             | ～から                 | 否                         | N21            |
| [koukishin-ousei-na-koneko](../entries/koukishin-ousei-na-koneko.yml) | 好奇心旺盛な子猫       | 是                         | WF23           |
| [landing-general](../entries/landing-general.yml)                     | landing                | 是                         | WF10、N07、R15 |
| [landmark-html](../entries/landmark-html.yml)                         | landmark               | 是                         | N16            |
| [moeagaru](../entries/moeagaru.yml)                                   | 燃え上がる             | 否                         | N22            |
| [moushimasu](../entries/moushimasu.yml)                               | 申します               | 否                         | N24            |
| [multiline-range-tech](../entries/multiline-range-tech.yml)           | multiline range        | 是                         | WF11           |
| [nani-ga-da-rhetorical](../entries/nani-ga-da-rhetorical.yml)         | 何が～だ               | 是                         | WF27           |
| [nihongo-ga-hanasemasen](../entries/nihongo-ga-hanasemasen.yml)       | 日本語が話せません。   | 是                         | WF25、R25      |
| [ordered-list-html](../entries/ordered-list-html.yml)                 | ordered list           | 是                         | R26 反向       |
| [raw-food](../entries/raw-food.yml)                                   | raw                    | 是                         | N12            |
| [reattach-general](../entries/reattach-general.yml)                   | reattach               | 否                         | R12 反向       |
| [registry-docker](../entries/registry-docker.yml)                     | registry               | 是                         | N15、R20       |
| [regression-test-tech](../entries/regression-test-tech.yml)           | regression test        | 是                         | WF17           |
| [repository-git](../entries/repository-git.yml)                       | repository             | 是                         | R20 反向       |
| [rerender-react](../entries/rerender-react.yml)                       | re-render              | 是                         | WF12           |
| [retry-workflow](../entries/retry-workflow.yml)                       | retry                  | 是                         | WF18           |
| [ridge-terrain](../entries/ridge-terrain.yml)                         | ridge                  | 否                         | N09            |
| [sail-watercraft](../entries/sail-watercraft.yml)                     | sail                   | 是                         | N17            |
| [shitsurei-desu-ga](../entries/shitsurei-desu-ga.yml)                 | 失礼ですが             | 是                         | WF26           |
| [sono-kata](../entries/sono-kata.yml)                                 | その方                 | 是                         | N23、R24       |
| [subagent-ai](../entries/subagent-ai.yml)                             | subagent               | 是                         | WF13           |
| [unordered-list-html](../entries/unordered-list-html.yml)             | unordered list         | 是                         | WF14、R26      |

### 审批与检查

- 常用词族待批范围为 WF06–WF19。
- 易混义项待批范围为 N05–N18。
- 日语词形与句型待批范围为 WF20–WF27 和 N19–N24。
- 词源单独待批范围为 ETY04–ETY05。
- 关联单独待批范围为 R11–R26。

获批后沿用上文实施程序。以获批编号为字段白名单，比较两次构建、现有导出字段、直接双向关联、词条内容及媒体哈希。

静态检查通过，检查了本节 187 处本地链接、来源、编号、字段、计数及两轮关系去重。原计划正文逐字节保留。除本计划外，3761 个已有文件与基线哈希匹配。本轮未构建、生成音频或执行 Anki 实机验证。

## 实施记录

全量实施范围为 WF01–WF27、ETY01–ETY05、N01–N24 和 R01–R26。

- 修改 59 个 YAML，共 91 个获批字段值。写入 56 项记忆辅助、11 条出处及 24 个关联列表。WF01 按原案移出 Notes 末尾短语，其余原 Notes 保留。
- 26 对新增关联涉及 50 个端点，导出新增 52 条直接邻接显示。逐行验证了原有 146 对和新增关联合计 172 对。
- Go 测试、vet、词库验证、两次导出构建及 Git 差异检查通过。两次构建字节完全相同。
- 对照实施前导出并验证 1833 行、18 列及其顺序、稳定 ID 集合、Audio、StudyOrder、全部辅助值、空辅助字段和正反向关系。已静态核对卡背的条件渲染。
- 词条改动符合获批字段白名单。所有 1875 个现有 MP3 的 SHA-256 匹配。已复查计划链接和行尾格式。人工听审、实际 Anki 导入、渲染和播放未核验。
