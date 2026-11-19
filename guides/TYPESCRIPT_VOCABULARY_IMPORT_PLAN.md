# TypeScript vocabulary import plan

**Status** — Approved and implemented on 2026-11-19. See the [implementation record](#implementation-record).

Build a core-first English topic using the existing `typescript` tag. Start with types and functions, then introduce type checking, narrowing, generics, modules, and selected type transformations. The tiers describe introduction order within the topic. See the [learning-order guide](LEARNING_ORDER.md).

## Scope and accounting

The checkout contains 1,320 vocabulary entries. This plan proposes **66 distinct learning targets**, with **48 new entries** and **18 existing entries reused**. Of the reused entries, 16 have proposed metadata updates, and 2 already fit as written. The 14 tag additions are included in the metadata update count.

The Handbook pages below are reference documents. Rows in the candidate tables count selected learning targets, so overlapping coverage across pages counts once. The four user examples are included in the foundations. Specialized vocabulary already tagged `typescript`, such as *variance* and *global augmentation*, remains available as later study material outside this plan’s accounting.

New entries use `language: en`, concise Chinese meanings, and `typescript` plus appropriate existing tags. Use the proposed stable ID as the new YAML filename stem. A dash in a prerequisite column means an empty list for planning purposes. Omit the field if no dependency is needed. Prerequisites that point to proposed new IDs require those entries to exist before validation. See the [entry schema](ENTRY_SCHEMA.md) and [maintenance prompt](MAINTENANCE_PROMPTS.md#collect-vocabulary-from-sources).

## Source inventory

- S1 — [Everyday Types](https://www.typescriptlang.org/docs/handbook/2/everyday-types.html)
- S2 — [The Basics](https://www.typescriptlang.org/docs/handbook/2/basic-types.html)
- S3 — [Narrowing](https://www.typescriptlang.org/docs/handbook/2/narrowing.html)
- S4 — [More on Functions](https://www.typescriptlang.org/docs/handbook/2/functions.html)
- S5 — [Object Types](https://www.typescriptlang.org/docs/handbook/2/objects.html)
- S6 — [Generics](https://www.typescriptlang.org/docs/handbook/2/generics.html)
- S7 — [Modules](https://www.typescriptlang.org/docs/handbook/2/modules.html)
- S8 — [Creating Types from Types](https://www.typescriptlang.org/docs/handbook/2/types-from-types.html)
- S9 — [Keyof Type Operator](https://www.typescriptlang.org/docs/handbook/2/keyof-types.html)
- S10 — [Typeof Type Operator](https://www.typescriptlang.org/docs/handbook/2/typeof-types.html)
- S11 — [Indexed Access Types](https://www.typescriptlang.org/docs/handbook/2/indexed-access-types.html)
- S12 — [Conditional Types](https://www.typescriptlang.org/docs/handbook/2/conditional-types.html)
- S13 — [Template Literal Types](https://www.typescriptlang.org/docs/handbook/2/template-literal-types.html)
- S14 — [Utility Types](https://www.typescriptlang.org/docs/handbook/utility-types.html)
- S15 — [Type Compatibility](https://www.typescriptlang.org/docs/handbook/type-compatibility.html)

## New candidates

Priority 1 covers 12 foundations, priority 2 covers 25 applied items, and priority 3 covers 11 specialized items. Each table connects a source group to concrete entry proposals. Examples are original drafts.

### Values and basic types

Source coverage for these candidates comes from S1, S2, S4.

| Word           | Proposed ID               | Chinese meaning                    | Priority | Prerequisites         | Proposed example                                        |
| :------------- | :------------------------ | :--------------------------------- | :------- | :-------------------- | :------------------------------------------------------ |
| `boolean`      | `boolean-typescript`      | 布尔类型，值为 true 或 false       | 1        | `type-tech`           | The enabled flag has type boolean.                      |
| `string`       | `string-typescript`       | 字符串类型                         | 1        | `type-tech`           | The name parameter has type string.                     |
| `number`       | `number-typescript`       | 数值类型                           | 1        | `type-tech`           | The count variable has type number.                     |
| `null`         | `null-typescript`         | 表示有意缺省的空值及其类型         | 1        | —                     | The lookup returns null when the record is missing.     |
| `undefined`    | `undefined-typescript`    | 未定义值及其类型                   | 1        | —                     | An omitted optional argument can be undefined.          |
| `object`       | `object-typescript`       | 非原始值的类型                     | 1        | `type-tech`           | The object type accepts arrays and functions.           |
| `tuple`        | `tuple-typescript`        | 元组，按位置定义元素类型的数组类型 | 1        | `array-tech`          | This tuple stores a name followed by a count.           |
| `literal type` | `literal-type-typescript` | 字面量类型，表示某个特定字面值     | 2        | `literal-programming` | The literal type "ready" allows just that string value. |

### Functions

Source coverage for these candidates comes from S1, S3, S4.

| Word                 | Proposed ID                     | Chinese meaning                          | Priority | Prerequisites                        | Proposed example                                                          |
| :------------------- | :------------------------------ | :--------------------------------------- | :------- | :----------------------------------- | :------------------------------------------------------------------------ |
| `function`           | `function-tech`                 | 函数，可调用并执行计算或操作的代码单元   | 1        | —                                    | This function doubles its input.                                          |
| `argument`           | `argument-tech`                 | 实参，调用时传给函数的值                 | 1        | `function-tech`                      | We pass 42 as the second argument.                                        |
| `function signature` | `function-signature-typescript` | 函数签名，描述可调用形式的参数和返回类型 | 2        | `parameter-tech`, `return-type-tech` | The function signature shows what callers can pass and what they receive. |
| `optional parameter` | `optional-parameter-typescript` | 可选形参，调用时可以省略的参数           | 2        | `parameter-tech`                     | The optional parameter lets callers supply a label.                       |
| `rest parameter`     | `rest-parameter-typescript`     | 剩余参数，将剩余实参收集为数组的形参     | 2        | `parameter-tech`, `array-tech`       | The rest parameter collects the remaining arguments into an array.        |
| `void`               | `void-typescript`               | 表示调用方不应使用返回值的类型           | 1        | `function-tech`                      | The callback has a void return type.                                      |

### Describing and checking types

Source coverage for these candidates comes from S1, S2, S5, S15.

| Word                | Proposed ID                    | Chinese meaning                              | Priority | Prerequisites              | Proposed example                                                              |
| :------------------ | :----------------------------- | :------------------------------------------- | :------- | :------------------------- | :---------------------------------------------------------------------------- |
| `type annotation`   | `type-annotation-typescript`   | 类型注解，写在代码中的显式类型标记           | 1        | `type-tech`                | The type annotation specifies the shape of the input.                         |
| `type inference`    | `type-inference-typescript`    | 类型推断，根据代码和上下文自动确定类型       | 1        | `type-tech`                | Type inference determines the type from the initializer.                      |
| `type alias`        | `type-alias-typescript`        | 类型别名，为类型指定名称的声明               | 2        | `type-tech`                | Use a type alias to name a union.                                             |
| `interface`         | `interface-typescript`         | 接口，用来描述对象形状的声明                 | 2        | `property-tech`            | The interface describes the fields required by the API.                       |
| `optional property` | `optional-property-typescript` | 可选属性，对象中可以缺省的属性               | 2        | `property-tech`            | An optional property can be absent from the response.                         |
| `readonly`          | `readonly-typescript`          | 只读修饰符，限制通过该类型写入属性           | 2        | `property-tech`            | The readonly modifier prevents reassignment through this reference.           |
| `intersection type` | `intersection-type-typescript` | 交叉类型，要求值同时满足各组成类型           | 2        | `type-tech`                | The intersection type requires fields from both inputs.                       |
| `assignability`     | `assignability-typescript`     | 可赋值性，源类型能否用于目标类型所要求的位置 | 2        | `type-tech`                | Assignability determines whether this value can be passed to the parameter.   |
| `structural typing` | `structural-typing-typescript` | 结构类型系统，按类型结构判断兼容性           | 2        | `assignability-typescript` | Structural typing allows objects with compatible members to be used together. |
| `any`               | `any-typescript`               | 允许跳过相关静态类型检查的类型               | 2        | `type-tech`                | A value typed as any can be used without checking its properties first.       |
| `unknown`           | `unknown-typescript`           | 可接收任意值，使用前需确认具体类型的类型     | 2        | `type-tech`                | Treat the external input as unknown until you check its shape.                |
| `never`             | `never-typescript`             | 无值可属于的类型                             | 2        | `type-tech`                | A function that always throws can have a never return type.                   |

### Narrowing and generics

Source coverage for these candidates comes from S3, S4, S6.

| Word                      | Proposed ID                          | Chinese meaning                                | Priority | Prerequisites                                        | Proposed example                                                           |
| :------------------------ | :----------------------------------- | :--------------------------------------------- | :------- | :--------------------------------------------------- | :------------------------------------------------------------------------- |
| `narrowing`               | `narrowing-typescript`               | 类型收窄，根据检查和控制流缩小可能的类型       | 2        | `union-tech`                                         | Narrowing lets us treat the value as a string after the check.             |
| `type guard`              | `type-guard-typescript`              | 类型守卫，用于帮助收窄类型的检查               | 2        | `narrowing-typescript`                               | The typeof check acts as a type guard.                                     |
| `type predicate`          | `type-predicate-typescript`          | 类型谓词，用返回类型形式声明函数判定的类型条件 | 2        | `type-guard-typescript`                              | The type predicate tells the compiler what a successful check establishes. |
| `discriminated union`     | `discriminated-union-typescript`     | 可辨识联合，用共同的字面量属性区分成员         | 2        | `union-tech`, `literal-type-typescript`              | The discriminated union uses a kind field to distinguish its cases.        |
| `exhaustiveness checking` | `exhaustiveness-checking-typescript` | 穷尽性检查，确认代码覆盖了所有可能的情况       | 2        | `discriminated-union-typescript`, `never-typescript` | Exhaustiveness checking catches a case missing from the switch.            |
| `generic`                 | `generic-typescript`                 | 泛型，使用类型参数描述可复用的函数或类型       | 2        | `type-tech`                                          | A generic function can preserve the type of its input.                     |
| `type parameter`          | `type-parameter-typescript`          | 类型形参，在泛型定义中代表待确定的类型         | 2        | `generic-typescript`                                 | The type parameter connects the element type to the result.                |
| `generic constraint`      | `generic-constraint-typescript`      | 泛型约束，限制类型参数可接受的类型             | 2        | `type-parameter-typescript`                          | The generic constraint requires a length property.                         |

### Modules and type manipulation

Source coverage for these candidates comes from S7–S14.

| Word                    | Proposed ID                          | Chinese meaning                                  | Priority | Prerequisites                                             | Proposed example                                                     |
| :---------------------- | :----------------------------------- | :----------------------------------------------- | :------- | :-------------------------------------------------------- | :------------------------------------------------------------------- |
| `module`                | `module-typescript`                  | 模块，具有独立作用域并可显式导入或导出的代码单元 | 2        | —                                                         | The module exports a parser and its configuration type.              |
| `type-only import`      | `type-only-import-typescript`        | 仅类型导入，导入供类型检查使用的绑定             | 2        | `module-typescript`                                       | Use a type-only import for the options interface.                    |
| `declaration file`      | `declaration-file-typescript`        | 声明文件，为代码提供类型信息的 .d.ts 文件        | 2        | `module-typescript`                                       | The library ships a declaration file alongside its JavaScript.       |
| `keyof`                 | `keyof-typescript`                   | 生成类型的属性键类型的运算符                     | 3        | `property-tech`, `union-tech`                             | Use keyof to restrict the key to known properties.                   |
| `typeof type operator`  | `typeof-type-operator-typescript`    | 在类型位置获取值或属性的静态类型的运算符         | 3        | `type-inference-typescript`                               | Use the typeof type operator to reuse the type of an existing value. |
| `indexed access type`   | `indexed-access-type-typescript`     | 索引访问类型，按键取得另一个类型中的属性类型     | 3        | `property-tech`                                           | Use an indexed access type to obtain the type of the id property.    |
| `conditional type`      | `conditional-type-typescript`        | 条件类型，根据类型关系选择结果类型               | 3        | `assignability-typescript`, `generic-typescript`          | The conditional type chooses a result based on its input type.       |
| `template literal type` | `template-literal-type-typescript`   | 模板字面量类型，按字符串模板构造字符串字面量类型 | 3        | `literal-type-typescript`                                 | A template literal type describes the allowed event names.           |
| `utility type`          | `utility-type-typescript`            | 工具类型，用于常见类型变换的内置泛型类型         | 3        | `generic-typescript`                                      | The utility type derives a smaller type from the original.           |
| `Partial`               | `partial-utility-type-typescript`    | 将类型的所有属性变为可选的工具类型               | 3        | `utility-type-typescript`, `optional-property-typescript` | Use Partial when a patch may omit any field.                         |
| `Pick`                  | `pick-utility-type-typescript`       | 从类型中保留指定属性的工具类型                   | 3        | `utility-type-typescript`, `keyof-typescript`             | Use Pick to keep only the id and name fields.                        |
| `Omit`                  | `omit-utility-type-typescript`       | 从类型中移除指定属性的工具类型                   | 3        | `utility-type-typescript`, `keyof-typescript`             | Use Omit to remove the internal field from the public type.          |
| `Record`                | `record-utility-type-typescript`     | 以给定键类型和值类型构造对象类型的工具类型       | 3        | `utility-type-typescript`, `keyof-typescript`             | Use Record to map each status to a display label.                    |
| `ReturnType`            | `returntype-utility-type-typescript` | 取得函数类型的返回类型的工具类型                 | 3        | `utility-type-typescript`, `return-type-tech`             | Use ReturnType to derive the result type from the function.          |

## Existing entries to reuse

Keep the existing priority and prerequisites in this table. The added prerequisites and tags are proposed changes for the eventual import. The reused entries cover the corresponding concepts in the source groups above.

| Word                 | Existing ID                                                       | Priority | Current prerequisites             | Add prerequisites               | `typescript` tag |
| :------------------- | :---------------------------------------------------------------- | :------- | :-------------------------------- | :------------------------------ | :--------------- |
| `type`               | [type-tech](../entries/type-tech.yml)                             | 1        | —                                 | —                               | Add              |
| `array`              | [array-tech](../entries/array-tech.yml)                           | 1        | —                                 | —                               | Add              |
| `primitive`          | [primitive-javascript](../entries/primitive-javascript.yml)       | 2        | —                                 | —                               | Add              |
| `literal`            | [literal-programming](../entries/literal-programming.yml)         | 2        | —                                 | —                               | Add              |
| `BigInt`             | [bigint-javascript](../entries/bigint-javascript.yml)             | 2        | `primitive-javascript`            | —                               | Add              |
| `Symbol`             | [symbol-javascript](../entries/symbol-javascript.yml)             | 2        | `primitive-javascript`            | —                               | Add              |
| `parameter`          | [parameter-tech](../entries/parameter-tech.yml)                   | 1        | —                                 | `function-tech`                 | Add              |
| `return type`        | [return-type-tech](../entries/return-type-tech.yml)               | 2        | `type-tech`                       | `function-tech`                 | Add              |
| `property`           | [property-tech](../entries/property-tech.yml)                     | 1        | —                                 | —                               | Add              |
| `union`              | [union-tech](../entries/union-tech.yml)                           | 2        | —                                 | `type-tech`                     | Add              |
| `type assertion`     | [type-assertion-tech](../entries/type-assertion-tech.yml)         | 3        | `cast-tech`                       | —                               | Add              |
| `overload`           | [overload-tech](../entries/overload-tech.yml)                     | 2        | —                                 | `function-signature-typescript` | Add              |
| `type argument`      | [type-argument-tech](../entries/type-argument-tech.yml)           | 3        | —                                 | `type-parameter-typescript`     | Add              |
| `type instantiation` | [type-instantiation-tech](../entries/type-instantiation-tech.yml) | 3        | `type-tech`, `type-argument-tech` | —                               | Keep             |
| `module specifier`   | [module-specifier-tech](../entries/module-specifier-tech.yml)     | 2        | —                                 | `module-typescript`             | Keep             |
| `declaration emit`   | [declaration-emit-tech](../entries/declaration-emit-tech.yml)     | 3        | `emit-tech`                       | `declaration-file-typescript`   | Add              |
| `mapped type`        | [mapped-type-typescript](../entries/mapped-type-typescript.yml)   | 3        | `type-tech`, `property-tech`      | `keyof-typescript`              | Keep             |
| `enum`               | [enum-typescript](../entries/enum-typescript.yml)                 | 2        | —                                 | —                               | Keep             |

## Grouping and distinctions

- Teach lowercase `boolean`, `string`, and `number`. Explain uppercase wrapper names in Notes when useful. Reuse *primitive*, *BigInt*, and *Symbol* and explain lowercase `bigint` and `symbol` in TypeScript examples or Notes where needed. Record a proposed wording change before editing a reused entry.
- Teach lowercase `object` as the TypeScript type for non-primitive values. Use object type syntax and object shapes in the *interface* and *property* examples. Group the general idea of primitive types with the existing *primitive* card.
- Keep *literal* and *literal type* separate, and use the former as a prerequisite. Keep *parameter* and *argument* separate, as well as *type parameter* and *type argument*.
- Keep `void`, `undefined`, and `never` separate. Explain in `void` Notes that a function assigned to a callback returning `void` may produce a value which the caller ignores. Use `never` to connect impossible values with exhaustive checks.
- Explain that `readonly` restricts writes through the type. It provides neither runtime freezing nor deep immutability. Keep the type-position `typeof` with its TypeScript meaning and mention the JavaScript runtime operator in Notes.
- Use the existing *union* entry for union types, *overload* for function overloading, and *type assertion* for `as`. Keep assertion and runtime validation distinct. Each named utility in scope has a separate practical use. The case-sensitive utility `Omit` is a distinct sense from the existing verb `omit` (`omit-tech`). Expect a same-word warning for this intentional pair. Variants and extra utilities can stay in examples or wait for a later batch.
- Defer a full pass over classes, compiler flags, module resolution, declaration authoring, compiler diagnostics, and the remaining utilities. Aim for useful reading vocabulary and avoid separate cards for incidental syntax or identifiers.

## Proposed memory aids and associations

Consider a Notes comparison for `any`, `unknown`, and `never`. A value typed `any` bypasses checks on its use, `unknown` needs a check before most operations, and `never` describes a type with no possible values. Use the proposed examples to reinforce these distinctions.

This plan proposes zero `related_entries` associations and zero changes to existing mnemonic fields. Prerequisites describe introduction order.

## Import and verification after approval

1. Recheck the stable IDs, senses, and source pages against the approved rows. Record the relevant Handbook title and URL in Notes. Verify IPA separately using the [pronunciation conventions](ENTRY_SCHEMA.md#pronunciation). Leave unsupported IPA unset. Audio generation is outside this batch.
2. Add the approved YAML files and apply only approved metadata to reused entries. Keep IDs, existing recordings, and unrelated entry fields as written.
3. Run `go test ./...`, `go vet ./...`, `go run scripts/validate/main.go`, and `git diff --check`. Build twice with `go run scripts/build/main.go` and compare the resulting bytes.
4. Compare exported IDs and confirm `typescript` on the approved topic entries. Check every declared prerequisite against its dependent’s generated `StudyOrder`. Review warnings and verify preservation of original media and unrelated entry content.
5. Report the implementation results and hand off the complete export. Apply the ranks to existing new cards using the documented workflow if desired. [Apply the order in Anki](LEARNING_ORDER.md#apply-the-order-in-anki)

Planning checks cover the proposed IDs and filenames, reuse links, counts, and the combined prerequisite graph. Import, audio generation, and Anki changes require their own authorization.


## Implementation record

The approved batch added 48 entries and reused 18, with 16 reused files receiving the planned metadata changes. The export grew from 1,320 to 1,368 rows. All 66 targets carry the `typescript` tag. The tag also includes earlier TypeScript vocabulary outside this batch.

Verified US IPA is present in 13 new entries, including all four original examples. The other 35 entries await a reliable pronunciation source. All existing MP3 hashes matched the baseline. This batch adds no audio or associations.

Go tests, vet, vocabulary validation, and diff checks passed. Two builds produced identical bytes. The 17-column export contains unique ranks and places every declared prerequisite before its dependent. The one new duplicate-word warning is the approved `Omit` utility alongside the existing verb `omit`. The other 15 warnings were already present.

## Prerequisite follow-up

The tag now includes *cast*, *emit*, and lexical *scope*. This keeps the existing prerequisites of *type assertion*, *declaration emit*, and *namespace* available when studying by topic. The current `tag:typescript` filter includes 73 cards.
