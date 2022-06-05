package utility

type errorResponseSchema struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
}

type successResponseSchema struct {
	StatusCode int         `json:"status_code"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data"`
}

func ErrorResponse(statusCode int, message string) errorResponseSchema {
	return errorResponseSchema{
		StatusCode: statusCode,
		Message:    message,
	}
}

func SuccessResponse(message string, data interface{}) successResponseSchema {
	return successResponseSchema{
		StatusCode: 200,
		Message:    message,
		Data:       data,
	}
}
