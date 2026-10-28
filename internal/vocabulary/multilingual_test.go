package vocabulary

import (
	"bytes"
	"io"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestLanguageParsing(t *testing.T) {
	for input, want := range map[string]string{"en-QQ": "en-QQ", "EN-us": "en-US", "iw": "he", "ja": "ja", "ko": "ko", "fr": "fr", "zh-Hant": "zh-Hant", "x-personal": "x-personal", "en-u-ca-gregory": "en-u-ca-gregory"} {
		t.Run(input, func(t *testing.T) {
			data := strings.Replace(validEntry, "language: en", "language: "+input, 1)
			entry, err := parse([]byte(data))
			if err != nil || entry.Language != want {
				t.Fatalf("%q: %+v %v", input, entry, err)
			}
			after, err := insertAudio([]byte(data), "idiomatic-tech.mp3")
			if err != nil || !bytes.Contains(after, []byte("language: "+input+"\n")) {
				t.Fatalf("source spelling changed: %s %v", after, err)
			}
		})
	}
	for _, input := range []string{"en_US", "en--US", "zzzz", "en-unknownvariant"} {
		if _, err := parse([]byte(strings.Replace(validEntry, "language: en", "language: "+input, 1))); err == nil {
			t.Errorf("accepted %s", input)
		}
	}
	if _, err := parse([]byte(strings.Replace(validEntry, "language: en\n", "", 1))); err == nil {
		t.Fatal("accepted missing language")
	}
}

func TestMultilingualCollection(t *testing.T) {
	root := t.TempDir()
	for i, lang := range []string{"ja", "ko", "JA"} {
		data := strings.Replace(validEntry, "language: en", "language: "+lang, 1)
		data = strings.ReplaceAll(data, "idiomatic-tech", []string{"ja-one", "ko-one", "ja-two"}[i])
		data = strings.Replace(data, "word: idiomatic", "word: 確認", 1)
		data += "pronunciation: かくにん\ntags: [lang-ja]\n"
		writeFile(t, root, "entries/"+[]string{"a", "b", "c"}[i]+".yml", data)
	}
	entries, warnings, err := Load(root)
	if err != nil || len(warnings) != 1 {
		t.Fatalf("%v %v", warnings, err)
	}
	data, err := render(entries)
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Contains(data, []byte("ja-one\tja\t確認\tかくにん")) || bytes.Contains(data, []byte("lang-ja lang-ja")) {
		t.Fatalf("bad export: %s", data)
	}
	// IDs remain global even when languages differ.
	writeFile(t, root, "entries/d.yml", strings.Replace(validEntry, "idiomatic-tech", "ja-one", 1))
	if _, _, err := Load(root); err == nil {
		t.Fatal("accepted cross-language duplicate ID")
	}
}

func TestVoiceSelection(t *testing.T) {
	voices := parseVoices("Samantha en_US # hello\nAlex en_US\nDaniel en_GB\nKyoko ja_JP\nYuna ko_KR\nThomas fr_FR\nTest Voice fr_CA\n")
	cases := []struct{ lang, override, want string }{
		{"en", "", "Samantha"}, {"ja", "", "Kyoko"}, {"ko", "", "Yuna"},
		{"en-GB", "", "Daniel"}, {"fr-CA", "", "Test Voice"}, {"en-CA", "Daniel", "Daniel"},
		{"en-US", "", ""}, {"fr", "", ""}, {"en-CA", "", ""}, {"ja", "Samantha", ""}, {"ko", "Missing", ""},
	}
	for _, tc := range cases {
		got, err := selectSpeech(Entry{Language: tc.lang, Word: "test"}, AudioOptions{Voice: tc.override}, voices)
		if tc.want == "" {
			if err == nil {
				t.Errorf("expected error for %+v", tc)
			}
			continue
		}
		if err != nil || got.voice != tc.want {
			t.Errorf("%+v: %+v %v", tc, got, err)
		}
	}
	for _, reading := range []string{"かくにん", "カクニン", "コーヒー"} {
		got, err := selectSpeech(Entry{Language: "ja", Word: "確認", Pronunciation: reading}, AudioOptions{TextSource: "pronunciation", Rate: 180}, voices)
		if err != nil || got.text != reading || got.rate != 180 {
			t.Fatalf("%+v %v", got, err)
		}
	}
	for _, entry := range []Entry{{Language: "ja"}, {Language: "en", Pronunciation: "/test/"}, {Language: "ja", Pronunciation: "kakunin"}} {
		if _, err := selectSpeech(entry, AudioOptions{TextSource: "pronunciation"}, voices); err == nil {
			t.Fatal("accepted invalid kana speech")
		}
	}
}

func TestAudioPreflightAndUnsupportedReuse(t *testing.T) {
	root := t.TempDir()
	writeFile(t, root, "entries/a.yml", validEntry)
	writeFile(t, root, "entries/b.yml", strings.Replace(strings.Replace(validEntry, "idiomatic-tech", "other", 1), "language: en", "language: fr", 1))
	_, _, err := generateAudio(root, AudioOptions{All: true}, io.Discard, fakeAudioTools(t, func(string, []string) error { t.Fatal("preflight generated audio"); return nil }))
	if err == nil {
		t.Fatal("accepted missing voice")
	}
	if _, err := os.Stat(filepath.Join(root, "media")); !os.IsNotExist(err) {
		t.Fatal("preflight created media")
	}
	writeFile(t, root, "media/other.mp3", "existing MP3")
	summary, _, err := generateAudio(root, AudioOptions{ID: "other"}, io.Discard, fakeAudioTools(t, nil))
	if err != nil || summary.Reused != 1 {
		t.Fatalf("%+v %v", summary, err)
	}
}

func TestMultilingualAudioArguments(t *testing.T) {
	for _, tc := range []struct {
		lang, word, reading, voice string
		rate                       int
	}{
		{"ja", "確認", "かくにん", "Kyoko", 180}, {"ko", "확인", "", "Yuna", 0},
	} {
		t.Run(tc.lang, func(t *testing.T) {
			root := t.TempDir()
			data := strings.Replace(validEntry, "language: en", "language: "+tc.lang, 1)
			data = strings.Replace(data, "word: idiomatic", "word: "+tc.word, 1)
			options := AudioOptions{ID: "idiomatic-tech", Rate: tc.rate}
			if tc.reading != "" {
				data += "pronunciation: " + tc.reading + "\n"
				options.TextSource = "pronunciation"
			}
			writeFile(t, root, "entries/word.yml", data)
			called := false
			prepare := fakeAudioTools(t, func(name string, args []string) error {
				if name != "say" {
					return nil
				}
				called = true
				if args[1] != tc.voice {
					t.Fatalf("wrong voice: %v", args)
				}
				index := 2
				if tc.rate > 0 {
					if args[2] != "-r" || args[3] != "180" {
						t.Fatalf("wrong rate: %v", args)
					}
					index = 4
				}
				if args[index] != "-f" {
					t.Fatalf("wrong input args: %v", args)
				}
				input, err := os.ReadFile(args[index+1])
				if err != nil {
					return err
				}
				want := tc.word
				if tc.reading != "" {
					want = tc.reading
				}
				if string(input) != want {
					t.Fatalf("wrong speech input %q", input)
				}
				return nil
			})
			summary, _, err := generateAudio(root, options, io.Discard, func() (audioTools, error) {
				tools, err := prepare()
				tools.voices = parseVoices("Kyoko ja_JP\nYuna ko_KR")
				return tools, err
			})
			if err != nil || !called || summary.Generated != 1 {
				t.Fatalf("%+v %v", summary, err)
			}
		})
	}
}

func TestAudioOverridesRejectInvalidInput(t *testing.T) {
	root := t.TempDir()
	writeFile(t, root, "entries/a.yml", validEntry+"audio: idiomatic-tech.mp3\n")
	writeFile(t, root, "media/idiomatic-tech.mp3", "existing")
	for _, options := range []AudioOptions{
		{ID: "idiomatic-tech", Rate: -1}, {All: true, Voice: "Kyoko"}, {All: true, Rate: 180},
		{All: true, TextSource: "pronunciation"}, {ID: "idiomatic-tech", TextSource: "other"},
		{ID: "idiomatic-tech", TextSource: "pronunciation"},
	} {
		_, _, err := generateAudio(root, options, io.Discard, func() (audioTools, error) { t.Fatal("invalid override reached tools"); return audioTools{}, nil })
		if err == nil {
			t.Fatalf("accepted %+v", options)
		}
	}
}

func TestKanaUnicodeForms(t *testing.T) {
	voices := parseVoices("Kyoko ja_JP")
	for _, reading := range []string{"か\u3099", "ｶﾞ"} {
		entry := Entry{Language: "ja", Pronunciation: reading}
		speech, err := selectSpeech(entry, AudioOptions{TextSource: "pronunciation"}, voices)
		if err != nil {
			t.Fatalf("valid kana %q rejected: %v", reading, err)
		}
		if speech.text != reading || entry.Pronunciation != reading {
			t.Fatalf("kana was rewritten: %q", speech.text)
		}
	}
	for _, reading := range []string{"\u3099", "ﾞ", "ｶﾞa"} {
		if err := validateKanaReading(Entry{Language: "ja", Pronunciation: reading}); err == nil {
			t.Errorf("accepted invalid kana %q", reading)
		}
	}
}
