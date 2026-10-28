# PR vocabulary import

This collection comes from a read-only review of 27 Markdown documents under `docs/PR`: 21 PR descriptions, four reports or plans, one scheduler document, and the contributor guide. The two image assets were not text sources. Historical statements were used as language examples, not as current project or PR status.

## Coverage and grouping

- Added 277 entries, with no fixed count limit; the collection now contains 283 entries.
- Related inflections and derivations share a headword where they express the same learning item. Related forms and sentence patterns are recorded in entry notes.
- Distinct meanings have separate IDs: work scope and lexical scope; literal text and programming literals; resolving a reference and resolving a problem; editorial revisions and Git revisions; workflow artifacts and build artifacts.
- Test assertions and type assertions are separate learning items. Opposite actions such as mount/unmount and encode/decode are also separate.
- Standalone expressions receive entries. Patterns with X/Y placeholders belong in headword notes; audio speaks only the entry's word or phrase.
- Chinese definitions and English examples are written for the indicated sense. Examples are learning examples, not claims that a historical check ran or a PR remains current. IPA is omitted unless already present in an existing entry.

## Existing entries reused

- `fallback-tech`: covers fallback and the verb pattern "fall back to X"; existing notes already distinguish noun and verb.
- `upstream-git`: covers the upstream repository sense.
- `artifact-tech`: remains the build-output sense; `artifact-workflow` adds task-related documents and outputs.
- `deprecated-tech`, `idiomatic-tech`, and `vocabulary-general`: retained unchanged; not new candidates from this import.

All six original YAML files and recordings are preserved.

## Grouped inventory items

| Inventory item or pattern                       | Entry                                                             |
| :---------------------------------------------- | :---------------------------------------------------------------- |
| restore / restoration                           | [restore-tech](../entries/restore-tech.yml)                       |
| reject / rejection                              | [reject-tech](../entries/reject-tech.yml)                         |
| distinguish / distinction; distinguish X from Y | [distinguish-tech](../entries/distinguish-tech.yml)               |
| infer / inferred                                | [infer-tech](../entries/infer-tech.yml)                           |
| emit / emission                                 | [emit-tech](../entries/emit-tech.yml)                             |
| propagate / propagation                         | [propagate-tech](../entries/propagate-tech.yml)                   |
| prune / pruning                                 | [prune-tech](../entries/prune-tech.yml)                           |
| resolve / resolution                            | [resolve-tech](../entries/resolve-tech.yml)                       |
| invalidate / invalidation                       | [invalidate-tech](../entries/invalidate-tech.yml)                 |
| precede / preceding / preceded by               | [precede-tech](../entries/precede-tech.yml)                       |
| explicit / explicitly                           | [explicit-tech](../entries/explicit-tech.yml)                     |
| investigate / investigation                     | [investigate-tech](../entries/investigate-tech.yml)               |
| inspect / inspection                            | [inspect-tech](../entries/inspect-tech.yml)                       |
| reproduce / reproduction                        | [reproduce-tech](../entries/reproduce-tech.yml)                   |
| mock / mocked                                   | [mock-tech](../entries/mock-tech.yml)                             |
| synthesize / synthesized                        | [synthesize-tech](../entries/synthesize-tech.yml)                 |
| iterable / non-iterable                         | [iterable-tech](../entries/iterable-tech.yml)                     |
| sanitizer / sanitization                        | [sanitize-tech](../entries/sanitize-tech.yml)                     |
| encode / encoding                               | [encode-tech](../entries/encode-tech.yml)                         |
| decode / decoder                                | [decode-tech](../entries/decode-tech.yml)                         |
| register / registration                         | [register-tech](../entries/register-tech.yml)                     |
| exit status / exit code                         | [exit-status-tech](../entries/exit-status-tech.yml)               |
| detach / detachment                             | [detach-git](../entries/detach-git.yml)                           |
| preserve X while Y; without changing X          | [preserve-tech](../entries/preserve-tech.yml)                     |
| consistent with X                               | [consistent-tech](../entries/consistent-tech.yml)                 |
| coverage gap                                    | [coverage-tech](../entries/coverage-tech.yml)                     |
| including runtime evaluation                    | [runtime-evaluation-tech](../entries/runtime-evaluation-tech.yml) |
| outside this fix / out of scope                 | [scope-workflow](../entries/scope-workflow.yml)                   |
| leave validation to X                           | [validation-tech](../entries/validation-tech.yml)                 |
| stale result                                    | [stale-tech](../entries/stale-tech.yml)                           |
| attribute X to Y                                | [attribute-to-general](../entries/attribute-to-general.yml)       |
| supports the reported cause                     | [support-general](../entries/support-general.yml)                 |
| was not independently reproduced                | [independently-tech](../entries/independently-tech.yml)           |
| does not justify X                              | [justify-general](../entries/justify-general.yml)                 |
| unless the scope materially changes             | [material-tech](../entries/material-tech.yml)                     |

## Exclusions and corrected labels

- Excluded `implicit`, `underlying`, `substantiate`, `reconcile`, and `state-safe`: earlier assistant suggestions or explanatory labels without source support for this inventory.
- Excluded `at least`: it appeared in the earlier inventory, but a complete literal search of the 27 documents found no occurrence. `at most` is supported and included.
- Normalized "constrain through restrict" to `restrict`, "substantiate through support" to `support`, "reconcile through resolve" to `resolve`, and "state-safe idea expressed as read-only" to `read-only`.
- Excluded incorrect forms from correction-table Before columns, code identifiers, URLs, and command flags as standalone vocabulary. Corrected terms such as `toolchain` and `callback` are included.

