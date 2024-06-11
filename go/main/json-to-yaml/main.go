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
}

func init() {
	utils.ParseFlagsAndCheckErrors(defineAndParseArgs, checkArgs, 1)
}

func defineAndParseArgs() {
	flag.StringVar(&args.inputFilename, "i", "", "input json file (optional; if not specified, read from stdin)")
	flag.StringVar(&args.outputFilename, "o", "", "output yaml file (optional; if not specified, write to stdout)")

	flag.Parse()
}

func checkArgs(errs *[]error) {}

func main() {
	jsonBytes, err := utils.ReadFile(args.inputFilename)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(2)
	}

	yamlBytes, err := utils.JSONToYAML(jsonBytes)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(2)
	}

	_, err = utils.WriteFile(args.outputFilename, yamlBytes)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(2)
	}
}
