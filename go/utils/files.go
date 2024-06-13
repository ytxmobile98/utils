package utils

import (
	"io"
	"os"
)

const (
	BufferSize uint = 16 << 20
)

// Read a file by filename.
// If filename is not specified, read from stdin.
func ReadFile(filename string) (bytes []byte, err error) {
	if filename != "" {
		return os.ReadFile(filename)
	} else {
		return io.ReadAll(os.Stdin)
	}
}

// Write the given bytes to a file.
// If filename is not specified, write to stdout.
func WriteFile(filename string, bytes []byte) (int, error) {
	if filename != "" {
		// create or truncate a file, and then write into it
		file, err := os.Create(filename)
		if err != nil {
			return 0, err
		}
		defer file.Close()

		return file.Write(bytes)
	} else {
		// filename is not specified, write to stdout
		return os.Stdout.Write(bytes)
	}
}
