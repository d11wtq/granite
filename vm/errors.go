package vm

import (
	"fmt"
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
