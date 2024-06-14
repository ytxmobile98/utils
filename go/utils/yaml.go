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
