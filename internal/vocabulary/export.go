package vocabulary

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"html"
	"os"
	"path/filepath"
	"slices"
	"sort"
	"strings"
)

// Build validates all input before replacing the export. Invalid input leaves
// the previous export untouched. The temporary file stays on the same filesystem.
func Build(root string) (int, []string, error) {
	entries, warnings, err := Load(root)
	if err != nil {
		return 0, warnings, err
	}
	data, err := render(entries)
	if err != nil {
		return 0, warnings, fmt.Errorf("encode TSV: %w", err)
	}
	dir := filepath.Join(root, "dist")
	if err := os.MkdirAll(dir, 0755); err != nil {
		return 0, warnings, fmt.Errorf("dist: %w", err)
	}
	file, err := os.CreateTemp(dir, ".vocabulary-*.tsv")
	if err != nil {
		return 0, warnings, fmt.Errorf("dist: %w", err)
	}
	defer os.Remove(file.Name())
	defer file.Close()
	if err := file.Chmod(0644); err != nil {
		return 0, warnings, fmt.Errorf("set export permissions: %w", err)
	}
	if _, err := file.Write(data); err != nil {
		return 0, warnings, fmt.Errorf("write export: %w", err)
	}
	if err := file.Close(); err != nil {
		return 0, warnings, fmt.Errorf("close export: %w", err)
	}
	if err := os.Rename(file.Name(), filepath.Join(dir, "vocabulary.tsv")); err != nil {
		return 0, warnings, fmt.Errorf("replace dist/vocabulary.tsv: %w", err)
	}
	return len(entries), warnings, nil
}

func render(entries []Entry) ([]byte, error) {
	var output bytes.Buffer
	output.WriteString("#separator:Tab\n#html:true\n#notetype:LexForge\n#tags column:11\n")
	output.WriteString("#columns:ID\tLanguage\tWord\tPronunciation\tPartOfSpeech\tMeaning\tContext\tExample\tNotes\tAudio\tTags\n")
	writer := csv.NewWriter(&output)
	writer.Comma = '\t'
	for _, entry := range entries {
		fields := []string{entry.ID, entry.Language, entry.Word, entry.Pronunciation, entry.PartOfSpeech, entry.Meaning, entry.Context, entry.Example, entry.Notes}
		for i := range fields {
			fields[i] = displayText(fields[i])
		}
		audio := ""
		if entry.Audio != "" {
			audio = "[sound:" + entry.Audio + "]"
		}
		tags := append([]string(nil), entry.Tags...)
		tags = append(tags, "lang-"+strings.ToLower(entry.Language))
		sort.Strings(tags)
		tags = slices.Compact(tags)
		fields = append(fields, audio, strings.Join(tags, " "))
		if err := writer.Write(fields); err != nil {
			return nil, err
		}
	}
	writer.Flush()
	return output.Bytes(), writer.Error()
}

func displayText(value string) string {
	value = strings.ReplaceAll(value, "\r\n", "\n")
	value = strings.ReplaceAll(value, "\r", "\n")
	value = html.EscapeString(value)
	// Keep literal Anki media markers in vocabulary text from becoming audio.
	value = strings.ReplaceAll(value, "[", "&#91;")
	return strings.ReplaceAll(value, "\n", "<br>")
}
