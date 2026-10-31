package vocabulary

import (
	"bytes"
	"encoding/csv"
	"strings"
	"testing"
)

func TestOptionalMemoryText(t *testing.T) {
	base := "id: concept\nlanguage: en\nword: concept\nmeaning: 原理\n"
	for _, field := range []string{"context", "example", "etymology", "word_formation", "mnemonic"} {
		for _, value := range []string{"", "null", `""`, `"  \t\n"`, `"保留 <text> [sound:test.mp3]"`} {
			if _, err := parse([]byte(base + field + ": " + value + "\n")); err != nil {
				t.Fatalf("%s=%s: %v", field, value, err)
			}
		}
		for _, value := range []string{"42", "true", "[]", "{}", `"\v"`, "*target"} {
			if _, err := parse([]byte(base + "notes: &target text\n" + field + ": " + value + "\n")); err == nil {
				t.Fatalf("accepted %s=%s", field, value)
			}
		}
	}
	for _, field := range []string{"pronunciation", "part_of_speech", "notes", "audio"} {
		if _, err := parse([]byte(base + field + ": null\n")); err == nil {
			t.Fatalf("relaxed %s", field)
		}
	}
}

func TestRelatedValidation(t *testing.T) {
	for _, value := range []string{"null", "[]", "{}", "other", "[42]", "[Bad-ID]", "[self]", "[missing]", "[other, other]"} {
		t.Run(value, func(t *testing.T) {
			root := t.TempDir()
			writeFile(t, root, "entries/self.yml", "id: self\nlanguage: en\nword: Self\nmeaning: 自身\nrelated_entries: "+value+"\n")
			writeFile(t, root, "entries/other.yml", "id: other\nlanguage: en\nword: Other\nmeaning: 其他\n")
			writeFile(t, root, "dist/vocabulary.tsv", "previous export")
			if _, _, err := Build(root); err == nil {
				t.Fatal("accepted invalid relationship")
			}
			if string(readExport(t, root)) != "previous export" {
				t.Fatal("replaced export on invalid reference")
			}
		})
	}
	// Reciprocal declarations and cycles are valid, independently of prerequisites.
	entries := []Entry{{ID: "a", RelatedEntries: []string{"b", "c"}}, {ID: "b", RelatedEntries: []string{"a", "c"}}, {ID: "c", RelatedEntries: []string{"a"}}}
	if err := validateRelatedEntries(entries); err != nil {
		t.Fatal(err)
	}
}

func TestRelatedDisplay(t *testing.T) {
	entries := []Entry{
		{ID: "a", Language: "en", Word: "A", Meaning: "甲", RelatedEntries: []string{"c", "b"}, Etymology: "<origin>", WordFormation: "root\nform", Mnemonic: "[sound:x.mp3]", Notes: "详细\n说明"},
		{ID: "b", Language: "en", Word: "<B>", Meaning: "乙", RelatedEntries: []string{"a", "d"}},
		{ID: "c", Language: "ja", Word: "語", Meaning: "丙"},
		{ID: "d", Language: "en", Word: "D", Meaning: "丁"},
	}
	got := relatedWords(entries)
	if got["a"] != "<B> — 乙\n語 (ja) — 丙" || got["d"] != "<B> — 乙" {
		t.Fatalf("unexpected neighbors: %#v", got)
	}
	data, err := render(entries)
	if err != nil {
		t.Fatal(err)
	}
	reader := csv.NewReader(bytes.NewReader(data))
	reader.Comma = '\t'
	reader.Comment = '#'
	reader.FieldsPerRecord = 17
	rows, err := reader.ReadAll()
	if err != nil {
		t.Fatal(err)
	}
	if rows[0][12] != "&lt;origin&gt;" || rows[0][13] != "root<br>form" || rows[0][14] != "&#91;sound:x.mp3]" || rows[0][15] != "&lt;B&gt; — 乙<br>語 (ja) — 丙" || rows[0][8] != "详细<br>说明" {
		t.Fatalf("incorrect fields: %#v", rows[0])
	}
	entries[0].RelatedEntries = []string{"b", "c"}
	again, _ := render(entries)
	if !bytes.Equal(data, again) {
		t.Fatal("declaration order changed export")
	}
	entries[1].Word = "Updated"
	if !strings.Contains(relatedWords(entries)["a"], "Updated") {
		t.Fatal("stale word")
	}
	entries[0].RelatedEntries = []string{"c"}
	entries[1].RelatedEntries = []string{"d"}
	if strings.Contains(relatedWords(entries)["a"], "Updated") {
		t.Fatal("removed link remains")
	}
}

func TestAudioInsertionPreservesMemoryFields(t *testing.T) {
	before := []byte("id: concept\nlanguage: en\nword: concept\nmeaning: 原理\ncontext: null\nexample: ''\netymology: origin\nword_formation: formation\nmnemonic: memory\nrelated_entries: [other]\nlearning_priority: 1\nprerequisites: [base]\nnotes: |\n  Longer notes.\n  Another line.\n")
	after, err := insertAudio(before, "concept.mp3")
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(bytes.Replace(after, []byte("audio: concept.mp3\n"), nil, 1), before) {
		t.Fatal("insertion changed source bytes")
	}
	if _, err := checkAudioInsertion(before, after, "concept.mp3"); err != nil {
		t.Fatal(err)
	}
	for _, pair := range [][2]string{{"origin", "changed"}, {"formation", "changed"}, {"memory", "changed"}, {"[other]", "[changed]"}, {"priority: 1", "priority: 2"}, {"[base]", "[changed]"}} {
		modified := bytes.Replace(after, []byte(pair[0]), []byte(pair[1]), 1)
		if _, err := checkAudioInsertion(before, modified, "concept.mp3"); err == nil {
			t.Fatalf("missed change to %s", pair[0])
		}
	}
}
