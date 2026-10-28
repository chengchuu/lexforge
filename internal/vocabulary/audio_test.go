package vocabulary

import (
	"bytes"
	"errors"
	"io"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"testing"
)

func fakeAudioTools(t *testing.T, hook func(string, []string) error) func() (audioTools, error) {
	t.Helper()
	return func() (audioTools, error) {
		return audioTools{say: "say", lame: "lame", decoder: "mpg123", run: func(name string, args ...string) ([]byte, error) {
			if hook != nil {
				if err := hook(name, args); err != nil {
					return nil, err
				}
			}
			if name == "lame" {
				return nil, os.WriteFile(args[len(args)-1], []byte("new MP3"), 0644)
			}
			if name == "mpg123" {
				return []byte{0, 0}, nil
			}
			return nil, nil
		}}, nil
	}
}

func TestAudioSelection(t *testing.T) {
	root := t.TempDir()
	writeFile(t, root, "entries/idiomatic.yml", validEntry)
	for _, options := range []AudioOptions{{}, {All: true, ID: "idiomatic-tech"}, {ID: "unknown"}} {
		_, _, err := generateAudio(root, options, io.Discard, func() (audioTools, error) {
			t.Fatal("invalid selector reached dependency checks")
			return audioTools{}, nil
		})
		if err == nil {
			t.Fatalf("accepted %+v", options)
		}
	}
}

func TestAudioPrerequisites(t *testing.T) {
	lookup := func(name string) (string, error) { return name, nil }
	run := func(string, ...string) ([]byte, error) { return []byte("Samantha en_US # Hello\n"), nil }
	if _, err := findAudioTools("linux", lookup, run); err == nil {
		t.Fatal("accepted unsupported platform")
	}
	for _, missing := range []string{"say", "lame", "mpg123"} {
		_, err := findAudioTools("darwin", func(name string) (string, error) {
			if name == missing {
				return "", os.ErrNotExist
			}
			return name, nil
		}, run)
		if err == nil || !strings.Contains(err.Error(), missing) {
			t.Fatalf("missing %s: %v", missing, err)
		}
	}
	for _, output := range []string{"", "Samantha2 en_US", "Samantha en_GB"} {
		if _, err := findAudioTools("darwin", lookup, func(string, ...string) ([]byte, error) { return []byte(output), nil }); err == nil {
			t.Fatalf("accepted voice list %q", output)
		}
	}
	if _, err := findAudioTools("darwin", lookup, func(string, ...string) ([]byte, error) { return nil, errors.New("voice failure") }); err == nil {
		t.Fatal("ignored voice-list failure")
	}
	if _, err := findAudioTools("darwin", lookup, run); err != nil {
		t.Fatal(err)
	}
}

func TestInsertAudioPreservesYAML(t *testing.T) {
	for _, data := range []string{
		validEntry,
		strings.TrimSuffix(validEntry, "\n"),
		"---\n# entry\n" + validEntry + "notes: |\n  Keep this text.\n# tags\ntags: [tech] # inline\n...\n",
		validEntry + "notes: 'quoted: text'\n... # end\n",
		validEntry + "notes: |+\n  Keep the trailing blank line.\n\n",
		validEntry + "notes: |\n  ...\n  This is literal text.\n",
		"{id: idiomatic-tech, word: idiomatic, meaning: natural, context: usage, example: Use idiomatic Go.}\n",
		"!!map # a comment containing {\n{id: idiomatic-tech, word: idiomatic, meaning: natural, context: usage, example: Use idiomatic Go.}\n",
	} {
		updated, err := insertAudio([]byte(data), "idiomatic-tech.mp3")
		if err != nil {
			t.Fatal(err)
		}
		before, _ := parse([]byte(data))
		after, err := parse(updated)
		if err != nil {
			t.Fatal(err)
		}
		after.Audio = ""
		if !reflect.DeepEqual(before, after) {
			t.Fatalf("changed entry: %s", updated)
		}
		removed := strings.Replace(string(updated), "audio: idiomatic-tech.mp3\n", "", 1)
		removed = strings.Replace(removed, "audio: idiomatic-tech.mp3, ", "", 1)
		if strings.TrimSuffix(removed, "\n") != strings.TrimSuffix(data, "\n") {
			t.Fatalf("changed formatting: %s", updated)
		}
	}
}

