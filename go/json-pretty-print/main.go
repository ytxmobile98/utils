package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/ytxmobile98/utils/go/utils"
)

type Args struct {
	inputFilename  string
	outputFilename string

	indent uint
}

const defaultIndent uint = 4

var args Args

func init() {
	parseArgs := func() {
		flag.StringVar(&args.inputFilename, "i", "", "input json file (required)")
		flag.StringVar(&args.outputFilename, "o", "", "output json file (optional; if not specified, write to stdout)")

		flag.UintVar(&args.indent, "p", defaultIndent, fmt.Sprintf("number of spaces used for pretty indent, max: %d; default: %d", utils.PrettyPrintMaxIndent, defaultIndent))

		flag.Parse()
	}

	checkArgs := func(errs *[]error) {
		if args.inputFilename == "" {
			*errs = append(*errs, fmt.Errorf("input json file is required"))
		}
	}

	utils.ParseFlagsAndCheckErrors(parseArgs, checkArgs, 1)
}

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "Usage: json-pretty-print <input_filename>")
		os.Exit(1)
	}

	inputFilename := args.inputFilename
	output, err := utils.PrettyPrintJSONFile(inputFilename, args.indent)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(2)
	}
	if args.outputFilename != "" {
		err = os.WriteFile(args.outputFilename, output, 0644)
	} else {
		_, err = os.Stdout.Write(output)
	}
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(2)
	}
}
