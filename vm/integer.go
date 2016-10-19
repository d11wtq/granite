package vm

import (
	"fmt"
)

type Integer int64

func (Integer) Type() uint8 {
	return V_INT
}

func (a Integer) String() string {
	return fmt.Sprintf("%d", a)
}

func (a Integer) Add(b Value) Value {
	switch t := b.(type) {
	case Integer:
		return a + t
	default:
		return nil
	}
}

func (a Integer) Sub(b Value) Value {
	switch t := b.(type) {
	case Integer:
		return a - t
	default:
		return nil
	}
}

func (a Integer) Mul(b Value) Value {
	switch t := b.(type) {
	case Integer:
		return a * t
	default:
		return nil
	}
}

func (a Integer) Div(b Value) Value {
	switch t := b.(type) {
	case Integer:
		return a / t
	default:
		return nil
	}
}
