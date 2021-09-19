package util

import (
	"io"
	"io/ioutil"
	"os"
)

const (
	filePermissions   = 0200
	fileOperationMode = os.O_RDWR | os.O_CREATE | os.O_APPEND
)

// ClearFileContent Clears the content of the file on the address parameter.
// Errors are explicitly ignored since the input file path should be existing.
func ClearFileContent(path string) {
	_ = ioutil.WriteFile(path, []byte(""), filePermissions)
}

// OpenFile Opens and returns an output stream to the file on the address parameter.
// Error is returned on IO errors.
func OpenFile(path string) (io.Writer, error) {
	fileWriter, err := os.OpenFile(path, fileOperationMode, filePermissions)
	if err != nil {
		return fileWriter, err
	}
	return fileWriter, nil
}
