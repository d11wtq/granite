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
		return Nil
	}
}

func (a Integer) Sub(b Value) Value {
	switch t := b.(type) {
	case Integer:
		return a - t
	default:
		return Nil
	}
}

func (a Integer) Mul(b Value) Value {
	switch t := b.(type) {
	case Integer:
		return a * t
	default:
		return Nil
	}
}

func (a Integer) Div(b Value) Value {
	switch t := b.(type) {
	case Integer:
		return a / t
	default:
		return Nil
	}
}

func (a Integer) Eq(b Value) bool {
	return (a == b)
}

func (a Integer) Lt(b Value) bool {
	if b.Type() == V_INT {
		return a < b.(Integer)
	}

	return a.Type() < b.Type()
}

func (a Integer) Lte(b Value) bool {
	if b.Type() == V_INT {
		return a <= b.(Integer)
	}

	return a.Type() < b.Type()
}

func (Integer) Len() uint64 {
	return 0
}

func (Integer) Append(Value) Value {
	return Nil
}

func (Integer) Get(Value) Value {
	return Nil
}
