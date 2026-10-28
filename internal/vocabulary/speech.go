package vocabulary

import (
	"fmt"
	"strings"
	"unicode"

	"golang.org/x/text/language"
	"golang.org/x/text/unicode/norm"
)

type speechVoice struct {
	name string
	tag  language.Tag
}
type speechSettings struct {
	voice, text string
	rate        int
}

func parseVoices(data string) []speechVoice {
	var voices []speechVoice
	for _, line := range strings.Split(data, "\n") {
		fields := strings.Fields(strings.SplitN(line, "#", 2)[0])
		if len(fields) < 2 {
			continue
		}
		tag, err := language.Parse(fields[len(fields)-1])
		if err != nil {
			continue
		}
		voices = append(voices, speechVoice{strings.Join(fields[:len(fields)-1], " "), tag})
	}
	return voices
}

func selectSpeech(entry Entry, options AudioOptions, voices []speechVoice) (speechSettings, error) {
	result := speechSettings{text: entry.Word, rate: options.Rate}
	tag, err := language.Parse(entry.Language)
	if err != nil {
		return result, err
	}
	base, script, region := tag.Raw()
	if options.TextSource == "pronunciation" {
		if err := validateKanaReading(entry); err != nil {
			return result, err
		}
		result.text = entry.Pronunciation
	}
	if result.rate == 0 && base.String() == "en" {
		result.rate = 150
	}
	defaults := map[string]string{"en": "Samantha", "ja": "Kyoko", "ko": "Yuna"}
	var matches []speechVoice
	for _, voice := range voices {
		vb, _, vr := voice.tag.Raw()
		vs, _ := voice.tag.Script()
		sameLanguage := base.String() != "und" && base == vb && (script.String() == "Zzzz" || script == vs)
		if options.Voice != "" {
			if voice.name == options.Voice {
				if !sameLanguage {
					return result, fmt.Errorf("voice %q does not match language %s", options.Voice, entry.Language)
				}
				result.voice = voice.name
				return result, nil
			}
			continue
		}
		if !sameLanguage || (region.String() != "ZZ" && region != vr) {
			continue
		}
		if voice.name == defaults[entry.Language] {
			result.voice = voice.name
			return result, nil
		}
		matches = append(matches, voice)
	}
	if options.Voice != "" {
		return result, fmt.Errorf("voice %q is not installed; inspect say -v '?'", options.Voice)
	}
	if len(matches) == 1 {
		result.voice = matches[0].name
		return result, nil
	}
	names := make([]string, len(matches))
	for i, v := range matches {
		names[i] = v.name + " (" + v.tag.String() + ")"
	}
	return result, fmt.Errorf("language %s has %d matching voices [%s]; install a matching voice or select one with --id and --voice", entry.Language, len(matches), strings.Join(names, ", "))
}

func validateKanaReading(entry Entry) error {
	if strings.Split(entry.Language, "-")[0] != "ja" || entry.Pronunciation == "" {
		return fmt.Errorf("--text-source pronunciation requires a Japanese entry with a kana reading")
	}
	hasKana := false
	// Accept decomposed and half-width kana without rewriting source or speech text.
	for _, r := range norm.NFKC.String(entry.Pronunciation) {
		if unicode.In(r, unicode.Hiragana, unicode.Katakana) {
			hasKana = true
			continue
		}
		if unicode.IsSpace(r) || strings.ContainsRune("ー・", r) {
			continue
		}
		return fmt.Errorf("pronunciation must contain kana only for speech; found %q", r)
	}
	if !hasKana {
		return fmt.Errorf("pronunciation must contain a kana reading")
	}
	return nil
}
