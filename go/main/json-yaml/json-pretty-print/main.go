package main

import (
	"flag"
	"fmt"

	"github.com/ytxmobile98/utils/go/utils"
	"github.com/ytxmobile98/utils/go/utils/converters"
)

var args struct {
	inputFilename  string
	inputJSONText  string
	outputFilename string

	indent uint
}

const defaultIndent = utils.PrettyPrintDefaultIndent

func init() {
	utils.ParseFlagsAndCheckErrors(defineAndParseArgs, checkArgs, 1)
}

func defineAndParseArgs() {
	flag.StringVar(&args.inputFilename, "i", "", "input JSON file (optional)")
	flag.StringVar(&args.inputJSONText, "t", "", "input JSON text (optional； if both `-i` and `-t` are not specified, read from stdin)")
	flag.StringVar(&args.outputFilename, "o", "", "output JSON file (optional; if not specified, write to stdout)")

	flag.UintVar(&args.indent, "p", defaultIndent, fmt.Sprintf("number of spaces used for pretty indent, max: %d; default: %d", utils.PrettyPrintMaxIndent, defaultIndent))

	flag.Parse()
}

func checkArgs(errs *[]error) {}

func main() {
	var converter converters.Converter = converters.GetJSONLayoutConverter(args.indent)

	var err = func() (err error) {
		if args.inputFilename != "" {
			// If input filename is set, read from the file.
			_, err = converters.ConvertFile(
				args.inputFilename, args.outputFilename, converter)
			return
		} else if args.inputJSONText != "" {
			// If input JSON text is set, convert it.
			_, err = converters.ConvertBytes(
				[]byte(args.inputJSONText), args.outputFilename, converter)
			return
		} else {
			// Read from stdin.
			var bytes []byte
			bytes, err = utils.ReadFile("")
			if err != nil {
				return
			}
			_, err = converters.ConvertBytes(
				bytes, args.outputFilename, converter)
			return
		}
	}()

	if err != nil {
		panic(err)
	}
}