func TestAudioGenerationReuseAndSkip(t *testing.T) {
	root := t.TempDir()
	writeFile(t, root, "entries/z.yml", strings.Replace(validEntry, "id: idiomatic-tech", "id: a-tech", 1))
	writeFile(t, root, "entries/a.yml", strings.Replace(validEntry, "id: idiomatic-tech", "id: z-tech", 1))
	writeFile(t, root, "media/z-tech.mp3", "existing MP3")
	var output bytes.Buffer
	var spoken []string
	tools := fakeAudioTools(t, func(name string, args []string) error {
		if name == "say" {
			if !reflect.DeepEqual(args[:5], []string{"-v", "Samantha", "-r", "150", "-f"}) {
				t.Fatalf("unexpected speech settings: %v", args)
			}
			data, err := os.ReadFile(args[5])
			if err != nil {
				return err
			}
			spoken = append(spoken, string(data))
		}
		return nil
	})
	summary, _, err := generateAudio(root, AudioOptions{All: true}, &output, tools)
	if err != nil || summary != (AudioSummary{Generated: 1, Reused: 1}) {
		t.Fatalf("summary=%+v err=%v", summary, err)
	}
	if output.String() != "a-tech: generated\nz-tech: reused\n" || !reflect.DeepEqual(spoken, []string{"idiomatic"}) {
		t.Fatalf("order or speech: %s %v", output.String(), spoken)
	}
	if _, _, err := Load(root); err != nil {
		t.Fatal(err)
	}
	before, _ := os.ReadFile(filepath.Join(root, "media/z-tech.mp3"))
	summary, _, err = generateAudio(root, AudioOptions{All: true}, io.Discard, fakeAudioTools(t, func(string, []string) error {
		t.Fatal("skipped audio invoked a subprocess")
		return nil
	}))
	after, _ := os.ReadFile(filepath.Join(root, "media/z-tech.mp3"))
	if err != nil || summary.Skipped != 2 || !bytes.Equal(before, after) {
		t.Fatalf("rerun: %+v %v", summary, err)
	}
	summary, _, err = generateAudio(root, AudioOptions{ID: "z-tech", Overwrite: true}, io.Discard, tools)
	after, _ = os.ReadFile(filepath.Join(root, "media/z-tech.mp3"))
	if err != nil || summary.Generated != 1 || string(after) != "new MP3" {
		t.Fatalf("overwrite: %+v %v %q", summary, err, after)
	}
}

func TestAudioFailuresPreserveEntryAndRecording(t *testing.T) {
	for _, existing := range []bool{false, true} {
		for _, failure := range []string{"say", "lame", "mpg123", "source-change"} {
			t.Run(failure+fmtBool(existing), func(t *testing.T) {
				root := t.TempDir()
				data := validEntry
				if existing {
					data += "audio: idiomatic-tech.mp3\n"
					writeFile(t, root, "media/idiomatic-tech.mp3", "original MP3")
				}
				writeFile(t, root, "entries/idiomatic.yml", data)
				_, _, err := generateAudio(root, AudioOptions{All: true, Overwrite: true}, io.Discard, fakeAudioTools(t, func(name string, args []string) error {
					if name == failure {
						return errors.New("subprocess failed")
					}
					if failure == "source-change" && name == "mpg123" {
						writeFile(t, root, "entries/idiomatic.yml", data+"# user edit\n")
					}
					return nil
				}))
				if err == nil || !strings.Contains(err.Error(), "idiomatic-tech") {
					t.Fatalf("error: %v", err)
				}
				got, _ := os.ReadFile(filepath.Join(root, "entries/idiomatic.yml"))
				want := data
				if failure == "source-change" {
					want += "# user edit\n"
				}
				if string(got) != want {
					t.Fatalf("source changed: %s", got)
				}
				mp3, err := os.ReadFile(filepath.Join(root, "media/idiomatic-tech.mp3"))
				if existing && (err != nil || string(mp3) != "original MP3") {
					t.Fatalf("lost recording: %q %v", mp3, err)
				}
				if !existing && !os.IsNotExist(err) {
					t.Fatalf("published failed recording: %v", err)
				}
				if _, _, err := Load(root); err != nil {
					t.Fatal(err)
				}
				files, _ := filepath.Glob(filepath.Join(root, "media/.audio-*"))
				if len(files) != 0 {
					t.Fatalf("temporary files remain: %v", files)
				}
			})
		}
	}
}

