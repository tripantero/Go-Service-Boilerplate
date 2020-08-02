package helpers

import (
	"io"

	jsoniter "github.com/json-iterator/go"
)

var (
	json = jsoniter.ConfigCompatibleWithStandardLibrary
)

// ByteToStruct provide functionalities to transform array of byte
// to struct by json tags, i don't use the standard go json library because performance issues
// return error if parsed didn't work.
func ByteToStruct(bytes []byte, reference interface{}) error {
	err := json.Unmarshal(bytes, reference)
	return err
}

// JSONNewEncoder generate new JSON encoder from writer
func JSONNewEncoder(w io.Writer) *jsoniter.Encoder {
	return json.NewEncoder(w)
}
