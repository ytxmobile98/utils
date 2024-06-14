package converters

import (
	"github.com/ytxmobile98/utils/go/utils"
	"gopkg.in/yaml.v3"
)

type Converter func([]byte) ([]byte, error)

func Convert(inputFilename string, outputFilename string, convert Converter) (n int, err error) {
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

// Convert JSON to YAML
func JSONToYAML(input []byte) ([]byte, error) {
	data, err := utils.ReadJSONData[any](input)
	if err != nil {
		return nil, err
	}
	return yaml.Marshal(data)
}

// Convert YAML to JSON
// If prettyIndent is specified, use it to pretty print the JSON output
func YAMLToJSON(input []byte, prettyIndent uint) ([]byte, error) {
	data, err := utils.ReadYAMLData[any](input)
	if err != nil {
		return nil, err
	}

	if prettyIndent > 0 {
		return utils.PrettyPrintJSON(data, prettyIndent)
	}
	return utils.MarshalJSONData(data)
}

// Get the converter function by specifying the pretty print indent
func GetYAMLToJSONConverter(prettyIndent uint) Converter {
	return func(input []byte) ([]byte, error) {
		return YAMLToJSON(input, prettyIndent)
	}
}
