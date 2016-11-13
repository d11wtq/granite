package vm

import (
	"fmt"
)

type String string

func (String) Type() uint8 {
	return V_STR
}

func (a String) String() string {
	return fmt.Sprintf("%#v", a)
}

func (a String) Eq(b Value) bool {
	return (a == b)
}

func (a String) Lt(b Value) bool {
	if b.Type() == V_STR {
		return a < b.(String)
	}

	return a.Type() < b.Type()
}

func (a String) Lte(b Value) bool {
	if b.Type() == V_STR {
		return a <= b.(String)
	}

	return a.Type() < b.Type()
}

func (a String) Add(b Value) (Value, error) {
	switch t := b.(type) {
	case String:
		return a + t, nil
	default:
		return nil, NewOpError("+", a, b)
	}
}

func (a String) Sub(b Value) (Value, error) {
	return nil, NewOpError("-", a, b)
}

func (a String) Mul(b Value) (Value, error) {
	switch t := b.(type) {
	case Integer:
		var r String
		for t > 0 {
			r += a
			t--
		}
		return r, nil
	default:
		return nil, NewOpError("*", a, b)
	}
}

func (a String) Div(b Value) (Value, error) {
	return nil, NewOpError("/", a, b)
}

func (a String) Len() (Value, error) {
	return Integer(len(a)), nil
}

func (a String) Append(b Value) (Value, error) {
	return nil, NewOpError("append", a, b)
}

func (a String) Get(b Value) (Value, error) {
	return nil, NewOpError("get", a, b)
}

func (a String) Call(b Value) (Value, error) {
	return nil, NewOpError("call", a, b)
}
