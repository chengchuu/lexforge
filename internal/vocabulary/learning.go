package vocabulary

import (
	"fmt"
	"slices"
	"sort"
	"strings"
)

// Prerequisites describe introduction order, never Anki learning state.
func validatePrerequisites(entries []Entry) error {
	byID := make(map[string]Entry, len(entries))
	for _, entry := range entries {
		byID[entry.ID] = entry
	}
	for _, entry := range entries {
		seen := make(map[string]bool)
		for _, id := range entry.Prerequisites {
			if id == entry.ID {
				return fmt.Errorf("%s: field %q: entry cannot depend on itself", entry.sourcePath, "prerequisites")
			}
			if seen[id] {
				return fmt.Errorf("%s: field %q: duplicate ID %q; keep one occurrence", entry.sourcePath, "prerequisites", id)
			}
			if _, ok := byID[id]; !ok {
				return fmt.Errorf("%s: field %q: unknown ID %q; reference an existing entry", entry.sourcePath, "prerequisites", id)
			}
			seen[id] = true
		}
	}

	done := make(map[string]bool)
	var path []string
	var visit func(string) error
	visit = func(id string) error {
		if start := slices.Index(path, id); start >= 0 {
			cycle := append(slices.Clone(path[start:]), id)
			return fmt.Errorf("%s: field %q: prerequisite cycle: %s", byID[id].sourcePath, "prerequisites", strings.Join(cycle, " -> "))
		}
		if done[id] {
			return nil
		}
		path = append(path, id)
		dependencies := slices.Clone(byID[id].Prerequisites)
		sort.Strings(dependencies)
		for _, dependency := range dependencies {
			if err := visit(dependency); err != nil {
				return err
			}
		}
		path = path[:len(path)-1]
		done[id] = true
		return nil
	}
	for _, entry := range entries {
		if err := visit(entry.ID); err != nil {
			return err
		}
	}
	return nil
}

// learningOrder requires the complete, validated collection. At each step it
// selects the lowest-priority-number ready entry, breaking ties by stable ID.
func learningOrder(entries []Entry) []Entry {
	pending := slices.Clone(entries)
	sort.Slice(pending, func(i, j int) bool {
		if pending[i].LearningPriority != pending[j].LearningPriority {
			return pending[i].LearningPriority < pending[j].LearningPriority
		}
		return pending[i].ID < pending[j].ID
	})
	introduced := make(map[string]bool, len(entries))
	ordered := make([]Entry, 0, len(entries))
	for len(pending) > 0 {
		for i, entry := range pending {
			ready := true
			for _, id := range entry.Prerequisites {
				if !introduced[id] {
					ready = false
					break
				}
			}
			if ready {
				ordered = append(ordered, entry)
				introduced[entry.ID] = true
				pending = slices.Delete(pending, i, i+1)
				break
			}
		}
	}
	return ordered
}
