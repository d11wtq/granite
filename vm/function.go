package vm

import (
	"fmt"
)

// Function closures.
type Function struct {
	// The VM that created this function.
	VM *VM
	// The environment in which the function was created (registers).
	Env [256]Value
	// The location of the first instruction in the function.
	IP int64
}

// Create a new function, which starts at ip in vm.
func NewFunction(vm *VM, ip int64) *Function {
	return &Function{
		VM: vm,
		IP: ip,
	}
}

// Create a closure over the given env.
// The environment is the set of registers in the current frame.
func (a *Function) Close(env [256]Value) {
	copy(a.Env[:], env[:])
}

func (*Function) Type() uint8 {
	return V_FUN
}

func (a *Function) String() string {
	return fmt.Sprintf("<Function %p:%08x>", a.VM, a.IP)
}

func (a *Function) Eq(b Value) bool {
	return (a == b)
}

func (a *Function) Lt(b Value) bool {
	return false
}

func (*Function) Lte(Value) bool {
	return false
}

func (*Function) Add(Value) (Value, error) {
	return Nil, nil
}

func (*Function) Sub(Value) (Value, error) {
	return Nil, nil
}

func (*Function) Mul(Value) (Value, error) {
	return Nil, nil
}

func (*Function) Div(Value) (Value, error) {
	return Nil, nil
}

func (*Function) Len() (Value, error) {
	return Nil, nil
}

func (*Function) Append(Value) (Value, error) {
	return Nil, nil
}

func (*Function) Get(Value) (Value, error) {
	return Nil, nil
}

func (a *Function) Call(args Value) (Value, error) {
	f := NewFrame(a.IP)
	f.Restore(a.Env)
	f.Set(uint8(ArgumentRegister), args)
	return a.VM.Run(f)
}
