package data

type Error struct {
	Cause    error
	Metadata map[string]string
}

func NewError(cause error, metadata map[string]string) *Error {
	return &Error{
		Cause:    cause,
		Metadata: metadata,
	}
}

func (e *Error) Error() string {
	return e.Cause.Error()
}
