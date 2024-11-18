package http_helpers

import (
	"encoding/json"
	"net/http"
	"url_shortener/pkg/constants/headers"
	mediatypes "url_shortener/pkg/constants/media_types"
)

func WriteJsonRes(resWriter http.ResponseWriter, res interface{}, statusCode int) error {
	resWriter.WriteHeader(statusCode)
	resWriter.Header().Set(headers.ContentType, mediatypes.ApplicationJson)
	err := json.NewEncoder(resWriter).Encode(res)

	return err
}
