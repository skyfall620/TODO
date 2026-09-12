package req

import (
	"encoding/json"
	"io"
)

func Decode[T any](data io.ReadCloser) (*T, error) {
	var payload T

	if err := json.NewDecoder(data).Decode(&payload); err != nil {
		return nil, err
	}

	return &payload, nil
}
