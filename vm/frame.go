package vm

// Frame of execution in the VM.
type Frame struct {
	// Registers of the currently executing frame.
	Registers [256]Value
	// Stack of POP addresses to return to
	LinkRegister []int64
	// The current instruction pointer.
	IP int64
}

// Create a new stack frame starting at instruction IP.
func NewFrame(ip int64) *Frame {
	return &Frame{IP: ip, LinkRegister: make([]int64, 0, 16)}
}

// Load the set of registers from the given env.
func (f *Frame) Restore(env [256]Value) {
	copy(f.Registers[:], env[:])
}

// Set the value in a given register.
func (f *Frame) Set(regNum uint8, v Value) {
	f.Registers[regNum] = v
}

// Push the IP at a given offset onto the link register for later return.
func (f *Frame) PushLink(offset int64) {
	f.LinkRegister = append(f.LinkRegister, f.IP+offset)
}

// Pop the top of stack IP from the link register and jmp to it
func (f *Frame) PopLink() {
	f.IP = f.LinkRegister[len(f.LinkRegister)-1]
	f.LinkRegister = f.LinkRegister[:len(f.LinkRegister)-1]
}
