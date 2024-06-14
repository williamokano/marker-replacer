package stringhelper

import (
	"fmt"
	"io"
	"strings"

	"github.com/spf13/afero"
)

func ReadAllLines(fs afero.Fs, filename string) (string, error) {
	file, err := fs.Open(filename)
	if err != nil {
		return "", fmt.Errorf("error opening file %s: %w", filename, err)
	}

	defer file.Close()

	var sb strings.Builder
	buffer := make([]byte, 1024)
	for {
		n, err := file.Read(buffer)
		if err != nil {
			if err == io.EOF {
				sb.Write(buffer[:n])
				break
			}
			return "", fmt.Errorf("error reading file %s: %w", filename, err)
		}
		sb.Write(buffer[:n])
	}

	// Convert the builder to a string
	return sb.String(), nil
}
