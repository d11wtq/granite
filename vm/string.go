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

func (a String) Add(b Value) Value {
	switch t := b.(type) {
	case String:
		return a + t
	default:
		return Nil
	}
}

func (a String) Sub(b Value) Value {
	return Nil
}

func (a String) Mul(b Value) Value {
	switch t := b.(type) {
	case Integer:
		var r String
		for t > 0 {
			r += a
			t--
		}
		return r
	default:
		return Nil
	}
}

func (a String) Div(b Value) Value {
	return Nil
}

func (a String) Eq(b Value) bool {
	return (a == b)
}

func (a String) Lt(b Value) bool {
	if b.Type() == V_STR {
		return string(a) < string(b.(String))
	}

	return a.Type() < b.Type()
}

func (a String) Lte(b Value) bool {
	if b.Type() == V_STR {
		return string(a) <= string(b.(String))
	}

	return a.Type() < b.Type()
}

func (a String) Len() uint64 {
	return uint64(len(string(a)))
}

func (String) Append(Value) Value {
	return Nil
}

func (String) Get(Value) Value {
	return Nil
}