func fmtBool(value bool) string {
	if value {
		return "-existing"
	}
	return "-new"
}

func TestAudioBatchFailureCanResume(t *testing.T) {
	root := t.TempDir()
	writeFile(t, root, "entries/a.yml", strings.Replace(validEntry, "id: idiomatic-tech", "id: a-tech", 1))
	writeFile(t, root, "entries/b.yml", strings.Replace(validEntry, "id: idiomatic-tech", "id: b-tech", 1))
	calls := 0
	summary, _, err := generateAudio(root, AudioOptions{All: true}, io.Discard, fakeAudioTools(t, func(name string, args []string) error {
		if name == "say" {
			calls++
			if calls == 2 {
				return errors.New("interrupted")
			}
		}
		return nil
	}))
	if err == nil || summary.Generated != 1 {
		t.Fatalf("%+v %v", summary, err)
	}
	if _, _, err := Load(root); err != nil {
		t.Fatal(err)
	}
	summary, _, err = generateAudio(root, AudioOptions{All: true}, io.Discard, fakeAudioTools(t, nil))
	if err != nil || summary.Generated != 1 || summary.Skipped != 1 {
		t.Fatalf("%+v %v", summary, err)
	}
}

func TestAudioRejectsInvalidExistingMedia(t *testing.T) {
	for _, content := range []string{"", "invalid MP3"} {
		root := t.TempDir()
		writeFile(t, root, "entries/idiomatic.yml", validEntry)
		writeFile(t, root, "media/idiomatic-tech.mp3", content)
		_, _, err := generateAudio(root, AudioOptions{All: true}, io.Discard, fakeAudioTools(t, func(name string, args []string) error {
			if name == "mpg123" {
				return errors.New("invalid audio")
			}
			t.Fatalf("unexpected synthesis: %s", name)
			return nil
		}))
		if err == nil {
			t.Fatal("accepted invalid MP3")
		}
		data, _ := os.ReadFile(filepath.Join(root, "entries/idiomatic.yml"))
		if string(data) != validEntry {
			t.Fatalf("attached invalid MP3: %s", data)
		}
	}
}

func TestAudioConcurrentFileCreationIsNotOverwritten(t *testing.T) {
	root := t.TempDir()
	writeFile(t, root, "entries/idiomatic.yml", validEntry)
	_, _, err := generateAudio(root, AudioOptions{All: true}, io.Discard, fakeAudioTools(t, func(name string, args []string) error {
		if name == "mpg123" {
			writeFile(t, root, "media/idiomatic-tech.mp3", "another recording")
		}
		return nil
	}))
	if err == nil {
		t.Fatal("replaced a concurrently created recording")
	}
	data, _ := os.ReadFile(filepath.Join(root, "media/idiomatic-tech.mp3"))
	if string(data) != "another recording" {
		t.Fatalf("recording changed: %q", data)
	}
	data, _ = os.ReadFile(filepath.Join(root, "entries/idiomatic.yml"))
	if string(data) != validEntry {
		t.Fatal("source changed")
	}
}

