package vm

// The singular nil value
type NilType struct{}

var Nil = (*NilType)(nil)

func (*NilType) Type() uint8 {
	return V_NIL
}

func (*NilType) String() string {
	return "nil"
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

func (*NilType) Add(Value) (Value, error) {
	return Nil, nil
}

func (*NilType) Sub(Value) (Value, error) {
	return Nil, nil
}

func (*NilType) Mul(Value) (Value, error) {
	return Nil, nil
}

func (*NilType) Div(Value) (Value, error) {
	return Nil, nil
}

func (*NilType) Len() (Value, error) {
	return Nil, nil
}

func (*NilType) Append(Value) (Value, error) {
	return Nil, nil
}

func (*NilType) Get(Value) (Value, error) {
	return Nil, nil
}

func (*NilType) Call(Value) (Value, error) {
	return Nil, nil
}
