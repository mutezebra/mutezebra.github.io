package merror

import (
	"fmt"
	"io"
)

type merror struct {
	msg   string
	extra int64
	*stack
}

type Merror interface {
	Error() string
	Extra() any
}

func New(msg string, extra ...int64) error {
	e := &merror{
		msg:   msg,
		stack: callers(),
	}
	if extra != nil {
		e.extra = extra[0]
	}
	return e
}

func (e *merror) Error() string {
	return e.msg
}

func (e *merror) Extra() any {
	return e.extra
}

func (e *merror) Format(state fmt.State, verb rune) {
	switch verb {
	case 's', 'v':
		io.WriteString(state, fmt.Sprintf("msg: %s,extra: %d", e.msg, e.extra))
		switch {
		case state.Flag('+'):
			io.WriteString(state, fmt.Sprintf("%+v", e.stack)) // here！
		}
	}
}

func (e *merror) StackTrace() any {
	return e.stack
}
