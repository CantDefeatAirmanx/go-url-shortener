package http_helpers

import (
	"encoding/json"
	"io"
)

func DecodeJson[T any](body io.ReadCloser) (*T, error) {
	var payload T
	jsonParseErr := json.NewDecoder(body).Decode(&payload)

	if jsonParseErr != nil {
		return nil, jsonParseErr
	}

	return &payload, nil
}