## New entries and sources

Source paths below are relative to `docs/PR` in the source repository. Each YAML entry records its source in `notes`. A headword may be normalized from an inflected or derived form; source wording is noted where useful.

### Changes and behavior

| Word or expression | Entry                                               | Source                                                 |
| :----------------- | :-------------------------------------------------- | :----------------------------------------------------- |
| preserve           | [preserve-tech](../entries/preserve-tech.yml)       | react/react-37721.md                                   |
| retain             | [retain-tech](../entries/retain-tech.yml)           | react/react-9-rust-compiler.md                         |
| restore            | [restore-tech](../entries/restore-tech.yml)         | react/react-d-compiler.md                              |
| prevent            | [prevent-tech](../entries/prevent-tech.yml)         | react/react-37836.md                                   |
| reject             | [reject-tech](../entries/reject-tech.yml)           | react/react-9-rust-compiler.md                         |
| distinguish        | [distinguish-tech](../entries/distinguish-tech.yml) | typescript/reports/typescript-wedoct28.md              |
| infer              | [infer-tech](../entries/infer-tech.yml)             | typescript/typescript-a-62854.md                       |
| derive             | [derive-tech](../entries/derive-tech.yml)           | typescript/typescript-a-62854.md                       |
| invoke             | [invoke-tech](../entries/invoke-tech.yml)           | typescript/typescript-a-62854.md                       |
| emit               | [emit-tech](../entries/emit-tech.yml)               | typescript/typescript-64952.md                         |
| defer              | [defer-tech](../entries/defer-tech.yml)             | react/react-37721.md                                   |
| bypass             | [bypass-tech](../entries/bypass-tech.yml)           | react/react-c-rust-compiler.md                         |
| invalidate         | [invalidate-tech](../entries/invalidate-tech.yml)   | typescript/typescript-64936.md                         |
| overwrite          | [overwrite-tech](../entries/overwrite-tech.yml)     | typescript/schedulers/daily-1am-issue-investigation.md |
| synchronize        | [synchronize-tech](../entries/synchronize-tech.yml) | react/react-37859.md                                   |
| propagate          | [propagate-tech](../entries/propagate-tech.yml)     | react/react-37995.md                                   |
| prune              | [prune-tech](../entries/prune-tech.yml)             | typescript/typescript-64936.md                         |
| coerce             | [coerce-tech](../entries/coerce-tech.yml)           | react/reports/react-candidates.md                      |
| resolve            | [resolve-tech](../entries/resolve-tech.yml)         | typescript/reports/typescript-wedoct28.md              |
| resolve            | [resolve-workflow](../entries/resolve-workflow.yml) | typescript/schedulers/daily-1am-issue-investigation.md |
| consume            | [consume-tech](../entries/consume-tech.yml)         | react/reports/react-candidates.md                      |
| alternate          | [alternate-tech](../entries/alternate-tech.yml)     | react/react-37995.md                                   |
| toggle             | [toggle-tech](../entries/toggle-tech.yml)           | react/react-37995.md                                   |
| precede            | [precede-tech](../entries/precede-tech.yml)         | react/react-c-rust-compiler.md                         |
| exclude            | [exclude-tech](../entries/exclude-tech.yml)         | go101/go101-321.md                                     |
| omit               | [omit-tech](../entries/omit-tech.yml)               | AGENTS.md                                              |
| introduce          | [introduce-tech](../entries/introduce-tech.yml)     | react/reports/react-candidates.md                      |
| expose             | [expose-tech](../entries/expose-tech.yml)           | react/reports/react-candidates.md                      |
| overlap            | [overlap-tech](../entries/overlap-tech.yml)         | typescript/typescript-a-62854.md                       |
| restrict           | [restrict-tech](../entries/restrict-tech.yml)       | typescript/typescript-64965.md                         |
| corrupt            | [corrupt-tech](../entries/corrupt-tech.yml)         | react/reports/react-b-plan.md                          |

### Conditions and limitations

