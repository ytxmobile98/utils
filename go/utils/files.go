package utils

import (
	"io"
	"os"
)

// Read a file by filename.
// If filename is not specified, read from stdin.
func ReadFile(filename string) (bytes []byte, err error) {
	if filename != "" {
		return os.ReadFile(filename)
	} else {
		// read from stdin
		buffer := [16 << 20]byte{}
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
