package vm

type String string

func (String) Type() uint8 {
	return V_STR
}

func (a String) Add(b Value) Value {
	return nil
}
