package fibers

type Violation struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

type Error struct {
	Code       int         `json:"code,omitempty"`
	Message    string      `json:"message,omitempty"`
	Type       string      `json:"type,omitempty"`
	Violations []Violation `json:"violations,omitempty"`
}

func NewError(code int, message string, errType string) *Error {
	return &Error{
		Code:    code,
		Message: message,
		Type:    errType,
	}
}
func NewViolationErrors(violations []Violation) *Error {
	return &Error{
		Code:       442,
		Message:    "unprocessable entity",
		Type:       "unprocessable_entity",
		Violations: violations,
	}
}

func (e *Error) Error() string {
	return e.Message
}
