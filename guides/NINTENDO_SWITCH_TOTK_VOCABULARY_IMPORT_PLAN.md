# Nintendo Switch and TOTK vocabulary import

## Status and scope

Implemented on 2026-11-20 from the user-provided Nintendo Switch and Zelda TOTK vocabulary list. Of the 101 supplied candidates, 65 received new entries, 35 reused existing entries, and one, **phenomena**, was grouped with **phenomenon**. Six approved foundation words bring the new-entry total to 71. All new entries use English.

| Topic                          | New | Reused | Grouped |
| :----------------------------- | :-- | :----- | :------ |
| Nintendo Switch and hardware   | 11  | 0      | 0       |
| Gliding and movement           | 8   | 5      | 0       |
| Combat                         | 8   | 4      | 0       |
| Exploration and locations      | 7   | 9      | 1       |
| Construction and Zonai devices | 14  | 7      | 0       |
| Items and survival             | 7   | 7      | 0       |
| Game systems and quests        | 10  | 3      | 0       |
| Additional foundations         | 6   | 0      | 0       |

## Editorial decisions

- Retain separate physical senses of **drop**, **guard**, **hook**, and **emit**. The existing entries still cover discarding an item, software checks or hooks, or compiler output.
- Reuse the general sense of **restore** in `restore-tech` and retain its existing technical example. Keep **main quest** as the general singular term and **Main Quests** as the existing TOTK interface category.
- Teach **phenomena** in the notes of `phenomenon-story`. Keep **Regional Phenomena** as a separate card for the named quest.
- Use **HDMI Type A connector** and retain the supplied spelling **Type-A connector** in the notes. Use **TV Mode** for the original Switch and OLED Model. Nintendo Switch Lite lacks TV Mode. [Nintendo Switch support](https://en-americas-support.nintendo.com/app/answers/detail/a_id/22372/p/989/c/188) · [HDMI Licensing Administrator](https://www.hdmi.org/spec/passive)
- Use the Simplified Chinese game labels **完美回避**, **突击**, **调查四地区的异变**, **电池**, and **附加电池**. Preserve the existing **电池制造机** for **Crystal Refinery**. Each newly verified game label records its matching English and Simplified Chinese text key in a pinned version 1.2.1 game-text snapshot. The snapshot is a third-party extraction. Nintendo’s combat guide uses **Flurry Rush** and **Perfect Guard**. The game’s action prompt spells the former **Flurry rush**. [Paired game text](https://github.com/sh0wer1ee/TotkText/tree/1c66ac1a02edf2d73d10b0dc6f0f2de29da7b119) · [Nintendo combat guide](https://www.nintendo.com/en-gb/News/2023/July/More-tips-for-your-The-Legend-of-Zelda-Tears-of-the-Kingdom-adventures--2411783.html)
- Teach **quicksave**, **auto-save**, and **manual save** as general gameplay vocabulary. Their cards make no claim about dedicated save modes or exact labels in TOTK.
- The six previously existing files edited in this batch received one targeted change each. **Paraglider**, **emitter**, **detach**, and **dispenser** gained prerequisites. **Intact** gained the existing game and TOTK tags, and **phenomenon** gained the plural note.

## Entry index

Priority and prerequisites below are the values after import. New entries without a prerequisite use the global priority and stable-ID ordering.

### Nintendo Switch and hardware

| Learning item                            | Stable ID                                                     | Decision | Priority | Prerequisites                          |
| :--------------------------------------- | :------------------------------------------------------------ | :------- | :------- | :------------------------------------- |
| dock                                     | [dock-hardware](../entries/dock-hardware.yml)                 | New      | 1        | —                                      |
| monitor                                  | [monitor-hardware](../entries/monitor-hardware.yml)           | New      | 1        | —                                      |
| screen                                   | [screen-display](../entries/screen-display.yml)               | New      | 1        | —                                      |
| HDMI cable                               | [hdmi-cable-hardware](../entries/hdmi-cable-hardware.yml)     | New      | 2        | `hdmi-hardware`, `cable-hardware`      |
| connector                                | [connector-hardware](../entries/connector-hardware.yml)       | New      | 2        | `connect-general`                      |
| male-to-male                             | [male-to-male-hardware](../entries/male-to-male-hardware.yml) | New      | 2        | `connector-hardware`                   |
| Type-A connector → HDMI Type A connector | [hdmi-type-a-connector](../entries/hdmi-type-a-connector.yml) | New      | 3        | `hdmi-hardware`, `connector-hardware`  |
| TV Mode                                  | [tv-mode-switch](../entries/tv-mode-switch.yml)               | New      | 2        | `dock-hardware`, `hdmi-cable-hardware` |
| connect                                  | [connect-general](../entries/connect-general.yml)             | New      | 1        | —                                      |
| lose                                     | [lose-general](../entries/lose-general.yml)                   | New      | 1        | —                                      |
| drop                                     | [drop-fall-general](../entries/drop-fall-general.yml)         | New      | 1        | `fall-general`                         |

### Gliding and movement

| Learning item | Stable ID                                           | Decision | Priority | Prerequisites     |
| :------------ | :-------------------------------------------------- | :------- | :------- | :---------------- |
| updraft       | [updraft-flight](../entries/updraft-flight.yml)     | New      | 2        | —                 |
| gust          | [gust-weather](../entries/gust-weather.yml)         | New      | 2        | —                 |
| descent       | [descent-flight](../entries/descent-flight.yml)     | New      | 2        | `descend-general` |
| free fall     | [free-fall-flight](../entries/free-fall-flight.yml) | New      | 2        | `fall-general`    |
| land          | [land-motion](../entries/land-motion.yml)           | New      | 1        | —                 |
| fall damage   | [fall-damage-game](../entries/fall-damage-game.yml) | New      | 2        | `fall-general`    |
| runway        | [runway-flight](../entries/runway-flight.yml)       | New      | 2        | —                 |
| launch        | [launch-motion](../entries/launch-motion.yml)       | New      | 2        | —                 |
| glide         | [glide-general](../entries/glide-general.yml)       | Reuse    | 1        | —                 |
| Paraglider    | [paraglider-zelda](../entries/paraglider-zelda.yml) | Reuse    | 1        | `glide-general`   |
| stamina       | [stamina-general](../entries/stamina-general.yml)   | Reuse    | 1        | —                 |
| ascend        | [ascend-general](../entries/ascend-general.yml)     | Reuse    | 1        | —                 |
| descend       | [descend-general](../entries/descend-general.yml)   | Reuse    | 1        | —                 |

### Combat

| Learning item | Stable ID                                                 | Decision | Priority | Prerequisites                          |
| :------------ | :-------------------------------------------------------- | :------- | :------- | :------------------------------------- |
| side hop      | [side-hop-combat](../entries/side-hop-combat.yml)         | New      | 2        | `dodge-general`                        |
| backflip      | [backflip-movement](../entries/backflip-movement.yml)     | New      | 2        | —                                      |
| Perfect Dodge | [perfect-dodge-zelda](../entries/perfect-dodge-zelda.yml) | New      | 2        | `dodge-general`                        |
| Flurry Rush   | [flurry-rush-zelda](../entries/flurry-rush-zelda.yml)     | New      | 3        | `perfect-dodge-zelda`                  |
| guard         | [guard-combat](../entries/guard-combat.yml)               | New      | 1        | —                                      |
| Perfect Guard | [perfect-guard-zelda](../entries/perfect-guard-zelda.yml) | New      | 2        | `guard-combat`                         |
| weak point    | [weak-point-combat](../entries/weak-point-combat.yml)     | New      | 2        | —                                      |
| core          | [core-general](../entries/core-general.yml)               | New      | 1        | —                                      |
| dodge         | [dodge-general](../entries/dodge-general.yml)             | Reuse    | 1        | —                                      |
| fuse          | [fuse-general](../entries/fuse-general.yml)               | Reuse    | 1        | —                                      |
| emitter       | [emitter-engineering](../entries/emitter-engineering.yml) | Reuse    | 2        | `emit-physical`                        |
| Flame Emitter | [flame-emitter-zelda](../entries/flame-emitter-zelda.yml) | Reuse    | 2        | `flame-general`, `emitter-engineering` |

### Exploration and locations

| Learning item      | Stable ID                                                             | Decision | Priority | Prerequisites                                       |
| :----------------- | :-------------------------------------------------------------------- | :------- | :------- | :-------------------------------------------------- |
| observation deck   | [observation-deck-general](../entries/observation-deck-general.yml)   | New      | 2        | —                                                   |
| bottomless         | [bottomless-general](../entries/bottomless-general.yml)               | New      | 2        | —                                                   |
| pit                | [pit-terrain](../entries/pit-terrain.yml)                             | New      | 1        | —                                                   |
| mine cart          | [mine-cart-engineering](../entries/mine-cart-engineering.yml)         | New      | 2        | `cart-engineering`                                  |
| railway track      | [railway-track-engineering](../entries/railway-track-engineering.yml) | New      | 2        | —                                                   |
| opposite shore     | [opposite-shore-terrain](../entries/opposite-shore-terrain.yml)       | New      | 2        | —                                                   |
| Regional Phenomena | [regional-phenomena-zelda](../entries/regional-phenomena-zelda.yml)   | New      | 3        | `regional-general`, `phenomenon-story`              |
| shrine             | [shrine-architecture](../entries/shrine-architecture.yml)             | Reuse    | 1        | —                                                   |
| temple             | [temple-architecture](../entries/temple-architecture.yml)             | Reuse    | 1        | —                                                   |
| lookout            | [lookout-general](../entries/lookout-general.yml)                     | Reuse    | 1        | —                                                   |
| landing            | [landing-general](../entries/landing-general.yml)                     | Reuse    | 1        | —                                                   |
| Lookout Landing    | [lookout-landing-zelda](../entries/lookout-landing-zelda.yml)         | Reuse    | 2        | `lookout-general`, `landing-general`                |
| Skyview Tower      | [skyview-tower-zelda](../entries/skyview-tower-zelda.yml)             | Reuse    | 1        | `sky-terrain`, `view-general`, `tower-architecture` |
| cavern             | [cavern-architecture](../entries/cavern-architecture.yml)             | Reuse    | 2        | —                                                   |
| regional           | [regional-general](../entries/regional-general.yml)                   | Reuse    | 2        | —                                                   |
| phenomenon         | [phenomenon-story](../entries/phenomenon-story.yml)                   | Reuse    | 2        | —                                                   |
| phenomena          | [phenomenon-story](../entries/phenomenon-story.yml)                   | Grouped  | 2        | —                                                   |

### Construction and Zonai devices

| Learning item    | Stable ID                                                       | Decision | Priority | Prerequisites                                 |
| :--------------- | :-------------------------------------------------------------- | :------- | :------- | :-------------------------------------------- |
| raft             | [raft-watercraft](../entries/raft-watercraft.yml)               | New      | 1        | —                                             |
| sail             | [sail-watercraft](../entries/sail-watercraft.yml)               | New      | 1        | —                                             |
| propulsion       | [propulsion-engineering](../entries/propulsion-engineering.yml) | New      | 3        | —                                             |
| disassemble      | [disassemble-general](../entries/disassemble-general.yml)       | New      | 2        | `assemble-general`                            |
| reattach         | [reattach-general](../entries/reattach-general.yml)             | New      | 2        | `attach-general`                              |
| decompose        | [decompose-general](../entries/decompose-general.yml)           | New      | 2        | —                                             |
| wiggle           | [wiggle-general](../entries/wiggle-general.yml)                 | New      | 2        | —                                             |
| hook             | [hook-hardware](../entries/hook-hardware.yml)                   | New      | 1        | —                                             |
| suspend          | [suspend-hang-general](../entries/suspend-hang-general.yml)     | New      | 2        | —                                             |
| dispense         | [dispense-general](../entries/dispense-general.yml)             | New      | 1        | —                                             |
| capsule          | [capsule-general](../entries/capsule-general.yml)               | New      | 2        | —                                             |
| charge           | [charge-electricity](../entries/charge-electricity.yml)         | New      | 1        | —                                             |
| crystallized     | [crystallized-general](../entries/crystallized-general.yml)     | New      | 2        | `crystallize-general`                         |
| illuminate       | [illuminate-general](../entries/illuminate-general.yml)         | New      | 2        | —                                             |
| assemble         | [assemble-general](../entries/assemble-general.yml)             | Reuse    | 1        | —                                             |
| attach           | [attach-general](../entries/attach-general.yml)                 | Reuse    | 1        | —                                             |
| detach           | [detach-general](../entries/detach-general.yml)                 | Reuse    | 1        | `attach-general`                              |
| intact           | [intact-tech](../entries/intact-tech.yml)                       | Reuse    | 1        | —                                             |
| dispenser        | [dispenser-engineering](../entries/dispenser-engineering.yml)   | Reuse    | 1        | `device-engineering`, `dispense-general`      |
| Device Dispenser | [device-dispenser-zelda](../entries/device-dispenser-zelda.yml) | Reuse    | 2        | `device-engineering`, `dispenser-engineering` |
| refinery         | [refinery-architecture](../entries/refinery-architecture.yml)   | Reuse    | 2        | —                                             |

### Items and survival

| Learning item       | Stable ID                                                             | Decision | Priority | Prerequisites                                |
| :------------------ | :-------------------------------------------------------------------- | :------- | :------- | :------------------------------------------- |
| Heart Container     | [heart-container-zelda](../entries/heart-container-zelda.yml)         | New      | 2        | —                                            |
| vitality            | [vitality-general](../entries/vitality-general.yml)                   | New      | 2        | —                                            |
| Brightbloom Seed    | [brightbloom-seed-zelda](../entries/brightbloom-seed-zelda.yml)       | New      | 2        | `seed-general`                               |
| Energy Cell         | [energy-cell-zelda](../entries/energy-cell-zelda.yml)                 | New      | 2        | —                                            |
| Energy Well         | [energy-well-zelda](../entries/energy-well-zelda.yml)                 | New      | 3        | `energy-cell-zelda`                          |
| Crystallized Charge | [crystallized-charge-zelda](../entries/crystallized-charge-zelda.yml) | New      | 3        | `crystallized-general`, `charge-electricity` |
| campfire            | [campfire-general](../entries/campfire-general.yml)                   | New      | 1        | —                                            |
| blessing            | [blessing-general](../entries/blessing-general.yml)                   | Reuse    | 1        | —                                            |
| Light of Blessing   | [light-of-blessing-zelda](../entries/light-of-blessing-zelda.yml)     | Reuse    | 2        | `light-engineering`, `blessing-general`      |
| resistance          | [resistance-general](../entries/resistance-general.yml)               | Reuse    | 1        | —                                            |
| cold resistance     | [cold-resistance-combat](../entries/cold-resistance-combat.yml)       | Reuse    | 1        | `resistance-general`                         |
| Spicy Pepper        | [spicy-pepper-zelda](../entries/spicy-pepper-zelda.yml)               | Reuse    | 2        | `spicy-general`, `pepper-vegetable`          |
| Crystal Refinery    | [crystal-refinery-zelda](../entries/crystal-refinery-zelda.yml)       | Reuse    | 2        | `crystal-general`, `refinery-architecture`   |
| restore             | [restore-tech](../entries/restore-tech.yml)                           | Reuse    | 1        | —                                            |

### Game systems and quests

| Learning item | Stable ID                                             | Decision | Priority | Prerequisites |
| :------------ | :---------------------------------------------------- | :------- | :------- | :------------ |
| fast travel   | [fast-travel-game](../entries/fast-travel-game.yml)   | New      | 2        | —             |
| quicksave     | [quicksave-game](../entries/quicksave-game.yml)       | New      | 2        | `save-game`   |
| auto-save     | [auto-save-game](../entries/auto-save-game.yml)       | New      | 2        | `save-game`   |
| manual save   | [manual-save-game](../entries/manual-save-game.yml)   | New      | 2        | `save-game`   |
| optional      | [optional-general](../entries/optional-general.yml)   | New      | 1        | —             |
| mandatory     | [mandatory-general](../entries/mandatory-general.yml) | New      | 1        | —             |
| collectible   | [collectible-game](../entries/collectible-game.yml)   | New      | 2        | —             |
| main quest    | [main-quest-game](../entries/main-quest-game.yml)     | New      | 2        | `quest-game`  |
| detour        | [detour-general](../entries/detour-general.yml)       | New      | 2        | —             |
| progression   | [progression-game](../entries/progression-game.yml)   | New      | 2        | —             |
| recall        | [recall-general](../entries/recall-general.yml)       | Reuse    | 2        | —             |
| unlock        | [unlock-general](../entries/unlock-general.yml)       | Reuse    | 1        | —             |
| objective     | [objective-ui](../entries/objective-ui.yml)           | Reuse    | 1        | —             |

### Additional foundations

| Learning item | Stable ID                                                 | Decision | Priority | Prerequisites |
| :------------ | :-------------------------------------------------------- | :------- | :------- | :------------ |
| HDMI          | [hdmi-hardware](../entries/hdmi-hardware.yml)             | New      | 1        | —             |
| cable         | [cable-hardware](../entries/cable-hardware.yml)           | New      | 1        | —             |
| fall          | [fall-general](../entries/fall-general.yml)               | New      | 1        | —             |
| emit          | [emit-physical](../entries/emit-physical.yml)             | New      | 1        | —             |
| crystallize   | [crystallize-general](../entries/crystallize-general.yml) | New      | 1        | —             |
| save          | [save-game](../entries/save-game.yml)                     | New      | 1        | —             |

## Pronunciation and audio

Added 47 sourced American IPA transcriptions and one labeled British transcription for **reattach**. Notes contain dictionary references and access dates. Verified full-item IPA was unavailable for 23 entries during this pass. Those fields are unset:

**HDMI cable**, **male-to-male**, **HDMI Type A connector**, **TV Mode**, **fall damage**, **side hop**, **Perfect Dodge**, **Flurry Rush**, **Perfect Guard**, **weak point**, **mine cart**, **railway track**, **opposite shore**, **Regional Phenomena**, **Heart Container**, **Brightbloom Seed**, **Energy Cell**, **Energy Well**, **Crystallized Charge**, **quicksave**, **auto-save**, **manual save**, **main quest**.

Existing recordings and archived audio were preserved byte for byte. The 71 new entries have no recording. IPA and the exported example-speech locale are independent of recorded word audio.

## Verification

- `go test ./...`, `go vet ./...`, and vocabulary validation passed.
- Two builds produced identical TSV bytes. The export contains 1,757 data rows with 17 columns and unique ranks. All 934 prerequisite links place the foundation before the dependent card.
- `tag:nintendo-switch` contains 14 entries and `tag:zelda-totk` contains 612 entries. Both filters contain every prerequisite of their members. Their shared card is **fall**.
- Validation reported 30 repeated-word warnings, including four additional warnings for the intended senses of **drop**, **guard**, **hook**, and **emit**.
- Byte comparison found only the six documented edits among existing entry, active media, and archive files. The import did not touch Anki.

Build SHA-256

```text
9226395f579d5522ecee299b86692afab04e3bb3106713e1ec125c3de6b5b1eb
```
