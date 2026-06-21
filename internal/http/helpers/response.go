package helpers

type standardResponse struct {
	Success      bool   `json:"success"`           // `true` for successful response and `false` for error response
	Message      string `json:"message"`           // a message that gives more information about response
	Data         any    `json:"data,omitempty"`    // presents only for successful response (`success=true`)
	ErrorDetails any    `json:"details,omitempty"` // presents only for error response (`success=false`)
}

func GenerateSuccessfulResponse(message string, data any) *standardResponse {
	return &standardResponse{
		Success: true,
		Message: message,
		Data:    data,
	}
}

func GenerateErrorResponse(message string, details any) *standardResponse {
	return &standardResponse{
		Success:      false,
		Message:      message,
		ErrorDetails: details,
	}
}
