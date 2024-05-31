package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/ytxmobile98/utils/go/utils"
)

func init() {
	parseCmdLineArgs()
	checkArgs()
}

type Args struct {
	inputFileName  string
	outputFilename string

	prettyPrintIndent uint
}

var args Args

func parseCmdLineArgs() {
	flag.StringVar(&args.inputFileName, "i", "", "input yaml file (required)")
	flag.StringVar(&args.outputFilename, "o", "", "output json file (optional; if not specified, write to stdout)")

	flag.UintVar(&args.prettyPrintIndent, "p", 0, fmt.Sprintf("number of spaces used for pretty indent, max: %d", utils.PrettyPrintMaxIndent))

	flag.Parse()
}

func checkArgs() {
	errors := make([]error, 0)
	defer func() {
		if len(errors) > 0 {
			fmt.Fprintln(os.Stderr, "Error(s):")
			for _, err := range errors {
				fmt.Fprintln(os.Stderr, "*", err)
			}

			fmt.Fprintln(os.Stderr)
			flag.Usage()
			os.Exit(1)
		}
	}()

	if args.inputFileName == "" {
		errors = append(errors, fmt.Errorf("input yaml file is required"))
	}
}

func main() {
	yamlBytes, err := os.ReadFile(args.inputFileName)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(2)
	}

	jsonBytes, err := utils.YAMLToJSON(yamlBytes, args.prettyPrintIndent)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(2)
	}

	// write to output
	// if output file not specified, write to stdout
	if args.outputFilename != "" {
		err = os.WriteFile(args.outputFilename, jsonBytes, 0644)
	} else {
		_, err = os.Stdout.Write(jsonBytes)
	}
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
}