| Word or expression | Entry                                                   | Source                                                 |
| :----------------- | :------------------------------------------------------ | :----------------------------------------------------- |
| subsequent         | [subsequent-tech](../entries/subsequent-tech.yml)       | react/react-37836.md                                   |
| eligible           | [eligible-tech](../entries/eligible-tech.yml)           | react/react-37836.md                                   |
| stale              | [stale-tech](../entries/stale-tech.yml)                 | react/reports/react-candidates.md                      |
| legacy             | [legacy-tech](../entries/legacy-tech.yml)               | react/react-d-compiler.md                              |
| consistent         | [consistent-tech](../entries/consistent-tech.yml)       | react/react-37953.md                                   |
| intact             | [intact-tech](../entries/intact-tech.yml)               | react/react-c-rust-compiler.md                         |
| explicit           | [explicit-tech](../entries/explicit-tech.yml)           | typescript/schedulers/daily-1am-issue-investigation.md |
| genuine            | [genuine-tech](../entries/genuine-tech.yml)             | react/react-c-rust-compiler.md                         |
| literal            | [literal-text](../entries/literal-text.yml)             | react/react-c-rust-compiler.md                         |
| adjacent           | [adjacent-tech](../entries/adjacent-tech.yml)           | react/react-c-rust-compiler.md                         |
| nested             | [nested-tech](../entries/nested-tech.yml)               | react/react-a-compiler.md                              |
| malformed          | [malformed-tech](../entries/malformed-tech.yml)         | react/reports/react-b-plan.md                          |
| truncated          | [truncated-tech](../entries/truncated-tech.yml)         | react/reports/react-b-plan.md                          |
| inaccessible       | [inaccessible-tech](../entries/inaccessible-tech.yml)   | typescript/typescript-64952.md                         |
| uninitialized      | [uninitialized-tech](../entries/uninitialized-tech.yml) | typescript/typescript-64965.md                         |
| mutable            | [mutable-tech](../entries/mutable-tech.yml)             | typescript/typescript-64965.md                         |
| anonymous          | [anonymous-tech](../entries/anonymous-tech.yml)         | typescript/typescript-a-62854.md                       |
| bounded            | [bounded-tech](../entries/bounded-tech.yml)             | typescript/reports/typescript-wedoct28.md              |
| exhaustive         | [exhaustive-tech](../entries/exhaustive-tech.yml)       | react/reports/react-candidates.md                      |
| actionable         | [actionable-tech](../entries/actionable-tech.yml)       | react/reports/react-candidates.md                      |
| concrete           | [concrete-tech](../entries/concrete-tech.yml)           | react/reports/react-candidates.md                      |
| suspected          | [suspected-tech](../entries/suspected-tech.yml)         | typescript/schedulers/daily-1am-issue-investigation.md |
| confirmed          | [confirmed-tech](../entries/confirmed-tech.yml)         | react/reports/react-candidates.md                      |
| speculative        | [speculative-tech](../entries/speculative-tech.yml)     | typescript/schedulers/daily-1am-issue-investigation.md |
| outstanding        | [outstanding-tech](../entries/outstanding-tech.yml)     | react/reports/react-candidates.md                      |
| pending            | [pending-tech](../entries/pending-tech.yml)             | typescript/reports/typescript-wedoct28.md              |
| applicable         | [applicable-tech](../entries/applicable-tech.yml)       | typescript/schedulers/daily-1am-issue-investigation.md |
| material           | [material-tech](../entries/material-tech.yml)           | AGENTS.md                                              |
| pre-existing       | [pre-existing-tech](../entries/pre-existing-tech.yml)   | typescript/schedulers/daily-1am-issue-investigation.md |
| reproducible       | [reproducible-tech](../entries/reproducible-tech.yml)   | typescript/schedulers/daily-1am-issue-investigation.md |
| maintainable       | [maintainable-tech](../entries/maintainable-tech.yml)   | typescript/schedulers/daily-1am-issue-investigation.md |
| independently      | [independently-tech](../entries/independently-tech.yml) | typescript/reports/typescript-tueoct27.md              |

### Investigation and testing

| Word or expression | Entry                                                             | Source                                                 |
| :----------------- | :---------------------------------------------------------------- | :----------------------------------------------------- |
| investigate        | [investigate-tech](../entries/investigate-tech.yml)               | typescript/schedulers/daily-1am-issue-investigation.md |
| inspect            | [inspect-tech](../entries/inspect-tech.yml)                       | typescript/schedulers/daily-1am-issue-investigation.md |
| reproduce          | [reproduce-tech](../entries/reproduce-tech.yml)                   | typescript/schedulers/daily-1am-issue-investigation.md |
| regression         | [regression-tech](../entries/regression-tech.yml)                 | react/react-37871.md                                   |
| regression test    | [regression-test-tech](../entries/regression-test-tech.yml)       | react/react-37898.md                                   |
| validation         | [validation-tech](../entries/validation-tech.yml)                 | typescript/typescript-a-62854.md                       |
| verification       | [verification-tech](../entries/verification-tech.yml)             | react/react-9-rust-compiler.md                         |
| evidence           | [evidence-tech](../entries/evidence-tech.yml)                     | typescript/schedulers/daily-1am-issue-investigation.md |
| finding            | [finding-tech](../entries/finding-tech.yml)                       | typescript/schedulers/daily-1am-issue-investigation.md |
| candidate          | [candidate-tech](../entries/candidate-tech.yml)                   | typescript/reports/typescript-wedoct28.md              |
| trigger            | [trigger-tech](../entries/trigger-tech.yml)                       | react/reports/react-candidates.md                      |
| probe              | [probe-tech](../entries/probe-tech.yml)                           | react/react-c-rust-compiler.md                         |
| audit              | [audit-tech](../entries/audit-tech.yml)                           | typescript/reports/typescript-tueoct27.md              |
| coverage           | [coverage-tech](../entries/coverage-tech.yml)                     | react/react-d-compiler.md                              |
| fixture            | [fixture-tech](../entries/fixture-tech.yml)                       | react/react-a-compiler.md                              |
| assertion          | [assertion-testing](../entries/assertion-testing.yml)             | react/reports/react-b-plan.md                          |
| baseline           | [baseline-tech](../entries/baseline-tech.yml)                     | typescript/typescript-64965.md                         |
| snapshot           | [snapshot-tech](../entries/snapshot-tech.yml)                     | react/react-37871.md                                   |
| test suite         | [test-suite-tech](../entries/test-suite-tech.yml)                 | react/react-37757.md                                   |
| harness            | [harness-tech](../entries/harness-tech.yml)                       | react/react-a-compiler.md                              |
| runtime evaluation | [runtime-evaluation-tech](../entries/runtime-evaluation-tech.yml) | react/react-37995.md                                   |
| integration test   | [integration-test-tech](../entries/integration-test-tech.yml)     | react/react-d-compiler.md                              |
| mock               | [mock-tech](../entries/mock-tech.yml)                             | react/react-d-compiler.md                              |
| static inspection  | [static-inspection-tech](../entries/static-inspection-tech.yml)   | react/reports/react-candidates.md                      |
| round trip         | [round-trip-tech](../entries/round-trip-tech.yml)                 | react/reports/react-b-plan.md                          |
| benchmark          | [benchmark-tech](../entries/benchmark-tech.yml)                   | AGENTS.md                                              |
| workload           | [workload-tech](../entries/workload-tech.yml)                     | AGENTS.md                                              |
| synthetic          | [synthetic-tech](../entries/synthetic-tech.yml)                   | AGENTS.md                                              |
| edge case          | [edge-case-tech](../entries/edge-case-tech.yml)                   | typescript/reports/typescript-tueoct27.md              |
| false diagnostic   | [false-diagnostic-tech](../entries/false-diagnostic-tech.yml)     | typescript/typescript-64965.md                         |
| correctness        | [correctness-tech](../entries/correctness-tech.yml)               | typescript/reports/typescript-wedoct28.md              |
| limitation         | [limitation-tech](../entries/limitation-tech.yml)                 | AGENTS.md                                              |

