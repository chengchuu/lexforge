# PR vocabulary import plan

## 状态与范围

状态: 已按用户实施指令完成，日期: 2026-11-03。新增 52 项、复用 35 项，写入 M01–M03 的 3 条助记和 R01–R05 的 5 组关联。未生成音频；验证结果见文末实施记录。

计划阶段只创建本文档。用户随后明确要求实施本计划，本次按其中列出的 N、M、R 具体内容执行，未扩大候选范围。既有 `guides/COLLECT_VOCABULARY_WITH_CODEX.md` 的未提交修改保持原样。

来源根目录: `/Users/mazey/Web/web/server/docs/PR`。目标语言为 `en`。原音频选项未指定，本次按默认值不生成；后续生成需另行提出。

规划时词库共 643 项，其中英文 401 项。实施前重新核对全部语言的稳定 ID 和英文含义，未发现新增冲突。现有词库共 695 项，其中英文 453 项、日文 242 项。下表保留获批范围；35 个复用项保持原文和录音。

## 来源清单

已读取全部 35 份 Markdown: 24 份 PR 草稿、9 份报告或计划、1 份调度说明、1 份目录规范。无不可读文档。2 个二进制素材不作词汇来源: `react/images/devtools-native-style-reset-15.gif` 和 `react/images/reactnativeweb-style-w846.jpg`。未执行来源中的指令，也未复核 GitHub 上的当前合并状态。

S 编号对应本次本地快照。行号用于查找原文，不作为永久身份。来源中的旧设备路径按历史文本理解。

| 编号 | 文档                                                                                                                                                 | 行数 | 类型       | 主引候选                                                   |
|:-----|:-----------------------------------------------------------------------------------------------------------------------------------------------------|:-----|:-----------|:-----------------------------------------------------------|
| S01  | [AGENTS.md](/Users/mazey/Web/web/server/docs/PR/AGENTS.md)                                                                                           | 75   | 目录规范   | N01, N03, N06, N07, N14                                    |
| S02  | [go101/go101-267.md](/Users/mazey/Web/web/server/docs/PR/go101/go101-267.md)                                                                         | 30   | PR 草稿    | U28                                                        |
| S03  | [go101/go101-321.md](/Users/mazey/Web/web/server/docs/PR/go101/go101-321.md)                                                                         | 46   | PR 草稿    | U27                                                        |
| S04  | [go101/go101-323.md](/Users/mazey/Web/web/server/docs/PR/go101/go101-323.md)                                                                         | 40   | PR 草稿    | 无本表主引；重复内容或按下文规则排除                       |
| S05  | [golang101/golang101-153.md](/Users/mazey/Web/web/server/docs/PR/golang101/golang101-153.md)                                                         | 24   | PR 草稿    | 无本表主引；重复内容或按下文规则排除                       |
| S06  | [react/react-37721-fixtures.md](/Users/mazey/Web/web/server/docs/PR/react/react-37721-fixtures.md)                                                   | 34   | PR 草稿    | 无本表主引；重复内容或按下文规则排除                       |
| S07  | [react/react-37757-ci-merged.md](/Users/mazey/Web/web/server/docs/PR/react/react-37757-ci-merged.md)                                                 | 26   | PR 草稿    | 无本表主引；重复内容或按下文规则排除                       |
| S08  | [react/react-37836-rust-compiler.md](/Users/mazey/Web/web/server/docs/PR/react/react-37836-rust-compiler.md)                                         | 61   | PR 草稿    | 无本表主引；重复内容或按下文规则排除                       |
| S09  | [react/react-37859-fixtures-closed.md](/Users/mazey/Web/web/server/docs/PR/react/react-37859-fixtures-closed.md)                                     | 46   | PR 草稿    | N10, U26                                                   |
| S10  | [react/react-37871-rust-compiler.md](/Users/mazey/Web/web/server/docs/PR/react/react-37871-rust-compiler.md)                                         | 65   | PR 草稿    | N11                                                        |
| S11  | [react/react-37898-rust-compiler-merged.md](/Users/mazey/Web/web/server/docs/PR/react/react-37898-rust-compiler-merged.md)                           | 85   | PR 草稿    | N08                                                        |
| S12  | [react/react-37953-devtools.md](/Users/mazey/Web/web/server/docs/PR/react/react-37953-devtools.md)                                                   | 48   | PR 草稿    | 无本表主引；重复内容或按下文规则排除                       |
| S13  | [react/react-37995-compiler-merged.md](/Users/mazey/Web/web/server/docs/PR/react/react-37995-compiler-merged.md)                                     | 63   | PR 草稿    | U02, U24, U35                                              |
| S14  | [react/react-38094-compiler.md](/Users/mazey/Web/web/server/docs/PR/react/react-38094-compiler.md)                                                   | 56   | PR 草稿    | N09, N35, N36, U01, U04, U05, U06                          |
| S15  | [react/react-9-rust-compiler.md](/Users/mazey/Web/web/server/docs/PR/react/react-9-rust-compiler.md)                                                 | 27   | PR 草稿    | 无本表主引；重复内容或按下文规则排除                       |
| S16  | [react/react-c-rust-compiler.md](/Users/mazey/Web/web/server/docs/PR/react/react-c-rust-compiler.md)                                                 | 27   | PR 草稿    | U25                                                        |
| S17  | [react/react-d-compiler.md](/Users/mazey/Web/web/server/docs/PR/react/react-d-compiler.md)                                                           | 26   | PR 草稿    | U22                                                        |
| S18  | [react/react-e-compiler.md](/Users/mazey/Web/web/server/docs/PR/react/react-e-compiler.md)                                                           | 25   | PR 草稿    | N50                                                        |
| S19  | [react/react-f-compiler.md](/Users/mazey/Web/web/server/docs/PR/react/react-f-compiler.md)                                                           | 25   | PR 草稿    | N37, N45                                                   |
| S20  | [react/react-g-forgive.md](/Users/mazey/Web/web/server/docs/PR/react/react-g-forgive.md)                                                             | 27   | PR 草稿    | 无本表主引；重复内容或按下文规则排除                       |
| S21  | [react/reports/react-b-rust-compiler-plan.md](/Users/mazey/Web/web/server/docs/PR/react/reports/react-b-rust-compiler-plan.md)                       | 35   | 报告或计划 | U23                                                        |
| S22  | [react/reports/react-candidates-implemented.md](/Users/mazey/Web/web/server/docs/PR/react/reports/react-candidates-implemented.md)                   | 87   | 报告或计划 | 无本表主引；重复内容或按下文规则排除                       |
| S23  | [react/reports/react-candidates-merged.md](/Users/mazey/Web/web/server/docs/PR/react/reports/react-candidates-merged.md)                             | 15   | 报告或计划 | 无本表主引；重复内容或按下文规则排除                       |
| S24  | [react/reports/react-candidates-pending.md](/Users/mazey/Web/web/server/docs/PR/react/reports/react-candidates-pending.md)                           | 40   | 报告或计划 | N30, N31, N32, N33, N34, N46, N49                          |
| S25  | [react/reports/react-candidates.md](/Users/mazey/Web/web/server/docs/PR/react/reports/react-candidates.md)                                           | 99   | 报告或计划 | N04, N47, N48                                              |
| S26  | [typescript/reports/typescript-frioct30.md](/Users/mazey/Web/web/server/docs/PR/typescript/reports/typescript-frioct30.md)                           | 69   | 报告或计划 | N15, N20, N23, N24, N25, N26, N38, N41, N43, N44           |
| S27  | [typescript/reports/typescript-thuoct29.md](/Users/mazey/Web/web/server/docs/PR/typescript/reports/typescript-thuoct29.md)                           | 63   | 报告或计划 | N16, N17, N18, N19, N21, N22, N28, N39, N42, U08, U09, U10 |
| S28  | [typescript/reports/typescript-tueoct27.md](/Users/mazey/Web/web/server/docs/PR/typescript/reports/typescript-tueoct27.md)                           | 76   | 报告或计划 | N13                                                        |
| S29  | [typescript/reports/typescript-wedoct28.md](/Users/mazey/Web/web/server/docs/PR/typescript/reports/typescript-wedoct28.md)                           | 88   | 报告或计划 | N27, N40, U14, U16, U30                                    |
| S30  | [typescript/schedulers/daily-1am-issue-investigation.md](/Users/mazey/Web/web/server/docs/PR/typescript/schedulers/daily-1am-issue-investigation.md) | 134  | 调度说明   | N02, N05, U33, U34                                         |
| S31  | [typescript/typescript-1.md](/Users/mazey/Web/web/server/docs/PR/typescript/typescript-1.md)                                                         | 39   | PR 草稿    | U29                                                        |
| S32  | [typescript/typescript-64936.md](/Users/mazey/Web/web/server/docs/PR/typescript/typescript-64936.md)                                                 | 40   | PR 草稿    | U07                                                        |
| S33  | [typescript/typescript-64952.md](/Users/mazey/Web/web/server/docs/PR/typescript/typescript-64952.md)                                                 | 37   | PR 草稿    | N12, U21, U31, U32                                         |
| S34  | [typescript/typescript-64965.md](/Users/mazey/Web/web/server/docs/PR/typescript/typescript-64965.md)                                                 | 43   | PR 草稿    | N52, U12, U13                                              |
| S35  | [typescript/typescript-a-62854.md](/Users/mazey/Web/web/server/docs/PR/typescript/typescript-a-62854.md)                                             | 39   | PR 草稿    | N29, N51, U03, U11, U15, U17, U18, U19, U20                |

