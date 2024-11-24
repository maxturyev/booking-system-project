package data

import (
	"encoding/json"
	"io"
)

// FromJSON deserializes the struct from JSON
func (h *Hotel) FromJSON(w io.Reader) error {
	e := json.NewDecoder(w)
	return e.Decode(h)
}

// ToJSON serializes the collection to JSON
func (h *Hotels) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(h)
}
