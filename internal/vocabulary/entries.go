// Package vocabulary validates source entries and builds Anki import data.
package vocabulary

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
	"unicode"

	"go.yaml.in/yaml/v3"
	"golang.org/x/text/language"
)

type Entry struct {
	ID            string
	Language      string
	Word          string
	Pronunciation string
	PartOfSpeech  string
	Meaning       string
	Context       string
	Example       string
	Notes         string
	Audio         string
	Tags          []string
	sourcePath    string
	sourceData    []byte
}

var kebabCase = regexp.MustCompile(`^[a-z0-9]+(-[a-z0-9]+)*$`)

// Load reads entries once, reports errors in filename order, and sorts by ID.
// Repeated words require human review; they may represent distinct meanings.
func Load(root string) ([]Entry, []string, error) {
	files, err := os.ReadDir(filepath.Join(root, "entries"))
	if err != nil {
		return nil, nil, fmt.Errorf("entries: %w", err)
	}
	var entries []Entry
	var warnings []string
	var problems []error
	ids := make(map[string]string)
	words := make(map[string]string)
	for _, file := range files {
		name := file.Name()
		if strings.HasPrefix(name, ".") {
			continue
		}
		path := filepath.Join("entries", name)
		if !strings.HasSuffix(name, ".yml") || !kebabCase.MatchString(strings.TrimSuffix(name, ".yml")) {
			problems = append(problems, fmt.Errorf("%s: invalid filename: use lowercase kebab-case with a .yml extension", path))
			continue
		}
		if !file.Type().IsRegular() {
			problems = append(problems, fmt.Errorf("%s: expected a regular file; directories and symlinks are not supported", path))
			continue
		}
		data, err := os.ReadFile(filepath.Join(root, path))
		if err != nil {
			problems = append(problems, fmt.Errorf("%s: %w", path, err))
			continue
		}
		entry, err := parse(data)
		if err == nil {
			err = validate(entry, root)
		}
		if err != nil {
			problems = append(problems, fmt.Errorf("%s: %w", path, err))
			continue
		}
		entry.sourcePath = path
		entry.sourceData = data
		if previous, ok := ids[entry.ID]; ok {
			problems = append(problems, fmt.Errorf("%s: duplicate id %q (also in %s); use a unique stable ID", path, entry.ID, previous))
		} else {
			ids[entry.ID] = path
		}
		word := entry.Language + "\x00" + strings.ToLower(strings.Join(strings.Fields(entry.Word), " "))
		if previous, ok := words[word]; ok {
			warnings = append(warnings, fmt.Sprintf("%s: word %q also appears in %s; review meaning and context before keeping both entries", path, entry.Word, previous))
		} else {
			words[word] = path
		}
		entries = append(entries, entry)
	}
	if len(problems) > 0 {
		return nil, warnings, errors.Join(problems...)
	}
	if len(entries) == 0 {
		return nil, warnings, errors.New("entries: no vocabulary found; add a lowercase kebab-case .yml file")
	}
	sort.Slice(entries, func(i, j int) bool { return entries[i].ID < entries[j].ID })
	return entries, warnings, nil
}

