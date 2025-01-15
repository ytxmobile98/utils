package main

import (
	"flag"
	"fmt"

	"github.com/ytxmobile98/utils/go/utils"
	"github.com/ytxmobile98/utils/go/utils/converters"
)

var args struct {
	inputFilename  string
	inputJsonText  string
	outputFilename string

	indent uint
}

const defaultIndent = utils.PrettyPrintDefaultIndent

func init() {
	utils.ParseFlagsAndCheckErrors(defineAndParseArgs, checkArgs, 1)
}

func defineAndParseArgs() {
	flag.StringVar(&args.inputFilename, "i", "", "input json file")
	flag.StringVar(&args.inputJsonText, "t", "", "input json text (either -i or -t is required)")
	flag.StringVar(&args.outputFilename, "o", "", "output json file (optional; if not specified, write to stdout)")

	flag.UintVar(&args.indent, "p", defaultIndent, fmt.Sprintf("number of spaces used for pretty indent, max: %d; default: %d", utils.PrettyPrintMaxIndent, defaultIndent))

	flag.Parse()
}

func checkArgs(errs *[]error) {}

func main() {
	var converter converters.Converter = converters.GetJSONLayoutConverter(args.indent)

	var err error
	if args.inputFilename != "" {
		_, err = converters.ConvertFile(args.inputFilename, args.outputFilename, converter)
	} else if args.inputJsonText != "" {
		var bytes []byte
		bytes, err = converters.ConvertBytes([]byte(args.inputJsonText), converter)
		fmt.Println(string(bytes))
	} else {
		err = fmt.Errorf("no JSON input specified")
	}
	if err != nil {
		panic(err)
	}
}
