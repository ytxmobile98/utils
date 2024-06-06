package utils

import (
	"os"

	"gopkg.in/yaml.v3"
)

func ReadYAMLData[T any](bytes []byte) (*T, error) {
	var result T
	err := yaml.Unmarshal(bytes, &result)
	return &result, err
}

func ReadYAMLFromFile[T any](filename string) (result *T, err error) {
	bytes, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return ReadYAMLData[T](bytes)
}

// Convert YAML to JSON
// If prettyIndent is specified, use it to pretty print the JSON output
func YAMLToJSON(input []byte, prettyIndent uint) (output []byte, err error) {
	data, err := ReadYAMLData[any](input)
	if err != nil {
		return nil, err
	}

	if prettyIndent > 0 {
		return PrettyPrintJSON(data, prettyIndent)
	}
	return MarshalJSONData(data)
}

// Convert JSON to YAML
func JSONToYAML(input []byte) (output []byte, err error) {
	data, err := ReadJSONData[any](input)
	if err != nil {
		return nil, err
	}
	return yaml.Marshal(data)
}
