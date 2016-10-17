package vm

const (
	V_NIL uint8 = iota
	V_INT
	V_STR
)

type Value interface {
	Type() uint8
	Add(Value) Value
}
