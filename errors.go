package gotokens

// Error code partially adapted from article at https://middlemost.com/failure-is-your-domain/

import (
	"bytes"
	"fmt"
)

// application error codes
const (
	ENOERROR       string = ""
	EINTERNAL      string = "E_INTERNAL"
	EFILENOTFOUND  string = "E_FILE_NOT_FOUND"
	ENOSEARCHPATH  string = "E_NO_SEARCH_PATH"
	EMALFORMEDJSON string = "E_MALFORMED_JSON"
)

// Error is our error type
type Error struct {
	Code    string
	Message string
	Context string
	Op      string
	Err     error
}

// Error returns the string representation of the error message.
func (e *Error) Error() string {
	var buf bytes.Buffer

	if e.Op != "" {
		fmt.Fprintf(&buf, "%s: ", e.Op)
	}

	if e.Err != nil {
		buf.WriteString(e.Err.Error())
	} else {
		if e.Code != ENOERROR {
			fmt.Fprintf(&buf, "<%s> ", e.Code)
		}
		buf.WriteString(e.Message)
	}
	return buf.String()
}

// ErrorCode returns relevant error code
func ErrorCode(err error) string {
	if err == nil {
		return ENOERROR
	} else if e, ok := err.(*Error); ok && e.Code != ENOERROR {
		return e.Code
	} else if ok && e.Err != nil {
		return ErrorCode(e.Err)
	}
	return EINTERNAL
}

// ErrorMessage returns relevant error message
func ErrorMessage(err error) string {
	if err == nil {
		return ""
	} else if e, ok := err.(*Error); ok && e.Message != "" {
		return e.Message
	} else if ok && e.Err != nil {
		return ErrorMessage(e.Err)
	}
	return "An internal error has occurred"
}