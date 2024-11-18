package http_helpers

import (
	"net/http"

	"github.com/go-playground/validator/v10"
)

func ValidateReq[T any](resWriter *http.ResponseWriter, reqPayload T) error {
	validate := validator.New(validator.WithRequiredStructEnabled())
	validationErr := validate.Struct(reqPayload)

	if validationErr != nil {
		WriteJsonRes(*resWriter, validationErr.Error(), 400)
		return validationErr
	}

	return nil
}
