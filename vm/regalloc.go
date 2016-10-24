package vm

// Resource pool of all registers available for allocation.
type RegisterPool struct {
	// The set of currently available registers
	Available []uint32
	// The set of currently in-use registers
	Live map[uint32]bool
}

// Create a new register pool of size.
func NewRegisterPool(size uint32) *RegisterPool {
	pool := &RegisterPool{
		Available: make([]uint32, 0, size),
		Live:      make(map[uint32]bool),
	}
	for size > 0 {
		size--
		pool.Available = append(pool.Available, size)
		pool.Live[size] = false
	}
	return pool
}

// Reserve a register from the pool.
// The compiler has exclusive use of this register until it is released.
func (p *RegisterPool) Reserve() uint32 {
	if len(p.Available) == 0 {
		panic("out of registers")
	}

	regIdx := p.Available[len(p.Available)-1]
	p.Available = p.Available[:len(p.Available)-1]
	p.Live[regIdx] = true

	return regIdx
}

// Release a register back into the pool.
func (p *RegisterPool) Release(regIdx uint32) {
	if p.Live[regIdx] == false {
		panic("register not in use")
	}

	p.Live[regIdx] = false
	p.Available = append(p.Available, regIdx)
}
