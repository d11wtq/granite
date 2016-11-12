package vm

import (
	"fmt"
	"strings"
)

const (
	E_BADMATCH = iota
	E_NAME
)

type NameError string

func (e NameError) Error() string {
	return fmt.Sprintf(
		"*** NameError: unbound variable `%s'",
		string(e),
	)
}

type BadMatch struct {
	Right Value
}

func (e *BadMatch) Error() string {
	return fmt.Sprintf(
		"*** BadMatch: no match for right-hand-side of expression `%s'",
		e.Right.String(),
	)
}

type OpError struct {
	Op       string
	Operands []Value
}

func NewOpError(op string, operands ...Value) *OpError {
	return &OpError{op, operands}
}

func (e *OpError) Error() string {
	s := make([]string, 0, len(e.Operands))
	for _, v := range e.Operands {
		s = append(s, v.String())
	}
	return fmt.Sprintf(
		"*** OpError: unsupported operation `%s': %s",
		e.Op,
		strings.Join(s, ", "),
	)
}