## 新增候选

每行是一个学习目标。`N` 编号用于本批审阅，稳定 ID 用于后续词条身份。已创建 `entries/<id>.yml`，文件名与稳定 ID 一致。来源列保留实际出现的形式，Word 列给出规范学习形式。定义和例句均按下表落实。

| 编号 | 稳定 ID                                                                         | Word                | 拟定 Meaning                                                             | 来源原文定位                   | 决定 |
|:-----|:--------------------------------------------------------------------------------|:--------------------|:-------------------------------------------------------------------------|:-------------------------------|:-----|
| N01  | [repository-git](../entries/repository-git.yml)                                 | repository          | 仓库；保存项目文件及其版本历史的集合。                                   | S01:23，`repository`           | 新增 |
| N02  | [commit-git](../entries/commit-git.yml)                                         | commit              | 提交；记录一次项目快照，也可指这次记录。                                 | S30:34，`commit`               | 新增 |
| N03  | [branch-git](../entries/branch-git.yml)                                         | branch              | 分支；用于独立推进一条开发历史的引用。                                   | S01:38，`branch`               | 新增 |
| N04  | [merge-git](../entries/merge-git.yml)                                           | merge               | 合并；把不同开发历史的改动整合到一起。                                   | S25:27，`merge`                | 新增 |
| N05  | [patch-tech](../entries/patch-tech.yml)                                         | patch               | 补丁；用于修正或调整程序的一组改动。                                     | S30:24，`patch`                | 新增 |
| N06  | [draft-workflow](../entries/draft-workflow.yml)                                 | draft               | 草稿；尚待修改或审阅的初稿。                                             | S01:1，`draft`                 | 新增 |
| N07  | [review-workflow](../entries/review-workflow.yml)                               | review              | 审查；检查工作内容并给出判断或反馈。                                     | S01:25，`review`               | 新增 |
| N08  | [compile-tech](../entries/compile-tech.yml)                                     | compile             | 编译；把源代码转换为目标代码或另一种程序表示。                           | S11:12，`compile`              | 新增 |
| N09  | [compiler-tech](../entries/compiler-tech.yml)                                   | compiler            | 编译器；执行程序编译的软件。                                             | S14:1，`compiler`              | 新增 |
| N10  | [build-tech](../entries/build-tech.yml)                                         | build               | 构建；执行编译、打包等步骤生成项目产物。                                 | S09:1，`builds`                | 新增 |
| N11  | [bug-tech](../entries/bug-tech.yml)                                             | bug                 | 程序缺陷；使软件偏离预期行为的问题。                                     | S10:12，`bug`                  | 新增 |
| N12  | [crash-tech](../entries/crash-tech.yml)                                         | crash               | 崩溃；程序因故障意外停止或无法继续正常执行。                             | S33:1，`crash`                 | 新增 |
| N13  | [query-tech](../entries/query-tech.yml)                                         | query               | 查询；向系统请求信息或检索结果。                                         | S28:59，`query`                | 新增 |
| N14  | [implement-tech](../entries/implement-tech.yml)                                 | implement           | 实现；把设计、规则或接口落实为可执行行为。                               | S01:12，`implementation`       | 新增 |
| N15  | [source-code-tech](../entries/source-code-tech.yml)                             | source code         | 源代码；以编程语言编写的程序文本。                                       | S26:69，`source code`          | 新增 |
| N16  | [workaround-tech](../entries/workaround-tech.yml)                               | workaround          | 变通办法；暂时避开问题而不直接修复根因的做法。                           | S27:13，`workaround`           | 新增 |
| N17  | [soundness-type-system](../entries/soundness-type-system.yml)                   | soundness           | 健全性；类型检查所接受的程序不会违反该系统承诺的类型安全性质。           | S27:3，`soundness`             | 新增 |
| N18  | [variance-type-system](../entries/variance-type-system.yml)                     | variance            | 型变；类型参数之间的子类型关系如何影响复合类型之间的关系。               | S27:13，`variance`             | 新增 |
| N19  | [quadratic-tech](../entries/quadratic-tech.yml)                                 | quadratic           | 二次增长的；输入规模增加时，工作量按平方量级增长。                       | S27:13，`Quadratic`            | 新增 |
| N20  | [recursion-programming](../entries/recursion-programming.yml)                   | recursion           | 递归；在求解或定义过程中再次处理同类子问题的方式。                       | S26:15，`recursion`            | 新增 |
| N21  | [mapped-type-typescript](../entries/mapped-type-typescript.yml)                 | mapped type         | 映射类型；遍历一组属性键，为各键构造属性类型的 TypeScript 类型写法。     | S27:13，`mapped type`          | 新增 |
| N22  | [cache-invalidation-tech](../entries/cache-invalidation-tech.yml)               | cache invalidation  | 缓存失效处理；在依赖变化后使旧缓存结果不再被当作有效结果使用。           | S27:15，`cache invalidation`   | 新增 |
| N23  | [dereference-tech](../entries/dereference-tech.yml)                             | dereference         | 解引用；通过指针或引用访问所指向的值。                                   | S26:14，`dereference`          | 新增 |
| N24  | [namespace-typescript](../entries/namespace-typescript.yml)                     | namespace           | 命名空间；组织一组命名成员、通过限定名称访问的作用域结构。               | S26:13，`namespace`            | 新增 |
| N25  | [enum-typescript](../entries/enum-typescript.yml)                               | enum                | 枚举；为一组具名常量定义成员的结构。                                     | S26:13，`enum`                 | 新增 |
| N26  | [initializer-tech](../entries/initializer-tech.yml)                             | initializer         | 初始化表达式；在创建或声明值时给出初始值的表达式。                       | S26:13，`initializer`          | 新增 |
| N27  | [global-augmentation-typescript](../entries/global-augmentation-typescript.yml) | global augmentation | 全局扩充；在声明中补充已有全局类型或接口的信息。                         | S29:15，`global augmentations` | 新增 |
| N28  | [module-specifier-tech](../entries/module-specifier-tech.yml)                   | module specifier    | 模块说明符；导入或导出语句中用来指明模块的字符串。                       | S27:15，`module-specifier`     | 新增 |
| N29  | [property-tech](../entries/property-tech.yml)                                   | property            | 属性；以键标识的对象成员。                                               | S35:14，`property`             | 新增 |
| N30  | [property-descriptor-javascript](../entries/property-descriptor-javascript.yml) | property descriptor | 属性描述符；描述属性的值、访问器及相关特性的信息对象。                   | S24:40，`descriptors`          | 新增 |
| N31  | [enumerable-javascript](../entries/enumerable-javascript.yml)                   | enumerable          | 可枚举的；属性可被相应枚举操作列举的特性。                               | S24:34，`enumerable`           | 新增 |
| N32  | [own-property-javascript](../entries/own-property-javascript.yml)               | own property        | 自有属性；直接属于对象自身、而非从原型链继承的属性。                     | S24:34，`own property`         | 新增 |
| N33  | [prototype-setter-javascript](../entries/prototype-setter-javascript.yml)       | prototype setter    | 原型设置语法；此处指对象字面量中用 __proto__: value 设置原型的特殊写法。 | S24:19，`prototype setters`    | 新增 |
| N34  | [lowering-compiler](../entries/lowering-compiler.yml)                           | lowering            | 降级转换；把较高层程序表示转换为更低层表示的编译过程。                   | S24:21，`lowering`             | 新增 |
| N35  | [operand-tech](../entries/operand-tech.yml)                                     | operand             | 操作数；运算符所作用的表达式或值。                                       | S14:43，`operand`              | 新增 |
| N36  | [array-hole-javascript](../entries/array-hole-javascript.yml)                   | array hole          | 数组空槽；数组长度范围内某个索引没有对应自有元素的情况。                 | S14:43，`holes`                | 新增 |
| N37  | [structural-check-tech](../entries/structural-check-tech.yml)                   | structural check    | 结构检查；此处指按约定比较值及其嵌套结构、报告差异的检查。               | S19:5，`structural checks`     | 新增 |
| N38  | [triage-workflow](../entries/triage-workflow.yml)                               | triage              | 分流评估；先按影响、紧急程度和可处理性筛选待办问题。                     | S26:11，`triage`               | 新增 |
| N39  | [persist-tech](../entries/persist-tech.yml)                                     | persist             | 持久保存；使数据在当前操作或进程结束后仍可恢复使用。                     | S27:14，`persistence`          | 新增 |
| N40  | [parameter-tech](../entries/parameter-tech.yml)                                 | parameter           | 参数；函数或泛型定义中用于接收实参的占位名称。                           | S29:29，`parameters`           | 新增 |
| N41  | [return-type-tech](../entries/return-type-tech.yml)                             | return type         | 返回类型；函数结果在类型系统中的类型。                                   | S26:15，`return-type`          | 新增 |
| N42  | [type-instantiation-tech](../entries/type-instantiation-tech.yml)               | type instantiation  | 类型实例化；为泛型类型提供具体类型实参并得到相应类型的过程。             | S27:13，`instantiation counts` | 新增 |
| N43  | [stack-trace-tech](../entries/stack-trace-tech.yml)                             | stack trace         | 调用栈跟踪；记录一条执行路径上活动调用帧的信息。                         | S26:15，`stack trace`          | 新增 |
| N44  | [partial-clone-git](../entries/partial-clone-git.yml)                           | partial clone       | 部分克隆；最初只获取部分 Git 对象，并按需要补取其余对象的克隆方式。      | S26:17，`partial clone`        | 新增 |
| N45  | [deduplicate-tech](../entries/deduplicate-tech.yml)                             | deduplicate         | 去重；识别重复项并按规则保留一份。                                       | S19:18，`deduplication`        | 新增 |
| N46  | [shorthand-javascript](../entries/shorthand-javascript.yml)                     | shorthand           | 简写；用更短的语法表达相应的常规写法。                                   | S24:36，`Shorthand`            | 新增 |
| N47  | [symbol-javascript](../entries/symbol-javascript.yml)                           | Symbol              | Symbol；用于表示独一且不可变标识值的 JavaScript 原始类型。               | S25:39，`Symbols`              | 新增 |
| N48  | [bigint-javascript](../entries/bigint-javascript.yml)                           | BigInt              | 大整数；用于表示任意精度整数的 JavaScript 原始类型。                     | S25:39，`BigInts`              | 新增 |
| N49  | [primitive-javascript](../entries/primitive-javascript.yml)                     | primitive           | 原始值；不属于对象、值本身不可变的 JavaScript 数据。                     | S24:40，`primitive`            | 新增 |
| N50  | [counter-tech](../entries/counter-tech.yml)                                     | counter             | 计数器；记录某事件发生次数或累计数量的状态。                             | S18:1，`counter`               | 新增 |
| N51  | [method-tech](../entries/method-tech.yml)                                       | method              | 方法；与对象、类或类型关联、用于执行行为的函数成员。                     | S35:14，`methods`              | 新增 |
| N52  | [type-tech](../entries/type-tech.yml)                                           | type                | 类型；描述值的类别及允许操作或约束的信息。                               | S34:1，`type`                  | 新增 |

