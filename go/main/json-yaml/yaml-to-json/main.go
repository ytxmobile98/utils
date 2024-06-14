package main

import (
	"flag"
	"fmt"

	"github.com/ytxmobile98/utils/go/utils"
	"github.com/ytxmobile98/utils/go/utils/converters"
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
	var converter converters.Converter = converters.GetYAMLToJSONConverter(args.prettyPrintIndent)

	_, err := converters.Convert(args.inputFilename, args.outputFilename, converter)
	if err != nil {
		panic(err)
	}
}
