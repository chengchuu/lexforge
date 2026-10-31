# 面试知识词条导入计划

状态: 已整理候选，尚未执行导入。

记录日期: 2026-10-31。

## 目标与范围

以英文单词、词组或术语作为记忆锚点。Meaning 用中文概括核心含义或原理。Notes 承载步骤、边界、辨析和来源。

本计划保存已讨论的 57 个候选点，分为六类。候选数量不等于新增词条数量。导入前仍需逐项核对现有语义和稳定 ID。

本次授权仅创建计划文件。新增或更新词条、生成音频和写入关联均未执行。后续明确批准导入后，再按批准范围实施。

## 来源

- 原始材料: `/Users/mazey/Web/learning/Interviews/AprilSkill.md`。
- 补充材料: `/Users/mazey/Web/learning/Interviews/March.md`。

March 的六个主题在 AprilSkill 中也有涉及，应合并整理。重复出现的问题不创建重复卡片。原文作为选题来源，不作为未经核实的技术结论。

本文保留候选选题，不覆盖原文每道题。源文件位于仓库外，其他机器需自行提供原文。实施记录应补充对应章节、最终 ID 和采用的核实来源。

## 卡片组织

遵循 [Entry schema](ENTRY_SCHEMA.md) 和 [Add vocabulary with Codex](ADD_VOCABULARY_WITH_CODEX.md)。使用现有字段，不增加问答字段或生成基础设施。

- `language`: 使用 `en`；`word` 保留英文概念名称。
- `meaning`: 支持快速回忆，提供自足的中文释义或简短原理。
- `context`、`example`: 仅在有用时提供；例句使用英语。
- `notes`: 保留详细思路、边界、来源和中文演示。不设置长度上限。
- `etymology`、`word_formation`、`mnemonic`: 按需提供，区分事实与助记。
- `pronunciation`: 仅填可靠 IPA；不根据拼写猜测。
- `audio`: 仅在明确要求生成且文件存在后添加。

Notes 使用纯文本段落或步骤。不要依赖 Markdown 渲染。中文推演、符号过程和代码说明不放入 Example 供 TTS 朗读。

一个词条聚焦一个学习目标。已有词汇适合例句时自然复用，不强求覆盖。框架专属含义要明确适用环境。

## 候选清单

以下 Meaning 是内容方向，不是已核实的最终导入文本。Notes 列列出后续展开范围。实现前应核实技术细节及版本边界。

### JavaScript 与内存

| 编号 | 英文锚点                 | Meaning 方向                               | Notes 展开范围                        |
|:-----|:-------------------------|:-------------------------------------------|:--------------------------------------|
| 1    | call stack               | 调用栈；记录当前函数调用及返回位置         | 调用帧、递归和栈溢出                  |
| 2    | heap memory              | 堆内存；运行时管理动态分配数据的内存区域   | 区分概念模型与引擎实现                |
| 3    | garbage collection       | 垃圾回收；回收不再可达的对象所占内存       | 可达性、引用和内存泄漏                |
| 4    | scope                    | 作用域；名称在程序中可被访问的范围         | 复用 scope-lexical，补充词法作用域    |
| 5    | closure                  | 闭包；函数及其可访问的外部词法环境         | 复用 closure-tech，补充变量保留与共享 |
| 6    | prototype chain          | 原型链；沿原型关系查找对象属性             | 自有属性、继承和查找终点              |
| 7    | hoisting                 | 提升；执行前建立声明绑定的行为             | 区分 var、函数声明和词法声明          |
| 8    | temporal dead zone       | 暂时性死区；词法绑定初始化前不可访问的阶段 | 作用域起点、初始化和 ReferenceError   |
| 9    | higher-order function    | 高阶函数；接收函数或返回函数的函数         | 回调、组合和职责分离                  |
| 10   | shallow copy             | 浅拷贝；复制外层结构并保留嵌套引用         | 共享引用带来的修改影响                |
| 11   | deep copy                | 深拷贝；复制嵌套数据以隔离相应引用         | 支持类型、循环引用和复制边界          |
| 12   | floating-point precision | 浮点精度；有限位数表示带来的精度限制       | 二进制小数、舍入和比较                |

