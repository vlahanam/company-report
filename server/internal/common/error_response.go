package common

import "fmt"

type ErrorResponse struct {
	Status  string                 `json:"status"`
	Code    string                 `json:"code"`
	Message string                 `json:"message"`
	Debug   string                 `json:"debug,omitempty"`
	Details map[string]interface{} `json:"details,omitempty"`
}

func (r ErrorResponse) Error() string {
	return fmt.Sprintf("[%s] %s", r.Code, r.Message)
}

func NewErrorResponse(code, message string) ErrorResponse {
	return ErrorResponse{
		Status:  "ERROR",
		Code:    code,
		Message: message,
	}
}

func (r ErrorResponse) WithDebug(debug string) ErrorResponse {
	r.Debug = debug
	return r
}

func (r ErrorResponse) WithDetail(details map[string]interface{}) ErrorResponse {
	r.Details = details
	return r
}

func (r ErrorResponse) WithMessage(message string) ErrorResponse {
	r.Message = message
	return r
}

var (
	ErrCodeInternal     = NewErrorResponse("INTERNAL_ERROR", "Internal system error")
	ErrCodeBadRequest   = NewErrorResponse("BAD_REQUEST", "Invalid request")
	ErrCodeValidation   = NewErrorResponse("VALIDATION_ERROR", "Validation failed")
	ErrCodeUnauthorized = NewErrorResponse("UNAUTHORIZED", "Unauthorized access")
	ErrCodeForbidden    = NewErrorResponse("FORBIDDEN", "Access forbidden")
	ErrCodeNotFound     = NewErrorResponse("NOT_FOUND", "Resource not found")
)
