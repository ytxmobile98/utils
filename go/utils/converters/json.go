package converters

import (
	"github.com/ytxmobile98/utils/go/utils"
)

// Convert JSON to YAML
func JSONToYAML(input []byte) ([]byte, error) {
	data, err := utils.ReadJSONData[any](input)
	if err != nil {
		return nil, err
	}
	return utils.MarshalYAML(data)
}

// Convert YAML to JSON.
// If `indent` is specified, use it to pretty print the JSON output.
func YAMLToJSON(input []byte, indent uint) ([]byte, error) {
	data, err := utils.ReadYAMLData[any](input)
	if err != nil {
		return nil, err
	}
	return utils.MarshalJSON(data, indent)
}

// Get the converter function using the specified indent
func GetYAMLToJSONConverter(prettyIndent uint) Converter {
	return func(input []byte) ([]byte, error) {
		return YAMLToJSON(input, prettyIndent)
	}
}

// Convert JSON layout.
// If `indent` is 0, the output will be compressed.
// Otherwise, the output will be pretty printed with the specified number of spaces.
// Maximum indent is `utils.PrettyPrintMaxIndent`.
func JSONLayoutConverter(input []byte, indent uint) ([]byte, error) {
	data, err := utils.ReadJSONData[any](input)
	if err != nil {
		return nil, err
	}
	return utils.MarshalJSON(data, indent)
}

// Get the converter function using the specified indent
func GetJSONLayoutConverter(indent uint) Converter {
	return func(input []byte) ([]byte, error) {
		return JSONLayoutConverter(input, indent)
	}
}
