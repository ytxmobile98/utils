package utils

import (
	"encoding/json"
	"os"
)

const PrettyPrintMaxIndent = 10

func ReadJSONData[T any](bytes []byte) (*T, error) {
	var result T
	err := json.Unmarshal(bytes, &result)
	return &result, err
}

func ReadJSONFile[T any](filename string) (*T, error) {
	bytes, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return ReadJSONData[T](bytes)
}

func PrettyPrintJSON(data any, indent uint) ([]byte, error) {
	// generate spaces according to indent
	spaces := func(indent uint) string {
		indent = min(indent, PrettyPrintMaxIndent)

		spaces := make([]byte, indent)
		for i := range spaces {
			spaces[i] = ' '
		}
		return string(spaces)
	}(indent)

	bytes, err := json.MarshalIndent(data, "", spaces)
	if err != nil {
		return nil, err
	}
	return bytes, nil
}

func PrettyPrintJSONFile(filename string, indent uint) ([]byte, error) {
	data, err := ReadJSONFile[any](filename)
	if err != nil {
		return nil, err
	}
	return PrettyPrintJSON(data, indent)
}

func MarshalJSONData(data any) ([]byte, error) {
	return json.Marshal(data)
}