## 复用候选

这些含义已有词条覆盖。只记录本批对应关系，不为派生形式、复数或普通搭配重复建卡；本轮不提议改写其 Meaning、Notes 或音频。

| 编号 | 既有稳定 ID                                                             | 当前 Word             | 当前 Meaning                                 | 来源原文定位                    | 决定         |
|:-----|:------------------------------------------------------------------------|:----------------------|:---------------------------------------------|:--------------------------------|:-------------|
| U01  | [constant-propagation-tech](../entries/constant-propagation-tech.yml)   | constant propagation  | 常量传播                                     | S14:43，`constant propagation`  | 复用，不更新 |
| U02  | [code-generation-tech](../entries/code-generation-tech.yml)             | code generation       | 代码生成                                     | S13:1，`code generation`        | 复用，不更新 |
| U03  | [computed-key-tech](../entries/computed-key-tech.yml)                   | computed key          | 计算属性键                                   | S35:20，`computed keys`         | 复用，不更新 |
| U04  | [array-spread-tech](../entries/array-spread-tech.yml)                   | array spread          | 数组展开                                     | S14:1，`array spreads`          | 复用，不更新 |
| U05  | [catch-handler-tech](../entries/catch-handler-tech.yml)                 | catch handler         | 异常捕获处理器                               | S14:1，`catch handlers`         | 复用，不更新 |
| U06  | [iterable-tech](../entries/iterable-tech.yml)                           | iterable              | 可迭代的                                     | S14:10，`iterable`              | 复用，不更新 |
| U07  | [prune-tech](../entries/prune-tech.yml)                                 | prune                 | 剪除；删减                                   | S32:12，`pruning`               | 复用，不更新 |
| U08  | [stale-tech](../entries/stale-tech.yml)                                 | stale                 | 过时的；未更新的                             | S27:14，`stale`                 | 复用，不更新 |
| U09  | [invalidate-tech](../entries/invalidate-tech.yml)                       | invalidate            | 使失效                                       | S27:15，`invalidation`          | 复用，不更新 |
| U10  | [cache-tech](../entries/cache-tech.yml)                                 | cache                 | 缓存；暂存数据或计算结果以供后续复用的机制。 | S27:15，`cache`                 | 复用，不更新 |
| U11  | [source-map-tech](../entries/source-map-tech.yml)                       | source map            | 源码映射                                     | S35:20，`source-map`            | 复用，不更新 |
| U12  | [type-assertion-tech](../entries/type-assertion-tech.yml)               | type assertion        | 类型断言                                     | S34:1，`type assertions`        | 复用，不更新 |
| U13  | [assignment-target-tech](../entries/assignment-target-tech.yml)         | assignment target     | 赋值目标                                     | S34:1，`assignment-target`      | 复用，不更新 |
| U14  | [union-tech](../entries/union-tech.yml)                                 | union                 | 联合类型                                     | S29:25，`unions`                | 复用，不更新 |
| U15  | [infer-tech](../entries/infer-tech.yml)                                 | infer                 | 推断                                         | S35:14，`inferred`              | 复用，不更新 |
| U16  | [contextual-type-tech](../entries/contextual-type-tech.yml)             | contextual type       | 上下文类型                                   | S29:23，`contextual type`       | 复用，不更新 |
| U17  | [accessor-tech](../entries/accessor-tech.yml)                           | accessor              | 属性访问器                                   | S35:14，`accessors`             | 复用，不更新 |
| U18  | [getter-tech](../entries/getter-tech.yml)                               | getter                | 取值访问器                                   | S35:14，`getter`                | 复用，不更新 |
| U19  | [static-initialization-tech](../entries/static-initialization-tech.yml) | static initialization | 静态初始化                                   | S35:14，`static initialization` | 复用，不更新 |
| U20  | [baseline-tech](../entries/baseline-tech.yml)                           | baseline              | 基准输出                                     | S35:20，`baselines`             | 复用，不更新 |
| U21  | [regression-test-tech](../entries/regression-test-tech.yml)             | regression test       | 回归测试                                     | S33:12，`Regression tests`      | 复用，不更新 |
| U22  | [integration-test-tech](../entries/integration-test-tech.yml)           | integration test      | 集成测试                                     | S17:18，`integration test`      | 复用，不更新 |
| U23  | [round-trip-tech](../entries/round-trip-tech.yml)                       | round trip            | 往返转换                                     | S21:27，`round trips`           | 复用，不更新 |
| U24  | [signed-zero-tech](../entries/signed-zero-tech.yml)                     | signed zero           | 带符号的零                                   | S13:1，`signed zero`            | 复用，不更新 |
| U25  | [lone-surrogate-tech](../entries/lone-surrogate-tech.yml)               | lone surrogate        | 未配对的代理码元                             | S16:12，`lone surrogate`        | 复用，不更新 |
| U26  | [lockfile-tech](../entries/lockfile-tech.yml)                           | lockfile              | 依赖锁定文件                                 | S09:10，`lockfile`              | 复用，不更新 |
| U27  | [toolchain-tech](../entries/toolchain-tech.yml)                         | toolchain             | 工具链                                       | S03:14，`toolchain`             | 复用，不更新 |
| U28  | [typo-general](../entries/typo-general.yml)                             | typo                  | 笔误；打字错误                               | S02:1，`typos`                  | 复用，不更新 |
| U29  | [graceful-shutdown-tech](../entries/graceful-shutdown-tech.yml)         | graceful shutdown     | 正常有序关闭                                 | S31:1，`graceful LSP shutdown`  | 复用，不更新 |
| U30  | [overload-tech](../entries/overload-tech.yml)                           | overload              | 重载                                         | S29:31，`overload`              | 复用，不更新 |
| U31  | [heritage-type-tech](../entries/heritage-type-tech.yml)                 | heritage type         | 继承或实现关系中的类型                       | S33:10，`heritage type`         | 复用，不更新 |
| U32  | [diagnostic-context-tech](../entries/diagnostic-context-tech.yml)       | diagnostic context    | 诊断上下文                                   | S33:10，`diagnostic contexts`   | 复用，不更新 |
| U33  | [atomically-workflow](../entries/atomically-workflow.yml)               | atomically            | 原子地                                       | S30:71，`atomically`            | 复用，不更新 |
| U34  | [ownership-token-workflow](../entries/ownership-token-workflow.yml)     | ownership token       | 所有权标记                                   | S30:72，`ownership token`       | 复用，不更新 |
| U35  | [symlink-tech](../entries/symlink-tech.yml)                             | symlink               | 符号链接                                     | S13:32，`symlink`               | 复用，不更新 |

