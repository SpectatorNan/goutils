package errorx

// ForbiddenError represents a 403 Forbidden error
type ForbiddenError struct {
	Message string
	Reason  string
}

func (e *ForbiddenError) Error() string {
	return e.Message
}

// NewForbiddenError creates a new forbidden error
func NewForbiddenError(message string, reason string) error {
	return &ForbiddenError{
		Message: message,
		Reason:  reason,
	}
}
