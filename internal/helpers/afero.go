package helpers

import (
	"fmt"

	"github.com/spf13/afero"
)

func MockFilePath(fs afero.Fs, fileName string, contents string) error {
	file, err := fs.Create(fileName)
	if err != nil {
		return fmt.Errorf("failed to create file %s, err: %v", fileName, err)
	}

	_, err = file.WriteString(contents)
	if err != nil {
		return fmt.Errorf("failed to write file %s, err: %v", fileName, err)
	}

	err = file.Close()
	if err != nil {
		return fmt.Errorf("failed to close file %s, err: %v", fileName, err)
	}

	return nil
}
