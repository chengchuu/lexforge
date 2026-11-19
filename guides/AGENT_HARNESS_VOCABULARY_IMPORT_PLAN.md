# 跨工具智能体运行框架词汇导入计划

状态: H01–H22、T01–T18、R01–R06 已全部批准并实施，静态及导出验证通过。规划与实施日期: 2026-11-19。

## 授权与范围

规划阶段只创建本计划。用户选择先整理跨工具通用词汇，不限定产品；随后明确回复 "approve all of them; implement it"，批准本批全部内容、复用更新与关联。以下表格保留规划基线，实际实施结果见文末。实施仅修改获批词条与本计划，并通过现有构建器更新导出；不生成音频，不更改实现、模板、依赖或配置，也不提交或操作 Anki。

依据当前 [AGENTS.md](../AGENTS.md)、[README](../README.md)、[词条规范](ENTRY_SCHEMA.md)、[学习顺序](LEARNING_ORDER.md) 和 [维护提示](MAINTENANCE_PROMPTS.md) 编写。现有 Codex 批次见 [已实施计划](CODEX_CLI_VOCABULARY_IMPORT_PLAN.md)。

| 项目           | 当前基线或拟议范围                                                       |
|:---------------|:-------------------------------------------------------------------------|
| 实际词库       | 1,664 个稳定 ID；英语 1,421 项、日语 243 项                              |
| 新增           | H01–H22，共 22 项英语学习目标                                            |
| 复用更新       | T01–T18，共 18 项，只追加 agent-harness 标签                             |
| 独立主题       | 40 项，包含全部直接和间接先修                                            |
| 新先修声明     | 17 条，分布于 12 个新条目；保留 5 条既有先修边                           |
| 关联候选       | R01–R06，共 6 对，单独批准                                               |
| 已有音频       | 1,642 个活动录音引用；另有 43 个归档 MP3                                 |
| 缺少音频       | 22 个上批新条目尚无 audio；本批默认不生成音频                            |
| 发音草案       | 17 项核实美式 IPA；5 项暂不填写 pronunciation                            |
| 获批实施后预计 | 1,686 项；英语 1,443 项、日语 243 项                                     |
| 记忆辅助字段   | 本批不新增或改写 etymology、word_formation、mnemonic；辨析写入新词 Notes |


规划开始时工作树已有上一批的 8 个修改词条、22 个未跟踪新词条及 Codex 计划。本轮保留这些内容，不能把它们算作本轮新增。已在仓库外记录原文件 SHA-256、词条解析内容和 Git 状态，覆盖 1,685 个活动或归档 MP3；导出是否存在也纳入原文件清单。

本批以稳定 ID 查重，而不是文件名。只选择有独立学习价值的词语，不逐条收集菜单和状态文案。全部英文例句均为编写的学习示例，不是本机操作记录。

## 来源与边界

以下 14 份官方文档已读取与候选相关的正文，访问日期为 2026-11-19。它们说明通用概念的实际用法，不构成所有产品的统一实现规范。未安装、启动或实测 Claude Code、Gemini CLI 或 OpenCode，也未核实本机产品版本与功能开关。文档会更新；产品专有配置和命令不作为学习词头。

