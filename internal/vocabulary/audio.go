package vocabulary

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"slices"
	"strconv"
	"strings"

	"go.yaml.in/yaml/v3"
)

type AudioOptions struct {
	ID         string
	All        bool
	Overwrite  bool
	Voice      string
	Rate       int
	TextSource string
}

type AudioSummary struct {
	Generated int
	Reused    int
	Skipped   int
}

// GenerateAudio uses local macOS speech tools; validation and export stay portable.
func GenerateAudio(root string, options AudioOptions, output io.Writer) (AudioSummary, []string, error) {
	return generateAudio(root, options, output, func() (audioTools, error) {
		return findAudioTools(runtime.GOOS, exec.LookPath, runAudioTool)
	})
}

type audioTools struct {
	say, lame, decoder string
	voices             []speechVoice
	run                func(string, ...string) ([]byte, error)
}

func runAudioTool(name string, args ...string) ([]byte, error) {
	command := exec.Command(name, args...)
	var stderr bytes.Buffer
	command.Stderr = &stderr
	data, err := command.Output()
	if err != nil {
		return nil, fmt.Errorf("%s: %w: %s", filepath.Base(name), err, strings.TrimSpace(stderr.String()))
	}
	return data, nil
}

func (t audioTools) decode(path string) error {
	// Test mode can exit successfully without decoding a single MPEG frame.
	pcm, err := t.run(t.decoder, "-q", "-s", path)
	if err != nil {
		return err
	}
	if len(pcm) == 0 {
		return errors.New("decoder produced no audio samples")
	}
	return nil
}

func findAudioTools(platform string, lookup func(string) (string, error), run func(string, ...string) ([]byte, error)) (audioTools, error) {
	t := audioTools{run: run}
	if platform != "darwin" {
		return t, errors.New("audio generation requires macOS")
	}
	for _, tool := range []struct {
		name string
		path *string
	}{{"say", &t.say}, {"lame", &t.lame}, {"mpg123", &t.decoder}} {
		path, err := lookup(tool.name)
		if err != nil {
			return t, fmt.Errorf("missing %s executable; see README.md audio prerequisites: %w", tool.name, err)
		}
		*tool.path = path
	}
	voices, err := run(t.say, "-v", "?")
	if err != nil {
		return t, fmt.Errorf("list macOS voices: %w", err)
	}
	t.voices = parseVoices(string(voices))
	return t, nil
}

func generateAudio(root string, options AudioOptions, output io.Writer, prepare func() (audioTools, error)) (AudioSummary, []string, error) {
	var summary AudioSummary
	if options.All == (options.ID != "") {
		return summary, nil, errors.New("select exactly one of --id <stable-id> or --all")
	}
	if options.Rate < 0 {
		return summary, nil, errors.New("--rate must be positive or zero for the default")
	}
	if options.TextSource == "" {
		options.TextSource = "word"
	}
	if options.TextSource != "word" && options.TextSource != "pronunciation" {
		return summary, nil, errors.New("--text-source must be word or pronunciation")
	}
	if options.All && (options.Voice != "" || options.Rate != 0 || options.TextSource != "word") {
		return summary, nil, errors.New("--voice, --rate, and --text-source pronunciation require --id")
	}
	entries, warnings, err := Load(root)
	if err != nil {
		return summary, warnings, err
	}
	if !options.All {
		var selected []Entry
		for _, entry := range entries {
			if entry.ID == options.ID {
				selected = append(selected, entry)
			}
		}
		if len(selected) == 0 {
			return summary, warnings, fmt.Errorf("unknown entry ID %q", options.ID)
		}
		entries = selected
	}
	if options.TextSource == "pronunciation" {
		if err := validateKanaReading(entries[0]); err != nil {
			return summary, warnings, err
		}
	}
	t, err := prepare()
	if err != nil {
		return summary, warnings, err
	}
	settings := make(map[string]speechSettings)
	for _, entry := range entries {
		_, statErr := os.Lstat(filepath.Join(root, "media", entry.ID+".mp3"))
		if !options.Overwrite && (entry.Audio != "" || statErr == nil) {
			continue
		}
		if statErr != nil && !os.IsNotExist(statErr) {
			return summary, warnings, fmt.Errorf("%s: %w", entry.ID, statErr)
		}
		speech, err := selectSpeech(entry, options, t.voices)
		if err != nil {
			return summary, warnings, fmt.Errorf("%s: %w", entry.ID, err)
		}
		settings[entry.ID] = speech
	}
	for _, entry := range entries {
		status, err := t.process(root, entry, options.Overwrite, settings[entry.ID])
		if err != nil {
			return summary, warnings, fmt.Errorf("%s (%s): %w", entry.ID, entry.sourcePath, err)
		}
		switch status {
		case "generated":
			summary.Generated++
		case "reused":
			summary.Reused++
		case "skipped":
			summary.Skipped++
		}
		fmt.Fprintf(output, "%s: %s\n", entry.ID, status)
	}
	return summary, warnings, nil
}

