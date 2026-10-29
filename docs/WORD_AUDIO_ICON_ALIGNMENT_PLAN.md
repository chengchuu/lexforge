# Word Audio Icon Alignment Plan

## Status

Planned. Creating this document does not authorize implementation or changes to live Anki templates.

## Problem

The word-audio icon appears slightly above the pronunciation on AnkiDroid. On Mac Anki, alignment appears closer, but the icon is large relative to the pronunciation.

The current pronunciation row already uses flex alignment. Native replay-control dimensions, margins, and internal inline-element spacing may account for the difference. This is a working explanation based on the screenshots; inspect the rendered replay markup and computed styles before choosing the final selectors.

## Intended Result

- Keep pronunciation and word audio beside each other on both card sides.
- Use a 24px square visible word-audio icon on both clients.
- Center the icon vertically against the pronunciation container without a client-specific downward offset.
- Preserve a clickable area of at least 44px square. Center the visible icon within that area; recognize that the larger target may increase row height slightly.
- Allow long readings to wrap within their flex item while the icon remains beside the reading block, vertically centered against the whole block.
- Preserve clean layouts when pronunciation, audio, or both are absent.

## Implementation Approach

Update only the card styling documented in `README.md`, unless inspection shows that a small template adjustment is necessary.

1. Inspect native replay anchors and icon elements in Mac Anki and AnkiDroid. Record client versions and distinguish actual device inspection from mocked markup.
2. Scope all new sizing and alignment rules to the word-audio container inside `.pronunciation-row`. Support the actual native replay classes used by the target clients.
3. Center the replay control's contents with flex alignment. Normalize its vertical margin, padding, and line-height where they cause extra spacing; retain the intended touch-target dimensions.
4. Size the native icon element to 24px by 24px and remove inline baseline spacing, using block display where appropriate. Preserve its native artwork and theme colors.
5. Retain the existing pronunciation typography, horizontal gap, wrapping behavior, accessible control semantics, and visible keyboard focus.

Do not replace native playback handlers, duplicate audio controls, or introduce fixed pixel translations to compensate for one screenshot. Keep word autoplay and example playback unchanged. Do not modify the example-audio icon, vocabulary, stable IDs, note fields, TSV format, or recordings.

## Verification

- Check front and back templates with pronunciation and audio both present, either field absent, and both absent.
- Check short English IPA, Japanese kana, and long unbroken readings at narrow phone and desktop widths.
- Confirm the visible icon is 24px square, the target remains at least 44px square, and the icon and pronunciation container share a vertical center.
- Confirm no horizontal overflow, clipped icons, unintended empty-row spacing, or changes to example-audio styling.
- Check light and dark modes, keyboard focus, and readable text at increased display or font sizes.
- Verify word audio still plays once per side and can be replayed manually; example speech remains tap-triggered through its separate control.
- Run `git diff --check` and inspect the final diff for changes outside the intended template/style scope. Use isolated renderer checks where available; no TSV rebuild or audio generation is required for a styling-only change.
- Verify the final appearance and playback on Mac Anki and the Pixel 8 Pro running AnkiDroid. If either device is unavailable, report that live verification as outstanding rather than treating mocked layout checks as equivalent.

## Applying the Change

After implementation, copy the updated Styling into the existing LexForge note type and sync Anki Desktop and AnkiDroid. If the implementation also changes template markup, copy the affected templates as well. Do not recreate notes or reimport the TSV for this presentation-only change.

Keep the previous styling available for rollback. Do not change the user's live Anki collection automatically.
