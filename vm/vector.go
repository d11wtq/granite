package vm

import (
	"fmt"
	"github.com/d11wtq/persistent/vector"
	"strings"
)

type Vector vector.Vector

var EmptyVector *Vector

func init() {
	EmptyVector = (*Vector)(vector.New())
}

func (*Vector) Type() uint8 {
	return V_VEC
}

func (a *Vector) cast() *vector.Vector {
	return (*vector.Vector)(a)
}

func (a *Vector) String() string {
	if a.cast().Count() == 0 {
		return "[]"
	}

	var v vector.Value

	elems := make([]string, a.cast().Count())
	for i := uint32(0); i < a.cast().Count(); i++ {
		v, _ = a.cast().Get(i)
		elems[i] = v.(Value).String()
	}

	return fmt.Sprintf("[ %s ]", strings.Join(elems, ", "))
}

func (*Vector) Add(Value) Value {
	return Nil
}

func (*Vector) Sub(Value) Value {
	return Nil
}

func (*Vector) Mul(Value) Value {
	return Nil
}

func (*Vector) Div(Value) Value {
	return Nil
}

func (*Vector) Eq(b Value) bool {
	return (b == Nil)
}

func (*Vector) Lt(b Value) bool {
	return b != Nil
}

func (*Vector) Lte(Value) bool {
	return true
}

func (a *Vector) Append(b Value) Value {
	return (*Vector)((*vector.Vector)(a).Append(b))
}
