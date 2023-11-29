package domain

type Response struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data,omitempty"`
	Error   *ErrorInfo  `json:"error,omitempty"`
}

// ErrorInfo holds detailed info about an error.
type ErrorInfo struct {
	Code    int    `json:"code,omitempty"`
	Type    string `json:"type,omitempty"`
	Message string `json:"message"`
	Details string `json:"details,omitempty"` // Optional for additional error details
}

// SuccessResponse creates a generic successful response with provided data.
func SuccessResponse(data interface{}) *Response {
	return &Response{
		Success: true,
		Data:    data,
	}
}

// ErrorResponse creates a generic error response with an error code and message.
func ErrorResponse(code int, errType, message string) *Response {
	return &Response{
		Success: false,
		Error: &ErrorInfo{
			Code:    code,
			Type:    errType,
			Message: message,
		},
	}
}

// DetailedErrorResponse allows adding more details to the error response.
func DetailedErrorResponse(code int, errType, message, details string) *Response {
	return &Response{
		Success: false,
		Error: &ErrorInfo{
			Code:    code,
			Type:    errType,
			Message: message,
			Details: details,
		},
	}
}
