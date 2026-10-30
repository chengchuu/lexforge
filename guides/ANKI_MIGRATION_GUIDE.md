# Anki Migration Guide

Use this guide to update an existing LexForge note type for example speech and learning order. For a new collection, follow [Set up Anki once](../README.md#set-up-anki-once).

Update your existing note type in Anki; editing repository documentation does not update installed templates. Keep existing notes and their `ID` values intact. Deleting and recreating notes discards their review history.

1. In **Tools → Manage Note Types → LexForge → Fields**, add `TTSLanguage` after `Audio` and `StudyOrder` after `TTSLanguage` if missing. Select `StudyOrder` as the browser sort field. Keep `ID` first and confirm the field order shown in [Set up Anki once](../README.md#set-up-anki-once).
2. Replace the front, back, and styling with the templates in that section. Follow the [example-speech setup](../README.md#example-speech) for audio options and voice requirements.
3. [Rebuild and reimport the TSV](../README.md#import-and-update) with **update existing notes** and **note type** match scope. Verify all twelve fields and the tags column before importing. [Apply the learning order](../README.md#apply-the-order-in-anki) to new cards.
4. Sync Anki Desktop and AnkiDroid, then perform the [playback checks](../README.md#example-speech) on both clients.
