package protolexError

import (
	"github.com/fatih/color"
)

type Error struct {
	Code    int
	Message string
	Status  string
}

func (e *Error) Error() string {
	panic("unimplemented")
}

var ErrorCode = map[string]int{
	"NotFound":    1,
	"SyntaxError": 2,
	"TypeError":   3,
	"Info":        4,
}

func NewError(code string, message string) *Error {
	return &Error{
		Code:    ErrorCode[code],
		Message: message,
		Status:  code,
	}
}

func (e *Error) LogError() string {
	switch e.Code {
	case 1:
		color.Red("Error: %s", e.Message)
	case 2:
		color.Red("Syntax Error: %s", e.Message)
	case 3:
		color.Red("Type Error: %s", e.Message)
	case 4:
		color.Yellow("Info: %s", e.Message)
	}
	return e.Message
}
