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
	return nil
}

func (a String) Sub(b Value) Value {
	return nil
}

func (a String) Mul(b Value) Value {
	return nil
}

func (a String) Div(b Value) Value {
	return nil
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
