package vocabulary

import (
	"bytes"
	"encoding/csv"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"testing"
)

const validEntry = `id: idiomatic-tech
word: idiomatic
meaning: 地道的；惯用的
context: 符合某种语言惯用方式的。
example: Prefer idiomatic Go code.
`

func writeFile(t *testing.T, root, name, content string) {
	t.Helper()
	path := filepath.Join(root, name)
	if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(path, []byte(content), 0644); err != nil {
		t.Fatal(err)
	}
}

func readExport(t *testing.T, root string) []byte {
	t.Helper()
	data, err := os.ReadFile(filepath.Join(root, "dist", "vocabulary.tsv"))
	if err != nil {
		t.Fatal(err)
	}
	return data
}

func TestValidation(t *testing.T) {
	cases := []struct {
		name, data, want string
	}{
		{"malformed", "word: [", "malformed YAML"},
		{"empty file", "", "empty YAML"},
		{"root sequence", "- word", "expected a YAML mapping"},
		{"null document", "null", "expected a YAML mapping"},
		{"multiple documents", validEntry + "---\n" + validEntry, "exactly one YAML document"},
		{"empty second document", validEntry + "---\n", "exactly one YAML document"},
		{"missing field", strings.ReplaceAll(validEntry, "context: 符合某种语言惯用方式的。\n", ""), `missing required field "context"`},
		{"empty required", strings.ReplaceAll(validEntry, "word: idiomatic", `word: "  "`), `field "word": must not be empty`},
		{"null required", strings.ReplaceAll(validEntry, "word: idiomatic", "word: null"), `field "word": expected a string`},
		{"numeric text", strings.ReplaceAll(validEntry, "word: idiomatic", "word: 42"), `field "word": expected a string`},
		{"boolean text", strings.ReplaceAll(validEntry, "word: idiomatic", "word: true"), `field "word": expected a string`},
		{"mapping text", strings.ReplaceAll(validEntry, "word: idiomatic", "word: {a: b}"), `field "word": expected a string`},
		{"unknown field", validEntry + "meaning_en: natural\n", `unknown field "meaning_en"`},
		{"duplicate key", validEntry + "word: other\n", `duplicate field "word"`},
		{"numeric key", validEntry + "12: value\n", "field names must be strings"},
		{"empty optional", validEntry + "notes: ''\n", `field "notes": must not be empty`},
		{"null optional", validEntry + "notes: null\n", `field "notes": expected a string`},
		{"Anki field separator", validEntry + "notes: \"text\\x1ftext\"\n", `field "notes": unsupported control character U+001F`},
		{"bad id", strings.ReplaceAll(validEntry, "idiomatic-tech", "Idiomatic-Tech"), `field "id": must use lowercase kebab-case`},
		{"padded id", strings.ReplaceAll(validEntry, "id: idiomatic-tech", `id: " idiomatic-tech "`), `field "id"`},
		{"empty tags", validEntry + "tags: []\n", `field "tags": expected a nonempty list`},
		{"scalar tags", validEntry + "tags: tech\n", `field "tags": expected a nonempty list`},
		{"null tags", validEntry + "tags: null\n", `field "tags": expected a nonempty list`},
		{"numeric tag", validEntry + "tags: [42]\n", "each tag must be a string"},
		{"bad tag", validEntry + "tags: [Tech]\n", `invalid tag "Tech"`},
		{"empty tag", validEntry + "tags: ['']\n", `invalid tag ""`},
		{"duplicate tag", validEntry + "tags: [tech, tech]\n", `duplicate tag "tech"`},
		{"missing audio", validEntry + "audio: idiomatic-tech.mp3\n", "cannot read media/idiomatic-tech.mp3"},
		{"wrong audio id", validEntry + "audio: another.mp3\n", `must be "idiomatic-tech.mp3"`},
		{"audio path", validEntry + "audio: ../idiomatic-tech.mp3\n", "without a directory"},
		{"alias", strings.ReplaceAll(validEntry, "word: idiomatic", "word: &word idiomatic") + "notes: *word\n", "aliases are not supported"},
		{"merge key", validEntry + "<<: {notes: merged}\n", "field names must be strings"},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			root := t.TempDir()
			writeFile(t, root, "entries/idiomatic.yml", tc.data)
			_, _, err := Load(root)
			if err == nil || !strings.Contains(err.Error(), "entries/idiomatic.yml:") || !strings.Contains(err.Error(), tc.want) {
				t.Fatalf("want file and %q in error, got %v", tc.want, err)
			}
		})
	}
}