func readRegular(path string) ([]byte, os.FileMode, error) {
	info, err := os.Lstat(path)
	if err != nil {
		return nil, 0, err
	}
	if !info.Mode().IsRegular() {
		return nil, 0, fmt.Errorf("%s: expected a regular file, not a directory or symlink", path)
	}
	data, err := os.ReadFile(path)
	return data, info.Mode().Perm(), err
}

func sourceUnchanged(path string, expected []byte) error {
	current, _, err := readRegular(path)
	if err != nil {
		return err
	}
	if !bytes.Equal(current, expected) {
		return fmt.Errorf("source changed during audio generation; rerun after reviewing %s", path)
	}
	return nil
}

func (t audioTools) process(root string, entry Entry, overwrite bool, speech speechSettings) (string, error) {
	source := filepath.Join(root, entry.sourcePath)
	if err := sourceUnchanged(source, entry.sourceData); err != nil {
		return "", err
	}
	if entry.Audio != "" && !overwrite {
		return "skipped", nil
	}
	filename := entry.ID + ".mp3"
	media := filepath.Join(root, "media")
	target := filepath.Join(media, filename)
	previous, mode, err := readRegular(target)
	exists := err == nil
	if err != nil && !os.IsNotExist(err) {
		return "", err
	}
	if exists && len(previous) == 0 && !overwrite {
		return "", fmt.Errorf("%s is empty; remove it or explicitly use --overwrite", target)
	}
	updated := entry.sourceData
	if entry.Audio == "" {
		updated, err = insertAudio(entry.sourceData, filename)
		if err != nil {
			return "", err
		}
	}
	if exists && !overwrite {
		if err := t.decode(target); err != nil {
			return "", fmt.Errorf("existing audio failed decoding: %w", err)
		}
		current, _, err := readRegular(target)
		if err != nil || !bytes.Equal(current, previous) {
			return "", errors.New("audio changed during decoding; source was not modified")
		}
		return "reused", replaceSource(source, entry.sourceData, updated)
	}
	if speech.voice == "" {
		return "", errors.New("audio disappeared after preflight; rerun to select a voice")
	}
	if err := os.MkdirAll(media, 0755); err != nil {
		return "", err
	}
	tmp, err := os.MkdirTemp(media, ".audio-")
	if err != nil {
		return "", err
	}
	defer os.RemoveAll(tmp)
	textPath := filepath.Join(tmp, "word.txt")
	aiff := filepath.Join(tmp, "speech.aiff")
	mp3 := filepath.Join(tmp, "speech.mp3")
	if err := os.WriteFile(textPath, []byte(speech.text), 0600); err != nil {
		return "", err
	}
	args := []string{"-v", speech.voice}
	if speech.rate > 0 {
		args = append(args, "-r", strconv.Itoa(speech.rate))
	}
	args = append(args, "-f", textPath, "-o", aiff)
	if _, err := t.run(t.say, args...); err != nil {
		return "", err
	}
	if _, err := t.run(t.lame, "--silent", "-m", "m", "-b", "96", aiff, mp3); err != nil {
		return "", err
	}
	data, _, err := readRegular(mp3)
	if err != nil {
		return "", err
	}
	if len(data) == 0 {
		return "", errors.New("encoder produced empty audio")
	}
	if err := t.decode(mp3); err != nil {
		return "", fmt.Errorf("generated audio failed decoding: %w", err)
	}
	if err := sourceUnchanged(source, entry.sourceData); err != nil {
		return "", err
	}
	if err := os.Chmod(mp3, 0644); err != nil {
		return "", err
	}
	// A hard link publishes a new file without replacing a concurrently created one.
	if !exists {
		err = os.Link(mp3, target)
	} else {
		current, _, readErr := readRegular(target)
		if readErr != nil || !bytes.Equal(current, previous) {
			return "", errors.New("audio changed during generation; existing audio was not replaced")
		}
		err = os.Rename(mp3, target)
	}
	if err != nil {
		return "", err
	}
	if err := replaceSource(source, entry.sourceData, updated); err != nil {
		var rollback error
		if exists {
			backup := filepath.Join(tmp, "previous.mp3")
			rollback = os.WriteFile(backup, previous, mode)
			if rollback == nil {
				rollback = os.Rename(backup, target)
			}
		} else {
			rollback = os.Remove(target)
		}
		return "", errors.Join(err, rollback)
	}
	return "generated", nil
}

