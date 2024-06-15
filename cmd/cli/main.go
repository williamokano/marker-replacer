package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/afero"
	"github.com/williamokano/marker-replacer/pkg/replacer"
)

func main() {
	if err := run(os.Args); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run(runArgs []string) error {
	flagSet := flag.NewFlagSet(runArgs[0], flag.ExitOnError)

	filePath := flagSet.String("file", "", "Path to the file")
	marker := flagSet.String("marker", "", "Marker to identify the section to replace")

	flagSet.Parse(runArgs[1:])

	if *filePath == "" || *marker == "" {
		fmt.Printf("Usage: %s -file <file_path> -marker <marker> <new_content>\n", os.Args[0])
		flag.PrintDefaults()
		os.Exit(1)
	}

	args := flagSet.Args()
	if len(args) == 0 {
		fmt.Printf("Usage: %s -file <file_path> -marker <marker> <new_content>\n", os.Args[0])
		flag.PrintDefaults()
		os.Exit(1)
	}

	newContent := strings.Join(args, " ")

	fileReplacer := replacer.NewFileReplacer(afero.NewOsFs(), *filePath)

	output, err := fileReplacer.Replace(*marker, newContent)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error replacing %s: %s\n", *marker, err)
	}

	fmt.Fprintf(os.Stdout, "%s", output)

	return nil
}
