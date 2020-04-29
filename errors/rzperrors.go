package errors

import "net/http"

//RZPError ...
type RZPError struct {
	InternalErrorCode string `json:"internal_error_code"`
	Field             string `json:"field"`
	Description       string `json:"description"`
	Code              string `json:"code"`
}

//RZPErrorJSON ...
type RZPErrorJSON struct {
	ErrorData RZPError `json:"error"`
}

//BadRequestError ...
// type BadRequestError struct {
// 	Message string
// 	Err     error
// }

// //ServerError ...
// type ServerError struct {
// 	Message string
// 	Err     error
// }

// //GatewayError ...
// type GatewayError struct {
// 	Message string
// 	Err     error
// }

// //SignatureVerificationError ...
// type SignatureVerificationError struct {
// 	Message string
// 	Err     error
// }

// func handleError(msg string, desc string) string {
// 	errorMessage := ""
// 	if msg != "" {
// 		errorMessage += msg
// 	}
// 	errorMessage += desc
// 	return errorMessage
// }

//Error ...
// func (s *BadRequestError) Error() string {
// 	return handleError(s.Message, s.Err.Error())
// }

// func (s *GatewayError) Error() string {
// 	return handleError(s.Message, s.Err.Error())

// }

// func (s *ServerError) Error() string {
// 	return handleError(s.Message, s.Err.Error())
// }

// func (s *SignatureVerificationError) Error() string {
// 	return handleError(s.Message, s.Err.Error())
// }

// Error error struct
type Error struct {
	Code    int
	Message string
	Header  int
}

// Error implement error method
func (e *Error) Error() string {
	return e.Message
}

// New create a custom error
func New(code int, message string, header ...int) error {
	var hc = http.StatusBadRequest
	if len(header) > 0 && header[0] < 1000 {
		hc = header[0]
	}
	return &Error{Code: code, Message: message, Header: hc}
}
