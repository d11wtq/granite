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

func (a Integer) Add(b Value) (Value, error) {
	switch t := b.(type) {
	case Integer:
		return a + t, nil
	default:
		return nil, NewOpError("+", a, b)
	}
}

func (a Integer) Sub(b Value) (Value, error) {
	switch t := b.(type) {
	case Integer:
		return a - t, nil
	default:
		return nil, NewOpError("-", a, b)
	}
}

func (a Integer) Mul(b Value) (Value, error) {
	switch t := b.(type) {
	case Integer:
		return a * t, nil
	default:
		return nil, NewOpError("*", a, b)
	}
}

func (a Integer) Div(b Value) (Value, error) {
	switch t := b.(type) {
	case Integer:
		return a / t, nil
	default:
		return nil, NewOpError("/", a, b)
	}
}

func (a Integer) Len() (Value, error) {
	return nil, NewOpError("len", a)
}

func (a Integer) Append(b Value) (Value, error) {
	return nil, NewOpError("append", a, b)
}

func (a Integer) Get(b Value) (Value, error) {
	return nil, NewOpError("get", a, b)
}

func (a Integer) Call(b Value) (Value, error) {
	return nil, NewOpError("call", a, b)
}