func TestAudioRejectsEmptyEncoderOutput(t *testing.T) {
	root := t.TempDir()
	writeFile(t, root, "entries/idiomatic.yml", validEntry)
	_, _, err := generateAudio(root, AudioOptions{All: true}, io.Discard, func() (audioTools, error) {
		return audioTools{say: "say", lame: "lame", decoder: "mpg123", run: func(name string, args ...string) ([]byte, error) {
			if name == "lame" {
				return nil, os.WriteFile(args[len(args)-1], nil, 0644)
			}
			if name == "mpg123" {
				t.Fatal("attempted to decode empty output")
			}
			return nil, nil
		}}, nil
	})
	if err == nil || !strings.Contains(err.Error(), "empty audio") {
		t.Fatalf("error: %v", err)
	}
	if _, err := os.Stat(filepath.Join(root, "media/idiomatic-tech.mp3")); !os.IsNotExist(err) {
		t.Fatal("published empty audio")
	}
}

func TestReplaceSourceRejectsConcurrentEdit(t *testing.T) {
	root := t.TempDir()
	path := filepath.Join(root, "entry.yml")
	writeFile(t, root, "entry.yml", validEntry+"# new edit\n")
	err := replaceSource(path, []byte(validEntry), []byte(validEntry+"audio: idiomatic-tech.mp3\n"))
	if err == nil {
		t.Fatal("overwrote a changed source")
	}
	data, _ := os.ReadFile(path)
	if string(data) != validEntry+"# new edit\n" {
		t.Fatalf("source changed: %s", data)
	}
}

func TestAudioFailureToAttachRollsBackRecording(t *testing.T) {
	for _, existing := range []bool{false, true} {
		t.Run(fmtBool(existing), func(t *testing.T) {
			root := t.TempDir()
			writeFile(t, root, "entries/idiomatic.yml", validEntry)
			if existing {
				writeFile(t, root, "media/idiomatic-tech.mp3", "original MP3")
			}
			dir := filepath.Join(root, "entries")
			if err := os.Chmod(dir, 0555); err != nil {
				t.Fatal(err)
			}
			t.Cleanup(func() { os.Chmod(dir, 0755) })
			probe, err := os.CreateTemp(dir, ".permission-check-")
			if err == nil {
				probe.Close()
				os.Remove(probe.Name())
				t.Skip("filesystem permits writes despite directory permissions")
			}
			_, _, err = generateAudio(root, AudioOptions{All: true, Overwrite: true}, io.Discard, fakeAudioTools(t, nil))
			if err == nil {
				t.Fatal("source replacement unexpectedly succeeded")
			}
			data, _ := os.ReadFile(filepath.Join(dir, "idiomatic.yml"))
			if string(data) != validEntry {
				t.Fatalf("changed source: %s", data)
			}
			data, err = os.ReadFile(filepath.Join(root, "media/idiomatic-tech.mp3"))
			if existing && (err != nil || string(data) != "original MP3") {
				t.Fatalf("lost original audio: %s %v", data, err)
			}
			if !existing && !os.IsNotExist(err) {
				t.Fatalf("failed audio was retained: %v", err)
			}
		})
	}
}

func TestAudioRejectsDecoderWithoutSamples(t *testing.T) {
	for _, existing := range []bool{false, true} {
		t.Run(fmtBool(existing), func(t *testing.T) {
			root := t.TempDir()
			writeFile(t, root, "entries/idiomatic.yml", validEntry)
			if existing {
				writeFile(t, root, "media/idiomatic-tech.mp3", "not an MP3")
			}
			prepare := fakeAudioTools(t, nil)
			_, _, err := generateAudio(root, AudioOptions{All: true}, io.Discard, func() (audioTools, error) {
				tools, err := prepare()
				run := tools.run
				tools.run = func(name string, args ...string) ([]byte, error) {
					if name == "mpg123" {
						return nil, nil // mpg123 can succeed without decoding any frames.
					}
					return run(name, args...)
				}
				return tools, err
			})
			if err == nil {
				t.Fatal("accepted a recording with no decoded audio samples")
			}
			data, _ := os.ReadFile(filepath.Join(root, "entries/idiomatic.yml"))
			if string(data) != validEntry {
				t.Fatal("linked an undecodable recording")
			}
			data, readErr := os.ReadFile(filepath.Join(root, "media/idiomatic-tech.mp3"))
			if existing && (readErr != nil || string(data) != "not an MP3") {
				t.Fatal("changed the existing recording")
			}
			if !existing && !os.IsNotExist(readErr) {
				t.Fatal("published an undecodable recording")
			}
		})
	}
}