### Compiler and language terms

| Word or expression    | Entry                                                                   | Source                                    |
| :-------------------- | :---------------------------------------------------------------------- | :---------------------------------------- |
| diagnostic            | [diagnostic-tech](../entries/diagnostic-tech.yml)                       | typescript/typescript-64952.md            |
| diagnostic context    | [diagnostic-context-tech](../entries/diagnostic-context-tech.yml)       | typescript/typescript-64952.md            |
| declaration emit      | [declaration-emit-tech](../entries/declaration-emit-tech.yml)           | typescript/typescript-64952.md            |
| heritage type         | [heritage-type-tech](../entries/heritage-type-tech.yml)                 | typescript/typescript-64952.md            |
| clause                | [clause-tech](../entries/clause-tech.yml)                               | typescript/typescript-64952.md            |
| type argument         | [type-argument-tech](../entries/type-argument-tech.yml)                 | typescript/typescript-64952.md            |
| type assertion        | [type-assertion-tech](../entries/type-assertion-tech.yml)               | typescript/typescript-64965.md            |
| assignment target     | [assignment-target-tech](../entries/assignment-target-tech.yml)         | typescript/typescript-64965.md            |
| compound assignment   | [compound-assignment-tech](../entries/compound-assignment-tech.yml)     | typescript/typescript-64965.md            |
| binding               | [binding-tech](../entries/binding-tech.yml)                             | typescript/typescript-64965.md            |
| cast                  | [cast-tech](../entries/cast-tech.yml)                                   | typescript/typescript-64965.md            |
| flow analysis         | [flow-analysis-tech](../entries/flow-analysis-tech.yml)                 | typescript/typescript-64965.md            |
| contextual type       | [contextual-type-tech](../entries/contextual-type-tech.yml)             | typescript/reports/typescript-wedoct28.md |
| union                 | [union-tech](../entries/union-tech.yml)                                 | typescript/reports/typescript-wedoct28.md |
| synthesize            | [synthesize-tech](../entries/synthesize-tech.yml)                       | typescript/reports/typescript-wedoct28.md |
| overload              | [overload-tech](../entries/overload-tech.yml)                           | typescript/reports/typescript-wedoct28.md |
| accessor              | [accessor-tech](../entries/accessor-tech.yml)                           | typescript/typescript-a-62854.md          |
| getter                | [getter-tech](../entries/getter-tech.yml)                               | typescript/typescript-a-62854.md          |
| decorator             | [decorator-tech](../entries/decorator-tech.yml)                         | typescript/typescript-a-62854.md          |
| computed key          | [computed-key-tech](../entries/computed-key-tech.yml)                   | typescript/typescript-a-62854.md          |
| static initialization | [static-initialization-tech](../entries/static-initialization-tech.yml) | typescript/typescript-a-62854.md          |
| constant propagation  | [constant-propagation-tech](../entries/constant-propagation-tech.yml)   | react/react-37995.md                      |
| signed zero           | [signed-zero-tech](../entries/signed-zero-tech.yml)                     | react/react-37995.md                      |
| equality semantics    | [equality-semantics-tech](../entries/equality-semantics-tech.yml)       | react/react-37995.md                      |
| array spread          | [array-spread-tech](../entries/array-spread-tech.yml)                   | react/react-a-compiler.md                 |
| iterable              | [iterable-tech](../entries/iterable-tech.yml)                           | react/react-a-compiler.md                 |
| catch handler         | [catch-handler-tech](../entries/catch-handler-tech.yml)                 | react/react-a-compiler.md                 |
| nonthrowing           | [nonthrowing-tech](../entries/nonthrowing-tech.yml)                     | react/react-a-compiler.md                 |
| directive             | [directive-tech](../entries/directive-tech.yml)                         | react/react-37871.md                      |
| annotation            | [annotation-tech](../entries/annotation-tech.yml)                       | react/react-37871.md                      |
| prefilter             | [prefilter-tech](../entries/prefilter-tech.yml)                         | react/react-37871.md                      |
| dynamic gating        | [dynamic-gating-tech](../entries/dynamic-gating-tech.yml)               | react/react-9-rust-compiler.md            |
| closure               | [closure-tech](../entries/closure-tech.yml)                             | react/react-37836.md                      |
| scope                 | [scope-lexical](../entries/scope-lexical.yml)                           | react/react-37995.md                      |
| transformation        | [transformation-tech](../entries/transformation-tech.yml)               | react/react-37836.md                      |
| code generation       | [code-generation-tech](../entries/code-generation-tech.yml)             | react/react-37995.md                      |
| source map            | [source-map-tech](../entries/source-map-tech.yml)                       | typescript/typescript-a-62854.md          |
| serialization         | [serialization-tech](../entries/serialization-tech.yml)                 | react/reports/react-b-plan.md             |
| traversal             | [traversal-tech](../entries/traversal-tech.yml)                         | typescript/typescript-64936.md            |
| breadth-first         | [breadth-first-tech](../entries/breadth-first-tech.yml)                 | typescript/typescript-64936.md            |
| fast path             | [fast-path-tech](../entries/fast-path-tech.yml)                         | typescript/typescript-64936.md            |
| linear                | [linear-tech](../entries/linear-tech.yml)                               | typescript/typescript-64936.md            |
| incremental build     | [incremental-build-tech](../entries/incremental-build-tech.yml)         | typescript/reports/typescript-tueoct27.md |
| allowlist             | [allowlist-tech](../entries/allowlist-tech.yml)                         | typescript/reports/typescript-tueoct27.md |
| literal               | [literal-programming](../entries/literal-programming.yml)               | typescript/reports/typescript-wedoct28.md |

