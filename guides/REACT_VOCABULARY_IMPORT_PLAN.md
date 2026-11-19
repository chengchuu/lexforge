# React vocabulary import plan

**Status** — Approved and implemented on 2026-11-19.

The core-first batch adds 44 English entries and reuses nine, for 53 cards under `tag:react` at import. Seven existing cards received the topic tag. The existing Effect card now follows *commit phase*, and *reconciliation* also follows *render*.

The [learning-order guide](LEARNING_ORDER.md#apply-the-order-in-anki) explains how to apply the export order to existing Anki new cards.

## Sources

- S1 — [Quick Start](https://react.dev/learn)
- S2 — [Render and Commit](https://react.dev/learn/render-and-commit)
- S3 — [Keeping Components Pure](https://react.dev/learn/keeping-components-pure)
- S4 — [Preserving and Resetting State](https://react.dev/learn/preserving-and-resetting-state)
- S5 — [Managing State](https://react.dev/learn/managing-state)
- S6 — [State as a Snapshot](https://react.dev/learn/state-as-a-snapshot)
- S7 — [Queueing a Series of State Updates](https://react.dev/learn/queueing-a-series-of-state-updates)
- S8 — [Escape Hatches](https://react.dev/learn/escape-hatches)
- S9 — [Lifecycle of Reactive Effects](https://react.dev/learn/lifecycle-of-reactive-effects)
- S10 — [Rules of React](https://react.dev/reference/rules)
- S11 — [React Reference Overview](https://react.dev/reference/react)
- S12 — [StrictMode](https://react.dev/reference/react/StrictMode)
- S13 — [Component — Catching rendering errors with an error boundary](https://react.dev/reference/react/Component#catching-rendering-errors-with-an-error-boundary)
- S14 — [createPortal](https://react.dev/reference/react-dom/createPortal)
- S15 — [Suspense](https://react.dev/reference/react/Suspense)
- S16 — [useTransition](https://react.dev/reference/react/useTransition)
- S17 — [Server React DOM APIs](https://react.dev/reference/react-dom/server)
- S18 — [hydrateRoot](https://react.dev/reference/react-dom/client/hydrateRoot)

## New entries

The linked cards contain Chinese meanings, English examples, and provenance. A dash indicates an empty prerequisite list.

### Components and JSX

| Word                    | Stable ID                                                                 | Priority | Prerequisites                      | Source |
| :---------------------- | :------------------------------------------------------------------------ | :------- | :--------------------------------- | :----- |
| `React`                 | [react-library](../entries/react-library.yml)                             | 1        | —                                  | S1     |
| `component`             | [component-react](../entries/component-react.yml)                         | 1        | `react-library`                    | S1     |
| `function component`    | [function-component-react](../entries/function-component-react.yml)       | 1        | `component-react`, `function-tech` | S1     |
| `JSX`                   | [jsx-react](../entries/jsx-react.yml)                                     | 1        | `component-react`                  | S1     |
| `props`                 | [props-react](../entries/props-react.yml)                                 | 1        | `component-react`                  | S1     |
| `children`              | [children-react](../entries/children-react.yml)                           | 1        | `props-react`                      | S1     |
| `Fragment`              | [fragment-react](../entries/fragment-react.yml)                           | 1        | `jsx-react`                        | S11    |
| `event handler`         | [event-handler-react](../entries/event-handler-react.yml)                 | 1        | `function-tech`                    | S1     |
| `conditional rendering` | [conditional-rendering-react](../entries/conditional-rendering-react.yml) | 1        | `jsx-react`, `render-react`        | S1     |
| `list rendering`        | [list-rendering-react](../entries/list-rendering-react.yml)               | 1        | `jsx-react`, `render-react`        | S1     |
| `key`                   | [key-react](../entries/key-react.yml)                                     | 1        | `list-rendering-react`             | S1     |

### Rendering and component identity

| Word                 | Stable ID                                                           | Priority | Prerequisites                             | Source |
| :------------------- | :------------------------------------------------------------------ | :------- | :---------------------------------------- | :----- |
| `render`             | [render-react](../entries/render-react.yml)                         | 1        | `component-react`                         | S2     |
| `re-render`          | [rerender-react](../entries/rerender-react.yml)                     | 1        | `render-react`                            | S2     |
| `commit phase`       | [commit-phase-react](../entries/commit-phase-react.yml)             | 1        | `render-react`                            | S2     |
| `pure component`     | [pure-component-react](../entries/pure-component-react.yml)         | 2        | `render-react`, `props-react`             | S3     |
| `component identity` | [component-identity-react](../entries/component-identity-react.yml) | 2        | `component-react`, `key-react`            | S4     |
| `state preservation` | [state-preservation-react](../entries/state-preservation-react.yml) | 2        | `component-identity-react`, `state-react` | S4     |

### State and data flow

| Word                     | Stable ID                                                                   | Priority | Prerequisites                                | Source |
| :----------------------- | :-------------------------------------------------------------------------- | :------- | :------------------------------------------- | :----- |
| `state`                  | [state-react](../entries/state-react.yml)                                   | 1        | `component-react`                            | S5     |
| `state setter`           | [state-setter-react](../entries/state-setter-react.yml)                     | 2        | `state-react`, `hook-react`                  | S7     |
| `updater function`       | [updater-function-react](../entries/updater-function-react.yml)             | 2        | `state-snapshot-react`, `state-setter-react` | S7     |
| `state snapshot`         | [state-snapshot-react](../entries/state-snapshot-react.yml)                 | 2        | `state-react`, `render-react`                | S6     |
| `batching`               | [batching-react](../entries/batching-react.yml)                             | 2        | `updater-function-react`                     | S7     |
| `lifting state up`       | [lifting-state-up-react](../entries/lifting-state-up-react.yml)             | 2        | `state-react`, `props-react`                 | S5     |
| `controlled component`   | [controlled-component-react](../entries/controlled-component-react.yml)     | 2        | `props-react`, `state-react`                 | S5     |
| `uncontrolled component` | [uncontrolled-component-react](../entries/uncontrolled-component-react.yml) | 2        | `controlled-component-react`                 | S5     |
| `derived state`          | [derived-state-react](../entries/derived-state-react.yml)                   | 2        | `state-react`, `props-react`                 | S5     |
| `context`                | [context-react](../entries/context-react.yml)                               | 2        | `props-react`                                | S5     |
| `context provider`       | [context-provider-react](../entries/context-provider-react.yml)             | 2        | `context-react`                              | S5     |
| `reducer`                | [reducer-react](../entries/reducer-react.yml)                               | 2        | `state-react`                                | S5     |

### Hooks and escape hatches

| Word                  | Stable ID                                                             | Priority | Prerequisites                  | Source |
| :-------------------- | :-------------------------------------------------------------------- | :------- | :----------------------------- | :----- |
| `Hook`                | [hook-react](../entries/hook-react.yml)                               | 1        | `function-component-react`     | S10    |
| `custom Hook`         | [custom-hook-react](../entries/custom-hook-react.yml)                 | 2        | `hook-react`                   | S11    |
| `Rules of Hooks`      | [rules-of-hooks-react](../entries/rules-of-hooks-react.yml)           | 2        | `hook-react`                   | S10    |
| `ref`                 | [ref-react](../entries/ref-react.yml)                                 | 2        | `hook-react`                   | S8     |
| `reactive dependency` | [reactive-dependency-react](../entries/reactive-dependency-react.yml) | 2        | `effect-react`                 | S9     |
| `dependency array`    | [dependency-array-react](../entries/dependency-array-react.yml)       | 2        | `reactive-dependency-react`    | S9     |
| `Effect cleanup`      | [effect-cleanup-react](../entries/effect-cleanup-react.yml)           | 2        | `effect-react`, `cleanup-tech` | S9     |

### Common advanced concepts

| Word               | Stable ID                                                       | Priority | Prerequisites                     | Source |
| :----------------- | :-------------------------------------------------------------- | :------- | :-------------------------------- | :----- |
| `Strict Mode`      | [strict-mode-react](../entries/strict-mode-react.yml)           | 2        | `render-react`, `effect-react`    | S12    |
| `error boundary`   | [error-boundary-react](../entries/error-boundary-react.yml)     | 2        | `component-react`, `render-react` | S13    |
| `portal`           | [portal-react](../entries/portal-react.yml)                     | 2        | `component-react`                 | S14    |
| `Suspense`         | [suspense-react](../entries/suspense-react.yml)                 | 3        | `fallback-ui-react`               | S15    |
| `fallback UI`      | [fallback-ui-react](../entries/fallback-ui-react.yml)           | 2        | `conditional-rendering-react`     | S15    |
| `transition`       | [transition-react](../entries/transition-react.yml)             | 3        | `state-setter-react`              | S16    |
| `server rendering` | [server-rendering-react](../entries/server-rendering-react.yml) | 3        | `render-react`                    | S17    |
| `hydration`        | [hydration-react](../entries/hydration-react.yml)               | 3        | `server-rendering-react`          | S18    |

## Reused entries

| Word                | Stable ID                                                               | Priority | Prerequisites after import             | Added prerequisites  |
| :------------------ | :---------------------------------------------------------------------- | :------- | :------------------------------------- | :------------------- |
| `effect`            | [effect-react](../entries/effect-react.yml)                             | 2        | `commit-phase-react`                   | `commit-phase-react` |
| `reconciliation`    | [reconciliation-react](../entries/reconciliation-react.yml)             | 2        | `virtual-dom-frontend`, `render-react` | `render-react`       |
| `mount`             | [mount-tech](../entries/mount-tech.yml)                                 | 2        | —                                      | —                    |
| `unmount`           | [unmount-tech](../entries/unmount-tech.yml)                             | 2        | `mount-tech`                           | —                    |
| `one-way data flow` | [one-way-data-flow-frontend](../entries/one-way-data-flow-frontend.yml) | 1        | —                                      | —                    |
| `cleanup`           | [cleanup-tech](../entries/cleanup-tech.yml)                             | 1        | —                                      | —                    |
| `callback`          | [callback-tech](../entries/callback-tech.yml)                           | 2        | —                                      | —                    |
| `memoization`       | [memoization-programming](../entries/memoization-programming.yml)       | 2        | —                                      | —                    |
| `virtual DOM`       | [virtual-dom-frontend](../entries/virtual-dom-frontend.yml)             | 2        | —                                      | —                    |

## Grouping decisions

- Teach React’s render and commit phases separately. The TypeScript, CSS, Git, and general vocabulary retain their own senses. React *context* and *transition* have separate cards.
- Include `useState`, `useEffect`, `useRef`, `useContext`, and `useReducer` in Notes for the concepts they teach. Keep API names as supporting material. When studying the reused *memoization* card, consider `memo` for skipping component renders and `useMemo` for caching calculations.
- Defer a reference-wide pass over Server Components, Client Components, class lifecycles, and React Compiler terms. This batch adds no associations.

## Implementation and verification

Go tests, vet, vocabulary validation, and diff checks passed. Two builds produced identical bytes. The export contains 1,541 rows with 17 columns. All 759 declared prerequisite links follow the generated `StudyOrder`.

Verified US IPA is present in 12 new entries. The other 32 await a reliable pronunciation source. Existing MP3 hashes matched the baseline. This batch generated no audio.

The 22 existing same-word warnings remain alongside two new warnings for the distinct React senses of *context* and *transition*.

## Prerequisite follow-up

The shared *function* card now carries the React tag so it can appear before *function component* and *event handler*. The current `tag:react` filter includes 54 cards.