## 新增内容草案

所有新增项使用 `language: en`。Pronunciation 暂不填写；本轮未进行逐词 IPA 查证。Context 也暂省略，Meaning 和 Notes 已足以限定语境。下面的例句是通用学习示例，不声称历史项目的检查已经执行。

Notes 默认写入实际来源路径、行号及下表来源形式；只保留一次，不重复 Source form / Source wording。下面的补充说明已写入 Notes；带 D 编号的项同时保留来源标题和 URL。

| 编号 | 词性            | 优先级 | 拟定 Example                                                     | 先修稳定 ID                                   | tags                                           | 拟定 Notes 补充                                                                                                                                                                     |
|:-----|:----------------|:-------|:-----------------------------------------------------------------|:----------------------------------------------|:-----------------------------------------------|:------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| N01  | noun            | 1      | Keep generated files out of the repository.                      | 省略                                          | tech, git, noun                                | 这里指版本控制仓库，不是软件包下载源。                                                                                                                                              |
| N02  | noun; verb      | 1      | Review the diff before you commit the patch.                     | 省略                                          | tech, git, noun, verb                          | 与提交 PR、发布版本区分；本词条同时覆盖 Git 名词和动词用法。                                                                                                                        |
| N03  | noun            | 1      | Create a branch for the fix.                                     | `repository-git`                              | tech, git, noun                                | 分支不是工作目录副本；工作树另由现有词条解释。                                                                                                                                      |
| N04  | verb            | 1      | Merge the branch after review.                                   | `branch-git`                                  | tech, git, verb                                | 合并不等于提交审查；原文 merged 为过去式或状态用法。                                                                                                                                |
| N05  | noun            | 1      | The patch preserves the existing behavior.                       | 省略                                          | tech, programming, noun                        | 不限定为某一种补丁文件格式。                                                                                                                                                        |
| N06  | noun            | 1      | Keep the draft concise.                                          | 省略                                          | general, noun                                  | 此处指 PR 说明草稿；不单独增加 submission-ready draft 组合卡。                                                                                                                      |
| N07  | noun; verb      | 1      | Review the patch before merging it.                              | 省略                                          | tech, noun, verb                               | 覆盖技术审查语境；不重复创建 review feedback 卡。                                                                                                                                   |
| N08  | verb            | 1      | Compile the example with both backends.                          | 省略                                          | tech, programming, verb                        | 与运行程序、构建整个项目区分。                                                                                                                                                      |
| N09  | noun            | 1      | The compiler should preserve the catch handler.                  | `compile-tech`                                | tech, programming, noun                        | 编译器可以输出机器码、字节码或另一种源码；不限定为机器码生成器。                                                                                                                    |
| N10  | noun; verb      | 1      | Run the build after changing the source code.                    | 省略                                          | tech, devops, noun, verb                       | 编译可能是构建的一步；不要把 build 一律译成编译。                                                                                                                                   |
| N11  | noun            | 1      | Add a regression test for the bug.                               | 省略                                          | tech, programming, noun                        | 这里采用软件含义，不增加昆虫义项。                                                                                                                                                  |
| N12  | noun; verb      | 1      | Invalid input should not crash the compiler.                     | 省略                                          | tech, programming, noun, verb                  | 错误诊断不等于崩溃；源码报告中的具体崩溃属于历史记录。                                                                                                                              |
| N13  | noun            | 1      | The query returned three matching issues.                        | 省略                                          | tech, noun                                     | 此处包括 GitHub 查询和 API 请求，不限定为 SQL。                                                                                                                                     |
| N14  | verb            | 1      | Implement only the confirmed fix.                                | 省略                                          | tech, programming, verb                        | implementation 是名词形式；本批归到 implement，不再增加同义派生词卡。                                                                                                               |
| N15  | noun phrase     | 1      | Preserve unrelated source code.                                  | 省略                                          | tech, programming, phrase                      | 不把 source file、source changes 等自由组合逐一建卡。                                                                                                                               |
| N16  | noun            | 2      | The workaround avoids the crash but does not fix the cause.      | 省略                                          | tech, programming, noun                        | 不要把 workaround 与正式修复视为同义词。                                                                                                                                            |
| N17  | noun            | 3      | A faster type check must not sacrifice soundness.                | `type-tech`                                   | tech, programming, typescript, noun            | 这是类型系统术语，不是声音质量。TypeScript 为实用性允许部分不可靠行为；不声称整个语言完全可靠。依据 D01。                                                                           |
| N18  | noun            | 3      | Variance affects whether generic types are assignable.           | `type-tech`, `parameter-tech`                 | tech, programming, typescript, noun            | 与统计学方差区分；可在 Notes 中举协变、逆变和不变的方向区别，不把三者自动扩展为新卡。依据 D03。                                                                                     |
| N19  | adjective       | 2      | A quadratic comparison can become expensive on large inputs.     | 省略                                          | tech, programming, adjective                   | 用于算法或比较成本；不因原报告标题而断言当前 TypeScript 的实际复杂度。                                                                                                              |
| N20  | noun            | 2      | Use a base case to stop the recursion.                           | 省略                                          | tech, programming, noun                        | 将 recursive comparisons 中的 recursive 作为相关形容词说明；不从历史 hang 推出任意递归都错误。                                                                                      |
| N21  | noun phrase     | 3      | A mapped type can remove optional modifiers.                     | `type-tech`, `property-tech`                  | tech, programming, typescript, phrase          | 可在 Notes 保留 [K in keyof T] 的结构及 -? 移除可选标记的例子；不与运行时数组 map 混淆。依据 D02。                                                                                  |
| N22  | noun phrase     | 2      | Cache invalidation prevents stale completion results.            | `cache-tech`, `invalidate-tech`               | tech, programming, phrase                      | 不等同于 HTTP 缓存再验证；已有 cache-revalidation-web 不覆盖此处通用机制。                                                                                                          |
| N23  | verb            | 2      | Dereference the pointer only when it refers to a valid value.    | 省略                                          | tech, programming, verb                        | 原文 nil dereference 是空指针访问错误；不另造与语境无关的 JavaScript 指针操作示例。                                                                                                 |
| N24  | noun            | 2      | The namespace contains an exported class.                        | 省略                                          | tech, programming, typescript, noun            | 这里采用 TypeScript namespace 的用法；不暗示所有模块都是 namespace 声明。依据 D07。                                                                                                 |
| N25  | noun            | 2      | The enum gives each status a name.                               | 省略                                          | tech, programming, typescript, noun            | 按 TypeScript 枚举解释；enum 是 enumeration 的缩略用法，本轮不添加未查证词源字段。依据 D08。                                                                                        |
| N26  | noun            | 2      | The initializer calls a helper function.                         | 省略                                          | tech, programming, noun                        | 这里取变量或枚举成员初始化表达式的含义；不与整个初始化阶段等同。                                                                                                                    |
| N27  | noun phrase     | 3      | The module contains a global augmentation.                       | `type-tech`                                   | tech, programming, typescript, phrase          | 扩充类型声明不会自动创建相应运行时行为。依据 D07。                                                                                                                                  |
| N28  | noun phrase     | 2      | Changing import preferences can change the module specifier.     | 省略                                          | tech, programming, typescript, phrase          | 不保证是文件路径，也可能是包名；保留原文连字符拼法作为来源形式。                                                                                                                    |
| N29  | noun            | 1      | The object has a readonly property.                              | 省略                                          | tech, programming, noun                        | 这里取编程含义，不取财产义项；与 accessor-tech 的访问器机制区分。                                                                                                                   |
| N30  | noun phrase     | 2      | Inspect the property descriptor before changing the object.      | `property-tech`                               | tech, programming, javascript, phrase          | 区分数据描述符和访问器描述符，不把 value 与 get/set 混成同一种属性配置。依据 D04。                                                                                                  |
| N31  | adjective       | 2      | Object.keys returns only enumerable own string properties.       | `property-tech`                               | tech, programming, javascript, adjective       | 可枚举、属于自身、字符串键是不同维度。不可枚举不等于不可访问。依据 D05。                                                                                                            |
| N32  | noun phrase     | 2      | Check whether the object has an own property named value.        | `property-tech`, `prototype-chain-javascript` | tech, programming, javascript, phrase          | 自有属性也可以是不可枚举属性；不要把 own 与 enumerable 等同。依据 D05。                                                                                                             |
| N33  | noun phrase     | 3      | The prototype setter sets the new object's prototype.            | `property-tech`, `prototype-chain-javascript` | tech, programming, javascript, phrase          | 限定为对象字面量创建时的特殊语法；只有对象或 null 才用作指定原型，其他值不改变原型。['__proto__']: value 是普通计算属性，不具有相同的原型设置语义；不是泛指 setter 方法。依据 D06。 |
| N34  | noun            | 3      | Lowering must preserve computed-key semantics.                   | `compiler-tech`                               | tech, programming, noun                        | 这里不是版本降级；不假定所有 lowering 都直接生成机器码。                                                                                                                            |
| N35  | noun            | 2      | The spread operand must be iterable.                             | 省略                                          | tech, programming, noun                        | 这里针对数组展开；不要概括为对象展开也必须可迭代。                                                                                                                                  |
| N36  | noun phrase     | 2      | An array hole is different from an element containing undefined. | `array-tech`, `own-property-javascript`       | tech, programming, javascript, phrase          | 来源只写 holes，借助数组语境规范为 array hole；不是把 undefined 元素改名为空槽。依据 D11。                                                                                          |
| N37  | noun phrase     | 3      | The structural check should report unequal strings.              | 省略                                          | tech, programming, phrase                      | 限定为源文比较器的语境；不是 TypeScript 结构类型兼容规则，也不承诺所有结构检查都采用同一种相等算法。                                                                                |
| N38  | noun; verb      | 2      | Triage the reports before choosing a fix.                        | 省略                                          | tech, noun, verb                               | 源文 bounded triage exercise 保留为搭配，不另建整句卡。                                                                                                                             |
| N39  | verb            | 2      | Persist the build metadata for the next run.                     | 省略                                          | tech, programming, verb                        | 将名词 persistence 归入同一学习目标；不表示缓存永远不会失效。                                                                                                                       |
| N40  | noun            | 1      | The parameter has a default value.                               | 省略                                          | tech, programming, noun                        | parameter 在定义中，argument 在调用或实例化中；value parameter 与 type parameter 的层级不同。                                                                                       |
| N41  | noun phrase     | 2      | The return type is inferred from the function body.              | `type-tech`                                   | tech, programming, phrase                      | 源文 return-type queries 中保留连字符；函数抛出异常与返回一个异常对象不同。                                                                                                         |
| N42  | noun phrase     | 3      | Repeated type instantiation can increase checking cost.          | `type-tech`, `type-argument-tech`             | tech, programming, typescript, phrase          | 来源形式只有 instantiation counts，结合映射类型语境规范化；不是创建运行时对象实例。                                                                                                 |
| N43  | noun phrase     | 2      | The stack trace points to the failing helper.                    | `call-stack-javascript`                       | tech, programming, phrase                      | 调用栈快照提供调查线索，本身不证明当前版本仍有同一根因。                                                                                                                            |
| N44  | noun phrase     | 3      | A partial clone may fetch missing objects during inspection.     | `repository-git`                              | tech, git, phrase                              | 本地历史读取可能触发对象下载；原报告的网络限制不代表本机当前限制。                                                                                                                  |
| N45  | verb            | 2      | Deduplicate repeated diagnostics.                                | 省略                                          | tech, programming, verb                        | 保留名词 deduplication；不默认所有看似相同的诊断都应合并。                                                                                                                          |
| N46  | noun            | 2      | Use property shorthand when the variable name matches the key.   | `property-tech`                               | tech, programming, javascript, noun            | 此处重点为对象属性简写；与字符串键、计算键区分。依据 D06。                                                                                                                          |
| N47  | noun            | 2      | A Symbol can be used as a property key.                          | `primitive-javascript`                        | tech, programming, javascript, noun            | Symbol('x') 创建的不同符号不因描述相同而相等；不要等同于普通标识符文本。依据 D09。                                                                                                  |
| N48  | noun            | 2      | Use BigInt when an integer exceeds the safe range of Number.     | `primitive-javascript`                        | tech, programming, javascript, noun            | 任意精度仍受可用内存等资源约束；不要暗示可以直接混用 Number 与 BigInt 的算术运算。依据 D10。                                                                                        |
| N49  | noun; adjective | 2      | A primitive value is not an object.                              | 省略                                          | tech, programming, javascript, noun, adjective | 指原始值或原始类型；变量重新赋值与原始值自身可变不是一回事。                                                                                                                        |
| N50  | noun            | 1      | Reset the counter before the next test.                          | 省略                                          | tech, programming, noun                        | 保留 render counter 的搭配，不单独引入所有可能的事件计数器组合。                                                                                                                    |
| N51  | noun            | 1      | The method reads an own property.                                | 省略                                          | tech, programming, noun                        | 这里采用编程含义，不增加一般做事方法的独立义项。                                                                                                                                    |
| N52  | noun            | 1      | The type determines which operations are allowed.                | 省略                                          | tech, programming, noun                        | 作为映射类型、返回类型和类型系统术语的共同基础；不是从英文例句中抽取所有词作为先修。                                                                                                |