### Unicode and text processing

| Word or expression  | Entry                                                               | Source                            |
| :------------------ | :------------------------------------------------------------------ | :-------------------------------- |
| identifier          | [identifier-tech](../entries/identifier-tech.yml)                   | react/react-9-rust-compiler.md    |
| code unit           | [code-unit-tech](../entries/code-unit-tech.yml)                     | react/react-37898.md              |
| byte offset         | [byte-offset-tech](../entries/byte-offset-tech.yml)                 | react/react-37898.md              |
| surrogate pair      | [surrogate-pair-tech](../entries/surrogate-pair-tech.yml)           | react/react-c-rust-compiler.md    |
| lone surrogate      | [lone-surrogate-tech](../entries/lone-surrogate-tech.yml)           | react/react-c-rust-compiler.md    |
| combining mark      | [combining-mark-tech](../entries/combining-mark-tech.yml)           | react/react-9-rust-compiler.md    |
| joiner              | [joiner-tech](../entries/joiner-tech.yml)                           | react/react-9-rust-compiler.md    |
| accented text       | [accented-text-tech](../entries/accented-text-tech.yml)             | react/react-37898.md              |
| multibyte           | [multibyte-tech](../entries/multibyte-tech.yml)                     | react/reports/react-b-plan.md     |
| escape              | [escape-tech](../entries/escape-tech.yml)                           | react/react-c-rust-compiler.md    |
| escaped backslash   | [escaped-backslash-tech](../entries/escaped-backslash-tech.yml)     | react/react-c-rust-compiler.md    |
| escape parity       | [escape-parity-tech](../entries/escape-parity-tech.yml)             | react/reports/react-candidates.md |
| sanitize            | [sanitize-tech](../entries/sanitize-tech.yml)                       | react/react-c-rust-compiler.md    |
| marker              | [marker-tech](../entries/marker-tech.yml)                           | react/reports/react-b-plan.md     |
| reserved prefix     | [reserved-prefix-tech](../entries/reserved-prefix-tech.yml)         | react/reports/react-candidates.md |
| collision           | [collision-tech](../entries/collision-tech.yml)                     | react/reports/react-candidates.md |
| encode              | [encode-tech](../entries/encode-tech.yml)                           | react/react-c-rust-compiler.md    |
| decode              | [decode-tech](../entries/decode-tech.yml)                           | react/reports/react-b-plan.md     |
| negative lookbehind | [negative-lookbehind-tech](../entries/negative-lookbehind-tech.yml) | react/reports/react-candidates.md |
| concatenation       | [concatenation-tech](../entries/concatenation-tech.yml)             | react/reports/react-b-plan.md     |
| misalign            | [misalign-tech](../entries/misalign-tech.yml)                       | react/react-37898.md              |
| underline           | [underline-tech](../entries/underline-tech.yml)                     | react/react-37898.md              |
| multiline range     | [multiline-range-tech](../entries/multiline-range-tech.yml)         | react/react-37898.md              |
| transport metadata  | [transport-metadata-tech](../entries/transport-metadata-tech.yml)   | react/reports/react-candidates.md |

### Runtime and builds

