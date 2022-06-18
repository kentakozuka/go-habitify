package habitify

import (
	"fmt"
	"net/http"
)

type ErrPreconditionFailed struct{ msg string }

func (e *ErrPreconditionFailed) Error() string { return e.msg }

type ErrInternal struct{ msg string }

func (e *ErrInternal) Error() string { return e.msg }

func newError(statusCode int, msg string) error {
	switch statusCode {
	case http.StatusPreconditionFailed:
		return &ErrPreconditionFailed{msg: msg}
	}
	return fmt.Errorf("unknown error: %s", msg)
}
