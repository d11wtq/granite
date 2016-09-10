package parser

import (
	"fmt"
)

// Meaningful syntax error
type BijouParseError struct {
	// The lexer at the time the error occurred
	lexer *BijouLex
}

func (e *BijouParseError) Error() string {
	return fmt.Sprintf(
		"%s in %s line %d, column %d",
		e.errorMessage(),
		e.currentFile(),
		e.currentLine(),
		e.currentColumn(),
	)
}

func (e *BijouParseError) errorMessage() string {
	return e.lexer.error
}

func (e *BijouParseError) currentFile() string {
	return e.lexer.filename
}

func (e *BijouParseError) currentLine() int {
	return e.lexer.lineNo
}

func (e *BijouParseError) currentColumn() int {
	return e.lexer.columnNo
}