| Word or expression | Entry                                                           | Source                            |
| :----------------- | :-------------------------------------------------------------- | :-------------------------------- |
| runtime            | [runtime-tech](../entries/runtime-tech.yml)                     | react/react-d-compiler.md         |
| backend            | [backend-tech](../entries/backend-tech.yml)                     | react/react-a-compiler.md         |
| adapter            | [adapter-tech](../entries/adapter-tech.yml)                     | react/react-9-rust-compiler.md    |
| bridge             | [bridge-tech](../entries/bridge-tech.yml)                       | react/reports/react-b-plan.md     |
| native addon       | [native-addon-tech](../entries/native-addon-tech.yml)           | react/reports/react-b-plan.md     |
| internals          | [internals-tech](../entries/internals-tech.yml)                 | react/react-d-compiler.md         |
| dispatcher         | [dispatcher-tech](../entries/dispatcher-tech.yml)               | react/react-d-compiler.md         |
| guard              | [guard-tech](../entries/guard-tech.yml)                         | react/react-37836.md              |
| re-entry           | [re-entry-tech](../entries/re-entry-tech.yml)                   | react/react-37836.md              |
| mount              | [mount-tech](../entries/mount-tech.yml)                         | react/react-37721.md              |
| unmount            | [unmount-tech](../entries/unmount-tech.yml)                     | react/react-37721.md              |
| lifecycle          | [lifecycle-tech](../entries/lifecycle-tech.yml)                 | react/reports/react-candidates.md |
| effect replay      | [effect-replay-tech](../entries/effect-replay-tech.yml)         | react/react-37721.md              |
| setup              | [setup-tech](../entries/setup-tech.yml)                         | react/react-37721.md              |
| cleanup            | [cleanup-tech](../entries/cleanup-tech.yml)                     | react/reports/react-candidates.md |
| cancellation       | [cancellation-tech](../entries/cancellation-tech.yml)           | typescript/typescript-1.md        |
| graceful shutdown  | [graceful-shutdown-tech](../entries/graceful-shutdown-tech.yml) | typescript/typescript-1.md        |
| transport failure  | [transport-failure-tech](../entries/transport-failure-tech.yml) | typescript/typescript-1.md        |
| exit status        | [exit-status-tech](../entries/exit-status-tech.yml)             | typescript/typescript-1.md        |
| recovery           | [recovery-tech](../entries/recovery-tech.yml)                   | react/react-d-compiler.md         |
| forwarding         | [forwarding-tech](../entries/forwarding-tech.yml)               | react/react-d-compiler.md         |
| register           | [register-tech](../entries/register-tech.yml)                   | react/reports/react-candidates.md |
| unregister         | [unregister-tech](../entries/unregister-tech.yml)               | react/reports/react-candidates.md |
| registry           | [registry-tech](../entries/registry-tech.yml)                   | react/reports/react-candidates.md |
| profiling          | [profiling-tech](../entries/profiling-tech.yml)                 | react/react-9-rust-compiler.md    |
| asynchronous       | [asynchronous-tech](../entries/asynchronous-tech.yml)           | react/reports/react-candidates.md |
| out-of-order       | [out-of-order-tech](../entries/out-of-order-tech.yml)           | react/reports/react-candidates.md |
| field presence     | [field-presence-tech](../entries/field-presence-tech.yml)       | react/react-d-compiler.md         |
| hashing            | [hashing-tech](../entries/hashing-tech.yml)                     | react/react-37859.md              |
| lockfile           | [lockfile-tech](../entries/lockfile-tech.yml)                   | react/react-37859.md              |
| manifest           | [manifest-tech](../entries/manifest-tech.yml)                   | react/react-37859.md              |
| version pin        | [version-pin-tech](../entries/version-pin-tech.yml)             | react/react-9-rust-compiler.md    |
| symlink            | [symlink-tech](../entries/symlink-tech.yml)                     | react/react-37995.md              |
| toolchain          | [toolchain-tech](../entries/toolchain-tech.yml)                 | go101/go101-321.md                |
| callback           | [callback-tech](../entries/callback-tech.yml)                   | go101/go101-321.md                |

### Contribution workflow

