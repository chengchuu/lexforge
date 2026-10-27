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
		fmt.Fprintln(os.Stderr, "usage: go run scripts/build/main.go (from the repository root)")
		os.Exit(2)
	}
	count, warnings, err := vocabulary.Build(".")
	for _, warning := range warnings {
		fmt.Fprintln(os.Stderr, "Warning:", warning)
	}
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	fmt.Printf("Built dist/vocabulary.tsv: %d entries\n", count)
}
