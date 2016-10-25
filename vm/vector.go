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

	elems := make([]string, a.Len())
	for i := uint64(0); i < a.Len(); i++ {
		elems[i] = a.Get(Integer(i)).String()
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

func (a *Vector) Len() uint64 {
	return uint64(a.cast().Count())
}

func (a *Vector) Append(b Value) Value {
	return (*Vector)(a.cast().Append(b))
}

func (a *Vector) Get(b Value) Value {
	switch t := b.(type) {
	case Integer:
		v, err := a.cast().Get(uint32(t))
		if err != nil {
			return Nil
		}
		return v.(Value)
	default:
		return Nil
	}
}
