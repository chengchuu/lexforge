# Codex CLI 词汇导入审阅计划

状态: **C01 至 C22、T01 至 T08、R01 至 R03 已全部批准并实施，静态与导出验证完成。** 查证及实施日期: 2026-11-19。本机版本基线: `codex-cli 0.169.0`。

目标是读懂常用操作、过程提示与结果状态。以英文词语为学习锚点，在 Notes 中解释真实使用场景。规划阶段仅创建本计划。用户随后明确批准全部 C/T/R 项目并要求实施，现已完成。本批不生成音频。

内容与操作边界遵循 [AGENTS.md](../AGENTS.md)、[README.md](../README.md)、[字段指南](ENTRY_SCHEMA.md)、[学习顺序](LEARNING_ORDER.md)和[维护提示](MAINTENANCE_PROMPTS.md)。

## 导航

- [范围与基线](#范围与基线)
- [来源与查证边界](#来源与查证边界)
- [候选与复用决定](#候选与复用决定)
- [新词具体文案](#新词具体文案)
- [操作参考](#操作参考)
- [分组与排除决定](#分组与排除决定)
- [关联候选单独批准](#关联候选单独批准)
- [实施与验证步骤](#实施与验证步骤)
- [审阅与规划检查](#审阅与规划检查)
- [实施与验证记录](#实施与验证记录)

## 范围与基线

以下表格保留规划时基线和批准范围，当前结果见文末实施记录。规划时已按稳定 ID 检索全库，并检查候选的同形词和近义项目。这里只提出有实际学习价值的第一批，不收集全部帮助文字，也不规定填充比例。


| 项目                       | 实际范围                                               |
|:---------------------------|:-------------------------------------------------------|
| 现有词库                   | 1,642 项，英语 1,399 项，日语 243 项                   |
| 新增学习目标               | 22 项，编号 C01 至 C22，均为英语                       |
| 复用现有条目               | 8 项，编号 T01 至 T08，仅拟追加 codex-cli 标签         |
| 主题范围                   | 共 30 个稳定 ID，含先修基础 agent-ai                   |
| 新增先修声明               | 6 条，分布于 5 个新词条；不是掌握程度解锁              |
| 关联候选                   | 3 对，编号 R01 至 R03，须单独批准                      |
| 音频                       | 本批未请求，默认不生成；既有录音不变                   |
| 拟增加或改写的记忆辅助字段 | 0；本批的操作说明放在 Notes                            |
| 获批后的预计词库           | 1,664 项，英语 1,421 项，日语 243 项                   |
| 当前修改权限               | 仅本计划；不改词条、录音、实现、配置、模板、依赖或导出 |


规划时 Git 工作树干净。现有 `codex-cli` 标签使用次数为 0。拟为 22 个新条目与 8 个复用条目使用该主题标签，不新增近似主题标签。

新词使用 `language: en`。Meaning、Context 与 Notes 使用中文，Example 使用英文。全部例句是编写的学习示例，不是本机测试结果或历史事件。词条内容保持纯文本，不增加字段，不依赖 Markdown 渲染。

## 来源与查证边界

以下本地帮助仅执行了只读查询，未启动新会话或执行产品操作。


| 来源 | 核实范围                       | 证据位置                                                                                                                   |
|:-----|:-------------------------------|:---------------------------------------------------------------------------------------------------------------------------|
| L1   | 本机版本、终端子命令和选项     | codex --version；codex --help                                                                                              |
| L2   | 恢复、分支和审查的命令语法     | codex resume --help；codex fork --help；codex review --help                                                                |
| S1   | 0.169.0 的斜杠命令名与用途     | [SlashCommand](https://github.com/openai/codex/blob/rust-v0.169.0/codex-rs/tui/src/slash_command.rs)                       |
| S2   | 目标状态的实际界面文案         | [goal_status_indicator_line](https://github.com/openai/codex/blob/rust-v0.169.0/codex-rs/tui/src/bottom_pane/footer.rs)    |
| S3   | Worked for 后的已用时间显示    | [FinalMessageSeparator](https://github.com/openai/codex/blob/rust-v0.169.0/codex-rs/tui/src/history_cell/separators.rs)    |
| S4   | 回顾、上下文压缩和操作流程     | [Developer commands: CLI](https://learn.chatgpt.com/docs/developer-commands?surface=cli)                                   |
| D1   | 21 个词语的普通词义及美式 IPA  | 下文各条目的 Collins 具体词页                                                                                              |
| D2   | retry 的计算机场景义及美式 IPA | [Oxford Advanced American Dictionary: retry](https://www.oxfordlearnersdictionaries.com/definition/american_english/retry) |


S1、S2 和 S3 固定到与本机版本对应的 `rust-v0.169.0` 标签。S4 是持续更新的官方文档，不能假定所有新功能均已在当前会话启用。未实测任何交互命令，不声称所有 UI 路径均已现场验证。

音标使用已读词典的完整美式转写，不加 `US ` 前缀。保留各词典的 /i/、/u/、/r/ 和重音记法。词典支持普通英语词形的读音，CLI 技术含义另由官方资料支持。部分词页包含多个同形词或词性，不能取页面第一个音标就结束核对。

- recap 采用概述义的美式第二分区，不采用翻新轮胎义。
- resume 采用继续进行的动词分区，不取 résumé 的简历读音。
- compact 采用明确标为动词的 /kəmˈpækt/，不是名词的首音节重音形式。
- stall 采用美式第一分区中停止推进的动词义，不取故意拖延的另一分区作为本项释义。
- retry 的 Collins 页面主要解释法庭重审，未用于本项。Cambridge 页面返回 403。已读 Oxford 美式页面明确列出计算机场景的再次尝试义。
- 本批不新增词源或构词解释，不将可见拼写或命令名称当作历史词源证据。

## 候选与复用决定

新增 ID 和文件名均已检查全局唯一性。下表的词义与优先级是待批准内容，具体字段见下一节。


| 编号 | 稳定 ID / Word                   | 词性            | 拟写 Meaning                                       | 优先级 |
|:-----|:---------------------------------|:----------------|:---------------------------------------------------|:-------|
| C01  | `recap-workflow` / recap         | noun; verb      | 回顾；对前文要点的简要概述，或概述这些要点。       | 1      |
| C02  | `work-general` / work            | verb            | 工作；投入时间和精力处理任务。                     | 1      |
| C03  | `goal-general` / goal            | noun            | 目标；希望通过行动达成的结果。                     | 1      |
| C04  | `achieve-general` / achieve      | verb            | 达成；通过努力取得预期的目标或结果。               | 1      |
| C05  | `resume-workflow` / resume       | verb            | 恢复继续；在中断或暂停后继续进行。                 | 1      |
| C06  | `fork-session-tech` / fork       | verb            | 分出新会话；从已有对话建立可继续推进的独立分支。   | 2      |
| C07  | `compact-context-tech` / compact | verb            | 压缩；将较长上下文整理为更精简的表示。             | 2      |
| C08  | `session-tech` / session         | noun            | 会话；可继续交互并保存相关对话历史的一次交流过程。 | 1      |
| C09  | `prompt-ai` / prompt             | noun            | 提示词；向模型提供的任务要求和相关输入。           | 1      |
| C10  | `command-cli` / command          | noun            | 命令；要求程序执行某项操作的指令。                 | 1      |
| C11  | `status-workflow` / status       | noun            | 状态；某项任务、会话或资源当前所处的情况。         | 1      |
| C12  | `permission-tech` / permission   | noun            | 权限；允许执行某种操作或访问某项资源的许可。       | 2      |
| C13  | `approval-workflow` / approval   | noun            | 批准；对某项提议或操作给予同意。                   | 2      |
| C14  | `sandbox-tech` / sandbox         | noun            | 沙箱；限制程序可访问资源和可执行操作的受约束环境。 | 2      |
| C15  | `reasoning-ai` / reasoning       | noun            | 推理；依据已知信息进行分析并形成结论的过程。       | 2      |
| C16  | `interrupt-workflow` / interrupt | verb            | 中断；使正在进行的活动暂时或提前停止。             | 2      |
| C17  | `retry-workflow` / retry         | verb            | 重试；在先前尝试失败后再次进行尝试。               | 2      |
| C18  | `pause-workflow` / pause         | verb            | 暂停；暂时停止推进，保留稍后继续的可能。           | 1      |
| C19  | `stall-workflow` / stall         | verb            | 停滞；因障碍而无法继续推进。                       | 2      |
| C20  | `complete-workflow` / complete   | verb; adjective | 完成；已结束所要求的工作，或使工作达到完成状态。   | 1      |
| C21  | `summary-language` / summary     | noun            | 摘要；保留主要内容的简短概述。                     | 1      |
| C22  | `transcript-tech` / transcript   | noun            | 文字记录；按交流过程保留的会话内容。               | 2      |


复用决定如下。当前释义、例句、Notes、音标、录音、学习元数据及关系全部保留，实施时只追加一次 `codex-cli` 标签。操作说明保留在本计划中，不自动覆盖既有内容。


| 编号 | 现有词条                                                                     | 当前 Meaning                                                       | 当前 Example                                                |
|:-----|:-----------------------------------------------------------------------------|:-------------------------------------------------------------------|:------------------------------------------------------------|
| T01  | [review-workflow](../entries/review-workflow.yml) / review                   | 审查；检查工作内容并给出判断或反馈。                               | Review the patch before merging it.                         |
| T02  | [context-language](../entries/context-language.yml) / context                | 语境；上下文；帮助理解含义的背景。                                 | The context makes the meaning clear.                        |
| T03  | [token-llm](../entries/token-llm.yml) / token                                | 词元；语言模型将文本编码后使用的基本处理单位。                     | Keep the token count within the model's context limit.      |
| T04  | [pending-tech](../entries/pending-tech.yml) / pending                        | 待处理的；待完成的                                                 | The framework tests are still pending.                      |
| T05  | [read-only-workflow](../entries/read-only-workflow.yml) / read-only          | 只读的                                                             | Keep the investigation read-only.                           |
| T06  | [elapsed-time-workflow](../entries/elapsed-time-workflow.yml) / elapsed time | 已用时间                                                           | Check elapsed time before starting another build.           |
| T07  | [codex-ai](../entries/codex-ai.yml) / Codex                                  | OpenAI 的 AI 编程智能体。                                          | Ask Codex to review the diff before committing the changes. |
| T08  | [agent-ai](../entries/agent-ai.yml) / agent                                  | 智能体；围绕目标选择步骤、调用工具，并根据执行结果继续行动的系统。 | The agent inspects the repository before editing files.     |

| 编号 | 本主题中的复用理由或操作说明                                                 |
|:-----|:-----------------------------------------------------------------------------|
| T01  | /review 检查当前改动；也可在终端使用 codex review。审查结论不代替运行测试。  |
| T02  | 在本主题中理解对话背景、文件和要求；复用通用上下文义，不复用 context-react。 |
| T03  | 结合 /status 阅读用量和上下文信息；保留现有词元解释。                        |
| T04  | 用于理解仍待处理的状态；不推断当前版本一定显示独立的 Pending 标签。          |
| T05  | 结合 --sandbox read-only 理解只读边界；保留现有释义与例句。                  |
| T06  | 对应 Worked for 后的耗时信息；不将耗时当作测试或任务成功证据。               |
| T07  | 复用现有产品条目，不把 CLI 客户端单独当作同一品牌的新义项。                  |
| T08  | 作为 codex-ai 的既有先修基础，加入同一主题，保持筛选后的先修闭包。           |


复用范围包含 `codex-ai` 的先修 `agent-ai`。8 个现有条目的全部直接、间接先修均在本主题内。新词的先修也全部指向这 30 个拟纳入主题的稳定 ID，便于按标签独立学习。

## 新词具体文案

每个 YAML 块是批准后拟写入的完整条目草案。仅供审阅，本轮不创建这些文件。未使用的可选字段省略，Notes 中的来源也属于审阅文案。


### C01 recap

拟新增文件: `entries/recap-workflow.yml`。

```yaml
id: "recap-workflow"
language: "en"
word: "recap"
pronunciation: "/ˈriˌkæp/"
part_of_speech: "noun; verb"
meaning: "回顾；对前文要点的简要概述，或概述这些要点。"
context: "在 Codex CLI 中回顾对话要点。"
example: "Recap the decisions before starting the next task."
notes: |-
  /recap 生成当前对话摘要，适合重新进入一段对话时使用。与 /compact 的用途不同，/recap 不用于压缩当前上下文。这里采用概述义，不采用翻新轮胎义。
  学习语境: Codex CLI 0.169.0。
  操作参考: OpenAI Codex — SlashCommand. https://github.com/openai/codex/blob/rust-v0.169.0/codex-rs/tui/src/slash_command.rs
  操作参考: OpenAI — Developer commands (CLI). https://learn.chatgpt.com/docs/developer-commands?surface=cli
  Pronunciation source: Collins Dictionary — recap; recap in American English 2. https://www.collinsdictionary.com/dictionary/english/recap (accent: US; accessed 2026-11-19).
learning_priority: 1
tags:
  - general
  - codex-cli
  - noun
  - verb
```


### C02 work

拟新增文件: `entries/work-general.yml`。

```yaml
id: "work-general"
language: "en"
word: "work"
pronunciation: "/wɜrk/"
part_of_speech: "verb"
meaning: "工作；投入时间和精力处理任务。"
context: "理解界面中的处理耗时。"
example: "Codex worked on the task for two minutes."
notes: |-
  Worked for 2m 表示这轮处理已用两分钟。worked 是 work 的过去式，不单独建卡。该提示不证明修改正确或测试通过。It worked 中的 work 可表示奏效，属于另一用法。
  学习语境: Codex CLI 0.169.0。
  操作参考: OpenAI Codex — FinalMessageSeparator. https://github.com/openai/codex/blob/rust-v0.169.0/codex-rs/tui/src/history_cell/separators.rs
  Pronunciation source: Collins Dictionary — work; work in American English. https://www.collinsdictionary.com/dictionary/english/work (accent: US; accessed 2026-11-19).
learning_priority: 1
tags:
  - general
  - codex-cli
  - verb
```


### C03 goal

拟新增文件: `entries/goal-general.yml`。

```yaml
id: "goal-general"
language: "en"
word: "goal"
pronunciation: "/ɡoʊl/"
part_of_speech: "noun"
meaning: "目标；希望通过行动达成的结果。"
context: "理解持续任务的目标和状态。"
example: "The goal is to fix the bug without changing the API."
notes: |-
  /goal 用于查看当前目标，/goal 后接目标文本可设置目标。Goal achieved 表示目标被标记为完成；仍需查看具体结果与验证证据。当前词库已有 objective，英语词形不同，保留各自学习价值。
  学习语境: Codex CLI 0.169.0。
  操作参考: OpenAI Codex — SlashCommand. https://github.com/openai/codex/blob/rust-v0.169.0/codex-rs/tui/src/slash_command.rs
  操作参考: OpenAI Codex — goal_status_indicator_line. https://github.com/openai/codex/blob/rust-v0.169.0/codex-rs/tui/src/bottom_pane/footer.rs
  操作参考: OpenAI — Developer commands (CLI). https://learn.chatgpt.com/docs/developer-commands?surface=cli
  Pronunciation source: Collins Dictionary — goal; goal in American English. https://www.collinsdictionary.com/dictionary/english/goal (accent: US; accessed 2026-11-19).
learning_priority: 1
tags:
  - general
  - codex-cli
  - noun
```


### C04 achieve

拟新增文件: `entries/achieve-general.yml`。

```yaml
id: "achieve-general"
language: "en"
word: "achieve"
pronunciation: "/əˈtʃiv/"
part_of_speech: "verb"
meaning: "达成；通过努力取得预期的目标或结果。"
context: "理解 Goal achieved 中的动作含义。"
example: "We achieved the goal after the tests passed."
notes: |-
  常见搭配为 achieve a goal。Goal achieved 是表示目标已达成的简短状态提示，achieved 是过去分词。保留完整提示作为用法，不另建整句卡片。
  学习语境: Codex CLI 0.169.0。
  操作参考: OpenAI Codex — goal_status_indicator_line. https://github.com/openai/codex/blob/rust-v0.169.0/codex-rs/tui/src/bottom_pane/footer.rs
  Pronunciation source: Collins Dictionary — achieve; achieve in American English. https://www.collinsdictionary.com/dictionary/english/achieve (accent: US; accessed 2026-11-19).
learning_priority: 1
tags:
  - general
  - codex-cli
  - verb
```


### C05 resume

拟新增文件: `entries/resume-workflow.yml`。

```yaml
id: "resume-workflow"
language: "en"
word: "resume"
pronunciation: "/rɪˈzum/"
part_of_speech: "verb"
meaning: "恢复继续；在中断或暂停后继续进行。"
context: "继续已有会话，或恢复目标推进。"
example: "Resume the session after reviewing the recap."
notes: |-
  在终端运行 codex resume 可选择已有会话，codex resume --last 可继续最近会话。交互界面中的 /resume 用于选择已有会话；/goal resume 则恢复当前目标推进。这里不是表示简历的 résumé。已有 resumption-workflow 是名词，本项学习常用动词与操作含义。
  学习语境: Codex CLI 0.169.0。
  操作参考: OpenAI Codex — SlashCommand. https://github.com/openai/codex/blob/rust-v0.169.0/codex-rs/tui/src/slash_command.rs
  操作参考: OpenAI — Developer commands (CLI). https://learn.chatgpt.com/docs/developer-commands?surface=cli
  Pronunciation source: Collins Dictionary — resume; resume in American English. https://www.collinsdictionary.com/dictionary/english/resume (accent: US; accessed 2026-11-19).
learning_priority: 1
tags:
  - general
  - codex-cli
  - verb
```


### C06 fork

拟新增文件: `entries/fork-session-tech.yml`。

```yaml
id: "fork-session-tech"
language: "en"
word: "fork"
pronunciation: "/fɔrk/"
part_of_speech: "verb"
meaning: "分出新会话；从已有对话建立可继续推进的独立分支。"
context: "尝试另一种方案，同时保留原对话。"
example: "Fork the session to explore another approach."
notes: |-
  /fork 从当前对话分出新会话。终端中的 codex fork 用于选择已有会话再分支。会话分支与 Git 仓库派生、进程 fork 的含义不同。分出会话不等于隔离工作文件；是否使用新 worktree 应核对界面选择。
  学习语境: Codex CLI 0.169.0。
  操作参考: OpenAI Codex — SlashCommand. https://github.com/openai/codex/blob/rust-v0.169.0/codex-rs/tui/src/slash_command.rs
  操作参考: OpenAI — Developer commands (CLI). https://learn.chatgpt.com/docs/developer-commands?surface=cli
  Pronunciation source: Collins Dictionary — fork; fork in American English. https://www.collinsdictionary.com/dictionary/english/fork (accent: US; accessed 2026-11-19).
learning_priority: 2
prerequisites:
  - session-tech
tags:
  - tech
  - codex-cli
  - verb
```


### C07 compact

拟新增文件: `entries/compact-context-tech.yml`。

```yaml
id: "compact-context-tech"
language: "en"
word: "compact"
pronunciation: "/kəmˈpækt/"
part_of_speech: "verb"
meaning: "压缩；将较长上下文整理为更精简的表示。"
context: "长对话接近上下文容量时，压缩已积累的上下文。"
example: "Compact the context before continuing the task."
notes: |-
  /compact 对对话上下文作摘要压缩，以释放上下文空间。/recap 用于回顾，/compact 会影响供后续处理使用的上下文。压缩摘要不等于逐字保留全部记录。当前词性为动词，不取表示协议或粉盒的名词读音。
  学习语境: Codex CLI 0.169.0。
  操作参考: OpenAI Codex — SlashCommand. https://github.com/openai/codex/blob/rust-v0.169.0/codex-rs/tui/src/slash_command.rs
  操作参考: OpenAI — Developer commands (CLI). https://learn.chatgpt.com/docs/developer-commands?surface=cli
  Pronunciation source: Collins Dictionary — compact; compact in American English 1; verb. https://www.collinsdictionary.com/dictionary/english/compact (accent: US; accessed 2026-11-19).
learning_priority: 2
prerequisites:
  - context-language
  - summary-language
tags:
  - tech
  - codex-cli
  - verb
```


### C08 session

拟新增文件: `entries/session-tech.yml`。

```yaml
id: "session-tech"
language: "en"
word: "session"
pronunciation: "/ˈsɛʃən/"
part_of_speech: "noun"
meaning: "会话；可继续交互并保存相关对话历史的一次交流过程。"
context: "区分当前会话、历史会话和单次回复。"
example: "Resume the previous session."
notes: |-
  session 指会话，不等于一次模型回复、一个终端进程或当前目录。codex resume 用于继续已有会话，codex fork 用于分出会话。工作目录与会话历史需要分别理解。
  学习语境: Codex CLI 0.169.0。
  操作参考: OpenAI Codex — SlashCommand. https://github.com/openai/codex/blob/rust-v0.169.0/codex-rs/tui/src/slash_command.rs
  操作参考: OpenAI — Developer commands (CLI). https://learn.chatgpt.com/docs/developer-commands?surface=cli
  Pronunciation source: Collins Dictionary — session; session in American English. https://www.collinsdictionary.com/dictionary/english/session (accent: US; accessed 2026-11-19).
learning_priority: 1
tags:
  - tech
  - codex-cli
  - noun
```


### C09 prompt

拟新增文件: `entries/prompt-ai.yml`。

```yaml
id: "prompt-ai"
language: "en"
word: "prompt"
pronunciation: "/prɑmpt/"
part_of_speech: "noun"
meaning: "提示词；向模型提供的任务要求和相关输入。"
context: "表达目标、约束和验收标准。"
example: "Write a clear prompt before starting the task."
notes: |-
  Codex CLI 接受任务提示作为输入。提示词可描述任务目标、修改边界和检查要求。这是撰写建议，不是 CLI 强制的固定格式。这里不是表示及时的形容词，也不专指 Shell 命令提示符。
  学习语境: Codex CLI 0.169.0。
  操作参考: OpenAI — Developer commands (CLI). https://learn.chatgpt.com/docs/developer-commands?surface=cli
  Pronunciation source: Collins Dictionary — prompt; prompt in American English. https://www.collinsdictionary.com/dictionary/english/prompt (accent: US; accessed 2026-11-19).
learning_priority: 1
tags:
  - tech
  - codex-cli
  - noun
```


### C10 command

拟新增文件: `entries/command-cli.yml`。

```yaml
id: "command-cli"
language: "en"
word: "command"
pronunciation: "/kəˈmænd/"
part_of_speech: "noun"
meaning: "命令；要求程序执行某项操作的指令。"
context: "区分终端子命令与交互界面的斜杠命令。"
example: "Check the command before running it."
notes: |-
  codex resume 是终端命令，/resume 是 Codex 交互界面的斜杠命令。普通英语提示不是自动可执行的命令。codex --help 和 codex resume --help 可查看本机帮助。其他子命令可按同样方式查询。
  学习语境: Codex CLI 0.169.0。
  操作参考: OpenAI Codex — SlashCommand. https://github.com/openai/codex/blob/rust-v0.169.0/codex-rs/tui/src/slash_command.rs
  操作参考: OpenAI — Developer commands (CLI). https://learn.chatgpt.com/docs/developer-commands?surface=cli
  Pronunciation source: Collins Dictionary — command; command in American English. https://www.collinsdictionary.com/dictionary/english/command (accent: US; accessed 2026-11-19).
learning_priority: 1
tags:
  - tech
  - codex-cli
  - noun
```


### C11 status

拟新增文件: `entries/status-workflow.yml`。

```yaml
id: "status-workflow"
language: "en"
word: "status"
pronunciation: "/ˈstætəs/"
part_of_speech: "noun"
meaning: "状态；某项任务、会话或资源当前所处的情况。"
context: "读取会话信息和目标状态。"
example: "Check the current status before resuming the goal."
notes: |-
  /status 显示当前会话配置和 token 用量。Goal paused、Goal stalled 和 Goal achieved 则是目标状态提示。已有 exit-status-tech 表示进程退出状态，本项为更通用的状态含义。
  学习语境: Codex CLI 0.169.0。
  操作参考: OpenAI Codex — SlashCommand. https://github.com/openai/codex/blob/rust-v0.169.0/codex-rs/tui/src/slash_command.rs
  操作参考: OpenAI Codex — goal_status_indicator_line. https://github.com/openai/codex/blob/rust-v0.169.0/codex-rs/tui/src/bottom_pane/footer.rs
  操作参考: OpenAI — Developer commands (CLI). https://learn.chatgpt.com/docs/developer-commands?surface=cli
  Pronunciation source: Collins Dictionary — status; status in American English. https://www.collinsdictionary.com/dictionary/english/status (accent: US; accessed 2026-11-19).
learning_priority: 1
tags:
  - tech
  - codex-cli
  - noun
```


### C12 permission

拟新增文件: `entries/permission-tech.yml`。

```yaml
id: "permission-tech"
language: "en"
word: "permission"
pronunciation: "/pərˈmɪʃən/"
part_of_speech: "noun"
meaning: "权限；允许执行某种操作或访问某项资源的许可。"
context: "理解允许的文件访问与操作范围。"
example: "The agent does not have permission to edit this directory."
notes: |-
  /permissions 用于查看或选择允许 Codex 执行的操作范围。界面常用复数 permissions，词条采用单数 permission。权限范围与是否需要审批是两个维度，不应把无需询问等同于任意文件访问权。
  学习语境: Codex CLI 0.169.0。
  操作参考: OpenAI Codex — SlashCommand. https://github.com/openai/codex/blob/rust-v0.169.0/codex-rs/tui/src/slash_command.rs
  操作参考: OpenAI — Developer commands (CLI). https://learn.chatgpt.com/docs/developer-commands?surface=cli
  Pronunciation source: Collins Dictionary — permission; permission in American English. https://www.collinsdictionary.com/dictionary/english/permission (accent: US; accessed 2026-11-19).
learning_priority: 2
tags:
  - tech
  - codex-cli
  - noun
```


### C13 approval

拟新增文件: `entries/approval-workflow.yml`。

```yaml
id: "approval-workflow"
language: "en"
word: "approval"
pronunciation: "/əˈpruvəl/"
part_of_speech: "noun"
meaning: "批准；对某项提议或操作给予同意。"
context: "理解执行前的审批请求。"
example: "The command is waiting for your approval."
notes: |-
  approval 关注某次操作是否获准，permission 关注允许的访问和操作范围。审批方式受会话策略影响。这里学习批准操作的含义，不把某个审批策略作为所有会话的固定默认值。
  学习语境: Codex CLI 0.169.0。
  操作参考: OpenAI — Developer commands (CLI). https://learn.chatgpt.com/docs/developer-commands?surface=cli
  Pronunciation source: Collins Dictionary — approval; approval in American English. https://www.collinsdictionary.com/dictionary/english/approval (accent: US; accessed 2026-11-19).
learning_priority: 2
tags:
  - tech
  - codex-cli
  - noun
```


### C14 sandbox

拟新增文件: `entries/sandbox-tech.yml`。

```yaml
id: "sandbox-tech"
language: "en"
word: "sandbox"
pronunciation: "/ˈsændˌbɑks/"
part_of_speech: "noun"
meaning: "沙箱；限制程序可访问资源和可执行操作的受约束环境。"
context: "理解文件系统与执行环境的限制。"
example: "The sandbox limits where the command can write."
notes: |-
  Codex CLI 的 --sandbox 选项选择执行隔离策略。本机帮助列出 read-only、workspace-write 和 danger-full-access。这里学习策略名称和边界含义，不建议为了方便自动扩大权限。沙箱策略与审批策略需要分开核对。
  学习语境: Codex CLI 0.169.0。
  操作参考: OpenAI — Developer commands (CLI). https://learn.chatgpt.com/docs/developer-commands?surface=cli
  Pronunciation source: Collins Dictionary — sandbox; sandbox in American English. https://www.collinsdictionary.com/dictionary/english/sandbox (accent: US; accessed 2026-11-19).
learning_priority: 2
prerequisites:
  - permission-tech
tags:
  - tech
  - codex-cli
  - noun
```


### C15 reasoning

拟新增文件: `entries/reasoning-ai.yml`。

```yaml
id: "reasoning-ai"
language: "en"
word: "reasoning"
pronunciation: "/ˈrizənɪŋ/"
part_of_speech: "noun"
meaning: "推理；依据已知信息进行分析并形成结论的过程。"
context: "理解模型设置中的 reasoning effort。"
example: "This task may require more reasoning."
notes: |-
  /model 的本机版本描述包含模型和 reasoning effort 的选择。reasoning effort 表达推理投入设置，不等于已完成的工作量，也不保证结论正确。可见回复不代表完整内部推理记录。
  学习语境: Codex CLI 0.169.0。
  操作参考: OpenAI Codex — SlashCommand. https://github.com/openai/codex/blob/rust-v0.169.0/codex-rs/tui/src/slash_command.rs
  Pronunciation source: Collins Dictionary — reasoning; reasoning in American English. https://www.collinsdictionary.com/dictionary/english/reasoning (accent: US; accessed 2026-11-19).
learning_priority: 2
prerequisites:
  - agent-ai
tags:
  - tech
  - codex-cli
  - noun
```


### C16 interrupt

拟新增文件: `entries/interrupt-workflow.yml`。

```yaml
id: "interrupt-workflow"
language: "en"
word: "interrupt"
pronunciation: "/ˌɪntəˈrʌpt/"
part_of_speech: "verb"
meaning: "中断；使正在进行的活动暂时或提前停止。"
context: "在需要检查或调整方向时停止当前推进。"
example: "The user interrupted the task to review the proposed changes."
notes: |-
  interrupt 表达中断活动，不是本批核实到的内建斜杠命令名。具体中断方式以当前界面的快捷键提示为准。中断一轮任务不等于删除会话，也不保证撤销已完成的文件修改。
  学习语境: Codex CLI 0.169.0。
  操作参考: OpenAI — Developer commands (CLI). https://learn.chatgpt.com/docs/developer-commands?surface=cli
  Pronunciation source: Collins Dictionary — interrupt; interrupt in American English. https://www.collinsdictionary.com/dictionary/english/interrupt (accent: US; accessed 2026-11-19).
learning_priority: 2
tags:
  - tech
  - codex-cli
  - verb
```


### C17 retry

拟新增文件: `entries/retry-workflow.yml`。

```yaml
id: "retry-workflow"
language: "en"
word: "retry"
pronunciation: "/ˌriˈtraɪ/"
part_of_speech: "verb"
meaning: "重试；在先前尝试失败后再次进行尝试。"
context: "先处理失败原因，再考虑重新执行。"
example: "Retry after fixing the missing dependency."
notes: |-
  retry 是操作说明中的普通动词。本机内建命令清单未列出 /retry，不应凭词义编造命令。重试不保证自动回滚，也不意味着重复执行一定没有副作用。当前采用计算机场景的再次尝试义，不取法庭重审义作为释义依据。
  学习语境: Codex CLI 0.169.0。
  操作参考: OpenAI — Developer commands (CLI). https://learn.chatgpt.com/docs/developer-commands?surface=cli
  Pronunciation source: Oxford Advanced American Dictionary — retry; retry, sense 2 (NAmE). https://www.oxfordlearnersdictionaries.com/definition/american_english/retry (accent: US; accessed 2026-11-19).
learning_priority: 2
tags:
  - tech
  - codex-cli
  - verb
```


### C18 pause

拟新增文件: `entries/pause-workflow.yml`。

```yaml
id: "pause-workflow"
language: "en"
word: "pause"
pronunciation: "/pɔz/"
part_of_speech: "verb"
meaning: "暂停；暂时停止推进，保留稍后继续的可能。"
context: "暂时停止目标的持续推进。"
example: "Pause the goal while you review the changes."
notes: |-
  /goal pause 暂停目标，/goal resume 请求恢复推进。Goal paused 是对应的界面提示。暂停目标不应当作所有后台进程已终止的凭据；需要时单独检查运行状态。
  学习语境: Codex CLI 0.169.0。
  操作参考: OpenAI Codex — goal_status_indicator_line. https://github.com/openai/codex/blob/rust-v0.169.0/codex-rs/tui/src/bottom_pane/footer.rs
  操作参考: OpenAI — Developer commands (CLI). https://learn.chatgpt.com/docs/developer-commands?surface=cli
  Pronunciation source: Collins Dictionary — pause; pause in American English. https://www.collinsdictionary.com/dictionary/english/pause (accent: US; accessed 2026-11-19).
learning_priority: 1
tags:
  - general
  - codex-cli
  - verb
```


### C19 stall

拟新增文件: `entries/stall-workflow.yml`。

```yaml
id: "stall-workflow"
language: "en"
word: "stall"
pronunciation: "/stɔl/"
part_of_speech: "verb"
meaning: "停滞；因障碍而无法继续推进。"
context: "理解 Goal stalled 与主动暂停的区别。"
example: "The goal stalled because a required file was missing."
notes: |-
  Goal stalled 表示目标推进受阻，与用户主动暂停的 Goal paused 不同。先确认并解决阻碍，再考虑 /goal resume。源码中的状态名为 Blocked，显示文字为 Goal stalled；不把内部枚举名当作界面原文。
  学习语境: Codex CLI 0.169.0。
  操作参考: OpenAI Codex — goal_status_indicator_line. https://github.com/openai/codex/blob/rust-v0.169.0/codex-rs/tui/src/bottom_pane/footer.rs
  Pronunciation source: Collins Dictionary — stall; stall in American English 1; sense 9. https://www.collinsdictionary.com/dictionary/english/stall (accent: US; accessed 2026-11-19).
learning_priority: 2
tags:
  - tech
  - codex-cli
  - verb
```


### C20 complete

拟新增文件: `entries/complete-workflow.yml`。

```yaml
id: "complete-workflow"
language: "en"
word: "complete"
pronunciation: "/kəmˈplit/"
part_of_speech: "verb; adjective"
meaning: "完成；已结束所要求的工作，或使工作达到完成状态。"
context: "理解完成状态与验收证据。"
example: "Complete the review before making more changes."
notes: |-
  completed 是 complete 的过去式或过去分词。目标完成状态与 Goal achieved 相对应，但这不表示界面一定显示单独的 Complete。完成某一步与完成整个目标应分别判断，并检查适用的验证结果。
  学习语境: Codex CLI 0.169.0。
  操作参考: OpenAI Codex — goal_status_indicator_line. https://github.com/openai/codex/blob/rust-v0.169.0/codex-rs/tui/src/bottom_pane/footer.rs
  Pronunciation source: Collins Dictionary — complete; complete in American English. https://www.collinsdictionary.com/dictionary/english/complete (accent: US; accessed 2026-11-19).
learning_priority: 1
tags:
  - general
  - codex-cli
  - verb
  - adjective
```


### C21 summary

拟新增文件: `entries/summary-language.yml`。

```yaml
id: "summary-language"
language: "en"
word: "summary"
pronunciation: "/ˈsʌməri/"
part_of_speech: "noun"
meaning: "摘要；保留主要内容的简短概述。"
context: "阅读回顾结果或压缩后的要点。"
example: "Include the remaining tasks in the summary."
notes: |-
  summary 是摘要结果的名称。/recap 会生成用于回顾的摘要，/compact 用摘要压缩上下文。摘要不等于完整会话记录，也不能保证保留每一个细节。
  学习语境: Codex CLI 0.169.0。
  操作参考: OpenAI Codex — SlashCommand. https://github.com/openai/codex/blob/rust-v0.169.0/codex-rs/tui/src/slash_command.rs
  操作参考: OpenAI — Developer commands (CLI). https://learn.chatgpt.com/docs/developer-commands?surface=cli
  Pronunciation source: Collins Dictionary — summary; summary in American English; noun. https://www.collinsdictionary.com/dictionary/english/summary (accent: US; accessed 2026-11-19).
learning_priority: 1
tags:
  - general
  - codex-cli
  - noun
```


### C22 transcript

拟新增文件: `entries/transcript-tech.yml`。

```yaml
id: "transcript-tech"
language: "en"
word: "transcript"
pronunciation: "/ˈtrænˌskrɪpt/"
part_of_speech: "noun"
meaning: "文字记录；按交流过程保留的会话内容。"
context: "回看或导出对话记录。"
example: "Read the transcript to check the earlier instructions."
notes: |-
  /export 导出会话记录为 Markdown。transcript 与只保留要点的 summary 不同。可见对话记录不等于完整内部推理或所有系统日志；本计划不执行导出，也不复制真实会话内容。
  学习语境: Codex CLI 0.169.0。
  操作参考: OpenAI Codex — SlashCommand. https://github.com/openai/codex/blob/rust-v0.169.0/codex-rs/tui/src/slash_command.rs
  操作参考: OpenAI — Developer commands (CLI). https://learn.chatgpt.com/docs/developer-commands?surface=cli
  Pronunciation source: Collins Dictionary — transcript; transcript in American English. https://www.collinsdictionary.com/dictionary/english/transcript (accent: US; accessed 2026-11-19).
learning_priority: 2
prerequisites:
  - session-tech
tags:
  - tech
  - codex-cli
  - noun
```


## 操作参考

下表是使用时的查阅说明，不执行命令，也不为每一种命令组合创建卡片。终端命令与 Codex 对话输入框中的斜杠命令分别列出。


| 位置     | 输入                          | 用途                           | 需要区分的边界                             |
|:---------|:------------------------------|:-------------------------------|:-------------------------------------------|
| 交互界面 | `/recap`                      | 回顾当前对话要点。             | 用于重新了解进展，不等同于释放上下文空间。 |
| 交互界面 | `/compact`                    | 摘要压缩当前上下文。           | 可能省略细节，不代替完整记录。             |
| 交互界面 | `/resume`                     | 选择一个已有会话继续。         | 与恢复目标的 /goal resume 区分。           |
| 终端     | `codex resume --last`         | 继续最近的已有会话。           | 选择范围以本机帮助和当前目录为准。         |
| 交互界面 | `/fork`                       | 从当前对话分出新会话。         | 会话分支与文件隔离分开理解。               |
| 终端     | `codex fork`                  | 选择一个已有会话作为分支起点。 | 不是直接修改 Git 远程仓库。                |
| 交互界面 | `/goal`                       | 查看当前目标。                 | 后接目标文本用于设置目标，不在本轮执行。   |
| 交互界面 | `/goal pause`、`/goal resume` | 暂停或恢复目标推进。           | 恢复推进不代替解决缺失依赖等阻碍。         |
| 交互界面 | `/status`                     | 查看会话配置与 token 用量。    | 与某条任务的完成状态区分。                 |
| 交互界面 | `/permissions`                | 查看或选择操作权限。           | 不在本轮调整权限或审批策略。               |
| 交互界面 | `/review`                     | 审查当前改动。                 | 审查结论与实际测试结果分别判断。           |
| 交互界面 | `/export`                     | 导出对话为 Markdown。          | 可能写入文件或复制内容，本轮不执行。       |
| 交互界面 | `/model`                      | 选择模型与可用推理投入设置。   | 本轮不改变模型、配置或推理设置。           |
| 终端     | `codex --help`                | 查看本机命令帮助。             | 优先核实本机版本，避免照搬新版本文档。     |


## 分组与排除决定

- 用户给出的 3 个示例不是 3 个强制新卡。recap 对应 C01；worked 归入 C02 work；Goal achieved 由 C03 goal 与 C04 achieve 解释。T06 elapsed time 支持理解耗时。
- worked、achieved、paused 和 stalled 保留在例句或 Notes 中，不另建屈折形式卡片。permissions 归入 permission。
- recap、summary、transcript 分别用于回顾、摘要结果和对话记录。保留独立学习目标，不因中文释义相近而合并。
- 新词 resume 与既有 resumption-workflow 分别学习动词与名词。原条目不改写，也不自动创建关联。
- fork-session-tech 与 branch-git、isolated-worktree-workflow 语境不同。会话分支、Git 引用和文件隔离不混为同一词义。
- status-workflow 不复用更专门的 exit-status-tech。context-language 可支持一般上下文理解，context-react 则保留 React 专义。
- goal 与 objective-ui 是不同英语词形。后者保留其现有游戏界面语境，不改成 Codex 条目。
- complete 是有独立学习价值的编辑补充目标。源码使用 Complete 状态名，界面则显示 Goal achieved；不声称 UI 一定出现单独的 Complete。
- /recap、/compact 和其他斜杠写法放在 Notes 或操作表。普通英语词语、界面状态、Shell 子命令不按同一种结构建卡。
- 不引入 /worked、/achieved、/interrupt 或 /retry 这样的推测命令。本机核实的内建清单没有这些命令名。
- 暂不收集全部快捷键、实验功能、错误信息或选项组合。也不推测用户配置的默认审批策略，不改动现有 Codex 配置。

3 个用户示例、22 个新增词、8 个复用条目和 14 行操作参考属于不同计数维度。它们有交叉映射，不能相加作为源材料总数。本批没有逐条复制整份官方命令参考。

## 关联候选单独批准

以下仅为 R01 至 R03 候选。批准 C/T 条目不等于批准关系。当前拟用端点都在本批 22 个新条目内，实施前须先创建获批端点。每对只由左端声明右端，保留其余既有成员。


| 编号 | 声明端 ID / Word / Meaning                                                    | 另一端 ID / Word / Meaning                                        | 联想或辨析价值                                               | 混淆风险                                                       |
|:-----|:------------------------------------------------------------------------------|:------------------------------------------------------------------|:-------------------------------------------------------------|:---------------------------------------------------------------|
| R01  | `recap-workflow` / recap / 回顾；对前文要点的简要概述，或概述这些要点。       | `summary-language` / summary / 摘要；保留主要内容的简短概述。     | recap 可以表示回顾动作或回顾内容，summary 主要命名摘要结果。 | 两者有重叠含义，但不能据此把 /recap 与 /compact 当作同一操作。 |
| R02  | `permission-tech` / permission / 权限；允许执行某种操作或访问某项资源的许可。 | `approval-workflow` / approval / 批准；对某项提议或操作给予同意。 | 对比允许的操作范围与某次操作是否获准。                       | 获得一次批准不等于获得任意资源访问权限。                       |
| R03  | `pause-workflow` / pause / 暂停；暂时停止推进，保留稍后继续的可能。           | `stall-workflow` / stall / 停滞；因障碍而无法继续推进。           | 对比主动暂停与因障碍停滞。                                   | 两种状态都可能停止推进，但原因不同；resume 不能替代排除阻碍。  |


全部批准时新增 3 条单端声明，涉及 6 个新词条，导出产生 6 个直接邻居展示 (含 3 个反向展示)。本轮没有伴随的 mnemonic 改写提案。未批准时全部省略 related_entries，词条和操作说明仍可独立实施。

## 实施与验证步骤

1. 明确批准 C01 至 C22 的字段文案和 T01 至 T08 的标签更新。R01 至 R03 单独批准。音频默认不生成；若后续要求音频，需要明确扩大该实施范围。若只批准部分条目，缺少获批先修或端点时暂缓依赖项，不自动扩大范围。
2. 重读本计划、当前 AGENTS.md 和 Git 状态。按稳定 ID 重新检查复用与新 ID 冲突。保存现有词条原始字节、解析值和所有现有媒体哈希。
3. 仅创建获批的 22 个新词条，并为获批复用项追加 codex-cli 标签。保留现有字段、录音与顺序，不借此改写旧例句或补全其他音标。获批的新 ID 在同批创建完成后才用于先修与关联引用。
4. R 组若单独获批，仅按上表方向写一次关系，保留其他成员；不推导额外关联、传递关联或助记。
5. 检查 codex-cli 的直接与间接先修闭包。全部 C/T 获批时共 30 项，6 条新先修与 codex-ai 的既有先修都应包含在内。
6. 运行下列现有命令。仅通过构建生成 dist，不手动编辑导出。

```bash
go test ./...
go vet ./...
go run scripts/validate/main.go
go run scripts/build/main.go
git diff --check
```

7. 新文件另查 UTF-8、LF、尾随空格与字段类型。连续构建两次并比较字节。全部获批时核对 1,664 个稳定 ID 和 17 列导出，不丢失任何旧 ID。
8. 新增词条会改变全局 StudyOrder，因此不要求既有排序编号不变。应核对排序确定性，并检查 codex-cli 筛选后的每项先修都先于依赖项。
9. 全部新条目默认没有 audio 字段，导出音频列为空。现有 1,642 个音频引用和录音字节保持不变。未获批的关联列不应变化；获批时只产生相应直接双向展示。
10. 比较修改允许清单，报告新词、复用标签、可选关联和实际验证结果。更新本计划状态。规划检查不代表词库构建或真实 Codex、Anki 流程已经验证。

不执行任何 Codex 会话操作，不安装工具或语音，不改 Anki 集合，不暂存、提交或推送。本计划无需 schema、CLI、模板、导出器或依赖变更。

## 审阅与规划检查

原审阅分组 (现已全部批准):

- C01 至 C22: 新词完整字段文案、音标、例句、Notes、优先级和先修。
- T01 至 T08: 仅为既有词条追加 codex-cli 标签。
- R01 至 R03: 可选关联，单独确认；不包含 mnemonic 改写。

音频默认不生成。22 个新词的音标均有具体来源，本批不留下待猜写音标。是否需要其他命令词、快捷键或更多状态表达，不属于当前 30 个学习目标的实施范围。

规划检查已通过:

- 22 个 YAML 草案均可解析，字段和值与逐项提案一致。新 ID 全局唯一。
- 8 个复用条目的基线值已核对，既有内容未改写。
- 30 项主题的直接与间接先修闭包完整，无循环。
- 3 对关联的端点和声明方向明确，未实际写入关系。
- 22 个本地链接及锚点、7 张表格的显示宽度和新文件格式检查通过。
- 与规划快照相比，仅新增本计划。1,642 个既有词条、1,685 个既有 MP3、实现、模板、配置和导出均未改变。
- git diff --check 通过；未跟踪文件另查换行与尾随空格。

规划阶段未运行构建、音频生成或实际交互命令。词典和源码查证不代表已经验证本机交互流程、发音质量或真实 Anki 播放。

## 实施与验证记录

用户已明确批准全部项目并要求实施。C01 至 C22、T01 至 T08 和 R01 至 R03 全部完成，失败 0 项。原始提案、例句、音标、来源及关系方向保留为批准记录。

| 检查项目     | 实际结果                                                                       |
|:-------------|:-------------------------------------------------------------------------------|
| 词库规模     | 1,664 项，英语 1,421 项，日语 243 项                                           |
| 新增词条     | 22 个文件，完整字段与获批 YAML 草案一致，另加入获批的 3 条单端关联声明         |
| 复用更新     | 8 个既有词条仅追加 codex-cli 标签，其他字段和文本字节保持不变                  |
| 其他旧词条   | 1,634 个文件与实施前快照逐字节一致                                             |
| 主题独立学习 | codex-cli 共 30 项，全部直接和间接先修均在主题内                               |
| 先修顺序     | 6 条新先修加 1 条既有先修，筛选后的导出顺序全部正确                            |
| 关联展示     | 3 对关系产生 6 个直接邻居展示，全部旧关系也与导出一致                          |
| 既有录音     | 1,642 个音频引用不变；media 与 archives 中全部 1,685 个 MP3 哈希不变           |
| 新词录音     | 按批准范围不生成，22 个新词没有 audio 字段，导出音频列为空                     |
| Go 检查      | go test ./...、go vet ./... 通过                                               |
| 词库验证     | go run scripts/validate/main.go 通过；保留 24 条既有同词警告，无新增同词警告   |
| 导出         | 1,664 行数据、17 列；旧条目除 StudyOrder 和获批的 8 项 Tags 外，其余导出列不变 |
| 确定性       | 两次无源变化构建逐字节一致                                                     |
| 导出 SHA-256 | 2d609d390d8114f88d116dfc067d7fa6a28729e881be9a6f9ab31357b536b7c6               |

通过现有命令重建导出，未手工修改 dist。现有实现、模板、依赖、其他文档和配置保持不变。新文件格式及 git diff --check 均已通过检查。

上述检查验证源数据、主题闭包、排序与导出关系。未验证真实 Codex 交互路径、人工发音质量或 Anki 导入、显示和播放。未执行音频生成、Codex 配置修改、Git 暂存、提交、推送或 Anki 集合操作。