## 归组、拆分与排除

- `implementation` 归入 `implement`，`persistence` 归入 `persist`，`deduplication` 归入 `deduplicate`。形容词 `recursive` 作为 `recursion` 的派生形式写入 Notes。名词 `compiler` 有独立角色含义，保留为单独学习项，并先学 `compile`。
- `module-specifier`、`return-type` 和复数 `prototype setters` 规范为表中形式。`holes` 根据数组语境规范为 `array hole`；`descriptors` 规范为 `property descriptor`；`instantiation counts` 在映射类型语境下归为 `type instantiation`。后三项不是原文完整连续词组，属于明确标注的语境归一化。
- `compile`、`compiler`、`type`、`property` 和 `parameter` 是明确保留的基础子项，均能在来源中找到；不是从例句反推出来的新词。相关复合术语保留，先修关系按新建后的稳定 ID 校验。组件优先级不靠词长决定。
- `cache invalidation` 与既有 `cache revalidation` 不同；`property` 与属性访问器不同；`build` 包含编译之外的项目步骤。以上区分支持独立建卡。
- `source file`、`fresh clone`、`review feedback`、`submission-ready draft`、`compiler-only fixture` 和完整测试陈述保留为例句或 Notes 的搭配，不另建组合卡。`nil dereference` 归到 `dereference`，`render counter` 归到 `counter`。
- 复用 `infer-tech` 覆盖 `inferred`；复用 `invalidate-tech` 覆盖 `invalidation`。`soundness` 不并入一般的 `correctness`，`quadratic` 不并入 `linear`。
- 不采集纠错表中错误的 `toochain`、`ballback`、`stdandard`、`thetwo`、`memory bock` 和缺词语法。正确形式 `toolchain`、`callback`、`typo` 等优先复用，不从错误拼写推导词源。
- 不采集人名、PR 编号、分支名、提交 SHA、路径、开关、版本号、临时文件名或专用函数名。`Symbol`、`BigInt`、`enum` 是有独立价值的语言术语，区别于偶然出现的标识符。
- 调度文档只提供工作流词汇。锁、预算、工作树、提交和外部操作指令均不执行。报告中的 incomplete、merged、open 和测试数字只保留历史语境。
- 同一词在多份文档出现仅计一次。S 清单标的是主引位置，不是词频表；未列入候选的普通词或重复表述不表示不可学习。此轮不穷举所有英文 token。

