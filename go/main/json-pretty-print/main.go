package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/ytxmobile98/utils/go/utils"
)

var args struct {
	inputFilename  string
	outputFilename string

	indent uint
}

const defaultIndent = utils.PrettyPrintDefaultIndent

func init() {
	utils.ParseFlagsAndCheckErrors(defineAndParseArgs, checkArgs, 1)
}

func defineAndParseArgs() {
	flag.StringVar(&args.inputFilename, "i", "", "input json file (optional; if not specified, read from stdin)")
	flag.StringVar(&args.outputFilename, "o", "", "output json file (optional; if not specified, write to stdout)")

	flag.UintVar(&args.indent, "p", defaultIndent, fmt.Sprintf("number of spaces used for pretty indent, max: %d; default: %d", utils.PrettyPrintMaxIndent, defaultIndent))

	flag.Parse()
}

func checkArgs(errs *[]error) {}

func main() {
	inputFilename := args.inputFilename
	output, err := utils.PrettyPrintJSONFile(inputFilename, args.indent)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(2)
	}

	_, err = utils.WriteFile(args.outputFilename, output)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(2)
	}
}