// insertAudio preserves source bytes, including comments and scalar styles.
func insertAudio(data []byte, filename string) ([]byte, error) {
	var document yaml.Node
	if err := yaml.Unmarshal(data, &document); err != nil {
		return nil, err
	}
	mapping := document.Content[0]
	lines := bytes.SplitAfter(data, []byte("\n"))
	if mapping.Style&yaml.FlowStyle != 0 {
		offset := len(bytes.Join(lines[:mapping.Line-1], nil))
		found := false
		for _, line := range lines[mapping.Line-1:] {
			content := bytes.SplitN(line, []byte("#"), 2)[0]
			if index := bytes.IndexByte(content, '{'); index >= 0 {
				offset += index + 1
				found = true
				break
			}
			offset += len(line)
		}
		if !found {
			return nil, errors.New("cannot locate the YAML mapping for audio insertion")
		}
		result := append([]byte(nil), data[:offset]...)
		result = append(result, []byte("audio: "+filename+", ")...)
		return checkAudioInsertion(data, append(result, data[offset:]...), filename)
	}
	line := len(lines)
	for i := 0; i < len(mapping.Content); i += 2 {
		if mapping.Content[i].Value == "tags" {
			line = mapping.Content[i].Line - 1
			break
		}
	}
	if line == len(lines) {
		for i, value := range lines {
			if bytes.HasPrefix(value, []byte("...")) {
				rest := strings.TrimSpace(string(value[3:]))
				if rest == "" || strings.HasPrefix(rest, "#") {
					line = i
					break
				}
			}
		}
	}
	keyLine := lines[mapping.Content[0].Line-1]
	indent := string(keyLine[:len(keyLine)-len(bytes.TrimLeft(keyLine, " "))])
	newline := "\n"
	if bytes.Contains(data, []byte("\r\n")) {
		newline = "\r\n"
	}
	prefix := bytes.Join(lines[:line], nil)
	if len(prefix) > 0 && prefix[len(prefix)-1] != '\n' {
		prefix = append(prefix, newline...)
	}
	result := append(prefix, []byte(indent+"audio: "+filename+newline)...)
	result = append(result, bytes.Join(lines[line:], nil)...)
	return checkAudioInsertion(data, result, filename)
}

func checkAudioInsertion(before, after []byte, filename string) ([]byte, error) {
	original, err := parse(before)
	if err != nil {
		return nil, err
	}
	updated, err := parse(after)
	if err != nil {
		return nil, fmt.Errorf("cannot insert audio without changing YAML structure: %w", err)
	}
	original.Audio = filename
	if original.ID != updated.ID || original.Language != updated.Language || original.Word != updated.Word ||
		original.Pronunciation != updated.Pronunciation || original.PartOfSpeech != updated.PartOfSpeech ||
		original.Meaning != updated.Meaning || original.Context != updated.Context ||
		original.Example != updated.Example || original.Notes != updated.Notes ||
		original.Audio != updated.Audio || !slices.Equal(original.Tags, updated.Tags) {
		return nil, errors.New("audio insertion would change other YAML values; source was not modified")
	}
	return after, nil
}

func replaceSource(path string, before, after []byte) error {
	if err := sourceUnchanged(path, before); err != nil {
		return err
	}
	if bytes.Equal(before, after) {
		return nil
	}
	_, mode, err := readRegular(path)
	if err != nil {
		return err
	}
	file, err := os.CreateTemp(filepath.Dir(path), ".audio-entry-")
	if err != nil {
		return err
	}
	defer os.Remove(file.Name())
	defer file.Close()
	if err := file.Chmod(mode); err != nil {
		return err
	}
	if _, err := file.Write(after); err != nil {
		return err
	}
	if err := file.Close(); err != nil {
		return err
	}
	if err := sourceUnchanged(path, before); err != nil {
		return err
	}
	return os.Rename(file.Name(), path)
}
