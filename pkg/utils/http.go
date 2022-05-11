package utils

import (
	"fmt"
	"net/http"

	"github.com/richardsahvic/jamtangan/domain/model"
)

// RequestRequired returns a response with a request required message.
func RequestRequired(fieldName string) (int, *model.BaseResponse) {
	return http.StatusBadRequest, &model.BaseResponse{RawMessage: fmt.Sprintf("%s is required", fieldName)}
}

// RequestInvalid returns a response with a request invalid message.
func RequestInvalid(fieldName string) (int, *model.BaseResponse) {
	return http.StatusBadRequest, &model.BaseResponse{RawMessage: fmt.Sprintf("%s is invalid", fieldName)}
}
