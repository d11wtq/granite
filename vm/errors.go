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
	Left  Value
	Right Value
}

func (e *BadMatch) Error() string {
	return fmt.Sprintf(
		"*** BadMatch: no match for right hand side: %s = %s",
		e.Left.String(),
		e.Right.String(),
	)
}