### 异步与事件

| 编号 | 英文锚点         | Meaning 方向                           | Notes 展开范围                           |
|:-----|:-----------------|:---------------------------------------|:-----------------------------------------|
| 13   | event loop       | 事件循环；协调任务执行的运行机制       | 区分浏览器与 Node.js 环境                |
| 14   | microtask        | 微任务；在微任务检查点处理的任务       | Promise 回调与任务队列的区别             |
| 15   | Promise          | 表示异步操作最终结果的对象             | 状态、链式调用和错误传播                 |
| 16   | async/await      | 组织基于 Promise 的异步控制流的语法    | 暂停当前异步函数，不阻塞整个线程         |
| 17   | Promise.all      | 聚合多个结果；全部成功时按输入顺序返回 | 空输入、提前拒绝和未完成任务             |
| 18   | event capturing  | 事件捕获；沿事件路径向目标传播的阶段   | 监听阶段与事件目标                       |
| 19   | event bubbling   | 事件冒泡；事件从目标向祖先传播的阶段   | 并非所有事件都会冒泡                     |
| 20   | event delegation | 事件委托；通过祖先监听器处理后代事件   | target、currentTarget 和传播阶段         |
| 21   | debounce         | 防抖；合并短时间内的连续调用           | 等待窗口、首尾触发和取消                 |
| 22   | worker pool      | 工作池；复用有限工作单元处理任务       | 并发上限、排队与补位；区分异步任务和线程 |

### 浏览器、网络与布局

| 编号 | 英文锚点                      | Meaning 方向                           | Notes 展开范围                        |
|:-----|:------------------------------|:---------------------------------------|:--------------------------------------|
| 23   | same-origin policy            | 同源策略；限制不同源之间的资源访问     | 协议、主机、端口和访问边界            |
| 24   | cross-origin resource sharing | 跨源资源共享；服务器声明跨源读取权限   | CORS 响应头和凭据请求                 |
| 25   | preflight request             | 预检请求；实际跨源请求前的权限检查     | OPTIONS 与预检触发条件                |
| 26   | HTTP caching                  | HTTP 缓存；按缓存规则复用响应          | 新鲜度、私有缓存与共享缓存            |
| 27   | cache revalidation            | 缓存再验证；向服务器确认缓存是否仍有效 | ETag、Last-Modified 和 304            |
| 28   | web storage                   | Web 存储；浏览器提供的键值存储机制     | localStorage 与 sessionStorage 的范围 |
| 29   | storage event                 | 存储事件；通知其他相关文档的存储变更   | 触发范围；修改来源窗口不接收自身事件  |
| 30   | block formatting context      | 块级格式化上下文；独立的块布局环境     | 浮动包含、外边距和建立条件            |
| 31   | flexbox                       | 弹性布局；沿主轴分配空间的一维布局     | 主轴、交叉轴、伸缩和换行              |
| 32   | grid layout                   | 网格布局；按行和列组织内容的二维布局   | 轨道、网格区域与尺寸分配              |
| 33   | layout                        | 布局；计算元素几何位置和尺寸           | 布局失效及与绘制的区别                |
| 34   | paint                         | 绘制；生成元素视觉内容的过程           | 与布局、合成的区别                    |

### 框架与工程

| 编号 | 英文锚点             | Meaning 方向                            | Notes 展开范围                   |
|:-----|:---------------------|:----------------------------------------|:---------------------------------|
| 35   | virtual DOM          | 虚拟 DOM；用数据结构描述界面            | 与真实 DOM 的关系及更新成本      |
| 36   | reconciliation       | 协调；比较界面描述并确定更新            | 元素类型、key 和状态保留         |
| 37   | one-way data flow    | 单向数据流；沿明确方向传递状态          | 状态来源与更新入口               |
| 38   | two-way binding      | 双向绑定；关联状态与界面输入的更新      | 区分绑定语法与底层数据流         |
| 39   | memoization          | 记忆化；复用相同输入对应的计算结果      | 比较规则、成本和缓存失效         |
| 40   | effect               | 副作用；在 React 中用于同步外部系统     | 依赖、清理和开发环境检查         |
| 41   | middleware           | 中间件；在处理链中组合请求处理逻辑      | 上下文、顺序和错误处理           |
| 42   | onion model          | 洋葱模型；下游返回后继续执行上游逻辑    | Koa 的 await next() 与异常传播   |
| 43   | abstract syntax tree | 抽象语法树；表示程序语法结构的树        | 解析、遍历和代码转换             |
| 44   | tree shaking         | 摇树优化；移除未使用代码的构建优化      | 静态分析、副作用和构建配置       |
| 45   | semantic versioning  | 语义化版本；表达公共 API 变化的版本约定 | 主版本、次版本、补丁与预发布版本 |

