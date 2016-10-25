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

func (a Boolean) Add(b Value) Value {
	return nil
}

func (a Boolean) Sub(b Value) Value {
	return nil
}

func (a Boolean) Mul(b Value) Value {
	return nil
}

func (a Boolean) Div(b Value) Value {
	return nil
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

func (Boolean) Len() uint64 {
	return 0
}

func (Boolean) Append(Value) Value {
	return Nil
}

func (Boolean) Get(Value) Value {
	return Nil
}
