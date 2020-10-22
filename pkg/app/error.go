package app

// Error is a generic Application error struct
type Error struct {
	Message     string      `json:"message"`
	Details     interface{} `json:"details"`
	Description string      `json:"description"`
	Code        string      `json:"code"`
}

func (e Error) Error() string {
	return e.Message
}

// FormatAppError generates a new Error struct from params
func FormatAppError(msg, code, desc string, details interface{}) *Error {
	return &Error{
		Message:     msg,
		Details:     details,
		Description: desc,
		Code:        code,
	}
}
