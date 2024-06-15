package replacer

import (
	"fmt"
	"strings"

	"github.com/spf13/afero"
	"github.com/williamokano/marker-replacer/pkg/stringhelper"
)

type FileReplacer struct {
	Filename string
	fs       afero.Fs
}

func NewFileReplacer(fs afero.Fs, filename string) *FileReplacer {
	return &FileReplacer{
		Filename: filename,
		fs:       fs,
	}
}

func (r *FileReplacer) Replace(marker string, newContent string) (string, error) {

	file, err := r.fs.Open(r.Filename)
	if err != nil {
		return "", fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	startMarker := fmt.Sprintf("<!--%s-->", marker)
	endMarker := fmt.Sprintf("<!--/%s-->", marker)

	input, err := stringhelper.ReadAllLines(r.fs, r.Filename)
	if err != nil {
		return "", fmt.Errorf("failed to read input: %w", err)
	}

	startIndex := strings.Index(input, startMarker)
	endIndex := strings.Index(input, endMarker)

	if startIndex == -1 || endIndex == -1 {
		// If the markers are not found, return the original input
		return input, nil
	}

	startIndex += len(startMarker)

	// Preserve the newline character after the start marker if it exists
	if startIndex < len(input) && input[startIndex] == '\n' {
		startIndex++
	}

	// Preserve the newline character before the end marker if it exists
	if endIndex > 0 && input[endIndex-1] == '\n' {
		endIndex--
	}

	return input[:startIndex] + newContent + input[endIndex:], nil
}
