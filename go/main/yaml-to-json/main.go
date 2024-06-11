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

	prettyPrintIndent uint
}

func init() {
	utils.ParseFlagsAndCheckErrors(defineAndParseArgs, checkArgs, 1)
}

func defineAndParseArgs() {
	flag.StringVar(&args.inputFilename, "i", "", "input yaml file (optional; if not specified, read from stdin)")
	flag.StringVar(&args.outputFilename, "o", "", "output json file (optional; if not specified, write to stdout)")

	flag.UintVar(&args.prettyPrintIndent, "p", 0, fmt.Sprintf("number of spaces used for pretty indent, max: %d", utils.PrettyPrintMaxIndent))

	flag.Parse()
}

func checkArgs(errs *[]error) {}

func main() {
	yamlBytes, err := utils.ReadFile(args.inputFilename)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(2)
	}

	jsonBytes, err := utils.YAMLToJSON(yamlBytes, args.prettyPrintIndent)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(2)
	}

	_, err = utils.WriteFile(args.outputFilename, jsonBytes)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(2)
	}
}
