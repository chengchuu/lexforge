# 词汇记忆辅助内容补充计划

## 状态与确认边界

审查日期: 2026-11-03。状态: M01–M16 助记文案、R01–R34 新增关联，以及 WF01–WF84、ETY01–ETY18、N01–N23 的具体文案均已确认；已完成本计划的内容实施与本地验证。

用户已确认采用"全库审查、按价值分批补充"的方向，并于 2026-11-03 明确确认 M01–M16 的具体助记文案及其两端目标、WF01–WF84 的构成说明补充范围、ETY01–ETY18 的词源候选范围、R01–R34 的新增关联，以及 N01–N23 独立助记的价值和补充范围。WF、ETY 和 N 的逐条草案、来源与限定已补入[文案审阅稿](#逐条文案审阅稿)，具体文案已于 2026-11-03 获确认，包括来源补充和 N03 留空决定；ETY13 的直接借词对应另获用户确认，词典正文未复核的事实仍保留，但不再阻塞该项内容实施。R 组的确认仅覆盖关系，不扩大到额外助记文案或其他候选内容。

- 最初创建本计划只授权记录审查结果；本轮内容实施由用户后续的明确执行指令授权。录音、导出器和已安装的 Anki 模板不在本轮修改范围内。
- M01–M16 的具体 `mnemonic` 文案已确认，涉及 32 个词条；已按确认范围实施。
- R01–R34 的新增 `related_entries` 关系已于 2026-11-03 获确认，涉及 64 个词条；已按确认范围实施。批准关系不连带批准助记改写。
- WF01–WF84 的构成说明补充范围已于 2026-11-03 获确认，涉及 84 个词条；逐条具体文案与来源补充已于 2026-11-03 获确认，已按确认范围实施。
- ETY01–ETY18 的词源候选范围已于 2026-11-03 获确认，涉及 18 个词条；逐条具体文案、来源补充及不确定性限定已于 2026-11-03 获确认；ETY13 的直接借词对应已由用户补充确认，不再阻塞实施；仍不得声称已复核词典正文。已按确认范围实施。
- N01–N23 独立助记的价值和补充范围已于 2026-11-03 获确认，涉及 23 个词条；22 条具体助记文案、来源补充和 N03 留空决定已于 2026-11-03 获确认，已按确认范围实施。
- 审阅编号只用于本计划，不是词条 ID，也不是新增 YAML 字段。后续实施以明确的执行指令和已确认范围为准。

沿用原有[内容确认流程](VOCABULARY_MEMORY_AIDS_PLAN.md#human-confirmation-workflow)。该历史计划中的 21 组关系已经实施；本计划不重新创建这些关系，也不扩大此前批准范围。

## 目标与当前基线

让词源、构成说明、助记和关联帮助理解、记忆与辨析。每项补充都应提供学习价值；不要求每个词条填满四类信息，也不设置覆盖比例。

本次检查了全部 643 个词条的词形、当前释义、语境和 Notes，覆盖 401 个英语词条及 242 个日语词条。审查时:

| 项目            | 当前数量                              |
|:----------------|:--------------------------------------|
| etymology       | 0 个词条                              |
| word_formation  | 0 个词条                              |
| mnemonic        | 0 个词条                              |
| related_entries | 21 组既有关系，覆盖 42 个双向展示端点 |

初筛得到 203 个不同的候选词条，其中英语 120 个、日语 83 个。其余 440 个本轮建议维持现状，已有内容和关联保留。

| 补充方向           | 候选数量              | 内容状态                                  |
|:-------------------|:----------------------|:------------------------------------------|
| Etymology          | 18 个词条             | 18 条文案已确认；ETY13 正文待复核         |
| Word Formation     | 84 个词条             | 84 条具体文案已确认                       |
| Mnemonic           | 118 个词条            | M、N 文案已确认；N03 留空，R 不含额外助记 |
| 新增 Related Words | 34 组，涉及 64 个词条 | 34 组关联已确认，不包含额外助记文案       |

各类候选有重叠，数量不能直接相加。118 个助记候选包括 M 组的 32 个端点、R 组的 64 个端点及 N 组的 23 个词条；pronunciation-general 在 R16 和 N04 中重复出现，因此去重后为 118 个。四类方向合并去重后为 203 个，并非要求最终修改 203 个词条。

这些是带日期的审查基线。实施前重新检查实际词库和 Git 状态；若词条、释义或关系已改变，应更新受影响的候选及计数，不直接沿用旧快照。

## 第一批: M01–M16 助记文案

这 16 组关系已经存在，涉及 32 个不同词条。本批内容范围仅为添加 `mnemonic`，不新增关系、不改释义。

每行的已确认文案已分别写入该行两个词条的 `mnemonic`。如某一端已有内容或实施前发生变化，应先重新审阅，不覆盖已有文案。M01–M16 已于 2026-11-03 获用户确认，已于 2026-11-03 实施。

| 编号 | 词条 A: 稳定 ID、词语、当前释义                                              | 词条 B: 稳定 ID、词语、当前释义                                                   | 拟写入两端的 mnemonic 文案                                                                               | 状态   |
|:-----|:-----------------------------------------------------------------------------|:----------------------------------------------------------------------------------|:---------------------------------------------------------------------------------------------------------|:-------|
| M01  | [artifact-tech](../entries/artifact.yml) · artifact · 构建产物               | [artifact-workflow](../entries/artifact-workflow.yml) · artifact · 工作成果文件   | 构建产物是工作成果的一种；两个 artifact 的范围不同，不要把所有工作文件都理解为编译输出。                 | 已实施 |
| M02  | [closure-tech](../entries/closure-tech.yml) · closure · 闭包                 | [disclosure-workflow](../entries/disclosure-workflow.yml) · disclosure · 披露说明 | closure 在这里是闭包，disclosure 是披露说明；拼写相近，但当前两个义项不是反义词。                        | 已实施 |
| M03  | [encode-tech](../entries/encode-tech.yml) · encode · 编码                    | [decode-tech](../entries/decode-tech.yml) · decode · 解码                         | encode 按规则改变数据的表示，decode 按对应规则解释或还原；按转换方向联想。                               | 已实施 |
| M04  | [infer-tech](../entries/infer-tech.yml) · infer · 推断                       | [imply-general](../entries/imply-general.yml) · imply · 暗示                      | imply 是表达者暗示，infer 是接收者根据线索推断；注意信息方向。                                           | 已实施 |
| M05  | [literal-programming](../entries/literal-programming.yml) · literal · 字面量 | [literal-text](../entries/literal-text.yml) · literal · 按原样解释的；字面的      | 名词 literal 指源码中的字面量，形容词 literal 强调按原样解释；结合词性和语境记忆。                       | 已实施 |
| M06  | [mount-tech](../entries/mount-tech.yml) · mount · 挂载                       | [unmount-tech](../entries/unmount-tech.yml) · unmount · 卸载                      | mount 把组件接入渲染环境，unmount 移除并清理；这里不是软件包的安装与卸载。                               | 已实施 |
| M07  | [preserve-tech](../entries/preserve-tech.yml) · preserve · 保留；保持不变    | [retain-tech](../entries/retain-tech.yml) · retain · 保留；继续维持               | 两者都可表示保留；preserve 常强调保持已有状态，retain 常强调继续保有。结合对象和搭配理解，不设互斥规则。 | 已实施 |
| M08  | [register-tech](../entries/register-tech.yml) · register · 注册              | [unregister-tech](../entries/unregister-tech.yml) · unregister · 取消注册         | register 将对象或处理器加入系统集合，unregister 将其移除；按加入与清理两个方向联想。                     | 已实施 |
| M09  | [reproduce-tech](../entries/reproduce-tech.yml) · reproduce · 复现           | [reproducible-tech](../entries/reproducible-tech.yml) · reproducible · 可复现的   | reproduce 表示复现的动作，reproducible 表示能够复现的性质；区分动词与形容词。                            | 已实施 |
| M10  | [resolve-tech](../entries/resolve-tech.yml) · resolve · 解析；确定所指对象   | [resolve-workflow](../entries/resolve-workflow.yml) · resolve · 解决；处理完毕    | resolve 引用时是确定所指对象，resolve 问题时是解决问题；先看处理的对象。                                 | 已实施 |
| M11  | [revision-git](../entries/revision-git.yml) · revision · 代码版本；提交版本  | [revision-general](../entries/revision-general.yml) · revision · 修订             | Git 中的 revision 可以是代码或提交版本，普通工作语境中的 revision 可以是修订；区分版本对象与修订。       | 已实施 |
| M12  | [scope-lexical](../entries/scope-lexical.yml) · scope · 作用域               | [scope-workflow](../entries/scope-workflow.yml) · scope · 工作范围                | 代码中的 scope 是名称或绑定的可见范围，工作中的 scope 是任务边界；先判断讨论的是代码还是工作范围。       | 已实施 |
| M13  | [staged-git](../entries/staged-git.yml) · staged · 已暂存的                  | [unstaged-git](../entries/unstaged-git.yml) · unstaged · 未暂存的                 | staged 表示更改已进入 Git 索引，unstaged 表示更改未进入索引；未暂存不等于未跟踪。                        | 已实施 |
| M14  | [hanasu](../entries/hanasu.yml) · 話す · 说；讲话                            | [hanasemasen](../entries/hanasemasen.yml) · 話せません · 不会说；不能说           | 話す表示说话这一行为；話せません表示不能或不会说，区别于单纯表示不说的話しません。                       | 已实施 |
| M15  | [kyoushi](../entries/kyoushi.yml) · 教師 · 教师                              | [sensei](../entries/sensei.yml) · 先生 · 老师等称呼                               | 教師是职业名称，先生是称呼或尊称，也可用于医生等人；通常不用先生自称。                                   | 已实施 |
| M16  | [kaishain](../entries/kaishain.yml) · 会社員 · 公司职员                      | [shain](../entries/shain.yml) · 社員 · 公司成员；职员                             | 会社員侧重职业类别，社員侧重所属公司的成员身份；不要按两个完全不同的职业记忆。                           | 已实施 |

其余五组既有关联本轮不新增 Mnemonic:

- 誰／どなた、何歳／おいくつ、あの人／あの方: 当前 Context 已较直接地说明礼貌区别，先避免重复展示。
- 救う／救える、やる／やれる: 优先整理为构成说明，原有关联继续保留。

## 第二批: WF01–WF84 构成说明候选

确认记录: 用户于 2026-11-03 确认 WF01–WF84 构成说明的补充范围及对应的 84 个词条。此表保留目标词条和当前释义；具体构成文案与依据见[文案审阅稿](#逐条文案审阅稿)，已于同日获用户确认。已于 2026-11-03 实施。

`word_formation` 用于有依据的现代构词、缩略、复合或语法组合。优先从已有 Notes 提炼短说明，原有详细解释与出处保留。表格只登记候选，不能据此直接写入未经核对的拆分。

名称类只解释可识别的成分，不将字面拆解表述为品牌、机构或人物的历史命名来源。日语角色口语必须保留作品语境与特殊形式限定，不推广为普遍适用的变形规则。

### 英语派生与前缀 (23 个)

核对派生基础、词缀功能与实际技术义，不仅按拼写切分。

| 编号 | 现有词条                                                                | 当前释义                                               |
|:-----|:------------------------------------------------------------------------|:-------------------------------------------------------|
| WF01 | [actionable-tech](../entries/actionable-tech.yml) · actionable          | 可据以采取行动的                                       |
| WF02 | [accessor-tech](../entries/accessor-tech.yml) · accessor                | 属性访问器                                             |
| WF03 | [adapter-tech](../entries/adapter-tech.yml) · adapter                   | 适配器                                                 |
| WF04 | [blocker-workflow](../entries/blocker-workflow.yml) · blocker           | 阻碍事项                                               |
| WF05 | [correctness-tech](../entries/correctness-tech.yml) · correctness       | 正确性                                                 |
| WF06 | [maintainable-tech](../entries/maintainable-tech.yml) · maintainable    | 易维护的                                               |
| WF07 | [misalign-tech](../entries/misalign-tech.yml) · misalign                | 使错位；未正确对齐                                     |
| WF08 | [malformed-tech](../entries/malformed-tech.yml) · malformed             | 格式错误的                                             |
| WF09 | [multibyte-tech](../entries/multibyte-tech.yml) · multibyte             | 多字节的                                               |
| WF10 | [microtask-javascript](../entries/microtask-javascript.yml) · microtask | 微任务；在规定的检查点处理、通常先于下一个任务的工作。 |
| WF11 | [pre-existing-tech](../entries/pre-existing-tech.yml) · pre-existing    | 原先就存在的                                           |
| WF12 | [prefilter-tech](../entries/prefilter-tech.yml) · prefilter             | 预筛选器                                               |
| WF13 | [re-entry-tech](../entries/re-entry-tech.yml) · re-entry                | 再次进入                                               |
| WF14 | [uninitialized-tech](../entries/uninitialized-tech.yml) · uninitialized | 未初始化的                                             |
| WF15 | [unclaimed-workflow](../entries/unclaimed-workflow.yml) · unclaimed     | 尚未被承接的                                           |
| WF16 | [unmount-tech](../entries/unmount-tech.yml) · unmount                   | 卸载                                                   |
| WF17 | [unregister-tech](../entries/unregister-tech.yml) · unregister          | 取消注册                                               |
| WF18 | [unstaged-git](../entries/unstaged-git.yml) · unstaged                  | 未暂存的                                               |
| WF19 | [untracked-git](../entries/untracked-git.yml) · untracked               | 未跟踪的                                               |
| WF20 | [inaccessible-tech](../entries/inaccessible-tech.yml) · inaccessible    | 不可访问的                                             |
| WF21 | [nonthrowing-tech](../entries/nonthrowing-tech.yml) · nonthrowing       | 不会抛出异常的                                         |
| WF22 | [asynchronous-tech](../entries/asynchronous-tech.yml) · asynchronous    | 异步的                                                 |
| WF23 | [independently-tech](../entries/independently-tech.yml) · independently | 独立地                                                 |

### 英语复合与缩略 (16 个)

核对复合、缩略及名词与动词短语的形式区别，避免重复 Meaning。

| 编号 | 现有词条                                                                       | 当前释义                                         |
|:-----|:-------------------------------------------------------------------------------|:-------------------------------------------------|
| WF24 | [fallback-tech](../entries/fallback.yml) · fallback                            | 备用方案；回退方案                               |
| WF25 | [setup-tech](../entries/setup-tech.yml) · setup                                | 初始化设置                                       |
| WF26 | [cleanup-tech](../entries/cleanup-tech.yml) · cleanup                          | 清理                                             |
| WF27 | [callback-tech](../entries/callback-tech.yml) · callback                       | 回调函数                                         |
| WF28 | [handoff-workflow](../entries/handoff-workflow.yml) · handoff                  | 交接                                             |
| WF29 | [tradeoff-general](../entries/tradeoff-general.yml) · tradeoff                 | 取舍                                             |
| WF30 | [lifecycle-tech](../entries/lifecycle-tech.yml) · lifecycle                    | 生命周期                                         |
| WF31 | [lockfile-tech](../entries/lockfile-tech.yml) · lockfile                       | 依赖锁定文件                                     |
| WF32 | [toolchain-tech](../entries/toolchain-tech.yml) · toolchain                    | 工具链                                           |
| WF33 | [workload-tech](../entries/workload-tech.yml) · workload                       | 工作负载                                         |
| WF34 | [upstream-git](../entries/upstream.yml) · upstream                             | 上游的                                           |
| WF35 | [backtracking-algorithm](../entries/backtracking-algorithm.yml) · backtracking | 回溯；逐步尝试选择，撤销不合适的选择并继续搜索。 |
| WF36 | [breadth-first-tech](../entries/breadth-first-tech.yml) · breadth-first        | 广度优先的                                       |
| WF37 | [quicksort-algorithm](../entries/quicksort-algorithm.yml) · quicksort          | 快速排序；按基准值分区，再递归排序各分区。       |
| WF38 | [symlink-tech](../entries/symlink-tech.yml) · symlink                          | 符号链接                                         |
| WF39 | [typo-general](../entries/typo-general.yml) · typo                             | 笔误；打字错误                                   |

### 名称中的可识别成分 (4 个)

沿用现有 Notes 中的普通成分解释，不把组件含义当作命名史或产品属性的证据。

| 编号 | 现有词条                                             | 当前释义                                                        |
|:-----|:-----------------------------------------------------|:----------------------------------------------------------------|
| WF40 | [copilot-ai](../entries/copilot-ai.yml) · Copilot    | 此处指 GitHub Copilot，用于辅助软件开发的 AI 助手。             |
| WF41 | [deepseek-ai](../entries/deepseek-ai.yml) · DeepSeek | 深度求索；提供 AI 模型及相关服务的品牌。                        |
| WF42 | [dots-ai](../entries/dots-ai.yml) · dots             | OpenAI 的持续运行型个人智能体产品名称，可围绕用户目标推进任务。 |
| WF43 | [openclaw-ai](../entries/openclaw-ai.yml) · OpenClaw | 可自行部署的开源 AI 助手项目。                                  |

### 日语复合与缩略 (18 个)

核对词语组合、后缀或缩略依据，保留原写法与读法。

| 编号 | 现有词条                                                           | 当前释义             |
|:-----|:-------------------------------------------------------------------|:---------------------|
| WF44 | [anime](../entries/anime.yml) · アニメ                             | 动画                 |
| WF45 | [anime-guzzu](../entries/anime-guzzu.yml) · アニメグッズ           | 动漫周边             |
| WF46 | [burajiru-eaa](../entries/burajiru-eaa.yml) · ブラジルエアー       | 名称；ブラジルエアー |
| WF47 | [chuugokujin](../entries/chuugokujin.yml) · 中国人                 | 中国人               |
| WF48 | [denwa-bangou](../entries/denwa-bangou.yml) · 電話番号             | 电话号码             |
| WF49 | [fisshu-baagaa](../entries/fisshu-baagaa.yml) · フィッシュバーガー | 鱼肉汉堡             |
| WF50 | [gaikokujin](../entries/gaikokujin.yml) · 外国人                   | 外国人               |
| WF51 | [gentei-guzzu](../entries/gentei-guzzu.yml) · 限定グッズ           | 限定周边             |
| WF52 | [kaishain](../entries/kaishain.yml) · 会社員                       | 公司职员             |
| WF53 | [kenkyuusha](../entries/kenkyuusha.yml) · 研究者                   | 研究人员             |
| WF54 | [koneko](../entries/koneko.yml) · 子猫                             | 小猫                 |
| WF55 | [neko-chan](../entries/neko-chan.yml) · 猫ちゃん                   | 猫猫；对猫的亲昵称呼 |
| WF56 | [nihongo](../entries/nihongo.yml) · 日本語                         | 日语                 |
| WF57 | [nihonjin](../entries/nihonjin.yml) · 日本人                       | 日本人               |
| WF58 | [pawaa-denki](../entries/pawaa-denki.yml) · パワー電気             | 名称；パワー電気     |
| WF59 | [sakura-no-ki](../entries/sakura-no-ki.yml) · 桜の木               | 樱花树               |
| WF60 | [shinsekai](../entries/shinsekai.yml) · 新世界                     | 新世界               |
| WF61 | [yomikata](../entries/yomikata.yml) · 読み方                       | 读法                 |

### 日语屈折与语法组合 (23 个)

核对动词形式、修饰关系和完整句法环境；不把敬语片段当作可独立替换的普通动词。

| 编号 | 现有词条                                                                                                               | 当前释义                                            |
|:-----|:-----------------------------------------------------------------------------------------------------------------------|:----------------------------------------------------|
| WF62 | [hanasemasen](../entries/hanasemasen.yml) · 話せません                                                                 | 不会说；不能说                                      |
| WF63 | [sukueru](../entries/sukueru.yml) · 救える                                                                             | 能够拯救                                            |
| WF64 | [yareru](../entries/yareru.yml) · やれる                                                                               | 能做；能够做到                                      |
| WF65 | [shinde-miru](../entries/shinde-miru.yml) · 死んでみる                                                                 | 试着死一次看看                                      |
| WF66 | [yareru-koto](../entries/yareru-koto.yml) · やれること                                                                 | 能做到的事                                          |
| WF67 | [yaritai-dake-da-yo](../entries/yaritai-dake-da-yo.yml) · やりたいだけだよ                                             | 只是想做而已                                        |
| WF68 | [ni-nareru-grammar](../entries/ni-nareru-grammar.yml) · ～になれる                                                     | 能够成为……                                          |
| WF69 | [to-moushimasu-grammar](../entries/to-moushimasu-grammar.yml) · ～と申します                                           | 我叫……；我是……                                      |
| WF70 | [to-nanoru-grammar](../entries/to-nanoru-grammar.yml) · ～と名乗る                                                     | 自称是……                                            |
| WF71 | [gimu-toka-janai-yo](../entries/gimu-toka-janai-yo.yml) · 義務とかじゃないよ                                           | 这不是什么义务                                      |
| WF72 | [ore-wa-ningen-dakara](../entries/ore-wa-ningen-dakara.yml) · 俺は人間だから                                           | 因为我是人类                                        |
| WF73 | [hikari-ni-nareru-nda](../entries/hikari-ni-nareru-nda.yml) · 光になれるんだ                                           | 能够成为光                                          |
| WF74 | [jibun-jishin-no-chikara-de](../entries/jibun-jishin-no-chikara-de.yml) · 自分自身の力で                               | 凭自己的力量                                        |
| WF75 | [kotoba-wa-magenee](../entries/kotoba-wa-magenee.yml) · 言葉は曲げねェ                                                 | 不违背自己说过的话                                  |
| WF76 | [zettee-shinanee](../entries/zettee-shinanee.yml) · ぜってェ死なねー                                                   | 绝对不会死                                          |
| WF77 | [hokage-ni-naru-made](../entries/hokage-ni-naru-made.yml) · 火影になるまで                                             | 直到成为火影                                        |
| WF78 | [tantei-sa](../entries/tantei-sa.yml) · 探偵さ！                                                                       | 是侦探！                                            |
| WF79 | [kimi-ga-yobu-namae-yume-no-tsuzuki](../entries/kimi-ga-yobu-namae-yume-no-tsuzuki.yml) · きみが呼ぶなまえ～夢のつづき | 你呼唤的名字～梦的延续；《続 夏目友人帳》原声曲名。 |
| WF80 | [boku-wa-shinsekai-no-kami-ni-naru](../entries/boku-wa-shinsekai-no-kami-ni-naru.yml) · 僕は新世界の神になる           | 我将成为新世界的神                                  |
| WF81 | [tsuki-ni-kawatte-oshioki-yo](../entries/tsuki-ni-kawatte-oshioki-yo.yml) · 月にかわっておしおきよ！                   | 我要代表月亮惩罚你！                                |
| WF82 | [fuuin-kaijo-reriizu](../entries/fuuin-kaijo-reriizu.yml) · 封印解除（レリーズ）！                                     | 封印解除（Release）！                               |
| WF83 | [omachigae](../entries/omachigae.yml) · お間違え                                                                       | 弄错；「間違える」的礼貌表达形式                    |
| WF84 | [osagashi](../entries/osagashi.yml) · お探し                                                                           | 寻找；「探す」的礼貌表达形式                        |

优先整理的现成材料:

- [fallback-tech](../entries/fallback.yml): Notes 已区分名词 fallback 与动词 fall back。
- [hanasemasen](../entries/hanasemasen.yml): Notes 已有「話す → 話せる → 話せません」。
- [shinde-miru](../entries/shinde-miru.yml): Notes 已有「死ぬ → 死んで + みる」。
- [kimi-ga-yobu-namae-yume-no-tsuzuki](../entries/kimi-ga-yobu-namae-yume-no-tsuzuki.yml): Notes 已说明「きみが呼ぶ」修饰「なまえ」，可提炼修饰关系并保留详细分析。
- [deepseek-ai](../entries/deepseek-ai.yml)、[pawaa-denki](../entries/pawaa-denki.yml) 等名称: 可识别成分与"不推定命名来源"的限定应同时保留。

## 第三批: ETY01–ETY18 词源候选

确认记录: 用户于 2026-11-03 确认 ETY01–ETY18 词源候选范围。候选范围确认后，用户于同日确认了逐条具体文案及来源补充；ETY13 的直接借词对应另获用户确认；词典正文仍未复核。逐条文案及来源限制见[文案审阅稿](#逐条文案审阅稿)。已于 2026-11-03 实施。

`etymology` 只记录有可靠依据的历史来源、借入路径或语义演变。来源标题与 URL 放在 Notes；保留"可能""重建形式"等不确定性，不从拼写反推历史。

下表保留初筛时的方向、依据与核实状态，作为历史记录；本轮具体文本和来源核对结果以[文案审阅稿](#逐条文案审阅稿)为准。文案已获用户确认；ETY13 的词典证据限于检索摘要，后续另获用户确认直接借词对应。

| 编号  | 现有词条                                                                | 拟查证方向                         | 状态与限制                                       | 来源或核查入口                                                                       |
|:------|:------------------------------------------------------------------------|:-----------------------------------|:-------------------------------------------------|:-------------------------------------------------------------------------------------|
| ETY01 | [escape-tech](../entries/escape-tech.yml) · escape                      | 历史来源                           | 已找到词典依据；须保留重建形式限定               | [Merriam-Webster: escape](https://www.merriam-webster.com/dictionary/escape)         |
| ETY02 | [escape-parity-tech](../entries/escape-parity-tech.yml) · escape parity | 短语中 parity 的历史来源与当前义项 | 已找到词典依据；当前必须保持奇偶关系义           | [Merriam-Webster: parity](https://www.merriam-webster.com/dictionary/parity)         |
| ETY03 | [cache-tech](../entries/cache-tech.yml) · cache                         | 法语来源与藏放、储存的联系         | 已找到词典依据；不替代现代缓存定义               | [Merriam-Webster: cache](https://www.merriam-webster.com/dictionary/cache)           |
| ETY04 | [candidate-tech](../entries/candidate-tech.yml) · candidate             | 历史词形与候选人含义的联系         | 已找到词典依据；需避免把历史背景当成现代使用条件 | [Merriam-Webster: candidate](https://www.merriam-webster.com/dictionary/candidate)   |
| ETY05 | [etymology-language](../entries/etymology-language.yml) · etymology     | 希腊语来源及组成                   | 已找到词典依据；待整理                           | [Merriam-Webster: etymology](https://www.merriam-webster.com/dictionary/etymology)   |
| ETY06 | [mnemonic-memory](../entries/mnemonic-memory.yml) · mnemonic            | 与记忆有关的希腊语来源             | 已找到词典依据；历史来源与助记提示分开           | [Merriam-Webster: mnemonic](https://www.merriam-webster.com/dictionary/mnemonic)     |
| ETY07 | [vocabulary-general](../entries/vocabulary.yml) · vocabulary            | 传入路径及相关历史词形             | 已找到词典依据；保留词典中的 probably 限定       | [Merriam-Webster: vocabulary](https://www.merriam-webster.com/dictionary/vocabulary) |
| ETY08 | [syntax-tech](../entries/syntax-tech.yml) · syntax                      | 与共同排列有关的历史来源           | 已找到词典依据；不替代编程语法的当前定义         | [Merriam-Webster: syntax](https://www.merriam-webster.com/dictionary/syntax)         |
| ETY09 | [artifact-tech](../entries/artifact.yml) · artifact                     | 与人为制作有关的历史来源           | 已找到词典依据；保留当前构建产物义               | [Merriam-Webster: artifact](https://www.merriam-webster.com/dictionary/artifact)     |
| ETY10 | [idiomatic-tech](../entries/idiomatic.yml) · idiomatic                  | 经 idiom 追溯相关词形与语义        | 待核实派生链，不按拼写强拆                       | 现有词条与可靠英语词典                                                               |
| ETY11 | [gochisou](../entries/gochisou.yml) · ごちそう                          | 从奔走、周旋到款待的语义演变       | 已找到词典依据；待整理                           | [精選版 日本国語大辞典: 御馳走](https://kotobank.jp/word/御馳走-502226)              |
| ETY12 | [chisou](../entries/chisou.yml) · 馳走                                  | 与ご馳走相关的历史词义             | 已找到词典依据；避免添加无依据的故事细节         | [デジタル大辞泉等: 馳走](https://kotobank.jp/word/馳走-565923)                       |
| ETY13 | [aitemu](../entries/aitemu.yml) · アイテム                              | 借词来源                           | 待查证                                           | 可靠日语词典                                                                         |
| ETY14 | [eaa](../entries/eaa.yml) · エアー                                      | 借词来源                           | 已有 Notes 来源，待本轮重读核实                  | 词条 Notes 所列小学馆《デジタル大辞泉》「エア」                                      |
| ETY15 | [fisshu](../entries/fisshu.yml) · フィッシュ                            | 借词来源                           | 待查证，保留日语中的使用范围                     | 可靠日语词典                                                                         |
| ETY16 | [guzzu](../entries/guzzu.yml) · グッズ                                  | 借词来源                           | 待查证，区分英语与日语中的用法                   | 可靠日语词典                                                                         |
| ETY17 | [pawaa](../entries/pawaa.yml) · パワー                                  | 借词来源                           | 已有 Notes 来源，待本轮重读核实                  | 词条 Notes 所列小学馆《デジタル大辞泉》「パワー」                                    |
| ETY18 | [taoru](../entries/taoru.yml) · タオル                                  | 借词来源                           | 待查证                                           | 可靠日语词典                                                                         |

特别注意:

- escape 的词源材料涉及重建的拉丁语形式；最终内容须保留其重建性质，不写成全部有直接文献记录。
- [escape-parity-tech](../entries/escape-parity-tech.yml) 指连续反斜杠数量的奇偶关系，不改写为功能对等。
- 普通词的来源不自动证明同形 AI 品牌的命名来源。Codex、Cursor、Gemini、Grok、Hermes、Muse 等现有名称说明不自动改成品牌词源。
- 不为每个词根、同源词或借词原形新建词条。

## 第四批: R01–R34 新增关联候选

确认记录: 用户于 2026-11-03 确认全部 34 组新增关联。该确认仅覆盖关系，不授权额外改写助记文案。已于 2026-11-03 实施。

以下 34 组新增关系已确认，涉及 64 个现有词条。两端词语和释义来自本次审查时的源文件；理由只解释关联价值及风险，不是已经批准的 Mnemonic 文案。

批准 R 编号只确认对应关系，不授权改写 `mnemonic`、Meaning、Notes、先修关系或学习优先级。新助记内容应另行提供具体文案并确认。

| 编号 | 词条 A: 稳定 ID、词语、当前释义                                                                                                                  | 词条 B: 稳定 ID、词语、当前释义                                                                                                                                                | 关联理由与混淆风险                                                       | 状态   |
|:-----|:-------------------------------------------------------------------------------------------------------------------------------------------------|:-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|:-------------------------------------------------------------------------|:-------|
| R01  | [harness-tech](../entries/harness-tech.yml) · harness · 测试驱动程序                                                                             | [harness-ai](../entries/harness-ai.yml) · harness · 智能体运行框架；连接模型、上下文和工具，并协调执行循环与会话状态的软件层。                                                 | 同形不同技术义；测试辅助程序与智能体运行层不能混用。                     | 已实施 |
| R02  | [token-llm](../entries/token-llm.yml) · token · 词元；语言模型将文本编码后使用的基本处理单位。                                                   | [ownership-token-workflow](../entries/ownership-token-workflow.yml) · ownership token · 所有权标记                                                                             | 都含 token，但词元与资源所有权标记用途不同。                             | 已实施 |
| R03  | [deep-copy-javascript](../entries/deep-copy-javascript.yml) · deep copy · 深拷贝；复制嵌套数据，使相应对象引用与原数据隔离。                     | [shallow-copy-javascript](../entries/shallow-copy-javascript.yml) · shallow copy · 浅拷贝；复制外层结构，嵌套对象仍共享引用。                                                  | 区别嵌套对象引用是否共享；不能把 JSON 往返当成通用深拷贝。               | 已实施 |
| R04  | [flexbox-css](../entries/flexbox-css.yml) · flexbox · 弹性布局；主要沿一个维度分配空间和排列项目的布局。                                         | [grid-layout-css](../entries/grid-layout-css.yml) · grid layout · 网格布局；通过行和列组织项目的二维布局。                                                                     | 一维与二维布局侧重点不同；不应记成绝对互斥方案。                         | 已实施 |
| R05  | [graceful-degradation-web](../entries/graceful-degradation-web.yml) · graceful degradation · 优雅降级；在能力受限或部分功能失败时保留核心体验。  | [progressive-enhancement-web](../entries/progressive-enhancement-web.yml) · progressive enhancement · 渐进增强；先提供基础可用能力，再为更强环境增加功能。                     | 设计起点不同，可以共同使用；不是简单的好坏或反义关系。                   | 已实施 |
| R06  | [event-bubbling-dom](../entries/event-bubbling-dom.yml) · event bubbling · 事件冒泡；可冒泡事件从目标向祖先传播的阶段。                          | [event-capturing-dom](../entries/event-capturing-dom.yml) · event capturing · 事件捕获；事件沿传播路径向目标传递的阶段。                                                       | 传播方向不同；并非所有事件都冒泡。                                       | 已实施 |
| R07  | [event-delegation-dom](../entries/event-delegation-dom.yml) · event delegation · 事件委托；在祖先上监听并处理后代元素的事件。                    | [event-bubbling-dom](../entries/event-bubbling-dom.yml) · event bubbling · 事件冒泡；可冒泡事件从目标向祖先传播的阶段。                                                        | 委托是监听策略，冒泡是传播阶段；二者不是同义词。                         | 已实施 |
| R08  | [hoisting-javascript](../entries/hoisting-javascript.yml) · hoisting · 提升；声明在执行前影响所在作用域的行为。                                  | [temporal-dead-zone-javascript](../entries/temporal-dead-zone-javascript.yml) · temporal dead zone · 暂时性死区；词法绑定从作用域开始到初始化前不可访问的阶段。                | 声明影响作用域不等于已初始化；不能记成源码被移动。                       | 已实施 |
| R09  | [one-way-data-flow-frontend](../entries/one-way-data-flow-frontend.yml) · one-way data flow · 单向数据流；让状态沿明确方向传递并由指定入口更新。 | [two-way-binding-frontend](../entries/two-way-binding-frontend.yml) · two-way binding · 双向绑定；让状态值与界面输入保持相互同步的绑定方式。                                   | 表单双向绑定不等于取消整体数据流约束；不是同层面的绝对对立。             | 已实施 |
| R10  | [layout-browser-rendering](../entries/layout-browser-rendering.yml) · layout · 布局；根据样式与约束计算页面元素的位置和尺寸。                    | [paint-browser-rendering](../entries/paint-browser-rendering.yml) · paint · 绘制；把页面元素的视觉内容记录或绘制出来的过程。                                                   | 计算位置尺寸与绘制视觉内容不同；不保证所有更新都走完整流程。             | 已实施 |
| R11  | [async-await-javascript](../entries/async-await-javascript.yml) · async/await · 异步函数语法；用 await 组织基于 Promise 的控制流。               | [promise-javascript](../entries/promise-javascript.yml) · Promise · Promise；表示异步操作最终成功或失败结果的对象。                                                            | 语法与异步结果对象不同；await 不阻塞整个线程。                           | 已实施 |
| R12  | [cache-tech](../entries/cache-tech.yml) · cache · 缓存；暂存数据或计算结果以供后续复用的机制。                                                   | [memoization-programming](../entries/memoization-programming.yml) · memoization · 记忆化；缓存计算结果，在输入满足复用条件时避免重复计算。                                     | 通用缓存机制与按计算输入复用结果的策略不同；不要当成完全同义词。         | 已实施 |
| R13  | [type-assertion-tech](../entries/type-assertion-tech.yml) · type assertion · 类型断言                                                            | [cast-tech](../entries/cast-tech.yml) · cast · 类型转换写法；类型断言写法                                                                                                      | TypeScript 断言不增加运行时转换；cast 的实际含义取决于语言。             | 已实施 |
| R14  | [actual-behavior-workflow](../entries/actual-behavior-workflow.yml) · actual behavior · 实际行为                                                 | [expected-behavior-workflow](../entries/expected-behavior-workflow.yml) · expected behavior · 预期行为                                                                         | 应当发生与实际发生不同；预期不能代替观察。                               | 已实施 |
| R15  | [formation-general](../entries/formation-general.yml) · formation · 形成；构成                                                                   | [transformation-tech](../entries/transformation-tech.yml) · transformation · 转换                                                                                              | 形成或构成与转换不同；相似词形不等于同义。                               | 已实施 |
| R16  | [audio-general](../entries/audio-general.yml) · audio · 音频；录制或传输的声音。                                                                 | [pronunciation-general](../entries/pronunciation.yml) · pronunciation · 发音；读音                                                                                             | 声音资料与发音方式不同；有录音不证明发音正确。                           | 已实施 |
| R17  | [machigaeru](../entries/machigaeru.yml) · 間違える · 弄错；搞错                                                                                  | [machigai](../entries/machigai.yml) · 間違い · 错误；弄错                                                                                                                      | 动词与名词用法不同；不是把間違える机械去尾得到間違い。                   | 已实施 |
| R18  | [sagasu](../entries/sagasu.yml) · 探す · 寻找                                                                                                    | [osagashi](../entries/osagashi.yml) · お探し · 寻找；「探す」的礼貌表达形式                                                                                                    | 普通动词与礼貌结构形式不同；お探し不能独立替换所有探す。                 | 已实施 |
| R19  | [koneko](../entries/koneko.yml) · 子猫 · 小猫                                                                                                    | [neko-chan](../entries/neko-chan.yml) · 猫ちゃん · 猫猫；对猫的亲昵称呼                                                                                                        | 年龄与亲昵称呼不同；猫ちゃん不一定是幼猫。                               | 已实施 |
| R20  | [nihongo](../entries/nihongo.yml) · 日本語 · 日语                                                                                                | [nihonjin](../entries/nihonjin.yml) · 日本人 · 日本人                                                                                                                          | 語表示语言，人表示人；注意末尾与读音。                                   | 已实施 |
| R21  | [kimi](../entries/kimi.yml) · 君 · 你；用于较亲近、同辈或晚辈之间的称呼。                                                                        | [anata](../entries/anata.yml) · あなた · 你                                                                                                                                    | 称呼取决于人际关系与场合；不能按固定礼貌等级机械替换。                   | 已实施 |
| R22  | [boku](../entries/boku.yml) · 僕 · 我；男性自称                                                                                                  | [ore](../entries/ore.yml) · 俺 · 我；男性较随意或粗犷的自称                                                                                                                    | 自称语气和关系不同；不能仅凭说话者性别决定用法。                         | 已实施 |
| R23  | [gochisousama](../entries/gochisousama.yml) · ごちそうさま · 多谢款待                                                                            | [itadakimasu](../entries/itadakimasu.yml) · いただきます · 我开动了                                                                                                            | 饭后感谢与饭前表达不同；按使用时机联想。                                 | 已实施 |
| R24  | [gochisousama](../entries/gochisousama.yml) · ごちそうさま · 多谢款待                                                                            | [gochisousama-deshita](../entries/gochisousama-deshita.yml) · ごちそうさまでした · 多谢款待；谢谢这顿饭                                                                        | 都是饭后感谢，后者较礼貌；不误解为不同食物或不同对象。                   | 已实施 |
| R25  | [gomen-ne](../entries/gomen-ne.yml) · ごめんね · 抱歉啦；对不起                                                                                  | [gomennasai](../entries/gomennasai.yml) · ごめんなさい · 对不起；请原谅                                                                                                        | 熟人随意道歉与较郑重道歉不同；不把两者当成所有场合可互换。               | 已实施 |
| R26  | [gomennasai](../entries/gomennasai.yml) · ごめんなさい · 对不起；请原谅                                                                          | [sumimasen](../entries/sumimasen.yml) · すみません · 对不起；劳驾；不好意思                                                                                                    | 用法有重叠，但すみません还用于引起注意或麻烦别人；不是完全同义。         | 已实施 |
| R27  | [moushimasu](../entries/moushimasu.yml) · 申します · 叫作；自称（谦逊语）                                                                        | [nanoru](../entries/nanoru.yml) · 名乗る · 自报姓名；自称                                                                                                                      | 谦逊自我介绍与自报姓名或身份不同；自称不证明身份真实。                   | 已实施 |
| R28  | [ni-naru-grammar](../entries/ni-naru-grammar.yml) · 〜になる · 成为……；变成……                                                                    | [ni-nareru-grammar](../entries/ni-nareru-grammar.yml) · ～になれる · 能够成为……                                                                                                | 成为与能够成为不同；区别一般形式与可能形式。                             | 已实施 |
| R29  | [ni-naru-grammar](../entries/ni-naru-grammar.yml) · 〜になる · 成为……；变成……                                                                    | [ni-kawatte-grammar](../entries/ni-kawatte-grammar.yml) · ～にかわって · 代替……；代表……                                                                                        | 成为与代替或代表不同；不要因都含に就互换。                               | 已实施 |
| R30  | [kara-reason-grammar](../entries/kara-reason-grammar.yml) · ～から · 因为……                                                                      | [kara-kimashita-origin](../entries/kara-kimashita-origin.yml) · ～から来ました · 来自……                                                                                        | 原因与来源不同；按句法和语境识别から。                                   | 已实施 |
| R31  | [janai-grammar](../entries/janai-grammar.yml) · ～じゃない · 不是……                                                                              | [dewa-arimasen-grammar](../entries/dewa-arimasen-grammar.yml) · ～ではありません · 不是……                                                                                      | 名词或な形容词谓语的口语否定与礼貌否定不同；不推广为所有动词的转换规则。 | 已实施 |
| R32  | [fisshu](../entries/fisshu.yml) · フィッシュ · fish；鱼类相关借词                                                                                | [sakana-general](../entries/sakana-general.yml) · 魚 · 鱼                                                                                                                      | 借词多见商品或复合名称，魚是日常普通词；不要机械互换。                   | 已实施 |
| R33  | [lone-surrogate-tech](../entries/lone-surrogate-tech.yml) · lone surrogate · 未配对的代理码元                                                    | [surrogate-pair-tech](../entries/surrogate-pair-tech.yml) · surrogate pair · 代理对                                                                                            | 未配对的代理码元与成对表示不同；不要把一个码元直接当作完整字符。         | 已实施 |
| R34  | [same-origin-policy-web](../entries/same-origin-policy-web.yml) · same-origin policy · 同源策略；限制脚本访问不同源资源的浏览器安全机制。        | [cross-origin-resource-sharing-web](../entries/cross-origin-resource-sharing-web.yml) · cross-origin resource sharing · 跨源资源共享；服务器通过响应头授权浏览器进行跨源读取。 | 同源策略施加访问限制，CORS 授权特定跨源读取；不等于禁止所有跨源请求。    | 已实施 |

确认后的关系沿用现有机制: 一端声明一次，两边显示当前词语和释义；仅展示直接邻居，不自动补全传递关系。关联不改变先修顺序。

在本次基线不变且 R01–R34 全部获批、实施的前提下，唯一关系数将从 21 增至 55，RelatedWords 的非空端点将从 42 增至 106。若只实施部分候选，应按实际确认集合计算预期结果。

## 后续: N01–N23 独立助记候选

确认记录: 用户于 2026-11-03 确认 N01–N23 独立助记有学习价值，并确认其补充范围。下表保留初筛方向；22 条具体助记文案及 N03 的留空决定见[文案审阅稿](#逐条文案审阅稿)，已于同日获用户确认，已于 2026-11-03 实施。

这些候选不要求新增关联，目标是提炼容易遗漏的拼写、语气、概念边界或使用限制。具体文案形成后按内容确认流程处理。N04 与 R16 涉及同一词条，实施时不得重复或互相覆盖。

| 编号 | 现有词条                                                                                                             | 待整理方向                                                             |
|:-----|:---------------------------------------------------------------------------------------------------------------------|:-----------------------------------------------------------------------|
| N01  | [invaluable-general](../entries/invaluable-general.yml) · invaluable                                                 | 区分极有价值与无价值，不按词首机械推断否定义。                         |
| N02  | [deprecated-tech](../entries/deprecated.yml) · deprecated                                                            | 提炼现有 Notes 中"不等于已移除或损坏"的限定。                          |
| N03  | [fallback-tech](../entries/fallback.yml) · fallback                                                                  | 突出名词 fallback 与动词 fall back 的形式区别。                        |
| N04  | [pronunciation-general](../entries/pronunciation.yml) · pronunciation                                                | 突出 pronunciation 与 pronounce 的拼写变化，避免 pronounciation 误拼。 |
| N05  | [mnemonic-memory](../entries/mnemonic-memory.yml) · mnemonic                                                         | 提炼现有 Notes 中词首 m 不发音的提示，与词源说明分开。                 |
| N06  | [field-presence-tech](../entries/field-presence-tech.yml) · field presence                                           | 区分字段存在与字段值为真。                                             |
| N07  | [escaped-backslash-tech](../entries/escaped-backslash-tech.yml) · escaped backslash                                  | 区分源码中的转义表示与实际字符数量。                                   |
| N08  | [hito-o-norowaba-ana-futatsu](../entries/hito-o-norowaba-ana-futatsu.yml) · 人を呪わば穴二つ                         | 区分文语条件呪わば与现代条件形式，不机械类推。                         |
| N09  | [fuuin-kaijo-reriizu](../entries/fuuin-kaijo-reriizu.yml) · 封印解除（レリーズ）！                                   | 区分作品读法レリーズ与普通读法ふういんかいじょ。                       |
| N10  | [kosumo-saint-seiya](../entries/kosumo-saint-seiya.yml) · 小宇宙                                                     | 保留作品读法コスモ的限定，不推及所有小宇宙的读法。                     |
| N11  | [daijoubu](../entries/daijoubu.yml) · 大丈夫                                                                         | 强调依语境判断回答含义；如补充拒绝用法，先查证再写。                   |
| N12  | [onamae-wa](../entries/onamae-wa.yml) · お名前は                                                                     | 解释礼貌问句的省略，不将助词本身当作疑问标记。                         |
| N13  | [san-name-suffix](../entries/san-name-suffix.yml) · ～さん                                                           | 区分姓名后缀、对自己的称呼与数字三。                                   |
| N14  | [jin-nationality-suffix](../entries/jin-nationality-suffix.yml) · ～人                                               | 保留国籍后缀じん的读法限定，不混用其他人字读法。                       |
| N15  | [sa-ending-grammar](../entries/sa-ending-grammar.yml) · ～さ                                                         | 区分句尾语气与构成名词的接尾词さ。                                     |
| N16  | [no-ka-grammar](../entries/no-ka-grammar.yml) · ～のか                                                               | 区分疑问与反问，不把所有のか都译成强烈反问。                           |
| N17  | [promise-all-javascript](../entries/promise-all-javascript.yml) · Promise.all                                        | 提炼按输入顺序返回结果，以及拒绝不自动取消其他任务。                   |
| N18  | [block-formatting-context-css](../entries/block-formatting-context-css.yml) · block formatting context               | 避免把 BFC 记成内部所有外边距都不会折叠。                              |
| N19  | [floating-point-precision-javascript](../entries/floating-point-precision-javascript.yml) · floating-point precision | 保留有限二进制表示与量级条件，不把 EPSILON 当万能阈值。                |
| N20  | [virtual-dom-frontend](../entries/virtual-dom-frontend.yml) · virtual DOM                                            | 保留维护虚拟结构的成本，不声称总比直接操作 DOM 更快。                  |
| N21  | [http-caching-web](../entries/http-caching-web.yml) · HTTP caching                                                   | 突出 no-cache 与 no-store 的区别。                                     |
| N22  | [least-recently-used-cache](../entries/least-recently-used-cache.yml) · least recently used                          | 区分最近访问、访问频次与有效期策略。                                   |
| N23  | [preflight-request-web](../entries/preflight-request-web.yml) · preflight request                                    | 区分预检与实际请求，以及预检条件与凭据条件。                           |

## 逐条文案审阅稿

整理日期: 2026-11-03。本节覆盖 WF01–WF84、ETY01–ETY18、N01–N23 共 125 个审阅编号，其中已补充 84 条 `word_formation`、18 条 `etymology` 和 22 条 `mnemonic`；N03 已按决定留空。用户已于 2026-11-03 确认以下具体文本、来源补充和 N03 留空决定，确认原文为"已确认 WF01–WF84 → ETY01–ETY18 → N01–N23"。ETY13 另获用户确认直接借词对应，词典正文仍未复核；用户后续已明确要求执行，本轮已写入下列批准文案；N03 的说明仅作为留空记录，不写入字段。

### 阅读与来源约定

- "完整草案"栏是拟展示在卡片背面的纯文本，不含审阅编号、来源编号或审批说明。N03 明确标为留空建议，不把该行说明写入字段。
- WF 的大部分切分属于根据现代词义、词缀或语法规则作出的分析。依据列同时注明技术含义或作品限制；不把分析结果写成某作者的历史造词过程。
- ETY 记录历史来源或借入关系，保留重建形式和不确定路径。词源帮助记忆，不能取代当前释义或推导现代技术行为。
- N 是依据可核对事实拟写的记忆提示，文案组织属于本计划的提炼，不当作词典原文引用，也不新增关联。
- 名称与作品表达以现有词条中已确认的拼写、读法和语境为输入。本次未重核所有原作台词、音轨或品牌命名史；只在需要的地方查普通词义和语法，不补写未知设定。
- 每行的 S 编号链接到具体外部页面；完整标题、URL 和所支持的范围见本节末尾的来源清单。ETY13 的词典证据限于检索摘要，另有用户确认直接借词对应；两种来源分别记录，不能视作词典正文复核通过。
- 本轮已将该行使用的来源按"来源: 标题。URL"追加到 Notes，保留已有内容与出处；相同来源不重复追加。ETY13 单独标明用户确认与仅核对词典摘要的限制。作品信息已在 Notes 中的，保留原文，不另编来源。

### WF01–WF23: 英语派生与前缀

| 编号 | 目标词条                                                                | 完整草案或留空建议                                                                                                             | 核对依据与限定                                                                                                                                                                                                                              |
|:-----|:------------------------------------------------------------------------|:-------------------------------------------------------------------------------------------------------------------------------|:--------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| WF01 | [actionable-tech](../entries/actionable-tech.yml) · actionable          | action + -able。-able 在这里表示可被采取行动处理的；actionable 强调信息能转化为具体行动，不只是值得关注。                      | [S001](https://www.merriam-webster.com/dictionary/actionable)。按来源与当前义项整理；已写入词条。                                                                                                                                           |
| WF02 | [accessor-tech](../entries/accessor-tech.yml) · accessor                | access + -or。-or 可构成执行某种操作的人或事物的名称；这里的 accessor 是提供属性访问的机制，结合 getter、setter 理解。         | [S002](https://www.merriam-webster.com/dictionary/-or)、[S003](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Functions/get)。现代词形分析；MDN 核对技术义，不声称 MDN 记载造词历史。                                    |
| WF03 | [adapter-tech](../entries/adapter-tech.yml) · adapter                   | adapt + -er。adapt 表示使适应，-er 在这里指承担这种作用的事物；adapter 将一种接口适配到另一种接口。                            | [S004](https://www.merriam-webster.com/dictionary/adapter)。按来源与当前义项整理；已写入词条。                                                                                                                                              |
| WF04 | [blocker-workflow](../entries/blocker-workflow.yml) · blocker           | block + -er。block 表示阻碍，-er 在这里指造成阻碍的事物；工作中的 blocker 是阻止任务继续推进的问题。                           | [S005](https://www.merriam-webster.com/dictionary/blocker)。按来源与当前义项整理；已写入词条。                                                                                                                                              |
| WF05 | [correctness-tech](../entries/correctness-tech.yml) · correctness       | correct + -ness。-ness 把形容词转为表示性质或状态的名词；correctness 关注结果或行为是否正确。                                  | [S006](https://www.merriam-webster.com/dictionary/correctness)、[S007](https://stella.glasgow.ac.uk/grammar/unit-4-3-affixes.html)。按来源与当前义项整理；已写入词条。                                                                      |
| WF06 | [maintainable-tech](../entries/maintainable-tech.yml) · maintainable    | maintain + -able。-able 表示能够被……的；maintainable 关注代码或系统能否方便地维护。                                            | [S008](https://www.merriam-webster.com/dictionary/maintainable)。按来源与当前义项整理；已写入词条。                                                                                                                                         |
| WF07 | [misalign-tech](../entries/misalign-tech.yml) · misalign                | mis- + align。mis- 在这里表示错误或不当；misalign 是没有正确对齐，既可用于位置，也可用于目标或假设。                           | [S009](https://www.merriam-webster.com/dictionary/misalign)。按来源与当前义项整理；已写入词条。                                                                                                                                             |
| WF08 | [malformed-tech](../entries/malformed-tech.yml) · malformed             | mal- + formed。mal- 表示不良或错误；malformed 指形式或结构不符合要求，不只是内容不合预期。                                     | [S010](https://www.merriam-webster.com/dictionary/malformed)、[S011](https://www.merriam-webster.com/dictionary/mal-)。按来源与当前义项整理；已写入词条。                                                                                   |
| WF09 | [multibyte-tech](../entries/multibyte-tech.yml) · multibyte             | multi- + byte。multi- 表示多个；multibyte 指由多个字节表示或处理，字节数不等于字符数。                                         | [S012](https://www.merriam-webster.com/dictionary/multi-)。现代前缀分析；字节与字符的区分不等于具体编码的字节长度公式。                                                                                                                     |
| WF10 | [microtask-javascript](../entries/microtask-javascript.yml) · microtask | micro- + task。micro- 提供"小"的词形线索；JavaScript 中 microtask 是调度类别，不保证任务执行时间很短。                         | [S013](https://developer.mozilla.org/en-US/docs/Web/API/HTML_DOM_API/Microtask_guide)。按来源与当前义项整理；已写入词条。                                                                                                                   |
| WF11 | [pre-existing-tech](../entries/pre-existing-tech.yml) · pre-existing    | pre- + existing。pre- 表示在前；pre-existing 指在当前事件或改动之前就已经存在的。                                              | [S014](https://www.merriam-webster.com/dictionary/preexisting)。按来源与当前义项整理；已写入词条。                                                                                                                                          |
| WF12 | [prefilter-tech](../entries/prefilter-tech.yml) · prefilter             | pre- + filter。pre- 表示预先；当前名词 prefilter 是在后续处理之前先做过滤的预筛选器。                                          | [S015](https://www.merriam-webster.com/dictionary/prefilter)。按来源与当前义项整理；已写入词条。                                                                                                                                            |
| WF13 | [re-entry-tech](../entries/re-entry-tech.yml) · re-entry                | re- + entry。re- 表示再次；re-entry 指再次进入，具体是否允许或安全要看程序状态与调用约定。                                     | [S016](https://www.merriam-webster.com/dictionary/reentry)。按来源与当前义项整理；已写入词条。                                                                                                                                              |
| WF14 | [uninitialized-tech](../entries/uninitialized-tech.yml) · uninitialized | un- + initialized。un- 在这里否定已初始化的状态；uninitialized 表示尚未初始化，不等于已经初始化为 null。                       | [S017](https://www.merriam-webster.com/dictionary/un-)。现代否定状态分析，不补写语言未规定的初始值。                                                                                                                                        |
| WF15 | [unclaimed-workflow](../entries/unclaimed-workflow.yml) · unclaimed     | un- + claimed。un- 在这里表示尚未；unclaimed 在当前工作流语境中指任务还没有被认领。                                            | [S018](https://www.merriam-webster.com/dictionary/unclaimed)、[S017](https://www.merriam-webster.com/dictionary/un-)。按来源与当前义项整理；已写入词条。                                                                                    |
| WF16 | [unmount-tech](../entries/unmount-tech.yml) · unmount                   | un- + mount。un- 在这里表示撤销原动作；组件 unmount 是从渲染环境中移除，不是卸载软件包。                                       | [S017](https://www.merriam-webster.com/dictionary/un-)。与既有 M06 分别说明构成和语义对照，不改 M06。                                                                                                                                       |
| WF17 | [unregister-tech](../entries/unregister-tech.yml) · unregister          | un- + register。un- 在这里表示撤销原动作；unregister 是取消已有注册，例如移除处理器的注册。                                    | [S017](https://www.merriam-webster.com/dictionary/un-)。与既有 M08 分别说明构成和动作方向，不改 M08。                                                                                                                                       |
| WF18 | [unstaged-git](../entries/unstaged-git.yml) · unstaged                  | un- + staged。un- 在这里否定已暂存状态；unstaged 说的是更改尚未暂存，不等于文件未被 Git 跟踪。                                 | [S017](https://www.merriam-webster.com/dictionary/un-)、[S019](https://git-scm.com/book/en/v2/Git-Basics-Recording-Changes-to-the-Repository)。保留 staged/unstaged 与 tracked/untracked 两个维度。                                         |
| WF19 | [untracked-git](../entries/untracked-git.yml) · untracked               | un- + tracked。un- 在这里否定已跟踪状态；untracked 指 Git 尚未跟踪的文件，不能与已有文件的未暂存更改混为一谈。                 | [S017](https://www.merriam-webster.com/dictionary/un-)、[S019](https://git-scm.com/book/en/v2/Git-Basics-Recording-Changes-to-the-Repository)。不采用 Merriam-Webster untrack 的体育用法；Git 义以 Pro Git 为据。                           |
| WF20 | [inaccessible-tech](../entries/inaccessible-tech.yml) · inaccessible    | in- + accessible。in- 在这里表示否定；inaccessible 指无法访问，原因可能来自权限、位置或可达性。                                | [S020](https://www.merriam-webster.com/dictionary/inaccessible)。按来源与当前义项整理；已写入词条。                                                                                                                                         |
| WF21 | [nonthrowing-tech](../entries/nonthrowing-tech.yml) · nonthrowing       | non- + throwing。non- 表示不具有后面的性质；nonthrowing 在当前技术语境中指不抛出异常，不代表没有副作用。                       | [S021](https://www.merriam-webster.com/dictionary/non-)。现代前缀分析，不将"不抛异常"扩大为"纯函数"。                                                                                                                                       |
| WF22 | [asynchronous-tech](../entries/asynchronous-tech.yml) · asynchronous    | a- + synchronous。a- 在这里表示否定；asynchronous 表示异步，不要求调用方按同步方式等待完成，不等于自动并行或完全没有顺序约束。 | [S022](https://www.merriam-webster.com/dictionary/a-)、[S023](https://www.merriam-webster.com/dictionary/asynchronous)、[S024](https://developer.mozilla.org/en-US/docs/Glossary/Callback_function)。现代构成分析；否定同步不等于保证并行。 |
| WF23 | [independently-tech](../entries/independently-tech.yml) · independently | independent + -ly。-ly 将形容词转为副词；independently 表示独立地执行或判断，不依赖所讨论的另一项。                            | [S025](https://www.merriam-webster.com/dictionary/independently)、[S007](https://stella.glasgow.ac.uk/grammar/unit-4-3-affixes.html)。按来源与当前义项整理；已写入词条。                                                                    |

### WF24–WF39: 英语复合与缩略

| 编号 | 目标词条                                                                       | 完整草案或留空建议                                                                                                                | 核对依据与限定                                                                                                  |
|:-----|:-------------------------------------------------------------------------------|:----------------------------------------------------------------------------------------------------------------------------------|:----------------------------------------------------------------------------------------------------------------|
| WF24 | [fallback-tech](../entries/fallback.yml) · fallback                            | fall + back → fallback。动词短语 fall back 表示退回或转而依靠；名词 fallback 是常规方案不可用时的后备方案。                       | [S026](https://www.merriam-webster.com/dictionary/fallback)。与 N03 去重，建议只在此处写入。                    |
| WF25 | [setup-tech](../entries/setup-tech.yml) · setup                                | set + up → setup。动词短语 set up 表示设置或准备；当前名词 setup 指配置、准备工作或安排。                                         | [S027](https://www.merriam-webster.com/dictionary/setup)。按来源与当前义项整理；已写入词条。                    |
| WF26 | [cleanup-tech](../entries/cleanup-tech.yml) · cleanup                          | clean + up → cleanup。动词短语 clean up 表示清理；名词 cleanup 指清理工作，例如释放资源或移除临时状态。                           | [S028](https://www.merriam-webster.com/dictionary/cleanup)。按来源与当前义项整理；已写入词条。                  |
| WF27 | [callback-tech](../entries/callback-tech.yml) · callback                       | call + back → callback。可以按"交给另一段代码回过头来调用"理解；回调既可能同步调用，也可能异步调用。                              | [S024](https://developer.mozilla.org/en-US/docs/Glossary/Callback_function)。按来源与当前义项整理；已写入词条。 |
| WF28 | [handoff-workflow](../entries/handoff-workflow.yml) · handoff                  | hand + off → handoff。动词短语 hand off 表示移交；名词 handoff 指任务、责任或控制权的交接。                                       | [S029](https://www.merriam-webster.com/dictionary/handoff)。按来源与当前义项整理；已写入词条。                  |
| WF29 | [tradeoff-general](../entries/tradeoff-general.yml) · tradeoff                 | trade + off → tradeoff。trade off 表示以一方面换取另一方面；名词 tradeoff 指相互制约条件下的取舍。                                | [S030](https://www.merriam-webster.com/dictionary/trade-off)。按来源与当前义项整理；已写入词条。                |
| WF30 | [lifecycle-tech](../entries/lifecycle-tech.yml) · lifecycle                    | life + cycle。life cycle 把对象从开始到结束的阶段组织成一个周期；软件中可指组件或资源的生命周期，不保证对象一定重复经历所有阶段。 | [S031](https://www.merriam-webster.com/dictionary/life%20cycle)。按来源与当前义项整理；已写入词条。             |
| WF31 | [lockfile-tech](../entries/lockfile-tech.yml) · lockfile                       | lock + file。lock 提供"锁定"的线索；依赖管理中的 lockfile 记录解析后的依赖信息，不能只按互斥锁文件理解。                          | [S032](https://docs.npmjs.com/cli/v11/configuring-npm/package-lock-json/)。按来源与当前义项整理；已写入词条。   |
| WF32 | [toolchain-tech](../entries/toolchain-tech.yml) · toolchain                    | tool + chain。chain 表示相互衔接的一组环节；toolchain 是配合使用的一组开发工具，不必是单一程序。                                  | [S033](https://rust-lang.github.io/rustup/concepts/toolchains.html)。按来源与当前义项整理；已写入词条。         |
| WF33 | [workload-tech](../entries/workload-tech.yml) · workload                       | work + load。load 表示承载的量；workload 指需要执行的工作量或系统承受的任务负载。                                                 | [S034](https://www.merriam-webster.com/dictionary/workload)。按来源与当前义项整理；已写入词条。                 |
| WF34 | [upstream-git](../entries/upstream.yml) · upstream                             | up + stream。借河流上游的方向理解来源端；Git 协作中的 upstream 可指更上游的项目或分支，不是必须使用的远程名称。                   | [S035](https://www.merriam-webster.com/dictionary/upstream)。按来源与当前义项整理；已写入词条。                 |
| WF35 | [backtracking-algorithm](../entries/backtracking-algorithm.yml) · backtracking | back + track + -ing。backtrack 是沿先前路径返回；算法中的 backtracking 会撤销当前选择，退回分支点再尝试其他选择。                 | [S036](https://www.merriam-webster.com/dictionary/backtrack)。按来源与当前义项整理；已写入词条。                |
| WF36 | [breadth-first-tech](../entries/breadth-first-tech.yml) · breadth-first        | breadth + first。breadth 表示广度，first 表示优先；breadth-first 强调先扩展同一层的节点，再进入更深一层。                         | [S037](https://xlinux.nist.gov/dads/HTML/breadthfirst.html)。按来源与当前义项整理；已写入词条。                 |
| WF37 | [quicksort-algorithm](../entries/quicksort-algorithm.yml) · quicksort          | quick + sort。sort 表示排序，quick 是名称中的速度线索；quicksort 的核心是分区后继续处理子区间，不保证所有输入下都最快。           | [S038](https://xlinux.nist.gov/dads/HTML/quicksort.html)。按来源与当前义项整理；已写入词条。                    |
| WF38 | [symlink-tech](../entries/symlink-tech.yml) · symlink                          | symlink 是 symbolic link 的缩略形式；它指向另一个路径，不是把目标文件内容复制一份。                                               | [S039](https://man7.org/linux/man-pages/man7/symlink.7.html)。按来源与当前义项整理；已写入词条。                |
| WF39 | [typo-general](../entries/typo-general.yml) · typo                             | typo 是 typographical error 的缩略形式，指输入或排印造成的错误；不是所有代码缺陷都叫 typo。                                       | [S040](https://www.merriam-webster.com/dictionary/typo)。按来源与当前义项整理；已写入词条。                     |

### WF40–WF43: 名称中的可识别成分

| 编号 | 目标词条                                             | 完整草案或留空建议                                                                                                                 | 核对依据与限定                                                                                                                                             |
|:-----|:-----------------------------------------------------|:-----------------------------------------------------------------------------------------------------------------------------------|:-----------------------------------------------------------------------------------------------------------------------------------------------------------|
| WF40 | [copilot-ai](../entries/copilot-ai.yml) · Copilot    | 普通英语 copilot 可看作 co- + pilot，指协助主驾驶的副驾驶；这里只帮助识别 Copilot 名称中的词义，不据此断言产品命名史。             | [S041](https://www.merriam-webster.com/dictionary/copilot)。普通词义线索，不查认或编写品牌命名史。                                                         |
| WF41 | [deepseek-ai](../entries/deepseek-ai.yml) · DeepSeek | DeepSeek 的拼写中可识别 deep (深的、深入的) 与 seek (寻找、寻求)。这里只提供普通词义线索，不将"深入寻找"当作已核实的品牌命名来历。 | [S042](https://www.merriam-webster.com/dictionary/deep)、[S043](https://www.merriam-webster.com/dictionary/seek)。普通词义线索，不将直译当作官方命名说明。 |
| WF42 | [dots-ai](../entries/dots-ai.yml) · dots             | dots 是普通名词 dot (点) 的复数形式。这里只帮助记住产品名称的拼写，不推断这些"点"代表什么或官方为何这样命名。                      | [S044](https://www.merriam-webster.com/dictionary/dot)。只分析复数形式，不推断产品隐喻。                                                                   |
| WF43 | [openclaw-ai](../entries/openclaw-ai.yml) · OpenClaw | OpenClaw 的拼写中可识别 open (开放的、打开) 与 claw (爪)。这里只分析可识别的英语成分，不据此推断项目能力、标志含义或命名史。       | [S045](https://www.merriam-webster.com/dictionary/open)、[S046](https://www.merriam-webster.com/dictionary/claw)。普通词义线索，不以拆词推断产品能力。     |

### WF44–WF61: 日语复合与缩略

| 编号 | 目标词条                                                           | 完整草案或留空建议                                                                                                               | 核对依据与限定                                                                                                                       |
|:-----|:-------------------------------------------------------------------|:---------------------------------------------------------------------------------------------------------------------------------|:-------------------------------------------------------------------------------------------------------------------------------------|
| WF44 | [anime](../entries/anime.yml) · アニメ                             | アニメ 是 アニメーション 的缩略形式；保留的是前半部分，不是每个音节都要对应一个独立词根。                                        | [S047](https://kotobank.jp/word/あにめ-3206684)。按来源与当前义项整理；已写入词条。                                                  |
| WF45 | [anime-guzzu](../entries/anime-guzzu.yml) · アニメグッズ           | アニメ + グッズ。前项限定商品所属的领域，整体指动漫相关商品或周边，不是另一种动画形式。                                          | [S047](https://kotobank.jp/word/あにめ-3206684)、[S048](https://kotobank.jp/word/ぐつず-3210057)。按来源与当前义项整理；已写入词条。 |
| WF46 | [burajiru-eaa](../entries/burajiru-eaa.yml) · ブラジルエアー       | ブラジル + エアー。可识别为"巴西"与源自 air 的成分；这里只分析名称中的普通成分，不确认具体航空公司或命名来历。                   | [S049](https://kotobank.jp/word/えあ-1509216)。名称成分分析；库内已有说明保留，未确认机构身份。                                      |
| WF47 | [chuugokujin](../entries/chuugokujin.yml) · 中国人                 | 中国 + 人 (じん)。国名后加 人 可表示该国的人；此处读 ちゅうごくじん，不把 じん 推广为所有 人 字的读法。                          | [S050](https://kotobank.jp/word/人-536577)。人 的具体读音受组合限制，不机械类推。                                                    |
| WF48 | [denwa-bangou](../entries/denwa-bangou.yml) · 電話番号             | 電話 + 番号。整体是电话所用的号码；按两个完整词理解，不需要把每个汉字另拆成词。                                                  | [S051](https://kotobank.jp/word/電話)。依据已确认完整词形与 電話 的普通词义分析，不作历史词源主张。                                  |
| WF49 | [fisshu-baagaa](../entries/fisshu-baagaa.yml) · フィッシュバーガー | フィッシュ + バーガー。前项说明汉堡的鱼肉馅料，整体指鱼肉汉堡。                                                                  | [S052](https://kotobank.jp/word/ふいつしゆ-3185695)。依据已确认完整词形与 フィッシュ 的普通词义分析，不额外推导借入年代。            |
| WF50 | [gaikokujin](../entries/gaikokujin.yml) · 外国人                   | 外国 + 人 (じん)。外国 表示别的国家，整体指外国人；日常理解需以说话语境中的国家为参照。                                          | [S053](https://kotobank.jp/word/外国人)、[S050](https://kotobank.jp/word/人-536577)。按来源与当前义项整理；已写入词条。              |
| WF51 | [gentei-guzzu](../entries/gentei-guzzu.yml) · 限定グッズ           | 限定 + グッズ。限定 修饰商品的供应范围或条件；具体是数量、时间还是活动限定，要看实际说明。                                       | [S054](https://kotobank.jp/word/限定)、[S048](https://kotobank.jp/word/ぐつず-3210057)。按来源与当前义项整理；已写入词条。           |
| WF52 | [kaishain](../entries/kaishain.yml) · 会社員                       | 会社 + 員 (いん)。員 表示某组织中的成员；会社員 指受雇于公司工作的职员，不特指管理层。                                           | [S055](https://kotobank.jp/word/会社員)、[S056](https://kotobank.jp/word/員)。按来源与当前义项整理；已写入词条。                     |
| WF53 | [kenkyuusha](../entries/kenkyuusha.yml) · 研究者                   | 研究 + 者 (しゃ)。者 表示从事某种活动的人；研究者 指从事研究的人。                                                               | [S057](https://kotobank.jp/word/研究)、[S058](https://kotobank.jp/word/者)。按来源与当前义项整理；已写入词条。                       |
| WF54 | [koneko](../entries/koneko.yml) · 子猫                             | 子 (こ) + 猫 (ねこ)。子 在这里表示年幼，合起来读 こねこ，指幼猫。                                                                | [S059](https://kotobank.jp/word/子)、[S060](https://kotobank.jp/word/猫)。现代词形分析；子 在此表示幼小，不涉及姓名读音。            |
| WF55 | [neko-chan](../entries/neko-chan.yml) · 猫ちゃん                   | 猫 + ちゃん。ちゃん 是表示亲昵的称呼后缀；猫ちゃん 带亲近语气，不表示小猫的年龄。                                                | [S061](https://kokugo.jitenon.jp/word/p32740)。按来源与当前义项整理；已写入词条。                                                    |
| WF56 | [nihongo](../entries/nihongo.yml) · 日本語                         | 日本 + 語 (ご)。語 在这里表示语言；日本語 读 にほんご，指日语。                                                                  | [S062](https://kotobank.jp/word/日本語)。按来源与当前义项整理；已写入词条。                                                          |
| WF57 | [nihonjin](../entries/nihonjin.yml) · 日本人                       | 日本 + 人 (じん)。整体表示日本人；当前词条读 にほんじん，不能把 人 一律读作 ひと。                                               | [S063](https://kotobank.jp/word/日本人)、[S050](https://kotobank.jp/word/人-536577)。按来源与当前义项整理；已写入词条。              |
| WF58 | [pawaa-denki](../entries/pawaa-denki.yml) · パワー電気             | パワー + 電気。パワー 是源自 power 的外来语，電気 是普通日语词；这里只辨认名称中的成分，不据此确认名称含义、组织身份或业务范围。 | [S064](https://kotobank.jp/word/ぱわー-3165055)、[S065](https://kotobank.jp/word/電気)。名称成分分析；未确认命名者意图或机构业务。   |
| WF59 | [sakura-no-ki](../entries/sakura-no-ki.yml) · 桜の木               | 桜 + の + 木。の 连接两个名词，后项 木 是中心词；整体是樱花树，不是单指樱花花朵。                                                | [S066](https://www.coelang.tufs.ac.jp/mt/ja/gmod/courses/c01/lesson03/step1/card/008.html)。按来源与当前义项整理；已写入词条。       |
| WF60 | [shinsekai](../entries/shinsekai.yml) · 新世界                     | 新 (しん) + 世界 (せかい)。整体读 しんせかい，意为新的世界；具体指什么世界由语境决定。                                           | [S067](https://kotobank.jp/word/新)、[S068](https://kotobank.jp/word/世界)。现代词形组合分析，不解释作品中的唯一世界观。             |
| WF61 | [yomikata](../entries/yomikata.yml) · 読み方                       | 読む → 読み + 方 (かた)。読み 是 読む 的连用形；方 在这里表示方式，整体指读法或阅读方式。                                        | [S069](https://kotobank.jp/word/読み方)、[S070](https://kotobank.jp/word/方)。按来源与当前义项整理；已写入词条。                     |

### WF62–WF84: 日语屈折与语法组合

| 编号 | 目标词条                                                                                                               | 完整草案或留空建议                                                                                                                                | 核对依据与限定                                                                                                                                                                                                                                                                                                   |
|:-----|:-----------------------------------------------------------------------------------------------------------------------|:--------------------------------------------------------------------------------------------------------------------------------------------------|:-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| WF62 | [hanasemasen](../entries/hanasemasen.yml) · 話せません                                                                 | 話す → 話せる → 話せません。先变成可能形，再用礼貌否定；表达不会说或不能说，区别于单纯的 話しません (不说)。                                      | [S071](https://www.coelang.tufs.ac.jp/mt/ja/gmod/courses/c01/lesson18/step3/card/074.html)、[S072](https://www.coelang.tufs.ac.jp/mt/ja/gmod/courses/c01/lesson07/step1/card/023.html)。按来源与当前义项整理；已写入词条。                                                                                       |
| WF63 | [sukueru](../entries/sukueru.yml) · 救える                                                                             | 救う → 救える。五段动词的词尾 う 在可能形中变为 える，表示能够救助。                                                                              | [S071](https://www.coelang.tufs.ac.jp/mt/ja/gmod/courses/c01/lesson18/step3/card/074.html)。按来源与当前义项整理；已写入词条。                                                                                                                                                                                   |
| WF64 | [yareru](../entries/yareru.yml) · やれる                                                                               | やる → やれる。五段动词的词尾 る 在可能形中变为 れる；这里是能够做，不是另一个独立词根。                                                          | [S071](https://www.coelang.tufs.ac.jp/mt/ja/gmod/courses/c01/lesson18/step3/card/074.html)。按来源与当前义项整理；已写入词条。                                                                                                                                                                                   |
| WF65 | [shinde-miru](../entries/shinde-miru.yml) · 死んでみる                                                                 | 死ぬ → 死んで + みる。死んで 是て形，后接 みる 表示试着做；这是所给作品表达的结构分析，语气结合台词理解。                                         | [S073](https://www.coelang.tufs.ac.jp/mt/ja/gmod/contents/explanation/088.html)。作品文字沿用现有条目，本次未重核原作音轨。                                                                                                                                                                                      |
| WF66 | [yareru-koto](../entries/yareru-koto.yml) · やれること                                                                 | やれる + こと。やれる 是 やる 的可能形；こと 将前面的内容名词化，整体指能做的事。                                                                 | [S071](https://www.coelang.tufs.ac.jp/mt/ja/gmod/courses/c01/lesson18/step3/card/074.html)、[S074](https://www.coelang.tufs.ac.jp/mt/ja/gmod/courses/c01/lesson20/step3/card/080.html)。按来源与当前义项整理；已写入词条。                                                                                       |
| WF67 | [yaritai-dake-da-yo](../entries/yaritai-dake-da-yo.yml) · やりたいだけだよ                                             | やる → やり + たい + だけ + だ + よ。やりたい 表示想做，だけ 限定为"只是"；整体强调只是想做而已，よ 补充告知或强调语气。                          | [S075](https://www.coelang.tufs.ac.jp/mt/ja/gmod/courses/c01/lesson17/step1/card/067.html)。按来源与当前义项整理；已写入词条。                                                                                                                                                                                   |
| WF68 | [ni-nareru-grammar](../entries/ni-nareru-grammar.yml) · ～になれる                                                     | 名词或な形容词词干 + に + なれる。なれる 是 なる 的可能形；与单纯的 になる (成为) 区别在于"能够成为"。                                            | [S071](https://www.coelang.tufs.ac.jp/mt/ja/gmod/courses/c01/lesson18/step3/card/074.html)。なれる 在此是 なる 的可能形，不与 慣れる 混同。                                                                                                                                                                      |
| WF69 | [to-moushimasu-grammar](../entries/to-moushimasu-grammar.yml) · ～と申します                                           | 姓名 + と + 申します。と 引出所说的名字，申します 是 申す 的礼貌形式；整体用于郑重地自我介绍。                                                    | [S076](https://kotobank.jp/word/申す)。自称姓名的郑重用法，不把 申す 用作尊称他人的动作。                                                                                                                                                                                                                        |
| WF70 | [to-nanoru-grammar](../entries/to-nanoru-grammar.yml) · ～と名乗る                                                     | 名称或身份 + と + 名乗る。と 引出报出的名称或身份，名乗る 表示报上或自称；不单凭这个结构判断所称身份是否属实。                                    | [S077](https://kotobank.jp/word/名乗る)。按来源与当前义项整理；已写入词条。                                                                                                                                                                                                                                      |
| WF71 | [gimu-toka-janai-yo](../entries/gimu-toka-janai-yo.yml) · 義務とかじゃないよ                                           | 義務 + とか + じゃない + よ。とか 在这里将"义务之类"概括提出，じゃない 是 ではない 的口语形式；整体是否定，并非在列举两项义务。                   | [S079](https://www.irodori.jpf.go.jp/assets/data/starter/pdf/X_L03.pdf)。否定形式据《いろどり》；とか 的当前用法结合已确认词条语境分析，不展开所有语法义项。                                                                                                                                                     |
| WF72 | [ore-wa-ningen-dakara](../entries/ore-wa-ningen-dakara.yml) · 俺は人間だから                                           | 俺 + は + 人間 + だから。は 提出话题，だから 连接"是人类"这一理由；这里的理由在前，不把 だから 一律译成句首的"所以"。                             | [S078](https://www.coelang.tufs.ac.jp/mt/ja/gmod/courses/c01/lesson21/step1/card/081.html)。だから 的接续另据理由表达；具体作品语气以现有语境为限。                                                                                                                                                              |
| WF73 | [hikari-ni-nareru-nda](../entries/hikari-ni-nareru-nda.yml) · 光になれるんだ                                           | 光 + に + なれる + んだ。に 指变化的结果，なれる 表示能够成为，んだ 为解释或强调语气；保留所给作品的比喻语境。                                    | [S071](https://www.coelang.tufs.ac.jp/mt/ja/gmod/courses/c01/lesson18/step3/card/074.html)。现代结构分析；んだ 的解释语气限于当前句，不推测剧情。                                                                                                                                                                |
| WF74 | [jibun-jishin-no-chikara-de](../entries/jibun-jishin-no-chikara-de.yml) · 自分自身の力で                               | 自分自身 + の + 力 + で。自分自身 强调自己，の 连接所属，で 表示凭借的手段；整体是凭自己的力量。                                                  | [S066](https://www.coelang.tufs.ac.jp/mt/ja/gmod/courses/c01/lesson03/step1/card/008.html)。按来源与当前义项整理；已写入词条。                                                                                                                                                                                   |
| WF75 | [kotoba-wa-magenee](../entries/kotoba-wa-magenee.yml) · 言葉は曲げねェ                                                 | 言葉 + は + 曲げねェ。曲げねェ 对应 曲げない，是所给台词中的粗犷口语写法；这里用"不弯曲说过的话"表达不违背承诺，不把 ねェ 当作通用书面否定。      | 按现有 Notes 已确认的 曲げねェ → 曲げない 对应分析；本次未重核原作音轨，不建立普遍音变规则。                                                                                                                                                                                                                     |
| WF76 | [zettee-shinanee](../entries/zettee-shinanee.yml) · ぜってェ死なねー                                                   | ぜってェ + 死なねー。当前台词中分别对应 絶対 与 死なない；死なない 来自 死ぬ 的否定形。保留角色化拼写，不照搬到正式表达。                         | 按现有 Notes 已确认的角色口语对应分析；不假定所有日语说话者都这样发音。                                                                                                                                                                                                                                          |
| WF77 | [hokage-ni-naru-made](../entries/hokage-ni-naru-made.yml) · 火影になるまで                                             | 火影 + に + なる + まで。に 标记成为的身份，まで 表示持续到该事件；整体是直到成为火影为止。                                                       | [S080](https://www.coelang.tufs.ac.jp/mt/ja/gmod/courses/c01/lesson20/step2/card/079.html)。现代句法分析，まで 的终点是成为火影这一事件，不新增词条。                                                                                                                                                            |
| WF78 | [tantei-sa](../entries/tantei-sa.yml) · 探偵さ！                                                                       | 探偵 + さ。さ 是句尾助词，给"是侦探"加上断定或强调的口语语气；不是形容词后用于构成名词的 -さ。                                                    | [S081](https://kotobank.jp/word/さ-506834)。按来源与当前义项整理；已写入词条。                                                                                                                                                                                                                                   |
| WF79 | [kimi-ga-yobu-namae-yume-no-tsuzuki](../entries/kimi-ga-yobu-namae-yume-no-tsuzuki.yml) · きみが呼ぶなまえ～夢のつづき | きみが呼ぶ + なまえ；夢 + の + つづき。前半是"你呼唤的名字"，后半是"梦的延续"；きみが呼ぶ 修饰 なまえ，标题中的 ～ 分隔两部分，不表示"从……到……"。 | [S080](https://www.coelang.tufs.ac.jp/mt/ja/gmod/courses/c01/lesson20/step2/card/079.html)、[S066](https://www.coelang.tufs.ac.jp/mt/ja/gmod/courses/c01/lesson03/step1/card/008.html)、[S082](https://www.aniplex.co.jp/natsume/cd/02.html)。官方来源核对标题；修饰关系为语法分析，不声称官方解释了标题的寓意。 |
| WF80 | [boku-wa-shinsekai-no-kami-ni-naru](../entries/boku-wa-shinsekai-no-kami-ni-naru.yml) · 僕は新世界の神になる           | 僕 + は + 新世界 + の + 神 + に + なる。新世界の神 是"新世界的神"，に 标记成为的身份；整体是"我要成为新世界的神"。                                | [S066](https://www.coelang.tufs.ac.jp/mt/ja/gmod/courses/c01/lesson03/step1/card/008.html)。现代句法分析，不将内容扩写为作品设定。                                                                                                                                                                               |
| WF81 | [tsuki-ni-kawatte-oshioki-yo](../entries/tsuki-ni-kawatte-oshioki-yo.yml) · 月にかわっておしおきよ！                   | 月 + にかわって + おしおき + よ。にかわって 表示代表或代替，おしおき 是惩罚，よ 加上告知或强调语气；不是"变成月亮"。                              | [S083](https://www2.ninjal.ac.jp/verbhandbook/headwords/変わる・替わる・換わる・代わる.html)。代表义据动词手册，台词文字与作品归属沿用库内记录。                                                                                                                                                                 |
| WF82 | [fuuin-kaijo-reriizu](../entries/fuuin-kaijo-reriizu.yml) · 封印解除（レリーズ）！                                     | 封印 + 解除。普通汉字词组合读 ふういんかいじょ；当前咒语采用作品读法 レリーズ，不能从两个普通词的读音直接拼出这一作品读法。                       | 拆分普通汉字词；レリーズ 的限定来自已确认词条，不宣称听验原作。                                                                                                                                                                                                                                                  |
| WF83 | [omachigae](../entries/omachigae.yml) · お間違え                                                                       | お + 間違え。間違え 是 間違える 的连用形，お 加在其前，用于 お間違えではないでしょうか 等礼貌结构；不要与名词 間違い 的词形混同。                 | [S084](https://www.coelang.tufs.ac.jp/mt/ja/gmod/courses/c01/lesson24/step4/card/092.html)。这是礼貌结构中的连用形成分，不是可任意替换的独立动词。                                                                                                                                                               |
| WF84 | [osagashi](../entries/osagashi.yml) · お探し                                                                           | お + 探し。探し 是 探す 的连用形，お 加在其前，用于 お探しですか 等礼貌结构；需要结合完整句式理解。                                               | [S084](https://www.coelang.tufs.ac.jp/mt/ja/gmod/courses/c01/lesson24/step4/card/092.html)。结合 お探しですか 等完整结构理解，不把它当作能独立替换所有 探す 的动词。                                                                                                                                             |

### ETY01–ETY18: 历史来源与借入关系

| 编号  | 目标词条                                                                | 完整草案或留空建议                                                                                                                                                                                                           | 核对依据与限定                                                                                                                                                                                                               |
|:------|:------------------------------------------------------------------------|:-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|:-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| ETY01 | [escape-tech](../entries/escape-tech.yml) · escape                      | 经由盎格鲁法语 escaper、eschaper 进入中古英语，可追溯到通俗拉丁语的重建形式 *excappare，由 ex- (向外) 与晚期拉丁语 cappa (头巾、斗篷) 构成。"从斗篷中脱身"有助于理解逃离的历史意象；现代编程中的转义规则仍须按具体语法学习。 | [S085](https://www.merriam-webster.com/dictionary/escape)。* 标明重建形式；技术转义的含义不由斗篷故事推出。                                                                                                                  |
| ETY02 | [escape-parity-tech](../entries/escape-parity-tech.yml) · escape parity | parity 源自拉丁语 paritas，其基础 par 表示相等。现代英语也用 parity 表示奇偶性；当前 escape parity 指反斜杠数量的奇偶，而不是两段代码的功能对等，不能仅凭历史原义推导技术含义。                                              | [S086](https://www.merriam-webster.com/dictionary/parity)。当前词条是奇偶义；历史来源与现代技术定义分别处理。                                                                                                                |
| ETY03 | [cache-tech](../entries/cache-tech.yml) · cache                         | 来自北美法语 cache，意为藏物处；相关法语动词 cacher 表示隐藏。这个历史上的存放意象可以辅助记忆，但计算机缓存是否有效仍取决于缓存策略，不是只要存起来就能一直复用。                                                           | [S087](https://www.merriam-webster.com/dictionary/cache)。仅采用可确认的法语近源，不展开未证实的远源细节。                                                                                                                   |
| ETY04 | [candidate-tech](../entries/candidate-tech.yml) · candidate             | 来自拉丁语 candidatus，原指穿白衣者，也用于指竞逐公职的人；它与 candidus (明亮的、白色的) 有关，背景是古罗马候选人穿白色托加袍的习俗。现代技术语境沿用"候选者"的意义，不含颜色要求。                                         | [S088](https://www.merriam-webster.com/dictionary/candidate)。按来源与当前义项整理；已写入词条。                                                                                                                             |
| ETY05 | [etymology-language](../entries/etymology-language.yml) · etymology     | 经由法语和拉丁语进入英语，最终来自希腊语 etymon (词按其来源的本义) 与 -logia (论述、研究)。词源研究追踪词的历史；历史本义不等于现代语境中唯一正确的含义。                                                                    | [S089](https://www.merriam-webster.com/dictionary/etymology)。按来源与当前义项整理；已写入词条。                                                                                                                             |
| ETY06 | [mnemonic-memory](../entries/mnemonic-memory.yml) · mnemonic            | 来自希腊语 mnemonikos (与记忆有关的)，关联 mnemon (记得的、留心的) 与表示记忆的动词 mimneskesthai。这解释了它与记忆的联系，不必把现代拼写机械拆成几个英语词。                                                                | [S090](https://www.merriam-webster.com/dictionary/mnemonic)。按来源与当前义项整理；已写入词条。                                                                                                                              |
| ETY07 | [vocabulary-general](../entries/vocabulary.yml) · vocabulary            | 来自中古法语 vocabulaire，可能经由中世纪拉丁语 vocabularium，最终关联拉丁语 vocabulum (词、名称)。借入路径中的"可能"需保留；现代 vocabulary 通常指词汇或词汇量，不是一个单词。                                               | [S091](https://www.merriam-webster.com/dictionary/vocabulary)。保留词典的 probably，不将路径写成完全确定。                                                                                                                   |
| ETY08 | [syntax-tech](../entries/syntax-tech.yml) · syntax                      | 经由中古法语或晚期拉丁语 syntaxis，最终关联希腊语 syntassein (安排在一起)，其中 syn- 表示共同，tassein 表示安排。可据此联想词语或代码成分的组合规则，不把 syntax 中的 tax 当作现代英语的税。                                 | [S092](https://www.merriam-webster.com/dictionary/syntax)。按来源与当前义项整理；已写入词条。                                                                                                                                |
| ETY09 | [artifact-tech](../entries/artifact.yml) · artifact                     | 来自新拉丁语 arte factus (经人为制作的)，其中 arte 源自表示技艺的 ars，factus 是 facere (制作、做) 的过去分词。技术中的 artifact 延伸为生成的产物；不要拆成现代英语的 art 加 fact。                                          | [S093](https://www.merriam-webster.com/dictionary/artifact)。按来源与当前义项整理；已写入词条。                                                                                                                              |
| ETY10 | [idiomatic-tech](../entries/idiomatic.yml) · idiomatic                  | 最终来自希腊语 idiomatikos (特有的、有其特点的)，关联 idioma (特有的表达方式) 与 idios (自己的、特有的)。这能帮助理解"符合某种语言或生态惯用方式"，不表示只能由某个个人使用。                                                | [S094](https://www.oxfordlearnersdictionaries.com/definition/english/idiomatic)。改用 Oxford 的 idiomatic 本词 Word Origin，不只用 idiom 反推。                                                                              |
| ETY11 | [gochisou](../entries/gochisou.yml) · ごちそう                          | ごちそう 对应 御馳走，是在 馳走 前加 ご 的形式。馳走 原有奔走、周旋的意思，后用于用心招待，尤其是提供饭食；后来加 ご 的形式逐渐普遍。用这条语义联系理解款待，不另编宗教或人物故事。                                          | [S095](https://kotobank.jp/word/御馳走-502226)、[S096](https://kotobank.jp/word/馳走-565923)。不编造寺庙、人物或具体年代故事。                                                                                               |
| ETY12 | [chisou](../entries/chisou.yml) · 馳走                                  | 馳走 原可表示奔走，后有为他人张罗、照料的意思，再用于用心招待及所准备的食物。这里的历史联系是为款待而奔走，不是现代语境中见到 馳走 都译成跑步。                                                                              | [S096](https://kotobank.jp/word/馳走-565923)。不把几个历史义项写成可精确断代的单一路线。                                                                                                                                     |
| ETY13 | [aitemu](../entries/aitemu.yml) · アイテム                              | アイテム 是借自英语 item 的外来语。日语中可指物品、项目或游戏道具，具体义项取决于语境；不因是外来语就认为它与英语 item 的所有用法完全对应。                                                                                  | [S097](https://dictionary.sanseido-publ.co.jp/wordpress/wp-content/themes/dictionaryandbeyond/pdf/ssd11064.pdf)。词典证据限于检索摘要，试读 PDF 正文抓取返回 502；直接借词对应另获用户确认，不再阻塞实施，不声称正文已复核。 |
| ETY14 | [eaa](../entries/eaa.yml) · エアー                                      | エアー 是英语 air 的借入形式，也是 エア 的另一种表记。当前可结合空气、压缩空气和航空等语境理解，不把名称中出现 エアー 当作确认航空公司身份的证据。                                                                           | [S049](https://kotobank.jp/word/えあ-1509216)。按来源与当前义项整理；已写入词条。                                                                                                                                            |
| ETY15 | [fisshu](../entries/fisshu.yml) · フィッシュ                            | フィッシュ 借自英语 fish，表示鱼或鱼类，常作为复合词的构成成分，例如 フィッシュバーガー。这里说明的是借词来源，不据此推导所有英语发音到片假名的转换规则。                                                                    | [S052](https://kotobank.jp/word/ふいつしゆ-3185695)。按来源与当前义项整理；已写入词条。                                                                                                                                      |
| ETY16 | [guzzu](../entries/guzzu.yml) · グッズ                                  | グッズ 借自英语 goods，表示商品、物品；日语也常用它指与人物或活动相关的周边。不要把它反向理解为英语 good (好的) 的日语复数变化。                                                                                             | [S048](https://kotobank.jp/word/ぐつず-3210057)。按来源与当前义项整理；已写入词条。                                                                                                                                          |
| ETY17 | [pawaa](../entries/pawaa.yml) · パワー                                  | パワー 借自英语 power，可用于人的力量、机器动力等语境。可借英语词形帮助记忆，但不据此确认含有 パワー 的名称如何命名或经营什么业务。                                                                                          | [S064](https://kotobank.jp/word/ぱわー-3165055)。按来源与当前义项整理；已写入词条。                                                                                                                                          |
| ETY18 | [taoru](../entries/taoru.yml) · タオル                                  | タオル 借自英语 towel，指毛巾或毛巾织物。本项只记录日语的直接借入来源，不把相关欧洲语言中的远源混写成日语直接借自那些语言。                                                                                                  | [S098](https://kotobank.jp/word/たおる-3157938)。仅采用日语直接借入英语的记录；页面上的欧洲远源说法不混为同一层借入。                                                                                                        |

### N01–N23: 独立助记

| 编号 | 目标词条                                                                                                             | 完整草案或留空建议                                                                                                                                                               | 核对依据与限定                                                                                                                                                                                                                                                        |
|:-----|:---------------------------------------------------------------------------------------------------------------------|:---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|:----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| N01  | [invaluable-general](../entries/invaluable-general.yml) · invaluable                                                 | invaluable 记成"价值大到难以估量"，所以是极有价值；不要看见 in- 就记成 worthless (无价值的)。                                                                                    | [S099](https://www.merriam-webster.com/dictionary/invaluable)。按来源与当前义项整理；已写入词条。                                                                                                                                                                     |
| N02  | [deprecated-tech](../entries/deprecated.yml) · deprecated                                                            | deprecated 是"不再推荐使用"；可能仍然可用，不等于已经删除、不能运行或一定存在缺陷。                                                                                              | [S100](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Deprecated_and_obsolete_features)。按来源与当前义项整理；已写入词条。                                                                                                                        |
| N03  | [fallback-tech](../entries/fallback.yml) · fallback                                                                  | 建议留空，不写入 mnemonic。名词 fallback 连写、动词短语 fall back 分写的提示已纳入 WF24；保留本编号用于确认去重决定。                                                            | [S026](https://www.merriam-webster.com/dictionary/fallback)。已按决定留空；与 WF24 的说明完全重叠，无需两个区块同时显示。                                                                                                                                             |
| N04  | [pronunciation-general](../entries/pronunciation.yml) · pronunciation                                                | 动词 pronounce 中有 ou，名词 pronunciation 对应部分只有 u；标准拼写不是 pronounciation。                                                                                         | [S101](https://www.merriam-webster.com/dictionary/pronunciation)。按来源与当前义项整理；已写入词条。                                                                                                                                                                  |
| N05  | [mnemonic-memory](../entries/mnemonic-memory.yml) · mnemonic                                                         | mnemonic 的词首 m 不发音，发音从 n 开始；记住"写 m，读 n 开头"，不要按每个字母逐个拼读。                                                                                         | [S090](https://www.merriam-webster.com/dictionary/mnemonic)。仅提醒标准英语读音；历史来源另见 ETY06。                                                                                                                                                                 |
| N06  | [field-presence-tech](../entries/field-presence-tech.yml) · field presence                                           | 先问"字段在不在"，再问"值是什么"。即使值为 false、0、空字符串、null 或 undefined，字段仍可能存在；还要区分自身属性与继承属性。                                                   | [S102](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Object/hasOwn)。按来源与当前义项整理；已写入词条。                                                                                                                            |
| N07  | [escaped-backslash-tech](../entries/escaped-backslash-tech.yml) · escaped backslash                                  | 先分清源码与运行时值。普通 JavaScript 字符串字面量中，两个连续反斜杠构成转义序列，运行时表示一个反斜杠；多套一层字符串编码就多一层规则。                                         | [S103](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Lexical_grammar)。按来源与当前义项整理；已写入词条。                                                                                                                                         |
| N08  | [hito-o-norowaba-ana-futatsu](../entries/hito-o-norowaba-ana-futatsu.yml) · 人を呪わば穴二つ                         | 呪わば 是谚语保留的文语条件形式。现代条件可见 呪えば、呪ったら；不要把词尾换成 わば 当作现代动词的通用变形。                                                                     | [S104](https://k-rain.repo.nii.ac.jp/record/2001962/files/bunkenronsyu_052_005.pdf)。按来源与当前义项整理；已写入词条。                                                                                                                                               |
| N09  | [fuuin-kaijo-reriizu](../entries/fuuin-kaijo-reriizu.yml) · 封印解除（レリーズ）！                                   | 这张卡按《カードキャプターさくら》的作品读法读 レリーズ；普通汉字词 封印解除 则读 ふういんかいじょ。看见相同汉字时先判断是否处于作品语境。                                       | 依据现有已确认作品读法及 Notes；本次不重核原作音轨。                                                                                                                                                                                                                  |
| N10  | [kosumo-saint-seiya](../entries/kosumo-saint-seiya.yml) · 小宇宙                                                     | 这张卡的 小宇宙 来自《聖闘士星矢》，按作品读法读 コスモ；不要把这一读法推广到其他语境中的所有 小宇宙。                                                                           | 依据现有已确认作品读法及 Notes；不发明普通读法或扩写设定。                                                                                                                                                                                                            |
| N11  | [daijoubu](../entries/daijoubu.yml) · 大丈夫                                                                         | 大丈夫 的核心是"没问题"；回答邀约或询问是否需要某物时，也可能表示"不用了"。不要脱离问题、语气和动作，仅凭这个词判断答应还是拒绝。                                                | [S105](https://www.caa.go.jp/policies/policy/consumer_education/consumer_education/nurture/assets/caa_futurer101_250508_02.pdf)。按来源与当前义项整理；已写入词条。                                                                                                   |
| N12  | [onamae-wa](../entries/onamae-wa.yml) · お名前は                                                                     | お名前は 是省略的礼貌问句，可补想成"您的名字是……？"；は 提出话题，疑问来自完整语境与语调，不是 は 本身等于问号。                                                                 | [S079](https://www.irodori.jpf.go.jp/assets/data/starter/pdf/X_L03.pdf)。按来源与当前义项整理；已写入词条。                                                                                                                                                           |
| N13  | [san-name-suffix](../entries/san-name-suffix.yml) · ～さん                                                           | 姓名后的 さん 是敬称；普通自我介绍时不给自己的姓名加 さん。它与数字 三 的 さん 发音相同，但作用不同。                                                                            | [S106](https://kotobank.jp/word/さん-70293)、[S079](https://www.irodori.jpf.go.jp/assets/data/starter/pdf/X_L03.pdf)。指普通自我介绍惯例，不声称任何戏仿、转述中都不能出现自己的名字加敬称。                                                                          |
| N14  | [jin-nationality-suffix](../entries/jin-nationality-suffix.yml) · ～人                                               | 国籍表达中的 ～人 读 じん，如 日本人、中国人；人 单独作普通名词时常读 ひと。按具体词记读法，不把一个读音套到所有组合。                                                           | [S050](https://kotobank.jp/word/人-536577)。按来源与当前义项整理；已写入词条。                                                                                                                                                                                        |
| N15  | [sa-ending-grammar](../entries/sa-ending-grammar.yml) · ～さ                                                         | 探偵さ 中的 さ 在句尾补充语气；高さ 中的 さ 则构成表示程度的名词。同一个假名，不代表同一种语法作用。                                                                             | [S081](https://kotobank.jp/word/さ-506834)。按来源与当前义项整理；已写入词条。                                                                                                                                                                                        |
| N16  | [no-ka-grammar](../entries/no-ka-grammar.yml) · ～のか                                                               | のか 可以用来询问、确认，也能带反问语气；先看上下文和说话语气，不把每一个 のか 都记成强烈反问。                                                                                  | [S107](https://kotobank.jp/word/か-456254)。按来源与当前义项整理；已写入词条。                                                                                                                                                                                        |
| N17  | [promise-all-javascript](../entries/promise-all-javascript.yml) · Promise.all                                        | Promise.all 的结果按输入顺序排列，不按完成先后排列；其中一项拒绝会使聚合拒绝，但不会自动取消其他已经启动的任务。                                                                 | [S108](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Promise/all)。按来源与当前义项整理；已写入词条。                                                                                                                              |
| N18  | [block-formatting-context-css](../entries/block-formatting-context-css.yml) · block formatting context               | 把 BFC 记成容器与内容的布局边界: 它能包含内部浮动，阻止容器与其流内子元素的外边距折叠；内部相邻普通块之间仍可能折叠。                                                            | [S109](https://developer.mozilla.org/en-US/docs/Web/CSS/Guides/Display/Block_formatting_context)。[S115](https://developer.mozilla.org/en-US/docs/Web/CSS/Guides/Box_model/Margin_collapsing)。明确为容器与子元素的边界，不概括为所有内部或外部相邻块的外边距都独立。 |
| N19  | [floating-point-precision-javascript](../entries/floating-point-precision-javascript.yml) · floating-point precision | 有限的二进制位不一定能精确表示十进制小数。比较容差要结合数值量级和业务精度；Number.EPSILON 是 1 附近的间距，不是所有数的通用误差阈值。                                           | [S110](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Number/EPSILON)。按来源与当前义项整理；已写入词条。                                                                                                                           |
| N20  | [virtual-dom-frontend](../entries/virtual-dom-frontend.yml) · virtual DOM                                            | virtual DOM 是内存中的界面描述，更新后再同步真实 DOM；创建和比较这层描述也有成本，不能仅凭"virtual"断言它总比直接操作 DOM 更快。                                                 | [S111](https://vuejs.org/guide/extras/rendering-mechanism.html)。可从文档中的分配、遍历和 patch 推知有运行成本；不是性能基准结论。                                                                                                                                    |
| N21  | [http-caching-web](../entries/http-caching-web.yml) · HTTP caching                                                   | no-cache 可存储，但作为 HTTP 缓存复用前要重新验证；no-store 要求不要存储该响应。记"先验证"和"不存储"，不要把 no-cache 理解成完全不缓存，也不要把 no-store 当作删除旧缓存的指令。 | [S112](https://developer.mozilla.org/en-US/docs/Web/HTTP/Guides/Caching)。按来源与当前义项整理；已写入词条。                                                                                                                                                          |
| N22  | [least-recently-used-cache](../entries/least-recently-used-cache.yml) · least recently used                          | LRU 看"上次访问距今多久"，LFU 看"访问了多少次"，TTL 看"是否到期"；最近最少使用不是使用次数最少，具体系统还可能采用近似 LRU。                                                     | [S113](https://redis.io/docs/latest/develop/reference/eviction/)。按来源与当前义项整理；已写入词条。                                                                                                                                                                  |
| N23  | [preflight-request-web](../entries/preflight-request-web.yml) · preflight request                                    | 预检用 OPTIONS 先询问允许的方法和请求头，实际请求随后另发；携带凭据与触发预检是两项不同判断。简单跨源请求也可能携带凭据，不能把两者绑定。                                        | [S114](https://developer.mozilla.org/en-US/docs/Web/HTTP/Guides/CORS)。按来源与当前义项整理；已写入词条。                                                                                                                                                             |

### 来源清单与使用范围

以下为本次草案采用的词典、教材、标准相关文档和项目官方文档。URL 中的作品、代码或词形保持原样；来源只支持其列明的范围，不等于支持同一聚合页面中的所有义项。ETY13 的正文抓取限制见其条目。

- **S001** [Merriam-Webster Dictionary: actionable](https://www.merriam-webster.com/dictionary/actionable). 对应词义、词形或 Word History；现代构成分析不冒充词典明载的历史推导。
- **S002** [Merriam-Webster Dictionary: -or](https://www.merriam-webster.com/dictionary/-or). 对应词义、词形或 Word History；现代构成分析不冒充词典明载的历史推导。
- **S003** [MDN: get](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Functions/get). getter 与 accessor property 的程序语义。
- **S004** [Merriam-Webster Dictionary: adapter](https://www.merriam-webster.com/dictionary/adapter). 对应词义、词形或 Word History；现代构成分析不冒充词典明载的历史推导。
- **S005** [Merriam-Webster Dictionary: blocker](https://www.merriam-webster.com/dictionary/blocker). 对应词义、词形或 Word History；现代构成分析不冒充词典明载的历史推导。
- **S006** [Merriam-Webster Dictionary: correctness](https://www.merriam-webster.com/dictionary/correctness). 对应词义、词形或 Word History；现代构成分析不冒充词典明载的历史推导。
- **S007** [University of Glasgow, STELLA: Unit 4.3 Affixes](https://stella.glasgow.ac.uk/grammar/unit-4-3-affixes.html). 派生词缀、-ness、-ly 等词类变化的一般规则。
- **S008** [Merriam-Webster Dictionary: maintainable](https://www.merriam-webster.com/dictionary/maintainable). 对应词义、词形或 Word History；现代构成分析不冒充词典明载的历史推导。
- **S009** [Merriam-Webster Dictionary: misalign](https://www.merriam-webster.com/dictionary/misalign). 对应词义、词形或 Word History；现代构成分析不冒充词典明载的历史推导。
- **S010** [Merriam-Webster Dictionary: malformed](https://www.merriam-webster.com/dictionary/malformed). 对应词义、词形或 Word History；现代构成分析不冒充词典明载的历史推导。
- **S011** [Merriam-Webster Dictionary: mal-](https://www.merriam-webster.com/dictionary/mal-). 对应词义、词形或 Word History；现代构成分析不冒充词典明载的历史推导。
- **S012** [Merriam-Webster Dictionary: multi-](https://www.merriam-webster.com/dictionary/multi-). 对应词义、词形或 Word History；现代构成分析不冒充词典明载的历史推导。
- **S013** [MDN: Using microtasks in JavaScript with queueMicrotask()](https://developer.mozilla.org/en-US/docs/Web/API/HTML_DOM_API/Microtask_guide). 微任务调度及持续添加微任务的风险。
- **S014** [Merriam-Webster Dictionary: preexisting](https://www.merriam-webster.com/dictionary/preexisting). 对应词义、词形或 Word History；现代构成分析不冒充词典明载的历史推导。
- **S015** [Merriam-Webster Dictionary: prefilter](https://www.merriam-webster.com/dictionary/prefilter). 对应词义、词形或 Word History；现代构成分析不冒充词典明载的历史推导。
- **S016** [Merriam-Webster Dictionary: reentry](https://www.merriam-webster.com/dictionary/reentry). 对应词义、词形或 Word History；现代构成分析不冒充词典明载的历史推导。
- **S017** [Merriam-Webster Dictionary: un-](https://www.merriam-webster.com/dictionary/un-). 对应词义、词形或 Word History；现代构成分析不冒充词典明载的历史推导。
- **S018** [Merriam-Webster Dictionary: unclaimed](https://www.merriam-webster.com/dictionary/unclaimed). 对应词义、词形或 Word History；现代构成分析不冒充词典明载的历史推导。
- **S019** [Pro Git: Recording Changes to the Repository](https://git-scm.com/book/en/v2/Git-Basics-Recording-Changes-to-the-Repository). 区分 tracked、untracked、staged 与未暂存更改。
- **S020** [Merriam-Webster Dictionary: inaccessible](https://www.merriam-webster.com/dictionary/inaccessible). 对应词义、词形或 Word History；现代构成分析不冒充词典明载的历史推导。
- **S021** [Merriam-Webster Dictionary: non-](https://www.merriam-webster.com/dictionary/non-). 对应词义、词形或 Word History；现代构成分析不冒充词典明载的历史推导。
- **S022** [Merriam-Webster Dictionary: a-](https://www.merriam-webster.com/dictionary/a-). 对应词义、词形或 Word History；现代构成分析不冒充词典明载的历史推导。
- **S023** [Merriam-Webster Dictionary: asynchronous](https://www.merriam-webster.com/dictionary/asynchronous). 对应词义、词形或 Word History；现代构成分析不冒充词典明载的历史推导。
- **S024** [MDN: Callback function](https://developer.mozilla.org/en-US/docs/Glossary/Callback_function). 回调可能同步或异步。
- **S025** [Merriam-Webster Dictionary: independently](https://www.merriam-webster.com/dictionary/independently). 对应词义、词形或 Word History；现代构成分析不冒充词典明载的历史推导。
- **S026** [Merriam-Webster Dictionary: fallback](https://www.merriam-webster.com/dictionary/fallback). 对应词义、词形或 Word History；现代构成分析不冒充词典明载的历史推导。
- **S027** [Merriam-Webster Dictionary: setup](https://www.merriam-webster.com/dictionary/setup). 对应词义、词形或 Word History；现代构成分析不冒充词典明载的历史推导。
- **S028** [Merriam-Webster Dictionary: cleanup](https://www.merriam-webster.com/dictionary/cleanup). 对应词义、词形或 Word History；现代构成分析不冒充词典明载的历史推导。
- **S029** [Merriam-Webster Dictionary: handoff](https://www.merriam-webster.com/dictionary/handoff). 对应词义、词形或 Word History；现代构成分析不冒充词典明载的历史推导。
- **S030** [Merriam-Webster Dictionary: trade-off](https://www.merriam-webster.com/dictionary/trade-off). 对应词义、词形或 Word History；现代构成分析不冒充词典明载的历史推导。
- **S031** [Merriam-Webster Dictionary: life cycle](https://www.merriam-webster.com/dictionary/life%20cycle). 对应词义、词形或 Word History；现代构成分析不冒充词典明载的历史推导。
- **S032** [npm Docs: package-lock.json](https://docs.npmjs.com/cli/v11/configuring-npm/package-lock-json/). 依赖锁文件记录解析后的依赖树，不是操作系统互斥锁。
- **S033** [The rustup book: Toolchains](https://rust-lang.github.io/rustup/concepts/toolchains.html). 以 Rust 工具组说明 toolchain 的现代技术用法。
- **S034** [Merriam-Webster Dictionary: workload](https://www.merriam-webster.com/dictionary/workload). 对应词义、词形或 Word History；现代构成分析不冒充词典明载的历史推导。
- **S035** [Merriam-Webster Dictionary: upstream](https://www.merriam-webster.com/dictionary/upstream). 对应词义、词形或 Word History；现代构成分析不冒充词典明载的历史推导。
- **S036** [Merriam-Webster Dictionary: backtrack](https://www.merriam-webster.com/dictionary/backtrack). 对应词义、词形或 Word History；现代构成分析不冒充词典明载的历史推导。
- **S037** [NIST DADS: breadth-first search](https://xlinux.nist.gov/dads/HTML/breadthfirst.html). 先同层再下一层的广度优先搜索。
- **S038** [NIST DADS: quicksort](https://xlinux.nist.gov/dads/HTML/quicksort.html). 分区、递归及最坏情况复杂度，不保证所有输入都最快。
- **S039** [Linux manual pages: symlink(7)](https://man7.org/linux/man-pages/man7/symlink.7.html). symbolic link、symlink 的对应及路径链接语义。
- **S040** [Merriam-Webster Dictionary: typo](https://www.merriam-webster.com/dictionary/typo). 对应词义、词形或 Word History；现代构成分析不冒充词典明载的历史推导。
- **S041** [Merriam-Webster Dictionary: copilot](https://www.merriam-webster.com/dictionary/copilot). 对应词义、词形或 Word History；现代构成分析不冒充词典明载的历史推导。
- **S042** [Merriam-Webster Dictionary: deep](https://www.merriam-webster.com/dictionary/deep). 仅支持名称中可识别的普通英语成分，不证明品牌命名史。
- **S043** [Merriam-Webster Dictionary: seek](https://www.merriam-webster.com/dictionary/seek). 对应词义、词形或 Word History；现代构成分析不冒充词典明载的历史推导。
- **S044** [Merriam-Webster Dictionary: dot](https://www.merriam-webster.com/dictionary/dot). 对应词义、词形或 Word History；现代构成分析不冒充词典明载的历史推导。
- **S045** [Merriam-Webster Dictionary: open](https://www.merriam-webster.com/dictionary/open). 仅支持名称中可识别的普通英语成分，不证明品牌命名史。
- **S046** [Merriam-Webster Dictionary: claw](https://www.merriam-webster.com/dictionary/claw). 对应词义、词形或 Word History；现代构成分析不冒充词典明载的历史推导。
- **S047** [小学館《デジタル大辞泉》或《精選版 日本国語大辞典》: アニメ (コトバンク)](https://kotobank.jp/word/あにめ-3206684). アニメーション 的缩略。
- **S048** [小学館《デジタル大辞泉》或《精選版 日本国語大辞典》: グッズ (コトバンク)](https://kotobank.jp/word/ぐつず-3210057). 英语 goods 的借入，以及商品、周边的义项。
- **S049** [小学館《デジタル大辞泉》或《精選版 日本国語大辞典》: エア (コトバンク)](https://kotobank.jp/word/えあ-1509216). air 的借入、エアー 异表记、空气与航空义。
- **S050** [小学館《デジタル大辞泉》或《精選版 日本国語大辞典》: 人 (コトバンク)](https://kotobank.jp/word/人-536577). じん 的组合用法及 人 的多种读音。
- **S051** [小学館《デジタル大辞泉》或《精選版 日本国語大辞典》: 電話 (コトバンク)](https://kotobank.jp/word/電話). 電話 的普通词义；完整组合按库内已确认词形分析。
- **S052** [小学館《デジタル大辞泉》或《精選版 日本国語大辞典》: フィッシュ (コトバンク)](https://kotobank.jp/word/ふいつしゆ-3185695). fish 的借入，鱼类及复合词成分用法。
- **S053** [小学館《デジタル大辞泉》或《精選版 日本国語大辞典》: 外国人 (コトバンク)](https://kotobank.jp/word/外国人). 现代词义与构成，不展开法律身份判断。
- **S054** [小学館《デジタル大辞泉》或《精選版 日本国語大辞典》: 限定 (コトバンク)](https://kotobank.jp/word/限定). 限制范围或数量。
- **S055** [小学館《デジタル大辞泉》或《精選版 日本国語大辞典》: 会社員 (コトバンク)](https://kotobank.jp/word/会社員). 会社に雇われ働く人，词形为 かいしゃ‐いん。
- **S056** [小学館《デジタル大辞泉》或《精選版 日本国語大辞典》: 員 (コトバンク)](https://kotobank.jp/word/員). 成员、人员的构成成分。
- **S057** [小学館《デジタル大辞泉》或《精選版 日本国語大辞典》: 研究 (コトバンク)](https://kotobank.jp/word/研究). 普通词义"研究"。
- **S058** [小学館《デジタル大辞泉》或《精選版 日本国語大辞典》: 者 (コトバンク)](https://kotobank.jp/word/者). 用来指人的构成成分；不推导为所有 者 字的读法。
- **S059** [小学館《デジタル大辞泉》: 子 (コトバンク)](https://kotobank.jp/word/子). 普通词义及现代构成成分，不推导姓名读音。
- **S060** [小学館《デジタル大辞泉》: 猫 (コトバンク)](https://kotobank.jp/word/猫). 普通词义及现代构成成分，不推导姓名读音。
- **S061** [国語辞典オンライン: ちゃん](https://kokugo.jitenon.jp/word/p32740). 亲昵称呼后缀；只采用基本亲爱义，不扩大为年龄规则。
- **S062** [小学館《デジタル大辞泉》或《精選版 日本国語大辞典》: 日本語 (コトバンク)](https://kotobank.jp/word/日本語). 日本語 的词形与日语义项；不引用未定的语言谱系讨论。
- **S063** [小学館《デジタル大辞泉》或《精選版 日本国語大辞典》: 日本人 (コトバンク)](https://kotobank.jp/word/日本人). 仅采用日本国民义项；不采用该聚合页面的旧人类学分类。
- **S064** [小学館《デジタル大辞泉》或《精選版 日本国語大辞典》: パワー (コトバンク)](https://kotobank.jp/word/ぱわー-3165055). power 的借入，力量与动力的义项。
- **S065** [小学館《デジタル大辞泉》或《精選版 日本国語大辞典》: 電気 (コトバンク)](https://kotobank.jp/word/電気). 電気 的普通词义，不作为机构身份或命名来源证据。
- **S066** [東京外国語大学言語モジュール: NのN](https://www.coelang.tufs.ac.jp/mt/ja/gmod/courses/c01/lesson03/step1/card/008.html). 名词间 の 的修饰关系。
- **S067** [小学館《デジタル大辞泉》或《精選版 日本国語大辞典》: 新 (コトバンク)](https://kotobank.jp/word/新). しん 表示新的构成线索。
- **S068** [小学館《デジタル大辞泉》或《精選版 日本国語大辞典》: 世界 (コトバンク)](https://kotobank.jp/word/世界). 世界 的普通词义。
- **S069** [小学館《デジタル大辞泉》或《精選版 日本国語大辞典》: 読み方 (コトバンク)](https://kotobank.jp/word/読み方). 读法、阅读方式；配合 方 接尾用法作现代分析。
- **S070** [小学館《デジタル大辞泉》或《精選版 日本国語大辞典》: 方 (コトバンク)](https://kotobank.jp/word/方). 动词连用形后表示方式的接尾用法；访问时重定向至 方-44281。
- **S071** [東京外国語大学言語モジュール: 可能形](https://www.coelang.tufs.ac.jp/mt/ja/gmod/courses/c01/lesson18/step3/card/074.html). 五段动词可能形和礼貌否定；具体词形按该规则分析。
- **S072** [東京外国語大学言語モジュール: Vます／Vません](https://www.coelang.tufs.ac.jp/mt/ja/gmod/courses/c01/lesson07/step1/card/023.html). 动词礼貌否定，配合可能形。
- **S073** [東京外国語大学言語モジュール: Vておきます／Vてみます: 解説](https://www.coelang.tufs.ac.jp/mt/ja/gmod/contents/explanation/088.html). て形后接 みる 的试做含义。
- **S074** [東京外国語大学言語モジュール: Vこと・Vの](https://www.coelang.tufs.ac.jp/mt/ja/gmod/courses/c01/lesson20/step3/card/080.html). こと、の 的名词化。
- **S075** [東京外国語大学言語モジュール: Vたいです／Vたくありません](https://www.coelang.tufs.ac.jp/mt/ja/gmod/courses/c01/lesson17/step1/card/067.html). 愿望表达；具体接续结合连用形分析。
- **S076** [小学館《デジタル大辞泉》或《精選版 日本国語大辞典》: 申す (コトバンク)](https://kotobank.jp/word/申す). 言う 的郑重表达，含自称姓名例句。
- **S077** [小学館《デジタル大辞泉》或《精選版 日本国語大辞典》: 名乗る (コトバンク)](https://kotobank.jp/word/名乗る). 报出姓名或身份。
- **S078** [東京外国語大学言語モジュール: 理由](https://www.coelang.tufs.ac.jp/mt/ja/gmod/courses/c01/lesson21/step1/card/081.html). 名词谓语后的 だから 等理由表达。
- **S079** [国際交流基金《いろどり》入門 第3課: よろしくお願いします](https://www.irodori.jpf.go.jp/assets/data/starter/pdf/X_L03.pdf). お名前は 的省略、话题与语调，以及自我介绍语境。
- **S080** [東京外国語大学言語モジュール: V＋N](https://www.coelang.tufs.ac.jp/mt/ja/gmod/courses/c01/lesson20/step2/card/079.html). 分句修饰名词。
- **S081** [小学館《デジタル大辞泉》或《精選版 日本国語大辞典》: さ (コトバンク)](https://kotobank.jp/word/さ-506834). 句尾助词与接尾词的不同功能。
- **S082** [Aniplex: 《続 夏目友人帳音楽集 いとうるわしきもの》CD 曲目表](https://www.aniplex.co.jp/natsume/cd/02.html). 核对曲名及其为标题；语法切分是另行分析，不是官方剧情解说。
- **S083** [国立国語研究所: 基本動詞ハンドブック「変わる・替わる・換わる・代わる」](https://www2.ninjal.ac.jp/verbhandbook/headwords/変わる・替わる・換わる・代わる.html). 代替、代表角色的 に代わって；不等于变成某物。
- **S084** [東京外国語大学言語モジュール: 敬語](https://www.coelang.tufs.ac.jp/mt/ja/gmod/courses/c01/lesson24/step4/card/092.html). お + 连用形等敬语构成；不当作任意替换规则。
- **S085** [Merriam-Webster Dictionary: escape](https://www.merriam-webster.com/dictionary/escape). 对应词义、词形或 Word History；现代构成分析不冒充词典明载的历史推导。
- **S086** [Merriam-Webster Dictionary: parity](https://www.merriam-webster.com/dictionary/parity). 对应词义、词形或 Word History；现代构成分析不冒充词典明载的历史推导。
- **S087** [Merriam-Webster Dictionary: cache](https://www.merriam-webster.com/dictionary/cache). 对应词义、词形或 Word History；现代构成分析不冒充词典明载的历史推导。
- **S088** [Merriam-Webster Dictionary: candidate](https://www.merriam-webster.com/dictionary/candidate). 对应词义、词形或 Word History；现代构成分析不冒充词典明载的历史推导。
- **S089** [Merriam-Webster Dictionary: etymology](https://www.merriam-webster.com/dictionary/etymology). 对应词义、词形或 Word History；现代构成分析不冒充词典明载的历史推导。
- **S090** [Merriam-Webster Dictionary: mnemonic](https://www.merriam-webster.com/dictionary/mnemonic). 对应词义、词形或 Word History；现代构成分析不冒充词典明载的历史推导。
- **S091** [Merriam-Webster Dictionary: vocabulary](https://www.merriam-webster.com/dictionary/vocabulary). 对应词义、词形或 Word History；现代构成分析不冒充词典明载的历史推导。
- **S092** [Merriam-Webster Dictionary: syntax](https://www.merriam-webster.com/dictionary/syntax). 对应词义、词形或 Word History；现代构成分析不冒充词典明载的历史推导。
- **S093** [Merriam-Webster Dictionary: artifact](https://www.merriam-webster.com/dictionary/artifact). 对应词义、词形或 Word History；现代构成分析不冒充词典明载的历史推导。
- **S094** [Oxford Advanced Learner's Dictionary: idiomatic](https://www.oxfordlearnersdictionaries.com/definition/english/idiomatic). Word Origin 的希腊语来源；不采用未经核实的额外拆解。
- **S095** [小学館《デジタル大辞泉》或《精選版 日本国語大辞典》: 御馳走 (コトバンク)](https://kotobank.jp/word/御馳走-502226). ご 前缀、款待义以及丁寧语形逐渐普遍的语志。
- **S096** [小学館《デジタル大辞泉》或《精選版 日本国語大辞典》: 馳走 (コトバンク)](https://kotobank.jp/word/馳走-565923). 奔走、张罗照料、款待和饭食的语义联系。
- **S097** [三省堂《コンサイスカタカナ語辞典》第5版: 试读](https://dictionary.sanseido-publ.co.jp/wordpress/wp-content/themes/dictionaryandbeyond/pdf/ssd11064.pdf). 检索摘要列出 アイテム [item]；本次 PDF 正文抓取返回 502，只据摘要核对直接借词对应，不扩写远源。
- **S098** [小学館《デジタル大辞泉》或《精選版 日本国語大辞典》: タオル (コトバンク)](https://kotobank.jp/word/たおる-3157938). 采用《精選版 日本国語大辞典》的英语直接借入记录，不混合页面上不同百科的欧洲远源叙述。
- **S099** [Merriam-Webster Dictionary: invaluable](https://www.merriam-webster.com/dictionary/invaluable). 对应词义、词形或 Word History；现代构成分析不冒充词典明载的历史推导。
- **S100** [MDN: Deprecated and obsolete features](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Deprecated_and_obsolete_features). deprecated 与 obsolete 的区别。
- **S101** [Merriam-Webster Dictionary: pronunciation](https://www.merriam-webster.com/dictionary/pronunciation). 对应词义、词形或 Word History；现代构成分析不冒充词典明载的历史推导。
- **S102** [MDN: Object.hasOwn()](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Object/hasOwn). 属性存在与值、原型链是不同判断。
- **S103** [MDN: Lexical grammar](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Lexical_grammar). 字符串转义序列与运行时字符。
- **S104** [國學院大學: 「現代でも用いられる文語的表現」リスト作成の試み](https://k-rain.repo.nii.ac.jp/record/2001962/files/bunkenronsyu_052_005.pdf). 将 人を呪わば穴二つ 列为文语未然ば的假定条件。
- **S105** [消費者庁: 実践ガイド (若年者向け)](https://www.caa.go.jp/policies/policy/consumer_education/consumer_education/nurture/assets/caa_futurer101_250508_02.pdf). 拒绝练习指出 大丈夫です 也可能被对方理解为 OK；用于说明歧义，不作为正式拒绝话术推荐。
- **S106** [小学館《デジタル大辞泉》或《精選版 日本国語大辞典》: さん (コトバンク)](https://kotobank.jp/word/さん-70293). 仅采用 さん 接尾词义项，不引用同页 sun 等其他同音词。
- **S107** [小学館《デジタル大辞泉》或《精選版 日本国語大辞典》: か (コトバンク)](https://kotobank.jp/word/か-456254). 询问、反问及确认等语用。
- **S108** [MDN: Promise.all()](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Promise/all). 结果按输入顺序、任一拒绝及不自动取消。
- **S109** [MDN: Block formatting context](https://developer.mozilla.org/en-US/docs/Web/CSS/Guides/Display/Block_formatting_context). BFC 边界、浮动与外边距折叠。
- **S110** [MDN: Number.EPSILON](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Number/EPSILON). 二进制精度及按量级、输入精度选择容差。
- **S111** [Vue: Rendering Mechanism](https://vuejs.org/guide/extras/rendering-mechanism.html). 虚拟节点、patch 及运行时遍历与分配成本。
- **S112** [MDN: HTTP caching](https://developer.mozilla.org/en-US/docs/Web/HTTP/Guides/Caching). HTTP 缓存 no-cache 与 no-store 的语义。
- **S113** [Redis Docs: Key eviction](https://redis.io/docs/latest/develop/reference/eviction/). LRU、LFU、过期和近似 LRU。
- **S114** [MDN: Cross-Origin Resource Sharing (CORS)](https://developer.mozilla.org/en-US/docs/Web/HTTP/Guides/CORS). 预检、实际请求及凭据的不同条件。
- **S115** [MDN: Mastering margin collapsing](https://developer.mozilla.org/en-US/docs/Web/CSS/Guides/Box_model/Margin_collapsing). 父元素建立 BFC 可阻止与子元素的外边距折叠；兄弟元素的折叠条件另行判断。

## 用户补充的名称与借词信息

补充日期: 2026-11-03。以下信息来自用户对学习材料的说明，记录其明确程度；不冒充官网、工商资料、学校资料或词典正文的独立核验结果。补充记录时仅更新计划；用户后续要求执行后，以下说明已落实到词条，名称拼写改动限于明确确认的「神戸病院」「富士大学」。

| 原词条                                                       | 用户补充                                  | 后续采用范围与保留事项                                                                                     |
|:-------------------------------------------------------------|:------------------------------------------|:-----------------------------------------------------------------------------------------------------------|
| [koube-byouin](../entries/koube-byouin.yml) · こうべ病院     | 「こうべ」应写成「神戸」。                | 用户确认的名称写法为「神戸病院」；具体是哪家医院、地址等仍未提供。与独立地名「神戸」区别记录。             |
| [sakura-daigaku](../entries/sakura-daigaku.yml) · さくら大学 | 「さくら」是樱花。                        | 可以按樱花含义理解名称成分；用户未要求把正式写法改成「桜大学」，也未提供实际学校身份。                     |
| [fuji-daigaku](../entries/fuji-daigaku.yml) · ふじ大学       | 是「富士大学」。                          | 采用用户确认的汉字写法「富士大学」；不据名称自行补充所在地、官网、历史或命名来历。                         |
| [ee-kee-shii](../entries/ee-kee-shii.yml) · エー・ケー・シー | 是名为 A・K・C 的组织。                   | 可以记录它是组织名称及给定的日语字母读音；不扩写缩写全称、业务或具体法律实体。                             |
| [pawaa-denki](../entries/pawaa-denki.yml) · パワー電気       | 大概就是类似于 Power 电气这样的电力公司。 | 保留"大概、类似于"的推测性质。可以帮助理解学习场景，但不能写成已确认的电力公司身份或业务事实。             |
| [burajiru-eaa](../entries/burajiru-eaa.yml) · ブラジルエアー | 就是 Brazil 航空公司。                    | 按用户提供的航空公司名称语境理解；不自行对应某家正式注册的航空公司，也不补写运营状态、航线或官方英文商号。 |

这些补充明确了学习材料中的名称含义或写法，并不同时证明名称背后的全部现实身份。前文保留的"不推定"原则不否定本节的用户补充；仍未确定的细节继续留空。既有稳定 ID、文件名和录音不因这些补充自动改变。

ETY13 的补充: 用户已明确说明「アイテム is item (English Word)」，确认「アイテム」与英语 item 的直接借词对应。后续可以将其记录为用户确认信息；此前出版社试读 PDF 抓取失败不再作为该项内容实施的阻塞。该来源的正文仍未复核，不得写成"已完成词典正文复核"，也不扩写 item 更早的词源。

## 本轮维持现状的内容

- 对初筛未选中的 440 个词条，除用户后续明确补充的名称条目外维持原样；不因覆盖比例而添加字段。
- 简单食材、释义已足够直接的普通词及工作短语，不额外堆叠同义解释。
- 来源不明确的人名、地名、学校或机构名，不猜测命名史、不补写未知汉字或身份。
- 日语单字读音不推导为姓名的通用读法。
- 库内没有独立"节流"词条，本轮不为防抖建立指向不存在条目的关系，也不为凑齐关系而新建词条。
- Context、Example 和 Notes 不做自动清理或截断；保留详细说明和来源。
- 不新增词条、词根卡片、模式字段或审批状态字段。

## 审阅方式与后续实施

可逐项确认、要求修改或明确拒绝；也可明确批准完整批次。示例:

- "M01–M16 文案全部确认": 确认第一批表格中的具体文本及其两端目标，不包含 R、WF、ETY 或 N。
- "WF01–WF84 构成说明确认": 这是此前的范围确认；若批准本轮文本，应明确说“具体文案确认”。
- "ETY01–ETY18 词源候选确认": 确认表中 18 个词条的查证范围，不将候选方向视为已证实的词源或最终文案。
- "R01–R34 关联全部确认": 仅确认 34 组关系，不包含新的助记文案。
- "N01–N23 独立助记有价值": 确认表中 23 个词条的补充价值和范围，不等于逐条最终文案已获批。
- "M07 修改为……；其他 M 文案确认": 记录指定修改，并保持未涉及批次的状态。
- "确认 WF01–WF84、ETY01–ETY18、N01–N23 本轮具体文案，接受 N03 留空": 确认本轮草案与去重建议；ETY13 的直接借词对应另获用户确认，词典正文抓取失败不再阻塞实施，不声称正文已复核。本次已收到这三个批次的文案确认；此前的文案确认本身不授权改动；本轮实施另由用户的明确执行指令授权。

M01–M16 的具体助记文案及其 32 个词条目标、R01–R34 的新增关联，以及 WF01–WF84、ETY01–ETY18、N01–N23 的补充范围均已于 2026-11-03 获批。WF、ETY 和 N 的逐条草案现已列出，具体文案已于同日获批，包括来源补充和 N03 留空决定；ETY13 的直接借词对应另获用户确认，不再阻塞实施；词典正文仍未复核。R 组未包含额外助记文案的批准。用户后续已下达明确执行指令，本轮仅实施批准文案、关联和补充名称说明。

实施时:

1. 重读当前词条与约定，检查 Git 状态并保存源文件、录音的保留基线。
2. M 组仅写入已确认的 `mnemonic`；R 组仅新增已确认的 `related_entries`，保留已有声明和直接邻居。
3. WF、ETY、N 按已确认的具体文本处理；ETY13 的用户确认与词典摘要依据分别记录，不能标成正文已核实。来源标题和 URL 保留在 Notes；仅作经确认的内容补充，不删除现有细节。
4. 保持稳定 ID、读音、标签、优先级、先修关系和现有录音不变。Word、Meaning、Context 和 Example 仅按用户后续确认的名称说明调整；其余内容不变。名称变更不改变 ID、文件名或录音。
5. 使用现有验证和构建命令生成完整导出，不手工编辑 dist/，不生成或替换音频。
6. 报告实际改动、未处理项、警告与验证限制；不暂存、提交、推送或修改 Anki 集合。

## 实施记录 (2026-11-03)

本轮已完成内容实施。实际修改 178 个词条，其中 148 个词条获得记忆辅助字段；未新增或删除词条。其余改动为已批准的关联声明及用户补充的名称信息。下列数量是本轮结果，不替换前文带日期的初筛基线。

| 项目          | 实施结果                                                                                      |
|:--------------|:----------------------------------------------------------------------------------------------|
| M01–M16       | 32 个端点分别写入批准的 mnemonic                                                              |
| WF01–WF84     | 84 条 word_formation，与批准文本逐条一致                                                      |
| ETY01–ETY18   | 18 条 etymology；ETY13 保留用户确认及词典正文未复核的限定                                     |
| N01–N23       | 22 条 mnemonic；N03 不写入，与 WF24 去重                                                      |
| Mnemonic 合计 | 54 个词条，不为 R 组额外生成助记                                                              |
| R01–R34       | 新增 34 组关系，单端声明，导出双向显示                                                        |
| 全库关联      | 原有 21 组保留，共 55 组，覆盖 106 个展示端点                                                 |
| 来源与 Notes  | 106 个词条更新 Notes；原有无关说明保留，相同来源不重复追加                                    |
| 名称补充      | 六项用户说明已落实；两项名称改为「神戸病院」「富士大学」，相关展示文字同步更新，ID 和录音不变 |
| 全库与导出    | 643 个词条，643 行十七列 TSV；两次构建字节一致                                                |
| 保留检查      | 643 个原有 MP3 哈希不变；ID、语言、读音、词性、音频引用、标签、学习优先级及先修关系均不变     |
| 学习顺序      | 与实施前完整构建逐行比较，ID 顺序及 StudyOrder 全部一致                                       |

已运行 `go test ./...`、`go vet ./...`、词库验证和两次完整构建，均通过。原有 6 条同词不同义警告保留，涉及 artifact、harness、literal、resolve、revision、scope；本轮未新增同词警告。另按计划逐项检查全部批准文本、关联、来源记录、空字段与导出列，并独立核对保护字段及录音哈希。

原有 README 改动保留；未修改 Go、依赖、schema、导出格式或模板，也未生成音频。未暂存、提交、推送或修改用户的 Anki 集合。实际 Anki Desktop、AnkiDroid 的导入、显示和播放未在本轮测试。

导出 SHA-256: `9205d5da9702884b299a54edb625d6610b5c2ae1f97b7daf8235589223e47c66`。

## 用户验收 (2026-11-03)

用户确认已完成人工听审及真实 Anki 导入、显示、播放验证。确认原文:

> 我已经完成 人工听审及真实 Anki 导入、显示、播放验证

这些项目记为用户完成的人工验收，来源为用户反馈，不改写为 Codex 自动化测试或独立设备测试的结果。客户端版本及逐项测试记录未提供；本次确认不扩展为 AnkiWeb 同步、迁移时空值清除或复习历史保留等专项验证。

本节补充此前实施记录中的验证边界，不覆盖原有历史结果。人工听审不补写缺失的 IPA，也不等于完成 ETY13 的词典正文复核。

## 验证与验收

创建本计划时仅做文档和候选检查；内容实施与本地检查结果见实施记录。人工听审及真实 Anki 导入、显示、播放验证现已由用户确认完成，见上方用户验收记录。

本轮已运行以下命令；后续内容更新时沿用:

```bash
go test ./...
go vet ./...
go run scripts/validate/main.go
go run scripts/build/main.go
git diff --check
```

同时检查:

- 分开检查新文件格式，因为 git diff --check 不覆盖未跟踪文件。
- M 组文本与确认版本一致，只作用于指定端点；未确认项保持不变。
- R 组不存在缺失目标、自关联、重复声明或重复关系；原有 21 组关系保留。
- 对确认关系检查双向直接展示、邻居顺序及当前释义；不推导传递关系。
- 无新增词条时，词条数保持实施前实际基线；本次历史基线为 643。
- 导出保持十六个笔记字段和标签列，共十七列；ID 仍在第一列。
- 构建两次并比较字节；核对原有学习顺序未被记忆辅助内容改变。
- 原有录音哈希不变；源文件只发生已确认的记忆辅助内容和必要来源补充。
- 静态模板和导出检查不等于已验证实际 Anki 显示。用户后续可完整重导入来刷新两端 RelatedWords；不得自动操作用户的 Anki 集合。