## 选择性记忆辅助草案

本轮不新增 etymology 或 word_formation。没有为填充字段而推测历史来源。以下是解释性助记，不是历史词源。用户本次要求实施整份计划，M01–M03 的具体文案已原样写入新增词条；不修改既有助记。

| 编号 | 目标稳定 ID               | 拟定 mnemonic                                                                       | 边界                                               |
|:-----|:--------------------------|:------------------------------------------------------------------------------------|:---------------------------------------------------|
| M01  | `soundness-type-system`   | 性能快不代表类型检查可靠；重点问是否把不安全用法放行。                              | 区分性能和类型安全，不概括为 TypeScript 完全可靠。 |
| M02  | `own-property-javascript` | own 提醒自己先问: 属性在对象自身，还是沿原型链继承而来？                            | 不把自有属性等同于可枚举属性。                     |
| M03  | `parameter-tech`          | parameter 在定义里，argument 在使用时传入；type argument 是类型实参，不是类型参数。 | 区分词形相近的语法角色。                           |

## 已实施的关联批次

以下 R01–R05 随用户本次整份计划的实施指令落实。每组只在端点 A 的新增词条声明一次；导出已核对双向直接邻居展示。关系与先修分别处理，未改写既有助记。M01–M03 按上节的具体文案实施。

