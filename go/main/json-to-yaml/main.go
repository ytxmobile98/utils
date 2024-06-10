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
	flag.StringVar(&args.inputFilename, "i", "", "input json file (required)")
	flag.StringVar(&args.outputFilename, "o", "", "output yaml file (optional; if not specified, write to stdout)")

	flag.Parse()
}

func checkArgs(errs *[]error) {
	if args.inputFilename == "" {
		*errs = append(*errs, fmt.Errorf("input json file is required"))
	}
}

func main() {
	jsonBytes, err := os.ReadFile(args.inputFilename)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(2)
	}

	yamlBytes, err := utils.JSONToYAML(jsonBytes)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(2)
	}

	// write to output
	// if output file not specified, write to stdout
	if args.outputFilename != "" {
		err = os.WriteFile(args.outputFilename, yamlBytes, 0644)
	} else {
		_, err = os.Stdout.Write(yamlBytes)
	}
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(2)
	}
}
