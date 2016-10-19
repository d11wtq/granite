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
