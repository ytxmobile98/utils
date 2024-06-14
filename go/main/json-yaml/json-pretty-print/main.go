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
	var converter converters.Converter = func(bytes []byte) ([]byte, error) {
		data, err := utils.ReadJSONData[any](bytes)
		if err != nil {
			return nil, err
		}
		return utils.PrettyPrintJSON(data, args.indent)
	}

	_, err := converters.Convert(args.inputFilename, args.outputFilename, converter)
	if err != nil {
		panic(err)
	}
}