func TestMissingOrEmptyEntries(t *testing.T) {
	root := t.TempDir()
	if _, _, err := Load(root); err == nil || !strings.Contains(err.Error(), "entries:") {
		t.Fatalf("missing directory: %v", err)
	}
	writeFile(t, root, "entries/.gitkeep", "")
	if _, _, err := Load(root); err == nil || !strings.Contains(err.Error(), "no vocabulary found") {
		t.Fatalf("empty directory: %v", err)
	}
}

func TestInvalidFilenames(t *testing.T) {
	for _, name := range []string{"Bad.yml", "bad_name.yml", "word.yaml", "word.yml.bak", "-word.yml", "word--name.yml"} {
		t.Run(name, func(t *testing.T) {
			root := t.TempDir()
			writeFile(t, root, "entries/"+name, validEntry)
			_, _, err := Load(root)
			if err == nil || !strings.Contains(err.Error(), "invalid filename") {
				t.Fatalf("got %v", err)
			}
		})
	}
}

func TestDirectoryIsNotEntry(t *testing.T) {
	root := t.TempDir()
	if err := os.MkdirAll(filepath.Join(root, "entries", "word.yml"), 0755); err != nil {
		t.Fatal(err)
	}
	if _, _, err := Load(root); err == nil || !strings.Contains(err.Error(), "expected a regular file") {
		t.Fatalf("got %v", err)
	}
}

func TestDuplicateIDsAndDeterministicErrors(t *testing.T) {
	root := t.TempDir()
	writeFile(t, root, "entries/b.yml", validEntry)
	writeFile(t, root, "entries/a.yml", validEntry)
	writeFile(t, root, "entries/c.yml", validEntry+"typo: value\n")
	_, _, err := Load(root)
	want := "entries/b.yml: duplicate id \"idiomatic-tech\" (also in entries/a.yml); use a unique stable ID\nentries/c.yml: line 6: unknown field \"typo\"; check the schema in README.md"
	if err == nil || err.Error() != want {
		t.Fatalf("want %q, got %v", want, err)
	}
	_, _, repeated := Load(root)
	if repeated == nil || repeated.Error() != err.Error() {
		t.Fatalf("nondeterministic errors: %v", repeated)
	}
}

func TestDistinctMeaningsAndIDSorting(t *testing.T) {
	root := t.TempDir()
	writeFile(t, root, "entries/a.yml", "id: commit-git\nword: commit\nmeaning: 提交\ncontext: 保存一组 Git 更改。\nexample: Commit the fix.\n")
	writeFile(t, root, "entries/z.yml", "id: commit-general\nword: Commit\nmeaning: 承诺\ncontext: 承诺投入时间或资源。\nexample: Commit to a schedule.\n")
	entries, warnings, err := Load(root)
	if err != nil {
		t.Fatal(err)
	}
	if len(entries) != 2 || entries[0].ID != "commit-general" || entries[1].ID != "commit-git" {
		t.Fatalf("entries not sorted by ID: %+v", entries)
	}
	if len(warnings) != 1 || !strings.Contains(warnings[0], "review meaning and context") {
		t.Fatalf("want a review warning, got %v", warnings)
	}
}

func TestAudio(t *testing.T) {
	root := t.TempDir()
	writeFile(t, root, "entries/idiomatic.yml", validEntry+"audio: idiomatic-tech.mp3\n")
	writeFile(t, root, "media/idiomatic-tech.mp3", "")
	if _, _, err := Load(root); err == nil || !strings.Contains(err.Error(), "nonempty regular file") {
		t.Fatalf("empty audio accepted: %v", err)
	}
	// This tests references and file checks, not MP3 decoding.
	writeFile(t, root, "media/idiomatic-tech.mp3", "audio fixture")
	entries, _, err := Load(root)
	if err != nil || entries[0].Audio != "idiomatic-tech.mp3" {
		t.Fatalf("valid reference: %v", err)
	}
}