func TestInsertAudioAcceptsValidBlockLayouts(t *testing.T) {
	cases := map[string]string{
		"indented":               "  " + strings.ReplaceAll(strings.TrimSuffix(validEntry, "\n"), "\n", "\n  ") + "\n  tags: [tech]\n",
		"indented without tags":  "  " + strings.ReplaceAll(strings.TrimSuffix(validEntry, "\n"), "\n", "\n  ") + "\n",
		"end marker with spaces": validEntry + "...   # document end\n",
		"end marker with tab":    validEntry + "...\t# document end\n",
		"CRLF":                   strings.ReplaceAll(validEntry+"tags: [tech]\n", "\n", "\r\n"),
	}
	for name, data := range cases {
		t.Run(name, func(t *testing.T) {
			before, err := parse([]byte(data))
			if err != nil {
				t.Fatalf("invalid test input: %v", err)
			}
			updated, err := insertAudio([]byte(data), "idiomatic-tech.mp3")
			if err != nil {
				t.Fatal(err)
			}
			after, err := parse(updated)
			if err != nil {
				t.Fatal(err)
			}
			if after.Audio != "idiomatic-tech.mp3" {
				t.Fatal("missing audio field")
			}
			after.Audio = ""
			if !reflect.DeepEqual(before, after) {
				t.Fatal("changed other YAML values")
			}
			if name == "CRLF" && bytes.Contains(bytes.ReplaceAll(updated, []byte("\r\n"), nil), []byte("\n")) {
				t.Fatal("introduced mixed line endings")
			}
		})
	}
}

func TestAudioReuseRejectsRecordingChangedDuringDecoding(t *testing.T) {
	for _, removed := range []bool{false, true} {
		t.Run(fmtBool(removed), func(t *testing.T) {
			root := t.TempDir()
			writeFile(t, root, "entries/idiomatic.yml", validEntry)
			writeFile(t, root, "media/idiomatic-tech.mp3", "original MP3")
			_, _, err := generateAudio(root, AudioOptions{All: true}, io.Discard, fakeAudioTools(t, func(name string, args []string) error {
				if name == "mpg123" {
					if removed {
						return os.Remove(filepath.Join(root, "media/idiomatic-tech.mp3"))
					}
					writeFile(t, root, "media/idiomatic-tech.mp3", "unverified replacement")
				}
				return nil
			}))
			if err == nil {
				t.Fatal("attached changed or deleted audio")
			}
			data, _ := os.ReadFile(filepath.Join(root, "entries/idiomatic.yml"))
			if string(data) != validEntry {
				t.Fatal("changed source")
			}
		})
	}
}

func TestRunAudioToolSeparatesDiagnostics(t *testing.T) {
	if os.Getenv("LEXFORGE_TEST_AUDIO_STDERR") == "1" {
		os.Stderr.WriteString("decoder warning without PCM output\n")
		os.Exit(0)
	}
	t.Setenv("LEXFORGE_TEST_AUDIO_STDERR", "1")
	data, err := runAudioTool(os.Args[0], "-test.run=^TestRunAudioToolSeparatesDiagnostics$")
	if err != nil {
		t.Fatal(err)
	}
	if len(data) != 0 {
		t.Fatalf("diagnostics were returned as decoded samples: %q", data)
	}
}
