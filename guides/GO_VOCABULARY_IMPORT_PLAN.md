# Go vocabulary import plan

**Status** — Approved and implemented on 2026-11-19.

This batch introduces English Go vocabulary using the `go` tag. It covers everyday language features, concurrency, modules, and testing. The batch starts from 1,368 entries and adds 57 learning targets. It also reuses 18 existing entries, for 75 tagged targets in all. Existing priorities and prerequisites on reused cards stay as written. The tag lets you find the topic in Anki with `tag:go`.

Priorities and prerequisite links guide new-card introductions. Follow the [learning-order guide](LEARNING_ORDER.md#apply-the-order-in-anki) when applying ranks to Anki. Use the [entry schema](ENTRY_SCHEMA.md) for future additions.

## Source inventory

- S1 — [Go specification — Types](https://go.dev/ref/spec#Types)
- S2 — [Go specification — Variables](https://go.dev/ref/spec#Variables)
- S3 — [Go specification](https://go.dev/ref/spec)
- S4 — [Go specification — Short variable declarations](https://go.dev/ref/spec#Short_variable_declarations)
- S5 — [Go specification — Built-in functions](https://go.dev/ref/spec#Built-in_functions)
- S6 — [Go specification — For statements with range clause](https://go.dev/ref/spec#For_range)
- S7 — [Go specification — Function types](https://go.dev/ref/spec#Function_types)
- S8 — [Go specification — Method sets](https://go.dev/ref/spec#Method_sets)
- S9 — [Go specification — Interface types](https://go.dev/ref/spec#Interface_types)
- S10 — [Go specification — Type switches](https://go.dev/ref/spec#Type_switches)
- S11 — [Go specification — Type parameter declarations](https://go.dev/ref/spec#Type_parameter_declarations)
- S12 — [Error handling and Go](https://go.dev/blog/error-handling-and-go)
- S13 — [Working with Errors in Go 1.13](https://go.dev/blog/go1.13-errors)
- S14 — [Go specification — Defer statements](https://go.dev/ref/spec#Defer_statements)
- S15 — [Go specification — Handling panics](https://go.dev/ref/spec#Handling_panics)
- S16 — [Go specification — Go statements](https://go.dev/ref/spec#Go_statements)
- S17 — [Go specification — Channel types](https://go.dev/ref/spec#Channel_types)
- S18 — [Go specification — Select statements](https://go.dev/ref/spec#Select_statements)
- S19 — [Go sync package](https://pkg.go.dev/sync)
- S20 — [Go Concurrency Patterns: Context](https://go.dev/blog/context)
- S21 — [The Go Memory Model](https://go.dev/ref/mem)
- S22 — [Create a Go module](https://go.dev/doc/tutorial/create-module)
- S23 — [go.mod file reference](https://go.dev/doc/modules/gomod-ref)
- S24 — [Managing dependencies](https://go.dev/doc/modules/managing-dependencies)
- S25 — [Go Wiki — TableDrivenTests](https://go.dev/wiki/TableDrivenTests)
- S26 — [Go testing package](https://pkg.go.dev/testing)
- S27 — [Go Fuzzing](https://go.dev/doc/security/fuzz/)
- S28 — [Data Race Detector](https://go.dev/doc/articles/race_detector)

## New entries

Each linked entry records the meaning, an English example, and its source. The tables record priority, direct prerequisites, and source coverage. A dash indicates that a card has no declared prerequisites.

### Values and declarations

| Word                         | Stable ID                                                                     | Chinese meaning                                            | Priority | Prerequisites   | Source |
| :--------------------------- | :---------------------------------------------------------------------------- | :--------------------------------------------------------- | :------- | :-------------- | :----- |
| `bool`                       | [bool-go](../entries/bool-go.yml)                                             | 布尔类型，取值为 true 或 false                             | 1        | `type-tech`     | S1     |
| `int`                        | [int-go](../entries/int-go.yml)                                               | 有符号整数类型，在 Go 中为 32 位或 64 位                   | 1        | `type-tech`     | S1     |
| `byte`                       | [byte-go](../entries/byte-go.yml)                                             | 字节，Go 中 uint8 的别名                                   | 1        | `int-go`        | S1     |
| `rune`                       | [rune-go](../entries/rune-go.yml)                                             | Go 中 int32 的别名，通常用来表示 Unicode 码点              | 1        | `int-go`        | S1     |
| `variable`                   | [variable-go](../entries/variable-go.yml)                                     | 变量，具有类型且可以存储值的位置                           | 1        | —               | S2     |
| `constant`                   | [constant-go](../entries/constant-go.yml)                                     | 常量，在编译时确定的值                                     | 1        | —               | S3     |
| `zero value`                 | [zero-value-go](../entries/zero-value-go.yml)                                 | 零值，未显式初始化时由类型决定的默认值                     | 1        | `type-tech`     | S2     |
| `nil`                        | [nil-go](../entries/nil-go.yml)                                               | 表示某些类型零值的预声明标识符，例如指针、切片和通道的零值 | 1        | `zero-value-go` | S2     |
| `short variable declaration` | [short-variable-declaration-go](../entries/short-variable-declaration-go.yml) | 在函数内使用 := 声明并初始化变量的语法                     | 2        | `variable-go`   | S4     |

### Functions and data structures

| Word                     | Stable ID                                                             | Chinese meaning                                              | Priority | Prerequisites        | Source |
| :----------------------- | :-------------------------------------------------------------------- | :----------------------------------------------------------- | :------- | :------------------- | :----- |
| `slice`                  | [slice-go](../entries/slice-go.yml)                                   | 切片，引用底层数组中一段连续元素的描述符                     | 1        | `array-tech`         | S1     |
| `map`                    | [map-go](../entries/map-go.yml)                                       | 映射，按唯一键存取值的无序集合                               | 1        | `type-tech`          | S1     |
| `pointer`                | [pointer-go](../entries/pointer-go.yml)                               | 指针，保存变量地址的值                                       | 1        | `type-tech`          | S1     |
| `struct`                 | [struct-go](../entries/struct-go.yml)                                 | 结构体，将具名字段组合为一个类型                             | 1        | `type-tech`          | S1     |
| `length`                 | [length-go](../entries/length-go.yml)                                 | 长度，表示切片等值当前包含的元素数                           | 1        | `slice-go`           | S5     |
| `capacity`               | [capacity-go](../entries/capacity-go.yml)                             | 容量，对切片表示在不更换底层数组的情况下允许扩展到的最大长度 | 2        | `length-go`          | S5     |
| `backing array`          | [backing-array-go](../entries/backing-array-go.yml)                   | 切片所引用的底层数组                                         | 2        | `slice-go`           | S1     |
| `append`                 | [append-go](../entries/append-go.yml)                                 | 向切片追加元素并返回结果切片的内置函数                       | 2        | `slice-go`           | S5     |
| `make`                   | [make-go](../entries/make-go.yml)                                     | 用于创建切片、映射和通道的内置函数，返回初始化后的值         | 2        | `slice-go`, `map-go` | S5     |
| `new`                    | [new-go](../entries/new-go.yml)                                       | 分配存储并返回指针的内置函数                                 | 2        | `pointer-go`         | S5     |
| `range`                  | [range-go](../entries/range-go.yml)                                   | 用于按规则遍历值的 for 子句关键字                            | 2        | `slice-go`           | S6     |
| `multiple return values` | [multiple-return-values-go](../entries/multiple-return-values-go.yml) | 多返回值，一次函数调用返回的多个结果                         | 2        | `return-type-tech`   | S7     |
| `named result parameter` | [named-result-parameter-go](../entries/named-result-parameter-go.yml) | 命名结果参数，在函数签名中声明名称的返回参数                 | 2        | `return-type-tech`   | S7     |

### Methods and interfaces

| Word                      | Stable ID                                                               | Chinese meaning                                            | Priority | Prerequisites               | Source |
| :------------------------ | :---------------------------------------------------------------------- | :--------------------------------------------------------- | :------- | :-------------------------- | :----- |
| `receiver`                | [receiver-go](../entries/receiver-go.yml)                               | 接收者，在方法声明中指定方法所关联类型的参数               | 2        | `method-tech`               | S8     |
| `pointer receiver`        | [pointer-receiver-go](../entries/pointer-receiver-go.yml)               | 指针接收者，类型为指针的方法接收者                         | 2        | `receiver-go`, `pointer-go` | S8     |
| `method set`              | [method-set-go](../entries/method-set-go.yml)                           | 方法集，由类型及相关规则决定的方法集合                     | 2        | `receiver-go`               | S8     |
| `interface`               | [interface-go](../entries/interface-go.yml)                             | 接口，通过类型集规定哪些类型符合要求，常用一组方法描述行为 | 2        | `method-set-go`             | S9     |
| `implicit implementation` | [implicit-implementation-go](../entries/implicit-implementation-go.yml) | 隐式实现，类型满足接口要求即可实现该接口                   | 2        | `interface-go`              | S9     |
| `embedding`               | [embedding-go](../entries/embedding-go.yml)                             | 嵌入，在结构体中声明嵌入字段或在接口中包含其他接口         | 2        | `struct-go`                 | S1     |
| `type switch`             | [type-switch-go](../entries/type-switch-go.yml)                         | 类型选择，根据接口值的动态类型选择分支                     | 2        | `interface-go`              | S10    |
| `type constraint`         | [type-constraint-go](../entries/type-constraint-go.yml)                 | 类型约束，限定类型形参可以接受的类型实参                   | 2        | `type-parameter-typescript` | S11    |
| `comparable`              | [comparable-go](../entries/comparable-go.yml)                           | 预声明约束，表示类型实参必须满足 Go 的可比较性要求         | 3        | `type-constraint-go`        | S9     |

### Errors and cleanup

| Word             | Stable ID                                             | Chinese meaning                                           | Priority | Prerequisites          | Source |
| :--------------- | :---------------------------------------------------- | :-------------------------------------------------------- | :------- | :--------------------- | :----- |
| `error`          | [error-go](../entries/error-go.yml)                   | 内置错误接口，要求实现 Error 方法来提供错误描述           | 1        | `type-tech`            | S12    |
| `error wrapping` | [error-wrapping-go](../entries/error-wrapping-go.yml) | 错误包装，添加上下文并保留原错误供后续检查                | 2        | `error-go`             | S13    |
| `sentinel error` | [sentinel-error-go](../entries/sentinel-error-go.yml) | 哨兵错误，用预定义错误值表示可识别的情况                  | 2        | `error-go`             | S13    |
| `defer`          | [defer-go](../entries/defer-go.yml)                   | 延迟调用，在外围函数返回前按后进先出顺序运行已安排的调用  | 2        | `function-tech`        | S14    |
| `panic`          | [panic-go](../entries/panic-go.yml)                   | 触发 panic，停止当前正常执行并开始展开调用栈              | 2        | `function-tech`        | S15    |
| `recover`        | [recover-go](../entries/recover-go.yml)               | 在延迟调用的函数中直接使用时可接管同一 goroutine 的 panic | 3        | `panic-go`, `defer-go` | S15    |

### Concurrency

| Word               | Stable ID                                                 | Chinese meaning                                                    | Priority | Prerequisites               | Source |
| :----------------- | :-------------------------------------------------------- | :----------------------------------------------------------------- | :------- | :-------------------------- | :----- |
| `goroutine`        | [goroutine-go](../entries/goroutine-go.yml)               | Go 运行时管理的并发执行单元                                        | 2        | `function-tech`             | S16    |
| `channel`          | [channel-go](../entries/channel-go.yml)                   | 通道，用于在并发执行单元之间发送和接收指定类型的值                 | 2        | `goroutine-go`              | S17    |
| `buffered channel` | [buffered-channel-go](../entries/buffered-channel-go.yml) | 缓冲通道，有容量可暂存待接收的值                                   | 2        | `channel-go`, `capacity-go` | S17    |
| `select`           | [select-go](../entries/select-go.yml)                     | 选择语句，从可执行的通道发送或接收操作中选择一项                   | 2        | `channel-go`                | S18    |
| `mutex`            | [mutex-go](../entries/mutex-go.yml)                       | 互斥锁，用于协调对共享资源的访问                                   | 2        | `goroutine-go`              | S19    |
| `wait group`       | [wait-group-go](../entries/wait-group-go.yml)             | 等待组，等待一组任务完成的同步机制                                 | 2        | `goroutine-go`              | S19    |
| `deadline`         | [deadline-go](../entries/deadline-go.yml)                 | 截止时间，指定操作最晚应结束的时刻                                 | 2        | —                           | S20    |
| `data race`        | [data-race-go](../entries/data-race-go.yml)               | 数据竞争，针对同一内存位置发生未正确同步且至少有一次写入的并发访问 | 2        | `goroutine-go`              | S21    |
| `deadlock`         | [deadlock-go](../entries/deadlock-go.yml)                 | 死锁，相关执行单元一直等待而无法继续                               | 3        | `goroutine-go`              | S21    |
| `happens-before`   | [happens-before-go](../entries/happens-before-go.yml)     | 先行发生关系，用于描述按内存模型可观测的操作顺序                   | 3        | `goroutine-go`              | S21    |

### Project work and testing

| Word                | Stable ID                                                   | Chinese meaning                                        | Priority | Prerequisites                 | Source |
| :------------------ | :---------------------------------------------------------- | :----------------------------------------------------- | :------- | :---------------------------- | :----- |
| `package`           | [package-go](../entries/package-go.yml)                     | 包，在同一目录中共同构建的一组 Go 源文件，通常共享包名 | 1        | —                             | S22    |
| `Go module`         | [go-module](../entries/go-module.yml)                       | Go 模块，一组一起发布和管理版本的 Go 包                | 2        | `package-go`                  | S22    |
| `module path`       | [module-path-go](../entries/module-path-go.yml)             | 模块路径，标识模块并构成其包导入路径的前缀             | 2        | `go-module`                   | S23    |
| `dependency`        | [dependency-go](../entries/dependency-go.yml)               | 依赖，项目需要使用的其他软件组件                       | 2        | `go-module`                   | S24    |
| `go.mod`            | [go-mod-go](../entries/go-mod-go.yml)                       | Go 模块定义文件，记录模块路径和依赖等信息              | 2        | `go-module`, `module-path-go` | S23    |
| `go.sum`            | [go-sum-go](../entries/go-sum-go.yml)                       | 记录依赖模块内容和 go.mod 文件校验值的文件             | 2        | `go-mod-go`                   | S24    |
| `table-driven test` | [table-driven-test-go](../entries/table-driven-test-go.yml) | 表驱动测试，用一组数据描述并运行多个测试用例           | 2        | `test-tech`                   | S25    |
| `subtest`           | [subtest-go](../entries/subtest-go.yml)                     | 子测试，在一个测试内用 Run 组织的具名测试              | 2        | `test-tech`                   | S26    |
| `fuzz testing`      | [fuzz-testing-go](../entries/fuzz-testing-go.yml)           | 模糊测试，自动变异输入以查找意外行为                   | 3        | `test-tech`                   | S27    |
| `race detector`     | [race-detector-go](../entries/race-detector-go.yml)         | 竞争检测器，检查运行时执行路径中的数据竞争             | 3        | `data-race-go`, `test-tech`   | S28    |

## Reused entries

The shared cards received the `go` tag.

| Word             | Stable ID                                                             | Priority | Existing prerequisites       |
| :--------------- | :-------------------------------------------------------------------- | :------- | :--------------------------- |
| `type`           | [type-tech](../entries/type-tech.yml)                                 | 1        | —                            |
| `string`         | [string-typescript](../entries/string-typescript.yml)                 | 1        | `type-tech`                  |
| `function`       | [function-tech](../entries/function-tech.yml)                         | 1        | —                            |
| `array`          | [array-tech](../entries/array-tech.yml)                               | 1        | —                            |
| `parameter`      | [parameter-tech](../entries/parameter-tech.yml)                       | 1        | `function-tech`              |
| `argument`       | [argument-tech](../entries/argument-tech.yml)                         | 1        | `function-tech`              |
| `return type`    | [return-type-tech](../entries/return-type-tech.yml)                   | 2        | `type-tech`, `function-tech` |
| `method`         | [method-tech](../entries/method-tech.yml)                             | 1        | —                            |
| `generic`        | [generic-typescript](../entries/generic-typescript.yml)               | 2        | `type-tech`                  |
| `type parameter` | [type-parameter-typescript](../entries/type-parameter-typescript.yml) | 2        | `generic-typescript`         |
| `type argument`  | [type-argument-tech](../entries/type-argument-tech.yml)               | 3        | `type-parameter-typescript`  |
| `cancellation`   | [cancellation-tech](../entries/cancellation-tech.yml)                 | 2        | —                            |
| `worker pool`    | [worker-pool-concurrency](../entries/worker-pool-concurrency.yml)     | 3        | —                            |
| `compile`        | [compile-tech](../entries/compile-tech.yml)                           | 1        | —                            |
| `compiler`       | [compiler-tech](../entries/compiler-tech.yml)                         | 1        | `compile-tech`               |
| `toolchain`      | [toolchain-tech](../entries/toolchain-tech.yml)                       | 2        | —                            |
| `test`           | [test-tech](../entries/test-tech.yml)                                 | 1        | —                            |
| `benchmark`      | [benchmark-tech](../entries/benchmark-tech.yml)                       | 2        | —                            |

## Grouping decisions

- Store Go’s `bool` spelling on its own card. Explain functions without results in Notes and reuse the shared cards for *string* and *function*.
- Preserve distinct senses for `map`, `interface`, `defer`, `recover`, and `select`. Use *Go module* for module versioning and distribution. Teach it after *package*.
- Defer a reference-wide pass over escape analysis, reflection, build constraints, workspaces, and runtime scheduling. The batch contains no new associations.

## Implementation and verification

Go tests, vet, vocabulary validation, and diff checks passed. Two builds produced identical bytes. The export contains 1,425 rows with 17 columns. Each of the 619 declared prerequisites has a lower `StudyOrder` than its dependent.

Verified US IPA is present in 23 new entries. The other 34 await a reliable pronunciation source. Existing MP3 hashes matched the pre-import snapshot. The batch generated no audio.

Validation reported 21 same-word warnings. The 16 existing warnings remain alongside the five intentional Go sense pairs listed above.
