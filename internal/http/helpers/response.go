package helpers

type standardResponse struct {
	Success      bool   `json:"success"` // can be ... or ...
	Message      string `json:"message"`
	Data         any    `json:"data,omitempty"`    // presents only if `Success = true`
	ErrorDetails any    `json:"details,omitempty"` // presents only if `Success = false`
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
