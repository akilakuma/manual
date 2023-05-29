package json

import (
	// "errors"

	"github.com/buger/jsonparser"
)

var (
	KeyPathNotFoundError       = jsonparser.KeyPathNotFoundError
	UnknownValueTypeError      = jsonparser.UnknownValueTypeError
	MalformedJsonError         = jsonparser.MalformedJsonError
	MalformedStringError       = jsonparser.MalformedStringError
	MalformedArrayError        = jsonparser.MalformedArrayError
	MalformedObjectError       = jsonparser.MalformedObjectError
	MalformedValueError        = jsonparser.MalformedValueError
	OverflowIntegerError       = jsonparser.OverflowIntegerError
	MalformedStringEscapeError = jsonparser.MalformedStringEscapeError
)
