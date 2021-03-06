package vm

const (
	// The nil value
	V_NIL uint8 = iota
	// Boolean values
	V_BLN
	// Integer values
	V_INT
	// String values
	V_STR
	// Vectors
	V_VEC
	// Closures
	V_FUN
)

// Runtime values, stored in VM registers.
type Value interface {
	// The type tag of the value.
	// This is used to identify constants in the byte-code.
	Type() uint8
	// The string representation of the value.
	String() string
	// Implementation for the EQ opcode.
	Eq(Value) bool
	// Implementation for the LT opcode.
	Lt(Value) bool
	// Implementation for the LTE opcode.
	Lte(Value) bool
	// Implementation for the ADD opcode.
	Add(Value) (Value, error)
	// Implementation for the SUB opcode.
	Sub(Value) (Value, error)
	// Implementation for the MUL opcode.
	Mul(Value) (Value, error)
	// Implementation for the DIV opcode.
	Div(Value) (Value, error)
	// Implementation for the LEN opcode.
	Len() (Value, error)
	// Implementation for the APPEND opcode.
	Append(Value) (Value, error)
	// Implementation for the GET opcode.
	Get(Value) (Value, error)
	// Implementation for the CALL opcode.
	Call(Value) (Value, error)
}