| Word or expression     | Entry                                                                     | Source                                                 |
| :--------------------- | :------------------------------------------------------------------------ | :----------------------------------------------------- |
| scope                  | [scope-workflow](../entries/scope-workflow.yml)                           | typescript/schedulers/daily-1am-issue-investigation.md |
| handoff                | [handoff-workflow](../entries/handoff-workflow.yml)                       | typescript/schedulers/daily-1am-issue-investigation.md |
| invocation             | [invocation-workflow](../entries/invocation-workflow.yml)                 | typescript/schedulers/daily-1am-issue-investigation.md |
| run budget             | [run-budget-workflow](../entries/run-budget-workflow.yml)                 | typescript/schedulers/daily-1am-issue-investigation.md |
| reserve                | [reserve-workflow](../entries/reserve-workflow.yml)                       | typescript/schedulers/daily-1am-issue-investigation.md |
| elapsed time           | [elapsed-time-workflow](../entries/elapsed-time-workflow.yml)             | typescript/schedulers/daily-1am-issue-investigation.md |
| resumption             | [resumption-workflow](../entries/resumption-workflow.yml)                 | typescript/reports/typescript-tueoct27.md              |
| blocker                | [blocker-workflow](../entries/blocker-workflow.yml)                       | typescript/schedulers/daily-1am-issue-investigation.md |
| ownership token        | [ownership-token-workflow](../entries/ownership-token-workflow.yml)       | typescript/schedulers/daily-1am-issue-investigation.md |
| atomically             | [atomically-workflow](../entries/atomically-workflow.yml)                 | typescript/schedulers/daily-1am-issue-investigation.md |
| acquire                | [acquire-workflow](../entries/acquire-workflow.yml)                       | typescript/schedulers/daily-1am-issue-investigation.md |
| release                | [release-workflow](../entries/release-workflow.yml)                       | typescript/schedulers/daily-1am-issue-investigation.md |
| occupied               | [occupied-workflow](../entries/occupied-workflow.yml)                     | typescript/schedulers/daily-1am-issue-investigation.md |
| abandoned              | [abandoned-workflow](../entries/abandoned-workflow.yml)                   | typescript/schedulers/daily-1am-issue-investigation.md |
| isolated worktree      | [isolated-worktree-workflow](../entries/isolated-worktree-workflow.yml)   | typescript/schedulers/daily-1am-issue-investigation.md |
| primary checkout       | [primary-checkout-workflow](../entries/primary-checkout-workflow.yml)     | typescript/schedulers/daily-1am-issue-investigation.md |
| detach                 | [detach-git](../entries/detach-git.yml)                                   | typescript/schedulers/daily-1am-issue-investigation.md |
| remote-tracking branch | [remote-tracking-branch-git](../entries/remote-tracking-branch-git.yml)   | typescript/schedulers/daily-1am-issue-investigation.md |
| staged                 | [staged-git](../entries/staged-git.yml)                                   | typescript/schedulers/daily-1am-issue-investigation.md |
| unstaged               | [unstaged-git](../entries/unstaged-git.yml)                               | typescript/schedulers/daily-1am-issue-investigation.md |
| untracked              | [untracked-git](../entries/untracked-git.yml)                             | AGENTS.md                                              |
| amend                  | [amend-git](../entries/amend-git.yml)                                     | typescript/schedulers/daily-1am-issue-investigation.md |
| discard                | [discard-workflow](../entries/discard-workflow.yml)                       | typescript/schedulers/daily-1am-issue-investigation.md |
| overlap prevention     | [overlap-prevention-workflow](../entries/overlap-prevention-workflow.yml) | typescript/schedulers/daily-1am-issue-investigation.md |
| competing work         | [competing-work-workflow](../entries/competing-work-workflow.yml)         | typescript/reports/typescript-tueoct27.md              |
| unclaimed              | [unclaimed-workflow](../entries/unclaimed-workflow.yml)                   | typescript/reports/typescript-wedoct28.md              |
| submission             | [submission-workflow](../entries/submission-workflow.yml)                 | typescript/schedulers/daily-1am-issue-investigation.md |
| disclosure             | [disclosure-workflow](../entries/disclosure-workflow.yml)                 | AGENTS.md                                              |
| artifact               | [artifact-workflow](../entries/artifact-workflow.yml)                     | typescript/schedulers/daily-1am-issue-investigation.md |
| pinned checkout        | [pinned-checkout-workflow](../entries/pinned-checkout-workflow.yml)       | typescript/reports/typescript-wedoct28.md              |
| read-only              | [read-only-workflow](../entries/read-only-workflow.yml)                   | react/reports/react-candidates.md                      |

### Writing and review

| Word or expression | Entry                                                         | Source                                                 |
| :----------------- | :------------------------------------------------------------ | :----------------------------------------------------- |
| wording            | [wording-general](../entries/wording-general.yml)             | go101/go101-321.md                                     |
| clarify            | [clarify-general](../entries/clarify-general.yml)             | go101/go101-321.md                                     |
| corresponding      | [corresponding-general](../entries/corresponding-general.yml) | go101/go101-321.md                                     |
| typo               | [typo-general](../entries/typo-general.yml)                   | go101/go101-321.md                                     |
| concise            | [concise-general](../entries/concise-general.yml)             | AGENTS.md                                              |
| tradeoff           | [tradeoff-general](../entries/tradeoff-general.yml)           | AGENTS.md                                              |
| caveat             | [caveat-general](../entries/caveat-general.yml)               | typescript/schedulers/daily-1am-issue-investigation.md |
| imperative         | [imperative-general](../entries/imperative-general.yml)       | typescript/schedulers/daily-1am-issue-investigation.md |
| factual            | [factual-general](../entries/factual-general.yml)             | typescript/schedulers/daily-1am-issue-investigation.md |
| claim              | [claim-general](../entries/claim-general.yml)                 | AGENTS.md                                              |
| support            | [support-general](../entries/support-general.yml)             | typescript/reports/typescript-wedoct28.md              |
| generalize         | [generalize-general](../entries/generalize-general.yml)       | AGENTS.md                                              |
| imply              | [imply-general](../entries/imply-general.yml)                 | AGENTS.md                                              |
| attribute to       | [attribute-to-general](../entries/attribute-to-general.yml)   | typescript/reports/typescript-tueoct27.md              |
| justify            | [justify-general](../entries/justify-general.yml)             | typescript/reports/typescript-wedoct28.md              |
| merit              | [merit-general](../entries/merit-general.yml)                 | typescript/reports/typescript-tueoct27.md              |
| invaluable         | [invaluable-general](../entries/invaluable-general.yml)       | go101/go101-267.md                                     |
| gratitude          | [gratitude-general](../entries/gratitude-general.yml)         | go101/go101-267.md                                     |
| appreciate         | [appreciate-general](../entries/appreciate-general.yml)       | go101/go101-267.md                                     |
| contribution       | [contribution-general](../entries/contribution-general.yml)   | typescript/reports/typescript-tueoct27.md              |
| revision           | [revision-general](../entries/revision-general.yml)           | typescript/typescript-64952.md                         |
| revision           | [revision-git](../entries/revision-git.yml)                   | react/react-37859.md                                   |

