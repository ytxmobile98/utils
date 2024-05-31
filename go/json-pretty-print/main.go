package main

import (
	"fmt"
	"os"

	"github.com/ytxmobile98/utils/go/utils"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "Usage: json-pretty-print <input_filename>")
		os.Exit(1)
	}

	inputFilename := os.Args[1]
	output, err := utils.PrettyPrintJSONFile(inputFilename, 4)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(2)
	}
	fmt.Println(output)
}