func parse(data []byte) (Entry, error) {
	var entry Entry
	decoder := yaml.NewDecoder(bytes.NewReader(data))
	var document yaml.Node
	if err := decoder.Decode(&document); err != nil {
		if errors.Is(err, io.EOF) {
			return entry, errors.New("empty YAML: expected one vocabulary mapping")
		}
		return entry, fmt.Errorf("malformed YAML: %w", err)
	}
	var extra yaml.Node
	if err := decoder.Decode(&extra); !errors.Is(err, io.EOF) {
		if err != nil {
			return entry, fmt.Errorf("malformed YAML: %w", err)
		}
		return entry, errors.New("expected exactly one YAML document per file")
	}
	if len(document.Content) != 1 || document.Content[0].Kind != yaml.MappingNode || document.Content[0].Tag != "!!map" {
		return entry, errors.New("expected a YAML mapping of vocabulary fields")
	}
	fields := map[string]*string{
		"id": &entry.ID, "language": &entry.Language, "word": &entry.Word, "pronunciation": &entry.Pronunciation,
		"part_of_speech": &entry.PartOfSpeech, "meaning": &entry.Meaning,
		"context": &entry.Context, "example": &entry.Example, "notes": &entry.Notes,
		"audio": &entry.Audio,
	}
	seen := make(map[string]bool)
	nodes := document.Content[0].Content
	for i := 0; i < len(nodes); i += 2 {
		key, value := nodes[i], nodes[i+1]
		if key.Kind != yaml.ScalarNode || key.Tag != "!!str" {
			return entry, fmt.Errorf("line %d: field names must be strings", key.Line)
		}
		name := key.Value
		if seen[name] {
			return entry, fmt.Errorf("line %d: duplicate field %q; keep one value", key.Line, name)
		}
		seen[name] = true
		if name == "tags" {
			if value.Kind != yaml.SequenceNode || value.Tag != "!!seq" || len(value.Content) == 0 {
				return entry, errors.New(`field "tags": expected a nonempty list of strings; omit unused optional fields`)
			}
			for _, tag := range value.Content {
				if tag.Kind != yaml.ScalarNode || tag.Tag != "!!str" {
					return entry, errors.New(`field "tags": each tag must be a string`)
				}
				entry.Tags = append(entry.Tags, tag.Value)
			}
			continue
		}
		destination, ok := fields[name]
		if !ok {
			return entry, fmt.Errorf("line %d: unknown field %q; check the schema in README.md", key.Line, name)
		}
		if value.Kind != yaml.ScalarNode || value.Tag != "!!str" {
			return entry, fmt.Errorf("field %q: expected a string (quote numeric or boolean text); aliases are not supported", name)
		}
		if strings.TrimSpace(value.Value) == "" {
			return entry, fmt.Errorf("field %q: must not be empty; omit unused optional fields", name)
		}
		for _, character := range value.Value {
			if unicode.IsControl(character) && character != '\t' && character != '\n' && character != '\r' {
				return entry, fmt.Errorf("field %q: unsupported control character %U; remove it", name, character)
			}
		}
		*destination = value.Value
	}
	for _, name := range []string{"id", "language", "word", "meaning", "context", "example"} {
		if !seen[name] {
			return entry, fmt.Errorf("missing required field %q", name)
		}
	}
	if strings.Contains(entry.Language, "_") {
		return entry, errors.New(`field "language": use BCP 47 hyphens, not underscores`)
	}
	tag, err := language.Parse(entry.Language)
	if err != nil {
		return entry, fmt.Errorf("field %q: %w", "language", err)
	}
	entry.Language = tag.String()
	return entry, nil
}

func validate(entry Entry, root string) error {
	if !kebabCase.MatchString(entry.ID) {
		return errors.New(`field "id": must use lowercase kebab-case (letters, digits, and single hyphens)`)
	}
	seen := make(map[string]bool)
	for _, tag := range entry.Tags {
		if !kebabCase.MatchString(tag) {
			return fmt.Errorf("field %q: invalid tag %q; use lowercase kebab-case", "tags", tag)
		}
		if seen[tag] {
			return fmt.Errorf("field %q: duplicate tag %q; keep one occurrence", "tags", tag)
		}
		seen[tag] = true
	}
	if entry.Audio != "" {
		if entry.Audio != entry.ID+".mp3" {
			return fmt.Errorf("field %q: must be %q, without a directory", "audio", entry.ID+".mp3")
		}
		path := filepath.Join(root, "media", entry.Audio)
		info, err := os.Lstat(path)
		if err != nil {
			return fmt.Errorf("field %q: cannot read media/%s: %w", "audio", entry.Audio, err)
		}
		if !info.Mode().IsRegular() || info.Size() == 0 {
			return fmt.Errorf("field %q: media/%s must be a nonempty regular file", "audio", entry.Audio)
		}
	}
	return nil
}
