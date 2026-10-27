package main

import (
	"flag"
	"fmt"
	"os"

	"lexforge/internal/vocabulary"
)

func main() {
	flag.Parse()
	if flag.NArg() != 0 {
		fmt.Fprintln(os.Stderr, "usage: go run scripts/validate/main.go (from the repository root)")
		os.Exit(2)
	}
	entries, warnings, err := vocabulary.Load(".")
	for _, warning := range warnings {
		fmt.Fprintln(os.Stderr, "Warning:", warning)
	}
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	fmt.Printf("Validation successful: %d entries\n", len(entries))
}