| 来源 | 已读文档                                                                                                              | 支持范围与限制                                                                             |
|:-----|:----------------------------------------------------------------------------------------------------------------------|:-------------------------------------------------------------------------------------------|
| S1   | [Visual Studio Code — Understand agent harnesses](https://code.visualstudio.com/docs/agents/concepts/agent-harnesses) | 模型、运行框架、请求、工具、执行环境与审批的职责；采用通用职责，不照搬具体产品菜单。       |
| S2   | [Model Context Protocol — Tools (2025-06-18)](https://modelcontextprotocol.io/specification/2025-06-18/server/tools)  | 工具定义、tools/call 请求与结果返回；固定协议版本，不声称所有工具都使用 MCP。              |
| S3   | [Agent Skills — Agent Skills Overview](https://agentskills.io/home)                                                   | 技能由说明和可选脚本、参考资料等组成，按需加载；不是单个可执行工具的同义词。               |
| S4   | [Claude Code — Hooks reference](https://code.claude.com/docs/en/hooks)                                                | 生命周期事件与配置的处理动作；各工具事件名、处理方式和阻止行为不同。                       |
| S5   | [Claude Code — Create custom subagents](https://code.claude.com/docs/en/sub-agents)                                   | 子智能体、委派、独立上下文与返回结果的一个产品实例；不推导所有工具均有相同隔离或并行机制。 |
| S6   | [Gemini CLI — Checkpointing](https://geminicli.com/docs/cli/checkpointing/)                                           | 检查点与恢复的产品实例；不把某工具的保存范围推广为通用保证。                               |
| S7   | [Gemini CLI — Headless mode reference](https://geminicli.com/docs/cli/headless/)                                      | 无交互终端界面运行、结构化输出和逐条事件输出。                                             |
| S8   | [Gemini CLI — Extensions](https://geminicli.com/docs/extensions/)                                                     | 扩展可以打包提示、工具连接、命令和技能等内容。                                             |
| S9   | [OpenCode — Plugins](https://opencode.ai/docs/plugins/)                                                               | 插件包、局部资源与宿主加载配置；不规定跨产品的插件结构。                                   |
| S10  | [OpenCode — Providers](https://opencode.ai/docs/providers/)                                                           | 选择模型与选择 provider 的区别；服务端或本地后端的配置。                                   |
| S11  | [Claude Code — How Claude remembers your project](https://code.claude.com/docs/en/memory)                             | 跨会话说明、已保存记忆与当前上下文的区别；文本说明不等于强制执行配置。                     |
| S12  | [Claude Code — How Claude Code works](https://code.claude.com/docs/en/how-claude-code-works)                          | 模型与工具、当前上下文以及协调执行的工作循环。                                             |
| S13  | [Gemini CLI — Policy engine](https://geminicli.com/docs/reference/policy-engine/)                                     | 规则决定允许、拒绝或请求确认；不是单纯在提示词中表达愿望。                                 |
| S14  | [Gemini CLI — System Prompt Override](https://geminicli.com/docs/cli/system-prompt/)                                  | system prompt 设置系统行为的实例；不把该工具的覆盖方式推广为通用操作。                     |


技术释义依据上述文档整理；普通词形的读音另查词典，不把词典中的人体、职业或日常义直接当成 AI 机制。MCP 来源固定到 2025-06-18；其他来源未固定版本。配置、窗口容量、权限、检查点恢复、自动记忆和子智能体机制均以实际宿主为准。

## 候选清单

| 编号 | 稳定 ID             | Word           | 拟定 Meaning                                                       | 来源 | 决定 |
|:-----|:--------------------|:---------------|:-------------------------------------------------------------------|:-----|:-----|
| H01  | model-ai            | model          | 模型；根据输入生成回答或工具请求的 AI 系统。                       | S1   | 新增 |
| H02  | tool-tech           | tool           | 工具；可由程序或智能体调用以执行特定操作的能力。                   | S12  | 新增 |
| H03  | result-general      | result         | 结果；某个操作、过程或事件产生的输出或后果。                       | S2   | 新增 |
| H04  | tool-call-ai        | tool call      | 工具调用；请求某个工具按给定参数执行操作。                         | S2   | 新增 |
| H05  | tool-result-ai      | tool result    | 工具结果；工具执行后返回给调用方的内容或错误信息。                 | S2   | 新增 |
| H06  | instruction-tech    | instruction    | 指令；告诉系统或执行者应做什么、如何行动的要求。                   | S11  | 新增 |
| H07  | system-prompt-ai    | system prompt  | 系统提示；用于设定模型角色、行为要求和任务边界的系统层输入。       | S14  | 新增 |
| H08  | context-window-ai   | context window | 上下文窗口；模型一次处理可容纳的上下文范围或容量，通常以词元计量。 | S12  | 新增 |
| H09  | memory-ai           | memory         | 记忆；智能体保存并可在后续交互中重新使用的信息。                   | S11  | 新增 |
| H10  | hook-tech           | hook           | 钩子；在预定生命周期事件发生时触发配置动作的扩展点。               | S4   | 新增 |
| H11  | skill-ai            | skill          | 技能；为特定任务打包的可复用说明、知识和可选资源。                 | S3   | 新增 |
| H12  | extension-tech      | extension      | 扩展；为宿主增加功能或资源的附加组件或资源包。                     | S8   | 新增 |
| H13  | plugin-tech         | plugin         | 插件；由宿主按其扩展机制加载、用于增加功能的组件。                 | S9   | 新增 |
| H14  | provider-ai         | provider       | 提供方；为工具提供模型访问接口的服务或后端。                       | S10  | 新增 |
| H15  | subagent-ai         | subagent       | 子智能体；受委派处理部分任务并返回结果的智能体。                   | S5   | 新增 |
| H16  | delegation-workflow | delegation     | 委派；把部分工作交给其他执行者完成。                               | S5   | 新增 |
| H17  | orchestration-ai    | orchestration  | 协调执行；安排工具、智能体或步骤的执行顺序与结果流转。             | S12  | 新增 |
| H18  | checkpoint-tech     | checkpoint     | 检查点；用于检查进度或支持恢复的阶段性状态记录。                   | S6   | 新增 |
| H19  | streaming-tech      | streaming      | 流式输出或传输；把数据逐段交付，而非等待全部生成后一次返回。       | S7   | 新增 |
| H20  | headless-tech       | headless       | 无交互界面的；不依赖交互式终端或图形界面运行的。                   | S7   | 新增 |
| H21  | configuration-tech  | configuration  | 配置；决定系统行为的设置及其组合。                                 | S9   | 新增 |
| H22  | policy-tech         | policy         | 策略；决定操作是否允许、拒绝或需要确认的执行规则。                 | S13  | 新增 |


H02 tool 与 H03 result 是独立有用的基础项，服务于 tool call / tool result。它们同时出现在已读工具文档中，不是凭空添加来源词。没有为了组成短语另建 system、window 或 sub 卡片。

## 复用决定

以下每项只在原 tags 列表末尾追加一次 agent-harness。保留原 ID、文件路径、Word、Meaning、Notes、发音、录音、学习优先级、先修和已有关联；不以本批 Notes 覆盖旧条目。

| 编号 | 稳定 ID 与源文件                                                   | Word       | 当前 Meaning                                                               | 复用理由                                       |
|:-----|:-------------------------------------------------------------------|:-----------|:---------------------------------------------------------------------------|:-----------------------------------------------|
| T01  | harness-ai ([文件](../entries/harness-ai.yml))                     | harness    | 智能体运行框架；连接模型、上下文和工具，并协调执行循环与会话状态的软件层。 | 复用运行框架义，保留既有测试 harness 关联      |
| T02  | agent-ai ([文件](../entries/agent-ai.yml))                         | agent      | 智能体；围绕目标选择步骤、调用工具，并根据执行结果继续行动的系统。         | 复用智能体义                                   |
| T03  | prompt-ai ([文件](../entries/prompt-ai.yml))                       | prompt     | 提示词；向模型提供的任务要求和相关输入。                                   | 复用提示词义                                   |
| T04  | context-language ([文件](../entries/context-language.yml))         | context    | 语境；上下文；帮助理解含义的背景。                                         | 复用通用上下文基础，保留原语言语境内容         |
| T05  | token-llm ([文件](../entries/token-llm.yml))                       | token      | 词元；语言模型将文本编码后使用的基本处理单位。                             | 复用词元义                                     |
| T06  | session-tech ([文件](../entries/session-tech.yml))                 | session    | 会话；可继续交互并保存相关对话历史的一次交流过程。                         | 复用会话义                                     |
| T07  | permission-tech ([文件](../entries/permission-tech.yml))           | permission | 权限；允许执行某种操作或访问某项资源的许可。                               | 复用权限义                                     |
| T08  | approval-workflow ([文件](../entries/approval-workflow.yml))       | approval   | 批准；对某项提议或操作给予同意。                                           | 复用批准义                                     |
| T09  | sandbox-tech ([文件](../entries/sandbox-tech.yml))                 | sandbox    | 沙箱；限制程序可访问资源和可执行操作的受约束环境。                         | 复用沙箱义并包含权限先修                       |
| T10  | compact-context-tech ([文件](../entries/compact-context-tech.yml)) | compact    | 压缩；将较长上下文整理为更精简的表示。                                     | 复用上下文压缩动词，不另建 compaction 重叠卡片 |
| T11  | summary-language ([文件](../entries/summary-language.yml))         | summary    | 摘要；保留主要内容的简短概述。                                             | 复用摘要基础                                   |
| T12  | snapshot-tech ([文件](../entries/snapshot-tech.yml))               | snapshot   | 快照                                                                       | 复用状态快照义                                 |
| T13  | artifact-workflow ([文件](../entries/artifact-workflow.yml))       | artifact   | 工作成果文件                                                               | 复用工作成果文件义，不混同仅构建产物义         |
| T14  | handoff-workflow ([文件](../entries/handoff-workflow.yml))         | handoff    | 交接                                                                       | 复用交接义                                     |
| T15  | retry-workflow ([文件](../entries/retry-workflow.yml))             | retry      | 重试；在先前尝试失败后再次进行尝试。                                       | 复用重试义                                     |
| T16  | fallback-tech ([文件](../entries/fallback.yml))                    | fallback   | 备用方案；回退方案                                                         | 复用备用方案义                                 |
| T17  | invoke-tech ([文件](../entries/invoke-tech.yml))                   | invoke     | 调用                                                                       | 复用调用动词，不另建 call 同义目标             |
| T18  | invocation-workflow ([文件](../entries/invocation-workflow.yml))   | invocation | 一次调用                                                                   | 复用一次调用义，保留 invoke 先修和已有关联     |


注意 fallback-tech 位于 entries/fallback.yml。context-language 保留其通用语境基础义，不改写成专有 AI 机制；context window 单独解释容量。复用 invoke 与 invocation 后，不再新增同义的 call；tool call 保留为有独立辨识价值的模型交互术语，而 invocation 仍覆盖更广的一次调用或启动过程。

## 新词具体文案

下列 YAML 为逐项审阅草案，不是已创建文件。未列出的字段保持省略，尤其不提前写入 audio 或 related_entries。Context 未单独添加，因为这些草案的 Meaning 和 Notes 已提供足够语境。每个新文件拟为 entries/<id>.yml。

### H01 model

```yaml
id: "model-ai"
language: "en"
word: "model"
pronunciation: "/ˈmɑdəl/"
part_of_speech: "noun"
meaning: "模型；根据输入生成回答或工具请求的 AI 系统。"
example: "Choose a model that fits the task."
notes: |-
  这里学习语言模型的技术含义。模型不等于管理工具执行和会话状态的 harness，也不等于提供模型访问接口的 provider。
  Reference: Visual Studio Code — Understand agent harnesses. https://code.visualstudio.com/docs/agents/concepts/agent-harnesses (accessed 2026-11-19).
  Pronunciation source: Collins Dictionary — model; model in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/model (accent: US; accessed 2026-11-19).
learning_priority: 2
tags:
  - tech
  - agent-harness
  - noun
```

### H02 tool

```yaml
id: "tool-tech"
language: "en"
word: "tool"
pronunciation: "/tul/"
part_of_speech: "noun"
meaning: "工具；可由程序或智能体调用以执行特定操作的能力。"
example: "The agent uses a tool to inspect the files."
notes: |-
  这里学习软件工具义，范围可包含检索、文件操作和命令执行。工具是否可用以及能访问什么，由实际运行环境和权限决定。
  Reference: Claude Code — How Claude Code works. https://code.claude.com/docs/en/how-claude-code-works (accessed 2026-11-19).
  Pronunciation source: Collins Dictionary — tool; tool in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/tool (accent: US; accessed 2026-11-19).
learning_priority: 1
tags:
  - tech
  - agent-harness
  - noun
```

### H03 result

```yaml
id: "result-general"
language: "en"
word: "result"
pronunciation: "/rɪˈzʌlt/"
part_of_speech: "noun"
meaning: "结果；某个操作、过程或事件产生的输出或后果。"
example: "Check the result before continuing."
notes: |-
  result 不保证表示成功；错误信息也可能是一次操作返回的结果。这里保留通用义，不限定某种协议的字段名。
  Reference: Model Context Protocol — Tools (2025-06-18). https://modelcontextprotocol.io/specification/2025-06-18/server/tools (accessed 2026-11-19).
  Pronunciation source: Collins Dictionary — result; result in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/result (accent: US; accessed 2026-11-19).
learning_priority: 1
tags:
  - general
  - agent-harness
  - noun
```

### H04 tool call

```yaml
id: "tool-call-ai"
language: "en"
word: "tool call"
part_of_speech: "noun phrase"
meaning: "工具调用；请求某个工具按给定参数执行操作。"
example: "Inspect the tool call before approving it."
notes: |-
  一次请求通常包含工具名称和参数。模型提出请求不代表操作已经成功执行；具体审批和执行方式由运行框架决定。不是所有工具调用都采用 MCP。
  Reference: Model Context Protocol — Tools (2025-06-18). https://modelcontextprotocol.io/specification/2025-06-18/server/tools (accessed 2026-11-19).
learning_priority: 2
prerequisites:
  - tool-tech
  - invoke-tech
tags:
  - tech
  - agent-harness
  - phrase
```

### H05 tool result

```yaml
id: "tool-result-ai"
language: "en"
word: "tool result"
part_of_speech: "noun phrase"
meaning: "工具结果；工具执行后返回给调用方的内容或错误信息。"
example: "The tool result contains an error message."
notes: |-
  区分请求执行的 tool call 与执行后返回的 tool result。结果可以包含文本或其他结构化内容；收到结果不等于操作成功。
  Reference: Model Context Protocol — Tools (2025-06-18). https://modelcontextprotocol.io/specification/2025-06-18/server/tools (accessed 2026-11-19).
learning_priority: 2
prerequisites:
  - tool-call-ai
  - result-general
tags:
  - tech
  - agent-harness
  - phrase
```

### H06 instruction

```yaml
id: "instruction-tech"
language: "en"
word: "instruction"
pronunciation: "/ɪnˈstrʌkʃən/"
part_of_speech: "noun"
meaning: "指令；告诉系统或执行者应做什么、如何行动的要求。"
example: "Follow the instruction to preserve existing files."
notes: |-
  这里侧重自然语言任务要求，而不是处理器机器指令。常见复数 instructions 表示一组要求；本条用单数词头。文本指令与实际执行权限不同。
  Reference: Claude Code — How Claude remembers your project. https://code.claude.com/docs/en/memory (accessed 2026-11-19).
  Pronunciation source: Collins Dictionary — instruction; instruction in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/instruction (accent: US; accessed 2026-11-19).
learning_priority: 1
tags:
  - tech
  - agent-harness
  - noun
```

### H07 system prompt

```yaml
id: "system-prompt-ai"
language: "en"
word: "system prompt"
part_of_speech: "noun phrase"
meaning: "系统提示；用于设定模型角色、行为要求和任务边界的系统层输入。"
example: "The system prompt defines the assistant’s role."
notes: |-
  具体消息层级和可配置方式依工具及模型接口而异。系统提示不是操作系统权限或沙箱，不能只靠提示文本保证操作一定被阻止。
  Reference: Gemini CLI — System Prompt Override. https://geminicli.com/docs/cli/system-prompt/ (accessed 2026-11-19).
learning_priority: 2
prerequisites:
  - prompt-ai
  - instruction-tech
tags:
  - tech
  - agent-harness
  - phrase
```

### H08 context window

```yaml
id: "context-window-ai"
language: "en"
word: "context window"
part_of_speech: "noun phrase"
meaning: "上下文窗口；模型一次处理可容纳的上下文范围或容量，通常以词元计量。"
example: "The conversation is close to the context window limit."
notes: |-
  当前上下文可能包括指令、对话和工具结果。保存的完整历史或记忆文件不一定全部进入当前窗口；实际可用容量及计算方式取决于模型和工具。
  Reference: Claude Code — How Claude Code works. https://code.claude.com/docs/en/how-claude-code-works (accessed 2026-11-19).
learning_priority: 2
prerequisites:
  - context-language
  - token-llm
tags:
  - tech
  - agent-harness
  - phrase
```

### H09 memory

```yaml
id: "memory-ai"
language: "en"
word: "memory"
pronunciation: "/ˈmɛməri/"
part_of_speech: "noun"
meaning: "记忆；智能体保存并可在后续交互中重新使用的信息。"
example: "Review the saved memory before reusing it."
notes: |-
  这里是智能体的信息保留机制，不是故事中的个人回忆。它不保证逐字保存全部对话，也不保证内容永远正确；保存、加载和更新方式依工具而异。
  Reference: Claude Code — How Claude remembers your project. https://code.claude.com/docs/en/memory (accessed 2026-11-19).
  Pronunciation source: Collins Dictionary — memory; memory in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/memory (accent: US; accessed 2026-11-19).
learning_priority: 2
prerequisites:
  - session-tech
tags:
  - tech
  - agent-harness
  - noun
```

### H10 hook

```yaml
id: "hook-tech"
language: "en"
word: "hook"
pronunciation: "/hʊk/"
part_of_speech: "noun"
meaning: "钩子；在预定生命周期事件发生时触发配置动作的扩展点。"
example: "Run the formatter from a hook after an edit."
notes: |-
  不同工具支持的事件和处理动作不同，不一定只执行 shell 命令。此义与 React Hook 不同，不能套用 React 的调用规则；一个 hook 是否能阻止操作要看宿主约定。
  Reference: Claude Code — Hooks reference. https://code.claude.com/docs/en/hooks (accessed 2026-11-19).
  Pronunciation source: Collins Dictionary — hook; hook in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/hook (accent: US; accessed 2026-11-19).
learning_priority: 2
tags:
  - tech
  - agent-harness
  - noun
```

### H11 skill

```yaml
id: "skill-ai"
language: "en"
word: "skill"
pronunciation: "/skɪl/"
part_of_speech: "noun"
meaning: "技能；为特定任务打包的可复用说明、知识和可选资源。"
example: "Load the skill when its workflow fits the task."
notes: |-
  在 Agent Skills 格式中，核心是 SKILL.md；还可包含脚本和参考材料。技能指导如何完成任务，不等于工具本身，也不保证加载后每个脚本都会执行。
  Reference: Agent Skills — Agent Skills Overview. https://agentskills.io/home (accessed 2026-11-19).
  Pronunciation source: Collins Dictionary — skill; skill in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/skill (accent: US; accessed 2026-11-19).
learning_priority: 2
prerequisites:
  - instruction-tech
tags:
  - tech
  - agent-harness
  - noun
```

### H12 extension

```yaml
id: "extension-tech"
language: "en"
word: "extension"
pronunciation: "/ɛkˈstɛnʃən/"
part_of_speech: "noun"
meaning: "扩展；为宿主增加功能或资源的附加组件或资源包。"
example: "Disable the extension if it is no longer needed."
notes: |-
  这里学习软件扩展义，不是文件扩展名。不同产品可能把提示、工具连接或技能等打包为扩展；extension 与 plugin 的边界依宿主约定，不能推定固定包含关系。
  Reference: Gemini CLI — Extensions. https://geminicli.com/docs/extensions/ (accessed 2026-11-19).
  Pronunciation source: Collins Dictionary — extension; extension in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/extension (accent: US; accessed 2026-11-19).
learning_priority: 2
tags:
  - tech
  - agent-harness
  - noun
```

### H13 plugin

```yaml
id: "plugin-tech"
language: "en"
word: "plugin"
pronunciation: "/ˈplʌɡˌɪn/"
part_of_speech: "noun"
meaning: "插件；由宿主按其扩展机制加载、用于增加功能的组件。"
example: "Check the plugin’s compatibility before loading it."
notes: |-
  plugin 也写作 plug-in；本条以软件界面常见的 plugin 为词头。实际加载形式依宿主而异，不必是单个脚本文件。已有 JavaScript plugin 条目侧重 Bootstrap，本条学习通用插件概念。
  Reference: OpenCode — Plugins. https://opencode.ai/docs/plugins/ (accessed 2026-11-19).
  Pronunciation source: Collins Dictionary — plug-in; plug-in in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/plug-in (accent: US; accessed 2026-11-19).
learning_priority: 2
tags:
  - tech
  - agent-harness
  - noun
```

### H14 provider

```yaml
id: "provider-ai"
language: "en"
word: "provider"
pronunciation: "/prəˈvaɪdər/"
part_of_speech: "noun"
meaning: "提供方；为工具提供模型访问接口的服务或后端。"
example: "Select a provider, then choose an available model."
notes: |-
  这里按模型接入语境理解；provider 不一定指模型开发公司，也可能接入本地运行后端。切换 provider 与切换模型、运行框架是不同选择。
  Reference: OpenCode — Providers. https://opencode.ai/docs/providers/ (accessed 2026-11-19).
  Pronunciation source: Collins Dictionary — provider; provider in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/provider (accent: US; accessed 2026-11-19).
learning_priority: 2
prerequisites:
  - model-ai
tags:
  - tech
  - agent-harness
  - noun
```

### H15 subagent

```yaml
id: "subagent-ai"
language: "en"
word: "subagent"
pronunciation: "/sʌbˈeɪdʒənt/"
part_of_speech: "noun"
meaning: "子智能体；受委派处理部分任务并返回结果的智能体。"
example: "Ask the subagent to review one module."
notes: |-
  各工具的上下文继承、工具权限和通信机制不同。子智能体不必同时运行，也不自动拥有主智能体的全部权限；本条例句是学习示例，不授权实际委派。
  Reference: Claude Code — Create custom subagents. https://code.claude.com/docs/en/sub-agents (accessed 2026-11-19).
  Pronunciation source: Collins Dictionary — subagent; subagent in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/subagent (accent: US; accessed 2026-11-19).
learning_priority: 2
prerequisites:
  - agent-ai
tags:
  - tech
  - agent-harness
  - noun
```

### H16 delegation

```yaml
id: "delegation-workflow"
language: "en"
word: "delegation"
pronunciation: "/ˌdɛləˈɡeɪʃən/"
part_of_speech: "noun"
meaning: "委派；把部分工作交给其他执行者完成。"
example: "Use delegation for a focused review task."
notes: |-
  任务委派应说明范围、预期结果和约束。它不同于 DOM 的 event delegation；也不必意味着交接整个任务或启动并行执行。
  Reference: Claude Code — Create custom subagents. https://code.claude.com/docs/en/sub-agents (accessed 2026-11-19).
  Pronunciation source: Collins Dictionary — delegation; delegation in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/delegation (accent: US; accessed 2026-11-19).
learning_priority: 2
prerequisites:
  - agent-ai
tags:
  - tech
  - agent-harness
  - noun
```

### H17 orchestration

```yaml
id: "orchestration-ai"
language: "en"
word: "orchestration"
part_of_speech: "noun"
meaning: "协调执行；安排工具、智能体或步骤的执行顺序与结果流转。"
example: "The harness handles orchestration across tools and agents."
notes: |-
  这里学习技术中的协调义，不是音乐配器。协调可以是顺序执行，也可以涉及并行；不能因出现 orchestration 就推定系统使用了多个智能体。
  Reference: Claude Code — How Claude Code works. https://code.claude.com/docs/en/how-claude-code-works (accessed 2026-11-19).
learning_priority: 3
prerequisites:
  - delegation-workflow
  - tool-call-ai
tags:
  - tech
  - agent-harness
  - noun
```

### H18 checkpoint

```yaml
id: "checkpoint-tech"
language: "en"
word: "checkpoint"
pronunciation: "/ˈtʃɛkˌpɔɪnt/"
part_of_speech: "noun"
meaning: "检查点；用于检查进度或支持恢复的阶段性状态记录。"
example: "Create a checkpoint before applying the changes."
notes: |-
  保存和恢复哪些内容依工具而异。检查点不一定是项目 Git 提交，也不保证撤销网络请求、外部服务操作或所有文件变更；使用前应核对范围。
  Reference: Gemini CLI — Checkpointing. https://geminicli.com/docs/cli/checkpointing/ (accessed 2026-11-19).
  Pronunciation source: Collins Dictionary — checkpoint; checkpoint in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/checkpoint (accent: US; accessed 2026-11-19).
learning_priority: 2
prerequisites:
  - snapshot-tech
tags:
  - tech
  - agent-harness
  - noun
```

### H19 streaming

```yaml
id: "streaming-tech"
language: "en"
word: "streaming"
pronunciation: "/ˈstrimɪŋ/"
part_of_speech: "noun"
meaning: "流式输出或传输；把数据逐段交付，而非等待全部生成后一次返回。"
example: "Streaming lets the client process events as they arrive."
notes: |-
  输出可能包含消息片段、工具事件和最终结果。已显示部分内容不表示任务已完成；实际事件格式和结束标志依接口而异。
  Reference: Gemini CLI — Headless mode reference. https://geminicli.com/docs/cli/headless/ (accessed 2026-11-19).
  Pronunciation source: Collins Dictionary — streaming; streaming in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/streaming (accent: US; accessed 2026-11-19).
learning_priority: 2
tags:
  - tech
  - agent-harness
  - noun
```

### H20 headless

```yaml
id: "headless-tech"
language: "en"
word: "headless"
pronunciation: "/ˈhɛdlɪs/"
part_of_speech: "adjective"
meaning: "无交互界面的；不依赖交互式终端或图形界面运行的。"
example: "Run the agent in headless mode for automation."
notes: |-
  这里学习软件运行模式义，常用于脚本或流水线。headless 不等于没有输出，也不代表自动批准操作或取消权限限制。
  Reference: Gemini CLI — Headless mode reference. https://geminicli.com/docs/cli/headless/ (accessed 2026-11-19).
  Pronunciation source: Collins Dictionary — headless; headless in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/headless (accent: US; accessed 2026-11-19).
learning_priority: 2
tags:
  - tech
  - agent-harness
  - adjective
```

### H21 configuration

```yaml
id: "configuration-tech"
language: "en"
word: "configuration"
pronunciation: "/kənˌfɪɡjəˈreɪʃən/"
part_of_speech: "noun"
meaning: "配置；决定系统行为的设置及其组合。"
example: "Review the configuration before starting the session."
notes: |-
  配置可能控制模型选择、插件和权限等。配置文件、命令行参数等如何合并和覆盖依工具而异，不把单个产品的优先级当作通用规则。
  Reference: OpenCode — Plugins. https://opencode.ai/docs/plugins/ (accessed 2026-11-19).
  Pronunciation source: Collins Dictionary — configuration; configuration in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/configuration (accent: US; accessed 2026-11-19).
learning_priority: 1
tags:
  - tech
  - agent-harness
  - noun
```

### H22 policy

```yaml
id: "policy-tech"
language: "en"
word: "policy"
pronunciation: "/ˈpɑləsi/"
part_of_speech: "noun"
meaning: "策略；决定操作是否允许、拒绝或需要确认的执行规则。"
example: "The policy requires approval for this operation."
notes: |-
  此处不是保险单。策略规则、某次审批结果和自然语言指令是不同概念；具体规则是否生效、在哪一层执行，应查实际工具实现。
  Reference: Gemini CLI — Policy engine. https://geminicli.com/docs/reference/policy-engine/ (accessed 2026-11-19).
  Pronunciation source: Collins Dictionary — policy; policy in American English; Webster’s New World College Dictionary, 5th Digital Edition. https://www.collinsdictionary.com/dictionary/english/policy (accent: US; accessed 2026-11-19).
learning_priority: 2
prerequisites:
  - permission-tech
tags:
  - tech
  - agent-harness
  - noun
```

## 发音核实与暂缓项

17 项均使用 Collins 的 American English 栏、Webster’s New World College Dictionary 转写，选择该词形的一个完整变体，不拼接词根或口音。全为美式，故不加 US 前缀。plugin 采用同一软件词形的词典拼写 plug-in；Notes 明示对应。技术义另由官方文档支持。

| 编号 | Word          | pronunciation      | 词典来源                                                                                     |
|:-----|:--------------|:-------------------|:---------------------------------------------------------------------------------------------|
| H01  | model         | /ˈmɑdəl/           | [Collins: model](https://www.collinsdictionary.com/dictionary/english/model)                 |
| H02  | tool          | /tul/              | [Collins: tool](https://www.collinsdictionary.com/dictionary/english/tool)                   |
| H03  | result        | /rɪˈzʌlt/          | [Collins: result](https://www.collinsdictionary.com/dictionary/english/result)               |
| H06  | instruction   | /ɪnˈstrʌkʃən/      | [Collins: instruction](https://www.collinsdictionary.com/dictionary/english/instruction)     |
| H09  | memory        | /ˈmɛməri/          | [Collins: memory](https://www.collinsdictionary.com/dictionary/english/memory)               |
| H10  | hook          | /hʊk/              | [Collins: hook](https://www.collinsdictionary.com/dictionary/english/hook)                   |
| H11  | skill         | /skɪl/             | [Collins: skill](https://www.collinsdictionary.com/dictionary/english/skill)                 |
| H12  | extension     | /ɛkˈstɛnʃən/       | [Collins: extension](https://www.collinsdictionary.com/dictionary/english/extension)         |
| H13  | plugin        | /ˈplʌɡˌɪn/         | [Collins: plug-in](https://www.collinsdictionary.com/dictionary/english/plug-in)             |
| H14  | provider      | /prəˈvaɪdər/       | [Collins: provider](https://www.collinsdictionary.com/dictionary/english/provider)           |
| H15  | subagent      | /sʌbˈeɪdʒənt/      | [Collins: subagent](https://www.collinsdictionary.com/dictionary/english/subagent)           |
| H16  | delegation    | /ˌdɛləˈɡeɪʃən/     | [Collins: delegation](https://www.collinsdictionary.com/dictionary/english/delegation)       |
| H18  | checkpoint    | /ˈtʃɛkˌpɔɪnt/      | [Collins: checkpoint](https://www.collinsdictionary.com/dictionary/english/checkpoint)       |
| H19  | streaming     | /ˈstrimɪŋ/         | [Collins: streaming](https://www.collinsdictionary.com/dictionary/english/streaming)         |
| H20  | headless      | /ˈhɛdlɪs/          | [Collins: headless](https://www.collinsdictionary.com/dictionary/english/headless)           |
| H21  | configuration | /kənˌfɪɡjəˈreɪʃən/ | [Collins: configuration](https://www.collinsdictionary.com/dictionary/english/configuration) |
| H22  | policy        | /ˈpɑləsi/          | [Collins: policy](https://www.collinsdictionary.com/dictionary/english/policy)               |


| 编号 | 暂不填写的词   | 查证结果与后续边界                                                                                              |
|:-----|:---------------|:----------------------------------------------------------------------------------------------------------------|
| H04  | tool call      | 已尝试 Collins 完整词页，未取得可核实的整词/整短语美式或英式 IPA；不拼接单词读音。                              |
| H05  | tool result    | 已尝试 Collins 完整词页，未取得可核实的整词/整短语美式或英式 IPA；不拼接单词读音。                              |
| H07  | system prompt  | 已尝试 Collins 完整词页，未取得可核实的整词/整短语美式或英式 IPA；不拼接单词读音。                              |
| H08  | context window | 已尝试 Collins 完整词页，未取得可核实的整词/整短语美式或英式 IPA；不拼接单词读音。                              |
| H17  | orchestration  | Collins 仅取得未明确标注口音的 COBUILD 显示；Cambridge 和 Oxford 请求失败。暂不将其认作已核实美式或改写为英式。 |


初查 Cambridge 的 model、tool、call、result 发音页也返回 403，随后改查 Collins 成功。访问失败不证明某种口音不存在。5 项缺少可靠 IPA 不阻止审阅含义，但不能把它们标为发音已完成；后续如要补具体音标，需提供可靠证据与审阅文本。未试听词典音频或既有 MP3。

## 分组与排除决定

- 保留 tool call、tool result、system prompt 和 context window 这些有独立技术含义的完整术语，不机械拆成所有组合。
- 复用 compact 动词及其既有摘要先修，不另增 compaction；普通完整句子只作例句。
- hook-react、memory-story、event-delegation-dom、javascript-plugin-bootstrap、same-origin-policy-web 是需辨识的既有专门义，不能直接改写为本批通用概念。
- 本批不收集产品名、命令标志、协议字段名、配置文件名和安装步骤；MCP 只作来源背景，未新增缩写卡片。agent loop、routing、telemetry、worktree 等可在后批按实际需要选择，本批不宣称已穷尽所有 harness 词汇。
- 不补未经查证的词源或随意联想，不给复用项改写助记。新词 Notes 的辨析属于本批具体文案；关联批准不另外授权修改它们。

计数以本批审阅目标为单位。来源不是用户提供的固定单词表，因此不存在固定源表行数。实际读取 14 份产品/标准文档，另核实 17 个词形的美式 IPA；选出 40 个唯一学习目标，其中 22 个新项、18 个复用更新，复用与更新在此为同一组，不相加。tool 和 result 两个基础组件包含在 22 个新项内。17 个发音提案、6 对关联与 40 个学习目标也不是互斥集合。

## 学习顺序与主题闭包

所有新项与上述复用项使用 agent-harness 主题标签；它不是对某个产品专属行为的断言。优先级按用途和复杂度设置，先修保证组件先介绍，不保证相邻或按掌握程度解锁。

| 新目标              | 优先级 | 拟议直接先修                      |
|:--------------------|:-------|:----------------------------------|
| model-ai            | 2      | 无                                |
| tool-tech           | 1      | 无                                |
| result-general      | 1      | 无                                |
| tool-call-ai        | 2      | tool-tech, invoke-tech            |
| tool-result-ai      | 2      | tool-call-ai, result-general      |
| instruction-tech    | 1      | 无                                |
| system-prompt-ai    | 2      | prompt-ai, instruction-tech       |
| context-window-ai   | 2      | context-language, token-llm       |
| memory-ai           | 2      | session-tech                      |
| hook-tech           | 2      | 无                                |
| skill-ai            | 2      | instruction-tech                  |
| extension-tech      | 2      | 无                                |
| plugin-tech         | 2      | 无                                |
| provider-ai         | 2      | model-ai                          |
| subagent-ai         | 2      | agent-ai                          |
| delegation-workflow | 2      | agent-ai                          |
| orchestration-ai    | 3      | delegation-workflow, tool-call-ai |
| checkpoint-tech     | 2      | snapshot-tech                     |
| streaming-tech      | 2      | 无                                |
| headless-tech       | 2      | 无                                |
| configuration-tech  | 1      | 无                                |
| policy-tech         | 2      | permission-tech                   |


保留既有边: harness-ai → agent-ai；sandbox-tech → permission-tech；compact-context-tech → context-language、summary-language；invocation-workflow → invoke-tech。这 5 条加 17 条新声明构成主题内 22 条直接先修边。40 项主题集对全部直接和间接先修闭合。

## 关联候选单独批准

R01–R06 已由用户本次全部批准，按下表方向实施。下表箭头只表示拟在左端 related_entries 中声明右端 ID，不表示显示单向或前置依赖。右端既有 YAML 无需为反向展示增加声明；保留双方全部既有列表成员。没有附带 mnemonic 改写。

| 编号 | 拟声明端 ID / Word / Meaning                                                    | 目标 ID / Word / Meaning                                                                          | 学习价值                             | 混淆风险                                                 |
|:-----|:--------------------------------------------------------------------------------|:--------------------------------------------------------------------------------------------------|:-------------------------------------|:---------------------------------------------------------|
| R01  | model-ai / model / 模型；根据输入生成回答或工具请求的 AI 系统。                 | harness-ai / harness / 智能体运行框架；连接模型、上下文和工具，并协调执行循环与会话状态的软件层。 | 区分模型和组织执行的运行框架         | 更换模型不等于更换 harness。                             |
| R02  | skill-ai / skill / 技能；为特定任务打包的可复用说明、知识和可选资源。           | instruction-tech / instruction / 指令；告诉系统或执行者应做什么、如何行动的要求。                 | 联系可复用的任务说明包与组成它的指令 | 技能不等于单条指令，也不等于可执行工具。                 |
| R03  | hook-tech / hook / 钩子；在预定生命周期事件发生时触发配置动作的扩展点。         | hook-react / Hook / Hook，在 React 组件或自定义 Hook 中按规则使用的函数，可接入状态等能力         | 区分同形技术词的两个义项             | 生命周期钩子与 React Hook 的语义、调用规则不同。         |
| R04  | checkpoint-tech / checkpoint / 检查点；用于检查进度或支持恢复的阶段性状态记录。 | snapshot-tech / snapshot / 快照                                                                   | 区分检查点的用途和状态快照的概念     | 快照不一定可恢复；检查点也不保证包含所有外部状态。       |
| R05  | extension-tech / extension / 扩展；为宿主增加功能或资源的附加组件或资源包。     | plugin-tech / plugin / 插件；由宿主按其扩展机制加载、用于增加功能的组件。                         | 辨识宿主常用的两种扩展称呼           | 不能据此断言二者等价或有固定层级。                       |
| R06  | delegation-workflow / delegation / 委派；把部分工作交给其他执行者完成。         | handoff-workflow / handoff / 交接                                                                 | 区分委派部分工作和交接工作状态或责任 | 委派后原执行者可能仍负责统筹；交接也不一定涉及子智能体。 |


6 对均涉及至少一个新 ID，不重复既有双向直接关系。若全部获批，增加 6 条单端声明，涉及 6 个声明端、12 个不同端点，派生 12 条直接邻接展示 (其中 6 条为反向展示)。4 个既有端点为 harness-ai、hook-react、snapshot-tech、handoff-workflow；hook-react 位于本批 40 项主题范围之外，仅导出反向文字发生变化，不自动追加主题标签、先修或改写说明。其原先修 function-component-react 不因关联进入本批主题。

## 音频范围

本批尚未请求音频，默认不生成、不重用附着、不覆盖既有录音。若按当前范围实施且不另批音频，缺少 audio 的条目会由 22 项增至 44 项；这与 5 项缺 IPA 是不同计数。当前活动引用和归档 MP3 字节均须保持不变。

如果后续请求生成，另按现有本地生成器检查 macOS、say、lame、mpg123 和匹配语种声音；英语从 word 生成，不从 IPA 生成。缺少 IPA 不等于缺少声音。不得自动安装、切换远程服务或再生有效录音。

## 批准范围与实施程序

- 内容组 (已批准并实施): H01–H22 的全部字段、具体例句与 Notes，含 17 个 IPA、17 条先修；接受所列 5 项 IPA 暂缺。
- 复用组 (已批准并实施): T01–T18 仅追加 agent-harness 标签，保留旧语境和内容。
- 关联组 (已明确全部批准并实施): R01–R06 按单端声明实现双向展示；本次批准包含关联组，不附带额外助记改写。
- 音频: 默认不生成；本批不执行安装或产品配置。

创建本计划不授予实施权限。部分批准时，只实施被批准且引用目标已存在或同时获批的项目，缺少先修或关联端点则暂缓依赖项并报告，不自行扩大范围。

明确批准后的步骤:

1. 重读计划与当前代码，检查 Git 状态，核对稳定 ID、词义和字段变化，重新快照词条原文和活动/归档媒体哈希。仓库变化后需重新计算预计总数。
2. 只创建获批的新 YAML；复用项仅追加获批主题标签；只声明单独批准的关联。音频、旧字段、旧列表成员、其他工作不变。
3. 运行以下现有命令；不手改 dist，不更改 schema、CLI、模板、依赖，也不修改 Anki。

```bash
go test ./...
go vet ./...
go run scripts/validate/main.go
go run scripts/build/main.go
git diff --check
```

4. 不改变源文件再构建一次，比较两次导出字节。按当前基线全批应为 1,686 个 17 列数据行，ID 增量恰为 22。检查 40 个 agent-harness 行及先修闭包、顺序和全部获批双向直接展示；全局 StudyOrder 秩可因新增项变化。未批准关联不得出现。
5. 精确比较新词字段与获批文本、复用文件字段白名单及旧记录；验证所有既有录音哈希不变。保留已有关联，记录合法重复词警告，如 hook、memory；不自动合并不同义项。
6. 另查未跟踪文本格式与本地链接。更新本计划真实批准、实施、校验结果；静态字段和导出检查不等于真人发音听审或真实 Anki 导入、显示、播放。完成后报告变更文件及剩余待核实项，不提交或推送。

## 规划检查记录

2026-11-19 已完成以下规划检查:

- 22 个新 ID 与文件名均无冲突；18 个复用 ID 及其实际文件路径存在。
- 22 段 YAML 草案经 Ruby Psych 解析成功；字段属于当前实现范围。
- 40 项主题集包含全部直接和间接先修；17 条新边与 5 条既有边共 22 条，无环；12 个新条目声明先修。
- 6 对关联无重复，端点存在或明确属于本批新词草案；共 12 个不同端点，未将既有反向展示误作缺失关系。
- 24 个本地 Markdown 链接目标存在。本计划无本地锚点链接；外部来源按正文访问记录核对，访问失败的词典页未作为已读证据。
- 新计划单独检查行尾、制表符、换行、表格显示宽度及代码围栏；git diff --check 通过。
- 与规划快照比较，原有 3,397 个文件未改变，包括全部词条、上一批未提交工作、实现、模板、依赖、导出和 1,685 个 MP3。仓库唯一新增文件为本计划。

本轮未运行 Go 测试、vet、词库验证器、构建或音频生成；未做真人发音听审或真实 Anki 验证。以上为计划静态检查，不代表候选已实施。

## 实施与验证记录

2026-11-19 用户明确批准本计划全部项目并要求实施。实施前重新读取计划和当前实现，核对 1,664 个稳定 ID 及原字段，确认规划后词条未变化，并在仓库外保存全部词条原文、Git 状态及既有媒体哈希。

- 新增 22 个文件: entries/<id>.yml，ID 与 H01–H22 完全一致。逐字段核对批准的 Meaning、例句、Notes、17 个美式 IPA、优先级和 17 条先修声明，除已批准关联外没有额外字段。
- 更新 18 个既有文件: T01–T18 表所列实际路径，仅在 tags 末尾追加 agent-harness。与实施前原文比较，其他字节和解析字段保持不变，包括上一批已存在的未提交工作。
- 写入 R01–R06 的 6 条单端关联声明，均位于新文件，保留所有既有关联。全库导出逐项核对直接邻居、排序和双向展示，6 对新增关系派生 12 条邻接显示。
- 当前总数为 1,686 项: 英语 1,443 项、日语 243 项。agent-harness 筛选恰含 40 项；22 条主题内先修边均闭合，全部先修在导出中先于依赖项。全库先修顺序亦已核对。
- 1,642 个既有 audio 引用保持不变；当前 44 项无 audio，包含上一批 22 项和本批 22 项。本批未生成、附着、解码或试听音频，1,685 个活动与归档 MP3 哈希均与实施前一致。
- 17 个 IPA 已按批准文本加入；tool call、tool result、system prompt、context window、orchestration 的 IPA 仍按批准范围暂缺。未重新独立查证词典发音或改变既有录音。
- go test ./...、go vet ./...、go run scripts/validate/main.go、go run scripts/build/main.go 和 git diff --check 均通过。未改动实现，因此未新增只镜像数据内容的回归测试；另用批准字段、原始字节和完整导出做了专项核对。
- 两次构建字节完全相同，均为 1,686 个 17 列数据行。SHA-256: 999b16ccc4fe92fe6f8232ab4934d27ac9d80e201b25538c438dafe6f4c6e31e。
- 与本批实施前的完整构建比较，既有行只出现预期的 StudyOrder 重排、18 项 Tags 更新及 4 个既有端点的 RelatedWords 反向展示；其他列保持不变。所有新项的音频列为空。
- 验证和构建报告 26 条重复词警告；原基线为 24 条，新增 hook、memory 两条均为已审核的不同义项，保留独立稳定 ID。没有自动合并词条。
- 本计划 24 个本地链接目标存在；新建与修改文本另外检查格式，覆盖 git diff --check 不检查的未跟踪文件。实现、README 模板、依赖、其他指南及与本批无关的源文件均未改变。

以上为本地测试、字段及导出检查。未验证真实 Anki 导入、显示、播放，也未做真人发音听审。未安装工具，未提交、推送、复制媒体或修改 Anki 集合。全部获批内容已完成；后续音频生成及 5 项 IPA 补充需另行提出。