### 算法思路

| 编号 | 英文锚点            | Meaning 方向                           | Notes 展开范围               |
|:-----|:--------------------|:---------------------------------------|:-----------------------------|
| 46   | binary search       | 二分查找；在有序范围中逐步缩小搜索区间 | 区间不变量、中点、终止和边界 |
| 47   | quicksort           | 快速排序；分区后递归处理子区间         | 基准选择、重复值和最坏情况   |
| 48   | backtracking        | 回溯；尝试选择并撤销选择以搜索解空间   | 状态、约束、终止和剪枝       |
| 49   | linked list merge   | 链表合并；按顺序连接两个有序链表的节点 | 游标推进、空链表和剩余部分   |
| 50   | postorder traversal | 后序遍历；对子树处理后再访问根节点     | 左右根顺序与递归栈           |
| 51   | least recently used | 最近最少使用；淘汰最久未被访问的数据   | 访问顺序维护与容量边界       |

### 设计与协作

| 编号 | 英文锚点                       | Meaning 方向                                | Notes 展开范围                 |
|:-----|:-------------------------------|:--------------------------------------------|:-------------------------------|
| 52   | progressive enhancement        | 渐进增强；先提供基础能力，再逐步增强体验    | 基础可用性与能力检测           |
| 53   | graceful degradation           | 优雅降级；能力不足时保留核心功能            | 核心路径与降级策略             |
| 54   | project onboarding             | 项目熟悉过程；建立对项目和工作流程的理解    | 运行项目、业务边界、测试和风险 |
| 55   | cross-functional collaboration | 跨职能协作；围绕共同目标协调不同角色        | 目标、接口、责任和反馈         |
| 56   | trade-off                      | 权衡；比较收益、成本和约束后作出选择        | 保留决策依据与适用条件         |
| 57   | STAR method                    | STAR 方法；按情境、任务、行动和结果组织经历 | 仅使用真实经历，不虚构结果     |

## 现有词条复用

已核对两项可复用词条:

- [closure-tech](../entries/closure-tech.yml): 当前 Meaning 为"闭包"。保留 ID、原有来源、音频及已批准关联。面试原理可补充到 Notes。
- [scope-lexical](../entries/scope-lexical.yml): 当前 Meaning 为"作用域"。保留词法语义，不与工作范围合并。

这两项是已确认的复用方向，不代表其余 55 项均需新增。实施时扫描整个集合，核对同义词、缩写和不同技术语境。尤其要区分 React effect 与现有其他 effect 词条。

## 原文风险与核实要求

已识别的风险应在导入时处理，保留原始材料不变:

1. 栈与堆的说明混用了不同语言的内存管理方式。JavaScript 的说明应围绕运行时和可达性展开。
2. Fetch 的凭据行为需要区分同源与跨源，不能概括为默认不携带 Cookie。
3. 原文的 Promise.all 实现按完成顺序追加结果。最后一个索引完成也不代表全部完成；空数组无法完成。
4. 二分查找右侧分支使用 `start + 1`。这不能保证每轮排除半个区间，不应沿用其复杂度说明。
5. 链表节点定义使用 `element`，合并代码读取 `val`。字段不一致，不能直接作为正确示例。
6. 事件委托不应限制为只能通过冒泡实现。需说明捕获与冒泡的适用边界。
7. 缓存说明需区分 `no-cache` 与 `no-store`，避免把再验证等同于禁止存储。
8. React effect、memoization、Vue 行为及构建优化需核实版本。不要把性能建议写成语义保证。

以下链接供实施时复核。原文没有完整解释的候选，仍需补充一手依据:

