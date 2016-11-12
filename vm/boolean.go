package vm

type Boolean bool

func (Boolean) Type() uint8 {
	return V_BLN
}

func (a Boolean) String() string {
	if bool(a) {
		return "true"
	}

	return "false"
}

func (a Boolean) Eq(b Value) bool {
	return (a == b)
}

func (a Boolean) Lt(b Value) bool {
	if b.Type() == V_BLN {
		return a != b && bool(b.(Boolean)) != false
	}

	return a.Type() < b.Type()
}

func (a Boolean) Lte(b Value) bool {
	if b.Type() == V_BLN {
		return bool(b.(Boolean)) != false
	}

	return a.Type() < b.Type()
}

func (a Boolean) Add(b Value) (Value, error) {
	return nil, NewOpError("+", a, b)
}

func (a Boolean) Sub(b Value) (Value, error) {
	return nil, NewOpError("-", a, b)
}

func (a Boolean) Mul(b Value) (Value, error) {
	return nil, NewOpError("*", a, b)
}

func (a Boolean) Div(b Value) (Value, error) {
	return nil, NewOpError("/", a, b)
}

func (a Boolean) Len() (Value, error) {
	return nil, NewOpError("len", a)
}

func (a Boolean) Append(b Value) (Value, error) {
	return nil, NewOpError("append", a, b)
}

func (a Boolean) Get(b Value) (Value, error) {
	return nil, NewOpError("get", a, b)
}
