package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/spf13/afero"
	"github.com/williamokano/marker-replacer/pkg/replacer"
)

func isStdinAvailable() bool {
	info, err := os.Stdin.Stat()
	if err != nil {
		return false
	}
	return info.Mode()&os.ModeCharDevice == 0
}

func readStdin() (string, error) {
	var builder strings.Builder
	reader := bufio.NewReader(os.Stdin)
	_, err := io.Copy(&builder, reader)
	if err != nil {
		return "", err
	}
	return builder.String(), nil
}

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
		flagSet.PrintDefaults()
		os.Exit(1)
	}

	args := flagSet.Args()

	var newContent string
	if isStdinAvailable() {
		stdinContent, err := readStdin()
		if err != nil {
			return fmt.Errorf("failed to read from stdin: %w", err)
		}
		newContent = stdinContent
	} else if len(args) > 0 {
		newContent = strings.Join(args, " ")
	} else {
		fmt.Printf("Usage: %s -file <file_path> -marker <marker> <new_content>\n", os.Args[0])
		flagSet.PrintDefaults()
		os.Exit(1)
	}

	fileReplacer := replacer.NewFileReplacer(afero.NewOsFs(), *filePath)

	output, err := fileReplacer.Replace(*marker, newContent)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error replacing %s: %s\n", *marker, err)
	}

	fmt.Fprintf(os.Stdout, "%s", output)

	return nil
}
