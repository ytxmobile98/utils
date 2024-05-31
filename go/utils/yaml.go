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

func YAMLToJSON(input []byte) (output []byte, err error) {
	var data interface{}

	data, err = ReadYAMLData[interface{}](input)
	if err != nil {
		return nil, err
	}

	output, err = MarshalJsonData(data)
	if err != nil {
		return nil, err
	}

	return output, nil
}
