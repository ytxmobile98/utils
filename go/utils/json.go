package utils

import (
	"encoding/json"
)

const (
	PrettyPrintDefaultIndent uint = 4
	PrettyPrintMaxIndent     uint = 10
)

func ReadJSONData[T any](bytes []byte) (*T, error) {
	var result T
	err := json.Unmarshal(bytes, &result)
	return &result, err
}

func ReadJSONFile[T any](filename string) (*T, error) {
	bytes, err := ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return ReadJSONData[T](bytes)
}

// Marshal JSON data.
// If indent is 0, then simply marshal the data.
// If indent is greater than 0, then pretty print the JSON data using the specified indent.
// Maximum indent is `PrettyPrintMaxIndent`.
func MarshalJSON(data any, indent uint) ([]byte, error) {
	if indent == 0 {
		return json.Marshal(data)
	} else {
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
}

func PrettyPrintJSONFile(filename string, indent uint) ([]byte, error) {
	data, err := ReadJSONFile[any](filename)
	if err != nil {
		return nil, err
	}
	return MarshalJSON(data, indent)
}
