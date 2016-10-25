package vm

const (
	V_NIL uint8 = iota
	V_BLN
	V_INT
	V_STR
	V_VEC
)

type Value interface {
	Type() uint8
	String() string
	Add(Value) Value
	Sub(Value) Value
	Mul(Value) Value
	Div(Value) Value
	Eq(Value) bool
	Lt(Value) bool
	Lte(Value) bool
	Len() uint64
	Append(Value) Value
	Get(Value) Value
}