- [MDN: Memory management](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Guide/Memory_management)
- [MDN: Using Fetch](https://developer.mozilla.org/en-US/docs/Web/API/Fetch_API/Using_Fetch)
- [MDN: Promise.all](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Promise/all)
- [MDN: CORS](https://developer.mozilla.org/en-US/docs/Web/HTTP/Guides/CORS)
- [MDN: Event bubbling](https://developer.mozilla.org/en-US/docs/Learn_web_development/Core/Scripting/Event_bubbling)
- [MDN: Cache-Control](https://developer.mozilla.org/en-US/docs/Web/HTTP/Reference/Headers/Cache-Control)
- [React: Synchronizing with Effects](https://react.dev/learn/synchronizing-with-effects)
- [React: useMemo](https://react.dev/reference/react/useMemo)
- [Koa guide](https://github.com/koajs/koa/blob/master/docs/guide.md)

## 学习顺序与关联

采用现有 [Learning order](LEARNING_ORDER.md) 机制。先理解基础概念，再引入依赖这些概念的应用。章节顺序和表格编号不是最终 StudyOrder。

可评估的顺序包括作用域到闭包、Promise 到 Promise.all。同源策略可先于 CORS 和预检请求。实施时判断直接依赖，按真实稳定 ID 写入 `prerequisites`。

`learning_priority` 根据实用性和概念复杂度确定。不按英文长度或章节批量赋值。没有明确依赖时不强加连线。

记忆关联与先修关系分别处理。浅拷贝与深拷贝等辨析可作为关联候选。先列出双方 ID、当前释义、理由和混淆点。只有人工确认后的关系才写入 `related_entries`。

本计划没有批准新的关联对，也不改变已有的 21 组关系。导出器继续负责双向展示。

## 暂不收录

- 薪资、离职原因、Offer 状态和团队人数等个人信息。
- 未经用户提供的工作成果或面试经历。
- 仅有宽泛标题、尚未形成单一学习目标的数据库或系统设计题。
- 为覆盖原文而机械拆出的词语，以及组合已有词汇的普通句子。

STAR 等通用方法可以收录，个人故事必须基于真实材料。后续需要拓展主题时，另行整理具体候选。

## 实施步骤

1. 确认导入范围。重新读取当前 schema、现有词条和原文。
2. 对 57 项逐一记录复用、新增、合并或暂缓及理由。
3. 为新增项确定唯一稳定 ID。复用项保留已有身份与来源。
4. 核实原理、边界和版本，撰写中文 Meaning 与 Notes。
5. 按需添加 Context、英语 Example 和可靠 IPA。
6. 设置选择性的学习优先级和先修关系。新增关联另行确认。
7. 音频默认保留现状。若获授权，则生成英文锚点读音并更新元数据。
8. 验证、完整构建，并记录实际数量与未核实事项。

不修改 schema、导出格式、卡片模板或 Anki 调度逻辑。不自动改动本地 Anki 集合，不提交或推送 Git。

## 验收

导入完成时应满足:

- 57 个候选均有明确去向；分别统计新增、复用、合并和暂缓。
- 未新增重复含义，既有稳定 ID 和无关内容保持不变。
- 原理经过核实，来源可追溯，已知错误未照搬。
- 缺省的 Context 和 Example 不产生占位内容。
- Notes 保留所需深度，未把推演材料误设为朗读例句。
- 关联仅包含已批准的关系，先修图通过验证。
- 未授权的音频保持不变；新增音频验证解码及对应关系。

实施后的仓库检查:

```bash
go test ./...
go vet ./...
go run scripts/validate/main.go
go run scripts/build/main.go
git diff --check
```

连续构建两次并比较导出字节。检查每行 17 列、稳定 ID、学习顺序和音频引用。另行检查新增文件的格式；报告既有重复词警告。

静态验证不能证明读音质量或真实设备播放正常。未做人工试听或 Anki 检查时，应明确说明。

## 当前进度

- 已完成: 六类候选整理、复用方向、风险记录和实施验收计划。
- 待执行: 明确批准后的内容核实与词条导入。
- 尚未执行: YAML 修改、音频生成、导出重建和 Anki 操作。
