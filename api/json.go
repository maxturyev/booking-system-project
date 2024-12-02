package api

import (
	"encoding/json"
	"io"
)

// FromJSON deserializes the struct from JSON
func FromJSON(i interface{}, r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(i)
}

// ToJSON serializes the collection to JSON
func ToJSON(i interface{}, w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(i)
}
