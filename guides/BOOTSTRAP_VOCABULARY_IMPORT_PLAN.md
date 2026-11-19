# Bootstrap vocabulary import plan

**Status** — Approved and implemented on 2026-11-19.

The core-first English batch adds 47 entries and reuses 11 CSS entries. The original import put 58 cards under `tag:bootstrap`. The batch uses [Bootstrap 5.3 documentation](https://getbootstrap.com/docs/5.3/getting-started/introduction/) for its vocabulary and examples. See the [Bootstrap release index](https://getbootstrap.com/docs/versions/) for other versions.

Use the [learning-order guide](LEARNING_ORDER.md#apply-the-order-in-anki) to apply the generated ranks to existing Anki new cards.

## New entries

Each linked entry contains a Chinese meaning, an English example, and the source. A dash indicates that the entry has no declared prerequisite.

### Grid and responsive layout

| Word                   | Stable ID                                                                       | Priority | Prerequisites                           | Bootstrap 5.3 source                                                            |
| :--------------------- | :------------------------------------------------------------------------------ | :------- | :-------------------------------------- | :------------------------------------------------------------------------------ |
| `Bootstrap`            | [bootstrap-framework](../entries/bootstrap-framework.yml)                       | 1        | —                                       | [Introduction](https://getbootstrap.com/docs/5.3/getting-started/introduction/) |
| `grid system`          | [grid-system-bootstrap](../entries/grid-system-bootstrap.yml)                   | 1        | `bootstrap-framework`, `flexbox-css`    | [Grid](https://getbootstrap.com/docs/5.3/layout/grid/)                          |
| `container`            | [container-bootstrap](../entries/container-bootstrap.yml)                       | 1        | `bootstrap-framework`                   | [Containers](https://getbootstrap.com/docs/5.3/layout/containers/)              |
| `fluid container`      | [fluid-container-bootstrap](../entries/fluid-container-bootstrap.yml)           | 2        | `container-bootstrap`                   | [Containers](https://getbootstrap.com/docs/5.3/layout/containers/)              |
| `responsive container` | [responsive-container-bootstrap](../entries/responsive-container-bootstrap.yml) | 2        | `container-bootstrap`, `breakpoint-css` | [Containers](https://getbootstrap.com/docs/5.3/layout/containers/)              |
| `row`                  | [row-bootstrap](../entries/row-bootstrap.yml)                                   | 1        | `grid-system-bootstrap`                 | [Grid](https://getbootstrap.com/docs/5.3/layout/grid/)                          |
| `column`               | [column-bootstrap](../entries/column-bootstrap.yml)                             | 1        | `row-bootstrap`                         | [Grid](https://getbootstrap.com/docs/5.3/layout/grid/)                          |
| `gutter`               | [gutter-bootstrap](../entries/gutter-bootstrap.yml)                             | 2        | `column-bootstrap`, `padding-css`       | [Gutters](https://getbootstrap.com/docs/5.3/layout/gutters/)                    |
| `column offset`        | [column-offset-bootstrap](../entries/column-offset-bootstrap.yml)               | 2        | `column-bootstrap`                      | [Columns](https://getbootstrap.com/docs/5.3/layout/columns/)                    |

### Utilities

| Word                 | Stable ID                                                                   | Priority | Prerequisites                                          | Bootstrap 5.3 source                                            |
| :------------------- | :-------------------------------------------------------------------------- | :------- | :----------------------------------------------------- | :-------------------------------------------------------------- |
| `utility class`      | [utility-class-bootstrap](../entries/utility-class-bootstrap.yml)           | 1        | `bootstrap-framework`                                  | [Utility API](https://getbootstrap.com/docs/5.3/utilities/api/) |
| `responsive utility` | [responsive-utility-bootstrap](../entries/responsive-utility-bootstrap.yml) | 2        | `utility-class-bootstrap`, `breakpoint-css`            | [Utility API](https://getbootstrap.com/docs/5.3/utilities/api/) |
| `spacing utility`    | [spacing-utility-bootstrap](../entries/spacing-utility-bootstrap.yml)       | 2        | `utility-class-bootstrap`, `padding-css`, `margin-css` | [Spacing](https://getbootstrap.com/docs/5.3/utilities/spacing/) |
| `display utility`    | [display-utility-bootstrap](../entries/display-utility-bootstrap.yml)       | 2        | `utility-class-bootstrap`, `display-css`               | [Display](https://getbootstrap.com/docs/5.3/utilities/display/) |
| `utility API`        | [utility-api-bootstrap](../entries/utility-api-bootstrap.yml)               | 3        | `utility-class-bootstrap`                              | [Utility API](https://getbootstrap.com/docs/5.3/utilities/api/) |

### Common components

| Word           | Stable ID                                                       | Priority | Prerequisites         | Bootstrap 5.3 source                                                       |
| :------------- | :-------------------------------------------------------------- | :------- | :-------------------- | :------------------------------------------------------------------------- |
| `navbar`       | [navbar-bootstrap](../entries/navbar-bootstrap.yml)             | 2        | `bootstrap-framework` | [Navbar](https://getbootstrap.com/docs/5.3/components/navbar/)             |
| `dropdown`     | [dropdown-bootstrap](../entries/dropdown-bootstrap.yml)         | 2        | `bootstrap-framework` | [Dropdowns](https://getbootstrap.com/docs/5.3/components/dropdowns/)       |
| `collapse`     | [collapse-bootstrap](../entries/collapse-bootstrap.yml)         | 2        | `bootstrap-framework` | [Collapse](https://getbootstrap.com/docs/5.3/components/collapse/)         |
| `accordion`    | [accordion-bootstrap](../entries/accordion-bootstrap.yml)       | 2        | `collapse-bootstrap`  | [Accordion](https://getbootstrap.com/docs/5.3/components/accordion/)       |
| `modal`        | [modal-bootstrap](../entries/modal-bootstrap.yml)               | 2        | `bootstrap-framework` | [Modal](https://getbootstrap.com/docs/5.3/components/modal/)               |
| `offcanvas`    | [offcanvas-bootstrap](../entries/offcanvas-bootstrap.yml)       | 2        | `bootstrap-framework` | [Offcanvas](https://getbootstrap.com/docs/5.3/components/offcanvas/)       |
| `card`         | [card-bootstrap](../entries/card-bootstrap.yml)                 | 2        | `bootstrap-framework` | [Card](https://getbootstrap.com/docs/5.3/components/card/)                 |
| `badge`        | [badge-bootstrap](../entries/badge-bootstrap.yml)               | 2        | `bootstrap-framework` | [Badge](https://getbootstrap.com/docs/5.3/components/badge/)               |
| `alert`        | [alert-bootstrap](../entries/alert-bootstrap.yml)               | 2        | `bootstrap-framework` | [Alerts](https://getbootstrap.com/docs/5.3/components/alerts/)             |
| `toast`        | [toast-bootstrap](../entries/toast-bootstrap.yml)               | 2        | `bootstrap-framework` | [Toasts](https://getbootstrap.com/docs/5.3/components/toasts/)             |
| `tooltip`      | [tooltip-bootstrap](../entries/tooltip-bootstrap.yml)           | 2        | `bootstrap-framework` | [Tooltips](https://getbootstrap.com/docs/5.3/components/tooltips/)         |
| `popover`      | [popover-bootstrap](../entries/popover-bootstrap.yml)           | 2        | `tooltip-bootstrap`   | [Popovers](https://getbootstrap.com/docs/5.3/components/popovers/)         |
| `breadcrumb`   | [breadcrumb-bootstrap](../entries/breadcrumb-bootstrap.yml)     | 2        | `bootstrap-framework` | [Breadcrumb](https://getbootstrap.com/docs/5.3/components/breadcrumb/)     |
| `pagination`   | [pagination-bootstrap](../entries/pagination-bootstrap.yml)     | 2        | `bootstrap-framework` | [Pagination](https://getbootstrap.com/docs/5.3/components/pagination/)     |
| `list group`   | [list-group-bootstrap](../entries/list-group-bootstrap.yml)     | 2        | `bootstrap-framework` | [List group](https://getbootstrap.com/docs/5.3/components/list-group/)     |
| `button group` | [button-group-bootstrap](../entries/button-group-bootstrap.yml) | 2        | `bootstrap-framework` | [Button group](https://getbootstrap.com/docs/5.3/components/button-group/) |
| `carousel`     | [carousel-bootstrap](../entries/carousel-bootstrap.yml)         | 2        | `bootstrap-framework` | [Carousel](https://getbootstrap.com/docs/5.3/components/carousel/)         |
| `progress bar` | [progress-bar-bootstrap](../entries/progress-bar-bootstrap.yml) | 2        | `bootstrap-framework` | [Progress](https://getbootstrap.com/docs/5.3/components/progress/)         |
| `spinner`      | [spinner-bootstrap](../entries/spinner-bootstrap.yml)           | 2        | `bootstrap-framework` | [Spinners](https://getbootstrap.com/docs/5.3/components/spinners/)         |

### Forms and accessibility

| Word                  | Stable ID                                                                     | Priority | Prerequisites                   | Bootstrap 5.3 source                                                          |
| :-------------------- | :---------------------------------------------------------------------------- | :------- | :------------------------------ | :---------------------------------------------------------------------------- |
| `form control`        | [form-control-bootstrap](../entries/form-control-bootstrap.yml)               | 2        | `bootstrap-framework`           | [Form control](https://getbootstrap.com/docs/5.3/forms/form-control/)         |
| `input group`         | [input-group-bootstrap](../entries/input-group-bootstrap.yml)                 | 2        | `form-control-bootstrap`        | [Input group](https://getbootstrap.com/docs/5.3/forms/input-group/)           |
| `validation feedback` | [validation-feedback-bootstrap](../entries/validation-feedback-bootstrap.yml) | 2        | `form-control-bootstrap`        | [Validation](https://getbootstrap.com/docs/5.3/forms/validation/)             |
| `valid state`         | [valid-state-bootstrap](../entries/valid-state-bootstrap.yml)                 | 2        | `validation-feedback-bootstrap` | [Validation](https://getbootstrap.com/docs/5.3/forms/validation/)             |
| `invalid state`       | [invalid-state-bootstrap](../entries/invalid-state-bootstrap.yml)             | 2        | `validation-feedback-bootstrap` | [Validation](https://getbootstrap.com/docs/5.3/forms/validation/)             |
| `visually hidden`     | [visually-hidden-bootstrap](../entries/visually-hidden-bootstrap.yml)         | 2        | `bootstrap-framework`           | [Visually hidden](https://getbootstrap.com/docs/5.3/helpers/visually-hidden/) |

### Behavior and customization

| Word                | Stable ID                                                                 | Priority | Prerequisites                                             | Bootstrap 5.3 source                                                        |
| :------------------ | :------------------------------------------------------------------------ | :------- | :-------------------------------------------------------- | :-------------------------------------------------------------------------- |
| `data attribute`    | [data-attribute-bootstrap](../entries/data-attribute-bootstrap.yml)       | 2        | `bootstrap-framework`                                     | [Javascript](https://getbootstrap.com/docs/5.3/getting-started/javascript/) |
| `JavaScript plugin` | [javascript-plugin-bootstrap](../entries/javascript-plugin-bootstrap.yml) | 2        | `bootstrap-framework`                                     | [Javascript](https://getbootstrap.com/docs/5.3/getting-started/javascript/) |
| `toggle trigger`    | [toggle-trigger-bootstrap](../entries/toggle-trigger-bootstrap.yml)       | 2        | `data-attribute-bootstrap`, `javascript-plugin-bootstrap` | [Javascript](https://getbootstrap.com/docs/5.3/getting-started/javascript/) |
| `backdrop`          | [backdrop-bootstrap](../entries/backdrop-bootstrap.yml)                   | 2        | `modal-bootstrap`                                         | [Modal](https://getbootstrap.com/docs/5.3/components/modal/)                |
| `theming`           | [theming-bootstrap](../entries/theming-bootstrap.yml)                     | 2        | `bootstrap-framework`                                     | [Css variables](https://getbootstrap.com/docs/5.3/customize/css-variables/) |
| `color mode`        | [color-mode-bootstrap](../entries/color-mode-bootstrap.yml)               | 2        | `theming-bootstrap`                                       | [Color modes](https://getbootstrap.com/docs/5.3/customize/color-modes/)     |
| `Sass variable`     | [sass-variable-bootstrap](../entries/sass-variable-bootstrap.yml)         | 3        | `theming-bootstrap`                                       | [Sass](https://getbootstrap.com/docs/5.3/customize/sass/)                   |
| `Sass mixin`        | [sass-mixin-bootstrap](../entries/sass-mixin-bootstrap.yml)               | 3        | `sass-variable-bootstrap`                                 | [Sass](https://getbootstrap.com/docs/5.3/customize/sass/)                   |

## Reused entries

The shared cards received the `bootstrap` tag. Their priorities, prerequisites, and examples are the same as before.

| Word                | Stable ID                                                     | Priority | Prerequisites              |
| :------------------ | :------------------------------------------------------------ | :------- | :------------------------- |
| `breakpoint`        | [breakpoint-css](../entries/breakpoint-css.yml)               | 2        | `responsive-design-css`    |
| `responsive design` | [responsive-design-css](../entries/responsive-design-css.yml) | 2        | `viewport-css`             |
| `flexbox`           | [flexbox-css](../entries/flexbox-css.yml)                     | 1        | `layout-browser-rendering` |
| `margin`            | [margin-css](../entries/margin-css.yml)                       | 1        | `box-model-css`            |
| `padding`           | [padding-css](../entries/padding-css.yml)                     | 1        | `box-model-css`            |
| `display`           | [display-css](../entries/display-css.yml)                     | 1        | `layout-browser-rendering` |
| `alignment`         | [alignment-css](../entries/alignment-css.yml)                 | 2        | `layout-browser-rendering` |
| `gap`               | [gap-css](../entries/gap-css.yml)                             | 2        | `layout-browser-rendering` |
| `focus indicator`   | [focus-indicator-css](../entries/focus-indicator-css.yml)     | 2        | `pseudo-class-css`         |
| `reduced motion`    | [reduced-motion-css](../entries/reduced-motion-css.yml)       | 2        | `media-query-css`          |
| `custom property`   | [custom-property-css](../entries/custom-property-css.yml)     | 2        | `property-css`             |

## Grouping decisions

- Bootstrap 5.3’s default grid builds on Flexbox. Keep CSS Grid as a separate learning concept. Explain Bootstrap gutters alongside the reused padding card.
- Use class names and data attributes as examples of the concepts they teach. Include examples such as `.mt-3`, `.d-md-flex`, `data-bs-toggle`, and `data-bs-theme` on their relevant cards.
- Keep component names that describe distinct roles. For example, *tooltip* explains contextual hints and *popover* introduces a richer floating panel.
- Defer the reference-wide pass on Sass maps, build integration, plugin events, and migration terminology. This batch adds no associations.

## Implementation and verification

Go tests, vet, vocabulary validation, and diff checks passed. Two builds produced identical bytes. The export contains 1,588 rows with 17 columns. All 813 declared prerequisite links follow the generated `StudyOrder`.

Verified US IPA is present in 10 new entries. The other 37 await a reliable pronunciation source. Existing MP3 hashes matched the baseline. This batch generated no audio.

Validation reported the same 24 same-word warnings as before the import.

## Prerequisite follow-up

The tag now includes the HTML foundations for *form control* and *data attribute*, plus the CSS foundations used by the shared layout and accessibility cards. The current `tag:bootstrap` filter includes 73 cards.
