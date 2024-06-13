package main

import (
	"flag"

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
	flag.StringVar(&args.outputFilename, "o", "", "output json file (optional; if not specified, write to stdout)")

	flag.Parse()
}

func checkArgs(errs *[]error) {}

func main() {
	var converter utils.Converter = func(bytes []byte) ([]byte, error) {
		data, err := utils.ReadJSONData[any](bytes)
		if err != nil {
			return nil, err
		}
		return utils.MarshalJSONData(data)
	}

	_, err := utils.Convert(args.inputFilename, args.outputFilename, converter)
	if err != nil {
		panic(err)
	}
}
