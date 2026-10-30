# American English pronunciation sources

## Status

Research pass complete with three unresolved entries; status reviewed on 2026-10-29. `accessor-tech`, `iterable-tech`, and `symlink-tech` still have no IPA. Counts below cover the original 284-entry collection, not later vocabulary batches. This record documents transcription evidence, not independent verification of generated audio pronunciation.

This pass accounts for the 282 entries that previously lacked pronunciation: **176 direct transcriptions**, **103 composed or derived transcriptions**, and **3 unresolved entries**. Together with the two preserved pronunciations, 281 of 284 entries have IPA. All 284 retain their existing audio references.

## Method and notation

- **Direct** means a consulted dictionary prints the complete word or expression, including an explicitly transcribed derived word or an equivalent spaced/hyphenated spelling. Prefer Oxford Advanced American Dictionary, then Oxford Advanced Learner’s Dictionary’s North American pronunciation and Cambridge’s US pronunciation.
- **Composed** means a transcription assembled from verified American component forms, with reviewed inflection, derivation, or compound stress. It is not a dictionary transcription of the complete expression. The entry index below records this distinction; card notes omit the repeated composition label.
- **Unresolved** means the available evidence did not support a reliable American transcription. The pronunciation field stays absent; audio remains available.
- Use broad rhotic IPA between slashes. Normalize Oxford’s /y/ to IPA /j/, /e/ to /ɛ/ except in /eɪ/, rhotic /ɚ, ɝ, ɜːr/ to /ər/, and marked flapped /t/ to broad /t/. Omit phonetic vowel-length marks and syllable separators. These notation changes do not select British variants.
- Retain the cited American variant rather than forcing every vowel reduction to match one speaker. For composed expressions, retain component word stresses; they do not prescribe a unique sentence-level accent. Use clear function-word forms, with ordinary unstressed /əv/ for of and /ðə/ for the before consonants (/ði/ before the vowel in original).
- Match noun/verb forms explicitly where stress differs: alternate is a verb; attribute in attribute to is a verb; overlap in overlap prevention and transport in transport metadata/failure are nouns. Inflectional -ed uses /ɪd/ after /t, d/, /t/ after other voiceless sounds, and /d/ elsewhere; plural -s uses /z/ for conclusions and internals.
- The existing IPA and complete source files for idiomatic-tech and pronunciation-general are preserved. The entries below cover only the 282 candidates for this pass.

## Unresolved entries

No American IPA was established for **accessor**, **iterable**, or **symlink**. Their indexed records below identify the evidence gap. They remain audio-only on the front until reliable pronunciation evidence is available.

## Apply the updated content in Anki

