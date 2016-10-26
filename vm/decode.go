package vm

// Decode opcode with three operands.
func decodeAxBxCx(inst uint32, ax *uint32, bx *uint32, cx *uint8) {
	// 00000000 00000000 00000000 00111111 // op (0x3F) = 6 bits
	// 11111100 00000000 00000000 00000000 // op << 26

	// 00000000 00000000 00000001 11111111 // ax (0x1FF) = 9 bits
	// 00000011 11111110 00000000 00000000 // ax << 17

	// 00000000 00000000 00000001 11111111 // bx (0x1FF) = 9 bits
	// 00000000 00000001 11111111 00000000 // bx << 8

	// 00000000 00000000 00000000 11111111 // cx (0xFF) = 8 bits

	*ax = uint32((inst >> 17) & 0x1FF) // 9 bits
	*bx = uint32((inst >> 8) & 0x1FF)  // 9 bits
	*cx = uint8(inst & 0xFF)           // 8 bits
}

// Decode opcode with two operands.
func decodeAxBx(inst uint32, ax *uint32, bx *uint32) {
	// 00000000 00000000 00000000 00111111 // op (0x3F) = 6 bits
	// 11111100 00000000 00000000 00000000 // op << 26

	// 00000000 00000000 00000000 11111111 // ax (0xFF) = 8 bits
	// 00000011 11111100 00000000 00000000 // ax << 18

	// 00000000 00000011 11111111 11111111 // bx (0x3FFFF) = 18 bits

	*ax = uint32((inst >> 18) & 0xFF) // 8 bits
	*bx = uint32(inst & 0x3FFFF)      // 18 bits
}

// Decode opcode with a single operand.
func decodeAx(inst uint32, ax *uint32) {
	// 00000000 00000000 00000000 00111111 // op (0x3F) = 6 bits
	// 11111100 00000000 00000000 00000000 // op << 26

	// 00000011 11111111 11111111 11111111 // ax (0xFFFFF) = 26 bits

	*ax = uint32(inst & 0xFFFFF) // 20 bits
}
