# CSS vocabulary import plan

**Status** — Approved and implemented on 2026-11-19.

The core-first batch adds 72 English entries and reuses five. Three already carried the `css` tag. The browser *layout* and *paint* cards received the tag to bring all 77 targets together. Find the topic in Anki using `tag:css`.

Priorities follow the [learning-order guide](LEARNING_ORDER.md). Prerequisites order introductions. Follow [Apply the order in Anki](LEARNING_ORDER.md#apply-the-order-in-anki) to import the complete export and reposition existing new cards.

## Sources

- S1 — [MDN — CSS styling basics](https://developer.mozilla.org/en-US/docs/Learn_web_development/Core/Styling_basics)
- S2 — [MDN — CSS reference](https://developer.mozilla.org/en-US/docs/Web/CSS/Reference)
- S3 — [W3C — Selectors Level 4](https://www.w3.org/TR/selectors-4/)
- S4 — [W3C — CSS Cascading and Inheritance Level 5](https://www.w3.org/TR/css-cascade-5/)
- S5 — [W3C — CSS Box Model Module Level 4](https://www.w3.org/TR/css-box-4/)
- S6 — [MDN — CSS layout](https://developer.mozilla.org/en-US/docs/Learn_web_development/Core/CSS_layout)
- S7 — [W3C — CSS Flexible Box Layout Module Level 1](https://www.w3.org/TR/css-flexbox-1/)
- S8 — [W3C — CSS Grid Layout Module Level 2](https://www.w3.org/TR/css-grid-2/)
- S9 — [MDN — CSS container queries](https://developer.mozilla.org/en-US/docs/Web/CSS/Guides/Containment/Container_queries)
- S10 — [W3C — CSS Custom Properties for Cascading Variables](https://www.w3.org/TR/css-variables-1/)
- S11 — [MDN — prefers-reduced-motion](https://developer.mozilla.org/en-US/docs/Web/CSS/Reference/At-rules/@media/prefers-reduced-motion)

## New entries

The linked cards contain concise Chinese meanings, English examples, and source references. Each row records the priority, direct prerequisites, and primary source. A dash indicates an empty prerequisite list.

### Stylesheets and syntax

| Word                 | Stable ID                                                       | Priority | Prerequisites                        | Source |
| :------------------- | :-------------------------------------------------------------- | :------- | :----------------------------------- | :----- |
| `CSS`                | [css-language](../entries/css-language.yml)                     | 1        | —                                    | S1     |
| `stylesheet`         | [stylesheet-css](../entries/stylesheet-css.yml)                 | 1        | `css-language`                       | S1     |
| `style rule`         | [style-rule-css](../entries/style-rule-css.yml)                 | 1        | `selector-css`, `declaration-css`    | S1     |
| `selector`           | [selector-css](../entries/selector-css.yml)                     | 1        | `css-language`                       | S3     |
| `declaration`        | [declaration-css](../entries/declaration-css.yml)               | 1        | `property-css`, `property-value-css` | S1     |
| `CSS property`       | [property-css](../entries/property-css.yml)                     | 1        | `css-language`                       | S2     |
| `property value`     | [property-value-css](../entries/property-value-css.yml)         | 1        | `property-css`                       | S1     |
| `at-rule`            | [at-rule-css](../entries/at-rule-css.yml)                       | 2        | `stylesheet-css`                     | S2     |
| `shorthand property` | [shorthand-property-css](../entries/shorthand-property-css.yml) | 2        | `longhand-property-css`              | S4     |
| `longhand property`  | [longhand-property-css](../entries/longhand-property-css.yml)   | 2        | `property-css`                       | S4     |

### Selectors and the cascade

| Word                 | Stable ID                                                       | Priority | Prerequisites                 | Source |
| :------------------- | :-------------------------------------------------------------- | :------- | :---------------------------- | :----- |
| `type selector`      | [type-selector-css](../entries/type-selector-css.yml)           | 1        | `selector-css`                | S3     |
| `class selector`     | [class-selector-css](../entries/class-selector-css.yml)         | 1        | `selector-css`                | S3     |
| `ID selector`        | [id-selector-css](../entries/id-selector-css.yml)               | 1        | `selector-css`                | S3     |
| `attribute selector` | [attribute-selector-css](../entries/attribute-selector-css.yml) | 2        | `selector-css`                | S3     |
| `combinator`         | [combinator-css](../entries/combinator-css.yml)                 | 2        | `selector-css`                | S3     |
| `pseudo-class`       | [pseudo-class-css](../entries/pseudo-class-css.yml)             | 2        | `selector-css`                | S3     |
| `pseudo-element`     | [pseudo-element-css](../entries/pseudo-element-css.yml)         | 2        | `selector-css`                | S3     |
| `cascade`            | [cascade-css](../entries/cascade-css.yml)                       | 1        | `declaration-css`             | S4     |
| `inheritance`        | [inheritance-css](../entries/inheritance-css.yml)               | 2        | `property-css`                | S4     |
| `specificity`        | [specificity-css](../entries/specificity-css.yml)               | 2        | `selector-css`, `cascade-css` | S4     |
| `source order`       | [source-order-css](../entries/source-order-css.yml)             | 2        | `cascade-css`                 | S4     |
| `cascade layer`      | [cascade-layer-css](../entries/cascade-layer-css.yml)           | 3        | `cascade-css`                 | S4     |
| `!important`         | [important-css](../entries/important-css.yml)                   | 2        | `cascade-css`                 | S4     |

### Boxes and flow

| Word                | Stable ID                                                     | Priority | Prerequisites                                  | Source |
| :------------------ | :------------------------------------------------------------ | :------- | :--------------------------------------------- | :----- |
| `box model`         | [box-model-css](../entries/box-model-css.yml)                 | 1        | `layout-browser-rendering`                     | S5     |
| `content box`       | [content-box-css](../entries/content-box-css.yml)             | 1        | `box-model-css`                                | S5     |
| `padding`           | [padding-css](../entries/padding-css.yml)                     | 1        | `box-model-css`                                | S5     |
| `border`            | [border-css](../entries/border-css.yml)                       | 1        | `box-model-css`                                | S5     |
| `margin`            | [margin-css](../entries/margin-css.yml)                       | 1        | `box-model-css`                                | S5     |
| `box-sizing`        | [box-sizing-css](../entries/box-sizing-css.yml)               | 2        | `content-box-css`, `padding-css`, `border-css` | S2     |
| `normal flow`       | [normal-flow-css](../entries/normal-flow-css.yml)             | 1        | `layout-browser-rendering`                     | S6     |
| `block-level box`   | [block-level-box-css](../entries/block-level-box-css.yml)     | 2        | `normal-flow-css`                              | S6     |
| `inline box`        | [inline-box-css](../entries/inline-box-css.yml)               | 2        | `normal-flow-css`                              | S6     |
| `display`           | [display-css](../entries/display-css.yml)                     | 1        | `layout-browser-rendering`                     | S2     |
| `overflow`          | [overflow-css](../entries/overflow-css.yml)                   | 2        | `box-model-css`                                | S2     |
| `margin collapsing` | [margin-collapsing-css](../entries/margin-collapsing-css.yml) | 2        | `margin-css`, `normal-flow-css`                | S5     |

### Layout and positioning

| Word                   | Stable ID                                                           | Priority | Prerequisites                             | Source |
| :--------------------- | :------------------------------------------------------------------ | :------- | :---------------------------------------- | :----- |
| `flex container`       | [flex-container-css](../entries/flex-container-css.yml)             | 2        | `flexbox-css`, `display-css`              | S7     |
| `flex item`            | [flex-item-css](../entries/flex-item-css.yml)                       | 2        | `flex-container-css`                      | S7     |
| `main axis`            | [main-axis-css](../entries/main-axis-css.yml)                       | 2        | `flex-container-css`                      | S7     |
| `cross axis`           | [cross-axis-css](../entries/cross-axis-css.yml)                     | 2        | `main-axis-css`                           | S7     |
| `grid container`       | [grid-container-css](../entries/grid-container-css.yml)             | 2        | `grid-layout-css`, `display-css`          | S8     |
| `grid track`           | [grid-track-css](../entries/grid-track-css.yml)                     | 2        | `grid-container-css`                      | S8     |
| `grid area`            | [grid-area-css](../entries/grid-area-css.yml)                       | 2        | `grid-container-css`                      | S8     |
| `gap`                  | [gap-css](../entries/gap-css.yml)                                   | 2        | `layout-browser-rendering`                | S8     |
| `alignment`            | [alignment-css](../entries/alignment-css.yml)                       | 2        | `layout-browser-rendering`                | S6     |
| `containing block`     | [containing-block-css](../entries/containing-block-css.yml)         | 2        | `layout-browser-rendering`                | S6     |
| `relative positioning` | [relative-positioning-css](../entries/relative-positioning-css.yml) | 2        | `normal-flow-css`                         | S6     |
| `absolute positioning` | [absolute-positioning-css](../entries/absolute-positioning-css.yml) | 2        | `containing-block-css`, `normal-flow-css` | S6     |
| `fixed positioning`    | [fixed-positioning-css](../entries/fixed-positioning-css.yml)       | 2        | `absolute-positioning-css`                | S6     |
| `sticky positioning`   | [sticky-positioning-css](../entries/sticky-positioning-css.yml)     | 2        | `normal-flow-css`                         | S6     |
| `stacking context`     | [stacking-context-css](../entries/stacking-context-css.yml)         | 3        | `layout-browser-rendering`                | S6     |
| `z-index`              | [z-index-css](../entries/z-index-css.yml)                           | 3        | `stacking-context-css`                    | S2     |

### Responsive design and reusable values

| Word                | Stable ID                                                     | Priority | Prerequisites                 | Source |
| :------------------ | :------------------------------------------------------------ | :------- | :---------------------------- | :----- |
| `responsive design` | [responsive-design-css](../entries/responsive-design-css.yml) | 2        | `viewport-css`                | S6     |
| `viewport`          | [viewport-css](../entries/viewport-css.yml)                   | 1        | —                             | S6     |
| `media query`       | [media-query-css](../entries/media-query-css.yml)             | 2        | `at-rule-css`, `viewport-css` | S6     |
| `breakpoint`        | [breakpoint-css](../entries/breakpoint-css.yml)               | 2        | `responsive-design-css`       | S6     |
| `container query`   | [container-query-css](../entries/container-query-css.yml)     | 3        | `responsive-design-css`       | S9     |
| `relative unit`     | [relative-unit-css](../entries/relative-unit-css.yml)         | 2        | `property-value-css`          | S2     |
| `rem`               | [rem-css](../entries/rem-css.yml)                             | 2        | `relative-unit-css`           | S2     |
| `fr`                | [fr-css](../entries/fr-css.yml)                               | 3        | `grid-track-css`              | S8     |
| `custom property`   | [custom-property-css](../entries/custom-property-css.yml)     | 2        | `property-css`                | S10    |
| `fallback value`    | [fallback-value-css](../entries/fallback-value-css.yml)       | 2        | `var-function-css`            | S10    |
| `var()`             | [var-function-css](../entries/var-function-css.yml)           | 2        | `custom-property-css`         | S10    |

### Typography, motion, and accessibility

| Word              | Stable ID                                                 | Priority | Prerequisites        | Source |
| :---------------- | :-------------------------------------------------------- | :------- | :------------------- | :----- |
| `font stack`      | [font-stack-css](../entries/font-stack-css.yml)           | 2        | `property-value-css` | S2     |
| `line height`     | [line-height-css](../entries/line-height-css.yml)         | 2        | `inline-box-css`     | S2     |
| `text wrapping`   | [text-wrapping-css](../entries/text-wrapping-css.yml)     | 2        | `inline-box-css`     | S2     |
| `transform`       | [transform-css](../entries/transform-css.yml)             | 2        | `property-css`       | S2     |
| `transition`      | [transition-css](../entries/transition-css.yml)           | 2        | `property-css`       | S2     |
| `animation`       | [animation-css](../entries/animation-css.yml)             | 2        | `property-css`       | S2     |
| `keyframe`        | [keyframe-css](../entries/keyframe-css.yml)               | 2        | `animation-css`      | S2     |
| `easing`          | [easing-css](../entries/easing-css.yml)                   | 3        | `transition-css`     | S2     |
| `focus indicator` | [focus-indicator-css](../entries/focus-indicator-css.yml) | 2        | `pseudo-class-css`   | S2     |
| `reduced motion`  | [reduced-motion-css](../entries/reduced-motion-css.yml)   | 2        | `media-query-css`    | S11    |

## Reused entries

Reuse the existing cards and their learning metadata.

| Word                       | Stable ID                                                                   | Priority | Prerequisites              | Topic tag |
| :------------------------- | :-------------------------------------------------------------------------- | :------- | :------------------------- | :-------- |
| `block formatting context` | [block-formatting-context-css](../entries/block-formatting-context-css.yml) | 2        | —                          | Existing  |
| `flexbox`                  | [flexbox-css](../entries/flexbox-css.yml)                                   | 1        | `layout-browser-rendering` | Existing  |
| `grid layout`              | [grid-layout-css](../entries/grid-layout-css.yml)                           | 2        | —                          | Existing  |
| `layout`                   | [layout-browser-rendering](../entries/layout-browser-rendering.yml)         | 1        | —                          | Added     |
| `paint`                    | [paint-browser-rendering](../entries/paint-browser-rendering.yml)           | 2        | `layout-browser-rendering` | Added     |

## Grouping decisions

- Keep the CSS *margin* separate from the finance sense. Use *CSS property* for a stylesheet declaration’s property and retain the object-property card as a separate sense.
- Teach *cascade* before *specificity* and *grid layout* before *grid track*. Examples show concrete CSS property values and syntax.
- Defer a reference-wide pass over nesting, subgrid, containment, logical properties, and modern color functions. This batch adds no associations.

## Implementation and verification

Go tests, vet, vocabulary validation, and diff checks passed. Two builds produced identical bytes. The export contains 1,497 rows with 17 columns. Every one of the 699 declared prerequisite links appears in the correct `StudyOrder`.

Verified US IPA is present in 13 new entries. The other 59 await a reliable pronunciation source. Existing MP3 hashes matched the baseline. This batch generated no audio.

The 21 existing same-word warnings remain. The one new warning is the intentional pairing of CSS *margin* with financial *margin*.