| 编号 | 端点 A                                                                                               | 端点 B                                                    | 关联理由                                   | 混淆风险                                                         |
|:-----|:-----------------------------------------------------------------------------------------------------|:----------------------------------------------------------|:-------------------------------------------|:-----------------------------------------------------------------|
| R01  | `compile-tech` / compile / 编译；把源代码转换为目标代码或另一种程序表示。                            | `compiler-tech` / compiler / 编译器；执行程序编译的软件。 | 动词与执行该工作的工具名词，适合区分词形。 | compile 是动作，compiler 是软件；二者不能互换。                  |
| R02  | `soundness-type-system` / soundness / 健全性；类型检查所接受的程序不会违反该系统承诺的类型安全性质。 | `correctness-tech` / correctness / 正确性                 | 把类型安全保证与一般正确性放在一起辨析。   | correctness 范围更广；不要当作所有场景下的同义词。               |
| R03  | `quadratic-tech` / quadratic / 二次增长的；输入规模增加时，工作量按平方量级增长。                    | `linear-tech` / linear / 线性增长的                       | 对照规模增长的平方量级和线性量级。         | 这是复杂度增长趋势，不是某次实测耗时的比例。                     |
| R04  | `property-tech` / property / 属性；以键标识的对象成员。                                              | `accessor-tech` / accessor / 属性访问器                   | 区分对象成员和属性访问机制。               | 访问器是属性的一种实现方式，不等同于所有属性。                   |
| R05  | `parameter-tech` / parameter / 参数；函数或泛型定义中用于接收实参的占位名称。                        | `type-argument-tech` / type argument / 类型实参           | 对照参数占位与类型实参。                   | parameter 可指值参数或类型参数；不能把类型实参与值实参混为一谈。 |

## 查证来源

以下官方资料用于校对概念边界，不用于证明历史 PR 当前状态。复述采用本计划自己的简短表述。来源文档中的引用和本轮独立查证分开记录。