Use the existing multilingual LexForge note type and the current field mapping from the [README](../README.md#set-up-anki-once): twelve note fields, including derived `TTSLanguage` and `StudyOrder`, followed by tags in column 13. Reimport the rebuilt `dist/vocabulary.tsv`, choose to update existing notes using the first `ID` field and note-type match scope, then sync Anki Desktop and AnkiDroid. Keep the existing note type and notes to preserve scheduling. The front template already includes conditional pronunciation and audio fields; no template change or new media copy is needed for this pronunciation update once the current README setup is in place.

## Verification

- Vocabulary validation passed for all 284 entries. The five existing same-word warnings still describe distinct senses.
- `go test ./...`, `go vet ./...`, and `git diff --check` passed. Two export builds produced identical bytes.
- Source comparison confirmed that only 279 pronunciation fields and 103 appended composition labels changed. Existing note text and all other fields were preserved; the two previously populated entries and all 284 MP3 files remained byte-identical.
- Every exported row has eleven columns, the expected pronunciation value, and its original audio reference. The source index accounts for all 282 candidates.
- Front-template field substitution was checked for a single word, an expression, a long expression, and an unresolved entry. This was a template check, not an Anki renderer or device test.
- No Anki import, synchronization, device playback, or independent listening review was performed in this pass. The IPA describes a selected American reading, not an acoustic transcription of Samantha's recordings. Pronunciation agreement with those recordings remains unverified.

## Entry index

### abandoned-workflow

- Entry: [abandoned](../entries/abandoned-workflow.yml)
- IPA: /əˈbændənd/
- Basis: Direct
- Sources: [abandoned](https://www.oxfordlearnersdictionaries.com/definition/american_english/abandoned)

### accented-text-tech

- Entry: [accented text](../entries/accented-text-tech.yml)
- IPA: /ˈæksɛntəd tɛkst/
- Basis: Composed
- Review: Component citation forms; retain word-level stresses, without claiming a unique sentence-level accent.
- Sources: [accented](https://www.oxfordlearnersdictionaries.com/definition/american_english/accented); [text](https://www.oxfordlearnersdictionaries.com/definition/american_english/text_1)

### accessor-tech

- Entry: [accessor](../entries/accessor-tech.yml)
- IPA: unset
- Basis: Unresolved
- Review: Oxford and Cambridge lookups did not provide a matching American transcription. The supplementary Wiktionary entry has no English IPA; its pronunciation section belongs to Latin. Do not reuse that Latin transcription or infer English stress from accessory.
- Sources: [supplementary lookup (not American IPA evidence)](https://en.wiktionary.org/wiki/accessor#English)

### acquire-workflow

- Entry: [acquire](../entries/acquire-workflow.yml)
- IPA: /əˈkwaɪər/
- Basis: Direct
- Sources: [acquire](https://www.oxfordlearnersdictionaries.com/definition/american_english/acquire)

### actionable-tech

- Entry: [actionable](../entries/actionable-tech.yml)
- IPA: /ˈækʃənəbl/
- Basis: Direct
- Sources: [actionable](https://www.oxfordlearnersdictionaries.com/definition/american_english/actionable)

### actual-behavior-workflow

- Entry: [actual behavior](../entries/actual-behavior-workflow.yml)
- IPA: /ˈæktʃuəl bɪˈheɪvjər/
- Basis: Composed
- Review: Component citation forms; retain word-level stresses, without claiming a unique sentence-level accent.
- Sources: [actual](https://www.oxfordlearnersdictionaries.com/definition/american_english/actual); [behavior](https://www.oxfordlearnersdictionaries.com/definition/american_english/behavior)

### adapter-tech

- Entry: [adapter](../entries/adapter-tech.yml)
- IPA: /əˈdæptər/
- Basis: Direct
- Sources: [adapter](https://www.oxfordlearnersdictionaries.com/definition/american_english/adapter)

### adjacent-tech

- Entry: [adjacent](../entries/adjacent-tech.yml)
- IPA: /əˈdʒeɪsnt/
- Basis: Direct
- Sources: [adjacent](https://www.oxfordlearnersdictionaries.com/definition/american_english/adjacent)

### allowlist-tech

- Entry: [allowlist](../entries/allowlist-tech.yml)
- IPA: /əˈlaʊˌlɪst/
- Basis: Composed
- Review: Closed compound: allow + list; retain stress in allow and secondary stress on list.
- Sources: [allow](https://www.oxfordlearnersdictionaries.com/definition/american_english/allow); [list](https://www.oxfordlearnersdictionaries.com/definition/american_english/list_1)

### alternate-tech

- Entry: [alternate](../entries/alternate-tech.yml)
- IPA: /ˈɔltərˌneɪt/
- Basis: Direct
- Sources: [alternate](https://www.oxfordlearnersdictionaries.com/definition/american_english/alternate_2)

### amend-git

- Entry: [amend](../entries/amend-git.yml)
- IPA: /əˈmɛnd/
- Basis: Direct
- Sources: [amend](https://www.oxfordlearnersdictionaries.com/definition/american_english/amend)

### annotation-tech

- Entry: [annotation](../entries/annotation-tech.yml)
- IPA: /ˌænəˈteɪʃn/
- Basis: Direct
- Sources: [annotation](https://www.oxfordlearnersdictionaries.com/definition/english/annotation)

### anonymous-tech

- Entry: [anonymous](../entries/anonymous-tech.yml)
- IPA: /əˈnɑnəməs/
- Basis: Direct
- Sources: [anonymous](https://www.oxfordlearnersdictionaries.com/definition/american_english/anonymous)

### applicable-tech

- Entry: [applicable](../entries/applicable-tech.yml)
- IPA: /ˈæplɪkəbl/
- Basis: Direct
- Sources: [applicable](https://www.oxfordlearnersdictionaries.com/definition/american_english/applicable)

### appreciate-general

- Entry: [appreciate](../entries/appreciate-general.yml)
- IPA: /əˈpriʃiˌeɪt/
- Basis: Direct
- Sources: [appreciate](https://www.oxfordlearnersdictionaries.com/definition/american_english/appreciate)

### array-spread-tech

- Entry: [array spread](../entries/array-spread-tech.yml)
- IPA: /əˈreɪ sprɛd/
- Basis: Composed
- Review: Component citation forms; retain word-level stresses, without claiming a unique sentence-level accent.
- Sources: [array](https://www.oxfordlearnersdictionaries.com/definition/american_english/array_1); [spread](https://www.oxfordlearnersdictionaries.com/definition/american_english/spread_1)

### artifact-tech

- Entry: [artifact](../entries/artifact.yml)
- IPA: /ˈɑrtəˌfækt/
- Basis: Direct
- Sources: [artifact](https://www.oxfordlearnersdictionaries.com/definition/american_english/artifact)

### artifact-workflow

- Entry: [artifact](../entries/artifact-workflow.yml)
- IPA: /ˈɑrtəˌfækt/
- Basis: Direct
- Sources: [artifact](https://www.oxfordlearnersdictionaries.com/definition/american_english/artifact)

### as-proof-of-workflow

- Entry: [as proof of](../entries/as-proof-of-workflow.yml)
- IPA: /æz pruf əv/
- Basis: Composed
- Review: Component citation forms; retain word-level stresses, without claiming a unique sentence-level accent.
- Sources: [as](https://www.oxfordlearnersdictionaries.com/definition/american_english/as_1); [proof](https://www.oxfordlearnersdictionaries.com/definition/american_english/proof_1); [of](https://www.oxfordlearnersdictionaries.com/definition/american_english/of)

### assertion-testing

- Entry: [assertion](../entries/assertion-testing.yml)
- IPA: /əˈsərʃn/
- Basis: Direct
- Sources: [assertion](https://www.oxfordlearnersdictionaries.com/definition/american_english/assertion)

### assignment-target-tech

- Entry: [assignment target](../entries/assignment-target-tech.yml)
- IPA: /əˈsaɪnmənt ˈtɑrɡət/
- Basis: Composed
- Review: Component citation forms; retain word-level stresses, without claiming a unique sentence-level accent.
- Sources: [assignment](https://www.oxfordlearnersdictionaries.com/definition/american_english/assignment); [target](https://www.oxfordlearnersdictionaries.com/definition/american_english/target_1)

### asynchronous-tech

- Entry: [asynchronous](../entries/asynchronous-tech.yml)
- IPA: /eɪˈsɪŋkrənəs/
- Basis: Direct
- Sources: [asynchronous](https://www.oxfordlearnersdictionaries.com/definition/american_english/asynchronous)

### at-most-workflow

- Entry: [at most](../entries/at-most-workflow.yml)
- IPA: /æt moʊst/
- Basis: Composed
- Review: Component citation forms; retain word-level stresses, without claiming a unique sentence-level accent.
- Sources: [at](https://www.oxfordlearnersdictionaries.com/definition/american_english/at); [most](https://www.oxfordlearnersdictionaries.com/definition/american_english/most_1)

### atomically-workflow

- Entry: [atomically](../entries/atomically-workflow.yml)
- IPA: /əˈtɑmɪkəli/
- Basis: Direct
- Sources: [atomically](https://dictionary.cambridge.org/dictionary/english/atomically)

### attribute-to-general

- Entry: [attribute to](../entries/attribute-to-general.yml)
- IPA: /əˈtrɪbjut tu/
- Basis: Composed
- Review: Component citation forms; retain word-level stresses, without claiming a unique sentence-level accent.
- Sources: [attribute](https://www.oxfordlearnersdictionaries.com/definition/american_english/attribute_1); [to](https://www.oxfordlearnersdictionaries.com/definition/american_english/to_1)

### audit-tech

- Entry: [audit](../entries/audit-tech.yml)
- IPA: /ˈɔdət/
- Basis: Direct
- Sources: [audit](https://www.oxfordlearnersdictionaries.com/definition/american_english/audit_1)

### await-review-workflow

- Entry: [await review](../entries/await-review-workflow.yml)
- IPA: /əˈweɪt rɪˈvju/
- Basis: Composed
- Review: Component citation forms; retain word-level stresses, without claiming a unique sentence-level accent.
- Sources: [await](https://www.oxfordlearnersdictionaries.com/definition/american_english/await); [review](https://www.oxfordlearnersdictionaries.com/definition/american_english/review_1)

### backend-tech

- Entry: [backend](../entries/backend-tech.yml)
- IPA: /ˈbæk ˌɛnd/
- Basis: Direct
- Review: Same spoken compound as dictionary spelling back end.
- Sources: [back end](https://dictionary.cambridge.org/dictionary/english/back-end)

### based-on-workflow

- Entry: [based on](../entries/based-on-workflow.yml)
- IPA: /beɪst ɑn/
- Basis: Composed
- Review: Component citation forms; retain word-level stresses, without claiming a unique sentence-level accent.
- Sources: [based](https://www.oxfordlearnersdictionaries.com/definition/american_english/based); [on](https://www.oxfordlearnersdictionaries.com/definition/american_english/on_1)

### baseline-tech

- Entry: [baseline](../entries/baseline-tech.yml)
- IPA: /ˈbeɪslaɪn/
- Basis: Direct
- Sources: [baseline](https://www.oxfordlearnersdictionaries.com/definition/american_english/baseline)

### before-and-after-the-fix-workflow

- Entry: [before and after the fix](../entries/before-and-after-the-fix-workflow.yml)
- IPA: /bɪˈfɔr ænd ˈæftər ðə fɪks/
- Basis: Composed
- Review: Component citation forms; retain word-level stresses, without claiming a unique sentence-level accent.
- Sources: [before](https://www.oxfordlearnersdictionaries.com/definition/american_english/before_1); [and](https://www.oxfordlearnersdictionaries.com/definition/american_english/and); [after](https://www.oxfordlearnersdictionaries.com/definition/american_english/after_1); [the](https://www.oxfordlearnersdictionaries.com/definition/american_english/the); [fix](https://www.oxfordlearnersdictionaries.com/definition/american_english/fix_1)

### before-drawing-conclusions-workflow

- Entry: [before drawing conclusions](../entries/before-drawing-conclusions-workflow.yml)
- IPA: /bɪˈfɔr ˈdrɔɪŋ kənˈkluʒnz/
- Basis: Composed
- Review: Component citation forms; retain word-level stresses, without claiming a unique sentence-level accent. conclusions: Plural conclusion + /z/.
- Sources: [before](https://www.oxfordlearnersdictionaries.com/definition/american_english/before_1); [drawing](https://www.oxfordlearnersdictionaries.com/definition/american_english/drawing); [conclusion](https://www.oxfordlearnersdictionaries.com/definition/american_english/conclusion); [-s](https://dictionary.cambridge.org/dictionary/english/s)

### benchmark-tech

- Entry: [benchmark](../entries/benchmark-tech.yml)
- IPA: /ˈbɛntʃmɑrk/
- Basis: Direct
- Sources: [benchmark](https://www.oxfordlearnersdictionaries.com/definition/american_english/benchmark_1)

### binding-tech

- Entry: [binding](../entries/binding-tech.yml)
- IPA: /ˈbaɪndɪŋ/
- Basis: Direct
- Sources: [binding](https://www.oxfordlearnersdictionaries.com/definition/american_english/binding_2)

### blocker-workflow

- Entry: [blocker](../entries/blocker-workflow.yml)
- IPA: /ˈblɑkər/
- Basis: Direct
- Sources: [blocker](https://dictionary.cambridge.org/dictionary/english/blocker)

### bounded-tech

- Entry: [bounded](../entries/bounded-tech.yml)
- IPA: /ˈbaʊndɪd/
- Basis: Composed
- Review: Bound + -ed /ɪd/ after /d/.
- Sources: [bound](https://www.oxfordlearnersdictionaries.com/definition/american_english/bound_2); [-ed](https://dictionary.cambridge.org/dictionary/english/ed)

### breadth-first-tech

- Entry: [breadth-first](../entries/breadth-first-tech.yml)
- IPA: /ˌbrɛdθ ˈfərst/
- Basis: Composed
- Review: Compound modifier: breadth + first; primary stress on first.
- Sources: [breadth](https://www.oxfordlearnersdictionaries.com/definition/american_english/breadth); [first](https://www.oxfordlearnersdictionaries.com/definition/american_english/first_1)

### bridge-tech

- Entry: [bridge](../entries/bridge-tech.yml)
- IPA: /brɪdʒ/
- Basis: Direct
- Sources: [bridge](https://www.oxfordlearnersdictionaries.com/definition/american_english/bridge_1)

### bypass-tech

- Entry: [bypass](../entries/bypass-tech.yml)
- IPA: /ˈbaɪpæs/
- Basis: Direct
- Sources: [bypass](https://www.oxfordlearnersdictionaries.com/definition/american_english/bypass_2)

### byte-offset-tech

- Entry: [byte offset](../entries/byte-offset-tech.yml)
- IPA: /baɪt ˈɔfsɛt/
- Basis: Composed
- Review: Component citation forms; retain word-level stresses, without claiming a unique sentence-level accent.
- Sources: [byte](https://www.oxfordlearnersdictionaries.com/definition/american_english/byte); [offset](https://www.oxfordlearnersdictionaries.com/definition/american_english/offset_1)

### callback-tech

- Entry: [callback](../entries/callback-tech.yml)
- IPA: /ˈkɔlbæk/
- Basis: Direct
- Sources: [callback](https://www.oxfordlearnersdictionaries.com/definition/american_english/callback)

### cancellation-tech

- Entry: [cancellation](../entries/cancellation-tech.yml)
- IPA: /ˌkænsəˈleɪʃn/
- Basis: Direct
- Sources: [cancellation](https://www.oxfordlearnersdictionaries.com/definition/american_english/cancellation)

### candidate-tech

- Entry: [candidate](../entries/candidate-tech.yml)
- IPA: /ˈkændəˌdeɪt/
- Basis: Direct
- Sources: [candidate](https://www.oxfordlearnersdictionaries.com/definition/american_english/candidate)

### cast-tech

- Entry: [cast](../entries/cast-tech.yml)
- IPA: /kæst/
- Basis: Direct
- Sources: [cast](https://www.oxfordlearnersdictionaries.com/definition/american_english/cast_2)

### catch-handler-tech

- Entry: [catch handler](../entries/catch-handler-tech.yml)
- IPA: /kætʃ ˈhændlər/
- Basis: Composed
- Review: Component citation forms; retain word-level stresses, without claiming a unique sentence-level accent.
- Sources: [catch](https://www.oxfordlearnersdictionaries.com/definition/american_english/catch_1); [handler](https://www.oxfordlearnersdictionaries.com/definition/american_english/handler)

### caveat-general

- Entry: [caveat](../entries/caveat-general.yml)
- IPA: /ˈkæviˌɑt/
- Basis: Direct
- Sources: [caveat](https://www.oxfordlearnersdictionaries.com/definition/american_english/caveat)

### claim-general

- Entry: [claim](../entries/claim-general.yml)
- IPA: /kleɪm/
- Basis: Direct
- Sources: [claim](https://www.oxfordlearnersdictionaries.com/definition/american_english/claim_2)

### clarify-general

- Entry: [clarify](../entries/clarify-general.yml)
- IPA: /ˈklærəˌfaɪ/
- Basis: Direct
- Sources: [clarify](https://www.oxfordlearnersdictionaries.com/definition/american_english/clarify)

### clause-tech

- Entry: [clause](../entries/clause-tech.yml)
- IPA: /klɔz/
- Basis: Direct
- Sources: [clause](https://www.oxfordlearnersdictionaries.com/definition/american_english/clause)

### cleanup-tech

- Entry: [cleanup](../entries/cleanup-tech.yml)
- IPA: /ˈklinʌp/
- Basis: Direct
- Sources: [cleanup](https://dictionary.cambridge.org/dictionary/english/cleanup)

### closure-tech

- Entry: [closure](../entries/closure-tech.yml)
- IPA: /ˈkloʊʒər/
- Basis: Direct
- Sources: [closure](https://www.oxfordlearnersdictionaries.com/definition/american_english/closure)

### code-generation-tech

- Entry: [code generation](../entries/code-generation-tech.yml)
- IPA: /koʊd ˌdʒɛnəˈreɪʃn/
- Basis: Composed
- Review: Component citation forms; retain word-level stresses, without claiming a unique sentence-level accent.
- Sources: [code](https://www.oxfordlearnersdictionaries.com/definition/american_english/code_1); [generation](https://www.oxfordlearnersdictionaries.com/definition/american_english/generation)

### code-unit-tech

- Entry: [code unit](../entries/code-unit-tech.yml)
- IPA: /koʊd ˈjunət/
- Basis: Composed
- Review: Component citation forms; retain word-level stresses, without claiming a unique sentence-level accent.
- Sources: [code](https://www.oxfordlearnersdictionaries.com/definition/american_english/code_1); [unit](https://www.oxfordlearnersdictionaries.com/definition/american_english/unit)

### coerce-tech

- Entry: [coerce](../entries/coerce-tech.yml)
- IPA: /koʊˈərs/
- Basis: Direct
- Sources: [coerce](https://www.oxfordlearnersdictionaries.com/definition/american_english/coerce)

### collision-tech

- Entry: [collision](../entries/collision-tech.yml)
- IPA: /kəˈlɪʒn/
- Basis: Direct
- Sources: [collision](https://www.oxfordlearnersdictionaries.com/definition/american_english/collision)

### combining-mark-tech

- Entry: [combining mark](../entries/combining-mark-tech.yml)
- IPA: /kəmˈbaɪnɪŋ mɑrk/
- Basis: Composed
- Review: Component citation forms; retain word-level stresses, without claiming a unique sentence-level accent. combining: Combine + -ing; retain stem stress.
- Sources: [combine](https://www.oxfordlearnersdictionaries.com/definition/american_english/combine_1); [-ing](https://dictionary.cambridge.org/dictionary/english/ing); [mark](https://www.oxfordlearnersdictionaries.com/definition/american_english/mark_1)

### competing-work-workflow

- Entry: [competing work](../entries/competing-work-workflow.yml)
- IPA: /kəmˈpitɪŋ wərk/
- Basis: Composed
- Review: Component citation forms; retain word-level stresses, without claiming a unique sentence-level accent.
- Sources: [competing](https://www.oxfordlearnersdictionaries.com/definition/american_english/competing); [work](https://www.oxfordlearnersdictionaries.com/definition/american_english/work_1)

### compound-assignment-tech

- Entry: [compound assignment](../entries/compound-assignment-tech.yml)
- IPA: /ˈkɑmpaʊnd əˈsaɪnmənt/
- Basis: Composed
- Review: Component citation forms; retain word-level stresses, without claiming a unique sentence-level accent.
- Sources: [compound](https://www.oxfordlearnersdictionaries.com/definition/american_english/compound_2); [assignment](https://www.oxfordlearnersdictionaries.com/definition/american_english/assignment)

### computed-key-tech

- Entry: [computed key](../entries/computed-key-tech.yml)
- IPA: /kəmˈpjutɪd ki/
- Basis: Composed
- Review: Component citation forms; retain word-level stresses, without claiming a unique sentence-level accent. computed: Compute + -ed /ɪd/ after /t/.
- Sources: [compute](https://www.oxfordlearnersdictionaries.com/definition/american_english/compute); [-ed](https://dictionary.cambridge.org/dictionary/english/ed); [key](https://www.oxfordlearnersdictionaries.com/definition/american_english/key_1)

### concatenation-tech

- Entry: [concatenation](../entries/concatenation-tech.yml)
- IPA: /kənˌkætnˈeɪʃn/
- Basis: Direct
- Sources: [concatenation](https://www.oxfordlearnersdictionaries.com/definition/american_english/concatenation)

### concise-general

- Entry: [concise](../entries/concise-general.yml)
- IPA: /kənˈsaɪs/
- Basis: Direct
- Sources: [concise](https://www.oxfordlearnersdictionaries.com/definition/american_english/concise)

### concrete-tech

- Entry: [concrete](../entries/concrete-tech.yml)
- IPA: /kɑnˈkrit/
- Basis: Direct
- Sources: [concrete](https://www.oxfordlearnersdictionaries.com/definition/american_english/concrete_1)

### confirmed-tech

- Entry: [confirmed](../entries/confirmed-tech.yml)
- IPA: /kənˈfərmd/
- Basis: Direct
- Sources: [confirmed](https://www.oxfordlearnersdictionaries.com/definition/american_english/confirmed)

### consistent-tech

- Entry: [consistent](../entries/consistent-tech.yml)
- IPA: /kənˈsɪstənt/
- Basis: Direct
- Sources: [consistent](https://www.oxfordlearnersdictionaries.com/definition/american_english/consistent)

### constant-propagation-tech

- Entry: [constant propagation](../entries/constant-propagation-tech.yml)
- IPA: /ˈkɑnstənt ˌprɑpəˈɡeɪʃn/
- Basis: Composed
- Review: Component citation forms; retain word-level stresses, without claiming a unique sentence-level accent.
- Sources: [constant](https://www.oxfordlearnersdictionaries.com/definition/american_english/constant_1); [propagation](https://www.oxfordlearnersdictionaries.com/definition/american_english/propagate)

### consume-tech

- Entry: [consume](../entries/consume-tech.yml)
- IPA: /kənˈsum/
- Basis: Direct
- Sources: [consume](https://www.oxfordlearnersdictionaries.com/definition/american_english/consume)

### contextual-type-tech

- Entry: [contextual type](../entries/contextual-type-tech.yml)
- IPA: /kənˈtɛkstʃuəl taɪp/
- Basis: Composed
- Review: Component citation forms; retain word-level stresses, without claiming a unique sentence-level accent.
- Sources: [contextual](https://www.oxfordlearnersdictionaries.com/definition/american_english/contextual); [type](https://www.oxfordlearnersdictionaries.com/definition/american_english/type_1)

### contribution-general

- Entry: [contribution](../entries/contribution-general.yml)
- IPA: /ˌkɑntrəˈbjuʃn/
- Basis: Direct
- Sources: [contribution](https://www.oxfordlearnersdictionaries.com/definition/american_english/contribution)

### correctness-tech

- Entry: [correctness](../entries/correctness-tech.yml)
- IPA: /kəˈrɛktnəs/
- Basis: Direct
- Sources: [correctness](https://www.oxfordlearnersdictionaries.com/definition/english/correctness)

### corresponding-general

- Entry: [corresponding](../entries/corresponding-general.yml)
- IPA: /ˌkɔrəˈspɑndɪŋ/
- Basis: Direct
- Sources: [corresponding](https://www.oxfordlearnersdictionaries.com/definition/american_english/corresponding)

### corrupt-tech

- Entry: [corrupt](../entries/corrupt-tech.yml)
- IPA: /kəˈrʌpt/
- Basis: Direct
- Sources: [corrupt](https://www.oxfordlearnersdictionaries.com/definition/american_english/corrupt_2)

### coverage-tech

- Entry: [coverage](../entries/coverage-tech.yml)
- IPA: /ˈkʌvrɪdʒ/
- Basis: Direct
- Sources: [coverage](https://www.oxfordlearnersdictionaries.com/definition/american_english/coverage)

### declaration-emit-tech

- Entry: [declaration emit](../entries/declaration-emit-tech.yml)
- IPA: /ˌdɛkləˈreɪʃn ɪˈmɪt/
- Basis: Composed
- Review: Component citation forms; retain word-level stresses, without claiming a unique sentence-level accent.
- Sources: [declaration](https://www.oxfordlearnersdictionaries.com/definition/american_english/declaration); [emit](https://www.oxfordlearnersdictionaries.com/definition/american_english/emit)

### decode-tech

- Entry: [decode](../entries/decode-tech.yml)
- IPA: /diˈkoʊd/
- Basis: Direct
- Sources: [decode](https://www.oxfordlearnersdictionaries.com/definition/american_english/decode)

### decorator-tech

- Entry: [decorator](../entries/decorator-tech.yml)
- IPA: /ˈdɛkəˌreɪtər/
- Basis: Direct
- Sources: [decorator](https://www.oxfordlearnersdictionaries.com/definition/american_english/decorator)

### defer-tech

- Entry: [defer](../entries/defer-tech.yml)
- IPA: /dɪˈfər/
- Basis: Direct
- Sources: [defer](https://www.oxfordlearnersdictionaries.com/definition/american_english/defer)

### deprecated-tech

- Entry: [deprecated](../entries/deprecated.yml)
- IPA: /ˈdɛprəˌkeɪtɪd/
- Basis: Composed
- Review: Deprecate + -ed /ɪd/ after /t/.
- Sources: [deprecate](https://www.oxfordlearnersdictionaries.com/definition/american_english/deprecate); [-ed](https://dictionary.cambridge.org/dictionary/english/ed)

### derive-tech

- Entry: [derive](../entries/derive-tech.yml)
- IPA: /dɪˈraɪv/
- Basis: Direct
- Sources: [derive](https://www.oxfordlearnersdictionaries.com/definition/american_english/derive)

### detach-git

- Entry: [detach](../entries/detach-git.yml)
- IPA: /dɪˈtætʃ/
- Basis: Direct
- Sources: [detach](https://www.oxfordlearnersdictionaries.com/definition/american_english/detach)

### diagnostic-context-tech

- Entry: [diagnostic context](../entries/diagnostic-context-tech.yml)
- IPA: /ˌdaɪəɡˈnɑstɪk ˈkɑntɛkst/
- Basis: Composed
- Review: Component citation forms; retain word-level stresses, without claiming a unique sentence-level accent.
- Sources: [diagnostic](https://www.oxfordlearnersdictionaries.com/definition/american_english/diagnostic_1); [diagnostic](https://dictionary.cambridge.org/dictionary/english/diagnostic); [context](https://www.oxfordlearnersdictionaries.com/definition/american_english/context)

### diagnostic-tech

- Entry: [diagnostic](../entries/diagnostic-tech.yml)
- IPA: /ˌdaɪəɡˈnɑstɪk/
- Basis: Direct
- Review: Use Cambridge US /ɪ/ in the final syllable; Oxford also supplies the noun entry.
- Sources: [diagnostic](https://www.oxfordlearnersdictionaries.com/definition/american_english/diagnostic_2); [diagnostic](https://dictionary.cambridge.org/dictionary/english/diagnostic)

### directive-tech

- Entry: [directive](../entries/directive-tech.yml)
- IPA: /dəˈrɛktɪv/
- Basis: Direct
- Sources: [directive](https://www.oxfordlearnersdictionaries.com/definition/american_english/directive_1)

### discard-workflow

- Entry: [discard](../entries/discard-workflow.yml)
- IPA: /dɪˈskɑrd/
- Basis: Direct
- Sources: [discard](https://www.oxfordlearnersdictionaries.com/definition/american_english/discard_1)

### disclosure-workflow

- Entry: [disclosure](../entries/disclosure-workflow.yml)
- IPA: /dɪsˈkloʊʒər/
- Basis: Direct
- Sources: [disclosure](https://www.oxfordlearnersdictionaries.com/definition/american_english/disclosure)

### dispatcher-tech

- Entry: [dispatcher](../entries/dispatcher-tech.yml)
- IPA: /dɪˈspætʃər/
- Basis: Direct
- Sources: [dispatcher](https://www.oxfordlearnersdictionaries.com/definition/american_english/dispatcher)

### distinguish-tech

- Entry: [distinguish](../entries/distinguish-tech.yml)
- IPA: /dɪˈstɪŋɡwɪʃ/
- Basis: Direct
- Sources: [distinguish](https://www.oxfordlearnersdictionaries.com/definition/american_english/distinguish)

### dynamic-gating-tech

- Entry: [dynamic gating](../entries/dynamic-gating-tech.yml)
- IPA: /daɪˈnæmɪk ˈɡeɪtɪŋ/
- Basis: Composed
- Review: Component citation forms; retain word-level stresses, without claiming a unique sentence-level accent. gating: Gate + -ing.
- Sources: [dynamic](https://www.oxfordlearnersdictionaries.com/definition/american_english/dynamic_1); [gate](https://www.oxfordlearnersdictionaries.com/definition/american_english/gate); [-ing](https://dictionary.cambridge.org/dictionary/english/ing)

### edge-case-tech

- Entry: [edge case](../entries/edge-case-tech.yml)
- IPA: /ˈɛdʒ keɪs/
- Basis: Direct
- Sources: [edge case](https://www.oxfordlearnersdictionaries.com/definition/english/edge-case)

### effect-replay-tech

- Entry: [effect replay](../entries/effect-replay-tech.yml)
- IPA: /ɪˈfɛkt ˈripleɪ/
- Basis: Composed
- Review: Component citation forms; retain word-level stresses, without claiming a unique sentence-level accent.
- Sources: [effect](https://www.oxfordlearnersdictionaries.com/definition/american_english/effect_1); [replay](https://www.oxfordlearnersdictionaries.com/definition/american_english/replay_1)

### elapsed-time-workflow

- Entry: [elapsed time](../entries/elapsed-time-workflow.yml)
- IPA: /ɪˌlæpst ˈtaɪm/
- Basis: Direct
- Sources: [elapsed time](https://www.oxfordlearnersdictionaries.com/definition/english/elapsed-time)

### eligible-tech

- Entry: [eligible](../entries/eligible-tech.yml)
- IPA: /ˈɛlədʒəbl/
- Basis: Direct
- Sources: [eligible](https://www.oxfordlearnersdictionaries.com/definition/american_english/eligible)

### emit-tech

- Entry: [emit](../entries/emit-tech.yml)
- IPA: /ɪˈmɪt/
- Basis: Direct
- Sources: [emit](https://www.oxfordlearnersdictionaries.com/definition/american_english/emit)

### encode-tech

- Entry: [encode](../entries/encode-tech.yml)
- IPA: /ɪnˈkoʊd/
- Basis: Direct
- Sources: [encode](https://www.oxfordlearnersdictionaries.com/definition/american_english/encode)

### equality-semantics-tech

- Entry: [equality semantics](../entries/equality-semantics-tech.yml)
- IPA: /ɪˈkwɑləti səˈmæntɪks/
- Basis: Composed
- Review: Component citation forms; retain word-level stresses, without claiming a unique sentence-level accent.
- Sources: [equality](https://www.oxfordlearnersdictionaries.com/definition/american_english/equality); [semantics](https://www.oxfordlearnersdictionaries.com/definition/american_english/semantics)

### escape-parity-tech

- Entry: [escape parity](../entries/escape-parity-tech.yml)
- IPA: /ɪˈskeɪp ˈpærəti/
- Basis: Composed
- Review: Component citation forms; retain word-level stresses, without claiming a unique sentence-level accent.
- Sources: [escape](https://www.oxfordlearnersdictionaries.com/definition/american_english/escape_1); [parity](https://www.oxfordlearnersdictionaries.com/definition/american_english/parity)

### escape-tech

- Entry: [escape](../entries/escape-tech.yml)
- IPA: /ɪˈskeɪp/
- Basis: Direct
- Sources: [escape](https://www.oxfordlearnersdictionaries.com/definition/american_english/escape_2)

### escaped-backslash-tech

- Entry: [escaped backslash](../entries/escaped-backslash-tech.yml)
- IPA: /ɪˈskeɪpt ˈbækslæʃ/
- Basis: Composed
- Review: Component citation forms; retain word-level stresses, without claiming a unique sentence-level accent.
- Sources: [escaped](https://www.oxfordlearnersdictionaries.com/definition/american_english/escaped); [backslash](https://www.oxfordlearnersdictionaries.com/definition/american_english/backslash)

### evidence-tech

- Entry: [evidence](../entries/evidence-tech.yml)
- IPA: /ˈɛvədəns/
- Basis: Direct
- Sources: [evidence](https://www.oxfordlearnersdictionaries.com/definition/american_english/evidence_1)

### exclude-tech

- Entry: [exclude](../entries/exclude-tech.yml)
- IPA: /ɪkˈsklud/
- Basis: Direct
- Sources: [exclude](https://www.oxfordlearnersdictionaries.com/definition/american_english/exclude)

### exhaustive-tech

- Entry: [exhaustive](../entries/exhaustive-tech.yml)
- IPA: /ɪɡˈzɔstɪv/
- Basis: Direct
- Sources: [exhaustive](https://www.oxfordlearnersdictionaries.com/definition/american_english/exhaustive)

### exit-status-tech

- Entry: [exit status](../entries/exit-status-tech.yml)
- IPA: /ˈɛɡzət ˈstætəs/
- Basis: Composed
- Review: Component citation forms; retain word-level stresses, without claiming a unique sentence-level accent.
- Sources: [exit](https://www.oxfordlearnersdictionaries.com/definition/american_english/exit_1); [status](https://www.oxfordlearnersdictionaries.com/definition/american_english/status)

### expected-behavior-workflow

- Entry: [expected behavior](../entries/expected-behavior-workflow.yml)
- IPA: /ɪkˈspɛktəd bɪˈheɪvjər/
- Basis: Composed
- Review: Component citation forms; retain word-level stresses, without claiming a unique sentence-level accent.
- Sources: [expected](https://www.oxfordlearnersdictionaries.com/definition/american_english/expected); [behavior](https://www.oxfordlearnersdictionaries.com/definition/american_english/behavior)

### explicit-tech

- Entry: [explicit](../entries/explicit-tech.yml)
- IPA: /ɪkˈsplɪsət/
- Basis: Direct
- Sources: [explicit](https://www.oxfordlearnersdictionaries.com/definition/american_english/explicit)

### expose-tech

- Entry: [expose](../entries/expose-tech.yml)
- IPA: /ɪkˈspoʊz/
- Basis: Direct
- Sources: [expose](https://www.oxfordlearnersdictionaries.com/definition/american_english/expose_1)

### factual-general

- Entry: [factual](../entries/factual-general.yml)
- IPA: /ˈfæktʃuəl/
- Basis: Direct
- Sources: [factual](https://www.oxfordlearnersdictionaries.com/definition/american_english/factual)

### fail-against-the-original-implementation-workflow

- Entry: [fail against the original implementation](../entries/fail-against-the-original-implementation-workflow.yml)
- IPA: /feɪl əˈɡɛnst ði əˈrɪdʒənl ˌɪmpləmənˈteɪʃn/
- Basis: Composed
- Review: Component citation forms; retain word-level stresses, without claiming a unique sentence-level accent.
- Sources: [fail](https://www.oxfordlearnersdictionaries.com/definition/american_english/fail_1); [against](https://www.oxfordlearnersdictionaries.com/definition/american_english/against); [the](https://www.oxfordlearnersdictionaries.com/definition/american_english/the); [original](https://www.oxfordlearnersdictionaries.com/definition/american_english/original_1); [implementation](https://www.oxfordlearnersdictionaries.com/definition/american_english/implement_1)

### fallback-tech

- Entry: [fallback](../entries/fallback.yml)
- IPA: /ˈfɔlbæk/
- Basis: Direct
- Sources: [fallback](https://www.oxfordlearnersdictionaries.com/definition/american_english/fallback)

### false-diagnostic-tech

- Entry: [false diagnostic](../entries/false-diagnostic-tech.yml)
- IPA: /fɔls ˌdaɪəɡˈnɑstɪk/
- Basis: Composed
- Review: Component citation forms; retain word-level stresses, without claiming a unique sentence-level accent.
- Sources: [false](https://www.oxfordlearnersdictionaries.com/definition/american_english/false); [diagnostic](https://www.oxfordlearnersdictionaries.com/definition/american_english/diagnostic_1); [diagnostic](https://dictionary.cambridge.org/dictionary/english/diagnostic)

### fast-path-tech

- Entry: [fast path](../entries/fast-path-tech.yml)
- IPA: /fæst pæθ/
- Basis: Composed
- Review: Component citation forms; retain word-level stresses, without claiming a unique sentence-level accent.
- Sources: [fast](https://www.oxfordlearnersdictionaries.com/definition/american_english/fast_1); [path](https://www.oxfordlearnersdictionaries.com/definition/american_english/path)

### field-presence-tech

- Entry: [field presence](../entries/field-presence-tech.yml)
- IPA: /fild ˈprɛzns/
- Basis: Composed
- Review: Component citation forms; retain word-level stresses, without claiming a unique sentence-level accent.
- Sources: [field](https://www.oxfordlearnersdictionaries.com/definition/american_english/field_1); [presence](https://www.oxfordlearnersdictionaries.com/definition/american_english/presence)

### finding-tech

- Entry: [finding](../entries/finding-tech.yml)
- IPA: /ˈfaɪndɪŋ/
- Basis: Direct
- Sources: [finding](https://www.oxfordlearnersdictionaries.com/definition/american_english/finding)

### fixture-tech

- Entry: [fixture](../entries/fixture-tech.yml)
- IPA: /ˈfɪkstʃər/
- Basis: Direct
- Sources: [fixture](https://www.oxfordlearnersdictionaries.com/definition/american_english/fixture)

### flow-analysis-tech

- Entry: [flow analysis](../entries/flow-analysis-tech.yml)
- IPA: /floʊ əˈnæləsəs/
- Basis: Composed
- Review: Component citation forms; retain word-level stresses, without claiming a unique sentence-level accent.
- Sources: [flow](https://www.oxfordlearnersdictionaries.com/definition/american_english/flow_1); [analysis](https://www.oxfordlearnersdictionaries.com/definition/american_english/analysis)

### followed-by-workflow

- Entry: [followed by](../entries/followed-by-workflow.yml)
- IPA: /ˈfɑloʊd baɪ/
- Basis: Composed
- Review: Component citation forms; retain word-level stresses, without claiming a unique sentence-level accent. followed: Follow + -ed /d/ after a vowel.
- Sources: [follow](https://www.oxfordlearnersdictionaries.com/definition/american_english/follow); [-ed](https://dictionary.cambridge.org/dictionary/english/ed); [by](https://www.oxfordlearnersdictionaries.com/definition/american_english/by_1)

### forwarding-tech

- Entry: [forwarding](../entries/forwarding-tech.yml)
- IPA: /ˈfɔrwərdɪŋ/
- Basis: Composed
- Review: Forward + -ing; retain stem stress.
- Sources: [forward](https://www.oxfordlearnersdictionaries.com/definition/american_english/forward_3); [-ing](https://dictionary.cambridge.org/dictionary/english/ing)

### generalize-general

- Entry: [generalize](../entries/generalize-general.yml)
- IPA: /ˈdʒɛnərəˌlaɪz/
- Basis: Direct
- Sources: [generalize](https://www.oxfordlearnersdictionaries.com/definition/american_english/generalize)

### generated-output-only-workflow

- Entry: [generated output only](../entries/generated-output-only-workflow.yml)
- IPA: /ˈdʒɛnəˌreɪtɪd ˈaʊtpʊt ˈoʊnli/
- Basis: Composed
- Review: Component citation forms; retain word-level stresses, without claiming a unique sentence-level accent. generated: Generate + -ed /ɪd/ after /t/.
- Sources: [generate](https://www.oxfordlearnersdictionaries.com/definition/american_english/generate); [-ed](https://dictionary.cambridge.org/dictionary/english/ed); [output](https://www.oxfordlearnersdictionaries.com/definition/american_english/output_1); [only](https://www.oxfordlearnersdictionaries.com/definition/american_english/only_1)

### genuine-tech

- Entry: [genuine](../entries/genuine-tech.yml)
- IPA: /ˈdʒɛnjuən/
- Basis: Direct
- Sources: [genuine](https://www.oxfordlearnersdictionaries.com/definition/american_english/genuine)

### getter-tech

- Entry: [getter](../entries/getter-tech.yml)
- IPA: /ˈɡɛtər/
- Basis: Composed
- Review: Get + agent suffix -er; broad /t/ rather than a narrow flap.
- Sources: [get](https://www.oxfordlearnersdictionaries.com/definition/american_english/get); [-er](https://dictionary.cambridge.org/dictionary/english/er)

### graceful-shutdown-tech

- Entry: [graceful shutdown](../entries/graceful-shutdown-tech.yml)
- IPA: /ˈɡreɪsfl ˈʃʌtdaʊn/
- Basis: Composed
- Review: Component citation forms; retain word-level stresses, without claiming a unique sentence-level accent.
- Sources: [graceful](https://www.oxfordlearnersdictionaries.com/definition/american_english/graceful); [shutdown](https://www.oxfordlearnersdictionaries.com/definition/american_english/shutdown)

### gratitude-general

- Entry: [gratitude](../entries/gratitude-general.yml)
- IPA: /ˈɡrætəˌtud/
- Basis: Direct
- Sources: [gratitude](https://www.oxfordlearnersdictionaries.com/definition/american_english/gratitude)

### guard-tech

- Entry: [guard](../entries/guard-tech.yml)
- IPA: /ɡɑrd/
- Basis: Direct
- Sources: [guard](https://www.oxfordlearnersdictionaries.com/definition/american_english/guard_1)

### handoff-workflow

- Entry: [handoff](../entries/handoff-workflow.yml)
- IPA: /ˈhændɔf/
- Basis: Direct
- Sources: [handoff](https://www.oxfordlearnersdictionaries.com/definition/american_english/handoff)

### harness-tech

- Entry: [harness](../entries/harness-tech.yml)
- IPA: /ˈhɑrnəs/
- Basis: Direct
- Sources: [harness](https://www.oxfordlearnersdictionaries.com/definition/american_english/harness_1)

### hashing-tech

- Entry: [hashing](../entries/hashing-tech.yml)
- IPA: /ˈhæʃɪŋ/
- Basis: Composed
- Review: Hash + -ing.
- Sources: [hash](https://www.oxfordlearnersdictionaries.com/definition/american_english/hash_2); [-ing](https://dictionary.cambridge.org/dictionary/english/ing)

### heritage-type-tech

- Entry: [heritage type](../entries/heritage-type-tech.yml)
- IPA: /ˈhɛrətɪdʒ taɪp/
- Basis: Composed
- Review: Component citation forms; retain word-level stresses, without claiming a unique sentence-level accent.
- Sources: [heritage](https://www.oxfordlearnersdictionaries.com/definition/american_english/heritage); [type](https://www.oxfordlearnersdictionaries.com/definition/american_english/type_1)

### identifier-tech

- Entry: [identifier](../entries/identifier-tech.yml)
- IPA: /aɪˈdɛntəˌfaɪər/
- Basis: Direct
- Sources: [identifier](https://www.oxfordlearnersdictionaries.com/definition/american_english/identifier)

### imperative-general

- Entry: [imperative](../entries/imperative-general.yml)
- IPA: /ɪmˈpɛrətɪv/
- Basis: Direct
- Sources: [imperative](https://www.oxfordlearnersdictionaries.com/definition/american_english/imperative_1)

### imply-general

- Entry: [imply](../entries/imply-general.yml)
- IPA: /ɪmˈplaɪ/
- Basis: Direct
- Sources: [imply](https://www.oxfordlearnersdictionaries.com/definition/american_english/imply)

### in-one-pass-workflow

- Entry: [in one pass](../entries/in-one-pass-workflow.yml)
- IPA: /ɪn wʌn pæs/
- Basis: Composed
- Review: Component citation forms; retain word-level stresses, without claiming a unique sentence-level accent.
- Sources: [in](https://www.oxfordlearnersdictionaries.com/definition/english/in_1); [one](https://www.oxfordlearnersdictionaries.com/definition/american_english/one_1); [pass](https://www.oxfordlearnersdictionaries.com/definition/american_english/pass_1)

### in-order-workflow

- Entry: [in order](../entries/in-order-workflow.yml)
- IPA: /ɪn ˈɔrdər/
- Basis: Composed
- Review: Component citation forms; retain word-level stresses, without claiming a unique sentence-level accent.
- Sources: [in](https://www.oxfordlearnersdictionaries.com/definition/english/in_1); [order](https://www.oxfordlearnersdictionaries.com/definition/american_english/order_1)

### inaccessible-tech

- Entry: [inaccessible](../entries/inaccessible-tech.yml)
- IPA: /ˌɪnəkˈsɛsəbl/
- Basis: Direct
- Sources: [inaccessible](https://www.oxfordlearnersdictionaries.com/definition/american_english/inaccessible)

### incremental-build-tech

- Entry: [incremental build](../entries/incremental-build-tech.yml)
- IPA: /ˌɪŋkrəˈmɛntl bɪld/
- Basis: Composed
- Review: Component citation forms; retain word-level stresses, without claiming a unique sentence-level accent.
- Sources: [incremental](https://www.oxfordlearnersdictionaries.com/definition/english/incremental); [build](https://www.oxfordlearnersdictionaries.com/definition/american_english/build_1)

### independently-tech

- Entry: [independently](../entries/independently-tech.yml)
- IPA: /ˌɪndɪˈpɛndəntli/
- Basis: Direct
- Sources: [independently](https://www.oxfordlearnersdictionaries.com/definition/english/independently)

### infer-tech

- Entry: [infer](../entries/infer-tech.yml)
- IPA: /ɪnˈfər/
- Basis: Direct
- Sources: [infer](https://www.oxfordlearnersdictionaries.com/definition/american_english/infer)

### inspect-tech

- Entry: [inspect](../entries/inspect-tech.yml)
- IPA: /ɪnˈspɛkt/
- Basis: Direct
- Sources: [inspect](https://www.oxfordlearnersdictionaries.com/definition/american_english/inspect)

### intact-tech

- Entry: [intact](../entries/intact-tech.yml)
- IPA: /ɪnˈtækt/
- Basis: Direct
- Sources: [intact](https://www.oxfordlearnersdictionaries.com/definition/american_english/intact)

### integration-test-tech

- Entry: [integration test](../entries/integration-test-tech.yml)
- IPA: /ˌɪntəˈɡreɪʃn tɛst/
- Basis: Composed
- Review: Component citation forms; retain word-level stresses, without claiming a unique sentence-level accent.
- Sources: [integration](https://www.oxfordlearnersdictionaries.com/definition/american_english/integration); [test](https://www.oxfordlearnersdictionaries.com/definition/american_english/test_1)

### internals-tech

- Entry: [internals](../entries/internals-tech.yml)
- IPA: /ɪnˈtərnlz/
- Basis: Composed
- Review: Nominalized plural of internal; retain stem stress and add /z/.
- Sources: [internal](https://www.oxfordlearnersdictionaries.com/definition/american_english/internal); [-s](https://dictionary.cambridge.org/dictionary/english/s)

### introduce-tech

- Entry: [introduce](../entries/introduce-tech.yml)
- IPA: /ˌɪntrəˈdus/
- Basis: Direct
- Sources: [introduce](https://www.oxfordlearnersdictionaries.com/definition/american_english/introduce)

### invalidate-tech

- Entry: [invalidate](../entries/invalidate-tech.yml)
- IPA: /ɪnˈvæləˌdeɪt/
- Basis: Direct
- Sources: [invalidate](https://www.oxfordlearnersdictionaries.com/definition/american_english/invalidate)

### invaluable-general

- Entry: [invaluable](../entries/invaluable-general.yml)
- IPA: /ɪnˈvæljəbl/
- Basis: Direct
- Sources: [invaluable](https://www.oxfordlearnersdictionaries.com/definition/american_english/invaluable)

### investigate-tech

- Entry: [investigate](../entries/investigate-tech.yml)
- IPA: /ɪnˈvɛstəˌɡeɪt/
- Basis: Direct
- Sources: [investigate](https://www.oxfordlearnersdictionaries.com/definition/american_english/investigate)

### invocation-workflow

- Entry: [invocation](../entries/invocation-workflow.yml)
- IPA: /ˌɪnvəˈkeɪʃn/
- Basis: Direct
- Sources: [invocation](https://www.oxfordlearnersdictionaries.com/definition/american_english/invocation)

### invoke-tech

- Entry: [invoke](../entries/invoke-tech.yml)
- IPA: /ɪnˈvoʊk/
- Basis: Direct
- Sources: [invoke](https://www.oxfordlearnersdictionaries.com/definition/american_english/invoke)

### isolated-worktree-workflow

- Entry: [isolated worktree](../entries/isolated-worktree-workflow.yml)
- IPA: /ˈaɪsəˌleɪtəd ˈwərkˌtri/
- Basis: Composed
- Review: Component citation forms; retain word-level stresses, without claiming a unique sentence-level accent. worktree: Closed noun compound: work + tree; primary stress on work.
- Sources: [isolated](https://www.oxfordlearnersdictionaries.com/definition/american_english/isolated); [work](https://www.oxfordlearnersdictionaries.com/definition/american_english/work_1); [tree](https://www.oxfordlearnersdictionaries.com/definition/american_english/tree)

### iterable-tech

- Entry: [iterable](../entries/iterable-tech.yml)
- IPA: unset
- Basis: Unresolved
- Review: Oxford and Cambridge lookups did not provide a matching American transcription. Wiktionary supplies IPA only for Spanish. The English stem change from iterate does not establish a complete verified pronunciation.
- Sources: [supplementary lookup (not American IPA evidence)](https://en.wiktionary.org/wiki/iterable#English)

### joiner-tech

- Entry: [joiner](../entries/joiner-tech.yml)
- IPA: /ˈdʒɔɪnər/
- Basis: Direct
- Sources: [joiner](https://www.oxfordlearnersdictionaries.com/definition/american_english/joiner)

### justify-general

- Entry: [justify](../entries/justify-general.yml)
- IPA: /ˈdʒʌstəˌfaɪ/
- Basis: Direct
- Sources: [justify](https://www.oxfordlearnersdictionaries.com/definition/american_english/justify)

### leave-untouched-workflow

- Entry: [leave untouched](../entries/leave-untouched-workflow.yml)
- IPA: /liv ʌnˈtʌtʃt/
- Basis: Composed
- Review: Component citation forms; retain word-level stresses, without claiming a unique sentence-level accent.
- Sources: [leave](https://www.oxfordlearnersdictionaries.com/definition/american_english/leave_1); [untouched](https://www.oxfordlearnersdictionaries.com/definition/american_english/untouched)

### legacy-tech

- Entry: [legacy](../entries/legacy-tech.yml)
- IPA: /ˈlɛɡəsi/
- Basis: Direct
- Sources: [legacy](https://www.oxfordlearnersdictionaries.com/definition/american_english/legacy_2)

### lifecycle-tech

- Entry: [lifecycle](../entries/lifecycle-tech.yml)
- IPA: /ˈlaɪf ˌsaɪkəl/
- Basis: Direct
- Review: Same spoken compound as dictionary spelling life cycle.
- Sources: [life cycle](https://dictionary.cambridge.org/dictionary/english/life-cycle?q=lifecycle)

### limitation-tech

- Entry: [limitation](../entries/limitation-tech.yml)
- IPA: /ˌlɪməˈteɪʃn/
- Basis: Direct
- Sources: [limitation](https://www.oxfordlearnersdictionaries.com/definition/american_english/limitation)

### linear-tech

- Entry: [linear](../entries/linear-tech.yml)
- IPA: /ˈlɪniər/
- Basis: Direct
- Sources: [linear](https://www.oxfordlearnersdictionaries.com/definition/american_english/linear)

### literal-programming

- Entry: [literal](../entries/literal-programming.yml)
- IPA: /ˈlɪtərəl/
- Basis: Composed
- Review: Noun use of the dictionary-transcribed adjective; retain its word stress (zero derivation).
- Sources: [literal](https://www.oxfordlearnersdictionaries.com/definition/american_english/literal)

### literal-text

- Entry: [literal](../entries/literal-text.yml)
- IPA: /ˈlɪtərəl/
- Basis: Direct
- Sources: [literal](https://www.oxfordlearnersdictionaries.com/definition/american_english/literal)

### lockfile-tech

- Entry: [lockfile](../entries/lockfile-tech.yml)
- IPA: /ˈlɑkˌfaɪl/
- Basis: Composed
- Review: Closed noun compound: lock + file; primary stress on lock.
- Sources: [lock](https://www.oxfordlearnersdictionaries.com/definition/american_english/lock_1); [file](https://www.oxfordlearnersdictionaries.com/definition/american_english/file_1)

### lone-surrogate-tech

- Entry: [lone surrogate](../entries/lone-surrogate-tech.yml)
- IPA: /loʊn ˈsərəɡət/
- Basis: Composed
- Review: Component citation forms; retain word-level stresses, without claiming a unique sentence-level accent.
- Sources: [lone](https://www.oxfordlearnersdictionaries.com/definition/american_english/lone); [surrogate](https://www.oxfordlearnersdictionaries.com/definition/american_english/surrogate)

### maintainable-tech

- Entry: [maintainable](../entries/maintainable-tech.yml)
- IPA: /meɪnˈteɪnəbl/
- Basis: Composed
- Review: Maintain + unstressed -able; retain stem stress.
- Sources: [maintain](https://www.oxfordlearnersdictionaries.com/definition/american_english/maintain); [-able](https://dictionary.cambridge.org/dictionary/english/able)

### malformed-tech

- Entry: [malformed](../entries/malformed-tech.yml)
- IPA: /ˌmælˈfɔrmd/
- Basis: Direct
- Sources: [malformed](https://www.oxfordlearnersdictionaries.com/definition/american_english/malformed)

### manifest-tech

- Entry: [manifest](../entries/manifest-tech.yml)
- IPA: /ˈmænəˌfɛst/
- Basis: Direct
- Sources: [manifest](https://www.oxfordlearnersdictionaries.com/definition/american_english/manifest_3)

### marker-tech

- Entry: [marker](../entries/marker-tech.yml)
- IPA: /ˈmɑrkər/
- Basis: Direct
- Sources: [marker](https://www.oxfordlearnersdictionaries.com/definition/american_english/marker)

### material-tech

- Entry: [material](../entries/material-tech.yml)
- IPA: /məˈtɪriəl/
- Basis: Direct
- Sources: [material](https://www.oxfordlearnersdictionaries.com/definition/american_english/material_2)

### merit-general

- Entry: [merit](../entries/merit-general.yml)
- IPA: /ˈmɛrət/
- Basis: Direct
- Sources: [merit](https://www.oxfordlearnersdictionaries.com/definition/american_english/merit_2)

### misalign-tech

- Entry: [misalign](../entries/misalign-tech.yml)
- IPA: /ˌmɪsəˈlaɪn/
- Basis: Direct
- Sources: [misalign](https://dictionary.cambridge.org/dictionary/english/misalign)

### mock-tech

- Entry: [mock](../entries/mock-tech.yml)
- IPA: /mɑk/
- Basis: Direct
- Sources: [mock](https://www.oxfordlearnersdictionaries.com/definition/american_english/mock_1)

### mount-tech

- Entry: [mount](../entries/mount-tech.yml)
- IPA: /maʊnt/
- Basis: Direct
- Sources: [mount](https://www.oxfordlearnersdictionaries.com/definition/american_english/mount_1)

### multibyte-tech

- Entry: [multibyte](../entries/multibyte-tech.yml)
- IPA: /ˌmʌltiˈbaɪt/
- Basis: Composed
- Review: Prefix multi- + byte; select the /mʌlti/ variant.
- Sources: [multi-](https://dictionary.cambridge.org/dictionary/english/multi); [byte](https://www.oxfordlearnersdictionaries.com/definition/american_english/byte)

### multiline-range-tech

- Entry: [multiline range](../entries/multiline-range-tech.yml)
- IPA: /ˌmʌltiˈlaɪn reɪndʒ/
- Basis: Composed
- Review: Component citation forms; retain word-level stresses, without claiming a unique sentence-level accent.
- Sources: [multi-line](https://dictionary.cambridge.org/dictionary/english/multi-line?q=multiline); [range](https://www.oxfordlearnersdictionaries.com/definition/american_english/range_1)

### mutable-tech

- Entry: [mutable](../entries/mutable-tech.yml)
- IPA: /ˈmjutəbl/
- Basis: Direct
- Sources: [mutable](https://www.oxfordlearnersdictionaries.com/definition/american_english/mutable)

### native-addon-tech

- Entry: [native addon](../entries/native-addon-tech.yml)
- IPA: /ˈneɪtɪv ˈædɑn/
- Basis: Composed
- Review: Component citation forms; retain word-level stresses, without claiming a unique sentence-level accent.
- Sources: [native](https://www.oxfordlearnersdictionaries.com/definition/american_english/native_1); [add-on](https://dictionary.cambridge.org/dictionary/english/add-on)

### negative-lookbehind-tech

- Entry: [negative lookbehind](../entries/negative-lookbehind-tech.yml)
- IPA: /ˈnɛɡətɪv ˈlʊkbɪˌhaɪnd/
- Basis: Composed
- Review: Component citation forms; retain word-level stresses, without claiming a unique sentence-level accent. lookbehind: Technical noun compound: look + behind; primary stress on look.
- Sources: [negative](https://www.oxfordlearnersdictionaries.com/definition/american_english/negative_1); [look](https://www.oxfordlearnersdictionaries.com/definition/american_english/look_1); [behind](https://www.oxfordlearnersdictionaries.com/definition/american_english/behind_1)

### nested-tech

- Entry: [nested](../entries/nested-tech.yml)
- IPA: /ˈnɛstɪd/
- Basis: Composed
- Review: Nest + -ed /ɪd/ after /t/.
- Sources: [nest](https://www.oxfordlearnersdictionaries.com/definition/american_english/nest_1); [-ed](https://dictionary.cambridge.org/dictionary/english/ed)

### nonthrowing-tech

- Entry: [nonthrowing](../entries/nonthrowing-tech.yml)
- IPA: /ˌnɑnˈθroʊɪŋ/
- Basis: Composed
- Review: Non- + throw + -ing; retain stress on throw.
- Sources: [non-](https://dictionary.cambridge.org/dictionary/english/non); [throw](https://www.oxfordlearnersdictionaries.com/definition/american_english/throw_1); [-ing](https://dictionary.cambridge.org/dictionary/english/ing)

### observed-result-workflow

- Entry: [observed result](../entries/observed-result-workflow.yml)
- IPA: /əbˈzərvd rɪˈzʌlt/
- Basis: Composed
- Review: Component citation forms; retain word-level stresses, without claiming a unique sentence-level accent. observed: Observe + -ed /d/ after /v/.
- Sources: [observe](https://www.oxfordlearnersdictionaries.com/definition/american_english/observe); [-ed](https://dictionary.cambridge.org/dictionary/english/ed); [result](https://www.oxfordlearnersdictionaries.com/definition/american_english/result_1)

### occupied-workflow

- Entry: [occupied](../entries/occupied-workflow.yml)
- IPA: /ˈɑkjəˌpaɪd/
- Basis: Direct
- Sources: [occupied](https://www.oxfordlearnersdictionaries.com/definition/american_english/occupied)

### omit-tech

- Entry: [omit](../entries/omit-tech.yml)
- IPA: /oʊˈmɪt/
- Basis: Direct
- Sources: [omit](https://www.oxfordlearnersdictionaries.com/definition/american_english/omit)

### on-every-exit-path-workflow

- Entry: [on every exit path](../entries/on-every-exit-path-workflow.yml)
- IPA: /ɑn ˈɛvri ˈɛɡzət pæθ/
- Basis: Composed
- Review: Component citation forms; retain word-level stresses, without claiming a unique sentence-level accent.
- Sources: [on](https://www.oxfordlearnersdictionaries.com/definition/american_english/on_1); [every](https://www.oxfordlearnersdictionaries.com/definition/american_english/every); [exit](https://www.oxfordlearnersdictionaries.com/definition/american_english/exit_1); [path](https://www.oxfordlearnersdictionaries.com/definition/american_english/path)

### out-of-order-tech

- Entry: [out-of-order](../entries/out-of-order-tech.yml)
- IPA: /ˌaʊt əv ˈɔrdər/
- Basis: Composed
- Review: Compound modifier: primary stress on order; ordinary unstressed of /əv/.
- Sources: [out](https://www.oxfordlearnersdictionaries.com/definition/american_english/out_1); [of](https://www.oxfordlearnersdictionaries.com/definition/american_english/of); [order](https://www.oxfordlearnersdictionaries.com/definition/american_english/order_1)

### outstanding-tech

- Entry: [outstanding](../entries/outstanding-tech.yml)
- IPA: /aʊtˈstændɪŋ/
- Basis: Direct
- Sources: [outstanding](https://www.oxfordlearnersdictionaries.com/definition/american_english/outstanding)

### overlap-prevention-workflow

- Entry: [overlap prevention](../entries/overlap-prevention-workflow.yml)
- IPA: /ˈoʊvərˌlæp prɪˈvɛnʃn/
- Basis: Composed
- Review: Component citation forms; retain word-level stresses, without claiming a unique sentence-level accent.
- Sources: [overlap](https://www.oxfordlearnersdictionaries.com/definition/american_english/overlap_2); [prevention](https://www.oxfordlearnersdictionaries.com/definition/american_english/prevention)

### overlap-tech

- Entry: [overlap](../entries/overlap-tech.yml)
- IPA: /ˌoʊvərˈlæp/
- Basis: Direct
- Sources: [overlap](https://www.oxfordlearnersdictionaries.com/definition/american_english/overlap_1)

### overload-tech

- Entry: [overload](../entries/overload-tech.yml)
- IPA: /ˈoʊvərˌloʊd/
- Basis: Direct
- Sources: [overload](https://www.oxfordlearnersdictionaries.com/definition/american_english/overload_2)

### overwrite-tech

- Entry: [overwrite](../entries/overwrite-tech.yml)
- IPA: /ˌoʊvərˈraɪt/
- Basis: Direct
- Sources: [overwrite](https://www.oxfordlearnersdictionaries.com/definition/american_english/overwrite)

### ownership-token-workflow

- Entry: [ownership token](../entries/ownership-token-workflow.yml)
- IPA: /ˈoʊnərˌʃɪp ˈtoʊkən/
- Basis: Composed
- Review: Component citation forms; retain word-level stresses, without claiming a unique sentence-level accent.
- Sources: [ownership](https://www.oxfordlearnersdictionaries.com/definition/american_english/ownership); [token](https://www.oxfordlearnersdictionaries.com/definition/american_english/token_1)

### pass-with-the-fix-workflow

- Entry: [pass with the fix](../entries/pass-with-the-fix-workflow.yml)
- IPA: /pæs wɪθ ðə fɪks/
- Basis: Composed
- Review: Component citation forms; retain word-level stresses, without claiming a unique sentence-level accent.
- Sources: [pass](https://www.oxfordlearnersdictionaries.com/definition/american_english/pass_1); [with](https://www.oxfordlearnersdictionaries.com/definition/american_english/with); [the](https://www.oxfordlearnersdictionaries.com/definition/american_english/the); [fix](https://www.oxfordlearnersdictionaries.com/definition/american_english/fix_1)

### pending-tech

- Entry: [pending](../entries/pending-tech.yml)
- IPA: /ˈpɛndɪŋ/
- Basis: Direct
- Sources: [pending](https://www.oxfordlearnersdictionaries.com/definition/american_english/pending_2)

### pinned-checkout-workflow

- Entry: [pinned checkout](../entries/pinned-checkout-workflow.yml)
- IPA: /pɪnd ˈtʃɛkaʊt/
- Basis: Composed
- Review: Component citation forms; retain word-level stresses, without claiming a unique sentence-level accent. pinned: Pin + -ed /d/ after /n/.
- Sources: [pin](https://www.oxfordlearnersdictionaries.com/definition/american_english/pin_1); [-ed](https://dictionary.cambridge.org/dictionary/english/ed); [checkout](https://www.oxfordlearnersdictionaries.com/definition/american_english/checkout)

### pre-existing-tech

- Entry: [pre-existing](../entries/pre-existing-tech.yml)
- IPA: /ˌpri ɪɡˈzɪstɪŋ/
- Basis: Direct
- Sources: [pre-existing](https://www.oxfordlearnersdictionaries.com/definition/english/pre-existing)

### precede-tech

- Entry: [precede](../entries/precede-tech.yml)
- IPA: /prɪˈsid/
- Basis: Direct
- Sources: [precede](https://www.oxfordlearnersdictionaries.com/definition/american_english/precede)

### prefilter-tech

- Entry: [prefilter](../entries/prefilter-tech.yml)
- IPA: /ˈpriˌfɪltər/
- Basis: Composed
- Review: Noun compound pre- + filter; initial primary stress, distinct from a verb reading.
- Sources: [pre-](https://dictionary.cambridge.org/dictionary/english/pre); [filter](https://www.oxfordlearnersdictionaries.com/definition/american_english/filter_1)

### preserve-tech

- Entry: [preserve](../entries/preserve-tech.yml)
- IPA: /prɪˈzərv/
- Basis: Direct
- Sources: [preserve](https://www.oxfordlearnersdictionaries.com/definition/american_english/preserve_1)

### prevent-tech

- Entry: [prevent](../entries/prevent-tech.yml)
- IPA: /prɪˈvɛnt/
- Basis: Direct
- Sources: [prevent](https://www.oxfordlearnersdictionaries.com/definition/american_english/prevent)

### primary-checkout-workflow

- Entry: [primary checkout](../entries/primary-checkout-workflow.yml)
- IPA: /ˈpraɪˌmɛri ˈtʃɛkaʊt/
- Basis: Composed
- Review: Component citation forms; retain word-level stresses, without claiming a unique sentence-level accent.
- Sources: [primary](https://www.oxfordlearnersdictionaries.com/definition/american_english/primary_1); [checkout](https://www.oxfordlearnersdictionaries.com/definition/american_english/checkout)

### probe-tech

- Entry: [probe](../entries/probe-tech.yml)
- IPA: /proʊb/
- Basis: Direct
- Sources: [probe](https://www.oxfordlearnersdictionaries.com/definition/american_english/probe_2)

### profiling-tech

- Entry: [profiling](../entries/profiling-tech.yml)
- IPA: /ˈproʊfaɪlɪŋ/
- Basis: Direct
- Sources: [profiling](https://www.oxfordlearnersdictionaries.com/definition/american_english/profiling)

### propagate-tech

- Entry: [propagate](../entries/propagate-tech.yml)
- IPA: /ˈprɑpəˌɡeɪt/
- Basis: Direct
- Sources: [propagate](https://www.oxfordlearnersdictionaries.com/definition/american_english/propagate)

### prune-tech

- Entry: [prune](../entries/prune-tech.yml)
- IPA: /prun/
- Basis: Direct
- Sources: [prune](https://www.oxfordlearnersdictionaries.com/definition/american_english/prune_2)

### rather-than-workflow

- Entry: [rather than](../entries/rather-than-workflow.yml)
- IPA: /ˈræðər ðən/
- Basis: Composed
- Review: Component citation forms; retain word-level stresses, without claiming a unique sentence-level accent.
- Sources: [rather](https://www.oxfordlearnersdictionaries.com/definition/american_english/rather); [than](https://www.oxfordlearnersdictionaries.com/definition/american_english/than)

### re-entry-tech

- Entry: [re-entry](../entries/re-entry-tech.yml)
- IPA: /ˌriˈɛntri/
- Basis: Direct
- Sources: [re-entry](https://www.oxfordlearnersdictionaries.com/definition/american_english/re-entry)

### read-only-workflow

- Entry: [read-only](../entries/read-only-workflow.yml)
- IPA: /ˌridˈoʊnli/
- Basis: Direct
- Sources: [read-only](https://dictionary.cambridge.org/dictionary/english/read-only)

### ready-for-review-workflow

- Entry: [ready for review](../entries/ready-for-review-workflow.yml)
- IPA: /ˈrɛdi fɔr rɪˈvju/
- Basis: Composed
- Review: Component citation forms; retain word-level stresses, without claiming a unique sentence-level accent.
- Sources: [ready](https://www.oxfordlearnersdictionaries.com/definition/american_english/ready_1); [for](https://www.oxfordlearnersdictionaries.com/definition/american_english/for_1); [review](https://www.oxfordlearnersdictionaries.com/definition/american_english/review_1)

### recovery-tech

- Entry: [recovery](../entries/recovery-tech.yml)
- IPA: /rɪˈkʌvəri/
- Basis: Direct
- Sources: [recovery](https://www.oxfordlearnersdictionaries.com/definition/american_english/recovery)

### register-tech

- Entry: [register](../entries/register-tech.yml)
- IPA: /ˈrɛdʒəstər/
- Basis: Direct
- Sources: [register](https://www.oxfordlearnersdictionaries.com/definition/american_english/register_1)

### registry-tech

- Entry: [registry](../entries/registry-tech.yml)
- IPA: /ˈrɛdʒəstri/
- Basis: Direct
- Sources: [registry](https://www.oxfordlearnersdictionaries.com/definition/american_english/registry)

### regression-tech

- Entry: [regression](../entries/regression-tech.yml)
- IPA: /rɪˈɡrɛʃn/
- Basis: Direct
- Sources: [regression](https://www.oxfordlearnersdictionaries.com/definition/american_english/regression)

### regression-test-tech

- Entry: [regression test](../entries/regression-test-tech.yml)
- IPA: /rɪˈɡrɛʃn tɛst/
- Basis: Composed
- Review: Component citation forms; retain word-level stresses, without claiming a unique sentence-level accent.
- Sources: [regression](https://www.oxfordlearnersdictionaries.com/definition/american_english/regression); [test](https://www.oxfordlearnersdictionaries.com/definition/american_english/test_1)

### reject-tech

- Entry: [reject](../entries/reject-tech.yml)
- IPA: /rɪˈdʒɛkt/
- Basis: Direct
- Sources: [reject](https://www.oxfordlearnersdictionaries.com/definition/american_english/reject_1)

### release-workflow

- Entry: [release](../entries/release-workflow.yml)
- IPA: /rɪˈlis/
- Basis: Direct
- Sources: [release](https://www.oxfordlearnersdictionaries.com/definition/american_english/release_1)

### remain-intact-workflow

- Entry: [remain intact](../entries/remain-intact-workflow.yml)
- IPA: /rɪˈmeɪn ɪnˈtækt/
- Basis: Composed
- Review: Component citation forms; retain word-level stresses, without claiming a unique sentence-level accent.
- Sources: [remain](https://www.oxfordlearnersdictionaries.com/definition/american_english/remain); [intact](https://www.oxfordlearnersdictionaries.com/definition/american_english/intact)

### remain-unchanged-workflow

- Entry: [remain unchanged](../entries/remain-unchanged-workflow.yml)
- IPA: /rɪˈmeɪn ʌnˈtʃeɪndʒd/
- Basis: Composed
- Review: Component citation forms; retain word-level stresses, without claiming a unique sentence-level accent.
- Sources: [remain](https://www.oxfordlearnersdictionaries.com/definition/american_english/remain); [unchanged](https://www.oxfordlearnersdictionaries.com/definition/american_english/unchanged)

### remote-tracking-branch-git

- Entry: [remote-tracking branch](../entries/remote-tracking-branch-git.yml)
- IPA: /rɪˈmoʊt ˈtrækɪŋ bræntʃ/
- Basis: Composed
- Review: Component citation forms; retain word-level stresses, without claiming a unique sentence-level accent. remote-tracking: Remote + track + -ing; retain component word stresses.
- Sources: [remote](https://www.oxfordlearnersdictionaries.com/definition/american_english/remote_1); [track](https://www.oxfordlearnersdictionaries.com/definition/american_english/track_2); [-ing](https://dictionary.cambridge.org/dictionary/english/ing); [branch](https://www.oxfordlearnersdictionaries.com/definition/american_english/branch_1)

### reported-behavior-workflow

- Entry: [reported behavior](../entries/reported-behavior-workflow.yml)
- IPA: /rɪˈpɔrtɪd bɪˈheɪvjər/
- Basis: Composed
- Review: Component citation forms; retain word-level stresses, without claiming a unique sentence-level accent. reported: Report + -ed /ɪd/ after /t/.
- Sources: [report](https://www.oxfordlearnersdictionaries.com/definition/american_english/report_1); [-ed](https://dictionary.cambridge.org/dictionary/english/ed); [behavior](https://www.oxfordlearnersdictionaries.com/definition/american_english/behavior)

### reproduce-tech

- Entry: [reproduce](../entries/reproduce-tech.yml)
- IPA: /ˌriprəˈdus/
- Basis: Direct
- Sources: [reproduce](https://www.oxfordlearnersdictionaries.com/definition/american_english/reproduce)

### reproducible-tech

- Entry: [reproducible](../entries/reproducible-tech.yml)
- IPA: /ˌriprəˈdusəbl/
- Basis: Direct
- Sources: [reproducible](https://www.oxfordlearnersdictionaries.com/definition/english/reproducible)

### reserve-workflow

- Entry: [reserve](../entries/reserve-workflow.yml)
- IPA: /rɪˈzərv/
- Basis: Direct
- Sources: [reserve](https://www.oxfordlearnersdictionaries.com/definition/american_english/reserve_2)

### reserved-prefix-tech

- Entry: [reserved prefix](../entries/reserved-prefix-tech.yml)
- IPA: /rɪˈzərvd ˈprifɪks/
- Basis: Composed
- Review: Component citation forms; retain word-level stresses, without claiming a unique sentence-level accent.
- Sources: [reserved](https://www.oxfordlearnersdictionaries.com/definition/american_english/reserved); [prefix](https://www.oxfordlearnersdictionaries.com/definition/american_english/prefix_1)

### resolve-tech

- Entry: [resolve](../entries/resolve-tech.yml)
- IPA: /rɪˈzɑlv/
- Basis: Direct
- Sources: [resolve](https://www.oxfordlearnersdictionaries.com/definition/american_english/resolve_1)

### resolve-workflow

- Entry: [resolve](../entries/resolve-workflow.yml)
- IPA: /rɪˈzɑlv/
- Basis: Direct
- Sources: [resolve](https://www.oxfordlearnersdictionaries.com/definition/american_english/resolve_1)

### restore-tech

- Entry: [restore](../entries/restore-tech.yml)
- IPA: /rɪˈstɔr/
- Basis: Direct
- Sources: [restore](https://www.oxfordlearnersdictionaries.com/definition/american_english/restore)

### restrict-tech

- Entry: [restrict](../entries/restrict-tech.yml)
- IPA: /rɪˈstrɪkt/
- Basis: Direct
- Sources: [restrict](https://www.oxfordlearnersdictionaries.com/definition/american_english/restrict)

### resumption-workflow

- Entry: [resumption](../entries/resumption-workflow.yml)
- IPA: /rɪˈzʌmpʃn/
- Basis: Direct
- Sources: [resumption](https://www.oxfordlearnersdictionaries.com/definition/american_english/resumption)

### retain-tech

- Entry: [retain](../entries/retain-tech.yml)
- IPA: /rɪˈteɪn/
- Basis: Direct
- Sources: [retain](https://www.oxfordlearnersdictionaries.com/definition/american_english/retain)

### revision-general

- Entry: [revision](../entries/revision-general.yml)
- IPA: /rɪˈvɪʒn/
- Basis: Direct
- Sources: [revision](https://www.oxfordlearnersdictionaries.com/definition/american_english/revision)

### revision-git

- Entry: [revision](../entries/revision-git.yml)
- IPA: /rɪˈvɪʒn/
- Basis: Direct
- Sources: [revision](https://www.oxfordlearnersdictionaries.com/definition/american_english/revision)

### round-trip-tech

- Entry: [round trip](../entries/round-trip-tech.yml)
- IPA: /ˌraʊnd ˈtrɪp/
- Basis: Direct
- Sources: [round trip](https://www.oxfordlearnersdictionaries.com/definition/english/round-trip_1)

### run-budget-workflow

- Entry: [run budget](../entries/run-budget-workflow.yml)
- IPA: /rʌn ˈbʌdʒət/
- Basis: Composed
- Review: Component citation forms; retain word-level stresses, without claiming a unique sentence-level accent.
- Sources: [run](https://www.oxfordlearnersdictionaries.com/definition/american_english/run_1); [budget](https://www.oxfordlearnersdictionaries.com/definition/american_english/budget_1)

### run-in-parallel-workflow

- Entry: [run in parallel](../entries/run-in-parallel-workflow.yml)
- IPA: /rʌn ɪn ˈpærəˌlɛl/
- Basis: Composed
- Review: Component citation forms; retain word-level stresses, without claiming a unique sentence-level accent.
- Sources: [run](https://www.oxfordlearnersdictionaries.com/definition/american_english/run_1); [in](https://www.oxfordlearnersdictionaries.com/definition/english/in_1); [parallel](https://www.oxfordlearnersdictionaries.com/definition/american_english/parallel_1)

### runtime-evaluation-tech

- Entry: [runtime evaluation](../entries/runtime-evaluation-tech.yml)
- IPA: /ˈrʌntaɪm ɪˌvæljuˈeɪʃn/
- Basis: Composed
- Review: Component citation forms; retain word-level stresses, without claiming a unique sentence-level accent.
- Sources: [runtime](https://www.oxfordlearnersdictionaries.com/definition/english/runtime); [evaluation](https://www.oxfordlearnersdictionaries.com/definition/american_english/evaluate)

### runtime-tech

- Entry: [runtime](../entries/runtime-tech.yml)
- IPA: /ˈrʌntaɪm/
- Basis: Direct
- Sources: [runtime](https://www.oxfordlearnersdictionaries.com/definition/english/runtime)

### sanitize-tech

- Entry: [sanitize](../entries/sanitize-tech.yml)
- IPA: /ˈsænəˌtaɪz/
- Basis: Direct
- Sources: [sanitize](https://www.oxfordlearnersdictionaries.com/definition/american_english/sanitize)

### scope-lexical

- Entry: [scope](../entries/scope-lexical.yml)
- IPA: /skoʊp/
- Basis: Direct
- Sources: [scope](https://www.oxfordlearnersdictionaries.com/definition/american_english/scope_1)

### scope-workflow

- Entry: [scope](../entries/scope-workflow.yml)
- IPA: /skoʊp/
- Basis: Direct
- Sources: [scope](https://www.oxfordlearnersdictionaries.com/definition/american_english/scope_1)

### serialization-tech

- Entry: [serialization](../entries/serialization-tech.yml)
- IPA: /ˌsɪriələˈzeɪʃn/
- Basis: Direct
- Sources: [serialization](https://www.oxfordlearnersdictionaries.com/definition/english/serialization)

### setup-tech

- Entry: [setup](../entries/setup-tech.yml)
- IPA: /ˈsɛtʌp/
- Basis: Direct
- Sources: [setup](https://www.oxfordlearnersdictionaries.com/definition/american_english/setup)

### signed-zero-tech

- Entry: [signed zero](../entries/signed-zero-tech.yml)
- IPA: /saɪnd ˈzɪroʊ/
- Basis: Composed
- Review: Component citation forms; retain word-level stresses, without claiming a unique sentence-level accent. signed: Sign + -ed /d/ after /n/; silent written g remains silent.
- Sources: [sign](https://www.oxfordlearnersdictionaries.com/definition/american_english/sign_2); [-ed](https://dictionary.cambridge.org/dictionary/english/ed); [zero](https://www.oxfordlearnersdictionaries.com/definition/american_english/zero_1)

### snapshot-tech

- Entry: [snapshot](../entries/snapshot-tech.yml)
- IPA: /ˈsnæpʃɑt/
- Basis: Direct
- Sources: [snapshot](https://www.oxfordlearnersdictionaries.com/definition/american_english/snapshot)

### source-map-tech

- Entry: [source map](../entries/source-map-tech.yml)
- IPA: /sɔrs mæp/
- Basis: Composed
- Review: Component citation forms; retain word-level stresses, without claiming a unique sentence-level accent.
- Sources: [source](https://www.oxfordlearnersdictionaries.com/definition/american_english/source_1); [map](https://www.oxfordlearnersdictionaries.com/definition/american_english/map_1)

### speculative-tech

- Entry: [speculative](../entries/speculative-tech.yml)
- IPA: /ˈspɛkjələtɪv/
- Basis: Direct
- Sources: [speculative](https://www.oxfordlearnersdictionaries.com/definition/american_english/speculative)

### staged-git

- Entry: [staged](../entries/staged-git.yml)
- IPA: /steɪdʒd/
- Basis: Direct
- Sources: [staged](https://www.oxfordlearnersdictionaries.com/definition/english/staged)

### stale-tech

- Entry: [stale](../entries/stale-tech.yml)
- IPA: /steɪl/
- Basis: Direct
- Sources: [stale](https://www.oxfordlearnersdictionaries.com/definition/american_english/stale)

### static-initialization-tech

- Entry: [static initialization](../entries/static-initialization-tech.yml)
- IPA: /ˈstætɪk ɪˌnɪʃələˈzeɪʃn/
- Basis: Composed
- Review: Component citation forms; retain word-level stresses, without claiming a unique sentence-level accent.
- Sources: [static](https://www.oxfordlearnersdictionaries.com/definition/american_english/static_1); [initialization](https://www.oxfordlearnersdictionaries.com/definition/american_english/initialize)

### static-inspection-tech

- Entry: [static inspection](../entries/static-inspection-tech.yml)
- IPA: /ˈstætɪk ɪnˈspɛkʃn/
- Basis: Composed
- Review: Component citation forms; retain word-level stresses, without claiming a unique sentence-level accent.
- Sources: [static](https://www.oxfordlearnersdictionaries.com/definition/american_english/static_1); [inspection](https://www.oxfordlearnersdictionaries.com/definition/american_english/inspection)

### submission-workflow

- Entry: [submission](../entries/submission-workflow.yml)
- IPA: /səbˈmɪʃn/
- Basis: Direct
- Sources: [submission](https://www.oxfordlearnersdictionaries.com/definition/american_english/submission)

### subsequent-tech

- Entry: [subsequent](../entries/subsequent-tech.yml)
- IPA: /ˈsʌbsəkwənt/
- Basis: Direct
- Sources: [subsequent](https://www.oxfordlearnersdictionaries.com/definition/american_english/subsequent)

### support-general

- Entry: [support](../entries/support-general.yml)
- IPA: /səˈpɔrt/
- Basis: Direct
- Sources: [support](https://www.oxfordlearnersdictionaries.com/definition/american_english/support_1)

### surrogate-pair-tech

- Entry: [surrogate pair](../entries/surrogate-pair-tech.yml)
- IPA: /ˈsərəɡət pɛr/
- Basis: Composed
- Review: Component citation forms; retain word-level stresses, without claiming a unique sentence-level accent.
- Sources: [surrogate](https://www.oxfordlearnersdictionaries.com/definition/american_english/surrogate); [pair](https://www.oxfordlearnersdictionaries.com/definition/american_english/pair_1)

### suspected-tech

- Entry: [suspected](../entries/suspected-tech.yml)
- IPA: /səˈspɛktɪd/
- Basis: Direct
- Sources: [suspected](https://www.oxfordlearnersdictionaries.com/definition/english/suspected)

### symlink-tech

- Entry: [symlink](../entries/symlink-tech.yml)
- IPA: unset
- Basis: Unresolved
- Review: Oxford and Cambridge lookups did not provide a matching American transcription. Wiktionary lists UK audio but no American IPA. A transcription of symbolic link would not verify the clipped word.
- Sources: [supplementary lookup (not American IPA evidence)](https://en.wiktionary.org/wiki/symlink#English)

### synchronize-tech

- Entry: [synchronize](../entries/synchronize-tech.yml)
- IPA: /ˈsɪŋkrəˌnaɪz/
- Basis: Direct
- Sources: [synchronize](https://www.oxfordlearnersdictionaries.com/definition/american_english/synchronize)

### synthesize-tech

- Entry: [synthesize](../entries/synthesize-tech.yml)
- IPA: /ˈsɪnθəˌsaɪz/
- Basis: Direct
- Sources: [synthesize](https://www.oxfordlearnersdictionaries.com/definition/american_english/synthesize)

### synthetic-tech

- Entry: [synthetic](../entries/synthetic-tech.yml)
- IPA: /sɪnˈθɛtɪk/
- Basis: Direct
- Sources: [synthetic](https://www.oxfordlearnersdictionaries.com/definition/american_english/synthetic_1)

### test-suite-tech

- Entry: [test suite](../entries/test-suite-tech.yml)
- IPA: /tɛst swit/
- Basis: Composed
- Review: Component citation forms; retain word-level stresses, without claiming a unique sentence-level accent.
- Sources: [test](https://www.oxfordlearnersdictionaries.com/definition/american_english/test_1); [suite](https://www.oxfordlearnersdictionaries.com/definition/american_english/suite)

### toggle-tech

- Entry: [toggle](../entries/toggle-tech.yml)
- IPA: /ˈtɑɡl/
- Basis: Direct
- Sources: [toggle](https://www.oxfordlearnersdictionaries.com/definition/american_english/toggle_2)

### toolchain-tech

- Entry: [toolchain](../entries/toolchain-tech.yml)
- IPA: /ˈtulˌtʃeɪn/
- Basis: Composed
- Review: Closed noun compound: tool + chain; primary stress on tool.
- Sources: [tool](https://www.oxfordlearnersdictionaries.com/definition/american_english/tool_1); [chain](https://www.oxfordlearnersdictionaries.com/definition/american_english/chain_1)

### tradeoff-general

- Entry: [tradeoff](../entries/tradeoff-general.yml)
- IPA: /ˈtreɪdˌɔf/
- Basis: Direct
- Sources: [tradeoff](https://dictionary.cambridge.org/dictionary/english/tradeoff)

### transformation-tech

- Entry: [transformation](../entries/transformation-tech.yml)
- IPA: /ˌtrænsfərˈmeɪʃn/
- Basis: Direct
- Sources: [transformation](https://www.oxfordlearnersdictionaries.com/definition/american_english/transformation)

### transport-failure-tech

- Entry: [transport failure](../entries/transport-failure-tech.yml)
- IPA: /ˈtrænspɔrt ˈfeɪljər/
- Basis: Composed
- Review: Component citation forms; retain word-level stresses, without claiming a unique sentence-level accent.
- Sources: [transport](https://www.oxfordlearnersdictionaries.com/definition/american_english/transport_2); [failure](https://www.oxfordlearnersdictionaries.com/definition/american_english/failure)

### transport-metadata-tech

- Entry: [transport metadata](../entries/transport-metadata-tech.yml)
- IPA: /ˈtrænspɔrt ˈmɛtəˌdeɪtə/
- Basis: Composed
- Review: Component citation forms; retain word-level stresses, without claiming a unique sentence-level accent.
- Sources: [transport](https://www.oxfordlearnersdictionaries.com/definition/american_english/transport_2); [metadata](https://www.oxfordlearnersdictionaries.com/definition/american_english/metadata)

### traversal-tech

- Entry: [traversal](../entries/traversal-tech.yml)
- IPA: /trəˈvərsəl/
- Basis: Composed
- Review: Traverse + unstressed nominal suffix -al; retain the American stem stress.
- Sources: [traverse](https://www.oxfordlearnersdictionaries.com/definition/american_english/traverse_1); [-al](https://dictionary.cambridge.org/dictionary/english/al)

### trigger-tech

- Entry: [trigger](../entries/trigger-tech.yml)
- IPA: /ˈtrɪɡər/
- Basis: Direct
- Sources: [trigger](https://www.oxfordlearnersdictionaries.com/definition/american_english/trigger_1)

### truncated-tech

- Entry: [truncated](../entries/truncated-tech.yml)
- IPA: /ˈtrʌŋkeɪtɪd/
- Basis: Direct
- Sources: [truncated](https://dictionary.cambridge.org/dictionary/english/truncated)

### type-argument-tech

- Entry: [type argument](../entries/type-argument-tech.yml)
- IPA: /taɪp ˈɑrɡjəmənt/
- Basis: Composed
- Review: Component citation forms; retain word-level stresses, without claiming a unique sentence-level accent.
- Sources: [type](https://www.oxfordlearnersdictionaries.com/definition/american_english/type_1); [argument](https://www.oxfordlearnersdictionaries.com/definition/american_english/argument)

### type-assertion-tech

- Entry: [type assertion](../entries/type-assertion-tech.yml)
- IPA: /taɪp əˈsərʃn/
- Basis: Composed
- Review: Component citation forms; retain word-level stresses, without claiming a unique sentence-level accent.
- Sources: [type](https://www.oxfordlearnersdictionaries.com/definition/american_english/type_1); [assertion](https://www.oxfordlearnersdictionaries.com/definition/american_english/assertion)

### typo-general

- Entry: [typo](../entries/typo-general.yml)
- IPA: /ˈtaɪpoʊ/
- Basis: Direct
- Sources: [typo](https://www.oxfordlearnersdictionaries.com/definition/american_english/typo)

### unclaimed-workflow

- Entry: [unclaimed](../entries/unclaimed-workflow.yml)
- IPA: /ʌnˈkleɪmd/
- Basis: Direct
- Sources: [unclaimed](https://www.oxfordlearnersdictionaries.com/definition/american_english/unclaimed)

### underline-tech

- Entry: [underline](../entries/underline-tech.yml)
- IPA: /ˈʌndərˌlaɪn/
- Basis: Composed
- Review: Noun use of the dictionary-transcribed verb; retain its word stress (zero derivation).
- Sources: [underline](https://www.oxfordlearnersdictionaries.com/definition/american_english/underline)

### uninitialized-tech

- Entry: [uninitialized](../entries/uninitialized-tech.yml)
- IPA: /ˌʌnɪˈnɪʃəˌlaɪzd/
- Basis: Composed
- Review: Un- + initialize + -ed /d/ after /z/.
- Sources: [un-](https://dictionary.cambridge.org/dictionary/english/un); [initialize](https://www.oxfordlearnersdictionaries.com/definition/american_english/initialize); [-ed](https://dictionary.cambridge.org/dictionary/english/ed)

### union-tech

- Entry: [union](../entries/union-tech.yml)
- IPA: /ˈjunjən/
- Basis: Direct
- Sources: [union](https://www.oxfordlearnersdictionaries.com/definition/american_english/union)

### unmount-tech

- Entry: [unmount](../entries/unmount-tech.yml)
- IPA: /ˌʌnˈmaʊnt/
- Basis: Direct
- Sources: [unmount](https://dictionary.cambridge.org/dictionary/english/unmount)

### unregister-tech

- Entry: [unregister](../entries/unregister-tech.yml)
- IPA: /ˌʌnˈrɛdʒəstər/
- Basis: Composed
- Review: Un- + register; retain stem stress.
- Sources: [un-](https://dictionary.cambridge.org/dictionary/english/un); [register](https://www.oxfordlearnersdictionaries.com/definition/american_english/register_1)

### unstaged-git

- Entry: [unstaged](../entries/unstaged-git.yml)
- IPA: /ˌʌnˈsteɪdʒd/
- Basis: Composed
- Review: Un- + stage + -ed /d/ after /dʒ/.
- Sources: [un-](https://dictionary.cambridge.org/dictionary/english/un); [stage](https://www.oxfordlearnersdictionaries.com/definition/american_english/stage_2); [-ed](https://dictionary.cambridge.org/dictionary/english/ed)

### untracked-git

- Entry: [untracked](../entries/untracked-git.yml)
- IPA: /ˌʌnˈtrækt/
- Basis: Composed
- Review: Un- + track + -ed /t/ after /k/.
- Sources: [un-](https://dictionary.cambridge.org/dictionary/english/un); [track](https://www.oxfordlearnersdictionaries.com/definition/american_english/track_2); [-ed](https://dictionary.cambridge.org/dictionary/english/ed)

### upstream-git

- Entry: [upstream](../entries/upstream.yml)
- IPA: /ˌʌpˈstrim/
- Basis: Direct
- Sources: [upstream](https://www.oxfordlearnersdictionaries.com/definition/american_english/upstream)

### validation-tech

- Entry: [validation](../entries/validation-tech.yml)
- IPA: /ˌvælɪˈdeɪʃn/
- Basis: Direct
- Sources: [validation](https://www.oxfordlearnersdictionaries.com/definition/english/validation)

### verification-tech

- Entry: [verification](../entries/verification-tech.yml)
- IPA: /ˌvɛrɪfɪˈkeɪʃn/
- Basis: Direct
- Sources: [verification](https://www.oxfordlearnersdictionaries.com/definition/english/verification)

### version-pin-tech

- Entry: [version pin](../entries/version-pin-tech.yml)
- IPA: /ˈvərʒn pɪn/
- Basis: Composed
- Review: Component citation forms; retain word-level stresses, without claiming a unique sentence-level accent.
- Sources: [version](https://www.oxfordlearnersdictionaries.com/definition/american_english/version); [pin](https://www.oxfordlearnersdictionaries.com/definition/american_english/pin_1)

### vocabulary-general

- Entry: [vocabulary](../entries/vocabulary.yml)
- IPA: /voʊˈkæbjəˌlɛri/
- Basis: Direct
- Sources: [vocabulary](https://www.oxfordlearnersdictionaries.com/definition/american_english/vocabulary)

### when-applicable-workflow

- Entry: [when applicable](../entries/when-applicable-workflow.yml)
- IPA: /wɛn ˈæplɪkəbl/
- Basis: Composed
- Review: Component citation forms; retain word-level stresses, without claiming a unique sentence-level accent.
- Sources: [when](https://www.oxfordlearnersdictionaries.com/definition/american_english/when_1); [applicable](https://www.oxfordlearnersdictionaries.com/definition/american_english/applicable)

### where-practical-workflow

- Entry: [where practical](../entries/where-practical-workflow.yml)
- IPA: /wɛr ˈpræktɪkl/
- Basis: Composed
- Review: Component citation forms; retain word-level stresses, without claiming a unique sentence-level accent.
- Sources: [where](https://www.oxfordlearnersdictionaries.com/definition/american_english/where_1); [practical](https://www.oxfordlearnersdictionaries.com/definition/american_english/practical)

### wording-general

- Entry: [wording](../entries/wording-general.yml)
- IPA: /ˈwərdɪŋ/
- Basis: Direct
- Sources: [wording](https://www.oxfordlearnersdictionaries.com/definition/american_english/wording)

### workload-tech

- Entry: [workload](../entries/workload-tech.yml)
- IPA: /ˈwərkloʊd/
- Basis: Direct
- Sources: [workload](https://www.oxfordlearnersdictionaries.com/definition/american_english/workload)
