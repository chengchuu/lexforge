# Anki Migration Guide

Use this guide to update an existing LexForge note type for example speech, learning order, and memory aids. For a new collection, follow [Set up Anki once](../README.md#set-up-anki-once).

Update your existing note type in Anki; editing repository documentation does not update installed templates. Keep existing notes and their `ID` values intact. Deleting and recreating notes discards their review history.

1. In **Tools → Manage Note Types → LexForge → Fields**, add `TTSLanguage` after `Audio` and `StudyOrder` after `TTSLanguage` if missing. Append `Etymology`, `WordFormation`, `Mnemonic`, and `RelatedWords` after `StudyOrder`, in that order. Select `StudyOrder` as the browser sort field. Keep `ID` first and confirm the field order shown in [Set up Anki once](../README.md#set-up-anki-once).
2. Replace the front, back, and styling with the templates in that section. Follow the [example-speech setup](../README.md#example-speech) for audio options and voice requirements.
3. [Rebuild and reimport the TSV](../README.md#import-and-update) with **update existing notes** and **note type** match scope. Verify all sixteen fields and tags in column 17 before importing. [Apply the learning order](LEARNING_ORDER.md#apply-the-order-in-anki) to new cards.
4. Sync Anki Desktop and AnkiDroid, then perform the [playback checks](../README.md#example-speech) on both clients.

Test the migration in a separate Anki profile before updating your main collection. Import twice and confirm that note and card counts stay unchanged. Check that an empty Context, Example, or memory-aid value clears the previous value: disable any importer option that preserves existing fields when incoming fields are empty. Verify the result in Browse rather than assuming blanks were applied.

Reimport the full export after changing related entries or their text so both display directions refresh. Inspect a card with no Context or Example: it should show neither empty sections nor an example replay control. Detailed Notes remain visible in full. Preserve ID matching, scheduling, and review history; no audio regeneration is needed.
