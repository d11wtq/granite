package vm

// Decode opcode with a single operand.
func decodeAx(inst uint32, ax *uint32) {
	*ax = uint32(inst & 0xFFFFF) // 20 bits
}

// Decode opcode with two operands.
func decodeAxBx(inst uint32, ax *uint32, bx *uint16) {
	*ax = uint32((inst >> 13) & 0x1FFF) // 13 bits
	*bx = uint16(inst & 0x1FFF)         // 13 bits
}

// Decode opcode with three operands.
func decodeAxBxCx(inst uint32, ax *uint32, bx *uint16, cx *uint8) {
	*ax = uint32((inst >> 17) & 0x1FF) // 9 bits
	*bx = uint16((inst >> 8) & 0x1FF)  // 9 bits
	*cx = uint8(inst & 0xFF)           // 8 bits
}
