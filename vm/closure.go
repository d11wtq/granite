package vm

import (
	"fmt"
)

// Function closures.
type Closure struct {
	VM  *VM
	Env [256]Value
	IP  int64
}

func NewClosure(vm *VM, ip int64) *Closure {
	return &Closure{
		VM: vm,
		IP: ip,
	}
}

func (*Closure) Type() uint8 {
	return V_FUN
}

func (a *Closure) String() string {
	return fmt.Sprintf("<Closure %p:%08x>", a.VM, a.IP)
}

func (a *Closure) Eq(b Value) bool {
	return (a == b)
}

func (a *Closure) Lt(b Value) bool {
	return false
}

func (*Closure) Lte(Value) bool {
	return false
}

func (*Closure) Add(Value) (Value, error) {
	return Nil, nil
}

func (*Closure) Sub(Value) (Value, error) {
	return Nil, nil
}

func (*Closure) Mul(Value) (Value, error) {
	return Nil, nil
}

func (*Closure) Div(Value) (Value, error) {
	return Nil, nil
}

func (*Closure) Len() (Value, error) {
	return Nil, nil
}

func (*Closure) Append(Value) (Value, error) {
	return Nil, nil
}

func (*Closure) Get(Value) (Value, error) {
	return Nil, nil
}

func (a *Closure) Call(Value) (Value, error) {
	f := NewFrame(a.IP)
	copy(f.Registers[:], a.Env[:])
	return a.VM.Run(f)
}
