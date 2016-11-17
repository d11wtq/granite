package vm

// Convert Ax to the signed variant of it.
// Signed integers use 2's complement ^(n - 1).
// -1 in 32 bits is 0xFFFFFFFF.
// Only the mask has been applied to this, which strips off the topmost 6 bits.
// We check for the sign, then OR the bits back into the 32 bit number.
func sAx(ax uint32) int32 {
	if ax < uint32((((^0xFFFFF) >> 1) & 0xFFFFF)) {
		return int32(ax)
	}

	return int32(ax) | (^0xFFFFF)
}

// Convert Bx to the signed variant of it.
// Signed integers use 2's complement ^(n - 1).
// -1 in 32 bits is 0xFFFFFFFF.
// Only the mask has been applied to this, which strips off the topmost 14 bits.
// We check for the sign, then OR the bits back into the 32 bit number.
func sBx(bx uint32) int32 {
	if bx < uint32((((^0x3FFFF) >> 1) & 0x3FFFF)) {
		return int32(bx)
	}

	return int32(bx) | (^0x3FFFF)
}

// Encode a three operand instruction.
// All arguments are unsigned integers.
// Values too large will be bitwise truncated without warning.
//
//   11111111 - 10000000 - 00111111 - 11000000
//   [    Ax    ][    Bx    ][   Cx    ][ OP ]
//
//   ax = 9 bits
//   bx = 9 bits
//   cx = 8 bits
//   op = 6 bits
func encodeAxBxCx(op, ax, bx, cx uint32) uint32 {
	// 00000000 00000000 00000001 11111111 // ax (0x1FF) = 9 bits
	// 11111111 10000000 00000000 00000000 // ax << 23

	// 00000000 00000000 00000001 11111111 // bx (0x1FF) = 9 bits
	// 00000000 01111111 11000000 00000000 // bx << 14

	// 00000000 00000000 00000000 11111111 // cx (0xFF) = 8 bits
	// 00000000 00000000 00111111 11000000 // cx << 6

	// 00000000 00000000 00000000 00111111 // op (0x3F) = 6 bits
	return ((ax & 0x1FF) << 23) | ((bx & 0x1FF) << 14) | ((cx & 0xFF) << 6) | (op & 0x3F)
}

// Encode a two operand instruction.
// All arguments are unsigned integers, except for Bx, which may be signed.
// Signedness is dependent on the context in which the instruction is used.
// Signed Bx values must first be cast as their unsigned counterparts.
// Values too large will be bitwise truncated without warning.
//
//   11111111 - 00000000 - 00000000 - 00111111
//   [  Ax  ]   [         Bx           ][ OP ]
//
//   ax = 8 bits
//   bx = 18 bits
//   op = 6 bits
func encodeAxBx(op, ax, bx uint32) uint32 {
	// 00000000 00000000 00000000 11111111 // ax (0xFF) = 8 bits
	// 11111111 00000000 00000000 00000000 // ax << 24

	// 00000000 00000011 11111111 11111111 // bx (0x3FFFF) = 18 bits
	// 00000000 11111111 11111111 11000000 // bx << 6

	// 00000000 00000000 00000000 00111111 // op (0x3F) = 6 bits
	return ((ax & 0xFF) << 24) | ((bx & 0x3FFFF) << 6) | (op & 0x3F)
}

// Encode a one operand instruction.
// Ax may be signed or unsigned.
// Signedness is dependent on the context in which the instruction is used.
// Signed Ax values must first be cast as their unsigned counterparts.
// Values too large will be bitwise truncated without warning.
//
//   11111111 - 11111111 - 11111111 - 11000000
//   [               Ax                ][ OP ]
//
//   ax = 26 bits
//   op = 6 bits
func encodeAx(op, ax uint32) uint32 {
	// 00000011 11111111 11111111 11111111 // ax (0xFFFFF) = 26 bits
	// 11111111 11111111 11111111 11000000 // ax << 6

	// 00000000 00000000 00000000 00111111 // op (0x3F) = 6 bits
	return ((ax & 0xFFFFF) << 6) | (op & 0x3F)
}

// Decode a three operand instruction.
// See encodeAxBxCx() for the layout.
// Values are returned by reference.
func decodeAxBxCx(inst uint32, ax *uint32, bx *uint32, cx *uint8) {
	*ax = uint32((inst >> 23) & 0x1FF) // 9 bits
	*bx = uint32((inst >> 14) & 0x1FF) // 9 bits
	*cx = uint8((inst >> 6) & 0xFF)    // 8 bits
}

// Decode a two operand instruction.
// See encodeAxBx() for the layout.
// If Bx is signed, it can be converted to int32 with sBx(bx).
// Values are returned by reference.
func decodeAxBx(inst uint32, ax *uint32, bx *uint32) {
	*ax = uint32((inst >> 24) & 0xFF)   // 8 bits
	*bx = uint32((inst >> 6) & 0x3FFFF) // 18 bits
}

// Decode a one operand instruction.
// See encodeAx() for the layout.
// If Ax is signed, it can be converted to int32 with sAx(ax).
// Values are returned by reference.
func decodeAx(inst uint32, ax *uint32) {
	*ax = uint32((inst >> 6) & 0xFFFFF) // 20 bits
}
