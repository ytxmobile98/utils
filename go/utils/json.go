package utils

import (
	"bytes"
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
	// By default, json.Marshal() escapes selected HTML characters.
	// To avoid this, we use `json.Encoder` and set `SetEscapeHTML(false)`.
	// Reference: https://pkg.go.dev/encoding/json#Encoder.SetEscapeHTML
	buf := bytes.Buffer{}
	encoder := json.NewEncoder(&buf)
	encoder.SetEscapeHTML(false)
	encoder.SetIndent("", getSpaces(indent))

	err := encoder.Encode(data)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// generate spaces according to indent
func getSpaces(indent uint) string {
	if indent == 0 {
		return ""
	}

	indent = min(indent, PrettyPrintMaxIndent)
	spaces := bytes.Buffer{}
	for i := uint(0); i < indent; i++ {
		spaces.WriteByte(' ')
	}
	return spaces.String()
}
