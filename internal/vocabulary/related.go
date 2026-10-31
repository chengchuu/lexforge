package vocabulary

import (
	"errors"
	"fmt"
	"sort"
	"strings"
)

func validateRelatedEntries(entries []Entry) error {
	ids := make(map[string]bool, len(entries))
	for _, entry := range entries {
		ids[entry.ID] = true
	}
	var problems []error
	for _, entry := range entries {
		seen := make(map[string]bool)
		for _, id := range entry.RelatedEntries {
			switch {
			case id == entry.ID:
				problems = append(problems, fmt.Errorf("%s: related_entries: self-reference %q", entry.sourcePath, id))
			case seen[id]:
				problems = append(problems, fmt.Errorf("%s: related_entries: duplicate ID %q", entry.sourcePath, id))
			case !ids[id]:
				problems = append(problems, fmt.Errorf("%s: related_entries: missing entry %q", entry.sourcePath, id))
			}
			seen[id] = true
		}
	}
	return errors.Join(problems...)
}

// Associations are direct and undirected; they do not affect learning order.
func relatedWords(entries []Entry) map[string]string {
	byID := make(map[string]Entry, len(entries))
	neighbors := make(map[string]map[string]bool, len(entries))
	for _, entry := range entries {
		byID[entry.ID] = entry
		neighbors[entry.ID] = make(map[string]bool)
	}
	for _, entry := range entries {
		for _, id := range entry.RelatedEntries {
			neighbors[entry.ID][id] = true
			neighbors[id][entry.ID] = true
		}
	}
	result := make(map[string]string, len(entries))
	for _, entry := range entries {
		ids := make([]string, 0, len(neighbors[entry.ID]))
		for id := range neighbors[entry.ID] {
			ids = append(ids, id)
		}
		sort.Strings(ids)
		lines := make([]string, 0, len(ids))
		for _, id := range ids {
			neighbor := byID[id]
			word := neighbor.Word
			if neighbor.Language != entry.Language {
				word += " (" + neighbor.Language + ")"
			}
			lines = append(lines, word+" — "+neighbor.Meaning)
		}
		result[entry.ID] = strings.Join(lines, "\n")
	}
	return result
}
