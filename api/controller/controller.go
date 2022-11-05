package controller

import (
	"encoding/json"
	"io"
)

func ReadJSON(r io.Reader, data any) error {
	dec := json.NewDecoder(r)
	err := dec.Decode(data)
	return err
}
