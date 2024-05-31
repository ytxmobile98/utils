package main

import (
	"fmt"
	"os"

	"github.com/ytxmobile98/utils/go/utils"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "Usage: \"%s\" <input.yaml> [output.json]\n", os.Args[0])
		fmt.Fprintln(os.Stderr, "If output file not specified, output to stdout.")

		os.Exit(1)
	}

	inputFilename := os.Args[1]
	yamlBytes, err := os.ReadFile(inputFilename)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(2)
	}
	jsonBytes, err := utils.YAMLToJSON(yamlBytes)
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}

	// write to output
	// if output file not specified, write to stdout
	if len(os.Args) > 2 {
		outFile := os.Args[2]
		err = os.WriteFile(outFile, jsonBytes, 0644)
	} else {
		_, err = os.Stdout.Write(jsonBytes)
	}
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
}