| 编号 | 来源标题和 URL                                                                                                                                                                                        | 用途                            |
|:-----|:------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|:--------------------------------|
| D01  | [TypeScript Handbook: Type Compatibility / A Note on Soundness](https://www.typescriptlang.org/docs/handbook/type-compatibility.html#a-note-on-soundness)                                             | soundness                       |
| D02  | [TypeScript Handbook: Mapped Types](https://www.typescriptlang.org/docs/handbook/2/mapped-types.html)                                                                                                 | mapped type                     |
| D03  | [TypeScript 4.7: Optional Variance Annotations for Type Parameters](https://www.typescriptlang.org/docs/handbook/release-notes/typescript-4-7.html#optional-variance-annotations-for-type-parameters) | variance                        |
| D04  | [MDN: Object.getOwnPropertyDescriptor()](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Object/getOwnPropertyDescriptor)                                            | property descriptor             |
| D05  | [MDN: Enumerability and ownership of properties](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Guide/Enumerability_and_ownership_of_properties)                                             | enumerable / own property       |
| D06  | [MDN: Object initializer](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Operators/Object_initializer)                                                                             | prototype setter / shorthand    |
| D07  | [TypeScript Handbook: Declaration Merging](https://www.typescriptlang.org/docs/handbook/declaration-merging.html)                                                                                     | namespace / global augmentation |
| D08  | [TypeScript Handbook: Enums](https://www.typescriptlang.org/docs/handbook/enums.html)                                                                                                                 | enum                            |
| D09  | [MDN: Symbol](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Symbol)                                                                                                | Symbol                          |
| D10  | [MDN: BigInt](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/BigInt)                                                                                                | BigInt                          |
| D11  | [MDN: Indexed collections](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Guide/Indexed_collections)                                                                                         | array hole                      |

## 数量与来源记账

| 指标                   | 数量 | 计数口径                                                     |
|:-----------------------|:-----|:-------------------------------------------------------------|
| 来源文件               | 37   | 35 份 Markdown + 2 个二进制素材                              |
| 读取的文本来源         | 35   | 全部 Markdown，无不可读文件                                  |
| 本批选中的规范学习目标 | 87   | 新增和复用稳定 ID 的并集                                     |
| 来源形式记录           | 87   | 每个候选保留一个主引原文形式；不统计全目录所有出现次数       |
| 来源支持的学习目标     | 87   | 包含明确标注的语境归一化；没有来源之外额外添加的目标         |
| 新增词条候选           | 52   | N01–N52                                                      |
| 复用词条               | 35   | U01–U35，保留原内容                                          |
| 拟更新既有词条         | 0    | 不追加 Notes，也不改录音                                     |
| 明确保留的基础子项     | 5    | compile、compiler、type、property、parameter；已计入新增总数 |
| 来源之外的额外组件     | 0    | 不因先修图补齐而扩充名单                                     |
| 本批助记               | 3    | M01–M03，附属内容，不增加卡数                                |
| 本批关联               | 5    | R01–R05，关系不计为词条                                      |
| 实施后的词库总数       | 695  | 643 + 52 = 695；英文 453，日文 242                           |

这些数字是筛选结果，不是配额。未把来源目录按词数穷举，也不声称所有普通词均已入表。上面 N/U 行完成主引来源到学习目标的对应；归组和排除规则解释其余常见组合。

## 音频与内容边界

- 默认不生成音频；不添加尚不存在的 audio 引用。若明确改为生成，再检查 macOS、say、lame、mpg123 和匹配英文的已安装语音。用现有单 ID 命令从 word 朗读，不覆盖既有文件，不安装工具或改用远程服务。
- 新增项暂缺 IPA。文本内容获批不等于发音已查证；不能用 TTS 文件替代 IPA 证据。
- N01–N52 已按本次实施指令加入词库；复用项 U01–U35 未改动。
- 仅写入本计划列出的 M01–M03 和 R01–R05；未扩展其他助记或关系。
- `array hole`、`property descriptor` 和 `type instantiation` 按已列出的语境归一化保留。`structural check` 和 `prototype setter` 采用明确限定的源文含义。
- 运行时和类型层面的语义区别保留在 Notes；不顺带创建协变、逆变、原型污染或每个编译阶段的新卡。

## 获批后的实施与验证

1. 再读当前 AGENTS、schema 和词库，检查 Git 状态及稳定 ID 冲突。若词库变化，先更新复用决定和数量。
2. 保存全部既有词条原文及录音哈希。按批准后的 N 行创建 YAML；保留既有词条及录音。完成新增组件后，再引用其稳定 ID。
3. 只写入获批的助记和关联。先修、关联、标签分开处理；不增加源字段，也不修改导出器或模板。
4. 音频默认跳过。若后续获批，使用既有生成器逐项生成或复用；解码、人工听审和真实 Anki 播放分别记录。
5. 执行下面的检查。两次 build 间复制或哈希比较生成文件，确认字节一致。不得手工修改 dist。

```bash
go test ./...
go vet ./...
go run scripts/validate/main.go
go run scripts/build/main.go
git diff --check
```

6. 若全部 52 个新增项获批，核对 695 条十七列数据，英文 453 条、日文 242 条。检查 ID 首列、派生 lang-en 标签、空音频列及所有先修顺序。关联数量仅按获批关系计算。
7. 对比既有词条字节和录音哈希，确认没有附带修改。单独检查新文件，因为 git diff --check 不涵盖未跟踪文件。报告实际完成、失败或待办项，不提交、推送或操作 Anki。

## 计划阶段检查

- 已检查全部来源文档、现有稳定 ID、同词含义、主引原文与行号、候选计数和先修目标。
- 已检查本计划的本地文件链接、格式和新文件内容；既有工作与来源文档保持不变。
- 本轮未运行导出、测试或音频命令。上面的命令是获批后的步骤，不是已完成的验证记录。
- 上述检查描述最初的计划创建阶段；后续实施和验证见下节。

## 实施记录 (2026-11-03)

用户要求实施本文计划，本次完成全部 N01–N52、M01–M03 和 R01–R05。没有额外收集词汇，也没有执行来源中的调度、提交或发布指令。

- 新增 52 个英文 YAML 文件，见 N 表的源文件链接；复用 35 个既有词条，未更新其内容。
- 按计划写入 Meaning、英文 Example、词性、优先级、标签和 28 条先修链接。把 Notes 中的拟稿指令转为学习说明，保留含义、来源路径、行号、原文形式及查证标题和 URL。
- 3 条 mnemonic 与 M 表文案完全一致。5 组关联各声明一次，均在新增文件中；原有 55 组关系保持不变，全库现为 60 组。
- 52 个新增项均省略 Context、Pronunciation、Audio、Etymology 和 WordFormation。音频默认不生成，IPA 未作逐词查证；这些是明确保留的内容边界。
- 词库实际为 695 项: 英文 453 项、日文 242 项。导出为 695 条十七列数据，ID 在首列，新增项均带 lang-en 标签和 en_US 示例 TTS locale；新增 Audio 列为空。

验证结果:

- `go test ./...` 通过，Go 报告已有测试缓存命中；脚本目录没有测试文件。
- `go vet ./...` 通过。
- `go run scripts/validate/main.go` 通过，共 695 项；6 条既有同词警告未增加，分别涉及 artifact、harness、literal、resolve、revision 和 scope。
- 两次 `go run scripts/build/main.go` 均通过，导出字节完全一致。
- 核对全库 258 条先修链接，均排在目标之前；进一步核对全部行的优先级与稳定 ID 决胜顺序。
- 核对全库 RelatedWords 与源文件的直接双向邻居一致，包括新增 5 组关系的两个方向。未新增自定义 Anki 字段，也未修改导出器或模板。
- 所有原有 643 个词条字节保持不变；643 个既有 MP3 哈希及 media 中其他文件保持不变。实现、依赖、README、AGENTS、来源文档和已有 prompt 修改均保持原样。
- 新增文件单独检查 UTF-8、LF、末尾换行、无尾随空格、字段内容及文件名；计划的本地源文件链接也已检查。`git diff --check` 通过。

导出 SHA-256: `136024d8a3aa3b22431d473731fd2c115c724000964005cedbde90d88d8ff62d`。

本次没有生成、替换或试听录音，没有进行真实 Anki 导入、显示、TTS 或播放检查。已有录音的字节保留检查不等于本次重新解码或听审。没有暂存、提交、推送或修改 Anki 集合。