### Reusable expressions

| Word or expression                       | Entry                                                                                                                 | Source                                                 |
| :--------------------------------------- | :-------------------------------------------------------------------------------------------------------------------- | :----------------------------------------------------- |
| remain unchanged                         | [remain-unchanged-workflow](../entries/remain-unchanged-workflow.yml)                                                 | react/react-37898.md                                   |
| remain intact                            | [remain-intact-workflow](../entries/remain-intact-workflow.yml)                                                       | react/react-c-rust-compiler.md                         |
| fail against the original implementation | [fail-against-the-original-implementation-workflow](../entries/fail-against-the-original-implementation-workflow.yml) | react/reports/react-candidates.md                      |
| pass with the fix                        | [pass-with-the-fix-workflow](../entries/pass-with-the-fix-workflow.yml)                                               | react/react-c-rust-compiler.md                         |
| before and after the fix                 | [before-and-after-the-fix-workflow](../entries/before-and-after-the-fix-workflow.yml)                                 | react/react-37898.md                                   |
| generated output only                    | [generated-output-only-workflow](../entries/generated-output-only-workflow.yml)                                       | react/react-d-compiler.md                              |
| leave untouched                          | [leave-untouched-workflow](../entries/leave-untouched-workflow.yml)                                                   | typescript/reports/typescript-tueoct27.md              |
| run in parallel                          | [run-in-parallel-workflow](../entries/run-in-parallel-workflow.yml)                                                   | typescript/typescript-64936.md                         |
| in order                                 | [in-order-workflow](../entries/in-order-workflow.yml)                                                                 | react/react-c-rust-compiler.md                         |
| in one pass                              | [in-one-pass-workflow](../entries/in-one-pass-workflow.yml)                                                           | react/reports/react-b-plan.md                          |
| followed by                              | [followed-by-workflow](../entries/followed-by-workflow.yml)                                                           | react/react-c-rust-compiler.md                         |
| rather than                              | [rather-than-workflow](../entries/rather-than-workflow.yml)                                                           | react/reports/react-candidates.md                      |
| as proof of                              | [as-proof-of-workflow](../entries/as-proof-of-workflow.yml)                                                           | typescript/schedulers/daily-1am-issue-investigation.md |
| based on                                 | [based-on-workflow](../entries/based-on-workflow.yml)                                                                 | typescript/reports/typescript-tueoct27.md              |
| before drawing conclusions               | [before-drawing-conclusions-workflow](../entries/before-drawing-conclusions-workflow.yml)                             | typescript/reports/typescript-wedoct28.md              |
| at most                                  | [at-most-workflow](../entries/at-most-workflow.yml)                                                                   | typescript/schedulers/daily-1am-issue-investigation.md |
| when applicable                          | [when-applicable-workflow](../entries/when-applicable-workflow.yml)                                                   | typescript/schedulers/daily-1am-issue-investigation.md |
| where practical                          | [where-practical-workflow](../entries/where-practical-workflow.yml)                                                   | react/reports/react-candidates.md                      |
| on every exit path                       | [on-every-exit-path-workflow](../entries/on-every-exit-path-workflow.yml)                                             | typescript/schedulers/daily-1am-issue-investigation.md |
| await review                             | [await-review-workflow](../entries/await-review-workflow.yml)                                                         | typescript/schedulers/daily-1am-issue-investigation.md |
| ready for review                         | [ready-for-review-workflow](../entries/ready-for-review-workflow.yml)                                                 | typescript/schedulers/daily-1am-issue-investigation.md |
| reported behavior                        | [reported-behavior-workflow](../entries/reported-behavior-workflow.yml)                                               | typescript/reports/typescript-wedoct28.md              |
| observed result                          | [observed-result-workflow](../entries/observed-result-workflow.yml)                                                   | typescript/schedulers/daily-1am-issue-investigation.md |
| expected behavior                        | [expected-behavior-workflow](../entries/expected-behavior-workflow.yml)                                               | react/reports/react-candidates.md                      |
| actual behavior                          | [actual-behavior-workflow](../entries/actual-behavior-workflow.yml)                                                   | react/reports/react-candidates.md                      |

## Audio and verification

The audio command speaks each entry's `word` using Samantha at 150 words per minute and LAME CLI at 96 kbps mono. Existing recordings are reused. Decoder checks verify playable MP3 data; they do not certify the pronunciation of every specialized term. Filenames follow stable IDs; audio bytes may vary across operating-system or encoder versions.

Validation performed for this import:

- `go test ./...` and `go vet ./...` passed, including audio selection, generation failures, source-edit detection, YAML preservation, overwrite protection, and batch-resume cases.
- Validation and export passed for all 283 entries. The five repeated-word warnings represent the separate meanings documented above.
- All 283 MP3s decoded successfully. A second audio run generated nothing and skipped all 283 recordings; every MP3 hash remained unchanged.
- Two export builds produced identical bytes. The six original entries and six original recordings remained byte-identical to their tracked versions.
- Source references and audio links were checked for all 277 new entries. New text files passed LF, final-newline, and trailing-whitespace checks; the Markdown entry links were checked locally.
- Local playback commands completed for `preserve`, `constant propagation`, and `remain unchanged`. This confirms sample playback, not an independent assessment of pronunciation accuracy. Specialized pronunciations still need listening review.

No Anki import or synchronization was performed. The PR source documents were not changed, and no files were staged or committed.
