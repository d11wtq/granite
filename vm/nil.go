package vm

type NilType struct {
}

var Nil = (*NilType)(nil)

func (*NilType) Type() uint8 {
	return V_NIL
}

func (*NilType) String() string {
	return "nil"
}

func (*NilType) Add(Value) Value {
	return Nil
}

func (*NilType) Sub(Value) Value {
	return Nil
}

func (*NilType) Mul(Value) Value {
	return Nil
}

func (*NilType) Div(Value) Value {
	return Nil
}

func (*NilType) Eq(b Value) bool {
	return (b == Nil)
}

func (*NilType) Lt(b Value) bool {
	return b != Nil
}

func (*NilType) Lte(Value) bool {
	return true
}

func (*NilType) Append(Value) Value {
	return Nil
}
