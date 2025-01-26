package converters

import (
	"github.com/ytxmobile98/utils/go/utils"
)

type Converter func([]byte) ([]byte, error)

func ConvertFile(inputFilename string, outputFilename string, convert Converter) (n int, err error) {
	bytes, err := utils.ReadFile(inputFilename)
	if err != nil {
		return
	}

	bytes, err = convert(bytes)
	if err != nil {
		return
	}

	return utils.WriteFile(outputFilename, bytes)
}

func ConvertBytes(input []byte, outputFilename string, convert Converter) ([]byte, error) {
	bytes, err := convert(input)
	if err != nil {
		return nil, err
	}

	_, err = utils.WriteFile(outputFilename, bytes)
	if err != nil {
		return nil, err
	}

	return bytes, nil
}
