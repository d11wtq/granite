package vm

type UndefinedType struct {
}

var Undefined = (*UndefinedType)(nil)

func (*UndefinedType) Type() uint8 {
	return uint8(1 << 7)
}

func (*UndefinedType) String() string {
	return "undefined"
}

func (*UndefinedType) Add(Value) Value {
	return Undefined
}

func (*UndefinedType) Sub(Value) Value {
	return Undefined
}

func (*UndefinedType) Mul(Value) Value {
	return Undefined
}

func (*UndefinedType) Div(Value) Value {
	return Undefined
}

func (*UndefinedType) Eq(b Value) bool {
	return false
}

func (*UndefinedType) Lt(b Value) bool {
	return false
}

func (*UndefinedType) Lte(Value) bool {
	return true
}

func (*UndefinedType) Append(Value) Value {
	return Undefined
}
