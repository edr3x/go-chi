package utils

type ErrorTypeStruct struct {
	Message    string
	StatusCode int
}

func (e *ErrorTypeStruct) Error() string {
	return e.Message
}

func CustomError(statusCode int, message string) *ErrorTypeStruct {
	return &ErrorTypeStruct{
		StatusCode: statusCode,
		Message:    message,
	}
}
