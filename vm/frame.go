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

// Push the IP at a given offset onto the link register for later return.
func (f *Frame) pushLink(offset int64) {
	f.LinkRegister = append(f.LinkRegister, f.IP+offset)
}

// Pop the top of stack IP from the link register and jmp to it
func (f *Frame) popLink() {
	f.IP = f.LinkRegister[len(f.LinkRegister)-1]
	f.LinkRegister = f.LinkRegister[:len(f.LinkRegister)-1]
}
