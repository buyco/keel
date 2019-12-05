package app

// Generic Application error
type Error struct {
	Message     string      `json:"message"`
	Details     interface{} `json:"details"`
	Description string      `json:"description"`
	Code        string      `json:"code"`
}

func FormatAppError(msg, code, desc string, details interface{}) *Error {
	return &Error{
		Message:     msg,
		Details:     details,
		Description: desc,
		Code:        code,
	}
}
