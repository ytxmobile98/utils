package utils

import (
	"gopkg.in/yaml.v3"
)

func ReadYAMLData[T any](bytes []byte) (*T, error) {
	var result T
	err := yaml.Unmarshal(bytes, &result)
	return &result, err
}

func ReadYAMLFromFile[T any](filename string) (*T, error) {
	bytes, err := ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return ReadYAMLData[T](bytes)
}

// Convert YAML to JSON
// If prettyIndent is specified, use it to pretty print the JSON output
func YAMLToJSON(input []byte, prettyIndent uint) ([]byte, error) {
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
func JSONToYAML(input []byte) ([]byte, error) {
	data, err := ReadJSONData[any](input)
	if err != nil {
		return nil, err
	}
	return yaml.Marshal(data)
}
