package main

import (
	"flag"
	"fmt"
	"os"

	"lexforge/internal/vocabulary"
)

func main() {
	var options vocabulary.AudioOptions
	flag.StringVar(&options.ID, "id", "", "generate audio for one stable entry ID")
	flag.BoolVar(&options.All, "all", false, "process all entries in stable-ID order")
	flag.BoolVar(&options.Overwrite, "overwrite", false, "explicitly replace existing recordings")
	flag.StringVar(&options.Voice, "voice", "", "installed voice override (requires --id)")
	flag.IntVar(&options.Rate, "rate", 0, "speech rate; zero uses the language default (requires --id for overrides)")
	flag.StringVar(&options.TextSource, "text-source", "word", "word or pronunciation (Japanese kana; requires --id)")
	flag.Parse()
	if flag.NArg() != 0 {
		fmt.Fprintln(os.Stderr, "usage: go run scripts/audio/main.go (--id <stable-id> | --all) [--overwrite] [--voice <name>] [--rate <wpm>] [--text-source word|pronunciation]")
		os.Exit(2)
	}
	summary, warnings, err := vocabulary.GenerateAudio(".", options, os.Stdout)
	for _, warning := range warnings {
		fmt.Fprintln(os.Stderr, "Warning:", warning)
	}
	fmt.Printf("Audio: %d generated, %d reused, %d skipped\n", summary.Generated, summary.Reused, summary.Skipped)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
