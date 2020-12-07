package error_custom

type errorContext struct {
	Field   string
	Message string
}

// AddErrorContext adds a context to an error
func AddErrorContext(err error, field, message string) error {
	context := errorContext{Field: field, Message: message}
	if customErr, ok := err.(CustomError); ok {
		return CustomError{
			errorType:    customErr.errorType,
			wrappedError: customErr.wrappedError,
			errorContext: context,
		}
	}

	return CustomError{
		errorType:    Error,
		wrappedError: err,
		errorContext: context,
	}
}

// GetErrorContext returns the error context
func GetErrorContext(err error) map[string]string {
	emptyContext := errorContext{}
	if customErr, ok := err.(CustomError); ok || customErr.errorContext != emptyContext {

		return map[string]string{
			"field":   customErr.errorContext.Field,
			"message": customErr.errorContext.Message,
		}
	}

	return nil
}

// GetType returns the error type
func GetType(err error) ErrorType {
	if customErr, ok := err.(CustomError); ok {
		return customErr.errorType
	}

	return Error
}
