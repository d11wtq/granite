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

	var (
		elems = make([]string, a.cast().Count())
		v     Value
	)

	for i := uint32(0); i < a.cast().Count(); i++ {
		v, _ = a.Get(Integer(i))
		elems[i] = v.String()
	}

	return fmt.Sprintf("[ %s ]", strings.Join(elems, ", "))
}

func (a *Vector) Eq(b Value) bool {
	var (
		av, bv Value
		idx    Integer
	)

	if b.Type() != V_VEC {
		return false
	}

	av, _ = a.Len()
	bv, _ = b.Len()

	if av != bv {
		return false
	}

	idx = av.(Integer) - 1

	for {
		av, _ = a.Get(idx)
		bv, _ = b.Get(idx)

		if !av.Eq(bv) {
			return false
		}

		if idx == 0 {
			break
		}

		idx--
	}

	return true
}

func (*Vector) Lt(b Value) bool {
	return b != Nil
}

func (*Vector) Lte(Value) bool {
	return true
}

func (a *Vector) Add(b Value) (Value, error) {
	return nil, NewOpError("+", a, b)
}

func (a *Vector) Sub(b Value) (Value, error) {
	return nil, NewOpError("-", a, b)
}

func (a *Vector) Mul(b Value) (Value, error) {
	return nil, NewOpError("*", a, b)
}

func (a *Vector) Div(b Value) (Value, error) {
	return nil, NewOpError("/", a, b)
}

func (a *Vector) Len() (Value, error) {
	return Integer(a.cast().Count()), nil
}

func (a *Vector) Append(b Value) (Value, error) {
	return (*Vector)(a.cast().Append(b)), nil
}

func (a *Vector) Get(b Value) (Value, error) {
	switch t := b.(type) {
	case Integer:
		v, err := a.cast().Get(uint32(t))
		if err != nil {
			return Nil, nil // FIXME: What type of error to return
		}
		return v.(Value), nil
	default:
		return Nil, nil // FIXME: What type of error to return
	}
}

func (a *Vector) Call(b Value) (Value, error) {
	return nil, NewOpError("call", a, b)
}
