package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/ytxmobile98/utils/go/utils"
)

type Args struct {
	inputFileName  string
	outputFilename string

	prettyPrintIndent uint
}

var args Args

func init() {
	parseArgs := func() {
		flag.StringVar(&args.inputFileName, "i", "", "input yaml file (required)")
		flag.StringVar(&args.outputFilename, "o", "", "output json file (optional; if not specified, write to stdout)")

		flag.UintVar(&args.prettyPrintIndent, "p", 0, fmt.Sprintf("number of spaces used for pretty indent, max: %d", utils.PrettyPrintMaxIndent))

		flag.Parse()
	}

	checkArgs := func(errs *[]error) {
		if args.inputFileName == "" {
			*errs = append(*errs, fmt.Errorf("input yaml file is required"))
		}
	}

	utils.ParseFlagsAndCheckErrors(parseArgs, checkArgs, 1)
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
