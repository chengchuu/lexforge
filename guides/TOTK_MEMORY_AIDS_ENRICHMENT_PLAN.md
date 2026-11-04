# TOTK memory-aid enrichment plan

Status: **Completed — all proposals approved by the user and implemented.**

This plan reviews the 546 English entries added by the [TOTK vocabulary import](TOTK_VOCABULARY_IMPORT.md). It proposes exact Chinese card text for useful memory aids and additional direct associations. It does not seek a filling percentage: an entry can keep any or all optional fields empty.

This file was initially created under planning-only authorization. The user subsequently approved the complete plan with "I approve it." All WF01–WF48, ETY01–ETY18, N01–N20, and R10–R29 are now implemented. Existing import relationships R01–R09 remain intact. No schema, implementation, template, dependency, pronunciation-audio, Git-index, or Anki changes were made.

## Review navigation

- [Scope and counts](#scope-and-counts)
- [Keep the existing aids](#keep-the-existing-aids)
- [Word Formation drafts](#word-formation-drafts)
- [Etymology drafts](#etymology-drafts)
- [Mnemonic drafts](#mnemonic-drafts)
- [Additional related pairs](#additional-related-pairs)
- [Deferred naming questions](#deferred-naming-questions)
- [Source register](#source-register)
- [Complete review ledger](#complete-review-ledger)
- [Approval and implementation boundaries](#approval-and-implementation-boundaries)
- [Verification record](#verification-record)

## Scope and counts

The planning snapshot contains 1,241 entries. Counts in this section retain the planning baseline and proposed scope; actual implementation results are recorded under Verification record. The reviewed batch is the 546 new IDs linked from the import record, excluding its seven reused IDs: abandoned-workflow, curry-food, mushroom-food, restore-tech, rice-food, roast-food, and seafood-food. These seven and the rest of the collection were consulted as context, not selected for broad rewriting.

| Item                                                     | Count and meaning                                                                         |
| :------------------------------------------------------- | :---------------------------------------------------------------------------------------- |
| Batch entries screened                                   | 546; current Word, Meaning, optional aid fields, and Notes read for suitability           |
| Planning baseline: Etymology / Word Formation / Mnemonic | 0 / 6 / 0 populated fields in this batch                                                  |
| Planning baseline: approved relationships                | 9 direct pairs; display on 16 new entries and one existing entry                          |
| Proposed Word Formation                                  | 48 absent fields: 47 additions and one scoped move from Notes                             |
| Proposed Etymology                                       | 18 absent fields; historical research focused on these selected words                     |
| Proposed Mnemonic                                        | 20 absent fields; original imagery or usage distinctions                                  |
| Unique entries with text proposals                       | 81; categories overlap, so 48 + 18 + 20 does not equal the unique-entry count             |
| Additional relationship proposals                        | 20 pairs; 37 new-batch endpoints and emit-tech from the existing collection               |
| Unique new-batch entries with any proposal               | 98; includes text and relationship-only proposals                                         |
| New-batch entries with no proposed edit                  | 448; includes existing aids retained unchanged and entries left blank                     |
| Evidence register                                        | 79 external reference records, plus explicit local-entry evidence                         |
| Initial planning stage                                   | Only this Markdown plan was created; subsequent approved implementation is recorded below |

Screening all 546 entries does not mean researching every historical origin. Historical claims below were checked against the cited headword entries; ordinary definitions in the existing collection serve as local evidence for other editorial comparisons. A blank decision means no addition is justified in this pass, not that a word has no history or could never benefit from an aid.

## Keep the existing aids

Retain the following six Word Formation values exactly. They are not replacements or new approvals in this plan.

| Stable ID                                                   | Word           | Current text — keep                                                                       |
| :---------------------------------------------------------- | :------------- | :---------------------------------------------------------------------------------------- |
| [broadsword-weapon](../entries/broadsword-weapon.yml)       | broadsword     | broad（宽阔的）+ sword（剑）。                                                            |
| [durability-equipment](../entries/durability-equipment.yml) | durability     | durable 是形容词，durability 是对应的抽象名词；-ble 结尾的形容词常有 -bility 形式的名词。 |
| [emitter-engineering](../entries/emitter-engineering.yml)   | emitter        | emit（发出）+ -er（执行者或装置）；这里指发射装置。                                       |
| [lightroot-zelda](../entries/lightroot-zelda.yml)           | Lightroot      | light（光）+ root（根）；这是拼写构成说明，不是历史词源。                                 |
| [sky-island-zelda](../entries/sky-island-zelda.yml)         | sky island     | sky（天空）+ island（岛屿）。                                                             |
| [steering-stick-zelda](../entries/steering-stick-zelda.yml) | Steering Stick | steering（控制方向）+ stick（杆）；组合表示操纵杆。                                       |

Also preserve the nine approved import associations R01–R09, all current prerequisites, learning priorities, stable IDs, meanings, examples, and audio references. A prerequisite does not itself create a Related Words display. Associations remain direct neighbors, with no transitive expansion.

## Word Formation drafts

Each block below is the exact proposed plain-text `word_formation` value. An absent field is added only after its proposal is approved. These are modern patterns, word-family correspondences, or visibly recognizable compound structures; a constituent analysis does not establish when or why Nintendo coined a name. External sources support the stated affix or lexical sense; local entries and pinned game keys support the current target spelling and localized meaning.

Review batches **WF01–WF24** and **WF25–WF48** were both approved and implemented. The exact reviewed values below are preserved.

### WF01 — Autobuild

Target: [autobuild-zelda](../entries/autobuild-zelda.yml). Current meaning: 蓝图；复现已记录组合设计的能力。

Action: Move the existing construction explanation out of Notes; see the exact removal below.

```plain
auto- (自动) + build (建造)。这是可辨认的英语构成说明，不表示已核实命名史。
```

Evidence: [P09], [G01]. Local: [autobuild-zelda](../entries/autobuild-zelda.yml).

WF01 approval includes removing only this existing standalone paragraph from Notes, leaving all source and game-text references intact. The new value preserves its caution about naming history:

```plain
auto- 表示自动，build 表示建造；这里不推断虚构技术的历史来源。
```

### WF02 — Ultrahand

Target: [ultrahand-zelda](../entries/ultrahand-zelda.yml). Current meaning: 究极手；可抓取、移动和粘接许多物体的能力。

Action: Add the currently absent `word_formation` field.

```plain
ultra- (超出通常范围) + hand (手)。用来理解名称的表面构成，不推断官方命名来源。
```

Evidence: [P08], [G02]. Local: [ultrahand-zelda](../entries/ultrahand-zelda.yml).

### WF03 — unfuse

Target: [unfuse-general](../entries/unfuse-general.yml). Current meaning: 解除融合

Action: Add the currently absent `word_formation` field.

```plain
un- (逆转某动作) + fuse (融合)，表示解除融合；仅凭词义不能推出是否返还材料。
```

Evidence: [P01]. Local: [unfuse-general](../entries/unfuse-general.yml).

### WF04 — unlock

Target: [unlock-general](../entries/unlock-general.yml). Current meaning: 解锁

Action: Add the currently absent `word_formation` field.

```plain
un- (解除或逆转) + lock (锁住)，表示解锁或解除限制。
```

Evidence: [P01]. Local: [unlock-general](../entries/unlock-general.yml).

### WF05 — recharge

Target: [recharge-general](../entries/recharge-general.yml). Current meaning: 充能

Action: Add the currently absent `word_formation` field.

```plain
re- (再次) + charge (充电、装载能量)，表示重新充能。
```

Evidence: [P02], [L32]. Local: [recharge-general](../entries/recharge-general.yml).

### WF06 — underground

Target: [underground-general](../entries/underground-general.yml). Current meaning: 地下的

Action: Add the currently absent `word_formation` field.

```plain
under (在……下面) + ground (地面)，表示在地下或地下的。
```

Evidence: [L19]. Local: [underground-general](../entries/underground-general.yml).

### WF07 — undead

Target: [undead-combat](../entries/undead-combat.yml). Current meaning: 不死的

Action: Add the currently absent `word_formation` field.

```plain
un- (非、不) + dead (死的)。奇幻语境指亡灵般的非正常生存状态，不简单等于普通的活着。
```

Evidence: [P01], [L18]. Local: [undead-combat](../entries/undead-combat.yml).

### WF08 — fearless

Target: [fearless-general](../entries/fearless-general.yml). Current meaning: 无畏的

Action: Add the currently absent `word_formation` field.

```plain
fear (恐惧) + -less (没有……的)，表示无畏的。
```

Evidence: [P03], [L30]. Local: [fearless-general](../entries/fearless-general.yml).

### WF09 — selfless

Target: [selfless-general](../entries/selfless-general.yml). Current meaning: 无私的

Action: Add the currently absent `word_formation` field.

```plain
self (自我) + -less；表示不以自身利益为中心，不是字面上的没有自我。
```

Evidence: [P03], [L31]. Local: [selfless-general](../entries/selfless-general.yml).

### WF10 — ruthless

Target: [ruthless-general](../entries/ruthless-general.yml). Current meaning: 无情的

Action: Add the currently absent `word_formation` field.

```plain
ruth (怜悯，现已少见的名词) + -less (没有……的)，表示无情的；这里的 ruth 不是人名 Ruth。
```

Evidence: [L15]. Local: [ruthless-general](../entries/ruthless-general.yml).

### WF11 — powerful

Target: [powerful-general](../entries/powerful-general.yml). Current meaning: 强大的

Action: Add the currently absent `word_formation` field.

```plain
power (力量、能力) + -ful (具有……特质的)，表示强大的。
```

Evidence: [P06]. Local: [powerful-general](../entries/powerful-general.yml).

### WF12 — skilled

Target: [skilled-general](../entries/skilled-general.yml). Current meaning: 熟练的、有技巧的

Action: Add the currently absent `word_formation` field.

```plain
skill (技能) + -ed (具有……的)，表示有技能的、熟练的。
```

Evidence: [P10]. Local: [skilled-general](../entries/skilled-general.yml).

### WF13 — armored

Target: [armored-combat](../entries/armored-combat.yml). Current meaning: 装甲覆盖的

Action: Add the currently absent `word_formation` field.

```plain
armor (护甲) + -ed (具有……的)，表示覆盖护甲的。
```

Evidence: [P10]. Local: [armored-combat](../entries/armored-combat.yml).

### WF14 — regional

Target: [regional-general](../entries/regional-general.yml). Current meaning: 地区性的

Action: Add the currently absent `word_formation` field.

```plain
region (地区) + -al (与……有关的)，表示地区性的。
```

Evidence: [P11]. Local: [regional-general](../entries/regional-general.yml).

### WF15 — headband

Target: [headband-clothing](../entries/headband-clothing.yml). Current meaning: 头带

Action: Add the currently absent `word_formation` field.

```plain
head (头) + band (带)，表示戴在头上的带状物。
```

Evidence: [L01]. Local: [headband-clothing](../entries/headband-clothing.yml).

### WF16 — headdress

Target: [headdress-clothing](../entries/headdress-clothing.yml). Current meaning: 头饰

Action: Add the currently absent `word_formation` field.

```plain
head (头) + dress (装束)，表示头部的覆盖物或装饰；这里的 dress 不限于连衣裙。
```

Evidence: [L02]. Local: [headdress-clothing](../entries/headdress-clothing.yml).

### WF17 — waistguard

Target: [waistguard-clothing](../entries/waistguard-clothing.yml). Current meaning: 护腰

Action: Add the currently absent `word_formation` field.

```plain
waist (腰部) + guard (防护物)，表示保护腰部的装备。
```

Evidence: [L28], [L29]. Local: [waistguard-clothing](../entries/waistguard-clothing.yml).

### WF18 — highlands

Target: [highlands-terrain](../entries/highlands-terrain.yml). Current meaning: 高地

Action: Add the currently absent `word_formation` field.

```plain
high (高的) + land (土地) + 复数 -s，表示高地区域。
```

Evidence: [L13]. Local: [highlands-terrain](../entries/highlands-terrain.yml).

### WF19 — wetlands

Target: [wetlands-terrain](../entries/wetlands-terrain.yml). Current meaning: 湿地

Action: Add the currently absent `word_formation` field.

```plain
wet (湿的) + land (土地) + 复数 -s，表示湿地区域。
```

Evidence: [L14]. Local: [wetlands-terrain](../entries/wetlands-terrain.yml).

### WF20 — waterfall

Target: [waterfall-terrain](../entries/waterfall-terrain.yml). Current meaning: 瀑布

Action: Add the currently absent `word_formation` field.

```plain
water (水) + fall (落下)，表示从高处跌落的水流，即瀑布。
```

Evidence: [L11]. Local: [waterfall-terrain](../entries/waterfall-terrain.yml).

### WF21 — staircase

Target: [staircase-architecture](../entries/staircase-architecture.yml). Current meaning: 楼梯

Action: Add the currently absent `word_formation` field.

```plain
stair (楼梯踏步) + case (框架、围护结构这一义项)，表示楼梯整体；case 不是案件义。
```

Evidence: [L12]. Local: [staircase-architecture](../entries/staircase-architecture.yml).

### WF22 — foothill

Target: [foothill-terrain](../entries/foothill-terrain.yml). Current meaning: 山麓

Action: Add the currently absent `word_formation` field.

```plain
foot (山脚、底部) + hill (丘陵)，表示高山或山脉脚下的丘陵。
```

Evidence: [L03]. Local: [foothill-terrain](../entries/foothill-terrain.yml).

### WF23 — airborne

Target: [airborne-combat](../entries/airborne-combat.yml). Current meaning: 空中的

Action: Add the currently absent `word_formation` field.

```plain
air (空气、空中) + borne (bear 表示承载时的过去分词)，表示在空中的或经空中运送的；不是 born (出生)。
```

Evidence: [L04]. Local: [airborne-combat](../entries/airborne-combat.yml).

### WF24 — homing

Target: [homing-engineering](../entries/homing-engineering.yml). Current meaning: 自动导引的；能自动追踪目标的。

Action: Add the currently absent `word_formation` field.

```plain
home 用作动词可表示朝目标归向；homing 是其 -ing 形式，此处表示自动朝目标导引，不是返回住宅。
```

Evidence: [L08]. Local: [homing-engineering](../entries/homing-engineering.yml).

### WF25 — stabilizer

Target: [stabilizer-engineering](../entries/stabilizer-engineering.yml). Current meaning: 稳定器

Action: Add the currently absent `word_formation` field.

```plain
stabilize (使稳定) 去掉词尾 e 后接 -er (执行者或装置)，构成 stabilizer；这里表示稳定装置。
```

Evidence: [L06], [P04]. Local: [stabilizer-engineering](../entries/stabilizer-engineering.yml).

### WF26 — dispenser

Target: [dispenser-engineering](../entries/dispenser-engineering.yml). Current meaning: 分配器；按需放出物品的装置。

Action: Add the currently absent `word_formation` field.

```plain
dispense (分配、发放) + -er (执行者或装置)，去掉词尾 e 后构成 dispenser。
```

Evidence: [L07], [P04]. Local: [dispenser-engineering](../entries/dispenser-engineering.yml).

### WF27 — refinery

Target: [refinery-architecture](../entries/refinery-architecture.yml). Current meaning: 精炼设施

Action: Add the currently absent `word_formation` field.

```plain
refine (精炼) + -ery (场所)，拼写为 refinery，表示进行精炼的设施。
```

Evidence: [L05]. Local: [refinery-architecture](../entries/refinery-architecture.yml).

### WF28 — miner

Target: [miner-general](../entries/miner-general.yml). Current meaning: 矿工；采矿的人。

Action: Add the currently absent `word_formation` field.

```plain
mine (开矿) + -er (从事者)，去掉词尾 e，表示矿工。
```

Evidence: [P04]. Local: [miner-general](../entries/miner-general.yml).

### WF29 — traveler

Target: [traveler-general](../entries/traveler-general.yml). Current meaning: 旅行者；旅人。

Action: Add the currently absent `word_formation` field.

```plain
travel (旅行) + -er (从事者)，表示旅行者。
```

Evidence: [P04]. Local: [traveler-general](../entries/traveler-general.yml).

### WF30 — enhancement

Target: [enhancement-general](../entries/enhancement-general.yml). Current meaning: 强化、增强

Action: Add the currently absent `word_formation` field.

```plain
enhance (增强) + -ment (动作、结果或状态)，表示增强或改进。
```

Evidence: [L16], [P05]. Local: [enhancement-general](../entries/enhancement-general.yml).

### WF31 — settlement

Target: [settlement-architecture](../entries/settlement-architecture.yml). Current meaning: 定居点、聚落

Action: Add the currently absent `word_formation` field.

```plain
settle (定居) + -ment；此处指人们定居形成的聚落，不取结算或和解义。
```

Evidence: [L17], [P05]. Local: [settlement-architecture](../entries/settlement-architecture.yml).

### WF32 — kingdom

Target: [kingdom-general](../entries/kingdom-general.yml). Current meaning: 王国；由国王或女王统治的国家。

Action: Add the currently absent `word_formation` field.

```plain
king (国王) + -dom (统治领域)，表示王国；现实中也可由女王统治。
```

Evidence: [P07]. Local: [kingdom-general](../entries/kingdom-general.yml).

### WF33 — descendant

Target: [descendant-general](../entries/descendant-general.yml). Current meaning: 后裔

Action: Add the currently absent `word_formation` field.

```plain
与 descend 的传承义对应，带有 -ant 名词词尾；可沿族谱向下延续来理解后裔，不把它等同于下降动作。
```

Evidence: [L09]. Local: [descendant-general](../entries/descendant-general.yml).

### WF34 — mysterious

Target: [mysterious-general](../entries/mysterious-general.yml). Current meaning: 神秘的

Action: Add the currently absent `word_formation` field.

```plain
词形对应: mystery (谜、神秘之事) → mysterious (神秘的)；可识别 -ous 形容词词尾。这里说明现代词形，不替代历史词源。
```

Evidence: [L10], [P12]. Local: [mysterious-general](../entries/mysterious-general.yml).

### WF35 — courageous

Target: [courageous-general](../entries/courageous-general.yml). Current meaning: 勇敢的

Action: Add the currently absent `word_formation` field.

```plain
courage (勇气) + -ous (具有……特质的)，表示勇敢的。
```

Evidence: [P12]. Local: [courageous-general](../entries/courageous-general.yml).

### WF36 — critical hit

Target: [critical-hit-combat](../entries/critical-hit-combat.yml). Current meaning: 暴击

Action: Add the currently absent `word_formation` field.

```plain
critical 修饰 hit (命中、打击)；整体是暴击这一术语，不能从构成说明推出具体伤害倍率。
```

Evidence: Local entry evidence; editorial construction/learning image. Local: [critical-hit-combat](../entries/critical-hit-combat.yml).

### WF37 — attack power

Target: [attack-power-combat](../entries/attack-power-combat.yml). Current meaning: 攻击力

Action: Add the currently absent `word_formation` field.

```plain
attack (攻击) + power (力量、能力)，整体表示攻击力。
```

Evidence: Local entry evidence; editorial construction/learning image. Local: [attack-power-combat](../entries/attack-power-combat.yml), [powerful-general](../entries/powerful-general.yml).

### WF38 — cooking pot

Target: [cooking-pot-cooking](../entries/cooking-pot-cooking.yml). Current meaning: 烹饪锅

Action: Add the currently absent `word_formation` field.

```plain
cooking (烹饪) 修饰 pot (锅)，整体表示烹饪用锅。
```

Evidence: Local entry evidence; editorial construction/learning image. Local: [cooking-pot-cooking](../entries/cooking-pot-cooking.yml), [cook-cooking](../entries/cook-cooking.yml), [pot-general](../entries/pot-general.yml).

### WF39 — one-handed

Target: [one-handed-combat](../entries/one-handed-combat.yml). Current meaning: 单手的

Action: Add the currently absent `word_formation` field.

```plain
one (一) + hand (手) + -ed (具有……的)，表示单手操作的。
```

Evidence: [P10]. Local: [one-handed-combat](../entries/one-handed-combat.yml).

### WF40 — two-handed

Target: [two-handed-combat](../entries/two-handed-combat.yml). Current meaning: 双手的

Action: Add the currently absent `word_formation` field.

```plain
two (二) + hand (手) + -ed (具有……的)，表示双手操作的。
```

Evidence: [P10]. Local: [two-handed-combat](../entries/two-handed-combat.yml).

### WF41 — Hover Stone

Target: [hover-stone-zelda](../entries/hover-stone-zelda.yml). Current meaning: 浮游石；《王国之泪》中的左纳乌装置。

Action: Add the currently absent `word_formation` field.

```plain
hover (悬停) + stone (石头)。说明英文名称构成；游戏简体中文名为浮游石。
```

Evidence: [G03]. Local: [hover-stone-zelda](../entries/hover-stone-zelda.yml), [hover-general](../entries/hover-general.yml), [stone-general](../entries/stone-general.yml).

### WF42 — Flame Emitter

Target: [flame-emitter-zelda](../entries/flame-emitter-zelda.yml). Current meaning: 火龙头；《王国之泪》中的左纳乌装置。

Action: Add the currently absent `word_formation` field.

```plain
flame (火焰) + emitter (发射装置)。字面可理解为火焰发射装置；游戏简体中文名为火龙头。
```

Evidence: [G04]. Local: [flame-emitter-zelda](../entries/flame-emitter-zelda.yml), [flame-general](../entries/flame-general.yml), [emitter-engineering](../entries/emitter-engineering.yml).

### WF43 — Frost Emitter

Target: [frost-emitter-zelda](../entries/frost-emitter-zelda.yml). Current meaning: 冰龙头；《王国之泪》中的左纳乌装置。

Action: Add the currently absent `word_formation` field.

```plain
frost (冰霜) + emitter (发射装置)。说明名称构成；游戏简体中文名为冰龙头。
```

Evidence: [G05]. Local: [frost-emitter-zelda](../entries/frost-emitter-zelda.yml), [frost-general](../entries/frost-general.yml), [emitter-engineering](../entries/emitter-engineering.yml).

### WF44 — Shock Emitter

Target: [shock-emitter-zelda](../entries/shock-emitter-zelda.yml). Current meaning: 雷龙头；《王国之泪》中的左纳乌装置。

Action: Add the currently absent `word_formation` field.

```plain
shock 在此取电击义，emitter 表示发射装置；游戏简体中文名为雷龙头，不取惊讶义。
```

Evidence: [G06]. Local: [shock-emitter-zelda](../entries/shock-emitter-zelda.yml), [emitter-engineering](../entries/emitter-engineering.yml).

### WF45 — Beam Emitter

Target: [beam-emitter-zelda](../entries/beam-emitter-zelda.yml). Current meaning: 光线头；《王国之泪》中的左纳乌装置。

Action: Add the currently absent `word_formation` field.

```plain
beam 在此指光束，emitter 表示发射装置；游戏简体中文名为光线头。
```

Evidence: [G07]. Local: [beam-emitter-zelda](../entries/beam-emitter-zelda.yml), [beam-engineering](../entries/beam-engineering.yml), [emitter-engineering](../entries/emitter-engineering.yml).

### WF46 — Skyview Tower

Target: [skyview-tower-zelda](../entries/skyview-tower-zelda.yml). Current meaning: 鸟望台；用于调查区域地图的塔状设施。

Action: Add the currently absent `word_formation` field.

```plain
名称中可识别 sky (天空)、view (视野) 和 tower (塔)；游戏简体中文名为鸟望台。此处不推断命名史。
```

Evidence: [G08]. Local: [skyview-tower-zelda](../entries/skyview-tower-zelda.yml).

### WF47 — Construct Head

Target: [construct-head-zelda](../entries/construct-head-zelda.yml). Current meaning: 魔像的头部；《王国之泪》中的左纳乌装置。

Action: Add the currently absent `word_formation` field.

```plain
Construct 在此是魔像这一名词，head 表示头部；不是动词 construct (建造)。
```

Evidence: [G09]. Local: [construct-head-zelda](../entries/construct-head-zelda.yml), [construct-being-zelda](../entries/construct-being-zelda.yml), [head-general](../entries/head-general.yml).

### WF48 — Device Dispenser

Target: [device-dispenser-zelda](../entries/device-dispenser-zelda.yml). Current meaning: 左纳乌装置制造机；《王国之泪》中的地点或设施。

Action: Add the currently absent `word_formation` field.

```plain
device (装置) + dispenser (发放装置)，帮助理解英文名称；游戏简体中文名为左纳乌装置制造机。
```

Evidence: [G10]. Local: [device-dispenser-zelda](../entries/device-dispenser-zelda.yml), [device-engineering](../entries/device-engineering.yml), [dispenser-engineering](../entries/dispenser-engineering.yml).

## Etymology drafts

Each block is the exact proposed `etymology` value. All ETY proposals were approved and implemented. The evidence is Douglas Harper’s Online Etymology Dictionary, using the headword origin paragraph rather than unrelated entries, advertisements, or speculative deeper roots. These Chinese summaries are original paraphrases. Qualifications such as “possibly,” and alternative transmission routes, are preserved in the card text.

Review batch: **ETY01–ETY18**. No fictional naming origin is inferred from ordinary-word history.

### ETY01 — stamina

Target: [stamina-general](../entries/stamina-general.yml). Current meaning: 耐力；持续活动的体力；游戏中对应精力。

```plain
来自拉丁语 stamina，是 stamen (线、经线) 的复数；英语后来发展出支撑持续活动的耐力这一含义。更具体的比喻演变存在解释空间，不把线的图像当作唯一确定路径。
```

Evidence: [E01]. Existing sense: [stamina-general](../entries/stamina-general.yml), [durability-equipment](../entries/durability-equipment.yml).

### ETY02 — integrity

Target: [integrity-general](../entries/integrity-general.yml). Current meaning: 正直；诚信；坚持诚实和道德原则。

```plain
与拉丁语 integritas (完整、健全) 和 integer (完整的) 有关；完整无损的含义也发展为品格健全、正直。这里是历史联系，不是把现代词拆成 in + tegrity。
```

Evidence: [E02]. Existing sense: [integrity-general](../entries/integrity-general.yml).

### ETY03 — resilient

Target: [resilient-combat](../entries/resilient-combat.yml). Current meaning: 有韧性的；遭受打击后能恢复或继续应对的。

```plain
来自拉丁语 resilire (弹回、跳回) 的分词形式。英语从回弹的物理含义发展出遭遇困难后恢复的引申义。
```

Evidence: [E03]. Existing sense: [resilient-combat](../entries/resilient-combat.yml).

### ETY04 — elixir

Target: [elixir-cooking](../entries/elixir-cooking.yml). Current meaning: 药剂

```plain
经中世纪拉丁语来自阿拉伯语 al-iksir，原与炼金术中的点金石有关；更早可能来自希腊语 xerion (使伤口干燥的粉末)。
```

Evidence: [E04]. Existing sense: [elixir-cooking](../entries/elixir-cooking.yml).

### ETY05 — recipe

Target: [recipe-cooking](../entries/recipe-cooking.yml). Current meaning: 食谱

```plain
来自拉丁语 recipe (取用吧)，是 recipere 的命令式，早先用于药方开头；英语后来也用它指准备食物的配方。
```

Evidence: [E05]. Existing sense: [recipe-cooking](../entries/recipe-cooking.yml).

### ETY06 — melee

Target: [melee-combat](../entries/melee-combat.yml). Current meaning: 近战

```plain
借自法语 mêlée，与表示混合的动词 mêler 有关；原本描写多人混在一起的混战，有助于理解近身交战这一用法。
```

Evidence: [E06]. Existing sense: [melee-combat](../entries/melee-combat.yml).

### ETY07 — hydrant

Target: [hydrant-engineering](../entries/hydrant-engineering.yml). Current meaning: 消防栓；连接供水管道的取水装置。

```plain
构词使用源于希腊语 hydor (水) 的 hydr- 和后缀 -ant；原指从供水主管取水的装置。这不是从英语 hydrate 删除字母得到的词。
```

Evidence: [E07]. Existing sense: [hydrant-engineering](../entries/hydrant-engineering.yml).

### ETY08 — aquatic

Target: [aquatic-combat](../entries/aquatic-combat.yml). Current meaning: 水生的

```plain
经法语 aquatique 来自拉丁语 aquaticus，源头是 aqua (水)；因此与水有关或生活在水中的意思相连。
```

Evidence: [E08]. Existing sense: [aquatic-combat](../entries/aquatic-combat.yml).

### ETY09 — vulnerable

Target: [vulnerable-combat](../entries/vulnerable-combat.yml). Current meaning: 易受攻击的

```plain
来自晚期拉丁语 vulnerabilis，与 vulnerare (使受伤) 及 vulnus (伤口) 有关；由可受伤理解易受攻击或伤害。
```

Evidence: [E09]. Existing sense: [vulnerable-combat](../entries/vulnerable-combat.yml).

### ETY10 — hostile

Target: [hostile-combat](../entries/hostile-combat.yml). Current meaning: 敌对的

```plain
经法语或直接来自拉丁语 hostilis (敌人的、敌对的)，与古典拉丁语 hostis (敌人) 有关；不按现代英语 host + -ile 来推义。
```

Evidence: [E10]. Existing sense: [hostile-combat](../entries/hostile-combat.yml), [aggressive-combat](../entries/aggressive-combat.yml).

### ETY11 — trident

Target: [trident-weapon](../entries/trident-weapon.yml). Current meaning: 三叉戟

```plain
来自拉丁语 tridens，含 tri- (三) 和 dens (牙齿)；三个齿状尖端形成三叉戟这一形象。
```

Evidence: [E11]. Existing sense: [trident-weapon](../entries/trident-weapon.yml).

### ETY12 — sacrifice

Target: [sacrifice-story](../entries/sacrifice-story.yml). Current meaning: 牺牲

```plain
经古法语来自拉丁语 sacrificium；历史构成与 sacer (神圣的) 和 facere (做、使成为) 有关，原与献祭相连，后来也表示为更重要目标而放弃某物。
```

Evidence: [E12]. Existing sense: [sacrifice-story](../entries/sacrifice-story.yml).

### ETY13 — malevolent

Target: [malevolent-general](../entries/malevolent-general.yml). Current meaning: 邪恶的、怀有恶意的

```plain
与拉丁语 male (坏地) 和 velle (希望) 的分词形式有关；核心形象是希望别人遭遇坏事，因此表示怀有恶意的。
```

Evidence: [E13]. Existing sense: [malevolent-general](../entries/malevolent-general.yml).

### ETY14 — courageous

Target: [courageous-general](../entries/courageous-general.yml). Current meaning: 勇敢的

```plain
来自古法语 corageus，与 corage (内心、精神、勇气) 有关，更早与拉丁语 cor (心) 相连；心与内在力量的历史联系有助于理解勇敢。
```

Evidence: [E14]. Existing sense: [courageous-general](../entries/courageous-general.yml).

### ETY15 — enthusiastic

Target: [enthusiastic-general](../entries/enthusiastic-general.yml). Current meaning: 热情的

```plain
来自希腊语中表示受神启发的词族；更早的 entheos 与 en (在……之内) 和 theos (神) 有关。现代热情、热切的含义不要求宗教背景。
```

Evidence: [E15]. Existing sense: [enthusiastic-general](../entries/enthusiastic-general.yml).

### ETY16 — sauté

Target: [saute-cooking](../entries/saute-cooking.yml). Current meaning: 用少量油快速翻炒或煎炒。

```plain
借自法语 sauté，是 sauter (跳) 的过去分词；词源中的跳与烹饪时颠动食物的动作有关。现代英语中可作快速煎炒这一动词使用。
```

Evidence: [E16]. Existing sense: [saute-cooking](../entries/saute-cooking.yml).

### ETY17 — paella

Target: [paella-cooking](../entries/paella-cooking.yml). Current meaning: 西班牙海鲜饭或烩饭；以米饭配合肉类、海鲜或蔬菜烹制的料理。

```plain
来自加泰罗尼亚语 paella (烹饪锅)，菜名取自烹制它的锅具；不要仅凭名称认定配料必须全是海鲜。
```

Evidence: [E17]. Existing sense: [paella-cooking](../entries/paella-cooking.yml).

### ETY18 — rupee

Target: [rupee-game](../entries/rupee-game.yml). Current meaning: 卢比；本作的货币单位。

```plain
经印地语或乌尔都语词形追溯到梵语 rupyah (加工过的银)。这是现实货币名称的词源，不证明游戏卢比的材质或世界观起源。
```

Evidence: [E18]. Existing sense: [rupee-game](../entries/rupee-game.yml).

## Mnemonic drafts

Each block is the exact proposed `mnemonic` value. All N proposals were explicitly approved along with the complete plan and implemented; pair approval alone would not authorize them. Images and simplified contrasts below are editorial learning devices, not claims about word origins, actual gameplay events, or rigid classifications. Sources support the underlying meanings; they do not author or validate the invented imagery. The core definitions and existing examples remain intact.

Review batch: **N01–N20**.

### N01 — stamina

Target: [stamina-general](../entries/stamina-general.yml). Current meaning: 耐力；持续活动的体力；游戏中对应精力。

```plain
用法助记: 想象角色跑步时逐渐消耗的体力，把 stamina 记在人身上；装备承受使用的耐久度则是 durability。
```

Lexical evidence: [stamina-general](../entries/stamina-general.yml), [durability-equipment](../entries/durability-equipment.yml). Local entry evidence; editorial construction/learning image. The image or contrast is an original proposal.

### N02 — resilient

Target: [resilient-combat](../entries/resilient-combat.yml). Current meaning: 有韧性的；遭受打击后能恢复或继续应对的。

```plain
图像助记 (非词源): 想象弹簧被压下后又弹回来，把 resilient 的重点放在受挫后恢复，而不只是硬或结实。
```

Lexical evidence: [resilient-combat](../entries/resilient-combat.yml). [E03] The image or contrast is an original proposal.

### N03 — sturdy

Target: [sturdy-general](../entries/sturdy-general.yml). Current meaning: 坚固的

```plain
图像助记 (非词源): 想象一张结构结实、不容易摇晃的桌子。sturdy 强调坚固；resilient 更容易让人想到受压或受挫后的恢复。
```

Lexical evidence: [sturdy-general](../entries/sturdy-general.yml), [resilient-combat](../entries/resilient-combat.yml). [L21], [E03] The image or contrast is an original proposal.

### N04 — vulnerable

Target: [vulnerable-combat](../entries/vulnerable-combat.yml). Current meaning: 易受攻击的

```plain
图像助记 (非词源): 想象护甲间露出的空隙，把 vulnerable 和容易受到伤害联系起来；这不是对任何具体怪物弱点的声明。
```

Lexical evidence: [vulnerable-combat](../entries/vulnerable-combat.yml). [E09] The image or contrast is an original proposal.

### N05 — susceptible

Target: [susceptible-combat](../entries/susceptible-combat.yml). Current meaning: 易受某种影响的

```plain
用法助记: 看到 susceptible to，就追问容易受什么影响；它突出对某种影响的易感性，不只局限于直接攻击。
```

Lexical evidence: [susceptible-combat](../entries/susceptible-combat.yml). [L20] The image or contrast is an original proposal.

### N06 — hostile

Target: [hostile-combat](../entries/hostile-combat.yml). Current meaning: 敌对的

```plain
辨析助记: hostile 先想敌对态度，aggressive 先想主动进攻或强势行为；两种特征可以同时存在。
```

Lexical evidence: [hostile-combat](../entries/hostile-combat.yml), [aggressive-combat](../entries/aggressive-combat.yml). [L22], [L23] The image or contrast is an original proposal.

### N07 — reckless

Target: [reckless-general](../entries/reckless-general.yml). Current meaning: 鲁莽的

```plain
辨析助记: 想象没有顾及后果就冲上去的人，记 reckless；面对危险的勇气用 courageous 或 brave，不能直接把鲁莽当作勇敢。
```

Lexical evidence: [reckless-general](../entries/reckless-general.yml), [courageous-general](../entries/courageous-general.yml), [brave-general](../entries/brave-general.yml). [L25] The image or contrast is an original proposal.

### N08 — heir

Target: [heir-general](../entries/heir-general.yml). Current meaning: 继承人

```plain
用法助记: 追问谁继承财产、位置或称号，想到 heir；追问谁是谁的后辈，想到 descendant。后裔不自动等于继承人。
```

Lexical evidence: [heir-general](../entries/heir-general.yml), [descendant-general](../entries/descendant-general.yml). [L24], [L09] The image or contrast is an original proposal.

### N09 — pristine

Target: [pristine-general](../entries/pristine-general.yml). Current meaning: 完好如初的；未受损或未被改变的。

```plain
图像助记 (非词源): 想象一件仍完好如初的物品，再对照 decayed 的衰败形象；pristine 不等于攻击力一定更高。
```

Lexical evidence: [pristine-general](../entries/pristine-general.yml), [decayed-general](../entries/decayed-general.yml). Local entry evidence; editorial construction/learning image. The image or contrast is an original proposal.

### N10 — summit

Target: [summit-terrain](../entries/summit-terrain.yml). Current meaning: 山顶

```plain
图像助记 (非词源): 用点、线、面作简化图像: summit 想山顶，ridge 想延伸的山脊，plateau 想较平展的高地；不是严格地貌分类。
```

Lexical evidence: [summit-terrain](../entries/summit-terrain.yml), [ridge-terrain](../entries/ridge-terrain.yml), [plateau-terrain](../entries/plateau-terrain.yml). Local entry evidence; editorial construction/learning image. The image or contrast is an original proposal.

### N11 — chasm

Target: [chasm-terrain](../entries/chasm-terrain.yml). Current meaning: 深渊；深而宽的裂隙。

```plain
图像助记 (非词源): chasm 想脚下裂开的深口，canyon 想沿两侧陡壁延伸的长谷；只帮助区分，不限定游戏地穴的全部形状。
```

Lexical evidence: [chasm-terrain](../entries/chasm-terrain.yml), [canyon-terrain](../entries/canyon-terrain.yml). [L33] The image or contrast is an original proposal.

### N12 — cavern

Target: [cavern-architecture](../entries/cavern-architecture.yml). Current meaning: 大型洞穴

```plain
图像助记 (非词源): 想象走进较大的洞内空间来记 cavern；cave 是更通用的洞穴词，两词并非互斥类别。
```

Lexical evidence: [cavern-architecture](../entries/cavern-architecture.yml), [cave-architecture](../entries/cave-architecture.yml). [L34] The image or contrast is an original proposal.

### N13 — simmer

Target: [simmer-cooking](../entries/simmer-cooking.yml). Current meaning: 小火慢煮；使液体保持轻微沸腾。

```plain
图像助记 (非词源): 想象锅中只有细小气泡、轻微翻动，以此记住 simmer 的小火慢煮。
```

Lexical evidence: [simmer-cooking](../entries/simmer-cooking.yml). [L35] The image or contrast is an original proposal.

### N14 — sear

Target: [sear-cooking](../entries/sear-cooking.yml). Current meaning: 以高温迅速煎烤食物表面。

```plain
图像助记 (非词源): 先想象食物表面在高温下迅速上色，再记 sear 的表面处理重点。
```

Lexical evidence: [sear-cooking](../entries/sear-cooking.yml). [L36] The image or contrast is an original proposal.

### N15 — steam

Target: [steam-cooking](../entries/steam-cooking.yml). Current meaning: 蒸；利用蒸汽加热食物。

```plain
图像助记 (非词源): 想象锅盖下的热蒸汽包围食物，把 steam 与蒸汽加热联系起来。
```

Lexical evidence: [steam-cooking](../entries/steam-cooking.yml). Local entry evidence; editorial construction/learning image. The image or contrast is an original proposal.

### N16 — grill

Target: [grill-cooking](../entries/grill-cooking.yml). Current meaning: 炙烤；在烤架上烹饪。

```plain
图像助记 (非词源): 想象食物放在烤架上受热，用烤架帮助记住 grill，而不是把所有加热方式都叫 grill。
```

Lexical evidence: [grill-cooking](../entries/grill-cooking.yml). [L37] The image or contrast is an original proposal.

### N17 — bass

Target: [bass-food](../entries/bass-food.yml). Current meaning: 鲈鱼。

```plain
辨义助记: 跟着 Hyrule Bass 想鱼，把这里的 bass 记成鲈鱼；音乐中的低音是另一义项，同样拼写也不保证同音。
```

Lexical evidence: [bass-food](../entries/bass-food.yml), [hyrule-bass-zelda](../entries/hyrule-bass-zelda.yml). [L26] The image or contrast is an original proposal.

### N18 — tear

Target: [tear-general](../entries/tear-general.yml). Current meaning: 眼泪；泪水。

```plain
辨义助记: 跟着 Tears of the Kingdom 想眼睛里的泪水；不要把这里的 tear 自动套成撕裂这一另义。
```

Lexical evidence: [tear-general](../entries/tear-general.yml). [L27] The image or contrast is an original proposal.

### N19 — inventory

Target: [inventory-ui](../entries/inventory-ui.yml). Current meaning: 物品清单；库存；游戏中的物品栏。

```plain
图像助记 (非词源): inventory 先想物品清单，pouch 先想装东西的袋子；游戏界面可能借袋子的名称表示物品栏。
```

Lexical evidence: [inventory-ui](../entries/inventory-ui.yml), [pouch-ui](../entries/pouch-ui.yml). Local entry evidence; editorial construction/learning image. The image or contrast is an original proposal.

### N20 — spring

Target: [spring-water](../entries/spring-water.yml). Current meaning: 泉；泉水；地下水自然涌出之处。

```plain
图像助记 (非词源): 一幅图是地下涌出的泉水，另一幅图是受压后回弹的弹簧；用两幅图分开记同形的 spring。
```

Lexical evidence: [spring-water](../entries/spring-water.yml), [spring-mechanical](../entries/spring-mechanical.yml). Local entry evidence; editorial construction/learning image. The image or contrast is an original proposal.

## Additional related pairs

**R10–R29 were approved and implemented as 20 additional pairs.** R01–R09 refer to the already approved import batch and are not being resubmitted. Current words and meanings below are read from the collection. The rationale and risk are review notes, not text automatically written into Mnemonic or Notes.

If a pair is approved, declare the second ID once in the first ID’s `related_entries`, preserving any existing list members. The exporter supplies the reverse display from current Word and Meaning; no custom RelatedWords text is written. The declaration table specifies the exact ID addition. Approval of a pair does not approve any N draft, prerequisite, or other edit.

| Proposal | Endpoint A: stable ID / Word / current Meaning                                           | Endpoint B: stable ID / Word / current Meaning                                                           | Learning value                                                 | Confusion risk                                              |
| :------- | :--------------------------------------------------------------------------------------- | :------------------------------------------------------------------------------------------------------- | :------------------------------------------------------------- | :---------------------------------------------------------- |
| R10      | [courageous-general](../entries/courageous-general.yml); courageous; 勇敢的              | [brave-general](../entries/brave-general.yml); brave; 勇敢的                                             | 中文均为勇敢，可连接近义表达。                                 | 相近不代表所有搭配都能互换；不额外推断人物性格。            |
| R11      | [brave-general](../entries/brave-general.yml); brave; 勇敢的                             | [fearless-general](../entries/fearless-general.yml); fearless; 无畏的                                    | 区分面对危险的勇气与突出无畏的表达。                           | brave 不必表示没有恐惧；fearless 也可能是修辞强调。         |
| R12      | [courageous-general](../entries/courageous-general.yml); courageous; 勇敢的              | [reckless-general](../entries/reckless-general.yml); reckless; 鲁莽的                                    | 区分勇敢与不顾后果的鲁莽。                                     | 不将二者当同义词，也不把一次冒险自动判断为鲁莽。            |
| R13      | [hostile-combat](../entries/hostile-combat.yml); hostile; 敌对的                         | [aggressive-combat](../entries/aggressive-combat.yml); aggressive; 攻击性强的                            | 区分敌对态度与主动进攻、强势行为。                             | 可同时成立；不是互斥分类，aggressive 在别的语境也可偏中性。 |
| R14      | [vulnerable-combat](../entries/vulnerable-combat.yml); vulnerable; 易受攻击的            | [susceptible-combat](../entries/susceptible-combat.yml); susceptible; 易受某种影响的                     | 连接易受伤害与易受特定影响两个相近概念。                       | 语义有重叠；不要固定成 vulnerable 只能接攻击。              |
| R15      | [sturdy-general](../entries/sturdy-general.yml); sturdy; 坚固的                          | [resilient-combat](../entries/resilient-combat.yml); resilient; 有韧性的；遭受打击后能恢复或继续应对的。 | 对照结构坚固与受压或受挫后恢复。                               | 结实不等于有弹性；resilient 也可用于人而非材料。            |
| R16      | [heir-general](../entries/heir-general.yml); heir; 继承人                                | [descendant-general](../entries/descendant-general.yml); descendant; 后裔                                | 对照继承位置与血缘后裔。                                       | 不是所有后裔都是继承人；heir 也可指非血缘的接班者。         |
| R17      | [summit-terrain](../entries/summit-terrain.yml); summit; 山顶                            | [ridge-terrain](../entries/ridge-terrain.yml); ridge; 山脊                                               | 区分山顶与山脊，建立地形图像。                                 | 点与线只是助记；实际地形不能只用几何形状定义。              |
| R18      | [cave-architecture](../entries/cave-architecture.yml); cave; 洞穴                        | [cavern-architecture](../entries/cavern-architecture.yml); cavern; 大型洞穴                              | 连接通用洞穴词和常指较大洞穴的词。                             | 不是两个互斥的洞穴等级，也不能据此推断游戏洞窟大小。        |
| R19      | [chasm-terrain](../entries/chasm-terrain.yml); chasm; 深渊；深而宽的裂隙。               | [canyon-terrain](../entries/canyon-terrain.yml); canyon; 峡谷                                            | 区分深裂口或深穴与延伸的峡谷。                                 | 实际地形可能重叠；不把一般词义硬套成 UI 分类。              |
| R20      | [grill-cooking](../entries/grill-cooking.yml); grill; 炙烤；在烤架上烹饪。               | [sear-cooking](../entries/sear-cooking.yml); sear; 以高温迅速煎烤食物表面。                              | 对照烤架烹饪与高温快速处理表面。                               | 同一食物可能先后涉及两种处理，不是互斥方法。                |
| R21      | [simmer-cooking](../entries/simmer-cooking.yml); simmer; 小火慢煮；使液体保持轻微沸腾。  | [stew-cooking](../entries/stew-cooking.yml); stew; 炖菜                                                  | 连接小火慢煮方法与炖菜。                                       | 本库 simmer 是动词，stew 是名词；不能直接互换词性。         |
| R22      | [inventory-ui](../entries/inventory-ui.yml); inventory; 物品清单；库存；游戏中的物品栏。 | [pouch-ui](../entries/pouch-ui.yml); pouch; 小袋；随身袋。                                               | 区分物品清单或物品栏与随身袋的形象。                           | 界面隐喻有关联，不表示两个英文词在所有场景等义。            |
| R23      | [armor-equipment](../entries/armor-equipment.yml); armor; 护甲                           | [armory-architecture](../entries/armory-architecture.yml); armory; 武器库                                | 辨析装备与存放武器的设施，防止近似拼写混淆。                   | armory 不是护甲的复数，也不意味着只能存放护甲。             |
| R24      | [helm-clothing](../entries/helm-clothing.yml); helm; 头盔                                | [helmet-clothing](../entries/helmet-clothing.yml); helmet; 头盔                                          | 连接两个表示头盔的名词。                                       | helm 另有舵等义项，本次仅关联头盔义。                       |
| R25      | [headband-clothing](../entries/headband-clothing.yml); headband; 头带                    | [headdress-clothing](../entries/headdress-clothing.yml); headdress; 头饰                                 | 区分具体头带与较宽泛的头部装束或装饰。                         | 不能把所有 headdress 都翻成头带。                           |
| R26      | [ingredient-cooking](../entries/ingredient-cooking.yml); ingredient; 食材、原料          | [recipe-cooking](../entries/recipe-cooking.yml); recipe; 食谱                                            | 连接材料与组织材料、步骤的配方。                               | 食材不等于食谱；关联不创造配方依赖或烹饪规则。              |
| R27      | [endure-general](../entries/endure-general.yml); endure; 持久；忍受；持续存在。          | [energize-general](../entries/energize-general.yml); energize; 使充满活力；补充能量。                    | 连接两种料理前缀背后的普通动词，便于回看持续与补充活力的差别。 | 保留各自 Context 的游戏说明，不把一般英语词义当作数值机制。 |
| R28      | [emitter-engineering](../entries/emitter-engineering.yml); emitter; 发射器、发出装置     | [emit-tech](../entries/emit-tech.yml); emit; 生成；输出                                                  | 跨工程与软件语境连接发出装置和输出动作。                       | 现有 emit-tech 学习生成、输出；不把编译输出理解为物理射线。 |
| R29      | [surface-terrain](../entries/surface-terrain.yml); surface; 地表                         | [the-depths-zelda](../entries/the-depths-zelda.yml); The Depths; 地底；本作位于地表下方的广大区域。      | 对照本作地表与地底区域。                                       | The Depths 是本作区域名，不能把所有 depths 都译作该区域。   |

| Proposal | Declare on                                                | Add exactly this ID   |
| :------- | :-------------------------------------------------------- | :-------------------- |
| R10      | [courageous-general](../entries/courageous-general.yml)   | `brave-general`       |
| R11      | [brave-general](../entries/brave-general.yml)             | `fearless-general`    |
| R12      | [courageous-general](../entries/courageous-general.yml)   | `reckless-general`    |
| R13      | [hostile-combat](../entries/hostile-combat.yml)           | `aggressive-combat`   |
| R14      | [vulnerable-combat](../entries/vulnerable-combat.yml)     | `susceptible-combat`  |
| R15      | [sturdy-general](../entries/sturdy-general.yml)           | `resilient-combat`    |
| R16      | [heir-general](../entries/heir-general.yml)               | `descendant-general`  |
| R17      | [summit-terrain](../entries/summit-terrain.yml)           | `ridge-terrain`       |
| R18      | [cave-architecture](../entries/cave-architecture.yml)     | `cavern-architecture` |
| R19      | [chasm-terrain](../entries/chasm-terrain.yml)             | `canyon-terrain`      |
| R20      | [grill-cooking](../entries/grill-cooking.yml)             | `sear-cooking`        |
| R21      | [simmer-cooking](../entries/simmer-cooking.yml)           | `stew-cooking`        |
| R22      | [inventory-ui](../entries/inventory-ui.yml)               | `pouch-ui`            |
| R23      | [armor-equipment](../entries/armor-equipment.yml)         | `armory-architecture` |
| R24      | [helm-clothing](../entries/helm-clothing.yml)             | `helmet-clothing`     |
| R25      | [headband-clothing](../entries/headband-clothing.yml)     | `headdress-clothing`  |
| R26      | [ingredient-cooking](../entries/ingredient-cooking.yml)   | `recipe-cooking`      |
| R27      | [endure-general](../entries/endure-general.yml)           | `energize-general`    |
| R28      | [emitter-engineering](../entries/emitter-engineering.yml) | `emit-tech`           |
| R29      | [surface-terrain](../entries/surface-terrain.yml)         | `the-depths-zelda`    |

Both endpoints’ linked entries are local evidence. Additional dictionary checks support the vulnerable/susceptible distinction ([L20]), hostile/aggressive ([L22], [L23]), sturdy/resilient ([L21], [E03]), heir/descendant ([L24], [L09]), reckless ([L25]), cavern ([L34]), canyon ([L33]), simmer/sear/grill ([L35], [L36], [L37]), and helm/helmet ([L38], [L39]). R27’s cooking-prefix context is the existing Context on endure-general and energize-general; this plan adds no new game-effect parameter claims. R28 deliberately bridges the existing software sense of emit-tech to a physical device noun, without treating the senses as identical.

## Deferred naming questions

These decisions preserve useful words while avoiding unsupported enrichment. They are not blockers for independently approved drafts above.

| Targets                                                                                                                                                                                                                                                       | Decision and boundary                                                                                                                                                                                |
| :------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | :--------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| [sundelion-zelda](../entries/sundelion-zelda.yml), [puffshroom-zelda](../entries/puffshroom-zelda.yml), [brightcap-zelda](../entries/brightcap-zelda.yml), [skyshroom-zelda](../entries/skyshroom-zelda.yml), [evermean-zelda](../entries/evermean-zelda.yml) | Leave historical naming fields empty. Possible English wordplay requires evidence or an explicitly qualified future proposal; resemblance alone does not establish a developer’s intended blend.     |
| [draconification-zelda](../entries/draconification-zelda.yml)                                                                                                                                                                                                 | Keep the verified game spelling and current plot explanation. No invented Latin history or official coinage account is proposed.                                                                     |
| [link-zelda](../entries/link-zelda.yml), [ganondorf-zelda](../entries/ganondorf-zelda.yml), [zonai-zelda](../entries/zonai-zelda.yml), [gerudo-zelda](../entries/gerudo-zelda.yml), [bokoblin-zelda](../entries/bokoblin-zelda.yml)                           | Retain names, localization, and provenance. Do not derive identities, race histories, or naming origins from similar English strings.                                                                |
| [waistguard-clothing](../entries/waistguard-clothing.yml)                                                                                                                                                                                                     | WF17 is an openly labeled constituent analysis from waist and guard plus the existing target. A whole-word dictionary page was unavailable; no dictionary-attested historical derivation is claimed. |
| Other unselected ordinary words                                                                                                                                                                                                                               | Keep empty fields unless a later concrete explanation adds value; do not add arbitrary imagery or a list of near-synonyms merely to raise coverage.                                                  |

## Source register

References were inspected during this planning task on 2026-11-04. Dictionary claims use editor-written definitions or origin sections; automatically collected example sentences are not treated as authoritative. Cambridge pages returned HTTP 403 and are excluded from the evidence. Unavailable whole-word or guessed URLs are not listed as successfully checked sources.

Pinned game-name references G01–G10 were checked against the retained bilingual 1.2.1 snapshot at commit `1c66ac1a02edf2d73d10b0dc6f0f2de29da7b119`. That collection is a third-party transcription, not a Nintendo-published etymological source. Its name keys establish wording and localization, not naming intent. No network accessibility claim is made for those GitHub URLs in this task; their retained contents were inspected locally.

For approved WF, ETY, and N text proposals only, preserve the complete existing Notes text except the explicitly approved WF01 paragraph move. Append one plain-text source line for each distinct cited external URL not already present, in source-code order, using the following exact format with the title and URL from this register:

```plain
Memory-aid source: <source title>. <source URL>
```

The bracketed source codes are review-document references, not new YAML fields and not literal card content. Do not put Markdown links or source codes into the card text. When several approved proposals for one entry cite the same URL, append its source line once. Local-entry comparisons do not authorize new relationships or copies of whole entries into Notes.

| Source | Title                                                                      | What it supports                                                                                                                                                                                                                                   |
| :----- | :------------------------------------------------------------------------- | :------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| [P01]  | Merriam-Webster.com Dictionary — un-                                       | Negative and reversing prefix senses; not a whole-word history.                                                                                                                                                                                    |
| [P02]  | Merriam-Webster.com Dictionary — re-                                       | Again and back senses.                                                                                                                                                                                                                             |
| [P03]  | Merriam-Webster.com Dictionary — -less                                     | Lacking a quality; not a universal literal rule.                                                                                                                                                                                                   |
| [P04]  | Merriam-Webster.com Dictionary — -er                                       | Agent and instrument noun suffix; not the comparative suffix.                                                                                                                                                                                      |
| [P05]  | Merriam-Webster.com Dictionary — -ment                                     | Action, result, state, or place nouns.                                                                                                                                                                                                             |
| [P06]  | Merriam-Webster.com Dictionary — -ful                                      | Adjectives characterized by a quality.                                                                                                                                                                                                             |
| [P07]  | Merriam-Webster.com Dictionary — -dom                                      | Realm/jurisdiction sense, including kingdom.                                                                                                                                                                                                       |
| [P08]  | Merriam-Webster.com Dictionary — ultra-                                    | Beyond ordinary limits; does not establish Nintendo naming intent.                                                                                                                                                                                 |
| [P09]  | Merriam-Webster.com Dictionary — auto-                                     | Automatic/self-acting sense.                                                                                                                                                                                                                       |
| [P10]  | Merriam-Webster.com Dictionary — -ed                                       | Having/characterized-by adjective sense.                                                                                                                                                                                                           |
| [P11]  | Merriam-Webster.com Dictionary — -al                                       | Adjective suffix meaning relating to.                                                                                                                                                                                                              |
| [P12]  | Merriam-Webster.com Dictionary — -ous                                      | Adjective suffix meaning having a quality.                                                                                                                                                                                                         |
| [L01]  | Merriam-Webster.com Dictionary — headband                                  | Selected dictionary sense or explicitly stated construction; modern analysis is identified separately from historical origin.                                                                                                                      |
| [L02]  | Merriam-Webster.com Dictionary — headdress                                 | Selected dictionary sense or explicitly stated construction; modern analysis is identified separately from historical origin.                                                                                                                      |
| [L03]  | Merriam-Webster.com Dictionary — foothill                                  | Selected dictionary sense or explicitly stated construction; modern analysis is identified separately from historical origin.                                                                                                                      |
| [L04]  | Online Etymology Dictionary — airborne                                     | Selected dictionary sense or explicitly stated construction; modern analysis is identified separately from historical origin.                                                                                                                      |
| [L05]  | Online Etymology Dictionary — refinery                                     | Selected dictionary sense or explicitly stated construction; modern analysis is identified separately from historical origin.                                                                                                                      |
| [L06]  | Online Etymology Dictionary — stabilizer                                   | Selected dictionary sense or explicitly stated construction; modern analysis is identified separately from historical origin.                                                                                                                      |
| [L07]  | Merriam-Webster.com Dictionary — dispenser                                 | Selected dictionary sense or explicitly stated construction; modern analysis is identified separately from historical origin.                                                                                                                      |
| [L08]  | Merriam-Webster.com Dictionary — home (verb; homed, homing)                | Selected dictionary sense or explicitly stated construction; modern analysis is identified separately from historical origin.                                                                                                                      |
| [L09]  | Merriam-Webster.com Dictionary — descendant                                | Selected dictionary sense or explicitly stated construction; modern analysis is identified separately from historical origin.                                                                                                                      |
| [L10]  | Online Etymology Dictionary — mysterious                                   | Selected dictionary sense or explicitly stated construction; modern analysis is identified separately from historical origin.                                                                                                                      |
| [L11]  | Online Etymology Dictionary — waterfall                                    | Selected dictionary sense or explicitly stated construction; modern analysis is identified separately from historical origin.                                                                                                                      |
| [L12]  | Online Etymology Dictionary — staircase                                    | Selected dictionary sense or explicitly stated construction; modern analysis is identified separately from historical origin.                                                                                                                      |
| [L13]  | Online Etymology Dictionary — highland                                     | Selected dictionary sense or explicitly stated construction; modern analysis is identified separately from historical origin.                                                                                                                      |
| [L14]  | Online Etymology Dictionary — wetland                                      | Selected dictionary sense or explicitly stated construction; modern analysis is identified separately from historical origin.                                                                                                                      |
| [L15]  | Online Etymology Dictionary — ruthless                                     | Selected dictionary sense or explicitly stated construction; modern analysis is identified separately from historical origin.                                                                                                                      |
| [L16]  | Online Etymology Dictionary — enhancement                                  | Selected dictionary sense or explicitly stated construction; modern analysis is identified separately from historical origin.                                                                                                                      |
| [L17]  | Online Etymology Dictionary — settlement                                   | Selected dictionary sense or explicitly stated construction; modern analysis is identified separately from historical origin.                                                                                                                      |
| [L18]  | Online Etymology Dictionary — undead                                       | Selected dictionary sense or explicitly stated construction; modern analysis is identified separately from historical origin.                                                                                                                      |
| [L19]  | Online Etymology Dictionary — underground                                  | Selected dictionary sense or explicitly stated construction; modern analysis is identified separately from historical origin.                                                                                                                      |
| [L20]  | Merriam-Webster.com Dictionary — susceptible                               | Selected dictionary sense or explicitly stated construction; modern analysis is identified separately from historical origin.                                                                                                                      |
| [L21]  | Merriam-Webster.com Dictionary — sturdy                                    | Selected dictionary sense or explicitly stated construction; modern analysis is identified separately from historical origin.                                                                                                                      |
| [L22]  | Merriam-Webster.com Dictionary — hostile                                   | Selected dictionary sense or explicitly stated construction; modern analysis is identified separately from historical origin.                                                                                                                      |
| [L23]  | Merriam-Webster.com Dictionary — aggressive                                | Selected dictionary sense or explicitly stated construction; modern analysis is identified separately from historical origin.                                                                                                                      |
| [L24]  | Merriam-Webster.com Dictionary — heir                                      | Selected dictionary sense or explicitly stated construction; modern analysis is identified separately from historical origin.                                                                                                                      |
| [L25]  | Merriam-Webster.com Dictionary — reckless                                  | Selected dictionary sense or explicitly stated construction; modern analysis is identified separately from historical origin.                                                                                                                      |
| [L26]  | Merriam-Webster.com Dictionary — bass                                      | Selected dictionary sense or explicitly stated construction; modern analysis is identified separately from historical origin.                                                                                                                      |
| [L27]  | Merriam-Webster.com Dictionary — tear                                      | Selected dictionary sense or explicitly stated construction; modern analysis is identified separately from historical origin.                                                                                                                      |
| [L28]  | Merriam-Webster.com Dictionary — waist                                     | Selected dictionary sense or explicitly stated construction; modern analysis is identified separately from historical origin.                                                                                                                      |
| [L29]  | Merriam-Webster.com Dictionary — guard                                     | Selected dictionary sense or explicitly stated construction; modern analysis is identified separately from historical origin.                                                                                                                      |
| [L30]  | Merriam-Webster.com Dictionary — fearless                                  | Selected dictionary sense or explicitly stated construction; modern analysis is identified separately from historical origin.                                                                                                                      |
| [L31]  | Merriam-Webster.com Dictionary — selfless                                  | Selected dictionary sense or explicitly stated construction; modern analysis is identified separately from historical origin.                                                                                                                      |
| [L32]  | Online Etymology Dictionary — recharge                                     | Selected dictionary sense or explicitly stated construction; modern analysis is identified separately from historical origin.                                                                                                                      |
| [G01]  | TOTK paired game text 1.2.1 — Obj_AutoBuilder_Name                         | Existing pinned bilingual text key: Obj_AutoBuilder_Name. Evidence for the displayed English name and localized Chinese label; not evidence for naming history. Name checked in retained snapshot: Autobuild / 蓝图.                               |
| [G02]  | TOTK paired game text 1.2.1 — Obj_UltraHand_Name                           | Existing pinned bilingual text key: Obj_UltraHand_Name. Evidence for the displayed English name and localized Chinese label; not evidence for naming history. Name checked in retained snapshot: Ultrahand / 究极手.                               |
| [G03]  | TOTK paired game text 1.2.1 — SpObj_FloatingStone_Capsule_A_01_Name        | Existing pinned bilingual text key: SpObj_FloatingStone_Capsule_A_01_Name. Evidence for the displayed English name and localized Chinese label; not evidence for naming history. Name checked in retained snapshot: Hover Stone / 浮游石.          |
| [G04]  | TOTK paired game text 1.2.1 — SpObj_FlameThrower_Capsule_A_01_Name         | Existing pinned bilingual text key: SpObj_FlameThrower_Capsule_A_01_Name. Evidence for the displayed English name and localized Chinese label; not evidence for naming history. Name checked in retained snapshot: Flame Emitter / 火龙头.         |
| [G05]  | TOTK paired game text 1.2.1 — SpObj_SnowMachine_Capsule_A_01_Name          | Existing pinned bilingual text key: SpObj_SnowMachine_Capsule_A_01_Name. Evidence for the displayed English name and localized Chinese label; not evidence for naming history. Name checked in retained snapshot: Frost Emitter / 冰龙头.          |
| [G06]  | TOTK paired game text 1.2.1 — SpObj_ElectricBoxGenerator_Capsule_A_01_Name | Existing pinned bilingual text key: SpObj_ElectricBoxGenerator_Capsule_A_01_Name. Evidence for the displayed English name and localized Chinese label; not evidence for naming history. Name checked in retained snapshot: Shock Emitter / 雷龙头. |
| [G07]  | TOTK paired game text 1.2.1 — SpObj_Beamos_Capsule_A_01_Name               | Existing pinned bilingual text key: SpObj_Beamos_Capsule_A_01_Name. Evidence for the displayed English name and localized Chinese label; not evidence for naming history. Name checked in retained snapshot: Beam Emitter / 光线头.                |
| [G08]  | TOTK paired game text 1.2.1 — P_Cannon                                     | Existing pinned bilingual text key: P_Cannon. Evidence for the displayed English name and localized Chinese label; not evidence for naming history. Name checked in retained snapshot: Skyview Tower / 鸟望台.                                     |
| [G09]  | TOTK paired game text 1.2.1 — SpObj_GolemHead_Capsule_A_01_Name            | Existing pinned bilingual text key: SpObj_GolemHead_Capsule_A_01_Name. Evidence for the displayed English name and localized Chinese label; not evidence for naming history. Name checked in retained snapshot: Construct Head / 魔像的头部.       |
| [G10]  | TOTK paired game text 1.2.1 — RecycleBox                                   | Existing pinned bilingual text key: RecycleBox. Evidence for the displayed English name and localized Chinese label; not evidence for naming history. Name checked in retained snapshot: Device Dispenser / 左纳乌装置制造机.                      |
| [E01]  | Online Etymology Dictionary — stamina                                      | Headword origin paragraph; draft omits unsupported naming intent, exact dates, and speculative deeper roots.                                                                                                                                       |
| [E02]  | Online Etymology Dictionary — integrity                                    | Headword origin paragraph; draft omits unsupported naming intent, exact dates, and speculative deeper roots.                                                                                                                                       |
| [E03]  | Online Etymology Dictionary — resilient                                    | Headword origin paragraph; draft omits unsupported naming intent, exact dates, and speculative deeper roots.                                                                                                                                       |
| [E04]  | Online Etymology Dictionary — elixir                                       | Headword origin paragraph; draft omits unsupported naming intent, exact dates, and speculative deeper roots.                                                                                                                                       |
| [E05]  | Online Etymology Dictionary — recipe                                       | Headword origin paragraph; draft omits unsupported naming intent, exact dates, and speculative deeper roots.                                                                                                                                       |
| [E06]  | Online Etymology Dictionary — melee                                        | Headword origin paragraph; draft omits unsupported naming intent, exact dates, and speculative deeper roots.                                                                                                                                       |
| [E07]  | Online Etymology Dictionary — hydrant                                      | Headword origin paragraph; draft omits unsupported naming intent, exact dates, and speculative deeper roots.                                                                                                                                       |
| [E08]  | Online Etymology Dictionary — aquatic                                      | Headword origin paragraph; draft omits unsupported naming intent, exact dates, and speculative deeper roots.                                                                                                                                       |
| [E09]  | Online Etymology Dictionary — vulnerable                                   | Headword origin paragraph; draft omits unsupported naming intent, exact dates, and speculative deeper roots.                                                                                                                                       |
| [E10]  | Online Etymology Dictionary — hostile                                      | Headword origin paragraph; draft omits unsupported naming intent, exact dates, and speculative deeper roots.                                                                                                                                       |
| [E11]  | Online Etymology Dictionary — trident                                      | Headword origin paragraph; draft omits unsupported naming intent, exact dates, and speculative deeper roots.                                                                                                                                       |
| [E12]  | Online Etymology Dictionary — sacrifice                                    | Headword origin paragraph; draft omits unsupported naming intent, exact dates, and speculative deeper roots.                                                                                                                                       |
| [E13]  | Online Etymology Dictionary — malevolent                                   | Headword origin paragraph; draft omits unsupported naming intent, exact dates, and speculative deeper roots.                                                                                                                                       |
| [E14]  | Online Etymology Dictionary — courageous                                   | Headword origin paragraph; draft omits unsupported naming intent, exact dates, and speculative deeper roots.                                                                                                                                       |
| [E15]  | Online Etymology Dictionary — enthusiastic                                 | Headword origin paragraph; draft omits unsupported naming intent, exact dates, and speculative deeper roots.                                                                                                                                       |
| [E16]  | Online Etymology Dictionary — saute                                        | Headword origin paragraph; draft omits unsupported naming intent, exact dates, and speculative deeper roots.                                                                                                                                       |
| [E17]  | Online Etymology Dictionary — paella                                       | Headword origin paragraph; draft omits unsupported naming intent, exact dates, and speculative deeper roots.                                                                                                                                       |
| [E18]  | Online Etymology Dictionary — rupee                                        | Headword origin paragraph; draft omits unsupported naming intent, exact dates, and speculative deeper roots.                                                                                                                                       |
| [L33]  | Merriam-Webster.com Dictionary — canyon                                    | Editor-written definition for the selected sense; automatically collected example sentences are not evidence.                                                                                                                                      |
| [L34]  | Merriam-Webster.com Dictionary — cavern                                    | Editor-written definition for the selected sense; automatically collected example sentences are not evidence.                                                                                                                                      |
| [L35]  | Merriam-Webster.com Dictionary — simmer                                    | Editor-written definition for the selected sense; automatically collected example sentences are not evidence.                                                                                                                                      |
| [L36]  | Merriam-Webster.com Dictionary — sear                                      | Editor-written definition for the selected sense; automatically collected example sentences are not evidence.                                                                                                                                      |
| [L37]  | Merriam-Webster.com Dictionary — grill                                     | Editor-written definition for the selected sense; automatically collected example sentences are not evidence.                                                                                                                                      |
| [L38]  | Merriam-Webster.com Dictionary — helm                                      | Editor-written definition for the selected sense; automatically collected example sentences are not evidence.                                                                                                                                      |
| [L39]  | Merriam-Webster.com Dictionary — helmet                                    | Editor-written definition for the selected sense; automatically collected example sentences are not evidence.                                                                                                                                      |

## Complete review ledger

Every reviewed batch ID appears exactly once below. A missing local related_entries field does not imply an absent Related Words display: reverse associations are derived during export. The ledger preserves the planning decisions and baseline source-field inventory. Listed proposal codes are now approved and applied, including derived reverse displays on endpoints with no YAML edit. Existing aids are retained. “Leave optional fields empty” is a scope decision after screening, not a claim that no useful etymology exists. This ledger is a suitability audit; the source register identifies the narrower set actually researched in depth.

| Stable ID                                                                               | Word                         | Approved proposals    | Existing source fields        | Review decision                                                                                       |
| :-------------------------------------------------------------------------------------- | :--------------------------- | :-------------------- | :---------------------------- | :---------------------------------------------------------------------------------------------------- |
| [activate-general](../entries/activate-general.yml)                                     | activate                     | None                  | No locally declared aid field | Leave optional fields empty or retain current aids; no additional explanation selected for this pass. |
| [adventure-general](../entries/adventure-general.yml)                                   | adventure                    | None                  | No locally declared aid field | Leave optional fields empty or retain current aids; no additional explanation selected for this pass. |
| [adventure-log-zelda](../entries/adventure-log-zelda.yml)                               | Adventure Log                | None                  | No locally declared aid field | Keep name/localization and Notes; no verified naming explanation selected.                            |
| [aerocuda-zelda](../entries/aerocuda-zelda.yml)                                         | Aerocuda                     | None                  | No locally declared aid field | Keep name/localization and Notes; no verified naming explanation selected.                            |
| [aggressive-combat](../entries/aggressive-combat.yml)                                   | aggressive                   | R13                   | No locally declared aid field | Applied listed enrichment or derived display; retained all other fields.                              |
| [aim-general](../entries/aim-general.yml)                                               | aim                          | None                  | No locally declared aid field | Leave optional fields empty or retain current aids; no additional explanation selected for this pass. |
| [airborne-combat](../entries/airborne-combat.yml)                                       | airborne                     | WF23                  | No locally declared aid field | Applied listed enrichment or derived display; retained all other fields.                              |
| [album-ui](../entries/album-ui.yml)                                                     | album                        | None                  | No locally declared aid field | Leave optional fields empty or retain current aids; no additional explanation selected for this pass. |
| [altar-architecture](../entries/altar-architecture.yml)                                 | altar                        | None                  | No locally declared aid field | Leave optional fields empty or retain current aids; no additional explanation selected for this pass. |
| [ambitious-general](../entries/ambitious-general.yml)                                   | ambitious                    | None                  | No locally declared aid field | Leave optional fields empty or retain current aids; no additional explanation selected for this pass. |
| [ancestor-general](../entries/ancestor-general.yml)                                     | ancestor                     | None                  | existing declared association | Leave optional fields empty or retain current aids; no additional explanation selected for this pass. |
| [ancient-arowana-zelda](../entries/ancient-arowana-zelda.yml)                           | Ancient Arowana              | None                  | No locally declared aid field | Keep name/localization and Notes; no verified naming explanation selected.                            |
| [ancient-general](../entries/ancient-general.yml)                                       | ancient                      | None                  | No locally declared aid field | Leave optional fields empty or retain current aids; no additional explanation selected for this pass. |
| [apple-food](../entries/apple-food.yml)                                                 | apple                        | None                  | No locally declared aid field | Leave optional fields empty or retain current aids; no additional explanation selected for this pass. |
| [aquatic-combat](../entries/aquatic-combat.yml)                                         | aquatic                      | ETY08                 | No locally declared aid field | Applied listed enrichment or derived display; retained all other fields.                              |
| [archipelago-terrain](../entries/archipelago-terrain.yml)                               | archipelago                  | None                  | No locally declared aid field | Leave optional fields empty or retain current aids; no additional explanation selected for this pass. |
| [armor-equipment](../entries/armor-equipment.yml)                                       | armor                        | R23                   | No locally declared aid field | Applied listed enrichment or derived display; retained all other fields.                              |
| [armored-combat](../entries/armored-combat.yml)                                         | armored                      | WF13                  | No locally declared aid field | Applied listed enrichment or derived display; retained all other fields.                              |
| [armory-architecture](../entries/armory-architecture.yml)                               | armory                       | R23                   | No locally declared aid field | Applied listed enrichment or derived display; retained all other fields.                              |
| [arowana-general](../entries/arowana-general.yml)                                       | arowana                      | None                  | No locally declared aid field | Leave optional fields empty or retain current aids; no additional explanation selected for this pass. |
| [arrow-weapon](../entries/arrow-weapon.yml)                                             | arrow                        | None                  | No locally declared aid field | Leave optional fields empty or retain current aids; no additional explanation selected for this pass. |
| [ascend-general](../entries/ascend-general.yml)                                         | ascend                       | None                  | existing declared association | Leave optional fields empty or retain current aids; no additional explanation selected for this pass. |
| [assemble-general](../entries/assemble-general.yml)                                     | assemble                     | None                  | No locally declared aid field | Leave optional fields empty or retain current aids; no additional explanation selected for this pass. |
| [attach-general](../entries/attach-general.yml)                                         | attach                       | None                  | No locally declared aid field | Leave optional fields empty or retain current aids; no additional explanation selected for this pass. |
| [attack-power-combat](../entries/attack-power-combat.yml)                               | attack power                 | WF37                  | No locally declared aid field | Applied listed enrichment or derived display; retained all other fields.                              |
| [attack-up-zelda](../entries/attack-up-zelda.yml)                                       | Attack Up                    | None                  | No locally declared aid field | Keep name/localization and Notes; no verified naming explanation selected.                            |
| [autobuild-zelda](../entries/autobuild-zelda.yml)                                       | Autobuild                    | WF01                  | No locally declared aid field | Applied listed enrichment or derived display; retained all other fields.                              |
| [awaken-story](../entries/awaken-story.yml)                                             | awaken                       | None                  | No locally declared aid field | Leave optional fields empty or retain current aids; no additional explanation selected for this pass. |
| [axe-weapon](../entries/axe-weapon.yml)                                                 | axe                          | None                  | No locally declared aid field | Leave optional fields empty or retain current aids; no additional explanation selected for this pass. |
| [balloon-engineering](../entries/balloon-engineering.yml)                               | balloon                      | None                  | No locally declared aid field | Leave optional fields empty or retain current aids; no additional explanation selected for this pass. |
| [banana-food](../entries/banana-food.yml)                                               | banana                       | None                  | No locally declared aid field | Leave optional fields empty or retain current aids; no additional explanation selected for this pass. |
| [bandanna-clothing](../entries/bandanna-clothing.yml)                                   | bandanna                     | None                  | No locally declared aid field | Leave optional fields empty or retain current aids; no additional explanation selected for this pass. |
| [bass-food](../entries/bass-food.yml)                                                   | bass                         | N17                   | No locally declared aid field | Applied listed enrichment or derived display; retained all other fields.                              |
| [battery-engineering](../entries/battery-engineering.yml)                               | battery                      | None                  | No locally declared aid field | Leave optional fields empty or retain current aids; no additional explanation selected for this pass. |
| [beam-emitter-zelda](../entries/beam-emitter-zelda.yml)                                 | Beam Emitter                 | WF45                  | No locally declared aid field | Applied listed enrichment or derived display; retained all other fields.                              |
| [beam-engineering](../entries/beam-engineering.yml)                                     | beam                         | None                  | No locally declared aid field | Leave optional fields empty or retain current aids; no additional explanation selected for this pass. |
| [bee-general](../entries/bee-general.yml)                                               | bee                          | None                  | No locally declared aid field | Leave optional fields empty or retain current aids; no additional explanation selected for this pass. |
| [big-battery-zelda](../entries/big-battery-zelda.yml)                                   | Big Battery                  | None                  | No locally declared aid field | Keep name/localization and Notes; no verified naming explanation selected.                            |
| [big-wheel-zelda](../entries/big-wheel-zelda.yml)                                       | Big Wheel                    | None                  | No locally declared aid field | Keep name/localization and Notes; no verified naming explanation selected.                            |
| [blessing-general](../entries/blessing-general.yml)                                     | blessing                     | None                  | No locally declared aid field | Leave optional fields empty or retain current aids; no additional explanation selected for this pass. |
| [block-general](../entries/block-general.yml)                                           | block                        | None                  | No locally declared aid field | Leave optional fields empty or retain current aids; no additional explanation selected for this pass. |
| [bokoblin-zelda](../entries/bokoblin-zelda.yml)                                         | Bokoblin                     | None                  | No locally declared aid field | Keep name/localization and Notes; no verified naming explanation selected.                            |
| [bomb-weapon](../entries/bomb-weapon.yml)                                               | bomb                         | None                  | No locally declared aid field | Leave optional fields empty or retain current aids; no additional explanation selected for this pass. |
| [boomerang-weapon](../entries/boomerang-weapon.yml)                                     | boomerang                    | None                  | No locally declared aid field | Leave optional fields empty or retain current aids; no additional explanation selected for this pass. |
| [boots-clothing](../entries/boots-clothing.yml)                                         | boots                        | None                  | No locally declared aid field | Leave optional fields empty or retain current aids; no additional explanation selected for this pass. |
| [boss-bokoblin-zelda](../entries/boss-bokoblin-zelda.yml)                               | Boss Bokoblin                | None                  | No locally declared aid field | Keep name/localization and Notes; no verified naming explanation selected.                            |
| [bow-weapon](../entries/bow-weapon.yml)                                                 | bow                          | None                  | No locally declared aid field | Leave optional fields empty or retain current aids; no additional explanation selected for this pass. |
| [brave-general](../entries/brave-general.yml)                                           | brave                        | R10, R11              | No locally declared aid field | Applied listed enrichment or derived display; retained all other fields.                              |
| [bridge-architecture](../entries/bridge-architecture.yml)                               | bridge                       | None                  | No locally declared aid field | Leave optional fields empty or retain current aids; no additional explanation selected for this pass. |
| [brightcap-zelda](../entries/brightcap-zelda.yml)                                       | Brightcap                    | None                  | No locally declared aid field | Keep name/localization and Notes; no verified naming explanation selected.                            |
| [broadsword-weapon](../entries/broadsword-weapon.yml)                                   | broadsword                   | None                  | existing Word Formation       | Keep the existing construction; no additional field justified in this pass.                           |
| [build-construction](../entries/build-construction.yml)                                 | build                        | None                  | No locally declared aid field | Leave optional fields empty or retain current aids; no additional explanation selected for this pass. |
| [butter-food](../entries/butter-food.yml)                                               | butter                       | None                  | No locally declared aid field | Leave optional fields empty or retain current aids; no additional explanation selected for this pass. |
| [calamity-story](../entries/calamity-story.yml)                                         | calamity                     | None                  | No locally declared aid field | Leave optional fields empty or retain current aids; no additional explanation selected for this pass. |
| [cancel-general](../entries/cancel-general.yml)                                         | cancel                       | None                  | No locally declared aid field | Leave optional fields empty or retain current aids; no additional explanation selected for this pass. |
| [cannon-engineering](../entries/cannon-engineering.yml)                                 | cannon                       | None                  | No locally declared aid field | Leave optional fields empty or retain current aids; no additional explanation selected for this pass. |
| [canyon-terrain](../entries/canyon-terrain.yml)                                         | canyon                       | R19                   | No locally declared aid field | Applied listed enrichment or derived display; retained all other fields.                              |
| [captain-construct-zelda](../entries/captain-construct-zelda.yml)                       | Captain Construct            | None                  | No locally declared aid field | Keep name/localization and Notes; no verified naming explanation selected.                            |
| [carrot-food](../entries/carrot-food.yml)                                               | carrot                       | None                  | No locally declared aid field | Leave optional fields empty or retain current aids; no additional explanation selected for this pass. |
| [cart-engineering](../entries/cart-engineering.yml)                                     | cart                         | None                  | No locally declared aid field | Leave optional fields empty or retain current aids; no additional explanation selected for this pass. |
| [castle-architecture](../entries/castle-architecture.yml)                               | castle                       | None                  | No locally declared aid field | Leave optional fields empty or retain current aids; no additional explanation selected for this pass. |
| [cave-architecture](../entries/cave-architecture.yml)                                   | cave                         | R18                   | No locally declared aid field | Applied listed enrichment or derived display; retained all other fields.                              |
| [cavern-architecture](../entries/cavern-architecture.yml)                               | cavern                       | N12, R18              | No locally declared aid field | Applied listed enrichment or derived display; retained all other fields.                              |
| [ceiling-architecture](../entries/ceiling-architecture.yml)                             | ceiling                      | None                  | No locally declared aid field | Leave optional fields empty or retain current aids; no additional explanation selected for this pass. |
| [chamber-architecture](../entries/chamber-architecture.yml)                             | chamber                      | None                  | No locally declared aid field | Leave optional fields empty or retain current aids; no additional explanation selected for this pass. |
| [champion-general](../entries/champion-general.yml)                                     | champion                     | None                  | No locally declared aid field | Leave optional fields empty or retain current aids; no additional explanation selected for this pass. |
| [character-profiles-zelda](../entries/character-profiles-zelda.yml)                     | Character Profiles           | None                  | No locally declared aid field | Keep name/localization and Notes; no verified naming explanation selected.                            |
| [chasm-terrain](../entries/chasm-terrain.yml)                                           | chasm                        | N11, R19              | No locally declared aid field | Applied listed enrichment or derived display; retained all other fields.                              |
| [cheese-food](../entries/cheese-food.yml)                                               | cheese                       | None                  | No locally declared aid field | Leave optional fields empty or retain current aids; no additional explanation selected for this pass. |
| [chilly-general](../entries/chilly-general.yml)                                         | chilly                       | None                  | No locally declared aid field | Leave optional fields empty or retain current aids; no additional explanation selected for this pass. |
| [chuchu-zelda](../entries/chuchu-zelda.yml)                                             | Chuchu                       | None                  | No locally declared aid field | Keep name/localization and Notes; no verified naming explanation selected.                            |
| [citadel-architecture](../entries/citadel-architecture.yml)                             | citadel                      | None                  | No locally declared aid field | Leave optional fields empty or retain current aids; no additional explanation selected for this pass. |
| [civilization-general](../entries/civilization-general.yml)                             | civilization                 | None                  | No locally declared aid field | Leave optional fields empty or retain current aids; no additional explanation selected for this pass. |
| [claymore-weapon](../entries/claymore-weapon.yml)                                       | claymore                     | None                  | No locally declared aid field | Leave optional fields empty or retain current aids; no additional explanation selected for this pass. |
| [cliff-terrain](../entries/cliff-terrain.yml)                                           | cliff                        | None                  | No locally declared aid field | Leave optional fields empty or retain current aids; no additional explanation selected for this pass. |
| [climb-general](../entries/climb-general.yml)                                           | climb                        | None                  | No locally declared aid field | Leave optional fields empty or retain current aids; no additional explanation selected for this pass. |
| [climbing-gear-zelda](../entries/climbing-gear-zelda.yml)                               | Climbing Gear                | None                  | No locally declared aid field | Keep name/localization and Notes; no verified naming explanation selected.                            |
| [closed-general](../entries/closed-general.yml)                                         | closed                       | None                  | No locally declared aid field | Leave optional fields empty or retain current aids; no additional explanation selected for this pass. |
| [club-weapon](../entries/club-weapon.yml)                                               | club                         | None                  | No locally declared aid field | Leave optional fields empty or retain current aids; no additional explanation selected for this pass. |
| [clue-story](../entries/clue-story.yml)                                                 | clue                         | None                  | No locally declared aid field | Leave optional fields empty or retain current aids; no additional explanation selected for this pass. |
| [cold-resistance-combat](../entries/cold-resistance-combat.yml)                         | cold resistance              | None                  | existing declared association | Leave optional fields empty or retain current aids; no additional explanation selected for this pass. |
| [colgera-zelda](../entries/colgera-zelda.yml)                                           | Colgera                      | None                  | No locally declared aid field | Keep name/localization and Notes; no verified naming explanation selected.                            |
| [colossal-combat](../entries/colossal-combat.yml)                                       | colossal                     | None                  | No locally declared aid field | Leave optional fields empty or retain current aids; no additional explanation selected for this pass. |
| [compendium-general](../entries/compendium-general.yml)                                 | compendium                   | None                  | No locally declared aid field | Leave optional fields empty or retain current aids; no additional explanation selected for this pass. |
| [construct-being-zelda](../entries/construct-being-zelda.yml)                           | Construct                    | None                  | No locally declared aid field | Keep name/localization and Notes; no verified naming explanation selected.                            |
| [construct-factory-zelda](../entries/construct-factory-zelda.yml)                       | Construct Factory            | None                  | No locally declared aid field | Keep name/localization and Notes; no verified naming explanation selected.                            |
| [construct-general](../entries/construct-general.yml)                                   | construct                    | None                  | No locally declared aid field | Leave optional fields empty or retain current aids; no additional explanation selected for this pass. |
| [construct-head-zelda](../entries/construct-head-zelda.yml)                             | Construct Head               | WF47                  | No locally declared aid field | Applied listed enrichment or derived display; retained all other fields.                              |
| [cook-cooking](../entries/cook-cooking.yml)                                             | cook                         | None                  | No locally declared aid field | Leave optional fields empty or retain current aids; no additional explanation selected for this pass. |
| [cooking-pot-cooking](../entries/cooking-pot-cooking.yml)                               | cooking pot                  | WF38                  | No locally declared aid field | Applied listed enrichment or derived display; retained all other fields.                              |
| [coordinates-ui](../entries/coordinates-ui.yml)                                         | coordinates                  | None                  | No locally declared aid field | Leave optional fields empty or retain current aids; no additional explanation selected for this pass. |
| [corridor-architecture](../entries/corridor-architecture.yml)                           | corridor                     | None                  | No locally declared aid field | Leave optional fields empty or retain current aids; no additional explanation selected for this pass. |
| [corrupted-combat](../entries/corrupted-combat.yml)                                     | corrupted                    | None                  | No locally declared aid field | Leave optional fields empty or retain current aids; no additional explanation selected for this pass. |
| [courageous-general](../entries/courageous-general.yml)                                 | courageous                   | WF35, ETY14, R10, R12 | No locally declared aid field | Applied listed enrichment or derived display; retained all other fields.                              |
| [courser-bee-honey-zelda](../entries/courser-bee-honey-zelda.yml)                       | Courser Bee Honey            | None                  | No locally declared aid field | Keep name/localization and Notes; no verified naming explanation selected.                            |
| [crater-terrain](../entries/crater-terrain.yml)                                         | crater                       | None                  | No locally declared aid field | Leave optional fields empty or retain current aids; no additional explanation selected for this pass. |
| [crisis-story](../entries/crisis-story.yml)                                             | crisis                       | None                  | No locally declared aid field | Leave optional fields empty or retain current aids; no additional explanation selected for this pass. |
| [critical-hit-combat](../entries/critical-hit-combat.yml)                               | critical hit                 | WF36                  | No locally declared aid field | Applied listed enrichment or derived display; retained all other fields.                              |
| [crouch-general](../entries/crouch-general.yml)                                         | crouch                       | None                  | No locally declared aid field | Leave optional fields empty or retain current aids; no additional explanation selected for this pass. |
| [crystal-general](../entries/crystal-general.yml)                                       | crystal                      | None                  | No locally declared aid field | Leave optional fields empty or retain current aids; no additional explanation selected for this pass. |
| [crystal-refinery-zelda](../entries/crystal-refinery-zelda.yml)                         | Crystal Refinery             | None                  | No locally declared aid field | Keep name/localization and Notes; no verified naming explanation selected.                            |
| [cunning-general](../entries/cunning-general.yml)                                       | cunning                      | None                  | No locally declared aid field | Leave optional fields empty or retain current aids; no additional explanation selected for this pass. |
| [curious-general](../entries/curious-general.yml)                                       | curious                      | None                  | No locally declared aid field | Leave optional fields empty or retain current aids; no additional explanation selected for this pass. |
| [dagger-weapon](../entries/dagger-weapon.yml)                                           | dagger                       | None                  | No locally declared aid field | Leave optional fields empty or retain current aids; no additional explanation selected for this pass. |
| [damage-combat](../entries/damage-combat.yml)                                           | damage                       | None                  | No locally declared aid field | Leave optional fields empty or retain current aids; no additional explanation selected for this pass. |
| [death-mountain-zelda](../entries/death-mountain-zelda.yml)                             | Death Mountain               | None                  | No locally declared aid field | Keep name/localization and Notes; no verified naming explanation selected.                            |
| [decayed-general](../entries/decayed-general.yml)                                       | decayed                      | None                  | existing declared association | Leave optional fields empty or retain current aids; no additional explanation selected for this pass. |
| [defense-combat](../entries/defense-combat.yml)                                         | defense                      | None                  | No locally declared aid field | Leave optional fields empty or retain current aids; no additional explanation selected for this pass. |
| [defense-up-zelda](../entries/defense-up-zelda.yml)                                     | Defense Up                   | None                  | No locally declared aid field | Keep name/localization and Notes; no verified naming explanation selected.                            |
| [deity-story](../entries/deity-story.yml)                                               | deity                        | None                  | No locally declared aid field | Leave optional fields empty or retain current aids; no additional explanation selected for this pass. |
| [demon-dragon-zelda](../entries/demon-dragon-zelda.yml)                                 | Demon Dragon                 | None                  | No locally declared aid field | Keep name/localization and Notes; no verified naming explanation selected.                            |
| [demon-king-ganondorf-zelda](../entries/demon-king-ganondorf-zelda.yml)                 | Demon King Ganondorf         | None                  | No locally declared aid field | Keep name/localization and Notes; no verified naming explanation selected.                            |
| [demon-king-zelda](../entries/demon-king-zelda.yml)                                     | Demon King                   | None                  | No locally declared aid field | Keep name/localization and Notes; no verified naming explanation selected.                            |
| [demon-kings-bow-zelda](../entries/demon-kings-bow-zelda.yml)                           | Demon King's Bow             | None                  | No locally declared aid field | Keep name/localization and Notes; no verified naming explanation selected.                            |
| [demon-story](../entries/demon-story.yml)                                               | demon                        | None                  | No locally declared aid field | Leave optional fields empty or retain current aids; no additional explanation selected for this pass. |
| [depths-terrain](../entries/depths-terrain.yml)                                         | depths                       | None                  | No locally declared aid field | Leave optional fields empty or retain current aids; no additional explanation selected for this pass. |
| [descend-general](../entries/descend-general.yml)                                       | descend                      | None                  | No locally declared aid field | Leave optional fields empty or retain current aids; no additional explanation selected for this pass. |
| [descendant-general](../entries/descendant-general.yml)                                 | descendant                   | WF33, R16             | No locally declared aid field | Applied listed enrichment or derived display; retained all other fields.                              |
| [desert-terrain](../entries/desert-terrain.yml)                                         | desert                       | None                  | No locally declared aid field | Leave optional fields empty or retain current aids; no additional explanation selected for this pass. |
| [destination-ui](../entries/destination-ui.yml)                                         | destination                  | None                  | No locally declared aid field | Leave optional fields empty or retain current aids; no additional explanation selected for this pass. |
| [destroy-general](../entries/destroy-general.yml)                                       | destroy                      | None                  | No locally declared aid field | Leave optional fields empty or retain current aids; no additional explanation selected for this pass. |
| [detach-general](../entries/detach-general.yml)                                         | detach                       | None                  | No locally declared aid field | Leave optional fields empty or retain current aids; no additional explanation selected for this pass. |
| [determined-general](../entries/determined-general.yml)                                 | determined                   | None                  | No locally declared aid field | Leave optional fields empty or retain current aids; no additional explanation selected for this pass. |
| [device-dispenser-zelda](../entries/device-dispenser-zelda.yml)                         | Device Dispenser             | WF48                  | No locally declared aid field | Applied listed enrichment or derived display; retained all other fields.                              |
| [device-engineering](../entries/device-engineering.yml)                                 | device                       | None                  | No locally declared aid field | Leave optional fields empty or retain current aids; no additional explanation selected for this pass. |
| [diligent-general](../entries/diligent-general.yml)                                     | diligent                     | None                  | No locally declared aid field | Leave optional fields empty or retain current aids; no additional explanation selected for this pass. |
| [disguised-combat](../entries/disguised-combat.yml)                                     | disguised                    | None                  | No locally declared aid field | Leave optional fields empty or retain current aids; no additional explanation selected for this pass. |
| [dispenser-engineering](../entries/dispenser-engineering.yml)                           | dispenser                    | WF26                  | No locally declared aid field | Applied listed enrichment or derived display; retained all other fields.                              |
| [distinctive-combat](../entries/distinctive-combat.yml)                                 | distinctive                  | None                  | No locally declared aid field | Leave optional fields empty or retain current aids; no additional explanation selected for this pass. |
| [dive-general](../entries/dive-general.yml)                                             | dive                         | None                  | No locally declared aid field | Leave optional fields empty or retain current aids; no additional explanation selected for this pass. |
| [dodge-general](../entries/dodge-general.yml)                                           | dodge                        | None                  | No locally declared aid field | Leave optional fields empty or retain current aids; no additional explanation selected for this pass. |
| [door-architecture](../entries/door-architecture.yml)                                   | door                         | None                  | No locally declared aid field | Leave optional fields empty or retain current aids; no additional explanation selected for this pass. |
| [draconification-zelda](../entries/draconification-zelda.yml)                           | draconification              | None                  | No locally declared aid field | Keep name/localization and Notes; no verified naming explanation selected.                            |
| [dragon-story](../entries/dragon-story.yml)                                             | dragon                       | None                  | No locally declared aid field | Leave optional fields empty or retain current aids; no additional explanation selected for this pass. |
| [draw-general](../entries/draw-general.yml)                                             | draw                         | None                  | No locally declared aid field | Leave optional fields empty or retain current aids; no additional explanation selected for this pass. |
| [drop-general](../entries/drop-general.yml)                                             | drop                         | None                  | No locally declared aid field | Leave optional fields empty or retain current aids; no additional explanation selected for this pass. |
| [durability-equipment](../entries/durability-equipment.yml)                             | durability                   | None                  | existing Word Formation       | Keep the existing construction; no additional field justified in this pass.                           |
| [durable-general](../entries/durable-general.yml)                                       | durable                      | None                  | existing declared association | Leave optional fields empty or retain current aids; no additional explanation selected for this pass. |
| [eagle-general](../entries/eagle-general.yml)                                           | eagle                        | None                  | No locally declared aid field | Leave optional fields empty or retain current aids; no additional explanation selected for this pass. |
| [egg-food](../entries/egg-food.yml)                                                     | egg                          | None                  | No locally declared aid field | Leave optional fields empty or retain current aids; no additional explanation selected for this pass. |
| [electro-cooking](../entries/electro-cooking.yml)                                       | Electro                      | None                  | No locally declared aid field | Leave optional fields empty or retain current aids; no additional explanation selected for this pass. |
| [elixir-cooking](../entries/elixir-cooking.yml)                                         | elixir                       | ETY04                 | No locally declared aid field | Applied listed enrichment or derived display; retained all other fields.                              |
| [elusive-combat](../entries/elusive-combat.yml)                                         | elusive                      | None                  | No locally declared aid field | Leave optional fields empty or retain current aids; no additional explanation selected for this pass. |
| [emitter-engineering](../entries/emitter-engineering.yml)                               | emitter                      | R28                   | existing Word Formation       | Applied listed enrichment or derived display; retained all other fields.                              |
| [endura-carrot-zelda](../entries/endura-carrot-zelda.yml)                               | Endura Carrot                | None                  | No locally declared aid field | Keep name/localization and Notes; no verified naming explanation selected.                            |
| [endura-shroom-zelda](../entries/endura-shroom-zelda.yml)                               | Endura Shroom                | None                  | No locally declared aid field | Keep name/localization and Notes; no verified naming explanation selected.                            |
| [endure-general](../entries/endure-general.yml)                                         | endure                       | R27                   | No locally declared aid field | Applied listed enrichment or derived display; retained all other fields.                              |
| [energize-general](../entries/energize-general.yml)                                     | energize                     | R27                   | No locally declared aid field | Applied listed enrichment or derived display; retained all other fields.                              |
| [enhancement-general](../entries/enhancement-general.yml)                               | enhancement                  | WF30                  | No locally declared aid field | Applied listed enrichment or derived display; retained all other fields.                              |
| [enormous-combat](../entries/enormous-combat.yml)                                       | enormous                     | None                  | No locally declared aid field | Leave optional fields empty or retain current aids; no additional explanation selected for this pass. |
| [enthusiastic-general](../entries/enthusiastic-general.yml)                             | enthusiastic                 | ETY15                 | No locally declared aid field | Applied listed enrichment or derived display; retained all other fields.                              |
| [entrance-architecture](../entries/entrance-architecture.yml)                           | entrance                     | None                  | No locally declared aid field | Leave optional fields empty or retain current aids; no additional explanation selected for this pass. |
| [equip-general](../entries/equip-general.yml)                                           | equip                        | None                  | No locally declared aid field | Leave optional fields empty or retain current aids; no additional explanation selected for this pass. |
| [equipment-equipment](../entries/equipment-equipment.yml)                               | equipment                    | None                  | No locally declared aid field | Leave optional fields empty or retain current aids; no additional explanation selected for this pass. |
| [evermean-zelda](../entries/evermean-zelda.yml)                                         | Evermean                     | None                  | No locally declared aid field | Keep name/localization and Notes; no verified naming explanation selected.                            |
| [examine-general](../entries/examine-general.yml)                                       | examine                      | None                  | No locally declared aid field | Leave optional fields empty or retain current aids; no additional explanation selected for this pass. |
| [exit-architecture](../entries/exit-architecture.yml)                                   | exit                         | None                  | No locally declared aid field | Leave optional fields empty or retain current aids; no additional explanation selected for this pass. |
| [factory-architecture](../entries/factory-architecture.yml)                             | factory                      | None                  | No locally declared aid field | Leave optional fields empty or retain current aids; no additional explanation selected for this pass. |
| [fan-engineering](../entries/fan-engineering.yml)                                       | fan                          | None                  | No locally declared aid field | Leave optional fields empty or retain current aids; no additional explanation selected for this pass. |
| [fearless-general](../entries/fearless-general.yml)                                     | fearless                     | WF08, R11             | No locally declared aid field | Applied listed enrichment or derived display; retained all other fields.                              |
| [fierce-deity-armor-zelda](../entries/fierce-deity-armor-zelda.yml)                     | Fierce Deity Armor           | None                  | No locally declared aid field | Keep name/localization and Notes; no verified naming explanation selected.                            |
| [fierce-general](../entries/fierce-general.yml)                                         | fierce                       | None                  | No locally declared aid field | Leave optional fields empty or retain current aids; no additional explanation selected for this pass. |
| [fifth-general](../entries/fifth-general.yml)                                           | fifth                        | None                  | No locally declared aid field | Leave optional fields empty or retain current aids; no additional explanation selected for this pass. |
| [find-general](../entries/find-general.yml)                                             | find                         | None                  | No locally declared aid field | Leave optional fields empty or retain current aids; no additional explanation selected for this pass. |
| [fire-general](../entries/fire-general.yml)                                             | fire                         | None                  | No locally declared aid field | Leave optional fields empty or retain current aids; no additional explanation selected for this pass. |
| [fire-temple-zelda](../entries/fire-temple-zelda.yml)                                   | Fire Temple                  | None                  | No locally declared aid field | Keep name/localization and Notes; no verified naming explanation selected.                            |
| [flame-emitter-zelda](../entries/flame-emitter-zelda.yml)                               | Flame Emitter                | WF42                  | No locally declared aid field | Applied listed enrichment or derived display; retained all other fields.                              |
| [flame-general](../entries/flame-general.yml)                                           | flame                        | None                  | No locally declared aid field | Leave optional fields empty or retain current aids; no additional explanation selected for this pass. |
| [flamebreaker-armor-zelda](../entries/flamebreaker-armor-zelda.yml)                     | Flamebreaker Armor           | None                  | No locally declared aid field | Keep name/localization and Notes; no verified naming explanation selected.                            |
| [flux-construct-zelda](../entries/flux-construct-zelda.yml)                             | Flux Construct               | None                  | No locally declared aid field | Keep name/localization and Notes; no verified naming explanation selected.                            |
| [foothill-terrain](../entries/foothill-terrain.yml)                                     | foothill                     | WF22                  | No locally declared aid field | Applied listed enrichment or derived display; retained all other fields.                              |
| [forest-terrain](../entries/forest-terrain.yml)                                         | forest                       | None                  | No locally declared aid field | Leave optional fields empty or retain current aids; no additional explanation selected for this pass. |
| [forge-architecture](../entries/forge-architecture.yml)                                 | forge                        | None                  | No locally declared aid field | Leave optional fields empty or retain current aids; no additional explanation selected for this pass. |
| [forget-general](../entries/forget-general.yml)                                         | forget                       | None                  | No locally declared aid field | Leave optional fields empty or retain current aids; no additional explanation selected for this pass. |
| [forgotten-temple-zelda](../entries/forgotten-temple-zelda.yml)                         | Forgotten Temple             | None                  | No locally declared aid field | Keep name/localization and Notes; no verified naming explanation selected.                            |
| [fortress-architecture](../entries/fortress-architecture.yml)                           | fortress                     | None                  | No locally declared aid field | Leave optional fields empty or retain current aids; no additional explanation selected for this pass. |
| [fountain-architecture](../entries/fountain-architecture.yml)                           | fountain                     | None                  | No locally declared aid field | Leave optional fields empty or retain current aids; no additional explanation selected for this pass. |
| [frog-general](../entries/frog-general.yml)                                             | frog                         | None                  | No locally declared aid field | Leave optional fields empty or retain current aids; no additional explanation selected for this pass. |
| [froggy-sleeve-zelda](../entries/froggy-sleeve-zelda.yml)                               | Froggy Sleeve                | None                  | No locally declared aid field | Keep name/localization and Notes; no verified naming explanation selected.                            |
| [frost-emitter-zelda](../entries/frost-emitter-zelda.yml)                               | Frost Emitter                | WF43                  | No locally declared aid field | Applied listed enrichment or derived display; retained all other fields.                              |
| [frost-general](../entries/frost-general.yml)                                           | frost                        | None                  | No locally declared aid field | Leave optional fields empty or retain current aids; no additional explanation selected for this pass. |
| [frox-zelda](../entries/frox-zelda.yml)                                                 | Frox                         | None                  | No locally declared aid field | Keep name/localization and Notes; no verified naming explanation selected.                            |
| [fruit-food](../entries/fruit-food.yml)                                                 | fruit                        | None                  | No locally declared aid field | Leave optional fields empty or retain current aids; no additional explanation selected for this pass. |
| [fry-cooking](../entries/fry-cooking.yml)                                               | fry                          | None                  | No locally declared aid field | Leave optional fields empty or retain current aids; no additional explanation selected for this pass. |
| [fuse-general](../entries/fuse-general.yml)                                             | fuse                         | None                  | existing declared association | Leave optional fields empty or retain current aids; no additional explanation selected for this pass. |
| [ganondorf-zelda](../entries/ganondorf-zelda.yml)                                       | Ganondorf                    | None                  | No locally declared aid field | Keep name/localization and Notes; no verified naming explanation selected.                            |
| [gate-architecture](../entries/gate-architecture.yml)                                   | gate                         | None                  | No locally declared aid field | Leave optional fields empty or retain current aids; no additional explanation selected for this pass. |
| [gear-equipment](../entries/gear-equipment.yml)                                         | gear                         | None                  | No locally declared aid field | Leave optional fields empty or retain current aids; no additional explanation selected for this pass. |
| [general-store-architecture](../entries/general-store-architecture.yml)                 | general store                | None                  | No locally declared aid field | Leave optional fields empty or retain current aids; no additional explanation selected for this pass. |
| [geoglyph-game](../entries/geoglyph-game.yml)                                           | geoglyph                     | None                  | No locally declared aid field | Leave optional fields empty or retain current aids; no additional explanation selected for this pass. |
| [gerudo-desert-zelda](../entries/gerudo-desert-zelda.yml)                               | Gerudo Desert                | None                  | No locally declared aid field | Keep name/localization and Notes; no verified naming explanation selected.                            |
| [gerudo-town-zelda](../entries/gerudo-town-zelda.yml)                                   | Gerudo Town                  | None                  | No locally declared aid field | Keep name/localization and Notes; no verified naming explanation selected.                            |
| [gerudo-zelda](../entries/gerudo-zelda.yml)                                             | Gerudo                       | None                  | No locally declared aid field | Keep name/localization and Notes; no verified naming explanation selected.                            |
| [gibdo-zelda](../entries/gibdo-zelda.yml)                                               | Gibdo                        | None                  | No locally declared aid field | Keep name/localization and Notes; no verified naming explanation selected.                            |
| [gleeok-zelda](../entries/gleeok-zelda.yml)                                             | Gleeok                       | None                  | No locally declared aid field | Keep name/localization and Notes; no verified naming explanation selected.                            |
| [glide-general](../entries/glide-general.yml)                                           | glide                        | None                  | No locally declared aid field | Leave optional fields empty or retain current aids; no additional explanation selected for this pass. |
| [glide-shirt-zelda](../entries/glide-shirt-zelda.yml)                                   | Glide Shirt                  | None                  | No locally declared aid field | Keep name/localization and Notes; no verified naming explanation selected.                            |
| [gloom-recovery-zelda](../entries/gloom-recovery-zelda.yml)                             | Gloom Recovery               | None                  | No locally declared aid field | Keep name/localization and Notes; no verified naming explanation selected.                            |
| [gloom-resistance-zelda](../entries/gloom-resistance-zelda.yml)                         | Gloom Resistance             | None                  | No locally declared aid field | Keep name/localization and Notes; no verified naming explanation selected.                            |
| [gloom-spawn-zelda](../entries/gloom-spawn-zelda.yml)                                   | Gloom Spawn                  | None                  | No locally declared aid field | Keep name/localization and Notes; no verified naming explanation selected.                            |
| [gloom-sword-zelda](../entries/gloom-sword-zelda.yml)                                   | Gloom Sword                  | None                  | No locally declared aid field | Keep name/localization and Notes; no verified naming explanation selected.                            |
| [gloom-zelda](../entries/gloom-zelda.yml)                                               | Gloom                        | None                  | No locally declared aid field | Keep name/localization and Notes; no verified naming explanation selected.                            |
| [golden-apple-zelda](../entries/golden-apple-zelda.yml)                                 | Golden Apple                 | None                  | No locally declared aid field | Keep name/localization and Notes; no verified naming explanation selected.                            |
| [golden-general](../entries/golden-general.yml)                                         | golden                       | None                  | No locally declared aid field | Leave optional fields empty or retain current aids; no additional explanation selected for this pass. |
| [goron-city-zelda](../entries/goron-city-zelda.yml)                                     | Goron City                   | None                  | No locally declared aid field | Keep name/localization and Notes; no verified naming explanation selected.                            |
| [goron-spice-zelda](../entries/goron-spice-zelda.yml)                                   | Goron Spice                  | None                  | No locally declared aid field | Keep name/localization and Notes; no verified naming explanation selected.                            |
| [goron-zelda](../entries/goron-zelda.yml)                                               | Goron                        | None                  | No locally declared aid field | Keep name/localization and Notes; no verified naming explanation selected.                            |
| [gourmet-food](../entries/gourmet-food.yml)                                             | gourmet                      | None                  | No locally declared aid field | Leave optional fields empty or retain current aids; no additional explanation selected for this pass. |
| [great-abandoned-central-mine-zelda](../entries/great-abandoned-central-mine-zelda.yml) | Great Abandoned Central Mine | None                  | No locally declared aid field | Keep name/localization and Notes; no verified naming explanation selected.                            |
| [great-deku-tree-zelda](../entries/great-deku-tree-zelda.yml)                           | Great Deku Tree              | None                  | No locally declared aid field | Keep name/localization and Notes; no verified naming explanation selected.                            |
| [great-eagle-bow-zelda](../entries/great-eagle-bow-zelda.yml)                           | Great Eagle Bow              | None                  | No locally declared aid field | Keep name/localization and Notes; no verified naming explanation selected.                            |
| [great-sky-island-zelda](../entries/great-sky-island-zelda.yml)                         | Great Sky Island             | None                  | No locally declared aid field | Keep name/localization and Notes; no verified naming explanation selected.                            |
| [greaves-clothing](../entries/greaves-clothing.yml)                                     | greaves                      | None                  | No locally declared aid field | Leave optional fields empty or retain current aids; no additional explanation selected for this pass. |
| [greens-food](../entries/greens-food.yml)                                               | greens                       | None                  | No locally declared aid field | Leave optional fields empty or retain current aids; no additional explanation selected for this pass. |
| [grill-cooking](../entries/grill-cooking.yml)                                           | grill                        | N16, R20              | No locally declared aid field | Applied listed enrichment or derived display; retained all other fields.                              |
| [guardian-general](../entries/guardian-general.yml)                                     | guardian                     | None                  | No locally declared aid field | Leave optional fields empty or retain current aids; no additional explanation selected for this pass. |
| [guidance-story](../entries/guidance-story.yml)                                         | guidance                     | None                  | No locally declared aid field | Leave optional fields empty or retain current aids; no additional explanation selected for this pass. |
| [halberd-weapon](../entries/halberd-weapon.yml)                                         | halberd                      | None                  | No locally declared aid field | Leave optional fields empty or retain current aids; no additional explanation selected for this pass. |
| [hammer-weapon](../entries/hammer-weapon.yml)                                           | hammer                       | None                  | No locally declared aid field | Leave optional fields empty or retain current aids; no additional explanation selected for this pass. |
| [hand-general](../entries/hand-general.yml)                                             | hand                         | None                  | No locally declared aid field | Leave optional fields empty or retain current aids; no additional explanation selected for this pass. |
| [hasty-general](../entries/hasty-general.yml)                                           | hasty                        | None                  | No locally declared aid field | Leave optional fields empty or retain current aids; no additional explanation selected for this pass. |
| [hateno-village-zelda](../entries/hateno-village-zelda.yml)                             | Hateno Village               | None                  | No locally declared aid field | Keep name/localization and Notes; no verified naming explanation selected.                            |
| [head-general](../entries/head-general.yml)                                             | head                         | None                  | No locally declared aid field | Leave optional fields empty or retain current aids; no additional explanation selected for this pass. |
| [headband-clothing](../entries/headband-clothing.yml)                                   | headband                     | WF15, R25             | No locally declared aid field | Applied listed enrichment or derived display; retained all other fields.                              |
| [headdress-clothing](../entries/headdress-clothing.yml)                                 | headdress                    | WF16, R25             | No locally declared aid field | Applied listed enrichment or derived display; retained all other fields.                              |
| [hearty-general](../entries/hearty-general.yml)                                         | hearty                       | None                  | No locally declared aid field | Leave optional fields empty or retain current aids; no additional explanation selected for this pass. |
| [hearty-radish-zelda](../entries/hearty-radish-zelda.yml)                               | Hearty Radish                | None                  | No locally declared aid field | Keep name/localization and Notes; no verified naming explanation selected.                            |
| [hearty-salmon-zelda](../entries/hearty-salmon-zelda.yml)                               | Hearty Salmon                | None                  | No locally declared aid field | Keep name/localization and Notes; no verified naming explanation selected.                            |
| [hearty-truffle-zelda](../entries/hearty-truffle-zelda.yml)                             | Hearty Truffle               | None                  | No locally declared aid field | Keep name/localization and Notes; no verified naming explanation selected.                            |
| [heat-resistance-combat](../entries/heat-resistance-combat.yml)                         | heat resistance              | None                  | No locally declared aid field | Leave optional fields empty or retain current aids; no additional explanation selected for this pass. |
| [heir-general](../entries/heir-general.yml)                                             | heir                         | N08, R16              | No locally declared aid field | Applied listed enrichment or derived display; retained all other fields.                              |
| [helm-clothing](../entries/helm-clothing.yml)                                           | helm                         | R24                   | No locally declared aid field | Applied listed enrichment or derived display; retained all other fields.                              |
| [helmet-clothing](../entries/helmet-clothing.yml)                                       | helmet                       | R24                   | No locally declared aid field | Applied listed enrichment or derived display; retained all other fields.                              |
| [hero-general](../entries/hero-general.yml)                                             | hero                         | None                  | No locally declared aid field | Leave optional fields empty or retain current aids; no additional explanation selected for this pass. |
| [hestu-zelda](../entries/hestu-zelda.yml)                                               | Hestu                        | None                  | No locally declared aid field | Keep name/localization and Notes; no verified naming explanation selected.                            |
| [hidden-general](../entries/hidden-general.yml)                                         | hidden                       | None                  | No locally declared aid field | Leave optional fields empty or retain current aids; no additional explanation selected for this pass. |
| [hideout-architecture](../entries/hideout-architecture.yml)                             | hideout                      | None                  | No locally declared aid field | Leave optional fields empty or retain current aids; no additional explanation selected for this pass. |
| [highlands-terrain](../entries/highlands-terrain.yml)                                   | highlands                    | WF18                  | No locally declared aid field | Applied listed enrichment or derived display; retained all other fields.                              |
| [hinox-zelda](../entries/hinox-zelda.yml)                                               | Hinox                        | None                  | No locally declared aid field | Keep name/localization and Notes; no verified naming explanation selected.                            |
| [hold-general](../entries/hold-general.yml)                                             | hold                         | None                  | No locally declared aid field | Leave optional fields empty or retain current aids; no additional explanation selected for this pass. |
| [homing-cart-zelda](../entries/homing-cart-zelda.yml)                                   | Homing Cart                  | None                  | No locally declared aid field | Keep name/localization and Notes; no verified naming explanation selected.                            |
| [homing-engineering](../entries/homing-engineering.yml)                                 | homing                       | WF24                  | No locally declared aid field | Applied listed enrichment or derived display; retained all other fields.                              |
| [honey-food](../entries/honey-food.yml)                                                 | honey                        | None                  | No locally declared aid field | Leave optional fields empty or retain current aids; no additional explanation selected for this pass. |
| [honorable-general](../entries/honorable-general.yml)                                   | honorable                    | None                  | No locally declared aid field | Leave optional fields empty or retain current aids; no additional explanation selected for this pass. |
| [hood-clothing](../entries/hood-clothing.yml)                                           | hood                         | None                  | No locally declared aid field | Leave optional fields empty or retain current aids; no additional explanation selected for this pass. |
| [horriblin-zelda](../entries/horriblin-zelda.yml)                                       | Horriblin                    | None                  | No locally declared aid field | Keep name/localization and Notes; no verified naming explanation selected.                            |
| [hostile-combat](../entries/hostile-combat.yml)                                         | hostile                      | ETY10, N06, R13       | No locally declared aid field | Applied listed enrichment or derived display; retained all other fields.                              |
| [hover-general](../entries/hover-general.yml)                                           | hover                        | None                  | No locally declared aid field | Leave optional fields empty or retain current aids; no additional explanation selected for this pass. |
| [hover-stone-zelda](../entries/hover-stone-zelda.yml)                                   | Hover Stone                  | WF41                  | No locally declared aid field | Applied listed enrichment or derived display; retained all other fields.                              |
| [hudson-zelda](../entries/hudson-zelda.yml)                                             | Hudson                       | None                  | No locally declared aid field | Keep name/localization and Notes; no verified naming explanation selected.                            |
| [hydrant-engineering](../entries/hydrant-engineering.yml)                               | hydrant                      | ETY07                 | No locally declared aid field | Applied listed enrichment or derived display; retained all other fields.                              |
| [hylian-hood-zelda](../entries/hylian-hood-zelda.yml)                                   | Hylian Hood                  | None                  | No locally declared aid field | Keep name/localization and Notes; no verified naming explanation selected.                            |
| [hylian-shield-zelda](../entries/hylian-shield-zelda.yml)                               | Hylian Shield                | None                  | No locally declared aid field | Keep name/localization and Notes; no verified naming explanation selected.                            |
| [hylian-tunic-zelda](../entries/hylian-tunic-zelda.yml)                                 | Hylian Tunic                 | None                  | No locally declared aid field | Keep name/localization and Notes; no verified naming explanation selected.                            |
| [hylian-zelda](../entries/hylian-zelda.yml)                                             | Hylian                       | None                  | No locally declared aid field | Keep name/localization and Notes; no verified naming explanation selected.                            |
| [hyrule-bass-zelda](../entries/hyrule-bass-zelda.yml)                                   | Hyrule Bass                  | None                  | No locally declared aid field | Keep name/localization and Notes; no verified naming explanation selected.                            |
| [hyrule-castle-zelda](../entries/hyrule-castle-zelda.yml)                               | Hyrule Castle                | None                  | No locally declared aid field | Keep name/localization and Notes; no verified naming explanation selected.                            |
| [hyrule-compendium-zelda](../entries/hyrule-compendium-zelda.yml)                       | Hyrule Compendium            | None                  | No locally declared aid field | Keep name/localization and Notes; no verified naming explanation selected.                            |
| [hyrule-zelda](../entries/hyrule-zelda.yml)                                             | Hyrule                       | None                  | No locally declared aid field | Keep name/localization and Notes; no verified naming explanation selected.                            |
| [impa-zelda](../entries/impa-zelda.yml)                                                 | Impa                         | None                  | No locally declared aid field | Keep name/localization and Notes; no verified naming explanation selected.                            |
| [imprison-story](../entries/imprison-story.yml)                                         | imprison                     | None                  | No locally declared aid field | Leave optional fields empty or retain current aids; no additional explanation selected for this pass. |
| [ingredient-cooking](../entries/ingredient-cooking.yml)                                 | ingredient                   | R26                   | No locally declared aid field | Applied listed enrichment or derived display; retained all other fields.                              |
| [inherit-story](../entries/inherit-story.yml)                                           | inherit                      | None                  | No locally declared aid field | Leave optional fields empty or retain current aids; no additional explanation selected for this pass. |
| [inn-architecture](../entries/inn-architecture.yml)                                     | inn                          | None                  | No locally declared aid field | Leave optional fields empty or retain current aids; no additional explanation selected for this pass. |
| [integrity-general](../entries/integrity-general.yml)                                   | integrity                    | ETY02                 | No locally declared aid field | Applied listed enrichment or derived display; retained all other fields.                              |
| [interact-general](../entries/interact-general.yml)                                     | interact                     | None                  | No locally declared aid field | Leave optional fields empty or retain current aids; no additional explanation selected for this pass. |
| [inventory-ui](../entries/inventory-ui.yml)                                             | inventory                    | N19, R22              | No locally declared aid field | Applied listed enrichment or derived display; retained all other fields.                              |
| [island-terrain](../entries/island-terrain.yml)                                         | island                       | None                  | No locally declared aid field | Leave optional fields empty or retain current aids; no additional explanation selected for this pass. |
| [jar-general](../entries/jar-general.yml)                                               | jar                          | None                  | No locally declared aid field | Leave optional fields empty or retain current aids; no additional explanation selected for this pass. |
| [josha-zelda](../entries/josha-zelda.yml)                                               | Josha                        | None                  | No locally declared aid field | Keep name/localization and Notes; no verified naming explanation selected.                            |
| [kakariko-village-zelda](../entries/kakariko-village-zelda.yml)                         | Kakariko Village             | None                  | No locally declared aid field | Keep name/localization and Notes; no verified naming explanation selected.                            |
| [keese-zelda](../entries/keese-zelda.yml)                                               | Keese                        | None                  | No locally declared aid field | Keep name/localization and Notes; no verified naming explanation selected.                            |
| [key-items-zelda](../entries/key-items-zelda.yml)                                       | Key Items                    | None                  | No locally declared aid field | Keep name/localization and Notes; no verified naming explanation selected.                            |
| [king-general](../entries/king-general.yml)                                             | king                         | None                  | No locally declared aid field | Leave optional fields empty or retain current aids; no additional explanation selected for this pass. |
| [king-gleeok-zelda](../entries/king-gleeok-zelda.yml)                                   | King Gleeok                  | None                  | No locally declared aid field | Keep name/localization and Notes; no verified naming explanation selected.                            |
| [kingdom-general](../entries/kingdom-general.yml)                                       | kingdom                      | WF32                  | No locally declared aid field | Applied listed enrichment or derived display; retained all other fields.                              |
| [knight-general](../entries/knight-general.yml)                                         | knight                       | None                  | No locally declared aid field | Leave optional fields empty or retain current aids; no additional explanation selected for this pass. |
| [knights-broadsword-zelda](../entries/knights-broadsword-zelda.yml)                     | Knight's Broadsword          | None                  | No locally declared aid field | Keep name/localization and Notes; no verified naming explanation selected.                            |
| [korok-forest-zelda](../entries/korok-forest-zelda.yml)                                 | Korok Forest                 | None                  | No locally declared aid field | Keep name/localization and Notes; no verified naming explanation selected.                            |
| [korok-seed-zelda](../entries/korok-seed-zelda.yml)                                     | Korok Seed                   | None                  | No locally declared aid field | Keep name/localization and Notes; no verified naming explanation selected.                            |
| [korok-zelda](../entries/korok-zelda.yml)                                               | Korok                        | None                  | No locally declared aid field | Keep name/localization and Notes; no verified naming explanation selected.                            |
| [lake-terrain](../entries/lake-terrain.yml)                                             | lake                         | None                  | No locally declared aid field | Leave optional fields empty or retain current aids; no additional explanation selected for this pass. |
| [landing-general](../entries/landing-general.yml)                                       | landing                      | None                  | No locally declared aid field | Leave optional fields empty or retain current aids; no additional explanation selected for this pass. |
| [lava-terrain](../entries/lava-terrain.yml)                                             | lava                         | None                  | No locally declared aid field | Leave optional fields empty or retain current aids; no additional explanation selected for this pass. |
| [leggings-clothing](../entries/leggings-clothing.yml)                                   | leggings                     | None                  | No locally declared aid field | Leave optional fields empty or retain current aids; no additional explanation selected for this pass. |
| [light-dragon-zelda](../entries/light-dragon-zelda.yml)                                 | Light Dragon                 | None                  | No locally declared aid field | Keep name/localization and Notes; no verified naming explanation selected.                            |
| [light-engineering](../entries/light-engineering.yml)                                   | light                        | None                  | No locally declared aid field | Leave optional fields empty or retain current aids; no additional explanation selected for this pass. |
| [light-of-blessing-zelda](../entries/light-of-blessing-zelda.yml)                       | Light of Blessing            | None                  | No locally declared aid field | Keep name/localization and Notes; no verified naming explanation selected.                            |
| [lightning-temple-zelda](../entries/lightning-temple-zelda.yml)                         | Lightning Temple             | None                  | No locally declared aid field | Keep name/localization and Notes; no verified naming explanation selected.                            |
| [lightning-terrain](../entries/lightning-terrain.yml)                                   | lightning                    | None                  | No locally declared aid field | Leave optional fields empty or retain current aids; no additional explanation selected for this pass. |
| [lightroot-zelda](../entries/lightroot-zelda.yml)                                       | Lightroot                    | None                  | existing Word Formation       | Keep the existing construction; no additional field justified in this pass.                           |
| [like-like-zelda](../entries/like-like-zelda.yml)                                       | Like Like                    | None                  | No locally declared aid field | Keep name/localization and Notes; no verified naming explanation selected.                            |
| [link-zelda](../entries/link-zelda.yml)                                                 | Link                         | None                  | No locally declared aid field | Keep name/localization and Notes; no verified naming explanation selected.                            |
| [lizalfos-zelda](../entries/lizalfos-zelda.yml)                                         | Lizalfos                     | None                  | No locally declared aid field | Keep name/localization and Notes; no verified naming explanation selected.                            |
| [lookout-general](../entries/lookout-general.yml)                                       | lookout                      | None                  | No locally declared aid field | Leave optional fields empty or retain current aids; no additional explanation selected for this pass. |
| [lookout-landing-zelda](../entries/lookout-landing-zelda.yml)                           | Lookout Landing              | None                  | No locally declared aid field | Keep name/localization and Notes; no verified naming explanation selected.                            |
| [loyal-general](../entries/loyal-general.yml)                                           | loyal                        | None                  | No locally declared aid field | Leave optional fields empty or retain current aids; no additional explanation selected for this pass. |
| [lurelin-village-zelda](../entries/lurelin-village-zelda.yml)                           | Lurelin Village              | None                  | No locally declared aid field | Keep name/localization and Notes; no verified naming explanation selected.                            |
| [lynel-zelda](../entries/lynel-zelda.yml)                                               | Lynel                        | None                  | No locally declared aid field | Keep name/localization and Notes; no verified naming explanation selected.                            |
| [magic-rod-zelda](../entries/magic-rod-zelda.yml)                                       | Magic Rod                    | None                  | No locally declared aid field | Keep name/localization and Notes; no verified naming explanation selected.                            |
| [main-quests-zelda](../entries/main-quests-zelda.yml)                                   | Main Quests                  | None                  | No locally declared aid field | Keep name/localization and Notes; no verified naming explanation selected.                            |
| [majestic-general](../entries/majestic-general.yml)                                     | majestic                     | None                  | No locally declared aid field | Leave optional fields empty or retain current aids; no additional explanation selected for this pass. |
| [malevolent-general](../entries/malevolent-general.yml)                                 | malevolent                   | ETY13                 | No locally declared aid field | Applied listed enrichment or derived display; retained all other fields.                              |
| [map-ui](../entries/map-ui.yml)                                                         | map                          | None                  | No locally declared aid field | Leave optional fields empty or retain current aids; no additional explanation selected for this pass. |
| [marbled-gohma-zelda](../entries/marbled-gohma-zelda.yml)                               | Marbled Gohma                | None                  | No locally declared aid field | Keep name/localization and Notes; no verified naming explanation selected.                            |
| [mask-clothing](../entries/mask-clothing.yml)                                           | mask                         | None                  | No locally declared aid field | Leave optional fields empty or retain current aids; no additional explanation selected for this pass. |
| [master-kohga-zelda](../entries/master-kohga-zelda.yml)                                 | Master Kohga                 | None                  | No locally declared aid field | Keep name/localization and Notes; no verified naming explanation selected.                            |
| [master-sword-zelda](../entries/master-sword-zelda.yml)                                 | Master Sword                 | None                  | No locally declared aid field | Keep name/localization and Notes; no verified naming explanation selected.                            |
| [material-general](../entries/material-general.yml)                                     | material                     | None                  | No locally declared aid field | Leave optional fields empty or retain current aids; no additional explanation selected for this pass. |
| [meal-cooking](../entries/meal-cooking.yml)                                             | meal                         | None                  | No locally declared aid field | Leave optional fields empty or retain current aids; no additional explanation selected for this pass. |
| [meat-food](../entries/meat-food.yml)                                                   | meat                         | None                  | No locally declared aid field | Leave optional fields empty or retain current aids; no additional explanation selected for this pass. |
| [melee-combat](../entries/melee-combat.yml)                                             | melee                        | ETY06                 | No locally declared aid field | Applied listed enrichment or derived display; retained all other fields.                              |
| [memory-story](../entries/memory-story.yml)                                             | memory                       | None                  | No locally declared aid field | Leave optional fields empty or retain current aids; no additional explanation selected for this pass. |
| [menu-ui](../entries/menu-ui.yml)                                                       | menu                         | None                  | No locally declared aid field | Leave optional fields empty or retain current aids; no additional explanation selected for this pass. |
| [mighty-bananas-zelda](../entries/mighty-bananas-zelda.yml)                             | Mighty Bananas               | None                  | No locally declared aid field | Keep name/localization and Notes; no verified naming explanation selected.                            |
| [mighty-general](../entries/mighty-general.yml)                                         | mighty                       | None                  | No locally declared aid field | Leave optional fields empty or retain current aids; no additional explanation selected for this pass. |
| [milk-food](../entries/milk-food.yml)                                                   | milk                         | None                  | No locally declared aid field | Leave optional fields empty or retain current aids; no additional explanation selected for this pass. |
| [mine-architecture](../entries/mine-architecture.yml)                                   | mine                         | None                  | No locally declared aid field | Leave optional fields empty or retain current aids; no additional explanation selected for this pass. |
| [miner-general](../entries/miner-general.yml)                                           | miner                        | WF28                  | No locally declared aid field | Applied listed enrichment or derived display; retained all other fields.                              |
| [miners-mask-zelda](../entries/miners-mask-zelda.yml)                                   | Miner's Mask                 | None                  | No locally declared aid field | Keep name/localization and Notes; no verified naming explanation selected.                            |
| [mineru-zelda](../entries/mineru-zelda.yml)                                             | Mineru                       | None                  | No locally declared aid field | Keep name/localization and Notes; no verified naming explanation selected.                            |
| [mirror-general](../entries/mirror-general.yml)                                         | mirror                       | None                  | No locally declared aid field | Leave optional fields empty or retain current aids; no additional explanation selected for this pass. |
| [moblin-zelda](../entries/moblin-zelda.yml)                                             | Moblin                       | None                  | No locally declared aid field | Keep name/localization and Notes; no verified naming explanation selected.                            |
| [mountain-terrain](../entries/mountain-terrain.yml)                                     | mountain                     | None                  | No locally declared aid field | Leave optional fields empty or retain current aids; no additional explanation selected for this pass. |
| [mucktorok-zelda](../entries/mucktorok-zelda.yml)                                       | Mucktorok                    | None                  | No locally declared aid field | Keep name/localization and Notes; no verified naming explanation selected.                            |
| [mysterious-general](../entries/mysterious-general.yml)                                 | mysterious                   | WF34                  | No locally declared aid field | Applied listed enrichment or derived display; retained all other fields.                              |
| [mystery-story](../entries/mystery-story.yml)                                           | mystery                      | None                  | No locally declared aid field | Leave optional fields empty or retain current aids; no additional explanation selected for this pass. |
| [noble-general](../entries/noble-general.yml)                                           | noble                        | None                  | No locally declared aid field | Leave optional fields empty or retain current aids; no additional explanation selected for this pass. |
| [objective-ui](../entries/objective-ui.yml)                                             | objective                    | None                  | No locally declared aid field | Leave optional fields empty or retain current aids; no additional explanation selected for this pass. |
| [octorok-zelda](../entries/octorok-zelda.yml)                                           | Octorok                      | None                  | No locally declared aid field | Keep name/localization and Notes; no verified naming explanation selected.                            |
| [oil-food](../entries/oil-food.yml)                                                     | oil                          | None                  | No locally declared aid field | Leave optional fields empty or retain current aids; no additional explanation selected for this pass. |
| [omelet-cooking](../entries/omelet-cooking.yml)                                         | omelet                       | None                  | No locally declared aid field | Leave optional fields empty or retain current aids; no additional explanation selected for this pass. |
| [one-handed-combat](../entries/one-handed-combat.yml)                                   | one-handed                   | WF39                  | No locally declared aid field | Applied listed enrichment or derived display; retained all other fields.                              |
| [pad-general](../entries/pad-general.yml)                                               | pad                          | None                  | No locally declared aid field | Leave optional fields empty or retain current aids; no additional explanation selected for this pass. |
| [paella-cooking](../entries/paella-cooking.yml)                                         | paella                       | ETY17                 | No locally declared aid field | Applied listed enrichment or derived display; retained all other fields.                              |
| [paraglider-zelda](../entries/paraglider-zelda.yml)                                     | Paraglider                   | None                  | No locally declared aid field | Keep name/localization and Notes; no verified naming explanation selected.                            |
| [passage-architecture](../entries/passage-architecture.yml)                             | passage                      | None                  | No locally declared aid field | Leave optional fields empty or retain current aids; no additional explanation selected for this pass. |
| [paya-zelda](../entries/paya-zelda.yml)                                                 | Paya                         | None                  | No locally declared aid field | Keep name/localization and Notes; no verified naming explanation selected.                            |
| [pepper-vegetable](../entries/pepper-vegetable.yml)                                     | pepper                       | None                  | No locally declared aid field | Leave optional fields empty or retain current aids; no additional explanation selected for this pass. |
| [phantom-ganon-zelda](../entries/phantom-ganon-zelda.yml)                               | Phantom Ganon                | None                  | No locally declared aid field | Keep name/localization and Notes; no verified naming explanation selected.                            |
| [phenomenon-story](../entries/phenomenon-story.yml)                                     | phenomenon                   | None                  | No locally declared aid field | Leave optional fields empty or retain current aids; no additional explanation selected for this pass. |
| [pick-up-general](../entries/pick-up-general.yml)                                       | pick up                      | None                  | No locally declared aid field | Leave optional fields empty or retain current aids; no additional explanation selected for this pass. |
| [plateau-terrain](../entries/plateau-terrain.yml)                                       | plateau                      | None                  | No locally declared aid field | Leave optional fields empty or retain current aids; no additional explanation selected for this pass. |
| [platform-architecture](../entries/platform-architecture.yml)                           | platform                     | None                  | No locally declared aid field | Leave optional fields empty or retain current aids; no additional explanation selected for this pass. |
| [porridge-cooking](../entries/porridge-cooking.yml)                                     | porridge                     | None                  | No locally declared aid field | Leave optional fields empty or retain current aids; no additional explanation selected for this pass. |
| [portable-pot-zelda](../entries/portable-pot-zelda.yml)                                 | Portable Pot                 | None                  | No locally declared aid field | Keep name/localization and Notes; no verified naming explanation selected.                            |
| [possessed-combat](../entries/possessed-combat.yml)                                     | possessed                    | None                  | No locally declared aid field | Leave optional fields empty or retain current aids; no additional explanation selected for this pass. |
| [pot-general](../entries/pot-general.yml)                                               | pot                          | None                  | No locally declared aid field | Leave optional fields empty or retain current aids; no additional explanation selected for this pass. |
| [pouch-ui](../entries/pouch-ui.yml)                                                     | pouch                        | R22                   | No locally declared aid field | Applied listed enrichment or derived display; retained all other fields.                              |
| [powerful-general](../entries/powerful-general.yml)                                     | powerful                     | WF11                  | No locally declared aid field | Applied listed enrichment or derived display; retained all other fields.                              |
| [prime-food](../entries/prime-food.yml)                                                 | prime                        | None                  | No locally declared aid field | Leave optional fields empty or retain current aids; no additional explanation selected for this pass. |
| [prince-general](../entries/prince-general.yml)                                         | prince                       | None                  | No locally declared aid field | Leave optional fields empty or retain current aids; no additional explanation selected for this pass. |
| [princess-general](../entries/princess-general.yml)                                     | princess                     | None                  | No locally declared aid field | Leave optional fields empty or retain current aids; no additional explanation selected for this pass. |
| [princess-zelda-zelda](../entries/princess-zelda-zelda.yml)                             | Princess Zelda               | None                  | No locally declared aid field | Keep name/localization and Notes; no verified naming explanation selected.                            |
| [pristine-general](../entries/pristine-general.yml)                                     | pristine                     | N09                   | No locally declared aid field | Applied listed enrichment or derived display; retained all other fields.                              |
| [profile-ui](../entries/profile-ui.yml)                                                 | profile                      | None                  | No locally declared aid field | Leave optional fields empty or retain current aids; no additional explanation selected for this pass. |
| [prophecy-story](../entries/prophecy-story.yml)                                         | prophecy                     | None                  | No locally declared aid field | Leave optional fields empty or retain current aids; no additional explanation selected for this pass. |
| [puffshroom-zelda](../entries/puffshroom-zelda.yml)                                     | Puffshroom                   | None                  | No locally declared aid field | Keep name/localization and Notes; no verified naming explanation selected.                            |
| [purah-pad-zelda](../entries/purah-pad-zelda.yml)                                       | Purah Pad                    | None                  | No locally declared aid field | Keep name/localization and Notes; no verified naming explanation selected.                            |
| [purah-zelda](../entries/purah-zelda.yml)                                               | Purah                        | None                  | No locally declared aid field | Keep name/localization and Notes; no verified naming explanation selected.                            |
| [queen-general](../entries/queen-general.yml)                                           | queen                        | None                  | No locally declared aid field | Leave optional fields empty or retain current aids; no additional explanation selected for this pass. |
| [queen-gibdo-zelda](../entries/queen-gibdo-zelda.yml)                                   | Queen Gibdo                  | None                  | No locally declared aid field | Keep name/localization and Notes; no verified naming explanation selected.                            |
| [quest-game](../entries/quest-game.yml)                                                 | quest                        | None                  | No locally declared aid field | Leave optional fields empty or retain current aids; no additional explanation selected for this pass. |
| [radish-food](../entries/radish-food.yml)                                               | radish                       | None                  | No locally declared aid field | Leave optional fields empty or retain current aids; no additional explanation selected for this pass. |
| [ranged-combat](../entries/ranged-combat.yml)                                           | ranged                       | None                  | No locally declared aid field | Leave optional fields empty or retain current aids; no additional explanation selected for this pass. |
| [rauru-zelda](../entries/rauru-zelda.yml)                                               | Rauru                        | None                  | No locally declared aid field | Keep name/localization and Notes; no verified naming explanation selected.                            |
| [raw-food](../entries/raw-food.yml)                                                     | raw                          | None                  | No locally declared aid field | Leave optional fields empty or retain current aids; no additional explanation selected for this pass. |
| [raw-gourmet-meat-zelda](../entries/raw-gourmet-meat-zelda.yml)                         | Raw Gourmet Meat             | None                  | No locally declared aid field | Keep name/localization and Notes; no verified naming explanation selected.                            |
| [raw-prime-meat-zelda](../entries/raw-prime-meat-zelda.yml)                             | Raw Prime Meat               | None                  | No locally declared aid field | Keep name/localization and Notes; no verified naming explanation selected.                            |
| [recall-general](../entries/recall-general.yml)                                         | recall                       | None                  | No locally declared aid field | Leave optional fields empty or retain current aids; no additional explanation selected for this pass. |
| [recharge-general](../entries/recharge-general.yml)                                     | recharge                     | WF05                  | No locally declared aid field | Applied listed enrichment or derived display; retained all other fields.                              |
| [recipe-cooking](../entries/recipe-cooking.yml)                                         | recipe                       | ETY05, R26            | No locally declared aid field | Applied listed enrichment or derived display; retained all other fields.                              |
| [reckless-general](../entries/reckless-general.yml)                                     | reckless                     | N07, R12              | No locally declared aid field | Applied listed enrichment or derived display; retained all other fields.                              |
| [recover-story](../entries/recover-story.yml)                                           | recover                      | None                  | existing declared association | Leave optional fields empty or retain current aids; no additional explanation selected for this pass. |
| [refinery-architecture](../entries/refinery-architecture.yml)                           | refinery                     | WF27                  | No locally declared aid field | Applied listed enrichment or derived display; retained all other fields.                              |
| [regional-general](../entries/regional-general.yml)                                     | regional                     | WF14                  | No locally declared aid field | Applied listed enrichment or derived display; retained all other fields.                              |
| [resilient-combat](../entries/resilient-combat.yml)                                     | resilient                    | ETY03, N02, R15       | No locally declared aid field | Applied listed enrichment or derived display; retained all other fields.                              |
| [resistance-general](../entries/resistance-general.yml)                                 | resistance                   | None                  | No locally declared aid field | Leave optional fields empty or retain current aids; no additional explanation selected for this pass. |
| [ridge-terrain](../entries/ridge-terrain.yml)                                           | ridge                        | R17                   | No locally declared aid field | Applied listed enrichment or derived display; retained all other fields.                              |
| [riju-zelda](../entries/riju-zelda.yml)                                                 | Riju                         | None                  | No locally declared aid field | Keep name/localization and Notes; no verified naming explanation selected.                            |
| [ring-general](../entries/ring-general.yml)                                             | ring                         | None                  | No locally declared aid field | Leave optional fields empty or retain current aids; no additional explanation selected for this pass. |
| [risotto-cooking](../entries/risotto-cooking.yml)                                       | risotto                      | None                  | No locally declared aid field | Leave optional fields empty or retain current aids; no additional explanation selected for this pass. |
| [rito-village-zelda](../entries/rito-village-zelda.yml)                                 | Rito Village                 | None                  | No locally declared aid field | Keep name/localization and Notes; no verified naming explanation selected.                            |
| [rito-zelda](../entries/rito-zelda.yml)                                                 | Rito                         | None                  | No locally declared aid field | Keep name/localization and Notes; no verified naming explanation selected.                            |
| [river-terrain](../entries/river-terrain.yml)                                           | river                        | None                  | No locally declared aid field | Leave optional fields empty or retain current aids; no additional explanation selected for this pass. |
| [robbie-zelda](../entries/robbie-zelda.yml)                                             | Robbie                       | None                  | No locally declared aid field | Keep name/localization and Notes; no verified naming explanation selected.                            |
| [robe-clothing](../entries/robe-clothing.yml)                                           | robe                         | None                  | No locally declared aid field | Leave optional fields empty or retain current aids; no additional explanation selected for this pass. |
| [rocket-engineering](../entries/rocket-engineering.yml)                                 | rocket                       | None                  | No locally declared aid field | Leave optional fields empty or retain current aids; no additional explanation selected for this pass. |
| [rod-general](../entries/rod-general.yml)                                               | rod                          | None                  | No locally declared aid field | Leave optional fields empty or retain current aids; no additional explanation selected for this pass. |
| [root-general](../entries/root-general.yml)                                             | root                         | None                  | No locally declared aid field | Leave optional fields empty or retain current aids; no additional explanation selected for this pass. |
| [rotate-general](../entries/rotate-general.yml)                                         | rotate                       | None                  | No locally declared aid field | Leave optional fields empty or retain current aids; no additional explanation selected for this pass. |
| [royal-broadsword-zelda](../entries/royal-broadsword-zelda.yml)                         | Royal Broadsword             | None                  | No locally declared aid field | Keep name/localization and Notes; no verified naming explanation selected.                            |
| [royal-general](../entries/royal-general.yml)                                           | royal                        | None                  | No locally declared aid field | Leave optional fields empty or retain current aids; no additional explanation selected for this pass. |
| [ruins-architecture](../entries/ruins-architecture.yml)                                 | ruins                        | None                  | No locally declared aid field | Leave optional fields empty or retain current aids; no additional explanation selected for this pass. |
| [rupee-game](../entries/rupee-game.yml)                                                 | rupee                        | ETY18                 | No locally declared aid field | Applied listed enrichment or derived display; retained all other fields.                              |
| [rusty-general](../entries/rusty-general.yml)                                           | rusty                        | None                  | No locally declared aid field | Leave optional fields empty or retain current aids; no additional explanation selected for this pass. |
| [ruthless-general](../entries/ruthless-general.yml)                                     | ruthless                     | WF10                  | No locally declared aid field | Applied listed enrichment or derived display; retained all other fields.                              |
| [sacred-general](../entries/sacred-general.yml)                                         | sacred                       | None                  | No locally declared aid field | Leave optional fields empty or retain current aids; no additional explanation selected for this pass. |
| [sacrifice-story](../entries/sacrifice-story.yml)                                       | sacrifice                    | ETY12                 | No locally declared aid field | Applied listed enrichment or derived display; retained all other fields.                              |
| [sage-general](../entries/sage-general.yml)                                             | sage                         | None                  | No locally declared aid field | Leave optional fields empty or retain current aids; no additional explanation selected for this pass. |
| [salmon-food](../entries/salmon-food.yml)                                               | salmon                       | None                  | No locally declared aid field | Leave optional fields empty or retain current aids; no additional explanation selected for this pass. |
| [salt-food](../entries/salt-food.yml)                                                   | salt                         | None                  | No locally declared aid field | Leave optional fields empty or retain current aids; no additional explanation selected for this pass. |
| [saute-cooking](../entries/saute-cooking.yml)                                           | sauté                        | ETY16                 | No locally declared aid field | Applied listed enrichment or derived display; retained all other fields.                              |
| [savage-general](../entries/savage-general.yml)                                         | savage                       | None                  | No locally declared aid field | Leave optional fields empty or retain current aids; no additional explanation selected for this pass. |
| [scimitar-of-the-seven-zelda](../entries/scimitar-of-the-seven-zelda.yml)               | Scimitar of the Seven        | None                  | No locally declared aid field | Keep name/localization and Notes; no verified naming explanation selected.                            |
| [scimitar-weapon](../entries/scimitar-weapon.yml)                                       | scimitar                     | None                  | No locally declared aid field | Leave optional fields empty or retain current aids; no additional explanation selected for this pass. |
| [sear-cooking](../entries/sear-cooking.yml)                                             | sear                         | N14, R20              | No locally declared aid field | Applied listed enrichment or derived display; retained all other fields.                              |
| [secret-stone-zelda](../entries/secret-stone-zelda.yml)                                 | Secret Stone                 | None                  | No locally declared aid field | Keep name/localization and Notes; no verified naming explanation selected.                            |
| [secret-story](../entries/secret-story.yml)                                             | secret                       | None                  | No locally declared aid field | Leave optional fields empty or retain current aids; no additional explanation selected for this pass. |
| [seed-general](../entries/seed-general.yml)                                             | seed                         | None                  | No locally declared aid field | Leave optional fields empty or retain current aids; no additional explanation selected for this pass. |
| [seized-construct-zelda](../entries/seized-construct-zelda.yml)                         | Seized Construct             | None                  | No locally declared aid field | Keep name/localization and Notes; no verified naming explanation selected.                            |
| [select-general](../entries/select-general.yml)                                         | select                       | None                  | No locally declared aid field | Leave optional fields empty or retain current aids; no additional explanation selected for this pass. |
| [selfless-general](../entries/selfless-general.yml)                                     | selfless                     | WF09                  | No locally declared aid field | Applied listed enrichment or derived display; retained all other fields.                              |
| [settlement-architecture](../entries/settlement-architecture.yml)                       | settlement                   | WF31                  | No locally declared aid field | Applied listed enrichment or derived display; retained all other fields.                              |
| [sheikah-zelda](../entries/sheikah-zelda.yml)                                           | Sheikah                      | None                  | No locally declared aid field | Keep name/localization and Notes; no verified naming explanation selected.                            |
| [shield-weapon](../entries/shield-weapon.yml)                                           | shield                       | None                  | No locally declared aid field | Leave optional fields empty or retain current aids; no additional explanation selected for this pass. |
| [shirt-clothing](../entries/shirt-clothing.yml)                                         | shirt                        | None                  | No locally declared aid field | Leave optional fields empty or retain current aids; no additional explanation selected for this pass. |
| [shock-emitter-zelda](../entries/shock-emitter-zelda.yml)                               | Shock Emitter                | WF44                  | No locally declared aid field | Applied listed enrichment or derived display; retained all other fields.                              |
| [shock-resistance-combat](../entries/shock-resistance-combat.yml)                       | shock resistance             | None                  | No locally declared aid field | Leave optional fields empty or retain current aids; no additional explanation selected for this pass. |
| [shrine-architecture](../entries/shrine-architecture.yml)                               | shrine                       | None                  | No locally declared aid field | Leave optional fields empty or retain current aids; no additional explanation selected for this pass. |
| [shrine-quests-zelda](../entries/shrine-quests-zelda.yml)                               | Shrine Quests                | None                  | No locally declared aid field | Keep name/localization and Notes; no verified naming explanation selected.                            |
| [side-adventures-zelda](../entries/side-adventures-zelda.yml)                           | Side Adventures              | None                  | No locally declared aid field | Keep name/localization and Notes; no verified naming explanation selected.                            |
| [side-quests-zelda](../entries/side-quests-zelda.yml)                                   | Side Quests                  | None                  | No locally declared aid field | Keep name/localization and Notes; no verified naming explanation selected.                            |
| [sidon-zelda](../entries/sidon-zelda.yml)                                               | Sidon                        | None                  | No locally declared aid field | Keep name/localization and Notes; no verified naming explanation selected.                            |
| [silent-princess-zelda](../entries/silent-princess-zelda.yml)                           | Silent Princess              | None                  | No locally declared aid field | Keep name/localization and Notes; no verified naming explanation selected.                            |
| [simmer-cooking](../entries/simmer-cooking.yml)                                         | simmer                       | N13, R21              | No locally declared aid field | Applied listed enrichment or derived display; retained all other fields.                              |
| [skewer-cooking](../entries/skewer-cooking.yml)                                         | skewer                       | None                  | No locally declared aid field | Leave optional fields empty or retain current aids; no additional explanation selected for this pass. |
| [skilled-general](../entries/skilled-general.yml)                                       | skilled                      | WF12                  | No locally declared aid field | Applied listed enrichment or derived display; retained all other fields.                              |
| [sky-island-zelda](../entries/sky-island-zelda.yml)                                     | sky island                   | None                  | existing Word Formation       | Keep the existing construction; no additional field justified in this pass.                           |
| [sky-terrain](../entries/sky-terrain.yml)                                               | sky                          | None                  | No locally declared aid field | Leave optional fields empty or retain current aids; no additional explanation selected for this pass. |
| [skyshroom-zelda](../entries/skyshroom-zelda.yml)                                       | Skyshroom                    | None                  | No locally declared aid field | Keep name/localization and Notes; no verified naming explanation selected.                            |
| [skyview-tower-zelda](../entries/skyview-tower-zelda.yml)                               | Skyview Tower                | WF46                  | No locally declared aid field | Applied listed enrichment or derived display; retained all other fields.                              |
| [sled-engineering](../entries/sled-engineering.yml)                                     | sled                         | None                  | No locally declared aid field | Leave optional fields empty or retain current aids; no additional explanation selected for this pass. |
| [sleeve-clothing](../entries/sleeve-clothing.yml)                                       | sleeve                       | None                  | No locally declared aid field | Leave optional fields empty or retain current aids; no additional explanation selected for this pass. |
| [slip-general](../entries/slip-general.yml)                                             | slip                         | None                  | No locally declared aid field | Leave optional fields empty or retain current aids; no additional explanation selected for this pass. |
| [slip-resistance-zelda](../entries/slip-resistance-zelda.yml)                           | Slip Resistance              | None                  | No locally declared aid field | Keep name/localization and Notes; no verified naming explanation selected.                            |
| [slope-terrain](../entries/slope-terrain.yml)                                           | slope                        | None                  | No locally declared aid field | Leave optional fields empty or retain current aids; no additional explanation selected for this pass. |
| [small-wheel-zelda](../entries/small-wheel-zelda.yml)                                   | Small Wheel                  | None                  | No locally declared aid field | Keep name/localization and Notes; no verified naming explanation selected.                            |
| [sneaky-general](../entries/sneaky-general.yml)                                         | sneaky                       | None                  | No locally declared aid field | Leave optional fields empty or retain current aids; no additional explanation selected for this pass. |
| [snowquill-tunic-zelda](../entries/snowquill-tunic-zelda.yml)                           | Snowquill Tunic              | None                  | No locally declared aid field | Keep name/localization and Notes; no verified naming explanation selected.                            |
| [soldier-construct-zelda](../entries/soldier-construct-zelda.yml)                       | Soldier Construct            | None                  | No locally declared aid field | Keep name/localization and Notes; no verified naming explanation selected.                            |
| [soldier-general](../entries/soldier-general.yml)                                       | soldier                      | None                  | No locally declared aid field | Leave optional fields empty or retain current aids; no additional explanation selected for this pass. |
| [soldiers-broadsword-zelda](../entries/soldiers-broadsword-zelda.yml)                   | Soldier's Broadsword         | None                  | No locally declared aid field | Keep name/localization and Notes; no verified naming explanation selected.                            |
| [sonia-zelda](../entries/sonia-zelda.yml)                                               | Sonia                        | None                  | No locally declared aid field | Keep name/localization and Notes; no verified naming explanation selected.                            |
| [sort-general](../entries/sort-general.yml)                                             | sort                         | None                  | No locally declared aid field | Leave optional fields empty or retain current aids; no additional explanation selected for this pass. |
| [spaulder-clothing](../entries/spaulder-clothing.yml)                                   | spaulder                     | None                  | No locally declared aid field | Leave optional fields empty or retain current aids; no additional explanation selected for this pass. |
| [spear-weapon](../entries/spear-weapon.yml)                                             | spear                        | None                  | No locally declared aid field | Leave optional fields empty or retain current aids; no additional explanation selected for this pass. |
| [special-controls-zelda](../entries/special-controls-zelda.yml)                         | Special Controls             | None                  | No locally declared aid field | Keep name/localization and Notes; no verified naming explanation selected.                            |
| [speed-up-zelda](../entries/speed-up-zelda.yml)                                         | Speed Up                     | None                  | No locally declared aid field | Keep name/localization and Notes; no verified naming explanation selected.                            |
| [spice-food](../entries/spice-food.yml)                                                 | spice                        | None                  | No locally declared aid field | Leave optional fields empty or retain current aids; no additional explanation selected for this pass. |
| [spicy-general](../entries/spicy-general.yml)                                           | spicy                        | None                  | No locally declared aid field | Leave optional fields empty or retain current aids; no additional explanation selected for this pass. |
| [spicy-pepper-zelda](../entries/spicy-pepper-zelda.yml)                                 | Spicy Pepper                 | None                  | No locally declared aid field | Keep name/localization and Notes; no verified naming explanation selected.                            |
| [spirit-story](../entries/spirit-story.yml)                                             | spirit                       | None                  | No locally declared aid field | Leave optional fields empty or retain current aids; no additional explanation selected for this pass. |
| [spirit-temple-zelda](../entries/spirit-temple-zelda.yml)                               | Spirit Temple                | None                  | No locally declared aid field | Keep name/localization and Notes; no verified naming explanation selected.                            |
| [spring-mechanical](../entries/spring-mechanical.yml)                                   | spring                       | None                  | No locally declared aid field | Leave optional fields empty or retain current aids; no additional explanation selected for this pass. |
| [spring-water](../entries/spring-water.yml)                                             | spring                       | N20                   | existing declared association | Applied listed enrichment or derived display; retained all other fields.                              |
| [sprint-general](../entries/sprint-general.yml)                                         | sprint                       | None                  | No locally declared aid field | Leave optional fields empty or retain current aids; no additional explanation selected for this pass. |
| [stabilizer-engineering](../entries/stabilizer-engineering.yml)                         | stabilizer                   | WF25                  | No locally declared aid field | Applied listed enrichment or derived display; retained all other fields.                              |
| [stable-architecture](../entries/stable-architecture.yml)                               | stable                       | None                  | No locally declared aid field | Leave optional fields empty or retain current aids; no additional explanation selected for this pass. |
| [staircase-architecture](../entries/staircase-architecture.yml)                         | staircase                    | WF21                  | No locally declared aid field | Applied listed enrichment or derived display; retained all other fields.                              |
| [stake-engineering](../entries/stake-engineering.yml)                                   | stake                        | None                  | No locally declared aid field | Leave optional fields empty or retain current aids; no additional explanation selected for this pass. |
| [stamina-general](../entries/stamina-general.yml)                                       | stamina                      | ETY01, N01            | existing declared association | Applied listed enrichment or derived display; retained all other fields.                              |
| [statue-architecture](../entries/statue-architecture.yml)                               | statue                       | None                  | No locally declared aid field | Leave optional fields empty or retain current aids; no additional explanation selected for this pass. |
| [stealth-mask-zelda](../entries/stealth-mask-zelda.yml)                                 | Stealth Mask                 | None                  | No locally declared aid field | Keep name/localization and Notes; no verified naming explanation selected.                            |
| [stealth-up-zelda](../entries/stealth-up-zelda.yml)                                     | Stealth Up                   | None                  | No locally declared aid field | Keep name/localization and Notes; no verified naming explanation selected.                            |
| [steam-cooking](../entries/steam-cooking.yml)                                           | steam                        | N15                   | No locally declared aid field | Applied listed enrichment or derived display; retained all other fields.                              |
| [steer-general](../entries/steer-general.yml)                                           | steer                        | None                  | No locally declared aid field | Leave optional fields empty or retain current aids; no additional explanation selected for this pass. |
| [steering-stick-zelda](../entries/steering-stick-zelda.yml)                             | Steering Stick               | None                  | existing Word Formation       | Keep the existing construction; no additional field justified in this pass.                           |
| [stew-cooking](../entries/stew-cooking.yml)                                             | stew                         | R21                   | No locally declared aid field | Applied listed enrichment or derived display; retained all other fields.                              |
| [stick-general](../entries/stick-general.yml)                                           | stick                        | None                  | No locally declared aid field | Leave optional fields empty or retain current aids; no additional explanation selected for this pass. |
| [stone-general](../entries/stone-general.yml)                                           | stone                        | None                  | No locally declared aid field | Leave optional fields empty or retain current aids; no additional explanation selected for this pass. |
| [stone-talus-zelda](../entries/stone-talus-zelda.yml)                                   | Stone Talus                  | None                  | No locally declared aid field | Keep name/localization and Notes; no verified naming explanation selected.                            |
| [sturdy-general](../entries/sturdy-general.yml)                                         | sturdy                       | N03, R15              | No locally declared aid field | Applied listed enrichment or derived display; retained all other fields.                              |
| [sugar-food](../entries/sugar-food.yml)                                                 | sugar                        | None                  | No locally declared aid field | Leave optional fields empty or retain current aids; no additional explanation selected for this pass. |
| [summit-terrain](../entries/summit-terrain.yml)                                         | summit                       | N10, R17              | No locally declared aid field | Applied listed enrichment or derived display; retained all other fields.                              |
| [summon-general](../entries/summon-general.yml)                                         | summon                       | None                  | No locally declared aid field | Leave optional fields empty or retain current aids; no additional explanation selected for this pass. |
| [sundelion-zelda](../entries/sundelion-zelda.yml)                                       | Sundelion                    | None                  | No locally declared aid field | Keep name/localization and Notes; no verified naming explanation selected.                            |
| [sunny-general](../entries/sunny-general.yml)                                           | sunny                        | None                  | No locally declared aid field | Leave optional fields empty or retain current aids; no additional explanation selected for this pass. |
| [surface-terrain](../entries/surface-terrain.yml)                                       | surface                      | R29                   | No locally declared aid field | Applied listed enrichment or derived display; retained all other fields.                              |
| [susceptible-combat](../entries/susceptible-combat.yml)                                 | susceptible                  | N05, R14              | No locally declared aid field | Applied listed enrichment or derived display; retained all other fields.                              |
| [swamp-terrain](../entries/swamp-terrain.yml)                                           | swamp                        | None                  | No locally declared aid field | Leave optional fields empty or retain current aids; no additional explanation selected for this pass. |
| [sword-weapon](../entries/sword-weapon.yml)                                             | sword                        | None                  | No locally declared aid field | Leave optional fields empty or retain current aids; no additional explanation selected for this pass. |
| [system-ui](../entries/system-ui.yml)                                                   | system                       | None                  | No locally declared aid field | Leave optional fields empty or retain current aids; no additional explanation selected for this pass. |
| [tarrey-town-zelda](../entries/tarrey-town-zelda.yml)                                   | Tarrey Town                  | None                  | No locally declared aid field | Keep name/localization and Notes; no verified naming explanation selected.                            |
| [tear-general](../entries/tear-general.yml)                                             | tear                         | N18                   | No locally declared aid field | Applied listed enrichment or derived display; retained all other fields.                              |
| [teba-zelda](../entries/teba-zelda.yml)                                                 | Teba                         | None                  | No locally declared aid field | Keep name/localization and Notes; no verified naming explanation selected.                            |
| [temple-architecture](../entries/temple-architecture.yml)                               | temple                       | None                  | No locally declared aid field | Leave optional fields empty or retain current aids; no additional explanation selected for this pass. |
| [temple-of-time-zelda](../entries/temple-of-time-zelda.yml)                             | Temple of Time               | None                  | No locally declared aid field | Keep name/localization and Notes; no verified naming explanation selected.                            |
| [the-depths-zelda](../entries/the-depths-zelda.yml)                                     | The Depths                   | R29                   | No locally declared aid field | Applied listed enrichment or derived display; retained all other fields.                              |
| [throw-general](../entries/throw-general.yml)                                           | throw                        | None                  | No locally declared aid field | Leave optional fields empty or retain current aids; no additional explanation selected for this pass. |
| [tights-clothing](../entries/tights-clothing.yml)                                       | tights                       | None                  | No locally declared aid field | Leave optional fields empty or retain current aids; no additional explanation selected for this pass. |
| [time-bomb-zelda](../entries/time-bomb-zelda.yml)                                       | Time Bomb                    | None                  | No locally declared aid field | Keep name/localization and Notes; no verified naming explanation selected.                            |
| [timid-general](../entries/timid-general.yml)                                           | timid                        | None                  | No locally declared aid field | Leave optional fields empty or retain current aids; no additional explanation selected for this pass. |
| [tomato-food](../entries/tomato-food.yml)                                               | tomato                       | None                  | No locally declared aid field | Leave optional fields empty or retain current aids; no additional explanation selected for this pass. |
| [torch-weapon](../entries/torch-weapon.yml)                                             | torch                        | None                  | No locally declared aid field | Leave optional fields empty or retain current aids; no additional explanation selected for this pass. |
| [tough-general](../entries/tough-general.yml)                                           | tough                        | None                  | No locally declared aid field | Leave optional fields empty or retain current aids; no additional explanation selected for this pass. |
| [tower-architecture](../entries/tower-architecture.yml)                                 | tower                        | None                  | No locally declared aid field | Leave optional fields empty or retain current aids; no additional explanation selected for this pass. |
| [trail-story](../entries/trail-story.yml)                                               | trail                        | None                  | No locally declared aid field | Leave optional fields empty or retain current aids; no additional explanation selected for this pass. |
| [traveler-general](../entries/traveler-general.yml)                                     | traveler                     | WF29                  | No locally declared aid field | Applied listed enrichment or derived display; retained all other fields.                              |
| [travelers-sword-zelda](../entries/travelers-sword-zelda.yml)                           | Traveler's Sword             | None                  | No locally declared aid field | Keep name/localization and Notes; no verified naming explanation selected.                            |
| [tree-terrain](../entries/tree-terrain.yml)                                             | tree                         | None                  | No locally declared aid field | Leave optional fields empty or retain current aids; no additional explanation selected for this pass. |
| [trident-weapon](../entries/trident-weapon.yml)                                         | trident                      | ETY11                 | No locally declared aid field | Applied listed enrichment or derived display; retained all other fields.                              |
| [trousers-clothing](../entries/trousers-clothing.yml)                                   | trousers                     | None                  | No locally declared aid field | Leave optional fields empty or retain current aids; no additional explanation selected for this pass. |
| [truffle-food](../entries/truffle-food.yml)                                             | truffle                      | None                  | No locally declared aid field | Leave optional fields empty or retain current aids; no additional explanation selected for this pass. |
| [tulin-zelda](../entries/tulin-zelda.yml)                                               | Tulin                        | None                  | No locally declared aid field | Keep name/localization and Notes; no verified naming explanation selected.                            |
| [tunic-clothing](../entries/tunic-clothing.yml)                                         | tunic                        | None                  | No locally declared aid field | Leave optional fields empty or retain current aids; no additional explanation selected for this pass. |
| [two-handed-combat](../entries/two-handed-combat.yml)                                   | two-handed                   | WF40                  | No locally declared aid field | Applied listed enrichment or derived display; retained all other fields.                              |
| [ultrahand-zelda](../entries/ultrahand-zelda.yml)                                       | Ultrahand                    | WF02                  | No locally declared aid field | Applied listed enrichment or derived display; retained all other fields.                              |
| [undead-combat](../entries/undead-combat.yml)                                           | undead                       | WF07                  | No locally declared aid field | Applied listed enrichment or derived display; retained all other fields.                              |
| [underground-general](../entries/underground-general.yml)                               | underground                  | WF06                  | No locally declared aid field | Applied listed enrichment or derived display; retained all other fields.                              |
| [unfuse-general](../entries/unfuse-general.yml)                                         | unfuse                       | WF03                  | No locally declared aid field | Applied listed enrichment or derived display; retained all other fields.                              |
| [unlock-general](../entries/unlock-general.yml)                                         | unlock                       | WF04                  | No locally declared aid field | Applied listed enrichment or derived display; retained all other fields.                              |
| [upgrade-general](../entries/upgrade-general.yml)                                       | upgrade                      | None                  | No locally declared aid field | Leave optional fields empty or retain current aids; no additional explanation selected for this pass. |
| [upheaval-story](../entries/upheaval-story.yml)                                         | upheaval                     | None                  | No locally declared aid field | Leave optional fields empty or retain current aids; no additional explanation selected for this pass. |
| [valley-terrain](../entries/valley-terrain.yml)                                         | valley                       | None                  | No locally declared aid field | Leave optional fields empty or retain current aids; no additional explanation selected for this pass. |
| [view-general](../entries/view-general.yml)                                             | view                         | None                  | No locally declared aid field | Leave optional fields empty or retain current aids; no additional explanation selected for this pass. |
| [village-architecture](../entries/village-architecture.yml)                             | village                      | None                  | No locally declared aid field | Leave optional fields empty or retain current aids; no additional explanation selected for this pass. |
| [volcanic-combat](../entries/volcanic-combat.yml)                                       | volcanic                     | None                  | No locally declared aid field | Leave optional fields empty or retain current aids; no additional explanation selected for this pass. |
| [vulnerable-combat](../entries/vulnerable-combat.yml)                                   | vulnerable                   | ETY09, N04, R14       | No locally declared aid field | Applied listed enrichment or derived display; retained all other fields.                              |
| [waistguard-clothing](../entries/waistguard-clothing.yml)                               | waistguard                   | WF17                  | No locally declared aid field | Applied listed enrichment or derived display; retained all other fields.                              |
| [warp-general](../entries/warp-general.yml)                                             | warp                         | None                  | No locally declared aid field | Leave optional fields empty or retain current aids; no additional explanation selected for this pass. |
| [warrior-general](../entries/warrior-general.yml)                                       | warrior                      | None                  | No locally declared aid field | Leave optional fields empty or retain current aids; no additional explanation selected for this pass. |
| [water-general](../entries/water-general.yml)                                           | water                        | None                  | No locally declared aid field | Leave optional fields empty or retain current aids; no additional explanation selected for this pass. |
| [water-temple-zelda](../entries/water-temple-zelda.yml)                                 | Water Temple                 | None                  | No locally declared aid field | Keep name/localization and Notes; no verified naming explanation selected.                            |
| [waterfall-terrain](../entries/waterfall-terrain.yml)                                   | waterfall                    | WF20                  | No locally declared aid field | Applied listed enrichment or derived display; retained all other fields.                              |
| [weapon-weapon](../entries/weapon-weapon.yml)                                           | weapon                       | None                  | No locally declared aid field | Leave optional fields empty or retain current aids; no additional explanation selected for this pass. |
| [well-architecture](../entries/well-architecture.yml)                                   | well                         | None                  | No locally declared aid field | Leave optional fields empty or retain current aids; no additional explanation selected for this pass. |
| [wetlands-terrain](../entries/wetlands-terrain.yml)                                     | wetlands                     | WF19                  | No locally declared aid field | Applied listed enrichment or derived display; retained all other fields.                              |
| [wheat-food](../entries/wheat-food.yml)                                                 | wheat                        | None                  | No locally declared aid field | Leave optional fields empty or retain current aids; no additional explanation selected for this pass. |
| [wheel-engineering](../entries/wheel-engineering.yml)                                   | wheel                        | None                  | No locally declared aid field | Leave optional fields empty or retain current aids; no additional explanation selected for this pass. |
| [wield-story](../entries/wield-story.yml)                                               | wield                        | None                  | No locally declared aid field | Leave optional fields empty or retain current aids; no additional explanation selected for this pass. |
| [wind-temple-zelda](../entries/wind-temple-zelda.yml)                                   | Wind Temple                  | None                  | No locally declared aid field | Keep name/localization and Notes; no verified naming explanation selected.                            |
| [wind-terrain](../entries/wind-terrain.yml)                                             | wind                         | None                  | No locally declared aid field | Leave optional fields empty or retain current aids; no additional explanation selected for this pass. |
| [wing-engineering](../entries/wing-engineering.yml)                                     | wing                         | None                  | No locally declared aid field | Leave optional fields empty or retain current aids; no additional explanation selected for this pass. |
| [wise-general](../entries/wise-general.yml)                                             | wise                         | None                  | No locally declared aid field | Leave optional fields empty or retain current aids; no additional explanation selected for this pass. |
| [wizzrobe-zelda](../entries/wizzrobe-zelda.yml)                                         | Wizzrobe                     | None                  | No locally declared aid field | Keep name/localization and Notes; no verified naming explanation selected.                            |
| [yiga-clan-hideout-zelda](../entries/yiga-clan-hideout-zelda.yml)                       | Yiga Clan Hideout            | None                  | No locally declared aid field | Keep name/localization and Notes; no verified naming explanation selected.                            |
| [yiga-clan-zelda](../entries/yiga-clan-zelda.yml)                                       | Yiga Clan                    | None                  | No locally declared aid field | Keep name/localization and Notes; no verified naming explanation selected.                            |
| [yona-zelda](../entries/yona-zelda.yml)                                                 | Yona                         | None                  | No locally declared aid field | Keep name/localization and Notes; no verified naming explanation selected.                            |
| [yunobo-zelda](../entries/yunobo-zelda.yml)                                             | Yunobo                       | None                  | No locally declared aid field | Keep name/localization and Notes; no verified naming explanation selected.                            |
| [zonai-device-zelda](../entries/zonai-device-zelda.yml)                                 | Zonai Device                 | None                  | No locally declared aid field | Keep name/localization and Notes; no verified naming explanation selected.                            |
| [zonai-zelda](../entries/zonai-zelda.yml)                                               | Zonai                        | None                  | No locally declared aid field | Keep name/localization and Notes; no verified naming explanation selected.                            |
| [zonaite-zelda](../entries/zonaite-zelda.yml)                                           | Zonaite                      | None                  | No locally declared aid field | Keep name/localization and Notes; no verified naming explanation selected.                            |
| [zora-armor-zelda](../entries/zora-armor-zelda.yml)                                     | Zora Armor                   | None                  | No locally declared aid field | Keep name/localization and Notes; no verified naming explanation selected.                            |
| [zora-zelda](../entries/zora-zelda.yml)                                                 | Zora                         | None                  | No locally declared aid field | Keep name/localization and Notes; no verified naming explanation selected.                            |
| [zoras-domain-zelda](../entries/zoras-domain-zelda.yml)                                 | Zora's Domain                | None                  | No locally declared aid field | Keep name/localization and Notes; no verified naming explanation selected.                            |

## Approval and implementation boundaries

You can approve or revise individual IDs or explicit ranges. Suggested review units are WF01–WF24, WF25–WF48, ETY01–ETY18, N01–N20, and R10–R29. These are review units rather than a required implementation sequence. Text and relationship decisions remain separate.

The approval record currently contains:

| Group                                        | Status                    | Authorized scope                                                                                              |
| :------------------------------------------- | :------------------------ | :------------------------------------------------------------------------------------------------------------ |
| WF01–WF48                                    | Approved; implemented     | Exact Word Formation values and source-line appends; WF01 additionally includes its stated one-paragraph move |
| ETY01–ETY18                                  | Approved; implemented     | Exact historical-origin values and source-line appends                                                        |
| N01–N20                                      | Approved; implemented     | Exact Mnemonic values and source-line appends; no inferred pair declarations                                  |
| R10–R29                                      | Approved; implemented     | Only the listed one-sided ID additions; no mnemonic rewrites                                                  |
| Existing six Word Formation values / R01–R09 | Already present; preserve | No replacement, removal, or expansion                                                                         |

Approved implementation procedure (completed in this task):

1. Re-read AGENTS.md, the current entries, and this plan. Inspect Git status and preserve any new user work. If an approved target or source value has changed, reconcile it before writing rather than overwriting concurrent edits.
2. Snapshot the current entry bytes and media hashes. Add only the approved optional fields and source lines; preserve existing Notes except the explicitly authorized WF01 move. Do not change Word, Meaning, pronunciation, examples, tags, stable IDs, learning metadata, or other entry fields as part of this enrichment.
3. Add only approved relationship pairs, retaining existing members and avoiding duplicate declarations. Do not add prerequisites or treat associations as introduction dependencies. Existing endpoints such as emit-tech may receive a generated reverse display without a YAML edit.
4. Run the standard repository checks from the root:

```bash
go test ./...
go vet ./...
go run scripts/validate/main.go
go run scripts/build/main.go
git diff --check
```

5. Build again without changing source and compare TSV bytes. Verify the same ID set and row count as the pre-implementation baseline (1,241 at planning time), and the unchanged 17-column export, unchanged audio fields, preserved StudyOrder ranks, omitted empty sections, and direct bidirectional RelatedWords displays for only the approved additions. Existing reverse displays must remain intact. Word and learning metadata are unchanged, so no new pronunciation recordings are needed.
6. Check untracked text files separately and compare all entry changes against the approved field allowlist. Confirm every pre-existing media hash remains unchanged. Update this plan with actual implementation and verification results. Do not stage, commit, push, copy media into Anki, or modify an Anki collection.

A complete rebuild and reimport will eventually refresh reverse displays in Anki. Installed templates are unchanged by this plan. Native rendering, pronunciation quality, and live playback require separate device verification; do not claim those checks from static source inspection.

## Verification record

Planning-stage checks (before the subsequent implementation approval):

- All 546 batch files were read and assigned a ledger decision.
- Proposed target IDs and relationship endpoints were resolved against the full current collection; proposed pairs were checked against direct declarations in both directions to avoid re-proposing existing pairs.
- Field names and conditional presentation were checked against the current schema, exporter, association implementation, and README template.
- External lexical and historical evidence is bounded by the source register; no uncertain fictional naming claim was promoted to verified history.
- Planning validation passed: 546 ledger rows, 86 exact field drafts, 20 unique new pairs, 79 source definitions, and consistent pending approval statuses.
- All 840 local link occurrences resolve, including 11 same-file navigation anchors; all 186 shortcut source-reference uses resolve to defined URLs. No link checker network requests were made beyond the documented source research.
- git diff --check and a separate UTF-8/whitespace/final-newline check of this untracked file passed. SHA-256 checks preserve all 2,502 snapshotted files and their original file sets, including entries, media, implementation, README, AGENTS.md, dependencies, and the generated export. Git status shows only this new plan.
- During the initial planning stage, tests, vocabulary validation, export builds, audio generation, and Anki operations were not run. Implementation results follow separately.


### Approved implementation results

The user approved the complete plan with "I approve it." All 86 exact field values and the 20 pair additions were checked against the reviewed wording and ID declarations in this file.

- Updated 89 YAML files. The planning scope covers 98 new-batch entries because some association endpoints receive only a generated reverse display; emit-tech also receives a reverse display without a YAML edit.
- Added 48 Word Formation values (including the approved Autobuild Notes move), 18 Etymology values, and 20 Mnemonic values. All six existing Word Formation values remain unchanged.
- Appended or preserved source citations as specified; 65 Notes fields changed. Only Autobuild’s explicitly quoted construction paragraph was removed from Notes, and its qualification remains in Word Formation. All original provenance remains intact.
- Added 20 direct pairs through 19 source lists. Existing pair declarations remain intact. Exactly 38 RelatedWords exports changed, with direct neighbors verified in both directions against the complete collection.
- All 1,241 IDs, language values, Word, Meaning, pronunciation, part of speech, Context, Example, tags, audio references, priorities, and prerequisites remain unchanged. All 1,152 untouched YAML files retain their original bytes; edits in the other 89 files are restricted to approved fields and source lines.
- All 1,241 existing MP3 SHA-256 hashes and the media file set are unchanged. No recording was generated, replaced, or deleted; decoding and listening were not repeated for this text-only enrichment.
- go test ./..., go vet ./..., go run scripts/validate/main.go, and go run scripts/build/main.go all pass. Validation and builds retain the same 14 intentional repeated-word warnings documented in the import record.
- Two consecutive builds produce identical output: 1,241 data rows, 17 columns, and unchanged StudyOrder ranks. Empty optional aid values remain empty in the export; the existing conditional README template is unchanged. Static checks do not establish native Anki rendering.
- git diff --check and a separate formatting check of this untracked plan pass. Local links and reference-style citations resolve. Implementation files, templates, dependencies, unrelated guides, and Git-index state remain unchanged.

Export SHA-256: e6ce5ec7221a063c720f06b4a8eb0bc1f048a44c5b0c7c8cdaedaadd101ea2a3.

No new tests were added because no implementation logic changed. The source-preservation, exact-wording, complete-neighbor, and export comparisons were performed as task-specific artifact checks. Human pronunciation review and live Anki import, display, and playback were not performed. No staging, commit, push, or Anki operation was performed.

[P01]: https://www.merriam-webster.com/dictionary/un- "Merriam-Webster.com Dictionary — un-"
[P02]: https://www.merriam-webster.com/dictionary/re- "Merriam-Webster.com Dictionary — re-"
[P03]: https://www.merriam-webster.com/dictionary/-less "Merriam-Webster.com Dictionary — -less"
[P04]: https://www.merriam-webster.com/dictionary/-er "Merriam-Webster.com Dictionary — -er"
[P05]: https://www.merriam-webster.com/dictionary/-ment "Merriam-Webster.com Dictionary — -ment"
[P06]: https://www.merriam-webster.com/dictionary/-ful "Merriam-Webster.com Dictionary — -ful"
[P07]: https://www.merriam-webster.com/dictionary/-dom "Merriam-Webster.com Dictionary — -dom"
[P08]: https://www.merriam-webster.com/dictionary/ultra- "Merriam-Webster.com Dictionary — ultra-"
[P09]: https://www.merriam-webster.com/dictionary/auto- "Merriam-Webster.com Dictionary — auto-"
[P10]: https://www.merriam-webster.com/dictionary/-ed "Merriam-Webster.com Dictionary — -ed"
[P11]: https://www.merriam-webster.com/dictionary/-al "Merriam-Webster.com Dictionary — -al"
[P12]: https://www.merriam-webster.com/dictionary/-ous "Merriam-Webster.com Dictionary — -ous"
[L01]: https://www.merriam-webster.com/dictionary/headband "Merriam-Webster.com Dictionary — headband"
[L02]: https://www.merriam-webster.com/dictionary/headdress "Merriam-Webster.com Dictionary — headdress"
[L03]: https://www.merriam-webster.com/dictionary/foothill "Merriam-Webster.com Dictionary — foothill"
[L04]: https://www.etymonline.com/word/airborne "Online Etymology Dictionary — airborne"
[L05]: https://www.etymonline.com/word/refinery "Online Etymology Dictionary — refinery"
[L06]: https://www.etymonline.com/word/stabilizer "Online Etymology Dictionary — stabilizer"
[L07]: https://www.merriam-webster.com/dictionary/dispenser "Merriam-Webster.com Dictionary — dispenser"
[L08]: https://www.merriam-webster.com/dictionary/homing "Merriam-Webster.com Dictionary — home (verb; homed, homing)"
[L09]: https://www.merriam-webster.com/dictionary/descendant "Merriam-Webster.com Dictionary — descendant"
[L10]: https://www.etymonline.com/word/mysterious "Online Etymology Dictionary — mysterious"
[L11]: https://www.etymonline.com/word/waterfall "Online Etymology Dictionary — waterfall"
[L12]: https://www.etymonline.com/word/staircase "Online Etymology Dictionary — staircase"
[L13]: https://www.etymonline.com/word/highland "Online Etymology Dictionary — highland"
[L14]: https://www.etymonline.com/word/wetland "Online Etymology Dictionary — wetland"
[L15]: https://www.etymonline.com/word/ruthless "Online Etymology Dictionary — ruthless"
[L16]: https://www.etymonline.com/word/enhancement "Online Etymology Dictionary — enhancement"
[L17]: https://www.etymonline.com/word/settlement "Online Etymology Dictionary — settlement"
[L18]: https://www.etymonline.com/word/undead "Online Etymology Dictionary — undead"
[L19]: https://www.etymonline.com/word/underground "Online Etymology Dictionary — underground"
[L20]: https://www.merriam-webster.com/dictionary/susceptible "Merriam-Webster.com Dictionary — susceptible"
[L21]: https://www.merriam-webster.com/dictionary/sturdy "Merriam-Webster.com Dictionary — sturdy"
[L22]: https://www.merriam-webster.com/dictionary/hostile "Merriam-Webster.com Dictionary — hostile"
[L23]: https://www.merriam-webster.com/dictionary/aggressive "Merriam-Webster.com Dictionary — aggressive"
[L24]: https://www.merriam-webster.com/dictionary/heir "Merriam-Webster.com Dictionary — heir"
[L25]: https://www.merriam-webster.com/dictionary/reckless "Merriam-Webster.com Dictionary — reckless"
[L26]: https://www.merriam-webster.com/dictionary/bass "Merriam-Webster.com Dictionary — bass"
[L27]: https://www.merriam-webster.com/dictionary/tear "Merriam-Webster.com Dictionary — tear"
[L28]: https://www.merriam-webster.com/dictionary/waist "Merriam-Webster.com Dictionary — waist"
[L29]: https://www.merriam-webster.com/dictionary/guard "Merriam-Webster.com Dictionary — guard"
[L30]: https://www.merriam-webster.com/dictionary/fearless "Merriam-Webster.com Dictionary — fearless"
[L31]: https://www.merriam-webster.com/dictionary/selfless "Merriam-Webster.com Dictionary — selfless"
[L32]: https://www.etymonline.com/word/recharge "Online Etymology Dictionary — recharge"
[G01]: https://github.com/sh0wer1ee/TotkText/blob/1c66ac1a02edf2d73d10b0dc6f0f2de29da7b119/data/chunks/ActorMsg/PouchContent/0020.json "TOTK paired game text 1.2.1 — Obj_AutoBuilder_Name"
[G02]: https://github.com/sh0wer1ee/TotkText/blob/1c66ac1a02edf2d73d10b0dc6f0f2de29da7b119/data/chunks/ActorMsg/PouchContent/0022.json "TOTK paired game text 1.2.1 — Obj_UltraHand_Name"
[G03]: https://github.com/sh0wer1ee/TotkText/blob/1c66ac1a02edf2d73d10b0dc6f0f2de29da7b119/data/chunks/ActorMsg/PouchContent/0023.json "TOTK paired game text 1.2.1 — SpObj_FloatingStone_Capsule_A_01_Name"
[G04]: https://github.com/sh0wer1ee/TotkText/blob/1c66ac1a02edf2d73d10b0dc6f0f2de29da7b119/data/chunks/ActorMsg/PouchContent/0023.json "TOTK paired game text 1.2.1 — SpObj_FlameThrower_Capsule_A_01_Name"
[G05]: https://github.com/sh0wer1ee/TotkText/blob/1c66ac1a02edf2d73d10b0dc6f0f2de29da7b119/data/chunks/ActorMsg/PouchContent/0023.json "TOTK paired game text 1.2.1 — SpObj_SnowMachine_Capsule_A_01_Name"
[G06]: https://github.com/sh0wer1ee/TotkText/blob/1c66ac1a02edf2d73d10b0dc6f0f2de29da7b119/data/chunks/ActorMsg/PouchContent/0023.json "TOTK paired game text 1.2.1 — SpObj_ElectricBoxGenerator_Capsule_A_01_Name"
[G07]: https://github.com/sh0wer1ee/TotkText/blob/1c66ac1a02edf2d73d10b0dc6f0f2de29da7b119/data/chunks/ActorMsg/PouchContent/0023.json "TOTK paired game text 1.2.1 — SpObj_Beamos_Capsule_A_01_Name"
[G08]: https://github.com/sh0wer1ee/TotkText/blob/1c66ac1a02edf2d73d10b0dc6f0f2de29da7b119/data/chunks/StaticMsg/Reference/0000.json "TOTK paired game text 1.2.1 — P_Cannon"
[G09]: https://github.com/sh0wer1ee/TotkText/blob/1c66ac1a02edf2d73d10b0dc6f0f2de29da7b119/data/chunks/ActorMsg/PouchContent/0023.json "TOTK paired game text 1.2.1 — SpObj_GolemHead_Capsule_A_01_Name"
[G10]: https://github.com/sh0wer1ee/TotkText/blob/1c66ac1a02edf2d73d10b0dc6f0f2de29da7b119/data/chunks/StaticMsg/Reference/0001.json "TOTK paired game text 1.2.1 — RecycleBox"
[E01]: https://www.etymonline.com/word/stamina "Online Etymology Dictionary — stamina"
[E02]: https://www.etymonline.com/word/integrity "Online Etymology Dictionary — integrity"
[E03]: https://www.etymonline.com/word/resilient "Online Etymology Dictionary — resilient"
[E04]: https://www.etymonline.com/word/elixir "Online Etymology Dictionary — elixir"
[E05]: https://www.etymonline.com/word/recipe "Online Etymology Dictionary — recipe"
[E06]: https://www.etymonline.com/word/melee "Online Etymology Dictionary — melee"
[E07]: https://www.etymonline.com/word/hydrant "Online Etymology Dictionary — hydrant"
[E08]: https://www.etymonline.com/word/aquatic "Online Etymology Dictionary — aquatic"
[E09]: https://www.etymonline.com/word/vulnerable "Online Etymology Dictionary — vulnerable"
[E10]: https://www.etymonline.com/word/hostile "Online Etymology Dictionary — hostile"
[E11]: https://www.etymonline.com/word/trident "Online Etymology Dictionary — trident"
[E12]: https://www.etymonline.com/word/sacrifice "Online Etymology Dictionary — sacrifice"
[E13]: https://www.etymonline.com/word/malevolent "Online Etymology Dictionary — malevolent"
[E14]: https://www.etymonline.com/word/courageous "Online Etymology Dictionary — courageous"
[E15]: https://www.etymonline.com/word/enthusiastic "Online Etymology Dictionary — enthusiastic"
[E16]: https://www.etymonline.com/word/saute "Online Etymology Dictionary — saute"
[E17]: https://www.etymonline.com/word/paella "Online Etymology Dictionary — paella"
[E18]: https://www.etymonline.com/word/rupee "Online Etymology Dictionary — rupee"
[L33]: https://www.merriam-webster.com/dictionary/canyon "Merriam-Webster.com Dictionary — canyon"
[L34]: https://www.merriam-webster.com/dictionary/cavern "Merriam-Webster.com Dictionary — cavern"
[L35]: https://www.merriam-webster.com/dictionary/simmer "Merriam-Webster.com Dictionary — simmer"
[L36]: https://www.merriam-webster.com/dictionary/sear "Merriam-Webster.com Dictionary — sear"
[L37]: https://www.merriam-webster.com/dictionary/grill "Merriam-Webster.com Dictionary — grill"
[L38]: https://www.merriam-webster.com/dictionary/helm "Merriam-Webster.com Dictionary — helm"
[L39]: https://www.merriam-webster.com/dictionary/helmet "Merriam-Webster.com Dictionary — helmet"
