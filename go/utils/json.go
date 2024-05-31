package utils

import (
	"encoding/json"
	"os"
)

const maxIndent = 10

func ReadJsonData[T interface{}](bytes []byte) (*T, error) {
	var result T
	err := json.Unmarshal(bytes, &result)
	return &result, err
}

func ReadJsonFile[T interface{}](filename string) (*T, error) {
	bytes, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return ReadJsonData[T](bytes)
}

func PrettyPrintJSON(data interface{}, indent uint) (string, error) {
	// generate spaces according to indent
	spaces := func(indent uint) string {
		if indent > maxIndent {
			indent = maxIndent
		}

		spaces := make([]byte, indent)
		for i := range spaces {
			spaces[i] = ' '
		}
		return string(spaces)
	}(indent)

	bytes, err := json.MarshalIndent(data, "", spaces)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func PrettyPrintJsonFile(filename string, indent uint) (string, error) {
	data, err := ReadJsonFile[interface{}](filename)
	if err != nil {
		return "", err
	}
	return PrettyPrintJSON(data, indent)
}

func MarshalJsonData(data interface{}) ([]byte, error) {
	return json.Marshal(data)
}
