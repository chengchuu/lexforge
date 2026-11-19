# HTML vocabulary import plan

**Status** — Approved and implemented on 2026-11-19.

The core-first English batch adds 54 entries and reuses five existing frontend entries. The original import put 59 cards under `tag:html`. The batch covers document syntax, semantics, forms, accessible content, embedded content, tables, and metadata.

Use the [learning-order guide](LEARNING_ORDER.md#apply-the-order-in-anki) to apply the generated ranks to existing Anki new cards.

## Sources

- S1 — [WHATWG HTML — Syntax](https://html.spec.whatwg.org/multipage/syntax.html)
- S2 — [WHATWG DOM Standard](https://dom.spec.whatwg.org/)
- S3 — [WHATWG HTML — Global attributes](https://html.spec.whatwg.org/multipage/dom.html#global-attributes)
- S4 — [WHATWG HTML — Boolean attributes](https://html.spec.whatwg.org/multipage/common-microsyntaxes.html#boolean-attributes)
- S5 — [WHATWG HTML — Content models](https://html.spec.whatwg.org/multipage/dom.html#content-models)
- S6 — [MDN — Structuring content with HTML](https://developer.mozilla.org/en-US/docs/Learn_web_development/Core/Structuring_content)
- S7 — [W3C WAI — Page Structure Tutorial](https://www.w3.org/WAI/tutorials/page-structure/)
- S8 — [WHATWG HTML — Links](https://html.spec.whatwg.org/multipage/links.html)
- S9 — [WHATWG HTML — Forms](https://html.spec.whatwg.org/multipage/forms.html)
- S10 — [WHATWG HTML — Form control infrastructure](https://html.spec.whatwg.org/multipage/form-control-infrastructure.html)
- S11 — [WHATWG HTML — The input element](https://html.spec.whatwg.org/multipage/input.html)
- S12 — [W3C WAI — Labeling Controls](https://www.w3.org/WAI/tutorials/forms/labels/)
- S13 — [W3C — Accessible Name and Description Computation](https://www.w3.org/TR/accname-1.2/)
- S14 — [WHATWG HTML — Embedded content](https://html.spec.whatwg.org/multipage/embedded-content.html)
- S15 — [W3C WAI — Images Tutorial](https://www.w3.org/WAI/tutorials/images/)
- S16 — [WHATWG HTML — Tabular data](https://html.spec.whatwg.org/multipage/tables.html)
- S17 — [W3C WAI — Tables Tutorial](https://www.w3.org/WAI/tutorials/tables/)
- S18 — [WHATWG HTML — The elements of HTML](https://html.spec.whatwg.org/multipage/semantics.html)
- S19 — [WHATWG HTML — Scripting](https://html.spec.whatwg.org/multipage/scripting.html)
- S20 — [WHATWG HTML — The iframe element](https://html.spec.whatwg.org/multipage/iframe-embed-object.html#the-iframe-element)

## New entries

Each linked card contains a Chinese meaning, an English example, and a source reference. A dash indicates an empty prerequisite list.

### Markup and document structure

| Word                  | Stable ID                                                           | Priority | Prerequisites                   | Source |
| :-------------------- | :------------------------------------------------------------------ | :------- | :------------------------------ | :----- |
| `HTML`                | [html-language](../entries/html-language.yml)                       | 1        | —                               | S6     |
| `markup`              | [markup-html](../entries/markup-html.yml)                           | 1        | `html-language`                 | S1     |
| `document`            | [document-html](../entries/document-html.yml)                       | 1        | `html-language`                 | S2     |
| `doctype`             | [doctype-html](../entries/doctype-html.yml)                         | 1        | `document-html`                 | S1     |
| `element`             | [element-html](../entries/element-html.yml)                         | 1        | `document-html`                 | S2     |
| `tag`                 | [tag-html](../entries/tag-html.yml)                                 | 1        | `markup-html`, `element-html`   | S1     |
| `start tag`           | [start-tag-html](../entries/start-tag-html.yml)                     | 1        | `tag-html`                      | S1     |
| `end tag`             | [end-tag-html](../entries/end-tag-html.yml)                         | 1        | `tag-html`                      | S1     |
| `attribute`           | [attribute-html](../entries/attribute-html.yml)                     | 1        | `element-html`                  | S1     |
| `global attribute`    | [global-attribute-html](../entries/global-attribute-html.yml)       | 2        | `attribute-html`                | S3     |
| `boolean attribute`   | [boolean-attribute-html](../entries/boolean-attribute-html.yml)     | 2        | `attribute-html`                | S4     |
| `void element`        | [void-element-html](../entries/void-element-html.yml)               | 2        | `element-html`, `end-tag-html`  | S1     |
| `nesting`             | [nesting-html](../entries/nesting-html.yml)                         | 1        | `element-html`                  | S1     |
| `character reference` | [character-reference-html](../entries/character-reference-html.yml) | 2        | `markup-html`                   | S1     |
| `DOM`                 | [dom-html](../entries/dom-html.yml)                                 | 1        | `document-html`, `element-html` | S2     |

### Semantics and navigation

| Word                  | Stable ID                                                           | Priority | Prerequisites    | Source |
| :-------------------- | :------------------------------------------------------------------ | :------- | :--------------- | :----- |
| `semantic HTML`       | [semantic-html](../entries/semantic-html.yml)                       | 2        | `element-html`   | S6     |
| `heading`             | [heading-html](../entries/heading-html.yml)                         | 1        | `element-html`   | S7     |
| `heading level`       | [heading-level-html](../entries/heading-level-html.yml)             | 2        | `heading-html`   | S7     |
| `paragraph`           | [paragraph-html](../entries/paragraph-html.yml)                     | 1        | `element-html`   | S6     |
| `ordered list`        | [ordered-list-html](../entries/ordered-list-html.yml)               | 1        | `element-html`   | S6     |
| `unordered list`      | [unordered-list-html](../entries/unordered-list-html.yml)           | 1        | `element-html`   | S6     |
| `list item`           | [list-item-html](../entries/list-item-html.yml)                     | 1        | `element-html`   | S6     |
| `hyperlink`           | [hyperlink-html](../entries/hyperlink-html.yml)                     | 1        | `attribute-html` | S8     |
| `fragment identifier` | [fragment-identifier-html](../entries/fragment-identifier-html.yml) | 2        | `hyperlink-html` | S8     |
| `landmark`            | [landmark-html](../entries/landmark-html.yml)                       | 2        | `semantic-html`  | S7     |

### Forms and accessibility

| Word                    | Stable ID                                                               | Priority | Prerequisites                                      | Source |
| :---------------------- | :---------------------------------------------------------------------- | :------- | :------------------------------------------------- | :----- |
| `form`                  | [form-html](../entries/form-html.yml)                                   | 1        | `element-html`                                     | S9     |
| `label`                 | [label-html](../entries/label-html.yml)                                 | 2        | `form-control-bootstrap`                           | S12    |
| `input type`            | [input-type-html](../entries/input-type-html.yml)                       | 2        | `attribute-html`, `form-control-bootstrap`         | S11    |
| `required field`        | [required-field-html](../entries/required-field-html.yml)               | 2        | `form-control-bootstrap`, `boolean-attribute-html` | S10    |
| `disabled control`      | [disabled-control-html](../entries/disabled-control-html.yml)           | 2        | `form-control-bootstrap`, `boolean-attribute-html` | S10    |
| `read-only control`     | [read-only-control-html](../entries/read-only-control-html.yml)         | 2        | `form-control-bootstrap`, `disabled-control-html`  | S11    |
| `autocomplete`          | [autocomplete-html](../entries/autocomplete-html.yml)                   | 2        | `form-control-bootstrap`, `attribute-html`         | S10    |
| `constraint validation` | [constraint-validation-html](../entries/constraint-validation-html.yml) | 2        | `required-field-html`, `input-type-html`           | S10    |
| `form submission`       | [form-submission-html](../entries/form-submission-html.yml)             | 2        | `form-html`                                        | S10    |
| `accessible name`       | [accessible-name-html](../entries/accessible-name-html.yml)             | 2        | `label-html`                                       | S13    |

### Images and embedded content

| Word               | Stable ID                                                     | Priority | Prerequisites                   | Source |
| :----------------- | :------------------------------------------------------------ | :------- | :------------------------------ | :----- |
| `alternative text` | [alternative-text-html](../entries/alternative-text-html.yml) | 2        | `attribute-html`                | S15    |
| `decorative image` | [decorative-image-html](../entries/decorative-image-html.yml) | 2        | `alternative-text-html`         | S15    |
| `responsive image` | [responsive-image-html](../entries/responsive-image-html.yml) | 2        | `attribute-html`                | S14    |
| `source set`       | [source-set-html](../entries/source-set-html.yml)             | 2        | `responsive-image-html`         | S14    |
| `lazy loading`     | [lazy-loading-html](../entries/lazy-loading-html.yml)         | 2        | `attribute-html`                | S14    |
| `inline frame`     | [inline-frame-html](../entries/inline-frame-html.yml)         | 2        | `document-html`, `element-html` | S20    |

### Data tables

| Word          | Stable ID                                           | Priority | Prerequisites     | Source |
| :------------ | :-------------------------------------------------- | :------- | :---------------- | :----- |
| `data table`  | [data-table-html](../entries/data-table-html.yml)   | 2        | `element-html`    | S17    |
| `caption`     | [caption-html](../entries/caption-html.yml)         | 2        | `data-table-html` | S16    |
| `table row`   | [table-row-html](../entries/table-row-html.yml)     | 2        | `data-table-html` | S16    |
| `header cell` | [header-cell-html](../entries/header-cell-html.yml) | 2        | `table-row-html`  | S17    |
| `data cell`   | [data-cell-html](../entries/data-cell-html.yml)     | 2        | `table-row-html`  | S16    |
| `column span` | [column-span-html](../entries/column-span-html.yml) | 2        | `data-cell-html`  | S16    |
| `row span`    | [row-span-html](../entries/row-span-html.yml)       | 2        | `data-cell-html`  | S16    |

### Metadata and document rules

| Word                  | Stable ID                                                           | Priority | Prerequisites                  | Source |
| :-------------------- | :------------------------------------------------------------------ | :------- | :----------------------------- | :----- |
| `metadata`            | [metadata-html](../entries/metadata-html.yml)                       | 2        | `document-html`                | S18    |
| `document title`      | [document-title-html](../entries/document-title-html.yml)           | 2        | `metadata-html`                | S18    |
| `character encoding`  | [character-encoding-html](../entries/character-encoding-html.yml)   | 2        | `metadata-html`                | S18    |
| `script`              | [script-html](../entries/script-html.yml)                           | 2        | `document-html`                | S19    |
| `content model`       | [content-model-html](../entries/content-model-html.yml)             | 3        | `element-html`, `nesting-html` | S5     |
| `interactive content` | [interactive-content-html](../entries/interactive-content-html.yml) | 3        | `content-model-html`           | S5     |

## Reused entries

Each card received the `html` tag.

| Word              | Stable ID                                                           | Priority | Prerequisites         |
| :---------------- | :------------------------------------------------------------------ | :------- | :-------------------- |
| `data attribute`  | [data-attribute-bootstrap](../entries/data-attribute-bootstrap.yml) | 2        | `bootstrap-framework` |
| `form control`    | [form-control-bootstrap](../entries/form-control-bootstrap.yml)     | 2        | `bootstrap-framework` |
| `focus indicator` | [focus-indicator-css](../entries/focus-indicator-css.yml)           | 2        | `pseudo-class-css`    |
| `stylesheet`      | [stylesheet-css](../entries/stylesheet-css.yml)                     | 1        | `css-language`        |
| `viewport`        | [viewport-css](../entries/viewport-css.yml)                         | 1        | —                     |

## Grouping decisions

- Teach the document tree’s *element* separately from the source syntax’s *tag*. Use `<p>`, `</p>`, and other markup in examples or notes.
- Teach *boolean attribute* after *attribute*. Use `disabled="false"` to show how presence works. Distinguish *disabled control* and *read-only control*.
- Reuse *form control* and teach *label* before *accessible name*. Keep HTML *table row* distinct from a Bootstrap layout row.
- Use `alt`, `srcset`, `sizes`, `scope`, `colspan`, and `rowspan` as supporting examples. Defer a broader pass on dialogs, popovers, templates, custom elements, and script-loading behavior. This batch adds no associations.

## Implementation and verification

Go tests, vet, vocabulary validation, and diff checks passed. Two builds produced identical bytes. The export contains 1,642 rows with 17 columns. All 877 declared prerequisite links follow the generated `StudyOrder`.

Verified US IPA is present in 16 new entries. The other 38 await a reliable pronunciation source. Existing MP3 hashes matched the baseline. This batch generated no audio.

Validation reported the same 24 same-word warnings as before the import.

## Prerequisite follow-up

The shared *form control* and *data attribute* cards now follow *form* and *attribute*, respectively. The tag also includes CSS, *selector*, and *pseudo-class* to support *stylesheet* and *focus indicator*. The current `tag:html` filter includes 62 cards.
