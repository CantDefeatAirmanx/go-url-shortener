package http_helpers

import (
	"net/http"
)

func HandleJsonReqBody[T any](resWriter *http.ResponseWriter, req *http.Request) (*T, error) {
	payload, jsonParseErr := DecodeJson[T](req.Body)

	if jsonParseErr != nil {
		return nil, jsonParseErr
	}

	validatioErr := ValidateReq(resWriter, payload)
	if validatioErr != nil {
		return nil, validatioErr
	}

	return payload, nil
}
