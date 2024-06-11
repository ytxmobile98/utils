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
