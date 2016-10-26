package vm

// Encode a three operand instruction.
func encodeAxBxCx(op, ax, bx, cx uint32) uint32 {
	// 00000000 00000000 00000000 00111111 // op (0x3F) = 6 bits
	// 11111100 00000000 00000000 00000000 // op << 26

	// 00000000 00000000 00000001 11111111 // ax (0x1FF) = 9 bits
	// 00000011 11111110 00000000 00000000 // ax << 17

	// 00000000 00000000 00000001 11111111 // bx (0x1FF) = 9 bits
	// 00000000 00000001 11111111 00000000 // bx << 8

	// 00000000 00000000 00000000 11111111 // cx (0xFF) = 8 bits
	return (op << 26) | (ax << 17) | (bx << 8) | cx
}

// Encode a two operand instruction.
func encodeAxBx(op, ax, bx uint32) uint32 {
	// 00000000 00000000 00000000 00111111 // op (0x3F) = 6 bits
	// 11111100 00000000 00000000 00000000 // op << 26

	// 00000000 00000000 00000000 11111111 // ax (0xFF) = 8 bits
	// 00000011 11111100 00000000 00000000 // ax << 18

	// 00000000 00000011 11111111 11111111 // bx (0x3FFFF) = 18 bits
	return (op << 26) | (ax << 18) | bx
}

// Encode a one operand instruction.
func encodeAx(op, ax uint32) uint32 {
	// 00000000 00000000 00000000 00111111 // op (0x3F) = 6 bits
	// 11111100 00000000 00000000 00000000 // op << 26

	// 00000011 11111111 11111111 11111111 // ax (0xFFFFF) = 26 bits
	return (op << 26) | ax
}