func TestExportEscapingAndMetadata(t *testing.T) {
	entries := []Entry{{
		ID: "example-tech", Word: "example", Pronunciation: "/test/", PartOfSpeech: "noun",
		Meaning: "示例", Context: "Use <div> & \"quotes\".\r\nNext\rLast",
		Example: "one\ttwo", Notes: "Literal [sound:other.mp3]", Audio: "example-tech.mp3",
		Tags: []string{"tech", "programming"},
	}}
	data, err := render(entries)
	if err != nil {
		t.Fatal(err)
	}
	headers := "#separator:Tab\n#html:true\n#notetype:LexForge\n#tags column:10\n#columns:ID\tEnglish\tPronunciation\tPartOfSpeech\tMeaning\tContext\tExample\tNotes\tAudio\tTags\n"
	if !bytes.HasPrefix(data, []byte(headers)) {
		t.Fatalf("wrong headers: %s", data)
	}
	reader := csv.NewReader(bytes.NewReader(data[len(headers):]))
	reader.Comma = '\t'
	reader.FieldsPerRecord = 10
	rows, err := reader.ReadAll()
	if err != nil {
		t.Fatal(err)
	}
	want := [][]string{{"example-tech", "example", "/test/", "noun", "示例", "Use &lt;div&gt; &amp; &#34;quotes&#34;.<br>Next<br>Last", "one\ttwo", "Literal &#91;sound:other.mp3]", "[sound:example-tech.mp3]", "programming tech"}}
	if !reflect.DeepEqual(rows, want) {
		t.Fatalf("want %#v, got %#v", want, rows)
	}
	if entries[0].Tags[0] != "tech" {
		t.Fatal("export changed source tag order")
	}
}

func TestBuildDeterminismAndFailurePreservesExport(t *testing.T) {
	root := t.TempDir()
	writeFile(t, root, "entries/z.yml", strings.ReplaceAll(validEntry, "id: idiomatic-tech", "id: a-tech"))
	writeFile(t, root, "entries/a.yml", strings.ReplaceAll(validEntry, "id: idiomatic-tech", "id: z-tech"))
	count, _, err := Build(root)
	if err != nil || count != 2 {
		t.Fatalf("build: count=%d, err=%v", count, err)
	}
	first := readExport(t, root)
	if bytes.Index(first, []byte("\na-tech\t")) > bytes.Index(first, []byte("\nz-tech\t")) {
		t.Fatal("export is not sorted by ID")
	}
	if !bytes.Contains(first, []byte("\tidiomatic\t\t\t")) {
		t.Fatal("missing optional fields shifted export columns")
	}
	if _, _, err := Build(root); err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(first, readExport(t, root)) {
		t.Fatal("repeated build differs")
	}
	writeFile(t, root, "entries/broken.yml", "word: [")
	if _, _, err := Build(root); err == nil {
		t.Fatal("build accepted invalid data")
	}
	if !bytes.Equal(first, readExport(t, root)) {
		t.Fatal("invalid input changed the existing export")
	}
	files, err := os.ReadDir(filepath.Join(root, "dist"))
	if err != nil || len(files) != 1 || files[0].Name() != "vocabulary.tsv" {
		t.Fatalf("temporary export files left behind: %v, %v", files, err)
	}
}

func TestInvalidBuildCreatesNoOutput(t *testing.T) {
	root := t.TempDir()
	writeFile(t, root, "entries/broken.yml", "word: [")
	if _, _, err := Build(root); err == nil {
		t.Fatal("build accepted invalid data")
	}
	if _, err := os.Stat(filepath.Join(root, "dist")); !os.IsNotExist(err) {
		t.Fatalf("invalid build created output: %v", err)
	}
}
