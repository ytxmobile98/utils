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
		// read from stdin
		buffer := [BufferSize]byte{}
		var n int
		for err == nil {
			n, err = os.Stdin.Read(buffer[:])
			if n > 0 {
				bytes = append(bytes, buffer[:n]...)
			}
		}
		if err == io.EOF {
			err = nil
		}
	}
	return
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
