# Anki Migration Guide

Use this guide to update an existing LexForge note type for images, code examples, example speech, learning order, and memory aids. For a new collection, follow [Set up Anki once](../README.md#set-up-anki-once).

Update your existing note type in Anki; editing repository documentation does not update installed templates. Keep existing notes and their `ID` values intact. Deleting and recreating notes discards their review history.

1. In **Tools → Manage Note Types → LexForge → Fields**, add `CodeExample` after `RelatedWords` if missing, then add `Image` after `CodeExample`. If migrating from an earlier layout, add any missing fields in the order shown in [Set up Anki once](../README.md#set-up-anki-once). Keep `ID` first and `StudyOrder` as the browser sort field.
2. Replace Back Template and Styling. If migrating from a note type predating example speech, also replace Front Template and follow the [example-speech setup](../README.md#example-speech).
3. [Copy the referenced images](../README.md#images) into the profile media folder. [Rebuild and reimport the TSV](../README.md#import-and-update) with **update existing notes** and **note type** match scope. Verify `CodeExample` in column 17, `Image` in column 18, and tags in column 19. [Apply the learning order](LEARNING_ORDER.md#apply-the-order-in-anki) to new cards.
4. Sync Anki Desktop and AnkiDroid, then perform the [image checks](../README.md#images) and [playback checks](../README.md#example-speech) on both clients.

Test the migration in a separate Anki profile before updating your main collection. Import twice and confirm that note and card counts stay unchanged. If your importer offers an option to preserve existing values when incoming fields are empty, disable it. In Browse, verify that blank Image, CodeExample, Context, Example, and memory-aid values clear the previous values. On a card with code, check indentation and horizontal scrolling and confirm that Read Aloud plays only the prose example.

Reimport the full export after changing related entries or their text so both display directions refresh. Inspect a card with no Context or Example: it should show neither empty sections nor an example replay control. Detailed Notes remain visible in full. Preserve ID matching, scheduling, and review history; no audio regeneration is needed.
