package vocabulary

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"reflect"
	"strings"
	"testing"
)

func TestLearningMetadata(t *testing.T) {
	entry, err := parse([]byte(validEntry))
	if err != nil || entry.LearningPriority != 2 || len(entry.Prerequisites) != 0 {
		t.Fatalf("omitted metadata: %+v, %v", entry, err)
	}
	for _, value := range []string{"0", "4", "-1", "2.0", `"1"`, "true", "null", "[]", "99999999999999999999999999"} {
		t.Run("priority-"+value, func(t *testing.T) {
			if _, err := parse([]byte(validEntry + "learning_priority: " + value + "\n")); err == nil || !strings.Contains(err.Error(), "learning_priority") {
				t.Fatalf("invalid priority accepted: %v", err)
			}
		})
	}
	for _, value := range []string{"[]", "null", "a-tech", "[1]", "[true]", "[Bad-ID]", `[""]`, "[[a-tech]]"} {
		t.Run("prerequisites-"+value, func(t *testing.T) {
			if _, err := parse([]byte(validEntry + "prerequisites: " + value + "\n")); err == nil || !strings.Contains(err.Error(), "prerequisites") {
				t.Fatalf("invalid prerequisites accepted: %v", err)
			}
		})
	}
	for priority := 1; priority <= 3; priority++ {
		entry, err := parse([]byte(validEntry + fmt.Sprintf("learning_priority: %d\nprerequisites: [base-tech]\n", priority)))
		if err != nil || entry.LearningPriority != priority || !reflect.DeepEqual(entry.Prerequisites, []string{"base-tech"}) {
			t.Fatalf("valid metadata: %+v, %v", entry, err)
		}
	}
}

func TestPrerequisiteErrorsPreserveExport(t *testing.T) {
	for _, tc := range []struct{ name, a, b, want string }{
		{"missing", "[missing]", "", `unknown ID "missing"`},
		{"self", "[a]", "", "cannot depend on itself"},
		{"duplicate", "[b, b]", "", `duplicate ID "b"`},
		{"cycle", "[b]", "[a]", "a -> b -> a"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			root := t.TempDir()
			for _, id := range []string{"a", "b"} {
				writeLearningEntry(t, root, id, "")
			}
			if _, _, err := Build(root); err != nil {
				t.Fatal(err)
			}
			before := readExport(t, root)
			writeLearningEntry(t, root, "a", "prerequisites: "+tc.a+"\n")
			if tc.b != "" {
				writeLearningEntry(t, root, "b", "prerequisites: "+tc.b+"\n")
			}
			if _, _, err := Load(root); err == nil || !strings.Contains(err.Error(), "entries/a.yml") || !strings.Contains(err.Error(), tc.want) {
				t.Fatalf("validation error: %v", err)
			}
			if _, _, err := Build(root); err == nil {
				t.Fatal("build accepted invalid prerequisites")
			}
			if !bytes.Equal(before, readExport(t, root)) {
				t.Fatal("invalid graph replaced previous export")
			}
		})
	}
}

func TestStudyOrder(t *testing.T) {
	root := t.TempDir()
	// The first-priority compound must wait for both lower-priority foundations.
	writeLearningEntry(t, root, "a-compound", "learning_priority: 1\nprerequisites: [z-base, y-base]\n")
	writeLearningEntry(t, root, "b-default", "")
	writeLearningEntry(t, root, "c-early", "learning_priority: 1\n")
	writeLearningEntry(t, root, "d-default", "")
	writeLearningEntry(t, root, "e-chain", "learning_priority: 1\nprerequisites: [a-compound, y-base]\n")
	writeLearningEntry(t, root, "y-base", "learning_priority: 3\n")
	writeLearningEntry(t, root, "z-base", "learning_priority: 3\n")
	// Include enough records to test text-sortable rank padding.
	for _, id := range []string{"f-default", "g-default", "h-default", "i-default"} {
		writeLearningEntry(t, root, id, "")
	}
	if _, _, err := Build(root); err != nil {
		t.Fatal(err)
	}
	first := readExport(t, root)
	reader := csv.NewReader(bytes.NewReader(bytes.SplitN(first, []byte("\n"), 6)[5]))
	reader.Comma = '\t'
	reader.FieldsPerRecord = 13
	rows, err := reader.ReadAll()
	if err != nil {
		t.Fatal(err)
	}
	want := []string{"c-early", "b-default", "d-default", "f-default", "g-default", "h-default", "i-default", "y-base", "z-base", "a-compound", "e-chain"}
	for i, row := range rows {
		if row[0] != want[i] || row[11] != fmt.Sprintf("%02d", i+1) {
			t.Fatalf("row %d: %v", i, row)
		}
	}
	// Reordering references has no effect.
	writeLearningEntry(t, root, "a-compound", "learning_priority: 1\nprerequisites: [y-base, z-base]\n")
	if _, _, err := Build(root); err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(first, readExport(t, root)) {
		t.Fatal("prerequisite list order changed export")
	}
	entries, _, err := Load(root)
	if err != nil || entries[0].ID != "a-compound" {
		t.Fatalf("Load must preserve ID order for audio batches: %v", err)
	}
	before := append([]Entry(nil), entries...)
	learningOrder(entries)
	if !reflect.DeepEqual(entries, before) {
		t.Fatal("ordering mutated source entries")
	}
}

func writeLearningEntry(t *testing.T, root, id, metadata string) {
	t.Helper()
	writeFile(t, root, "entries/"+id+".yml", strings.Replace(validEntry, "id: idiomatic-tech", "id: "+id, 1)+metadata)
}
